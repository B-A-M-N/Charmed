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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1)
    
*   [examples/styles/styles.go](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go)
    
*   [examples/styles/styles.tape](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.tape)
    
*   [go.mod](https://github.com/charmbracelet/log/blob/cf6e8671/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/log/blob/cf6e8671/go.sum)
    

This document provides an introduction to the charmbracelet/log library, a structured, leveled logging library for Go that offers human-readable, colorful output with various formatting options. The library is designed to provide a simple yet powerful API for logging in Go applications, with features including customizable styling, multiple output formats, and integration with standard Go logging libraries.

For detailed information about specific components, please refer to the corresponding wiki pages:

*   For details on the Logger structure, see [Core Architecture](https://deepwiki.com/charmbracelet/log/2-core-architecture)
    
*   For usage examples, see [Usage Guide](https://deepwiki.com/charmbracelet/log/3-usage-guide)
    
*   For information about formatters, see [Formatters](https://deepwiki.com/charmbracelet/log/4-formatters)
    

Purpose and Key Features
------------------------

The charmbracelet/log library aims to enhance Go application logging by providing:

1.  **Human-readable, colorful output** - Uses [Lip Gloss](https://github.com/charmbracelet/log/blob/cf6e8671/Lip%20Gloss)
     for styling and colorization
2.  **Leveled logging** - Debug, Info, Warn, Error, and Fatal levels
3.  **Structured logging** - Key-value pairs for contextual information
4.  **Multiple output formats** - Text, JSON, and Logfmt
5.  **Context integration** - Store and retrieve loggers from Go's context
6.  **Integration with standard logging systems** - `log` and `slog` compatibility

Sources: [README.md16-38](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L16-L38)

Core Components
---------------

The charmbracelet/log library is composed of several core components that work together to provide a flexible logging system. The diagram above illustrates these components and their relationships.

Sources: [README.md26-38](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L26-L38)

Logging Flow
------------

This diagram depicts the flow of a log message through the system. When a logging method is called, the logger first checks if the message level meets or exceeds the configured level. If the check passes, the message is formatted according to the selected formatter, and styles are applied for text output. Finally, the formatted message is written to the output.

Sources: [README.md110-134](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L110-L134)

Logger Hierarchy
----------------

Loggers can form a hierarchy through the `With` method, which creates child loggers that inherit settings from their parent while including additional context fields. The `WithPrefix` method creates loggers with specific prefixes.

Sources: [README.md235-245](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L235-L245)

Available Log Levels
--------------------

The library provides several log levels to categorize log messages by importance:

| Level | Method | Description |
| --- | --- | --- |
| `DebugLevel` | `Debug/Debugf` | Detailed information for debugging purposes |
| `InfoLevel` | `Info/Infof` | General operational information |
| `WarnLevel` | `Warn/Warnf` | Warning events that might cause issues |
| `ErrorLevel` | `Error/Errorf` | Error events that might still allow the application to continue |
| `FatalLevel` | `Fatal/Fatalf` | Critical errors that cause the application to exit |

The default level for the global logger is `InfoLevel`.

Sources: [README.md110-134](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L110-L134)

Formatters
----------

The library supports multiple output formats through its formatter system:

1.  **TextFormatter** (default) - Human-readable, colorful text output
2.  **JSONFormatter** - Structured JSON output for machine parsing
3.  **LogfmtFormatter** - Key-value pairs in the logfmt format

Each formatter has its own approach to rendering log entries, with the TextFormatter being the most visually appealing for human readers due to its styling capabilities.

Sources: [README.md196-202](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L196-L202)

Configuration Options
---------------------

The library provides a comprehensive set of configuration options through the `Options` struct:

| Option | Description |
| --- | --- |
| `Level` | Sets the minimum level at which logs will be output |
| `ReportTimestamp` | Whether to include timestamps in log output |
| `ReportCaller` | Whether to include caller information (file and line) |
| `Prefix` | A string prefix for all log entries |
| `TimeFunction` | Function to generate timestamps |
| `TimeFormat` | Format string for timestamps |
| `Formatter` | The formatter to use (Text, JSON, Logfmt) |
| `Fields` | Default key-value pairs to include in all log entries |

Here's an example configuration:

Sources: [README.md162-203](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L162-L203)

Styling System
--------------

The library uses [Lip Gloss](https://github.com/charmbracelet/log/blob/cf6e8671/Lip%20Gloss)
 to style log output when using the TextFormatter. The styling system allows customization of various elements:

Custom styles can be applied to various components of the log output, including level indicators, timestamps, keys, and values.

Sources: [README.md210-231](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L210-L231)
 [examples/styles/styles.go12-23](https://github.com/charmbracelet/log/blob/cf6e8671/examples/styles/styles.go#L12-L23)

Integration with Other Logging Systems
--------------------------------------

The library provides integration with Go's standard logging systems:

1.  **Standard Log Adapter** - Allows the logger to be used with packages expecting the standard `log.Logger` interface
2.  **Slog Handler** - Integrates with Go's structured logging (slog) package

These integrations allow the charmbracelet/log library to be used in places where the standard logging interfaces are expected.

Sources: [README.md320-342](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L320-L342)
 [README.md312-319](https://github.com/charmbracelet/log/blob/cf6e8671/README.md?plain=1#L312-L319)

Dependencies
------------

The library relies on several key dependencies:

| Dependency | Purpose |
| --- | --- |
| [github.com/charmbracelet/lipgloss](https://github.com/charmbracelet/log/blob/cf6e8671/github.com/charmbracelet/lipgloss) | Styling and colorization |
| [github.com/go-logfmt/logfmt](https://github.com/charmbracelet/log/blob/cf6e8671/github.com/go-logfmt/logfmt) | Logfmt formatter implementation |
| [github.com/muesli/termenv](https://github.com/charmbracelet/log/blob/cf6e8671/github.com/muesli/termenv) | Terminal environment detection |
| [golang.org/x/exp](https://golang.org/x/exp) | Experimental Go packages (for slog support) |

Sources: [go.mod5-11](https://github.com/charmbracelet/log/blob/cf6e8671/go.mod#L5-L11)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/log/1-overview#overview)
    
*   [Purpose and Key Features](https://deepwiki.com/charmbracelet/log/1-overview#purpose-and-key-features)
    
*   [Core Components](https://deepwiki.com/charmbracelet/log/1-overview#core-components)
    
*   [Logging Flow](https://deepwiki.com/charmbracelet/log/1-overview#logging-flow)
    
*   [Logger Hierarchy](https://deepwiki.com/charmbracelet/log/1-overview#logger-hierarchy)
    
*   [Available Log Levels](https://deepwiki.com/charmbracelet/log/1-overview#available-log-levels)
    
*   [Formatters](https://deepwiki.com/charmbracelet/log/1-overview#formatters)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/log/1-overview#configuration-options)
    
*   [Styling System](https://deepwiki.com/charmbracelet/log/1-overview#styling-system)
    
*   [Integration with Other Logging Systems](https://deepwiki.com/charmbracelet/log/1-overview#integration-with-other-logging-systems)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/log/1-overview#dependencies)
