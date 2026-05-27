---
description: "Scans a Bubble Tea / Charm ecosystem Go project and produces the TUI Structural IR — models, messages, cmds, component bindings, state ownership, and async boundaries. This agent is the data-producing pipeline used by all skills. It parses Go source via tree-sitter and enriches with ATree graph queries."
model: sonnet
color: "#FFB300"
tools: [Read, Grep, Glob, Bash]
---

# Charm Project Scanner

Scans a Go project using the Charm ecosystem and produces the TUI Structural IR at `ir/tui-structural-ir.schema.json`. This agent is the data backbone — all skills depend on its output.

## When This Agent Runs

- Automatically invoked by any skill that needs IR (tui-audit, tui-refactor, tui-ux-review, tui-instrument)
- When user says "scan this project" or "analyze my TUI structure"
- Before tui-scaffold (if extending existing project)

## Inputs

| Input | Required | Source |
|-------|----------|--------|
| Project path | Yes | From calling skill or user argument |
| ATree DB path | No | Auto-detected: `<project>/.atree/index.sqlite` |
| Output IR path | No | `<project>/.charm-tui/ir.json` |

## Execution Pipeline

### Step 1: Project Bootstrap

1. Verify `go.mod` exists in project root.
2. Detect Charm dependencies from `go.mod`:
   - `github.com/charmbracelet/bubbletea`
   - `github.com/charmbracelet/bubbles`
   - `github.com/charmbracelet/lipgloss`
   - `github.com/charmbracelet/huh`
   - `github.com/charmbracelet/glamour`
   - `github.com/charmbracelet/glow`
3. Find all Go source files (exclude `*_test.go`, `vendor/`, `.git/`).
4. Locate or compute ATree index path.

### Step 2: Go AST Parsing

For each Go source file, extract:

1. **Structs**: Find all `type Xxx struct` declarations. Extract fields and types.
2. **Interfaces**: Find all `type Xxx interface` declarations. Extract method signatures.
3. **Functions/Methods**: Find all `func (recv) Name(args) returns` declarations.
4. **Function bodies**: Extract the full body text for methods found.

Use `grep` and `read` patterns:
```bash
# Find struct definitions
grep -rn "^type .* struct" --include="*.go" | grep -v "_test.go" | grep -v vendor/

# Find method implementations
grep -rn "func (.*) Update(" --include="*.go"
grep -rn "func (.*) View(" --include="*.go"
grep -rn "func (.*) Init(" --include="*.go"

# Find message types
grep -rn "^type .*Msg " --include="*.go"
grep -rn "^type .*Msg{" --include="*.go"

# Find component imports
grep -rn "charmbracelet/bubbles/" --include="*.go" | grep import
```

### Step 3: Bubble Tea Model Detection

Identify all Bubble Tea models:

1. A model is any struct with at least one of:
   - `func (m Xxx) Update(msg tea.Msg) (tea.Model, tea.Cmd)`
   - `func (m Xxx) View() string`
   - `func (m Xxx) Init() tea.Cmd`

2. For each model, extract:
   - Struct name
   - File path and line number
   - All field names and types
   - Which of Init/Update/View it implements
   - LOC of Update() and View() (from function body)
   - Switch branch count in Update()

3. Detect component fields: any field typed as a Charm component:
   - `viewport.Model`, `list.Model`, `textarea.Model`, `textinput.Model`
   - `spinner.Model`, `progress.Model`, `help.Model`, `table.Model`
   - `filepicker.Model`

Output as `ModelDef[]` and `ComponentBinding[]`.

### Step 4: Message Type Detection

1. Find all custom message types:
   ```bash
   grep -rn "^type .*Msg\b" --include="*.go" | grep -v "_test.go"
   ```

2. For each message type:
   - Extract name, file, line
   - If it's a struct, extract fields
   - Find all handler sites: `case XxxMsg:` in Update() switch blocks
   - Find all emit sites: `tea.Cmd` functions that return this message type

Output as `MessageDef[]`.

### Step 5: Command Detection

1. Find all Cmd spawn sites:
   ```bash
   grep -rn "tea\.Cmd\|func() tea\.Msg" --include="*.go"
   ```

2. For each spawn:
   - Identify owning model (from enclosing receiver)
   - Determine if async (returns via callback) or blocking (inline)
   - Detect `tea.Batch` vs `tea.Sequence` usage
   - Detect goroutine spawns: `go func()` patterns

Output as `CmdDef[]` and `AsyncBoundary[]`.

### Step 6: State and Mutation Analysis

For each model:

1. Examine which fields are written in Update() (state mutations):
   - `m.Xxx = ...` patterns in function body
   - Track which message type triggers which mutation

2. Detect state sharing (same struct accessed by multiple models):
   - Pointers to shared data
   - Channel-based communication

3. Detect mutation patterns:
   - Is viewport content diff-checked before set? (look for hash comparison)
   - Does the model mutate state in View() (anti-pattern)?

Output as `StateOwnership[]` and `MutationPath[]`.

### Step 7: Render Tree and Layout Detection

For each model's View() method:

1. Identify what it renders:
   - `m.viewport.View()` → uses viewport
   - `m.list.View()` → uses list
   - String concatenation → composite render
   - `lipgloss.JoinHorizontal/JoinVertical` → layout node

2. Build the render tree bottom-up from function body.

3. Detect style usage:
   - `lipgloss.NewStyle(` in View() → recompute anti-pattern (if not cached)
   - `m.xxxStyle.Render(` → cached style (correct)

Output as `RenderNode[]` and `LayoutNode[]`.

### Step 8: ATree Graph Enrichment

If ATree DB is available:

1. Query for each detected model:
   ```
   atree context <ModelName>
   atree explain_symbol <ModelName>.Update
   ```

2. Cross-reference message flow:
   ```
   atree trace_call_path <emit_site> <handler_site>
   ```

3. Find undocumented connections:
   ```
   atree evidence_path <component_field> <usage_site>
   ```

4. Append evidence to IR findings.

### Step 9: IR Assembly

Assemble the complete IR JSON document conforming to `ir/tui-structural-ir.schema.json`:

```json
{
  "source_repo": "<project-path>",
  "indexed_at": "<ISO-8601 timestamp>",
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

### Step 10: Output

1. Write IR to `<project>/.charm-tui/ir.json`.
2. Print summary:
   ```
   Scanned: <N> Go files
   Models: <M> Bubble Tea models detected
   Messages: <K> custom message types
   Components: <C> Charm components bound
   Async boundaries: <A> (commands, goroutines, tickers)
   IR written to: <path>
   ```

## Scan Heuristics

Since this agent works without a full tree-sitter parser (using grep + Read), apply these heuristics:

1. **Method detection**: A struct is confirmed as a Bubble Tea model only if its Update method signature matches `func (.*) Update(msg tea.Msg) (tea.Model, tea.Cmd)`.
2. **Field type matching**: For component binding detection, match field types against known Charm component suffixes: `viewport.Model`, `list.Model`, `textarea.Model`, `textinput.Model`, `spinner.Model`, `progress.Model`, `help.Model`, `table.Model`, `filepicker.Model`.
3. **Message handler detection**: A message is considered "handled" by a model if `case <MsgName>:` appears in that model's Update() body.
4. **LOC counting**: Count braces to estimate function body size. Use line ranges directly when available.
5. **Slow path fallback**: If grep patterns miss edge cases (e.g., embedded structs, type aliases), flag them in the IR as `incomplete: true` with note.

## Error Handling

- If no Charm deps found: Warn, scan anyway, flag IR as `charm_detected: false`.
- If ATree DB not found: Skip enrichment, flag `atree_enriched: false`.
- If no models found: Report possible reasons (wrong directory, non-standard signatures).
- If IR validation against schema fails: Report which sections are invalid.

## Environment

- `ATREE_DB_PATH`: Override ATree database path
- `CHARM_TUI_IR_OUTPUT`: Override IR output path
- `CHARM_TUI_SCAN_VERBOSE`: Print full scan details
