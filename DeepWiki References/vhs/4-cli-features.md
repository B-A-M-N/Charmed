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

CLI Features
============

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
    

This document provides a comprehensive overview of the VHS command-line interface (CLI) features. VHS offers a rich set of commands that allow you to create, execute, and share terminal GIF recordings. This page focuses on the CLI itself and the commands you can use, rather than the tape file format (covered in [Tape Format](https://deepwiki.com/charmbracelet/vhs/2-tape-format)
) or the internal architecture (covered in [System Architecture](https://deepwiki.com/charmbracelet/vhs/3-system-architecture)
).

Command Overview
----------------

VHS provides several commands to create, run, and share terminal recordings. The main CLI commands include:

Sources: [main.go43-148](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L43-L148)
 [README.md159-216](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L159-L216)

Root Command: Running Tape Files
--------------------------------

The primary function of VHS is to execute tape files. The root command (`vhs <file>`) without any subcommands will run a tape file and generate the corresponding output.

The root command supports the following flags:

| Flag | Short | Description |
| --- | --- | --- |
| `--publish` | `-p` | Publish your GIF to vhs.charm.sh and get a shareable URL |
| `--output` | `-o` | Specify output filename(s) of video output |
| `--quiet` | `-q` | Quiet mode, do not log messages |

Sources: [main.go43-148](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L43-L148)
 [main.go257-260](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L257-L260)

Recording Tapes
---------------

The `record` command allows you to generate tape files by recording your terminal actions:

Usage:

    vhs record > example.tape
    

After recording, you can edit the generated tape file to add or modify settings, then run it with the root command to create your GIF.

Sources: [main.go173-178](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L173-L178)
 [README.md158-176](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L158-L176)

Publishing GIFs
---------------

VHS allows you to publish your GIFs for easy sharing:

You can publish a GIF in two ways:

1.  With a separate command: `vhs publish demo.gif`
2.  When running a tape file: `vhs demo.tape --publish`

When a GIF is published, VHS will provide several ways to share it:

*   Direct URL
*   Markdown embed code
*   HTML embed code with optional badge

Sources: [publish.go27-56](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go#L27-L56)
 [publish.go136-146](https://github.com/charmbracelet/vhs/blob/f28239f3/publish.go#L136-L146)
 [README.md176-186](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L176-L186)

SSH Server
----------

VHS includes a built-in SSH server that allows remote execution of tape files:

To start the server:

    vhs serve
    

The server can be configured with environment variables:

| Variable | Description | Default |
| --- | --- | --- |
| `VHS_PORT` | Port to listen on | `1976` |
| `VHS_HOST` | Host to listen on | `localhost` |
| `VHS_GID` | Group ID to run as | Current user's GID |
| `VHS_UID` | User ID to run as | Current user's UID |
| `VHS_KEY_PATH` | Path to SSH key | `.ssh/vhs_ed25519` |
| `VHS_AUTHORIZED_KEYS_PATH` | Path to authorized keys file | Empty (public access) |

To use VHS remotely:

    ssh vhs.example.com < demo.tape > demo.gif
    

Sources: [serve.go38-165](https://github.com/charmbracelet/vhs/blob/f28239f3/serve.go#L38-L165)
 [README.md187-216](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L187-L216)

Other CLI Commands
------------------

### Themes Command

Lists all available themes for terminal styling:

    vhs themes
    

The command outputs a list of theme names that can be used in tape files with the `Set Theme` command.

Sources: [main.go150-170](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L150-L170)

### Validate Command

Checks the syntax of tape files without executing them:

    vhs validate file1.tape file2.tape
    

This is useful for ensuring your tape files are valid before running them, especially in CI/CD pipelines.

Sources: [main.go205-240](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L205-L240)

### New Command

Creates a new tape file with example contents:

    vhs new example
    

This generates a file named `example.tape` with commented documentation and a simple example to help you get started.

Sources: [main.go181-203](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L181-L203)
 [examples/demo.tape1-70](https://github.com/charmbracelet/vhs/blob/f28239f3/examples/demo.tape#L1-L70)

Command Line Interface Architecture
-----------------------------------

The VHS CLI is built using the Cobra command framework, with a hierarchical command structure:

Sources: [main.go42-291](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L42-L291)

Integration with External Tools
-------------------------------

VHS relies on external dependencies to function correctly:

VHS checks for these dependencies when starting up and will provide helpful error messages if any are missing.

Sources: [main.go293-331](https://github.com/charmbracelet/vhs/blob/f28239f3/main.go#L293-L331)
 [README.md75-79](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L75-L79)

Common Workflows
----------------

Here are the most common workflows for using VHS:

1.  **Create and run a tape file**:
    
        vhs new demo.tape
        # Edit demo.tape
        vhs demo.tape
        
    
2.  **Record terminal actions**:
    
        vhs record > demo.tape
        # Perform actions, then exit
        vhs demo.tape
        
    
3.  **Share a GIF**:
    
        vhs demo.tape
        vhs publish demo.gif
        
    
4.  **Set up a VHS server**:
    
        vhs serve
        # From another machine:
        ssh your-server.example.com < demo.tape > output.gif
        
    

Each workflow leverages different aspects of the VHS CLI to create, record, and share terminal GIFs.

Sources: [README.md18-66](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L18-L66)
 [README.md158-215](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L158-L215)

Conclusion
----------

The VHS CLI provides a comprehensive set of features for creating, executing, and sharing terminal recordings. By combining the various commands and options, you can create professional-looking demonstrations of command-line tools, document workflows, or create engaging content for documentation.

For more detailed information about the tape file format, see [Tape Format](https://deepwiki.com/charmbracelet/vhs/2-tape-format)
, and to understand the underlying architecture, refer to [System Architecture](https://deepwiki.com/charmbracelet/vhs/3-system-architecture)
.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [CLI Features](https://deepwiki.com/charmbracelet/vhs/4-cli-features#cli-features)
    
*   [Command Overview](https://deepwiki.com/charmbracelet/vhs/4-cli-features#command-overview)
    
*   [Root Command: Running Tape Files](https://deepwiki.com/charmbracelet/vhs/4-cli-features#root-command-running-tape-files)
    
*   [Recording Tapes](https://deepwiki.com/charmbracelet/vhs/4-cli-features#recording-tapes)
    
*   [Publishing GIFs](https://deepwiki.com/charmbracelet/vhs/4-cli-features#publishing-gifs)
    
*   [SSH Server](https://deepwiki.com/charmbracelet/vhs/4-cli-features#ssh-server)
    
*   [Other CLI Commands](https://deepwiki.com/charmbracelet/vhs/4-cli-features#other-cli-commands)
    
*   [Themes Command](https://deepwiki.com/charmbracelet/vhs/4-cli-features#themes-command)
    
*   [Validate Command](https://deepwiki.com/charmbracelet/vhs/4-cli-features#validate-command)
    
*   [New Command](https://deepwiki.com/charmbracelet/vhs/4-cli-features#new-command)
    
*   [Command Line Interface Architecture](https://deepwiki.com/charmbracelet/vhs/4-cli-features#command-line-interface-architecture)
    
*   [Integration with External Tools](https://deepwiki.com/charmbracelet/vhs/4-cli-features#integration-with-external-tools)
    
*   [Common Workflows](https://deepwiki.com/charmbracelet/vhs/4-cli-features#common-workflows)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/vhs/4-cli-features#conclusion)
