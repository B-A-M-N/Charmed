---
name: tui-explain
description: "Knowledge corpus queries — lookup components, patterns, constraints, and timing contracts from the embedded Charm ontology. Structured semantic answers sourced from the IR schema, constraint definitions, timing contracts, and component knowledge base. Trigger when: user asks 'what is X', 'how does Y work', 'explain Z', 'what are the constraints on', 'what timing contracts apply to', or any question about Charm ecosystem architecture."
argument-hint: "<concept or question> [--depth quick|deep] [--with-code] [--with-graph]"
allowed-tools: [Read, Grep, Glob, Bash]
---

# tui-explain — Charm Ecosystem Knowledge Engine

Semantic explanation of Bubble Tea, Bubbles, Lip Gloss, and related Charm ecosystem concepts. Combines structured ontology knowledge from `knowledge.yaml` with live ATree graph queries for precise, context-aware answers. This is NOT a document search — it is architectural reasoning grounded in code graph evidence.

## Inputs

| Input | Required | Source |
|-------|----------|--------|
| Query/question | Yes | User's natural language |
| Ontology | Yes | `ontology/knowledge.yaml` |
| ATree DB path | Yes | Detected from project context or `ATREE_DB_PATH` |
| DeepWiki corpus | Yes | `../../deepwiki-charmbracelet/` if available |
| Depth | No | `--depth quick` (1-2 sentences) or `deep` (full explanation) |
| With code | No | `--with-code` — include code examples |
| With graph | No | `--with-graph` — include ATree evidence paths |

## Execution Pipeline

### Phase 1: Query Classification

Classify the user's query into one of these categories:

| Category | Example Queries | Primary Source |
|----------|----------------|----------------|
| Component | "What is viewport.Model?" | `ontology/knowledge.yaml` component defs |
| Pattern | "How does streaming_console work?" | Pattern relations in `knowledge.yaml` |
| Comparison | "viewport vs list for log display?" | Component specs + archetype mapping |
| Architecture | "How should I structure a multi-pane app?" | Ontology + archetypes |
| Debugging | "Why is my TUI flickering?" | Constraints + anti-patterns in `knowledge.yaml` |
| Performance | "Why is View() slow?" | Timing models + constraint rules |
| Rendering | "How does lipgloss layout work?" | Component specs + pairing rules |
| Message flow | "How do Cmd callbacks reach Update?" | Bubble Tea lifecycle + message routing |
| Dependency | "What does gum depend on?" | Charm library dependency graph |

### Phase 2: Knowledge Assembly

1. **Ontology lookup**: Find the concept in `ontology/knowledge.yaml`. Extract component semantics, constraints, pairings.
2. **ATree context query**: If a project context exists, query the ATree DB:
   - `atree context <symbol>` — where is this used in the project?
   - `atree explain_symbol <name>` — graph-level explanation
   - `atree trace_call_path <from> <to>` — how does control flow reach this?
3. **Cross-reference**: Link the concept to:
   - Related components (from `common_pairings` in knowledge.yaml)
   - Constraints that govern it (from `constraints.yaml`)
   - Timing models that apply (from `timing/timing-models.yaml`)
   - Anti-patterns that misuse it (from `knowledge.yaml` anti_patterns)

### Phase 3: Answer Structuring

Structure the answer based on category:

**Component explanation:**
```
## <Component.Name> (<Package>)

<One-line description from knowledge.yaml>

### Owned State
- <field>: <type> — <purpose>

### Lifecycle
1. <phase>: <what happens>
2. ...

### Constraints
- ⚠️ <constraint> — <what breaks if violated>

### Common With
- Pairs with: <component> (from common_pairings)
- Pattern: <pattern> (from pattern_relations)

### Anti-Patterns
- 🚫 <anti-pattern description>
```

**Pattern explanation:**
```
## <Pattern.Name>

<Description>

### Components Used
| Component | Role |
|-----------|------|
| <comp>   | <role> |

### Timing Model
<Which timing model this pattern obeys>

### Failure Modes
- <mode>: <what goes wrong and how to detect it>
```

**Comparison explanation:**
| Aspect | <Option A> | <Option B> |
|--------|-----------|-----------|
| Use case | ... | ... |
| State model | ... | ... |
| Pairings | ... | ... |
| Timing | ... | ... |
| When to choose | ... | ... |

### Phase 4: Evidence and Examples

If `--with-code` is set:
1. Use ATree to find real-world code examples from the indexed Charm repos:
   - `atree code_search "<component_name>" --lang go`
   - Extract relevant snippets (10-20 lines max per example)
2. Annotate code with inline comments explaining key lines

If `--with-graph` is set:
1. Use ATree to show graph relationships:
   - `atree evidence_path <concept> <related_concept>`
   - Render the path as a text graph:
     ```
     Model ──owns──► viewport.Model ──uses──► lipgloss.Style
       │                                        │
       └──handles──► WindowSizeMsg ◄─────────────┘
     ```

## Explanation Principles

1. **Always explain the WHY, not just the WHAT.** Don't just say "viewport has SetContent" — explain that SetContent replaces scroll position, so it must be diff-checked.
2. **Always mention the Constraints.** Every component/class comes with constraint rules from `constraints.yaml`. Surface them naturally.
3. **Cross-link to patterns.** When explaining a component, mention which patterns use it and why.
4. **Be opinionated about tradeoffs.** Say "use viewport for X, use list for Y" with reasoning.
5. **Ground in code, not docs.** Use ATree symbol queries to back up explanations with real references.

## Common Explanations (Pre-compiled)

### Bubble Tea Lifecycle
```
Init() → Cmd → [async work] → tea.Msg → Update() → (Model, Cmd) → View() → render
                                  ↑                              │
                                  └────────── loop ◄─────────────┘
```

### Message Flow
```
tea.Cmd callback → tea.Msg → Program.Send() → root.Update() → child.Update() → child.View()
```

### Viewport + Key + Help Pairing
```
viewport.Model (scrollable content)
    ↑ focused by
key.Binding (what keys scroll)
    ↑ described by
help.Model (visible keyboard help)
```

### Lip Gloss Style Caching Rule
```
Styles are pure functions of data → compute ONCE when data changes, not every View() frame.
Cache as struct fields: type Model struct { style lipgloss.Style }
Recompute in Update() only when the underlying data changes.
```

## Environment

- `ATREE_DB_PATH`: ATree database for code graph queries
- `DEEPWIKI_PATH`: Path to DeepWiki corpus (default: `../../deepwiki-charmbracelet/`)
