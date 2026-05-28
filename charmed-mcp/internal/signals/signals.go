// Package signals defines the signal types emitted by semantic passes
// and evaluated by the constraint precomputation pass.
package signals

// Kind is the type of signal emitted by analysis passes.
type Kind string

const (
	// Model signals
	ModelDetected       Kind = "model_detected"
	ModelRoot           Kind = "model_root"
	ModelChild          Kind = "model_child"
	ModelStandalone     Kind = "model_standalone"
	ModelUpdateLoc      Kind = "model_update_loc"
	ModelViewLoc        Kind = "model_view_loc"
	ModelInitLoc        Kind = "model_init_loc"
	ModelUpdateBranches Kind = "model_update_branches"

	// Message signals
	MessageHandled   Kind = "message_handled"
	MessageEmitted   Kind = "message_emitted"
	MessageUnhandled Kind = "message_unhandled"
	MessageOrphan    Kind = "message_orphan"

	// Component signals
	ComponentBound              Kind = "component_bound"
	ComponentCreatedInUpdate    Kind = "component_created_in_update"
	ComponentSetContentCall     Kind = "component_setcontent_call"
	ComponentMissingResize      Kind = "component_missing_resize"

	// Async signals
	AsyncBoundary      Kind = "async_boundary_detected"
	AsyncGoroutine     Kind = "async_goroutine"
	AsyncInLoop        Kind = "async_in_loop"
	AsyncBlockingUpdate Kind = "async_blocking_update"
	AsyncBatch         Kind = "async_batch"

	// State signals
	StateFieldMutated   Kind = "state_field_mutated"
	StateSharedPointer  Kind = "state_shared_pointer"
	StateMutationInView Kind = "state_mutation_in_view"
	ViewportNoDiffcheck Kind = "viewport_no_diffcheck"

	// View signals
	ViewRendersComponent  Kind = "view_renders_component"
	ViewStyleRecomputed   Kind = "view_style_recomputed"
	ViewSprintfCount      Kind = "view_sprintf_count"

	// Lifecycle signals
	AltScreenEnter Kind = "altscreen_enter"
	AltScreenExit  Kind = "altscreen_exit"
)

// Signal is a raw observation emitted by a semantic pass.
type Signal struct {
	Kind       Kind              `json:"kind"`
	File       string            `json:"file"`
	Line       int               `json:"line"`
	Model      string            `json:"model,omitempty"`
	Context    string            `json:"context,omitempty"` // "Update()" | "View()" | "Init()" | "loop" etc.
	Confidence float64           `json:"confidence"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

// Collector accumulates signals during analysis.
type Collector struct {
	signals []Signal
}

// NewCollector creates a new signal collector.
func NewCollector() *Collector {
	return &Collector{}
}

// Emit adds a signal.
func (c *Collector) Emit(s Signal) {
	c.signals = append(c.signals, s)
}

// Signals returns all collected signals.
func (c *Collector) Signals() []Signal {
	return c.signals
}

// ByKind filters signals by kind.
func (c *Collector) ByKind(kind Kind) []Signal {
	var result []Signal
	for _, s := range c.signals {
		if s.Kind == kind {
			result = append(result, s)
		}
	}
	return result
}

// ByModel filters signals by owning model.
func (c *Collector) ByModel(model string) []Signal {
	var result []Signal
	for _, s := range c.signals {
		if s.Model == model {
			result = append(result, s)
		}
	}
	return result
}

// CountByKind counts signals of a given kind.
func (c *Collector) CountByKind(kind Kind) int {
	count := 0
	for _, s := range c.signals {
		if s.Kind == kind {
			count++
		}
	}
	return count
}
