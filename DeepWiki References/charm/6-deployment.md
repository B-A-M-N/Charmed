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

Deployment
==========

Relevant source files

*   [README.md](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1)
    
*   [crypt/README.md](https://github.com/charmbracelet/charm/blob/aff3071d/crypt/README.md?plain=1)
    
*   [docker.md](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1)
    
*   [docs/backup-account.md](https://github.com/charmbracelet/charm/blob/aff3071d/docs/backup-account.md?plain=1)
    
*   [docs/restore-account.md](https://github.com/charmbracelet/charm/blob/aff3071d/docs/restore-account.md?plain=1)
    
*   [docs/self-hosting.md](https://github.com/charmbracelet/charm/blob/aff3071d/docs/self-hosting.md?plain=1)
    
*   [systemd.md](https://github.com/charmbracelet/charm/blob/aff3071d/systemd.md?plain=1)
    

This page documents the deployment options for the Charm server, which hosts the backend services that power Charm applications. Charm can be deployed either using the official cloud.charm.sh servers, or by self-hosting on your own infrastructure. This documentation focuses on the self-hosting deployment options, configuration, and management.

For detailed instructions on backing up and restoring Charm accounts and data, see [Backup and Restore](https://deepwiki.com/charmbracelet/charm/6.3-backup-and-restore)
.

Overview of Deployment Options
------------------------------

Charm is designed to be easily deployable in various environments. The server is packaged as a single binary that includes everything needed to run a complete Charm instance.

Sources: [README.md254-260](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L254-L260)
 [docker.md](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1)
 [systemd.md](https://github.com/charmbracelet/charm/blob/aff3071d/systemd.md?plain=1)

Server Configuration
--------------------

The Charm server can be configured using environment variables to customize aspects such as network interfaces, ports, data storage location, and TLS settings.

### Configuration Parameters

| Environment Variable | Description | Default Value |
| --- | --- | --- |
| `CHARM_SERVER_BIND_ADDRESS` | Network interface to listen on | `0.0.0.0` |
| `CHARM_SERVER_HOST` | Hostname to advertise | `localhost` |
| `CHARM_SERVER_SSH_PORT` | SSH server port | `35353` |
| `CHARM_SERVER_HTTP_PORT` | HTTP server port | `35354` |
| `CHARM_SERVER_STATS_PORT` | Stats server port | `35355` |
| `CHARM_SERVER_HEALTH_PORT` | Health server port | `35356` |
| `CHARM_SERVER_DATA_DIR` | Server data directory | `./data` |
| `CHARM_SERVER_USE_TLS` | Whether to use TLS | `false` |
| `CHARM_SERVER_TLS_KEY_FILE` | The TLS key file path | \-  |
| `CHARM_SERVER_TLS_CERT_FILE` | The TLS certificate file path | \-  |
| `CHARM_SERVER_PUBLIC_URL` | Server public URL | \-  |
| `CHARM_SERVER_ENABLE_METRICS` | Whether to enable Prometheus metrics | `false` |
| `CHARM_SERVER_USER_MAX_STORAGE` | Maximum filesystem storage per user | `0` (no limit) |

Sources: [README.md266-282](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L266-L282)

### Client Configuration

When self-hosting Charm, clients need to be configured to connect to your server instead of the default cloud.charm.sh. This is done by setting environment variables on the client side:

For applications that use Charm as a backend (like Skate or Glow), the same environment variables need to be set:

Sources: [README.md284-288](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L284-L288)
 [docs/self-hosting.md12-14](https://github.com/charmbracelet/charm/blob/aff3071d/docs/self-hosting.md?plain=1#L12-L14)
 [README.md10-25](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L10-L25)

Deployment Architecture
-----------------------

The following diagram illustrates the architecture of a deployed Charm server and how clients connect to it:

Sources: [README.md266-282](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L266-L282)
 [docs/self-hosting.md30-42](https://github.com/charmbracelet/charm/blob/aff3071d/docs/self-hosting.md?plain=1#L30-L42)

Deployment Methods
------------------

### Standalone Deployment

The simplest way to deploy Charm is to run the `charm serve` command directly:

This starts a Charm server with the default configuration. For production deployments, you should consider setting appropriate environment variables and ensuring the process runs continuously.

Sources: [README.md254-260](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L254-L260)
 [docs/self-hosting.md10-11](https://github.com/charmbracelet/charm/blob/aff3071d/docs/self-hosting.md?plain=1#L10-L11)

### Docker Deployment

Charm provides official Docker images that can be used to run the server in a container. The database is stored in the `/data` directory, which should be mounted as a volume to persist data across container restarts.

Docker run command:

Docker Compose example:

Sources: [docker.md1-41](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1#L1-L41)

### Systemd Deployment

For Linux systems using systemd, you can create a systemd service to manage the Charm server. This ensures the server starts automatically on boot and restarts if it crashes.

Create a file at `/etc/systemd/system/charm.service`:

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
    

After creating the service file, enable and start the service:

Sources: [systemd.md1-31](https://github.com/charmbracelet/charm/blob/aff3071d/systemd.md?plain=1#L1-L31)

TLS Configuration
-----------------

Charm supports TLS for secure communications between clients and the server. There are two main approaches to implementing TLS:

### Built-in TLS

Charm can handle TLS termination directly. To enable this:

1.  Set `CHARM_SERVER_USE_TLS=true`
2.  Provide paths to your TLS key and certificate:
    *   `CHARM_SERVER_TLS_KEY_FILE=/path/to/key.pem`
    *   `CHARM_SERVER_TLS_CERT_FILE=/path/to/cert.pem`
3.  Set `CHARM_SERVER_HOST` to your domain name

Sources: [README.md301-306](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L301-L306)
 [docs/self-hosting.md44-49](https://github.com/charmbracelet/charm/blob/aff3071d/docs/self-hosting.md?plain=1#L44-L49)

### Using a Reverse Proxy

Alternatively, you can use a reverse proxy like Nginx, Traefik, or Caddy to handle TLS:

When using a reverse proxy:

1.  The reverse proxy handles TLS termination for HTTP traffic
2.  Set `CHARM_SERVER_PUBLIC_URL` to the full public URL of your reverse proxy (e.g., `https://cloud.example.com:35354`)
3.  Ensure SSH traffic on port 35353 is forwarded correctly

Sources: [docs/self-hosting.md30-62](https://github.com/charmbracelet/charm/blob/aff3071d/docs/self-hosting.md?plain=1#L30-L62)
 [docker.md43-48](https://github.com/charmbracelet/charm/blob/aff3071d/docker.md?plain=1#L43-L48)

### TLS Best Practices

1.  Use strong certificates from a trusted CA or Let's Encrypt
2.  Configure appropriate cipher suites and TLS protocols
3.  Test your configuration with tools like SSL Labs
4.  Renew certificates before they expire

Storage Considerations
----------------------

### Data Directory

The Charm server stores all its data in the directory specified by `CHARM_SERVER_DATA_DIR` (default is `./data`). This directory contains:

1.  User accounts and authentication data
2.  Encrypted user data (KV store and filesystem)
3.  Server configuration

Ensure this directory:

*   Is backed up regularly
*   Has appropriate permissions
*   Has sufficient disk space

### Storage Limits

The official Charm Cloud limits users to 1GB of storage per account. When self-hosting, there is no storage limit by default, but you can set one using the `CHARM_SERVER_USER_MAX_STORAGE` environment variable.

Sources: [README.md294-299](https://github.com/charmbracelet/charm/blob/aff3071d/README.md?plain=1#L294-L299)
 [docs/self-hosting.md63-66](https://github.com/charmbracelet/charm/blob/aff3071d/docs/self-hosting.md?plain=1#L63-L66)

Monitoring and Health Checks
----------------------------

Charm provides built-in monitoring capabilities:

1.  A health check endpoint on port 35356
2.  Optional Prometheus metrics on port 35355 (enabled with `CHARM_SERVER_ENABLE_METRICS=true`)

For production deployments, consider:

*   Setting up regular health checks
*   Monitoring the Prometheus metrics
*   Configuring alerts for server issues

Conclusion
----------

Charm is designed to be easy to deploy and maintain, whether as a standalone server, in Docker, or as a systemd service. By following the guidance in this document, you can successfully deploy a Charm server for your own use or for your organization.

For information on backing up and restoring user accounts and data, refer to [Backup and Restore](https://deepwiki.com/charmbracelet/charm/6.3-backup-and-restore)
.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Deployment](https://deepwiki.com/charmbracelet/charm/6-deployment#deployment)
    
*   [Overview of Deployment Options](https://deepwiki.com/charmbracelet/charm/6-deployment#overview-of-deployment-options)
    
*   [Server Configuration](https://deepwiki.com/charmbracelet/charm/6-deployment#server-configuration)
    
*   [Configuration Parameters](https://deepwiki.com/charmbracelet/charm/6-deployment#configuration-parameters)
    
*   [Client Configuration](https://deepwiki.com/charmbracelet/charm/6-deployment#client-configuration)
    
*   [Deployment Architecture](https://deepwiki.com/charmbracelet/charm/6-deployment#deployment-architecture)
    
*   [Deployment Methods](https://deepwiki.com/charmbracelet/charm/6-deployment#deployment-methods)
    
*   [Standalone Deployment](https://deepwiki.com/charmbracelet/charm/6-deployment#standalone-deployment)
    
*   [Docker Deployment](https://deepwiki.com/charmbracelet/charm/6-deployment#docker-deployment)
    
*   [Systemd Deployment](https://deepwiki.com/charmbracelet/charm/6-deployment#systemd-deployment)
    
*   [TLS Configuration](https://deepwiki.com/charmbracelet/charm/6-deployment#tls-configuration)
    
*   [Built-in TLS](https://deepwiki.com/charmbracelet/charm/6-deployment#built-in-tls)
    
*   [Using a Reverse Proxy](https://deepwiki.com/charmbracelet/charm/6-deployment#using-a-reverse-proxy)
    
*   [TLS Best Practices](https://deepwiki.com/charmbracelet/charm/6-deployment#tls-best-practices)
    
*   [Storage Considerations](https://deepwiki.com/charmbracelet/charm/6-deployment#storage-considerations)
    
*   [Data Directory](https://deepwiki.com/charmbracelet/charm/6-deployment#data-directory)
    
*   [Storage Limits](https://deepwiki.com/charmbracelet/charm/6-deployment#storage-limits)
    
*   [Monitoring and Health Checks](https://deepwiki.com/charmbracelet/charm/6-deployment#monitoring-and-health-checks)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/charm/6-deployment#conclusion)
