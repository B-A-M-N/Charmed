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

Interactive Commands
====================

Relevant source files

*   [README.md](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1)
    
*   [examples/commit.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.sh)
    
*   [examples/commit.tape](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.tape)
    
*   [examples/demo.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh)
    
*   [examples/demo.tape](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.tape)
    
*   [gum.go](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go)
    
*   [main.go](https://github.com/charmbracelet/gum/blob/30cc5169/main.go)
    

Purpose and Scope
-----------------

This document provides an overview of **gum's** interactive commands, which are commands that present Terminal User Interface (TUI) applications for user input and selection. These commands leverage the [Bubble Tea](https://github.com/charmbracelet/gum/blob/30cc5169/Bubble%20Tea)
 framework to render interactive interfaces where users navigate, select, and input data using keyboard controls.

Interactive commands form the core of `gum`'s functionality, enabling shell scripts to gather user input in a visually appealing and user-friendly manner. This page covers the common characteristics, execution patterns, and architectural components shared across all interactive commands.

For detailed documentation of individual interactive commands, see:

*   **Choose**: [Choose Command](https://deepwiki.com/charmbracelet/gum/4.1-choose-command)
     — Select one or multiple options from a list.
*   **Confirm**: [Confirm Command](https://deepwiki.com/charmbracelet/gum/4.2-confirm-command)
     — Yes/No prompts for binary decisions.
*   **Filter**: [Filter Command](https://deepwiki.com/charmbracelet/gum/4.3-filter-command)
     — Fuzzy searching and filtering lists.
*   **Input**: [Input Command](https://deepwiki.com/charmbracelet/gum/4.4-input-command)
     — Single-line text input (including password mode).
*   **Write**: [Write Command](https://deepwiki.com/charmbracelet/gum/4.5-write-command)
     — Multi-line text input for long-form content.
*   **File**: [File Command](https://deepwiki.com/charmbracelet/gum/4.6-file-command)
     — Interactive file and directory selection.
*   **Table**: [Table Command](https://deepwiki.com/charmbracelet/gum/4.7-table-command)
     — Display and select from tabular CSV data.
*   **Pager**: [Pager Command](https://deepwiki.com/charmbracelet/gum/4.8-pager-command)
     — Scrollable content viewer with search.

**Sources:** [README.md11-14](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L11-L14)
 [README.md145-160](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L145-L160)
 [gum.go25-222](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L222)

Interactive Command Catalog
---------------------------

Gum provides eight interactive commands, each serving a specific input or selection use case:

| Command | Purpose | Primary Bubble Component | File Pointer |
| --- | --- | --- | --- |
| `choose` | Select items from a list | `viewport`, `paginator` | [gum.go35-46](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L35-L46) |
| `confirm` | Yes/No confirmation prompt | Custom model | [gum.go48-60](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L48-L60) |
| `filter` | Fuzzy search through items | `textinput`, `viewport` | [gum.go76-87](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L76-L87) |
| `input` | Single-line text input | `textinput` | [gum.go94-102](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L94-L102) |
| `write` | Multi-line text editor | `textarea` | [gum.go208-222](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L208-L222) |
| `file` | File system navigation | `filepicker` | [gum.go62-74](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L62-L74) |
| `table` | Tabular data selection | `table` | [gum.go183-198](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L183-L198) |
| `pager` | Scrollable content viewer | `viewport` | [gum.go122-139](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L122-L139) |

**Sources:** [gum.go25-222](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L222)
 [README.md145-160](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L145-L160)

Command Registration and Routing
--------------------------------

All interactive commands are registered in the `Gum` struct using [Kong](https://github.com/charmbracelet/gum/blob/30cc5169/Kong)
 CLI parser annotations. The `main` function in `main.go` parses the command line and routes execution to the appropriate command's `Run()` method.

### CLI to Code Entity Mapping

The following diagram bridges the natural language CLI commands to their specific Go `Options` structs and logic packages.

**Sources:** [gum.go25-222](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L222)
 [main.go46-75](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L46-L75)

Interactive Command Execution Pattern
-------------------------------------

All interactive commands follow a consistent execution pattern centered on the Bubble Tea framework's Model-View-Update (MVU) architecture.

**Sources:** [main.go76-90](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L90)
 [gum.go48-55](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L48-L55)

I/O Stream Separation
---------------------

Interactive commands strictly separate UI rendering from data output. This is critical for shell composability.

*   **Stderr**: Used for the TUI rendering (the "glamorous" interface). This ensures that the interface doesn't pollute the data pipeline.
*   **Stdout**: Used exclusively for the final result (e.g., the selected file path, the text input, or the chosen option).
*   **Exit Codes**: Used to signal success, user cancellation (`Ctrl+C`), or timeouts.

**Sources:** [main.go32](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L32-L32)
 [main.go76-90](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L90)
 [README.md34-36](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L34-L36)

Common Configuration Pattern
----------------------------

Each command uses a consistent priority resolution for configuration:

1.  **CLI Flags**: Highest priority (e.g., `--height 10`).
2.  **Environment Variables**: Prefix `GUM_<COMMAND>_` (e.g., `GUM_INPUT_WIDTH=80`).
3.  **Struct Tag Defaults**: Built-in defaults defined in the code.

**Sources:** [README.md161-189](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L161-L189)
 [main.go56-74](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L56-L74)

Shell Script Integration Examples
---------------------------------

Interactive commands are designed to be used in pipes and subshells.

### Capture Pattern

### Conditional Pattern

**Sources:** [examples/demo.sh4-27](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh#L4-L27)
 [examples/commit.sh17-28](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.sh#L17-L28)
 [README.md50-54](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L50-L54)

Error and Termination Handling
------------------------------

The `main.go` entry point centralizes the handling of termination states for all interactive commands:

*   **`tea.ErrInterrupted`**: Triggered by `Ctrl+C`. The application exits with status code `130` (`exit.StatusAborted`).
*   **`tea.ErrProgramKilled`**: Triggered when a `--timeout` is reached. The application prints "timed out" to `stderr` and exits with status `124` (`exit.StatusTimeout`).
*   **`exit.ErrExit`**: A custom error type used to force specific exit codes (e.g., in `confirm`).

**Sources:** [main.go76-90](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L90)
 [gum.go48-55](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L48-L55)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Interactive Commands](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#interactive-commands)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#purpose-and-scope)
    
*   [Interactive Command Catalog](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#interactive-command-catalog)
    
*   [Command Registration and Routing](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#command-registration-and-routing)
    
*   [CLI to Code Entity Mapping](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#cli-to-code-entity-mapping)
    
*   [Interactive Command Execution Pattern](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#interactive-command-execution-pattern)
    
*   [I/O Stream Separation](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#io-stream-separation)
    
*   [Common Configuration Pattern](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#common-configuration-pattern)
    
*   [Shell Script Integration Examples](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#shell-script-integration-examples)
    
*   [Capture Pattern](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#capture-pattern)
    
*   [Conditional Pattern](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#conditional-pattern)
    
*   [Error and Termination Handling](https://deepwiki.com/charmbracelet/gum/4-interactive-commands#error-and-termination-handling)
