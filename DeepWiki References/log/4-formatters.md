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

Formatters
==========

Relevant source files

*   [formatter.go](https://github.com/charmbracelet/log/blob/cf6e8671/formatter.go)
    
*   [json.go](https://github.com/charmbracelet/log/blob/cf6e8671/json.go)
    
*   [json\_test.go](https://github.com/charmbracelet/log/blob/cf6e8671/json_test.go)
    
*   [logfmt.go](https://github.com/charmbracelet/log/blob/cf6e8671/logfmt.go)
    
*   [text.go](https://github.com/charmbracelet/log/blob/cf6e8671/text.go)
    

This document describes the formatter system in the charmbracelet/log library, which controls how log messages are structured and presented in output. For information about styling the text output specifically, see [Styling and Customization](https://deepwiki.com/charmbracelet/log/5-styling-and-customization)
.

Overview
--------

The formatter system in charmbracelet/log provides three output format options, each suited for different use cases:

Sources: [formatter.go3-14](https://github.com/charmbracelet/log/blob/cf6e8671/formatter.go#L3-L14)

Formatter Types
---------------

The library provides three formatters, each represented as a constant of the `Formatter` type:

| Formatter | Description | Best For |
| --- | --- | --- |
| `TextFormatter` | Human-readable, styled text output | Console output, developer-facing logs |
| `JSONFormatter` | Structured JSON objects | Machine processing, log aggregation systems |
| `LogfmtFormatter` | Key-value pairs in logfmt format | Systems that parse logfmt, compromise between readability and structure |

Sources: [formatter.go3-14](https://github.com/charmbracelet/log/blob/cf6e8671/formatter.go#L3-L14)

Standard Log Field Keys
-----------------------

All formatters use a set of standard keys for common log fields, which can be customized globally:

Sources: [formatter.go16-27](https://github.com/charmbracelet/log/blob/cf6e8671/formatter.go#L16-L27)

Formatter Implementation
------------------------

Each formatter is implemented as a method on the `Logger` struct that processes key-value pairs and writes formatted output to the logger's buffer:

Sources: [formatter.go3-14](https://github.com/charmbracelet/log/blob/cf6e8671/formatter.go#L3-L14)
 [text.go167-272](https://github.com/charmbracelet/log/blob/cf6e8671/text.go#L167-L272)
 [json.go10-31](https://github.com/charmbracelet/log/blob/cf6e8671/json.go#L10-L31)
 [logfmt.go11-32](https://github.com/charmbracelet/log/blob/cf6e8671/logfmt.go#L11-L32)

Text Formatter
--------------

The text formatter produces human-readable output with optional styling. It's the default formatter in charmbracelet/log.

### Features

*   Color-coded log levels (when terminal supports it)
*   Configurable styles for different log elements
*   Special handling for multiline values
*   Proper escaping of non-printable characters

### Output Format Example

    12:34:56 INFO app: This is a log message key=value another=value
    

### Implementation Details

The text formatter processes each key-value pair differently based on the key:

*   Standard keys (timestamp, level, caller, prefix, message) get special formatting
*   Other keys are formatted as "key=value" pairs
*   Values containing newlines are specially formatted with indentation
*   Values are escaped if they contain non-printable characters

Sources: [text.go167-272](https://github.com/charmbracelet/log/blob/cf6e8671/text.go#L167-L272)

JSON Formatter
--------------

The JSON formatter produces structured JSON objects, one per log line, making it ideal for machine processing.

### Features

*   Preserves data types (numbers, booleans, arrays, objects)
*   Each log entry is a complete, valid JSON object
*   Special handling for errors and custom types
*   No color styling (purely data-focused)

### Output Format Example

    {"time":"12:34:56","level":"info","prefix":"app","msg":"This is a log message","key":"value","another":"value"}
    

### Implementation Details

The JSON formatter uses a custom `jsonWriter` to build the JSON structure:

*   Standard log fields are included with their standard keys
*   Complex Go types are properly encoded to JSON
*   Error objects are specially handled to include their messages
*   Failed encoding is gracefully handled with fallback values

Sources: [json.go10-151](https://github.com/charmbracelet/log/blob/cf6e8671/json.go#L10-L151)

Logfmt Formatter
----------------

The logfmt formatter produces output in the "logfmt" format - a line-oriented, key-value pair format that balances human and machine readability.

### Features

*   Simple key=value format
*   Good for grep and basic log analysis
*   Handles complex values by converting them to strings
*   No color styling

### Output Format Example

    time=12:34:56 level=info prefix=app msg="This is a log message" key=value another=value
    

### Implementation Details

The logfmt formatter uses the `go-logfmt/logfmt` package:

*   Formats timestamps according to the logger's time format
*   Converts keys to strings
*   Handles unsupported value types by converting them to strings with `fmt.Sprintf`

Sources: [logfmt.go11-32](https://github.com/charmbracelet/log/blob/cf6e8671/logfmt.go#L11-L32)

Output Comparison
-----------------

Sources: [text.go167-272](https://github.com/charmbracelet/log/blob/cf6e8671/text.go#L167-L272)
 [json.go10-151](https://github.com/charmbracelet/log/blob/cf6e8671/json.go#L10-L151)
 [logfmt.go11-32](https://github.com/charmbracelet/log/blob/cf6e8671/logfmt.go#L11-L32)

Setting a Formatter
-------------------

To change the output format of a logger, use the `SetFormatter` method:

The default formatter is `TextFormatter`.

Custom Keys
-----------

You can customize the standard keys used by formatters by modifying the global key variables:

Note that these changes affect all formatters and all loggers created after the change.

Sources: [formatter.go16-27](https://github.com/charmbracelet/log/blob/cf6e8671/formatter.go#L16-L27)
 [json\_test.go187-200](https://github.com/charmbracelet/log/blob/cf6e8671/json_test.go#L187-L200)

Formatter Selection Flow
------------------------

Sources: [text.go167-272](https://github.com/charmbracelet/log/blob/cf6e8671/text.go#L167-L272)
 [json.go10-31](https://github.com/charmbracelet/log/blob/cf6e8671/json.go#L10-L31)
 [logfmt.go11-32](https://github.com/charmbracelet/log/blob/cf6e8671/logfmt.go#L11-L32)

The formatter system is a crucial part of charmbracelet/log's flexibility, allowing the same logging code to produce output suitable for different environments and use cases - from developer-friendly console output to structured data for log aggregation systems.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Formatters](https://deepwiki.com/charmbracelet/log/4-formatters#formatters)
    
*   [Overview](https://deepwiki.com/charmbracelet/log/4-formatters#overview)
    
*   [Formatter Types](https://deepwiki.com/charmbracelet/log/4-formatters#formatter-types)
    
*   [Standard Log Field Keys](https://deepwiki.com/charmbracelet/log/4-formatters#standard-log-field-keys)
    
*   [Formatter Implementation](https://deepwiki.com/charmbracelet/log/4-formatters#formatter-implementation)
    
*   [Text Formatter](https://deepwiki.com/charmbracelet/log/4-formatters#text-formatter)
    
*   [Features](https://deepwiki.com/charmbracelet/log/4-formatters#features)
    
*   [Output Format Example](https://deepwiki.com/charmbracelet/log/4-formatters#output-format-example)
    
*   [Implementation Details](https://deepwiki.com/charmbracelet/log/4-formatters#implementation-details)
    
*   [JSON Formatter](https://deepwiki.com/charmbracelet/log/4-formatters#json-formatter)
    
*   [Features](https://deepwiki.com/charmbracelet/log/4-formatters#features-1)
    
*   [Output Format Example](https://deepwiki.com/charmbracelet/log/4-formatters#output-format-example-1)
    
*   [Implementation Details](https://deepwiki.com/charmbracelet/log/4-formatters#implementation-details-1)
    
*   [Logfmt Formatter](https://deepwiki.com/charmbracelet/log/4-formatters#logfmt-formatter)
    
*   [Features](https://deepwiki.com/charmbracelet/log/4-formatters#features-2)
    
*   [Output Format Example](https://deepwiki.com/charmbracelet/log/4-formatters#output-format-example-2)
    
*   [Implementation Details](https://deepwiki.com/charmbracelet/log/4-formatters#implementation-details-2)
    
*   [Output Comparison](https://deepwiki.com/charmbracelet/log/4-formatters#output-comparison)
    
*   [Setting a Formatter](https://deepwiki.com/charmbracelet/log/4-formatters#setting-a-formatter)
    
*   [Custom Keys](https://deepwiki.com/charmbracelet/log/4-formatters#custom-keys)
    
*   [Formatter Selection Flow](https://deepwiki.com/charmbracelet/log/4-formatters#formatter-selection-flow)
