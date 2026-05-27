Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/soft-serve](https://github.com/charmbracelet/soft-serve "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 18 April 2025 ([767bdd](https://github.com/charmbracelet/soft-serve/commits/767bdd5a)
)

*   [Overview](https://deepwiki.com/charmbracelet/soft-serve/1-overview)
    
*   [What is Soft Serve?](https://deepwiki.com/charmbracelet/soft-serve/1.1-what-is-soft-serve)
    
*   [Key Features](https://deepwiki.com/charmbracelet/soft-serve/1.2-key-features)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/soft-serve/1.3-system-architecture)
    
*   [Installation and Configuration](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration)
    
*   [Installation Methods](https://deepwiki.com/charmbracelet/soft-serve/2.1-installation-methods)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/soft-serve/2.2-configuration-options)
    
*   [System Service Setup](https://deepwiki.com/charmbracelet/soft-serve/2.3-system-service-setup)
    
*   [Core Components](https://deepwiki.com/charmbracelet/soft-serve/3-core-components)
    
*   [CLI Tool Structure](https://deepwiki.com/charmbracelet/soft-serve/3.1-cli-tool-structure)
    
*   [Server Components](https://deepwiki.com/charmbracelet/soft-serve/3.2-server-components)
    
*   [SSH Server](https://deepwiki.com/charmbracelet/soft-serve/3.2.1-ssh-server)
    
*   [Git Daemon](https://deepwiki.com/charmbracelet/soft-serve/3.2.2-git-daemon)
    
*   [HTTP Server](https://deepwiki.com/charmbracelet/soft-serve/3.2.3-http-server)
    
*   [Repository Management](https://deepwiki.com/charmbracelet/soft-serve/3.3-repository-management)
    
*   [Git Hooks and Webhooks](https://deepwiki.com/charmbracelet/soft-serve/3.4-git-hooks-and-webhooks)
    
*   [User Interfaces](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces)
    
*   [Terminal UI Architecture](https://deepwiki.com/charmbracelet/soft-serve/4.1-terminal-ui-architecture)
    
*   [UI Components](https://deepwiki.com/charmbracelet/soft-serve/4.2-ui-components)
    
*   [Command Line Interface](https://deepwiki.com/charmbracelet/soft-serve/4.3-command-line-interface)
    
*   [Developer Guide](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide)
    
*   [Project Dependencies](https://deepwiki.com/charmbracelet/soft-serve/5.1-project-dependencies)
    
*   [Build System](https://deepwiki.com/charmbracelet/soft-serve/5.2-build-system)
    
*   [Testing Framework](https://deepwiki.com/charmbracelet/soft-serve/5.3-testing-framework)
    
*   [Code Quality Tools](https://deepwiki.com/charmbracelet/soft-serve/5.4-code-quality-tools)
    
*   [Release Process](https://deepwiki.com/charmbracelet/soft-serve/5.5-release-process)
    

Menu

User Interfaces
===============

Relevant source files

*   [.gitignore](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.gitignore)
    
*   [cmd/soft/browse/browse.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/browse/browse.go)
    
*   [cmd/soft/main.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go)
    
*   [pkg/ssh/ui.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/ui.go)
    
*   [pkg/ui/components/code/code.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/components/code/code.go)
    
*   [pkg/ui/pages/repo/log.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/log.go)
    
*   [pkg/ui/pages/repo/repo.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/repo.go)
    
*   [pkg/ui/pages/repo/stash.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/stash.go)
    

This page documents the user interfaces of Soft Serve, explaining the different ways users can interact with the system. It covers the Terminal User Interface (TUI), Command Line Interface (CLI), and Git protocol interactions. For information about server components that power these interfaces, see [Server Components](https://deepwiki.com/charmbracelet/soft-serve/3.2-server-components)
.

Overview of User Interfaces
---------------------------

Soft Serve provides multiple interfaces for users to interact with the system:

1.  **Terminal User Interface (TUI)** - An interactive interface accessed via SSH
2.  **Command Line Interface (CLI)** - The `soft` command and its subcommands
3.  **Git Protocol Interfaces** - Standard git client interactions via git and HTTP protocols

The following diagram illustrates how users interact with the various interfaces:

Sources: [cmd/soft/main.go1-142](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go#L1-L142)
 [pkg/ssh/ui.go1-338](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/ui.go#L1-L338)

Terminal User Interface (TUI)
-----------------------------

The Terminal User Interface is the primary way users browse repositories when connecting via SSH. It provides an interactive, text-based interface for navigating repositories, viewing code, commit history, and more.

### TUI Architecture

The TUI is built using the Bubble Tea framework, which provides a Model-View-Update architecture:

Sources: [pkg/ssh/ui.go20-337](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/ui.go#L20-L337)
 [pkg/ui/pages/repo/repo.go1-431](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/repo.go#L1-L431)

### TUI Page Flow

When a user connects via SSH, they navigate through different pages and tabs:

Sources: [pkg/ssh/ui.go20-70](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/ui.go#L20-L70)
 [pkg/ui/pages/repo/repo.go21-26](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/repo.go#L21-L26)
 [pkg/ui/pages/repo/log.go25-31](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/log.go#L25-L31)

### Repository View Components

The repository view in the TUI is organized into tabs, each displaying different aspects of a repository:

| Tab Name | Purpose | Code Component |
| --- | --- | --- |
| Readme | Shows repository README | `repo.Readme` |
| Files | Browses repository files | `repo.Files` |
| Commits | Shows commit history | `repo.Log` |
| Refs | Displays branches and tags | `repo.Refs` |
| Stash | Shows stashed changes | `repo.Stash` |

Each tab is implemented as a `TabComponent` that is managed by the main `Repo` component. The `Repo` component handles tab switching, common keyboard shortcuts, and overall layout.

Sources: [pkg/ui/pages/repo/repo.go136-146](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/repo.go#L136-L146)
 [pkg/ssh/ui.go136-144](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/ui.go#L136-L144)

### TUI Message Flow

The TUI uses a message-based architecture for handling user interactions and UI updates:

Sources: [pkg/ui/pages/repo/repo.go148-290](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/repo.go#L148-L290)
 [pkg/ssh/ui.go168-261](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/ui.go#L168-L261)

Command Line Interface (CLI)
----------------------------

The Command Line Interface allows users to interact with Soft Serve through the `soft` command. This interface is primarily used for administration, browsing local repositories, and server management.

### CLI Command Structure

The CLI is structured as a set of subcommands using the Cobra framework:

Sources: [cmd/soft/main.go39-65](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go#L39-L65)
 [cmd/soft/main.go73-79](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go#L73-L79)

### CLI Browse Command

The `browse` command provides a TUI for browsing a local repository. It shares many components with the SSH TUI but is specialized for browsing local repositories:

Sources: [cmd/soft/browse/browse.go20-304](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/browse/browse.go#L20-L304)

UI Components
-------------

Soft Serve uses a component-based architecture for its UI, with reusable components that handle specific aspects of the interface.

### Code Component

The `Code` component is responsible for rendering code with syntax highlighting and other formatting:

Sources: [pkg/ui/components/code/code.go22-257](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/components/code/code.go#L22-L257)

### Repository Tabs

The repository view contains multiple tab components, each handling different aspects of a repository:

| Component | File | Purpose |
| --- | --- | --- |
| `Log` | [pkg/ui/pages/repo/log.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/log.go) | Displays commit history and diffs |
| `Stash` | [pkg/ui/pages/repo/stash.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/stash.go) | Shows stashed changes |
| `Files` | Not shown in excerpts | Browses repository files |
| `Readme` | Not shown in excerpts | Displays repository README |
| `Refs` | Not shown in excerpts | Shows branches and tags |

Each tab component implements the `TabComponent` interface:

Sources: [pkg/ui/pages/repo/log.go45-536](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/log.go#L45-L536)
 [pkg/ui/pages/repo/stash.go33-284](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/stash.go#L33-L284)

User Interface Interactions
---------------------------

Users interact with Soft Serve through different interfaces depending on their needs:

### SSH Terminal UI Flow

When a user connects via SSH, they follow this interaction flow:

1.  Connect to the Soft Serve server using SSH
2.  View the repository selection page
3.  Select a repository to browse
4.  Navigate through tabs (README, Files, Commits, etc.)
5.  Interact with repository content
6.  Return to repository selection or disconnect

### CLI Browsing Flow

When using the `soft browse` command:

1.  Run `soft browse [PATH]` to browse a local repository
2.  Navigate through repository tabs and content
3.  Exit the browser to return to the command line

### Git Protocol Flow

When interacting through standard git commands:

1.  Use standard git commands (clone, pull, push) with the Soft Serve server as remote
2.  The Soft Serve server handles the git protocol interactions
3.  Repository content is transferred according to standard git protocols

Key Components Implementation
-----------------------------

The Soft Serve UI system is built on several key components:

| Component | Purpose | Implementation |
| --- | --- | --- |
| UI Model | Main SSH UI controller | `UI` struct in [pkg/ssh/ui.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/ui.go) |
| Repository View | Repository browsing UI | `Repo` struct in [pkg/ui/pages/repo/repo.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/repo.go) |
| Tab Components | Individual repository views | `Log`, `Stash`, etc. |
| Bubble Tea | UI framework | Used throughout for Model-View-Update pattern |
| CLI Commands | Command-line interface | Cobra commands in [cmd/soft/main.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go) |

The UI system uses a message-passing architecture where user actions generate messages that flow through the component hierarchy, triggering updates to the UI state and view.

Sources: [pkg/ssh/ui.go1-338](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/ui.go#L1-L338)
 [pkg/ui/pages/repo/repo.go1-431](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ui/pages/repo/repo.go#L1-L431)
 [cmd/soft/main.go1-142](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go#L1-L142)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [User Interfaces](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#user-interfaces)
    
*   [Overview of User Interfaces](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#overview-of-user-interfaces)
    
*   [Terminal User Interface (TUI)](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#terminal-user-interface-tui)
    
*   [TUI Architecture](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#tui-architecture)
    
*   [TUI Page Flow](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#tui-page-flow)
    
*   [Repository View Components](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#repository-view-components)
    
*   [TUI Message Flow](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#tui-message-flow)
    
*   [Command Line Interface (CLI)](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#command-line-interface-cli)
    
*   [CLI Command Structure](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#cli-command-structure)
    
*   [CLI Browse Command](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#cli-browse-command)
    
*   [UI Components](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#ui-components)
    
*   [Code Component](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#code-component)
    
*   [Repository Tabs](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#repository-tabs)
    
*   [User Interface Interactions](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#user-interface-interactions)
    
*   [SSH Terminal UI Flow](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#ssh-terminal-ui-flow)
    
*   [CLI Browsing Flow](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#cli-browsing-flow)
    
*   [Git Protocol Flow](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#git-protocol-flow)
    
*   [Key Components Implementation](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces#key-components-implementation)
