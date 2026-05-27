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

Platform Abstraction Layer
==========================

Relevant source files

*   [conpty/conpty.go](https://github.com/charmbracelet/x/blob/7642919e/conpty/conpty.go)
    
*   [conpty/conpty\_other.go](https://github.com/charmbracelet/x/blob/7642919e/conpty/conpty_other.go)
    
*   [conpty/conpty\_windows.go](https://github.com/charmbracelet/x/blob/7642919e/conpty/conpty_windows.go)
    
*   [conpty/doc.go](https://github.com/charmbracelet/x/blob/7642919e/conpty/doc.go)
    
*   [conpty/exec\_windows.go](https://github.com/charmbracelet/x/blob/7642919e/conpty/exec_windows.go)
    
*   [conpty/go.mod](https://github.com/charmbracelet/x/blob/7642919e/conpty/go.mod)
    
*   [conpty/go.sum](https://github.com/charmbracelet/x/blob/7642919e/conpty/go.sum)
    
*   [input/go.mod](https://github.com/charmbracelet/x/blob/7642919e/input/go.mod)
    
*   [input/go.sum](https://github.com/charmbracelet/x/blob/7642919e/input/go.sum)
    
*   [term/go.mod](https://github.com/charmbracelet/x/blob/7642919e/term/go.mod)
    
*   [term/go.sum](https://github.com/charmbracelet/x/blob/7642919e/term/go.sum)
    
*   [termios/go.mod](https://github.com/charmbracelet/x/blob/7642919e/termios/go.mod)
    
*   [wcwidth/go.mod](https://github.com/charmbracelet/x/blob/7642919e/wcwidth/go.mod)
    
*   [wcwidth/go.sum](https://github.com/charmbracelet/x/blob/7642919e/wcwidth/go.sum)
    
*   [wcwidth/wcwidth.go](https://github.com/charmbracelet/x/blob/7642919e/wcwidth/wcwidth.go)
    
*   [xpty/go.mod](https://github.com/charmbracelet/x/blob/7642919e/xpty/go.mod)
    
*   [xpty/go.sum](https://github.com/charmbracelet/x/blob/7642919e/xpty/go.sum)
    

This document covers the cross-platform abstractions that enable consistent terminal functionality across operating systems. The Platform Abstraction Layer provides unified interfaces for pseudo-terminal (PTY) operations, Windows Console API access, and Unix terminal control, allowing terminal applications to work seamlessly on both Windows and Unix-like systems.

For specific input processing details, see [Cross-Platform Input Handling](https://deepwiki.com/charmbracelet/x/3.1-input-package)
. For terminal control sequences and ANSI processing, see [ANSI Processing and Control](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control)
.

Unified PTY Interface
---------------------

The core of the platform abstraction is the `Pty` interface defined in [xpty/xpty.go21-41](https://github.com/charmbracelet/x/blob/7642919e/xpty/xpty.go#L21-L41)
 which provides a common API for pseudo-terminal operations across platforms.

**Factory Pattern and Runtime Detection**

Sources: [xpty/xpty.go69-74](https://github.com/charmbracelet/x/blob/7642919e/xpty/xpty.go#L69-L74)

The `NewPty` function uses runtime OS detection to instantiate the appropriate implementation:

| OS Detection | Implementation | Factory Function |
| --- | --- | --- |
| `runtime.GOOS == "windows"` | `ConPty` | `NewConPty()` |
| All other platforms | `UnixPty` | `NewUnixPty()` |

The `Pty` interface defines these core methods:

*   `Read([]byte) (int, error)` - Read from PTY
*   `Write([]byte) (int, error)` - Write to PTY
*   `Close() error` - Close PTY
*   `Resize(width, height int) error` - Resize PTY
*   `Size() (width, height int, err error)` - Get PTY size
*   `Name() string` - Get PTY name
*   `Start(*exec.Cmd) error` - Start command on PTY

Sources: [xpty/xpty.go21-41](https://github.com/charmbracelet/x/blob/7642919e/xpty/xpty.go#L21-L41)

Unix PTY Implementation
-----------------------

The `UnixPty` struct implements the `Pty` interface using traditional Unix master/slave pseudo-terminal pairs.

**Platform-Specific Build Tags**

The Unix PTY implementation uses build constraints to ensure compatibility:

*   **Unix platforms**: [xpty/pty\_unix.go1-2](https://github.com/charmbracelet/x/blob/7642919e/xpty/pty_unix.go#L1-L2)
     - Supports Darwin, Linux, FreeBSD, etc.
*   **Unsupported platforms**: [xpty/pty\_other.go1-2](https://github.com/charmbracelet/x/blob/7642919e/xpty/pty_other.go#L1-L2)
     - Returns `ErrUnsupported`

**Core Unix PTY Operations**

| Method | Implementation | Platform APIs |
| --- | --- | --- |
| `NewUnixPty()` | [xpty/pty.go18-36](https://github.com/charmbracelet/x/blob/7642919e/xpty/pty.go#L18-L36) | `pty.Open()` |
| `Resize()` | [xpty/pty\_unix.go12-25](https://github.com/charmbracelet/x/blob/7642919e/xpty/pty_unix.go#L12-L25) | `termios.SetWinsize()` |
| `Size()` | [xpty/pty\_unix.go28-43](https://github.com/charmbracelet/x/blob/7642919e/xpty/pty_unix.go#L28-L43) | `termios.GetWinsize()` |
| `Start()` | [xpty/pty.go90-104](https://github.com/charmbracelet/x/blob/7642919e/xpty/pty.go#L90-L104) | Connects cmd stdio to slave |

Sources: [xpty/pty.go10-130](https://github.com/charmbracelet/x/blob/7642919e/xpty/pty.go#L10-L130)
 [xpty/pty\_unix.go1-44](https://github.com/charmbracelet/x/blob/7642919e/xpty/pty_unix.go#L1-L44)
 [xpty/pty\_other.go1-13](https://github.com/charmbracelet/x/blob/7642919e/xpty/pty_other.go#L1-L13)

Windows ConPTY Implementation
-----------------------------

The `ConPty` struct wraps the Windows Console Pseudo Console (ConPTY) API through the `conpty` package.

**Windows-Specific Process Handling**

The Windows implementation includes special process handling due to ConPTY limitations:

*   **Process Spawning**: [xpty/conpty\_windows.go15-37](https://github.com/charmbracelet/x/blob/7642919e/xpty/conpty_windows.go#L15-L37)
     - Uses `conpty.Spawn()` instead of standard `exec.Cmd.Start()`
*   **Process Waiting**: [xpty/xpty.go76-108](https://github.com/charmbracelet/x/blob/7642919e/xpty/xpty.go#L76-L108)
     - `WaitProcess()` function handles Windows ConPTY process waiting
*   **Error Recovery**: [xpty/conpty\_windows.go27-34](https://github.com/charmbracelet/x/blob/7642919e/xpty/conpty_windows.go#L27-L34)
     - Terminates process if `os.FindProcess()` fails

**Build Constraint Strategy**

| Build Tag | File | Implementation |
| --- | --- | --- |
| `//go:build windows` | [xpty/conpty\_windows.go1-2](https://github.com/charmbracelet/x/blob/7642919e/xpty/conpty_windows.go#L1-L2) | Full ConPTY support |
| `//go:build !windows` | [xpty/conpty\_other.go1-2](https://github.com/charmbracelet/x/blob/7642919e/xpty/conpty_other.go#L1-L2) | Returns `ErrUnsupported` |

Sources: [xpty/conpty.go1-41](https://github.com/charmbracelet/x/blob/7642919e/xpty/conpty.go#L1-L41)
 [xpty/conpty\_windows.go1-38](https://github.com/charmbracelet/x/blob/7642919e/xpty/conpty_windows.go#L1-L38)
 [xpty/conpty\_other.go1-11](https://github.com/charmbracelet/x/blob/7642919e/xpty/conpty_other.go#L1-L11)

Windows Console API Bindings
----------------------------

The `windows` package provides low-level bindings to the Windows Console API, supporting input processing and console manipulation.

**Input Record Processing**

The `InputRecord` struct [windows/types\_windows.go291-304](https://github.com/charmbracelet/x/blob/7642919e/windows/types_windows.go#L291-L304)
 serves as the central data structure for Windows console events:

| Event Type | Constant | Accessor Method | Data Structure |
| --- | --- | --- | --- |
| Key Events | `KEY_EVENT` | `KeyEvent()` | `KeyEventRecord` |
| Mouse Events | `MOUSE_EVENT` | `MouseEvent()` | `MouseEventRecord` |
| Focus Events | `FOCUS_EVENT` | `FocusEvent()` | `FocusEventRecord` |
| Window Size | `WINDOW_BUFFER_SIZE_EVENT` | `WindowBufferSizeEvent()` | `WindowBufferSizeRecord` |

**Generated System Call Interface**

The package uses Go's syscall generation system:

*   **Definition**: [windows/syscall\_windows.go9-12](https://github.com/charmbracelet/x/blob/7642919e/windows/syscall_windows.go#L9-L12)
     - `//sys` directives define API calls
*   **Generation**: [windows/doc.go3](https://github.com/charmbracelet/x/blob/7642919e/windows/doc.go#L3-L3)
     - `go generate` command creates bindings
*   **Implementation**: [windows/zsyscall\_windows.go38-77](https://github.com/charmbracelet/x/blob/7642919e/windows/zsyscall_windows.go#L38-L77)
     - Generated syscall wrappers

**Virtual Key Code Mappings**

The package defines comprehensive Windows virtual key code constants [windows/types\_windows.go11-173](https://github.com/charmbracelet/x/blob/7642919e/windows/types_windows.go#L11-L173)
 including:

*   Standard keys (`VK_RETURN`, `VK_ESCAPE`, `VK_SPACE`)
*   Function keys (`VK_F1` through `VK_F24`)
*   Navigation keys (`VK_LEFT`, `VK_RIGHT`, `VK_UP`, `VK_DOWN`)
*   Modifier keys (`VK_SHIFT`, `VK_CONTROL`, `VK_MENU`)

Sources: [windows/types\_windows.go1-352](https://github.com/charmbracelet/x/blob/7642919e/windows/types_windows.go#L1-L352)
 [windows/syscall\_windows.go1-13](https://github.com/charmbracelet/x/blob/7642919e/windows/syscall_windows.go#L1-L13)
 [windows/zsyscall\_windows.go1-78](https://github.com/charmbracelet/x/blob/7642919e/windows/zsyscall_windows.go#L1-L78)
 [windows/doc.go1-4](https://github.com/charmbracelet/x/blob/7642919e/windows/doc.go#L1-L4)

Platform-Specific Error Handling
--------------------------------

The platform abstraction layer implements consistent error handling across different operating systems.

**Error Constant Definitions**

*   **Unified Error**: [xpty/xpty.go18-19](https://github.com/charmbracelet/x/blob/7642919e/xpty/xpty.go#L18-L19)
     - `ErrUnsupported = pty.ErrUnsupported`
*   **Windows Errno**: [windows/zsyscall\_windows.go14-36](https://github.com/charmbracelet/x/blob/7642919e/windows/zsyscall_windows.go#L14-L36)
     - Error code translations

**Platform-Specific Error Scenarios**

| Scenario | Unix Behavior | Windows Behavior |
| --- | --- | --- |
| Unsupported OS | Returns `ErrUnsupported` | Returns `ErrUnsupported` |
| PTY Creation Failed | `pty.Open()` error | `conpty.New()` error |
| Process Start Failed | Standard `exec.Cmd` error | ConPTY spawn error + cleanup |
| Resize Failed | `termios` error | ConPTY resize error |

Sources: [xpty/xpty.go18-19](https://github.com/charmbracelet/x/blob/7642919e/xpty/xpty.go#L18-L19)
 [windows/zsyscall\_windows.go14-36](https://github.com/charmbracelet/x/blob/7642919e/windows/zsyscall_windows.go#L14-L36)

Cross-Platform Testing Infrastructure
-------------------------------------

The platform abstraction layer includes comprehensive testing infrastructure to validate behavior across platforms.

**Windows-Specific Test Cases**

The test suite [input/driver\_windows\_test.go15-165](https://github.com/charmbracelet/x/blob/7642919e/input/driver_windows_test.go#L15-L165)
 validates platform abstraction with scenarios including:

*   **Single Key Events**: Basic key press with virtual key codes
*   **Modifier Key Events**: Control, Alt, Shift key combinations
*   **UTF-16 Sequences**: Multi-byte Unicode character handling
*   **Mouse Events**: Click, release, and movement events
*   **Focus Events**: Window focus and blur events
*   **Window Resize**: Buffer size change events

**CI/CD Integration**

The Windows platform includes automated testing infrastructure:

*   **Windows Generation**: [.github/workflows/windows-generate.yml1-36](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/windows-generate.yml#L1-L36)
     - Automatically regenerates Windows syscalls
*   **Platform Matrix**: Tests run on `windows-latest` runners
*   **Dependency Management**: Go module caching and version management

Sources: [input/driver\_windows\_test.go1-266](https://github.com/charmbracelet/x/blob/7642919e/input/driver_windows_test.go#L1-L266)
 [.github/workflows/windows-generate.yml1-36](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/windows-generate.yml#L1-L36)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Platform Abstraction Layer](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer#platform-abstraction-layer)
    
*   [Unified PTY Interface](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer#unified-pty-interface)
    
*   [Unix PTY Implementation](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer#unix-pty-implementation)
    
*   [Windows ConPTY Implementation](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer#windows-conpty-implementation)
    
*   [Windows Console API Bindings](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer#windows-console-api-bindings)
    
*   [Platform-Specific Error Handling](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer#platform-specific-error-handling)
    
*   [Cross-Platform Testing Infrastructure](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer#cross-platform-testing-infrastructure)
