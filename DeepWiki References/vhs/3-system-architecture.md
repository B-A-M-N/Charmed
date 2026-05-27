Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/vhs](https://github.com/charmbracelet/vhs "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 18 April 2025 ([f28239](https://github.com/charmbracelet/vhs/commits/f28239f3)
)

*   [Overview](https://deepwiki.com/charmbracelet/vhs/1-overview)
    
*   [Installation and Setup](https://deepwiki.com/charmbracelet/vhs/1.1-installation-and-setup)
    
*   [Quick Start](https://deepwiki.com/charmbracelet/vhs/1.2-quick-start)
    
*   [Tape Format](https://deepwiki.com/charmbracelet/vhs/2-tape-format)
    
*   [Commands Reference](https://deepwiki.com/charmbracelet/vhs/2.1-commands-reference)
    
*   [Output Options](https://deepwiki.com/charmbracelet/vhs/2.2-output-options)
    
*   [Themes and Styling](https://deepwiki.com/charmbracelet/vhs/2.3-themes-and-styling)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/vhs/3-system-architecture)
    
*   [Command Execution](https://deepwiki.com/charmbracelet/vhs/3.1-command-execution)
    
*   [Terminal Emulation](https://deepwiki.com/charmbracelet/vhs/3.2-terminal-emulation)
    
*   [Recording and Rendering](https://deepwiki.com/charmbracelet/vhs/3.3-recording-and-rendering)
    
*   [CLI Features](https://deepwiki.com/charmbracelet/vhs/4-cli-features)
    
*   [Recording Mode](https://deepwiki.com/charmbracelet/vhs/4.1-recording-mode)
    
*   [SSH Server](https://deepwiki.com/charmbracelet/vhs/4.2-ssh-server)
    
*   [Publishing](https://deepwiki.com/charmbracelet/vhs/4.3-publishing)
    
*   [Examples and Use Cases](https://deepwiki.com/charmbracelet/vhs/5-examples-and-use-cases)
    
*   [Demos for CLI Tools](https://deepwiki.com/charmbracelet/vhs/5.1-demos-for-cli-tools)
    
*   [Shell Examples](https://deepwiki.com/charmbracelet/vhs/5.2-shell-examples)
    
*   [Advanced Tape Techniques](https://deepwiki.com/charmbracelet/vhs/5.3-advanced-tape-techniques)
    
*   [Development and Contribution](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution)
    
*   [Building from Source](https://deepwiki.com/charmbracelet/vhs/6.1-building-from-source)
    
*   [Project Structure](https://deepwiki.com/charmbracelet/vhs/6.2-project-structure)
    

Menu

System Architecture
===================

Relevant source files

*   [evaluator.go](https://github.com/charmbracelet/vhs/blob/f28239f3/evaluator.go)
    
*   [examples/demo.gif](https://github.com/charmbracelet/vhs/blob/f28239f3/examples/demo.gif)
    
*   [go.mod](https://github.com/charmbracelet/vhs/blob/f28239f3/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/vhs/blob/f28239f3/go.sum)
    
*   [main.go](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go)
    
*   [publish.go](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go)
    
*   [screenshot.go](https://github.com/charmbracelet/vhs/blob/f28239f3/screenshot.go)
    
*   [serve.go](https://github.com/charmbracelet/vhs/blob/f28239f3/serve.go)
    
*   [style.go](https://github.com/charmbracelet/vhs/blob/f28239f3/style.go)
    
*   [testing.go](https://github.com/charmbracelet/vhs/blob/f28239f3/testing.go)
    
*   [tty.go](https://github.com/charmbracelet/vhs/blob/f28239f3/tty.go)
    
*   [vhs.go](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go)
    
*   [video.go](https://github.com/charmbracelet/vhs/blob/f28239f3/video.go)
    

This page provides a technical overview of VHS's internal architecture, detailing how the system processes tape files, manages terminal recording, and generates output videos. This document is intended for developers who want to understand how VHS works internally or contribute to the codebase.

For information about the tape file format itself, see [Tape Format](https://deepwiki.com/charmbracelet/vhs/2-tape-format)
. For details about command execution implementation, see [Command Execution](https://deepwiki.com/charmbracelet/vhs/3.1-command-execution)
.

High-Level Architecture Overview
--------------------------------

VHS follows a pipeline architecture that transforms a declarative tape file into a recorded terminal session. The system is composed of several interconnected components that handle different aspects of the recording process.

Sources: [main.go243-254](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L243-L254)
 [vhs.go22-412](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go#L22-L412)
 [evaluator.go16-188](https://github.com/charmbracelet/vhs/blob/f28239f3/evaluator.go#L16-L188)
 [tty.go1-46](https://github.com/charmbracelet/vhs/blob/f28239f3/tty.go#L1-L46)
 [video.go1-185](https://github.com/charmbracelet/vhs/blob/f28239f3/video.go#L1-L185)

Core Components
---------------

### VHS Object

The central component of the system is the `VHS` struct, which maintains the state of the recording session and coordinates the interaction between different components.

Sources: [vhs.go22-36](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go#L22-L36)
 [vhs.go38-55](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go#L38-L55)

The `VHS` struct manages:

*   Terminal emulation through ttyd
*   Browser automation through go-rod
*   Frame recording and state
*   Video rendering through FFmpeg

### Command Processing Pipeline

VHS uses a lexer-parser-evaluator pipeline to process tape files:

Sources: [evaluator.go16-188](https://github.com/charmbracelet/vhs/blob/f28239f3/evaluator.go#L16-L188)
 [main.go218-239](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L218-L239)

The command processing flow works as follows:

1.  The lexer tokenizes the tape file content
2.  The parser converts tokens into command objects
3.  The evaluator processes commands in sequence
4.  Different command types affect different parts of the system:
    *   Setup commands configure options
    *   Shell commands interact with the terminal
    *   Control commands manage the recording process

### Terminal Emulation

VHS uses ttyd to create a browser-accessible terminal:

Sources: [tty.go19-46](https://github.com/charmbracelet/vhs/blob/f28239f3/tty.go#L19-L46)
 [vhs.go124-156](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go#L124-L156)

Key aspects of the terminal emulation:

*   A random port is selected for ttyd to avoid conflicts
*   ttyd creates a WebSocket server that provides terminal access
*   Special ttyd settings enable features like Sixel graphics and canvas rendering
*   The shell is configured based on user preferences or defaults

### Frame Recording Process

VHS captures frames from the terminal display using browser automation:

Sources: [vhs.go323-388](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go#L323-L388)

The recording process:

1.  VHS finds the xterm.js canvas elements for text and cursor layers
2.  At a regular interval (determined by framerate), VHS captures the canvas contents
3.  Frames are saved as separate text and cursor PNG files
4.  The recording can be paused and resumed with `PauseRecording()` and `ResumeRecording()`

### Video Generation

After recording, VHS uses FFmpeg to convert the captured frames into video outputs:

Sources: [video.go139-185](https://github.com/charmbracelet/vhs/blob/f28239f3/video.go#L139-L185)
 [video.go93-137](https://github.com/charmbracelet/vhs/blob/f28239f3/video.go#L93-L137)

The video generation process:

1.  VHS applies any loop offset to the frame sequence
2.  For each output format (GIF, MP4, WebM), VHS:
    *   Builds an FFmpeg command with appropriate filters
    *   Processes text and cursor frames separately
    *   Applies styling options (padding, window bar, etc.)
    *   Generates the output file

Command Execution System
------------------------

The command execution system is responsible for processing the commands parsed from the tape file:

Sources: [evaluator.go16-188](https://github.com/charmbracelet/vhs/blob/f28239f3/evaluator.go#L16-L188)

The command execution system works as follows:

1.  Commands from the parsed tape are processed sequentially
2.  Each command is mapped to a function that implements its behavior
3.  The function manipulates the VHS instance state or interacts with the terminal
4.  Setup commands are executed before recording starts
5.  Shell commands simulate user input in the terminal
6.  Control commands manage the recording process

Integration with External Systems
---------------------------------

VHS integrates with several external systems to provide its functionality:

Sources: [main.go313-331](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L313-L331)
 [serve.go37-165](https://github.com/charmbracelet/vhs/blob/f28239f3/serve.go#L37-L165)
 [publish.go27-196](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go#L27-L196)

Key external integrations:

*   ttyd: Creates a browser-accessible terminal
*   FFmpeg: Processes frames and generates video outputs
*   go-rod: Automates browser interactions with the terminal
*   SSH: Enables remote execution of tapes
*   charm.sh: Provides hosting for published recordings

Application Startup and Execution Flow
--------------------------------------

This diagram shows the flow of execution from starting VHS to generating output:

Sources: [main.go243-254](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L243-L254)
 [evaluator.go21-188](https://github.com/charmbracelet/vhs/blob/f28239f3/evaluator.go#L21-L188)
 [vhs.go124-247](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go#L124-L247)

Advanced Features
-----------------

### SSH Server Mode

VHS can run as an SSH server to enable remote execution of tapes:

Sources: [serve.go37-165](https://github.com/charmbracelet/vhs/blob/f28239f3/serve.go#L37-L165)

The SSH server mode:

1.  Starts an SSH server on a specified port
2.  Accepts connections from SSH clients
3.  Reads a tape file from stdin
4.  Evaluates the tape and generates output
5.  Returns the generated GIF via stdout

### Publishing System

VHS can publish recordings to the charm.sh server:

Sources: [publish.go149-196](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go#L149-L196)
 [publish.go27-56](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go#L27-L56)

The publishing system:

1.  Takes a local GIF file as input
2.  Establishes an SSH connection to ghost.charm.sh
3.  Uploads the GIF via the SSH connection
4.  Receives and returns a shareable URL

Data Flow During Recording
--------------------------

This diagram shows the detailed data flow during the recording process:

Sources: [evaluator.go21-188](https://github.com/charmbracelet/vhs/blob/f28239f3/evaluator.go#L21-L188)
 [vhs.go323-388](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go#L323-L388)
 [vhs.go222-247](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go#L222-L247)
 [video.go139-185](https://github.com/charmbracelet/vhs/blob/f28239f3/video.go#L139-L185)

Error Handling and Recovery
---------------------------

VHS implements several error handling mechanisms to ensure that resources are properly cleaned up and meaningful error messages are provided to users:

Sources: [vhs.go198-220](https://github.com/charmbracelet/vhs/blob/f28239f3/vhs.go#L198-L220)
 [evaluator.go125-186](https://github.com/charmbracelet/vhs/blob/f28239f3/evaluator.go#L125-L186)
 [main.go119-121](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L119-L121)

Summary
-------

VHS's architecture is built around a pipeline that transforms a declarative tape file into a recorded terminal session and then into video outputs. The system leverages several external tools (ttyd, go-rod, FFmpeg) and integrates them into a cohesive workflow.

Key architectural components include:

*   A lexer/parser system for tape files
*   A command execution system for interpreting commands
*   Terminal emulation via ttyd
*   Browser automation via go-rod
*   Frame recording from terminal canvases
*   Video generation via FFmpeg
*   Optional features like SSH server and publishing

This modular design allows for flexibility in how terminal sessions are recorded and reproduced, making VHS a powerful tool for creating terminal demos, documentation, and integration tests.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [System Architecture](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#system-architecture)
    
*   [High-Level Architecture Overview](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#high-level-architecture-overview)
    
*   [Core Components](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#core-components)
    
*   [VHS Object](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#vhs-object)
    
*   [Command Processing Pipeline](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#command-processing-pipeline)
    
*   [Terminal Emulation](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#terminal-emulation)
    
*   [Frame Recording Process](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#frame-recording-process)
    
*   [Video Generation](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#video-generation)
    
*   [Command Execution System](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#command-execution-system)
    
*   [Integration with External Systems](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#integration-with-external-systems)
    
*   [Application Startup and Execution Flow](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#application-startup-and-execution-flow)
    
*   [Advanced Features](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#advanced-features)
    
*   [SSH Server Mode](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#ssh-server-mode)
    
*   [Publishing System](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#publishing-system)
    
*   [Data Flow During Recording](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#data-flow-during-recording)
    
*   [Error Handling and Recovery](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#error-handling-and-recovery)
    
*   [Summary](https://deepwiki.com/charmbracelet/vhs/3-system-architecture#summary)
