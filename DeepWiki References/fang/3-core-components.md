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

Core Components
===============

Relevant source files

*   [fang.go](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go)
    
*   [help.go](https://github.com/charmbracelet/fang/blob/d89b30af/help.go)
    
*   [theme.go](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go)
    

This document provides an overview of the core architectural components that make up the fang library. It describes the main modules, their responsibilities, and how they interact to enhance Cobra CLI applications with styled output, custom help rendering, and additional features.

For detailed information about specific subsystems, see the child pages:

*   [fang.Execute Function](https://deepwiki.com/charmbracelet/fang/3.1-fang.execute-function)
     - Main orchestrator entry point
*   [Help Rendering System](https://deepwiki.com/charmbracelet/fang/3.2-help-rendering-system)
     - Help output generation and styling
*   [Styling and Themes](https://deepwiki.com/charmbracelet/fang/3.3-styling-and-themes)
     - Color schemes and style application
*   [Error Handling](https://deepwiki.com/charmbracelet/fang/3.4-error-handling)
     - Error output formatting

Component Overview
------------------

The fang library consists of five primary components organized in three source files:

| Component | Source File | Primary Types/Functions | Purpose |
| --- | --- | --- | --- |
| **Main Orchestrator** | `fang.go` | `Execute`, `settings`, `Option` | Entry point that coordinates all enhancements and executes commands |
| **Help Renderer** | `help.go` | `helpFn`, `styleUsage`, `styleExamples` | Generates styled help output with syntax highlighting |
| **Error Handler** | `help.go` | `DefaultErrorHandler`, `isUsageError` | Produces styled error messages with contextual hints |
| **Theming System** | `theme.go` | `ColorScheme`, `Styles`, `makeStyles` | Manages color schemes and style generation |
| **Version Builder** | `fang.go` | `buildVersion` | Extracts version information from build metadata |

Sources: [fang.go1-209](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L1-L209)
 [help.go1-482](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L1-L482)
 [theme.go1-209](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L1-L209)

Architecture Flow
-----------------

**Diagram: Component Interaction Flow**

This diagram shows how `fang.Execute` orchestrates the enhancement and execution phases. The function applies settings through various enhancement mechanisms before delegating to Cobra's execution engine, which then triggers the help or error handlers as needed.

Sources: [fang.go110-179](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L179)
 [help.go42-132](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L42-L132)
 [theme.go115-199](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L115-L199)

The Main Orchestrator
---------------------

The `Execute` function [fang.go110-179](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L179)
 serves as the library's entry point and main coordinator. It accepts a `context.Context`, a `*cobra.Command`, and variadic `Option` functions.

### Settings Configuration

The `settings` struct [fang.go27-36](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L27-L36)
 contains all configurable options:

    type settings struct {
        completions bool
        manpages    bool
        skipVersion bool
        version     string
        commit      string
        colorscheme ColorSchemeFunc
        errHandler  ErrorHandler
        signals     []os.Signal
    }
    

Configuration is applied through the options pattern using `Option` functions [fang.go39-107](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L39-L107)
:

| Option Function | Purpose | Default |
| --- | --- | --- |
| `WithoutCompletions()` | Disables shell completion command | Enabled |
| `WithoutManpage()` | Disables man page generation command | Enabled |
| `WithColorSchemeFunc(cs)` | Sets custom color scheme function | `DefaultColorScheme` |
| `WithVersion(version)` | Sets explicit version string | Auto-detected |
| `WithoutVersion()` | Disables version flag | Enabled |
| `WithCommit(commit)` | Sets commit SHA | Auto-detected |
| `WithErrorHandler(handler)` | Sets custom error handler | `DefaultErrorHandler` |
| `WithNotifySignal(signals...)` | Configures signal handling | None |

Sources: [fang.go27-107](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L27-L107)

### Enhancement Sequence

The `Execute` function performs the following enhancement steps in order:

**Diagram: Enhancement Sequence in fang.Execute**

Sources: [fang.go110-179](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L179)

Help Rendering Pipeline
-----------------------

The help rendering system is the library's most complex subsystem. The `helpFn` function [help.go42-108](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L42-L108)
 orchestrates a three-stage pipeline:

### Stage 1: Content Extraction

Three evaluator functions extract metadata from the `cobra.Command`:

| Function | Line Range | Returns | Purpose |
| --- | --- | --- | --- |
| `evalGroups` | [help.go435-444](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L435-L444) | `map[string]string`, `[]string` | Command group IDs and titles |
| `evalCmds` | [help.go416-433](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L416-L433) | `map[string]map[string]string`, `[]string` | Styled subcommand names and descriptions |
| `evalFlags` | [help.go366-412](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L366-L412) | `map[string]string`, `[]string` | Styled flag names and usage text |

### Stage 2: Content Styling

Styling functions apply terminal colors and formatting:

| Function | Line Range | Purpose |
| --- | --- | --- |
| `styleUsage` | [help.go162-222](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L162-L222) | Styles the usage line with program name, commands, args, flags |
| `styleExamples` | [help.go224-244](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L224-L244) | Splits and styles example block |
| `styleExample` | [help.go246-364](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L246-L364) | Token-based syntax highlighting for individual examples |

The `styleExample` function performs sophisticated parsing to distinguish between program names, commands, flags, arguments, quoted strings, redirections, and shell operators.

**Diagram: Help Rendering Pipeline Functions**

### Stage 3: Layout Rendering

Layout functions calculate spacing and render aligned output:

*   `calculateSpace` [help.go459-466](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L459-L466)
     determines column alignment based on the longest key
*   `renderGroup` [help.go447-457](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L447-L457)
     outputs a titled group with key-value pairs
*   `width` [help.go30-40](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L30-L40)
     returns terminal width (capped at 120 columns)

Sources: [help.go42-482](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L42-L482)

Theming Architecture
--------------------

The theming system implements a two-stage transformation: `ColorScheme` → `Styles`.

**Diagram: Theming System Transformation**

### ColorScheme Struct

The `ColorScheme` struct [theme.go16-33](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L16-L33)
 defines semantic colors:

| Field | Purpose |
| --- | --- |
| `Base` | Default text color |
| `Title` | Section titles (uppercase) |
| `Description` | Flag and command descriptions |
| `Codeblock` | Background for code examples |
| `Program` | Program name in examples |
| `DimmedArgument` | Optional arguments like `[flags]` |
| `Comment` | Comment lines starting with `#` |
| `Flag` | Flag names like `--help` |
| `FlagDefault` | Default values in flag descriptions |
| `Command` | Subcommand names |
| `QuotedString` | String literals in examples |
| `Argument` | Command arguments |
| `ErrorHeader` | Error label background/foreground |
| `ErrorDetails` | Error message text |

### Styles Struct

The `makeStyles` function [theme.go123-199](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L123-L199)
 transforms a `ColorScheme` into a `Styles` struct [theme.go85-95](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L85-L95)
 containing ready-to-use `lipgloss.Style` objects:

*   `Styles.Text` - Base text style
*   `Styles.Title` - Section titles (uppercase, bold, padded)
*   `Styles.ErrorHeader` - "ERROR" badge
*   `Styles.ErrorText` - Error message with titlecase transform
*   `Styles.FlagDescription` - Flag usage text
*   `Styles.FlagDefault` - Default value parenthetical
*   `Styles.Codeblock` - Nested styles for code blocks
*   `Styles.Program` - Nested styles for program elements

The `Codeblock` and `Program` nested structs [theme.go98-113](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L98-L113)
 provide context-specific styling (e.g., a flag looks different inside a code block vs. inline text).

Sources: [theme.go16-199](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L16-L199)

Error Handling
--------------

The `DefaultErrorHandler` function [help.go111-132](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L111-L132)
 handles error output with the following logic:

**Diagram: Error Handler Decision Flow**

The `isUsageError` function [help.go136-150](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L136-L150)
 detects common usage error patterns by checking if the error message starts with specific prefixes like `"unknown flag:"`, `"flag needs an argument:"`, or `"invalid argument"`.

Sources: [help.go111-150](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L111-L150)

Type Definitions
----------------

Key type aliases facilitate extensibility:

    type ErrorHandler = func(w io.Writer, styles Styles, err error)
    type ColorSchemeFunc = func(lipgloss.LightDarkFunc) ColorScheme
    

The `ErrorHandler` type [fang.go22](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L22-L22)
 allows custom error formatting. The `ColorSchemeFunc` type [fang.go25](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L25-L25)
 enables dynamic color scheme selection based on terminal capabilities.

Sources: [fang.go22-25](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L22-L25)

Version Management
------------------

The `buildVersion` function [fang.go181-196](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L181-L196)
 constructs version strings using this precedence:

1.  Explicit version from `WithVersion()` option
2.  Auto-detected from `debug.ReadBuildInfo()` if built as a module
3.  Fallback to `"unknown (built from source)"`

If a commit SHA is available (from `WithCommit()` or VCS build info), it appends the first 7 characters in parentheses.

The `getKey` helper function [fang.go198-208](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L198-L208)
 extracts build settings like `"vcs.revision"` from the build info.

Sources: [fang.go181-208](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L181-L208)

Integration with Cobra
----------------------

Fang integrates with Cobra through several hook points:

| Cobra Property | Fang Configuration | Purpose |
| --- | --- | --- |
| `root.SetHelpFunc()` | Custom `helpFunc` closure | Replaces default help with styled renderer |
| `root.Version` | `buildVersion(opts)` | Populates `--version` flag output |
| `root.SilenceUsage` | Set to `true` | Prevents duplicate usage output |
| `root.SilenceErrors` | Set to `true` | Allows custom error handler |
| `root.AddCommand()` | Hidden `man` command | Adds manpage generation |
| `root.CompletionOptions` | `DisableDefaultCmd` | Controls completion command visibility |

Sources: [fang.go135-165](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L135-L165)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core Components](https://deepwiki.com/charmbracelet/fang/3-core-components#core-components)
    
*   [Component Overview](https://deepwiki.com/charmbracelet/fang/3-core-components#component-overview)
    
*   [Architecture Flow](https://deepwiki.com/charmbracelet/fang/3-core-components#architecture-flow)
    
*   [The Main Orchestrator](https://deepwiki.com/charmbracelet/fang/3-core-components#the-main-orchestrator)
    
*   [Settings Configuration](https://deepwiki.com/charmbracelet/fang/3-core-components#settings-configuration)
    
*   [Enhancement Sequence](https://deepwiki.com/charmbracelet/fang/3-core-components#enhancement-sequence)
    
*   [Help Rendering Pipeline](https://deepwiki.com/charmbracelet/fang/3-core-components#help-rendering-pipeline)
    
*   [Stage 1: Content Extraction](https://deepwiki.com/charmbracelet/fang/3-core-components#stage-1-content-extraction)
    
*   [Stage 2: Content Styling](https://deepwiki.com/charmbracelet/fang/3-core-components#stage-2-content-styling)
    
*   [Stage 3: Layout Rendering](https://deepwiki.com/charmbracelet/fang/3-core-components#stage-3-layout-rendering)
    
*   [Theming Architecture](https://deepwiki.com/charmbracelet/fang/3-core-components#theming-architecture)
    
*   [ColorScheme Struct](https://deepwiki.com/charmbracelet/fang/3-core-components#colorscheme-struct)
    
*   [Styles Struct](https://deepwiki.com/charmbracelet/fang/3-core-components#styles-struct)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/fang/3-core-components#error-handling)
    
*   [Type Definitions](https://deepwiki.com/charmbracelet/fang/3-core-components#type-definitions)
    
*   [Version Management](https://deepwiki.com/charmbracelet/fang/3-core-components#version-management)
    
*   [Integration with Cobra](https://deepwiki.com/charmbracelet/fang/3-core-components#integration-with-cobra)
