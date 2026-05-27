---
name: tui-audit
description: "Architecture intelligence — deterministic constraint enforcement against the TUI Structural IR. Detects architectural violations (god model, blocking update, viewport churn, unbounded concurrency) with severity levels and fix patterns. Code-graph enriched via ATree. Trigger when: user says 'audit', 'check my TUI', 'find anti-patterns', 'lint', 'review architecture', 'detect violations', or asks about Bubble Tea model structure quality."
argument-hint: "[path-to-project] [--strict] [--rules rule1,rule2] [--format text|json]"
allowed-tools: [Read, Write, Bash, Grep, Glob, Agent]
---

# tui-audit — Charm Architecture Linter

Deterministic detection of anti-patterns, constraint violations, and architectural hazards in Bubble Tea projects, enriched with ATree graph queries for cross-file impact evidence.

## Inputs

| Input | Required | Source |
|-------|----------|--------|
| Project path | Yes | User argument (default: current directory) |
| constraints.yaml | Yes | `constraints/constraints.yaml` from plugin root |
| ATree DB path | No | Config: `atree.db.path` or auto-detected from `<project>/.atree/index.sqlite` |
| Strictness | No | `--strict` flag (default: warning threshold, strict: info threshold) |

## Execution Pipeline

### Phase 1: Project Discovery

1. Accept project path from argument or use `pwd`.
2. Verify it's a Go project: look for `go.mod`.
3. Detect Charm dependencies: grep `go.mod` for `charmbracelet/bubbletea`, `charmbracelet/bubbles`, `charmbracelet/lipgloss`.
4. If no Charm deps found, warn but continue — the project may have a non-standard module path.
5. Locate ATree index: `<project>/.atree/index.sqlite`, or `${ATREE_DB_PATH}`.

### Phase 2: TUI Structural IR Scan

For each `.go` file in the project (skip `*_test.go`, `vendor/`):

1. **Model Detection**: Find all structs implementing Bubble Tea interfaces:
   - Pattern: types with `Update(...) (Model, tea.Cmd)` method
   - Pattern: types with `View() string` method
   - Pattern: types with `Init() tea.Cmd` method
   - Extract: struct name, file, line, field list with types

2. **Message Type Detection**: Find all custom message types:
   - Pattern: `type XxxMsg struct` or `type XxxMsg interface`
   - Extract: handlers (which model's Update handles it), emitters (where tea.Cmd sends it)

3. **Update() Analysis**: For each model's Update method:
   - Count LOC
   - Count switch branches
   - Detect blocking patterns: `http.`, `ioutil.ReadFile`, `os.ReadFile`, `time.Sleep`, `sql.`
   - Detect viewport reconstruction: `viewport.New(` inside Update/View
   - Detect viewport churn: `SetContent(` call sites
   - Detect unbounded command spawning in loops
   - Detect `tea.Batch` inside state-dependent blocks

4. **Init() Analysis**: For each model's Init method:
   - Detect blocking I/O patterns
   - Verify it returns a `tea.Cmd` (not nil when work is needed)

5. **View() Analysis**: For each model's View method:
   - Count `fmt.Sprintf`, `lipgloss.NewStyle`, `.Width(`, `.Height(` calls
   - Detect style recomputation in hot path
   - Detect excessive string concatenation

6. **Component Bindings**: Detect which Charm components are used:
   - `bubbles/viewport.Model` → track SetContent sites, resize handling
   - `bubbles/list.Model` → track item count, filter usage
   - `bubbles/textarea.Model` → focus/blur lifecycle
   - `bubbles/textinput.Model` → validation patterns
   - `bubbles/help.Model` → key binding completeness
   - `lipgloss.Style` → cache vs recompute analysis

### Phase 3: Rule Matching

Load `constraints/constraints.yaml`. For each rule:

1. Match against Phase 2 findings using the rule's `match` pattern.
2. Record: `rule_id`, `severity`, `file`, `line`, `message`, `fix.suggestion`.
3. If ATree DB is available, enrich with:
   - `atree impact <symbol>` — what else is affected by the flagged code?
   - `atree evidence_path <from> <to>` — trace the call path to the violation.

Filter findings by severity threshold:
- Default: report `warning`, `error`, `fatal`
- `--strict`: report `info` and above

### Phase 4: Report

Output format:

```
╔══════════════════════════════════════════════════════════════╗
║  charm-tui audit — <project-name>                           ║
╠══════════════════════════════════════════════════════════════╣
║  Scanned: <N> files, <M> models, <K> message types          ║
║  Checks: <total> rules · <passed> passed · <failed> failed  ║
╠══════════════════════════════════════════════════════════════╣
║                                                            ║
║  🔴 FATAL (N)                                              ║
║  ────────────────────────────────────────────────────────  ║
║  [viewport_recreation] cmd/app.go:42                       ║
║    viewport.Model recreated inside Update().               ║
║    Fix: Create once in Init(), use SetContent() only.      ║
║                                                            ║
║  🛑 ERROR (N)                                              ║
║  ────────────────────────────────────────────────────────  ║
║  [god_model] cmd/app.go:87                                 ║
║    Update() is 340 lines. Decompose into child models.     ║
║                                                            ║
║  ⚠️  WARNING (N)                                           ║
║  ────────────────────────────────────────────────────────  ║
║  [missing_key_bindings_help]                               ║
║    No key.Binding import detected. Add bubbles/help.       ║
║                                                            ║
║  ℹ️  INFO (N) [strict mode only]                           ║
║  ────────────────────────────────────────────────────────  ║
║  [string_allocation_pressure] cmd/app.go:120               ║
║    Heavy string formatting inside View().                  ║
║                                                            ║
╚══════════════════════════════════════════════════════════════╝
```

With `--format json`, output the findings as JSON:
```json
{
  "project": "my-tui",
  "scanned_at": "2026-05-26T...",
  "summary": { "total_checks": 14, "passed": 8, "failed": 6, "by_severity": {"fatal":1,"error":2,"warning":2,"info":1} },
  "findings": [ ... ]
}
```

## Anti-Pattern Cheat Sheet

| Anti-Pattern | What to Look For | Severity |
|---|---|---|
| God model | Update() > 300 LOC, >20 switch branches | error/fatal |
| Viewport churn | SetContent() called every frame | error |
| Viewport recreation | viewport.New() in Update() or View() | fatal |
| Blocking Update | http.Get, time.Sleep, file I/O in Update() | error |
| Blocking Init | I/O in Init() instead of Cmd | error |
| Style recomputation | lipgloss.NewStyle() in View() | warning |
| Missing help | No key.Binding or help.Model import | warning |
| Unbounded cmds | tea.Cmd spawned in tight loop | error |
| Missing resize | viewport without WindowSizeMsg handler | error |
| Batch misuse | tea.Batch where order matters | info |
| Missing alt screen cleanup | EnterAltScreen without deferred ExitAltScreen | warning |

## Charm-Specific Patterns

When scanning Charm ecosystem code, also detect:

- **Bubble routing anti-pattern**: Messages not forwarded to child models
- **Viewport without key bindings**: Scrollable viewport without ↑↓/PgUp/PgDn bindings
- **Focus loss in multi-model**: Multiple textarea/textinput without clear focus state
- **Missing Spinner.Next**: Spinner animation not advanced via SpinnerMsg in Update

## Environment

- `ATREE_DB_PATH`: Override auto-detected ATree database path
- `CHARM_TUI_STRICT`: Set strict mode by default
- `CHARM_TUI_FORMAT`: Default output format (text|json)
