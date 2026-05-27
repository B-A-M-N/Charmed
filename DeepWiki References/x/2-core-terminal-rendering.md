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

Core Terminal Rendering
=======================

Relevant source files

*   [cellbuf/buffer.go](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/buffer.go)
    
*   [cellbuf/cell.go](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/cell.go)
    
*   [cellbuf/geom.go](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/geom.go)
    
*   [cellbuf/go.mod](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/go.mod)
    
*   [cellbuf/go.sum](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/go.sum)
    
*   [cellbuf/link.go](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/link.go)
    
*   [cellbuf/screen.go](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go)
    
*   [cellbuf/style.go](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/style.go)
    
*   [cellbuf/utils.go](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/utils.go)
    
*   [vt/buffer.go](https://github.com/charmbracelet/x/blob/7642919e/vt/buffer.go)
    
*   [vt/callbacks.go](https://github.com/charmbracelet/x/blob/7642919e/vt/callbacks.go)
    
*   [vt/cc.go](https://github.com/charmbracelet/x/blob/7642919e/vt/cc.go)
    
*   [vt/csi.go](https://github.com/charmbracelet/x/blob/7642919e/vt/csi.go)
    
*   [vt/csi\_cursor.go](https://github.com/charmbracelet/x/blob/7642919e/vt/csi_cursor.go)
    
*   [vt/csi\_mode.go](https://github.com/charmbracelet/x/blob/7642919e/vt/csi_mode.go)
    
*   [vt/csi\_screen.go](https://github.com/charmbracelet/x/blob/7642919e/vt/csi_screen.go)
    
*   [vt/dcs.go](https://github.com/charmbracelet/x/blob/7642919e/vt/dcs.go)
    
*   [vt/esc.go](https://github.com/charmbracelet/x/blob/7642919e/vt/esc.go)
    
*   [vt/handlers.go](https://github.com/charmbracelet/x/blob/7642919e/vt/handlers.go)
    
*   [vt/mode.go](https://github.com/charmbracelet/x/blob/7642919e/vt/mode.go)
    
*   [vt/osc.go](https://github.com/charmbracelet/x/blob/7642919e/vt/osc.go)
    
*   [vt/screen.go](https://github.com/charmbracelet/x/blob/7642919e/vt/screen.go)
    
*   [vt/terminal.go](https://github.com/charmbracelet/x/blob/7642919e/vt/terminal.go)
    
*   [vt/utf8.go](https://github.com/charmbracelet/x/blob/7642919e/vt/utf8.go)
    

Purpose and Scope
-----------------

The Core Terminal Rendering system provides two complementary approaches for terminal display management. The **cellbuf** package implements efficient screen buffer management with double-buffering and diff-based rendering, optimized for applications that control the entire screen state. The **vt** package implements a complete virtual terminal emulator that interprets ANSI escape sequences and maintains terminal state, suitable for applications that need to emulate a terminal session.

For detailed documentation of the cellbuf package, see [cellbuf Package](https://deepwiki.com/charmbracelet/x/2.1-cellbuf-package)
. For the vt package, see [Virtual Terminal Emulator (vt)](https://deepwiki.com/charmbracelet/x/2.2-virtual-terminal-emulator-(vt))
. For ANSI sequence parsing utilities used by both systems, see [ANSI Processing and Control](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control)
.

* * *

System Architecture
-------------------

The two rendering systems serve different use cases and can be used independently or together:

**Sources:** [cellbuf/screen.go1-611](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L1-L611)
 [vt/terminal.go1-57](https://github.com/charmbracelet/x/blob/7642919e/vt/terminal.go#L1-L57)
 [vt/screen.go1-330](https://github.com/charmbracelet/x/blob/7642919e/vt/screen.go#L1-L330)

* * *

cellbuf: Screen Buffer Management
---------------------------------

### Core Components

The cellbuf package centers on three primary structures:

| Component | Purpose | Key File |
| --- | --- | --- |
| `Screen` | Double-buffered screen renderer with diff algorithm | [cellbuf/screen.go363-386](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L363-L386) |
| `Buffer` | 2D grid of cells representing screen content | [cellbuf/buffer.go188-192](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/buffer.go#L188-L192) |
| `Cell` | Individual terminal cell with content, style, and link | [cellbuf/cell.go16-35](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/cell.go#L16-L35) |

**Sources:** [cellbuf/screen.go363-386](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L363-L386)
 [cellbuf/buffer.go188-192](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/buffer.go#L188-L192)

### Double-Buffering and Diff Algorithm

The `Screen` maintains two buffers and compares them to generate minimal output:

**Sources:** [cellbuf/screen.go1240-1265](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L1240-L1265)
 [cellbuf/screen.go924-1128](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L924-L1128)

### Cell Structure and Styling

Each `Cell` represents a single terminal character position:

**Sources:** [cellbuf/cell.go16-35](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/cell.go#L16-L35)
 [cellbuf/cell.go175-182](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/cell.go#L175-L182)
 [cellbuf/cell.go111-136](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/cell.go#L111-L136)

### Rendering Pipeline

The rendering process from application calls to terminal output:

**Sources:** [cellbuf/screen.go1261-1391](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L1261-L1391)
 [cellbuf/screen.go924-1128](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L924-L1128)
 [cellbuf/screen.go188-244](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L188-L244)

### Cursor Movement Optimization

The `Screen` optimizes cursor movement by trying multiple strategies:

| Strategy | Sequence | Use Case |
| --- | --- | --- |
| Absolute positioning | `CUP` (Cursor Position) | Long distance moves |
| Relative movement | `CUU`, `CUD`, `CUF`, `CUB` | Local movements |
| Carriage return + relative | `\r` + relative sequences | Beginning of line target |
| Hard tabs | `\t` or `CHT` | Horizontal alignment |
| Backspaces | `\b` | Short backward moves |
| Overwrite | Print actual characters | Cheaper than escape codes |

The algorithm at [cellbuf/screen.go188-244](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L188-L244)
 evaluates all enabled optimization methods and selects the shortest sequence:

**Sources:** [cellbuf/screen.go188-244](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L188-L244)
 [cellbuf/screen.go24-31](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L24-L31)
 [cellbuf/screen.go38-186](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L38-L186)

### Terminal Capabilities Detection

The `Screen` adapts to terminal capabilities using the `capabilities` bitmask at [cellbuf/screen.go505-536](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L505-L536)
:

| Capability | Sequence | Terminals Supporting |
| --- | --- | --- |
| `capVPA` | Vertical Position Absolute | Most modern terminals |
| `capHPA` | Horizontal Position Absolute | Most modern terminals |
| `capCHT` | Cursor Horizontal Tab | xterm-like (not Alacritty) |
| `capCBT` | Cursor Backward Tab | xterm-like, not VT100 |
| `capREP` | Repeat Previous Character | xterm-like, not screen |
| `capECH` | Erase Character | Most terminals |
| `capICH` | Insert Character | Most terminals |
| `capSD`/`capSU` | Scroll Down/Up | Most terminals |

The `xtermCaps()` function at [cellbuf/screen.go545-576](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L545-L576)
 determines capabilities based on the `$TERM` environment variable.

**Sources:** [cellbuf/screen.go505-576](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L505-L576)
 [cellbuf/screen.go578-611](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L578-L611)

* * *

vt: Virtual Terminal Emulation
------------------------------

### Terminal Interface

The `vt` package defines a comprehensive `Terminal` interface representing a full virtual terminal:

**Sources:** [vt/terminal.go10-56](https://github.com/charmbracelet/x/blob/7642919e/vt/terminal.go#L10-L56)

### Emulator Architecture

The virtual terminal emulator processes input through a state machine:

**Sources:** [vt/screen.go8-18](https://github.com/charmbracelet/x/blob/7642919e/vt/screen.go#L8-L18)
 [vt/handlers.go34-44](https://github.com/charmbracelet/x/blob/7642919e/vt/handlers.go#L34-L44)

### ANSI Sequence Processing

The emulator dispatches ANSI sequences through registered handlers:

**Sources:** [vt/utf8.go10-21](https://github.com/charmbracelet/x/blob/7642919e/vt/utf8.go#L10-L21)
 [vt/cc.go8-14](https://github.com/charmbracelet/x/blob/7642919e/vt/cc.go#L8-L14)
 [vt/esc.go8-21](https://github.com/charmbracelet/x/blob/7642919e/vt/esc.go#L8-L21)
 [vt/csi.go11-16](https://github.com/charmbracelet/x/blob/7642919e/vt/csi.go#L11-L16)
 [vt/osc.go14-19](https://github.com/charmbracelet/x/blob/7642919e/vt/osc.go#L14-L19)
 [vt/dcs.go5-11](https://github.com/charmbracelet/x/blob/7642919e/vt/dcs.go#L5-L11)

### Handler System

The handler system allows customization of escape sequence behavior:

**Sources:** [vt/handlers.go34-211](https://github.com/charmbracelet/x/blob/7642919e/vt/handlers.go#L34-L211)

### Default Handler Registration

Default handlers are registered for standard terminal operations:

| Handler Type | Examples | Registration Location |
| --- | --- | --- |
| Control Characters | `BEL`, `BS`, `CR`, `LF`, `HT` | [vt/handlers.go222-261](https://github.com/charmbracelet/x/blob/7642919e/vt/handlers.go#L222-L261) |
| CSI Sequences | `CUU`, `CUD`, `CUP`, `ED`, `IL`, `DL` | [vt/handlers.go459-954](https://github.com/charmbracelet/x/blob/7642919e/vt/handlers.go#L459-L954) |
| ESC Sequences | `DECSC`, `DECRC`, `IND`, `RI` | [vt/handlers.go346-456](https://github.com/charmbracelet/x/blob/7642919e/vt/handlers.go#L346-L456) |
| OSC Sequences | Window title, colors, hyperlinks | [vt/handlers.go305-343](https://github.com/charmbracelet/x/blob/7642919e/vt/handlers.go#L305-L343) |

**Sources:** [vt/handlers.go214-954](https://github.com/charmbracelet/x/blob/7642919e/vt/handlers.go#L214-L954)

### Screen State Management

The `Screen` structure in vt maintains terminal state:

**Sources:** [vt/screen.go8-330](https://github.com/charmbracelet/x/blob/7642919e/vt/screen.go#L8-L330)

### Grapheme and UTF-8 Handling

The emulator handles complex Unicode grapheme clusters:

**Sources:** [vt/utf8.go10-99](https://github.com/charmbracelet/x/blob/7642919e/vt/utf8.go#L10-L99)

* * *

Comparison: cellbuf vs vt
-------------------------

The two systems address different use cases:

| Aspect | cellbuf | vt  |
| --- | --- | --- |
| **Primary Use Case** | Direct screen control for TUI applications | Terminal emulation for shells/programs |
| **Input Model** | Application directly sets cells | Processes ANSI escape sequences |
| **State Management** | Application manages all state | Maintains terminal state (modes, cursor, etc.) |
| **Rendering Strategy** | Diff-based, double-buffered | State machine with callbacks |
| **Optimization** | Minimal ANSI output via diffing | Complete terminal emulation |
| **Typical Application** | Bubble Tea programs, custom TUIs | Terminal multiplexers, SSH clients |
| **Entry Points** | `Screen.SetCell()`, `Screen.Fill()` | `Terminal.Write()`, `Terminal.SendKey()` |
| **Dependencies** | Lightweight (ansi, term packages) | Depends on ultraviolet for buffer |

### Architecture Patterns

**cellbuf Pattern:**

    Application → SetCell() → newbuf → Render() → Diff → ANSI sequences → Terminal
    

**vt Pattern:**

    Raw bytes → Parser → Handler dispatch → Screen operations → Callbacks → Application
    

**Sources:** [cellbuf/screen.go363-386](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L363-L386)
 [vt/terminal.go10-56](https://github.com/charmbracelet/x/blob/7642919e/vt/terminal.go#L10-L56)
 [vt/screen.go8-18](https://github.com/charmbracelet/x/blob/7642919e/vt/screen.go#L8-L18)

* * *

Implementation Notes
--------------------

### Wide Character Handling

Both systems handle wide characters (CJK, emoji) by storing empty cells as placeholders:

*   In cellbuf: Wide cell at (x,y) has `Width > 1`, cells at (x+1,y)...(x+Width-1,y) are empty (`Width=0`)
*   In vt: Same approach via ultraviolet buffer

**Sources:** [cellbuf/buffer.go130-186](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/buffer.go#L130-L186)
 [cellbuf/cell.go74-78](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/cell.go#L74-L78)

### Phantom Cell State

Both systems track "phantom cursor" state for auto-wrap at screen edge:

*   When cursor reaches rightmost column, it enters phantom state
*   Next printable character triggers wrap to next line
*   Controlled by Auto Wrap Mode (`?7`)

**Sources:** [cellbuf/screen.go280-285](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L280-L285)
 [vt/utf8.go54-62](https://github.com/charmbracelet/x/blob/7642919e/vt/utf8.go#L54-L62)
 [vt/utf8.go91-94](https://github.com/charmbracelet/x/blob/7642919e/vt/utf8.go#L91-L94)

### Alternate Screen

Both support alternate screen buffer:

*   cellbuf: Via `EnterAltScreen()` / `ExitAltScreen()` methods
*   vt: Via `scrs[0]` (primary) and `scrs[1]` (alternate) array

**Sources:** [cellbuf/screen.go413-425](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L413-L425)
 [vt/csi\_mode.go38-58](https://github.com/charmbracelet/x/blob/7642919e/vt/csi_mode.go#L38-L58)

### Color Profile Support

The cellbuf `Screen` can downsample colors based on terminal capabilities:

*   `ConvertStyle()` and `ConvertLink()` convert to target color profile
*   Applied during `updatePen()` at [cellbuf/screen.go714-737](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L714-L737)
    
*   Supports TrueColor, ANSI256, ANSI16, ASCII profiles

**Sources:** [cellbuf/screen.go714-737](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L714-L737)
 [cellbuf/style.go7-31](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/style.go#L7-L31)
 [cellbuf/link.go7-14](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/link.go#L7-L14)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core Terminal Rendering](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#core-terminal-rendering)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#purpose-and-scope)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#system-architecture)
    
*   [cellbuf: Screen Buffer Management](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#cellbuf-screen-buffer-management)
    
*   [Core Components](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#core-components)
    
*   [Double-Buffering and Diff Algorithm](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#double-buffering-and-diff-algorithm)
    
*   [Cell Structure and Styling](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#cell-structure-and-styling)
    
*   [Rendering Pipeline](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#rendering-pipeline)
    
*   [Cursor Movement Optimization](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#cursor-movement-optimization)
    
*   [Terminal Capabilities Detection](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#terminal-capabilities-detection)
    
*   [vt: Virtual Terminal Emulation](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#vt-virtual-terminal-emulation)
    
*   [Terminal Interface](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#terminal-interface)
    
*   [Emulator Architecture](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#emulator-architecture)
    
*   [ANSI Sequence Processing](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#ansi-sequence-processing)
    
*   [Handler System](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#handler-system)
    
*   [Default Handler Registration](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#default-handler-registration)
    
*   [Screen State Management](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#screen-state-management)
    
*   [Grapheme and UTF-8 Handling](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#grapheme-and-utf-8-handling)
    
*   [Comparison: cellbuf vs vt](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#comparison-cellbuf-vs-vt)
    
*   [Architecture Patterns](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#architecture-patterns)
    
*   [Implementation Notes](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#implementation-notes)
    
*   [Wide Character Handling](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#wide-character-handling)
    
*   [Phantom Cell State](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#phantom-cell-state)
    
*   [Alternate Screen](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#alternate-screen)
    
*   [Color Profile Support](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering#color-profile-support)
