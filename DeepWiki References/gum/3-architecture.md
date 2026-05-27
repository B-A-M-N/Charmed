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

Architecture
============

Relevant source files

*   [go.mod](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/gum/blob/30cc5169/go.sum)
    
*   [gum.go](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go)
    
*   [main.go](https://github.com/charmbracelet/gum/blob/30cc5169/main.go)
    

Gum is structured as a layered command-line application that provides 15+ specialized subcommands for creating interactive terminal UIs and formatting text. The architecture follows a Model-View-Update (MVU) pattern for interactive commands via the Bubble Tea framework, while non-interactive commands perform direct stdin/stdout transformations.

Layered Architecture
--------------------

Gum's architecture consists of five primary layers that separate CLI parsing, command logic, TUI rendering, and shared utilities.

**Diagram: Layered Architecture with Code Entities**

Sources: [main.go31-91](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L31-L91)
 [gum.go25-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L228)
 [go.mod7-24](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L7-L24)

For details on the bootstrap process, see [Application Entry Point](https://deepwiki.com/charmbracelet/gum/3.1-application-entry-point)
.

Command Registry
----------------

The `Gum` struct in `gum.go` serves as the central registry for the application. Each field in the struct represents a subcommand, decorated with `kong` tags that define CLI behavior, help text, and environment variable mapping.

**Diagram: Command Registry Structure**

Sources: [gum.go25-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L228)

Each command implementation follows a consistent pattern: an `Options` struct containing flags and a `Run()` method that executes the command's logic. For details on these patterns, see [Command Patterns](https://deepwiki.com/charmbracelet/gum/3.3-command-patterns)
.

Execution Patterns
------------------

Gum implements three distinct execution models based on the nature of the command:

1.  **Interactive (TUI)**: Uses `bubbletea` to render an interface to `os.Stderr` while writing the final selection to `os.Stdout`. This ensures that the UI does not pollute the data pipeline in shell scripts.
2.  **Non-Interactive (Formatting)**: Directly transforms input (from arguments or stdin) and writes the result to `os.Stdout` using `lipgloss` or `glamour`.
3.  **Background (Execution)**: The `spin` command runs a subprocess while maintaining a TUI spinner on `os.Stderr`.

For details on the frameworks powering these patterns, see [Framework Components](https://deepwiki.com/charmbracelet/gum/3.2-framework-components)
.

Framework Dependencies
----------------------

Gum is built upon the Charmbracelet ecosystem, leveraging several key libraries:

*   **Bubble Tea**: The runtime for interactive commands, handling the terminal state and event loop.
*   **Bubbles**: Reusable TUI components like `textinput`, `viewport`, and `spinner` used across different commands.
*   **Lip Gloss**: The styling engine used for colors, borders, and layouts.
*   **Kong**: Used for declarative CLI argument parsing and help generation.

Sources: [go.mod7-24](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L7-L24)
 [main.go47-75](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L47-L75)

Internal Utilities
------------------

The application relies on internal utility packages to handle common terminal operations consistently:

*   **`internal/stdin`**: Handles reading data from pipes and standard input.
*   **`internal/exit`**: Provides standard exit codes (e.g., `StatusAborted` for `Ctrl+C`).
*   **`internal/timeout`**: Manages execution deadlines for interactive prompts.

Sources: [main.go12](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L12-L12)
 [main.go77-89](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L77-L89)

* * *

*   [Application Entry Point](https://deepwiki.com/charmbracelet/gum/3.1-application-entry-point)
     — Detail how main.go bootstraps the application.
*   [Framework Components](https://deepwiki.com/charmbracelet/gum/3.2-framework-components)
     — Explain the Charmbracelet ecosystem dependencies.
*   [Command Patterns](https://deepwiki.com/charmbracelet/gum/3.3-command-patterns)
     — Document common patterns across commands.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Architecture](https://deepwiki.com/charmbracelet/gum/3-architecture#architecture)
    
*   [Layered Architecture](https://deepwiki.com/charmbracelet/gum/3-architecture#layered-architecture)
    
*   [Command Registry](https://deepwiki.com/charmbracelet/gum/3-architecture#command-registry)
    
*   [Execution Patterns](https://deepwiki.com/charmbracelet/gum/3-architecture#execution-patterns)
    
*   [Framework Dependencies](https://deepwiki.com/charmbracelet/gum/3-architecture#framework-dependencies)
    
*   [Internal Utilities](https://deepwiki.com/charmbracelet/gum/3-architecture#internal-utilities)
