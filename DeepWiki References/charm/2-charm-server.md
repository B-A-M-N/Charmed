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

Charm Server
============

Relevant source files

*   [README.md](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1)
    
*   [crypt/README.md](https://github.com/charmbracelet/charm/blob/aff3071d/crypt/README.md?plain=1)
    
*   [docker.md](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1)
    
*   [server/server.go](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go)
    
*   [systemd.md](https://github.com/charmbracelet/charm/blob/aff3071d/systemd.md?plain=1)
    

The Charm Server is the backend component of the Charm Cloud system, providing authentication, data storage, and synchronization services for Charm clients. This document covers the architecture, components, and configuration options of the Charm Server.

For information about client-side components, see [Charm Client](https://deepwiki.com/charmbracelet/charm/3-charm-client)
.

Server Overview
---------------

The Charm Server consists of multiple server components that work together to provide a complete backend for terminal-based applications. It handles user authentication via SSH keys, stores encrypted user data, and provides APIs for key-value storage, file systems, and encrypted data.

**Server Architecture Diagram**

Sources: [server/server.go54-59](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L54-L59)
 [server/server.go151-166](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L151-L166)

Server Components
-----------------

The Charm Server is composed of several key components:

1.  **SSH Server**: Handles user authentication via SSH keys and account linking
2.  **HTTP Server**: Provides REST APIs for data access and manipulation
3.  **Stats Server**: Collects and exposes metrics (optional)
4.  **Health Server**: Provides health check endpoints
5.  **Database**: Stores user data, keys, and metadata (SQLite by default)
6.  **File Storage**: Stores user files

Each component can be configured independently and works together to provide the complete Charm Cloud experience.

### Component Interactions

**Server Component Interaction Flow**

Sources: [server/server.go138-148](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L138-L148)
 [server/server.go151-166](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L151-L166)

Server Configuration
--------------------

The Charm Server can be configured using environment variables. The default configuration uses SQLite for the database and local file storage for files.

### Environment Variables

| Variable | Description | Default |
| --- | --- | --- |
| CHARM\_SERVER\_BIND\_ADDRESS | Network interface to listen on | 0.0.0.0 |
| CHARM\_SERVER\_HOST | Hostname to advertise | localhost |
| CHARM\_SERVER\_SSH\_PORT | SSH server port | 35353 |
| CHARM\_SERVER\_HTTP\_PORT | HTTP server port | 35354 |
| CHARM\_SERVER\_STATS\_PORT | Stats server port | 35355 |
| CHARM\_SERVER\_HEALTH\_PORT | Health server port | 35356 |
| CHARM\_SERVER\_DATA\_DIR | Server data directory | ./data |
| CHARM\_SERVER\_USE\_TLS | Whether to use TLS | false |
| CHARM\_SERVER\_TLS\_KEY\_FILE | The TLS key file path |     |
| CHARM\_SERVER\_TLS\_CERT\_FILE | The TLS cert file path |     |
| CHARM\_SERVER\_PUBLIC\_URL | Server public URL |     |
| CHARM\_SERVER\_ENABLE\_METRICS | Whether to enable Prometheus metrics | false |
| CHARM\_SERVER\_USER\_MAX\_STORAGE | Maximum FS storage for a user (0 = no limit) | 0   |

Sources: [server/server.go28-42](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L28-L42)
 [README.md266-282](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L266-L282)

### Server Initialization Flow

**Server Initialization Flow**

Sources: [server/server.go129-149](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L129-L149)
 [server/server.go201-228](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L201-L228)

Data Storage and Security
-------------------------

The Charm Server stores all user data in an encrypted format. The encryption keys are managed by the client, and the server only stores encrypted data. This ensures that even server administrators cannot access the plaintext user data.

### Data Storage Components

1.  **SQLite Database**: Stores user accounts, metadata, and encrypted key-value data
2.  **File Storage**: Stores encrypted user files
3.  **JWT Authentication**: Secures API requests after initial SSH authentication

### End-to-End Encryption

All sensitive data sent to the Charm Server is encrypted on the client side before transmission. The server never has access to the encryption keys needed to decrypt user data.

**End-to-End Encryption Flow**

Sources: [crypt/README.md1-29](https://github.com/charmbracelet/charm/blob/aff3071d/crypt/README.md?plain=1#L1-L29)

Deployment Options
------------------

The Charm Server can be deployed in various ways, including standalone, Docker, and systemd.

### Docker Deployment

The official Docker image is available at `charmcli/charm:latest`. The server can be run with Docker using:

Sources: [docker.md1-48](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1#L1-L48)

### Systemd Deployment

To run the Charm Server as a systemd service, create a file at `/etc/systemd/system/charm.service`:

    [Unit]
    Description=The mystical Charm Cloud 🌟
    After=network.target
    StartLimitIntervalSec=0
    
    [Service]
    Type=simple
    Restart=always
    RestartSec=1
    Environment=CHARM_SERVER_DATA_DIR=/var/lib/charm
    ExecStart=/usr/bin/charm serve
    
    [Install]
    WantedBy=multi-user.target
    

Sources: [systemd.md1-31](https://github.com/charmbracelet/charm/blob/aff3071d/systemd.md?plain=1#L1-L31)

TLS Configuration
-----------------

The Charm Server supports TLS for secure communication. To enable TLS:

1.  Set `CHARM_SERVER_USE_TLS` to `true`
2.  Specify `CHARM_SERVER_TLS_KEY_FILE` and `CHARM_SERVER_TLS_CERT_FILE`
3.  Set `CHARM_SERVER_HOST` to your server's hostname

Alternatively, you can use a reverse proxy like Traefik or Caddy to handle TLS termination. In this case, set `CHARM_SERVER_PUBLIC_URL` to the public URL of your reverse proxy.

Sources: [README.md301-306](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L301-L306)
 [docker.md43-48](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1#L43-L48)

Metrics and Monitoring
----------------------

The Charm Server includes an optional metrics server that exposes Prometheus-compatible metrics. To enable metrics:

1.  Set `CHARM_SERVER_ENABLE_METRICS` to `true`
2.  Access metrics at `http://<CHARM_SERVER_HOST>:<CHARM_SERVER_STATS_PORT>/metrics`

The server also provides a health check endpoint on the health port (default: 35356).

Sources: [server/server.go223-228](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L223-L228)
 [README.md280](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L280-L280)

Storage Limitations
-------------------

The free Charm Cloud service provided by Charmbracelet, Inc. limits user storage to 1GB per account. When self-hosting, there is no default storage limit, but you can set one using the `CHARM_SERVER_USER_MAX_STORAGE` environment variable.

Sources: [README.md294-299](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L294-L299)

Server Component Relationships
------------------------------

**Server Component Relationships**

Sources: [server/server.go54-59](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L54-L59)
 [server/server.go27-52](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L27-L52)
 [server/server.go201-228](https://github.com/charmbracelet/charm/blob/aff3071d/server/server.go#L201-L228)

This diagram shows the relationship between the main server components and how they connect to the various interface implementations. The server uses dependency injection to allow for flexible configuration of the database, file storage, and statistics components.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Charm Server](https://deepwiki.com/charmbracelet/charm/2-charm-server#charm-server)
    
*   [Server Overview](https://deepwiki.com/charmbracelet/charm/2-charm-server#server-overview)
    
*   [Server Components](https://deepwiki.com/charmbracelet/charm/2-charm-server#server-components)
    
*   [Component Interactions](https://deepwiki.com/charmbracelet/charm/2-charm-server#component-interactions)
    
*   [Server Configuration](https://deepwiki.com/charmbracelet/charm/2-charm-server#server-configuration)
    
*   [Environment Variables](https://deepwiki.com/charmbracelet/charm/2-charm-server#environment-variables)
    
*   [Server Initialization Flow](https://deepwiki.com/charmbracelet/charm/2-charm-server#server-initialization-flow)
    
*   [Data Storage and Security](https://deepwiki.com/charmbracelet/charm/2-charm-server#data-storage-and-security)
    
*   [Data Storage Components](https://deepwiki.com/charmbracelet/charm/2-charm-server#data-storage-components)
    
*   [End-to-End Encryption](https://deepwiki.com/charmbracelet/charm/2-charm-server#end-to-end-encryption)
    
*   [Deployment Options](https://deepwiki.com/charmbracelet/charm/2-charm-server#deployment-options)
    
*   [Docker Deployment](https://deepwiki.com/charmbracelet/charm/2-charm-server#docker-deployment)
    
*   [Systemd Deployment](https://deepwiki.com/charmbracelet/charm/2-charm-server#systemd-deployment)
    
*   [TLS Configuration](https://deepwiki.com/charmbracelet/charm/2-charm-server#tls-configuration)
    
*   [Metrics and Monitoring](https://deepwiki.com/charmbracelet/charm/2-charm-server#metrics-and-monitoring)
    
*   [Storage Limitations](https://deepwiki.com/charmbracelet/charm/2-charm-server#storage-limitations)
    
*   [Server Component Relationships](https://deepwiki.com/charmbracelet/charm/2-charm-server#server-component-relationships)
