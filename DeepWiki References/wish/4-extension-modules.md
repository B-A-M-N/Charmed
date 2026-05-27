Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/wish](https://github.com/charmbracelet/wish "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 20 April 2025 ([19f432](https://github.com/charmbracelet/wish/commits/19f43236)
)

*   [Introduction to Wish](https://deepwiki.com/charmbracelet/wish/1-introduction-to-wish)
    
*   [Core SSH Server Framework](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework)
    
*   [Server Configuration and Options](https://deepwiki.com/charmbracelet/wish/2.1-server-configuration-and-options)
    
*   [Command Execution](https://deepwiki.com/charmbracelet/wish/2.2-command-execution)
    
*   [Middleware System](https://deepwiki.com/charmbracelet/wish/3-middleware-system)
    
*   [Access Control Middleware](https://deepwiki.com/charmbracelet/wish/3.1-access-control-middleware)
    
*   [Logging and Metrics Middleware](https://deepwiki.com/charmbracelet/wish/3.2-logging-and-metrics-middleware)
    
*   [Rate Limiting Middleware](https://deepwiki.com/charmbracelet/wish/3.3-rate-limiting-middleware)
    
*   [Extension Modules](https://deepwiki.com/charmbracelet/wish/4-extension-modules)
    
*   [Git Server Integration](https://deepwiki.com/charmbracelet/wish/4.1-git-server-integration)
    
*   [BubbleTea Integration](https://deepwiki.com/charmbracelet/wish/4.2-bubbletea-integration)
    
*   [File Transfer (SCP/SFTP)](https://deepwiki.com/charmbracelet/wish/4.3-file-transfer-(scpsftp))
    
*   [Authentication](https://deepwiki.com/charmbracelet/wish/5-authentication)
    
*   [Public Key Authentication](https://deepwiki.com/charmbracelet/wish/5.1-public-key-authentication)
    
*   [Other Authentication Methods](https://deepwiki.com/charmbracelet/wish/5.2-other-authentication-methods)
    
*   [Usage Examples](https://deepwiki.com/charmbracelet/wish/6-usage-examples)
    
*   [Simple SSH Server](https://deepwiki.com/charmbracelet/wish/6.1-simple-ssh-server)
    
*   [BubbleTea Examples](https://deepwiki.com/charmbracelet/wish/6.2-bubbletea-examples)
    
*   [Git Server Examples](https://deepwiki.com/charmbracelet/wish/6.3-git-server-examples)
    
*   [Development and Contribution](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution)
    
*   [Project Structure](https://deepwiki.com/charmbracelet/wish/7.1-project-structure)
    
*   [CI/CD and Releases](https://deepwiki.com/charmbracelet/wish/7.2-cicd-and-releases)
    

Menu

Extension Modules
=================

Relevant source files

*   [README.md](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1)
    
*   [bubbletea/tea.go](https://github.com/charmbracelet/wish/blob/19f43236/bubbletea/tea.go)
    
*   [git/git.go](https://github.com/charmbracelet/wish/blob/19f43236/git/git.go)
    
*   [options\_test.go](https://github.com/charmbracelet/wish/blob/19f43236/options_test.go)
    
*   [wish.go](https://github.com/charmbracelet/wish/blob/19f43236/wish.go)
    

This page provides an overview of the major functional extension modules in the Wish SSH framework. These modules leverage Wish's middleware system to add significant capabilities to SSH servers, making it easy to create feature-rich SSH applications.

For information about the underlying middleware system that powers these extensions, see [Middleware System](https://deepwiki.com/charmbracelet/wish/3-middleware-system)
.

Overview of Extension Modules
-----------------------------

Wish includes several key extension modules that provide specific functionality through the middleware pattern:

1.  **BubbleTea Integration**: Run terminal UI applications over SSH
2.  **Git Server Integration**: Host Git repositories over SSH
3.  **File Transfer (SCP/SFTP)**: Enable file transfer capabilities
4.  **Access Control**: Restrict commands and require active terminals
5.  **Logging**: Log SSH sessions and measure elapsed time

Sources: [README.md45-82](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L45-L82)
 [wish.go11-14](https://github.com/charmbracelet/wish/blob/19f43236/wish.go#L11-L14)

Extension Module Architecture
-----------------------------

Extension modules in Wish follow a middleware architecture pattern that allows composing functionality by chaining middleware components. Each middleware can process the SSH session, perform its specific function, and then pass the session to the next middleware in the chain.

Sources: [wish.go11-14](https://github.com/charmbracelet/wish/blob/19f43236/wish.go#L11-L14)
 [git/git.go67-118](https://github.com/charmbracelet/wish/blob/19f43236/git/git.go#L67-L118)
 [bubbletea/tea.go43-106](https://github.com/charmbracelet/wish/blob/19f43236/bubbletea/tea.go#L43-L106)

BubbleTea Integration
---------------------

The BubbleTea middleware enables SSH servers to host interactive terminal user interface (TUI) applications built with the Bubble Tea framework. This allows developers to create rich terminal applications that users can access remotely via SSH.

### How It Works

The BubbleTea middleware captures SSH session input/output and window resize events, and hooks them into the Bubble Tea program. This creates a seamless experience where the Bubble Tea application feels like it's running locally while actually being served over SSH.

### Key Features

*   **Handler Interface**: Provides a simple way to hook Bubble Tea applications into SSH
*   **Terminal Capabilities**: Automatically handles terminal capabilities and color profiles
*   **Window Resizing**: Captures terminal resize events and sends them to the Bubble Tea application
*   **Program Lifecycle**: Manages the program lifecycle, starting it when a session begins and terminating it when a session ends

### Usage Example

Sources: [bubbletea/tea.go1-169](https://github.com/charmbracelet/wish/blob/19f43236/bubbletea/tea.go#L1-L169)
 [README.md53-60](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L53-L60)

Git Server Integration
----------------------

The Git middleware turns your SSH server into a Git server, allowing users to push, pull, and clone repositories over SSH. It handles Git commands, repository access control, and provides hooks for custom behavior.

### How It Works

When a Git client connects over SSH to a server with the Git middleware, the middleware intercepts Git-specific commands (`git-receive-pack`, `git-upload-pack`, and `git-upload-archive`), validates access permissions, and executes the appropriate Git operation.

### Key Features

*   **Access Control**: Define custom authorization levels (NoAccess, ReadOnlyAccess, ReadWriteAccess, AdminAccess)
*   **Repository Management**: Automatically creates repositories if they don't exist
*   **Operation Hooks**: Execute custom code after successful Git operations
*   **Command Handling**: Properly handles Git-specific SSH commands

### Usage Example

Sources: [git/git.go1-240](https://github.com/charmbracelet/wish/blob/19f43236/git/git.go#L1-L240)
 [README.md62-68](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L62-L68)

File Transfer (SCP/SFTP)
------------------------

The SCP middleware enables file transfer capabilities over SSH, allowing clients to securely transfer files to and from the server.

### How It Works

The SCP middleware intercepts SCP commands and provides an implementation of the SCP protocol, enabling secure file copying between SSH clients and the server.

### Key Features

*   **Secure File Transfer**: Enables secure copying of files over SSH
*   **Protocol Implementation**: Implements the SCP protocol for compatibility with standard SSH clients
*   **File System Integration**: Manages file system operations for uploads and downloads

Sources: [README.md45-82](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L45-L82)

Access Control and Logging Middlewares
--------------------------------------

In addition to the major extension modules, Wish provides utilities for access control and logging:

### Access Control

The `accesscontrol` middleware allows you to restrict which commands can be executed through SSH. It's useful for limiting the functionality available to SSH clients.

### Active Terminal

The `activeterm` middleware ensures that only connections with an active terminal are accepted. This prevents non-interactive connections when your application requires an interactive terminal.

### Logging

The `logging` middleware provides detailed logging of SSH connections, including client address, commands, and session duration.

Sources: [README.md70-82](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L70-L82)

Integration with Core Wish Framework
------------------------------------

Extension modules integrate with the core Wish framework through the middleware system. When creating a Wish SSH server, you can add one or more extension modules as middleware to enhance the server's functionality.

### Middleware Composition

Multiple extension modules can be composed together to create a feature-rich SSH server. For example, you can combine the Git middleware with the Logging middleware to create a Git server with detailed logging.

Note that middlewares are composed from first to last, which means the last one is executed first in the chain.

Sources: [wish.go11-14](https://github.com/charmbracelet/wish/blob/19f43236/wish.go#L11-L14)
 [README.md45-52](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L45-L52)

Summary
-------

Extension modules in Wish provide powerful capabilities that transform a basic SSH server into a feature-rich application platform. By leveraging the middleware architecture, these modules can be easily composed to create SSH applications that support terminal UIs, Git hosting, file transfers, and more.

These extensions make it possible to build a wide range of SSH applications, from interactive terminal games to full-featured Git hosting services, all with a clean and consistent API.

For more detailed information about specific extension modules, refer to:

*   [Git Server Integration](https://deepwiki.com/charmbracelet/wish/4.1-git-server-integration)
    
*   [BubbleTea Integration](https://deepwiki.com/charmbracelet/wish/4.2-bubbletea-integration)
    
*   [File Transfer (SCP/SFTP)](https://deepwiki.com/charmbracelet/wish/4.3-file-transfer-(scpsftp))
    

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Extension Modules](https://deepwiki.com/charmbracelet/wish/4-extension-modules#extension-modules)
    
*   [Overview of Extension Modules](https://deepwiki.com/charmbracelet/wish/4-extension-modules#overview-of-extension-modules)
    
*   [Extension Module Architecture](https://deepwiki.com/charmbracelet/wish/4-extension-modules#extension-module-architecture)
    
*   [BubbleTea Integration](https://deepwiki.com/charmbracelet/wish/4-extension-modules#bubbletea-integration)
    
*   [How It Works](https://deepwiki.com/charmbracelet/wish/4-extension-modules#how-it-works)
    
*   [Key Features](https://deepwiki.com/charmbracelet/wish/4-extension-modules#key-features)
    
*   [Usage Example](https://deepwiki.com/charmbracelet/wish/4-extension-modules#usage-example)
    
*   [Git Server Integration](https://deepwiki.com/charmbracelet/wish/4-extension-modules#git-server-integration)
    
*   [How It Works](https://deepwiki.com/charmbracelet/wish/4-extension-modules#how-it-works-1)
    
*   [Key Features](https://deepwiki.com/charmbracelet/wish/4-extension-modules#key-features-1)
    
*   [Usage Example](https://deepwiki.com/charmbracelet/wish/4-extension-modules#usage-example-1)
    
*   [File Transfer (SCP/SFTP)](https://deepwiki.com/charmbracelet/wish/4-extension-modules#file-transfer-scpsftp)
    
*   [How It Works](https://deepwiki.com/charmbracelet/wish/4-extension-modules#how-it-works-2)
    
*   [Key Features](https://deepwiki.com/charmbracelet/wish/4-extension-modules#key-features-2)
    
*   [Access Control and Logging Middlewares](https://deepwiki.com/charmbracelet/wish/4-extension-modules#access-control-and-logging-middlewares)
    
*   [Access Control](https://deepwiki.com/charmbracelet/wish/4-extension-modules#access-control)
    
*   [Active Terminal](https://deepwiki.com/charmbracelet/wish/4-extension-modules#active-terminal)
    
*   [Logging](https://deepwiki.com/charmbracelet/wish/4-extension-modules#logging)
    
*   [Integration with Core Wish Framework](https://deepwiki.com/charmbracelet/wish/4-extension-modules#integration-with-core-wish-framework)
    
*   [Middleware Composition](https://deepwiki.com/charmbracelet/wish/4-extension-modules#middleware-composition)
    
*   [Summary](https://deepwiki.com/charmbracelet/wish/4-extension-modules#summary)
