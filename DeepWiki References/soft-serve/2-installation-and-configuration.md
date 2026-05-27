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

Installation and Configuration
==============================

Relevant source files

*   [.github/workflows/goreleaser.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/goreleaser.yml)
    
*   [.github/workflows/nightly.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/nightly.yml)
    
*   [.goreleaser.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.goreleaser.yml)
    
*   [.nfpm/postinstall.sh](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/postinstall.sh)
    
*   [.nfpm/postremove.sh](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/postremove.sh)
    
*   [.nfpm/soft-serve.conf](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/soft-serve.conf)
    
*   [.nfpm/soft-serve.service](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/soft-serve.service)
    
*   [.nfpm/sysusers.conf](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/sysusers.conf)
    
*   [.nfpm/tmpfiles.conf](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/tmpfiles.conf)
    
*   [Dockerfile](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/Dockerfile)
    
*   [README.md](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1)
    
*   [demo.tape](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/demo.tape)
    
*   [docker.md](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/docker.md?plain=1)
    
*   [pkg/config/config.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go)
    
*   [pkg/config/config\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config_test.go)
    
*   [pkg/config/testdata/config.yaml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/testdata/config.yaml)
    
*   [pkg/web/server.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/web/server.go)
    

This document describes how to install Soft Serve, a self-hostable Git server for the command line, and configure it for your environment. It covers various installation methods, configuration options, and setting up Soft Serve as a system service. For information about repository management and authentication, see [Core Components](https://deepwiki.com/charmbracelet/soft-serve/3-core-components)
.

Installation Methods
--------------------

Soft Serve is distributed as a single binary called `soft` and can be installed through various methods.

### Package Managers

Soft Serve can be installed from several package managers:

Sources: [README.md54-86](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L54-L86)

### Docker Installation

Soft Serve is available as a Docker image. You can run it as follows:

Alternatively, you can use Docker Compose:

**Note**: Ensure you run the image without a TTY, i.e., do not use the `--tty`/`-t` flags.

Sources: [docker.md1-56](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/docker.md?plain=1#L1-L56)
 [Dockerfile1-29](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/Dockerfile#L1-L29)

### Go Installation

You can also install Soft Serve using Go:

Sources: [README.md96-98](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L96-L98)

Installation Flow Diagram
-------------------------

Sources: [README.md54-86](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L54-L86)
 [.goreleaser.yml1-16](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.goreleaser.yml#L1-L16)
 [.github/workflows/goreleaser.yml1-29](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/goreleaser.yml#L1-L29)

Configuration Options
---------------------

After installation, Soft Serve needs to be configured. The configuration can be set through a configuration file (`config.yaml`) or environment variables.

### Data Directory

Soft Serve stores all its data in a directory specified by the `SOFT_SERVE_DATA_PATH` environment variable. If not set, it defaults to a `data` directory in the current working directory.

The data directory contains:

*   Repositories
*   SSH keys
*   Database
*   Configuration file

Sources: [README.md108-120](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L108-L120)

### Configuration File

When you run Soft Serve for the first time, it creates a `config.yaml` file in the data directory. This file contains all the configuration options and their default values. Here's an example configuration file:

Sources: [README.md136-235](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L136-L235)
 [pkg/config/config.go17-379](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L17-L379)

### Configuration Flow

Sources: [pkg/config/config.go130-161](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L130-L161)
 [pkg/config/config.go229-282](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L229-L282)

### Environment Variables

All configuration options can be overridden using environment variables. The environment variables start with `SOFT_SERVE_` followed by the setting name in uppercase. Here are some examples:

*   `SOFT_SERVE_NAME`: The name of the server that will appear in the TUI
*   `SOFT_SERVE_SSH_LISTEN_ADDR`: SSH listen address
*   `SOFT_SERVE_SSH_KEY_PATH`: SSH host key-pair path
*   `SOFT_SERVE_HTTP_LISTEN_ADDR`: HTTP listen address
*   `SOFT_SERVE_HTTP_PUBLIC_URL`: HTTP public URL used for cloning
*   `SOFT_SERVE_GIT_MAX_CONNECTIONS`: The number of simultaneous connections to Git daemon

When you run Soft Serve for the first time, you should set the `SOFT_SERVE_INITIAL_ADMIN_KEYS` environment variable to your SSH authorized key. This will create an admin user with full privileges.

Sources: [README.md239-247](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L239-L247)
 [README.md122-128](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L122-L128)
 [pkg/config/config.go166-211](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L166-L211)

### Database Configuration

Soft Serve supports both SQLite and PostgreSQL for its database. You can configure the database through the `db` section in the configuration file or using environment variables.

#### SQLite (Default)

By default, Soft Serve uses SQLite as the database. The SQLite database file is stored in the data directory.

#### PostgreSQL

To use PostgreSQL, you need to create a database first:

Then configure Soft Serve to use PostgreSQL:

Or using environment variables:

You can include a password in the data source URL: `postgres://myuser:dbpass@localhost:5432/my_soft_serve_db`.

Sources: [README.md248-276](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L248-L276)

### LFS Configuration

Soft Serve supports Git LFS (Large File Storage) with both HTTP and SSH protocols. You can configure LFS in the `lfs` section of the configuration file:

By default, the HTTP protocol is enabled, but the SSH protocol is disabled.

Sources: [README.md278-282](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L278-L282)

System Service Setup
--------------------

You can run Soft Serve as a system service using Systemd. When you install Soft Serve from the Debian/Ubuntu or Fedora/RHEL packages, they come with Systemd service units.

### Systemd Service File

Here's the Systemd service file included in the package:

Sources: [.nfpm/soft-serve.service1-48](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/soft-serve.service#L1-L48)

### Configuration with Systemd

When running as a Systemd service, Soft Serve uses the following configuration:

1.  Data path is set to `/var/lib/soft-serve`
2.  It runs as the `soft-serve` user and group
3.  It reads additional configuration from `/etc/soft-serve.conf`

You can override configuration settings by editing `/etc/soft-serve.conf`:

Sources: [.nfpm/soft-serve.conf1-10](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/soft-serve.conf#L1-L10)

### Package Post-Installation

When installing Soft Serve from a package, the post-installation script sets up the Systemd service:

1.  Creates the `soft-serve` user and group
2.  Creates the `/var/lib/soft-serve` directory
3.  Enables and starts the `soft-serve.service`

Sources: [.nfpm/postinstall.sh1-15](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/postinstall.sh#L1-L15)
 [.nfpm/sysusers.conf1-2](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/sysusers.conf#L1-L2)
 [.nfpm/tmpfiles.conf1-2](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/tmpfiles.conf#L1-L2)

Server Components Architecture
------------------------------

Sources: [README.md105-120](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L105-L120)
 [pkg/config/config.go324-380](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L324-L380)
 [pkg/web/server.go1-30](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/web/server.go#L1-L30)

Starting the Server
-------------------

To start Soft Serve, simply run:

This will start the server with the default configuration or the configuration from the `config.yaml` file in the data directory. You can override configuration using environment variables.

If you've installed Soft Serve from a package, you can start it as a system service:

Sources: [README.md106](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L106-L106)
 [.nfpm/soft-serve.service1-48](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.nfpm/soft-serve.service#L1-L48)

Configuration Options Summary
-----------------------------

| Configuration | Description | Default Value | Environment Variable |
| --- | --- | --- | --- |
| Data Path | Directory where Soft Serve stores data | `data` | `SOFT_SERVE_DATA_PATH` |
| Name | Server name displayed in UI | `Soft Serve` | `SOFT_SERVE_NAME` |
| SSH Listen Address | Address for SSH server | `:23231` | `SOFT_SERVE_SSH_LISTEN_ADDR` |
| SSH Public URL | Public URL for SSH cloning | `ssh://localhost:23231` | `SOFT_SERVE_SSH_PUBLIC_URL` |
| Git Listen Address | Address for Git daemon | `:9418` | `SOFT_SERVE_GIT_LISTEN_ADDR` |
| HTTP Listen Address | Address for HTTP server | `:23232` | `SOFT_SERVE_HTTP_LISTEN_ADDR` |
| HTTP Public URL | Public URL for HTTP cloning | `http://localhost:23232` | `SOFT_SERVE_HTTP_PUBLIC_URL` |
| Database Driver | Database type (`sqlite` or `postgres`) | `sqlite` | `SOFT_SERVE_DB_DRIVER` |
| LFS Enabled | Enable Git LFS support | `true` | `SOFT_SERVE_LFS_ENABLED` |
| Initial Admin Keys | SSH public keys for admin | None | `SOFT_SERVE_INITIAL_ADMIN_KEYS` |

Sources: [pkg/config/config.go330-378](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/config.go#L330-L378)
 [README.md239-247](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/README.md?plain=1#L239-L247)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Installation and Configuration](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#installation-and-configuration)
    
*   [Installation Methods](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#installation-methods)
    
*   [Package Managers](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#package-managers)
    
*   [Docker Installation](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#docker-installation)
    
*   [Go Installation](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#go-installation)
    
*   [Installation Flow Diagram](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#installation-flow-diagram)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#configuration-options)
    
*   [Data Directory](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#data-directory)
    
*   [Configuration File](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#configuration-file)
    
*   [Configuration Flow](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#configuration-flow)
    
*   [Environment Variables](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#environment-variables)
    
*   [Database Configuration](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#database-configuration)
    
*   [SQLite (Default)](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#sqlite-default)
    
*   [PostgreSQL](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#postgresql)
    
*   [LFS Configuration](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#lfs-configuration)
    
*   [System Service Setup](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#system-service-setup)
    
*   [Systemd Service File](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#systemd-service-file)
    
*   [Configuration with Systemd](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#configuration-with-systemd)
    
*   [Package Post-Installation](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#package-post-installation)
    
*   [Server Components Architecture](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#server-components-architecture)
    
*   [Starting the Server](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#starting-the-server)
    
*   [Configuration Options Summary](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration#configuration-options-summary)
