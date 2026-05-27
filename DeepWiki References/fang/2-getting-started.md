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

Getting Started
===============

Relevant source files

*   [LICENSE.md](https://github.com/charmbracelet/fang/blob/d89b30af/LICENSE.md?plain=1)
    
*   [README.md](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1)
    
*   [example/main.go](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go)
    
*   [fang.go](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go)
    

This guide demonstrates how to integrate Fang into a Cobra-based CLI application. It covers the minimal setup required to replace `cobra.Command.Execute()` with `fang.Execute()` and shows how to configure basic options.

For detailed information about configuration options, see [Configuration Options](https://deepwiki.com/charmbracelet/fang/4-configuration-options)
. For architectural details about `fang.Execute`, see [fang.Execute Function](https://deepwiki.com/charmbracelet/fang/3.1-fang.execute-function)
.

Overview
--------

Fang enhances Cobra applications by wrapping the standard `cobra.Command` execution flow. The integration requires minimal changes to existing Cobra code: replace the call to `cmd.Execute()` or `cmd.ExecuteContext()` with `fang.Execute()`.

### Integration Flow

The following diagram illustrates how user code integrates with Fang's execution model:

Sources: [example/main.go1-153](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L1-L153)
 [fang.go110-179](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L179)

Minimal Example
---------------

The simplest integration requires only wrapping your root `cobra.Command` with `fang.Execute`:

This minimal setup automatically provides:

*   Styled help output using the default color scheme
*   Styled error messages
*   Automatic `--version` flag (extracted from build info)
*   Hidden `man` command for manpage generation
*   Shell completion command

Sources: [README.md40-64](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L40-L64)
 [fang.go110-179](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L179)

Basic Command Structure
-----------------------

A typical Cobra command with Fang includes command metadata, flags, subcommands, and examples:

Sources: [example/main.go21-80](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L21-L80)
 [example/main.go104-134](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L104-L134)

### Code Example with Flags and Subcommands

The following pattern demonstrates a complete command structure with flags, subcommands, and groups:

| Component | Code Location | Purpose |
| --- | --- | --- |
| Root command | [example/main.go21-80](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L21-L80) | Main command definition with flags |
| Flag definitions | [example/main.go81-93](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L81-L93) | String, int, bool, duration flags |
| Command group | [example/main.go100-103](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L100-L103) | Organize subcommands visually |
| Subcommand | [example/main.go104-114](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L104-L114) | Child command with aliases |
| Nested subcommand | [example/main.go116-134](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L116-L134) | Multi-level command hierarchy |
| Fang integration | [example/main.go146-152](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L146-L152) | Execute with options |

Sources: [example/main.go1-153](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L1-L153)

Configuration with Options
--------------------------

Fang uses functional options to configure behavior. Options are passed as variadic arguments to `fang.Execute`:

### Available Options

Sources: [fang.go38-107](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L38-L107)
 [fang.go110-120](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L120)

### Common Configuration Patterns

**Signal Handling for Graceful Shutdown**

To enable context-aware signal handling, use `WithNotifySignal`:

This creates a signal-aware context that can be used in command `RunE` functions to handle graceful shutdown. See [example/main.go146-152](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L146-L152)
 and [example/main.go124-133](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L124-L133)
 for implementation.

**Custom Version Information**

If not provided, version information is automatically extracted from build metadata via `debug.ReadBuildInfo()`. See [fang.go181-196](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L181-L196)
 for version building logic.

**Minimal Configuration (Disable Features)**

Sources: [fang.go42-107](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L42-L107)
 [example/main.go146-152](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L146-L152)

Execution Lifecycle
-------------------

The following sequence shows the complete lifecycle from user invocation to output:

Sources: [fang.go110-179](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L179)
 [example/main.go146-152](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L146-L152)

Key Function Signatures
-----------------------

Understanding the primary function signatures helps when integrating Fang:

| Function | Signature | Location |
| --- | --- | --- |
| `Execute` | `func Execute(ctx context.Context, root *cobra.Command, options ...Option) error` | [fang.go110](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L110) |
| `Option` | `type Option func(*settings)` | [fang.go39](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L39-L39) |
| `WithVersion` | `func WithVersion(version string) Option` | [fang.go74-78](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L74-L78) |
| `WithNotifySignal` | `func WithNotifySignal(signals ...os.Signal) Option` | [fang.go103-107](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L103-L107) |
| `WithColorSchemeFunc` | `func WithColorSchemeFunc(cs ColorSchemeFunc) Option` | [fang.go56-60](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L56-L60) |
| `WithErrorHandler` | `func WithErrorHandler(handler ErrorHandler) Option` | [fang.go95-99](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L95-L99) |

Sources: [fang.go38-107](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L38-L107)
 [fang.go110](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L110)

Next Steps
----------

After integrating Fang into your application:

1.  **Customize appearance**: See [Styling and Themes](https://deepwiki.com/charmbracelet/fang/3.3-styling-and-themes)
     for color scheme customization
2.  **Configure features**: See [Configuration Options](https://deepwiki.com/charmbracelet/fang/4-configuration-options)
     for all available options
3.  **Add version info**: See [Version Management](https://deepwiki.com/charmbracelet/fang/5.1-version-management)
     for build-time version injection
4.  **Handle errors**: See [Error Handling](https://deepwiki.com/charmbracelet/fang/3.4-error-handling)
     for custom error formatting
5.  **Generate manpages**: See [Manpage Generation](https://deepwiki.com/charmbracelet/fang/5.2-manpage-generation)
     for documentation generation
6.  **Review example**: See [Example Application](https://deepwiki.com/charmbracelet/fang/7.2-example-application)
     for a complete reference implementation

The example application at [example/main.go1-153](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L1-L153)
 demonstrates all major features including command groups, nested subcommands, signal handling, and various flag types.

Sources: [example/main.go1-153](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L1-L153)
 [README.md1-93](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L1-L93)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Getting Started](https://deepwiki.com/charmbracelet/fang/2-getting-started#getting-started)
    
*   [Overview](https://deepwiki.com/charmbracelet/fang/2-getting-started#overview)
    
*   [Integration Flow](https://deepwiki.com/charmbracelet/fang/2-getting-started#integration-flow)
    
*   [Minimal Example](https://deepwiki.com/charmbracelet/fang/2-getting-started#minimal-example)
    
*   [Basic Command Structure](https://deepwiki.com/charmbracelet/fang/2-getting-started#basic-command-structure)
    
*   [Code Example with Flags and Subcommands](https://deepwiki.com/charmbracelet/fang/2-getting-started#code-example-with-flags-and-subcommands)
    
*   [Configuration with Options](https://deepwiki.com/charmbracelet/fang/2-getting-started#configuration-with-options)
    
*   [Available Options](https://deepwiki.com/charmbracelet/fang/2-getting-started#available-options)
    
*   [Common Configuration Patterns](https://deepwiki.com/charmbracelet/fang/2-getting-started#common-configuration-patterns)
    
*   [Execution Lifecycle](https://deepwiki.com/charmbracelet/fang/2-getting-started#execution-lifecycle)
    
*   [Key Function Signatures](https://deepwiki.com/charmbracelet/fang/2-getting-started#key-function-signatures)
    
*   [Next Steps](https://deepwiki.com/charmbracelet/fang/2-getting-started#next-steps)
