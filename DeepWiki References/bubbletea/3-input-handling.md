Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 8 March 2026 ([f25595](https://github.com/charmbracelet/bubbletea/commits/f25595a8)
)

*   [Overview](https://deepwiki.com/charmbracelet/bubbletea/1-overview)
    
*   [The Elm Architecture (MVU Pattern)](https://deepwiki.com/charmbracelet/bubbletea/1.1-the-elm-architecture-(mvu-pattern))
    
*   [Installation and Dependencies](https://deepwiki.com/charmbracelet/bubbletea/1.2-installation-and-dependencies)
    
*   [Core Components](https://deepwiki.com/charmbracelet/bubbletea/2-core-components)
    
*   [Program Lifecycle](https://deepwiki.com/charmbracelet/bubbletea/2.1-program-lifecycle)
    
*   [Model Interface](https://deepwiki.com/charmbracelet/bubbletea/2.2-model-interface)
    
*   [Message System](https://deepwiki.com/charmbracelet/bubbletea/2.3-message-system)
    
*   [Command System](https://deepwiki.com/charmbracelet/bubbletea/2.4-command-system)
    
*   [Program Options](https://deepwiki.com/charmbracelet/bubbletea/2.5-program-options)
    
*   [Input Handling](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling)
    
*   [Keyboard Input](https://deepwiki.com/charmbracelet/bubbletea/3.1-keyboard-input)
    
*   [Mouse Input](https://deepwiki.com/charmbracelet/bubbletea/3.2-mouse-input)
    
*   [Terminal Events](https://deepwiki.com/charmbracelet/bubbletea/3.3-terminal-events)
    
*   [Cross-platform Input](https://deepwiki.com/charmbracelet/bubbletea/3.4-cross-platform-input)
    
*   [Rendering System](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system)
    
*   [Renderer Interface and View Structure](https://deepwiki.com/charmbracelet/bubbletea/4.1-renderer-interface-and-view-structure)
    
*   [Cursed Renderer Implementation](https://deepwiki.com/charmbracelet/bubbletea/4.2-cursed-renderer-implementation)
    
*   [Terminal Management](https://deepwiki.com/charmbracelet/bubbletea/4.3-terminal-management)
    
*   [Terminal Features and Capabilities](https://deepwiki.com/charmbracelet/bubbletea/4.4-terminal-features-and-capabilities)
    
*   [Advanced Topics](https://deepwiki.com/charmbracelet/bubbletea/5-advanced-topics)
    
*   [Concurrency and Goroutines](https://deepwiki.com/charmbracelet/bubbletea/5.1-concurrency-and-goroutines)
    
*   [Error Handling and Recovery](https://deepwiki.com/charmbracelet/bubbletea/5.2-error-handling-and-recovery)
    
*   [Message Filtering](https://deepwiki.com/charmbracelet/bubbletea/5.3-message-filtering)
    
*   [Testing Bubble Tea Applications](https://deepwiki.com/charmbracelet/bubbletea/5.4-testing-bubble-tea-applications)
    
*   [External I/O and Side Effects](https://deepwiki.com/charmbracelet/bubbletea/5.5-external-io-and-side-effects)
    
*   [Platform-Specific Features](https://deepwiki.com/charmbracelet/bubbletea/5.6-platform-specific-features)
    
*   [Examples and Tutorials](https://deepwiki.com/charmbracelet/bubbletea/6-examples-and-tutorials)
    
*   [Simple Examples](https://deepwiki.com/charmbracelet/bubbletea/6.1-simple-examples)
    
*   [Multi-View Application Example](https://deepwiki.com/charmbracelet/bubbletea/6.2-multi-view-application-example)
    
*   [HTTP and Async Operations Tutorial](https://deepwiki.com/charmbracelet/bubbletea/6.3-http-and-async-operations-tutorial)
    
*   [Command Orchestration Examples](https://deepwiki.com/charmbracelet/bubbletea/6.4-command-orchestration-examples)
    
*   [Basics Tutorial (Shopping List)](https://deepwiki.com/charmbracelet/bubbletea/6.5-basics-tutorial-(shopping-list))
    
*   [Component Library Examples](https://deepwiki.com/charmbracelet/bubbletea/6.6-component-library-examples)
    
*   [Contributing](https://deepwiki.com/charmbracelet/bubbletea/7-contributing)
    
*   [Development Setup](https://deepwiki.com/charmbracelet/bubbletea/7.1-development-setup)
    
*   [Testing Guidelines](https://deepwiki.com/charmbracelet/bubbletea/7.2-testing-guidelines)
    
*   [Code Quality Standards](https://deepwiki.com/charmbracelet/bubbletea/7.3-code-quality-standards)
    
*   [CI/CD Pipeline](https://deepwiki.com/charmbracelet/bubbletea/7.4-cicd-pipeline)
    

Menu

Input Handling
==============

Relevant source files

*   [key.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go)
    
*   [mouse.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/mouse.go)
    

Input handling in Bubble Tea transforms raw terminal input into typed messages that are dispatched to the application's `Update` method. The system operates asynchronously through a dedicated `readLoop()` goroutine that reads from stdin (or an alternate input source), parses ANSI escape sequences using the ultraviolet library, and sends messages to the main event loop via a channel.

The input system handles three primary categories of events:

*   **Keyboard events** - Key presses and releases with modifier support ([Keyboard Input](https://deepwiki.com/charmbracelet/bubbletea/3.1-keyboard-input)
    )
*   **Mouse events** - Click, release, wheel, and motion events ([Mouse Input](https://deepwiki.com/charmbracelet/bubbletea/3.2-mouse-input)
    )
*   **Terminal events** - Window resize, focus/blur, and system signals ([Terminal Events](https://deepwiki.com/charmbracelet/bubbletea/3.3-terminal-events)
    )

Platform-specific handling for Unix and Windows systems is covered in [Cross-platform Input](https://deepwiki.com/charmbracelet/bubbletea/3.4-cross-platform-input)
.

Architecture Overview
---------------------

The input handling system consists of multiple layers that work together to transform raw terminal bytes into structured messages. The architecture is designed for non-blocking operation with graceful shutdown support.

**Input System Architecture**

The input pipeline operates asynchronously in a dedicated goroutine to prevent blocking the main event loop. Raw bytes flow through the ultraviolet library's parser, which handles ANSI escape sequence decoding and produces strongly-typed messages.

Sources: [key.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go)
 [mouse.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/mouse.go)

Message Types
-------------

Bubble Tea defines several message types for input events. Each message type corresponds to a specific category of terminal input:

| Message Type | Description | Child Page |
| --- | --- | --- |
| `KeyPressMsg` | Keyboard key press events with modifiers | [3.1 Keyboard Input](https://deepwiki.com/charmbracelet/bubbletea/3.1-keyboard-input) |
| `KeyReleaseMsg` | Keyboard key release events (Kitty protocol) | [3.1 Keyboard Input](https://deepwiki.com/charmbracelet/bubbletea/3.1-keyboard-input) |
| `MouseClickMsg` | Mouse button press events | [3.2 Mouse Input](https://deepwiki.com/charmbracelet/bubbletea/3.2-mouse-input) |
| `MouseReleaseMsg` | Mouse button release events | [3.2 Mouse Input](https://deepwiki.com/charmbracelet/bubbletea/3.2-mouse-input) |
| `MouseWheelMsg` | Mouse wheel scroll events | [3.2 Mouse Input](https://deepwiki.com/charmbracelet/bubbletea/3.2-mouse-input) |
| `MouseMotionMsg` | Mouse motion and drag events | [3.2 Mouse Input](https://deepwiki.com/charmbracelet/bubbletea/3.2-mouse-input) |
| `WindowSizeMsg` | Terminal window resize events | [3.3 Terminal Events](https://deepwiki.com/charmbracelet/bubbletea/3.3-terminal-events) |
| `FocusMsg` / `BlurMsg` | Terminal focus/blur events | [3.3 Terminal Events](https://deepwiki.com/charmbracelet/bubbletea/3.3-terminal-events) |
| `PasteMsg` | Bracketed paste events | [3.3 Terminal Events](https://deepwiki.com/charmbracelet/bubbletea/3.3-terminal-events) |

All messages implement the `Msg` interface and are sent through the `p.msgs` channel to the event loop. See [Message System](https://deepwiki.com/charmbracelet/bubbletea/2.3-message-system)
 for details on message flow and handling.

Sources: [key.go190-255](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L190-L255)
 [mouse.go44-144](https://github.com/charmbracelet/bubbletea/blob/f25595a8/mouse.go#L44-L144)

Input Reading Architecture
--------------------------

The input system consists of three layers: TTY initialization, the read loop, and message detection.

### TTY Initialization and Raw Mode

Before reading input, the terminal must be configured in raw mode to receive individual keypresses and escape sequences rather than line-buffered input.

**Platform-Specific Terminal Initialization**

Sources: [tty.go25-37](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L25-L37)
 [tty\_unix.go15-30](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_unix.go#L15-L30)
 [tty\_windows.go14-55](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_windows.go#L14-L55)

| Platform | File | Special Requirements |
| --- | --- | --- |
| Unix-like systems | [tty\_unix.go15-30](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_unix.go#L15-L30) | Sets raw mode via `term.MakeRaw()` |
| Windows | [tty\_windows.go14-55](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_windows.go#L14-L55) | Requires VT input/output mode flags: `ENABLE_VIRTUAL_TERMINAL_INPUT` and `ENABLE_VIRTUAL_TERMINAL_PROCESSING` |

The `initInput()` function saves the previous terminal state in `previousTtyInputState` so it can be restored when the program exits via `restoreInput()` [tty.go62-75](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L62-L75)

### Asynchronous Read Loop

The `readLoop()` goroutine [tty.go96-106](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L96-L106)
 runs continuously until cancelled, reading from `p.cancelReader` and forwarding parsed messages to `p.msgs` channel. The loop exits cleanly on cancellation without leaving blocked goroutines.

**Asynchronous Read Loop Lifecycle**

The `cancelreader` package (from `github.com/muesli/cancelreader`) provides non-blocking I/O with cancellation. When `Cancel()` is called, `Read()` returns `ErrCanceled`, allowing graceful shutdown.

Sources: [tty.go78-106](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L78-L106)
 [key.go558-607](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L558-L607)
 [input.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/input.go)

Message Detection and Parsing
-----------------------------

The `detectOneMsg()` function [key.go614-715](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L614-L715)
 is the core parser that examines raw input bytes and produces typed messages. It handles multiple input types in a specific precedence order.

### Detection Precedence

The `detectOneMsg()` function [key.go614-715](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L614-L715)
 implements a priority-based parsing algorithm. Detection order matters because some sequences are prefixes of others.

**detectOneMsg() Decision Tree with Code References**

| Priority | Check | Return Type | Code Reference |
| --- | --- | --- | --- |
| 1   | Mouse X10 (`\x1b<FileRef file-url="https://github.com/charmbracelet/bubbletea/blob/f25595a8/M`) | `MouseMsg` | \[key.go#L618-L621" min=618 max=621 file-path="M\`) |\
| 2   | Mouse SGR (`\x1b<FileRef file-url="https://github.com/charmbracelet/bubbletea/blob/f25595a8/<`) | `MouseMsg` | \[key.go#L622-L627" min=622 max=627 file-path="<\`) |\
| 3   | Focus/Blur (`\x1b<FileRef file-url="https://github.com/charmbracelet/bubbletea/blob/f25595a8/I`, `\\x1b[O`) | `FocusMsg`, `BlurMsg` | \[key.go#L632-L636" min=632 max=636 file-path="I`,` \\x1b\[O\`) |\
| 4   | Bracketed paste (`\x1b<FileRef file-url="https://github.com/charmbracelet/bubbletea/blob/f25595a8/200~`) | `KeyMsg` (Paste=true) | \[key.go#L639-L643" min=639 max=643 file-path="200~\`) |\
| 5   | Known sequences (from `extSequences` map) | `KeyMsg` | [key.go648-652](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L648-L652) |\
| 6   | NUL byte (`\x00`) | `KeyMsg` (Type=keyNUL) | [key.go664-666](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L664-L666) |\
| 7   | UTF-8 runes | `KeyMsg` (Type=KeyRunes) | [key.go668-703](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L668-L703) |\
| 8   | Invalid byte | `unknownInputByteMsg` | [key.go714](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L714-L714) |\
\
Sources: [key.go614-715](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L614-L715)\
\
### Input Buffer Management\
\
The `readAnsiInputs()` function [key.go558-607](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L558-L607)\
 handles partial sequences across read boundaries using a buffer strategy:\
\
| Variable | Purpose | Code Reference |\
| --- | --- | --- |\
| `buf [256]byte` | Fixed-size read buffer | [key.go559](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L559-L559) |\
| `leftOverFromPrevIteration []byte` | Partial sequence storage | [key.go561](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L561-L561) |\
| `canHaveMoreData bool` | `true` if buffer was full (may have more OS buffer data) | [key.go581](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L581-L581) |\
| `w int` | Width of detected message (0 means incomplete) | [key.go586-592](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L586-L592) |\
\
**Buffer Behavior**\
\
| Scenario | `canHaveMoreData` | Action |\
| --- | --- | --- |\
| Short read (`numBytes < 256`) | `false` | Process all bytes, no partial sequences possible |\
| Full buffer (`numBytes == 256`) | `true` | May have partial sequence at end, save if `w == 0` |\
| Parser returns `w=0` | N/A | Save `b[i:]` to `leftOverFromPrevIteration`, retry next iteration |\
\
This ensures multi-byte sequences like `\x1b<FileRef file-url="https://github.com/charmbracelet/bubbletea/blob/f25595a8/1;5A` (Ctrl+Up, 6 bytes) are fully read before parsing, even if split across read operations.\\n\\nSources#LNaN-LNaN" NaN file-path="1;5A\` (Ctrl+Up, 6 bytes) are fully read before parsing, even if split across read operations.\\n\\nSources">Hii\
\
Key Message Structure\
---------------------\
\
The `Key` and `KeyMsg` types represent keyboard events. `KeyMsg` is sent to the application's `Update` method.\
\
**Key Message Type Hierarchy**\
\
Sources: [key.go45-59](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L45-L59)\
 [key.go108](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L108-L108)\
\
### KeyType Constants\
\
Bubble Tea defines three categories of key types:\
\
#### Control Characters\
\
Control characters map directly to ASCII control codes (0x00-0x1F and 0x7F):\
\
| KeyType Constant | Value | Description |\
| --- | --- | --- |\
| `KeyNull` / `KeyCtrlAt` | 0   | NUL / Ctrl+@ |\
| `KeyCtrlA` - `KeyCtrlZ` | 1-26 | Control letter keys |\
| `KeyTab` / `KeyCtrlI` | 9   | Tab character |\
| `KeyEnter` / `KeyCtrlM` | 13  | Enter/Return |\
| `KeyEsc` / `KeyEscape` | 27  | Escape key |\
| `KeyBackspace` / `KeyCtrlQuestionMark` | 127 | Backspace/Delete |\
\
Full definitions: [key.go122-201](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L122-L201)\
\
#### Navigation and Special Keys\
\
Navigation keys are represented by negative `KeyType` values starting from -1:\
\
| KeyType | String Representation |\
| --- | --- |\
| `KeyRunes` | (varies by rune content) |\
| `KeyUp`, `KeyDown`, `KeyRight`, `KeyLeft` | "up", "down", "right", "left" |\
| `KeyHome`, `KeyEnd` | "home", "end" |\
| `KeyPgUp`, `KeyPgDown` | "pgup", "pgdown" |\
| `KeyShiftTab` | "shift+tab" |\
| `KeyDelete`, `KeyInsert` | "delete", "insert" |\
\
With modifier combinations:\
\
*   `KeyCtrlUp`, `KeyCtrlDown`, etc. - Control modifier\
*   `KeyShiftUp`, `KeyShiftDown`, etc. - Shift modifier\
*   `KeyCtrlShiftUp`, `KeyCtrlShiftDown`, etc. - Both modifiers\
\
Full definitions: [key.go204-258](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L204-L258)\
\
#### Function Keys\
\
Function keys F1-F20 are defined as negative `KeyType` values:\
\
    KeyF1 through KeyF20\
    \
\
Definitions: [key.go238-258](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L238-L258)\
\
### Key String Representation\
\
The `Key.String()` method [key.go67-92](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L67-L92)\
 produces human-readable strings suitable for key comparison:\
\
The bracketed format `[...]` for pasted text ensures that shortcuts defined via `bubbles/keys` don't accidentally trigger from pasted content [key.go74-85](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L74-L85)\
\
Sources: [key.go67-92](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L67-L92)\
\
ANSI Escape Sequence Parsing\
----------------------------\
\
Terminal input includes ANSI escape sequences that encode special keys, modifiers, and terminal capabilities. Bubble Tea uses a map-based longest-prefix-match algorithm to parse these sequences.\
\
### Sequence Database\
\
The `sequences` map [key.go354-532](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L354-L532)\
 contains 150+ ANSI escape sequence mappings to `Key` structs. The map is used at startup to build `extSequences` [key\_sequences.go14-37](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L14-L37)\
 which includes Alt-modified variants and control characters.\
\
**Sequence Database Structure**\
\
**Terminal Emulator Coverage**\
\
| Terminal Type | Example Sequences | Code Reference |\
| --- | --- | --- |\
| **xterm/VT100** | `\x1b<FileRef file-url="https://github.com/charmbracelet/bubbletea/blob/f25595a8/A` (Up), `\\x1b[H` (Home), `\\x1bOP` (F1) | \[key.go#L356-L365" min=356 max=365 file-path="A`(Up),`\\x1b\[H`(Home),`\\x1bOP\` (F1) |\
| **urxvt** | `\x1b<FileRef file-url="https://github.com/charmbracelet/bubbletea/blob/f25595a8/a` (Shift+Up), `\\x1b[7~` (Home), `\\x1b[11~` (F1) | \[key.go#L368-L371" min=368 max=371 file-path="a`(Shift+Up),`\\x1b\[7~`(Home),`\\x1b\[11~\` (F1) |\
| **Linux console** | `\x1b[[A]` through `\x1b[[E]` (F1-F5) | [key.go455-459](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L455-L459) |\
| **PowerShell** | `\x1bOA` (Up), `\x1bOB` (Down), `\x1bOC` (Right), `\x1bOD` (Left) | [key.go528-532](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L528-L532) |\
\
Sources: [key.go354-532](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L354-L532)\
 [key\_sequences.go14-37](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L14-L37)\
\
### Sequence Detection Algorithm\
\
The `detectSequence()` function [key\_sequences.go56-74](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L56-L74)\
 implements longest-prefix-match using a pre-computed hash map and length index.\
\
**detectSequence() Longest-Prefix-Match Algorithm**\
\
**Data Structures**\
\
| Variable | Type | Purpose | Initialization |\
| --- | --- | --- | --- |\
| `extSequences` | `map[string]Key` | Hash map of all sequences | [key\_sequences.go14-37](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L14-L37)<br> (init time) |\
| `seqLengths` | `[]int` | Valid sequence lengths, descending | [key\_sequences.go41-52](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L41-L52)<br> (init time) |\
| `unknownCSIRe` | `*regexp.Regexp` | Matches unrecognized CSI sequences | [key.go610](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L610-L610)<br> (compile time) |\
\
The `extSequences` map is built from:\
\
1.  Base `sequences` map [key.go354-532](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L354-L532)\
    \
2.  Alt-prefixed versions (`"\x1b" + seq`)\
3.  Control characters (0x01-0x1F, 0x7F) with and without Alt\
4.  Special cases: space, Alt+space, Alt+Escape [key\_sequences.go33-35](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L33-L35)\
    \
\
The `seqLengths` array enables O(L) algorithm where L is the number of distinct sequence lengths (~7), rather than O(N) where N is number of sequences (~300).\
\
Sources: [key\_sequences.go56-74](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L56-L74)\
 [key\_sequences.go14-52](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L14-L52)\
 [key.go610](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L610-L610)\
\
### Special Sequence Handling\
\
#### Bracketed Paste\
\
Bracketed paste mode wraps pasted text between escape sequences to distinguish it from typed input:\
\
    Start: \x1b<FileRef file-url="https://github.com/charmbracelet/bubbletea/blob/f25595a8/200~\n<pasted content>\nEnd#LNaN-LNaN" NaN  file-path="200~\n<pasted content>\nEnd">Hii</FileRef> parses this format:\
    \
    ```go\
    func detectBracketedPaste(input []byte) (hasBp bool, width int, msg Msg) {\
        const bpStart = "\x1b[200~"\
        if len(input) < len(bpStart) || string(input[:len(bpStart)]) != bpStart {\
            return false, 0, nil\
        }\
        \
        const bpEnd = "\x1b[201~"\
        idx := bytes.Index(input[len(bpStart):], []byte(bpEnd))\
        if idx == -1 {\
            // End marker not found, need more data\
            return true, 0, nil\
        }\
        \
        paste := input[len(bpStart):len(bpStart)+idx]\
        k := Key{Type: KeyRunes, Paste: true}\
        // Decode all runes from paste content\
        for len(paste) > 0 {\
            r, w := utf8.DecodeRune(paste)\
            if r != utf8.RuneError {\
                k.Runes = append(k.Runes, r)\
            }\
            paste = paste[w:]\
        }\
        return true, len(bpStart) + idx + len(bpEnd), KeyMsg(k)\
    }\
    \
\
**Partial Buffer Handling**\
\
If the end marker `\x1b<FileRef file-url="https://github.com/charmbracelet/bubbletea/blob/f25595a8/201~` is not found, the function returns `width=0` \[key\_sequences.go#L98-L102" min=98 max=102 file-path="201~`is not found, the function returns`width=0`[key_sequences.go">Hii</FileRef> signaling`readAnsiInputs()\` to save the buffer and retry with more data. This handles paste operations that exceed the 256-byte read buffer.\
\
**String Representation**\
\
Pasted text is wrapped in brackets to prevent matching key bindings [key.go74-85](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L74-L85)\
:\
\
Sources: [key\_sequences.go82-118](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L82-L118)\
 [key.go74-85](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L74-L85)\
\
#### Focus Events\
\
The `detectReportFocus()` function [key\_sequences.go121-129](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L121-L129)\
 detects terminal focus/blur events:\
\
| Sequence | Length | Message Type | Meaning |\
| --- | --- | --- | --- |\
| `\x1b[I` | 3 bytes | `FocusMsg{}` | Terminal window gained focus |\
| `\x1b[O` | 3 bytes | `BlurMsg{}` | Terminal window lost focus |\
\
Focus reporting must be enabled via `WithReportFocus()` program option or the renderer's `enableReportFocus()` method. See [Terminal Events](https://deepwiki.com/charmbracelet/bubbletea/3.3-terminal-events)\
 for details.\
\
Sources: [key\_sequences.go121-129](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_sequences.go#L121-L129)\
\
Rune Handling and UTF-8\
-----------------------\
\
When no escape sequence or control character is detected, the parser accumulates valid UTF-8 runes.\
\
### Multi-Rune Keys\
\
The parser collects consecutive valid runes into a single `KeyMsg`:\
\
The loop in `detectOneMsg()` [key.go671-685](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L671-L685)\
 continues until it encounters:\
\
*   A rune decode error\
*   A control character (≤ 0x1F or 0x7F)\
*   A space character\
*   The end of the input buffer (with `canHaveMoreData=true`)\
\
For Alt-modified keys, only a single rune is collected after the escape prefix [key.go680-684](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L680-L684)\
\
Sources: [key.go668-703](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L668-L703)\
\
### Invalid UTF-8 Handling\
\
Invalid UTF-8 bytes produce `unknownInputByteMsg` messages:\
\
These messages are generated but not handled by the default event loop, making them useful for debugging input issues [key.go534-542](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L534-L542)\
\
Sources: [key.go534-542](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L534-L542)\
 [key.go714](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key.go#L714-L714)\
\
Terminal Restoration and Suspension\
-----------------------------------\
\
The input system must properly restore terminal state when the program exits or suspends.\
\
### State Restoration\
\
The `restoreTerminalState()` function [tty.go41-60](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L41-L60)\
 reverses all terminal modifications:\
\
1.  Disable bracketed paste mode\
2.  Show cursor\
3.  Disable mouse reporting\
4.  Disable focus reporting\
5.  Exit alternate screen buffer (if used)\
6.  Restore original TTY state via `term.Restore()` [tty.go63-75](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L63-L75)\
    \
\
Sources: [tty.go41-75](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L41-L75)\
\
### Suspend/Resume on Unix\
\
On Unix-like systems, Bubble Tea supports suspend (Ctrl+Z) via `SIGTSTP`:\
\
**Unix Suspend/Resume Flow**\
\
The `suspendProcess()` function [tty\_unix.go43-49](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_unix.go#L43-L49)\
 sends `SIGTSTP` to the entire process group and blocks until `SIGCONT` is received. Windows does not support suspension (`suspendSupported = false` [tty\_windows.go66](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_windows.go#L66-L66)\
).\
\
Sources: [tty.go13-23](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L13-L23)\
 [tty\_unix.go40-49](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_unix.go#L40-L49)\
 [tty\_windows.go66-68](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_windows.go#L66-L68)\
\
Window Resize Detection\
-----------------------\
\
The `checkResize()` function [tty.go121-141](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L121-L141)\
 queries the terminal size and dispatches a `WindowSizeMsg`:\
\
This is called:\
\
*   On startup (if not using alternate screen)\
*   When `SIGWINCH` is received (Unix systems)\
*   Periodically in some terminal environments\
\
The resize message flows through the normal message channel and triggers a re-render with updated dimensions.\
\
Sources: [tty.go121-141](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L121-L141)\
\
Error Handling in Input Processing\
----------------------------------\
\
Input errors are communicated through the `p.errs` channel:\
\
**Input Error Handling Flow**\
\
The read loop distinguishes between expected errors (`io.EOF`, `cancelreader.ErrCanceled`) and unexpected errors. Expected errors during shutdown are silently ignored [tty.go100-105](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L100-L105)\
\
Sources: [tty.go96-106](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L96-L106)\
\
Platform-Specific Input Differences\
-----------------------------------\
\
### Unix Systems\
\
| Feature | Implementation |\
| --- | --- |\
| Raw mode | `term.MakeRaw()` on `/dev/tty` or stdin |\
| Suspension | `SIGTSTP` / `SIGCONT` signals supported |\
| Input source | `/dev/tty` opened directly via `openInputTTY()` |\
| Terminal type | Respects `$TERM` environment variable |\
\
Sources: [tty\_unix.go15-49](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_unix.go#L15-L49)\
\
### Windows\
\
| Feature | Implementation |\
| --- | --- |\
| Raw mode | `term.MakeRaw()` with VT mode flags |\
| VT processing | Requires `ENABLE_VIRTUAL_TERMINAL_INPUT` flag |\
| Input source | `CONIN$` pseudo-file |\
| Suspension | Not supported (`suspendSupported = false`) |\
| Console modes | Must enable VT output processing separately |\
\
Windows requires explicit VT mode enablement because older Windows versions used proprietary console APIs. The VT flags enable ANSI sequence processing [tty\_windows.go25-33](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_windows.go#L25-L33)\
\
Sources: [tty\_windows.go14-68](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty_windows.go#L14-L68)\
\
Testing Input Processing\
------------------------\
\
The test suite includes comprehensive input parsing tests:\
\
### Sequence Detection Tests\
\
The `buildBaseSeqTests()` function [key\_test.go69-111](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_test.go#L69-L111)\
 generates test cases covering:\
\
*   All sequences in the `sequences` map\
*   Alt-modified versions of all sequences\
*   All control characters (0x01-0x7F, excluding 0x1B)\
*   Unknown CSI sequences\
*   Space and Alt+Space\
\
Tests verify both `detectSequence()` and the full `detectOneMsg()` pipeline [key\_test.go113-225](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_test.go#L113-L225)\
\
### Randomized Testing\
\
The `genRandomData()` function [key\_test.go607-676](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_test.go#L607-L676)\
 generates randomized sequences of control characters and escape sequences to test parser robustness. This ensures correct parsing of sequences regardless of concatenation order.\
\
Sources: [key\_test.go69-225](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_test.go#L69-L225)\
 [key\_test.go607-717](https://github.com/charmbracelet/bubbletea/blob/f25595a8/key_test.go#L607-L717)\
\
Dismiss\
\
Refresh this wiki\
\
Enter email to refresh\
\
### On this page\
\
*   [Input Handling](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#input-handling)\
    \
*   [Architecture Overview](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#architecture-overview)\
    \
*   [Message Types](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#message-types)\
    \
*   [Input Reading Architecture](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#input-reading-architecture)\
    \
*   [TTY Initialization and Raw Mode](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#tty-initialization-and-raw-mode)\
    \
*   [Asynchronous Read Loop](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#asynchronous-read-loop)\
    \
*   [Message Detection and Parsing](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#message-detection-and-parsing)\
    \
*   [Detection Precedence](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#detection-precedence)\
    \
*   [Input Buffer Management](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#input-buffer-management)\
    \
*   [Key Message Structure](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#key-message-structure)\
    \
*   [KeyType Constants](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#keytype-constants)\
    \
*   [Control Characters](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#control-characters)\
    \
*   [Navigation and Special Keys](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#navigation-and-special-keys)\
    \
*   [Function Keys](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#function-keys)\
    \
*   [Key String Representation](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#key-string-representation)\
    \
*   [ANSI Escape Sequence Parsing](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#ansi-escape-sequence-parsing)\
    \
*   [Sequence Database](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#sequence-database)\
    \
*   [Sequence Detection Algorithm](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#sequence-detection-algorithm)\
    \
*   [Special Sequence Handling](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#special-sequence-handling)\
    \
*   [Bracketed Paste](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#bracketed-paste)\
    \
*   [Focus Events](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#focus-events)\
    \
*   [Rune Handling and UTF-8](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#rune-handling-and-utf-8)\
    \
*   [Multi-Rune Keys](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#multi-rune-keys)\
    \
*   [Invalid UTF-8 Handling](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#invalid-utf-8-handling)\
    \
*   [Terminal Restoration and Suspension](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#terminal-restoration-and-suspension)\
    \
*   [State Restoration](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#state-restoration)\
    \
*   [Suspend/Resume on Unix](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#suspendresume-on-unix)\
    \
*   [Window Resize Detection](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#window-resize-detection)\
    \
*   [Error Handling in Input Processing](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#error-handling-in-input-processing)\
    \
*   [Platform-Specific Input Differences](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#platform-specific-input-differences)\
    \
*   [Unix Systems](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#unix-systems)\
    \
*   [Windows](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#windows)\
    \
*   [Testing Input Processing](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#testing-input-processing)\
    \
*   [Sequence Detection Tests](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#sequence-detection-tests)\
    \
*   [Randomized Testing](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling#randomized-testing)
