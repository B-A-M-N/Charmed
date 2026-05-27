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

Deployment
==========

Relevant source files

*   [.github/workflows/goreleaser.yml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/goreleaser.yml)
    
*   [.gitignore](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.gitignore)
    
*   [.goreleaser.yml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.goreleaser.yml)
    
*   [Dockerfile](https://github.com/charmbracelet/wishlist/blob/5fe34a29/Dockerfile)
    

This document provides comprehensive information about deploying Wishlist in various environments. It covers binary installations, containerization, and deployment considerations. For information about configuring Wishlist after deployment, see [Configuration System](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system)
.

Deployment Options Overview
---------------------------

Wishlist can be deployed in several ways, depending on your requirements and environment:

1.  Binary installation via package managers
2.  Direct binary downloads
3.  Docker containerization
4.  Building from source

Sources: [.goreleaser.yml1-15](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.goreleaser.yml#L1-L15)
 [.github/workflows/goreleaser.yml1-30](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/goreleaser.yml#L1-L30)
 [Dockerfile1-4](https://github.com/charmbracelet/wishlist/blob/5fe34a29/Dockerfile#L1-L4)

Binary Installation
-------------------

Wishlist is distributed as a binary through various package managers and release channels, handled by GoReleaser as configured in the project.

### Package Managers

Wishlist can be installed using the following package managers:

| Package Manager | Platform | Installation Command |
| --- | --- | --- |
| Homebrew | macOS/Linux | `brew install charmbracelet/tap/wishlist` |
| Arch User Repository | Arch Linux | `yay -S wishlist` |
| APT | Debian/Ubuntu | `apt install wishlist` (after adding repository) |
| RPM | Fedora/RHEL/CentOS | `dnf install wishlist` (after adding repository) |

### Direct Download

You can download pre-compiled binaries for various platforms directly from the [GitHub Releases page](https://github.com/charmbracelet/wishlist/blob/5fe34a29/GitHub%20Releases%20page)

Sources: [.goreleaser.yml1-15](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.goreleaser.yml#L1-L15)
 [.github/workflows/goreleaser.yml1-30](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/goreleaser.yml#L1-L30)

Container Deployment
--------------------

Wishlist is available as a Docker container, which provides a consistent and isolated runtime environment.

### Using the Official Docker Image

The official Docker image is published to Docker Hub and can be used as follows:

### Dockerfile Details

The Dockerfile uses Google's distroless static image as the base, which provides a minimal container with just the essentials needed to run the Wishlist binary:

Sources: [Dockerfile1-4](https://github.com/charmbracelet/wishlist/blob/5fe34a29/Dockerfile#L1-L4)

### Docker Compose Example

For more complex deployments, a Docker Compose configuration can be used:

Release Process
---------------

Wishlist uses GitHub Actions and GoReleaser for its release process. When a version tag is pushed, the release workflow automatically builds, packages, and distributes Wishlist to various channels.

Sources: [.github/workflows/goreleaser.yml1-30](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/goreleaser.yml#L1-L30)
 [.goreleaser.yml1-15](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.goreleaser.yml#L1-L15)

Deployment Considerations
-------------------------

### Security

When deploying Wishlist, consider the following security aspects:

1.  **SSH Key Management**: Ensure proper management of SSH keys used for authentication
2.  **Network Configuration**: Properly configure firewalls to control access to the Wishlist server
3.  **Permissions**: When running Wishlist as a service, use appropriate system user permissions

### Configuration Management

Wishlist configuration can be managed in several ways:

1.  **Configuration file**: Place `wishlist.yaml` in the appropriate configuration directory
2.  **Environment variables**: Configure Wishlist behavior through environment variables
3.  **SSH config integration**: Leverage existing SSH configuration

For containerized deployments, mount the configuration file as a volume to make it persistent and easily updatable.

### Running as a Service

For persistent deployments, Wishlist can be configured as a system service:

#### Systemd Service Example

Cross-Platform Support
----------------------

Wishlist supports multiple platforms through its release process:

| Platform | Architecture | Package Format |
| --- | --- | --- |
| Linux | amd64, arm64, arm | Binary, DEB, RPM |
| macOS | amd64, arm64 | Binary, Homebrew |
| Windows | amd64 | Binary (.zip) |
| FreeBSD | amd64 | Binary |

Sources: [.goreleaser.yml1-15](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.goreleaser.yml#L1-L15)

Continuous Integration and Deployment
-------------------------------------

The release process is automated through GitHub Actions, which is triggered when a version tag is pushed. This ensures consistent and reliable releases across all distribution channels.

Sources: [.github/workflows/goreleaser.yml1-30](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/goreleaser.yml#L1-L30)

Build from Source
-----------------

For development purposes or custom builds, Wishlist can be built from source:

This produces a binary tailored to your specific environment and requirements.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Deployment](https://deepwiki.com/charmbracelet/wishlist/9-deployment#deployment)
    
*   [Deployment Options Overview](https://deepwiki.com/charmbracelet/wishlist/9-deployment#deployment-options-overview)
    
*   [Binary Installation](https://deepwiki.com/charmbracelet/wishlist/9-deployment#binary-installation)
    
*   [Package Managers](https://deepwiki.com/charmbracelet/wishlist/9-deployment#package-managers)
    
*   [Direct Download](https://deepwiki.com/charmbracelet/wishlist/9-deployment#direct-download)
    
*   [Container Deployment](https://deepwiki.com/charmbracelet/wishlist/9-deployment#container-deployment)
    
*   [Using the Official Docker Image](https://deepwiki.com/charmbracelet/wishlist/9-deployment#using-the-official-docker-image)
    
*   [Dockerfile Details](https://deepwiki.com/charmbracelet/wishlist/9-deployment#dockerfile-details)
    
*   [Docker Compose Example](https://deepwiki.com/charmbracelet/wishlist/9-deployment#docker-compose-example)
    
*   [Release Process](https://deepwiki.com/charmbracelet/wishlist/9-deployment#release-process)
    
*   [Deployment Considerations](https://deepwiki.com/charmbracelet/wishlist/9-deployment#deployment-considerations)
    
*   [Security](https://deepwiki.com/charmbracelet/wishlist/9-deployment#security)
    
*   [Configuration Management](https://deepwiki.com/charmbracelet/wishlist/9-deployment#configuration-management)
    
*   [Running as a Service](https://deepwiki.com/charmbracelet/wishlist/9-deployment#running-as-a-service)
    
*   [Systemd Service Example](https://deepwiki.com/charmbracelet/wishlist/9-deployment#systemd-service-example)
    
*   [Cross-Platform Support](https://deepwiki.com/charmbracelet/wishlist/9-deployment#cross-platform-support)
    
*   [Continuous Integration and Deployment](https://deepwiki.com/charmbracelet/wishlist/9-deployment#continuous-integration-and-deployment)
    
*   [Build from Source](https://deepwiki.com/charmbracelet/wishlist/9-deployment#build-from-source)
