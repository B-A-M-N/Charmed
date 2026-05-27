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

Architecture Overview
=====================

Relevant source files

*   [cmd/wishlist/main.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go)
    
*   [go.mod](https://github.com/charmbracelet/wishlist/blob/5fe34a29/go.mod)
    
*   [server.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go)
    
*   [wishlist.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go)
    

This document provides a comprehensive overview of the Wishlist architecture, explaining how its core components interact to form a cohesive SSH directory system. It covers the high-level structure, component relationships, and data flow within the application. For information on specific operation modes, see [Operation Modes](https://deepwiki.com/charmbracelet/wishlist/2.1-operation-modes)
.

Core Architecture
-----------------

Wishlist is designed as a flexible SSH management tool that can operate both as a local CLI tool and a SSH server. The architecture follows a modular design with clear separation of concerns between configuration, endpoint discovery, authentication, and user interface components.

Sources: [cmd/wishlist/main.go45-194](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L45-L194)
 [server.go22-156](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L22-L156)
 [wishlist.go27-47](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L27-L47)

Component Relationships
-----------------------

The following diagram illustrates the relationships between the core data structures and components in Wishlist. These structures form the foundation of the system's functionality.

Sources: [wishlist.go49-59](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L49-L59)
 [server.go22-96](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L22-L96)

Data Flow
---------

The following diagram illustrates the data flow through the Wishlist system, from user interaction to SSH connection establishment.

Sources: [cmd/wishlist/main.go96-105](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L96-L105)
 [cmd/wishlist/main.go308-392](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L308-L392)
 [cmd/wishlist/main.go395-430](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L395-L430)
 [server.go22-96](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L22-L96)

Key Components
--------------

### 1\. Configuration System

The configuration system is responsible for loading and parsing configuration from various sources. It supports both standard SSH config files and YAML configuration.

Sources: [cmd/wishlist/main.go308-392](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L308-L392)

The configuration system handles:

*   Parsing SSH config files via `sshconfig.ParseFile()`
*   Parsing YAML configuration with `yaml.Unmarshal()`
*   Applying hints to discovered endpoints
*   Providing a unified configuration structure for the application

### 2\. SSH Server

The SSH server component allows Wishlist to serve the TUI over SSH, enabling remote access to the SSH directory.

Sources: [server.go22-96](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L22-L96)
 [server.go98-121](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L98-L121)

The SSH server:

*   Listens on configured ports
*   Handles SSH connections
*   Applies middleware for different functionality
*   Manages user authentication
*   Provides graceful shutdown on signals

### 3\. Terminal UI

The Terminal UI is built using Bubble Tea and provides an interactive interface for browsing and connecting to SSH endpoints.

Sources: [wishlist.go27-47](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L27-L47)
 [wishlist.go114-231](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L114-L231)

The Terminal UI component:

*   Displays a list of available endpoints
*   Handles user interactions (selection, filtering)
*   Manages the connection to selected endpoints
*   Provides visual feedback and error handling

### 4\. Discovery Methods

Wishlist supports multiple methods for discovering SSH endpoints:

| Discovery Method | Description | Implementation |
| --- | --- | --- |
| Tailscale | Discovers endpoints in a Tailscale network | `tailscale.Endpoints()` |
| Zeroconf/mDNS | Uses multicast DNS for local network discovery | `zeroconf.Endpoints()` |
| SRV Records | Uses DNS SRV records to discover endpoints | `srv.Endpoints()` |

Each discovery method returns a list of endpoints that are then combined and enhanced with configuration information.

Sources: [cmd/wishlist/main.go340-366](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L340-L366)

SSH Connection Flow
-------------------

The following sequence diagram illustrates the SSH connection process from user selection to connection establishment:

Sources: [wishlist.go127-160](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L127-L160)
 [cmd/wishlist/main.go421-430](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L421-L430)

Execution Modes
---------------

Wishlist operates in two primary modes:

1.  **Local Mode**: Runs as a local CLI application, presenting a TUI for endpoint selection
2.  **Server Mode**: Runs as an SSH server, allowing remote users to access the TUI over SSH

Sources: [cmd/wishlist/main.go45-106](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L45-L106)
 [cmd/wishlist/main.go126-194](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L126-L194)
 [cmd/wishlist/main.go395-430](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L395-L430)

Middleware Architecture
-----------------------

The SSH server functionality is extended through a composable middleware architecture:

Sources: [server.go54-70](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go#L54-L70)
 [cmd/wishlist/main.go177-187](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L177-L187)

The middleware system:

*   Allows for separation of concerns in handling SSH sessions
*   Enables easy addition of new functionality
*   Provides a pipeline for processing SSH connections
*   Supports conditional middleware application

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Architecture Overview](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#architecture-overview)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#core-architecture)
    
*   [Component Relationships](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#component-relationships)
    
*   [Data Flow](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#data-flow)
    
*   [Key Components](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#key-components)
    
*   [1\. Configuration System](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#1-configuration-system)
    
*   [2\. SSH Server](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#2-ssh-server)
    
*   [3\. Terminal UI](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#3-terminal-ui)
    
*   [4\. Discovery Methods](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#4-discovery-methods)
    
*   [SSH Connection Flow](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#ssh-connection-flow)
    
*   [Execution Modes](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#execution-modes)
    
*   [Middleware Architecture](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview#middleware-architecture)
