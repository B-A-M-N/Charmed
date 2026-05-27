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

Input Processing System
=======================

Relevant source files

*   [.github/workflows/generate.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/generate.yml)
    
*   [.github/workflows/windows-generate.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/windows-generate.yml)
    
*   [input/go.mod](https://github.com/charmbracelet/x/blob/7642919e/input/go.mod)
    
*   [input/go.sum](https://github.com/charmbracelet/x/blob/7642919e/input/go.sum)
    
*   [term/go.mod](https://github.com/charmbracelet/x/blob/7642919e/term/go.mod)
    
*   [term/go.sum](https://github.com/charmbracelet/x/blob/7642919e/term/go.sum)
    
*   [windows/doc.go](https://github.com/charmbracelet/x/blob/7642919e/windows/doc.go)
    
*   [windows/go.mod](https://github.com/charmbracelet/x/blob/7642919e/windows/go.mod)
    
*   [windows/go.sum](https://github.com/charmbracelet/x/blob/7642919e/windows/go.sum)
    
*   [windows/syscall\_windows.go](https://github.com/charmbracelet/x/blob/7642919e/windows/syscall_windows.go)
    
*   [windows/types.go](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go)
    
*   [windows/types\_windows.go](https://github.com/charmbracelet/x/blob/7642919e/windows/types_windows.go)
    
*   [windows/zsyscall\_windows.go](https://github.com/charmbracelet/x/blob/7642919e/windows/zsyscall_windows.go)
    
*   [xpty/go.mod](https://github.com/charmbracelet/x/blob/7642919e/xpty/go.mod)
    
*   [xpty/go.sum](https://github.com/charmbracelet/x/blob/7642919e/xpty/go.sum)
    

Purpose and Scope
-----------------

The Input Processing System provides cross-platform terminal input handling for interactive applications in the charmbracelet/x repository. This system abstracts platform differences between Unix/Linux and Windows, converting raw terminal input (keyboard, mouse, window resize events) into structured, application-friendly event types.

This document covers the high-level architecture, platform abstraction strategy, and event flow through the input processing pipeline. For detailed documentation of the event-based API and ANSI sequence parsing, see [Input Package](https://deepwiki.com/charmbracelet/x/3.1-input-package)
. For Windows-specific console input handling details, see [Windows Console Input](https://deepwiki.com/charmbracelet/x/3.2-windows-console-input)
.

The input system integrates with the [Terminal I/O](https://deepwiki.com/charmbracelet/x/5.3-terminal-io-(termios-and-wcwidth))
 layer for raw mode control and the [ANSI Processing](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control)
 system for sequence parsing.

**Sources:** [input/go.mod1-22](https://github.com/charmbracelet/x/blob/7642919e/input/go.mod#L1-L22)
 [windows/go.mod1-6](https://github.com/charmbracelet/x/blob/7642919e/windows/go.mod#L1-L6)

System Overview
---------------

The Input Processing System serves as the bridge between raw operating system input mechanisms and application-level event handling. It performs three critical functions:

1.  **Platform Abstraction**: Provides a unified API that works identically on Unix/Linux (ANSI escape sequences) and Windows (Console API InputRecords)
2.  **Event Parsing**: Converts raw byte streams or OS-specific structures into typed event objects (key presses, mouse movements, window resizes)
3.  **Non-blocking I/O**: Supports cancellable read operations to enable graceful shutdown and timeout handling

The system is designed as a library component that applications use to receive input events through a simple `Read()` or `ReadEvents()` interface, hiding all platform-specific complexity.

**Sources:** [input/go.mod1-22](https://github.com/charmbracelet/x/blob/7642919e/input/go.mod#L1-L22)
 Diagram 2 and Diagram 6 from high-level architecture

Architecture
------------

### Component Structure

**Diagram: Input Processing System Component Architecture**

The architecture uses a driver pattern where `input.Reader` selects the appropriate platform implementation at runtime based on build tags (`//go:build !windows` vs `//go:build windows`).

**Sources:** [input/go.mod5-11](https://github.com/charmbracelet/x/blob/7642919e/input/go.mod#L5-L11)
 Diagram 4 from high-level architecture

### Key Dependencies

| Dependency | Version | Purpose |
| --- | --- | --- |
| `github.com/charmbracelet/x/ansi` | v0.11.4 | ANSI escape sequence parsing (Unix input) |
| `github.com/charmbracelet/x/windows` | v0.2.2 | Windows Console API wrappers |
| `github.com/muesli/cancelreader` | v0.2.2 | Cancellable I/O operations |
| `github.com/rivo/uniseg` | v0.4.7 | Unicode grapheme cluster segmentation |
| `github.com/xo/terminfo` | v0.0.0-20220910002029 | Terminfo database access |
| `golang.org/x/sys` | v0.40.0 | Low-level system calls (both platforms) |

**Sources:** [input/go.mod5-11](https://github.com/charmbracelet/x/blob/7642919e/input/go.mod#L5-L11)
 [input/go.sum1-24](https://github.com/charmbracelet/x/blob/7642919e/input/go.sum#L1-L24)

Platform-Specific Implementations
---------------------------------

### Unix/Linux Approach

On Unix-like systems, the input system reads raw bytes from stdin (or `/dev/tty` for direct terminal access) and parses ANSI escape sequences using a state machine:

**Diagram: Unix Input Processing Flow**

The Unix driver relies on the terminal being in raw mode (see [Terminal I/O](https://deepwiki.com/charmbracelet/x/5.3-terminal-io-(termios-and-wcwidth))
) to receive unprocessed input. It uses `cancelreader` to enable graceful cancellation of blocking reads.

**Sources:** [input/go.mod6-10](https://github.com/charmbracelet/x/blob/7642919e/input/go.mod#L6-L10)
 Diagram 2 and Diagram 4 from high-level architecture

### Windows Approach

On Windows, the input system uses the Console API's `ReadConsoleInput` function to receive structured `INPUT_RECORD` events:

**Diagram: Windows Input Processing Flow**

The Windows driver parses five event types defined in `INPUT_RECORD`: `KEY_EVENT`, `MOUSE_EVENT`, `WINDOW_BUFFER_SIZE_EVENT`, `FOCUS_EVENT`, and `MENU_EVENT`.

**Sources:** [windows/types\_windows.go10-97](https://github.com/charmbracelet/x/blob/7642919e/windows/types_windows.go#L10-L97)
 [windows/types.go211-266](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go#L211-L266)
 [windows/syscall\_windows.go11-14](https://github.com/charmbracelet/x/blob/7642919e/windows/syscall_windows.go#L11-L14)

### Build Tag Strategy

The platform-specific implementations are selected at compile time using Go build tags:

| Build Tag | Platform | Driver Implementation |
| --- | --- | --- |
| `!windows` | Unix/Linux/macOS | ANSI sequence parsing from stdin |
| `windows` | Windows | Console API `ReadConsoleInput` |

This ensures that only the relevant platform code is compiled, reducing binary size and eliminating unnecessary dependencies.

**Sources:** Diagram 4 from high-level architecture

Event Flow
----------

The following diagram illustrates the complete data flow from user input to application-level events:

**Diagram: End-to-End Input Event Flow**

**Sources:** Diagram 6 from high-level architecture

### Event Types

The system produces several event types that applications consume:

| Event Type | Description | Unix Source | Windows Source |
| --- | --- | --- | --- |
| `KeyPressEvent` | Keyboard key press/release | ANSI escape sequences | `KEY_EVENT` records |
| `MouseEvent` | Mouse movement/clicks | ANSI mouse sequences (CSI M) | `MOUSE_EVENT` records |
| `WindowSizeEvent` | Terminal resize | `SIGWINCH` signal | `WINDOW_BUFFER_SIZE_EVENT` |
| `FocusEvent` | Terminal focus change | ANSI focus sequences | `FOCUS_EVENT` records |
| `PasteEvent` | Bracketed paste | ANSI paste sequences | N/A (not supported) |

**Sources:** [windows/types.go211-221](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go#L211-L221)
 Diagram 6 from high-level architecture

Windows Event Structures
------------------------

The Windows implementation defines several structures that correspond to the Windows Console API:

### InputRecord Structure

[windows/types.go253-266](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go#L253-L266)
 defines `InputRecord`, which contains:

*   `EventType` (uint16): One of `KEY_EVENT`, `MOUSE_EVENT`, `WINDOW_BUFFER_SIZE_EVENT`, `FOCUS_EVENT`, or `MENU_EVENT`
*   `Event` (\[16\]byte): Raw event data, interpreted based on `EventType`

### KeyEventRecord Structure

[windows/types.go223-251](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go#L223-L251)
 defines `KeyEventRecord`:

*   `KeyDown` (bool): Key pressed or released
*   `RepeatCount` (uint16): Repeat count for held keys
*   `VirtualKeyCode` (uint16): Virtual key code (VK\_\* constants)
*   `VirtualScanCode` (uint16): Hardware scan code
*   `Char` (rune): Unicode character
*   `ControlKeyState` (uint32): Modifier keys (Shift, Ctrl, Alt)

### MouseEventRecord Structure

[windows/types\_windows.go25-41](https://github.com/charmbracelet/x/blob/7642919e/windows/types_windows.go#L25-L41)
 defines `MouseEventRecord`:

*   `MousePosition` (Coord): Cursor position in character cells
*   `ButtonState` (uint32): Which buttons are pressed
*   `ControlKeyState` (uint32): Modifier keys
*   `EventFlags` (uint32): Event type (moved, clicked, wheeled)

### WindowBufferSizeRecord Structure

[windows/types\_windows.go43-50](https://github.com/charmbracelet/x/blob/7642919e/windows/types_windows.go#L43-L50)
 defines `WindowBufferSizeRecord`:

*   `Size` (Coord): Console buffer dimensions in character cells

**Sources:** [windows/types.go223-266](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go#L223-L266)
 [windows/types\_windows.go10-97](https://github.com/charmbracelet/x/blob/7642919e/windows/types_windows.go#L10-L97)

System Calls and APIs
---------------------

### Unix System Calls

On Unix systems, the input driver uses standard POSIX I/O:

*   `read()` on stdin file descriptor for raw bytes
*   `ioctl()` with `TIOCGWINSZ` for terminal size (via `term` package)
*   Signal handling for `SIGWINCH` (window resize notifications)

### Windows System Calls

[windows/syscall\_windows.go11-14](https://github.com/charmbracelet/x/blob/7642919e/windows/syscall_windows.go#L11-L14)
 defines Windows Console API functions:

| Function | Purpose |
| --- | --- |
| `ReadConsoleInput` | Reads input events from console buffer |
| `PeekConsoleInput` | Checks for pending input without removing |
| `GetNumberOfConsoleInputEvents` | Returns count of pending events |
| `FlushConsoleInputBuffer` | Clears input buffer |

These are implemented via syscalls to `kernel32.dll` in [windows/zsyscall\_windows.go38-77](https://github.com/charmbracelet/x/blob/7642919e/windows/zsyscall_windows.go#L38-L77)

**Sources:** [windows/syscall\_windows.go11-14](https://github.com/charmbracelet/x/blob/7642919e/windows/syscall_windows.go#L11-L14)
 [windows/zsyscall\_windows.go38-77](https://github.com/charmbracelet/x/blob/7642919e/windows/zsyscall_windows.go#L38-L77)

Integration Points
------------------

The Input Processing System integrates with other systems in the charmbracelet/x ecosystem:

**Diagram: Input System Integration with Other Packages**

*   **[Term Package](https://deepwiki.com/charmbracelet/x/5.3-terminal-io-(termios-and-wcwidth))
    **: Provides raw mode control via `MakeRaw()` and terminal dimension queries
*   **[ANSI Package](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control)
    **: Parses escape sequences into structured data on Unix platforms
*   **[Windows Package](https://deepwiki.com/charmbracelet/x/3.2-windows-console-input)
    **: Wraps Console API for Windows input handling
*   **[Cellbuf Package](https://deepwiki.com/charmbracelet/x/2.1-cellbuf-package)
    **: Uses input events in its event loop for screen updates
*   **[Teatest Package](https://deepwiki.com/charmbracelet/x/7.2-testing-framework-(teatest))
    **: Provides simulated input for testing interactive applications

**Sources:** [input/go.mod5-11](https://github.com/charmbracelet/x/blob/7642919e/input/go.mod#L5-L11)
 Diagram 1 from high-level architecture

Virtual Key Codes
-----------------

The Windows implementation includes comprehensive virtual key code definitions in [windows/types.go3-169](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go#L3-L169)
 These map Windows keyboard input to standardized codes:

| Key Category | Code Range | Examples |
| --- | --- | --- |
| Function keys | 0x70-0x87 | `VK_F1` through `VK_F24` |
| Navigation | 0x21-0x28 | `VK_HOME`, `VK_END`, `VK_LEFT`, `VK_UP` |
| Modifiers | 0xA0-0xA5 | `VK_LSHIFT`, `VK_RSHIFT`, `VK_LCONTROL` |
| Alphanumeric | 0x30-0x5A | Standard keys (handled by Char field) |
| Numpad | 0x60-0x6F | `VK_NUMPAD0` through `VK_DIVIDE` |

**Sources:** [windows/types.go3-169](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go#L3-L169)

Modifier Key Handling
---------------------

Both platforms support modifier key detection, though through different mechanisms:

### Unix Modifiers

On Unix, modifiers are encoded in ANSI escape sequences:

*   Shift: Encoded in escape sequence parameters
*   Ctrl: May produce control characters (e.g., Ctrl+C = `\x03`)
*   Alt/Meta: Adds ESC prefix to sequences

### Windows Modifiers

Windows uses control key state flags in [windows/types.go183-198](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go#L183-L198)
:

| Flag | Mask | Meaning |
| --- | --- | --- |
| `SHIFT_PRESSED` | 0x0010 | Any Shift key |
| `LEFT_CTRL_PRESSED` | 0x0008 | Left Control |
| `RIGHT_CTRL_PRESSED` | 0x0004 | Right Control |
| `LEFT_ALT_PRESSED` | 0x0002 | Left Alt |
| `RIGHT_ALT_PRESSED` | 0x0001 | Right Alt |
| `CAPSLOCK_ON` | 0x0080 | Caps Lock state |
| `NUMLOCK_ON` | 0x0020 | Num Lock state |

**Sources:** [windows/types.go183-198](https://github.com/charmbracelet/x/blob/7642919e/windows/types.go#L183-L198)

Summary
-------

The Input Processing System provides a clean, cross-platform abstraction over terminal input handling. By using platform-specific drivers selected at compile time, it delivers optimal performance on each operating system while presenting a unified API to applications. The system handles keyboard input, mouse events, window resizes, and focus changes, enabling rich interactive terminal applications.

For implementation details of the main API, see [Input Package](https://deepwiki.com/charmbracelet/x/3.1-input-package)
. For Windows-specific console handling, see [Windows Console Input](https://deepwiki.com/charmbracelet/x/3.2-windows-console-input)
.

**Sources:** All files in input/, windows/, and high-level architecture diagrams

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Input Processing System](https://deepwiki.com/charmbracelet/x/3-input-processing-system#input-processing-system)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/x/3-input-processing-system#purpose-and-scope)
    
*   [System Overview](https://deepwiki.com/charmbracelet/x/3-input-processing-system#system-overview)
    
*   [Architecture](https://deepwiki.com/charmbracelet/x/3-input-processing-system#architecture)
    
*   [Component Structure](https://deepwiki.com/charmbracelet/x/3-input-processing-system#component-structure)
    
*   [Key Dependencies](https://deepwiki.com/charmbracelet/x/3-input-processing-system#key-dependencies)
    
*   [Platform-Specific Implementations](https://deepwiki.com/charmbracelet/x/3-input-processing-system#platform-specific-implementations)
    
*   [Unix/Linux Approach](https://deepwiki.com/charmbracelet/x/3-input-processing-system#unixlinux-approach)
    
*   [Windows Approach](https://deepwiki.com/charmbracelet/x/3-input-processing-system#windows-approach)
    
*   [Build Tag Strategy](https://deepwiki.com/charmbracelet/x/3-input-processing-system#build-tag-strategy)
    
*   [Event Flow](https://deepwiki.com/charmbracelet/x/3-input-processing-system#event-flow)
    
*   [Event Types](https://deepwiki.com/charmbracelet/x/3-input-processing-system#event-types)
    
*   [Windows Event Structures](https://deepwiki.com/charmbracelet/x/3-input-processing-system#windows-event-structures)
    
*   [InputRecord Structure](https://deepwiki.com/charmbracelet/x/3-input-processing-system#inputrecord-structure)
    
*   [KeyEventRecord Structure](https://deepwiki.com/charmbracelet/x/3-input-processing-system#keyeventrecord-structure)
    
*   [MouseEventRecord Structure](https://deepwiki.com/charmbracelet/x/3-input-processing-system#mouseeventrecord-structure)
    
*   [WindowBufferSizeRecord Structure](https://deepwiki.com/charmbracelet/x/3-input-processing-system#windowbuffersizerecord-structure)
    
*   [System Calls and APIs](https://deepwiki.com/charmbracelet/x/3-input-processing-system#system-calls-and-apis)
    
*   [Unix System Calls](https://deepwiki.com/charmbracelet/x/3-input-processing-system#unix-system-calls)
    
*   [Windows System Calls](https://deepwiki.com/charmbracelet/x/3-input-processing-system#windows-system-calls)
    
*   [Integration Points](https://deepwiki.com/charmbracelet/x/3-input-processing-system#integration-points)
    
*   [Virtual Key Codes](https://deepwiki.com/charmbracelet/x/3-input-processing-system#virtual-key-codes)
    
*   [Modifier Key Handling](https://deepwiki.com/charmbracelet/x/3-input-processing-system#modifier-key-handling)
    
*   [Unix Modifiers](https://deepwiki.com/charmbracelet/x/3-input-processing-system#unix-modifiers)
    
*   [Windows Modifiers](https://deepwiki.com/charmbracelet/x/3-input-processing-system#windows-modifiers)
    
*   [Summary](https://deepwiki.com/charmbracelet/x/3-input-processing-system#summary)
