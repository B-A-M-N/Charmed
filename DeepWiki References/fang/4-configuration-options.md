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

Configuration Options
=====================

Relevant source files

*   [fang.go](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go)
    
*   [fang\_test.go](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go)
    
*   [testdata/TestSetup/without\_version.golden](https://github.com/charmbracelet/fang/blob/d89b30af/testdata/TestSetup/without_version.golden)
    

This document covers the configuration system in the fang library, including all available options, the functional options pattern implementation, and how to customize CLI behavior through configuration. For information about styling and theming implementation details, see [Styling and Themes](https://deepwiki.com/charmbracelet/fang/3.3-styling-and-themes)
. For core execution mechanics, see [Execute Function](https://deepwiki.com/charmbracelet/fang/3.1-fang.execute-function)
.

Overview
--------

Fang uses the functional options pattern to provide flexible configuration of CLI enhancements. Configuration is applied during the `Execute` call and affects help rendering, error handling, feature enablement, and visual styling.

Configuration Architecture
--------------------------

The configuration system is built around a `settings` struct and `Option` functions that modify it:

Sources: [fang.go26-36](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L26-L36)
 [fang.go92-101](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L92-L101)

Core Settings Structure
-----------------------

The internal `settings` struct contains all configurable options:

| Field | Type | Default | Purpose |
| --- | --- | --- | --- |
| `completions` | `bool` | `true` | Enable shell completion generation |
| `manpages` | `bool` | `true` | Enable man page generation |
| `version` | `string` | auto-detected | Version string display |
| `commit` | `string` | auto-detected | Git commit SHA |
| `colorscheme` | `ColorSchemeFunc` | `DefaultColorScheme` | Color scheme provider |
| `errHandler` | `ErrorHandler` | `DefaultErrorHandler` | Error formatting function |

Sources: [fang.go26-33](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L26-L33)

Feature Control Options
-----------------------

### Completion Management

By default, fang enables cobra's default completion command. This option disables it entirely.

**Implementation:**

*   Sets `settings.completions = false`
*   Calls `root.CompletionOptions.DisableDefaultCmd = true` during execution

Sources: [fang.go38-43](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L38-L43)
 [fang.go131-135](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L131-L135)

### Man Page Generation

Controls whether fang adds a hidden `man` command for generating man pages using the mango-cobra integration.

**Implementation:**

*   Sets `settings.manpages = false`
*   Skips adding the hidden "man" command during execution

Sources: [fang.go45-50](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L45-L50)
 [fang.go110-129](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L110-L129)

Version Configuration
---------------------

### Version String

Overrides automatic version detection from build info.

**Auto-detection behavior:**

*   Attempts to read `debug.BuildInfo` for version information
*   Falls back to "unknown (built from source)" if unavailable
*   Appends commit SHA if available

Sources: [fang.go70-75](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L70-L75)
 [fang.go137-147](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L137-L147)

### Commit SHA

Provides commit information that gets appended to version strings. If commit is 7+ characters, only the first 7 are used.

Sources: [fang.go77-82](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L77-L82)
 [fang.go145-147](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L145-L147)

Visual Configuration
--------------------

### Color Scheme Function

Sources: [fang.go52-57](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L52-L57)
 [theme.go23-24](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L23-L24)
 [theme.go42-62](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L42-L62)

Accepts a function that receives a light/dark detector and returns a `ColorScheme`. Built-in options:

*   `DefaultColorScheme`: Rich colors with charm-tone palette
*   `AnsiColorScheme`: Basic ANSI colors for compatibility

**Legacy Support:**

Sources: [fang.go52-68](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L52-L68)
 [theme.go42-82](https://github.com/charmbracelet/fang/blob/d89b30af/theme.go#L42-L82)

Error Handling Configuration
----------------------------

### Custom Error Handler

Replaces the default error handler with custom formatting logic. The handler receives:

*   `w io.Writer`: Output destination
*   `styles Styles`: Current styling configuration
*   `err error`: The error to format

**Default Handler Behavior:**

*   Renders styled error header and message
*   Detects usage errors and suggests `--help`
*   Uses terminal-aware color formatting

Sources: [fang.go20-21](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L20-L21)
 [fang.go84-89](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L84-L89)
 [help.go109-141](https://github.com/charmbracelet/fang/blob/d89b30af/help.go#L109-L141)

Configuration Flow
------------------

Sources: [fang.go92-157](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L92-L157)

Usage Patterns
--------------

### Basic Configuration

### Version Management

### Visual Customization

### Error Handling

Sources: [fang.go92-157](https://github.com/charmbracelet/fang/blob/d89b30af/fang.go#L92-L157)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Configuration Options](https://deepwiki.com/charmbracelet/fang/4-configuration-options#configuration-options)
    
*   [Overview](https://deepwiki.com/charmbracelet/fang/4-configuration-options#overview)
    
*   [Configuration Architecture](https://deepwiki.com/charmbracelet/fang/4-configuration-options#configuration-architecture)
    
*   [Core Settings Structure](https://deepwiki.com/charmbracelet/fang/4-configuration-options#core-settings-structure)
    
*   [Feature Control Options](https://deepwiki.com/charmbracelet/fang/4-configuration-options#feature-control-options)
    
*   [Completion Management](https://deepwiki.com/charmbracelet/fang/4-configuration-options#completion-management)
    
*   [Man Page Generation](https://deepwiki.com/charmbracelet/fang/4-configuration-options#man-page-generation)
    
*   [Version Configuration](https://deepwiki.com/charmbracelet/fang/4-configuration-options#version-configuration)
    
*   [Version String](https://deepwiki.com/charmbracelet/fang/4-configuration-options#version-string)
    
*   [Commit SHA](https://deepwiki.com/charmbracelet/fang/4-configuration-options#commit-sha)
    
*   [Visual Configuration](https://deepwiki.com/charmbracelet/fang/4-configuration-options#visual-configuration)
    
*   [Color Scheme Function](https://deepwiki.com/charmbracelet/fang/4-configuration-options#color-scheme-function)
    
*   [Error Handling Configuration](https://deepwiki.com/charmbracelet/fang/4-configuration-options#error-handling-configuration)
    
*   [Custom Error Handler](https://deepwiki.com/charmbracelet/fang/4-configuration-options#custom-error-handler)
    
*   [Configuration Flow](https://deepwiki.com/charmbracelet/fang/4-configuration-options#configuration-flow)
    
*   [Usage Patterns](https://deepwiki.com/charmbracelet/fang/4-configuration-options#usage-patterns)
    
*   [Basic Configuration](https://deepwiki.com/charmbracelet/fang/4-configuration-options#basic-configuration)
    
*   [Version Management](https://deepwiki.com/charmbracelet/fang/4-configuration-options#version-management)
    
*   [Visual Customization](https://deepwiki.com/charmbracelet/fang/4-configuration-options#visual-customization)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/fang/4-configuration-options#error-handling)
