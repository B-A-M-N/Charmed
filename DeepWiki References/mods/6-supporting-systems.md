Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/mods](https://github.com/charmbracelet/mods "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 7 November 2025 ([07a05d](https://github.com/charmbracelet/mods/commits/07a05d5b)
)

*   [Introduction to Mods](https://deepwiki.com/charmbracelet/mods/1-introduction-to-mods)
    
*   [Installation and Setup](https://deepwiki.com/charmbracelet/mods/1.1-installation-and-setup)
    
*   [Quick Start Guide](https://deepwiki.com/charmbracelet/mods/1.2-quick-start-guide)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/mods/2-core-architecture)
    
*   [Application Entry Point](https://deepwiki.com/charmbracelet/mods/2.1-application-entry-point)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/mods/2.2-configuration-system)
    
*   [Bubble Tea UI Model](https://deepwiki.com/charmbracelet/mods/2.3-bubble-tea-ui-model)
    
*   [Message Stream Context](https://deepwiki.com/charmbracelet/mods/2.4-message-stream-context)
    
*   [Conversation Management](https://deepwiki.com/charmbracelet/mods/2.5-conversation-management)
    
*   [LLM Provider Integration](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration)
    
*   [Provider Configuration](https://deepwiki.com/charmbracelet/mods/3.1-provider-configuration)
    
*   [Client Initialization and Streaming](https://deepwiki.com/charmbracelet/mods/3.2-client-initialization-and-streaming)
    
*   [Model Resolution and Fallbacks](https://deepwiki.com/charmbracelet/mods/3.3-model-resolution-and-fallbacks)
    
*   [User Interface and Experience](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience)
    
*   [Terminal Capabilities and Styling](https://deepwiki.com/charmbracelet/mods/4.1-terminal-capabilities-and-styling)
    
*   [Animation System](https://deepwiki.com/charmbracelet/mods/4.2-animation-system)
    
*   [Output Rendering and Formatting](https://deepwiki.com/charmbracelet/mods/4.3-output-rendering-and-formatting)
    
*   [Data Persistence](https://deepwiki.com/charmbracelet/mods/5-data-persistence)
    
*   [Conversation Database](https://deepwiki.com/charmbracelet/mods/5.1-conversation-database)
    
*   [Cache System](https://deepwiki.com/charmbracelet/mods/5.2-cache-system)
    
*   [Supporting Systems](https://deepwiki.com/charmbracelet/mods/6-supporting-systems)
    
*   [Flag Parsing and Error Handling](https://deepwiki.com/charmbracelet/mods/6.1-flag-parsing-and-error-handling)
    
*   [Content Loading](https://deepwiki.com/charmbracelet/mods/6.2-content-loading)
    
*   [Message Utilities](https://deepwiki.com/charmbracelet/mods/6.3-message-utilities)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/mods/7-development-guide)
    
*   [Dependencies and Build System](https://deepwiki.com/charmbracelet/mods/7.1-dependencies-and-build-system)
    
*   [Testing Strategy](https://deepwiki.com/charmbracelet/mods/7.2-testing-strategy)
    
*   [Code Quality and CI/CD](https://deepwiki.com/charmbracelet/mods/7.3-code-quality-and-cicd)
    

Menu

Supporting Systems
==================

Relevant source files

*   [error.go](https://github.com/charmbracelet/mods/blob/07a05d5b/error.go)
    
*   [flag.go](https://github.com/charmbracelet/mods/blob/07a05d5b/flag.go)
    
*   [flag\_test.go](https://github.com/charmbracelet/mods/blob/07a05d5b/flag_test.go)
    
*   [load.go](https://github.com/charmbracelet/mods/blob/07a05d5b/load.go)
    
*   [load\_test.go](https://github.com/charmbracelet/mods/blob/07a05d5b/load_test.go)
    

This document covers auxiliary systems that support the core application functionality. These include flag parsing and validation, error handling and formatting, content loading from various sources (strings, files, URLs), and message preprocessing utilities. For information about the main CLI entry point and command structure, see [Application Entry Point](https://deepwiki.com/charmbracelet/mods/2.1-application-entry-point)
. For details on configuration parsing, see [Configuration System](https://deepwiki.com/charmbracelet/mods/2.2-configuration-system)
.

Overview
--------

The supporting systems provide foundational services used throughout the application:

**Sources:** [flag.go1-90](https://github.com/charmbracelet/mods/blob/07a05d5b/flag.go#L1-L90)
 [error.go1-24](https://github.com/charmbracelet/mods/blob/07a05d5b/error.go#L1-L24)
 [load.go1-34](https://github.com/charmbracelet/mods/blob/07a05d5b/load.go#L1-L34)

Flag Parsing and Error Handling
-------------------------------

The flag parsing system provides custom flag types and enhanced error messages for CLI flag validation failures.

### Custom Flag Types

The application defines custom flag types to extend cobra/pflag functionality:

| Type | Purpose | Key Methods |
| --- | --- | --- |
| `durationFlag` | Parse human-readable durations (e.g., "7d", "1h30m") | `Set()`, `String()`, `Type()` |
| `flagParseError` | Structured flag parsing errors | `Error()`, `ReasonFormat()`, `Flag()` |

The `durationFlag` type wraps `time.Duration` and uses the `github.com/caarlos0/duration` package for flexible parsing:

[flag.go69-89](https://github.com/charmbracelet/mods/blob/07a05d5b/flag.go#L69-L89)

### Flag Error Parsing Flow

When cobra reports a flag parsing error, `newFlagParseError` analyzes the error message to extract structured information:

**Sources:** [flag.go11-49](https://github.com/charmbracelet/mods/blob/07a05d5b/flag.go#L11-L49)

The `newFlagParseError` function handles four primary error patterns:

1.  **Missing argument**: "flag needs an argument: --delete" or "-d"
2.  **Unknown flag**: "unknown flag: --nope"
3.  **Unknown shorthand**: "unknown shorthand flag: 'x' in -xyz"
4.  **Invalid argument**: "invalid argument "foo" for "--max-tokens" flag: ..."

[flag.go14-43](https://github.com/charmbracelet/mods/blob/07a05d5b/flag.go#L14-L43)

### Error Structure

The `flagParseError` struct provides three accessor methods for error presentation:

**Sources:** [flag.go51-67](https://github.com/charmbracelet/mods/blob/07a05d5b/flag.go#L51-L67)

The `reason` field contains a format string (e.g., "Flag %s needs an argument.") where `%s` is replaced with the extracted flag name. This enables user-friendly error display in the UI.

Test coverage validates all error patterns:

[flag\_test.go10-45](https://github.com/charmbracelet/mods/blob/07a05d5b/flag_test.go#L10-L45)

Error Handling
--------------

The error handling system provides wrappers for user-facing errors with additional context.

### Error Types

**Sources:** [error.go1-24](https://github.com/charmbracelet/mods/blob/07a05d5b/error.go#L1-L24)

### modsError Wrapper

The `modsError` type wraps errors with additional context:

[error.go11-23](https://github.com/charmbracelet/mods/blob/07a05d5b/error.go#L11-L23)

*   `err`: The underlying error
*   `reason`: Human-readable context or explanation

The `Error()` method returns the original error message, while `Reason()` provides the additional context. This allows error handlers to display both technical details and user-friendly explanations.

### newUserErrorf Helper

The `newUserErrorf` function creates simple formatted errors for user display:

[error.go5-9](https://github.com/charmbracelet/mods/blob/07a05d5b/error.go#L5-L9)

This function exists primarily to satisfy linters that complain about error messages starting with capitalized letters. It wraps `fmt.Errorf` and allows user-facing errors to begin with capital letters for better readability.

Content Loading
---------------

The content loading system enables prompts to be loaded from multiple sources: plain strings, local files via `file://` URIs, and remote content via HTTP(S) URLs.

### Loading Flow

**Sources:** [load.go10-33](https://github.com/charmbracelet/mods/blob/07a05d5b/load.go#L10-L33)

### Protocol Handling

The `loadMsg` function checks for protocol prefixes in order:

1.  **HTTP(S) URLs**: [load.go11-22](https://github.com/charmbracelet/mods/blob/07a05d5b/load.go#L11-L22)
    
    *   Performs `http.Get` request (no context for simplicity)
    *   Reads entire response body
    *   Returns body as string
    *   Defers closing response body
2.  **File URLs**: [load.go24-30](https://github.com/charmbracelet/mods/blob/07a05d5b/load.go#L24-L30)
    
    *   Strips `file://` prefix
    *   Reads entire file with `os.ReadFile`
    *   Returns content as string
3.  **Plain text**: [load.go32](https://github.com/charmbracelet/mods/blob/07a05d5b/load.go#L32-L32)
    
    *   Returns input unchanged if no protocol prefix detected

### Usage Examples

The load tests demonstrate all three loading modes:

| Test Case | Input | Expected Behavior |
| --- | --- | --- |
| Normal message | "just text" | Returns input unchanged |
| Local file | "file:///tmp/foo.txt" | Reads and returns file content |
| HTTP URL | "[http://example.com/file](http://example.com/file)<br>" | Fetches and returns remote content |
| HTTPS URL | "[https://example.com/file](https://example.com/file)<br>" | Fetches and returns remote content |

[load\_test.go11-38](https://github.com/charmbracelet/mods/blob/07a05d5b/load_test.go#L11-L38)

The HTTP(S) test validates against a real GitHub URL to ensure live URL loading works:

[load\_test.go28-38](https://github.com/charmbracelet/mods/blob/07a05d5b/load_test.go#L28-L38)

Integration with Main Application
---------------------------------

These supporting systems integrate at multiple points in the application lifecycle:

**Sources:** [flag.go11-90](https://github.com/charmbracelet/mods/blob/07a05d5b/flag.go#L11-L90)
 [load.go10-33](https://github.com/charmbracelet/mods/blob/07a05d5b/load.go#L10-L33)
 [error.go11-23](https://github.com/charmbracelet/mods/blob/07a05d5b/error.go#L11-L23)

### Flag Parsing Integration

The `durationFlag` type is used for the `--delete-older-than` flag, allowing users to specify durations like:

*   `7d` (7 days)
*   `1h30m` (1 hour 30 minutes)
*   `14d` (14 days)

[flag.go69-89](https://github.com/charmbracelet/mods/blob/07a05d5b/flag.go#L69-L89)

### Content Loading Integration

The `loadMsg` function is called during prompt processing to support use cases like:

All three modes return a string that becomes part of the message context sent to the LLM provider.

Message Utilities
-----------------

While not included in the provided source files, the codebase includes utility functions for message preprocessing and manipulation:

| Function | Purpose |
| --- | --- |
| `lastPrompt()` | Extracts the most recent user prompt from message history |
| `firstLine()` | Gets the first line of a message for display/truncation |
| `cutPrompt()` | Truncates messages to fit within token limits |
| `removeWhitespace()` | Normalizes whitespace in message content |

These utilities are used in the message stream context setup (see [Message Stream Context](https://deepwiki.com/charmbracelet/mods/2.4-message-stream-context)
) to prepare messages for transmission to LLM providers while respecting token limits and formatting requirements.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Supporting Systems](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#supporting-systems)
    
*   [Overview](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#overview)
    
*   [Flag Parsing and Error Handling](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#flag-parsing-and-error-handling)
    
*   [Custom Flag Types](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#custom-flag-types)
    
*   [Flag Error Parsing Flow](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#flag-error-parsing-flow)
    
*   [Error Structure](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#error-structure)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#error-handling)
    
*   [Error Types](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#error-types)
    
*   [modsError Wrapper](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#modserror-wrapper)
    
*   [newUserErrorf Helper](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#newusererrorf-helper)
    
*   [Content Loading](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#content-loading)
    
*   [Loading Flow](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#loading-flow)
    
*   [Protocol Handling](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#protocol-handling)
    
*   [Usage Examples](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#usage-examples)
    
*   [Integration with Main Application](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#integration-with-main-application)
    
*   [Flag Parsing Integration](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#flag-parsing-integration)
    
*   [Content Loading Integration](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#content-loading-integration)
    
*   [Message Utilities](https://deepwiki.com/charmbracelet/mods/6-supporting-systems#message-utilities)
