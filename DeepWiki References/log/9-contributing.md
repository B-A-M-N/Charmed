Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/log](https://github.com/charmbracelet/log "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 20 April 2025 ([cf6e86](https://github.com/charmbracelet/log/commits/cf6e8671)
)

*   [Overview](https://deepwiki.com/charmbracelet/log/1-overview)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/log/2-core-architecture)
    
*   [Logger Structure](https://deepwiki.com/charmbracelet/log/2.1-logger-structure)
    
*   [Log Levels](https://deepwiki.com/charmbracelet/log/2.2-log-levels)
    
*   [Global Logger](https://deepwiki.com/charmbracelet/log/2.3-global-logger)
    
*   [Usage Guide](https://deepwiki.com/charmbracelet/log/3-usage-guide)
    
*   [Basic Logging](https://deepwiki.com/charmbracelet/log/3.1-basic-logging)
    
*   [Contextual Logging](https://deepwiki.com/charmbracelet/log/3.2-contextual-logging)
    
*   [Helper Functions](https://deepwiki.com/charmbracelet/log/3.3-helper-functions)
    
*   [Formatters](https://deepwiki.com/charmbracelet/log/4-formatters)
    
*   [Text Formatter](https://deepwiki.com/charmbracelet/log/4.1-text-formatter)
    
*   [JSON Formatter](https://deepwiki.com/charmbracelet/log/4.2-json-formatter)
    
*   [Logfmt Formatter](https://deepwiki.com/charmbracelet/log/4.3-logfmt-formatter)
    
*   [Styling and Customization](https://deepwiki.com/charmbracelet/log/5-styling-and-customization)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/log/6-configuration-options)
    
*   [Integration](https://deepwiki.com/charmbracelet/log/7-integration)
    
*   [Standard Library Integration](https://deepwiki.com/charmbracelet/log/7.1-standard-library-integration)
    
*   [Slog Integration](https://deepwiki.com/charmbracelet/log/7.2-slog-integration)
    
*   [Advanced Topics](https://deepwiki.com/charmbracelet/log/8-advanced-topics)
    
*   [Contributing](https://deepwiki.com/charmbracelet/log/9-contributing)
    

Menu

Contributing
============

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/build.yml)
    
*   [.github/workflows/lint-sync.yml](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/lint-sync.yml)
    
*   [.github/workflows/lint.yml](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/lint.yml)
    
*   [.golangci.yml](https://github.com/charmbracelet/log/blob/cf6e8671/.golangci.yml)
    

This document provides guidance for contributors to the `charmbracelet/log` project. It outlines the development workflow, coding standards, testing requirements, and process for submitting changes. For information about how to use the library, see [Usage Guide](https://deepwiki.com/charmbracelet/log/3-usage-guide)
.

Development Setup
-----------------

### Prerequisites

To contribute to the `charmbracelet/log` project, you'll need:

1.  Go 1.19 or newer
2.  Git
3.  A GitHub account
4.  Familiarity with Go modules and testing

### Getting Started

1.  **Fork the repository** on GitHub
2.  **Clone your fork**:
    
        git clone https://github.com/your-username/log.git
        cd log
        
    
3.  **Set up the upstream remote**:
    
        git remote add upstream https://github.com/charmbracelet/log.git
        
    
4.  **Create a feature branch**:
    
        git checkout -b feature/your-feature-name
        
    

Sources: [.github/workflows/build.yml1-21](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/build.yml#L1-L21)

Code Quality Standards
----------------------

The `charmbracelet/log` project maintains high code quality standards through automated linting and formatting tools. All contributions must pass these checks before they can be merged.

### Linting

The project uses golangci-lint with specific linters enabled:

To run linting locally:

1.  Install golangci-lint (if not already installed)
2.  Run: `golangci-lint run`

The configuration in `.golangci.yml` outlines all enabled linters and formatting rules.

Sources: [.golangci.yml1-50](https://github.com/charmbracelet/log/blob/cf6e8671/.golangci.yml#L1-L50)
 [.github/workflows/lint.yml1-9](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/lint.yml#L1-L9)

### Formatting

All code must be formatted according to the project's standards. The project uses:

*   gofumpt (stricter gofmt)
*   goimports (organizes imports)

These formatters run automatically as part of the linting process.

Sources: [.golangci.yml40-49](https://github.com/charmbracelet/log/blob/cf6e8671/.golangci.yml#L40-L49)

Testing
-------

All contributions should include appropriate tests to verify functionality and prevent regressions.

### Running Tests

To run tests locally:

    go test -v ./...
    

For coverage reporting:

    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out
    

### Test Requirements

New features should include:

1.  Unit tests that verify the functionality works as expected
2.  Integration tests where appropriate
3.  Test updates for any modified existing functionality

Sources: [.github/workflows/build.yml19-21](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/build.yml#L19-L21)

Continuous Integration
----------------------

The project uses GitHub Actions for CI/CD. When you submit a pull request, these workflows automatically run:

The CI process checks:

1.  Code passes all linters
2.  Code builds successfully
3.  Tests pass
4.  Test coverage is maintained

Sources: [.github/workflows/build.yml5-21](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/build.yml#L5-L21)
 [.github/workflows/lint.yml1-9](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/lint.yml#L1-L9)

Pull Request Process
--------------------

1.  **Create a pull request** from your feature branch to the main repository's `main` branch
2.  **Ensure CI passes** - all automated checks must pass
3.  **Await code review** - maintainers will review your changes
4.  **Address feedback** - make any requested changes
5.  **Approval and merge** - once approved, maintainers will merge your PR

### PR Guidelines

1.  Keep changes focused on a single issue/feature
2.  Include a clear description of the changes
3.  Reference any related issues
4.  Update documentation as needed
5.  Follow the existing code style

Contribution Workflow Architecture
----------------------------------

The following diagram illustrates how the contribution process interacts with the codebase:

Sources: [.github/workflows/build.yml1-21](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/build.yml#L1-L21)
 [.github/workflows/lint.yml1-9](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/lint.yml#L1-L9)
 [.github/workflows/lint-sync.yml1-15](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/lint-sync.yml#L1-L15)

Releasing
---------

The project uses automated workflows for creating snapshots and releases. Contributors typically don't need to worry about this process as it's handled by maintainers.

The build workflow includes a snapshot job that uses GoReleaser for generating snapshot releases during the development process.

Sources: [.github/workflows/build.yml14-17](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/build.yml#L14-L17)

Linting Standards Maintenance
-----------------------------

The project automatically syncs linting configurations weekly to ensure consistency across the Charmbracelet ecosystem:

This ensures that the project maintains consistent code quality standards with other Charmbracelet projects.

Sources: [.github/workflows/lint-sync.yml1-15](https://github.com/charmbracelet/log/blob/cf6e8671/.github/workflows/lint-sync.yml#L1-L15)

Summary
-------

Contributing to `charmbracelet/log` involves:

1.  Setting up a development environment
2.  Following the project's code quality standards
3.  Writing and running tests
4.  Submitting well-documented pull requests
5.  Addressing review feedback

By following these guidelines, you'll help maintain the quality and reliability of the `charmbracelet/log` library while expanding its capabilities.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Contributing](https://deepwiki.com/charmbracelet/log/9-contributing#contributing)
    
*   [Development Setup](https://deepwiki.com/charmbracelet/log/9-contributing#development-setup)
    
*   [Prerequisites](https://deepwiki.com/charmbracelet/log/9-contributing#prerequisites)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/log/9-contributing#getting-started)
    
*   [Code Quality Standards](https://deepwiki.com/charmbracelet/log/9-contributing#code-quality-standards)
    
*   [Linting](https://deepwiki.com/charmbracelet/log/9-contributing#linting)
    
*   [Formatting](https://deepwiki.com/charmbracelet/log/9-contributing#formatting)
    
*   [Testing](https://deepwiki.com/charmbracelet/log/9-contributing#testing)
    
*   [Running Tests](https://deepwiki.com/charmbracelet/log/9-contributing#running-tests)
    
*   [Test Requirements](https://deepwiki.com/charmbracelet/log/9-contributing#test-requirements)
    
*   [Continuous Integration](https://deepwiki.com/charmbracelet/log/9-contributing#continuous-integration)
    
*   [Pull Request Process](https://deepwiki.com/charmbracelet/log/9-contributing#pull-request-process)
    
*   [PR Guidelines](https://deepwiki.com/charmbracelet/log/9-contributing#pr-guidelines)
    
*   [Contribution Workflow Architecture](https://deepwiki.com/charmbracelet/log/9-contributing#contribution-workflow-architecture)
    
*   [Releasing](https://deepwiki.com/charmbracelet/log/9-contributing#releasing)
    
*   [Linting Standards Maintenance](https://deepwiki.com/charmbracelet/log/9-contributing#linting-standards-maintenance)
    
*   [Summary](https://deepwiki.com/charmbracelet/log/9-contributing#summary)
