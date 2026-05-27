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

Integration
===========

Relevant source files

*   [logger\_121.go](https://github.com/charmbracelet/log/blob/cf6e8671/logger_121.go)
    
*   [logger\_no121.go](https://github.com/charmbracelet/log/blob/cf6e8671/logger_no121.go)
    
*   [stdlog.go](https://github.com/charmbracelet/log/blob/cf6e8671/stdlog.go)
    
*   [stdlog\_test.go](https://github.com/charmbracelet/log/blob/cf6e8671/stdlog_test.go)
    

This page documents how to integrate Charmbracelet's log package with other logging systems in the Go ecosystem. The log package provides seamless integration with Go's standard library logging package (`log`) and the structured logging package (`slog`), enabling you to leverage Charmbracelet log's rich formatting and styling capabilities while maintaining compatibility with existing code.

For information about the core logging functionality, see [Core Architecture](https://deepwiki.com/charmbracelet/log/2-core-architecture)
. For usage guidelines and examples, see [Usage Guide](https://deepwiki.com/charmbracelet/log/3-usage-guide)
.

Integration Overview
--------------------

Charmbracelet's log package offers integration mechanisms for two primary logging systems:

1.  Go's standard `log` package through the Standard Log Adapter
2.  Go's structured logging interfaces through the Slog Handler

Sources: [stdlog.go1-68](https://github.com/charmbracelet/log/blob/cf6e8671/stdlog.go#L1-L68)
 [logger\_121.go1-71](https://github.com/charmbracelet/log/blob/cf6e8671/logger_121.go#L1-L71)
 [logger\_no121.go1-69](https://github.com/charmbracelet/log/blob/cf6e8671/logger_no121.go#L1-L69)

Standard Library Integration
----------------------------

The Charmbracelet log package provides a mechanism to create a standard `log.Logger` that uses the underlying Charmbracelet logger for output, allowing you to use familiar standard library logging while benefiting from Charmbracelet's enhanced formatting.

### How It Works

When integrating with the standard log package, the `stdLogWriter` adapter:

1.  Receives log messages from the standard `log.Logger`
2.  Optionally parses log levels from prefixes (DEBUG, INFO, WARN, ERROR)
3.  Routes the message to the appropriate level method on the Charmbracelet logger
4.  Adjusts the caller offset to maintain accurate caller information

Sources: [stdlog.go8-45](https://github.com/charmbracelet/log/blob/cf6e8671/stdlog.go#L8-L45)

### Using Standard Log Integration

To create a standard library logger that uses Charmbracelet log:

#### Level Detection

By default, the adapter attempts to infer log levels from message prefixes:

| Prefix | Mapped Level |
| --- | --- |
| "DEBUG" | DebugLevel |
| "INFO" | InfoLevel |
| "WARN" | WarnLevel |
| "ERROR" | ErrorLevel |
| "ERR" | ErrorLevel |
| _none_ | InfoLevel |

#### Forcing Log Levels

You can force all messages to a specific level using `StandardLogOptions`:

Sources: [stdlog.go47-67](https://github.com/charmbracelet/log/blob/cf6e8671/stdlog.go#L47-L67)
 [stdlog\_test.go15-119](https://github.com/charmbracelet/log/blob/cf6e8671/stdlog_test.go#L15-L119)

Structured Logging (slog) Integration
-------------------------------------

The Charmbracelet logger implements the `slog.Handler` interface, allowing it to be used with Go's structured logging package, either the standard library version (`log/slog` in Go 1.21+) or the experimental version (`golang.org/x/exp/slog` for older Go versions).

### Implementation Details

The Charmbracelet logger implements the following `slog.Handler` methods:

1.  `Enabled` - Checks if the logger should log at the given level
2.  `Handle` - Processes log records, converting attributes to fields
3.  `WithAttrs` - Creates a new handler with additional attributes
4.  `WithGroup` - Creates a new handler with the given group name as prefix

Sources: [logger\_121.go22-71](https://github.com/charmbracelet/log/blob/cf6e8671/logger_121.go#L22-L71)
 [logger\_no121.go23-68](https://github.com/charmbracelet/log/blob/cf6e8671/logger_no121.go#L23-L68)

### Using with slog

#### Go 1.21+

#### Go <1.21

### Level Mapping

The slog levels are mapped to Charmbracelet log levels as follows:

| slog Level | Charmbracelet Level |
| --- | --- |
| slog.LevelDebug | log.DebugLevel |
| slog.LevelInfo | log.InfoLevel |
| slog.LevelWarn | log.WarnLevel |
| slog.LevelError | log.ErrorLevel |

Integration Architecture
------------------------

Sources: [stdlog.go1-68](https://github.com/charmbracelet/log/blob/cf6e8671/stdlog.go#L1-L68)
 [logger\_121.go1-71](https://github.com/charmbracelet/log/blob/cf6e8671/logger_121.go#L1-L71)
 [logger\_no121.go1-69](https://github.com/charmbracelet/log/blob/cf6e8671/logger_no121.go#L1-L69)

Best Practices
--------------

### When to Use Standard Library Integration

*   When migrating existing code that uses the standard log package
*   When working with libraries that require a standard logger
*   When simplicity is preferred over structured logging capabilities

### When to Use slog Integration

*   When structured logging with key-value pairs is preferred
*   When working with modern Go code (especially Go 1.21+)
*   When you need to integrate with the slog ecosystem

### Choosing Between Direct and Indirect Integration

Sources: [stdlog.go1-68](https://github.com/charmbracelet/log/blob/cf6e8671/stdlog.go#L1-L68)
 [logger\_121.go1-71](https://github.com/charmbracelet/log/blob/cf6e8671/logger_121.go#L1-L71)
 [logger\_no121.go1-69](https://github.com/charmbracelet/log/blob/cf6e8671/logger_no121.go#L1-L69)

Further Reading
---------------

*   For details about the log levels available and how filtering works, see [Log Levels](https://deepwiki.com/charmbracelet/log/2.2-log-levels)
    
*   For formatter options available for structured logging, see [Formatters](https://deepwiki.com/charmbracelet/log/4-formatters)
    
*   For more detailed examples of standard logging integration, see [Standard Library Integration](https://deepwiki.com/charmbracelet/log/7.1-standard-library-integration)
    
*   For more detailed examples of slog integration, see [Slog Integration](https://deepwiki.com/charmbracelet/log/7.2-slog-integration)
    

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Integration](https://deepwiki.com/charmbracelet/log/7-integration#integration)
    
*   [Integration Overview](https://deepwiki.com/charmbracelet/log/7-integration#integration-overview)
    
*   [Standard Library Integration](https://deepwiki.com/charmbracelet/log/7-integration#standard-library-integration)
    
*   [How It Works](https://deepwiki.com/charmbracelet/log/7-integration#how-it-works)
    
*   [Using Standard Log Integration](https://deepwiki.com/charmbracelet/log/7-integration#using-standard-log-integration)
    
*   [Level Detection](https://deepwiki.com/charmbracelet/log/7-integration#level-detection)
    
*   [Forcing Log Levels](https://deepwiki.com/charmbracelet/log/7-integration#forcing-log-levels)
    
*   [Structured Logging (slog) Integration](https://deepwiki.com/charmbracelet/log/7-integration#structured-logging-slog-integration)
    
*   [Implementation Details](https://deepwiki.com/charmbracelet/log/7-integration#implementation-details)
    
*   [Using with slog](https://deepwiki.com/charmbracelet/log/7-integration#using-with-slog)
    
*   [Go 1.21+](https://deepwiki.com/charmbracelet/log/7-integration#go-121)
    
*   [Go <1.21](https://deepwiki.com/charmbracelet/log/7-integration#go-121-1)
    
*   [Level Mapping](https://deepwiki.com/charmbracelet/log/7-integration#level-mapping)
    
*   [Integration Architecture](https://deepwiki.com/charmbracelet/log/7-integration#integration-architecture)
    
*   [Best Practices](https://deepwiki.com/charmbracelet/log/7-integration#best-practices)
    
*   [When to Use Standard Library Integration](https://deepwiki.com/charmbracelet/log/7-integration#when-to-use-standard-library-integration)
    
*   [When to Use slog Integration](https://deepwiki.com/charmbracelet/log/7-integration#when-to-use-slog-integration)
    
*   [Choosing Between Direct and Indirect Integration](https://deepwiki.com/charmbracelet/log/7-integration#choosing-between-direct-and-indirect-integration)
    
*   [Further Reading](https://deepwiki.com/charmbracelet/log/7-integration#further-reading)
