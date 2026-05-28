# charmed

**A semantic operating model for terminal systems.**

charmed embeds structural knowledge from **23 Charm-ecosystem repositories** and cross-library runtime semantics into a constraint-aware architecture analysis engine. It understands how Bubble Tea, Lip Gloss, Bubbles, Glamour, Huh, Wish, Soft-Serve, and 17 more libraries interact — not as isolated packages, but as an interconnected terminal runtime.

```
Go Source
   ↓
ATree Semantic Graph             cross-file call resolution, type bindings, impact analysis
   ↓
TUI Structural IR                models, messages, commands, boundaries, ownership
   ↓
Constraint Engine                13 rules, 4 severity levels, 6 categories
   ↓
Timing Evaluation                perceptual contracts (smooth streaming, snappy focus, etc.)
   ↓
Architectural Suggestions        decomposition, boundary insertion, caching, autofix
```

This is not an AI coding assistant. It is not a prompt wrapper. `charmed` is architecture intelligence infrastructure for terminal interfaces.

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

## Table of Contents

1. [Semantic Ecosystem Model](#semantic-ecosystem-model)
2. [TUI Structural IR](#tui-structural-ir)
3. [Constraint Engine](#constraint-engine)
4. [Timing Contracts](#timing-contracts)
5. [Archetypes](#archetypes)
6. [Code Graph Reasoning](#code-graph-reasoning)
7. [Skills, Commands, and MCP](#skills-commands-and-mcp)
8. [Installation](#installation)

---

## Semantic Ecosystem Model

charmed does **not** treat the Charm ecosystem as isolated packages. It models the ecosystem as an interconnected runtime architecture — a cross-library semantic graph where component ownership, message flow, rendering behavior, and composition relationships are first-class knowledge.

This is the core asset. Everything else — the constraint engine, the timing evaluator, the refactoring tools — sits on top of this layer.

### 23 Repositories, Normalized into One Ontology

```
bubbletea ─┬─ bubbles ─── viewport
           │              ├── list
           │              ├── textarea
           │              ├── textinput
           │              ├── spinner
           │              ├── help
           │              ├── progress
           │              ├── table
           │              ├── filepicker
           │              └── key
           │
           ├── lipgloss ─── Style (layout, inheritance, render constraints)
           ├── glamour ──── Markdown rendering over Lip Gloss
           ├── glow ─────── Styled markdown viewer
           ├── huh ──────── Form / wizard primitive built on Bubble Tea
           ├── gum ──────── Interactive shell TUI toolkit
           ├── mods ─────── AI-powered terminal pipeline
           ├── wish ─────── SSH-based Bubble Tea host framework
           ├── soft-serve ─ Self-hostable Git server TUI
           └── vhs ──────── Terminal recording and playback
```

Plus: `charm`, `crush`, `fang`, `harmonica`, `log`, `pop`, `sequin`, `ultraviolet`, `wishlist`, `x`, `skate`, and `colorprofile`.

### Cross-Library Semantics

The ontology captures how these libraries **interact architecturally**:

```
viewport.Model
    ↕ scroll ownership, content lifecycle
lipgloss.Style
    ↕ render composition, layout constraints
glamour renderer
    ↕ markdown → styled terminal output
bubbletea.Cmd
    ↕ async orchestration, goroutine lifecycle
tea.Batch / tea.Sequence
    ↕ command ordering, concurrency semantics
bubbles/help.Model
    ↕ key binding discovery, discoverability contract
```

That is the killer feature. Not "we know these libraries exist" — but "we understand how they compose, conflict, and communicate at runtime."

### What's Represented

For every component across all 23 repositories:

- **Ownership** — which model owns which state, where mutations happen
- **Message flow** — emitter → type → handler, traced across files
- **Component relationships** — bounds, pairings, composition patterns
- **Async boundaries** — where concurrency starts, ends, and leaks
- **Render behavior** — what `View()` actually draws, how often it recomputes
- **State contracts** — lifecycle expectations, initialization requirements
- **Failure modes** — per-component known anti-patterns with fix patterns
- **Timing expectations** — perceptual contracts each component must obey
- **Styling semantics** — Lip Gloss style inheritance and render constraints
- **Transport/runtime integration** — SSH (wish), recording (vhs), TUI hosting (soft-serve)

See: [`ontology/knowledge.yaml`](ontology/knowledge.yaml) · [`ontology/ontology-contract.yaml`](ontology/ontology-contract.yaml)

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

The IR can be consumed independently — pipe it into custom analysis, visualization, or CI checks. It is the product surface for anything that needs to reason about TUI architecture programmatically.

Schema: [`ir/tui-structural-ir.schema.json`](ir/tui-structural-ir.schema.json)

---

## Constraint Engine

13 architectural lint rules enforced against the IR. Deterministic: same input, same findings, no stochastic variation.

| Rule | Severity | Category | What It Catches |
|------|----------|----------|-----------------|
| `god_model` | CRITICAL | ARCHITECTURE | 300+ line `Update()`, single struct owning everything |
| `giga_switch_update` | CRITICAL | ARCHITECTURE | 20+ message-type branches in one `Update()` |
| `viewport_recreation` | CRITICAL | PERFORMANCE | `viewport.New()` inside `View()` or `Update()` |
| `no_blocking_update` | ERROR | CORRECTNESS | HTTP/file/DB calls inside `Update()` |
| `viewport_churn` | ERROR | PERFORMANCE | `SetContent()` called multiple times per frame |
| `unbounded_cmd_spawning` | ERROR | CONCURRENCY | Goroutine storms from loop-spawned commands |
| `sync_io_in_init` | ERROR | CORRECTNESS | Blocking I/O delaying first render |
| `viewport_without_resize_handling` | ERROR | ROBUSTNESS | Viewport that never receives `WindowSizeMsg` |
| `style_recomputation_hot_path` | WARN | PERFORMANCE | `lipgloss.NewStyle()` inside `View()` |
| `string_allocation_pressure` | INFO | PERFORMANCE | `fmt.Sprintf` hot path in `View()` |
| `missing_key_bindings_help` | WARN | UX | No `help.Model` or key binding discovery |
| `missing_alt_screen_cleanup` | WARN | ROBUSTNESS | Alt screen entered, never exited |
| `tea_batch_misuse` | INFO | CORRECTNESS | `tea.Batch` where order matters |

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

Each rule includes: what was matched, why it's a problem, and the exact architectural pattern to apply.

See: [`constraints/constraints.yaml`](constraints/constraints.yaml)

---

## Timing Contracts

**SLOs for interaction systems.** Perceptual timing contracts define what "feels alive" in measurable terms.

Most TUI systems think in state and messages. charmed adds the **perceptual dimension** — how the interface *feels* to a human operator. This is the most original part of the project. Almost nobody formalizes perceptual systems engineering for TUIs.

| Contract | Target | Key Constraints |
|----------|--------|----------------|
| `smooth_streaming` | Progressive content display | 30fps max · 50ms cmd warn · 2 viewport mutations/frame |
| `snappy_focus` | Pane switching, navigation | 16ms focus change · 100ms key repeat · 30fps min |
| `responsive_input` | Typing, validation | 16ms echo · 50ms validation · 100ms autocomplete |
| `periodic_refresh` | Dashboard metrics | 500ms interval · 1 max frame drop · 200ms transitions |
| `instant_search` | Live filtering | 50ms debounce · 100ms result latency |
| `instant_validation` | Form feedback | 50ms validation · 0ms error display |
| `60fps_cap` | Render rate ceiling | 60fps max · 16ms frame budget |
| `ambient_refresh_1s` | Background status | 1s interval · 0 frame drops · low visual weight |

Every contract includes **perception traits** (e.g., "progressive disclosure", "no visual delays") and **interaction patterns** (e.g., "each typed character appears in <16ms").

See: [`timing/timing-models.yaml`](timing/timing-models.yaml)

---

## Archetypes

8 behavioral topology definitions — reference architecture templates for common TUI patterns. Each defines: required components, state topology, update flow, command orchestration, timing expectations, known failure modes, and perceptual benchmarks.

| Archetype | Layout | Density | Motion | Timing |
|---|---|---|---|---|
| `agent_console` | multi_pane | balanced | smooth | streaming + responsive input |
| `streaming_chat` | single_pane | balanced | smooth | smooth streaming |
| `ide_workspace` | multi_pane | dense | static | snappy focus |
| `monitoring_dashboard` | multi_pane | dense | static | periodic refresh |
| `cli_wizard` | single_pane | balanced | static | instant validation |
| `git_client` | multi_pane | dense | smooth | streaming + instant search |
| `log_explorer` | single_pane | dense | smooth | streaming + 60fps cap |
| `status_console` | single_pane | ambient | smooth | ambient refresh |

Archetypes are used for: intent classification (`tui-scaffold`), perceptual benchmarking (`tui-ux-review`), and architecture pattern enforcement (`tui-audit`).

Over time, archetypes are evolving toward complete reference architecture specifications — state topology diagrams, command orchestration patterns, and scaling constraint definitions.

See: [`archetypes/archetypes.yaml`](archetypes/archetypes.yaml)

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
# Install ATree (Rust project)
git clone https://github.com/Unity-Lab-AI/ATree.git
cd ATree && cargo build --release
cp target/release/atree ~/.local/bin/atree

# Index your project
export ATREE_DB_PATH=/path/to/your/project/.atree/index.sqlite
```

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

## Skills, Commands, and MCP

Skills, commands, and the MCP server are **interfaces** into the shared engine. The engine underneath is the same no matter which surface you invoke it through.

| Interface | What it does |
|-----------|-------------|
| **8 skills** | Conversational triggers for the engine |
| **6 commands** | Slash-command entry points (`/tui-audit`, `/tui-refactor`, etc.) |
| **charmed-mcp** | Programmatic MCP server for IDE integration |

### Skills

| Skill | Trigger | Primary Knowledge |
|---|---|---|
| **tui-audit** | "audit my TUI", "find anti-patterns", "lint architecture" | Constraints · IR |
| **tui-refactor** | "refactor", "split this model", "too complex", "decompose" | Constraints · Code Graph |
| **tui-explain** | "what is viewport.Model?", "explain timing contract" | Knowledge Corpus · Constraints |
| **tui-scaffold** | "create a TUI", "scaffold a terminal app" | Archetypes · Ontology |
| **tui-instrument** | "add tracing", "profile my TUI", "where is the lag" | Runtime Traces · Timing |
| **tui-ux-review** | "review UX", "how does this feel", "validate timing" | Timing Contracts · Archetypes |
| **evidence-extract** | "extract evidence", "show me proof", "what confirms this pattern" | Evidence Schema · Code Graph |
| **pattern-mine** | "mine patterns", "find recurring structures", "detect empirical patterns" | Patterns · Primitives |

Skills are delivery mechanisms. The ontology, IR, constraints, and timing models are the persistent platform.

---

## The Layered Architecture

```
Interface Layer                  (how you talk to the engine)
  ├── 8 skills                   (conversational triggers)
  ├── 6 slash commands           (/tui-audit, /tui-refactor, …)
  └── MCP server                 (charmed-mcp stdio server)

Semantic Layer                   (what the engine knows)
  ├── Semantic Ecosystem Model    (23 repos → normalized ontology)
  ├── TUI Structural IR          (models, messages, boundaries, ownership)
  ├── Archetype Catalog          (8 reference topology definitions)
  └── Cross-Library Semantics    (component interactions, conflicts, pairings)

Analysis Layer                   (what the engine does)
  ├── Constraint Engine          (13 rules, 4 severity levels, 6 categories)
  ├── Timing Contract Evaluator  (10 perceptual contracts with SLOs)
  ├── Runtime Trace Engine       (frame timing, message timeline, cadence)
  └── Suggestion Engine          (decomposition, boundary insertion, caching)

Evidence Layer                   (how the engine knows)
  └── ATree Semantic Graph       (call paths, type bindings, impact analysis)
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
├── .mcp.json                           # MCP server config (charmed-mcp + atree)
├── charmed-mcp/                        # Go MCP server (AST scanner, knowledge engine)
├── skills/                             # 8 skills (conversational entry points)
├── commands/                           # 6 slash commands
├── agents/                             # charm-project-scanner agent
├── scripts/                            # Binary build script
├── constraints/                        # 13 architectural lint rules + constraints-v2
├── timing/                             # 10 perceptual timing contracts
├── archetypes/                         # 8 behavioral topology definitions
├── ontology/                           # Component/pattern knowledge base
├── patterns/                           # Empirical pattern schema
├── primitives/                         # 32 behavioral primitives vocabulary
├── evidence/                           # Evidence schema
├── ir/                                 # TUI Structural IR schema
├── DeepWiki References/                # Semantic reference corpus (23 libraries)
└── README.md
```

## MCP Server Source

The `charmed-mcp` server lives at [`charmed-mcp/`](charmed-mcp/) inside this repository.

## License

MIT
