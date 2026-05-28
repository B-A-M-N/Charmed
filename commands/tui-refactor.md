---
name: charmed:refactor
description: "Architectural refactoring for Bubble Tea projects — god model decomposition, viewport churn fixes, with ATree impact analysis."
argument-hint: "[path-to-project] [--target <symbol>] [--dry-run] [--apply]"
allowed-tools: [Read, Write, Edit, Bash, Grep, Glob, Agent]
---

Invoke the tui-refactor skill to restructure a Bubble Tea project.

## Instructions

1. Accept project path and optional target symbol.
2. Run violation detection (tui-audit equivalent).
3. Produce minimal patches for FATAL and ERROR findings.
4. Default to dry-run. Only write changes with `--apply`.

## Usage

```
/charm-tui:refactor                                  # dry-run audit + patch plan
/charm-tui:refactor --target LayoutModel             # refactor specific model
/charm-tui:refactor --apply                          # write changes
/charm-tui:refactor /home/bamn/TalonCLI --target .   # full project refactor
```
