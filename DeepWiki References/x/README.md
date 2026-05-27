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

Overview
========

Relevant source files

*   [.github/dependabot.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/dependabot.yml)
    
*   [.github/workflows/gitignore.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/gitignore.yml)
    
*   [.github/workflows/vttest.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/vttest.yml)
    
*   [.gitignore](https://github.com/charmbracelet/x/blob/7642919e/.gitignore)
    
*   [LICENSE](https://github.com/charmbracelet/x/blob/7642919e/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1)
    
*   [exp/golden/testdata/TestTypes/SliceOfBytes.golden](https://github.com/charmbracelet/x/blob/7642919e/exp/golden/testdata/TestTypes/SliceOfBytes.golden)
    
*   [exp/golden/testdata/TestTypes/String.golden](https://github.com/charmbracelet/x/blob/7642919e/exp/golden/testdata/TestTypes/String.golden)
    
*   [scripts/dependabot](https://github.com/charmbracelet/x/blob/7642919e/scripts/dependabot)
    

Purpose and Scope
-----------------

The `charmbracelet/x` repository is a Go workspace containing experimental and utility packages for terminal-based applications. It serves as an incubation space where terminal-related functionality matures before potential graduation to standalone repositories within the Charm ecosystem.

The repository provides three primary capabilities:

1.  **Terminal Rendering**: Cell-based screen buffer management (`cellbuf.Screen`), virtual terminal emulation (`vt.Terminal`), and ANSI sequence generation
2.  **Input Processing**: Event-based input handling (`input.Reader`) with cross-platform keyboard, mouse, and window resize support
3.  **Platform Abstraction**: Unified interfaces for PTY operations (`xpty.PTY`), terminal I/O control (`term` package), and Windows/Unix-specific implementations

All packages are experimental with no backward compatibility guarantees. The monorepo structure contains 30+ independent Go modules, each with separate versioning managed through automated CI/CD workflows.

**Sources:** [README.md1-60](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1#L1-L60)
 [go.work1-31](https://github.com/charmbracelet/x/blob/7642919e/go.work#L1-L31)

Repository Structure
--------------------

The repository is organized as a Go 1.24.0 workspace with 30+ specialized modules managed through [go.work1-31](https://github.com/charmbracelet/x/blob/7642919e/go.work#L1-L31)
 Each module has independent versioning and dedicated CI/CD workflows.

### Module Organization

| Category | Modules | Key Types |
| --- | --- | --- |
| **Core Rendering** | `cellbuf`, `vt`, `ansi` | `cellbuf.Screen`, `vt.Terminal`, `ansi.Parser` |
| **Input/Control** | `input`, `xpty`, `term`, `termios` | `input.Reader`, `xpty.PTY`, `term.State` |
| **Platform Support** | `windows`, `conpty`, `wcwidth` | `windows.Console`, `conpty.Pty` |
| **Utilities** | `colors`, `sshkey`, `editor`, `errors`, `json` | `colors.Color`, `sshkey.Key` |
| **Experimental** | `exp/teatest`, `exp/golden`, `exp/slice`, etc. | `teatest.TestModel`, `golden.RequireEqual` |
| **Development** | `examples`, `mosaic`, `gitignore`, `vttest` | Example applications and testing tools |

**Sources:** [README.md15-48](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1#L15-L48)
 [go.work3-31](https://github.com/charmbracelet/x/blob/7642919e/go.work#L3-L31)
 [.github/dependabot.yml23-565](https://github.com/charmbracelet/x/blob/7642919e/.github/dependabot.yml#L23-L565)

System Architecture
-------------------

### Layered Architecture

The system follows a layered design where high-level abstractions depend on progressively lower-level platform primitives:

**Layered System Architecture**

Applications interact with high-level types (`cellbuf.Screen`, `vt.Terminal`, `input.Reader`) that delegate to the ANSI processing layer for parsing and generation. The PTY and terminal control layer provides unified interfaces that branch into platform-specific implementations at the bottom layers.

**Sources:** [README.md17-48](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1#L17-L48)
 [go.work3-31](https://github.com/charmbracelet/x/blob/7642919e/go.work#L3-L31)

### Core Component Integration

The following diagram maps concrete code entities used in example applications:

**Code Entity Integration Map**

**Sources:** [examples/cellbuf/main.go1-120](https://github.com/charmbracelet/x/blob/7642919e/examples/cellbuf/main.go#L1-L120)
 [examples/layout/main.go1-550](https://github.com/charmbracelet/x/blob/7642919e/examples/layout/main.go#L1-L550)

### Text Processing Pipeline

Text flows through multiple processing stages from raw input to terminal output:

**Text Processing Data Flow**

The parser tokenizes ANSI sequences using a state machine, width calculation handles Unicode grapheme clustering, text manipulation functions provide ANSI-aware operations, and rendering uses a double-buffer diff algorithm to generate minimal escape sequences.

**Sources:** [ansi/parser.go1-500](https://github.com/charmbracelet/x/blob/7642919e/ansi/parser.go#L1-L500)
 [cellbuf/screen.go1-300](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L1-L300)

Module Categories and Importance
--------------------------------

### Core Terminal Rendering (High Importance)

The rendering system forms the most critical layer of the ecosystem:

| Module | Importance | Primary Types | Purpose |
| --- | --- | --- | --- |
| `cellbuf` | ★★★★★ (227.99) | `Screen`, `Cell`, `Buffer`, `ScreenWriter` | Double-buffered screen management with diff-based rendering |
| `vt` | ★★★ (108.10) | `Terminal`, `State` | Virtual terminal emulator with ANSI sequence interpretation |
| `ansi` | ★★★ (87.14) | `Parser`, `Style`, `Cmd`, `StringWidth` | ANSI escape sequence parsing and text manipulation |

The `cellbuf.Screen` type is the highest-importance component, providing efficient cell-based rendering through a double-buffer mechanism that minimizes terminal I/O.

**Sources:** [README.md17-18](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1#L17-L18)
 [cellbuf/screen.go1-300](https://github.com/charmbracelet/x/blob/7642919e/cellbuf/screen.go#L1-L300)
 [vt/terminal.go1-400](https://github.com/charmbracelet/x/blob/7642919e/vt/terminal.go#L1-L400)

### Input & Terminal Control (High Importance)

Input processing provides event-based handling across platforms:

| Module | Importance | Primary Types | Purpose |
| --- | --- | --- | --- |
| `input` | ★★★★ (156.89) | `Reader`, `KeyPressEvent`, `MouseClickEvent`, `WindowSizeEvent` | Event-based terminal input with keyboard, mouse, and resize support |
| `xpty` | ★★★★ (156.89) | `PTY` interface | Cross-platform pseudo-terminal interface |
| `term` | ★★★★ (156.89) | `State`, `MakeRaw`, `GetSize`, `Restore` | Terminal state management and mode control |

**Sources:** [README.md28-42](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1#L28-L42)
 [input/reader.go1-500](https://github.com/charmbracelet/x/blob/7642919e/input/reader.go#L1-L500)
 [term/term.go1-200](https://github.com/charmbracelet/x/blob/7642919e/term/term.go#L1-L200)

### Platform-Specific Implementations

Platform abstraction isolates OS differences at the bottom layers:

| Module | Importance | Platform | Purpose |
| --- | --- | --- | --- |
| `windows` | ★★ (53.81) | Windows | Windows Console API bindings (`ReadConsoleInput`) |
| `conpty` | ★★ (49.32) | Windows | Windows ConPTY (`CreatePseudoConsole`) |
| `termios` | ★ (39.29) | Unix/Linux | Unix terminal I/O control (`tcgetattr`, `tcsetattr`) |

Build tags (`//go:build windows`, `//go:build !windows`) control platform-specific compilation.

**Sources:** [README.md46-47](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1#L46-L47)
 [windows/console.go1-300](https://github.com/charmbracelet/x/blob/7642919e/windows/console.go#L1-L300)
 [conpty/conpty.go1-200](https://github.com/charmbracelet/x/blob/7642919e/conpty/conpty.go#L1-L200)

### Utility Modules

Supporting utilities provide color handling, character width calculation, and SSH key management:

| Module | Importance | Primary Types | Purpose |
| --- | --- | --- | --- |
| `sshkey` | ★★★ (72.76) | `Key`, `ParseKey` | SSH key parsing with passphrase handling |
| `colors` | ★★★ (72.76) | `Color`, ANSI/X11 definitions | Terminal color definitions with lipgloss integration |
| `wcwidth` | ★ (39.29) | `RuneWidth`, `StringWidth` | Wide character width calculation |

**Sources:** [README.md20-46](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1#L20-L46)
 [sshkey/sshkey.go1-200](https://github.com/charmbracelet/x/blob/7642919e/sshkey/sshkey.go#L1-L200)
 [colors/colors.go1-300](https://github.com/charmbracelet/x/blob/7642919e/colors/colors.go#L1-L300)

Development Infrastructure
--------------------------

### Example Applications

The `examples` directory contains integration demonstrations that also serve as integration tests:

| Example | Primary Packages | Purpose |
| --- | --- | --- |
| `examples/cellbuf` | `cellbuf`, `input`, `term` | Basic cell-based rendering with mouse support |
| `examples/layout` | `cellbuf`, `lipgloss/v2` | Lipgloss layout integration |
| `examples/parserlog` | `ansi` | ANSI parser demonstration |
| `examples/vtgl` | `vt` | VT emulator with OpenGL rendering |

**Example Application Dependency Graph**

**Sources:** [examples/cellbuf/main.go1-120](https://github.com/charmbracelet/x/blob/7642919e/examples/cellbuf/main.go#L1-L120)
 [examples/layout/main.go1-550](https://github.com/charmbracelet/x/blob/7642919e/examples/layout/main.go#L1-L550)
 [README.md17-48](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1#L17-L48)

### CI/CD and Dependency Management

Automated workflows manage 30+ independent Go modules through GitHub Actions and Dependabot:

**CI/CD Infrastructure**

The [scripts/dependabot25-43](https://github.com/charmbracelet/x/blob/7642919e/scripts/dependabot#L25-L43)
 script automatically generates [.github/dependabot.yml1-566](https://github.com/charmbracelet/x/blob/7642919e/.github/dependabot.yml#L1-L566)
 by scanning for `go.mod` files. Each module gets dedicated workflow files that run builds and tests on all three platforms.

**Sources:** [scripts/dependabot1-44](https://github.com/charmbracelet/x/blob/7642919e/scripts/dependabot#L1-L44)
 [.github/dependabot.yml1-566](https://github.com/charmbracelet/x/blob/7642919e/.github/dependabot.yml#L1-L566)
 [.github/workflows/vttest.yml1-86](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/vttest.yml#L1-L86)
 [.github/workflows/gitignore.yml1-63](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/gitignore.yml#L1-L63)

### Testing Framework

The `exp/teatest` package provides testing utilities for Bubble Tea applications:

| Package | Purpose |
| --- | --- |
| `exp/teatest` | v1 testing framework with `TestModel`, `WaitFor` |
| `exp/teatest/v2` | v2 framework with enhanced multi-program support |
| `exp/golden` | Golden file testing with `RequireEqual` |

The `teatest.TestModel` function creates isolated test environments for Bubble Tea programs, while `exp/golden` provides snapshot-based testing for terminal output verification.

**Sources:** [go.work19-20](https://github.com/charmbracelet/x/blob/7642919e/go.work#L19-L20)
 [README.md39](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1#L39-L39)
 [exp/golden/testdata/TestTypes/String.golden1](https://github.com/charmbracelet/x/blob/7642919e/exp/golden/testdata/TestTypes/String.golden#L1-L1)
 [exp/golden/testdata/TestTypes/SliceOfBytes.golden1](https://github.com/charmbracelet/x/blob/7642919e/exp/golden/testdata/TestTypes/SliceOfBytes.golden#L1-L1)

Integration Patterns
--------------------

Key integration patterns demonstrate how modules work together:

1.  **Screen Management**: `cellbuf.NewScreen()` + `term.GetSize()` + `term.MakeRaw()`
2.  **Input Processing**: `input.NewReader()` + event type switching on `KeyPressEvent`/`MouseClickEvent`
3.  **ANSI Control**: `ansi.SetMode()` + `ansi.ResetMode()` for mouse events and terminal modes
4.  **Cross-Platform PTY**: `xpty` abstraction over platform-specific implementations

**Sources:** [examples/cellbuf/main.go33-52](https://github.com/charmbracelet/x/blob/7642919e/examples/cellbuf/main.go#L33-L52)
 [examples/layout/main.go387-405](https://github.com/charmbracelet/x/blob/7642919e/examples/layout/main.go#L387-L405)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/x#overview)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/x#purpose-and-scope)
    
*   [Repository Structure](https://deepwiki.com/charmbracelet/x#repository-structure)
    
*   [Module Organization](https://deepwiki.com/charmbracelet/x#module-organization)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/x#system-architecture)
    
*   [Layered Architecture](https://deepwiki.com/charmbracelet/x#layered-architecture)
    
*   [Core Component Integration](https://deepwiki.com/charmbracelet/x#core-component-integration)
    
*   [Text Processing Pipeline](https://deepwiki.com/charmbracelet/x#text-processing-pipeline)
    
*   [Module Categories and Importance](https://deepwiki.com/charmbracelet/x#module-categories-and-importance)
    
*   [Core Terminal Rendering (High Importance)](https://deepwiki.com/charmbracelet/x#core-terminal-rendering-high-importance)
    
*   [Input & Terminal Control (High Importance)](https://deepwiki.com/charmbracelet/x#input-terminal-control-high-importance)
    
*   [Platform-Specific Implementations](https://deepwiki.com/charmbracelet/x#platform-specific-implementations)
    
*   [Utility Modules](https://deepwiki.com/charmbracelet/x#utility-modules)
    
*   [Development Infrastructure](https://deepwiki.com/charmbracelet/x#development-infrastructure)
    
*   [Example Applications](https://deepwiki.com/charmbracelet/x#example-applications)
    
*   [CI/CD and Dependency Management](https://deepwiki.com/charmbracelet/x#cicd-and-dependency-management)
    
*   [Testing Framework](https://deepwiki.com/charmbracelet/x#testing-framework)
    
*   [Integration Patterns](https://deepwiki.com/charmbracelet/x#integration-patterns)
