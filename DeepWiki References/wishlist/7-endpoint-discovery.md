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

Endpoint Discovery
==================

Relevant source files

*   [README.md](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1)
    
*   [cmd/wishlist/main.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go)
    
*   [srv/srv.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/srv/srv.go)
    
*   [tailscale/tailscale.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/tailscale/tailscale.go)
    
*   [zeroconf/zeroconf.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/zeroconf/zeroconf.go)
    

Purpose and Scope
-----------------

This document details the endpoint discovery mechanisms in Wishlist, which allow the system to automatically find and configure SSH endpoints from various sources. Wishlist supports three primary discovery methods: Tailscale integration, Zeroconf/mDNS, and SRV records. This page covers how each discovery method works, how they're integrated into the application, and how the discovered endpoints can be customized with hints.

Overview of Endpoint Discovery
------------------------------

Endpoint discovery in Wishlist enables automatic detection of SSH servers on a network without requiring manual configuration for each server. This mechanism allows users to quickly connect to servers that may have dynamic addresses or that are newly provisioned.

### Discovery Flow

Sources: [cmd/wishlist/main.go340-367](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L340-L367)
 [cmd/wishlist/main.go259-306](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L259-L306)

Discovery Methods
-----------------

Wishlist provides three distinct methods for endpoint discovery, each suited for different network environments.

### Tailscale Integration

Tailscale integration allows Wishlist to discover SSH endpoints from a Tailscale network (tailnet). It queries the Tailscale API to retrieve a list of devices and adds them as SSH endpoints.

Sources: [tailscale/tailscale.go18-69](https://github.com/charmbracelet/wishlist/blob/5fe34a29/tailscale/tailscale.go#L18-L69)
 [tailscale/tailscale.go71-99](https://github.com/charmbracelet/wishlist/blob/5fe34a29/tailscale/tailscale.go#L71-L99)

### Zeroconf/mDNS Discovery

Zeroconf (also known as mDNS, Bonjour, or Avahi) discovery allows Wishlist to discover SSH services advertised on the local network.

Sources: [zeroconf/zeroconf.go18-45](https://github.com/charmbracelet/wishlist/blob/5fe34a29/zeroconf/zeroconf.go#L18-L45)

### SRV Record Discovery

SRV record discovery allows Wishlist to discover SSH services through DNS SRV records. It can also use TXT records to customize the endpoint names.

Sources: [srv/srv.go20-33](https://github.com/charmbracelet/wishlist/blob/5fe34a29/srv/srv.go#L20-L33)
 [srv/srv.go35-60](https://github.com/charmbracelet/wishlist/blob/5fe34a29/srv/srv.go#L35-L60)

Endpoint Discovery Configuration
--------------------------------

Each discovery method can be controlled through command-line flags in the Wishlist CLI.

### Command-line Flags

| Flag | Description | Default |
| --- | --- | --- |
| `--tailscale.net` | Tailscale tailnet name | ""  |
| `--tailscale.key` | Tailscale API key | ""  |
| `--tailscale.client.id` | Tailscale OAuth client ID | ""  |
| `--tailscale.client.secret` | Tailscale OAuth client secret | ""  |
| `--zeroconf.enabled` | Enable Zeroconf discovery | false |
| `--zeroconf.domain` | Domain for Zeroconf discovery | ""  |
| `--zeroconf.timeout` | Timeout for Zeroconf discovery | 1s  |
| `--srv.domain` | Domain(s) for SRV record discovery | \[\] |

Sources: [cmd/wishlist/main.go214-221](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L214-L221)

### Environment Variables

For Tailscale authentication, you can also use environment variables:

| Environment Variable | Corresponding Flag |
| --- | --- |
| `TAILSCALE_KEY` | `--tailscale.key` |
| `TAILSCALE_CLIENT_ID` | `--tailscale.client.id` |
| `TAILSCALE_CLIENT_SECRET` | `--tailscale.client.secret` |

Sources: [cmd/wishlist/main.go63-74](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L63-L74)

Endpoint Discovery Data Flow
----------------------------

The endpoint discovery process in Wishlist follows a well-defined flow from discovery to configuration.

Sources: [cmd/wishlist/main.go259-306](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L259-L306)

Tailscale Integration Details
-----------------------------

The Tailscale integration queries the Tailscale API to retrieve a list of devices in a tailnet and converts them to Wishlist endpoints.

### Authentication Methods

Wishlist supports two methods for authenticating with the Tailscale API:

1.  **API Key**: A bearer token provided through the `--tailscale.key` flag or `TAILSCALE_KEY` environment variable.
2.  **OAuth Client Credentials**: Client ID and secret provided through the `--tailscale.client.id` and `--tailscale.client.secret` flags or corresponding environment variables.

Sources: [tailscale/tailscale.go71-90](https://github.com/charmbracelet/wishlist/blob/5fe34a29/tailscale/tailscale.go#L71-L90)

### Endpoint Creation

The Tailscale integration creates endpoints with the following properties:

*   **Name**: The device name (up to the first period)
*   **Address**: The first IP address of the device with port 22

Sources: [tailscale/tailscale.go59-65](https://github.com/charmbracelet/wishlist/blob/5fe34a29/tailscale/tailscale.go#L59-L65)

Zeroconf/mDNS Integration Details
---------------------------------

The Zeroconf integration discovers SSH services advertised on the local network through mDNS.

### Service Discovery

The Zeroconf integration looks for services with the `_ssh._tcp` service type. It uses the `grandcat/zeroconf` library to perform the discovery.

Sources: [zeroconf/zeroconf.go16-32](https://github.com/charmbracelet/wishlist/blob/5fe34a29/zeroconf/zeroconf.go#L16-L32)

### Endpoint Creation

The Zeroconf integration creates endpoints with the following properties:

*   **Name**: The hostname of the service
*   **Address**: The hostname and port of the service

Sources: [zeroconf/zeroconf.go34-42](https://github.com/charmbracelet/wishlist/blob/5fe34a29/zeroconf/zeroconf.go#L34-L42)

SRV Record Integration Details
------------------------------

The SRV record integration discovers SSH services through DNS SRV records in a specified domain.

### Record Lookup

The SRV integration looks for `_ssh._tcp` SRV records in the specified domain using the Go standard library's DNS resolver.

Sources: [srv/srv.go22-29](https://github.com/charmbracelet/wishlist/blob/5fe34a29/srv/srv.go#L22-L29)

### Custom Naming with TXT Records

To customize the names of discovered endpoints, you can create TXT records in a special format:

    wishlist.name full.address:22=thename
    

This format allows you to specify a custom name for a specific address and port combination.

Sources: [srv/srv.go42-53](https://github.com/charmbracelet/wishlist/blob/5fe34a29/srv/srv.go#L42-L53)

### Endpoint Creation

The SRV integration creates endpoints with the following properties:

*   **Name**: Either the hostname from the SRV record or a custom name from a matching TXT record
*   **Address**: The hostname and port from the SRV record

Sources: [srv/srv.go35-58](https://github.com/charmbracelet/wishlist/blob/5fe34a29/srv/srv.go#L35-L58)

Hints System
------------

The hints system allows you to apply additional configuration to discovered endpoints based on glob patterns.

### Hint Definition

Hints are defined in the YAML configuration file under the `hints` key. Each hint includes:

*   A glob pattern to match against endpoint names
*   Configuration values to apply to matching endpoints

Sources: [cmd/wishlist/main.go259-306](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L259-L306)
 [README.md216-234](https://github.com/charmbracelet/wishlist/blob/5fe34a29/README.md?plain=1#L216-L234)

### Pattern Matching

Hints use glob patterns to match against endpoint names. For example:

*   `*` matches any number of characters
*   `?` matches a single character
*   `[abc]` matches any of a, b, or c

A hint with the pattern `server*` would match endpoints named "server1", "server-dev", etc.

Sources: [cmd/wishlist/main.go261-265](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L261-L265)

### Configuration Values

Hints can modify various aspects of an endpoint's configuration, including:

*   Port
*   User
*   ForwardAgent
*   RequestTTY
*   RemoteCommand
*   Description
*   Link details
*   ProxyJump
*   Environment variables
*   Authentication preferences
*   Identity files
*   Connection timeout

Sources: [cmd/wishlist/main.go270-302](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L270-L302)

Endpoint Refresh
----------------

In server mode, Wishlist can periodically refresh the discovered endpoints. This is useful for environments where endpoints may come and go dynamically.

### Refresh Interval

The refresh interval is configured with the `--endpoints.refresh.interval` flag. When set to a value greater than 0, Wishlist will refresh the endpoints at the specified interval.

Sources: [cmd/wishlist/main.go142-164](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L142-L164)

Integration with Configuration System
-------------------------------------

The endpoint discovery system integrates with Wishlist's configuration system, allowing discovered endpoints to be combined with manually configured endpoints.

### Configuration Loading

Wishlist checks for configuration files in various locations, including the `.wishlist` directory and the user's SSH config. The discovered endpoints are added to the configuration after applying hints.

Sources: [cmd/wishlist/main.go308-329](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L308-L329)
 [cmd/wishlist/main.go236-257](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L236-L257)

### YAML Configuration

When using a YAML configuration file, discovered endpoints are added to the endpoints defined in the configuration file.

Sources: [cmd/wishlist/main.go369-382](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L369-L382)

### SSH Configuration

When using an SSH configuration file, discovered endpoints are matched against hosts defined in the configuration file using the hints system.

Sources: [cmd/wishlist/main.go385-392](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go#L385-L392)

Summary
-------

Wishlist's endpoint discovery system provides a flexible and powerful way to automatically discover SSH endpoints from various sources. By combining different discovery methods and applying hints, users can create a seamless and dynamic SSH directory that adapts to changes in their network environment.

The three discovery methods—Tailscale, Zeroconf, and SRV records—cover a wide range of use cases, from local network discovery to cloud-based deployments. The hints system allows for customization of discovered endpoints, ensuring that they have the correct configuration for your environment.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Endpoint Discovery](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#endpoint-discovery)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#purpose-and-scope)
    
*   [Overview of Endpoint Discovery](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#overview-of-endpoint-discovery)
    
*   [Discovery Flow](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#discovery-flow)
    
*   [Discovery Methods](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#discovery-methods)
    
*   [Tailscale Integration](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#tailscale-integration)
    
*   [Zeroconf/mDNS Discovery](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#zeroconfmdns-discovery)
    
*   [SRV Record Discovery](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#srv-record-discovery)
    
*   [Endpoint Discovery Configuration](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#endpoint-discovery-configuration)
    
*   [Command-line Flags](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#command-line-flags)
    
*   [Environment Variables](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#environment-variables)
    
*   [Endpoint Discovery Data Flow](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#endpoint-discovery-data-flow)
    
*   [Tailscale Integration Details](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#tailscale-integration-details)
    
*   [Authentication Methods](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#authentication-methods)
    
*   [Endpoint Creation](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#endpoint-creation)
    
*   [Zeroconf/mDNS Integration Details](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#zeroconfmdns-integration-details)
    
*   [Service Discovery](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#service-discovery)
    
*   [Endpoint Creation](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#endpoint-creation-1)
    
*   [SRV Record Integration Details](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#srv-record-integration-details)
    
*   [Record Lookup](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#record-lookup)
    
*   [Custom Naming with TXT Records](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#custom-naming-with-txt-records)
    
*   [Endpoint Creation](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#endpoint-creation-2)
    
*   [Hints System](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#hints-system)
    
*   [Hint Definition](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#hint-definition)
    
*   [Pattern Matching](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#pattern-matching)
    
*   [Configuration Values](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#configuration-values)
    
*   [Endpoint Refresh](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#endpoint-refresh)
    
*   [Refresh Interval](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#refresh-interval)
    
*   [Integration with Configuration System](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#integration-with-configuration-system)
    
*   [Configuration Loading](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#configuration-loading)
    
*   [YAML Configuration](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#yaml-configuration)
    
*   [SSH Configuration](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#ssh-configuration)
    
*   [Summary](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery#summary)
