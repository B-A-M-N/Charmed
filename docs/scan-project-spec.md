# `scan_project` — Full Pipeline Spec

## 1. Purpose

Scan a Go project using the Charm ecosystem and produce the TUI Structural IR.
This is the data backbone for all downstream tools (audit, refactor, ux_review, etc.).

## 2. Tool Signature

```
scan_project(
  project_path: string        // absolute or relative path to Go project root
  include_vendor?: boolean    // default: false
  analyze_tests?: boolean     // default: false
  atree_db_path?: string      // override auto-detected ATree DB path
  depth?: "quick" | "full"   // default: "full"
) → ScanResult
```

### Depth Modes

| Mode | Stages | Use Case |
|---|---|---|
| `"quick"` | 1-4 (Bootstrap → Model Detection) | Fast model discovery, no enrichment |
| `"full"` | All 13 stages | Complete IR with all analysis |

## 3. Output: `ScanResult`

```json
{
  "ir": { /* TUI Structural IR — conforms to tui-structural-ir.schema.json */ },
  "scan_meta": {
    "charmed_version": "0.1.0",
    "ir_version": "1",
    "ontology_version": "1",
    "constraint_version": "1",
    "timing_model_version": "1",
    "scanned_at": "ISO-8601",
    "elapsed_ms": 1234,
    "files_scanned": 42,
    "go_files": 38,
    "charm_detected": true,
    "atree": {
      "available": true,
      "db_path": "/project/.atree/index.sqlite",
      "enrichments": {
        "symbol_resolution": true,
        "call_graph": true,
        "import_resolution": true,
        "method_signatures": true,
        "switch_case_bodies": true,
        "type_bindings": true,
        "assignment_tracking": true
      },
      "unresolved_symbols": 0,
      "confidence": 0.95
    }
  },
  "warnings": [
    {
      "type": "string",
      "message": "string",
      "confidence": 0.0
    }
  ]
}
```

## 4. Pipeline Stages

```
┌──────────────────────────────────────────────────────────────────────┐
│                        scan_project                                  │
│                                                                      │
│  1. Bootstrap           8. State & Mutation Analysis                 │
│  2. ATree Index           9. View Analysis                           │
│  3. Graph Normalization  10. Constraint Precomputation                │
│  4. Model Detection      11. IR Assembly                            │
│  5. Message Flow         12. Validation                              │
│  6. Component Binding    13. Output                                  │
│  7. Async Boundaries                                                  │
│                                                                      │
└──────────────────────────────────────────────────────────────────────┘
```

All stages after #3 consume ATree MCP responses exclusively.
No raw ATree formats leak past the normalization layer.

---

### Stage 1: Bootstrap

**Goal**: Verify project, detect Charm deps, locate ATree DB.

```
actions:
  - Verify go.mod exists at project_path
  - Parse go.mod for charmbracelet/* dependencies:
    - github.com/charmbracelet/bubbletea
    - github.com/charmbracelet/bubbles
    - github.com/charmbracelet/lipgloss
    - github.com/charmbracelet/huh
    - github.com/charmbracelet/glamour
    - github.com/charmbracelet/glow
    - github.com/charmbracelet/wish
  - Locate ATree DB (priority order):
    1. atree_db_path parameter
    2. <project_path>/.atree/index.sqlite
    3. $ATREE_DB_PATH env var
    4. Unavailable → atree_available=false, continue (stages 2,5-10 skip enrichment)
  - List Go source files (*.go minus *_test.go, minus vendor/)

outputs:
  project_root:     string
  go_mod_found:     bool
  charm_deps:       []string
  go_files:         []string
  atree_available:  bool
  atree_db_path:    string | null

warnings:
  - no_go_mod:       go.mod not found (continue)
  - no_charm_deps:   no charmbracelet imports (continue, charm_detected=false)
  - atree_unavailable: no ATree DB found (continue without enrichment)
```

### Stage 2: ATree Index

**Goal**: Ensure ATree DB is populated and up-to-date.

```
if atree_available:
  ATree MCP: index(path: project_path, force: false)

  If DB empty or stale:
    ATree MCP: index(path: project_path, force: true)

  Verify: ATree MCP: resolution_stats() → check unresolved rate < 5%
```

### Stage 3: Graph Normalization

**Goal**: Convert ATree MCP responses into a clean internal representation. This is the hard boundary — no ATree formats leak past this stage.

```go
type ProjectGraph struct {
    Files       map[string]*FileInfo
    Symbols     map[string]*Symbol
    Structs     []*StructInfo
    Methods     map[string][]*MethodInfo  // receiver_type → methods
    Interfaces  []*InterfaceInfo
    Calls       []*CallEdge
    Imports     map[string]*ImportMap     // file → local_name → full_path
    Heritage    []*HeritageEdge
    Assignments []*Assignment
    Scopes      []*ScopeInfo
}

type FileInfo struct {
    Path        string
    Language    string
    HasCharmDep bool
}

type Symbol struct {
    Name       string
    Qualified  string
    Tag        string   // "struct" | "method" | "function" | "interface" | "const" | "var" | "property"
    File       string
    Line       int
    Col        int
    IsExported bool
}

type StructInfo struct {
    Symbol
    Fields     []*FieldInfo
    Methods    []*MethodInfo
    Embeds     []string       // embedded type qualified names
    File       string
    Line       int
}

type FieldInfo struct {
    Name      string
    Type      string          // resolved type (import aliases expanded)
    StructTag string
    Line      int
    IsEmbed   bool
}

type MethodInfo struct {
    Symbol
    Receiver      string      // qualified receiver type name
    ReceiverIsPtr bool
    Params        []ParamInfo // name + type
    Returns       []string    // return types
    File          string
    Line          int
    BodyStart     int         // scope line_start
    BodyEnd       int         // scope line_end
    IsInterface   bool        // satisfies tea.Model interface?
}

type ParamInfo struct {
    Name string
    Type string
}

type InterfaceInfo struct {
    Symbol
    Methods    []*MethodInfo
    File       string
    Line       int
}

type CallEdge struct {
    Caller     string
    Callee     string
    Receiver   string    // method receiver if method call
    File       string
    Line       int
    Col        int
    Confidence float64
    Resolved   bool
}

type ImportMap struct {
    LocalName  string
    FullPath   string
    IsAlias    bool
}

type HeritageEdge struct {
    Child      string
    Parent     string
    Kind       string   // "extends" | "implements" | "embeds"
    File       string
    Line       int
}

type Assignment struct {
    Field      string
    Source     string    // what's assigned
    InFunction string    // which method/function
    File       string
    Line       int
}

type ScopeInfo struct {
    ID          int
    ParentID    int
    OwnerSym    string   // qualified symbol name
    Kind        string   // "method" | "function" | "struct" | "block"
    File        string
    LineStart   int
    LineEnd     int
}
```

**Normalization rules**:
1. Filter to Go files only
2. Expand import aliases: `vp viewport.Model` → field type `github.com/charmbracelet/bubbles/viewport.Model`
3. Build `MethodOwner` map: join `DefinitionMethod` symbols with `scopes.owner_symbol_id` → receiver type
4. Track unresolved symbols separately (for warnings)
5. Build full qualified names: `<module>/<package>.<Type>.<Method>`

**ATree MCP tools used**:
- `index` — ensure project is indexed
- `query` — bulk symbol retrieval

### Stage 4: Model Detection

**Goal**: Identify all Bubble Tea model implementations.

```
detection rules:
  A struct is a Bubble Tea model if ANY of:
    1. Direct: has Update + View + Init methods matching tea.Model signatures
    2. Partial: has at least one of Update/View/Init
    3. Via embedding: embeds a type that provides missing methods

  Signature matching (from ATree method signatures):
    Update: func (r Xxx) Update(msg tea.Msg) (tea.Model, tea.Cmd)
    View:   func (r Xxx) View() string
    Init:   func (r Xxx) Init() tea.Cmd

  For each detected model:
    - Classify: root_model | child_model | standalone
      root: referenced in tea.NewProgram() call
      child: field type on another model struct
      standalone: composed but not obviously parented

  Walk embedding chain:
    For embedded structs, check if they provide missing tea.Model methods
    Use HeritageEdge with Kind="embeds" + recursive struct inspection

  Output: []ModelDef with signals emitted:
    - model_detected
    - model_root / model_child / model_standalone
    - model_update_loc, model_view_loc, model_init_loc
    - model_update_branches (switch case count)
    - model_has_embedded_base
```

**ATree MCP tools used**:
- `find_entrypoints` — find tea.NewProgram calls
- `context` — get call relationships for classification

### Stage 5: Message Flow

**Goal**: Map custom message types to their handlers and emitters.

```
detection rules:
  A custom message type is:
    - Any type named *Msg (struct, interface, type alias)
    - Any type used in a case branch inside Update()'s type switch on tea.Msg

  Handler detection (from ATree switch_case_bodies):
    For each model's Update() method:
      Find: switch msg.(type)  or  switch msg := msg.(type)
      Extract: each case branch's type → handled message type
      Record: model, message_type, line, file

  Emitter detection (from ATree call graph + type_bindings):
    Find: func() tea.Cmd closures that construct *Msg types
    Find: return XxxMsg{...} statements
    Find: tea.Cmd functions returning message types
    Record: model (enclosing method's receiver), message_type, line, file

  Build message flow edges:
    emitter_model → message_type → handler_model

  Output: []MessageDef + message_flow_graph
  Signals emitted:
    - message_handled
    - message_emitted
    - message_unhandled (has emitters but no handlers)
    - message_orphan (has handlers but no emitters)
```

**ATree MCP tools used**:
- `query` — find type definitions matching *Msg pattern
- `context` — get method body semantics (switch cases, return sites)
- Semantic: switch_case_bodies, type_bindings

### Stage 6: Component Binding

**Goal**: Detect Charm components and their wiring.

```
known component types (from embedded ontology):
  bubbles/viewport.Model, bubbles/list.Model, bubbles/textarea.Model,
  bubbles/textinput.Model, bubbles/spinner.Model, bubbles/progress.Model,
  bubbles/help.Model, bubbles/table.Model, bubbles/filepicker.Model,
  key.Binding, lipgloss.Style

detection (from ATree field types + import resolution):
  For each model struct field:
    1. Resolve field type via ImportMap (expand aliases)
    2. Match against known component types
    3. Record: model, component, field_name, file, line, is_embedded

  Wiring analysis (from ATree call graph):
    For each component binding:
      - Creation site: struct literal | Init() | Update() | View()
      - SetContent() calls → viewport churn candidates
      - Resize handling: WindowSizeMsg propagation
      - Focus/Blur lifecycle for inputs

  Output: []ComponentBinding
  Signals emitted:
    - component_bound
    - component_created_in_update (viewport_recreation candidate)
    - component_setcontent_call (viewport_churn candidate)
    - component_missing_resize (viewport_without_resize_handling candidate)
```

**ATree MCP tools used**:
- `context` — method body call sites
- Semantic: assignments, call sites within function bodies

### Stage 7: Async Boundaries

**Goal**: Map all async execution boundaries.

```
detection (from ATree call graph):
  Async boundary types:
    command:    tea.Cmd creation (func() tea.Msg)
    batch:      tea.Batch(...)
    sequence:   tea.Sequence(...)
    goroutine:  go func() { ... }()
    ticker:     time.Tick, time.NewTicker
    channel:    <-ch, ch <-

  For each boundary:
    - Owning model (from enclosing method's receiver via scopes)
    - Classification (from ATree call graph)
    - Goroutine creation flag
    - Blocks render flag (synchronous call in Update)

  Output: []AsyncBoundary
  Signals emitted:
    - async_boundary_detected
    - async_goroutine
    - async_in_loop (unbounded_cmd_spawning candidate)
    - async_blocking_update (no_blocking_update candidate)
```

**ATree MCP tools used**:
- `concurrency_surface_detector` — detect goroutines, channels, locks
- `context` — method-level call relationships

### Stage 8: State & Mutation Analysis

**Goal**: Map state ownership and mutation paths.

```
detection (from ATree assignments + call graph):
  For each model:
    1. Field mutations in Update():
       - m.Xxx = ... patterns (from assignments)
       - Which message type triggers which mutation (from switch case → assignment mapping)

    2. State sharing:
       - Pointers to shared data (from type_bindings)
       - Channel-based communication (from concurrency_surface)

    3. Anti-pattern detection:
       - State mutation in View() (assignments in View scope)
       - Viewport content diff-check before SetContent

  Output: []StateOwnership + []MutationPath
  Signals emitted:
    - state_field_mutated
    - state_shared_pointer
    - state_mutation_in_view
    - viewport_no_diffcheck (viewport_churn candidate)
```

**ATree MCP tools used**:
- Semantic: assignments, scope information
- `side_effect_scanner` — detect side effects in methods

### Stage 9: View Analysis

**Goal**: Extract render tree and layout structure.

```
detection (from ATree call graph within View scopes):
  For each model's View() method:
    1. Rendered components (from method calls in View scope):
       - m.viewport.View() → viewport
       - m.list.View() → list
       - m.textarea.View() → textarea
       - lipgloss.JoinHorizontal/JoinVertical → layout node
       - lipgloss.NewStyle() → style creation
       - fmt.Sprintf() → string formatting

    2. Render tree (bottom-up):
       - Leaves: component .View() calls
       - Internal: lipgloss joins, string ops
       - Root: return value

    3. Style analysis:
       - Cached: m.xxxStyle.Render() → correct
       - Recomputed: lipgloss.NewStyle() in View → hot path

  Output: []RenderNode + []LayoutNode
  Signals emitted:
    - view_renders_component
    - view_style_recomputed (style_recomputation_hot_path candidate)
    - view_sprintf_count (string_allocation_pressure candidate)
```

**ATree MCP tools used**:
- Semantic: method call sites within View() scope
- `context` — method body semantics

### Stage 10: Constraint Precomputation

**Goal**: Evaluate constraint rules against accumulated signals. NOT a re-scan — purely signal evaluation.

```
For each constraint rule, evaluate against signals from stages 4-9:

  no_blocking_update:     async_blocking_update signal
  god_model:              model_update_loc > 300
  giga_switch_update:     model_update_branches > 20
  viewport_recreation:    component_created_in_update signal
  viewport_churn:         viewport_no_diffcheck + component_setcontent_call count
  style_recomputation:    view_style_recomputed signal
  unbounded_cmd_spawning: async_in_loop signal
  sync_io_in_init:        blocking call in Init scope (from side_effect_scanner)
  missing_key_bindings:   no key.Binding import detected
  viewport_no_resize:     component_missing_resize signal
  string_alloc_pressure:  view_sprintf_count > 5
  tea_batch_misuse:       async_batch in state-dependent context
  missing_alt_cleanup:    EnterAltScreen without ExitAltScreen

  Output: []ConstraintSignal
```

### Stage 11: IR Assembly

**Goal**: Assemble the complete TUI Structural IR document.

```
Assemble JSON conforming to tui-structural-ir.schema.json:

{
  "source_repo": "<project_path>",
  "indexed_at": "<ISO-8601>",
  "atree_db_path": "<path or null>",
  "models": [...],
  "messages": [...],
  "cmds": [...],
  "message_flow_graph": [...],
  "render_tree": [...],
  "layout_tree": [...],
  "state_ownership": [...],
  "mutation_paths": [...],
  "async_boundaries": [...],
  "component_bindings": [...]
}
```

### Stage 12: Validation

**Goal**: Validate IR completeness and consistency.

```
checks:
  - IR validates against tui-structural-ir.schema.json
  - Empty models array → warning "no_models_detected"
  - Models without Update → warning "incomplete_model"
  - Messages with no handlers → warning "unhandled_messages"
  - Components without resize handler → info "missing_resize"
  - Constraint signals with no source confidence → warning "low_confidence_signal"
  - Atree enrichment gaps → warning "partial_enrichment"

output:
  Return ScanResult JSON with ir + scan_meta + warnings
```

## 5. Signal Architecture

Stages 4-9 emit **signals** (raw observations). Stage 10 evaluates signals against constraint rules. This eliminates duplicate scanning.

```go
type Signal struct {
    Kind      string            // e.g., "model_detected", "viewport_new", "blocking_io"
    File      string
    Line      int
    Model     string            // owning model, if applicable
    Context   string            // "Update()", "View()", "Init()", "loop", etc.
    Confidence float64          // 0-1, from ATree confidence scores
    Metadata  map[string]string // additional context
}
```

Signal → Constraint mapping:

| Constraint | Required Signals |
|---|---|
| no_blocking_update | `async_blocking_update` in Update scope |
| god_model | `model_update_loc` > 300 |
| giga_switch_update | `model_update_branches` > 20 |
| viewport_recreation | `component_created_in_update` + component is viewport |
| viewport_churn | `viewport_no_diffcheck` + `component_setcontent_call` |
| style_recomputation | `view_style_recomputed` |
| unbounded_cmd_spawning | `async_in_loop` |
| sync_io_in_init | `async_blocking_update` in Init scope |
| missing_key_bindings | absence of key.Binding import |
| viewport_no_resize | `component_missing_resize` |
| string_alloc_pressure | `view_sprintf_count` > 5 |
| tea_batch_misuse | `async_batch` in state-dependent context |
| missing_alt_cleanup | `altscreen_enter` without `altscreen_exit` |

## 6. ATree Integration

### 6.1 ATree MCP Tools Used

| ATree Tool | Stages | Purpose |
|---|---|---|
| `index` | 2 | Ensure project is semantically indexed |
| `query` | 3, 5 | Symbol search, bulk retrieval |
| `context` | 4, 5, 6, 7, 9 | Symbol-level relationships |
| `find_entrypoints` | 4 | Find tea.NewProgram calls |
| `concurrency_surface_detector` | 7 | Goroutines, channels, locks |
| `side_effect_scanner` | 8 | I/O side effects in functions |

### 6.2 ATree Response Format (post-upgrade)

ATree provides structured semantic data per file:
```json
{
  "path": "cmd/app.go",
  "language": "Go",
  "defs": [
    {
      "name": "Update",
      "tag": "DefinitionMethod",
      "line": 42,
      "receiver": "AppModel",
      "params": [{"name": "msg", "type": "tea.Msg"}],
      "returns": ["tea.Model", "tea.Cmd"]
    }
  ],
  "switch_case_bodies": [
    {
      "scope": "AppModel.Update",
      "cases": [
        {"type": "tea.KeyMsg", "line": 44, "body_start": 45, "body_end": 60},
        {"type": "FetchMsg", "line": 61, "body_start": 62, "body_end": 80}
      ]
    }
  ],
  "assignments": [
    {"field": "m.items", "source": "msg.Items", "in": "AppModel.Update", "line": 70}
  ],
  "type_bindings": [
    {"var": "cmd", "type": "tea.Cmd", "line": 85}
  ],
  "calls": [
    {"name": "m.viewport.View()", "scope": "AppModel.View", "line": 92},
    {"name": "http.Get", "scope": "AppModel.Update", "line": 55}
  ],
  "imports": [
    {"local": "tea", "path": "github.com/charmbracelet/bubbletea"},
    {"local": "viewport", "path": "github.com/charmbracelet/bubbles/viewport"}
  ]
}
```

### 6.3 ATree MCP Client Mode

charmed-mcp connects to ATree MCP as a **client**. ATree runs as a separate server:

```
ATree MCP server ←── MCP client ──→ charmed-mcp ←── MCP client ──→ host (Zed, Cursor, etc.)
```

If ATree MCP is not available, charmed-mcp degrades gracefully:
- Skip Stages 2, 5-10 enrichment
- Mark `atree.available: false` in scan_meta
- Constraint precomputation runs on whatever signals were available from Stage 1-4

## 7. Versioning

All version fields pinned to prevent drift:

```yaml
charmed_version: "0.1.0"
ir_version: "1"
ontology_version: "1"
constraint_version: "1"
timing_model_version: "1"
```

Bump a knowledge base's version when its content changes. Clients detect mismatches.

## 8. Error Model

| Condition | Behavior |
|---|---|
| No go.mod | Warning `no_go_mod`. Continue. |
| No charm deps | Warning `no_charm_deps`. `charm_detected: false`. Continue. |
| No ATree DB | Warning `atree_unavailable`. Skip enrichment. Continue. |
| ATree index empty | Auto-trigger `index` tool. Retry. |
| No models found | Warning `no_models_detected`. Return empty IR. |
| Unresolved symbols | Include in `warnings[]` with confidence score. |
| Partial parse failure | Include failed files in `warnings[]`. IR has `incomplete: true`. |
| IR schema validation fail | Return error with details about failed sections. |

## 9. Validation Targets

| Repo | Expected | Notes |
|---|---|---|
| bubbletea/examples | 10+ models, all components | Canonical patterns |
| charmbracelet/glow | 1-2 models, viewport + list | Markdown viewer |
| charmbracelet/gum | Multiple subcommand models | Complex composition |
| charmbracelet/soft-serve | 3-5 models, dense components | Message flow testing |
| Random non-charm Go project | 0 models, no false positives | Negative test |
| Project with broken syntax | Partial IR with warnings | Robustness test |

### Success Criteria

1. All Bubble Tea models detected (no false negatives on validation targets)
2. All custom message types mapped to handlers + emitters
3. All Charm component bindings resolved with correct import paths
4. No false positive model detections
5. Constraint signals correctly precomputed from signals
6. IR validates against schema
