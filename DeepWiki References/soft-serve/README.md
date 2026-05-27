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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1)
    
*   [demo.tape](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/demo.tape)
    
*   [go.mod](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/go.sum)
    
*   [pkg/config/config.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go)
    
*   [pkg/config/config\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config_test.go)
    
*   [pkg/config/testdata/config.yaml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/testdata/config.yaml)
    
*   [pkg/web/server.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/web/server.go)
    

This document provides an introduction to Soft Serve, a self-hostable Git server that offers a terminal-based user interface and supports multiple protocols for Git operations. This overview explains the core architecture, key components, and main features of the system.

What is Soft Serve?
-------------------

Soft Serve is a self-hostable Git server specifically designed for the command line. It provides a terminal user interface (TUI) accessible over SSH, allowing users to browse and manage Git repositories. The server supports multiple protocols for Git operations including SSH, HTTP, and the Git protocol.

Sources: [README.md11-31](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L11-L31)

System Architecture
-------------------

Soft Serve consists of several server components that work together to provide Git repository hosting functionality.

Sources: [pkg/config/config.go19-63](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L19-L63)
 [pkg/config/config.go126-164](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L126-L164)
 [README.md18-32](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L18-L32)

Configuration System
--------------------

Soft Serve uses a robust configuration system that supports both configuration files and environment variables.

Sources: [pkg/config/config.go131-380](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L131-L380)
 [README.md136-141](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L136-L141)

### Core Configuration Elements

Soft Serve's configuration includes several key components:

| Component | Description | Default Port |
| --- | --- | --- |
| SSH Server | Provides SSH access for Git operations and TUI | 23231 |
| Git Daemon | Handles native Git protocol operations | 9418 |
| HTTP Server | Manages HTTP/HTTPS Git operations and LFS | 23232 |
| Stats Server | Exposes metrics and operational data | 23233 |
| Database | Stores user and repository metadata (SQLite or PostgreSQL) | N/A |
| LFS | Git Large File Storage support | Uses HTTP/SSH ports |

Sources: [pkg/config/config.go335-379](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L335-L379)
 [README.md142-273](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L142-L273)

Authentication and Authorization
--------------------------------

Soft Serve implements a granular access control system with authentication via SSH keys and user access tokens.

Sources: [README.md284-401](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L284-L401)

Repository Management
---------------------

Soft Serve provides comprehensive repository management functionality through its SSH command-line interface.

Sources: [README.md447-563](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L447-L563)
 [README.md644-667](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L644-L667)

Terminal User Interface
-----------------------

Soft Serve offers a rich terminal user interface accessible over SSH for browsing repositories.

Sources: [README.md669-693](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L669-L693)

User Interaction Flows
----------------------

The following sequence diagram illustrates the primary user interaction flows in Soft Serve:

Sources: [README.md33-53](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L33-L53)
 [pkg/web/server.go12-30](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/web/server.go#L12-L30)

Extending Soft Serve
--------------------

Soft Serve supports several extension mechanisms:

1.  **Git Hooks**: Server-side hooks like `pre-receive`, `update`, `post-update`, and `post-receive` for custom logic during Git operations
2.  **Webhooks**: Repository-specific webhooks for integrating with external systems
3.  **Custom Configuration**: Extensive configuration options via environment variables and configuration files

Sources: [README.md695-752](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L695-L752)

Summary
-------

Soft Serve provides a feature-rich, self-hostable Git server with a focus on terminal-based interaction. Its multi-protocol support, robust authentication system, and comprehensive repository management capabilities make it suitable for individual developers, teams, and organizations seeking a lightweight Git hosting solution.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/soft-serve#overview)
    
*   [What is Soft Serve?](https://deepwiki.com/charmbracelet/soft-serve#what-is-soft-serve)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/soft-serve#system-architecture)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/soft-serve#configuration-system)
    
*   [Core Configuration Elements](https://deepwiki.com/charmbracelet/soft-serve#core-configuration-elements)
    
*   [Authentication and Authorization](https://deepwiki.com/charmbracelet/soft-serve#authentication-and-authorization)
    
*   [Repository Management](https://deepwiki.com/charmbracelet/soft-serve#repository-management)
    
*   [Terminal User Interface](https://deepwiki.com/charmbracelet/soft-serve#terminal-user-interface)
    
*   [User Interaction Flows](https://deepwiki.com/charmbracelet/soft-serve#user-interaction-flows)
    
*   [Extending Soft Serve](https://deepwiki.com/charmbracelet/soft-serve#extending-soft-serve)
    
*   [Summary](https://deepwiki.com/charmbracelet/soft-serve#summary)
