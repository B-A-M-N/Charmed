# charmed — Charm Ecosystem TUI Intelligence

Claude Code plugin for architectural reasoning, anti-pattern detection, intent-aware scaffolding, and UX critique for Bubble Tea / Lip Gloss / Bubbles projects.

## What It Does

| Skill | Trigger | What It Does |
|-------|---------|-------------|
| **tui-audit** | "audit my TUI", "find anti-patterns" | Deterministic architecture linting — scans Go source for anti-patterns, constraint violations, performance hazards |
| **tui-scaffold** | "create a TUI", "scaffold a terminal app" | Intent-to-architecture generator — maps natural language to working Bubble Tea code |
| **tui-refactor** | "refactor", "split this model", "too complex" | Architectural refactoring — god model decomposition, viewport churn fixes |
| **tui-explain** | "what is X", "how does Y work" | Semantic explanation of Charm ecosystem components, patterns, and constraints |
| **tui-instrument** | "instrument", "add tracing", "profile TUI" | Generates runtime telemetry hooks — message timing, command latency, render duration |
| **tui-ux-review** | "review UX", "how does this feel", "critique" | Perceptual critique — pacing, density, focus behavior, responsiveness against timing contracts |

## ATree — Code Graph Engine

[`ATree`](https://github.com/Unity-Lab-AI/ATree) is the code graph server that powers `charmed-mcp`'s deep analysis. It provides cross-file semantic understanding — call path tracing, impact analysis, and evidence-based architectural reasoning.

**Required for:** `scan_project` full mode (model detection, message flow, component binding, async boundary analysis).

**Not required for:** Quick scans, knowledge lookups, archetypes/constraints/timing model listing.

Install ATree:

```bash
# Clone and build
git clone https://github.com/Unity-Lab-AI/ATree.git
cd ATree && go build -o ~/.local/bin/atree .

# Set the DB path for your project
export ATREE_DB_PATH=/path/to/your/project/.atree/index.sqlite
```

See [Unity-Lab-AI/ATree](https://github.com/Unity-Lab-AI/ATree) for full documentation.

## MCP Server

The `charmed-mcp` stdio server provides deep project scanning and knowledge base access:

```bash
# Build the binary (requires Go 1.25+)
bash scripts/build-charmed-mcp.sh

# Or manually:
cd ../charmed-mcp && go build -o ~/.local/bin/charmed-mcp .
```

The plugin's `.mcp.json` is pre-configured to launch `charmed-mcp` from PATH. After building, ensure `charmed-mcp` is accessible:

```bash
which charmed-mcp
# → /home/you/.local/bin/charmed-mcp
```

### MCP Tools

| Tool | Description |
|------|-------------|
| `scan_project` | Scan a Go project and produce the TUI Structural IR |
| `list_archetypes` | List TUI archetypes (agent_console, streaming_chat, ide_workspace, etc.) |
| `list_constraints` | List constraint rules (no_blocking_update, god_model, viewport_churn, etc.) |
| `list_timing_models` | List timing contracts (smooth_streaming, snappy_focus, etc.) |
| `explain_concept` | Explain a Charm component, pattern, or constraint |
| `list_deepwiki` | List available DeepWiki reference topics |

### MCP Resources

| Resource | Description |
|----------|-------------|
| `charmed://knowledge/components` | Component semantics (viewport, list, textarea, etc.) |
| `charmed://knowledge/patterns` | TUI architectural patterns |
| `charmed://knowledge/dependencies` | Charm library dependency graph |
| `charmed://archetypes` | TUI archetype catalog |
| `charmed://constraints` | Constraint rules |
| `charmed://timing-models` | Timing model contracts |
| `charmed://ir/schema` | TUI Structural IR JSON Schema |
| `charmed://deepwiki/{topic}` | DeepWiki reference (bubbles, bubbletea, lipgloss, etc.) |

## Agent

- **charm-project-scanner** — Scans a Go project and produces the TUI Structural IR (models, messages, cmds, component bindings, state ownership, async boundaries). Uses ATree for graph enrichment. Auto-invoked by skills that need IR data.

## Skills

Six skills activate automatically from conversational triggers:

1. **tui-scaffold** — "create a TUI for monitoring server logs" → generates complete Bubble Tea project
2. **tui-audit** — "audit my TUI project" → lints for anti-patterns and constraint violations
3. **tui-refactor** — "this model is too big" → proposes architectural decomposition
4. **tui-explain** — "what is viewport.Model?" → returns structured component knowledge
5. **tui-instrument** — "add timing traces" → generates telemetry hooks
6. **tui-ux-review** — "review my TUI's UX" → perceptual critique against timing contracts

## Knowledge Base

Embedded Charm ecosystem knowledge (no external API calls needed):

- **archetypes/** — 8 TUI archetypes (agent_console, streaming_chat, ide_workspace, monitoring_dashboard, cli_wizard, git_client, log_explorer, status_console)
- **constraints/** — 13 architectural lint rules (no_blocking_update, god_model, viewport_churn, etc.)
- **timing/** — 10 perceptual timing contracts (smooth_streaming, snappy_focus, responsive_input, etc.)
- **ontology/** — Component semantics and pattern relations for the Charm ecosystem
- **ir/** — TUI Structural IR JSON Schema
- **DeepWiki References/** — Full DeepWiki docs for bubbles, bubbletea, lipgloss, glamour, glow, gum, and more

## Installation

Add the SurvivalDev marketplace:

```bash
claude plugin marketplace add B-A-M-N/SurvivalDevPlugins
```

Install charmed:

```bash
claude plugin install charmed@survivaldev-plugins
```

Build the MCP binary:

```bash
bash scripts/build-charmed-mcp.sh
```

## Structure

```
charmed/
├── .claude-plugin/plugin.json    # Plugin manifest
├── .mcp.json                     # MCP server config (charmed-mcp stdio)
├── skills/                       # 6 skills
│   ├── tui-audit/SKILL.md
│   ├── tui-scaffold/SKILL.md
│   ├── tui-refactor/SKILL.md
│   ├── tui-explain/SKILL.md
│   ├── tui-instrument/SKILL.md
│   └── tui-ux-review/SKILL.md
├── commands/                     # 6 slash commands (/tui-audit, /tui-scaffold, etc.)
├── agents/                       # charm-project-scanner agent
├── scripts/                      # Build script for charmed-mcp binary
├── constraints/                  # Architectural lint rules
├── archetypes/                   # TUI archetype catalog
├── timing/                       # Perceptual timing contracts
├── ontology/                     # Component/pattern knowledge base
├── ir/                           # TUI Structural IR schema
├── DeepWiki References/          # Embedded DeepWiki docs
└── README.md                     # This file
```

## MCP Server Source

The `charmed-mcp` server is maintained at [`B-A-M-N/charmed-mcp`](https://github.com/B-A-M-N/charmed-mcp).

## License

MIT
