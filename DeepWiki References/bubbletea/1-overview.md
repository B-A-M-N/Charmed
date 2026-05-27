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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/bubbletea/blob/f25595a8/README.md?plain=1)
    
*   [examples/go.mod](https://github.com/charmbracelet/bubbletea/blob/f25595a8/examples/go.mod)
    
*   [examples/go.sum](https://github.com/charmbracelet/bubbletea/blob/f25595a8/examples/go.sum)
    
*   [go.mod](https://github.com/charmbracelet/bubbletea/blob/f25595a8/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/bubbletea/blob/f25595a8/go.sum)
    
*   [options.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go)
    
*   [tea.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go)
    
*   [tty.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go)
    
*   [tutorials/basics/README.md](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/README.md?plain=1)
    
*   [tutorials/basics/main.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/main.go)
    
*   [tutorials/commands/README.md](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/commands/README.md?plain=1)
    
*   [tutorials/commands/main.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/commands/main.go)
    

This document provides a high-level introduction to the Bubble Tea framework, covering its purpose, architectural patterns, repository organization, and ecosystem dependencies. Bubble Tea is a Terminal User Interface (TUI) framework for Go based on The Elm Architecture pattern.

For detailed information about The Elm Architecture pattern implementation, see [The Elm Architecture (MVU Pattern)](https://deepwiki.com/charmbracelet/bubbletea/1.1-the-elm-architecture-(mvu-pattern))
. For installation instructions and dependency management, see [Installation and Dependencies](https://deepwiki.com/charmbracelet/bubbletea/1.2-installation-and-dependencies)
.

What is Bubble Tea?
-------------------

Bubble Tea is a Go framework for building terminal applications using a functional, state-driven approach. It implements The Elm Architecture (also known as Model-View-Update or MVU pattern), providing a structured way to manage application state, handle events, and render UI.

The framework handles terminal control, input processing, rendering optimization, and concurrency, allowing developers to focus on application logic. Programs built with Bubble Tea can run inline, in full-window mode (alternate screen), or use a hybrid approach.

**Sources:** [tea.go1-10](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L1-L10)
 [README.md10-22](https://github.com/charmbracelet/bubbletea/blob/f25595a8/README.md?plain=1#L10-L22)

Core Architecture
-----------------

### The Elm Architecture Pattern

Bubble Tea applications are structured around four core concepts defined in the `Model` interface:

    Model Interface (tea.go)
    ├── Init() Cmd          - Returns initial command for startup
    ├── Update(Msg) (Model, Cmd) - Handles events and updates state
    └── View() View         - Renders UI based on current state
    

The `Program` type manages the application lifecycle and coordinates these components.

**Sources:** [tea.go52-65](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L52-L65)
 [README.md79-84](https://github.com/charmbracelet/bubbletea/blob/f25595a8/README.md?plain=1#L79-L84)

### Program Flow Diagram

**Sources:** [tea.go575-622](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L575-L622)
 [tea.go722-864](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L722-L864)
 [tea.go972-1155](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L972-L1155)

### Core Types

| Type | Definition | Purpose |
| --- | --- | --- |
| `Model` | Interface with Init/Update/View methods | Defines application state and behavior |
| `Msg` | `type Msg = uv.Event` | Any event/message type |
| `Cmd` | `type Cmd func() Msg` | I/O operation returning a message |
| `Program` | Struct at [tea.go406-533](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L406-L533) | Manages program lifecycle |
| `View` | Struct at [tea.go82-190](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L82-L190) | Declarative UI specification |
| `ProgramOption` | `type ProgramOption func(*Program)` | Configuration function |

**Sources:** [tea.go48-65](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L48-L65)
 [tea.go82-190](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L82-L190)
 [tea.go364-371](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L364-L371)
 [tea.go406-533](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L406-L533)

Repository Structure
--------------------

The repository is organized into the main module, examples, and tutorials:

### Directory Layout

| Path | Description |
| --- | --- |
| `/*.go` | Core framework implementation |
| `/examples/` | Standalone example applications with separate module |
| `/tutorials/basics/` | Basic tutorial - shopping list application |
| `/tutorials/commands/` | Commands tutorial - HTTP request handling |
| `/testdata/` | Test fixtures and golden files |

**Sources:** [go.mod1-32](https://github.com/charmbracelet/bubbletea/blob/f25595a8/go.mod#L1-L32)
 [examples/go.mod1-50](https://github.com/charmbracelet/bubbletea/blob/f25595a8/examples/go.mod#L1-L50)
 [tutorials/basics/README.md1-13](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/README.md?plain=1#L1-L13)
 [tutorials/commands/README.md1-13](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/commands/README.md?plain=1#L1-L13)

Module Dependencies
-------------------

### Charmbracelet Ecosystem

### Key Dependencies

| Package | Version | Purpose |
| --- | --- | --- |
| `github.com/charmbracelet/colorprofile` | v0.4.1 | Terminal color profile detection |
| `github.com/charmbracelet/ultraviolet` | v0.0.0-20260205113103 | Screen buffer management and diff rendering |
| `github.com/charmbracelet/x/ansi` | v0.11.6 | ANSI escape sequence generation |
| `github.com/charmbracelet/x/term` | v0.2.2 | Terminal feature detection |
| `github.com/muesli/cancelreader` | v0.2.2 | Interruptible I/O reader |
| `golang.org/x/sys` | v0.40.0 | Platform-specific system calls |

**Sources:** [go.mod9-18](https://github.com/charmbracelet/bubbletea/blob/f25595a8/go.mod#L9-L18)
 [go.sum1-39](https://github.com/charmbracelet/bubbletea/blob/f25595a8/go.sum#L1-L39)

Program Lifecycle
-----------------

### Execution Stages

### Goroutines Created

The `Program` runs several concurrent goroutines:

*   **readLoop** [tty.go84-94](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L84-L94)
    : Reads input events from `TerminalReader`
*   **handleSignals** [tea.go624-663](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L624-L663)
    : Listens for SIGINT/SIGTERM
*   **handleResize** [tea.go665-677](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L665-L677)
    : Listens for SIGWINCH (Unix only)
*   **handleCommands** [tea.go681-720](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L681-L720)
    : Executes `Cmd` functions
*   **startRenderer** [renderer.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/renderer.go)
    : Renders views at configured FPS

**Sources:** [tea.go972-1155](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L972-L1155)
 [tea.go624-720](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L624-L720)
 [tty.go84-94](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go#L84-L94)

Quick Start Example
-------------------

The simplest Bubble Tea program implements the `Model` interface:

**Sources:** [tutorials/basics/main.go1-91](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/main.go#L1-L91)
 [README.md92-256](https://github.com/charmbracelet/bubbletea/blob/f25595a8/README.md?plain=1#L92-L256)

Program Options
---------------

Programs can be configured using `ProgramOption` functions passed to `NewProgram`:

| Option | Function | Purpose |
| --- | --- | --- |
| Context | `WithContext(ctx)` | Cancellable context for lifecycle |
| I/O | `WithInput(r)`, `WithOutput(w)` | Custom input/output streams |
| Environment | `WithEnvironment(env)` | Environment variables (for SSH/remote) |
| Signals | `WithoutSignalHandler()` | Disable built-in signal handling |
| Renderer | `WithoutRenderer()` | Disable rendering (daemon mode) |
| Filter | `WithFilter(fn)` | Message filter/interceptor |
| Performance | `WithFPS(fps)` | Set rendering frame rate (default: 60) |
| Colors | `WithColorProfile(profile)` | Force specific color profile |
| Size | `WithWindowSize(w, h)` | Set initial window size |

**Sources:** [options.go1-169](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L1-L169)

Message Types
-------------

Bubble Tea defines several built-in message types:

### System Messages

| Type | Source | Purpose |
| --- | --- | --- |
| `QuitMsg` | `tea.Quit()` | Signals program termination |
| `InterruptMsg` | SIGINT/`tea.Interrupt()` | Signals interrupt (Ctrl+C) |
| `SuspendMsg` | `tea.Suspend()` | Signals suspend (Ctrl+Z on Unix) |
| `ResumeMsg` | Auto after suspend | Signals resume from suspend |
| `WindowSizeMsg` | SIGWINCH | Terminal resize event |

### Input Messages

| Type | Source | Purpose |
| --- | --- | --- |
| `KeyPressMsg` | Keyboard input | Key press event |
| `KeyReleaseMsg` | Keyboard input | Key release event (requires enhancement) |
| `MouseMsg` | Mouse input | Base type for mouse events |
| `MouseClickMsg` | Mouse input | Mouse button click |
| `MouseReleaseMsg` | Mouse input | Mouse button release |
| `MouseWheelMsg` | Mouse input | Mouse wheel scroll |
| `MouseMotionMsg` | Mouse input | Mouse movement |
| `FocusMsg` / `BlurMsg` | Terminal focus | Window focus change |
| `PasteMsg` | Bracketed paste | Paste operation |

### Configuration Messages

| Type | Source | Purpose |
| --- | --- | --- |
| `ColorProfileMsg` | Auto-detected | Terminal color capabilities |
| `KeyboardEnhancementsMsg` | Terminal response | Available keyboard features |
| `EnvMsg` | Program initialization | Environment variables |

**Sources:** [tea.go48-50](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L48-L50)
 [tea.go535-573](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L535-L573)
 [README.md148-149](https://github.com/charmbracelet/bubbletea/blob/f25595a8/README.md?plain=1#L148-L149)

View Structure
--------------

The `View` type declares the UI and terminal features:

**Sources:** [tea.go82-190](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L82-L190)

Next Steps
----------

*   **Learn the MVU pattern**: See [The Elm Architecture (MVU Pattern)](https://deepwiki.com/charmbracelet/bubbletea/1.1-the-elm-architecture-(mvu-pattern))
     for detailed explanation of Model, Init, Update, and View.
*   **Install and setup**: See [Installation and Dependencies](https://deepwiki.com/charmbracelet/bubbletea/1.2-installation-and-dependencies)
     for module setup and dependency management.
*   **Explore core components**: See [Core Components](https://deepwiki.com/charmbracelet/bubbletea/2-core-components)
     for details on Program, Model, Messages, Commands, and Options.
*   **Handle input**: See [Input Handling](https://deepwiki.com/charmbracelet/bubbletea/3-input-handling)
     for keyboard, mouse, and terminal events.
*   **Implement rendering**: See [Rendering System](https://deepwiki.com/charmbracelet/bubbletea/4-rendering-system)
     for view rendering and terminal management.
*   **Review examples**: The [tutorials/basics](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics)
     and [tutorials/commands](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/commands)
     directories provide complete working examples.

**Sources:** [README.md257-268](https://github.com/charmbracelet/bubbletea/blob/f25595a8/README.md?plain=1#L257-L268)
 [tutorials/basics/README.md1-243](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/README.md?plain=1#L1-L243)
 [tutorials/commands/README.md1-247](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/commands/README.md?plain=1#L1-L247)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/bubbletea/1-overview#overview)
    
*   [What is Bubble Tea?](https://deepwiki.com/charmbracelet/bubbletea/1-overview#what-is-bubble-tea)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/bubbletea/1-overview#core-architecture)
    
*   [The Elm Architecture Pattern](https://deepwiki.com/charmbracelet/bubbletea/1-overview#the-elm-architecture-pattern)
    
*   [Program Flow Diagram](https://deepwiki.com/charmbracelet/bubbletea/1-overview#program-flow-diagram)
    
*   [Core Types](https://deepwiki.com/charmbracelet/bubbletea/1-overview#core-types)
    
*   [Repository Structure](https://deepwiki.com/charmbracelet/bubbletea/1-overview#repository-structure)
    
*   [Directory Layout](https://deepwiki.com/charmbracelet/bubbletea/1-overview#directory-layout)
    
*   [Module Dependencies](https://deepwiki.com/charmbracelet/bubbletea/1-overview#module-dependencies)
    
*   [Charmbracelet Ecosystem](https://deepwiki.com/charmbracelet/bubbletea/1-overview#charmbracelet-ecosystem)
    
*   [Key Dependencies](https://deepwiki.com/charmbracelet/bubbletea/1-overview#key-dependencies)
    
*   [Program Lifecycle](https://deepwiki.com/charmbracelet/bubbletea/1-overview#program-lifecycle)
    
*   [Execution Stages](https://deepwiki.com/charmbracelet/bubbletea/1-overview#execution-stages)
    
*   [Goroutines Created](https://deepwiki.com/charmbracelet/bubbletea/1-overview#goroutines-created)
    
*   [Quick Start Example](https://deepwiki.com/charmbracelet/bubbletea/1-overview#quick-start-example)
    
*   [Program Options](https://deepwiki.com/charmbracelet/bubbletea/1-overview#program-options)
    
*   [Message Types](https://deepwiki.com/charmbracelet/bubbletea/1-overview#message-types)
    
*   [System Messages](https://deepwiki.com/charmbracelet/bubbletea/1-overview#system-messages)
    
*   [Input Messages](https://deepwiki.com/charmbracelet/bubbletea/1-overview#input-messages)
    
*   [Configuration Messages](https://deepwiki.com/charmbracelet/bubbletea/1-overview#configuration-messages)
    
*   [View Structure](https://deepwiki.com/charmbracelet/bubbletea/1-overview#view-structure)
    
*   [Next Steps](https://deepwiki.com/charmbracelet/bubbletea/1-overview#next-steps)
