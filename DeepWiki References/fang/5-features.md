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

Features
========

Relevant source files

*   [LICENSE.md](https://github.com/charmbracelet/fang/blob/d89b30af/LICENSE.md?plain=1)
    
*   [README.md](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1)
    
*   [fang.go](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go)
    

This document provides an overview of the key features that fang adds to Cobra applications. These features transform standard Cobra commands into polished, user-friendly CLI tools with consistent styling, automatic documentation generation, and enhanced usability.

For detailed information about specific features, see:

*   Version management implementation: [Version Management](https://deepwiki.com/charmbracelet/fang/5.1-version-management)
    
*   Manpage generation details: [Manpage Generation](https://deepwiki.com/charmbracelet/fang/5.2-manpage-generation)
    
*   Shell completion configuration: [Shell Completions](https://deepwiki.com/charmbracelet/fang/5.3-shell-completions)
    

For information about configuring these features, see [Configuration Options](https://deepwiki.com/charmbracelet/fang/4-configuration-options)
.

Feature Categories
------------------

Fang enhances Cobra applications across four main categories:

| Category | Features | Default State |
| --- | --- | --- |
| **Output Styling** | Styled help pages, styled error messages, custom theming | Always enabled |
| **Version Information** | Automatic version detection, commit hash display | Enabled (can disable) |
| **Documentation** | Manpage generation, shell completions | Enabled (can disable) |
| **Runtime Behavior** | Signal handling, silent usage on errors | Configurable |

Sources: [README.md18-28](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L18-L28)
 [fang.go27-36](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L27-L36)

Feature Architecture
--------------------

The following diagram shows how fang's features are organized and which code components implement them:

Sources: [fang.go27-36](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L27-L36)
 [fang.go110-179](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L179)
 [fang.go181-196](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L181-L196)

Core Features
-------------

### Styled Help Pages

Fang replaces Cobra's default help function with a custom renderer that provides:

*   **Syntax-highlighted examples**: Code examples are tokenized and styled with different colors for programs, flags, arguments, and operators
*   **Organized command groups**: Commands are displayed in user-defined groups with visual separation
*   **Aligned layouts**: Automatic spacing calculation ensures clean alignment across different terminal widths
*   **Rich typography**: Usage lines, descriptions, and section headers receive distinct styling

The help system is always enabled and uses the configured theme. Implementation details are in [Help Rendering System](https://deepwiki.com/charmbracelet/fang/3.2-help-rendering-system)
.

Sources: [fang.go130-133](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L130-L133)
 [README.md20](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L20-L20)

### Styled Error Messages

Error output receives consistent styling to distinguish errors from normal output:

*   **Error prefixes**: Errors are displayed with styled "Error:" prefixes
*   **Usage errors**: Automatically detects usage errors (invalid flags, missing arguments) and displays them with suggestions
*   **Custom handlers**: Applications can provide custom error handling logic via `WithErrorHandler`

The default error handler implementation is in [error.go](https://github.com/charmbracelet/fang/blob/d89b30af/error.go)

Sources: [fang.go173-176](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L173-L176)
 [README.md21](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L21-L21)

### Automatic Version Information

Fang automatically configures the `--version` and `-v` flags with version information:

The version string follows the format: `version (commit)` where the commit hash is truncated to 7 characters (defined by `shaLen` constant at [fang.go19](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L19-L19)
). If no version is provided via `WithVersion()`, fang attempts to read it from Go's `debug.BuildInfo()` using the `info.Main.Version` field and `vcs.revision` build setting.

Sources: [fang.go19](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L19-L19)
 [fang.go137-139](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L137-L139)
 [fang.go181-196](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L181-L196)
 [fang.go198-208](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L198-L208)
 [README.md22](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L22-L22)

### Silent Usage Output

Fang modifies Cobra's default behavior to provide a better user experience on errors:

*   Sets `root.SilenceUsage = true` to prevent help text from appearing after every error
*   Sets `root.SilenceErrors = true` to give fang control over error formatting
*   Errors are still printed, but through the styled error handler

This prevents the common issue where users encounter a wall of help text after making a simple mistake.

Sources: [fang.go135-136](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L135-L136)
 [README.md27](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L27-L27)

### Theming System

All styled output uses a configurable color scheme that adapts to terminal capabilities:

*   **Automatic detection**: Detects light/dark terminal backgrounds and terminal color support
*   **Two built-in schemes**: `DefaultColorScheme` (charmtone colors) and `AnsiColorScheme` (basic ANSI colors)
*   **Custom themes**: Applications can provide custom themes via `WithColorSchemeFunc` or deprecated `WithTheme`

Theme configuration is detailed in [Styling and Themes](https://deepwiki.com/charmbracelet/fang/3.3-styling-and-themes)
.

Sources: [fang.go55-70](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L55-L70)
 [README.md26](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L26-L26)

Optional Features
-----------------

The following features can be enabled or disabled through configuration options:

### Manpage Generation

A hidden `man` command is automatically added to generate Unix manual pages when `opts.manpages` is true (default):

| Command Property | Value |
| --- | --- |
| `Use` | `"man"` |
| `Short` | `"Generates manpages"` |
| `Hidden` | `true` |
| `DisableFlagsInUseLine` | `true` |
| `Args` | `cobra.NoArgs` |
| Implementation | Uses `mango.NewManPage()` and `roff.NewDocument()` |

The command generates a single manpage (section 1) for the entire command tree using [github.com/muesli/mango-cobra](https://github.com/charmbracelet/fang/blob/d89b30af/github.com/muesli/mango-cobra)
 and outputs it to stdout. Disable with `WithoutManpage()`. See [Manpage Generation](https://deepwiki.com/charmbracelet/fang/5.2-manpage-generation)
 for details.

Sources: [fang.go142-161](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L142-L161)
 [README.md23-24](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L23-L24)

### Shell Completions

Cobra's built-in completion command is enabled by default, allowing users to generate shell completion scripts for bash, zsh, fish, and PowerShell.

Disable with `WithoutCompletions()`. See [Shell Completions](https://deepwiki.com/charmbracelet/fang/5.3-shell-completions)
 for details.

Sources: [fang.go163-165](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L163-L165)
 [README.md25](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1#L25-L25)

### Version Flag

The automatic `--version` flag can be disabled entirely with `WithoutVersion()`, useful for applications that want to handle versioning differently.

Sources: [fang.go80-85](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L80-L85)
 [fang.go137-139](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L137-L139)

### Signal Handling

Applications can configure signals (like `SIGINT`, `SIGTERM`) to gracefully cancel the execution context:

When configured via `WithNotifySignal()`, fang wraps the provided context with `signal.NotifyContext()`. When any of the specified signals are received, the context passed to `root.ExecuteContext()` is automatically cancelled, allowing commands to detect and respond to interruption signals through the context's `Done()` channel.

Sources: [fang.go35](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L35-L35)
 [fang.go103-107](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L103-L107)
 [fang.go167-171](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L167-L171)

Feature Enablement Flow
-----------------------

The following diagram shows the precise order in which features are applied during `fang.Execute`:

Sources: [fang.go110-179](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L179)

Configuration Summary
---------------------

Features can be controlled through `Option` functions passed to `fang.Execute`:

| Feature | Enabled By Default | Disable Option | Configure Option |
| --- | --- | --- | --- |
| Styled Help | Yes (always) | N/A | `WithColorSchemeFunc` |
| Styled Errors | Yes (always) | N/A | `WithErrorHandler`, `WithColorSchemeFunc` |
| Version Flag | Yes | `WithoutVersion()` | `WithVersion`, `WithCommit` |
| Manpage Command | Yes | `WithoutManpage()` | N/A |
| Completion Command | Yes | `WithoutCompletions()` | N/A |
| Signal Handling | No  | N/A | `WithNotifySignal` |
| Theming | Yes (default theme) | N/A | `WithColorSchemeFunc`, `WithTheme` |

Sources: [fang.go38-107](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L38-L107)

Platform Support
----------------

### Windows Virtual Terminal Processing

On Windows platforms, fang automatically enables virtual terminal processing to ensure ANSI escape sequences render correctly:

This is a no-op on Unix-like systems. The implementation is platform-specific and located in [term\_windows.go](https://github.com/charmbracelet/fang/blob/d89b30af/term_windows.go)
 and [term\_unix.go](https://github.com/charmbracelet/fang/blob/d89b30af/term_unix.go)

Sources: [fang.go122-128](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L122-L128)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Features](https://deepwiki.com/charmbracelet/fang/5-features#features)
    
*   [Feature Categories](https://deepwiki.com/charmbracelet/fang/5-features#feature-categories)
    
*   [Feature Architecture](https://deepwiki.com/charmbracelet/fang/5-features#feature-architecture)
    
*   [Core Features](https://deepwiki.com/charmbracelet/fang/5-features#core-features)
    
*   [Styled Help Pages](https://deepwiki.com/charmbracelet/fang/5-features#styled-help-pages)
    
*   [Styled Error Messages](https://deepwiki.com/charmbracelet/fang/5-features#styled-error-messages)
    
*   [Automatic Version Information](https://deepwiki.com/charmbracelet/fang/5-features#automatic-version-information)
    
*   [Silent Usage Output](https://deepwiki.com/charmbracelet/fang/5-features#silent-usage-output)
    
*   [Theming System](https://deepwiki.com/charmbracelet/fang/5-features#theming-system)
    
*   [Optional Features](https://deepwiki.com/charmbracelet/fang/5-features#optional-features)
    
*   [Manpage Generation](https://deepwiki.com/charmbracelet/fang/5-features#manpage-generation)
    
*   [Shell Completions](https://deepwiki.com/charmbracelet/fang/5-features#shell-completions)
    
*   [Version Flag](https://deepwiki.com/charmbracelet/fang/5-features#version-flag)
    
*   [Signal Handling](https://deepwiki.com/charmbracelet/fang/5-features#signal-handling)
    
*   [Feature Enablement Flow](https://deepwiki.com/charmbracelet/fang/5-features#feature-enablement-flow)
    
*   [Configuration Summary](https://deepwiki.com/charmbracelet/fang/5-features#configuration-summary)
    
*   [Platform Support](https://deepwiki.com/charmbracelet/fang/5-features#platform-support)
    
*   [Windows Virtual Terminal Processing](https://deepwiki.com/charmbracelet/fang/5-features#windows-virtual-terminal-processing)
