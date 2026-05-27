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

Development
===========

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/build.yml)
    
*   [.github/workflows/goreleaser.yml](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/goreleaser.yml)
    
*   [go.mod](https://github.com/charmbracelet/charm/blob/aff3071d/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/charm/blob/aff3071d/go.sum)
    

This document provides information for developers who want to contribute to or work with the Charm codebase. It covers development environment setup, build processes, testing procedures, continuous integration workflows, and release processes. For information about the architectural design of Charm, see [Architecture Overview](https://deepwiki.com/charmbracelet/charm/1.2-architecture-overview)
.

Development Environment Setup
-----------------------------

### Prerequisites

*   Go 1.17 or later (as specified in [go.mod3](https://github.com/charmbracelet/charm/blob/aff3071d/go.mod#L3-L3)
    )
*   Git
*   SSH (for testing authentication functionality)

### Local Development Setup

1.  Clone the repository:

2.  Install dependencies:

The project uses Go modules for dependency management. Key dependencies include:

| Dependency | Purpose |
| --- | --- |
| github.com/charmbracelet/bubbles | TUI components |
| github.com/charmbracelet/bubbletea | TUI framework |
| github.com/charmbracelet/keygen | SSH key generation and management |
| github.com/charmbracelet/lipgloss | Terminal styling |
| github.com/charmbracelet/ssh | SSH server functionality |
| github.com/dgraph-io/badger/v3 | Key-value storage |
| github.com/spf13/cobra | CLI commands |
| modernc.org/sqlite | Database operations |

Sources: [go.mod5-36](https://github.com/charmbracelet/charm/blob/aff3071d/go.mod#L5-L36)

Build Process
-------------

### Development Workflow

Sources: [.github/workflows/build.yml](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/build.yml)
 [.github/workflows/goreleaser.yml](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/goreleaser.yml)

### Building Locally

To build the project:

This creates the `charm` executable in the current directory.

For cross-compilation (example for Linux ARM64):

### Go Build Tags

Certain components of Charm may use build tags to conditionally include code. Always ensure you're using the appropriate build tags if specific functionality is required.

Sources: [go.mod](https://github.com/charmbracelet/charm/blob/aff3071d/go.mod)

Testing
-------

### Running Tests

Run the full test suite:

For verbose output:

To test a specific package:

Sources: [go.mod](https://github.com/charmbracelet/charm/blob/aff3071d/go.mod)

Continuous Integration and Deployment
-------------------------------------

### Build Pipeline Diagram

Sources: [.github/workflows/build.yml](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/build.yml)
 [.github/workflows/goreleaser.yml](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/goreleaser.yml)

### Build Workflow

The build workflow runs on every push and pull request. It is defined in [.github/workflows/build.yml](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/build.yml)
 and uses reusable workflows from the `charmbracelet/meta` repository:

The build workflow performs standard checks like compilation, testing, and linting.

Sources: [.github/workflows/build.yml1-12](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/build.yml#L1-L12)

### Release Process

The release process is triggered automatically when a tag matching `v*.*.*` is pushed. It uses GoReleaser to:

1.  Build binaries for multiple platforms
2.  Create packages for various package managers
3.  Generate Docker images
4.  Upload artifacts to GitHub Releases
5.  Publish packages to repositories
6.  Push Docker images to Docker Hub

The configuration is defined in [.github/workflows/goreleaser.yml](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/goreleaser.yml)

To create a new release:

Sources: [.github/workflows/goreleaser.yml1-22](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/goreleaser.yml#L1-L22)

Code Structure for Development
------------------------------

Sources: [go.mod](https://github.com/charmbracelet/charm/blob/aff3071d/go.mod)

Contributing Guidelines
-----------------------

### Pull Request Process

1.  Fork the repository
2.  Create a feature branch (`git checkout -b feature/my-feature`)
3.  Make your changes
4.  Add or update tests
5.  Run tests locally (`go test ./...`)
6.  Commit your changes with descriptive messages
7.  Push to your fork (`git push origin feature/my-feature`)
8.  Create a pull request to the main repository

### Code Style

*   Follow standard Go code style and conventions
*   Use `gofmt` or `goimports` to format your code
*   Add comments for exported functions and types
*   Write clear commit messages

### Versioning

Charm follows semantic versioning (MAJOR.MINOR.PATCH). Version tags trigger the release workflow:

The tag format `v*.*.*` is recognized by the GoReleaser workflow as specified in [.github/workflows/goreleaser.yml3-6](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/goreleaser.yml#L3-L6)

Sources: [.github/workflows/goreleaser.yml1-22](https://github.com/charmbracelet/charm/blob/aff3071d/.github/workflows/goreleaser.yml#L1-L22)

Development Tools
-----------------

### Recommended Development Environment

For the best development experience with Charm, the following tools are recommended:

| Tool | Purpose |
| --- | --- |
| Visual Studio Code | Editor with Go support |
| Go extension for VS Code | Syntax highlighting, code navigation |
| Delve | Go debugger |
| golangci-lint | Go linter aggregator |
| mockgen | For generating mocks in tests |

### Debugging

For debugging the Charm client or server:

1.  Build with debug information: `go build -gcflags="all=-N -l"`
2.  Use Delve to attach to the process: `dlv attach [PID]`
3.  Or start the application directly with Delve: `dlv debug`

Sources: [go.mod](https://github.com/charmbracelet/charm/blob/aff3071d/go.mod)

Documentation
-------------

When making changes, ensure you update any relevant documentation. The main sources of documentation are:

1.  Code comments
2.  README files
3.  Wiki pages

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development](https://deepwiki.com/charmbracelet/charm/7-development#development)
    
*   [Development Environment Setup](https://deepwiki.com/charmbracelet/charm/7-development#development-environment-setup)
    
*   [Prerequisites](https://deepwiki.com/charmbracelet/charm/7-development#prerequisites)
    
*   [Local Development Setup](https://deepwiki.com/charmbracelet/charm/7-development#local-development-setup)
    
*   [Build Process](https://deepwiki.com/charmbracelet/charm/7-development#build-process)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/charm/7-development#development-workflow)
    
*   [Building Locally](https://deepwiki.com/charmbracelet/charm/7-development#building-locally)
    
*   [Go Build Tags](https://deepwiki.com/charmbracelet/charm/7-development#go-build-tags)
    
*   [Testing](https://deepwiki.com/charmbracelet/charm/7-development#testing)
    
*   [Running Tests](https://deepwiki.com/charmbracelet/charm/7-development#running-tests)
    
*   [Continuous Integration and Deployment](https://deepwiki.com/charmbracelet/charm/7-development#continuous-integration-and-deployment)
    
*   [Build Pipeline Diagram](https://deepwiki.com/charmbracelet/charm/7-development#build-pipeline-diagram)
    
*   [Build Workflow](https://deepwiki.com/charmbracelet/charm/7-development#build-workflow)
    
*   [Release Process](https://deepwiki.com/charmbracelet/charm/7-development#release-process)
    
*   [Code Structure for Development](https://deepwiki.com/charmbracelet/charm/7-development#code-structure-for-development)
    
*   [Contributing Guidelines](https://deepwiki.com/charmbracelet/charm/7-development#contributing-guidelines)
    
*   [Pull Request Process](https://deepwiki.com/charmbracelet/charm/7-development#pull-request-process)
    
*   [Code Style](https://deepwiki.com/charmbracelet/charm/7-development#code-style)
    
*   [Versioning](https://deepwiki.com/charmbracelet/charm/7-development#versioning)
    
*   [Development Tools](https://deepwiki.com/charmbracelet/charm/7-development#development-tools)
    
*   [Recommended Development Environment](https://deepwiki.com/charmbracelet/charm/7-development#recommended-development-environment)
    
*   [Debugging](https://deepwiki.com/charmbracelet/charm/7-development#debugging)
    
*   [Documentation](https://deepwiki.com/charmbracelet/charm/7-development#documentation)
