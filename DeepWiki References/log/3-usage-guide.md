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

Usage Guide
===========

Relevant source files

*   [README.md](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1)
    
*   [examples/styles/styles.go](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go)
    
*   [examples/styles/styles.tape](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.tape)
    
*   [logger\_test.go](https://github.com/charmbracelet/log/blob/cf6e8671/logger_test.go)
    
*   [pkg\_test.go](https://github.com/charmbracelet/log/blob/cf6e8671/pkg_test.go)
    

This page provides practical instructions for using the charmbracelet/log library. It covers everything from installation to basic logging operations and more advanced usage patterns. For more detailed information on specific topics, you can refer to [Basic Logging](https://deepwiki.com/charmbracelet/log/3.1-basic-logging)
, [Contextual Logging](https://deepwiki.com/charmbracelet/log/3.2-contextual-logging)
, and [Helper Functions](https://deepwiki.com/charmbracelet/log/3.3-helper-functions)
.

Installation
------------

To install the charmbracelet/log library, use Go's package manager:

Then import it in your Go files:

Sources: [README.md44-52](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L44-L52)

Logging Flow
------------

Before diving into specific usage patterns, it's helpful to understand how logging works in the library:

Sources: [README.md53-60](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L53-L60)

Basic Logging
-------------

The library comes with a global logger configured with reasonable defaults (timestamp on, level set to `info`).

### Using the Global Logger

Basic examples:

### Using Format Strings

You can use format strings (like in `fmt.Printf`) with the 'f' variants:

### Adding Key-Value Pairs

All logging functions accept optional key-value pairs:

Sources: [README.md53-60](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L53-L60)
 [README.md71-76](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L71-L76)
 [README.md138-148](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L138-L148)
 [pkg\_test.go137-184](https://github.com/charmbracelet/log/blob/cf6e8671/pkg_test.go#L137-L184)

Creating Custom Loggers
-----------------------

For more control, you can create custom loggers with different configurations.

### Basic Custom Logger

### Configured Custom Logger

### Modifying Logger Settings

You can modify a logger's configuration after creation:

Sources: [README.md91-107](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L91-L107)
 [README.md162-192](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L162-L192)
 [logger\_test.go178-205](https://github.com/charmbracelet/log/blob/cf6e8671/logger_test.go#L178-L205)

Output Formatters
-----------------

The library supports different output formats for logs:

Setting the formatter:

Examples of output:

| Formatter | Example Output |
| --- | --- |
| TextFormatter | `INFO Hello World!` |
| JSONFormatter | `{"level":"info","msg":"Hello World!"}` |
| LogfmtFormatter | `level=info msg="Hello World!"` |

Sources: [README.md194-203](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L194-L203)
 [pkg\_test.go216-225](https://github.com/charmbracelet/log/blob/cf6e8671/pkg_test.go#L216-L225)

Contextual Logging
------------------

You can create loggers with built-in context for cleaner logs in specific components.

### Using With() for Context

Example:

### Using WithPrefix

Sources: [README.md234-252](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L234-L252)
 [README.md181-205](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L181-L205)
 [pkg\_test.go186-213](https://github.com/charmbracelet/log/blob/cf6e8671/pkg_test.go#L186-L213)
 [logger\_test.go13-60](https://github.com/charmbracelet/log/blob/cf6e8671/logger_test.go#L13-L60)

Helper Functions
----------------

When creating logging utility functions, use `Helper()` to report the correct caller location:

Without `Helper()`, logs would show the line number within the helper function rather than where the helper was called from.

Sources: [README.md286-299](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L286-L299)

Style Customization
-------------------

The library uses [lipgloss](https://github.com/charmbracelet/log/blob/cf6e8671/lipgloss)
 for styling text output. You can customize these styles:

Note that styling only works when the output is a terminal, and only with the `TextFormatter`.

Sources: [README.md206-232](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L206-L232)
 [examples/styles/styles.go1-26](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go#L1-L26)

Integration with Other Logging Systems
--------------------------------------

The charmbracelet/log library provides adapters for integration with other Go logging systems.

### Standard Library Integration

### Slog Integration

Sources: [README.md312-342](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L312-L342)

Advanced Usage Patterns
-----------------------

For more advanced usage patterns, refer to the specific sub-pages:

*   [Basic Logging](https://deepwiki.com/charmbracelet/log/3.1-basic-logging)
     - Detailed examples of basic logging operations
*   [Contextual Logging](https://deepwiki.com/charmbracelet/log/3.2-contextual-logging)
     - In-depth guide to creating and using loggers with context
*   [Helper Functions](https://deepwiki.com/charmbracelet/log/3.3-helper-functions)
     - Comprehensive coverage of logging helpers
*   [Formatters](https://deepwiki.com/charmbracelet/log/4-formatters)
     - More information on output formatters
*   [Styling and Customization](https://deepwiki.com/charmbracelet/log/5-styling-and-customization)
     - Detailed guide to styling logs
*   [Configuration Options](https://deepwiki.com/charmbracelet/log/6-configuration-options)
     - Complete reference for all configuration options

By applying these usage patterns, you can create logs that are both human-readable and structured for machine processing, making debugging and system analysis much more efficient.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Usage Guide](https://deepwiki.com/charmbracelet/log/3-usage-guide#usage-guide)
    
*   [Installation](https://deepwiki.com/charmbracelet/log/3-usage-guide#installation)
    
*   [Logging Flow](https://deepwiki.com/charmbracelet/log/3-usage-guide#logging-flow)
    
*   [Basic Logging](https://deepwiki.com/charmbracelet/log/3-usage-guide#basic-logging)
    
*   [Using the Global Logger](https://deepwiki.com/charmbracelet/log/3-usage-guide#using-the-global-logger)
    
*   [Using Format Strings](https://deepwiki.com/charmbracelet/log/3-usage-guide#using-format-strings)
    
*   [Adding Key-Value Pairs](https://deepwiki.com/charmbracelet/log/3-usage-guide#adding-key-value-pairs)
    
*   [Creating Custom Loggers](https://deepwiki.com/charmbracelet/log/3-usage-guide#creating-custom-loggers)
    
*   [Basic Custom Logger](https://deepwiki.com/charmbracelet/log/3-usage-guide#basic-custom-logger)
    
*   [Configured Custom Logger](https://deepwiki.com/charmbracelet/log/3-usage-guide#configured-custom-logger)
    
*   [Modifying Logger Settings](https://deepwiki.com/charmbracelet/log/3-usage-guide#modifying-logger-settings)
    
*   [Output Formatters](https://deepwiki.com/charmbracelet/log/3-usage-guide#output-formatters)
    
*   [Contextual Logging](https://deepwiki.com/charmbracelet/log/3-usage-guide#contextual-logging)
    
*   [Using With() for Context](https://deepwiki.com/charmbracelet/log/3-usage-guide#using-with-for-context)
    
*   [Using WithPrefix](https://deepwiki.com/charmbracelet/log/3-usage-guide#using-withprefix)
    
*   [Helper Functions](https://deepwiki.com/charmbracelet/log/3-usage-guide#helper-functions)
    
*   [Style Customization](https://deepwiki.com/charmbracelet/log/3-usage-guide#style-customization)
    
*   [Integration with Other Logging Systems](https://deepwiki.com/charmbracelet/log/3-usage-guide#integration-with-other-logging-systems)
    
*   [Standard Library Integration](https://deepwiki.com/charmbracelet/log/3-usage-guide#standard-library-integration)
    
*   [Slog Integration](https://deepwiki.com/charmbracelet/log/3-usage-guide#slog-integration)
    
*   [Advanced Usage Patterns](https://deepwiki.com/charmbracelet/log/3-usage-guide#advanced-usage-patterns)
