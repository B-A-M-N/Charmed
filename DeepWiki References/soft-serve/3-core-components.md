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

Core Components
===============

Relevant source files

*   [.gitignore](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.gitignore)
    
*   [cmd/soft/main.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go)
    
*   [cmd/soft/serve/server.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/serve/server.go)
    
*   [pkg/backend/repo.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/backend/repo.go)
    
*   [pkg/daemon/daemon.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/daemon/daemon.go)
    
*   [pkg/daemon/daemon\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/daemon/daemon_test.go)
    
*   [pkg/ssh/cmd/cmd.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/cmd/cmd.go)
    
*   [pkg/utils/utils.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/utils/utils.go)
    
*   [testscript/testdata/config-servers-git\_disabled.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/config-servers-git_disabled.txtar)
    
*   [testscript/testdata/config-servers-http\_disabled.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/config-servers-http_disabled.txtar)
    
*   [testscript/testdata/config-servers-ssh\_disabled.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/config-servers-ssh_disabled.txtar)
    
*   [testscript/testdata/config-servers-stats\_disabled.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/config-servers-stats_disabled.txtar)
    

This page describes the main architectural components that make up Soft Serve's core functionality. It provides a technical overview of how the server is structured and how the various components interact with each other to provide Git hosting functionality.

For information about installing and configuring Soft Serve, see [Installation and Configuration](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration)
. For details about user interfaces, see [User Interfaces](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces)
.

Overview
--------

Soft Serve is composed of multiple server components that work together to provide Git repository hosting with SSH-based terminal UI, Git protocol support, HTTP access, and metrics collection. These components are orchestrated by a central server structure.

Sources: [cmd/soft/serve/server.go24-38](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/serve/server.go#L24-L38)

Server Orchestration
--------------------

The `Server` struct serves as the orchestrator for all core components. It's responsible for initializing, starting, and shutting down the various server components.

Sources: [cmd/soft/serve/server.go93-146](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/serve/server.go#L93-L146)

The `Server` component initializes each server with the appropriate context that contains configuration, backend, and database connections. Each server component can be enabled or disabled through configuration.

### Initialization and Lifecycle

The `NewServer` function creates instances of all server components but doesn't start them. The `Start` method launches each enabled server component in its own goroutine. The `Shutdown` and `Close` methods provide graceful and immediate shutdown options respectively.

    // Component lifecycle methods
    Start() - Starts all enabled server components
    Shutdown(ctx) - Gracefully shuts down all components
    Close() - Immediately closes all components
    

Sources: [cmd/soft/serve/server.go40-91](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/serve/server.go#L40-L91)
 [cmd/soft/serve/server.go148-187](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/serve/server.go#L148-L187)

Backend
-------

The `Backend` component is the core business logic layer that manages repositories, access control, and provides the data needed by the various server components.

Sources: [pkg/backend/repo.go28-489](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/backend/repo.go#L28-L489)
 [pkg/ssh/cmd/cmd.go110-197](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/cmd/cmd.go#L110-L197)

### Repository Management

The Backend provides comprehensive repository management functionality:

| Function | Description |
| --- | --- |
| `CreateRepository` | Creates a new Git repository |
| `ImportRepository` | Imports a repository from a remote URL |
| `DeleteRepository` | Deletes a repository |
| `RenameRepository` | Renames a repository |
| `Repositories` | Lists all repositories |
| `Repository` | Gets a repository by name |
| `SetDescription` | Updates repository description |
| `SetPrivate` | Updates repository visibility |
| `SetHidden` | Updates repository visibility in lists |

Sources: [pkg/backend/repo.go28-613](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/backend/repo.go#L28-L613)

### Repository Structure

Each repository is represented by the `repo` struct which provides access to the Git repository and its metadata:

Sources: [pkg/backend/repo.go638-758](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/backend/repo.go#L638-L758)

Protocol Servers
----------------

Soft Serve exposes repositories through multiple protocols, each handled by a dedicated server component.

### SSH Server

The SSH Server provides both Git access via SSH and an interactive terminal UI for browsing repositories. It authenticates users via SSH keys and enforces access controls.

Sources: [pkg/ssh/cmd/cmd.go19-197](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/cmd/cmd.go#L19-L197)

### Git Daemon

The Git Daemon serves repositories over the native Git protocol (git://), which provides unauthenticated read-only access to non-private repositories.

The `GitDaemon` manages TCP connections, handles Git protocol requests, and delegates to the appropriate Git service (upload-pack for cloning/fetching or upload-archive for archives).

Sources: [pkg/daemon/daemon.go44-349](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/daemon/daemon.go#L44-L349)

#### Connection Handling

The Git Daemon includes sophisticated connection management:

*   Limits the maximum number of concurrent connections
*   Implements connection timeouts (idle and maximum)
*   Handles temporary network errors with exponential backoff

Sources: [pkg/daemon/daemon.go74-134](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/daemon/daemon.go#L74-L134)

#### Request Processing

When a Git client connects, the daemon:

1.  Parses the Git protocol request
2.  Validates the repository name
3.  Checks access permissions
4.  Delegates to the appropriate Git service
5.  Collects metrics about the operation

Sources: [pkg/daemon/daemon.go144-310](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/daemon/daemon.go#L144-L310)

### HTTP Server

The HTTP Server provides:

*   Git HTTP protocol support (smart HTTP)
*   Git LFS (Large File Storage) support
*   API endpoints for programmatic access
*   Web hooks for integration with other systems

Sources: [cmd/soft/serve/server.go80-83](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/serve/server.go#L80-L83)
 [cmd/soft/serve/server.go120-128](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/serve/server.go#L120-L128)

### Stats Server

The Stats Server exposes metrics and status information for monitoring the Soft Serve instance, including:

*   Repository statistics
*   Connection metrics
*   System health information

Sources: [cmd/soft/serve/server.go85-88](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/serve/server.go#L85-L88)
 [cmd/soft/serve/server.go131-139](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/serve/server.go#L131-L139)

Command Line Interface
----------------------

The Soft Serve CLI provides both server management and client functionality.

Sources: [cmd/soft/main.go39-65](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go#L39-L65)
 [cmd/soft/main.go73-80](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go#L73-L80)

The CLI is built using Cobra, with custom usage templates for SSH-specific help formatting:

    Usage:
      soft serve [flags]
      ssh -p 23231 localhost
    
    Flags:
      --sync-hooks  Synchronize git hooks
      ...
    

Sources: [pkg/ssh/cmd/cmd.go30-60](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/ssh/cmd/cmd.go#L30-L60)

Configuration System
--------------------

Soft Serve uses a configuration system that supports multiple sources (files, environment variables, defaults) and provides typed access to settings.

Sources: [cmd/soft/main.go102-111](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/soft/main.go#L102-L111)

Individual server components can be enabled or disabled through configuration:

    SOFT_SERVE_SSH_ENABLED=true/false
    SOFT_SERVE_GIT_ENABLED=true/false
    SOFT_SERVE_HTTP_ENABLED=true/false
    SOFT_SERVE_STATS_ENABLED=true/false
    

Sources: [testscript/testdata/config-servers-ssh\_disabled.txtar4-7](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/config-servers-ssh_disabled.txtar#L4-L7)
 [testscript/testdata/config-servers-git\_disabled.txtar4-7](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/config-servers-git_disabled.txtar#L4-L7)
 [testscript/testdata/config-servers-http\_disabled.txtar4-7](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/config-servers-http_disabled.txtar#L4-L7)
 [testscript/testdata/config-servers-stats\_disabled.txtar4-7](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/config-servers-stats_disabled.txtar#L4-L7)

Repository Utilities
--------------------

Soft Serve includes utilities for repository management:

### Repository Name Validation

The `utils` package provides functions for sanitizing and validating repository names:

*   `SanitizeRepo(repo string)` - Normalizes repository paths
*   `ValidateRepo(repo string)` - Ensures repository names meet requirements

Sources: [pkg/utils/utils.go10-55](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/utils/utils.go#L10-L55)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core Components](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#core-components)
    
*   [Overview](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#overview)
    
*   [Server Orchestration](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#server-orchestration)
    
*   [Initialization and Lifecycle](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#initialization-and-lifecycle)
    
*   [Backend](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#backend)
    
*   [Repository Management](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#repository-management)
    
*   [Repository Structure](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#repository-structure)
    
*   [Protocol Servers](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#protocol-servers)
    
*   [SSH Server](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#ssh-server)
    
*   [Git Daemon](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#git-daemon)
    
*   [Connection Handling](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#connection-handling)
    
*   [Request Processing](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#request-processing)
    
*   [HTTP Server](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#http-server)
    
*   [Stats Server](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#stats-server)
    
*   [Command Line Interface](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#command-line-interface)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#configuration-system)
    
*   [Repository Utilities](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#repository-utilities)
    
*   [Repository Name Validation](https://deepwiki.com/charmbracelet/soft-serve/3-core-components#repository-name-validation)
