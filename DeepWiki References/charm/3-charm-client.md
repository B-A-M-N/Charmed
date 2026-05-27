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

Charm Client
============

Relevant source files

*   [README.md](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1)
    
*   [client/client.go](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go)
    
*   [cmd/post\_news.go](https://github.com/charmbracelet/charm/blob/aff3071d/cmd/post_news.go)
    
*   [cmd/serve.go](https://github.com/charmbracelet/charm/blob/aff3071d/cmd/serve.go)
    
*   [crypt/README.md](https://github.com/charmbracelet/charm/blob/aff3071d/crypt/README.md?plain=1)
    
*   [docker.md](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1)
    
*   [systemd.md](https://github.com/charmbracelet/charm/blob/aff3071d/systemd.md?plain=1)
    

The Charm Client is a core component of the Charm Cloud ecosystem that manages authentication, identity, and keys for users. It provides low-level HTTP and SSH APIs for accessing the Charm Cloud server and facilitates communication between terminal-based applications and Charm backend services. This document covers the client architecture, capabilities, and integration points with other Charm components.

For more specific information about client configuration, see [Client Configuration](https://deepwiki.com/charmbracelet/charm/3.1-client-configuration)
, and for authentication details, see [Authentication](https://deepwiki.com/charmbracelet/charm/3.2-authentication)
.

Overview
--------

The Charm Client serves as the foundation for applications needing to interact with Charm Cloud services. It provides a seamless authentication mechanism based on SSH keys, allowing terminal applications to easily leverage cloud storage and synchronization capabilities without complex authentication flows.

Sources: [client/client.go1-51](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go#L1-L51)

Core Capabilities
-----------------

The Charm Client provides several fundamental capabilities that form the basis for all Charm-powered applications:

1.  **Identity Management**: Creates and manages user identity based on SSH keys
2.  **Multi-Device Support**: Allows linking multiple devices to a single account
3.  **Transparent Authentication**: Handles authentication and JWT token management
4.  **API Access**: Provides structured access to Charm server APIs
5.  **User Profile Management**: Allows setting and retrieving user profile information

Sources: [client/client.go130-266](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go#L130-L266)

Client Architecture
-------------------

The Charm Client is structured around a central `Client` struct that maintains configuration, authentication state, and connections to the server. It uses SSH for authentication and HTTP for data transfer, with all sensitive data encrypted on the client side.

### Key Components

| Component | Purpose |
| --- | --- |
| Config | Stores client configuration from environment or explicit settings |
| Authentication | Manages SSH keys and JWT tokens for server access |
| HTTP Client | Handles data transfer to and from the server |
| SSH Client | Manages authentication and identity verification |
| Key Management | Creates, stores, and manages SSH and encryption keys |

Sources: [client/client.go28-50](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go#L28-L50)

### Data Flow Diagram

Sources: [client/client.go273-284](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go#L273-L284)
 [client/client.go130-142](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go#L130-L142)

Usage
-----

The Charm Client can be used programmatically by importing the client package, or via the command-line interface provided by the `charm` binary.

### Programmatic Usage

The client can be initialized with default settings that load from environment variables:

Or with explicit configuration:

Sources: [client/client.go61-128](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go#L61-L128)

### Command-Line Usage

The Charm Client functionality is also accessible through the `charm` command-line tool:

Sources: [README.md218-238](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L218-L238)

Integration with Other Charm Components
---------------------------------------

The Charm Client serves as the foundation for other Charm components, providing the authentication and communication layer needed by higher-level services.

| Component | Integration with Charm Client |
| --- | --- |
| Charm KV | Uses client for authentication and encrypted data transfer |
| Charm FS | Uses client to access cloud-based file system API |
| Charm Crypt | Uses client for key management and end-to-end encryption |
| Terminal UI | Uses client for account management and data access |

Sources: [README.md46-60](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L46-L60)

Self-Hosting Compatibility
--------------------------

While the client connects to cloud.charm.sh by default, it can be configured to work with a self-hosted Charm server:

Running a self-hosted Charm server is as simple as:

Sources: [README.md283-289](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L283-L289)
 [cmd/serve.go18-76](https://github.com/charmbracelet/charm/blob/aff3071d/cmd/serve.go#L18-L76)

Key Management
--------------

The client automatically handles SSH key generation and management. By default, it creates Ed25519 keys, though RSA keys are also supported. Keys are stored in the user's data directory, which defaults to an XDG-compliant location but can be customized.

Sources: [client/client.go70-114](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go#L70-L114)
 [client/client.go300-324](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go#L300-L324)

Data Storage Location
---------------------

The client stores user data (including keys, cached data, and local configuration) in a directory determined by the following logic:

1.  If `CHARM_DATA_DIR` is set, data is stored in `$CHARM_DATA_DIR/$CHARM_HOST`
2.  Otherwise, data is stored in the platform-specific user data directory under `charm/$CHARM_HOST`

This allows for clean separation of data when connecting to different Charm servers.

Sources: [client/client.go286-298](https://github.com/charmbracelet/charm/blob/aff3071d/client/client.go#L286-L298)

Limitations and Considerations
------------------------------

*   Authentication is based on SSH keys, so key security is paramount
*   Default connection is to cloud.charm.sh, but this can be changed through configuration
*   The client is designed to handle connection disruptions gracefully, but requires eventual connectivity for synchronization
*   Key backups should be created with `charm backup-keys` to ensure account recovery capability

Sources: [README.md212-216](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L212-L216)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Charm Client](https://deepwiki.com/charmbracelet/charm/3-charm-client#charm-client)
    
*   [Overview](https://deepwiki.com/charmbracelet/charm/3-charm-client#overview)
    
*   [Core Capabilities](https://deepwiki.com/charmbracelet/charm/3-charm-client#core-capabilities)
    
*   [Client Architecture](https://deepwiki.com/charmbracelet/charm/3-charm-client#client-architecture)
    
*   [Key Components](https://deepwiki.com/charmbracelet/charm/3-charm-client#key-components)
    
*   [Data Flow Diagram](https://deepwiki.com/charmbracelet/charm/3-charm-client#data-flow-diagram)
    
*   [Usage](https://deepwiki.com/charmbracelet/charm/3-charm-client#usage)
    
*   [Programmatic Usage](https://deepwiki.com/charmbracelet/charm/3-charm-client#programmatic-usage)
    
*   [Command-Line Usage](https://deepwiki.com/charmbracelet/charm/3-charm-client#command-line-usage)
    
*   [Integration with Other Charm Components](https://deepwiki.com/charmbracelet/charm/3-charm-client#integration-with-other-charm-components)
    
*   [Self-Hosting Compatibility](https://deepwiki.com/charmbracelet/charm/3-charm-client#self-hosting-compatibility)
    
*   [Key Management](https://deepwiki.com/charmbracelet/charm/3-charm-client#key-management)
    
*   [Data Storage Location](https://deepwiki.com/charmbracelet/charm/3-charm-client#data-storage-location)
    
*   [Limitations and Considerations](https://deepwiki.com/charmbracelet/charm/3-charm-client#limitations-and-considerations)
