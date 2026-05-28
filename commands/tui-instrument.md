---
name: charmed:instrument
description: "Generate runtime telemetry hooks for Bubble Tea — message timing, command latency, render duration."
argument-hint: "[path-to-project] [--target messages|cmds|render|all] [--dry-run] [--apply]"
allowed-tools: [Read, Write, Edit, Bash, Grep, Glob]
---

Invoke the tui-instrument skill to add tracing/telemetry to a Bubble Tea project.

## Instructions

1. Scan the project for message types, tea.Cmd sites, and View() methods.
2. Read `timing/timing-models.yaml` for instrumentation templates.
3. Generate Go source with tracing hooks. Default to all targets.
4. Use `--dry-run` to preview, `--apply` to write.

## Usage

```
/charm-tui:instrument                                  # dry-run, all targets
/charm-tui:instrument --target messages                # only message tracing
/charm-tui:instrument --apply                          # write changes
/charm-tui:instrument /home/bamn/TalonCLI --target render
```
