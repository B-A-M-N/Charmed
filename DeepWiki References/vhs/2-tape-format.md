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

Tape Format
===========

Relevant source files

*   [README.md](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1)
    
*   [error.go](https://github.com/charmbracelet/vhs/blob/f28239f3/error.go)
    
*   [examples/commands/alt.tape](https://github.com/charmbracelet/vhs/blob/f28239f3/examples/commands/alt.tape)
    
*   [examples/demo.tape](https://github.com/charmbracelet/vhs/blob/f28239f3/examples/demo.tape)
    
*   [lexer/lexer.go](https://github.com/charmbracelet/vhs/blob/f28239f3/lexer/lexer.go)
    
*   [lexer/lexer\_test.go](https://github.com/charmbracelet/vhs/blob/f28239f3/lexer/lexer_test.go)
    
*   [parser/parser.go](https://github.com/charmbracelet/vhs/blob/f28239f3/parser/parser.go)
    
*   [parser/parser\_test.go](https://github.com/charmbracelet/vhs/blob/f28239f3/parser/parser_test.go)
    
*   [record.go](https://github.com/charmbracelet/vhs/blob/f28239f3/record.go)
    
*   [record\_test.go](https://github.com/charmbracelet/vhs/blob/f28239f3/record_test.go)
    
*   [syntax.go](https://github.com/charmbracelet/vhs/blob/f28239f3/syntax.go)
    
*   [token/token.go](https://github.com/charmbracelet/vhs/blob/f28239f3/token/token.go)
    

The Tape Format is the core declarative scripting language used by VHS to define terminal recordings. Tape files (`.tape` extension) contain a sequence of commands that instruct VHS how to set up the terminal environment, what to type or display, how to interact with the terminal, and how to render the output. This document explains the structure and syntax of tape files.

For information about executing the specific commands available in tape files, see [Commands Reference](https://deepwiki.com/charmbracelet/vhs/2.1-commands-reference)
. For information about configuring output formats, see [Output Options](https://deepwiki.com/charmbracelet/vhs/2.2-output-options)
.

Overview of Tape Files
----------------------

A tape file is a plain text file that contains a series of commands, typically one per line. Each command instructs VHS to perform a specific action, such as typing text, pressing a key, or changing a setting. The commands are executed in sequence to produce the desired terminal recording.

Sources: [README.md32-34](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L32-L34)
 [examples/demo.tape1-57](https://github.com/charmbracelet/vhs/blob/f28239f3/examples/demo.tape#L1-L57)

Tape File Anatomy
-----------------

A typical tape file is organized into logical sections:

1.  **Header Comments**: Optional comments explaining the purpose of the tape
2.  **Output Definitions**: Specify the output file format and location
3.  **Requirements**: Declare dependencies needed for the tape to run
4.  **Settings**: Configure the terminal appearance and behavior
5.  **Actions**: The main sequence of commands that interact with the terminal
6.  **Control Flow**: Commands that affect how the recording progresses

Here's an example of a simple tape file structure:

    # My demo recording
    Output demo.gif
    
    Require echo
    
    Set FontSize 32
    Set Width 1200
    Set Height 600
    
    Type "echo 'Hello, world!'" 
    Sleep 500ms
    Enter
    
    Sleep 5s
    

Sources: [examples/demo.tape1-70](https://github.com/charmbracelet/vhs/blob/f28239f3/examples/demo.tape#L1-L70)

Lexical Structure
-----------------

Tape files have a simple lexical structure. Each command starts with a command name followed by arguments and options. The lexer processes the tape file and converts it into tokens that the parser can understand.

Sources: [lexer/lexer.go1-246](https://github.com/charmbracelet/vhs/blob/f28239f3/lexer/lexer.go#L1-L246)
 [parser/parser.go1-789](https://github.com/charmbracelet/vhs/blob/f28239f3/parser/parser.go#L1-L789)

### Comments

Comments in tape files start with `#` and continue until the end of the line. They are ignored during execution.

    # This is a comment
    

Sources: [lexer/lexer.go115-128](https://github.com/charmbracelet/vhs/blob/f28239f3/lexer/lexer.go#L115-L128)

### Tokens and Keywords

The tape format has a set of predefined tokens and keywords that correspond to commands, settings, and control structures. The lexer recognizes these tokens and converts them to the appropriate internal representation.

Key token types include:

| Token Type | Description | Examples |
| --- | --- | --- |
| Commands | Actions to perform | `Type`, `Sleep`, `Enter` |
| Settings | Configuration options | `FontSize`, `Width`, `Height` |
| Control | Flow control commands | `Hide`, `Show`, `Wait` |
| Literals | String and numeric values | `"text"`, `42`, `500ms` |
| Operators | Special characters | `@`, `+` |

Sources: [token/token.go1-226](https://github.com/charmbracelet/vhs/blob/f28239f3/token/token.go#L1-L226)
 [lexer/lexer\_test.go1-339](https://github.com/charmbracelet/vhs/blob/f28239f3/lexer/lexer_test.go#L1-L339)

Command Syntax
--------------

Commands in tape files follow a consistent syntax pattern, though there are variations depending on the command type.

### Basic Command Structure

The general structure of a command is:

    CommandName[@<speed>] [arguments]
    

*   `CommandName`: The name of the command (e.g., `Type`, `Sleep`, `Enter`)
*   `@<speed>`: Optional speed modifier that affects how quickly the command executes
*   `arguments`: Command-specific arguments and options

Sources: [parser/parser.go64-78](https://github.com/charmbracelet/vhs/blob/f28239f3/parser/parser.go#L64-L78)
 [README.md218-238](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L218-L238)

### Command Categories

Commands in tape files can be broadly categorized into several types:

1.  **Setup Commands**: Configure the environment
    
    *   `Output`: Define output file
    *   `Set`: Configure terminal settings
    *   `Require`: Specify required dependencies
2.  **Input Commands**: Simulate user input
    
    *   `Type`: Type text
    *   Key commands: `Enter`, `Backspace`, `Space`, etc.
    *   Modifier keys: `Ctrl+`, `Alt+`, `Shift+`
3.  **Timing Commands**: Control timing and flow
    
    *   `Sleep`: Pause for a specified duration
    *   `Wait`: Wait for specific conditions
4.  **Display Commands**: Control what is recorded
    
    *   `Hide`: Stop recording
    *   `Show`: Resume recording
    *   `Screenshot`: Capture current frame
5.  **Utility Commands**: Additional functionality
    
    *   `Source`: Include commands from another tape
    *   `Copy`/`Paste`: Clipboard operations
    *   `Env`: Set environment variables

Sources: [README.md218-238](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L218-L238)
 [parser/parser.go28-58](https://github.com/charmbracelet/vhs/blob/f28239f3/parser/parser.go#L28-L58)

Parsing and Execution
---------------------

When VHS processes a tape file, it goes through several stages:

Sources: [parser/parser.go115-132](https://github.com/charmbracelet/vhs/blob/f28239f3/parser/parser.go#L115-L132)
 [README.md32-56](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L32-L56)

### Lexical Analysis

The lexer reads the tape file character by character and produces a stream of tokens. It handles:

1.  Command identification
2.  String literal parsing
3.  Number and time value parsing
4.  Special character recognition

Sources: [lexer/lexer.go30-102](https://github.com/charmbracelet/vhs/blob/f28239f3/lexer/lexer.go#L30-L102)
 [token/token.go107-166](https://github.com/charmbracelet/vhs/blob/f28239f3/token/token.go#L107-L166)

### Parsing

The parser takes the token stream and builds command objects. It:

1.  Validates command syntax
2.  Interprets arguments and options
3.  Reports syntax errors
4.  Organizes commands in execution order

Sources: [parser/parser.go115-132](https://github.com/charmbracelet/vhs/blob/f28239f3/parser/parser.go#L115-L132)
 [error.go1-58](https://github.com/charmbracelet/vhs/blob/f28239f3/error.go#L1-L58)

### Command Execution

Once parsed, commands are executed in sequence:

1.  Setup commands initialize the environment
2.  Action commands interact with the terminal
3.  Control commands manage the recording process
4.  The resulting frames are captured and converted to the specified output format

Sources: [README.md65-71](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L65-L71)

Recording Mode
--------------

VHS can also generate tape files automatically by recording your terminal actions. This is useful for creating complex demos without manually writing tape files.

The recording process:

1.  Starts a pseudo-terminal for user interaction
2.  Records all key presses and terminal output
3.  Converts input into tape commands
4.  Generates a tape file that can be edited and reused

Sources: [record.go67-123](https://github.com/charmbracelet/vhs/blob/f28239f3/record.go#L67-L123)
 [README.md158-174](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L158-L174)

Example Tape File
-----------------

Here's an annotated example of a tape file:

    # Output file definition
    Output demo.gif
    
    # Required dependencies
    Require echo
    
    # Terminal setup
    Set Shell "bash"
    Set FontSize 32
    Set Width 1200
    Set Height 600
    Set Theme "Catppuccin Mocha"
    
    # Main actions
    Type "echo 'Welcome to VHS!'"
    Sleep 500ms
    Enter
    
    # Wait to see the output
    Sleep 5s
    

This tape file:

1.  Defines the output as a GIF file
2.  Requires the `echo` command to be available
3.  Configures the terminal appearance
4.  Types a command, waits, and presses Enter
5.  Waits for 5 seconds to show the output

Sources: [examples/demo.tape58-69](https://github.com/charmbracelet/vhs/blob/f28239f3/examples/demo.tape#L58-L69)
 [README.md37-55](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L37-L55)

Command Syntax Reference
------------------------

Here's a quick reference for common command syntaxes:

| Command Pattern | Description | Example |
| --- | --- | --- |
| `Output <path>` | Define output file | `Output demo.gif` |
| `Require <program>` | Specify dependency | `Require ffmpeg` |
| `Set <setting> <value>` | Configure setting | `Set FontSize 20` |
| `Type <text>` | Type text | `Type "ls -la"` |
| `Type@<speed> <text>` | Type with custom speed | `Type@50ms "slow typing"` |
| `<Key> [count]` | Press key (optional count) | `Enter 2` |
| `<Key>@<speed> [count]` | Press key with timing | `Tab@100ms 3` |
| `Sleep <time>` | Pause execution | `Sleep 1s` |
| `Wait[+Screen][+Line] /<regex>/` | Wait for match | `Wait /Done/` |
| `Hide/Show` | Control recording | `Hide` |
| `Screenshot <path>` | Take screenshot | `Screenshot frame.png` |

Sources: [README.md240-336](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L240-L336)
 [parser/parser.go135-186](https://github.com/charmbracelet/vhs/blob/f28239f3/parser/parser.go#L135-L186)

Time and Typing Speed
---------------------

Time values in tape files can be specified in various formats:

*   Seconds: `1`, `0.5`, `.5`
*   Milliseconds: `500ms`, `100ms`
*   Minutes: `1m`

The typing speed can be set globally or per command:

    # Global typing speed
    Set TypingSpeed 50ms
    
    # Command-specific typing speed
    Type@100ms "This types more slowly"
    

Sources: [parser/parser.go527-534](https://github.com/charmbracelet/vhs/blob/f28239f3/parser/parser.go#L527-L534)
 [README.md386-406](https://github.com/charmbracelet/vhs/blob/f28239f3/README.md?plain=1#L386-L406)

Command Line Generation
-----------------------

VHS can also generate tape files from command-line recordings:

    vhs record > demo.tape
    

The recorder captures user input and converts it to tape commands, handling:

*   Key presses
*   Special keys and combinations
*   Timing and pauses
*   Command execution

Sources: [record.go130-201](https://github.com/charmbracelet/vhs/blob/f28239f3/record.go#L130-L201)
 [record\_test.go1-109](https://github.com/charmbracelet/vhs/blob/f28239f3/record_test.go#L1-L109)

Best Practices
--------------

When working with tape files:

1.  Start with setup commands (Output, Set)
2.  Group related commands together
3.  Use comments to document sections
4.  Use Hide/Show to exclude setup/cleanup steps
5.  Test tapes with small changes before complex recordings
6.  Use Wait commands for unpredictable processes

This structured approach makes tape files easier to create, maintain, and reuse.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Tape Format](https://deepwiki.com/charmbracelet/vhs/2-tape-format#tape-format)
    
*   [Overview of Tape Files](https://deepwiki.com/charmbracelet/vhs/2-tape-format#overview-of-tape-files)
    
*   [Tape File Anatomy](https://deepwiki.com/charmbracelet/vhs/2-tape-format#tape-file-anatomy)
    
*   [Lexical Structure](https://deepwiki.com/charmbracelet/vhs/2-tape-format#lexical-structure)
    
*   [Comments](https://deepwiki.com/charmbracelet/vhs/2-tape-format#comments)
    
*   [Tokens and Keywords](https://deepwiki.com/charmbracelet/vhs/2-tape-format#tokens-and-keywords)
    
*   [Command Syntax](https://deepwiki.com/charmbracelet/vhs/2-tape-format#command-syntax)
    
*   [Basic Command Structure](https://deepwiki.com/charmbracelet/vhs/2-tape-format#basic-command-structure)
    
*   [Command Categories](https://deepwiki.com/charmbracelet/vhs/2-tape-format#command-categories)
    
*   [Parsing and Execution](https://deepwiki.com/charmbracelet/vhs/2-tape-format#parsing-and-execution)
    
*   [Lexical Analysis](https://deepwiki.com/charmbracelet/vhs/2-tape-format#lexical-analysis)
    
*   [Parsing](https://deepwiki.com/charmbracelet/vhs/2-tape-format#parsing)
    
*   [Command Execution](https://deepwiki.com/charmbracelet/vhs/2-tape-format#command-execution)
    
*   [Recording Mode](https://deepwiki.com/charmbracelet/vhs/2-tape-format#recording-mode)
    
*   [Example Tape File](https://deepwiki.com/charmbracelet/vhs/2-tape-format#example-tape-file)
    
*   [Command Syntax Reference](https://deepwiki.com/charmbracelet/vhs/2-tape-format#command-syntax-reference)
    
*   [Time and Typing Speed](https://deepwiki.com/charmbracelet/vhs/2-tape-format#time-and-typing-speed)
    
*   [Command Line Generation](https://deepwiki.com/charmbracelet/vhs/2-tape-format#command-line-generation)
    
*   [Best Practices](https://deepwiki.com/charmbracelet/vhs/2-tape-format#best-practices)
