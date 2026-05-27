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

Authentication
==============

Relevant source files

*   [README.md](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1)
    
*   [options.go](https://github.com/charmbracelet/wish/blob/19f43236/options.go)
    
*   [options\_test.go](https://github.com/charmbracelet/wish/blob/19f43236/options_test.go)
    

Purpose and Scope
-----------------

This document provides an overview of the authentication mechanisms supported by the wish SSH server framework. Authentication ensures that only authorized users can access your SSH applications. For specific details about public key authentication implementation, see [Public Key Authentication](https://deepwiki.com/charmbracelet/wish/5.1-public-key-authentication)
, and for password and keyboard-interactive methods, see [Other Authentication Methods](https://deepwiki.com/charmbracelet/wish/5.2-other-authentication-methods)
.

Authentication Methods Overview
-------------------------------

Wish supports the standard SSH authentication methods, providing flexible options for implementing each:

| Authentication Method | Description | Implementation Option |
| --- | --- | --- |
| Public Key | Authenticates users based on their SSH keys | `WithPublicKeyAuth`, `WithAuthorizedKeys`, `WithTrustedUserCAKeys` |
| Password | Simple username/password authentication | `WithPasswordAuth` |
| Keyboard Interactive | Challenge-response mechanism for complex auth flows | `WithKeyboardInteractiveAuth` |

Authentication Architecture
---------------------------

The authentication system in wish is implemented through option functions that configure the underlying SSH server. These options are applied when creating a new server with `wish.NewServer()`.

### Authentication Options Hierarchy

Sources: [options.go86-97](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L86-L97)
 [options.go99-134](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L99-L134)
 [options.go173-184](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L173-L184)

### Authentication Flow

Sources: [options.go173-184](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L173-L184)

Public Key Authentication
-------------------------

Public key authentication is the most secure method and is implemented in three ways:

1.  **Custom validation** with `WithPublicKeyAuth` - Complete control over key validation logic
2.  **Authorized keys file** with `WithAuthorizedKeys` - Similar to traditional SSH servers
3.  **Certificate validation** with `WithTrustedUserCAKeys` - Support for SSH certificates signed by trusted CAs

### Authorized Keys Implementation

The `WithAuthorizedKeys` function reads keys from a standard authorized\_keys file and creates a handler to validate client keys against this list:

Sources: [options.go86-97](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L86-L97)
 [options.go136-170](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L136-L170)

### Certificate Authentication

Certificate authentication uses the `WithTrustedUserCAKeys` function to validate SSH certificates:

Sources: [options.go99-134](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L99-L134)
 [options.go136-170](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L136-L170)
 [options\_test.go127-188](https://github.com/charmbracelet/wish/blob/19f43236/options_test.go#L127-L188)

Password and Keyboard Interactive Authentication
------------------------------------------------

Both methods allow custom authentication logic through handler functions:

*   `WithPasswordAuth` - Takes a function that validates username/password pairs
*   `WithKeyboardInteractiveAuth` - Takes a function for challenge-response authentication flows

Sources: [options.go177-179](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L177-L179)
 [options.go182-184](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L182-L184)

Configuration Example
---------------------

Here's an example showing how to configure a wish server with multiple authentication methods:

Sources: [options.go86-134](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L86-L134)
 [options.go173-184](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L173-L184)

Helper Functions
----------------

### The isAuthorized Function

Both `WithAuthorizedKeys` and `WithTrustedUserCAKeys` rely on the internal `isAuthorized` function, which:

1.  Opens and reads the specified keys file
2.  Parses each line for valid keys (skipping comments and empty lines)
3.  Applies a checker function to each key to determine authorization

Sources: [options.go136-170](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L136-L170)

Conclusion
----------

The wish library provides a flexible authentication system that supports all standard SSH authentication methods. The architecture allows for easy configuration through option functions while maintaining security best practices. You can choose from built-in mechanisms like authorized\_keys files or implement custom authentication logic as needed for your SSH application.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Authentication](https://deepwiki.com/charmbracelet/wish/5-authentication#authentication)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/wish/5-authentication#purpose-and-scope)
    
*   [Authentication Methods Overview](https://deepwiki.com/charmbracelet/wish/5-authentication#authentication-methods-overview)
    
*   [Authentication Architecture](https://deepwiki.com/charmbracelet/wish/5-authentication#authentication-architecture)
    
*   [Authentication Options Hierarchy](https://deepwiki.com/charmbracelet/wish/5-authentication#authentication-options-hierarchy)
    
*   [Authentication Flow](https://deepwiki.com/charmbracelet/wish/5-authentication#authentication-flow)
    
*   [Public Key Authentication](https://deepwiki.com/charmbracelet/wish/5-authentication#public-key-authentication)
    
*   [Authorized Keys Implementation](https://deepwiki.com/charmbracelet/wish/5-authentication#authorized-keys-implementation)
    
*   [Certificate Authentication](https://deepwiki.com/charmbracelet/wish/5-authentication#certificate-authentication)
    
*   [Password and Keyboard Interactive Authentication](https://deepwiki.com/charmbracelet/wish/5-authentication#password-and-keyboard-interactive-authentication)
    
*   [Configuration Example](https://deepwiki.com/charmbracelet/wish/5-authentication#configuration-example)
    
*   [Helper Functions](https://deepwiki.com/charmbracelet/wish/5-authentication#helper-functions)
    
*   [The isAuthorized Function](https://deepwiki.com/charmbracelet/wish/5-authentication#the-isauthorized-function)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/wish/5-authentication#conclusion)
