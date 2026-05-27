Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/vhs](https://github.com/charmbracelet/vhs "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 18 April 2025 ([f28239](https://github.com/charmbracelet/vhs/commits/f28239f3)
)

*   [Overview](https://deepwiki.com/charmbracelet/vhs/1-overview)
    
*   [Installation and Setup](https://deepwiki.com/charmbracelet/vhs/1.1-installation-and-setup)
    
*   [Quick Start](https://deepwiki.com/charmbracelet/vhs/1.2-quick-start)
    
*   [Tape Format](https://deepwiki.com/charmbracelet/vhs/2-tape-format)
    
*   [Commands Reference](https://deepwiki.com/charmbracelet/vhs/2.1-commands-reference)
    
*   [Output Options](https://deepwiki.com/charmbracelet/vhs/2.2-output-options)
    
*   [Themes and Styling](https://deepwiki.com/charmbracelet/vhs/2.3-themes-and-styling)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/vhs/3-system-architecture)
    
*   [Command Execution](https://deepwiki.com/charmbracelet/vhs/3.1-command-execution)
    
*   [Terminal Emulation](https://deepwiki.com/charmbracelet/vhs/3.2-terminal-emulation)
    
*   [Recording and Rendering](https://deepwiki.com/charmbracelet/vhs/3.3-recording-and-rendering)
    
*   [CLI Features](https://deepwiki.com/charmbracelet/vhs/4-cli-features)
    
*   [Recording Mode](https://deepwiki.com/charmbracelet/vhs/4.1-recording-mode)
    
*   [SSH Server](https://deepwiki.com/charmbracelet/vhs/4.2-ssh-server)
    
*   [Publishing](https://deepwiki.com/charmbracelet/vhs/4.3-publishing)
    
*   [Examples and Use Cases](https://deepwiki.com/charmbracelet/vhs/5-examples-and-use-cases)
    
*   [Demos for CLI Tools](https://deepwiki.com/charmbracelet/vhs/5.1-demos-for-cli-tools)
    
*   [Shell Examples](https://deepwiki.com/charmbracelet/vhs/5.2-shell-examples)
    
*   [Advanced Tape Techniques](https://deepwiki.com/charmbracelet/vhs/5.3-advanced-tape-techniques)
    
*   [Development and Contribution](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution)
    
*   [Building from Source](https://deepwiki.com/charmbracelet/vhs/6.1-building-from-source)
    
*   [Project Structure](https://deepwiki.com/charmbracelet/vhs/6.2-project-structure)
    

Menu

Development and Contribution
============================

Relevant source files

*   [.github/dependabot.yml](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/dependabot.yml)
    
*   [.github/workflows/build.yml](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/build.yml)
    
*   [.github/workflows/dependabot-sync.yml](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/dependabot-sync.yml)
    
*   [.github/workflows/goreleaser.yml](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/goreleaser.yml)
    
*   [.github/workflows/lint.yml](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/lint.yml)
    
*   [.github/workflows/nightly.yml](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/nightly.yml)
    
*   [.golangci.yml](https://github.com/charmbracelet/vhs/blob/f28239f3/.golangci.yml)
    
*   [.goreleaser.yml](https://github.com/charmbracelet/vhs/blob/f28239f3/.goreleaser.yml)
    

This page provides information for developers who want to contribute to VHS. It covers the development setup, project structure, contribution workflow, and CI/CD processes. For detailed instructions on building VHS from source, see [Building from Source](https://deepwiki.com/charmbracelet/vhs/6.1-building-from-source)
. For an in-depth look at the project's code organization, see [Project Structure](https://deepwiki.com/charmbracelet/vhs/6.2-project-structure)
.

Setting Up the Development Environment
--------------------------------------

Before contributing to VHS, you'll need to set up your development environment. VHS is written in Go and uses several dependencies.

### Prerequisites

*   Go 1.18 or higher
*   Git
*   ffmpeg (for rendering GIFs/videos)
*   ttyd (for terminal emulation)
*   A code editor with Go support (VSCode, GoLand, etc.)

### Getting the Source Code

Sources:

Project Structure and Organization
----------------------------------

VHS follows a modular architecture to handle the various aspects of terminal recording and video generation.

### Key Components

*   **cmd/**: Contains the main application entry point
*   **core/**: Core functionality and types
*   **recording/**: Components for capturing frames and rendering output
*   **terminal/**: Terminal emulation and browser interaction
*   **tape/**: Tape file parsing and command execution
*   **util/**: Utility functions used across the project
*   **examples/**: Example tape files

Sources:

Development Workflow
--------------------

The recommended workflow for contributing to VHS follows these steps:

### Development Commands

### Coding Standards

VHS uses `golangci-lint` for enforcing code quality standards. The configuration is defined in `.golangci.yml`.

Key linting rules include:

*   Code formatting with gofumpt and goimports
*   Error handling checks
*   Security checks with gosec
*   Various best practices for Go code

Sources: [.golangci.yml1-50](https://github.com/charmbracelet/vhs/blob/f28239f3/.golangci.yml#L1-L50)

CI/CD and GitHub Workflows
--------------------------

VHS uses GitHub Actions for continuous integration and deployment. Understanding these workflows is important for contributors.

### Main Workflows

1.  **Build and Test** - Triggered on every push and pull request
    
    *   Builds the project
    *   Runs tests
    *   Creates snapshot releases for testing
2.  **Lint** - Checks code quality
    
    *   Uses golangci-lint with project-specific rules
3.  **GoReleaser** - Triggered on new version tags
    
    *   Creates GitHub releases
    *   Builds binaries for multiple platforms
    *   Publishes Docker images
    *   Updates package managers (Homebrew, etc.)
4.  **Nightly** - Triggered on pushes to main
    
    *   Creates nightly development builds
5.  **Dependabot** - Keeps dependencies updated
    
    *   Automatically creates PRs for dependency updates
    *   Syncs dependabot configuration weekly

Sources: [.github/workflows/build.yml1-13](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/build.yml#L1-L13)
 [.github/workflows/lint.yml1-9](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/lint.yml#L1-L9)
 [.github/workflows/goreleaser.yml1-31](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/goreleaser.yml#L1-L31)
 [.github/workflows/nightly.yml1-20](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/nightly.yml#L1-L20)
 [.github/workflows/dependabot-sync.yml1-18](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/dependabot-sync.yml#L1-L18)
 [.github/dependabot.yml1-42](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/dependabot.yml#L1-L42)

Release Process
---------------

VHS uses GoReleaser for managing releases. The process is automated through GitHub Actions.

The release configuration is defined in `.goreleaser.yml`, which includes:

*   Building binaries for multiple platforms
*   Signing for MacOS
*   Docker image creation
*   Package manager integrations (Homebrew, etc.)

Sources: [.goreleaser.yml1-12](https://github.com/charmbracelet/vhs/blob/f28239f3/.goreleaser.yml#L1-L12)
 [.github/workflows/goreleaser.yml1-31](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/workflows/goreleaser.yml#L1-L31)

Dependency Management
---------------------

VHS uses Go modules for dependency management and Dependabot to keep dependencies updated.

### Dependabot Configuration

Dependabot is configured to automatically check for updates and create pull requests:

| Ecosystem | Directory | Schedule | Labels |
| --- | --- | --- | --- |
| Go modules | /   | Weekly (Monday) | dependencies |
| GitHub Actions | /   | Weekly (Monday) | dependencies |
| Docker | /   | Weekly (Monday) | dependencies |

Sources: [.github/dependabot.yml1-42](https://github.com/charmbracelet/vhs/blob/f28239f3/.github/dependabot.yml#L1-L42)

How to Contribute
-----------------

Here are some guidelines for contributing to VHS:

1.  **Find or Create an Issue**
    
    *   Check existing issues for things to work on
    *   Create a new issue to discuss significant changes before starting work
2.  **Fork and Clone**
    
    *   Fork the repository on GitHub
    *   Clone your fork locally
3.  **Branch**
    
    *   Create a new branch for your feature or bugfix
    *   Use a descriptive name (e.g., `feature/add-new-command`, `fix/recording-issue`)
4.  **Make Changes**
    
    *   Write your code following Go best practices
    *   Add or update tests as needed
    *   Update documentation when necessary
5.  **Test and Lint**
    
    *   Run tests locally (`go test ./...`)
    *   Run linters (`golangci-lint run`)
6.  **Commit and Push**
    
    *   Write clear commit messages
    *   Push your changes to your fork
7.  **Create a Pull Request**
    
    *   Submit a PR from your branch to the main VHS repository
    *   Include a description of your changes
    *   Link to any related issues
8.  **Address Feedback**
    
    *   Be responsive to review comments
    *   Make requested changes

Sources:

Common Development Tasks
------------------------

### Adding a New Command

To add a new command to the tape format:

1.  Define the command type in the tape package
2.  Implement the command execution function
3.  Register the command in the parser
4.  Add tests for the new command
5.  Update documentation

### Debugging Tips

*   Use the `--debug` flag for verbose logging
*   Examine the individual frames in the output directory
*   Test with simple tape files to isolate issues

### Running Integration Tests

VHS can be tested end-to-end by:

1.  Creating example tape files
2.  Running VHS on those files
3.  Verifying the output GIFs or videos

Sources:

Conclusion
----------

Contributing to VHS is an opportunity to work on a unique tool that bridges terminal interactions and video creation. The project's modular architecture and comprehensive testing and CI/CD infrastructure make it accessible for new contributors.

For more detailed information about the project structure and specific components, see [Project Structure](https://deepwiki.com/charmbracelet/vhs/6.2-project-structure)
. For step-by-step instructions on building from source, see [Building from Source](https://deepwiki.com/charmbracelet/vhs/6.1-building-from-source)
.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development and Contribution](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#development-and-contribution)
    
*   [Setting Up the Development Environment](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#setting-up-the-development-environment)
    
*   [Prerequisites](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#prerequisites)
    
*   [Getting the Source Code](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#getting-the-source-code)
    
*   [Project Structure and Organization](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#project-structure-and-organization)
    
*   [Key Components](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#key-components)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#development-workflow)
    
*   [Development Commands](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#development-commands)
    
*   [Coding Standards](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#coding-standards)
    
*   [CI/CD and GitHub Workflows](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#cicd-and-github-workflows)
    
*   [Main Workflows](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#main-workflows)
    
*   [Release Process](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#release-process)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#dependency-management)
    
*   [Dependabot Configuration](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#dependabot-configuration)
    
*   [How to Contribute](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#how-to-contribute)
    
*   [Common Development Tasks](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#common-development-tasks)
    
*   [Adding a New Command](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#adding-a-new-command)
    
*   [Debugging Tips](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#debugging-tips)
    
*   [Running Integration Tests](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#running-integration-tests)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/vhs/6-development-and-contribution#conclusion)
