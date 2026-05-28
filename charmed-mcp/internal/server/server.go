// Package server implements the MCP server for charmed.
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/B-A-M-N/charmed/charmed-mcp/internal/ir"
	"github.com/B-A-M-N/charmed/charmed-mcp/internal/scanner"
)

// CharmedServer is the MCP server for charmed.
type CharmedServer struct {
	scanner *scanner.Scanner
	mcp     *server.MCPServer
}

// New creates a new CharmedServer.
func New(s *scanner.Scanner) *CharmedServer {
	cs := &CharmedServer{
		scanner: s,
		mcp: server.NewMCPServer(
			"charmed",
			scanner.CharmedVersion,
		),
	}
	cs.registerTools()
	cs.registerResources()
	return cs
}

// registerTools registers all MCP tools.
func (cs *CharmedServer) registerTools() {
	// ===== EVIDENCE EXTRACTION =====

	// extract_evidence — deterministic structural extraction
	cs.mcp.AddTool(mcp.NewTool("extract_evidence",
		mcp.WithDescription("Extract architectural evidence from Go project — signals, topologies, primitive mappings. Deterministic extraction (no AI inference). Returns evidence graph conforming to evidence-schema.yaml."),
		mcp.WithString("project_path",
			mcp.Required(),
			mcp.Description("Absolute or relative path to the Go project root"),
		),
		mcp.WithString("extraction_method",
			mcp.Description("Extraction method: 'atree' (AST+ATree, best), 'ast' (Go parser only), 'grep' (pattern matching). Default: 'atree'"),
			mcp.Enum("atree", "ast", "grep"),
		),
		mcp.WithString("atree_db_path",
			mcp.Description("Override auto-detected ATree DB path"),
		),
	), cs.handleExtractEvidence)

	// scan_project — legacy compatibility (calls extract_evidence)
	cs.mcp.AddTool(mcp.NewTool("scan_project",
		mcp.WithDescription("Legacy scan tool. Use extract_evidence instead. Scans project and produces TUI Structural IR."),
		mcp.WithString("project_path",
			mcp.Required(),
			mcp.Description("Absolute or relative path to the Go project root"),
		),
		mcp.WithString("depth",
			mcp.Description("Scan depth: 'quick' or 'full'. Default: 'full'"),
			mcp.Enum("quick", "full"),
		),
	), cs.handleScanProject)

	// ===== PRIMITIVES =====

	// list_primitives
	cs.mcp.AddTool(mcp.NewTool("list_primitives",
		mcp.WithDescription("List all behavioral primitives — the normalized semantic vocabulary for TUI architecture. Primitives are stable; libraries implement them."),
		mcp.WithString("category",
			mcp.Description("Filter by category: lifecycle, async, rendering, state, coordination, infrastructure, testing"),
		),
	), cs.handleListPrimitives)

	// explain_primitive
	cs.mcp.AddTool(mcp.NewTool("explain_primitive",
		mcp.WithDescription("Explain a behavioral primitive — what it is, which libraries implement it, what signals indicate it, what constraints apply."),
		mcp.WithString("primitive_id",
			mcp.Required(),
			mcp.Description("Primitive ID from primitives.yaml, e.g. 'viewport', 'tick_source', 'storage_boundary'"),
		),
		mcp.WithBoolean("with_implementations",
			mcp.Description("Include library implementations. Default: true"),
		),
	), cs.handleExplainPrimitive)

	// ===== PATTERNS =====

	// list_patterns
	cs.mcp.AddTool(mcp.NewTool("list_patterns",
		mcp.WithDescription("List empirical patterns derived from ecosystem observations. Each pattern includes confidence score, observation count, and evidence references."),
		mcp.WithString("category",
			mcp.Description("Filter by category: lifecycle, async, rendering, state, coordination, infrastructure"),
		),
		mcp.WithNumber("min_confidence",
			mcp.Description("Minimum confidence threshold (0.0-1.0). Default: 0.75"),
		),
	), cs.handleListPatterns)

	// explain_pattern
	cs.mcp.AddTool(mcp.NewTool("explain_pattern",
		mcp.WithDescription("Explain an empirical pattern — what it is, confidence score, which repos use it, violations when absent, fixes that applied it."),
		mcp.WithString("pattern_id",
			mcp.Required(),
			mcp.Description("Pattern ID, e.g. 'viewport_content_diffing', 'storage_cmd_boundary'"),
		),
		mcp.WithBoolean("with_evidence",
			mcp.Description("Include full evidence (repos, issues, commits). Default: false"),
		),
	), cs.handleExplainPattern)

	// get_pattern_evidence
	cs.mcp.AddTool(mcp.NewTool("get_pattern_evidence",
		mcp.WithDescription("Get detailed evidence for a pattern — observations, violations, fixes with file paths, line numbers, commit SHAs."),
		mcp.WithString("pattern_id",
			mcp.Required(),
			mcp.Description("Pattern ID"),
		),
	), cs.handleGetPatternEvidence)

	// ===== CONSTRAINTS =====

	// list_constraints
	cs.mcp.AddTool(mcp.NewTool("list_constraints",
		mcp.WithDescription("List architectural constraints — both canonical (runtime semantics, confidence=1.0) and empirical (pattern-derived, confidence<1.0)."),
		mcp.WithString("type",
			mcp.Description("Filter by type: 'canonical', 'empirical', or 'all'. Default: 'all'"),
			mcp.Enum("canonical", "empirical", "all"),
		),
		mcp.WithNumber("min_confidence",
			mcp.Description("Minimum confidence threshold (0.0-1.0). Default: 0.75"),
		),
	), cs.handleListConstraints)

	// explain_constraint
	cs.mcp.AddTool(mcp.NewTool("explain_constraint",
		mcp.WithDescription("Explain a constraint — why it exists, confidence score, evidence (if empirical), violations, fixes, related patterns."),
		mcp.WithString("constraint_id",
			mcp.Required(),
			mcp.Description("Constraint ID, e.g. 'no_blocking_update', 'storage_must_be_async'"),
		),
		mcp.WithBoolean("with_evidence",
			mcp.Description("Include evidence for empirical constraints. Default: true"),
		),
	), cs.handleExplainConstraint)

	// ===== ARCHETYPES =====

	// list_archetypes
	cs.mcp.AddTool(mcp.NewTool("list_archetypes",
		mcp.WithDescription("List TUI archetypes — high-level intent categories with canonical components + empirical patterns."),
	), cs.handleListArchetypes)

	// get_archetype
	cs.mcp.AddTool(mcp.NewTool("get_archetype",
		mcp.WithDescription("Get full archetype definition including primitives, patterns, timing constraints, and typical topology."),
		mcp.WithString("archetype_id",
			mcp.Required(),
			mcp.Description("Archetype ID, e.g. 'agent_console', 'streaming_chat'"),
		),
	), cs.handleGetArchetype)

	// ===== TIMING MODELS =====

	// list_timing_models
	cs.mcp.AddTool(mcp.NewTool("list_timing_models",
		mcp.WithDescription("List timing models — perceptual contracts for TUI quality."),
	), cs.handleListTimingModels)

	// ===== LEGACY/COMPATIBILITY =====

	// explain_concept (legacy, redirects to primitive/pattern/constraint)
	cs.mcp.AddTool(mcp.NewTool("explain_concept",
		mcp.WithDescription("Legacy concept explanation. Use explain_primitive, explain_pattern, or explain_constraint instead."),
		mcp.WithString("concept",
			mcp.Required(),
			mcp.Description("Concept name"),
		),
	), cs.handleExplainConcept)

	// list_deepwiki
	cs.mcp.AddTool(mcp.NewTool("list_deepwiki",
		mcp.WithDescription("List DeepWiki reference documentation topics."),
		mcp.WithString("library",
			mcp.Description("Filter by library name"),
		),
	), cs.handleListDeepWiki)
}

// registerResources registers MCP resources.
func (cs *CharmedServer) registerResources() {
	// charmed://knowledge/components
	cs.mcp.AddResource(mcp.NewResource("charmed://knowledge/components",
		"Charm Component Knowledge",
		mcp.WithResourceDescription("Component semantics for the Charm ecosystem — viewport, list, textarea, etc."),
		mcp.WithMIMEType("application/json"),
	), cs.handleResourceComponents)

	// charmed://primitives
	cs.mcp.AddResource(mcp.NewResource("charmed://primitives",
		"Behavioral Primitives",
		mcp.WithResourceDescription("Normalized semantic vocabulary for TUI architecture — stable primitives that libraries implement."),
		mcp.WithMIMEType("application/json"),
	), cs.handleResourcePrimitives)

	// charmed://patterns (empirical)
	cs.mcp.AddResource(mcp.NewResource("charmed://patterns",
		"Empirical Patterns",
		mcp.WithResourceDescription("Evidence-based patterns derived from ecosystem observations with confidence scores."),
		mcp.WithMIMEType("application/json"),
	), cs.handleResourcePatterns)

	// charmed://knowledge/patterns (legacy compat)
	cs.mcp.AddResource(mcp.NewResource("charmed://knowledge/patterns",
		"TUI Pattern Knowledge (legacy)",
		mcp.WithResourceDescription("Legacy pattern resource. Use charmed://patterns instead."),
		mcp.WithMIMEType("application/json"),
	), cs.handleResourcePatterns)

	// charmed://knowledge/dependencies
	cs.mcp.AddResource(mcp.NewResource("charmed://knowledge/dependencies",
		"Charm Library Dependency Graph",
		mcp.WithResourceDescription("Dependency relationships between Charm ecosystem libraries."),
		mcp.WithMIMEType("application/json"),
	), cs.handleResourceDependencies)

	// charmed://archetypes
	cs.mcp.AddResource(mcp.NewResource("charmed://archetypes",
		"TUI Archetypes",
		mcp.WithResourceDescription("High-level TUI intent categories."),
		mcp.WithMIMEType("application/json"),
	), cs.handleResourceArchetypes)

	// charmed://constraints
	cs.mcp.AddResource(mcp.NewResource("charmed://constraints",
		"Constraint Rules",
		mcp.WithResourceDescription("Architectural lint rules for Bubble Tea projects."),
		mcp.WithMIMEType("application/json"),
	), cs.handleResourceConstraints)

	// charmed://timing-models
	cs.mcp.AddResource(mcp.NewResource("charmed://timing-models",
		"Timing Models",
		mcp.WithResourceDescription("Perceptual timing contracts for TUI quality."),
		mcp.WithMIMEType("application/json"),
	), cs.handleResourceTimingModels)

	// charmed://ir/schema
	cs.mcp.AddResource(mcp.NewResource("charmed://ir/schema",
		"TUI Structural IR Schema",
		mcp.WithResourceDescription("JSON Schema for the TUI Structural IR."),
		mcp.WithMIMEType("application/json"),
	), cs.handleResourceIRSchema)

	// DeepWiki resources (selective query into embedded docs)
	cs.mcp.AddResource(mcp.NewResource("charmed://deepwiki/{component}",
		"DeepWiki Component Reference",
		mcp.WithResourceDescription("DeepWiki documentation for a specific Charm ecosystem component or concept. Use 'list_deepwiki' to see available topics."),
		mcp.WithMIMEType("text/markdown"),
	), cs.handleResourceDeepWiki)
}

// Serve starts the MCP server on stdio.
func (cs *CharmedServer) Serve() error {
	return server.ServeStdio(cs.mcp)
}

// ServeContext starts the MCP server with a context.
func (cs *CharmedServer) ServeContext(ctx context.Context) error {
	s := server.NewStdioServer(cs.mcp)
	return s.Listen(ctx, os.Stdin, os.Stdout)
}

// ── Tool Handlers ──

func (cs *CharmedServer) handleScanProject(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid arguments")
	}

	projectPath, _ := args["project_path"].(string)
	if projectPath == "" {
		return nil, fmt.Errorf("project_path is required")
	}

	opts := scanner.ScanOptions{
		ProjectPath:   projectPath,
		IncludeVendor: getBoolArg(args, "include_vendor", false),
		AnalyzeTests:  getBoolArg(args, "analyze_tests", false),
		AtreeDBPath:   getStringArg(args, "atree_db_path", ""),
		Depth:         getStringArg(args, "depth", "full"),
	}

	result, err := cs.scanner.Scan(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal result: %w", err)
	}

	return mcp.NewToolResultText(string(data)), nil
}

func (cs *CharmedServer) handleListArchetypes(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	data, err := json.MarshalIndent(map[string]interface{}{
		"ontology_version": scanner.OntologyVersion,
		"archetypes":       cs.scanner.GetArchetypes(),
	}, "", "  ")
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}

func (cs *CharmedServer) handleListConstraints(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	data, err := json.MarshalIndent(map[string]interface{}{
		"constraint_version": scanner.ConstraintVersion,
		"constraints":        cs.scanner.GetConstraints(),
	}, "", "  ")
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}

func (cs *CharmedServer) handleListTimingModels(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	data, err := json.MarshalIndent(map[string]interface{}{
		"timing_model_version": scanner.TimingModelVersion,
		"timing_models":        cs.scanner.GetTimingModels(),
	}, "", "  ")
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}

func (cs *CharmedServer) handleExplainConcept(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	concept, _ := args["concept"].(string)
	if concept == "" {
		return nil, fmt.Errorf("concept is required")
	}
	withCode := getBoolArg(args, "with_code", false)

	onto := cs.scanner.GetOntology()
	if onto == nil {
		return mcp.NewToolResultText(fmt.Sprintf("Concept: %s\n\n(ontology not loaded)", concept)), nil
	}

	// Search components by name (case-insensitive, substring match)
	for _, c := range onto.Components {
		if strings.EqualFold(c.Name, concept) || strings.Contains(strings.ToLower(c.Name), strings.ToLower(concept)) {
			result := map[string]interface{}{
				"name":        c.Name,
				"package":     c.Package,
				"kind":        c.Kind,
				"owns_state":  c.OwnsState,
				"inputs":      c.Inputs,
				"outputs":     c.Outputs,
				"constraints": c.Constraints,
			}
			if withCode {
				result["anti_patterns"] = c.AntiPatterns
				result["common_pairings"] = c.CommonPairings
			}
			data, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(string(data)), nil
		}
	}

	// Search patterns by name
	for _, p := range onto.Patterns {
		if strings.EqualFold(p.Name, concept) || strings.Contains(strings.ToLower(p.Name), strings.ToLower(concept)) {
			data, err := json.MarshalIndent(p, "", "  ")
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(string(data)), nil
		}
	}

	return mcp.NewToolResultText(fmt.Sprintf("Concept: %s\n\nNot found in ontology. Try list_deepwiki for available topics.", concept)), nil
}

// ── Resource Handlers ──

func (cs *CharmedServer) handleResourceComponents(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	onto := cs.scanner.GetOntology()
	if onto == nil {
		return nil, fmt.Errorf("ontology not loaded")
	}
	data, err := json.MarshalIndent(onto.Components, "", "  ")
	if err != nil {
		return nil, err
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "charmed://knowledge/components",
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

func (cs *CharmedServer) handleResourcePatterns(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	onto := cs.scanner.GetOntology()
	if onto == nil {
		return nil, fmt.Errorf("ontology not loaded")
	}
	data, err := json.MarshalIndent(onto.Patterns, "", "  ")
	if err != nil {
		return nil, err
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "charmed://knowledge/patterns",
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

func (cs *CharmedServer) handleResourceDependencies(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	graph := cs.scanner.GetDependencyGraph()
	if graph == nil {
		return nil, fmt.Errorf("dependency graph not loaded")
	}
	data, err := json.MarshalIndent(graph, "", "  ")
	if err != nil {
		return nil, err
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "charmed://knowledge/dependencies",
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

func (cs *CharmedServer) handleResourceArchetypes(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	data, err := json.MarshalIndent(cs.scanner.GetArchetypes(), "", "  ")
	if err != nil {
		return nil, err
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "charmed://archetypes",
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

func (cs *CharmedServer) handleResourceConstraints(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	data, err := json.MarshalIndent(cs.scanner.GetConstraints(), "", "  ")
	if err != nil {
		return nil, err
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "charmed://constraints",
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

func (cs *CharmedServer) handleResourceTimingModels(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	data, err := json.MarshalIndent(cs.scanner.GetTimingModels(), "", "  ")
	if err != nil {
		return nil, err
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "charmed://timing-models",
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

func (cs *CharmedServer) handleResourceIRSchema(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	// Load the IR schema from the embedded copy
	data, err := ir.LoadIRSchema()
	if err != nil {
		return nil, fmt.Errorf("load IR schema: %w", err)
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "charmed://ir/schema",
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

func (cs *CharmedServer) handleResourceDeepWiki(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	uri := req.Params.URI
	data, err := resolveDeepWiki(uri)
	if err != nil {
		return nil, err
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      uri,
			MIMEType: "text/markdown",
			Text:     string(data),
		},
	}, nil
}

func (cs *CharmedServer) handleListDeepWiki(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	library, _ := args["library"].(string)

	entries := listDeepWiki(library)
	result := map[string]interface{}{
		"count":   len(entries),
		"library": library,
		"topics":  entries,
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}

// ── Helpers ──

func getBoolArg(args map[string]interface{}, key string, defaultVal bool) bool {
	if v, ok := args[key].(bool); ok {
		return v
	}
	return defaultVal
}

func getStringArg(args map[string]interface{}, key string, defaultVal string) string {
	if v, ok := args[key].(string); ok {
		return v
	}
	return defaultVal
}

// ── Implemented Handlers ──

func (cs *CharmedServer) handleExtractEvidence(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	projectPath := getStringArg(args, "project_path", "")
	if projectPath == "" {
		return nil, fmt.Errorf("project_path is required")
	}

	opts := scanner.ScanOptions{
		ProjectPath: projectPath,
		Depth:       "full",
		AtreeDBPath: getStringArg(args, "atree_db_path", ""),
	}

	result, err := cs.scanner.Scan(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("evidence extraction failed: %w", err)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}

func (cs *CharmedServer) handleListPrimitives(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	category := getStringArg(args, "category", "")

	primitives := cs.scanner.GetPrimitives()
	if category != "" {
		var filtered []scanner.Primitive
		for _, p := range primitives {
			if p.Category == category {
				filtered = append(filtered, p)
			}
		}
		primitives = filtered
	}

	data, err := json.MarshalIndent(map[string]interface{}{
		"count":      len(primitives),
		"primitives": primitives,
	}, "", "  ")
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}

func (cs *CharmedServer) handleExplainPrimitive(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	id := getStringArg(args, "primitive_id", "")
	if id == "" {
		return nil, fmt.Errorf("primitive_id is required")
	}
	for _, p := range cs.scanner.GetPrimitives() {
		if p.ID == id {
			data, err := json.MarshalIndent(p, "", "  ")
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(string(data)), nil
		}
	}
	return mcp.NewToolResultText(fmt.Sprintf(
		`{"error":"primitive '%s' not found","available_ids":%s}`,
		id, primitiveIDList(cs.scanner.GetPrimitives()),
	)), nil
}

func (cs *CharmedServer) handleListPatterns(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	category := getStringArg(args, "category", "")
	minConf := getNumberArg(args, "min_confidence", 0.75)

	var filtered []scanner.Pattern
	for _, p := range cs.scanner.GetPatterns() {
		if p.Confidence < minConf {
			continue
		}
		if category != "" && p.Category != category {
			continue
		}
		filtered = append(filtered, p)
	}

	data, err := json.MarshalIndent(map[string]interface{}{
		"count":    len(filtered),
		"patterns": filtered,
	}, "", "  ")
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}

func (cs *CharmedServer) handleExplainPattern(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	id := getStringArg(args, "pattern_id", "")
	if id == "" {
		return nil, fmt.Errorf("pattern_id is required")
	}
	for _, p := range cs.scanner.GetPatterns() {
		if p.ID == id {
			withEvidence := getBoolArg(args, "with_evidence", false)
			if withEvidence {
				data, err := json.MarshalIndent(p, "", "  ")
				if err != nil {
					return nil, err
				}
				return mcp.NewToolResultText(string(data)), nil
			}
			// Return summary without evidence
			summary := map[string]interface{}{
				"id":                      p.ID,
				"name":                    p.Name,
				"category":                p.Category,
				"primitives":              p.Primitives,
				"confidence":              p.Confidence,
				"confidence_factors":      p.ConfidenceFactors,
				"applies_to":              p.AppliesTo,
				"topology_signature":      p.TopologySignature,
				"synthesized_constraints": p.SynthesizedConstraints,
			}
			data, err := json.MarshalIndent(summary, "", "  ")
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(string(data)), nil
		}
	}
	return mcp.NewToolResultText(fmt.Sprintf(`{"error":"pattern '%s' not found"}`, id)), nil
}

func (cs *CharmedServer) handleGetPatternEvidence(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	id := getStringArg(args, "pattern_id", "")
	if id == "" {
		return nil, fmt.Errorf("pattern_id is required")
	}
	for _, p := range cs.scanner.GetPatterns() {
		if p.ID == id {
			data, err := json.MarshalIndent(map[string]interface{}{
				"pattern_id": p.ID,
				"confidence": p.Confidence,
				"evidence":   p.DerivedFrom,
			}, "", "  ")
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(string(data)), nil
		}
	}
	return mcp.NewToolResultText(fmt.Sprintf(`{"error":"pattern '%s' not found"}`, id)), nil
}

func (cs *CharmedServer) handleExplainConstraint(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	id := getStringArg(args, "constraint_id", "")
	if id == "" {
		return nil, fmt.Errorf("constraint_id is required")
	}

	// Search v2 constraints first
	for _, c := range cs.scanner.GetConstraintsV2() {
		if c.ID == id {
			withEvidence := getBoolArg(args, "with_evidence", true)
			if !withEvidence {
				c.Evidence = scanner.ConstraintV2Evidence{}
			}
			data, err := json.MarshalIndent(c, "", "  ")
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(string(data)), nil
		}
	}

	// Fall back to v1 constraints
	for _, c := range cs.scanner.GetConstraints() {
		if c.ID == id {
			data, err := json.MarshalIndent(c, "", "  ")
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(string(data)), nil
		}
	}

	return mcp.NewToolResultText(fmt.Sprintf(`{"error":"constraint '%s' not found"}`, id)), nil
}

func (cs *CharmedServer) handleGetArchetype(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := req.Params.Arguments.(map[string]interface{})
	id := getStringArg(args, "archetype_id", "")
	if id == "" {
		return nil, fmt.Errorf("archetype_id is required")
	}
	for _, a := range cs.scanner.GetArchetypes() {
		if a.ID == id {
			data, err := json.MarshalIndent(a, "", "  ")
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(string(data)), nil
		}
	}
	return mcp.NewToolResultText(fmt.Sprintf(
		`{"error":"archetype '%s' not found","available_ids":["agent_console","streaming_chat","ide_workspace","monitoring_dashboard","cli_wizard","git_client","log_explorer","status_console"]}`,
		id,
	)), nil
}

func (cs *CharmedServer) handleResourcePrimitives(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	data, err := json.MarshalIndent(map[string]interface{}{
		"count":      len(cs.scanner.GetPrimitives()),
		"primitives": cs.scanner.GetPrimitives(),
	}, "", "  ")
	if err != nil {
		return nil, err
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "charmed://primitives",
			MIMEType: "application/json",
			Text:     string(data),
		},
	}, nil
}

// ── Additional Helpers ──

func getNumberArg(args map[string]interface{}, key string, defaultVal float64) float64 {
	if v, ok := args[key].(float64); ok {
		return v
	}
	return defaultVal
}

func primitiveIDList(primitives []scanner.Primitive) string {
	ids := make([]string, len(primitives))
	for i, p := range primitives {
		ids[i] = `"` + p.ID + `"`
	}
	return "[" + strings.Join(ids, ",") + "]"
}
