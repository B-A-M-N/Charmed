---
name: charmed:audit
description: "Architecture linter for Bubble Tea projects — scans Go source for anti-patterns, constraint violations, and performance hazards."
argument-hint: "[path-to-project] [--strict] [--format text|json]"
allowed-tools: [Read, Write, Bash, Grep, Glob, Agent]
---

Invoke the tui-audit skill to scan a Charm ecosystem TUI project.

## Instructions

1. Accept the project path from the argument (default: current directory).
2. Invoke the tui-audit execution pipeline as described in `skills/tui-audit/SKILL.md`.
3. Read `constraints/constraints.yaml` for rule definitions.
4. Present findings in the specified format (text table or JSON).

## Usage

```
/charm-tui:audit                          # audit current directory
/charm-tui:audit /home/bamn/TalonCLI      # audit specific path
/charm-tui:audit --strict                 # include info-level findings
/charm-tui:audit --format json            # machine-readable output
```
