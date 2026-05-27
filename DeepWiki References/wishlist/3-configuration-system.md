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

Configuration System
====================

Relevant source files

*   [\_example/config](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/config)
    
*   [\_example/config.yaml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/config.yaml)
    
*   [cmd/wishlist/main\_test.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main_test.go)
    
*   [config.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go)
    

The Configuration System in Wishlist provides a flexible way to define SSH endpoints, user authentication, and system behavior. This page documents the structure of the configuration system, available configuration sources, and how configuration options are processed.

For details about how SSH configuration files are specifically integrated, see [SSH Config Integration](https://deepwiki.com/charmbracelet/wishlist/3.1-ssh-config-integration)
. For information on endpoint discovery and configuration, see [Endpoint Configuration](https://deepwiki.com/charmbracelet/wishlist/3.2-endpoint-configuration)
.

Overview
--------

Wishlist's configuration system allows you to define:

1.  Server listening address and port
2.  SSH endpoints to display and connect to
3.  Authorized users and their public keys
4.  Discovery hints for automatically detected endpoints
5.  Metrics collection settings

The configuration can be loaded from various file formats and locations, with YAML and SSH config formats being the primary supported formats.

Sources: [config.go146-158](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L146-L158)
 [\_example/config.yaml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/config.yaml)
 [\_example/config](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/config)

Configuration Structure
-----------------------

The configuration system is centered around the `Config` struct, which contains all settings needed to run Wishlist.

Sources: [config.go21-36](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L21-L36)
 [config.go38-56](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L38-L56)
 [config.go58-75](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L58-L75)
 [config.go146-158](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L146-L158)
 [config.go160-164](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L160-L164)
 [config.go166-171](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L166-L171)

Configuration Sources
---------------------

Wishlist looks for configuration in multiple locations, with the following precedence:

1.  Explicitly specified configuration file path
2.  Local directory configuration (`.wishlist/config.yaml`, `.wishlist/config.yml`, `.wishlist/config`)
3.  User configuration directory (`wishlist.yaml`, `wishlist.yml`, `wishlist`)
4.  Standard SSH configuration files (`~/.ssh/config`, `/etc/ssh/ssh_config`)

Sources: [cmd/wishlist/main\_test.go124-160](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main_test.go#L124-L160)

YAML Configuration
------------------

YAML configuration provides a comprehensive way to define all aspects of Wishlist's behavior. The example below shows the main configuration options:

| Section | Description |
| --- | --- |
| `listen` | Address to listen on (default: `0.0.0.0`) |
| `port` | Port to listen on (default: `22` or `2222`, whichever is open) |
| `endpoints` | List of SSH endpoints to display |
| `hints` | Rules for enhancing discovered endpoints |
| `users` | Authorized users with their public keys |
| `metrics` | Prometheus metrics configuration |

Key elements of the YAML configuration:

1.  **Server Settings**: Basic server configuration
    
2.  **Endpoint Configuration**: Defining SSH endpoints
    
3.  **User Authentication**: Defining allowed users
    

Sources: [\_example/config.yaml5-130](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/config.yaml#L5-L130)

Endpoint Configuration
----------------------

Endpoints are the core element of Wishlist, representing SSH connections that can be established. Each endpoint has a rich set of configuration options that control its behavior:

| Option | Description | Default |
| --- | --- | --- |
| `name` | Display name of the endpoint | Required |
| `address` | SSH address in `host:port` format | Required |
| `user` | Username for SSH connection | Current user |
| `forward_agent` | Whether to forward the SSH agent | `false` |
| `request_tty` | Whether to request a TTY | `true` if no `remote_command` |
| `remote_command` | Command to run on the remote server | Empty (request shell) |
| `description` | Description text for the endpoint | Empty |
| `link` | URL link associated with the endpoint | Empty |
| `proxy_jump` | Proxy server to connect through | Empty |
| `send_env` | Environment variables to send | `["LC_*", "LANG"]` |
| `set_env` | Environment variables to set | Empty |
| `identity_files` | SSH identity files to use | Empty |
| `connect_timeout` | Connection timeout | System default |

Environment variables are handled through the `Environment()` method, which processes both `SendEnv` and `SetEnv` configurations:

Sources: [config.go38-56](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L38-L56)
 [config.go79-129](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L79-L129)
 [\_example/config.yaml14-74](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/config.yaml#L14-L74)

Endpoint Hints
--------------

Endpoint hints provide a powerful way to customize discovered endpoints. When endpoints are discovered through methods like Tailscale, Zeroconf/mDNS, or SRV records, hints can be applied to modify their properties based on pattern matching.

A hint can set many properties of an endpoint, including:

*   SSH port
*   User name
*   Description
*   Authentication settings
*   Remote command to execute
*   Proxy jump configuration
*   And more

Sources: [config.go58-75](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L58-L75)
 [\_example/config.yaml75-108](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/config.yaml#L75-L108)
 [cmd/wishlist/main\_test.go162-221](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main_test.go#L162-L221)

User Authentication
-------------------

User authentication is configured through the `users` section in the YAML configuration. This defines which users are allowed to access the Wishlist SSH server:

Each user entry contains:

*   `name`: The username for authentication
*   `public-keys`: List of public keys allowed for this user (in OpenSSH authorized\_keys format)

Sources: [config.go160-164](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L160-L164)
 [\_example/config.yaml110-119](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/config.yaml#L110-L119)

Metrics Configuration
---------------------

Wishlist can expose Prometheus metrics through a dedicated HTTP endpoint. This is configured through the `metrics` section:

The metrics configuration consists of:

*   `enabled`: Whether metrics collection is enabled
*   `name`: Application name used in metrics labels
*   `address`: Address to bind the metrics HTTP server

Sources: [config.go166-171](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go#L166-L171)
 [\_example/config.yaml121-130](https://github.com/charmbracelet/wishlist/blob/5fe34a29/_example/config.yaml#L121-L130)

Configuration Processing Flow
-----------------------------

The overall flow of configuration processing in Wishlist follows this sequence:

Sources: [cmd/wishlist/main\_test.go13-123](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main_test.go#L13-L123)

Conclusion
----------

The Wishlist configuration system provides a flexible foundation for defining SSH endpoints, server behavior, and authentication. By supporting both YAML and standard SSH configuration formats, it offers familiarity to users while providing advanced features like endpoint hints and integrated metrics.

The design emphasizes a clear separation between configuration definition (in the `Config` struct) and the behavior it controls, allowing for easy extension and modification of the system's capabilities.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Configuration System](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#configuration-system)
    
*   [Overview](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#overview)
    
*   [Configuration Structure](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#configuration-structure)
    
*   [Configuration Sources](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#configuration-sources)
    
*   [YAML Configuration](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#yaml-configuration)
    
*   [Endpoint Configuration](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#endpoint-configuration)
    
*   [Endpoint Hints](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#endpoint-hints)
    
*   [User Authentication](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#user-authentication)
    
*   [Metrics Configuration](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#metrics-configuration)
    
*   [Configuration Processing Flow](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#configuration-processing-flow)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system#conclusion)
