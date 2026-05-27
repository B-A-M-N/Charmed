Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/sequin](https://github.com/charmbracelet/sequin "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 20 April 2025 ([025641](https://github.com/charmbracelet/sequin/commits/025641c1)
)

*   [Introduction to Sequin](https://deepwiki.com/charmbracelet/sequin/1-introduction-to-sequin)
    
*   [Installation](https://deepwiki.com/charmbracelet/sequin/1.1-installation)
    
*   [Basic Usage](https://deepwiki.com/charmbracelet/sequin/1.2-basic-usage)
    
*   [Architecture](https://deepwiki.com/charmbracelet/sequin/2-architecture)
    
*   [ANSI Sequence Processing](https://deepwiki.com/charmbracelet/sequin/2.1-ansi-sequence-processing)
    
*   [Command Execution](https://deepwiki.com/charmbracelet/sequin/2.2-command-execution)
    
*   [ANSI Sequence Handlers](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers)
    
*   [SGR (Select Graphic Rendition)](https://deepwiki.com/charmbracelet/sequin/3.1-sgr-(select-graphic-rendition))
    
*   [Cursor Commands](https://deepwiki.com/charmbracelet/sequin/3.2-cursor-commands)
    
*   [Screen and Line Commands](https://deepwiki.com/charmbracelet/sequin/3.3-screen-and-line-commands)
    
*   [Window Title and Colors](https://deepwiki.com/charmbracelet/sequin/3.4-window-title-and-colors)
    
*   [Terminal Modes](https://deepwiki.com/charmbracelet/sequin/3.5-terminal-modes)
    
*   [Pointer, Clipboard, and Links](https://deepwiki.com/charmbracelet/sequin/3.6-pointer-clipboard-and-links)
    
*   [FinalTerm Commands](https://deepwiki.com/charmbracelet/sequin/3.7-finalterm-commands)
    
*   [Theme System](https://deepwiki.com/charmbracelet/sequin/4-theme-system)
    
*   [Testing](https://deepwiki.com/charmbracelet/sequin/5-testing)
    
*   [Build and Development](https://deepwiki.com/charmbracelet/sequin/6-build-and-development)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/sequin/6.1-dependency-management)
    
*   [Release Process](https://deepwiki.com/charmbracelet/sequin/6.2-release-process)
    

Menu

ANSI Sequence Handlers
======================

Relevant source files

*   [cursor.go](https://github.com/charmbracelet/sequin/blob/025641c1/cursor.go)
    
*   [handlers.go](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go)
    
*   [kitty.go](https://github.com/charmbracelet/sequin/blob/025641c1/kitty.go)
    
*   [main\_test.go](https://github.com/charmbracelet/sequin/blob/025641c1/main_test.go)
    
*   [sgr.go](https://github.com/charmbracelet/sequin/blob/025641c1/sgr.go)
    

This document provides an overview of the ANSI sequence handler system in Sequin. This system is responsible for interpreting various types of ANSI escape sequences and translating them into human-readable explanations. For detailed information about specific handler types, see the child pages: [SGR (Select Graphic Rendition)](https://deepwiki.com/charmbracelet/sequin/3.1-sgr-(select-graphic-rendition))
, [Cursor Commands](https://deepwiki.com/charmbracelet/sequin/3.2-cursor-commands)
, and other specific handler documentation.

Handler System Overview
-----------------------

Sequin uses a dispatch-based architecture to route different types of ANSI escape sequences to their appropriate handlers. Each type of sequence (such as text styling, cursor movement, or window title commands) is processed by a specialized handler function that knows how to interpret its parameters.

Sources: [handlers.go10-86](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L10-L86)
 [main\_test.go262-280](https://github.com/charmbracelet/sequin/blob/025641c1/main_test.go#L262-L280)

Handler Registration System
---------------------------

The handler registration system is implemented through a set of maps that associate specific sequence identifiers with handler functions. These maps are defined in `handlers.go` and include:

Each handler function has the signature:

The map keys are calculated based on the sequence's final character, with additional bits set to represent intermediate characters and marker characters:

*   The base key is the final character (e.g., 'm' for SGR sequences)
*   Intermediate characters are shifted by 8 bits (`intermedShift`)
*   Marker characters (like '?' in private sequences) are shifted by 16 bits (`markerShift`)

This encoding system allows different handlers to be registered for sequences that share the same final character but have different prefixes or intermediate characters.

Sources: [handlers.go10-52](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L10-L52)
 [handlers.go54-70](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L54-L70)
 [handlers.go72-85](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L72-L85)
 [handlers.go92-98](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L92-L98)

Handler Types
-------------

Sequin includes handlers for various categories of ANSI sequences:

### CSI (Control Sequence Introducer) Handlers

CSI sequences start with `ESC <FileRef file-url="https://github.com/charmbracelet/sequin/blob/025641c1/` and include#LNaN-LNaN" NaN file-path="\` and include">Hii [sgr.go12-111](https://github.com/charmbracelet/sequin/blob/025641c1/sgr.go#L12-L111)
 [cursor.go10-80](https://github.com/charmbracelet/sequin/blob/025641c1/cursor.go#L10-L80)

### OSC (Operating System Command) Handlers

OSC sequences start with `ESC ]` and include:

| Handler | Parameter | Description |
| --- | --- | --- |
| Title Handler | 0, 1, 2 | Sets window/icon title |
| Working Directory Handler | 7   | Reports current working directory |
| Hyperlink Handler | 8   | Creates hyperlinks in terminal |
| Notify Handler | 9   | Sends desktop notifications |
| Terminal Color Handler | 10, 11, 12 | Sets/resets terminal colors |
| Pointer Shape Handler | 22  | Sets mouse pointer shape |
| Clipboard Handler | 52  | Controls clipboard operations |
| FinalTerm Handler | 133 | Handles FinalTerm shell integration |

Sources: [handlers.go54-70](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L54-L70)
 [main\_test.go195-254](https://github.com/charmbracelet/sequin/blob/025641c1/main_test.go#L195-L254)

### Other Handler Types

*   DCS (Device Control String) Handlers: For device control commands
*   ESC Handlers: For simple escape sequences like saving/restoring cursor position or keypad mode changes

Sources: [handlers.go72-85](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L72-L85)
 [main\_test.go255-259](https://github.com/charmbracelet/sequin/blob/025641c1/main_test.go#L255-L259)

Implementation Details
----------------------

### SGR Handler

The SGR (Select Graphic Rendition) handler is one of the most complex handlers, supporting:

*   Basic text styles (bold, italic, underline, etc.)
*   ANSI basic colors (8 colors)
*   ANSI bright colors (8 bright colors)
*   256-color palette
*   24-bit RGB colors
*   Underline styles (single, double, curly, dotted, dashed)

The handler parses SGR parameters and translates them into human-readable descriptions.

Sources: [sgr.go12-111](https://github.com/charmbracelet/sequin/blob/025641c1/sgr.go#L12-L111)
 [sgr.go113-122](https://github.com/charmbracelet/sequin/blob/025641c1/sgr.go#L113-L122)
 [sgr.go124-168](https://github.com/charmbracelet/sequin/blob/025641c1/sgr.go#L124-L168)

### Cursor Handler

The cursor handler interprets sequences that control cursor movement and appearance:

The cursor handler function examines the final character of the CSI sequence and the parameters to determine the specific cursor operation being performed.

Sources: [cursor.go10-60](https://github.com/charmbracelet/sequin/blob/025641c1/cursor.go#L10-L60)
 [cursor.go63-80](https://github.com/charmbracelet/sequin/blob/025641c1/cursor.go#L63-L80)

### Kitty Keyboard Handler

The Kitty Keyboard Protocol handler deserves special mention as it processes keyboard enhancement sequences specific to the Kitty terminal:

Sources: [kitty.go11-66](https://github.com/charmbracelet/sequin/blob/025641c1/kitty.go#L11-L66)

Handler Processing Flow
-----------------------

When Sequin processes input, handlers follow this sequence:

The process works as follows:

1.  The `ansi.Parser` identifies and parses an escape sequence
2.  The sequence is passed to the appropriate handler based on the maps in `handlers.go`
3.  The handler processes the sequence parameters and generates a human-readable description
4.  The description is formatted and displayed to the user

Sources: [handlers.go10-86](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L10-L86)
 [sgr.go12-111](https://github.com/charmbracelet/sequin/blob/025641c1/sgr.go#L12-L111)
 [cursor.go10-60](https://github.com/charmbracelet/sequin/blob/025641c1/cursor.go#L10-L60)

Error Handling
--------------

The handler system uses two main error types:

*   `errUnhandled`: Returned when a sequence is recognized but not yet implemented
*   `errInvalid`: Returned when a sequence is malformed or contains invalid parameters

Sources: [handlers.go87-90](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L87-L90)

Extending the Handler System
----------------------------

To add support for a new ANSI sequence type:

1.  Create a handler function with the signature `func(*ansi.Parser) (string, error)`
2.  Add an entry to the appropriate handler map in `handlers.go`
3.  Implement the logic to interpret the sequence parameters and generate a description

The helper function `printf` can be used for simple handlers that just return a fixed description:

Sources: [handlers.go92-98](https://github.com/charmbracelet/sequin/blob/025641c1/handlers.go#L92-L98)

This architecture makes the system easily extensible to support additional sequence types as needed.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [ANSI Sequence Handlers](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#ansi-sequence-handlers)
    
*   [Handler System Overview](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#handler-system-overview)
    
*   [Handler Registration System](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#handler-registration-system)
    
*   [Handler Types](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#handler-types)
    
*   [CSI (Control Sequence Introducer) Handlers](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#csi-control-sequence-introducer-handlers)
    
*   [OSC (Operating System Command) Handlers](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#osc-operating-system-command-handlers)
    
*   [Other Handler Types](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#other-handler-types)
    
*   [Implementation Details](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#implementation-details)
    
*   [SGR Handler](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#sgr-handler)
    
*   [Cursor Handler](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#cursor-handler)
    
*   [Kitty Keyboard Handler](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#kitty-keyboard-handler)
    
*   [Handler Processing Flow](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#handler-processing-flow)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#error-handling)
    
*   [Extending the Handler System](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers#extending-the-handler-system)
