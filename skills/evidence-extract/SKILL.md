---
name: evidence-extract
description: "Deterministic structural extraction — scans Go projects and extracts architecture evidence (signals, topologies, relationships) without AI inference. Maps implementation patterns to behavioral primitives. This is Stage 1 of the evidence-based analysis pipeline."
argument-hint: "[path-to-project] [--output evidence-db/] [--confidence-threshold 0.8]"
allowed-tools: [Read, Write, Bash, Grep, Glob, ATree, mcp__charmed__scan_project]
---

# evidence-extract — Structural Evidence Extraction

**This is deterministic. No LLM inference at this stage.**

Extracts observable architectural signals from Go source code and normalizes them into the evidence graph format defined in `evidence/evidence-schema.yaml`.

## Purpose

Evidence extraction is the foundation of the empirical constraint system. It produces structured, reproducible observations that pattern mining and constraint synthesis operate on.

## Pipeline Position

```
[Go Source] → **[Evidence Extraction]** → [Evidence DB] → [Pattern Mining] → [Constraint Synthesis]
```

This skill implements the **Evidence Extraction** stage.

## What Gets Extracted

### 1. Lifecycle Signals
- Init() implementations
- Update() message handlers
- View() rendering paths
- Quit/cleanup handlers
- WindowSizeMsg handling

### 2. Async Boundaries
- tea.Cmd creation sites
- Goroutine spawns
- Channel usage
- Ticker creation
- Cancellation handlers

### 3. Component Usage
- Viewport field declarations
- List/Textarea/TextInput usage
- Component initialization sites
- SetContent/Update call sites

### 4. Storage Boundaries
- charm/skate calls
- File I/O operations
- Database queries
- Where they occur (Init/Update/View/Cmd)

### 5. Rendering Patterns
- lipgloss.NewStyle calls
- Style caching vs recomputation
- Layout composition (JoinHorizontal/Vertical)
- String building strategies

### 6. Infrastructure Usage
- colorprofile detection
- harmonica spring physics
- wish SSH handling
- log setup

## Extraction Methods

### Primary: AST-based (via ATree)
```
Confidence: 0.95-1.0
Precision: High
Coverage: Complete for indexed projects
```

Uses ATree MCP to get:
- Symbol resolution
- Call graphs
- Type information
- Scope boundaries

### Fallback: grep-based
```
Confidence: 0.70-0.85
Precision: Medium
Coverage: Works without indexing
```

Pattern matching for:
- Function signatures
- Import statements
- Field types
- Method calls

## Output Format

Evidence entries conforming to `evidence/evidence-schema.yaml`:

```yaml
repo: example-repo
file: app/main.go
line: 45
context: AppModel.Update

signals:
  - id: sig_storage_001
    primitive: storage_boundary
    pattern: "skate.Get(key)"
    scope: method
    owner: loadConfigCmd

  - id: sig_cmd_002
    primitive: async_cmd
    pattern: "return func() tea.Msg"
    scope: method
    owner: loadConfigCmd

topology:
  structure_type: pipeline
  nodes: [sig_cmd_002, sig_storage_001]
  edges:
    - from: sig_cmd_002
      to: sig_storage_001
      relationship: calls
      confidence: 0.95

confidence: 0.95
extraction_method: atree
extracted_at: 2026-05-28T12:00:00Z
```

## Primitive Mapping

The key intelligence is mapping concrete code patterns to behavioral primitives:

### Example: Ticker Detection
```go
// Code pattern
tea.Tick(time.Second, func(time.Time) tea.Msg {
    return tickMsg{}
})
```

Maps to:
```yaml
primitive: tick_source
signals:
  - tea.Tick call
  - periodic interval
  - message emission
```

### Example: Storage in Cmd
```go
// Code pattern
func loadCmd() tea.Cmd {
    return func() tea.Msg {
        data := skate.Get("key")
        return loadedMsg{data}
    }
}
```

Maps to:
```yaml
primitives: [async_cmd, storage_boundary]
topology: pipeline
signals:
  - Cmd wrapper
  - Storage call inside Cmd
  - Result message emission
```

## Algorithm

### Stage 1: Project Bootstrap
1. Identify Go project root
2. Find all .go files
3. Detect Charm imports (bubbletea, bubbles, etc.)
4. Locate ATree DB or fall back to grep

### Stage 2: Model Discovery
1. Find all Bubble Tea models (structs with Init/Update/View)
2. Extract model fields and types
3. Map fields to primitives (viewport.Model → viewport primitive)
4. Record model ownership

### Stage 3: Lifecycle Extraction
For each model:
1. Extract Init() body
2. Extract Update() body with message branches
3. Extract View() body
4. Identify lifecycle phase signals

### Stage 4: Async Boundary Detection
1. Find all tea.Cmd returns
2. Trace goroutine spawns
3. Identify ticker creation
4. Find cancellation sites
5. Build async topology

### Stage 5: Integration Detection
1. Find storage calls (charm, skate, file I/O)
2. Find infrastructure usage (colorprofile, harmonica, log)
3. Determine where they occur (Update vs Cmd)
4. Flag violations (storage in Update)

### Stage 6: Topology Construction
1. Build signal graphs within each evidence entry
2. Link signals with relationships (calls, emits, depends_on)
3. Compute confidence based on extraction method
4. Serialize to evidence DB

## Usage

### Extract from single project
```
/evidence-extract /path/to/tui-project
```

Outputs:
```
evidence-db/
  tui-project/
    evidence-001.yaml
    evidence-002.yaml
    ...
    index.yaml
```

### Batch extraction
```
/evidence-extract --batch repos.txt
```

### With ATree
```
export ATREE_DB_PATH=/path/to/.atree/index.sqlite
/evidence-extract /path/to/project
```

### Quick mode (grep-only)
```
/evidence-extract --quick /path/to/project
```

## Confidence Scoring

### AST + ATree: 0.95-1.0
- Symbol resolution confirmed
- Call paths verified
- Type information complete

### AST only: 0.85-0.95
- Struct parsing confirmed
- Method signatures verified
- Limited cross-file resolution

### grep + heuristics: 0.70-0.85
- Pattern matching only
- No symbol resolution
- Higher false positive rate

### Manual review: 1.0
- Human-verified evidence
- Used for ground truth

## Output Structure

```
evidence-db/
  index.yaml                    # Master index of all evidence
  primitives-coverage.yaml      # Which primitives were observed
  
  repo-name/
    index.yaml                  # Repo-specific index
    evidence-001.yaml
    evidence-002.yaml
    ...
    
  violation-cases/              # Known anti-pattern evidence
    violation-001.yaml
    violation-002.yaml
    
  fix-cases/                    # Before/after evidence pairs
    fix-001.yaml
```

## Next Steps

After evidence extraction:
1. **Pattern mining** clusters similar evidence
2. **Constraint synthesis** derives rules from patterns
3. **Archetype enrichment** adds empirical patterns to archetypes

## Important

**This skill does not:**
- Infer patterns (that's pattern-mine)
- Generate constraints (that's constraint-synthesize)
- Make recommendations (that's tui-audit)

**This skill only:**
- Observes structure
- Normalizes to primitives
- Records evidence
- Computes confidence

The separation is critical for maintaining deterministic extraction and empirical rigor.
