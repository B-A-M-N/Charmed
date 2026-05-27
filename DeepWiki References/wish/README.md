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

Introduction to Wish
====================

Relevant source files

*   [README.md](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1)
    
*   [go.mod](https://github.com/charmbracelet/wish/blob/19f43236/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/wish/blob/19f43236/go.sum)
    
*   [options\_test.go](https://github.com/charmbracelet/wish/blob/19f43236/options_test.go)
    

Wish is a Go library that provides a framework for building SSH applications with minimal effort. This document introduces the core concepts of Wish, its architecture, and explains how to use it to create powerful, secure network applications accessible via SSH.

For detailed information on server configuration options, see [Server Configuration and Options](https://deepwiki.com/charmbracelet/wish/2.1-server-configuration-and-options)
. For information about specific extension modules like Git server integration, see [Extension Modules](https://deepwiki.com/charmbracelet/wish/4-extension-modules)
.

What is Wish?
-------------

Wish is an SSH server framework with sensible defaults and a collection of middlewares that makes building SSH applications straightforward. Unlike traditional SSH servers that provide shell access, Wish allows developers to create purpose-built applications accessible over SSH.

Sources: [README.md1-43](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L1-L43)

Why SSH Applications?
---------------------

SSH provides several advantages as an application protocol:

1.  **Built-in Security**: Secure communications without the complexity of HTTPS certificates
2.  **User Identification**: Authentication via SSH keys
3.  **Universal Access**: Available from any terminal
4.  **Established Protocol**: Robust and widely supported

Unlike traditional SSH servers that provide shell access, Wish enables developers to build custom applications that use SSH as a transport layer, including terminal user interfaces (TUIs), Git servers, and more.

Sources: [README.md14-42](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L14-L42)

Core Architecture
-----------------

Wish is built on top of the [gliderlabs/ssh](https://github.com/charmbracelet/wish/blob/19f43236/gliderlabs/ssh)
 library, which provides the underlying SSH server implementation. Wish extends this with a middleware system and configuration options that simplify building SSH applications.

Sources: [README.md44-52](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L44-L52)
 [options\_test.go16-91](https://github.com/charmbracelet/wish/blob/19f43236/options_test.go#L16-L91)

Middleware System
-----------------

The Wish middleware system is a core concept that allows for extending the SSH server's functionality. Middlewares are composed together to handle SSH sessions, with each middleware having the opportunity to process a session before passing it to the next.

Middlewares are composed from first to last in the chain, which means the last middleware specified is executed first in the chain. This behavior is similar to many HTTP middleware systems.

Sources: [README.md44-52](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L44-L52)

Key Middleware Components
-------------------------

Wish includes several built-in middlewares for common SSH application patterns:

1.  **BubbleTea Middleware**: Serves [Bubble Tea](https://github.com/charmbracelet/wish/blob/19f43236/Bubble%20Tea)
     terminal UI applications over SSH
2.  **Git Middleware**: Provides Git server functionality, supporting repository operations over SSH
3.  **Logging Middleware**: Logs connection information and session duration
4.  **Access Control Middleware**: Restricts access to allowed commands
5.  **ActiveTerm Middleware**: Only allows connections with active terminals

Sources: [README.md53-82](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L53-L82)

Basic Server Example
--------------------

Creating a basic Wish server involves initializing the server, configuring it with options, and adding middlewares as needed:

Server Configuration Options
----------------------------

Wish provides a rich set of options to configure the SSH server:

| Option Category | Examples | Purpose |
| --- | --- | --- |
| Basic Configuration | WithAddress, WithHostKeyPath, WithVersion | Configure server address, host key, and SSH version |
| Authentication | WithPublicKeyAuth, WithPasswordAuth, WithAuthorizedKeys | Configure authentication methods |
| Session Management | WithBanner, WithIdleTimeout, WithMaxTimeout | Configure session properties |
| Middleware | WithMiddleware | Add middleware functionality |

Sources: [options\_test.go16-91](https://github.com/charmbracelet/wish/blob/19f43236/options_test.go#L16-L91)
 [README.md84-89](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L84-L89)

Applications Built with Wish
----------------------------

Wish powers several popular SSH applications:

*   **Soft Serve**: A self-hosted Git server for the command line
*   **Wishlist**: A SSH directory server that lists and launches other SSH apps
*   **Various games and utilities**: Like SSHWordle, clidle, and others

These applications demonstrate the versatility of Wish for creating interactive SSH experiences.

Sources: [README.md94-101](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L94-L101)

Conclusion
----------

Wish simplifies the process of building SSH applications by providing a robust framework with sensible defaults and powerful extension capabilities. By leveraging SSH as an application protocol, developers can create secure, accessible applications that work from any terminal.

To learn more about specific components of Wish, refer to the other sections of this documentation:

*   [Core SSH Server Framework](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework)
    
*   [Middleware System](https://deepwiki.com/charmbracelet/wish/3-middleware-system)
    
*   [Extension Modules](https://deepwiki.com/charmbracelet/wish/4-extension-modules)
    
*   [Authentication](https://deepwiki.com/charmbracelet/wish/5-authentication)
    
*   [Usage Examples](https://deepwiki.com/charmbracelet/wish/6-usage-examples)
    

Sources: [README.md1-124](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L1-L124)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Introduction to Wish](https://deepwiki.com/charmbracelet/wish#introduction-to-wish)
    
*   [What is Wish?](https://deepwiki.com/charmbracelet/wish#what-is-wish)
    
*   [Why SSH Applications?](https://deepwiki.com/charmbracelet/wish#why-ssh-applications)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/wish#core-architecture)
    
*   [Middleware System](https://deepwiki.com/charmbracelet/wish#middleware-system)
    
*   [Key Middleware Components](https://deepwiki.com/charmbracelet/wish#key-middleware-components)
    
*   [Basic Server Example](https://deepwiki.com/charmbracelet/wish#basic-server-example)
    
*   [Server Configuration Options](https://deepwiki.com/charmbracelet/wish#server-configuration-options)
    
*   [Applications Built with Wish](https://deepwiki.com/charmbracelet/wish#applications-built-with-wish)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/wish#conclusion)
