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

Configuration Options
=====================

Relevant source files

*   [options.go](https://github.com/charmbracelet/log/blob/cf6e8671/options.go)
    
*   [options\_test.go](https://github.com/charmbracelet/log/blob/cf6e8671/options_test.go)
    

This page documents all the configuration options available for the Charmbracelet log library. The configuration system provides flexibility to customize logger behavior, appearance, and output format. For information about styling and visual customization specifically, see [Styling and Customization](https://deepwiki.com/charmbracelet/log/5-styling-and-customization)
.

Overview of Configuration Options
---------------------------------

The `charmbracelet/log` library provides a comprehensive set of configuration options through the `Options` struct. These options control various aspects of logging behavior, from log levels to timestamp formatting and caller information.

Sources: [options.go39-61](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L39-L61)

Creating a Logger with Custom Options
-------------------------------------

You can configure a logger either at creation time using `NewWithOptions()` or by using setter methods on an existing logger.

Sources: [options.go39-61](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L39-L61)
 [options\_test.go12-26](https://github.com/charmbracelet/log/blob/cf6e8671/options_test.go#L12-L26)

Available Configuration Options
-------------------------------

### Time Handling Options

Time-related options control how timestamps are formatted and processed in logs.

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `TimeFunction` | `func(time.Time) time.Time` | `time.Now` | Function that processes timestamps |
| `TimeFormat` | `string` | `"2006/01/02 15:04:05"` | Format string for timestamps |
| `ReportTimestamp` | `bool` | `false` | Whether to include timestamps in logs |

The library provides a built-in `NowUTC` function that can be used to convert timestamps to UTC time zone:

Sources: [options.go8-23](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L8-L23)
 [options.go42-50](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L42-L50)

### Log Content Options

These options control what information is included in log messages.

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `Level` | `Level` | `InfoLevel` | Minimum level at which logs are emitted |
| `Prefix` | `string` | `""` (empty string) | Prefix prepended to all log messages |
| `Fields` | `[]interface{}` | `nil` | Key-value pairs included with every log entry |

Sources: [options.go45-48](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L45-L48)
 [options.go57-58](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L57-L58)

### Caller Information Options

These options control whether and how information about the calling function is included in logs.

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `ReportCaller` | `bool` | `false` | Whether to include caller information |
| `CallerFormatter` | `func(string, int, string) string` | `ShortCallerFormatter` | Function to format caller information |
| `CallerOffset` | `int` | `0` | Stack frames to skip when determining caller |

The library provides two built-in caller formatters:

1.  `ShortCallerFormatter`: Shows the last 2 levels of the path and line number (e.g., `log/options_test.go:42`)
2.  `LongCallerFormatter`: Shows the full path and line number (e.g., `/home/user/project/log/options_test.go:42`)

Sources: [options.go25-37](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L25-L37)
 [options.go52-56](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L52-L56)
 [options\_test.go28-73](https://github.com/charmbracelet/log/blob/cf6e8671/options_test.go#L28-L73)

### Output Format Options

This option controls the format of the log output.

| Option | Type | Default | Description |
| --- | --- | --- | --- |
| `Formatter` | `Formatter` | `TextFormatter` | The formatter used to structure log output |

Available formatters include:

*   `TextFormatter`: Human-readable text with optional styling
*   `JSONFormatter`: Structured JSON format
*   `LogfmtFormatter`: Key-value pairs in the logfmt format

For more details on formatters, see [Formatters](https://deepwiki.com/charmbracelet/log/4-formatters)
.

Sources: [options.go59-61](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L59-L61)

Example: Configuring a Logger
-----------------------------

Here's a comprehensive example of creating a logger with custom options:

Sources: [options.go39-61](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L39-L61)
 [options\_test.go12-26](https://github.com/charmbracelet/log/blob/cf6e8671/options_test.go#L12-L26)

Configuration Flow and Effect
-----------------------------

The following diagram illustrates how configuration options affect the logging process:

Sources: [options.go39-61](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L39-L61)

Available Setter Methods
------------------------

The logger provides setter methods to modify configuration after creation:

| Method | Description |
| --- | --- |
| `SetLevel(Level)` | Sets the minimum log level |
| `SetTimeFormat(string)` | Sets the time format string |
| `SetTimeFunction(TimeFunction)` | Sets the time processing function |
| `SetPrefix(string)` | Sets the prefix for log messages |
| `SetReportTimestamp(bool)` | Enables/disables timestamp reporting |
| `SetReportCaller(bool)` | Enables/disables caller information |
| `SetCallerFormatter(CallerFormatter)` | Sets the formatter for caller information |
| `SetFormatter(Formatter)` | Sets the log formatter |

Example usage:

Sources: [options.go39-61](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L39-L61)

Relationship with Other Components
----------------------------------

The configuration options have direct effects on the behavior of formatters and the output of log messages:

Sources: [options.go39-61](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L39-L61)

Default Configuration
---------------------

When creating a logger without specifying options, the following defaults are used:

| Option | Default Value |
| --- | --- |
| `Level` | `InfoLevel` |
| `TimeFormat` | `"2006/01/02 15:04:05"` |
| `TimeFunction` | `time.Now` |
| `Prefix` | `""` (empty string) |
| `ReportTimestamp` | `false` |
| `ReportCaller` | `false` |
| `CallerFormatter` | `ShortCallerFormatter` |
| `CallerOffset` | `0` |
| `Fields` | `nil` |
| `Formatter` | `TextFormatter` |

Sources: [options.go39-61](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L39-L61)
 [options\_test.go12-26](https://github.com/charmbracelet/log/blob/cf6e8671/options_test.go#L12-L26)

Custom Caller Formatters
------------------------

You can create custom caller formatters to format the caller information in any way you need:

Sources: [options.go25-37](https://github.com/charmbracelet/log/blob/cf6e8671/options.go#L25-L37)
 [options\_test.go28-73](https://github.com/charmbracelet/log/blob/cf6e8671/options_test.go#L28-L73)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Configuration Options](https://deepwiki.com/charmbracelet/log/6-configuration-options#configuration-options)
    
*   [Overview of Configuration Options](https://deepwiki.com/charmbracelet/log/6-configuration-options#overview-of-configuration-options)
    
*   [Creating a Logger with Custom Options](https://deepwiki.com/charmbracelet/log/6-configuration-options#creating-a-logger-with-custom-options)
    
*   [Available Configuration Options](https://deepwiki.com/charmbracelet/log/6-configuration-options#available-configuration-options)
    
*   [Time Handling Options](https://deepwiki.com/charmbracelet/log/6-configuration-options#time-handling-options)
    
*   [Log Content Options](https://deepwiki.com/charmbracelet/log/6-configuration-options#log-content-options)
    
*   [Caller Information Options](https://deepwiki.com/charmbracelet/log/6-configuration-options#caller-information-options)
    
*   [Output Format Options](https://deepwiki.com/charmbracelet/log/6-configuration-options#output-format-options)
    
*   [Example: Configuring a Logger](https://deepwiki.com/charmbracelet/log/6-configuration-options#example-configuring-a-logger)
    
*   [Configuration Flow and Effect](https://deepwiki.com/charmbracelet/log/6-configuration-options#configuration-flow-and-effect)
    
*   [Available Setter Methods](https://deepwiki.com/charmbracelet/log/6-configuration-options#available-setter-methods)
    
*   [Relationship with Other Components](https://deepwiki.com/charmbracelet/log/6-configuration-options#relationship-with-other-components)
    
*   [Default Configuration](https://deepwiki.com/charmbracelet/log/6-configuration-options#default-configuration)
    
*   [Custom Caller Formatters](https://deepwiki.com/charmbracelet/log/6-configuration-options#custom-caller-formatters)
