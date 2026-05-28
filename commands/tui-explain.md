---
name: charmed:explain
description: "Graph-driven explanation of Charm ecosystem concepts — components, patterns, message flows, architectural tradeoffs."
argument-hint: "<concept or question> [--depth quick|deep] [--with-code]"
allowed-tools: [Read, Grep, Glob, Bash]
---

Invoke the tui-explain skill to get a semantic explanation of a Charm ecosystem concept.

## Instructions

1. Classify the user's query using the categories in `skills/tui-explain/SKILL.md`.
2. Read `ontology/knowledge.yaml` for the relevant knowledge base section.
3. If ATree DB is available, enrich with graph queries.
4. Present the explanation at the requested depth.

## Usage

```
/charm-tui:explain how does Bubble Tea Update loop work
/charm-tui:explain what is viewport churn --depth deep --with-code
/charm-tui:explain what's the difference between tea.Batch and tea.Sequence
```
