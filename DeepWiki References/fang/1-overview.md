Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/fang](https://github.com/charmbracelet/fang "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 15 March 2026 ([d89b30](https://github.com/charmbracelet/fang/commits/d89b30af)
)

*   [Overview](https://deepwiki.com/charmbracelet/fang/1-overview)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/fang/2-getting-started)
    
*   [Core Components](https://deepwiki.com/charmbracelet/fang/3-core-components)
    
*   [fang.Execute Function](https://deepwiki.com/charmbracelet/fang/3.1-fang.execute-function)
    
*   [Help Rendering System](https://deepwiki.com/charmbracelet/fang/3.2-help-rendering-system)
    
*   [Styling and Themes](https://deepwiki.com/charmbracelet/fang/3.3-styling-and-themes)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/fang/3.4-error-handling)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/fang/4-configuration-options)
    
*   [Features](https://deepwiki.com/charmbracelet/fang/5-features)
    
*   [Version Management](https://deepwiki.com/charmbracelet/fang/5.1-version-management)
    
*   [Manpage Generation](https://deepwiki.com/charmbracelet/fang/5.2-manpage-generation)
    
*   [Shell Completions](https://deepwiki.com/charmbracelet/fang/5.3-shell-completions)
    
*   [Signal Handling](https://deepwiki.com/charmbracelet/fang/5.4-signal-handling)
    
*   [Dependencies and Architecture](https://deepwiki.com/charmbracelet/fang/6-dependencies-and-architecture)
    
*   [Testing and Examples](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples)
    
*   [Golden File Testing](https://deepwiki.com/charmbracelet/fang/7.1-golden-file-testing)
    
*   [Example Application](https://deepwiki.com/charmbracelet/fang/7.2-example-application)
    
*   [Development](https://deepwiki.com/charmbracelet/fang/8-development)
    
*   [Code Quality](https://deepwiki.com/charmbracelet/fang/8.1-code-quality)
    
*   [Platform Support](https://deepwiki.com/charmbracelet/fang/8.2-platform-support)
    
*   [License](https://deepwiki.com/charmbracelet/fang/8.3-license)
    

Menu

Overview
========

Relevant source files

*   [LICENSE.md](https://github.com/charmbracelet/fang/blob/d89b30af/LICENSE.md?plain=1)
    
*   [README.md](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1)
    
*   [go.mod](https://github.com/charmbracelet/fang/blob/d89b30af/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/fang/blob/d89b30af/go.sum)
    
*   [testdata/TestSetup/with\_subcommands/help-sub-sub.golden](https://github.com/charmbracelet/fang/blob/d89b30af/testdata/TestSetup/with_subcommands/help-sub-sub.golden)
    

This document provides a comprehensive introduction to the Fang CLI framework, explaining its architecture, core components, and how it enhances Cobra-based command-line applications with styled output and additional features.

For practical integration instructions, see [Getting Started](https://deepwiki.com/charmbracelet/fang/2-getting-started)
. For detailed component documentation, see [Core Components](https://deepwiki.com/charmbracelet/fang/3-core-components)
. For configuration details, see [Configuration Options](https://deepwiki.com/charmbracelet/fang/4-configuration-options)
.

Purpose and Scope
-----------------

Fang is a batteries-included wrapper library for [spf13/cobra](https://github.com/charmbracelet/fang/blob/d89b30af/spf13/cobra)
 applications that adds:

*   **Styled terminal output** using lipgloss v2 for help text and error messages
*   **Automatic version management** with `--version` flag generation
*   **Manpage generation** via a hidden `man` command
*   **Enhanced error handling** with styled error output
*   **Theme customization** through color schemes
*   **Signal handling** for graceful shutdown
*   **Shell completions** management

Fang does not replace Cobra—it wraps and enhances existing `cobra.Command` structures. Applications define their commands using standard Cobra patterns, then pass the root command to `fang.Execute` for execution with enhanced features.

Sources: [README.md12-28](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L12-L28)
 [go.mod1-19](https://github.com/charmbracelet/fang/blob/d89b30af/go.mod#L1-L19)

Architecture Overview
---------------------

Fang operates as a presentation and orchestration layer on top of Cobra. The architecture consists of three primary layers:

**Key Architectural Principles:**

1.  **Non-invasive Enhancement**: Fang modifies `cobra.Command` objects before execution but does not replace Cobra's execution model
2.  **Functional Configuration**: All customization occurs through functional options passed to `fang.Execute`
3.  **Separation of Concerns**: Theming, help rendering, and error handling are independent subsystems
4.  **Terminal-Aware**: All output adapts to terminal capabilities via color profile detection

Sources: [README.md40-66](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L40-L66)
 [go.mod5-19](https://github.com/charmbracelet/fang/blob/d89b30af/go.mod#L5-L19)

Core Components and Code Mapping
--------------------------------

### Execution Entry Point

The entry point `fang.Execute` is the single function users call. It accepts:

*   A `context.Context` for cancellation and signal handling
*   A `*cobra.Command` root command
*   Variadic `Option` functions for configuration

Sources: [README.md42-64](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L42-L64)
 High-level diagrams

### Help Rendering System

| Component | Location | Purpose |
| --- | --- | --- |
| `helpFn` | `help.go` | Main help rendering function, replaces `cobra.Command.HelpFunc` |
| `styleUsage` | `help.go` | Formats usage lines with program name styling |
| `styleExamples` | `help.go` | Renders code examples with syntax highlighting |
| `evalFlags` | `help.go` | Processes and styles flag descriptions |
| `evalCmds` | `help.go` | Processes and styles subcommand listings |

The help system completely overrides Cobra's default help output. Each section (usage, examples, flags, commands) is processed through dedicated styling functions.

Sources: High-level Diagram 2, High-level Diagram 3

### Theme and Styling Engine

**Key Type Definitions:**

*   `ColorScheme`: Struct containing color values for different UI elements
*   `ColorSchemeFunc`: Function type `func(*colorprofile.Writer) ColorScheme` that returns colors based on terminal capabilities
*   `Styles`: Struct containing `lipgloss.Style` objects created from `ColorScheme`
*   `makeStyles`: Function that transforms `ColorScheme` into `Styles` by creating lipgloss style objects

The theme system separates color definition from style application. Users can provide custom `ColorSchemeFunc` implementations via `WithColorSchemeFunc` option.

Sources: High-level Diagram 3, [go.mod6](https://github.com/charmbracelet/fang/blob/d89b30af/go.mod#L6-L6)

### Error Handling

**Error Handler Signature:**

    ErrorHandler func(cmd *cobra.Command, err error)
    

The default error handler:

1.  Detects whether the error is a user error (invalid flags/arguments) or system error
2.  Suppresses usage output for user errors (UX improvement over default Cobra)
3.  Applies error styling from the theme system
4.  Writes to stderr with terminal capability detection

Users can provide custom error handlers via `WithErrorHandler` option.

Sources: [README.md27](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L27-L27)
 High-level Diagram 1

Execution Lifecycle
-------------------

**Phase Breakdown:**

1.  **Configuration Phase**: Functional options are applied to build the internal `config` struct
2.  **Enhancement Phase**: Cobra command tree is modified (help override, version injection, man command addition)
3.  **Execution Phase**: Control is delegated to `cobra.Command.ExecuteContext`
4.  **Output Phase**: Help or error handlers produce styled terminal output

Sources: High-level Diagram 5, [README.md55-63](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L55-L63)

Feature Integration Points
--------------------------

### Version Management

By default, Fang extracts version from `runtime/debug.BuildInfo`. Users can override via `WithVersion` or disable via `WithoutVersion`.

Sources: [README.md22](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L22-L22)
 [go.mod3](https://github.com/charmbracelet/fang/blob/d89b30af/go.mod#L3-L3)

### Manpage Generation

Integration with `muesli/mango-cobra` adds a hidden `man` command that generates manpages from Cobra command metadata. This can be disabled with `WithoutManpage`.

Sources: [README.md23-24](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L23-L24)
 [go.mod12-13](https://github.com/charmbracelet/fang/blob/d89b30af/go.mod#L12-L13)

### Shell Completions

Cobra provides shell completion by default. Fang allows disabling via `WithoutCompletions`.

Sources: [README.md25](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L25-L25)

### Signal Handling

`WithNotifySignal` option configures which OS signals trigger context cancellation, enabling graceful shutdown patterns.

Sources: High-level Diagram 1

Dependency Overview
-------------------

| Dependency | Purpose | Used By |
| --- | --- | --- |
| `spf13/cobra` | Command-line framework | Core execution engine |
| `lipgloss/v2` | Terminal styling | Theme system, help rendering, error formatting |
| `colorprofile` | Terminal capability detection | Adaptive color output |
| `muesli/mango-cobra` | Manpage generation | `man` command |
| `muesli/roff` | Roff formatting | Manpage output |
| `x/term` | Terminal control | Signal handling, output detection |
| `x/exp/charmtone` | Color palettes | Default color scheme |
| `x/exp/golden` | Golden file testing | Test infrastructure |

All dependencies are from two ecosystems: **Cobra** (command parsing) and **Charm** (terminal presentation).

Sources: [go.mod5-19](https://github.com/charmbracelet/fang/blob/d89b30af/go.mod#L5-L19)
 High-level Diagram 4

Testing Strategy
----------------

Fang uses **golden file testing** to validate all output:

Test scenarios cover:

*   Help output with various command structures
*   Error messages with different error types
*   Version information display
*   Manpage generation
*   Multiline descriptions and examples

Each scenario has corresponding `.golden` files in `testdata/` directory containing expected terminal output.

Sources: [testdata/TestSetup/with\_subcommands/help-sub-sub.golden1-16](https://github.com/charmbracelet/fang/blob/d89b30af/testdata/TestSetup/with_subcommands/help-sub-sub.golden#L1-L16)
 High-level Diagram 6

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/fang/1-overview#overview)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/fang/1-overview#purpose-and-scope)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/fang/1-overview#architecture-overview)
    
*   [Core Components and Code Mapping](https://deepwiki.com/charmbracelet/fang/1-overview#core-components-and-code-mapping)
    
*   [Execution Entry Point](https://deepwiki.com/charmbracelet/fang/1-overview#execution-entry-point)
    
*   [Help Rendering System](https://deepwiki.com/charmbracelet/fang/1-overview#help-rendering-system)
    
*   [Theme and Styling Engine](https://deepwiki.com/charmbracelet/fang/1-overview#theme-and-styling-engine)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/fang/1-overview#error-handling)
    
*   [Execution Lifecycle](https://deepwiki.com/charmbracelet/fang/1-overview#execution-lifecycle)
    
*   [Feature Integration Points](https://deepwiki.com/charmbracelet/fang/1-overview#feature-integration-points)
    
*   [Version Management](https://deepwiki.com/charmbracelet/fang/1-overview#version-management)
    
*   [Manpage Generation](https://deepwiki.com/charmbracelet/fang/1-overview#manpage-generation)
    
*   [Shell Completions](https://deepwiki.com/charmbracelet/fang/1-overview#shell-completions)
    
*   [Signal Handling](https://deepwiki.com/charmbracelet/fang/1-overview#signal-handling)
    
*   [Dependency Overview](https://deepwiki.com/charmbracelet/fang/1-overview#dependency-overview)
    
*   [Testing Strategy](https://deepwiki.com/charmbracelet/fang/1-overview#testing-strategy)
