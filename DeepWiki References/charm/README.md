Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/charm](https://github.com/charmbracelet/charm "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 21 April 2025 ([aff307](https://github.com/charmbracelet/charm/commits/aff3071d)
)

*   [Overview](https://deepwiki.com/charmbracelet/charm/1-overview)
    
*   [Core Concepts](https://deepwiki.com/charmbracelet/charm/1.1-core-concepts)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/charm/1.2-architecture-overview)
    
*   [Charm Server](https://deepwiki.com/charmbracelet/charm/2-charm-server)
    
*   [HTTP Server](https://deepwiki.com/charmbracelet/charm/2.1-http-server)
    
*   [SSH Server](https://deepwiki.com/charmbracelet/charm/2.2-ssh-server)
    
*   [Server Configuration](https://deepwiki.com/charmbracelet/charm/2.3-server-configuration)
    
*   [Database and Storage](https://deepwiki.com/charmbracelet/charm/2.4-database-and-storage)
    
*   [Metrics and Monitoring](https://deepwiki.com/charmbracelet/charm/2.5-metrics-and-monitoring)
    
*   [Charm Client](https://deepwiki.com/charmbracelet/charm/3-charm-client)
    
*   [Client Configuration](https://deepwiki.com/charmbracelet/charm/3.1-client-configuration)
    
*   [Authentication](https://deepwiki.com/charmbracelet/charm/3.2-authentication)
    
*   [Key Management](https://deepwiki.com/charmbracelet/charm/3.3-key-management)
    
*   [Terminal UI](https://deepwiki.com/charmbracelet/charm/4-terminal-ui)
    
*   [UI Components](https://deepwiki.com/charmbracelet/charm/4.1-ui-components)
    
*   [Key Management UI](https://deepwiki.com/charmbracelet/charm/4.2-key-management-ui)
    
*   [User Profile UI](https://deepwiki.com/charmbracelet/charm/4.3-user-profile-ui)
    
*   [Data Storage](https://deepwiki.com/charmbracelet/charm/5-data-storage)
    
*   [File System (FS)](https://deepwiki.com/charmbracelet/charm/5.1-file-system-(fs))
    
*   [Key-Value Store (KV)](https://deepwiki.com/charmbracelet/charm/5.2-key-value-store-(kv))
    
*   [Encryption](https://deepwiki.com/charmbracelet/charm/5.3-encryption)
    
*   [Deployment](https://deepwiki.com/charmbracelet/charm/6-deployment)
    
*   [Self-Hosting Guide](https://deepwiki.com/charmbracelet/charm/6.1-self-hosting-guide)
    
*   [Docker Deployment](https://deepwiki.com/charmbracelet/charm/6.2-docker-deployment)
    
*   [Backup and Restore](https://deepwiki.com/charmbracelet/charm/6.3-backup-and-restore)
    
*   [Development](https://deepwiki.com/charmbracelet/charm/7-development)
    
*   [Building and Testing](https://deepwiki.com/charmbracelet/charm/7.1-building-and-testing)
    
*   [CI/CD Workflows](https://deepwiki.com/charmbracelet/charm/7.2-cicd-workflows)
    
*   [Command-Line Interface](https://deepwiki.com/charmbracelet/charm/7.3-command-line-interface)
    

Menu

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1)
    
*   [crypt/README.md](https://github.com/charmbracelet/charm/blob/aff3071d/crypt/README.md?plain=1)
    
*   [docker.md](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1)
    
*   [go.mod](https://github.com/charmbracelet/charm/blob/aff3071d/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/charm/blob/aff3071d/go.sum)
    
*   [systemd.md](https://github.com/charmbracelet/charm/blob/aff3071d/systemd.md?plain=1)
    

Charm is a set of tools that provides backend services for terminal-based applications. It offers user account management, data storage, and encryption capabilities without requiring developers to implement these features from scratch. This document provides a technical overview of the Charm system, its architecture, and how its components interact.

For more detailed information about specific components, please refer to:

*   [Core Concepts](https://deepwiki.com/charmbracelet/charm/1.1-core-concepts)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/charm/1.2-architecture-overview)
    
*   [Server Implementation](https://deepwiki.com/charmbracelet/charm/2-charm-server)
    
*   [Client Implementation](https://deepwiki.com/charmbracelet/charm/3-charm-client)
    
*   [Data Storage](https://deepwiki.com/charmbracelet/charm/5-data-storage)
    

Core Components
---------------

Charm consists of four main components that work together to provide a complete backend solution:

1.  **Charm KV**: An encrypted, cloud-synced key-value store built on BadgerDB
2.  **Charm FS**: A cloud-based user filesystem compatible with Go's fs.FS interface
3.  **Charm Crypt**: End-to-end encryption for stored data and on-demand encryption
4.  **Charm Accounts**: Authentication system based on SSH keys

### Component Relationship Diagram

Sources: [README.md46-55](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L46-L55)
 [README.md102-177](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L102-L177)

Client-Server Architecture
--------------------------

Charm follows a client-server architecture with well-defined interfaces and communication protocols. The client components communicate with the server via SSH (for authentication) and HTTP (for data operations).

### Architecture Diagram

Sources: [README.md242-253](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L242-L253)
 [README.md267-281](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L267-L281)

Data Flow
---------

The data flow in Charm involves multiple layers to ensure security and proper synchronization between the client and server.

### Data Flow Diagram

Sources: [README.md102-177](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L102-L177)
 [crypt/README.md20-29](https://github.com/charmbracelet/charm/blob/aff3071d/crypt/README.md?plain=1#L20-L29)

Authentication System
---------------------

Charm uses SSH keys for authentication, providing a secure and passwordless experience. Users can link multiple devices to the same account, and each linked device can access the user's encrypted data.

### Authentication Flow

Sources: [README.md200-215](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L200-L215)

Encryption System
-----------------

All data in Charm is encrypted on the client side before being transmitted, ensuring end-to-end encryption. This means even the server operators cannot access the plaintext data.

### Encryption Flow

Sources: [crypt/README.md20-29](https://github.com/charmbracelet/charm/blob/aff3071d/crypt/README.md?plain=1#L20-L29)
 [README.md192-198](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L192-L198)

Server Implementation
---------------------

The Charm server provides the backend services for all Charm components through several interconnected subsystems.

### Server Components

Sources: [README.md256-299](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L256-L299)

Configuration
-------------

### Client Configuration

The Charm client can be configured using environment variables:

| Variable | Description | Default |
| --- | --- | --- |
| `CHARM_HOST` | Server public URL | cloud.charm.sh |
| `CHARM_SSH_PORT` | SSH port to connect to | 35353 |
| `CHARM_HTTP_PORT` | HTTP port to connect to | 35354 |
| `CHARM_DEBUG` | Enable debugging logs | false |
| `CHARM_LOGFILE` | Debug log file path |     |
| `CHARM_KEY_TYPE` | Key type for new users | ed25519 |
| `CHARM_DATA_DIR` | User data directory |     |
| `CHARM_IDENTITY_KEY` | Identity key path |     |

Sources: [README.md242-253](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L242-L253)

### Server Configuration

The Charm server can be configured using environment variables:

| Variable | Description | Default |
| --- | --- | --- |
| `CHARM_SERVER_BIND_ADDRESS` | Network interface | 0.0.0.0 |
| `CHARM_SERVER_HOST` | Hostname to advertise | localhost |
| `CHARM_SERVER_SSH_PORT` | SSH server port | 35353 |
| `CHARM_SERVER_HTTP_PORT` | HTTP server port | 35354 |
| `CHARM_SERVER_STATS_PORT` | Stats server port | 35355 |
| `CHARM_SERVER_HEALTH_PORT` | Health server port | 35356 |
| `CHARM_SERVER_DATA_DIR` | Server data directory | ./data |
| `CHARM_SERVER_USE_TLS` | Enable TLS | false |
| `CHARM_SERVER_TLS_KEY_FILE` | TLS key file path |     |
| `CHARM_SERVER_TLS_CERT_FILE` | TLS cert file path |     |
| `CHARM_SERVER_PUBLIC_URL` | Server public URL |     |
| `CHARM_SERVER_ENABLE_METRICS` | Enable Prometheus metrics | false |
| `CHARM_SERVER_USER_MAX_STORAGE` | Max user storage (0=unlimited) | 0   |

Sources: [README.md267-281](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L267-L281)

Deployment Options
------------------

Charm can be deployed in several ways:

1.  **Hosted Charm Cloud**: Using the Charmbracelet, Inc. hosted service at cloud.charm.sh (note: being sunset on November 29, 2024)
2.  **Self-hosted Server**: Running your own server with `charm serve`
3.  **Docker**: Using the official Docker image
4.  **Systemd**: Running as a systemd service

The `charm serve` command starts a server instance with the specified configuration.

### Storage Considerations

The maximum data storage on the official Charm Cloud service is 1GB per account. Self-hosted servers don't have a data storage limit by default, but administrators can set a limit using the `CHARM_SERVER_USER_MAX_STORAGE` environment variable.

Sources: [README.md1-3](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L1-L3)
 [README.md256-267](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L256-L267)
 [docker.md1-42](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1#L1-L42)
 [systemd.md1-31](https://github.com/charmbracelet/charm/blob/aff3071d/systemd.md?plain=1#L1-L31)

Using the Charm CLI
-------------------

The `charm` binary provides direct access to the functionality available in the libraries:

Sources: [README.md217-238](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L217-L238)

This overview provides a high-level understanding of the Charm system architecture and components. For more detailed information on using specific components, refer to the corresponding sections of the documentation.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/charm#overview)
    
*   [Core Components](https://deepwiki.com/charmbracelet/charm#core-components)
    
*   [Component Relationship Diagram](https://deepwiki.com/charmbracelet/charm#component-relationship-diagram)
    
*   [Client-Server Architecture](https://deepwiki.com/charmbracelet/charm#client-server-architecture)
    
*   [Architecture Diagram](https://deepwiki.com/charmbracelet/charm#architecture-diagram)
    
*   [Data Flow](https://deepwiki.com/charmbracelet/charm#data-flow)
    
*   [Data Flow Diagram](https://deepwiki.com/charmbracelet/charm#data-flow-diagram)
    
*   [Authentication System](https://deepwiki.com/charmbracelet/charm#authentication-system)
    
*   [Authentication Flow](https://deepwiki.com/charmbracelet/charm#authentication-flow)
    
*   [Encryption System](https://deepwiki.com/charmbracelet/charm#encryption-system)
    
*   [Encryption Flow](https://deepwiki.com/charmbracelet/charm#encryption-flow)
    
*   [Server Implementation](https://deepwiki.com/charmbracelet/charm#server-implementation)
    
*   [Server Components](https://deepwiki.com/charmbracelet/charm#server-components)
    
*   [Configuration](https://deepwiki.com/charmbracelet/charm#configuration)
    
*   [Client Configuration](https://deepwiki.com/charmbracelet/charm#client-configuration)
    
*   [Server Configuration](https://deepwiki.com/charmbracelet/charm#server-configuration)
    
*   [Deployment Options](https://deepwiki.com/charmbracelet/charm#deployment-options)
    
*   [Storage Considerations](https://deepwiki.com/charmbracelet/charm#storage-considerations)
    
*   [Using the Charm CLI](https://deepwiki.com/charmbracelet/charm#using-the-charm-cli)
