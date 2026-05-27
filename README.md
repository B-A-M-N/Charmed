# charmed

**Semantic architecture tooling for terminal interfaces.**

Compiler-grade architectural analysis, perceptual timing validation, deterministic constraint enforcement, and code-graph reasoning for the Charm ecosystem — Bubble Tea, Lip Gloss, Bubbles, Glamour, Glow, Gum, Huh, Mods, Soft-Serve, VHS, Wish, and 11+ more libraries.

---

## Why This Exists

Most TUI frameworks give you rendering primitives. Charm gives you the Model-View-Update pattern. But nothing tells you when your architecture is wrong — when a blocking call freezes the event loop, when a viewport resets on every frame, or when a 400-line `Update()` has silently become unmaintainable.

Generic coding assistants can't catch these because they don't understand **TUI-specific failure modes**. They don't know that `SetContent()` called more than once per frame causes visible flicker, or that a goroutine spawned inside a loop creates unbounded concurrency, or that a sync I/O call in `Init()` delays the first render by 300ms.

charmed does. It operates on a **structural intermediate representation** of your TUI — a semantic model extracted from your Go source — and evaluates it against constraint rules and perceptual timing contracts.

This is not a chatbot helper. It is a semantic analysis system for terminal interface architecture.

---

## What charmed Is NOT

```
charmed is NOT:
  - a general-purpose AI coding assistant
  - a replacement for Bubble Tea
  - a runtime framework
  - a UI renderer
  - a prompt wrapper
  - a vibe-coding tool

It IS:
  a semantic architecture analysis system for terminal interfaces.
```

If you want a coding assistant, use Claude. If you want to know whether your TUI architecture will collapse at scale — or whether it already has — use charmed.

---

## What "Semantic" Means

Semantic analysis means charmed understands:

- **ownership** — which model owns which state, where mutations happen
- **message flow** — emitter → type → handler, traced across files
- **component relationships** — which Bubbles components are bound, how they're used
- **async boundaries** — where concurrency starts, where it ends, where it leaks
- **render behavior** — what View() actually draws, how often it recomputes
- **architectural intent** — the structure the code implies, not just the syntax

— not just syntax.

---

## The Pipeline

```
Go Source
   ↓
ATree Semantic Graph            ← cross-file call resolution, type bindings, impact analysis
   ↓
TUI Structural IR               ← models, messages, commands, boundaries, ownership
   ↓
Constraint Engine               ← 13 rules, 4 severity levels, 6 categories
   ↓
Timing Evaluation               ← perceptual contracts (smooth streaming, snappy focus, etc.)
   ↓
Architectural Suggestions       ← decomposition, boundary insertion, caching, autofix
```

**The IR is the product.** Skills, commands, and the MCP server are interfaces into the engine. The semantic representation is the core platform — versioned, stable, and interface-agnostic.

---

## Architecture Intelligence

Deterministic detection of design flaws specific to terminal interface architecture.

- **Model complexity analysis** — LOC budgeting, switch branch counting, god-model detection
- **Component binding verification** — viewport, list, textarea field type matching against known Charm components
- **State ownership mapping** — which model owns which fields, mutation path tracing
- **Async boundary analysis** — goroutine spawn sites, command orchestration, channel detection
- **Render tree extraction** — identifies what each View() actually renders, detects recompute anti-patterns

Severity levels: `INFO` · `WARN` · `ERROR` · `CRITICAL`

Categories: `ARCHITECTURE` · `PERFORMANCE` · `CONCURRENCY` · `STATE` · `ROBUSTNESS` · `UX`

### Example: god model detection

```go
// VIOLATION: 340-line Update, 23 message types, 1 struct owning everything
type AppModel struct {
    viewport viewport.Model
    list     list.Model
    // ... 15 more fields
}
```

```
[CRITICAL] god_model: Update() is 340 lines of switch branches.
           Model owns too much state. Decompose into child models.

Fix: Extract list management into ListModel, viewport into ViewModel.
     Parent orchestrates message passing. Each child owns its own Update/View.
```

---

## Constraint Systems

13 architectural lint rules enforced against the IR with severity, category, and deterministic fix patterns.

| Rule | Severity | Category |
|------|----------|----------|
| `god_model` | CRITICAL | ARCHITECTURE |
| `giga_switch_update` | CRITICAL | ARCHITECTURE |
| `viewport_recreation` | CRITICAL | PERFORMANCE |
| `no_blocking_update` | ERROR | CORRECTNESS |
| `viewport_churn` | ERROR | PERFORMANCE |
| `unbounded_cmd_spawning` | ERROR | CONCURRENCY |
| `sync_io_in_init` | ERROR | CORRECTNESS |
| `viewport_without_resize_handling` | ERROR | ROBUSTNESS |
| `style_recomputation_hot_path` | WARN | PERFORMANCE |
| `string_allocation_pressure` | INFO | PERFORMANCE |
| `missing_key_bindings_help` | WARN | UX |
| `missing_alt_screen_cleanup` | WARN | ROBUSTNESS |
| `tea_batch_misuse` | INFO | CORRECTNESS |

Each rule includes: what was matched, why it's a problem, and the exact pattern to apply.

---

## Timing Contracts

**SLOs for interaction systems.** Perceptual timing contracts define what "feels alive" in measurable terms.

Most TUI systems think in state and messages. charmed adds the **perceptual dimension** — how the interface *feels* to a human operator.

### smooth_streaming
Content streams without jitter or blank frames.

| Constraint | Value |
|---|---|
| Max render frequency | 30 fps |
| Command latency warn | 50ms |
| Command latency error | 150ms |
| Viewport mutations/frame | 2 max |
| String allocations/frame | 3 max |

*Traits: progressive disclosure, incremental append, low jitter, immediate feedback*

### snappy_focus
Focus changes are instant and predictable.

| Constraint | Value |
|---|---|
| Focus change latency | 16ms |
| Key repeat delay | 100ms |
| Min render frequency | 30 fps |

*Traits: instant response, predictable navigation, no visual delays*

### responsive_input
Input feels realtime with no perceptible lag.

| Constraint | Value |
|---|---|
| Echo latency | 16ms |
| Validation latency | 50ms |
| Autocomplete latency | 100ms |

*Traits: character echo, live validation, immediate feedback*

### periodic_refresh
Dashboard metrics tick in without layout disruption.

| Constraint | Value |
|---|---|
| Refresh interval | 500ms |
| Max frame drops | 1 |
| Transition duration | 200ms |

*Traits: stable layout, smooth transitions, no layout jitter*

**Additional contracts:** `periodic_refresh_1s` · `ambient_refresh_1s` · `instant_search` · `instant_validation` · `60fps_cap`

Timing contracts are the most original part of charmed — almost nobody formalizes perceptual systems engineering for TUIs. This is where the project has the deepest differentiated potential.

---

## Perceptual Runtime Analysis

Architecture analysis filtered through human perception — not just "is the code correct" but "does the interface feel right."

- **Focus behavior** — does the interface acknowledge interactions within one frame?
- **Scroll stability** — is cursor position preserved during content append?
- **Pacing** — does the interface feel responsive or does it have perceptible lag?
- **Render cadence** — are frames consistent or does the interface visibly stutter?
- **Density profiling** — ambient, balanced, or dense — matched to archetype expectations
- **Motion philosophy** — static (forms, wizards) vs smooth (streaming, chat)

This is runtime perception engineering, not UX critique.

---

## Runtime Semantics

Understanding what your TUI *does*, not just what it *is*.

- **Message flow graph** — emitter → type → handler, traced across files via ATree
- **Command lifecycle** — spawn sites, ordering (Batch vs Sequence), throttling analysis
- **Side effect detection** — I/O inside Init/Update, network calls in hot paths
- **Concurrency surface** — goroutine leaks, unbounded spawning, channel misuse
- **Render behavior** — viewport churn detection, SetContent frequency analysis, style recompute paths

---

## Runtime Trace Engine

`tui-instrument` produces frame-level traces that can be replayed for perceptual analysis:

```
Frame Budget: 16.6ms (60fps)

Message Timeline:
  KeyMsg          t=0.0ms
  Update()        t=0.1ms  [42 LOC, 3 branches evaluated]
  Cmd Spawn       t=0.5ms  [fetchCmd → async]
  Render          t=1.2ms  [viewport.View() + lipgloss.Render()]
  
  Frame complete: t=2.1ms ✓

Next Frame (47ms later):
  HTTP Result     t=47.2ms
  Update()        t=47.3ms [120 LOC, 8 branches]
  Viewport Set    t=48.1ms [CHURN: SetContent called 3 times]
  Render          t=49.0ms

  ⚠ smooth_streaming breached: viewport mutations/frame = 3 (budget: 2)
```

Trace output feeds back into the constraint engine. Timing violations become architectural findings. Runtime behavior becomes structural evidence.

---

## Code Graph Reasoning

Powered by [ATree](https://github.com/Unity-Lab-AI/ATree) — a semantic code graph engine for Go.

ATree provides the cross-file understanding that makes deep analysis possible:

- **Call path tracing** — which function calls which, across package boundaries
- **Type binding resolution** — fully qualified types with import alias expansion
- **Evidence-based impact analysis** — "if I change this, what breaks?"
- **Heritage detection** — interface implementation, struct embedding, composition chains
- **Assignment tracking** — field mutations across function boundaries

Required for full `scan_project` mode. Not required for constraint lookups, archetypes, timing contract evaluation, or knowledge corpus queries.

```bash
# Install ATree
git clone https://github.com/Unity-Lab-AI/ATree.git
cd ATree && go build -o ~/.local/bin/atree .

# Index your project
export ATREE_DB_PATH=/path/to/your/project/.atree/index.sqlite
```

---

## TUI Structural IR

The intermediate representation at the heart of charmed. Versioned, stable, interface-agnostic.

```json
{
  "source_repo": "/path/to/project",
  "indexed_at": "2026-01-15T10:30:00Z",
  "models": [
    {
      "name": "AppModel",
      "file": "main.go",
      "line": 42,
      "loc_update": 180,
      "loc_view": 45,
      "switch_branches": 12,
      "component_fields": ["viewport.Model", "list.Model"],
      "implements": ["Init", "Update", "View"]
    }
  ],
  "messages": [...],
  "cmds": [...],
  "message_flow_graph": [...],
  "component_bindings": [...],
  "async_boundaries": [...],
  "state_ownership": [...],
  "mutation_paths": [...],
  "render_tree": [...],
  "layout_tree": [...]
}
```

Schema: [`ir/tui-structural-ir.schema.json`](ir/tui-structural-ir.schema.json)

The IR can be consumed independently of the skills — pipe it into custom analysis, visualization, or CI checks.

---

## Archetypes

8 behavioral topology definitions — reference architecture templates for common TUI patterns.

Each archetype defines: required components, state topology, update flow strategy, command orchestration pattern, timing expectations, known failure modes, and perceptual benchmarks.

| Archetype | Layout | Density | Motion | Key Traits |
|---|---|---|---|---|
| `agent_console` | multi_pane | balanced | smooth | streaming output, async orchestration, cursor stability |
| `streaming_chat` | single_pane | balanced | smooth | role-based rendering, incremental append, anchored scroll |
| `ide_workspace` | multi_pane | dense | static | persistent layout, resize propagation, focus management |
| `monitoring_dashboard` | multi_pane | dense | static | periodic refresh, data-driven, alert surfaces |
| `cli_wizard` | single_pane | balanced | static | state machine, step navigation, inline validation |
| `git_client` | multi_pane | dense | smooth | diff rendering, branch navigation, blame overlay |
| `log_explorer` | single_pane | dense | smooth | live tail, regex filtering, severity highlighting |
| `status_console` | single_pane | ambient | smooth | low visual weight, background polling, transient toast |

Archetypes are used for: intent classification (`tui-scaffold`), perceptual benchmarking (`tui-ux-review`), and architecture pattern enforcement (`tui-audit`).

Over time, archetypes are evolving toward complete reference architecture specifications — state topology diagrams, command orchestration patterns, and scaling constraint definitions.

---

## Architecture → Anti-pattern → Corrected

`tui-audit` and `tui-refactor` produce before/after examples for every violation found.

**no_blocking_update — CORRECT: async command boundary inserted:**
```go
// VIOLATION: HTTP call inside Update blocks the event loop
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    resp, _ := http.Get("https://api.example.com")
    return m, nil
}

// FIXED: command boundary
func fetchCmd() tea.Cmd {
    return func() tea.Msg {
        resp, _ := http.Get("https://api.example.com")
        return fetchResult{resp}
    }
}
```

**viewport_recreation — CORRECT: stable viewport, content set once:**
```go
// VIOLATION: viewport recreated every frame, scroll position resets
func (m Model) View() string {
    vp := viewport.New(80, 24)
    vp.SetContent(m.content)
    return vp.View()
}

// FIXED: viewport created in Init, content diffed before Set
func (m Model) Init() (Model, tea.Cmd) {
    m.viewport = viewport.New(80, 24)
    return m, nil
}
```

**god_model — CORRECT: decomposed into child models:**
```go
// VIOLATION: single model owns everything
type AppModel struct { /* 15 fields, 23 message types */ }

// FIXED: decomposed with focused ownership
type AppModel struct {
    list     ListModel
    detail   DetailModel
    // parent orchestrates, children own their domains
}
```

---

## Charm Knowledge Corpus

Embedded reference documentation for the Charm ecosystem — no external API calls needed.

- **Component semantics** — viewport, list, textarea, textinput, spinner, progress, help, table, filepicker, key.Binding, lipgloss.Style
- **Pattern relations** — streaming_console, multi_pane_layout, form_wizard, status_bar_pattern
- **Dependency graph** — bubbletea → bubbles → lipgloss → glamour → glow → gum → huh → mods → soft-serve → wish
- **Anti-patterns** — per-component known failure modes with fix patterns
- **Semantic Reference Corpus** — full DeepWiki documentation for 23 Charm libraries: bubbletea, bubbles, lipgloss, glamour, glow, gum, huh, mods, soft-serve, wish, vhs, and more

---

## Skills

Six skills provide conversational interfaces into the engine.

| Skill | Trigger | Pillar |
|---|---|---|
| **tui-audit** | "audit my TUI", "find anti-patterns", "lint architecture", "detect violations" | Architecture Intelligence · Constraint Systems |
| **tui-refactor** | "refactor", "split this model", "too complex", "decompose" | Architecture Intelligence · Code Graph Reasoning |
| **tui-explain** | "what is viewport.Model?", "what constraints apply to", "explain timing contract" | Constraint Systems · Knowledge Corpus |
| **tui-scaffold** | "create a TUI", "scaffold a terminal app", "I want a streaming log viewer" | Architecture Intelligence · Archetypes |
| **tui-instrument** | "add tracing", "profile my TUI", "where is the lag" | Runtime Semantics · Runtime Trace Engine |
| **tui-ux-review** | "review UX", "how does this feel", "validate timing contracts" | Timing Contracts · Perceptual Runtime |

Skills are interfaces — not the engine. The engine underneath is shared. You get the same constraint system whether you invoke it via `/tui-audit`, ask in natural language, or call the MCP server directly.

---

## The Layered Architecture

```
Interface Layer
  ├── skills         (conversational triggers)
  ├── commands       (/tui-audit, /tui-refactor, …)
  └── MCP server     (programmatic access)

Semantic Layer
  ├── TUI Structural IR       (models, messages, boundaries, ownership)
  ├── Charm Ontology          (components, patterns, dependency graph)
  ├── Archetype Catalog       (8 reference topology definitions)
  └── Semantic Reference Corpus (23-library DeepWiki knowledge base)

Analysis Layer
  ├── Constraint Engine       (13 rules, 4 severity levels, 6 categories)
  ├── Timing Contract Evaluator  (10 perceptual contracts with SLOs)
  ├── Runtime Trace Engine    (frame timing, message timeline, cadence analysis)
  └── Suggestion Engine       (decomposition, boundary insertion, caching)

Evidence Layer
  └── ATree Semantic Graph    (call paths, type bindings, impact analysis)
```

---

## Installation

```bash
# Add the marketplace
claude plugin marketplace add B-A-M-N/SurvivalDevPlugins

# Install
claude plugin install charmed@survivaldev-plugins

# Build the MCP server binary
bash scripts/build-charmed-mcp.sh

# Optional: install ATree for code graph reasoning
git clone https://github.com/Unity-Lab-AI/ATree.git
cd ATree && go build -o ~/.local/bin/atree .
```

## Structure

```
charmed/
├── .claude-plugin/plugin.json          # Plugin manifest
├── .mcp.json                           # MCP server config (charmed-mcp stdio)
├── skills/                             # 6 skills (conversational entry points)
├── commands/                           # 6 slash commands
├── agents/                             # charm-project-scanner agent
├── scripts/                            # Binary build script
├── constraints/                        # 13 architectural lint rules
├── timing/                             # 10 perceptual timing contracts
├── archetypes/                         # 8 behavioral topology definitions
├── ontology/                           # Component/pattern knowledge base
├── ir/                                 # TUI Structural IR schema
├── DeepWiki References/                # Semantic reference corpus (23 libraries)
└── README.md
```

## MCP Server Source

The `charmed-mcp` server is maintained at [`B-A-M-N/charmed-mcp`](https://github.com/B-A-M-N/charmed-mcp).

## License

MIT
