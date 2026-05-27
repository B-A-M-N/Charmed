---
name: tui-scaffold
description: "Intent-aware TUI scaffolding — maps natural language descriptions of terminal UIs to working Bubble Tea architecture. Matches user intent to an archetype, selects components, defines message types and model hierarchy, then generates complete Go source. Trigger when: user says 'create a TUI', 'scaffold', 'build a terminal app', 'generate bubble tea', 'I want a <description> TUI', 'make me a <thing> CLI', or describes a terminal interface goal."
argument-hint: "<intent-description> [--archetype <id>] [--path <output-dir>] [--dry-run]"
allowed-tools: [Read, Write, Bash, Glob, Grep]
---

# tui-scaffold — Intent-to-Architecture Generator

Maps a user's natural language description of a desired TUI to a working Charm architecture, then generates the full Go source. Does NOT just produce a template — it produces a structurally-sound TUI matching the user's specific intent.

## Inputs

| Input | Required | Source |
|-------|----------|--------|
| Intent description | Yes | User argument (natural language) |
| Archetype ID | No | Auto-detected or user override with `--archetype` |
| Constraints | Yes | `constraints/constraints.yaml` |
| Archetypes | Yes | `archetypes/archetypes.yaml` |
| Timing models | Yes | `timing/timing-models.yaml` |
| Existing IR | No | If extending an existing project |
| Output path | No | `--path` (default: current directory) |

## Execution Pipeline

### Phase 1: Intent Classification

Parse the user's description and match to an archetype from `archetypes/archetypes.yaml`:

1. Extract keywords and intent signals from the description.
2. Score each archetype by trait keyword overlap:
   - "streaming output", "agent responses", "LLM" → `agent_console`
   - "chat messages", "chat bubbles", "conversation" → `streaming_chat`
   - "file tree", "editor", "terminal", "multi-pane" → `ide_workspace`
   - "dashboard", "metrics", "gauges", "real-time" → `monitoring_dashboard`
   - "form", "wizard", "step by step", "configure" → `cli_wizard`
   - "git", "diff", "log", "commits" → `git_client`
   - "logs", "tail", "filter", "search logs" → `log_explorer`
   - "status", "background", "ambient" → `status_console`
3. Present the match with confidence score. If confidence < 0.7, ask the user to confirm or pick from top 2-3.
4. Once confirmed, load the full archetype definition: traits, required_components, layout_type, density_profile, motion_philosophy, timing_constraints.

### Phase 2: Architecture Planning

From the archetype, produce an `architecture_plan`:

```
archetype_match:
  id: <archetype_id>
  confidence: <0.0-1.0>
  traits: [...]

models:
  - name: "Model" (or custom name)
    role: root | child | leaf
    implements: [Init, Update, View]
    owns_state: [...]
    children: [...]

message_types:
  - name: "XxxMsg"
    emitted_by: <trigger>
    handled_by: <model>
    payload: <fields>

component_bindings:
  - component: "bubbles/viewport"
    model_field: "viewport"
    configuration: { width: dynamic, height: dynamic }

layout_tree:
  type: <horizontal|vertical|stack|table>
  children: [...]

timing_model: <from archetype>
```

**Key constraints during planning:**
- No model shall exceed 300 LOC in Update() (god_model prevention)
- Viewport must be created once, not per-frame
- Every interactive component must have key bindings
- Async work must be in `tea.Cmd`, never blocking
- WindowSizeMsg must be handled and propagated

### Phase 3: Code Generation

Generate the complete Go source for the project. Structure:

```
<project-name>/
├── main.go              # Entry point, tea.NewProgram()
├── model.go             # Root model struct + Init/Update/View
├── messages.go          # Custom message types
├── components/          # Optional: custom component wrappers
│   └── <component>.go
└── go.mod
```

**main.go** — Minimal entry point:
```go
package main

import (
    "os"
    tea "github.com/charmbracelet/bubbletea"
)

func main() {
    p := tea.NewProgram(NewModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
    if _, err := p.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
```

**model.go** — Root model with all child components composed in struct fields:
- Define the root struct with all component fields
- `Init()` returns `tea.Batch` of child inits + any startup commands
- `Update()`: handle `WindowSizeMsg` first (propagate to children), then switch on message type, delegate to children, forward tea.KeyMsg to focused component
- `View()`: use `lipgloss.JoinVertical`/`JoinHorizontal` per layout_tree

**messages.go** — Custom message types:
```go
type XxxMsg struct {
    // fields
}
type XxxCompleteMsg struct {
    Result <type>
    Err    error
}
```

**Per child model** (if multi-model):
- Own struct with focused state
- Own `Init/Update/View`
- Parent stores as field, delegates in switch

### Phase 4: Constraint Validation

After generating code, run a quick scan against the constraint rules:
- Verify no blocking patterns in Update()
- Verify viewport created in struct, not in Update()
- Verify all key bindings present
- Verify WindowSizeMsg handler exists

If any violations are found in the generated code, fix them before presenting.

### Phase 5: Output

**Dry run (default):** Show the architecture plan and generated files in the chat. Ask user to confirm before writing.

**With confirmation or `--apply`:** Write files to disk. Show file tree summary.

```
Scaffolded <project-name>/
├── main.go          (entry point)
├── model.go          (root model: <name>)
├── messages.go       (<N> message types)
└── go.mod

Archetype: <id> (confidence: <score>)
Layout: <layout-type>
Timing: <timing-model>

Run: cd <project-name> && go mod tidy && go run .
```

## Component Selection Guide

Based on archetype traits, select from:

| Need | Component | Import |
|------|-----------|--------|
| Scrollable content | `bubbles/viewport` | `github.com/charmbracelet/bubbles/viewport` |
| Item list | `bubbles/list` | `github.com/charmbracelet/bubbles/list` |
| Multi-line input | `bubbles/textarea` | `github.com/charmbracelet/bubbles/textarea` |
| Single-line input | `bubbles/textinput` | `github.com/charmbracelet/bubbles/textinput` |
| Loading indicator | `bubbles/spinner` | `github.com/charmbracelet/bubbles/spinner` |
| Progress bar | `bubbles/progress` | `github.com/charmbracelet/bubbles/progress` |
| Key help | `bubbles/help` | `github.com/charmbracelet/bubbles/help` |
| Table | `bubbles/table` | `github.com/charmbracelet/bubbles/table` |
| File picker | `bubbles/filepicker` | `github.com/charmbracelet/bubbles/filepicker` |
| Confirm dialog | `bubbles/confirm` or `huh` | `github.com/charmbracelet/huh` |

## Layout Archetypes

From `layout_type` in archetype:

- **single_pane**: One viewport or component fills the screen + optional status bar at bottom
- **multi_pane**: Split into 2-3 regions using `lipgloss.JoinHorizontal`/`JoinVertical` with fixed or dynamic ratios
- **tabbed**: Multiple pages, switch via key binding, only active page receives updates
- **overlay**: Base content + modal layer, modal captures all input

## Anti-Patterns to Prevent in Generated Code

- Never call `viewport.New()` in Update() or View() — create in struct or Init()
- Never do I/O in Update() — wrap in `func() tea.Cmd { ... }`
- Never create styles in View() — cache on model struct
- Never hardcode dimensions — derive from `tea.WindowSizeMsg`
- Never forget to forward messages to child models
- Never forget alt screen cleanup on quit

## Environment

- `CHARM_TUI_SCAFFOLD_PATH`: Default output directory
- `CHARM_TUI_AUTHOR`: Author name for generated go.mod
