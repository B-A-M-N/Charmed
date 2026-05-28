---
name: tui-explain
description: "Evidence-based knowledge queries — explains primitives, patterns, constraints, and their derivation. Shows why rules exist by referencing ecosystem observations, violations, and fixes. Answers architectural questions with empirical backing. Trigger when: user asks 'what is X', 'why does constraint Y exist', 'show me pattern Z', 'what's the evidence for', 'how do repos handle X', or questions about ecosystem practices."
argument-hint: "<concept or question> [--with-evidence] [--show-repos] [--primitives-only]"
allowed-tools: [Read, Grep, Glob, Bash, ATree, mcp__charmed__explain_primitive, mcp__charmed__explain_pattern, mcp__charmed__explain_constraint]
---

# tui-explain — Evidence-Based Knowledge Engine

Explains architectural concepts using **two-tier knowledge**:
- **Canonical knowledge**: Runtime semantics (how Bubble Tea works)
- **Empirical knowledge**: Observed patterns (how the ecosystem uses it)

## Query Types

### 1. Primitive Queries
"What is the viewport primitive?"
"Explain tick_source"

**Answers with:**
- Primitive definition from `primitives/primitives.yaml`
- Which libraries implement it
- Signals that indicate presence
- Constraints that apply

**Example response:**
```
Primitive: viewport
Category: rendering
Description: Scrollable content window with offset tracking

Implemented by:
- bubbles/viewport.Model
- Observed in 47 repos

Signals:
- viewport.Model field declaration
- SetContent() calls  
- Scroll offset management

Constraints:
- viewport_recreation (canonical, confidence: 1.0)
- viewport_churn (empirical, confidence: 0.94)

Common patterns:
- viewport_content_diffing (9 repos)
- viewport_with_resize_handling (23 repos)
```

### 2. Pattern Queries
"What is viewport_content_diffing?"
"Show me the storage_cmd_boundary pattern"

**Answers with:**
- Pattern definition from `patterns/`
- Confidence score and derivation
- Repos that use it
- Violations when absent
- Fixes that applied it

**Example response:**
```
Pattern: viewport_content_diffing
Confidence: 0.85
Category: rendering

Primitives: viewport + diffed_content
Topology: seq:diffed_content->viewport.SetContent

Observed in 9 repos:
- soft-serve (ui/dashboard.go:156)
- glow (ui/pager.go:89)
- mods (output.go:234)
... (6 more)

Violations (4 issues):
- early-tui-app#12: "Scroll position resets"
- buggy-viewer#8: "Flicker on every update"
... (2 more)

Fixes (4 PRs):
- early-tui-app PR#15: Added hash check before SetContent
- buggy-viewer PR#10: Content diffing eliminated flicker
... (2 more)

Why this works:
viewport.SetContent() resets scroll when content changes.
Diffing prevents unnecessary calls, maintaining scroll stability.
```

### 3. Constraint Queries
"Why does the no_blocking_update constraint exist?"
"Explain storage_must_be_async"

**Answers with:**
- Constraint definition
- Type (canonical vs empirical)
- Confidence and evidence
- Pattern it derives from (if empirical)
- Violation examples
- Fix recommendations

**Example response (empirical constraint):**
```
Constraint: storage_must_be_async
Type: empirical
Confidence: 0.94
Severity: error

Description:
Storage operations (charm, skate, file I/O) must be wrapped
in async commands, never called directly from Update() or View().

Derived from pattern:
storage_cmd_boundary (14 observations, 11 fixes)

Evidence:
Observations: 14 repos wrap storage in tea.Cmd
Violations: 12 repos had blocking issues
Fixes: 11 repos fixed by async wrapping
Fix success rate: 91.7%

Violation examples:
- blocking-app#8: "UI freezes when saving settings"
  Root cause: skate.Set() in Update()
  Fix: Moved to Cmd wrapper (PR#10)

- slow-tui#15: "Multi-second hangs on config load"
  Root cause: os.ReadFile() in Init()
  Fix: Async Cmd for loading (PR#18)

Why this matters:
Bubble Tea's event loop is single-threaded. Any blocking
call in Update() freezes the entire UI until it completes.
Storage I/O can take 10ms-1000ms depending on disk/network.
```

**Example response (canonical constraint):**
```
Constraint: no_blocking_update
Type: canonical
Confidence: 1.0
Severity: error

Description:
Blocking I/O or synchronous calls inside Update() block
the event loop and freeze the UI.

Runtime requirement:
Bubble Tea's architecture requires Update() to return
immediately. The event loop cannot process input, render,
or handle messages while Update() is executing.

This is a runtime semantic, not an observed pattern.
```

### 4. Library Usage Queries
"How do repos use harmonica?"
"What patterns exist for colorprofile?"

**Answers with:**
- Primitives from that library
- Observed usage patterns
- Integration strategies
- Common pairings

**Example response:**
```
Library: harmonica
Primitives: animated_transition

Usage patterns:
1. animation_loop_with_cleanup (17 repos)
   - Spring physics in goroutine
   - Tick messages drive updates
   - Cleanup on quit
   - Confidence: 0.87

2. smooth_scroll_interpolation (8 repos)
   - Spring for viewport offset
   - Read interpolated value in View()
   - Update spring in Update()
   - Confidence: 0.82

Common pairings:
- harmonica + viewport (smooth scrolling)
- harmonica + lipgloss (animated transitions)
- harmonica + tick_source (frame timing)

Constraints:
- no_spring_calculation_in_view (empirical, 0.88)
- must_cleanup_tickers (empirical, 0.91)
```

### 5. Evidence Queries
"Show me evidence for viewport diffing"
"Which repos use storage Cmds?"

**Answers with:**
- Direct evidence references
- Files and line numbers
- Commit links
- Issue/PR links

**Example response:**
```
Evidence for: viewport_content_diffing pattern

Observations (9):
1. soft-serve
   File: ui/dashboard.go:156
   Pattern: if hashContent(new) != m.hash { m.viewport.SetContent(new) }
   Evidence ID: softserve_001

2. glow
   File: ui/pager.go:89
   Pattern: if m.lastContent != content { m.vp.SetContent(content) }
   Evidence ID: glow_003

... (7 more)

Violations (4):
1. early-tui-app (issue #12)
   Symptom: "Viewport scroll position resets on every update"
   Root cause: SetContent called without diff check
   Fix: PR#15 added hash comparison
   
... (3 more)
```

## Execution Flow

1. **Parse query** - identify what user is asking about
2. **Classify** - primitive, pattern, constraint, library, or evidence
3. **Load knowledge** - read appropriate YAML files
4. **Check confidence** - distinguish canonical vs empirical
5. **Gather evidence** - if requested, load supporting observations
6. **Format response** - structured, with confidence and sources

## MCP Tool Usage

```
mcp__charmed__explain_primitive(name: "viewport")
mcp__charmed__explain_pattern(name: "storage_cmd_boundary")
mcp__charmed__explain_constraint(id: "viewport_churn")
mcp__charmed__list_evidence(pattern: "viewport_content_diffing")
```

## Output Formatting

### Quick mode (default)
2-3 paragraph explanation with confidence if empirical

### With evidence (--with-evidence)
Full response including repos, issues, commits

### Primitives only (--primitives-only)
Skip library-specific details, just show primitive semantics

### With repos (--show-repos)
List all repos exhibiting pattern/using primitive

## Important Distinctions

**Canonical answers:**
- Based on runtime semantics
- Confidence always 1.0
- No "observed in X repos"
- Authority: official docs + runtime behavior

**Empirical answers:**
- Based on ecosystem observations
- Confidence 0.75-0.95
- Show evidence count
- Authority: derived from pattern mining

Always make this distinction clear in responses.
