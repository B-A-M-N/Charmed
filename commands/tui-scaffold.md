---
name: charmed:scaffold
description: "Intent-to-architecture TUI scaffolding — maps natural language to working Bubble Tea Go code."
argument-hint: "<intent-description> [--archetype <id>] [--path <output-dir>] [--dry-run]"
allowed-tools: [Read, Write, Bash, Glob, Grep]
---

Invoke the tui-scaffold skill to generate a Bubble Tea TUI project.

## Instructions

1. Parse the user's intent description and optional flags.
2. Read `archetypes/archetypes.yaml` and `constraints/constraints.yaml`.
3. Invoke the tui-scaffold execution pipeline from `skills/tui-scaffold/SKILL.md`.
4. Generate complete Go source files. Use `--dry-run` to show the plan without writing.

## Usage

```
/charm-tui:scaffold a CLI dashboard with a log viewer and command input
/charm-tui:scaffold a git log explorer --archetype list-nav
/charm-tui:scaffold a chat TUI --path ./my-chat --dry-run
```
