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

Rendering System
================

Relevant source files

*   [cursed\_renderer.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/cursed_renderer.go)
    
*   [nil\_renderer.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/nil_renderer.go)
    
*   [renderer.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/renderer.go)
    
*   [screen.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/screen.go)
    

The Rendering System manages all terminal output in Bubbletea applications. It provides an abstraction layer for rendering UI content to the terminal, implementing framerate-based rendering with differential updates to avoid terminal overload. The system handles alternate screen buffers, cursor visibility, mouse modes, bracketed paste, focus reporting, and window title management.

For information about the input handling side of terminal I/O, see [Input Handling](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling)
. For terminal state management including raw mode setup and suspension, see section [4.3](https://deepwiki.com/charmbracelet/bubbletea/4.3-terminal-management)
.

Architecture Overview
---------------------

The rendering system follows an interface-based design with multiple implementations. The core abstraction is the `renderer` interface, which defines 26 methods for terminal output operations. The primary implementation is `standardRenderer`, which uses framerate-based rendering with differential updates for performance optimization.

**Sources:** [renderer.go1-89](https://github.com/charmbracelet/bubbletea/blob/f25595a8/renderer.go#L1-L89)
 [standard\_renderer.go1-791](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L1-L791)
 [nil\_renderer.go1-30](https://github.com/charmbracelet/bubbletea/blob/f25595a8/nil_renderer.go#L1-L30)
 [screen.go1-249](https://github.com/charmbracelet/bubbletea/blob/f25595a8/screen.go#L1-L249)

Renderer Interface
------------------

The `renderer` interface defines the contract for all renderer implementations. It provides 26 methods organized into functional groups: lifecycle management, output operations, screen control, cursor management, mouse configuration, paste handling, focus reporting, and window manipulation.

### Interface Methods

| Category | Methods | Description |
| --- | --- | --- |
| **Lifecycle** | `start()`, `stop()`, `kill()` | Control renderer execution. `stop()` renders final frame, `kill()` does not. |
| **Output** | `write(string)`, `repaint()`, `clearScreen()` | Buffer content for rendering and trigger repaints. |
| **Screen** | `altScreen()`, `enterAltScreen()`, `exitAltScreen()` | Manage alternate screen buffer state. |
| **Cursor** | `showCursor()`, `hideCursor()` | Control cursor visibility. |
| **Mouse Cell** | `enableMouseCellMotion()`, `disableMouseCellMotion()` | Enable mouse events when button is pressed (drag). |
| **Mouse All** | `enableMouseAllMotion()`, `disableMouseAllMotion()` | Enable mouse events regardless of button state (hover). |
| **Mouse SGR** | `enableMouseSGRMode()`, `disableMouseSGRMode()` | Enable extended mouse coordinate reporting. |
| **Paste** | `enableBracketedPaste()`, `disableBracketedPaste()`, `bracketedPasteActive()` | Control bracketed paste mode. |
| **Focus** | `enableReportFocus()`, `disableReportFocus()`, `reportFocus()` | Control focus event reporting. |
| **Window** | `setWindowTitle(string)` | Set terminal window title. |
| **Internal** | `resetLinesRendered()` | Reset line count for exec output. |

**Sources:** [renderer.go4-85](https://github.com/charmbracelet/bubbletea/blob/f25595a8/renderer.go#L4-L85)

### Message-Based Control

Applications control rendering features through command functions that return messages. The `tea.Program` processes these messages and invokes the appropriate renderer methods. This maintains separation between application logic and renderer control.

**Sources:** [screen.go12-165](https://github.com/charmbracelet/bubbletea/blob/f25595a8/screen.go#L12-L165)

### nilRenderer Implementation

The `nilRenderer` provides a no-op implementation of the `renderer` interface. All methods are empty stubs that perform no operations. This implementation is used for testing or when output is explicitly disabled via the `WithoutRenderer` program option.

**Sources:** [nil\_renderer.go3-30](https://github.com/charmbracelet/bubbletea/blob/f25595a8/nil_renderer.go#L3-L30)

Standard Renderer
-----------------

The `standardRenderer` is the primary renderer implementation, providing framerate-based rendering with sophisticated optimizations to minimize terminal I/O and avoid overloading terminal emulators.

### Core Structure

**Sources:** [standard\_renderer.go27-61](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L27-L61)
 [standard\_renderer.go86-158](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L86-L158)

### Initialization and Lifecycle

The renderer is created with `newRenderer()` which accepts an `io.Writer`, ANSI compression flag, and FPS setting. The framerate is clamped between 1 and 120 FPS (default 60 FPS).

**Sources:** [standard\_renderer.go65-83](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L65-L83)
 [standard\_renderer.go86-100](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L86-L100)
 [standard\_renderer.go103-124](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L103-L124)
 [standard\_renderer.go147-158](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L147-L158)

### Framerate-Based Rendering

The renderer uses a ticker-based goroutine (`listen()`) to trigger flushes at the configured framerate. This prevents excessive terminal updates that could overload the terminal emulator.

| FPS Setting | Frame Duration | Use Case |
| --- | --- | --- |
| 1-59 | Custom interval | Low-performance terminals or slow updates |
| 60 (default) | 16.67ms | Standard interactive applications |
| 61-120 | Custom interval | High-performance animations |

**Sources:** [standard\_renderer.go15-20](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L15-L20)
 [standard\_renderer.go147-158](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L147-L158)

### Differential Rendering

The `flush()` method implements differential rendering to minimize terminal I/O. It compares the current buffer content against `lastRender` and `lastRenderedLines` to avoid redundant writes.

**Sources:** [standard\_renderer.go161-291](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L161-L291)

### Optimization Techniques

The renderer employs several optimization strategies:

1.  **Early Return**: Exits immediately if buffer is empty or identical to `lastRender` [standard\_renderer.go165-168](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L165-L168)
    
2.  **Line-by-Line Comparison**: Skips writing lines that haven't changed since last render [standard\_renderer.go213-223](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L213-L223)
    
3.  **Width Truncation**: Truncates lines exceeding terminal width to prevent wrapping [standard\_renderer.go240-242](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L240-L242)
    
4.  **Selective Erasing**: Only appends `EraseLineRight` for lines shorter than terminal width [standard\_renderer.go244-252](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L244-L252)
    
5.  **Height Limiting**: Drops lines from the top when output exceeds terminal height [standard\_renderer.go186-188](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L186-L188)
    
6.  **Cursor Optimization**: Minimizes cursor movement sequences [standard\_renderer.go174-178](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L174-L178)
     [standard\_renderer.go274-281](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L274-L281)
    

**Sources:** [standard\_renderer.go161-291](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L161-L291)

### ANSI Sequence Generation

All ANSI escape sequences are generated through the `charmbracelet/x/ansi` library. The renderer calls functions like `ansi.CursorUp()`, `ansi.CursorHomePosition`, `ansi.EraseLineRight`, `ansi.EraseScreenBelow`, `ansi.Truncate()`, and `ansi.StringWidth()` to create terminal control sequences.

| ANSI Function | Purpose | Usage |
| --- | --- | --- |
| `ansi.CursorUp(n)` | Move cursor up n lines | Repositioning to start of render area |
| `ansi.CursorHomePosition` | Move cursor to (1,1) | Altscreen mode positioning |
| `ansi.CursorPosition(x, y)` | Move cursor to specific position | Final cursor placement |
| `ansi.EraseLineRight` | Clear from cursor to end of line | Remove previous content |
| `ansi.EraseScreenBelow` | Clear from cursor to end of screen | Remove leftover lines |
| `ansi.EraseEntireScreen` | Clear entire screen | Screen clearing operations |
| `ansi.StringWidth(s)` | Calculate visible width | Width calculations |
| `ansi.Truncate(s, width, tail)` | Truncate string to width | Prevent line wrapping |

**Sources:** [standard\_renderer.go11](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L11-L11)
 [standard\_renderer.go175-178](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L175-L178)
 [standard\_renderer.go195-263](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L195-L263)

### Thread Safety

All public methods of `standardRenderer` use mutex locks to ensure thread-safe access to internal state. The mutex is acquired at the start of each method and released via `defer` to guarantee unlocking even if panics occur.

**Sources:** [standard\_renderer.go28](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L28-L28)
 [standard\_renderer.go162-163](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L162-L163)
 [standard\_renderer.go304-305](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L304-L305)
 [standard\_renderer.go325-326](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L325-L326)

### Alternate Screen Buffer

The renderer tracks alternate screen buffer state with the `altScreenActive` boolean. When entering altscreen, it:

1.  Sets `altScreenActive = true`
2.  Executes `ansi.SetAltScreenSaveCursorMode`
3.  Clears the screen
4.  Resets cursor visibility based on `cursorHidden` state
5.  Resets `altLinesRendered` counter

Exiting altscreen reverses these operations.

**Sources:** [standard\_renderer.go341-397](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L341-L397)

### Message Handling

The renderer processes internal messages through `handleMessages()`, which responds to window size updates, repaint requests, and deprecated scroll area messages.

**Sources:** [standard\_renderer.go621-672](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L621-L672)

### Println Output

The renderer supports unmanaged output via `Println()` and `Printf()` commands. These messages are queued in `queuedMessageLines` and flushed during the next render cycle when not in altscreen mode.

**Sources:** [standard\_renderer.go32](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L32-L32)
 [standard\_renderer.go190-210](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L190-L210)
 [standard\_renderer.go663-670](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L663-L670)
 [standard\_renderer.go768-790](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L768-L790)

Terminal Features
-----------------

The renderer provides methods to control various terminal features. Each feature maintains internal state and uses ANSI sequences from the `x/ansi` library.

### Feature Control Methods

**Sources:** [standard\_renderer.go399-506](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L399-L506)

### ANSI Sequence Mapping

The following table maps renderer methods to their corresponding ANSI sequences:

| Renderer Method | ANSI Sequence Function | Terminal Effect |
| --- | --- | --- |
| `enterAltScreen()` | `ansi.SetAltScreenSaveCursorMode` | Enable alternate screen buffer |
| `exitAltScreen()` | `ansi.ResetAltScreenSaveCursorMode` | Disable alternate screen buffer |
| `hideCursor()` | `ansi.HideCursor` | Hide terminal cursor |
| `showCursor()` | `ansi.ShowCursor` | Show terminal cursor |
| `enableMouseCellMotion()` | `ansi.SetButtonEventMouseMode` | Enable button event mouse tracking (ESC<FileRef file-url="[https://github.com/charmbracelet/bubbletea/blob/f25595a8/?1002h](https://github.com/charmbracelet/bubbletea/blob/f25595a8/?1002h)<br>) |

Ignored Lines and Scrolling
---------------------------

The renderer provides deprecated functionality for high-performance scroll-based rendering through ignored line regions. This allows applications to bypass the normal rendering pipeline for specific line ranges.

The `setIgnoredLines(from, to)` method marks lines to be skipped during rendering. The `clearIgnoredLines()` method returns control of these lines to the normal rendering pipeline. These methods support deprecated scroll commands like `ScrollUp()`, `ScrollDown()`, and `SyncScrollArea()`.

**Sources:** [standard\_renderer.go510-589](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L510-L589)
 [standard\_renderer.go676-755](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L676-L755)

Testing Support
---------------

The `nilRenderer` provides a no-op implementation suitable for testing. All methods are empty stubs that perform no operations and return appropriate zero values.

**Sources:** [nil\_renderer.go3-30](https://github.com/charmbracelet/bubbletea/blob/f25595a8/nil_renderer.go#L3-L30)

Integration with Program
------------------------

The renderer is managed by `tea.Program` throughout the application lifecycle. The program creates the renderer during initialization, starts it before entering the event loop, calls `write()` after each `View()` invocation, and stops it during shutdown.

The program also processes screen control messages (from `screen.go`) by invoking corresponding renderer methods, maintaining the message-based architecture.

**Sources:** [standard\_renderer.go65-124](https://github.com/charmbracelet/bubbletea/blob/f25595a8/standard_renderer.go#L65-L124)
 [screen.go1-249](https://github.com/charmbracelet/bubbletea/blob/f25595a8/screen.go#L1-L249)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Rendering System](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#rendering-system)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#architecture-overview)
    
*   [Renderer Interface](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#renderer-interface)
    
*   [Interface Methods](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#interface-methods)
    
*   [Message-Based Control](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#message-based-control)
    
*   [nilRenderer Implementation](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#nilrenderer-implementation)
    
*   [Standard Renderer](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#standard-renderer)
    
*   [Core Structure](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#core-structure)
    
*   [Initialization and Lifecycle](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#initialization-and-lifecycle)
    
*   [Framerate-Based Rendering](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#framerate-based-rendering)
    
*   [Differential Rendering](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#differential-rendering)
    
*   [Optimization Techniques](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#optimization-techniques)
    
*   [ANSI Sequence Generation](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#ansi-sequence-generation)
    
*   [Thread Safety](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#thread-safety)
    
*   [Alternate Screen Buffer](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#alternate-screen-buffer)
    
*   [Message Handling](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#message-handling)
    
*   [Println Output](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#println-output)
    
*   [Terminal Features](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#terminal-features)
    
*   [Feature Control Methods](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#feature-control-methods)
    
*   [ANSI Sequence Mapping](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#ansi-sequence-mapping)
    
*   [Ignored Lines and Scrolling](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#ignored-lines-and-scrolling)
    
*   [Testing Support](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#testing-support)
    
*   [Integration with Program](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system#integration-with-program)
