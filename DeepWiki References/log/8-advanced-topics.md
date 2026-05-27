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

Advanced Topics
===============

Relevant source files

*   [logger.go](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go)
    
*   [pkg.go](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go)
    
*   [text.go](https://github.com/charmbracelet/log/blob/cf6e8671/text.go)
    

This document covers advanced usage scenarios and implementation details of the Charmbracelet Log library. It explores features that go beyond basic logging, focusing on customization, optimization, and deeper understanding of the internal mechanisms. For basic usage information, see [Usage Guide](https://deepwiki.com/charmbracelet/log/3-usage-guide)
, and for information on formatters, see [Formatters](https://deepwiki.com/charmbracelet/log/4-formatters)
.

Helper Functions for Stack Traces
---------------------------------

When debugging applications with extensive logging, it's crucial to see the correct source location where a log entry originated. The Helper functionality allows you to mark functions as helpers so they're skipped when reporting the caller location.

Example usage:

When `LogWrapper` is called, the caller of `LogWrapper` will be reported as the source, not `LogWrapper` itself.

Sources: [logger.go142-156](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L142-L156)
 [logger.go157-168](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L157-L168)
 [pkg.go167-172](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go#L167-L172)

Custom Caller Formatting
------------------------

The logger allows custom formatting of caller information through the `CallerFormatter` function type. This gives you control over how file paths, line numbers, and function names are displayed.

The library provides two built-in formatters:

*   `ShortCallerFormatter`: Shows only the final file name and line number
*   `LongCallerFormatter`: Shows the full file path and line number

You can create your own formatter:

Sources: [logger.go299-304](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L299-L304)
 [logger.go170-176](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L170-L176)
 [logger.go95-101](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L95-L101)

Custom Time Functions and Formats
---------------------------------

The logger provides flexibility in how timestamps are handled through custom time functions and format strings.

TimeFunction allows you to transform the time before it's formatted, which is useful for:

*   Time zone conversions
*   Timestamp rounding
*   Testing with fixed times

Example:

Sources: [logger.go256-268](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L256-L268)
 [pkg.go106-114](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go#L106-L114)

Thread Safety and Concurrency
-----------------------------

The Charmbracelet Log library is designed to be thread-safe, allowing multiple goroutines to log concurrently without corruption or data races.

Thread safety is ensured through:

1.  A read-write mutex protecting all logging operations
2.  Atomic variables for level checking and discard state
3.  Thread-local buffer usage to avoid contention

Sources: [logger.go25-29](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L25-L29)
 [logger.go125-126](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L125-L126)
 [logger.go236-240](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L236-L240)
 [pkg.go21-22](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go#L21-L22)

Performance Optimization
------------------------

Several techniques are used in the library to optimize performance:

### Early Filtering

The library provides fast paths to avoid unnecessary work:

*   Atomic check of discard state
*   Level check using atomic operations
*   Helper mapping using sync.Map for concurrent lookups

### Buffer Management

String building and formatting operations use efficient buffer pools and minimal allocations:

Sources: [logger.go57-64](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L57-L64)
 [text.go64-68](https://github.com/charmbracelet/log/blob/cf6e8671/text.go#L64-L68)
 [text.go76-79](https://github.com/charmbracelet/log/blob/cf6e8671/text.go#L76-L79)
 [logger.go138-139](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L138-L139)

Advanced Style Customization
----------------------------

The text formatter supports extensive style customization through the `Styles` struct, allowing fine-grained control over appearance.

You can customize styles for specific keys using the `Keys` and `Values` maps:

This allows for highlighting important fields or creating visually distinctive logs.

Sources: [logger.go319-327](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L319-L327)
 [text.go167-272](https://github.com/charmbracelet/log/blob/cf6e8671/text.go#L167-L272)

Logger Context Propagation
--------------------------

The Charmbracelet Log library supports creating derived loggers with additional context through the `With` and `WithPrefix` methods.

Context propagation happens by:

1.  Creating a new logger instance with a fresh mutex
2.  Copying and extending the existing fields slice
3.  Preserving all other settings from the parent logger

This creates a hierarchical logging structure where child loggers inherit all settings from their parent but include additional context.

Sources: [logger.go330-350](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L330-L350)

Custom Formatters Implementation
--------------------------------

To implement a custom formatter beyond the built-in ones, you need to understand the internal formatter selection flow:

While the library doesn't directly expose a formatter interface for extension, you can create a custom formatter by:

1.  Creating a new formatter enum value
2.  Extending the formatter switch case in `handle()`
3.  Implementing the formatting function

This requires forking or extending the library code.

Sources: [logger.go127-136](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L127-L136)
 [text.go167-272](https://github.com/charmbracelet/log/blob/cf6e8671/text.go#L167-L272)

Memory Management with Structured Logging
-----------------------------------------

When using structured logging extensively, understanding memory allocation patterns becomes important:

| Operation | Memory Behavior | Notes |
| --- | --- | --- |
| Basic logging | Minimal allocations | With level filtering, can be allocation-free |
| With() | Creates new logger | Copies and extends fields slice |
| Keys/Values | Stored as `[]interface{}` | Consider impact with high-volume logging |
| Text formatting | Uses buffer pooling | Minimizes allocations for string construction |

For high-performance applications, consider:

*   Using pre-created loggers with With() during initialization, not in hot paths
*   Checking IsLevelEnabled() before constructing expensive log messages
*   Using io.Discard for production debug logs

Sources: [logger.go54-82](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L54-L82)
 [logger.go330-343](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L330-L343)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Advanced Topics](https://deepwiki.com/charmbracelet/log/8-advanced-topics#advanced-topics)
    
*   [Helper Functions for Stack Traces](https://deepwiki.com/charmbracelet/log/8-advanced-topics#helper-functions-for-stack-traces)
    
*   [Custom Caller Formatting](https://deepwiki.com/charmbracelet/log/8-advanced-topics#custom-caller-formatting)
    
*   [Custom Time Functions and Formats](https://deepwiki.com/charmbracelet/log/8-advanced-topics#custom-time-functions-and-formats)
    
*   [Thread Safety and Concurrency](https://deepwiki.com/charmbracelet/log/8-advanced-topics#thread-safety-and-concurrency)
    
*   [Performance Optimization](https://deepwiki.com/charmbracelet/log/8-advanced-topics#performance-optimization)
    
*   [Early Filtering](https://deepwiki.com/charmbracelet/log/8-advanced-topics#early-filtering)
    
*   [Buffer Management](https://deepwiki.com/charmbracelet/log/8-advanced-topics#buffer-management)
    
*   [Advanced Style Customization](https://deepwiki.com/charmbracelet/log/8-advanced-topics#advanced-style-customization)
    
*   [Logger Context Propagation](https://deepwiki.com/charmbracelet/log/8-advanced-topics#logger-context-propagation)
    
*   [Custom Formatters Implementation](https://deepwiki.com/charmbracelet/log/8-advanced-topics#custom-formatters-implementation)
    
*   [Memory Management with Structured Logging](https://deepwiki.com/charmbracelet/log/8-advanced-topics#memory-management-with-structured-logging)
