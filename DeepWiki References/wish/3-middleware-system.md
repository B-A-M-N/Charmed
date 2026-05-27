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

Middleware System
=================

Relevant source files

*   [README.md](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1)
    
*   [options.go](https://github.com/charmbracelet/wish/blob/19f43236/options.go)
    
*   [options\_test.go](https://github.com/charmbracelet/wish/blob/19f43236/options_test.go)
    

Purpose and Scope
-----------------

This document describes the middleware system in the Wish library - a critical component that enables functional extensions to the SSH server. The middleware pattern allows developers to process SSH sessions through a chain of handlers, each adding specific functionality. For information on specific middleware implementations like BubbleTea or Git, see their respective pages ([Extension Modules](https://deepwiki.com/charmbracelet/wish/4-extension-modules)
).

Sources: [README.md44-52](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L44-L52)

What is Middleware?
-------------------

In Wish, middleware functions as a chain of handlers that process SSH sessions. Each middleware can perform operations before passing control to the next middleware, and then perform additional operations after control returns. This pattern is analogous to middleware systems found in HTTP frameworks, allowing clean separation of concerns and modular feature implementation.

Sources: [README.md44-48](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L44-L48)
 [options.go51-64](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L51-L64)

Middleware Implementation
-------------------------

The middleware system relies on function composition to create a processing chain. A middleware is a function that takes an SSH handler and returns a new handler with added functionality.

The core function that enables middleware composition is `WithMiddleware`, which assembles multiple middleware functions into a single SSH handler:

Sources: [options.go51-64](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L51-L64)

Middleware Execution Order
--------------------------

An important detail of the middleware system is the execution order. Middlewares are composed from first to last in the function call, but the last middleware in the list is the first to execute its pre-processing logic.

For example, when using `WithMiddleware(A, B, C)`:

1.  C's pre-processing executes first
2.  Then B's pre-processing
3.  Then A's pre-processing
4.  Then the final handler (if any)
5.  Then A's post-processing
6.  Then B's post-processing
7.  Finally, C's post-processing

Sources: [options.go55-56](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L55-L56)
 [README.md47-48](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L47-L48)

Using Middleware
----------------

To use middleware in your Wish SSH server, you apply the `WithMiddleware` option when creating the server. This function takes a variable number of middleware functions and composes them into a single handler.

The `WithMiddleware` function implementation looks like this:

Sources: [options.go51-64](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L51-L64)

Built-in Middlewares
--------------------

Wish provides several built-in middlewares that address common SSH server functionality:

| Middleware | Purpose | Description |
| --- | --- | --- |
| bubbletea | Terminal UI | Serves Bubble Tea applications over SSH |
| git | Git server | Adds Git repository functionality |
| logging | Connection logging | Logs connections with details and duration |
| accesscontrol | Command restriction | Restricts allowed commands |
| activeterm | Terminal requirement | Only allows connections with active terminals |

Each middleware can be imported from its respective package and added to your server configuration.

Sources: [README.md53-82](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L53-L82)

Creating Custom Middleware
--------------------------

You can create your own middleware to extend server functionality. A middleware function follows this pattern:

A custom middleware skeleton looks like this:

Sources: [options.go51-64](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L51-L64)
 [README.md44-48](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L44-L48)

Middleware Composition Example
------------------------------

Here's an example of how multiple middlewares are typically composed in a Wish server:

Remember that the order matters - the last middleware in the list will be the first to process an incoming SSH session.

Sources: [options.go51-64](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L51-L64)
 [README.md44-48](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L44-L48)

Middleware Chaining Implementation
----------------------------------

The implementation of middleware chaining in Wish is elegant and efficient. The `WithMiddleware` function builds the handler chain by starting with an empty handler and wrapping each middleware around it:

This implementation creates a nested set of function calls that correctly implements the middleware chain pattern.

Sources: [options.go56-62](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L56-L62)

Relationship to SSH Server Architecture
---------------------------------------

The middleware system is a key component in the overall SSH server architecture of Wish:

The middleware chain is established during server configuration and processes all incoming SSH sessions, making it a central part of the request handling process.

Sources: [options.go51-64](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L51-L64)
 [README.md44-52](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L44-L52)

Summary
-------

The middleware system forms the backbone of Wish's extensibility model, allowing modular and composable functionality. Understanding the middleware pattern is essential for effectively using the Wish library to build custom SSH applications, whether you're using the built-in middlewares or creating your own.

Sources: [README.md44-52](https://github.com/charmbracelet/wish/blob/19f43236/README.md?plain=1#L44-L52)
 [options.go51-64](https://github.com/charmbracelet/wish/blob/19f43236/options.go#L51-L64)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Middleware System](https://deepwiki.com/charmbracelet/wish/3-middleware-system#middleware-system)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/wish/3-middleware-system#purpose-and-scope)
    
*   [What is Middleware?](https://deepwiki.com/charmbracelet/wish/3-middleware-system#what-is-middleware)
    
*   [Middleware Implementation](https://deepwiki.com/charmbracelet/wish/3-middleware-system#middleware-implementation)
    
*   [Middleware Execution Order](https://deepwiki.com/charmbracelet/wish/3-middleware-system#middleware-execution-order)
    
*   [Using Middleware](https://deepwiki.com/charmbracelet/wish/3-middleware-system#using-middleware)
    
*   [Built-in Middlewares](https://deepwiki.com/charmbracelet/wish/3-middleware-system#built-in-middlewares)
    
*   [Creating Custom Middleware](https://deepwiki.com/charmbracelet/wish/3-middleware-system#creating-custom-middleware)
    
*   [Middleware Composition Example](https://deepwiki.com/charmbracelet/wish/3-middleware-system#middleware-composition-example)
    
*   [Middleware Chaining Implementation](https://deepwiki.com/charmbracelet/wish/3-middleware-system#middleware-chaining-implementation)
    
*   [Relationship to SSH Server Architecture](https://deepwiki.com/charmbracelet/wish/3-middleware-system#relationship-to-ssh-server-architecture)
    
*   [Summary](https://deepwiki.com/charmbracelet/wish/3-middleware-system#summary)
