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

Introduction to Wishlist
========================

Relevant source files

*   [README.md](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1)
    
*   [go.mod](https://github.com/charmbracelet/wishlist/blob/5fe34a29/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/wishlist/blob/5fe34a29/go.sum)
    

Wishlist is an SSH directory tool that provides a single entry point for managing multiple SSH endpoints. It allows users to list, discover, and connect to SSH servers through an interactive terminal interface. This document offers an overview of Wishlist's purpose, core features, and high-level architecture.

For detailed installation instructions, see [Getting Started](https://deepwiki.com/charmbracelet/wishlist/1.1-getting-started)
. For a comprehensive explanation of the system architecture, refer to [Architecture Overview](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview)
.

Purpose and Key Features
------------------------

Wishlist serves as a central hub for SSH connections with two primary modes of operation:

1.  **Local Mode**: A command-line interface that lists and connects to servers defined in your SSH configuration or YAML files
2.  **Server Mode**: A server application that aggregates multiple SSH endpoints into a single access point

Key features include:

*   Interactive terminal UI for browsing and connecting to SSH endpoints
*   Multiple endpoint discovery methods (Tailscale, Zeroconf/mDNS, SRV records)
*   Support for standard SSH configuration files
*   Flexible authentication handling with agent forwarding
*   Extensible middleware architecture

Sources: [README.md10-24](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L10-L24)

System Overview
---------------

Wishlist integrates several components to provide its functionality:

The architecture follows a logical separation of concerns between configuration, discovery, authentication, and user interfaces. This modular design allows Wishlist to be both flexible and extensible.

Sources: [README.md154-230](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L154-L230)

Operation Modes
---------------

Wishlist operates in two distinct modes, each serving different use cases:

### Local Mode

In this mode, Wishlist functions as a client application that:

*   Reads SSH configuration from various sources (`~/.ssh/config`, YAML files)
*   Discovers available SSH endpoints through multiple methods
*   Presents a terminal UI for selecting and connecting to endpoints
*   Handles authentication using SSH keys, agents, or other methods

Sources: [README.md95-104](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L95-L104)

### Server Mode

In server mode, Wishlist:

*   Runs as an SSH server that aggregates multiple endpoints
*   Presents a directory listing when clients connect
*   Forwards connections to the selected endpoint
*   Manages authentication for both itself and downstream connections

Sources: [README.md79-91](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L79-L91)

Data Flow
---------

The following diagram illustrates how data flows through the Wishlist system:

Sources: [README.md111-142](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L111-L142)

Configuration System
--------------------

Wishlist uses a flexible configuration system that supports multiple sources and formats:

| Configuration Source | Description | Priority |
| --- | --- | --- |
| `-config` flag | Command-line specified config file | Highest |
| `.wishlist/config.*` | Local project configuration | High |
| User config directory | System-wide user configuration | Medium |
| `~/.ssh/config` | Standard SSH configuration file | Low |
| `/etc/ssh/ssh_config` | System SSH configuration | Lowest |

The configuration system determines:

*   Available SSH endpoints
*   Authentication methods
*   Discovery settings
*   Server configuration

Sources: [README.md245-268](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L245-L268)

Component Architecture
----------------------

The following diagram shows the key components and their relationships within the Wishlist codebase:

The codebase is organized around these core components, with each handling a specific aspect of the system's functionality.

Sources: [README.md10-24](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L10-L24)
 [README.md154-230](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L154-L230)

Authentication Methods
----------------------

Wishlist supports multiple authentication methods:

### Local Mode Authentication

1.  Endpoint-specific identity files
2.  SSH agent (if available)
3.  Common key files in `~/.ssh/`

### Server Mode Authentication

1.  SSH agent forwarding
2.  Generated Ed25519 key (`.wishlist/client_ed25519`)

Wishlist provides automatic fallback between authentication methods, prioritizing security and convenience.

Sources: [README.md111-142](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L111-L142)

Discovery Methods
-----------------

Wishlist can discover SSH endpoints through several mechanisms:

### Tailscale Integration

Connects to the Tailscale API to discover nodes in your tailnet:

### Zeroconf/mDNS

Discovers SSH services advertised on the local network:

### SRV Records

Finds endpoints through DNS SRV records:

Sources: [README.md154-230](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L154-L230)

Summary
-------

Wishlist provides a unified way to organize, discover, and connect to SSH endpoints. Its flexible architecture supports various configuration sources, discovery methods, and authentication mechanisms. Whether used as a local client tool or as a server that aggregates multiple endpoints, Wishlist simplifies the management of SSH connections through an intuitive terminal interface.

For further information on specific components, refer to:

*   [Operation Modes](https://deepwiki.com/charmbracelet/wishlist/2.1-operation-modes)
     for detailed information about local and server modes
*   [Configuration System](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system)
     for configuration details
*   [SSH Client Implementation](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation)
     for client details
*   [SSH Server Implementation](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation)
     for server details
*   [Endpoint Discovery](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery)
     for more on discovery methods

Sources: [README.md1-336](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L1-L336)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Introduction to Wishlist](https://deepwiki.com/charmbracelet/wishlist#introduction-to-wishlist)
    
*   [Purpose and Key Features](https://deepwiki.com/charmbracelet/wishlist#purpose-and-key-features)
    
*   [System Overview](https://deepwiki.com/charmbracelet/wishlist#system-overview)
    
*   [Operation Modes](https://deepwiki.com/charmbracelet/wishlist#operation-modes)
    
*   [Local Mode](https://deepwiki.com/charmbracelet/wishlist#local-mode)
    
*   [Server Mode](https://deepwiki.com/charmbracelet/wishlist#server-mode)
    
*   [Data Flow](https://deepwiki.com/charmbracelet/wishlist#data-flow)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/wishlist#configuration-system)
    
*   [Component Architecture](https://deepwiki.com/charmbracelet/wishlist#component-architecture)
    
*   [Authentication Methods](https://deepwiki.com/charmbracelet/wishlist#authentication-methods)
    
*   [Local Mode Authentication](https://deepwiki.com/charmbracelet/wishlist#local-mode-authentication)
    
*   [Server Mode Authentication](https://deepwiki.com/charmbracelet/wishlist#server-mode-authentication)
    
*   [Discovery Methods](https://deepwiki.com/charmbracelet/wishlist#discovery-methods)
    
*   [Tailscale Integration](https://deepwiki.com/charmbracelet/wishlist#tailscale-integration)
    
*   [Zeroconf/mDNS](https://deepwiki.com/charmbracelet/wishlist#zeroconfmdns)
    
*   [SRV Records](https://deepwiki.com/charmbracelet/wishlist#srv-records)
    
*   [Summary](https://deepwiki.com/charmbracelet/wishlist#summary)
