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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1)
    
*   [examples/commit.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.sh)
    
*   [examples/commit.tape](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.tape)
    
*   [examples/demo.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh)
    
*   [examples/demo.tape](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.tape)
    
*   [go.mod](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod)
    
*   [gum.go](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go)
    
*   [main.go](https://github.com/charmbracelet/gum/blob/30cc5169/main.go)
    

Gum is a command-line tool for glamorous shell scripts. It provides ready-to-use utilities for creating interactive Terminal User Interfaces (TUIs) directly from shell scripts without writing Go code.

Purpose and Design Philosophy
-----------------------------

Gum enables shell script authors to leverage the [Bubble Tea](https://github.com/charmbracelet/gum/blob/30cc5169/Bubble%20Tea)
 TUI framework and [Lip Gloss](https://github.com/charmbracelet/gum/blob/30cc5169/Lip%20Gloss)
 styling library through simple command-line commands. The tool follows Unix philosophy principles:

*   **Composability**: Commands read from stdin and write to stdout, enabling pipeline integration.
*   **Single Responsibility**: Each command does one thing well (input, selection, filtering, styling).
*   **Data/UI Separation**: Interactive commands render UI to `stderr` while writing data results to `stdout`, preserving shell script composability [main.go32-33](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L32-L33)
    
*   **Configuration Flexibility**: Commands accept configuration through CLI flags, environment variables (following the `GUM_COMMAND_OPTION` pattern), and sensible defaults [README.md163-186](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L163-L186)
    

This design makes Gum commands highly composable in shell pipelines and scripts. For example:

Sources: [README.md11-14](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L11-L14)
 [main.go32-33](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L32-L33)
 [README.md393-459](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L393-L459)

System Architecture
-------------------

### Layered Architecture

**Layered Architecture of Gum**

Gum follows a layered architecture:

1.  **User Interface Layer**: Users invoke commands via CLI or shell scripts.
2.  **Application Entry Point**: `main.go` bootstraps the application, initializes the color profile using `termenv`, and uses `kong` to parse commands into the `Gum` struct [main.go31-75](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L31-L75)
    
3.  **Command Layer**: Each command is implemented as a separate package with an `Options` struct containing configuration and logic [gum.go25-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L228)
    
4.  **Framework Layer**: Commands depend on Charmbracelet ecosystem components (`bubbletea`, `bubbles`, `lipgloss`) for TUI and styling functionality [go.mod9-13](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L9-L13)
    

Sources: [main.go31-91](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L31-L91)
 [gum.go24-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L24-L228)
 [go.mod5-24](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L5-L24)

### Command Structure Pattern

**Command Structure Pattern**

Each command follows a consistent pattern:

*   Implemented as a separate package (e.g., `github.com/charmbracelet/gum/choose`) [gum.go6-21](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L6-L21)
    
*   Exports an `Options` struct with configuration fields tagged for `kong` parsing [gum.go25-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L228)
    
*   The `kong` context executes the command via the `ctx.Run()` call in `main.go` [main.go76](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L76)
    
*   Returns an error (handled specifically for `exit.ErrExit` to set exit codes) [main.go77-80](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L77-L80)
    

Sources: [gum.go24-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L24-L228)
 [main.go76-90](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L90)

### Command Execution Patterns

**Three Execution Patterns**

Gum commands follow three distinct execution patterns:

1.  **Interactive Commands** (`choose`, `filter`, `input`, `write`, `file`, `pager`, `confirm`, `table`):
    
    *   Initialize a Bubble Tea model.
    *   Run a `tea.Program` which renders the TUI to `stderr` [main.go32](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L32-L32)
        
    *   Write the final user selection to `stdout` upon completion [README.md34-36](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L34-L36)
        
2.  **Non-Interactive Commands** (`style`, `format`, `join`, `log`):
    
    *   Read input from stdin or arguments.
    *   Apply `lipgloss` styles or `glamour` formatting [gum.go89-92](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L89-L92)
        
    *   Write formatted output directly to `stdout`.
3.  **Background Commands** (`spin`):
    
    *   Start a background subprocess.
    *   Display a `spinner` bubble to `stderr` while the process runs [gum.go141-156](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L141-L156)
        
    *   Capture and relay the process output to `stdout`.

Sources: [main.go76-90](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L90)
 [README.md194-390](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L194-L390)
 [gum.go24-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L24-L228)

Command Categories
------------------

Gum provides 15+ commands organized by functionality:

### Interactive Selection Commands

| Command | Description | Key Features | Example Usage |
| --- | --- | --- | --- |
| `choose` | Select from a list of options | Multi-select, pagination | `gum choose "A" "B"` [gum.go46](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L46-L46) |
| `filter` | Fuzzy search through options | Exact/fuzzy matching | `gum filter < list.txt` [gum.go87](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L87-L87) |
| `confirm` | Yes/no confirmation prompt | Exit code 0/1 | `gum confirm "Sure?"` [gum.go60](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L60-L60) |
| `table` | Select from tabular CSV data | Row selection, borders | `gum table < data.csv` [gum.go183](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L183-L183) |
| `file` | Navigate and select files | File manager interface | `gum file .` [gum.go74](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L74-L74) |

### Input Commands

| Command | Description | Key Features | Example Usage |
| --- | --- | --- | --- |
| `input` | Single-line text input | Password mode, placeholders | `gum input --password` [gum.go102](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L102-L102) |
| `write` | Multi-line text input | Editor-like behavior | `gum write > note.txt` [gum.go221](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L221-L221) |

### Formatting and Display Commands

| Command | Description | Key Features | Example Usage |
| --- | --- | --- | --- |
| `style` | Apply colors/borders to text | Lipgloss integration | `gum style --foreground 212 "Hi"` [gum.go181](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L181-L181) |
| `format` | Render markdown/code/templates | Glamour rendering | `gum format -- "# Hello"` [gum.go92](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L92-L92) |
| `join` | Combine text horizontally/vertically | Layout management | `gum join --horizontal "A" "B"` [gum.go120](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L120-L120) |
| `pager` | Scroll through text content | Line numbers, vim keys | `gum pager < file.txt` [gum.go139](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L139-L139) |
| `log` | Structured logging | Log levels, timestamps | `gum log --level info "msg"` [gum.go114](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L114-L114) |

Sources: [gum.go24-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L24-L228)
 [README.md145-160](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L145-L160)

Data Flow and I/O Patterns
--------------------------

Gum follows Unix conventions to ensure it works well in scripts:

1.  **Configuration Cascade**: CLI flags > Environment Variables (`GUM_COMMAND_OPTION`) > Defaults defined in `main.go` [main.go56-74](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L56-L74)
    
2.  **Output Separation**:
    *   **stdout**: Used for the "result" of the command (e.g., the string selected in `choose`, the text typed in `input`).
    *   **stderr**: Used for the TUI interface itself, spinners, and error messages.
3.  **Exit Codes**:
    *   `0`: Success or affirmative (`confirm`).
    *   `1`: Error or negative (`confirm`).
    *   `130`: User aborted (Ctrl+C) [main.go81-83](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L81-L83)
        
    *   `124`: Command timed out [main.go84-87](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L84-L87)
        

Sources: [main.go76-90](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L90)
 [README.md163-189](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L163-L189)
 [gum.go48-60](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L48-L60)

Key Dependencies
----------------

Gum is built on the Charmbracelet TUI stack:

*   **Bubble Tea** (`github.com/charmbracelet/bubbletea`): The underlying TUI framework [go.mod10](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L10-L10)
    
*   **Bubbles** (`github.com/charmbracelet/bubbles`): Reusable UI components like the text input and spinner [go.mod9](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L9-L9)
    
*   **Lip Gloss** (`github.com/charmbracelet/lipgloss`): Styling and layout engine [go.mod12](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L12-L12)
    
*   **Kong** (`github.com/alecthomas/kong`): Command-line parser that maps CLI flags to Go structs [go.mod7](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L7-L7)
    

Sources: [go.mod5-24](https://github.com/charmbracelet/gum/blob/30cc5169/go.mod#L5-L24)
 [main.go10-13](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L10-L13)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/gum#overview)
    
*   [Purpose and Design Philosophy](https://deepwiki.com/charmbracelet/gum#purpose-and-design-philosophy)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/gum#system-architecture)
    
*   [Layered Architecture](https://deepwiki.com/charmbracelet/gum#layered-architecture)
    
*   [Command Structure Pattern](https://deepwiki.com/charmbracelet/gum#command-structure-pattern)
    
*   [Command Execution Patterns](https://deepwiki.com/charmbracelet/gum#command-execution-patterns)
    
*   [Command Categories](https://deepwiki.com/charmbracelet/gum#command-categories)
    
*   [Interactive Selection Commands](https://deepwiki.com/charmbracelet/gum#interactive-selection-commands)
    
*   [Input Commands](https://deepwiki.com/charmbracelet/gum#input-commands)
    
*   [Formatting and Display Commands](https://deepwiki.com/charmbracelet/gum#formatting-and-display-commands)
    
*   [Data Flow and I/O Patterns](https://deepwiki.com/charmbracelet/gum#data-flow-and-io-patterns)
    
*   [Key Dependencies](https://deepwiki.com/charmbracelet/gum#key-dependencies)
