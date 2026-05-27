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

SSH Server Implementation
=========================

Relevant source files

*   [\_example/main.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/main.go)
    
*   [config\_test.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config_test.go)
    
*   [middleware.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/middleware.go)
    
*   [server.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go)
    
*   [server\_test.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server_test.go)
    

This document provides a technical overview of the SSH server implementation in the Wishlist project. It describes the architecture, configuration, connection handling, and authentication mechanisms used to create and manage SSH servers for various endpoints.

For information about the middleware system that extends SSH server functionality, see [Middleware System](https://deepwiki.com/charmbracelet/wishlist/5.1-middleware-system)
.

Architecture Overview
---------------------

The SSH server component in Wishlist enables running multiple SSH servers, each handling connections to different endpoints. The implementation builds on the [github.com/charmbracelet/wish](https://github.com/charmbracelet/wishlist/blob/5fe34a29/github.com/charmbracelet/wish)
 and [github.com/charmbracelet/ssh](https://github.com/charmbracelet/wishlist/blob/5fe34a29/github.com/charmbracelet/ssh)
 libraries.

Sources: [server.go23-95](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L23-L95)
 [server.go98-121](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L98-L121)

Server Creation and Startup
---------------------------

The server creation process begins with the `Serve` function, which sets up signal handling for graceful shutdown, configures server parameters, and creates SSH servers for each valid endpoint.

The key steps in server creation are:

1.  Setting default values for unspecified configuration parameters
2.  Creating necessary directories
3.  Setting up broadcast relay for endpoint updates
4.  Creating and starting servers for each valid endpoint
5.  Waiting for termination signal to gracefully shut down

Sources: [server.go23-95](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L23-L95)

Server Configuration
--------------------

Server configuration is managed through the `Config` struct and the `Endpoint` struct:

The configuration specifies:

*   Global settings like listen address and port
*   Endpoint definitions that define what SSH services to expose
*   User authentication information
*   Factory function for creating SSH server instances

Sources: [server.go23-95](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L23-L95)
 [config\_test.go11-132](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config_test.go#L11-L132)

Connection Handling
-------------------

The `listenAndServe` function handles creating and starting an SSH server for a specific endpoint:

1.  The server is created using the factory function from the configuration
2.  Public key authentication is configured based on user definitions
3.  A TCP listener is created for the specified address
4.  The server is started in a separate goroutine
5.  A shutdown function is returned for graceful termination

Sources: [server.go98-121](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L98-L121)

Authentication
--------------

Authentication is primarily handled through public keys. The `publicKeyAccessOption` function creates a handler that validates SSH public keys against configured users:

Key points about authentication:

1.  If no users are configured, all connections are allowed
2.  If users are configured, only those with matching username and public key are permitted
3.  Keys are parsed and compared using SSH library functions
4.  Failed authentication attempts are logged

Sources: [server.go157-182](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L157-L182)

Port Management
---------------

The implementation includes utility functions for managing ports:

*   `getFirstOpenPort`: Finds the first available port from a list of candidates
*   `toAddress`: Formats a host and port into a network address string

If no port is specified in the configuration, the system attempts to use port 22 (standard SSH port) and falls back to port 2222 if necessary.

Sources: [server.go135-155](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L135-L155)

Error Handling and Graceful Shutdown
------------------------------------

Error handling and graceful shutdown are key aspects of the server implementation:

1.  The `Serve` function sets up signal handling for SIGINT and SIGTERM
2.  When termination is requested, all servers are shut down using their respective shutdown functions
3.  The `closeAll` function aggregates errors from all server shutdowns
4.  `ssh.ErrServerClosed` errors are specifically ignored as they are expected during shutdown
5.  Other errors are aggregated using `multierror` and returned

Sources: [server.go124-132](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L124-L132)
 [server\_test.go19-43](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server_test.go#L19-L43)

Example Usage
-------------

Here's how a typical server configuration might look:

A basic implementation would:

1.  Create a configuration with endpoints and authentication settings
2.  Implement a server factory function to customize server creation
3.  Call `wishlist.Serve(config)` to start the servers

Sources: [\_example/main.go21-101](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/main.go#L21-L101)

Conclusion
----------

The SSH server implementation in Wishlist provides a flexible and robust framework for running SSH servers for multiple endpoints. It leverages the Charm SSH and Wish libraries to handle the complexities of SSH protocol implementation while adding features like endpoint configuration, authentication management, and graceful shutdown handling.

The design allows for dynamic endpoint updates through channels and supports extending server functionality through middlewares, making it adaptable to various use cases.

Sources: [server.go23-95](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L23-L95)
 [server.go98-121](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L98-L121)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [SSH Server Implementation](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#ssh-server-implementation)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#architecture-overview)
    
*   [Server Creation and Startup](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#server-creation-and-startup)
    
*   [Server Configuration](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#server-configuration)
    
*   [Connection Handling](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#connection-handling)
    
*   [Authentication](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#authentication)
    
*   [Port Management](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#port-management)
    
*   [Error Handling and Graceful Shutdown](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#error-handling-and-graceful-shutdown)
    
*   [Example Usage](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#example-usage)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation#conclusion)
