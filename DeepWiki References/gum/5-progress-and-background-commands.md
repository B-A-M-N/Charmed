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

Progress and Background Commands
================================

Relevant source files

*   [spin/command.go](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go)
    
*   [spin/options.go](https://github.com/charmbracelet/gum/blob/30cc5169/spin/options.go)
    
*   [spin/spin.go](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go)
    
*   [spin/spinners.go](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spinners.go)
    

Purpose and Scope
-----------------

This document covers commands in gum that execute background processes while displaying progress indicators to the user. These commands have a unique execution model that combines non-interactive subprocess execution with a terminal user interface for progress feedback. The primary command in this category is `spin`, which runs shell commands in the background while displaying an animated spinner.

Unlike interactive commands that wait for user input (see [Interactive Commands](https://deepwiki.com/charmbracelet/gum/4-interactive-commands)
) or formatting commands that transform data (see [Formatting Commands](https://deepwiki.com/charmbracelet/gum/6-formatting-commands)
), progress and background commands execute external processes and capture their output while presenting real-time visual feedback.

For details, see [Spin Command](https://deepwiki.com/charmbracelet/gum/5.1-spin-command)
.

Command Category Overview
-------------------------

Progress and background commands represent a distinct execution pattern in gum's architecture. While most gum commands either present interactive UI elements for user input or transform text through pipelines, the `spin` command executes arbitrary shell commands as subprocesses and manages their lifecycle.

| Aspect | Progress/Background Commands | Interactive Commands | Formatting Commands |
| --- | --- | --- | --- |
| User Input | None (Ctrl+C to abort) | Keyboard/mouse for selections | None |
| Subprocess Execution | Yes - external commands | No  | No  |
| Output Timing | Async (during execution) | Sync (after user input) | Sync (immediate) |
| UI Rendering | Stderr (progress indicator) | Stderr (interactive elements) | Stdout (formatted text) |
| Data Output | Stdout (command results) | Stdout (user selections) | Stdout (transformed text) |
| Exit Code | Propagated from subprocess | Based on user choice | Based on processing success |

**Sources:** [spin/spin.go1-15](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L1-L15)
 [spin/command.go15-17](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L15-L17)

Execution Architecture
----------------------

The execution architecture for progress and background commands differs significantly from other gum commands. The system must orchestrate three concurrent concerns: displaying a progress indicator, executing a background process, and capturing that process's output.

**Diagram: Concurrent Execution Flow for Spin Command**

The diagram illustrates how `commandStart()` executes asynchronously while the Bubbletea program continuously updates the spinner animation. The `Init()` method at [spin/spin.go142-147](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L142-L147)
 launches both operations using `tea.Batch`, enabling true parallelism.

**Sources:** [spin/spin.go68-133](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L68-L133)
 [spin/command.go17-82](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L17-L82)
 [spin/spin.go142-147](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L142-L147)

Model Structure and State Management
------------------------------------

The `spin` command uses a specialized model structure that maintains state for both the UI rendering and the background process execution.

**Diagram: Model State and Message Flow**

The model at [spin/spin.go33-49](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L33-L49)
 maintains separate concerns: UI state (`spinner`, `title`, `padding`, `align`), process state (`command`, `status`, `err`), and output state (`stdout`, `stderr`, `output`). The global buffers at [spin/spin.go51-57](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L51-L57)
 are shared across the concurrent execution context.

**Sources:** [spin/spin.go33-49](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L33-L49)
 [spin/spin.go51-57](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L51-L57)
 [spin/spin.go59-66](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L59-L66)

PTY Support and Output Capture
------------------------------

A critical feature of the `spin` command is its ability to capture output from background processes using pseudo-terminals (PTY). This enables proper handling of terminal-aware applications that might behave differently when their output is redirected to a pipe.

**Diagram: PTY Allocation and Output Capture Strategy**

The PTY support at [spin/spin.go84-119](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L84-L119)
 handles three scenarios:

1.  **Non-TTY output**: When stdout is not a terminal, output goes directly to the actual stdout/stderr without capture [spin/spin.go115-119](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L115-L119)
    
2.  **Windows TTY**: Due to Git Bash PTY issues, uses `io.MultiWriter` to capture without PTY allocation [spin/spin.go84-87](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L84-L87)
    
3.  **Unix-like TTY**: Allocates separate PTYs for stdout and stderr using `openPty`, allowing proper terminal emulation for the subprocess [spin/spin.go88-114](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L88-L114)
    

The `openPty` function (imported from `github.com/charmbracelet/x/xpty`) creates a pseudo-terminal device. The slave side attaches to the subprocess, while the master side pipes into buffers through goroutines at [spin/spin.go108-109](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L108-L109)

**Sources:** [spin/spin.go68-133](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L68-L133)
 [spin/spin.go78-119](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L78-L119)

Output Handling and Display Options
-----------------------------------

The `spin` command provides granular control over which output streams to display and when to display them. This flexibility supports different use cases from silent background execution to verbose debugging.

| Option Flag | Environment Variable | Behavior | Use Case |
| --- | --- | --- | --- |
| `--show-output` | `GUM_SPIN_SHOW_OUTPUT` | Display both stdout and stderr during execution | Monitoring long-running processes |
| `--show-stdout` | `GUM_SPIN_SHOW_STDOUT` | Display only stdout during execution | Piping structured output |
| `--show-stderr` | `GUM_SPIN_SHOW_STDERR` | Display only stderr during execution | Monitoring error logs |
| `--show-error` | `GUM_SPIN_SHOW_ERROR` | Display all output only if command fails | Debugging failed commands |

The output routing logic at [spin/command.go51-79](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L51-L79)
 implements a priority system:

1.  If the command has an internal error (`m.err != nil`), write the error to stderr and exit with status 1 [spin/command.go54-58](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L54-L58)
    
2.  If the command succeeds (`m.status == 0`), determine which output to write to stdout based on flags [spin/command.go59-72](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L59-L72)
    
3.  If the command fails and `--show-error` is set, write all captured output to stdout [spin/command.go73-79](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L73-L79)
    

The `View()` method at [spin/spin.go149-175](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L149-L175)
 handles real-time display during execution. When output flags are enabled and the output is a TTY, it appends buffer contents to the spinner UI.

**Sources:** [spin/command.go51-79](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L51-L79)
 [spin/spin.go149-175](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L149-L175)
 [spin/options.go13-16](https://github.com/charmbracelet/gum/blob/30cc5169/spin/options.go#L13-L16)

Exit Code Propagation
---------------------

A distinguishing feature of the `spin` command is that it propagates the exit code of the executed subprocess, enabling shell script conditionals to react to command success or failure.

The exit code flow:

1.  **Process completion**: At [spin/spin.go121-124](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L121-L124)
     the process exit code is retrieved via `executing.ProcessState.ExitCode()`, with `-1` normalized to `1`.
2.  **Message delivery**: The status is packaged in `finishCommandMsg` at [spin/spin.go126-131](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L126-L131)
    
3.  **Model update**: The `Update()` method stores the status in `m.status` at [spin/spin.go183](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L183-L183)
    
4.  **Exit code return**: The `Run()` method returns `exit.ErrExit(m.status)` at [spin/command.go81](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L81-L81)
    

This propagation enables shell patterns like:

**Sources:** [spin/spin.go121-131](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L121-L131)
 [spin/spin.go179-184](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L179-L184)
 [spin/command.go81](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L81-L81)

Timeout Support
---------------

The `spin` command integrates with gum's timeout infrastructure to abort long-running commands. The timeout is specified via the `--timeout` flag or `GUM_SPIN_TIMEOUT` environment variable at [spin/options.go22](https://github.com/charmbracelet/gum/blob/30cc5169/spin/options.go#L22-L22)

At [spin/command.go37-38](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L37-L38)
 a context with timeout is created using `timeout.Context(o.Timeout)`. This context is passed to the Bubbletea program via `tea.WithContext(ctx)` at [spin/command.go43](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L43-L43)

The subprocess itself uses `exec.CommandContext(context.Background(), ...)` at [spin/spin.go75](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L75-L75)
 This design ensures the spinner UI terminates on timeout while allowing the subprocess to be managed by the context provided to `Run()`.

**Sources:** [spin/command.go37-43](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L37-L43)
 [spin/options.go22](https://github.com/charmbracelet/gum/blob/30cc5169/spin/options.go#L22-L22)
 [spin/spin.go75](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L75-L75)

Styling Configuration
---------------------

Like other gum commands, `spin` supports extensive styling customization through the `style.Styles` embedding pattern.

The `Options` struct at [spin/options.go9-24](https://github.com/charmbracelet/gum/blob/30cc5169/spin/options.go#L9-L24)
 embeds two style configurations:

*   **`SpinnerStyle`**: Controls the appearance of the spinner animation, prefixed with `spinner.` flags.
*   **`TitleStyle`**: Controls the appearance of the title text, prefixed with `title.` flags.

The initialization at [spin/command.go22-27](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L22-L27)
 applies these styles:

The layout is controlled by:

*   **`--align`**: Positions spinner left or right of the title at [spin/spin.go167-171](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L167-L171)
    
*   **`--padding`**: Applies padding around the entire display at [spin/spin.go172-174](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L172-L174)
    

**Sources:** [spin/options.go9-24](https://github.com/charmbracelet/gum/blob/30cc5169/spin/options.go#L9-L24)
 [spin/command.go22-27](https://github.com/charmbracelet/gum/blob/30cc5169/spin/command.go#L22-L27)
 [spin/spin.go167-174](https://github.com/charmbracelet/gum/blob/30cc5169/spin/spin.go#L167-L174)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Progress and Background Commands](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#progress-and-background-commands)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#purpose-and-scope)
    
*   [Command Category Overview](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#command-category-overview)
    
*   [Execution Architecture](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#execution-architecture)
    
*   [Model Structure and State Management](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#model-structure-and-state-management)
    
*   [PTY Support and Output Capture](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#pty-support-and-output-capture)
    
*   [Output Handling and Display Options](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#output-handling-and-display-options)
    
*   [Exit Code Propagation](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#exit-code-propagation)
    
*   [Timeout Support](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#timeout-support)
    
*   [Styling Configuration](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands#styling-configuration)
