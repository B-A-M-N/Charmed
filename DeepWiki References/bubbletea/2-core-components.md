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

Type Relationships Summary
--------------------------

The following table summarizes how the core types reference and interact with each other in the Bubble Tea architecture.

**Core Type Interactions**

| Type | Defined | Produces | Consumes | Managed By |
| --- | --- | --- | --- | --- |
| `Model` | User code | `Cmd` (from Init/Update)  <br>`View` (from View) | `Msg` (in Update) | User implements interface |
| `Msg` | [tea.go50](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L50-L50) | Nothing (data carrier) | N/A | Input system, commands, signals |
| `Cmd` | [tea.go371](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L371-L371) | `Msg` (when executed) | Nothing | `Program` executes in goroutines |
| `View` | [tea.go84-190](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L84-L190) | Terminal output specification | Model state | `renderer` processes for display |
| `Program` | [tea.go407-533](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L407-L533) | `Msg` (from input/signals)  <br>Execution context | `Model`, `Cmd`, `ProgramOption` | Framework manages lifecycle |

**Sources**: [tea.go50](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L50-L50)
 [tea.go52-65](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L52-L65)
 [tea.go84-190](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L84-L190)
 [tea.go371](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L371-L371)
 [tea.go407-533](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L407-L533)

Core Components
===============

Relevant source files

*   [options.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go)
    
*   [tea.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go)
    
*   [tty.go](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tty.go)
    

This document provides an overview of the fundamental building blocks that make up a Bubble Tea application. These components work together to implement The Elm Architecture pattern in Go, providing a structured approach to building terminal user interfaces.

**Scope**: This page introduces the four core types (`Model`, `Msg`, `Cmd`, `Program`) and their relationships. For detailed information about specific aspects:

*   Program lifecycle and execution: see [Program Lifecycle](https://deepwiki.com/charmbracelet/bubbletea/2.1-program-lifecycle)
    
*   Model interface implementation: see [Model Interface](https://deepwiki.com/charmbracelet/bubbletea/2.2-model-interface)
    
*   Message types and flow: see [Message System](https://deepwiki.com/charmbracelet/bubbletea/2.3-message-system)
    
*   Command composition and execution: see [Command System](https://deepwiki.com/charmbracelet/bubbletea/2.4-command-system)
    
*   Configuration options: see [Program Options](https://deepwiki.com/charmbracelet/bubbletea/2.5-program-options)
    

Architectural Overview
----------------------

Bubble Tea implements The Elm Architecture through four fundamental types that work in concert. The framework follows a unidirectional data flow pattern where messages trigger updates, which return new state and optional commands.

**Sources**: [tea.go43-56](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L43-L56)
 [tea.go39-65](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L39-L65)
 [README.md79-84](https://github.com/charmbracelet/bubbletea/blob/f25595a8/README.md?plain=1#L79-L84)

The Model Interface
-------------------

The `Model` interface is the contract that every Bubble Tea application must implement. It defines three methods that together describe the entire application lifecycle.

| Method | Signature | Purpose |
| --- | --- | --- |
| `Init` | `Init() Cmd` | Returns an optional initial command to execute on startup |
| `Update` | `Update(Msg) (Model, Cmd)` | Processes messages and returns updated state plus optional command |
| `View` | `View() string` | Renders the current state as a string for display |

The `Model` itself can be any Go type (typically a struct) that implements these three methods. It holds all application state.

**Sources**: [tea.go43-56](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L43-L56)
 [tutorials/basics/main.go10-14](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/main.go#L10-L14)
 [tutorials/basics/main.go27-29](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/main.go#L27-L29)
 [tutorials/basics/main.go31-56](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/main.go#L31-L56)
 [tutorials/basics/main.go58-78](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/main.go#L58-L78)

The Message Type
----------------

The `Msg` interface is defined as an empty interface, meaning any Go type can be a message. Messages represent events or data that flow through the system.

**Common Message Categories**:

| Category | Examples | Source |
| --- | --- | --- |
| Built-in Input | `KeyMsg`, `MouseMsg`, `WindowSizeMsg` | Generated by input system |
| Built-in Control | `QuitMsg`, `SuspendMsg`, `InterruptMsg`, `ResumeMsg` | Framework control flow |
| Internal Renderer | `clearScreenMsg`, `enterAltScreenMsg`, `showCursorMsg` | Renderer commands |
| User-defined | Any custom type | Returned from user commands |

The type definition is intentionally minimal to provide maximum flexibility:

Messages flow from three primary sources: user input (keyboard, mouse), internal commands (timers, I/O results), and system signals (resize, interrupt).

**Sources**: [tea.go39-41](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L39-L41)
 [tea.go201-238](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L201-L238)

The Command Function Type
-------------------------

The `Cmd` type represents asynchronous operations that produce messages. Commands are the only way to perform side effects in Bubble Tea while maintaining the pure functional update pattern.

**Key Characteristics**:

*   **Pure I/O**: Commands handle all side effects (HTTP, timers, file I/O)
*   **Asynchronous**: Executed in separate goroutines by the framework
*   **Message Producers**: Return messages that flow back to `Update`
*   **Composable**: Can be combined with `Batch` and `Sequence`

**Sources**: [tea.go58-65](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L58-L65)
 [tea.go333-372](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L333-L372)
 [tea.go364](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L364-L364)

The Program Struct
------------------

The `Program` struct is the central orchestrator that manages the event loop, coordinates all subsystems, and maintains the application lifecycle.

**Key Program Responsibilities**:

**Sources**: [tea.go135-198](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L135-L198)
 [tea.go241-271](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L241-L271)
 [tea.go578-746](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L578-L746)

### Creating and Running a Program

The standard pattern for initializing and running a Bubble Tea application:

The `NewProgram` function accepts a `Model` and zero or more `ProgramOption` functions. The `Run` method blocks until the program terminates, returning the final model state and any error.

**Sources**: [tea.go241-271](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L241-L271)
 [tutorials/basics/main.go80-86](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tutorials/basics/main.go#L80-L86)

Complete Interaction Flow
-------------------------

The following diagram shows how all four core components interact during a typical message processing cycle:

**Sources**: [tea.go382-505](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L382-L505)
 [tea.go493-502](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L493-L502)
 [tea.go333-372](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L333-L372)

Program Configuration
---------------------

Programs are configured through the `ProgramOption` function type. Options are applied during initialization via `NewProgram`.

**Common Configuration Options**:

| Option Function | Purpose | Reference |
| --- | --- | --- |
| `WithContext(ctx)` | Provide cancellation context | [options.go20-24](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L20-L24) |
| `WithInput(reader)` | Custom input source | [options.go38-43](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L38-L43) |
| `WithOutput(writer)` | Custom output destination | [options.go28-32](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L28-L32) |
| `WithAltScreen()` | Start in full-screen mode | [options.go109-113](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L109-L113) |
| `WithMouseCellMotion()` | Enable mouse input | [options.go137-142](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L137-L142) |
| `WithMouseAllMotion()` | Enable hover events | [options.go162-167](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L162-L167) |
| `WithFilter(fn)` | Intercept messages | [options.go226-230](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L226-L230) |
| `WithoutRenderer()` | Disable rendering | [options.go177-181](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L177-L181) |

Options modify the `Program` struct fields during construction. Multiple options can be combined:

For comprehensive documentation of all options, see [Program Options](https://deepwiki.com/charmbracelet/bubbletea/2.5-program-options)
.

**Sources**: [options.go9-15](https://github.com/charmbracelet/bubbletea/blob/f25595a8/options.go#L9-L15)
 [tea.go248-250](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L248-L250)

Type Relationships Summary
--------------------------

The following table summarizes how the core types reference and interact with each other:

| Type | Produces | Consumes | Managed By |
| --- | --- | --- | --- |
| `Model` | `Cmd` (from Init/Update), `string` (from View) | `Msg` (in Update) | User code |
| `Msg` | Nothing (data carrier) | N/A | Input system, commands |
| `Cmd` | `Msg` (when executed) | Nothing | Program (executes asynchronously) |
| `Program` | `Msg` (from input), execution context | `Model`, `Cmd`, options | Framework |

**Sources**: [tea.go43-65](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L43-L65)
 [tea.go135-198](https://github.com/charmbracelet/bubbletea/blob/f25595a8/tea.go#L135-L198)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Type Relationships Summary](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#type-relationships-summary)
    
*   [Core Components](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#core-components)
    
*   [Architectural Overview](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#architectural-overview)
    
*   [The Model Interface](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#the-model-interface)
    
*   [The Message Type](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#the-message-type)
    
*   [The Command Function Type](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#the-command-function-type)
    
*   [The Program Struct](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#the-program-struct)
    
*   [Creating and Running a Program](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#creating-and-running-a-program)
    
*   [Complete Interaction Flow](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#complete-interaction-flow)
    
*   [Program Configuration](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#program-configuration)
    
*   [Type Relationships Summary](https://deepwiki.com/charmbracelet/bubbletea/2-core-components#type-relationships-summary-1)
