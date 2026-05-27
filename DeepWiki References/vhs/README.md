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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1)
    
*   [examples/demo.tape](https://github.com/charmbracelet/vhs/blob/f28239f3/examples/demo.tape)
    
*   [go.mod](https://github.com/charmbracelet/vhs/blob/f28239f3/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/vhs/blob/f28239f3/go.sum)
    
*   [main.go](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go)
    
*   [publish.go](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go)
    
*   [screenshot.go](https://github.com/charmbracelet/vhs/blob/f28239f3/screenshot.go)
    
*   [serve.go](https://github.com/charmbracelet/vhs/blob/f28239f3/serve.go)
    
*   [style.go](https://github.com/charmbracelet/vhs/blob/f28239f3/style.go)
    

VHS (Video Home System) is a tool for writing terminal GIFs as code, enabling users to create reproducible terminal recordings for documentation, demos, and integration tests. This page provides a high-level introduction to VHS, explaining its purpose, architecture, and capabilities.

Purpose and Key Features
------------------------

VHS allows users to create and share terminal GIFs programmatically using a declarative "tape" file format. It addresses several needs:

*   Creating reproducible terminal recordings for documentation
*   Generating demo GIFs for showcasing CLI tools
*   Automating the process of terminal recording
*   Providing a way to share terminal sessions easily

Key features include:

*   Declarative tape file format for defining recordings
*   Support for multiple output formats (GIF, MP4, WebM)
*   Terminal emulation with customizable themes and styling
*   Recording mode for capturing real terminal sessions
*   SSH server for remote execution
*   Publishing capability for sharing recordings

Sources: [README.md12-16](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L12-L16)
 [README.md159-186](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L159-L186)

System Architecture
-------------------

The following diagram illustrates the high-level architecture of VHS:

VHS works by:

1.  Reading a tape file containing commands
2.  Executing those commands in a virtual terminal (using ttyd)
3.  Recording the terminal output as frames
4.  Converting those frames into GIFs or videos (using FFmpeg)

Sources: [main.go1-331](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L1-L331)
 [serve.go1-166](https://github.com/charmbracelet/vhs/blob/f28239f3/serve.go#L1-L166)
 [publish.go1-196](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go#L1-L196)

Tape Execution Process
----------------------

The following sequence diagram shows the process of executing a tape file:

Sources: [main.go56-146](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L56-L146)

Command System
--------------

VHS supports various commands in tape files, which can be categorized as follows:

Sources: [README.md219-236](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L219-L236)
 [examples/demo.tape1-70](https://github.com/charmbracelet/vhs/blob/f28239f3/examples/demo.tape#L1-L70)

Key Components
--------------

VHS consists of several key components that work together:

| Component | Description |
| --- | --- |
| CLI Interface | Processes user commands and options using Cobra |
| Tape Parser | Lexes and parses tape files into executable commands |
| Command Executor | Executes commands by controlling a virtual terminal |
| Terminal Emulation | Uses ttyd and browser automation (go-rod) to create a controllable terminal |
| Frame Recorder | Captures frames from the terminal display |
| Video Generator | Uses FFmpeg to convert frames to GIF, MP4, or WebM |
| SSH Server | Enables remote execution of tapes |
| Publishing | Allows sharing recordings via vhs.charm.sh |

Sources: [main.go1-331](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L1-L331)
 [serve.go1-166](https://github.com/charmbracelet/vhs/blob/f28239f3/serve.go#L1-L166)
 [publish.go1-196](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go#L1-L196)
 [screenshot.go1-101](https://github.com/charmbracelet/vhs/blob/f28239f3/screenshot.go#L1-L101)

Main Workflows
--------------

VHS supports several main workflows:

### Running a Tape File

This is the primary workflow for VHS. Users create a `.tape` file with a series of commands, then run it with VHS to generate a GIF or video.

Sources: [main.go56-146](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L56-L146)
 [README.md18-73](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L18-L73)

### Recording Mode

The recording mode allows users to generate tape files from their terminal actions, making it easier to create complex recordings.

Sources: [main.go172-178](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L172-L178)
 [README.md159-174](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L159-L174)

### SSH Server

VHS can run as an SSH server, allowing remote users to execute tape files without installing VHS locally.

Sources: [serve.go1-166](https://github.com/charmbracelet/vhs/blob/f28239f3/serve.go#L1-L166)
 [README.md187-215](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L187-L215)

### Publishing

VHS allows users to publish their GIFs to vhs.charm.sh for easy sharing with others.

Sources: [publish.go1-196](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go#L1-L196)
 [README.md176-186](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L176-L186)

Dependencies and Requirements
-----------------------------

VHS requires several external dependencies to function:

1.  **ttyd**: A web-based terminal emulator that VHS uses to create and control a virtual terminal
2.  **FFmpeg**: Used for converting captured frames into GIFs, MP4s, or WebMs
3.  **Shell**: A terminal shell (bash, zsh, etc.) for executing commands

VHS checks for these dependencies at startup and will fail with a clear error message if any are missing.

Sources: [main.go306-331](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L306-L331)

Examples of Use Cases
---------------------

VHS is particularly useful for:

*   Creating documentation for CLI tools
*   Generating demos for GitHub READMEs
*   Creating integration tests for terminal applications
*   Sharing terminal sessions with others

For more information on getting started with VHS, see the [Installation and Setup](https://deepwiki.com/charmbracelet/vhs/1.1-installation-and-setup)
 and [Quick Start](https://deepwiki.com/charmbracelet/vhs/1.2-quick-start)
 pages.

Sources: [README.md1-17](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L1-L17)

Environmental Configuration
---------------------------

VHS supports various environment variables for configuration, especially for the SSH server mode:

| Variable | Description | Default |
| --- | --- | --- |
| VHS\_PORT | SSH server port | 1976 |
| VHS\_HOST | SSH server host | localhost |
| VHS\_KEY\_PATH | Path to SSH key | .ssh/vhs\_ed25519 |
| VHS\_AUTHORIZED\_KEYS\_PATH | Path to authorized keys file | (empty, publicly accessible) |

Sources: [serve.go28-35](https://github.com/charmbracelet/vhs/blob/f28239f3/serve.go#L28-L35)
 [README.md199-209](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L199-L209)

This overview provides a high-level introduction to VHS. The subsequent pages will go into more detail on specific aspects of the tool, such as the tape format, command reference, and advanced features.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/vhs#overview)
    
*   [Purpose and Key Features](https://deepwiki.com/charmbracelet/vhs#purpose-and-key-features)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/vhs#system-architecture)
    
*   [Tape Execution Process](https://deepwiki.com/charmbracelet/vhs#tape-execution-process)
    
*   [Command System](https://deepwiki.com/charmbracelet/vhs#command-system)
    
*   [Key Components](https://deepwiki.com/charmbracelet/vhs#key-components)
    
*   [Main Workflows](https://deepwiki.com/charmbracelet/vhs#main-workflows)
    
*   [Running a Tape File](https://deepwiki.com/charmbracelet/vhs#running-a-tape-file)
    
*   [Recording Mode](https://deepwiki.com/charmbracelet/vhs#recording-mode)
    
*   [SSH Server](https://deepwiki.com/charmbracelet/vhs#ssh-server)
    
*   [Publishing](https://deepwiki.com/charmbracelet/vhs#publishing)
    
*   [Dependencies and Requirements](https://deepwiki.com/charmbracelet/vhs#dependencies-and-requirements)
    
*   [Examples of Use Cases](https://deepwiki.com/charmbracelet/vhs#examples-of-use-cases)
    
*   [Environmental Configuration](https://deepwiki.com/charmbracelet/vhs#environmental-configuration)
