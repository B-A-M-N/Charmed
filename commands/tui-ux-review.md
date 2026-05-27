---
name: charm-tui:ux-review
description: "Perceptual UX critique of a TUI — pacing, density, focus behavior, responsiveness against timing contracts."
argument-hint: "[path-to-project] [--archetype <id>] [--strictness normal|strict]"
allowed-tools: [Read, Grep, Glob, Bash, Agent]
---

Invoke the tui-ux-review skill to evaluate how a Bubble Tea TUI feels to use.

## Instructions

1. Scan the project and detect its archetype from `archetypes/archetypes.yaml`.
2. Read `timing/timing-models.yaml` for the relevant timing contracts.
3. Evaluate pacing, density, motion, focus behavior, and responsiveness.
4. Present findings with specific, actionable recommendations.

## Usage

```
/charm-tui:ux-review                                  # review current dir
/charm-tui:ux-review --strictness strict              # tighter thresholds
/charm-tui:ux-review /home/bamn/TalonCLI              # review specific path
/charm-tui:ux-review --archetype content-reader       # override archetype
```
