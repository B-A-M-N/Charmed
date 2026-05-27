# charm-tui — Charm Ecosystem TUI Intelligence

Claude Code plugin for architectural reasoning, anti-pattern detection, intent-aware scaffolding, and UX critique for Bubble Tea / Lip Gloss / Bubbles projects.

## What It Does

| Skill | Trigger | What It Does |
|-------|---------|-------------|
| **tui-audit** | "audit my TUI", "find anti-patterns" | Deterministic architecture linting — scans Go source for anti-patterns, constraint violations, performance hazards |
| **tui-scaffold** | "create a TUI", "scaffold a terminal app" | Intent-to-architecture generator — maps natural language to working Bubble Tea code |
| **tui-refactor** | "refactor", "split this model", "too complex" | Architectural refactoring with ATree impact analysis — god model decomposition, viewport churn fixes |
| **tui-explain** | "what is X", "how does Y work" | Graph-driven semantic explanation of Charm ecosystem concepts via ATree queries |
| **tui-instrument** | "instrument", "add tracing", "profile TUI" | Generates runtime telemetry hooks — message timing, command latency, render duration |
| **tui-ux-review** | "review UX", "how does this feel", "critique" | Perceptual critique — pacing, density, focus behavior, responsiveness against timing contracts |

## Agent

- **charm-project-scanner** — Scans a Go project and produces the TUI Structural IR (models, messages, cmds, component bindings, state ownership, async boundaries). Auto-invoked by skills that need IR data.

## MCP Integration

- **atree** — Code graph server for cross-file impact analysis, call path tracing, and evidence-based architectural reasoning. Requires `atree` binary on PATH and `ATREE_DB_PATH` set.

## Installation

Already installed as a local plugin. Skills activate automatically when you use trigger phrases in conversation.

## Structure

```
charm-tui/
├── .claude-plugin/plugin.json    # Manifest
├── .mcp.json                     # MCP server config (atree)
├── skills/                       # 6 skills (tui-audit, tui-scaffold, tui-refactor, tui-explain, tui-instrument, tui-ux-review)
├── agents/                       # charm-project-scanner agent
├── constraints/                  # Lint rules (constraints.yaml)
├── archetypes/                   # TUI archetype catalog
├── timing/                       # Timing model contracts
├── ontology/                     # Charm knowledge base
├── ir/                           # TUI Structural IR schema
├── prompts/                      # Skill contract templates
└── README.md                     # This file
```
