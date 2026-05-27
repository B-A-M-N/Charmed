---
name: tui-refactor
description: "Architectural refactoring for Bubble Tea projects — detects god models, viewport churn, blocking patterns, and unstable rendering, then produces safe, minimal patches with ATree impact analysis. Trigger when: user says 'refactor', 'split this model', 'fix architecture', 'too complex', 'clean up', 'decompose', 'extract component', 'optimize TUI', or when tui-audit findings warrant structural changes."
argument-hint: "[path-to-project] [--target <symbol>] [--dry-run] [--apply]"
allowed-tools: [Read, Write, Edit, Bash, Grep, Glob, Agent]
---

# tui-refactor — Architectural Refactoring Engine

Detects architectural violations in Bubble Tea projects and produces safe, minimal patches. All changes go through ATree impact analysis before application. Dry-run by default.

## Inputs

| Input | Required | Source |
|-------|----------|--------|
| Project path | Yes | User argument (default: current directory) |
| Constraints | Yes | `constraints/constraints.yaml` |
| ATree DB path | Yes | Auto-detected or `ATREE_DB_PATH` |
| Target symbol | No | `--target ModelName` (refactor specific model) |
| Dry run | Yes | Default on, `--apply` to write |

## Execution Pipeline

### Phase 1: Violation Detection

Run analysis equivalent to tui-audit, producing a violations list. Prioritize by:

1. FATAL first (viewport_recreation, giga_switch_update)
2. ERROR second (god_model, viewport_churn, blocking_update, unbounded_cmds)
3. WARNING third (missing_help, style_recomputation)

### Phase 2: Refactor Classification

For each violation, classify the required action:

| Violation | Action |
|---|---|
| god_model (>300 LOC Update) | `split_model` — extract child models |
| giga_switch_update (>20 branches) | `split_model` — route messages to children |
| viewport_recreation | `stabilize_viewport` — move to Init(), add diff check |
| viewport_churn | `stabilize_viewport` — add content diff before SetContent |
| blocking_update | `extract_component` — wrap in tea.Cmd |
| unbounded_cmd_spawning | `throttle_cmd` — add tick/debounce |
| viewport_without_resize_handling | `add_resize_handler` — propagate WindowSizeMsg |
| style_recomputation_hot_path | `cache_styles` — move to struct or computed field |
| missing_key_bindings_help | `add_help` — import and wire bubbles/help |

### Phase 3: ATree Impact Analysis

For each planned action, query ATree:

1. **`atree impact <target_symbol>`** — What depends on the symbol being refactored?
2. **`atree boundary_check <file>:<func>`** — What are the function's side effects?
3. **`atree scope_violations <symbol>`** — Will the refactoring break scope boundaries?
4. **`atree dep_cycles <package>`** — Any circular dependencies introduced?

If impact score > 10 affected symbols, warn user before proceeding.

### Phase 4: Patch Generation

For each action, produce a patch:

```
action: split_model
target_file: cmd/app.go
current_state: "single 350-line Update handling 14 message types"
target_state: "root Model + ChatModel + InputModel + StatusModel"
affected_by: atree impact — 3 methods, 2 interfaces
```

For each patch, show the diff:

```diff
- func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
-     switch msg := msg.(type) {
-     case ChatMsg: ...
-     case InputMsg: ...
-     ... (14 more cases)
+ func (m *Model) Update(msg tea.Cmd) (tea.Model, tea.Cmd) {
+     switch msg := msg.(type) {
+     case tea.WindowSizeMsg:
+         m.resize(msg)
+     case tea.KeyMsg:
+         return m, m.input.Update(msg)
+     default:
+         var cmd tea.Cmd
+         m.chat, cmd = m.chat.Update(msg)
+         return m, cmd
+     }
```

Generated patches always:
- Preserve the existing external interface (struct name, exported methods)
- Add new files for extracted child models
- Update imports
- Include the constraint that fixes the original violation
- Add ATree verification comments

### Phase 5: Output

**Dry run (default):** Show all planned actions with patches in the chat. No files modified.

**With `--apply`:** Write patches to disk sequentially. After each file write, verify the code still compiles (`go build ./...`). If compilation fails, roll back that file and report the error.

```
Refactor plan for <project>:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  [1/3] split_model: Extract ChatModel from root
          Impact: 3 symbols affected, 0 dep cycles
          Files: +components/chat.go, ~model.go

  [2/3] stabilize_viewport: Move viewport.New to Init
          Impact: 1 symbol, no side effects
          Files: ~model.go

  [3/3] add_resize_handler: Propagate WindowSizeMsg
          Impact: 4 symbols (model, viewport, list, status)
          Files: ~model.go

Apply? [y/N] (or use --apply flag)
```

## Splitting Strategies

### Extract Child Model

When `god_model` > 300 LOC:
1. Group related fields on the struct by message handling patterns
2. Each group becomes a child model struct with its own Update/View
3. Parent stores children as named fields
4. Parent Update routes messages to children based on type
5. Parent View composites children Views

### Stabilize Viewport

When detected `viewport.New(` in Update/View:
1. Move `viewport.New(w, h)` to Init() or struct initialization
2. In Update(), handle `WindowSizeMsg` → `viewport.SetSize(w, h)` on the model field
3. Add a content hash/diff check before calling `SetContent()`:
   ```go
   if newContent != m.viewportContent {
       m.viewport.SetContent(newContent)
       m.viewportContent = newContent
   }
   ```

### Blocking Decomposition

When blocking I/O in Update:
```diff
- resp, err := http.Get(url)
- // ... process resp
+ return func() tea.Msg {
+     resp, err := http.Get(url)
+     if err != nil {
+         return HttpErrorMsg{err}
+     }
+     return HttpRespMsg{resp}
+ }}
```

## Safety Checks

Before applying any patch:
1. Verify no existing tests break: `go test ./...`
2. Verify ATree impact scope is within expected bounds
3. Verify no new constraint violations introduced
4. Create a `.refactor-backup/` directory with original files

## Environment

- `ATREE_DB_PATH`: ATree database for impact analysis
- `CHARM_TUI_REFACTOR_BACKUP`: Create backups before applying (default: true)
