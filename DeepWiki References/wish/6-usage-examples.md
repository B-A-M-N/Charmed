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

Usage Examples
==============

Relevant source files

*   [examples/README.md](https://github.com/charmbracelet/wish/blob/19f43236/examples/README.md?plain=1)
    
*   [examples/banner/banner.txt](https://github.com/charmbracelet/wish/blob/19f43236/examples/banner/banner.txt)
    
*   [examples/banner/main.go](https://github.com/charmbracelet/wish/blob/19f43236/examples/banner/main.go)
    
*   [examples/forward/main.go](https://github.com/charmbracelet/wish/blob/19f43236/examples/forward/main.go)
    
*   [examples/graceful-shutdown/main.go](https://github.com/charmbracelet/wish/blob/19f43236/examples/graceful-shutdown/main.go)
    
*   [examples/simple/main.go](https://github.com/charmbracelet/wish/blob/19f43236/examples/simple/main.go)
    

This page provides concrete examples of how to use the Wish SSH server framework. It covers essential usage patterns, from creating basic SSH servers to implementing advanced features like middleware chains, graceful shutdowns, and integrations with other components. For specific integrations, see [BubbleTea Examples](https://deepwiki.com/charmbracelet/wish/6.2-bubbletea-examples)
 and [Git Server Examples](https://deepwiki.com/charmbracelet/wish/6.3-git-server-examples)
.

Basic SSH Server
----------------

The simplest use of Wish is creating a basic SSH server. This involves configuring server options and setting up a minimal middleware chain.

The minimal server setup requires:

1.  Creating a server with `wish.NewServer()`
2.  Configuring address with `wish.WithAddress()`
3.  Setting up host keys with `wish.WithHostKeyPath()`
4.  Adding at least one middleware
5.  Starting the server with `ListenAndServe()`

Here's an example from the codebase:

Sources: [examples/simple/main.go18-51](https://github.com/charmbracelet/wish/blob/19f43236/examples/simple/main.go#L18-L51)

Graceful Shutdown Pattern
-------------------------

For production applications, implementing graceful shutdown ensures active SSH sessions are properly terminated:

This pattern involves:

1.  Setting up a signal notification channel
2.  Starting the server in a goroutine
3.  Waiting for termination signals
4.  Using a timeout context for controlled shutdown
5.  Calling `srv.Shutdown(ctx)` to gracefully stop the server

Sources: [examples/graceful-shutdown/main.go23-73](https://github.com/charmbracelet/wish/blob/19f43236/examples/graceful-shutdown/main.go#L23-L73)

Server Banners and Custom Middleware
------------------------------------

SSH servers often display welcome banners and implement custom middleware to enhance functionality:

Banner handlers are executed before authentication, while middleware runs after successful authentication:

Sources: [examples/banner/main.go31-52](https://github.com/charmbracelet/wish/blob/19f43236/examples/banner/main.go#L31-L52)
 [examples/banner/banner.txt1-15](https://github.com/charmbracelet/wish/blob/19f43236/examples/banner/banner.txt#L1-L15)

Advanced Authentication
-----------------------

Wish supports multiple authentication methods, including:

*   Public key authentication
*   Password authentication
*   Keyboard-interactive authentication

Example configuration for password authentication:

Sources: [examples/banner/main.go38-40](https://github.com/charmbracelet/wish/blob/19f43236/examples/banner/main.go#L38-L40)

Reverse Port Forwarding
-----------------------

Wish supports SSH tunneling capabilities, including reverse port forwarding:

This example shows how to configure reverse port forwarding:

Sources: [examples/forward/main.go26-42](https://github.com/charmbracelet/wish/blob/19f43236/examples/forward/main.go#L26-L42)

Common Example Categories
-------------------------

The Wish codebase includes examples in these categories:

| Category | Description | Examples |
| --- | --- | --- |
| **Basics** | Fundamental server setup | Simple server, Graceful shutdown, Banners, Authentication |
| **SSH Apps** | Interactive applications | Bubble Tea integration, Cobra command apps, Multichat |
| **File Systems** | Storage and SCM | Git repository servers, SCP/SFTP file transfer |
| **Terminal Handling** | PTY allocation | PTY allocation, Program execution |

Sources: [examples/README.md1-30](https://github.com/charmbracelet/wish/blob/19f43236/examples/README.md?plain=1#L1-L30)

Implementation Pattern
----------------------

This diagram illustrates the typical implementation pattern for Wish applications:

Sources: [examples/simple/main.go18-51](https://github.com/charmbracelet/wish/blob/19f43236/examples/simple/main.go#L18-L51)
 [examples/graceful-shutdown/main.go23-73](https://github.com/charmbracelet/wish/blob/19f43236/examples/graceful-shutdown/main.go#L23-L73)
 [examples/banner/main.go31-52](https://github.com/charmbracelet/wish/blob/19f43236/examples/banner/main.go#L31-L52)
 [examples/forward/main.go26-42](https://github.com/charmbracelet/wish/blob/19f43236/examples/forward/main.go#L26-L42)

Typical Project Organization
----------------------------

When building SSH applications with Wish, a typical project organization includes:

The examples in the repository demonstrate these patterns, providing a foundation for building custom SSH applications.

Sources: [examples/simple/main.go18-51](https://github.com/charmbracelet/wish/blob/19f43236/examples/simple/main.go#L18-L51)
 [examples/README.md1-30](https://github.com/charmbracelet/wish/blob/19f43236/examples/README.md?plain=1#L1-L30)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Usage Examples](https://deepwiki.com/charmbracelet/wish/6-usage-examples#usage-examples)
    
*   [Basic SSH Server](https://deepwiki.com/charmbracelet/wish/6-usage-examples#basic-ssh-server)
    
*   [Graceful Shutdown Pattern](https://deepwiki.com/charmbracelet/wish/6-usage-examples#graceful-shutdown-pattern)
    
*   [Server Banners and Custom Middleware](https://deepwiki.com/charmbracelet/wish/6-usage-examples#server-banners-and-custom-middleware)
    
*   [Advanced Authentication](https://deepwiki.com/charmbracelet/wish/6-usage-examples#advanced-authentication)
    
*   [Reverse Port Forwarding](https://deepwiki.com/charmbracelet/wish/6-usage-examples#reverse-port-forwarding)
    
*   [Common Example Categories](https://deepwiki.com/charmbracelet/wish/6-usage-examples#common-example-categories)
    
*   [Implementation Pattern](https://deepwiki.com/charmbracelet/wish/6-usage-examples#implementation-pattern)
    
*   [Typical Project Organization](https://deepwiki.com/charmbracelet/wish/6-usage-examples#typical-project-organization)
