Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/gum](https://github.com/charmbracelet/gum "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 2 May 2026 ([30cc51](https://github.com/charmbracelet/gum/commits/30cc5169)
)

*   [Overview](https://deepwiki.com/charmbracelet/gum/1-overview)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/gum/2-getting-started)
    
*   [Architecture](https://deepwiki.com/charmbracelet/gum/3-architecture)
    
*   [Application Entry Point](https://deepwiki.com/charmbracelet/gum/3.1-application-entry-point)
    
*   [Framework Components](https://deepwiki.com/charmbracelet/gum/3.2-framework-components)
    
*   [Command Patterns](https://deepwiki.com/charmbracelet/gum/3.3-command-patterns)
    
*   [Interactive Commands](https://deepwiki.com/charmbracelet/gum/4-interactive-commands)
    
*   [Choose Command](https://deepwiki.com/charmbracelet/gum/4.1-choose-command)
    
*   [Confirm Command](https://deepwiki.com/charmbracelet/gum/4.2-confirm-command)
    
*   [Filter Command](https://deepwiki.com/charmbracelet/gum/4.3-filter-command)
    
*   [Input Command](https://deepwiki.com/charmbracelet/gum/4.4-input-command)
    
*   [Write Command](https://deepwiki.com/charmbracelet/gum/4.5-write-command)
    
*   [File Command](https://deepwiki.com/charmbracelet/gum/4.6-file-command)
    
*   [Table Command](https://deepwiki.com/charmbracelet/gum/4.7-table-command)
    
*   [Pager Command](https://deepwiki.com/charmbracelet/gum/4.8-pager-command)
    
*   [Progress and Background Commands](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands)
    
*   [Spin Command](https://deepwiki.com/charmbracelet/gum/5.1-spin-command)
    
*   [Formatting Commands](https://deepwiki.com/charmbracelet/gum/6-formatting-commands)
    
*   [Style Command](https://deepwiki.com/charmbracelet/gum/6.1-style-command)
    
*   [Format Command](https://deepwiki.com/charmbracelet/gum/6.2-format-command)
    
*   [Join Command](https://deepwiki.com/charmbracelet/gum/6.3-join-command)
    
*   [Log Command](https://deepwiki.com/charmbracelet/gum/6.4-log-command)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/gum/7-development-guide)
    
*   [Building and Testing](https://deepwiki.com/charmbracelet/gum/7.1-building-and-testing)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/gum/7.2-dependency-management)
    
*   [CI/CD Pipeline](https://deepwiki.com/charmbracelet/gum/7.3-cicd-pipeline)
    
*   [Shell Completion](https://deepwiki.com/charmbracelet/gum/7.4-shell-completion)
    
*   [Internal Utilities](https://deepwiki.com/charmbracelet/gum/8-internal-utilities)
    
*   [Glossary](https://deepwiki.com/charmbracelet/gum/9-glossary)
    

Menu

Glossary
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1)
    
*   [confirm/command.go](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/command.go)
    
*   [confirm/confirm.go](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/confirm.go)
    
*   [confirm/options.go](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/options.go)
    
*   [examples/commit.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.sh)
    
*   [examples/commit.tape](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.tape)
    
*   [examples/demo.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh)
    
*   [examples/demo.tape](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.tape)
    
*   [filter/filter.go](https://github.com/charmbracelet/gum/blob/30cc5169/filter/filter.go)
    
*   [filter/filter\_test.go](https://github.com/charmbracelet/gum/blob/30cc5169/filter/filter_test.go)
    
*   [go.mod](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod)
    
*   [gum.go](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go)
    
*   [main.go](https://github.com/charmbracelet/gum/blob/30cc5169/main.go)
    
*   [spin/command.go](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go)
    
*   [spin/options.go](https://github.com/charmbracelet/gum/blob/30cc5169/spin/options.go)
    
*   [spin/spin.go](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go)
    

This glossary defines technical terms, framework concepts, and codebase-specific jargon used within the `gum` project. It serves as a reference for onboarding engineers to understand the interaction between the Charmbracelet ecosystem and the `gum` command-line interface.

Charmbracelet Ecosystem
-----------------------

`gum` is built upon several foundational libraries from the Charmbracelet ecosystem.

| Term | Definition |
| --- | --- |
| **Bubble Tea** | The [The Elm Architecture (TEA)](https://guide.elm-lang.org/architecture/)<br> inspired framework used for building the TUI. It manages the application lifecycle through `Init`, `Update`, and `View` functions. [go.mod10](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L10-L10) |
| **Bubbles** | A library of reusable TUI components (widgets) for Bubble Tea, such as `spinner`, `textinput`, and `viewport`. [go.mod9](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L9-L9) |
| **Lip Gloss** | A terminal styling library used to define colors, borders, padding, and layout using a declarative API. [go.mod12](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L12-L12) |
| **Glamour** | A markdown rendering engine for the terminal, used by the `format` command. [go.mod11](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L11-L11) |

TUI Patterns & Lifecycle
------------------------

### Model-Update-View (TEA)

The core pattern for all interactive commands in `gum`.

*   **Model**: A `struct` that holds the application state (e.g., cursor position, user input, filtered matches). [filter/filter.go126-158](https://github.com/charmbracelet/gum/blob/30cc5169/filter/filter.go#L126-L158)
    
*   **Update**: A function that handles incoming `tea.Msg` (key presses, window size changes, custom messages) and returns a modified model and an optional command (`tea.Cmd`). [spin/spin.go177-200](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L177-L200)
    
*   **View**: A function that returns a string representation of the UI based on the current state of the model. [confirm/confirm.go133-168](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/confirm.go#L133-L168)
    

### Commands and Messages

*   **tea.Cmd**: A function that performs an I/O operation and returns a `tea.Msg`. In `gum`, this is often used to run background processes. [spin/spin.go68-133](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L68-L133)
    
*   **tea.Msg**: An interface representing an event. Common types include `tea.KeyMsg` for keyboard input and `tea.WindowSizeMsg` for terminal resizing. [confirm/confirm.go100-103](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/confirm.go#L100-L103)
    

### Code Entity Mapping: Lifecycle

The following diagram bridges the natural language concept of a "command execution" to the specific code entities involved in the Bubble Tea lifecycle within `gum`.

**Diagram: Bubble Tea Lifecycle in gum (e.g., Spin Command)**

**Sources:** [spin/spin.go33-49](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L33-L49)
 [spin/spin.go68-132](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L68-L132)
 [spin/spin.go142-147](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L142-L147)
 [spin/spin.go177-200](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L177-L200)
 [spin/command.go17-45](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L17-L45)

* * *

CLI & Framework Jargon
----------------------

### Kong

The command-line parser used by `gum`. It maps CLI arguments and environment variables directly to Go structs. [main.go47-75](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L47-L75)

*   **Options Struct**: A struct (e.g., `confirm.Options`) annotated with `kong` tags that defines the flags and arguments for a specific command. [confirm/options.go10-25](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/options.go#L10-L25)
    
*   **Embed/Prefix**: Kong features used to reuse styling flags across different commands (e.g., `prompt.Style`). [confirm/options.go17-21](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/options.go#L17-L21)
    

### PTY (Pseudo-Terminal)

A pair of virtual devices used to trick a process into thinking it is running in a real terminal. `gum` uses `xpty` to capture output from background commands in the `spin` command while maintaining TTY-like behavior. [spin/spin.go88-114](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L88-L114)

### Code Entity Mapping: CLI Parsing

This diagram shows how `main.go` routes CLI inputs to specific command implementations.

**Diagram: CLI Routing Entity Map**

**Sources:** [main.go46-75](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L46-L75)
 [gum.go25-60](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L60)
 [confirm/command.go18-71](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/command.go#L18-L71)
 [confirm/options.go10-25](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/options.go#L10-L25)

* * *

gum-Specific Concepts
---------------------

| Term | Definition |
| --- | --- |
| **Fuzzy Matching** | The algorithm used in the `filter` command (via `github.com/sahilm/fuzzy`) to find approximate string matches. [filter/filter.go24](https://github.com/charmbracelet/gum/blob/30cc5169/filter/filter.go#L24-L24)<br> [filter/filter.go180-231](https://github.com/charmbracelet/gum/blob/30cc5169/filter/filter.go#L180-L231) |
| **Indicator** | A visual marker (e.g., `>`) used in commands like `choose` and `filter` to show the currently focused item. [filter/filter.go191-197](https://github.com/charmbracelet/gum/blob/30cc5169/filter/filter.go#L191-L197) |
| **Aborted Status** | An exit code (130) returned when a user cancels a command via `Ctrl+C`. [main.go82](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L82-L82)<br> [internal/exit/exit.go](https://github.com/charmbracelet/gum/blob/30cc5169/internal/exit/exit.go) |
| **Styles Struct** | A common utility struct used in `gum` to wrap Lip Gloss styles, allowing them to be configured via Kong flags and environment variables. [confirm/options.go17-21](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/options.go#L17-L21) |
| **Multi-Writer** | A technique used in the `spin` command to simultaneously capture stdout/stderr to internal buffers while optionally displaying them to the user. [spin/spin.go85-86](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L85-L86) |

**Sources:** [main.go1-91](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L1-L91)
 [gum.go1-22](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L1-L22)
 [filter/filter.go1-250](https://github.com/charmbracelet/gum/blob/30cc5169/filter/filter.go#L1-L250)
 [spin/spin.go1-200](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L1-L200)
 [confirm/command.go1-71](https://github.com/charmbracelet/gum/blob/30cc5169/confirm/command.go#L1-L71)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Glossary](https://deepwiki.com/charmbracelet/gum/9-glossary#glossary)
    
*   [Charmbracelet Ecosystem](https://deepwiki.com/charmbracelet/gum/9-glossary#charmbracelet-ecosystem)
    
*   [TUI Patterns & Lifecycle](https://deepwiki.com/charmbracelet/gum/9-glossary#tui-patterns-lifecycle)
    
*   [Model-Update-View (TEA)](https://deepwiki.com/charmbracelet/gum/9-glossary#model-update-view-tea)
    
*   [Commands and Messages](https://deepwiki.com/charmbracelet/gum/9-glossary#commands-and-messages)
    
*   [Code Entity Mapping: Lifecycle](https://deepwiki.com/charmbracelet/gum/9-glossary#code-entity-mapping-lifecycle)
    
*   [CLI & Framework Jargon](https://deepwiki.com/charmbracelet/gum/9-glossary#cli-framework-jargon)
    
*   [Kong](https://deepwiki.com/charmbracelet/gum/9-glossary#kong)
    
*   [PTY (Pseudo-Terminal)](https://deepwiki.com/charmbracelet/gum/9-glossary#pty-pseudo-terminal)
    
*   [Code Entity Mapping: CLI Parsing](https://deepwiki.com/charmbracelet/gum/9-glossary#code-entity-mapping-cli-parsing)
    
*   [gum-Specific Concepts](https://deepwiki.com/charmbracelet/gum/9-glossary#gum-specific-concepts)
