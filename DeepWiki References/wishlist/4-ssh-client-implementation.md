Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/wishlist](https://github.com/charmbracelet/wishlist "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 23 April 2025 ([5fe34a](https://github.com/charmbracelet/wishlist/commits/5fe34a29)
)

*   [Introduction to Wishlist](https://deepwiki.com/charmbracelet/wishlist/1-introduction-to-wishlist)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/wishlist/1.1-getting-started)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview)
    
*   [Operation Modes](https://deepwiki.com/charmbracelet/wishlist/2.1-operation-modes)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system)
    
*   [SSH Config Integration](https://deepwiki.com/charmbracelet/wishlist/3.1-ssh-config-integration)
    
*   [Endpoint Configuration](https://deepwiki.com/charmbracelet/wishlist/3.2-endpoint-configuration)
    
*   [SSH Client Implementation](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation)
    
*   [Authentication Methods](https://deepwiki.com/charmbracelet/wishlist/4.1-authentication-methods)
    
*   [Proxy Jump Support](https://deepwiki.com/charmbracelet/wishlist/4.2-proxy-jump-support)
    
*   [SSH Server Implementation](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation)
    
*   [Middleware System](https://deepwiki.com/charmbracelet/wishlist/5.1-middleware-system)
    
*   [Terminal User Interface](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface)
    
*   [Endpoint Discovery](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery)
    
*   [Utility Components](https://deepwiki.com/charmbracelet/wishlist/8-utility-components)
    
*   [Deployment](https://deepwiki.com/charmbracelet/wishlist/9-deployment)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/wishlist/10-development-guide)
    

Menu

SSH Client Implementation
=========================

Relevant source files

*   [client.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client.go)
    
*   [client\_local.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_local.go)
    
*   [client\_remote.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_remote.go)
    

This document describes the SSH client implementation in the Wishlist project. The SSH client is responsible for establishing connections to remote SSH servers, managing sessions, handling terminal interactions, and executing commands. This page focuses specifically on the client-side SSH implementation; for details about authentication methods, see [Authentication Methods](https://deepwiki.com/charmbracelet/wishlist/4.1-authentication-methods)
 and for proxy jump support details, see [Proxy Jump Support](https://deepwiki.com/charmbracelet/wishlist/4.2-proxy-jump-support)
.

Overview of SSH Client Architecture
-----------------------------------

The SSH client in Wishlist is designed with flexibility in mind, supporting both local and remote operation modes. It provides a consistent interface for connecting to endpoints while handling different execution environments.

Sources: [client.go15-18](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client.go#L15-L18)
 [client\_local.go19-27](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_local.go#L19-L27)
 [client\_remote.go14-31](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_remote.go#L14-L31)

Client Interface
----------------

The core of the SSH client implementation is the `SSHClient` interface, which provides a simple abstraction for creating executable commands for SSH endpoints:

The interface defines a single method, `For`, which accepts an `Endpoint` and returns a `tea.ExecCommand`. This command can then be executed to establish an SSH connection to the specified endpoint.

Sources: [client.go15-18](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client.go#L15-L18)

Client Implementations
----------------------

Wishlist provides two implementations of the `SSHClient` interface, each designed for different operational contexts:

### Local Client

The `localClient` is designed for use when Wishlist is run directly on a user's machine. It creates SSH connections from the local machine to remote endpoints.

Key features of the local client:

*   Uses the local user's credentials and SSH agent
*   Can access the local environment variables
*   Manages terminal settings for the local terminal
*   Handles window size changes

Sources: [client\_local.go19-147](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_local.go#L19-L147)

### Remote Client

The `remoteClient` is designed for when Wishlist is running as an SSH server and users are connecting to it. In this case, the client creates "nested" SSH connections from the Wishlist server to other endpoints.

Key features of the remote client:

*   Forwards the user's session from the Wishlist server to the endpoint
*   Handles authentication in a remote context
*   Manages PTY forwarding
*   Forwards window size changes from the user's session to the endpoint

Sources: [client\_remote.go14-130](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_remote.go#L14-L130)

Session Creation and Management
-------------------------------

Both client implementations rely on common functionality for session creation and management. The `createSession` function is responsible for:

1.  Establishing an SSH connection to the endpoint
2.  Setting up proxy jumps if specified
3.  Creating an SSH session
4.  Configuring environment variables

After a session is created, depending on the endpoint configuration, the client will either:

*   Start an interactive shell with `shellAndWait`
*   Execute a specific command with `runAndWait`

The system also includes mechanisms for proper resource cleanup through the `closers` type, which aggregates cleanup functions and executes them when needed.

Sources: [client.go20-73](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client.go#L20-L73)
 [client.go75-96](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client.go#L75-L96)
 [client.go98-110](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client.go#L98-L110)

Terminal Handling
-----------------

A critical aspect of SSH client implementation is managing the terminal (PTY) interface. Wishlist implements:

1.  Terminal size detection and forwarding
2.  Window size change notifications
3.  Raw terminal mode setting
4.  Terminal reset functionality

For local clients, terminal handling involves:

*   Getting the current terminal size
*   Setting the terminal to raw mode
*   Monitoring window size changes

For remote clients, it involves:

*   Forwarding PTY information from the parent session
*   Monitoring window change events from the parent session
*   Notifying the endpoint of window changes

The `notifyWindowChanges` function in both client implementations monitors for window size changes and forwards them to the remote server, ensuring that the terminal display remains properly formatted.

Sources: [client\_local.go112-146](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_local.go#L112-L146)
 [client\_remote.go92-106](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_remote.go#L92-L106)
 [client\_remote.go114-130](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_remote.go#L114-L130)
 [client.go112-116](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client.go#L112-L116)

Integration with Wishlist System
--------------------------------

The SSH client implementation integrates with the broader Wishlist system in several ways:

1.  It uses endpoint configurations from the configuration system
2.  It leverages authentication methods from the authentication system
3.  It integrates with the Terminal UI for displaying connection status
4.  It supports the dual operation modes of Wishlist (local CLI and SSH server)

Sources: [client.go15-18](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client.go#L15-L18)
 [client\_local.go26-32](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_local.go#L26-L32)
 [client\_remote.go24-31](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_remote.go#L24-L31)

Summary
-------

The SSH client implementation in Wishlist provides a flexible, feature-rich way to establish and manage SSH connections to remote endpoints. By abstracting the connection process behind a simple interface and providing implementations for both local and remote usage, the system enables consistent SSH functionality regardless of how Wishlist is being used.

Key takeaways:

*   The `SSHClient` interface provides a unified abstraction for SSH connections
*   Two implementations (`localClient` and `remoteClient`) handle different operational contexts
*   Common session creation and management code ensures consistent behavior
*   Terminal handling ensures proper interactive sessions
*   Resource management ensures clean connection termination

The SSH client implementation works in tandem with other Wishlist components to provide a seamless SSH connection management experience.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [SSH Client Implementation](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#ssh-client-implementation)
    
*   [Overview of SSH Client Architecture](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#overview-of-ssh-client-architecture)
    
*   [Client Interface](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#client-interface)
    
*   [Client Implementations](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#client-implementations)
    
*   [Local Client](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#local-client)
    
*   [Remote Client](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#remote-client)
    
*   [Session Creation and Management](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#session-creation-and-management)
    
*   [Terminal Handling](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#terminal-handling)
    
*   [Integration with Wishlist System](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#integration-with-wishlist-system)
    
*   [Summary](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation#summary)
