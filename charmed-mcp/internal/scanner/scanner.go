// Package scanner implements the full scan_project pipeline.
package scanner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/B-A-M-N/charmed/charmed-mcp/internal/ir"
	"github.com/B-A-M-N/charmed/charmed-mcp/internal/signals"
	"gopkg.in/yaml.v3"
)

// CharmedVersion is the current version of the MCP server.
const CharmedVersion = "0.1.0"

// IRVersion is the TUI Structural IR schema version.
const IRVersion = "1"

// OntologyVersion is the knowledge base version.
const OntologyVersion = "1"

// ConstraintVersion is the constraints.yaml version.
const ConstraintVersion = "1"

// TimingModelVersion is the timing-models.yaml version.
const TimingModelVersion = "1"

// Scanner orchestrates the full scan pipeline.
type Scanner struct {
	// ATree MCP client (may be nil if unavailable)
	atree AtreeClient

	// Embedded knowledge bases
	constraints  []ConstraintRule
	archetypes   []ArchetypeDef
	ontology     *Ontology
	timingModels []TimingModel
	depGraph     map[string][]string

	// Extended knowledge bases (primitives, patterns, constraints-v2)
	primitives    []Primitive
	patterns      []Pattern
	constraintsV2 []ConstraintV2

	// Signal collector
	signals *signals.Collector
}

// AtreeClient is the interface for ATree MCP communication.
// Implemented by atree.Client (MCP client to ATree server).
type AtreeClient interface {
	Index(ctx context.Context, path string, force bool) error
	Query(ctx context.Context, query string) (*AtreeQueryResult, error)
	Context(ctx context.Context, name string) (*AtreeContextResult, error)
	FindEntrypoints(ctx context.Context) (*AtreeEntrypointsResult, error)
	ConcurrencySurface(ctx context.Context) (*AtreeConcurrencyResult, error)
	SideEffects(ctx context.Context, symbol string) (*AtreeSideEffectsResult, error)
	ResolutionStats(ctx context.Context) (*AtreeStatsResult, error)
	// Bulk graph extraction
	GetSemanticGraph(ctx context.Context, files []string) (*AtreeSemanticGraph, error)
}

// AtreeQueryResult is the result of an ATree query.
type AtreeQueryResult struct {
	Symbols    []AtreeSymbol    `json:"symbols"`
	Confidence float64          `json:"confidence"`
}

// AtreeSymbol is a symbol from ATree.
type AtreeSymbol struct {
	Name      string `json:"name"`
	Qualified string `json:"qualified_name"`
	Kind      string `json:"kind"` // "DefinitionStruct", "DefinitionMethod", etc.
	File      string `json:"file"`
	Line      int    `json:"line"`
	Col       int    `json:"col"`

	// Post-upgrade ATree fields
	Receiver   string       `json:"receiver,omitempty"`
	IsPtrRecv  bool         `json:"is_pointer_receiver,omitempty"`
	Params     []AtreeParam `json:"params,omitempty"`
	Returns    []string     `json:"returns,omitempty"`
	BodyStart  int          `json:"body_start,omitempty"`
	BodyEnd    int          `json:"body_end,omitempty"`
}

// AtreeParam is a method parameter.
type AtreeParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// AtreeContextResult is the result of an ATree context query.
type AtreeContextResult struct {
	Symbol     AtreeSymbol      `json:"symbol"`
	Callers    []AtreeCallEdge  `json:"callers"`
	Callees    []AtreeCallEdge  `json:"callees"`
	Heritage   []AtreeHeritage  `json:"heritage"`
	Confidence float64          `json:"confidence"`
}

// AtreeCallEdge is a call relationship.
type AtreeCallEdge struct {
	Symbol     string  `json:"symbol"`
	File       string  `json:"file"`
	Line       int     `json:"line"`
	Confidence float64 `json:"confidence"`
}

// AtreeHeritage is a type relationship.
type AtreeHeritage struct {
	Parent string `json:"parent"`
	Child  string `json:"child"`
	Kind   string `json:"kind"` // "extends", "embeds", "implements"
	File   string `json:"file"`
	Line   int    `json:"line"`
}

// AtreeEntrypointsResult contains entry point analysis.
type AtreeEntrypointsResult struct {
	Entrypoints []AtreeEntrypoint `json:"entrypoints"`
}

// AtreeEntrypoint is a code entry point.
type AtreeEntrypoint struct {
	Symbol string `json:"symbol"` // e.g., "main.AppModel"
	Kind   string `json:"kind"`   // "main", "tea.NewProgram", "handler"
	File   string `json:"file"`
	Line   int    `json:"line"`
}

// AtreeConcurrencyResult contains concurrency surface analysis.
type AtreeConcurrencyResult struct {
	Goroutines []AtreeGoroutine `json:"goroutines"`
	Channels   []AtreeChannel   `json:"channels"`
}

// AtreeGoroutine is a goroutine spawn site.
type AtreeGoroutine struct {
	InFunction string `json:"in_function"`
	File       string `json:"file"`
	Line       int    `json:"line"`
	IsInLoop   bool   `json:"is_in_loop"`
}

// AtreeChannel is a channel operation.
type AtreeChannel struct {
	Operation  string `json:"operation"` // "send" | "receive" | "close"
	InFunction string `json:"in_function"`
	File       string `json:"file"`
	Line       int    `json:"line"`
}

// AtreeSideEffectsResult contains side effect analysis.
type AtreeSideEffectsResult struct {
	Effects []AtreeSideEffect `json:"effects"`
}

// AtreeSideEffect is a detected side effect.
type AtreeSideEffect struct {
	Kind     string `json:"kind"` // "io", "network", "sleep", "db"
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
	Pattern  string `json:"pattern"` // matched code pattern
}

// AtreeStatsResult contains index statistics.
type AtreeStatsResult struct {
	Files          int     `json:"files"`
	Symbols        int     `json:"symbols"`
	Calls          int     `json:"calls"`
	ResolvedCalls  int     `json:"resolved_calls"`
	UnresolvedRate float64 `json:"unresolved_rate"`
}

// AtreeSemanticGraph is the full semantic graph for Go files.
type AtreeSemanticGraph struct {
	Files          []AtreeFileSem `json:"files"`
	Calls          []AtreeCall    `json:"calls"`
	Imports        []AtreeImport  `json:"imports"`
	Assignments    []AtreeAssign  `json:"assignments"`
	TypeBindings   []AtreeBind    `json:"type_bindings"`
}

// AtreeFileSem is semantic data for one file.
type AtreeFileSem struct {
	Path           string           `json:"path"`
	Language       string           `json:"language"`
	Defs           []AtreeSymbol    `json:"defs"`
	SwitchCases    []AtreeSwitch    `json:"switch_case_bodies"`
	FunctionCalls  []AtreeFuncCall  `json:"calls"`
}

// AtreeSwitch is a switch statement.
type AtreeSwitch struct {
	Scope   string        `json:"scope"`   // enclosing function
	VarType string        `json:"var_tea"` // "tea.Msg" if type switch on tea.Msg
	Cases   []AtreeCase   `json:"cases"`
	File    string        `json:"file"`
	Line    int           `json:"line"`
}

// AtreeCase is a switch case branch.
type AtreeCase struct {
	Type      string `json:"type"`     // matched type
	BodyStart int    `json:"body_start"`
	BodyEnd   int    `json:"body_end"`
	File      string `json:"file"`
	Line      int    `json:"line"`
}

// AtreeFuncCall is a function call site.
type AtreeFuncCall struct {
	Name     string `json:"name"`
	Recv     string `json:"receiver,omitempty"`
	Scope    string `json:"scope"`    // enclosing function
	File     string `json:"file"`
	Line     int    `json:"line"`
}

// AtreeCall is a resolved call edge.
type AtreeCall struct {
	CallerScope string  `json:"caller_scope"`
	CalleeName  string  `json:"callee_name"`
	Receiver    string  `json:"receiver,omitempty"`
	Resolved    bool    `json:"resolved"`
	Confidence  float64 `json:"confidence"`
	File        string  `json:"file"`
	Line        int     `json:"line"`
	Col         int     `json:"col"`
}

// AtreeImport is an import declaration.
type AtreeImport struct {
	File      string `json:"file"`
	LocalName string `json:"local_name"`
	FullPath  string `json:"full_path"`
	IsAlias   bool   `json:"is_alias"`
}

// AtreeAssign is an assignment.
type AtreeAssign struct {
	Field      string `json:"field"`
	Source     string `json:"source"`
	InFunction string `json:"in_function"`
	File       string `json:"file"`
	Line       int    `json:"line"`
}

// AtreeBind is a type binding.
type AtreeBind struct {
	VarName    string `json:"var"`
	Type       string `json:"type"`
	InFunction string `json:"in_function"`
	File       string `json:"file"`
	Line       int    `json:"line"`
}

// ConstraintRule is loaded from constraints.yaml.
type ConstraintRule struct {
	ID          string   `json:"id"`
	Severity    string   `json:"severity"`    // "info" | "warning" | "error" | "fatal"
	Category    string   `json:"category"`    // "correctness" | "performance" | etc.
	Description string   `json:"description"`
	Fix         string   `json:"fix"`
	// SignalKinds lists the signal kinds that trigger this constraint.
	SignalKinds []string `json:"signal_kinds"`
	// Threshold is an optional numeric threshold (e.g. LOC limit).
	Threshold  int    `json:"threshold,omitempty"`
}

// ArchetypeDef is loaded from archetypes.yaml.
type ArchetypeDef struct {
	ID               string   `json:"id"`
	Description      string   `json:"description"`
	Traits           []string `json:"traits"`
	RequiredComponents []string `json:"required_components"`
	LayoutType       string   `json:"layout_type"`
	DensityProfile   string   `json:"density_profile"`
	MotionPhilosophy string   `json:"motion_philosophy"`
	TimingConstraints []string `json:"timing_constraints"`
	FailureModes     []string `json:"failure_modes"`
}

// Ontology is the loaded knowledge.yaml.
type Ontology struct {
	Components []OntologyComponent `json:"components"`
	Patterns   []OntologyPattern   `json:"patterns"`
}

// OntologyComponent is a component definition.
type OntologyComponent struct {
	Name           string   `json:"name"`
	Package        string   `json:"package"`
	Kind           string   `json:"kind"`
	OwnsState      []string `json:"owns_state"`
	Inputs         []string `json:"inputs"`
	Outputs        []string `json:"outputs"`
	Constraints    []string `json:"constraints"`
	AntiPatterns   []string `json:"anti_patterns"`
	CommonPairings []string `json:"common_pairings"`
}

// OntologyPattern is a pattern definition.
type OntologyPattern struct {
	Name          string   `json:"name"`
	Uses          []string `json:"uses"`
	ComposesWith  []string `json:"composes_with"`
	Emit          []string `json:"emits"`
	FailureModes  []string `json:"failure_modes"`
}

// TimingModel is a timing contract.
type TimingModel struct {
	ID               string            `json:"id"`
	Description      string            `json:"description"`
	Constraints      map[string]int    `json:"constraints"`       // name → value in ms
	PerceptionTraits []string          `json:"perception_traits"`
}

// ScanOptions configures a scan.
type ScanOptions struct {
	ProjectPath    string
	IncludeVendor  bool
	AnalyzeTests   bool
	AtreeDBPath    string
	Depth          string // "quick" | "full"
}

// NewScanner creates a Scanner with embedded knowledge bases.
func NewScanner(atree AtreeClient) (*Scanner, error) {
	s := &Scanner{
		atree:   atree,
		signals: signals.NewCollector(),
	}

	// Load embedded knowledge bases
	if err := s.loadKnowledge(); err != nil {
		return nil, fmt.Errorf("load knowledge: %w", err)
	}

	return s, nil
}

// loadKnowledge loads embedded YAML knowledge bases from ../knowledge/.
func (s *Scanner) loadKnowledge() error {
	// These paths are relative to the compiled binary.
	// In development, the knowledge dir is at ../../knowledge/ from the binary.
	knowledgeDir := s.findKnowledgeDir()

	// Load constraints
	if data, err := os.ReadFile(filepath.Join(knowledgeDir, "constraints.yaml")); err == nil {
		s.loadConstraints(data)
	} else {
		return fmt.Errorf("constraints.yaml: %w", err)
	}

	// Load archetypes
	if data, err := os.ReadFile(filepath.Join(knowledgeDir, "archetypes.yaml")); err == nil {
		s.loadArchetypes(data)
	} else {
		return fmt.Errorf("archetypes.yaml: %w", err)
	}

	// Load ontology
	if data, err := os.ReadFile(filepath.Join(knowledgeDir, "knowledge.yaml")); err == nil {
		s.loadOntology(data)
	} else {
		return fmt.Errorf("knowledge.yaml: %w", err)
	}

	// Load timing models
	if data, err := os.ReadFile(filepath.Join(knowledgeDir, "timing-models.yaml")); err == nil {
		s.loadTimingModels(data)
	} else {
		return fmt.Errorf("timing-models.yaml: %w", err)
	}

		// Load dependency graph from knowledge.yaml
		if data, err := os.ReadFile(filepath.Join(knowledgeDir, "knowledge.yaml")); err == nil {
			s.loadDepGraph(data)
		}

		// Load extended knowledge bases (soft failures — optional)
		s.loadExtendedKnowledge(knowledgeDir)

		return nil
	}

// findKnowledgeDir locates the knowledge directory.
// Tries several paths relative to the working directory and binary location.
func (s *Scanner) findKnowledgeDir() string {
	candidates := []string{
		"knowledge",                           // same dir
		"../../knowledge",                     // from cmd/charmed-mcp/
		"../knowledge",                        // from charmed-mcp/
		"../../../charmed/knowledge",          // from charmed-mcp/cmd/charmed-mcp/
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	return "knowledge" // fallback
}

func (s *Scanner) loadConstraints(data []byte) {
	// Parse constraints.yaml into []ConstraintRule
	// Simple YAML parsing — each document separated by "---"
	// For now, use a structured approach
	var rules []ConstraintRule
	raw := splitYAMLDocuments(string(data))
	for _, doc := range raw {
		var r ConstraintRule
		if err := unmarshalYAML(doc, &r); err == nil && r.ID != "" {
			rules = append(rules, r)
		}
	}
	s.constraints = rules
}

func (s *Scanner) loadArchetypes(data []byte) {
	var defs []ArchetypeDef
	raw := splitYAMLDocuments(string(data))
	for _, doc := range raw {
		var d ArchetypeDef
		if err := unmarshalYAML(doc, &d); err == nil && d.ID != "" {
			defs = append(defs, d)
		}
	}
	s.archetypes = defs
}

func (s *Scanner) loadOntology(data []byte) {
	var o Ontology
	if err := unmarshalYAML(string(data), &o); err == nil {
		s.ontology = &o
	}
}

func (s *Scanner) loadTimingModels(data []byte) {
	var models []TimingModel
	raw := splitYAMLDocuments(string(data))
	for _, doc := range raw {
		var m TimingModel
		if err := unmarshalYAML(doc, &m); err == nil && m.ID != "" {
			models = append(models, m)
		}
	}
	s.timingModels = models
}

func (s *Scanner) loadDepGraph(data []byte) {
	// The dependency graph is at the end of knowledge.yaml as a YAML map.
	// Split on "---" to isolate it, then parse manually.
	text := string(data)
	// Find the bare map section (lines not in a YAML document with list items)
	lines := strings.Split(text, "\n")
	inDepGraph := false
	var graphLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Detect start of dep graph: key with empty list value, no leading "-"
		if !inDepGraph && strings.HasSuffix(trimmed, ": []") {
			inDepGraph = true
		}
		if inDepGraph {
			graphLines = append(graphLines, line)
		}
	}

	graph := make(map[string][]string)
	for _, line := range graphLines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		// Parse "lib:" or "lib: [dep1, dep2]"
		idx := strings.Index(trimmed, ":")
		if idx < 0 {
			continue
		}
		key := strings.TrimSpace(trimmed[:idx])
		val := strings.TrimSpace(trimmed[idx+1:])

		if val == "[]" || val == "" {
			graph[key] = []string{}
		} else if strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]") {
			// Parse [dep1, dep2]
			inner := strings.TrimSuffix(strings.TrimPrefix(val, "["), "]")
			inner = strings.TrimSuffix(inner, "]")
			var deps []string
			for _, d := range strings.Split(inner, ",") {
				d = strings.TrimSpace(d)
				if d != "" {
					deps = append(deps, d)
				}
			}
			graph[key] = deps
		}
	}
	s.depGraph = graph
}

// splitYAMLDocuments splits a multi-document YAML file.
func splitYAMLDocuments(s string) []string {
	var docs []string
	for _, d := range strings.Split(s, "\n---") {
		d = strings.TrimSpace(d)
		if d != "" {
			docs = append(docs, d)
		}
	}
	return docs
}

// unmarshalYAML unmarshals a YAML document into v.
func unmarshalYAML(data string, v interface{}) error {
	return yaml.Unmarshal([]byte(data), v)
}

// ── Accessors for MCP server handlers ──

// GetConstraints returns the loaded constraint rules.
func (s *Scanner) GetConstraints() []ConstraintRule {
	return s.constraints
}

// GetArchetypes returns the loaded archetype definitions.
func (s *Scanner) GetArchetypes() []ArchetypeDef {
	return s.archetypes
}

// GetTimingModels returns the loaded timing models.
func (s *Scanner) GetTimingModels() []TimingModel {
	return s.timingModels
}

// GetOntology returns the loaded ontology (components + patterns).
func (s *Scanner) GetOntology() *Ontology {
	return s.ontology
}

// GetDependencyGraph returns the Charm library dependency graph.
// Parsed from the bottom section of knowledge.yaml.
func (s *Scanner) GetDependencyGraph() map[string][]string {
	if s.ontology == nil {
		return nil
	}
	// The dependency graph is stored in the ontology's raw form.
	// We parse it from the same knowledge.yaml via a dedicated field.
	return s.depGraph
}

// Scan executes the full scan pipeline.
func (s *Scanner) Scan(ctx context.Context, opts ScanOptions) (*ir.ScanResult, error) {
	startTime := time.Now()
	result := &ir.ScanResult{
		Warnings: []ir.Warning{},
	}

	// Stage 1: Bootstrap
	boot, err := Bootstrap(opts.ProjectPath, opts.IncludeVendor, opts.AnalyzeTests)
	if err != nil {
		return nil, fmt.Errorf("bootstrap: %w", err)
	}

	for _, w := range boot.Warnings {
		result.Warnings = append(result.Warnings, ir.Warning{Type: w, Message: w})
	}

	// Build scan_meta early
	result.ScanMeta = &ir.ScanMeta{
		CharmedVersion:     CharmedVersion,
		IRVersion:          IRVersion,
		OntologyVersion:    OntologyVersion,
		ConstraintVersion:  ConstraintVersion,
		TimingModelVersion: TimingModelVersion,
		ScannedAt:          startTime,
		FilesScanned:       len(boot.GoFiles),
		GoFiles:            len(boot.GoFiles),
		CharmDetected:      len(boot.CharmDeps) > 0,
	}

	if boot.AtreeAvailable {
		result.ScanMeta.Atree = &ir.AtreeMeta{
			Available: true,
			DBPath:    boot.AtreeDBPath,
		}
	}

	// Quick mode: stop after bootstrap
	if opts.Depth == "quick" {
		result.ScanMeta.ElapsedMs = time.Since(startTime).Milliseconds()
		result.IR = &ir.TUI_IR{
			SourceRepo: boot.ProjectRoot,
			IndexedAt:  startTime.Format(time.RFC3339),
		}
		return result, nil
	}

	// Stage 2: ATree Index (optional enrichment — only if atree is available)
	if s.atree != nil {
		if err := s.atree.Index(ctx, boot.ProjectRoot, false); err != nil {
			result.Warnings = append(result.Warnings, ir.Warning{
				Type:    "atree_index_failed",
				Message: err.Error(),
			})
		}
	}

	// Stage 3: Graph Normalization
	pgraph, err := s.normalizeGraph(ctx, boot)
	if err != nil {
		return nil, fmt.Errorf("graph normalization: %w", err)
	}

	// Stage 4: Model Detection
	models, err := s.detectModels(ctx, pgraph, boot)
	if err != nil {
		return nil, fmt.Errorf("model detection: %w", err)
	}

	// Stage 5: Message Flow
	messages, flowGraph, err := s.detectMessageFlow(ctx, pgraph, models)
	if err != nil {
		return nil, fmt.Errorf("message flow: %w", err)
	}

	// Stage 6: Component Binding
	components, err := s.detectComponents(ctx, pgraph, models)
	if err != nil {
		return nil, fmt.Errorf("component binding: %w", err)
	}

	// Stage 7: Async Boundaries
	asyncBounds, err := s.detectAsyncBoundaries(ctx, pgraph, models)
	if err != nil {
		return nil, fmt.Errorf("async boundaries: %w", err)
	}

	// Stage 8: State & Mutation Analysis
	stateOwn, mutations, err := s.analyzeStateMutation(ctx, pgraph, models)
	if err != nil {
		return nil, fmt.Errorf("state mutation: %w", err)
	}

	// Stage 9: View Analysis
	renderTree, layoutTree, err := s.analyzeViews(ctx, pgraph, models)
	if err != nil {
		return nil, fmt.Errorf("view analysis: %w", err)
	}

	// Stage 10: Constraint Precomputation
	s.evaluateConstraints()

	// Stage 11: IR Assembly
	result.IR = &ir.TUI_IR{
		SourceRepo:        boot.ProjectRoot,
		IndexedAt:         startTime.Format(time.RFC3339),
		AtreeDBPath:       boot.AtreeDBPath,
		Models:            models,
		Messages:          messages,
		MessageFlowGraph:  flowGraph,
		ComponentBindings: components,
		AsyncBoundaries:   asyncBounds,
		StateOwnership:    stateOwn,
		MutationPaths:     mutations,
		RenderTree:        renderTree,
		LayoutTree:        layoutTree,
	}

	// Stage 12: Validation
	result.Warnings = append(result.Warnings, s.validateIR(result.IR)...)

	result.ScanMeta.ElapsedMs = time.Since(startTime).Milliseconds()

	return result, nil
}

// ── Pipeline stage implementations ──

func (s *Scanner) normalizeGraph(ctx context.Context, boot *BootstrapResult) (*AtreeSemanticGraph, error) {
	return buildGraphFromAST(boot.GoFiles)
}

func (s *Scanner) detectModels(ctx context.Context, graph *AtreeSemanticGraph, boot *BootstrapResult) ([]ir.ModelDef, error) {
	return detectModelsFromGraph(graph)
}

func (s *Scanner) detectMessageFlow(ctx context.Context, graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.MessageDef, []ir.MessageFlowEdge, error) {
	return detectMessageFlowFromGraph(graph, models)
}

func (s *Scanner) detectComponents(ctx context.Context, graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.ComponentBinding, error) {
	return detectComponentsFromGraph(graph, models)
}

func (s *Scanner) detectAsyncBoundaries(ctx context.Context, graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.AsyncBoundary, error) {
	return detectAsyncBoundariesFromGraph(graph, models)
}

func (s *Scanner) analyzeStateMutation(ctx context.Context, graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.StateOwnership, []ir.MutationPath, error) {
	return analyzeStateMutationFromGraph(graph, models)
}

func (s *Scanner) analyzeViews(ctx context.Context, graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.RenderNode, []ir.LayoutNode, error) {
	return analyzeViewsFromGraph(graph, models)
}

func (s *Scanner) evaluateConstraints() {
	// Stage 10: evaluate signals against constraint rules
}

func (s *Scanner) validateIR(irDoc *ir.TUI_IR) []ir.Warning {
	var warnings []ir.Warning
	if len(irDoc.Models) == 0 {
		warnings = append(warnings, ir.Warning{
			Type:    "no_models_detected",
			Message: "No Bubble Tea models detected in project",
		})
	}
	return warnings
}
