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

Core Architecture
=================

Relevant source files

*   [level.go](https://github.com/charmbracelet/log/blob/cf6e8671/level.go)
    
*   [level\_test.go](https://github.com/charmbracelet/log/blob/cf6e8671/level_test.go)
    
*   [logger.go](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go)
    
*   [pkg.go](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go)
    

This document outlines the architectural foundation of the `charmbracelet/log` package, a structured logging system for Go applications. It explains the core components, their relationships, and how log messages flow through the system. For specific details about the Logger structure, see [Logger Structure](https://deepwiki.com/charmbracelet/log/2.1-logger-structure)
, and for information about log levels, see [Log Levels](https://deepwiki.com/charmbracelet/log/2.2-log-levels)
.

Overview
--------

The `charmbracelet/log` package is built around a central `Logger` type that provides configurable, leveled logging capabilities with support for multiple output formats and styling options. The architecture is designed for flexibility, allowing both direct use of custom loggers and access through package-level functions via a global logger.

Sources: [logger.go24-48](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L24-L48)
 [pkg.go16-22](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go#L16-L22)
 [pkg.go25-37](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go#L25-L37)

Logger Structure and Initialization
-----------------------------------

The core of the architecture is the `Logger` struct defined in `logger.go`. It maintains internal state including the output writer, log level, formatting options, and context fields.

Loggers are created through the `New` or `NewWithOptions` functions in `pkg.go`. The `NewWithOptions` function accepts an `Options` struct that allows for detailed configuration of the logger's behavior.

Sources: [logger.go24-48](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L24-L48)
 [pkg.go44-84](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go#L44-L84)

Log Flow Architecture
---------------------

When a log message is created, it flows through several steps before being output. The architecture ensures that logs are properly filtered based on level, formatted according to the chosen formatter, and then written to the configured output.

Sources: [logger.go50-140](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L50-L140)

Level System
------------

The level system provides filtering capabilities to control which messages are logged. The architecture defines five standard levels (Debug, Info, Warn, Error, Fatal) with numeric values that allow comparison.

A logger only processes messages with a level greater than or equal to its configured level. This architecture allows efficient filtering of log messages at runtime.

Sources: [level.go10-65](https://github.com/charmbracelet/log/blob/cf6e8671/level.go#L10-L65)
 [logger.go61-64](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L61-L64)

Contextual Logging Architecture
-------------------------------

The architecture supports contextual logging through the creation of derived loggers. This allows for structured logging with context that follows call paths.

Sources: [logger.go330-350](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L330-L350)

Global Logger System
--------------------

The architecture includes a global logger that provides package-level logging functions. This simplifies the most common use case while still allowing for customization.

The global logger is initialized lazily when first needed. The architecture uses atomics and once-initialization to ensure thread safety.

Sources: [pkg.go16-22](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go#L16-L22)
 [pkg.go25-37](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go#L25-L37)
 [pkg.go174-244](https://github.com/charmbracelet/log/blob/cf6e8671/pkg.go#L174-L244)

Formatter System
----------------

The architecture supports multiple output formats through a pluggable formatter system. Each formatter handles the conversion of log data to a specific output format.

The formatter selection happens during logging and is configurable through the `SetFormatter` method or during logger creation.

Sources: [logger.go125-136](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L125-L136)
 [logger.go292-297](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L292-L297)

Caller Reporting Architecture
-----------------------------

The architecture includes an optional capability to report the caller of logging functions, which is useful for debugging. This system uses Go's runtime package to retrieve call stack information.

The architecture includes a "helper" system that allows marking certain functions to be skipped when determining the caller, similar to the approach used in Go's testing package.

Sources: [logger.go66-81](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L66-L81)
 [logger.go142-156](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L142-L156)
 [logger.go157-168](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L157-L168)
 [logger.go169-176](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L169-L176)

Integration Architecture
------------------------

The logging architecture is designed to integrate with other systems, both within the Go standard library and external ecosystems.

The architecture leverages the `lipgloss` library for styling, which in turn uses `termenv` for terminal capability detection. It also provides adapters for Go's standard logging facilities.

Sources: [logger.go3-16](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L3-L16)
 [logger.go270-290](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L270-L290)
 [logger.go315-327](https://github.com/charmbracelet/log/blob/cf6e8671/logger.go#L315-L327)

Summary
-------

The `charmbracelet/log` architecture is centered around a flexible `Logger` type that supports leveled, structured logging with multiple output formats and styling options. It provides both direct logger usage and a convenient global logger with package-level functions. The system's modular design allows for customization of virtually every aspect of the logging process, from filtering and formatting to styling and caller reporting.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core Architecture](https://deepwiki.com/charmbracelet/log/2-core-architecture#core-architecture)
    
*   [Overview](https://deepwiki.com/charmbracelet/log/2-core-architecture#overview)
    
*   [Logger Structure and Initialization](https://deepwiki.com/charmbracelet/log/2-core-architecture#logger-structure-and-initialization)
    
*   [Log Flow Architecture](https://deepwiki.com/charmbracelet/log/2-core-architecture#log-flow-architecture)
    
*   [Level System](https://deepwiki.com/charmbracelet/log/2-core-architecture#level-system)
    
*   [Contextual Logging Architecture](https://deepwiki.com/charmbracelet/log/2-core-architecture#contextual-logging-architecture)
    
*   [Global Logger System](https://deepwiki.com/charmbracelet/log/2-core-architecture#global-logger-system)
    
*   [Formatter System](https://deepwiki.com/charmbracelet/log/2-core-architecture#formatter-system)
    
*   [Caller Reporting Architecture](https://deepwiki.com/charmbracelet/log/2-core-architecture#caller-reporting-architecture)
    
*   [Integration Architecture](https://deepwiki.com/charmbracelet/log/2-core-architecture#integration-architecture)
    
*   [Summary](https://deepwiki.com/charmbracelet/log/2-core-architecture#summary)
