Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/log](https://github.com/charmbracelet/log "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 20 April 2025 ([cf6e86](https://github.com/charmbracelet/log/commits/cf6e8671)
)

*   [Overview](https://deepwiki.com/charmbracelet/log/1-overview)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/log/2-core-architecture)
    
*   [Logger Structure](https://deepwiki.com/charmbracelet/log/2.1-logger-structure)
    
*   [Log Levels](https://deepwiki.com/charmbracelet/log/2.2-log-levels)
    
*   [Global Logger](https://deepwiki.com/charmbracelet/log/2.3-global-logger)
    
*   [Usage Guide](https://deepwiki.com/charmbracelet/log/3-usage-guide)
    
*   [Basic Logging](https://deepwiki.com/charmbracelet/log/3.1-basic-logging)
    
*   [Contextual Logging](https://deepwiki.com/charmbracelet/log/3.2-contextual-logging)
    
*   [Helper Functions](https://deepwiki.com/charmbracelet/log/3.3-helper-functions)
    
*   [Formatters](https://deepwiki.com/charmbracelet/log/4-formatters)
    
*   [Text Formatter](https://deepwiki.com/charmbracelet/log/4.1-text-formatter)
    
*   [JSON Formatter](https://deepwiki.com/charmbracelet/log/4.2-json-formatter)
    
*   [Logfmt Formatter](https://deepwiki.com/charmbracelet/log/4.3-logfmt-formatter)
    
*   [Styling and Customization](https://deepwiki.com/charmbracelet/log/5-styling-and-customization)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/log/6-configuration-options)
    
*   [Integration](https://deepwiki.com/charmbracelet/log/7-integration)
    
*   [Standard Library Integration](https://deepwiki.com/charmbracelet/log/7.1-standard-library-integration)
    
*   [Slog Integration](https://deepwiki.com/charmbracelet/log/7.2-slog-integration)
    
*   [Advanced Topics](https://deepwiki.com/charmbracelet/log/8-advanced-topics)
    
*   [Contributing](https://deepwiki.com/charmbracelet/log/9-contributing)
    

Menu

Styling and Customization
=========================

Relevant source files

*   [README.md](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1)
    
*   [examples/styles/styles.go](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go)
    
*   [examples/styles/styles.tape](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.tape)
    
*   [styles.go](https://github.com/charmbracelet/log/blob/cf6e8671/styles.go)
    

This document covers how to customize the appearance of log output in the charmbracelet/log library. Styling is a key feature that makes logs more readable and visually appealing, but it only applies when using the text formatter—JSON and logfmt formatters ignore styling. For information about formatters, see [Formatters](https://deepwiki.com/charmbracelet/log/4-formatters)
.

Overview of Styling
-------------------

The charmbracelet/log library uses [Lipgloss](https://github.com/charmbracelet/log/blob/cf6e8671/Lipgloss)
 to provide colorful, styled log output. Styling allows you to customize the appearance of different log elements, such as level indicators, timestamps, keys, and values.

Sources: [styles.go9-40](https://github.com/charmbracelet/log/blob/cf6e8671/styles.go#L9-L40)

The Styles Struct
-----------------

The core of styling in charmbracelet/log is the `Styles` struct, which defines styles for all components of a log message.

Here's what each field controls:

| Field | Purpose |
| --- | --- |
| `Timestamp` | Styles timestamps when `ReportTimestamp` is enabled |
| `Caller` | Styles the caller information when `ReportCaller` is enabled |
| `Prefix` | Styles the logger prefix set with `WithPrefix` |
| `Message` | Styles the main log message |
| `Key` | Default style for all keys in key-value pairs |
| `Value` | Default style for all values in key-value pairs |
| `Separator` | Styles separators between parts of the log message |
| `Levels` | Map of styles for each log level's display (DEBUG, INFO, etc.) |
| `Keys` | Map to override styles for specific keys |
| `Values` | Map to override value styles for specific keys |

Sources: [styles.go9-40](https://github.com/charmbracelet/log/blob/cf6e8671/styles.go#L9-L40)

Default Styles
--------------

The library provides default styles through the `DefaultStyles()` function:

By default:

*   Level indicators are bold and colored according to severity
*   Timestamps have no special styling
*   Caller information is displayed with faint text
*   Prefixes are bold and faint
*   Keys in key-value pairs are displayed with faint text
*   Values have no special styling

Sources: [styles.go42-82](https://github.com/charmbracelet/log/blob/cf6e8671/styles.go#L42-L82)

Customizing Styles
------------------

You can customize styles by:

1.  Getting the default styles
2.  Modifying specific styles
3.  Applying the modified styles to a logger

For example, to customize the error level style and style for a specific key:

Sources: [examples/styles/styles.go12-23](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go#L12-L23)
 [README.md212-225](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L212-L225)

### Customizing Level Styles

Level styles control the appearance of level indicators (DEBUG, INFO, etc.). You can completely customize how they look:

Sources: [examples/styles/styles.go14-18](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go#L14-L18)
 [README.md214-218](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L214-L218)

### Customizing Key and Value Styles

You can set styles for specific keys and their values:

This allows you to highlight important fields in your logs.

Sources: [examples/styles/styles.go19-20](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go#L19-L20)
 [README.md219-220](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L219-L220)

Applying Custom Styles
----------------------

Apply your custom styles using the `SetStyles` method:

Sources: [examples/styles/styles.go21-22](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go#L21-L22)
 [README.md221-222](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L221-L222)

Style Inheritance in Logger Hierarchy
-------------------------------------

When you create child loggers using `With()` or `WithPrefix()`, they inherit the styles from their parent logger:

This means you only need to set up styles once, at the root logger level.

Sources: [README.md235-246](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L235-L246)

Terminal Detection
------------------

It's important to note that styling is automatically disabled when the output is not a terminal (TTY). This ensures that logs written to files or piped to other applications don't contain ANSI color codes or other formatting that could cause issues.

Sources: [README.md201-202](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L201-L202)

Example of Styled Output
------------------------

Here's what you can achieve with custom styles:

| Component | Default Style | Custom Example |
| --- | --- | --- |
| Level (ERROR) | Bold red "ERROR" | Red background with white text "ERROR!!" |
| Keys | Faint text | Red text for "err" key |
| Values | Regular text | Bold text for values of "err" key |
| Timestamp | Regular text | Italic gray text |
| Message | Regular text | Bold text |

By using these styling capabilities, you can create log output that is not only informative but also easy to scan visually, with important information highlighted appropriately.

Sources: [README.md212-225](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L212-L225)
 [examples/styles/styles.go12-23](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go#L12-L23)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Styling and Customization](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#styling-and-customization)
    
*   [Overview of Styling](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#overview-of-styling)
    
*   [The Styles Struct](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#the-styles-struct)
    
*   [Default Styles](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#default-styles)
    
*   [Customizing Styles](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#customizing-styles)
    
*   [Customizing Level Styles](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#customizing-level-styles)
    
*   [Customizing Key and Value Styles](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#customizing-key-and-value-styles)
    
*   [Applying Custom Styles](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#applying-custom-styles)
    
*   [Style Inheritance in Logger Hierarchy](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#style-inheritance-in-logger-hierarchy)
    
*   [Terminal Detection](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#terminal-detection)
    
*   [Example of Styled Output](https://deepwiki.com/charmbracelet/log/5-styling-and-customization#example-of-styled-output)
