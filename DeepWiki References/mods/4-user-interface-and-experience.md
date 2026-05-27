Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/mods](https://github.com/charmbracelet/mods "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 7 November 2025 ([07a05d](https://github.com/charmbracelet/mods/commits/07a05d5b)
)

*   [Introduction to Mods](https://deepwiki.com/charmbracelet/mods/1-introduction-to-mods)
    
*   [Installation and Setup](https://deepwiki.com/charmbracelet/mods/1.1-installation-and-setup)
    
*   [Quick Start Guide](https://deepwiki.com/charmbracelet/mods/1.2-quick-start-guide)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/mods/2-core-architecture)
    
*   [Application Entry Point](https://deepwiki.com/charmbracelet/mods/2.1-application-entry-point)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/mods/2.2-configuration-system)
    
*   [Bubble Tea UI Model](https://deepwiki.com/charmbracelet/mods/2.3-bubble-tea-ui-model)
    
*   [Message Stream Context](https://deepwiki.com/charmbracelet/mods/2.4-message-stream-context)
    
*   [Conversation Management](https://deepwiki.com/charmbracelet/mods/2.5-conversation-management)
    
*   [LLM Provider Integration](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration)
    
*   [Provider Configuration](https://deepwiki.com/charmbracelet/mods/3.1-provider-configuration)
    
*   [Client Initialization and Streaming](https://deepwiki.com/charmbracelet/mods/3.2-client-initialization-and-streaming)
    
*   [Model Resolution and Fallbacks](https://deepwiki.com/charmbracelet/mods/3.3-model-resolution-and-fallbacks)
    
*   [User Interface and Experience](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience)
    
*   [Terminal Capabilities and Styling](https://deepwiki.com/charmbracelet/mods/4.1-terminal-capabilities-and-styling)
    
*   [Animation System](https://deepwiki.com/charmbracelet/mods/4.2-animation-system)
    
*   [Output Rendering and Formatting](https://deepwiki.com/charmbracelet/mods/4.3-output-rendering-and-formatting)
    
*   [Data Persistence](https://deepwiki.com/charmbracelet/mods/5-data-persistence)
    
*   [Conversation Database](https://deepwiki.com/charmbracelet/mods/5.1-conversation-database)
    
*   [Cache System](https://deepwiki.com/charmbracelet/mods/5.2-cache-system)
    
*   [Supporting Systems](https://deepwiki.com/charmbracelet/mods/6-supporting-systems)
    
*   [Flag Parsing and Error Handling](https://deepwiki.com/charmbracelet/mods/6.1-flag-parsing-and-error-handling)
    
*   [Content Loading](https://deepwiki.com/charmbracelet/mods/6.2-content-loading)
    
*   [Message Utilities](https://deepwiki.com/charmbracelet/mods/6.3-message-utilities)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/mods/7-development-guide)
    
*   [Dependencies and Build System](https://deepwiki.com/charmbracelet/mods/7.1-dependencies-and-build-system)
    
*   [Testing Strategy](https://deepwiki.com/charmbracelet/mods/7.2-testing-strategy)
    
*   [Code Quality and CI/CD](https://deepwiki.com/charmbracelet/mods/7.3-code-quality-and-cicd)
    

Menu

User Interface and Experience
=============================

Relevant source files

*   [.gitignore](https://github.com/charmbracelet/mods/blob/07a05d5b/.gitignore)
    
*   [styles.go](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go)
    
*   [term.go](https://github.com/charmbracelet/mods/blob/07a05d5b/term.go)
    

Purpose and Scope
-----------------

This document covers the terminal user interface system in `mods`, including TTY (teletypewriter) detection, adaptive rendering strategies, and the lipgloss styling framework. The system provides styled terminal output that adapts to different terminal capabilities and usage contexts (interactive vs piped).

For information about the Bubble Tea state machine and UI model, see [Bubble Tea UI Model](https://deepwiki.com/charmbracelet/mods/2.3-bubble-tea-ui-model)
. For details about animation rendering during AI requests, see [Animation System](https://deepwiki.com/charmbracelet/mods/4.2-animation-system)
. For markdown rendering and viewport management, see [Output Rendering and Formatting](https://deepwiki.com/charmbracelet/mods/4.3-output-rendering-and-formatting)
.

Overview
--------

The `mods` terminal interface is designed to provide a rich interactive experience when run in a terminal while gracefully degrading to plain text when used in Unix pipes or scripts. The system uses three key components:

1.  **TTY Detection** - Determines if stdin/stdout are connected to a terminal
2.  **Lipgloss Styling** - Provides color, formatting, and layout capabilities
3.  **Dual Renderers** - Separate renderers for stdout and stderr with independent color profiles

**Diagram: Terminal UI Initialization Flow**

Sources: [term.go1-35](https://github.com/charmbracelet/mods/blob/07a05d5b/term.go#L1-L35)
 [styles.go1-65](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L1-L65)

TTY Detection
-------------

The application uses synchronized once-value functions to detect terminal capabilities at startup. These checks determine whether the application is running in an interactive terminal or is part of a Unix pipeline.

### Detection Functions

| Function | Purpose | Implementation |
| --- | --- | --- |
| `isInputTTY()` | Checks if stdin is a terminal | Uses `isatty.IsTerminal(os.Stdin.Fd())` |
| `isOutputTTY()` | Checks if stdout is a terminal | Uses `isatty.IsTerminal(os.Stdout.Fd())` |

Both functions use `sync.OnceValue` to ensure the terminal detection occurs exactly once and is cached for subsequent calls. This optimization prevents repeated system calls.

**Diagram: OnceValue Caching Pattern**

### Usage Context Detection

The TTY detection enables different rendering modes:

| Input TTY | Output TTY | Behavior |
| --- | --- | --- |
| ✓   | ✓   | Full interactive mode with Bubble Tea UI |
| ✗   | ✓   | Pipe input mode, formatted output |
| ✓   | ✗   | Output piped to another command, plain text |
| ✗   | ✗   | Full pipe mode, minimal formatting |

This adaptive behavior allows `mods` to work seamlessly in command chains like:

*   `echo "question" | mods` (piped input)
*   `mods "question" | less` (piped output)
*   `cat file.txt | mods "explain" | tee output.txt` (full pipeline)

Sources: [term.go12-18](https://github.com/charmbracelet/mods/blob/07a05d5b/term.go#L12-L18)

Lipgloss Renderer Architecture
------------------------------

The application maintains two independent lipgloss renderers to handle different output streams with appropriate color profiles:

### Dual Renderer System

**Diagram: Dual Renderer Architecture**

### Renderer Initialization Details

**Stdout Renderer** ([term.go20-22](https://github.com/charmbracelet/mods/blob/07a05d5b/term.go#L20-L22)
):

*   Uses `lipgloss.DefaultRenderer()` which automatically detects terminal capabilities
*   Handles main application output including AI responses
*   Used by confirmation messages via `printConfirmation`

**Stderr Renderer** ([term.go28-30](https://github.com/charmbracelet/mods/blob/07a05d5b/term.go#L28-L30)
):

*   Explicitly configured with `lipgloss.NewRenderer(os.Stderr)`
*   Includes `termenv.WithColorCache(true)` for performance optimization
*   Used exclusively for error messages and diagnostics

Both renderers are wrapped in `sync.OnceValue` to ensure singleton initialization.

Sources: [term.go20-34](https://github.com/charmbracelet/mods/blob/07a05d5b/term.go#L20-L34)

Style System
------------

The `styles` struct defines a comprehensive set of named styles used throughout the application. These styles are created by the `makeStyles` function which accepts a lipgloss renderer.

### Style Structure

**Diagram: Style Categories and Definitions**

### Complete Style Reference

The `styles` struct ([styles.go10-28](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L10-L28)
) contains the following styles:

| Style | Purpose | Visual Properties | Usage Context |
| --- | --- | --- | --- |
| `AppName` | Application name display | Bold | Help text, headers |
| `CliArgs` | CLI argument examples | Gray (#585858) | Documentation |
| `Comment` | Explanatory comments | Gray (#757575) | Help text |
| `CyclingChars` | Animation characters | Pink (#FF87D7) | Loading indicators |
| `ErrorHeader` | Error header badge | White on red (#FF5F87), bold, padded | Error messages |
| `ErrorDetails` | Error details text | Gray (reuses Comment) | Error descriptions |
| `ErrPadding` | Error message padding | Horizontal padding (2) | Error layout |
| `Flag` | CLI flag names | Adaptive teal/cyan, bold | Flag documentation |
| `FlagComma` | Flag list separator | Adaptive teal/brown | Flag lists |
| `FlagDesc` | Flag descriptions | Gray (reuses Comment) | Help text |
| `InlineCode` | Inline code snippets | Pink (#FF5F87) on dark gray (#3A3A3A) | Code examples |
| `Link` | Hyperlinks | Teal (#00AF87), underlined | URLs |
| `Pipe` | Pipe character in examples | Adaptive purple/indigo | Command pipelines |
| `Quote` | Quoted text | Adaptive pink | Examples |
| `ConversationList` | Conversation list entries | Padded (1) | List display |
| `SHA1` | Conversation IDs | Bold teal (reuses Flag) | Conversation references |
| `Timeago` | Relative timestamps | Adaptive gray | Conversation metadata |

### Adaptive Colors

Several styles use `lipgloss.AdaptiveColor` which provides different colors for light and dark terminal backgrounds:

This ensures readability across different terminal color schemes without manual configuration.

Sources: [styles.go30-50](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L30-L50)

Confirmation and Action Messages
--------------------------------

The application uses a specialized header style for confirmation messages that appear after successful operations (e.g., saving conversations to files).

### Output Header System

**Diagram: Confirmation Message Flow**

The `outputHeader` style ([styles.go56](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L56-L56)
):

*   Foreground: White (#F1F1F1)
*   Background: Purple (#6C50FF)
*   Bold with padding (0, 1) and right margin (1)
*   Dynamically set text via `SetString()`

The `printConfirmation` function ([styles.go58-64](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L58-L64)
) accepts an action verb (e.g., "WROTE", "SAVED") and content string, formatting them as:

    [ACTION] content-string
    

The default action is "WROTE" ([styles.go54](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L54-L54)
) if no action is specified.

Sources: [styles.go52-65](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L52-L65)

Style Creation Process
----------------------

The `makeStyles` function ([styles.go30-50](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L30-L50)
) creates all application styles from a single lipgloss renderer:

**Diagram: Style Creation Sequence**

Each style is constructed by:

1.  Creating a new base style from the renderer
2.  Chaining configuration methods (`.Foreground()`, `.Bold()`, `.Padding()`, etc.)
3.  Returning the fully configured style

The horizontal edge padding constant ([styles.go31](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L31-L31)
) is set to `2` for consistent spacing in error messages.

Sources: [styles.go30-50](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L30-L50)

Integration with Application Logic
----------------------------------

The terminal UI system integrates with the broader application through these entry points:

| Component | Access Pattern | Use Case |
| --- | --- | --- |
| Main execution | `isOutputTTY()` check | Determines if Bubble Tea UI should start |
| Error formatting | `stderrStyles()` | Formats error messages |
| Confirmation output | `stdoutStyles()` | Formats success messages |
| Animation rendering | `stdoutStyles().CyclingChars` | Colors loading spinners |

The caching pattern (`sync.OnceValue`) ensures that:

*   Terminal detection occurs once per execution
*   Renderers are initialized only when needed
*   Style objects are reused across the application
*   No unnecessary allocations or system calls occur

Sources: [term.go1-35](https://github.com/charmbracelet/mods/blob/07a05d5b/term.go#L1-L35)
 [styles.go1-65](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L1-L65)

Color Palette
-------------

The application uses a consistent color palette across all styles:

| Color | Hex Code | Usage |
| --- | --- | --- |
| Light Gray | #757575 | Comments, descriptions |
| Dark Gray | #585858 | CLI arguments |
| Bright Pink | #FF87D7 | Animations, cycling characters |
| Hot Pink | #FF5F87 | Errors, inline code |
| Teal/Cyan | #00AF87 - #3EEFCF | Links, flags, SHA1s |
| Purple/Indigo | #6C50FF - #8470FF | Action headers, pipes |
| Dark Background | #3A3A3A | Code blocks |
| Light Foreground | #F1F1F1 | Header text |

All colors are carefully chosen for WCAG contrast compliance and readability across different terminal emulators.

Sources: [styles.go1-65](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go#L1-L65)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [User Interface and Experience](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#user-interface-and-experience)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#purpose-and-scope)
    
*   [Overview](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#overview)
    
*   [TTY Detection](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#tty-detection)
    
*   [Detection Functions](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#detection-functions)
    
*   [Usage Context Detection](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#usage-context-detection)
    
*   [Lipgloss Renderer Architecture](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#lipgloss-renderer-architecture)
    
*   [Dual Renderer System](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#dual-renderer-system)
    
*   [Renderer Initialization Details](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#renderer-initialization-details)
    
*   [Style System](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#style-system)
    
*   [Style Structure](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#style-structure)
    
*   [Complete Style Reference](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#complete-style-reference)
    
*   [Adaptive Colors](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#adaptive-colors)
    
*   [Confirmation and Action Messages](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#confirmation-and-action-messages)
    
*   [Output Header System](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#output-header-system)
    
*   [Style Creation Process](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#style-creation-process)
    
*   [Integration with Application Logic](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#integration-with-application-logic)
    
*   [Color Palette](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience#color-palette)
