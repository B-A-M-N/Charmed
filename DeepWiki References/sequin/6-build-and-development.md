Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/sequin](https://github.com/charmbracelet/sequin "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 20 April 2025 ([025641](https://github.com/charmbracelet/sequin/commits/025641c1)
)

*   [Introduction to Sequin](https://deepwiki.com/charmbracelet/sequin/1-introduction-to-sequin)
    
*   [Installation](https://deepwiki.com/charmbracelet/sequin/1.1-installation)
    
*   [Basic Usage](https://deepwiki.com/charmbracelet/sequin/1.2-basic-usage)
    
*   [Architecture](https://deepwiki.com/charmbracelet/sequin/2-architecture)
    
*   [ANSI Sequence Processing](https://deepwiki.com/charmbracelet/sequin/2.1-ansi-sequence-processing)
    
*   [Command Execution](https://deepwiki.com/charmbracelet/sequin/2.2-command-execution)
    
*   [ANSI Sequence Handlers](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers)
    
*   [SGR (Select Graphic Rendition)](https://deepwiki.com/charmbracelet/sequin/3.1-sgr-(select-graphic-rendition))
    
*   [Cursor Commands](https://deepwiki.com/charmbracelet/sequin/3.2-cursor-commands)
    
*   [Screen and Line Commands](https://deepwiki.com/charmbracelet/sequin/3.3-screen-and-line-commands)
    
*   [Window Title and Colors](https://deepwiki.com/charmbracelet/sequin/3.4-window-title-and-colors)
    
*   [Terminal Modes](https://deepwiki.com/charmbracelet/sequin/3.5-terminal-modes)
    
*   [Pointer, Clipboard, and Links](https://deepwiki.com/charmbracelet/sequin/3.6-pointer-clipboard-and-links)
    
*   [FinalTerm Commands](https://deepwiki.com/charmbracelet/sequin/3.7-finalterm-commands)
    
*   [Theme System](https://deepwiki.com/charmbracelet/sequin/4-theme-system)
    
*   [Testing](https://deepwiki.com/charmbracelet/sequin/5-testing)
    
*   [Build and Development](https://deepwiki.com/charmbracelet/sequin/6-build-and-development)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/sequin/6.1-dependency-management)
    
*   [Release Process](https://deepwiki.com/charmbracelet/sequin/6.2-release-process)
    

Menu

Build and Development
=====================

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/build.yml)
    
*   [.github/workflows/lint.yml](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/lint.yml)
    
*   [.golangci.yml](https://github.com/charmbracelet/sequin/blob/025641c1/.golangci.yml)
    
*   [go.mod](https://github.com/charmbracelet/sequin/blob/025641c1/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/sequin/blob/025641c1/go.sum)
    

This page provides information for developers who want to build, lint, and contribute to Sequin. It covers development environment setup, build processes, and code quality tools. For information about the release process, see [Release Process](https://deepwiki.com/charmbracelet/sequin/6.2-release-process)
, and for details on dependency management, see [Dependency Management](https://deepwiki.com/charmbracelet/sequin/6.1-dependency-management)
.

Development Environment Setup
-----------------------------

### Go Version Requirements

Sequin requires Go 1.23.0 or later for development, with a preferred toolchain of Go 1.24.1. This requirement is specified in the project's `go.mod` file.

    module github.com/charmbracelet/sequin
    
    go 1.23.0
    
    toolchain go1.24.1
    

Sources: [go.mod1-5](https://github.com/charmbracelet/sequin/blob/025641c1/go.mod#L1-L5)

### Dependencies

Sequin relies on several key dependencies from the Charm ecosystem and other Go packages:

| Dependency | Purpose |
| --- | --- |
| `github.com/charmbracelet/colorprofile` | Terminal color profile detection |
| `github.com/charmbracelet/lipgloss/v2` | Terminal styling library |
| `github.com/charmbracelet/x/ansi` | ANSI sequence parsing |
| `github.com/charmbracelet/x/exp/golden` | Golden file testing |
| `github.com/charmbracelet/x/term` | Terminal interaction utilities |
| `github.com/charmbracelet/x/xpty` | Pseudo-terminal handling |
| `github.com/spf13/cobra` | Command-line interface framework |
| `github.com/stretchr/testify` | Testing assertions |

The full list of dependencies and their versions can be found in the `go.mod` file.

Sources: [go.mod7-16](https://github.com/charmbracelet/sequin/blob/025641c1/go.mod#L7-L16)

Build Process
-------------

### Local Development Workflow

Sources: [.github/workflows/build.yml1-17](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/build.yml#L1-L17)
 [.github/workflows/lint.yml1-15](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/lint.yml#L1-L15)

### Building Sequin Locally

To build Sequin locally, standard Go build commands can be used:

Sources: [go.mod1-36](https://github.com/charmbracelet/sequin/blob/025641c1/go.mod#L1-L36)

Code Quality Tools
------------------

### Linting Configuration

Sequin uses golangci-lint for code quality checks with an extensive set of linters enabled. The configuration is defined in `.golangci.yml`.

The linting configuration includes features like:

*   Code formatting with gofumpt and goimports
*   Error handling checks
*   Security scanning
*   Performance optimizations
*   Documentation style enforcement

Sources: [.golangci.yml1-50](https://github.com/charmbracelet/sequin/blob/025641c1/.golangci.yml#L1-L50)

### Enabled Linters

The project enables a comprehensive set of linters to ensure code quality:

| Category | Linters |
| --- | --- |
| Error handling | bodyclose, nilerr, rowserrcheck, sqlclosecheck |
| Code style | godot, whitespace, revive, misspell |
| Security | gosec, noctx |
| Performance | prealloc |
| Logic | exhaustive, nestif, nakedret, unconvert, unparam |
| Documentation | godox, goprintffuncname |

Sources: [.golangci.yml4-28](https://github.com/charmbracelet/sequin/blob/025641c1/.golangci.yml#L4-L28)

CI/CD Pipeline
--------------

Sequin uses GitHub Actions for continuous integration and delivery processes.

The CI pipeline includes:

1.  **Linting workflow**: Runs golangci-lint with the project's custom configuration
2.  **Build workflow**: Builds the project and runs tests

The workflows are configured in the `.github/workflows` directory.

Sources: [.github/workflows/build.yml1-17](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/build.yml#L1-L17)
 [.github/workflows/lint.yml1-15](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/lint.yml#L1-L15)

Contributing Guidelines
-----------------------

When contributing to Sequin, developers should follow these steps:

1.  Set up the development environment with the required Go version
2.  Fork the repository and create a feature branch
3.  Make code changes following the project's coding standards
4.  Ensure all linters pass locally using golangci-lint
5.  Add tests for new functionality
6.  Submit a pull request with a clear description of changes
7.  Wait for CI checks to pass and address any feedback

All pull requests trigger automatic builds and lint checks via GitHub Actions.

Sources: [.github/workflows/build.yml3-9](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/build.yml#L3-L9)
 [.github/workflows/lint.yml2-14](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/lint.yml#L2-L14)

Building for Different Platforms
--------------------------------

Being a Go application, Sequin can be cross-compiled for different platforms:

Sources: [go.mod1-36](https://github.com/charmbracelet/sequin/blob/025641c1/go.mod#L1-L36)

Development Tools Ecosystem
---------------------------

Sequin's development relies on an ecosystem of tools that work together to ensure code quality and project maintainability.

Sources: [go.mod1-36](https://github.com/charmbracelet/sequin/blob/025641c1/go.mod#L1-L36)
 [.github/workflows/build.yml1-17](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/build.yml#L1-L17)
 [.github/workflows/lint.yml1-15](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/lint.yml#L1-L15)

Summary
-------

Sequin follows modern Go development practices with a strong emphasis on code quality through comprehensive linting, testing, and CI/CD automation. Developers can easily contribute to the project by following the standard Go development workflow and ensuring their changes pass all the automated quality checks.

The project uses GitHub Actions for continuous integration, running linters and tests automatically for each pull request and push to the main branch, which helps maintain code quality throughout the development process.

Sources: [go.mod1-36](https://github.com/charmbracelet/sequin/blob/025641c1/go.mod#L1-L36)
 [.golangci.yml1-50](https://github.com/charmbracelet/sequin/blob/025641c1/.golangci.yml#L1-L50)
 [.github/workflows/build.yml1-17](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/build.yml#L1-L17)
 [.github/workflows/lint.yml1-15](https://github.com/charmbracelet/sequin/blob/025641c1/.github/workflows/lint.yml#L1-L15)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Build and Development](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#build-and-development)
    
*   [Development Environment Setup](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#development-environment-setup)
    
*   [Go Version Requirements](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#go-version-requirements)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#dependencies)
    
*   [Build Process](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#build-process)
    
*   [Local Development Workflow](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#local-development-workflow)
    
*   [Building Sequin Locally](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#building-sequin-locally)
    
*   [Code Quality Tools](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#code-quality-tools)
    
*   [Linting Configuration](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#linting-configuration)
    
*   [Enabled Linters](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#enabled-linters)
    
*   [CI/CD Pipeline](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#cicd-pipeline)
    
*   [Contributing Guidelines](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#contributing-guidelines)
    
*   [Building for Different Platforms](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#building-for-different-platforms)
    
*   [Development Tools Ecosystem](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#development-tools-ecosystem)
    
*   [Summary](https://deepwiki.com/charmbracelet/sequin/6-build-and-development#summary)
