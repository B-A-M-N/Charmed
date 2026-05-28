---
name: pattern-mine
description: "Pattern induction from evidence clusters — identifies recurring architectural topologies across evidence database, computes confidence scores, correlates violations with fixes. This is Stage 2 of the evidence-based analysis pipeline. AI assists clustering but does not define primitives."
argument-hint: "[--evidence-db evidence-db/] [--min-observations 8] [--output patterns/]"
allowed-tools: [Read, Write, Bash, Grep, Glob, Agent, ATree]
---

# pattern-mine — Evidence-Based Pattern Induction

Clusters evidence with equivalent topologies to discover recurring architectural patterns across the Charm ecosystem.

## Pipeline Position

```
[Evidence DB] → **[Pattern Mining]** → [Pattern Catalog] → [Constraint Synthesis]
```

## What This Does

Takes **deterministic evidence** from evidence-extract and finds:
- Recurring topology signatures
- Common primitive combinations
- Violation/fix correlations
- Ecosystem convergence patterns

## How It Works

### Stage 1: Evidence Normalization
Load all evidence entries and normalize topologies:

```yaml
# Evidence A: soft-serve dashboard
topology: seq:diffed_content->viewport.SetContent

# Evidence B: glow pager  
topology: seq:diffed_content->viewport.SetContent

# Evidence C: mods output
topology: seq:diffed_content->viewport.SetContent

→ Signature: "seq:diffed_content->viewport.SetContent"
→ Count: 3 repos
```

### Stage 2: Topology Clustering
Group evidence by signature similarity:
- Exact matches (confidence: 1.0)
- Structural equivalents (confidence: 0.9-1.0)
- Semantic variants (confidence: 0.7-0.9)

### Stage 3: Pattern Candidate Generation
For each cluster with ≥8 observations:

```yaml
pattern_candidate:
  topology_signature: "seq:diffed_content->viewport.SetContent"
  observation_count: 9
  repos: [soft-serve, glow, mods, ...]
  consistency_score: 0.92
```

### Stage 4: Violation Correlation
Link patterns to known violations:

```yaml
# Find evidence marked as violations
violation:
  symptom: "Viewport scroll resets"
  root_cause: "SetContent without diff check"
  
# Correlate with pattern absence
repos_with_pattern: [A, B, C, ...]  # No issues
repos_without_pattern: [X, Y, Z, ...]  # Had issues

correlation: 0.87
```

### Stage 5: Fix Correlation
Link violations to fixes:

```yaml
violation: viewport_scroll_reset_001
fix_commit: abc123
before_topology: "direct:viewport.SetContent"
after_topology: "seq:diffed_content->viewport.SetContent"

→ Pattern adoption fixed issue
→ Confidence boost: +0.15
```

### Stage 6: Confidence Computation

```python
confidence = (
    base_observation_score * 0.4 +
    consistency_score * 0.3 +
    violation_correlation * 0.2 +
    fix_correlation * 0.1
)

where:
  base_observation_score = min(observations / 20, 1.0)
  consistency_score = avg(topology_similarity)
  violation_correlation = violations_correlated / total_violations
  fix_correlation = fixes_adopting_pattern / total_fixes
```

### Stage 7: Pattern Promotion
If confidence ≥ 0.75 and observations ≥ 8:
→ Promote to pattern catalog

## AI Usage

**AI helps with:**
- Clustering similar topologies
- Identifying semantic equivalents
- Summarizing pattern properties
- Naming patterns

**AI does NOT:**
- Define primitives (humans do)
- Override confidence scores (formula does)
- Create patterns without evidence (threshold does)
- Infer patterns from vibes (clustering does)

## Output Format

Patterns conforming to `patterns/pattern-schema.yaml`:

```yaml
id: viewport_content_diffing
name: "Viewport Content Diffing"
category: rendering

primitives: [viewport, diffed_content]
topology_signature: "seq:diffed_content->viewport.SetContent"

confidence: 0.85
confidence_factors:
  observation_count: 9
  consistency_score: 0.92
  violation_count: 4
  fix_count: 4

derived_from:
  observations:
    - repo: soft-serve
      file: ui/dashboard.go
      line: 156
      evidence_id: softserve_001
    # ... 8 more

  violations:
    - repo: early-tui-app
      issue_url: https://github.com/example/issues/12
      symptom: "Scroll position resets"
      root_cause: "SetContent without diff check"

  fixes:
    - repo: early-tui-app
      commit_sha: def456
      improvement: "Scroll stable after adding hash check"

applies_to: [viewport]
generalizes_to: [streaming_console, log_viewer]
```

## Thresholds

From `constraints-v2.yaml`:

```yaml
min_observations: 8
min_violations: 3
min_fixes: 2
min_confidence: 0.75
violation_fix_ratio: 0.5
```

## Usage

### Mine from evidence database
```
/pattern-mine --evidence-db evidence-db/
```

### Focus on specific category
```
/pattern-mine --category async --evidence-db evidence-db/
```

### Lower thresholds for exploration
```
/pattern-mine --min-observations 5 --min-confidence 0.70
```

### Dry run (show candidates without writing)
```
/pattern-mine --dry-run
```

## Output Structure

```
patterns/
  pattern-schema.yaml           # Schema definition
  viewport_content_diffing.yaml
  storage_cmd_boundary.yaml
  animation_loop_with_cleanup.yaml
  ...
  
  index.yaml                    # Pattern catalog index
  confidence-report.yaml        # Confidence breakdown per pattern
```

## Pattern Evolution

Patterns are **continuously updated** as more evidence is extracted:

1. New repos analyzed → observation_count increases
2. New violations found → confidence adjusts
3. New fixes applied → confidence adjusts
4. Topology variants discovered → pattern may split or merge

This is **empirical evolution** - patterns track reality.

## Important Distinctions

**Pattern ≠ Library Component**
- A pattern is a behavioral topology
- It may involve multiple libraries
- It may appear across different component types

**Pattern ≠ Best Practice**
- A pattern is an observed regularity
- High confidence means "widely adopted and correlated with success"
- Not "architecturally pure" or "theoretically optimal"

**Pattern ≠ Rule**
- Patterns describe what exists
- Constraints prescribe what should exist
- Constraints are synthesized FROM patterns

## Next Steps

After pattern mining:
1. **Constraint synthesis** promotes high-confidence patterns to enforceable rules
2. **Archetype enrichment** adds patterns to TUI templates
3. **tui-audit** checks code against both canonical and empirical constraints
