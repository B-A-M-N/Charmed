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

Core SSH Server Framework
=========================

Relevant source files

*   [README.md](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1)
    
*   [cmd.go](https://github.com/charmbracelet/wish/blob/19f43236/cmd.go)
    
*   [cmd\_test.go](https://github.com/charmbracelet/wish/blob/19f43236/cmd_test.go)
    
*   [cmd\_unix.go](https://github.com/charmbracelet/wish/blob/19f43236/cmd_unix.go)
    
*   [cmd\_windows.go](https://github.com/charmbracelet/wish/blob/19f43236/cmd_windows.go)
    
*   [options.go](https://github.com/charmbracelet/wish/blob/19f43236/options.go)
    
*   [options\_test.go](https://github.com/charmbracelet/wish/blob/19f43236/options_test.go)
    

The Core SSH Server Framework provides the foundation for building SSH applications using the Wish library. It handles server creation, configuration, authentication, and command execution. This page focuses on the core components that make up the SSH server infrastructure. For information about middleware components, see [Middleware System](https://deepwiki.com/charmbracelet/wish/3-middleware-system)
, and for extension modules like Git and BubbleTea integration, see [Extension Modules](https://deepwiki.com/charmbracelet/wish/4-extension-modules)
.

Overview
--------

The Core SSH Server Framework consists of three primary components:

Sources: [options.go](https://github.com/charmbracelet/wish/blob/19f43236/options.go)
 [cmd.go](https://github.com/charmbracelet/wish/blob/19f43236/cmd.go)
 [README.md](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1)

Server Creation and Configuration
---------------------------------

Wish builds on top of [gliderlabs/ssh](https://github.com/charmbracelet/wish/blob/19f43236/gliderlabs/ssh)
 to provide an SSH server implementation with sensible defaults. The server creation and configuration process follows a functional options pattern, allowing for flexible and clear configuration.

Sources: [README.md24-26](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L24-L26)

Server Options System
---------------------

The options system defined in `options.go` provides a comprehensive set of configuration options for the SSH server. Each option is implemented as a function that returns an `ssh.Option` that modifies the server configuration.

### Basic Configuration Options

| Option | Description | Location |
| --- | --- | --- |
| `WithAddress(addr string)` | Sets the address to listen on | [options.go18-24](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L18-L24) |
| `WithVersion(version string)` | Sets the server version string | [options.go26-32](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L26-L32) |
| `WithHostKeyPath(path string)` | Sets host key file path (generates if doesn't exist) | [options.go68-78](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L68-L78) |
| `WithHostKeyPEM(pem []byte)` | Sets the host key from a PEM block | [options.go80-83](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L80-L83) |

### Authentication Options

| Option | Description | Location |
| --- | --- | --- |
| `WithPublicKeyAuth(h ssh.PublicKeyHandler)` | Sets public key authentication handler | [options.go173-175](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L173-L175) |
| `WithPasswordAuth(p ssh.PasswordHandler)` | Sets password authentication handler | [options.go178-180](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L178-L180) |
| `WithKeyboardInteractiveAuth(h)` | Sets keyboard-interactive authentication handler | [options.go183-185](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L183-L185) |
| `WithAuthorizedKeys(path string)` | Uses authorized\_keys file for authentication | [options.go86-97](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L86-L97) |
| `WithTrustedUserCAKeys(path string)` | Authorizes certificates signed by a CA | [options.go102-134](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L102-L134) |

### Session Management Options

| Option | Description | Location |
| --- | --- | --- |
| `WithBanner(banner string)` | Sets a static banner message | [options.go34-40](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L34-L40) |
| `WithBannerHandler(h ssh.BannerHandler)` | Sets a dynamic banner handler | [options.go43-49](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L43-L49) |
| `WithIdleTimeout(d time.Duration)` | Sets the connection's idle timeout | [options.go188-193](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L188-L193) |
| `WithMaxTimeout(d time.Duration)` | Sets the connection's absolute timeout | [options.go196-201](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L196-L201) |
| `WithSubsystem(key string, h ssh.SubsystemHandler)` | Registers a subsystem handler | [options.go205-213](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L205-L213) |

### Middleware Configuration

| Option | Description | Location |
| --- | --- | --- |
| `WithMiddleware(mw ...Middleware)` | Composes middleware functions to process SSH sessions | [options.go56-65](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L56-L65) |

Sources: [options.go18-213](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L18-L213)

Command Execution System
------------------------

The command execution system provides utilities for running commands in SSH sessions, handling both PTY (terminal) and non-PTY sessions correctly.

### Key Components

*   `Command(s ssh.Session, name string, args ...string)`: Creates a command using the session's context
*   `CommandContext(ctx context.Context, s ssh.Session, name string, args ...string)`: Creates a command with a specific context
*   `Cmd` struct: Wraps `exec.Cmd` and `ssh.Session` to manage command execution

The command execution process handles PTY and non-PTY sessions differently:

1.  **PTY Sessions**: Command I/O is connected to the PTY, which handles terminal emulation
2.  **Non-PTY Sessions**: Command I/O is connected directly to the SSH session

Platform-specific implementations exist for handling PTY sessions:

*   Unix platforms (Linux, macOS, etc.): [cmd\_unix.go8-13](https://github.com/charmbracelet/wish/blob/19f43236/cmd_unix.go#L8-L13)
    
*   Windows: [cmd\_windows.go13-29](https://github.com/charmbracelet/wish/blob/19f43236/cmd_windows.go#L13-L29)
    

Sources: [cmd.go12-71](https://github.com/charmbracelet/wish/blob/19f43236/cmd.go#L12-L71)
 [cmd\_unix.go8-13](https://github.com/charmbracelet/wish/blob/19f43236/cmd_unix.go#L8-L13)
 [cmd\_windows.go13-29](https://github.com/charmbracelet/wish/blob/19f43236/cmd_windows.go#L13-L29)

Middleware Integration
----------------------

The core SSH server framework is designed to integrate with middleware through the `WithMiddleware` function. This function composes middleware functions to create a processing chain for SSH sessions.

The middlewares are composed from first to last, which means the last one in the chain is executed first in the processing flow.

Sources: [options.go56-65](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L56-L65)
 [README.md44-51](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L44-L51)

Usage Examples
--------------

### Creating a Basic SSH Server

### Using Command Execution

For more detailed examples, see [Usage Examples](https://deepwiki.com/charmbracelet/wish/6-usage-examples)
.

Sources: [cmd\_test.go14-144](https://github.com/charmbracelet/wish/blob/19f43236/cmd_test.go#L14-L144)

Summary
-------

The Core SSH Server Framework provides a solid foundation for building SSH applications with Wish. It offers:

1.  Flexible server configuration through a functional options pattern
2.  Comprehensive authentication methods (public key, password, etc.)
3.  Command execution utilities that handle both PTY and non-PTY sessions
4.  Middleware integration for extending functionality

This framework is designed to be easy to use while providing the flexibility needed for a wide range of SSH applications. See [Server Configuration and Options](https://deepwiki.com/charmbracelet/wish/2.1-server-configuration-and-options)
 for more details on configuring servers and [Command Execution](https://deepwiki.com/charmbracelet/wish/2.2-command-execution)
 for more information about running commands within SSH sessions.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core SSH Server Framework](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#core-ssh-server-framework)
    
*   [Overview](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#overview)
    
*   [Server Creation and Configuration](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#server-creation-and-configuration)
    
*   [Server Options System](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#server-options-system)
    
*   [Basic Configuration Options](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#basic-configuration-options)
    
*   [Authentication Options](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#authentication-options)
    
*   [Session Management Options](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#session-management-options)
    
*   [Middleware Configuration](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#middleware-configuration)
    
*   [Command Execution System](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#command-execution-system)
    
*   [Key Components](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#key-components)
    
*   [Middleware Integration](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#middleware-integration)
    
*   [Usage Examples](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#usage-examples)
    
*   [Creating a Basic SSH Server](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#creating-a-basic-ssh-server)
    
*   [Using Command Execution](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#using-command-execution)
    
*   [Summary](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework#summary)
