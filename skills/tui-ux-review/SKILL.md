---
name: tui-ux-review
description: "Perceptual UX validation — evaluates the TUI against timing contracts (smooth streaming, snappy focus, responsive input, periodic refresh). Measures pacing, scroll stability, focus transition latency, and render cadence against perceptual benchmarks. Validates density profile and motion philosophy against archetype expectations. Trigger when: user says 'how does this TUI feel', 'review UX', 'critique', 'is this responsive', 'perceptual', 'pacing', 'does this feel right', 'UX review', or asks about TUI feel and responsiveness."
argument-hint: "[path-to-project] [--archetype <id>] [--timing-model <id>] [--strictness normal|strict]"
allowed-tools: [Read, Grep, Glob, Bash, Agent]
---

# tui-ux-review — Perceptual TUI Critique

Evaluates a Bubble Tea application's perceived UX quality — not whether the code is structurally sound (that's tui-audit), but whether it **feels right** to a user. Grounds subjective perception in measurable timing model contracts from `timing/timing-models.yaml`.

## Inputs

| Input | Required | Source |
|-------|----------|--------|
| Project path | Yes | User argument (default: current directory) |
| IR / scan results | Yes | TUI Structural IR from scanner (or auto-scan) |
| Timing models | Yes | `timing/timing-models.yaml` |
| Archetypes | Yes | `archetypes/archetypes.yaml` |
| Archetpe override | No | `--archetype <id>` (skip auto-detection) |
| Strictness | No | `--strictness normal|strict` |

## Execution Pipeline

### Phase 1: Archetype Detection

If no `--archetype` specified, classify the project's intent by scanning:

1. What components are used (viewport → content display; list → item selection; textarea → input)
2. What message types exist and their naming patterns
3. Layout structure from View() composition

Match to archetype using `archetypes/archetypes.yaml` trait scoring (same as tui-scaffold Phase 1).

The archetype determines:
- Which timing model contract to evaluate against
- Which perceptual dimensions matter
- What the user likely expects from the interaction

### Phase 2: Timing Model Contract Mapping

For the detected archetype, load the corresponding timing model(s) from the archetype's `timing_constraints`. Also load from `timing/timing-models.yaml`.

Example: if archetype is `agent_console` → timing model is `smooth_streaming` + `responsive_input`.

Extract the hard constraints:
- `render_frequency_max_fps`: 30
- `cmd_latency_warn_ms`: 50
- `cmd_latency_error_ms`: 150
- `viewport_mutation_budget_per_frame`: 2
- `echo_latency_ms`: 16
- etc.

These become the measurable contract. Every UX finding traces back to a specific constraint.

### Phase 3: Perceptual Dimension Analysis

Evaluate each dimension against the timing contract:

#### Pacing
- **Question**: Do things happen at the right rate?
- **Measure**: Message frequency, Update() call rate, Cmd completion time
- **Scan for**: Blocking operations that exceed timing thresholds, slow cmd callbacks
- **Flag if**: Any single Update() takes > 16ms, any Cmd exceeds `cmd_latency_warn_ms`

#### Density
- **Question**: Is the screen information-dense without being overwhelming?
- **Measure**: Number of component types used, visual elements per screen area
- **Scan for**: Too many components (cognitive overload), too little (wasted space)
- **Psychology**: 3-5 distinct visual regions max per screen (Miller's law)
- **Flag if**: > 5 distinct interactive regions, or < 2 for a multi-component app

#### Motion
- **Question**: Are transitions and animations smooth, not jarring?
- **Measure**: Viewport content change rate, re-render frequency
- **Scan for**: Viewport recreation (causes visual jump), style re-recomputation (flicker), spinner without Next()
- **Flag if**: Content replaced instead of appended, spinner doesn't animate

#### Focus
- **Question**: Is it always clear where input goes?
- **Measure**: Focus state management in model fields
- **Scan for**: Multiple focused inputs, no focus indicator, focus lost on resize
- **Flag if**: Focus state not tracked per-component, no visual focus indicator

#### Responsiveness
- **Question**: Does the UI react immediately to user input?
- **Measure**: KeyMsg handling path length, echo latency detection
- **Scan for**: Key handling through multiple layers, async key processing, missing AltScreen
- **Flag if**: Key press → View() update path exceeds 1 frame, no alt screen (scroll-through artifacts)

#### Readability
- **Question**: Can users read content without strain?
- **Measure**: Text width usage, style contrast, content layout
- **Scan for**: Overwide lines (> 120 chars), insufficient contrast, content outside viewport
- **Flag if**: Viewport width not constrained, no max-width on content

### Phase 4: Timing Violation Detection

For each timing model constraint, verify compliance:

```
smooth_streaming check:
  ✓ render_frequency: Update() returns < 33ms equivalent? (scan for heavy loops)
  ✓ cmd_latency: Cmd callbacks < 50ms average? (scan for blocking calls)
  ✓ viewport_mutation: SetContent called ≤ 2x/frame? (scan call sites)
  ✗ string_allocations: 3+ Sprintf in View()? [VIOLATION]

responsive_input check:
  ✓ echo_latency: KeyMsg reaches View() in 1 frame?
  ✓ no_blocking_key_handler: No I/O in KeyMsg branch?
```

### Phase 5: UX Report

```
╔══════════════════════════════════════════════════════════════╗
║  charm-tui UX Review — <project-name>                       ║
║  Archetype: <detected or specified>                          ║
╠══════════════════════════════════════════════════════════════╣
║                                                            ║
║  PERCEPTUAL SCORES                                         ║
║  ────────────────────────────────────────────────────────  ║
║  Pacing:        ████████░░  80%                            ║
║  Density:       █████████░  90%                            ║
║  Motion:        ██████░░░░  60%  ⚠️ viewport churn        ║
║  Focus:         ██████████ 100%                            ║
║  Responsiveness:███████░░░  70%  ⚠️ blocking cmd          ║
║  Readability:   ████████░░  80%                            ║
║                                                            ║
║  TIMING CONTRACT VIOLATIONS                                ║
║  ────────────────────────────────────────────────────────  ║
║  [smooth_streaming]                                        ║
║  ⚠️  View() computes 5 fmt.Sprintf per frame              ║
║      → 3 allocations over budget (max: 3, got: 5)         ║
║      → Fix: Cache formatted strings, use strings.Builder  ║
║                                                            ║
║  [responsive_input]                                        ║
║  ✗  File I/O in KeyMsg handler (cmd/app.go:156)           ║
║      → Blocks input echo by ~120ms                        ║
║      → Fix: Wrap in tea.Cmd, send result as message       ║
║                                                            ║
║  ARCHETYPE-SPECIFIC NOTES                                  ║
║  ────────────────────────────────────────────────────────  ║
║  As a <archetype>, users expect:                           ║
║  • Smooth content arrival (no blank frames)               ║
║  • Stable cursor position during streaming                ║
║  • Visible key hints for discoverability                  ║
║                                                            ║
║  ✓ Content streaming: viewport used with SetContent       ║
║  ✗ Key hints: no help.Model or key bindings visible       ║
║  ✓ Alt screen: EnterAltScreen configured                  ║
║                                                            ║
╚══════════════════════════════════════════════════════════════╝
```

## Archetype-Specific Checklists

### Agent Console (`agent_console`)
- [ ] Streaming content appears within 50ms of data arrival
- [ ] Scroll position preserved during content append
- [ ] No blank frames between content batches
- [ ] Input always visible and responsive
- [ ] Typing indicator or spinner visible during agent processing

### IDE Workspace (`ide_workspace`)
- [ ] Focus changes within one frame of keypress
- [ ] No ghost/flicker during pane switches
- [ ] All panes resize on WindowSizeMsg without layout pop
- [ ] Modal overlays capture all input
- [ ] Status bar updates without affecting pane content

### Monitoring Dashboard (`monitoring_dashboard`)
- [ ] Values tick in place (no repositioning)
- [ ] Layout structure stable across refreshes
- [ ] Old→new value transitions smooth (no blink)
- [ ] Alert surfaces visible but not disruptive
- [ ] Refresh does not block user input

### CLI Wizard (`cli_wizard`)
- [ ] Validation errors inline, not modal
- [ ] Field transition (Next/Prev) instant
- [ ] Step indicator visible
- [ ] Back button always available except first step
- [ ] Confirm/cancel clearly labeled

## Strictness Levels

- **normal**: Flag only clear violations (> 2x over threshold)
- **strict**: Flag anything over threshold, include suggestions for marginal cases

## Environment

- `CHARM_TUI_UX_STRICTNESS`: Default strictness level
- `CHARM_TUI_ARCHETYPE`: Override archetype detection
