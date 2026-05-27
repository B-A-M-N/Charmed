Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/gum](https://github.com/charmbracelet/gum "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 2 May 2026 ([30cc51](https://github.com/charmbracelet/gum/commits/30cc5169)
)

*   [Overview](https://deepwiki.com/charmbracelet/gum/1-overview)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/gum/2-getting-started)
    
*   [Architecture](https://deepwiki.com/charmbracelet/gum/3-architecture)
    
*   [Application Entry Point](https://deepwiki.com/charmbracelet/gum/3.1-application-entry-point)
    
*   [Framework Components](https://deepwiki.com/charmbracelet/gum/3.2-framework-components)
    
*   [Command Patterns](https://deepwiki.com/charmbracelet/gum/3.3-command-patterns)
    
*   [Interactive Commands](https://deepwiki.com/charmbracelet/gum/4-interactive-commands)
    
*   [Choose Command](https://deepwiki.com/charmbracelet/gum/4.1-choose-command)
    
*   [Confirm Command](https://deepwiki.com/charmbracelet/gum/4.2-confirm-command)
    
*   [Filter Command](https://deepwiki.com/charmbracelet/gum/4.3-filter-command)
    
*   [Input Command](https://deepwiki.com/charmbracelet/gum/4.4-input-command)
    
*   [Write Command](https://deepwiki.com/charmbracelet/gum/4.5-write-command)
    
*   [File Command](https://deepwiki.com/charmbracelet/gum/4.6-file-command)
    
*   [Table Command](https://deepwiki.com/charmbracelet/gum/4.7-table-command)
    
*   [Pager Command](https://deepwiki.com/charmbracelet/gum/4.8-pager-command)
    
*   [Progress and Background Commands](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands)
    
*   [Spin Command](https://deepwiki.com/charmbracelet/gum/5.1-spin-command)
    
*   [Formatting Commands](https://deepwiki.com/charmbracelet/gum/6-formatting-commands)
    
*   [Style Command](https://deepwiki.com/charmbracelet/gum/6.1-style-command)
    
*   [Format Command](https://deepwiki.com/charmbracelet/gum/6.2-format-command)
    
*   [Join Command](https://deepwiki.com/charmbracelet/gum/6.3-join-command)
    
*   [Log Command](https://deepwiki.com/charmbracelet/gum/6.4-log-command)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/gum/7-development-guide)
    
*   [Building and Testing](https://deepwiki.com/charmbracelet/gum/7.1-building-and-testing)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/gum/7.2-dependency-management)
    
*   [CI/CD Pipeline](https://deepwiki.com/charmbracelet/gum/7.3-cicd-pipeline)
    
*   [Shell Completion](https://deepwiki.com/charmbracelet/gum/7.4-shell-completion)
    
*   [Internal Utilities](https://deepwiki.com/charmbracelet/gum/8-internal-utilities)
    
*   [Glossary](https://deepwiki.com/charmbracelet/gum/9-glossary)
    

Menu

Getting Started
===============

Relevant source files

*   [.gitignore](https://github.com/charmbracelet/gum/blob/30cc5169/.gitignore)
    
*   [Dockerfile](https://github.com/charmbracelet/gum/blob/30cc5169/Dockerfile)
    
*   [README.md](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1)
    
*   [default.nix](https://github.com/charmbracelet/gum/blob/30cc5169/default.nix)
    
*   [examples/commit.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.sh)
    
*   [examples/commit.tape](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.tape)
    
*   [examples/demo.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh)
    
*   [examples/demo.tape](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.tape)
    
*   [examples/test.sh](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh)
    
*   [flake.lock](https://github.com/charmbracelet/gum/blob/30cc5169/flake.lock)
    
*   [flake.nix](https://github.com/charmbracelet/gum/blob/30cc5169/flake.nix)
    

This document provides installation instructions, basic usage examples, and a quick tour of gum's capabilities for new users. It covers the essential steps to install gum, understand its command execution model, and begin using it in shell scripts.

For architectural details about how gum is structured internally, see [Architecture](https://deepwiki.com/charmbracelet/gum/3-architecture)
. For comprehensive documentation of individual commands, see [Interactive Commands](https://deepwiki.com/charmbracelet/gum/4-interactive-commands)
, [Progress and Background Commands](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands)
, and [Formatting Commands](https://deepwiki.com/charmbracelet/gum/6-formatting-commands)
.

Installation
------------

Gum can be installed through multiple package managers, direct binary downloads, or built from source using Go.

### Package Manager Installation

**macOS and Linux:**

**Arch Linux:**

**Debian/Ubuntu:**

**Fedora/RHEL/EPEL 10:**

**Windows:**

**Go Installation:**

Sources: [README.md60-141](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L60-L141)
 [default.nix3-12](https://github.com/charmbracelet/gum/blob/30cc5169/default.nix#L3-L12)

### Verifying Installation

After installation, verify that gum is available:

This will display the installed version of gum, which is managed via `ldflags` during the build process.

Sources: [default.nix11](https://github.com/charmbracelet/gum/blob/30cc5169/default.nix#L11-L11)

Command Execution Model
-----------------------

Understanding how gum commands execute is essential for effective shell scripting integration.

**Key Execution Characteristics:**

| Aspect | Behavior | Example |
| --- | --- | --- |
| **Command Parsing** | Kong parser routes to command-specific `Run()` methods | `gum choose` → `choose.Run()` |
| **Input Sources** | stdin, CLI flags, environment variables | `echo "option" \| gum choose --height 10` |
| **UI Rendering** | Interactive commands render to stderr | Preserves stdout for data output |
| **Data Output** | Selected/processed data written to stdout | `gum input > result.txt` |
| **Exit Codes** | 0 for success, non-zero for errors | `gum confirm && echo "yes" \| echo "no"` |
| **Configuration Priority** | CLI flags > Environment variables > Defaults | `--height 20` overrides `GUM_CHOOSE_HEIGHT=10` |

Sources: [README.md145-190](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L145-L190)
 [examples/test.sh1-53](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L1-L53)

Available Commands
------------------

Gum provides 13 commands organized into functional categories:

| Category | Commands | Purpose |
| --- | --- | --- |
| **Interactive** | `choose`, `confirm`, `filter`, `input`, `write`, `file`, `table`, `pager` | Present TUI interfaces for user selection and input |
| **Progress** | `spin` | Display animated indicators while running background processes |
| **Formatting** | `style`, `format`, `join`, `log` | Transform and style text output |

### Command Overview Table

| Command | Primary Use Case | Output Location | Bubbletea-based |
| --- | --- | --- | --- |
| `choose` | Select from a list | stdout | Yes |
| `confirm` | Yes/No prompts | exit code | Yes |
| `filter` | Fuzzy search lists | stdout | Yes |
| `input` | Single-line text entry | stdout | Yes |
| `write` | Multi-line text entry | stdout | Yes |
| `file` | File/directory picker | stdout | Yes |
| `table` | Tabular data selection | stdout | Yes |
| `pager` | Scroll long content | display only | Yes |
| `spin` | Background task indicator | stdout (from subprocess) | Yes |
| `style` | Apply visual styling | stdout | No  |
| `format` | Render markdown/templates | stdout | No  |
| `join` | Combine text layouts | stdout | No  |
| `log` | Structured logging | stdout | No  |

Sources: [README.md147-159](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L147-L159)
 [examples/test.sh3-53](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L3-L53)

First Steps
-----------

### Example 1: Simple Input Capture

The most basic gum operation is capturing user input:

**Data Flow:**

1.  `gum input` renders prompt to stderr (visible to user).
2.  User types input.
3.  Result written to stdout.
4.  Shell captures stdout in `$NAME` variable.

Sources: [examples/demo.sh4](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh#L4-L4)
 [README.md198-199](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L198-L199)
 [examples/test.sh25-27](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L25-L27)

### Example 2: Multiple Choice Selection

Use `choose` for presenting options:

Sources: [examples/demo.sh16-17](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh#L16-L17)
 [README.md240-242](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L240-L242)
 [examples/test.sh4-7](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L4-L7)

### Example 3: Conditional Logic with confirm

The `confirm` command uses exit codes for conditional execution:

Sources: [README.md53](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L53-L53)
 [examples/commit.sh28](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.sh#L28-L28)
 [examples/test.sh10-11](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L10-L11)

Demo Script Walkthrough
-----------------------

The included demo script at `examples/demo.sh` demonstrates a complete interactive workflow.

**Script Breakdown:**

| Lines | Command | Purpose | Key Flags |
| --- | --- | --- | --- |
| [3](https://github.com/charmbracelet/gum/blob/30cc5169/3) | `style` | Welcome banner with border | `--border normal --margin "1" --padding "1 2"` |
| [4](https://github.com/charmbracelet/gum/blob/30cc5169/4) | `input` | Capture user name | `--placeholder "What is your name?"` |
| [12](https://github.com/charmbracelet/gum/blob/30cc5169/12) | `write` | Multi-line secret input | `--placeholder` (output to /dev/null) |
| [17](https://github.com/charmbracelet/gum/blob/30cc5169/17) | `choose` | Select multiple actions | `--no-limit` |
| [21-23](https://github.com/charmbracelet/gum/blob/30cc5169/21-23) | `spin` | Show progress for actions | `-s line/pulse/monkey --title` |
| [27](https://github.com/charmbracelet/gum/blob/30cc5169/27) | `filter` | Fuzzy search flavors | `--placeholder "Favorite flavor?"` |
| [35](https://github.com/charmbracelet/gum/blob/30cc5169/35) | `choose` | Single choice selection | `--item.foreground 250` |
| [45-47](https://github.com/charmbracelet/gum/blob/30cc5169/45-47) | `style` + `join` | Layout two styled boxes | `--border double --horizontal` |

Sources: [examples/demo.sh1-48](https://github.com/charmbracelet/gum/blob/30cc5169/examples/demo.sh#L1-L48)

Configuration Patterns
----------------------

Gum commands support two configuration methods that can be combined or used independently.

### Command-Line Flags

Flags provide direct, explicit configuration:

**Nested Options:** Many commands support nested option structures using dot notation for specific UI elements like `cursor`, `prompt`, or `selected` items.

Sources: [README.md170-176](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L170-L176)
 [examples/test.sh5-7](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L5-L7)

### Environment Variables

Environment variables follow the pattern `GUM_COMMAND_OPTION`:

Sources: [README.md181-188](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L181-L188)

### Configuration Priority

Sources: [README.md163-188](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L163-L188)

Practical Integration Examples
------------------------------

### Git Commit Message Helper

The `examples/commit.sh` script demonstrates a complete workflow for creating conventional commit messages:

Sources: [examples/commit.sh17-28](https://github.com/charmbracelet/gum/blob/30cc5169/examples/commit.sh#L17-L28)
 [README.md28-53](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L28-L53)

### Shell Pipelines

Gum excels at being part of a pipeline.

| Pattern | Example |
| --- | --- |
| **Fuzzy Selection** | `echo {1..500} \| sed 's/ /\n/g' \| gum filter` |
| **File Picker** | `$EDITOR $(gum file)` |
| **Table Display** | `gum table < table/example.csv` |
| **Pager** | `gum pager < README.md` |

Sources: [examples/test.sh15](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L15-L15)
 [examples/test.sh46](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L46-L46)
 [examples/test.sh49](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L49-L49)
 [examples/test.sh52](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L52-L52)

Common Patterns
---------------

Sources: [README.md219-222](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L219-L222)
 [README.md198-199](https://github.com/charmbracelet/gum/blob/30cc5169/README.md?plain=1#L198-L199)
 [examples/test.sh15](https://github.com/charmbracelet/gum/blob/30cc5169/examples/test.sh#L15-L15)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Getting Started](https://deepwiki.com/charmbracelet/gum/2-getting-started#getting-started)
    
*   [Installation](https://deepwiki.com/charmbracelet/gum/2-getting-started#installation)
    
*   [Package Manager Installation](https://deepwiki.com/charmbracelet/gum/2-getting-started#package-manager-installation)
    
*   [Verifying Installation](https://deepwiki.com/charmbracelet/gum/2-getting-started#verifying-installation)
    
*   [Command Execution Model](https://deepwiki.com/charmbracelet/gum/2-getting-started#command-execution-model)
    
*   [Available Commands](https://deepwiki.com/charmbracelet/gum/2-getting-started#available-commands)
    
*   [Command Overview Table](https://deepwiki.com/charmbracelet/gum/2-getting-started#command-overview-table)
    
*   [First Steps](https://deepwiki.com/charmbracelet/gum/2-getting-started#first-steps)
    
*   [Example 1: Simple Input Capture](https://deepwiki.com/charmbracelet/gum/2-getting-started#example-1-simple-input-capture)
    
*   [Example 2: Multiple Choice Selection](https://deepwiki.com/charmbracelet/gum/2-getting-started#example-2-multiple-choice-selection)
    
*   [Example 3: Conditional Logic with confirm](https://deepwiki.com/charmbracelet/gum/2-getting-started#example-3-conditional-logic-with-confirm)
    
*   [Demo Script Walkthrough](https://deepwiki.com/charmbracelet/gum/2-getting-started#demo-script-walkthrough)
    
*   [Configuration Patterns](https://deepwiki.com/charmbracelet/gum/2-getting-started#configuration-patterns)
    
*   [Command-Line Flags](https://deepwiki.com/charmbracelet/gum/2-getting-started#command-line-flags)
    
*   [Environment Variables](https://deepwiki.com/charmbracelet/gum/2-getting-started#environment-variables)
    
*   [Configuration Priority](https://deepwiki.com/charmbracelet/gum/2-getting-started#configuration-priority)
    
*   [Practical Integration Examples](https://deepwiki.com/charmbracelet/gum/2-getting-started#practical-integration-examples)
    
*   [Git Commit Message Helper](https://deepwiki.com/charmbracelet/gum/2-getting-started#git-commit-message-helper)
    
*   [Shell Pipelines](https://deepwiki.com/charmbracelet/gum/2-getting-started#shell-pipelines)
    
*   [Common Patterns](https://deepwiki.com/charmbracelet/gum/2-getting-started#common-patterns)
