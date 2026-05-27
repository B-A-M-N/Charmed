Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/skate](https://github.com/charmbracelet/skate "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 21 April 2025 ([098604](https://github.com/charmbracelet/skate/commits/09860417)
)

*   [Overview](https://deepwiki.com/charmbracelet/skate/1-overview)
    
*   [Installation](https://deepwiki.com/charmbracelet/skate/2-installation)
    
*   [User Guide](https://deepwiki.com/charmbracelet/skate/3-user-guide)
    
*   [Basic Commands](https://deepwiki.com/charmbracelet/skate/3.1-basic-commands)
    
*   [Working with Multiple Databases](https://deepwiki.com/charmbracelet/skate/3.2-working-with-multiple-databases)
    
*   [Advanced Features](https://deepwiki.com/charmbracelet/skate/3.3-advanced-features)
    
*   [Architecture](https://deepwiki.com/charmbracelet/skate/4-architecture)
    
*   [Command Processing](https://deepwiki.com/charmbracelet/skate/4.1-command-processing)
    
*   [Data Storage](https://deepwiki.com/charmbracelet/skate/4.2-data-storage)
    
*   [Terminal Interface](https://deepwiki.com/charmbracelet/skate/4.3-terminal-interface)
    
*   [Development](https://deepwiki.com/charmbracelet/skate/5-development)
    
*   [Build Process](https://deepwiki.com/charmbracelet/skate/5.1-build-process)
    
*   [Release Process](https://deepwiki.com/charmbracelet/skate/5.2-release-process)
    
*   [Code Quality](https://deepwiki.com/charmbracelet/skate/5.3-code-quality)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/skate/5.4-dependency-management)
    
*   [Reference](https://deepwiki.com/charmbracelet/skate/6-reference)
    
*   [Command Line Interface](https://deepwiki.com/charmbracelet/skate/6.1-command-line-interface)
    
*   [API Functions](https://deepwiki.com/charmbracelet/skate/6.2-api-functions)
    
*   [License](https://deepwiki.com/charmbracelet/skate/7-license)
    

Menu

Development
===========

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/skate/blob/09860417/.github/workflows/build.yml)
    
*   [go.mod](https://github.com/charmbracelet/skate/blob/09860417/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/skate/blob/09860417/go.sum)
    

This page provides guidance for developers who want to contribute to or modify the Skate project. It covers the development environment, build process, release pipeline, code quality practices, and dependency management for the Skate codebase. For information about Skate's architecture and design, see [Architecture](https://deepwiki.com/charmbracelet/skate/4-architecture)
.

Development Environment Setup
-----------------------------

To begin development on Skate, you'll need:

1.  **Go 1.23+** - Skate uses Go 1.23.0 with toolchain 1.24.1
2.  **Git** - For source code management
3.  **Editor** - Any code editor with Go support (VS Code, GoLand, etc.)

### Getting the Code

Sources: [go.mod3-5](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L3-L5)

Development Workflow
--------------------

The diagram below shows the typical development workflow for contributing to Skate:

Sources: [.github/workflows/build.yml1-12](https://github.com/charmbracelet/skate/blob/09860417/.github/workflows/build.yml#L1-L12)

Build Process
-------------

Skate uses a standard Go build process with GitHub Actions for CI/CD. The build workflow is triggered on both push events and pull requests.

### Local Build

To build Skate locally:

This will create a `skate` binary in your current directory.

### CI/CD Build Process

The diagram below illustrates the automated build process used for Skate:

The build process leverages reusable workflows from the `charmbracelet/meta` repository, ensuring consistency across Charm projects.

Sources: [.github/workflows/build.yml5-12](https://github.com/charmbracelet/skate/blob/09860417/.github/workflows/build.yml#L5-L12)

Release Process
---------------

Skate uses GoReleaser for managing releases. Releases are triggered by pushing a Git tag that follows semantic versioning (e.g., `v1.0.0`).

### Release Workflow

The release process creates binaries for multiple platforms and architectures, ensuring wide compatibility.

Sources: [.github/workflows/build.yml9-12](https://github.com/charmbracelet/skate/blob/09860417/.github/workflows/build.yml#L9-L12)

Dependency Management
---------------------

Skate uses Go modules for dependency management. Dependencies are defined in the `go.mod` file and locked in the `go.sum` file.

### Key Dependencies

Skate relies on several key dependencies:

| Dependency | Purpose | Version |
| --- | --- | --- |
| github.com/dgraph-io/badger/v4 | Embedded key-value database | v4.7.0 |
| github.com/spf13/cobra | Command-line interface framework | v1.9.1 |
| github.com/charmbracelet/lipgloss | Terminal styling | v1.1.0 |
| github.com/muesli/go-app-paths | Application paths management | v0.2.2 |
| github.com/agnivade/levenshtein | String similarity calculation | v1.2.1 |

### Dependency Architecture

The diagram below shows the key dependencies and their relationships in the Skate codebase:

Sources: [go.mod7-16](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L7-L16)

### Updating Dependencies

To update dependencies:

Code Quality
------------

Skate maintains code quality through automated checks and testing.

### Code Quality Tools

### Pull Request Process

1.  Create a branch for your changes
2.  Make your changes and commit them
3.  Ensure tests pass locally
4.  Submit a pull request
5.  Address any feedback from code review
6.  Once approved, your changes will be merged

Sources: [.github/workflows/build.yml1-12](https://github.com/charmbracelet/skate/blob/09860417/.github/workflows/build.yml#L1-L12)
 [go.mod1-50](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L1-L50)

Project Structure
-----------------

The diagram below illustrates the main components of the Skate codebase and their relationships:

This structure separates the command-line interface (in the `cmd/` directory) from the business logic (in the `internal/` directory), following Go best practices.

Troubleshooting Development Issues
----------------------------------

Common development issues and their solutions:

| Issue | Solution |
| --- | --- |
| Build errors | Ensure your Go version matches the one in go.mod |
| Test failures | Check that your changes maintain compatibility with existing functionality |
| Dependency conflicts | Run `go mod tidy` to resolve dependency issues |
| GitHub Actions failures | Check the workflow logs for detailed error information |

Contributing Guidelines
-----------------------

When contributing to Skate:

1.  Follow Go coding conventions
2.  Write tests for new functionality
3.  Ensure backward compatibility
4.  Update documentation as needed
5.  Keep pull requests focused on a single change
6.  Respect the existing code structure and architecture

Sources: [go.mod1-50](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L1-L50)
 [.github/workflows/build.yml1-12](https://github.com/charmbracelet/skate/blob/09860417/.github/workflows/build.yml#L1-L12)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development](https://deepwiki.com/charmbracelet/skate/5-development#development)
    
*   [Development Environment Setup](https://deepwiki.com/charmbracelet/skate/5-development#development-environment-setup)
    
*   [Getting the Code](https://deepwiki.com/charmbracelet/skate/5-development#getting-the-code)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/skate/5-development#development-workflow)
    
*   [Build Process](https://deepwiki.com/charmbracelet/skate/5-development#build-process)
    
*   [Local Build](https://deepwiki.com/charmbracelet/skate/5-development#local-build)
    
*   [CI/CD Build Process](https://deepwiki.com/charmbracelet/skate/5-development#cicd-build-process)
    
*   [Release Process](https://deepwiki.com/charmbracelet/skate/5-development#release-process)
    
*   [Release Workflow](https://deepwiki.com/charmbracelet/skate/5-development#release-workflow)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/skate/5-development#dependency-management)
    
*   [Key Dependencies](https://deepwiki.com/charmbracelet/skate/5-development#key-dependencies)
    
*   [Dependency Architecture](https://deepwiki.com/charmbracelet/skate/5-development#dependency-architecture)
    
*   [Updating Dependencies](https://deepwiki.com/charmbracelet/skate/5-development#updating-dependencies)
    
*   [Code Quality](https://deepwiki.com/charmbracelet/skate/5-development#code-quality)
    
*   [Code Quality Tools](https://deepwiki.com/charmbracelet/skate/5-development#code-quality-tools)
    
*   [Pull Request Process](https://deepwiki.com/charmbracelet/skate/5-development#pull-request-process)
    
*   [Project Structure](https://deepwiki.com/charmbracelet/skate/5-development#project-structure)
    
*   [Troubleshooting Development Issues](https://deepwiki.com/charmbracelet/skate/5-development#troubleshooting-development-issues)
    
*   [Contributing Guidelines](https://deepwiki.com/charmbracelet/skate/5-development#contributing-guidelines)
