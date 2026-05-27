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

Formatting Commands
===================

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

Formatting commands are a category of non-interactive gum commands that transform, style, and render text output for display in the terminal. Unlike interactive commands (see [Interactive Commands](https://deepwiki.com/charmbracelet/gum/4-interactive-commands)
) that present TUI interfaces and run event loops, formatting commands operate as simple Unix filters: they read input from stdin or command-line arguments, apply transformations, and write the result to stdout.

This page provides an overview of the four formatting commands in gum:

*   **style** - Apply Lipgloss styling (colors, borders, padding, margins) to text (see [Style Command](https://deepwiki.com/charmbracelet/gum/6.1-style-command)
    )
*   **format** - Render markdown, code, templates, and emojis (see [Format Command](https://deepwiki.com/charmbracelet/gum/6.2-format-command)
    )
*   **join** - Combine text horizontally or vertically with alignment (see [Join Command](https://deepwiki.com/charmbracelet/gum/6.3-join-command)
    )
*   **log** - Structured logging with levels and timestamps (see [Log Command](https://deepwiki.com/charmbracelet/gum/6.4-log-command)
    )

For commands that display progress indicators during background execution, see [Progress and Background Commands](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands)
.

**Sources:** [README.md145-160](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L145-L160)
 [gum.go89-217](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L89-L217)

Architectural Characteristics
-----------------------------

Formatting commands differ fundamentally from interactive commands in their execution model:

| Aspect | Formatting Commands | Interactive Commands |
| --- | --- | --- |
| **TUI Framework** | No Bubbletea dependency | Uses Bubbletea program lifecycle |
| **User Input** | None (non-interactive) | Keyboard/mouse events |
| **Output Channel** | stdout only | stderr for UI, stdout for data |
| **Execution Pattern** | Synchronous transform | Async event loop |
| **Dependencies** | Lipgloss, Glamour, charmbracelet/log | Bubbletea, Bubbles, Lipgloss |

All formatting commands implement a simple pattern: they read data, apply transformations, and write results without blocking for user interaction. This makes them ideal for shell pipelines and scripting.

**Sources:** [gum.go158-217](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L158-L217)
 [README.md307-390](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L307-L390)

Command Registration and Execution
----------------------------------

### System to Code Mapping

The following diagram bridges the high-level formatting commands to their specific implementation packages and options structures within the codebase.

The `Gum` struct in [gum.go25-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L228)
 registers each formatting command as a field with Kong struct tags. When the user invokes `gum style`, `gum format`, `gum join`, or `gum log`, Kong parses the arguments into the corresponding `Options` struct and calls its `Run()` method [main.go76](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L76)
 The command executes synchronously and writes results to stdout.

**Sources:** [gum.go25-228](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L25-L228)
 [main.go46-76](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L46-L76)

Data Flow Pattern
-----------------

### Pipeline Lifecycle

This diagram illustrates how formatting commands interact with shell input/output compared to internal rendering libraries.

Formatting commands follow a consistent data flow:

1.  **Input Acquisition**: Read from stdin (pipe) or command-line arguments.
2.  **Option Parsing**: Extract styling/formatting options from flags via Kong [main.go47-75](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L47-L75)
    
3.  **Transformation**: Apply Lipgloss styling, Glamour rendering, or layout operations.
4.  **Output**: Write result to stdout for shell pipeline composition.

Unlike interactive commands that separate UI (stderr) and data (stdout), formatting commands use stdout exclusively since they produce no interactive interface.

**Sources:** [README.md307-390](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L307-L390)
 [gum.go89-217](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L89-L217)

Command Catalog
---------------

### Style Command

The `style` command provides direct access to Lipgloss styling capabilities without writing Go code. It applies colors, borders, padding, margins, alignment, and other visual properties to text.

The Style struct is registered at [gum.go181](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L181-L181)
 with the command description spanning [gum.go158-181](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L158-L181)
 All Lipgloss styling properties are exposed as command-line flags.

For details, see [Style Command](https://deepwiki.com/charmbracelet/gum/6.1-style-command)
.

**Sources:** [README.md307-318](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L307-L318)
 [gum.go158-181](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L158-L181)
 [examples/demo.sh3](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh#L3-L3)

### Format Command

The `format` command processes text through different renderers: markdown (via Glamour), code syntax highlighting, template expansion (using termenv helpers), and emoji substitution.

The Format struct is registered at [gum.go92](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L92-L92)
 with usage details at [gum.go89-92](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L89-L92)
 It integrates the Glamour library for markdown rendering.

For details, see [Format Command](https://deepwiki.com/charmbracelet/gum/6.2-format-command)
.

**Sources:** [README.md341-366](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L341-L366)
 [gum.go89-92](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L89-L92)

### Join Command

The `join` command combines multiple text blocks horizontally or vertically with configurable alignment. It's commonly used with `gum style` output to build complex terminal layouts.

The Join struct is registered at [gum.go120](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L120-L120)
 with documentation at [gum.go104-120](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L104-L120)
 It wraps Lipgloss's `JoinHorizontal` and `JoinVertical` functions.

For details, see [Join Command](https://deepwiki.com/charmbracelet/gum/6.3-join-command)
.

**Sources:** [README.md320-339](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L320-L339)
 [gum.go104-120](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L104-L120)
 [examples/demo.sh45-47](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh#L45-L47)

### Log Command

The `log` command provides structured logging using the `charmbracelet/log` library. It supports different log levels (debug, info, warn, error), timestamps, and key-value structured fields.

The Log struct is registered at [gum.go217](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L217-L217)
 with description at [gum.go210-217](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L210-L217)
 It directly wraps the `charmbracelet/log` library.

For details, see [Log Command](https://deepwiki.com/charmbracelet/gum/6.4-log-command)
.

**Sources:** [README.md368-390](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L368-L390)
 [gum.go210-217](https://github.com/charmbracelet/gum/blob/30cc5169/gum.go#L210-L217)

Integration with Shell Pipelines
--------------------------------

Formatting commands are designed for composition. Their stdin/stdout-only interface enables chaining with other gum commands and standard Unix tools. The [examples/demo.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh)
 script demonstrates extensive use of `gum style` and `gum join` for building a multi-step interactive script with styled output:

**Sources:** [examples/demo.sh1-48](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh#L1-L48)
 [README.md322-336](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L322-L336)

Configuration and Customization
-------------------------------

Like all gum commands, formatting commands support configuration via both command-line flags and environment variables. The Kong parser in [main.go47-75](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L47-L75)
 sets up default values for styling options via the `kong.Vars` map at [main.go56-74](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L56-L74)

These defaults include `defaultHeight`, `defaultWidth`, `defaultAlign`, and `defaultBorder` [main.go59-62](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L59-L62)
 which apply across formatting commands that utilize Lipgloss.

**Sources:** [README.md158-188](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L158-L188)
 [main.go56-74](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L56-L74)

Error Handling
--------------

Formatting commands use the same error handling pattern as the rest of gum. They return errors that are processed by the main error handler in [main.go76-90](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L90)
:

1.  `exit.ErrExit` errors translate to specific exit codes [main.go77-80](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L77-L80)
    
2.  Other errors print to stderr and exit with code 1 [main.go88-89](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L88-L89)
    
3.  No timeout handling (unlike interactive commands) since execution is synchronous and does not use a Bubbletea event loop.

**Sources:** [main.go76-90](https://github.com/charmbracelet/gum/blob/30cc5169/main.go#L76-L90)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Formatting Commands](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#formatting-commands)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#purpose-and-scope)
    
*   [Architectural Characteristics](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#architectural-characteristics)
    
*   [Command Registration and Execution](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#command-registration-and-execution)
    
*   [System to Code Mapping](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#system-to-code-mapping)
    
*   [Data Flow Pattern](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#data-flow-pattern)
    
*   [Pipeline Lifecycle](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#pipeline-lifecycle)
    
*   [Command Catalog](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#command-catalog)
    
*   [Style Command](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#style-command)
    
*   [Format Command](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#format-command)
    
*   [Join Command](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#join-command)
    
*   [Log Command](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#log-command)
    
*   [Integration with Shell Pipelines](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#integration-with-shell-pipelines)
    
*   [Configuration and Customization](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#configuration-and-customization)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/gum/6-formatting-commands#error-handling)
