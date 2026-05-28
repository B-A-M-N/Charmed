// Package ir defines the TUI Structural IR types.
// These conform to tui-structural-ir.schema.json.
package ir

import (
	"embed"
	"time"
)

//go:embed tui-structural-ir.schema.json
var schemaFS embed.FS

// LoadIRSchema loads the embedded IR schema.
func LoadIRSchema() ([]byte, error) {
	return schemaFS.ReadFile("tui-structural-ir.schema.json")
}

// ScanResult is the top-level output of scan_project.
type ScanResult struct {
	IR         *TUI_IR    `json:"ir"`
	ScanMeta   *ScanMeta  `json:"scan_meta"`
	Warnings   []Warning  `json:"warnings"`
}

// ScanMeta carries scan metadata.
type ScanMeta struct {
	CharmedVersion    string        `json:"charmed_version"`
	IRVersion         string        `json:"ir_version"`
	OntologyVersion   string        `json:"ontology_version"`
	ConstraintVersion string        `json:"constraint_version"`
	TimingModelVersion string       `json:"timing_model_version"`
	ScannedAt         time.Time     `json:"scanned_at"`
	ElapsedMs         int64         `json:"elapsed_ms"`
	FilesScanned      int           `json:"files_scanned"`
	GoFiles           int           `json:"go_files"`
	CharmDetected     bool          `json:"charm_detected"`
	Atree             *AtreeMeta    `json:"atree,omitempty"`
}

// AtreeMeta describes ATree enrichment status.
type AtreeMeta struct {
	Available    bool              `json:"available"`
	DBPath       string            `json:"db_path,omitempty"`
	Enrichments  map[string]bool   `json:"enrichments"`
	Unresolved   int               `json:"unresolved_symbols"`
	Confidence   float64           `json:"confidence"`
}

// Warning is a non-fatal scan warning.
type Warning struct {
	Type       string  `json:"type"`
	Message    string  `json:"message"`
	Confidence float64 `json:"confidence,omitempty"`
}

// TUI_IR is the canonical intermediate representation.
type TUI_IR struct {
	SourceRepo       string             `json:"source_repo"`
	IndexedAt        string             `json:"indexed_at"`
	AtreeDBPath      string             `json:"atree_db_path,omitempty"`
	Models           []ModelDef         `json:"models"`
	Messages         []MessageDef       `json:"messages"`
	Cmds             []CmdDef           `json:"cmds"`
	MessageFlowGraph []MessageFlowEdge  `json:"message_flow_graph"`
	RenderTree       []RenderNode       `json:"render_tree"`
	LayoutTree       []LayoutNode       `json:"layout_tree"`
	StateOwnership   []StateOwnership   `json:"state_ownership"`
	MutationPaths    []MutationPath     `json:"mutation_paths"`
	AsyncBoundaries  []AsyncBoundary    `json:"async_boundaries"`
	ComponentBindings []ComponentBinding `json:"component_bindings"`
}

// ModelDef is a Bubble Tea model.
type ModelDef struct {
	Name            string   `json:"name"`
	File            string   `json:"file"`
	Line            int      `json:"line"`
	Fields          []Field  `json:"fields"`
	Implements      []string `json:"implements"` // "Init" | "Update" | "View"
	UpdateComplexity int     `json:"update_complexity"`
	ViewComplexity  int      `json:"view_complexity"`
	Role            string   `json:"role"` // "root" | "child" | "standalone"
}

// Field is a model struct field.
type Field struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	MutationAuthority string `json:"mutation_authority,omitempty"`
}

// MessageDef is a custom message type.
type MessageDef struct {
	Name    string       `json:"name"`
	File    string       `json:"file"`
	Line    int          `json:"line"`
	Handlers []Handler   `json:"handlers"`
	EmittedBy []Emitter  `json:"emitted_by"`
}

// Handler is a message handler site.
type Handler struct {
	Model    string `json:"model"`
	Pattern  string `json:"pattern"`
}

// Emitter is a message emission site.
type Emitter struct {
	Model    string `json:"model"`
	Function string `json:"function"`
}

// CmdDef is a command spawn site.
type CmdDef struct {
	Model      string `json:"model"`
	Line       int    `json:"line"`
	Function   string `json:"function"`
	IsAsync    bool   `json:"is_async"`
	IsBatch    bool   `json:"is_batch"`
	HasCallback bool  `json:"has_callback"`
}

// MessageFlowEdge is a message flow relationship.
type MessageFlowEdge struct {
	From string `json:"from"`
	To   string `json:"to"`
	Via  string `json:"via"`
}

// RenderNode is a node in the render tree.
type RenderNode struct {
	Model       string `json:"model"`
	OutputType  string `json:"output_type"` // "string" | "styled" | "component" | "composite"
	WidthRef    string `json:"width_ref,omitempty"`
	HeightRef   string `json:"height_ref,omitempty"`
	UsesLipgloss bool  `json:"uses_lipgloss"`
	UsesViewport bool  `json:"uses_viewport"`
}

// LayoutNode is a layout composition node.
type LayoutNode struct {
	Type     string       `json:"type"` // "horizontal" | "vertical" | "stack" | "table" | "custom"
	Children []LayoutNode `json:"children,omitempty"`
	Width    int          `json:"width,omitempty"`
	Height   int          `json:"height,omitempty"`
	Margin   []int        `json:"margin,omitempty"`
	Padding  []int        `json:"padding,omitempty"`
	StyleRefs []string   `json:"style_refs,omitempty"`
}

// StateOwnership tracks which model owns which state.
type StateOwnership struct {
	Field      string   `json:"field"`
	Owner      string   `json:"owner"`
	MutationRules []string `json:"mutation_rules,omitempty"`
	SharedWith []string  `json:"shared_with,omitempty"`
}

// MutationPath tracks how state is mutated.
type MutationPath struct {
	Field       string `json:"field"`
	Source      string `json:"source"`
	Trigger     string `json:"trigger"`
	IsDiffBased bool   `json:"is_diff_based"`
	Frequency   string `json:"frequency"`
}

// AsyncBoundary is an async execution boundary.
type AsyncBoundary struct {
	Model          string `json:"model"`
	Line           int    `json:"line"`
	Kind           string `json:"kind"` // "command" | "goroutine" | "channel" | "ticker" | "sequence"
	CreatesGoroutine bool `json:"creates_goroutine"`
	BlocksRender   bool   `json:"blocks_render"`
}

// ComponentBinding is a Charm component bound to a model.
type ComponentBinding struct {
	Model      string   `json:"model"`
	Component  string   `json:"component"`
	Package    string   `json:"package"`
	ConfigFields []string `json:"config_fields,omitempty"`
}
