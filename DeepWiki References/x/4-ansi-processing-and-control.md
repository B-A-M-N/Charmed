Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/x](https://github.com/charmbracelet/x "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 7 February 2026 ([764291](https://github.com/charmbracelet/x/commits/7642919e)
)

*   [Overview](https://deepwiki.com/charmbracelet/x/1-overview)
    
*   [Core Terminal Rendering](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering)
    
*   [cellbuf Package](https://deepwiki.com/charmbracelet/x/2.1-cellbuf-package)
    
*   [Virtual Terminal Emulator (vt)](https://deepwiki.com/charmbracelet/x/2.2-virtual-terminal-emulator-(vt))
    
*   [Input Processing System](https://deepwiki.com/charmbracelet/x/3-input-processing-system)
    
*   [Input Package](https://deepwiki.com/charmbracelet/x/3.1-input-package)
    
*   [Windows Console Input](https://deepwiki.com/charmbracelet/x/3.2-windows-console-input)
    
*   [ANSI Processing and Control](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control)
    
*   [Parser and Text Manipulation](https://deepwiki.com/charmbracelet/x/4.1-parser-and-text-manipulation)
    
*   [Terminal Control Sequences](https://deepwiki.com/charmbracelet/x/4.2-terminal-control-sequences)
    
*   [Styling and Colors](https://deepwiki.com/charmbracelet/x/4.3-styling-and-colors)
    
*   [Platform Abstraction Layer](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer)
    
*   [Pseudo-Terminal Interface (xpty)](https://deepwiki.com/charmbracelet/x/5.1-pseudo-terminal-interface-(xpty))
    
*   [Windows ConPTY](https://deepwiki.com/charmbracelet/x/5.2-windows-conpty)
    
*   [Terminal I/O (termios and wcwidth)](https://deepwiki.com/charmbracelet/x/5.3-terminal-io-(termios-and-wcwidth))
    
*   [Utility Packages](https://deepwiki.com/charmbracelet/x/6-utility-packages)
    
*   [SSH Keys and Colors](https://deepwiki.com/charmbracelet/x/6.1-ssh-keys-and-colors)
    
*   [Experimental Packages](https://deepwiki.com/charmbracelet/x/6.2-experimental-packages)
    
*   [Development and Testing](https://deepwiki.com/charmbracelet/x/7-development-and-testing)
    
*   [Examples and Integration](https://deepwiki.com/charmbracelet/x/7.1-examples-and-integration)
    
*   [Testing Framework (teatest)](https://deepwiki.com/charmbracelet/x/7.2-testing-framework-(teatest))
    
*   [Repository Structure and CI/CD](https://deepwiki.com/charmbracelet/x/7.3-repository-structure-and-cicd)
    

Menu

ANSI Processing and Control
===========================

Relevant source files

*   [ansi/ctrl.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/ctrl.go)
    
*   [ansi/cursor.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/cursor.go)
    
*   [ansi/go.mod](https://github.com/charmbracelet/x/blob/7642919e/ansi/go.mod)
    
*   [ansi/go.sum](https://github.com/charmbracelet/x/blob/7642919e/ansi/go.sum)
    
*   [ansi/keypad.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/keypad.go)
    
*   [ansi/method.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/method.go)
    
*   [ansi/mode.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/mode.go)
    
*   [ansi/mode\_deprecated.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/mode_deprecated.go)
    
*   [ansi/mode\_test.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/mode_test.go)
    
*   [ansi/parser.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser.go)
    
*   [ansi/parser\_decode.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser_decode.go)
    
*   [ansi/screen.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/screen.go)
    
*   [ansi/sgr.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/sgr.go)
    
*   [ansi/status.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/status.go)
    
*   [ansi/style.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/style.go)
    
*   [ansi/style\_test.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/style_test.go)
    
*   [ansi/termcap.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/termcap.go)
    
*   [ansi/truncate.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/truncate.go)
    
*   [ansi/truncate\_test.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/truncate_test.go)
    
*   [ansi/util.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/util.go)
    
*   [ansi/width.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/width.go)
    
*   [ansi/width\_test.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/width_test.go)
    
*   [ansi/wrap.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/wrap.go)
    
*   [ansi/wrap\_test.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/wrap_test.go)
    
*   [ansi/xterm.go](https://github.com/charmbracelet/x/blob/7642919e/ansi/xterm.go)
    

Purpose and Scope
-----------------

The `ansi` package provides comprehensive facilities for working with ANSI escape sequences in terminal applications. It covers three main areas: parsing ANSI sequences from terminal input or text, generating ANSI control sequences for terminal manipulation, and performing text operations that are aware of embedded ANSI codes. This page provides an overview of these capabilities.

For specific details on:

*   Text manipulation and parsing internals, see [Parser and Text Manipulation](https://deepwiki.com/charmbracelet/x/4.1-parser-and-text-manipulation)
    
*   Terminal control sequence generation, see [Terminal Control Sequences](https://deepwiki.com/charmbracelet/x/4.2-terminal-control-sequences)
    
*   Styling and color handling, see [Styling and Colors](https://deepwiki.com/charmbracelet/x/4.3-styling-and-colors)
    
*   Virtual terminal emulation, see [Virtual Terminal Emulator (vt)](https://deepwiki.com/charmbracelet/x/2.2-virtual-terminal-emulator-(vt))
    
*   Screen buffer management, see [cellbuf Package](https://deepwiki.com/charmbracelet/x/2.1-cellbuf-package)
    

Package Architecture
--------------------

The `ansi` package is organized around three core functional areas that work together to provide complete ANSI processing capabilities:

**Sources:** [ansi/parser\_decode.go1-526](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser_decode.go#L1-L526)
 [ansi/parser.go1-418](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser.go#L1-L418)
 [ansi/width.go1-108](https://github.com/charmbracelet/x/blob/7642919e/ansi/width.go#L1-L108)
 [ansi/wrap.go1-473](https://github.com/charmbracelet/x/blob/7642919e/ansi/wrap.go#L1-L473)
 [ansi/truncate.go1-291](https://github.com/charmbracelet/x/blob/7642919e/ansi/truncate.go#L1-L291)
 [ansi/style.go1-648](https://github.com/charmbracelet/x/blob/7642919e/ansi/style.go#L1-L648)
 [ansi/cursor.go1-638](https://github.com/charmbracelet/x/blob/7642919e/ansi/cursor.go#L1-L638)
 [ansi/mode.go1-691](https://github.com/charmbracelet/x/blob/7642919e/ansi/mode.go#L1-L691)
 [ansi/screen.go1-309](https://github.com/charmbracelet/x/blob/7642919e/ansi/screen.go#L1-L309)

Core Data Structures
--------------------

The package centers around several key types that represent ANSI processing state and operations:

**Sources:** [ansi/parser.go18-53](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser.go#L18-L53)
 [ansi/parser\_decode.go10-22](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser_decode.go#L10-L22)
 [ansi/parser\_decode.go453-525](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser_decode.go#L453-L525)
 [ansi/style.go14-59](https://github.com/charmbracelet/x/blob/7642919e/ansi/style.go#L14-L59)
 [ansi/mode.go8-224](https://github.com/charmbracelet/x/blob/7642919e/ansi/mode.go#L8-L224)

State Machine Parsing
---------------------

The package uses a DEC-compatible state machine for parsing ANSI escape sequences. The `DecodeSequence` function processes input byte-by-byte, transitioning between states and returning parsed sequences:

| State | Description | Transitions To |
| --- | --- | --- |
| `NormalState` | Ground state for printable text | `EscapeState`, `PrefixState`, `StringState` |
| `EscapeState` | After ESC byte (`\x1b`) | `PrefixState`, `StringState`, back to `NormalState` |
| `PrefixState` | Collecting CSI/DCS prefix (`<`, `=`, `>`, `?`) | `ParamsState` |
| `ParamsState` | Collecting numeric parameters | `IntermedState` |
| `IntermedState` | Collecting intermediate bytes ( to `/`) | `NormalState` (on final byte) |
| `StringState` | Collecting OSC, DCS, APC, PM, SOS data | `NormalState` (on terminator) |

The state machine ensures proper handling of:

*   **CSI sequences**: `ESC [ params intermediates final` (e.g., `\x1b[31m`)\
*   **OSC sequences**: `ESC ] command ; data ST` (e.g., `\x1b]8;;url\x07`)
*   **DCS sequences**: `ESC P params intermediates final data ST`
*   **Escape sequences**: `ESC intermediates final` (e.g., `\x1b7`)

**Sources:** [ansi/parser\_decode.go10-22](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser_decode.go#L10-L22)
 [ansi/parser\_decode.go124-352](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser_decode.go#L124-L352)

Text Processing Capabilities
----------------------------

### Width Calculation Methods

The package provides two methods for calculating text width, controlled by the `Method` type:

| Method | Description | Use Case |
| --- | --- | --- |
| `GraphemeWidth` | Treats text as grapheme clusters (default) | Modern terminal applications with emoji support |
| `WcWidth` | Uses character-based width calculation | Legacy compatibility, simpler width logic |

Both methods are ANSI-aware and ignore escape sequences when calculating display width.

### Text Manipulation Functions

All text manipulation functions preserve ANSI escape codes while operating on visible characters:

| Function | Purpose | Example |
| --- | --- | --- |
| `Strip(s)` | Remove all ANSI codes | `"\x1b<FileRef file-url="https://github.com/charmbracelet/x/blob/7642919e/31mRed\\x1b[0m\"` → `\"Red\"` |\
\
Sequence Generation Categories\
------------------------------\
\
**Sources:** [ansi/cursor.go1-638](https://github.com/charmbracelet/x/blob/7642919e/ansi/cursor.go#L1-L638)\
 [ansi/screen.go1-309](https://github.com/charmbracelet/x/blob/7642919e/ansi/screen.go#L1-L309)\
 [ansi/mode.go51-208](https://github.com/charmbracelet/x/blob/7642919e/ansi/mode.go#L51-L208)\
 [ansi/style.go1-648](https://github.com/charmbracelet/x/blob/7642919e/ansi/style.go#L1-L648)\
 [ansi/sgr.go1-80](https://github.com/charmbracelet/x/blob/7642919e/ansi/sgr.go#L1-L80)\
 [ansi/status.go1-169](https://github.com/charmbracelet/x/blob/7642919e/ansi/status.go#L1-L169)\
\
Parser Usage Patterns\
---------------------\
\
The package supports two parsing approaches:\
\
### Stateless Sequence Decoding\
\
`DecodeSequence` provides a lightweight, stateless parser for processing sequences one at a time:\
\
    var state byte = NormalState\
    p := NewParser(32, 1024)\
    input := []byte("\x1b[31mHello\x1b[0m")\
    \
    for len(input) > 0 {\
        seq, width, n, newState := DecodeSequence(input, state, p)\
        // Process seq based on p.Command(), p.Params(), p.Data()\
        state = newState\
        input = input[n:]\
    }\
    \
\
Key features:\
\
*   Zero allocations for sequence parsing\
*   Returns sequence slice, display width, bytes consumed, and new state\
*   Optional `Parser` parameter for collecting parameters and data\
\
**Sources:** [ansi/parser\_decode.go24-122](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser_decode.go#L24-L122)\
\
### Handler-Based Continuous Parsing\
\
The `Parser` struct with `Handler` interface enables event-driven processing:\
\
    type Handler struct {\
        Print      func(rune)\
        Execute    func(byte)\
        HandleCsi  func(Cmd, Params)\
        HandleOsc  func(int, []byte)\
        HandleDcs  func(Cmd, Params, []byte)\
        HandleEsc  func(Cmd)\
        HandleSos  func([]byte)\
        HandlePm   func([]byte)\
        HandleApc  func([]byte)\
    }\
    \
\
The parser calls handler methods as it encounters different sequence types, enabling streaming processing of terminal output.\
\
**Sources:** [ansi/parser.go1-418](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser.go#L1-L418)\
 [ansi/handler.go1-100](https://github.com/charmbracelet/x/blob/7642919e/ansi/handler.go#L1-L100)\
\
Command and Parameter Encoding\
------------------------------\
\
The package uses bit-packing to efficiently represent CSI/DCS commands with their modifiers:\
\
| Bits | Content | Accessor |\
| --- | --- | --- |\
| 0-7 | Final command byte | `Cmd.Final()` |\
| 8-15 | Intermediate byte | `Cmd.Intermediate()` |\
| 16-23 | Prefix character | `Cmd.Prefix()` |\
\
Parameters use similar packing with the `HasMoreFlag` to indicate sub-parameters separated by colons:\
\
| Bits | Content | Accessor |\
| --- | --- | --- |\
| 0-30 | Parameter value | `Param.Param(default)` |\
| 31  | HasMore flag (sub-parameter follows) | `Param.HasMore()` |\
\
This design enables efficient storage and manipulation without heap allocations.\
\
**Sources:** [ansi/parser\_decode.go453-525](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser_decode.go#L453-L525)\
 [ansi/parser.go22-53](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser.go#L22-L53)\
\
Integration with Terminal Components\
------------------------------------\
\
The ANSI processing package serves as a foundational layer for higher-level terminal components:\
\
The ANSI package provides the primitive operations that higher-level components build upon, handling the low-level details of escape sequence encoding and decoding.\
\
**Sources:** Based on architectural overview from diagrams\
\
Package Dependencies\
--------------------\
\
The `ansi` package relies on several external libraries for text processing:\
\
| Dependency | Purpose | Files |\
| --- | --- | --- |\
| `github.com/clipperhouse/displaywidth` | Grapheme-aware width calculation | [ansi/go.mod7](https://github.com/charmbracelet/x/blob/7642919e/ansi/go.mod#L7-L7) |\
| `github.com/clipperhouse/uax29/v2/graphemes` | Unicode grapheme clustering | [ansi/go.mod8](https://github.com/charmbracelet/x/blob/7642919e/ansi/go.mod#L8-L8) |\
| `github.com/mattn/go-runewidth` | Character width calculation (WcWidth mode) | [ansi/go.mod10](https://github.com/charmbracelet/x/blob/7642919e/ansi/go.mod#L10-L10) |\
| `github.com/lucasb-eyer/go-colorful` | Color space conversions | [ansi/go.mod9](https://github.com/charmbracelet/x/blob/7642919e/ansi/go.mod#L9-L9) |\
\
The package uses `displaywidth` and `graphemes` for the default `GraphemeWidth` method and `go-runewidth` for the `WcWidth` method, providing flexibility in handling Unicode text.\
\
**Sources:** [ansi/go.mod1-14](https://github.com/charmbracelet/x/blob/7642919e/ansi/go.mod#L1-L14)\
 [ansi/method.go1-101](https://github.com/charmbracelet/x/blob/7642919e/ansi/method.go#L1-L101)\
\
Dismiss\
\
Refresh this wiki\
\
Enter email to refresh\
\
### On this page\
\
*   [ANSI Processing and Control](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#ansi-processing-and-control)\
    \
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#purpose-and-scope)\
    \
*   [Package Architecture](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#package-architecture)\
    \
*   [Core Data Structures](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#core-data-structures)\
    \
*   [State Machine Parsing](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#state-machine-parsing)\
    \
*   [Text Processing Capabilities](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#text-processing-capabilities)\
    \
*   [Width Calculation Methods](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#width-calculation-methods)\
    \
*   [Text Manipulation Functions](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#text-manipulation-functions)\
    \
*   [Sequence Generation Categories](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#sequence-generation-categories)\
    \
*   [Parser Usage Patterns](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#parser-usage-patterns)\
    \
*   [Stateless Sequence Decoding](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#stateless-sequence-decoding)\
    \
*   [Handler-Based Continuous Parsing](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#handler-based-continuous-parsing)\
    \
*   [Command and Parameter Encoding](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#command-and-parameter-encoding)\
    \
*   [Integration with Terminal Components](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#integration-with-terminal-components)\
    \
*   [Package Dependencies](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control#package-dependencies)
