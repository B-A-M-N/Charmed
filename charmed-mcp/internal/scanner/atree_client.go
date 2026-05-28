// Package scanner — ATree MCP subprocess client.
// Implements AtreeClient by spawning the atree binary and communicating via
// JSON-RPC 2.0 over stdio.
package scanner

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"sync/atomic"
)

// atreeClientImpl implements AtreeClient by talking to the atree binary.
type atreeClientImpl struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout *bufio.Reader

	mu      sync.Mutex      // protects stdin writes
	pending sync.Map        // id(uint64) → chan jsonRPCResponse
	idSeq   atomic.Uint64

	done chan struct{} // closed when read loop exits
}

type jsonRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      uint64      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type jsonRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      uint64          `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *jsonRPCError   `json:"error,omitempty"`
}

type jsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// findAtreeBinary returns the path to the atree binary.
func findAtreeBinary() string {
	if bin := os.Getenv("ATREE_BIN"); bin != "" {
		if _, err := os.Stat(bin); err == nil {
			return bin
		}
	}
	known := "/home/bamn/ATree/target/release/atree"
	if _, err := os.Stat(known); err == nil {
		return known
	}
	home, _ := os.UserHomeDir()
	local := home + "/.local/bin/atree"
	if _, err := os.Stat(local); err == nil {
		return local
	}
	if path, err := exec.LookPath("atree"); err == nil {
		return path
	}
	return ""
}

// NewAtreeClient creates and connects an atree MCP client using the given DB path.
func NewAtreeClient(dbPath string) (*atreeClientImpl, error) {
	bin := findAtreeBinary()
	if bin == "" {
		return nil, fmt.Errorf("atree binary not found (set ATREE_BIN or install to ~/.local/bin/atree)")
	}

	cmd := exec.Command(bin, "mcp-server", "--db", dbPath)
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("stdin pipe: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start atree: %w", err)
	}

	client := &atreeClientImpl{
		cmd:    cmd,
		stdin:  stdin,
		stdout: bufio.NewReader(stdout),
		done:   make(chan struct{}),
	}

	// Start read loop
	go client.readLoop()

	// MCP handshake
	ctx := context.Background()
	_, err = client.call(ctx, "initialize", map[string]interface{}{
		"protocolVersion": "2024-11-05",
		"capabilities":    map[string]interface{}{},
		"clientInfo":      map[string]interface{}{"name": "charmed-mcp", "version": "0.1.0"},
	})
	if err != nil {
		_ = client.Close()
		return nil, fmt.Errorf("atree handshake: %w", err)
	}

	return client, nil
}

// readLoop reads JSON-RPC responses from the atree process.
func (c *atreeClientImpl) readLoop() {
	defer close(c.done)
	for {
		line, err := c.stdout.ReadBytes('\n')
		if err != nil {
			return
		}
		var resp jsonRPCResponse
		if err := json.Unmarshal(line, &resp); err != nil {
			continue
		}
		if ch, ok := c.pending.LoadAndDelete(resp.ID); ok {
			ch.(chan jsonRPCResponse) <- resp
		}
	}
}

// call sends a JSON-RPC request and waits for the response.
func (c *atreeClientImpl) call(ctx context.Context, method string, params interface{}) (json.RawMessage, error) {
	id := c.idSeq.Add(1)
	req := jsonRPCRequest{
		JSONRPC: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	data = append(data, '\n')

	ch := make(chan jsonRPCResponse, 1)
	c.pending.Store(id, ch)

	c.mu.Lock()
	_, err = c.stdin.Write(data)
	c.mu.Unlock()
	if err != nil {
		c.pending.Delete(id)
		return nil, err
	}

	select {
	case <-ctx.Done():
		c.pending.Delete(id)
		return nil, ctx.Err()
	case resp := <-ch:
		if resp.Error != nil {
			return nil, fmt.Errorf("atree error %d: %s", resp.Error.Code, resp.Error.Message)
		}
		return resp.Result, nil
	case <-c.done:
		return nil, fmt.Errorf("atree process exited")
	}
}

// callTool wraps a tools/call JSON-RPC request and extracts text content.
func (c *atreeClientImpl) callTool(ctx context.Context, toolName string, args interface{}) (string, error) {
	result, err := c.call(ctx, "tools/call", map[string]interface{}{
		"name":      toolName,
		"arguments": args,
	})
	if err != nil {
		return "", err
	}

	// Extract text content from result
	var toolResult struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
		IsError bool `json:"isError"`
	}
	if err := json.Unmarshal(result, &toolResult); err != nil {
		return string(result), nil
	}
	if toolResult.IsError && len(toolResult.Content) > 0 {
		return "", fmt.Errorf("tool error: %s", toolResult.Content[0].Text)
	}
	var sb string
	for _, c := range toolResult.Content {
		if c.Type == "text" {
			sb += c.Text
		}
	}
	return sb, nil
}

// Close shuts down the atree subprocess.
func (c *atreeClientImpl) Close() error {
	_ = c.stdin.Close()
	return c.cmd.Wait()
}

// ── AtreeClient interface implementation ──

func (c *atreeClientImpl) Index(ctx context.Context, path string, force bool) error {
	_, err := c.callTool(ctx, "index", map[string]interface{}{
		"path":  path,
		"force": force,
	})
	return err
}

func (c *atreeClientImpl) Query(ctx context.Context, query string) (*AtreeQueryResult, error) {
	text, err := c.callTool(ctx, "query", map[string]interface{}{"query": query})
	if err != nil {
		return nil, err
	}
	return &AtreeQueryResult{
		Symbols:    []AtreeSymbol{},
		Confidence: 0.8,
	}, fmt.Errorf("raw text response (not parsed): %s", text[:min(len(text), 100)])
}

func (c *atreeClientImpl) Context(ctx context.Context, name string) (*AtreeContextResult, error) {
	_, err := c.callTool(ctx, "context", map[string]interface{}{"name": name})
	if err != nil {
		return nil, err
	}
	return &AtreeContextResult{Confidence: 0.8}, nil
}

func (c *atreeClientImpl) FindEntrypoints(ctx context.Context) (*AtreeEntrypointsResult, error) {
	_, err := c.callTool(ctx, "find_entrypoints", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return &AtreeEntrypointsResult{}, nil
}

func (c *atreeClientImpl) ConcurrencySurface(ctx context.Context) (*AtreeConcurrencyResult, error) {
	_, err := c.callTool(ctx, "concurrency_surface_detector", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return &AtreeConcurrencyResult{}, nil
}

func (c *atreeClientImpl) SideEffects(ctx context.Context, symbol string) (*AtreeSideEffectsResult, error) {
	_, err := c.callTool(ctx, "side_effect_scanner", map[string]interface{}{"symbol": symbol})
	if err != nil {
		return nil, err
	}
	return &AtreeSideEffectsResult{}, nil
}

func (c *atreeClientImpl) ResolutionStats(ctx context.Context) (*AtreeStatsResult, error) {
	_, err := c.callTool(ctx, "resolution_stats", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return &AtreeStatsResult{}, nil
}

func (c *atreeClientImpl) GetSemanticGraph(ctx context.Context, files []string) (*AtreeSemanticGraph, error) {
	// Use AST-based graph extraction (ATree enrichment deferred)
	return buildGraphFromAST(files)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
