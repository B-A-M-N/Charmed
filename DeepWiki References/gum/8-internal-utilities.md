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

Internal Utilities
==================

Relevant source files

*   [format/command.go](https://github.com/charmbracelet/gum/blob/30cc5169/format/command.go)
    
*   [format/formats.go](https://github.com/charmbracelet/gum/blob/30cc5169/format/formats.go)
    
*   [format/options.go](https://github.com/charmbracelet/gum/blob/30cc5169/format/options.go)
    
*   [internal/exit/exit.go](https://github.com/charmbracelet/gum/blob/30cc5169/internal/exit/exit.go)
    
*   [internal/files/files.go](https://github.com/charmbracelet/gum/blob/30cc5169/internal/files/files.go)
    
*   [internal/stdin/stdin.go](https://github.com/charmbracelet/gum/blob/30cc5169/internal/stdin/stdin.go)
    
*   [internal/timeout/context.go](https://github.com/charmbracelet/gum/blob/30cc5169/internal/timeout/context.go)
    
*   [internal/tty/tty.go](https://github.com/charmbracelet/gum/blob/30cc5169/internal/tty/tty.go)
    

Purpose and Scope
-----------------

The `internal` package provides shared utility functions and abstractions used across all `gum` commands. These utilities handle cross-cutting concerns including input reading, timeout management, exit code handling, TTY detection, and file system operations. This standardization ensures consistent behavior across all commands and reduces code duplication.

Package Overview
----------------

The internal utilities are organized into specialized packages, each addressing a specific infrastructure concern:

Sources: [internal/stdin/stdin.go1-12](https://github.com/charmbracelet/gum/blob/30cc5169/internal/stdin/stdin.go#L1-L12)
 [internal/timeout/context.go1-7](https://github.com/charmbracelet/gum/blob/30cc5169/internal/timeout/context.go#L1-L7)
 [internal/exit/exit.go1-4](https://github.com/charmbracelet/gum/blob/30cc5169/internal/exit/exit.go#L1-L4)
 [internal/tty/tty.go1-11](https://github.com/charmbracelet/gum/blob/30cc5169/internal/tty/tty.go#L1-L11)
 [internal/files/files.go1-8](https://github.com/charmbracelet/gum/blob/30cc5169/internal/files/files.go#L1-L8)

stdin Package
-------------

### Purpose

The `stdin` package provides utilities for reading input from standard input with optional ANSI escape sequence stripping. This is essential for commands that accept pre-filled values from shell pipelines.

### Key Functions

| Function | Purpose | Implementation Detail |
| --- | --- | --- |
| `Read(...Option)` | Reads input from an stdin pipe | Uses `bufio.NewReader(os.Stdin)` [internal/stdin/stdin.go37-47](https://github.com/charmbracelet/gum/blob/30cc5169/internal/stdin/stdin.go#L37-L47) |
| `IsEmpty()` | Checks if stdin has data | Checks `os.ModeNamedPipe` and `stat.Size()` [internal/stdin/stdin.go80-91](https://github.com/charmbracelet/gum/blob/30cc5169/internal/stdin/stdin.go#L80-L91) |
| `StripANSI(bool)` | Option to clean input | Uses `ansi.Strip` from `github.com/charmbracelet/x/ansi` [internal/stdin/stdin.go23-27](https://github.com/charmbracelet/gum/blob/30cc5169/internal/stdin/stdin.go#L23-L27) |
| `SingleLine(bool)` | Option to read one line | Uses `reader.ReadLine()` [internal/stdin/stdin.go29-34](https://github.com/charmbracelet/gum/blob/30cc5169/internal/stdin/stdin.go#L29-L34) |

### Data Flow: Stdin Reading

Sources: [internal/stdin/stdin.go37-77](https://github.com/charmbracelet/gum/blob/30cc5169/internal/stdin/stdin.go#L37-L77)
 [format/command.go27](https://github.com/charmbracelet/gum/blob/30cc5169/format/command.go#L27-L27)

timeout Package
---------------

### Purpose

The `timeout` package provides context management for commands that need to enforce execution time limits. This prevents commands from hanging indefinitely and allows graceful cancellation.

### Key Function

| Function | Purpose | Return Value |
| --- | --- | --- |
| `Context(time.Duration)` | Creates a context with a timeout | `context.Context`, `context.CancelFunc` [internal/timeout/context.go10-16](https://github.com/charmbracelet/gum/blob/30cc5169/internal/timeout/context.go#L10-L16) |

If the duration is `0`, it returns a standard `context.Background()` and an empty cancel function [internal/timeout/context.go12-14](https://github.com/charmbracelet/gum/blob/30cc5169/internal/timeout/context.go#L12-L14)

Sources: [internal/timeout/context.go1-17](https://github.com/charmbracelet/gum/blob/30cc5169/internal/timeout/context.go#L1-L17)

exit Package
------------

### Purpose

The `exit` package defines standard exit codes used by `gum` to communicate command results to the calling shell.

### Constants and Types

*   `StatusTimeout = 124`: Exit code for timed out commands [internal/exit/exit.go7](https://github.com/charmbracelet/gum/blob/30cc5169/internal/exit/exit.go#L7-L7)
    
*   `StatusAborted = 130`: Exit code for aborted commands (e.g., Ctrl+C) [internal/exit/exit.go10](https://github.com/charmbracelet/gum/blob/30cc5169/internal/exit/exit.go#L10-L10)
    
*   `ErrExit`: A custom `int` type that implements the `error` interface, allowing functions to return specific exit codes as errors [internal/exit/exit.go13-16](https://github.com/charmbracelet/gum/blob/30cc5169/internal/exit/exit.go#L13-L16)
    

Sources: [internal/exit/exit.go1-17](https://github.com/charmbracelet/gum/blob/30cc5169/internal/exit/exit.go#L1-L17)

tty Package
-----------

### Purpose

The `tty` package provides TTY-aware printing. It detects if `stdout` is a terminal and handles ANSI sequence stripping when output is redirected to a file or pipe.

### Implementation

*   `isTTY`: A `sync.OnceValue` that caches the result of `term.IsTerminal(os.Stdout.Fd())` [internal/tty/tty.go13-15](https://github.com/charmbracelet/gum/blob/30cc5169/internal/tty/tty.go#L13-L15)
    
*   `Println(string)`: Prints the string to `stdout`. If the output is not a TTY, it automatically strips ANSI sequences using `ansi.Strip` [internal/tty/tty.go18-24](https://github.com/charmbracelet/gum/blob/30cc5169/internal/tty/tty.go#L18-L24)
    

Sources: [internal/tty/tty.go1-25](https://github.com/charmbracelet/gum/blob/30cc5169/internal/tty/tty.go#L1-L25)

files Package
-------------

### Purpose

The `files` package provides file system utilities, primarily for listing files in the current directory while ignoring common development directories.

### Functions

*   `List()`: Recursively walks the current directory using `filepath.Walk` and returns a slice of file paths [internal/files/files.go12-26](https://github.com/charmbracelet/gum/blob/30cc5169/internal/files/files.go#L12-L26)
    
*   `shouldIgnore(path)`: Internal helper that checks if a path starts with ignored prefixes like `node_modules`, `.git`, or `.` [internal/files/files.go30-37](https://github.com/charmbracelet/gum/blob/30cc5169/internal/files/files.go#L30-L37)
    

Sources: [internal/files/files.go1-38](https://github.com/charmbracelet/gum/blob/30cc5169/internal/files/files.go#L1-L38)

Integration Example: format Command
-----------------------------------

The `format` command illustrates how these utilities are composed in a functional command.

1.  **Stdin Handling**: It uses `stdin.Read` to fetch content if no template argument is provided [format/command.go27](https://github.com/charmbracelet/gum/blob/30cc5169/format/command.go#L27-L27)
    
2.  **ANSI Stripping**: It respects the `StripANSI` option when reading from stdin [format/command.go27](https://github.com/charmbracelet/gum/blob/30cc5169/format/command.go#L27-L27)
    
3.  **Formatting Logic**: It delegates to specialized functions (`code`, `emoji`, `template`, `markdown`) which utilize external libraries like `glamour` and `termenv` [format/formats.go13-67](https://github.com/charmbracelet/gum/blob/30cc5169/format/formats.go#L13-L67)
    

Sources: [format/command.go21-46](https://github.com/charmbracelet/gum/blob/30cc5169/format/command.go#L21-L46)
 [format/formats.go1-68](https://github.com/charmbracelet/gum/blob/30cc5169/format/formats.go#L1-L68)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Internal Utilities](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#internal-utilities)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#purpose-and-scope)
    
*   [Package Overview](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#package-overview)
    
*   [stdin Package](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#stdin-package)
    
*   [Purpose](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#purpose)
    
*   [Key Functions](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#key-functions)
    
*   [Data Flow: Stdin Reading](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#data-flow-stdin-reading)
    
*   [timeout Package](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#timeout-package)
    
*   [Purpose](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#purpose-1)
    
*   [Key Function](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#key-function)
    
*   [exit Package](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#exit-package)
    
*   [Purpose](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#purpose-2)
    
*   [Constants and Types](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#constants-and-types)
    
*   [tty Package](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#tty-package)
    
*   [Purpose](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#purpose-3)
    
*   [Implementation](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#implementation)
    
*   [files Package](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#files-package)
    
*   [Purpose](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#purpose-4)
    
*   [Functions](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#functions)
    
*   [Integration Example: format Command](https://deepwiki.com/charmbracelet/gum/8-internal-utilities#integration-example-format-command)
