---
name: tui-instrument
description: "Runtime semantics instrumentation — generates telemetry hooks for the message loop, command execution, render path, and async boundaries. Produces frame timing traces, command latency measurements, and focus transition logs that can be replayed for perceptual analysis. Trigger when: user says 'instrument', 'add tracing', 'add telemetry', 'profile TUI', 'measure performance', 'debug timing', 'where is the lag', or wants observability in a Bubble Tea app."
argument-hint: "[path-to-project] [--target messages|cmds|render|all] [--dry-run] [--apply]"
allowed-tools: [Read, Write, Edit, Bash, Grep, Glob, ATree]
---

# tui-instrument — TUI Telemetry Generator

Generates Go source code that instruments a Bubble Tea application with runtime tracing hooks. Can trace message flow timing, command execution latency, and render path duration. Output is deterministic (no LLM generation of arbitrary logic — fixed templates with project-specific bindings).

## Inputs

| Input | Required | Source |
|-------|----------|--------|
| Project path | Yes | User argument (default: current directory) |
| IR / scan results | Yes | TUI Structural IR from scanner (or auto-scan) |
| Timing models | Yes | `timing/timing-models.yaml` |
| Instrument targets | No | `--target messages|cmds|render|all` (default: all) |
| ATree DB path | No | For call graph analysis |

## Execution Pipeline

### Phase 1: Scanning

Run a lightweight scan of the project to identify:
1. All custom message types (for message tracing)
2. All `tea.Cmd` definitions and spawn sites (for command timing)
3. All `View()` methods (for render timing)
4. Root model struct and `Init()` signature (for tracing lifecycle)

Use same scanning approach as tui-audit Phase 2, focusing only on instrumentable targets.

### Phase 2: Plan

Produce an instrumentation plan:

```
instrumentation_plan:
  models:
    - model: "RootModel"
      add_fields:
        - name: "trace *trace.Span"
          type: "opentracing.Span"
        - name: "msgCount int"
          type: "int"
        - name: "lastRender time.Duration"
          type: "time.Duration"
      trace_points:
        - location: "Update() entry"
          trace: "msg_in"
          capture: "msg type, timestamp"
        - location: "Cmd spawn"
          trace: "cmd_start"
          capture: "cmd type"
        - location: "View() entry"
          trace: "render_start"
          capture: "-"
        - location: "View() exit"
          trace: "render_end"
          capture: "duration"
```

### Phase 3: Code Generation

Generate these files:

#### `tracing.go` — Tracing Message Types and Hooks

```go
package tracing

import (
    "fmt"
    "time"
    tea "github.com/charmbracelet/bubbletea"
)

// TraceMsg wraps any tea.Msg with timing metadata.
type TraceMsg struct {
    Inner    tea.Msg
    Entered  time.Time
    Source   string // which model/component sent this
}

func (t TraceMsg) String() string {
    return fmt.Sprintf("trace[%s → %T after %v]", t.Source, t.Inner, time.Since(t.Entered))
}

// CmdTimingMsg reports command execution duration.
type CmdTimingMsg struct {
    Name     string
    Duration time.Duration
    Err      error
}

// RenderTimingMsg reports View() render duration.
type RenderTimingMsg struct {
    Model    string
    Duration time.Duration
}

// TimedCmd wraps a tea.Cmd to measure its execution time.
func TimedCmd(name string, cmd tea.Cmd) tea.Cmd {
    return func() tea.Msg {
        start := time.Now()
        msg := cmd()
        return CmdTimingMsg{
            Name:     name,
            Duration: time.Since(start),
        }
    }
}

// WithTrace wraps a msg entering Update() with trace metadata.
func WithTrace(source string, msg tea.Msg) tea.Cmd {
    return func() tea.Msg {
        return TraceMsg{
            Inner:   msg,
            Entered: time.Now(),
            Source:  source,
        }
    }
}
```

#### `instrument_model.go` — Model Instrumentation Mixin

```go
package instrument

import (
    "time"
    "github.com/charmbracelet/bubbletea"
)

// InstrumentedModel wraps tracing around any model's Update/View.
type InstrumentedModel struct {
    Name      string
    MsgCount  int
    MaxMsgTime time.Duration
    TotalMsgTime time.Duration
}

func (im *InstrumentedModel) WrapUpdate(name string, msg tea.Msg, updateFn func(tea.Msg) (tea.Model, tea.Cmd)) (tea.Model, tea.Cmd) {
    start := time.Now()
    model, cmd := updateFn(msg)
    elapsed := time.Since(start)
    im.MsgCount++
    im.TotalMsgTime += elapsed
    if elapsed > im.MaxMsgTime {
        im.MaxMsgTime = elapsed
    }
    return model, cmd
}

func (im *InstrumentedModel) WrapView(name string, viewFn func() string) string {
    start := time.Now()
    result := viewFn()
    _ = time.Since(start) // captured or logged externally
    return result
}

func Stats() string {
    avg := time.Duration(0)
    if im.MsgCount > 0 {
        avg = im.TotalMsgTime / time.Duration(im.MsgCount)
    }
    return fmt.Sprintf("model[%s] msgs=%d avg=%v max=%v",
        im.Name, im.MsgCount, avg, im.MaxMsgTime)
}
```

#### `main.go` instrumentation patch

For each instrumented target, show a diff or modification marker:

```diff
+ import "your-project/tracing"

  func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
+     start := time.Now()
      switch msg := msg.(type) {
      // ... existing cases (unchanged)
      }
+     elapsed := time.Since(start)
+     if elapsed > 50*time.Millisecond {
+         log.Printf("slow Update: %T took %v", msg, elapsed)
+     }
  }
```

### Phase 4: Timing Validation

Cross-reference generated instrumentation with `timing/timing-models.yaml`:

| Timing Model | What to Instrument | Threshold to Flag |
|---|---|---|
| smooth_streaming | Cmd callbacks, SetContent calls | > 50ms latency |
| responsive_input | KeyMsg → echo in View | > 16ms |
| periodic_refresh | Tick handler duration | > 200ms render |
| 60fps_cap | View() total duration | > 16ms per frame |
| snappy_focus | Focus change → first render | > 16ms |

For each timing model applicable to the detected archetype, ensure thresholds are checked in generated code.

### Phase 5: Output

**Dry run (default):** Show the instrumentation plan and generated files without writing.

**With confirmation or `--apply`:** Write files to `<project>/internal/tracing/`, `<project>/internal/instrument/`. Show modified files list.

```
Instrumentation plan for <project>:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  Messages traced: <N> custom message types
  Cmds traced: <N> tea.Cmd spawn sites
  Render hooks: <N> View() methods

  Files to create:
    + internal/tracing/tracing.go    (TraceMsg, TimedCmd, WithTrace)
    + internal/tracing/types.go      (CmdTimingMsg, RenderTimingMsg)
    + internal/instrument/model.go   (InstrumentedModel mixin)

  Files to modify:
    ~ main.go                        (wrap tea.Cmd returns)
    ~ model.go                       (timed Update entry/exit)

  Timing thresholds:
    smooth_streaming: flag > 50ms cmd latency
    responsive_input: flag > 16ms key echo lag
    60fps_cap: flag > 16ms render

Apply? [y/N]
```

## Instrumentation Modes

### `--target messages`
Only trace message flow: wrap every tea.Msg through Update() with timing. Generate TraceMsg type and message log.

### `--target cmds`
Only trace command execution: wrap every tea.Cmd with TimedCmd. Generate CmdTimingMsg and latency histogram.

### `--target render`
Only trace rendering: time every View() call. Generate RenderTimingMsg and frame time log.

### `--target all` (default)
Full instrumentation: messages + cmds + render.

## Generated Tracing Output Format

All tracing output goes to a configurable `io.Writer` (default: `/tmp/tui-trace.log`):

```
2026-05-26T12:00:01.000Z  MSG    KeyMsg{runes:['j']}            dt=0.02ms
2026-05-26T12:00:01.050Z  CMD    fetchData     duration=48ms
2026-05-26T12:00:01.051Z  RENDER RootModel      duration=0.8ms
2026-05-26T12:00:01.200Z  MSG    DataMsg{items:5}               dt=0.01ms
2026-05-26T12:00:02.000Z  TICK   periodicRefresh               dt=0.1ms
```

## Environment

- `CHARM_TUI_TRACE_LOG`: Trace output path (default: `/tmp/tui-trace.log`)
- `CHARM_TUI_TRACE_LEVEL`: Trace verbosity (warn|fatal|all, default: warn — only logs violations)
