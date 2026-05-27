Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/wish](https://github.com/charmbracelet/wish "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 20 April 2025 ([19f432](https://github.com/charmbracelet/wish/commits/19f43236)
)

*   [Introduction to Wish](https://deepwiki.com/charmbracelet/wish/1-introduction-to-wish)
    
*   [Core SSH Server Framework](https://deepwiki.com/charmbracelet/wish/2-core-ssh-server-framework)
    
*   [Server Configuration and Options](https://deepwiki.com/charmbracelet/wish/2.1-server-configuration-and-options)
    
*   [Command Execution](https://deepwiki.com/charmbracelet/wish/2.2-command-execution)
    
*   [Middleware System](https://deepwiki.com/charmbracelet/wish/3-middleware-system)
    
*   [Access Control Middleware](https://deepwiki.com/charmbracelet/wish/3.1-access-control-middleware)
    
*   [Logging and Metrics Middleware](https://deepwiki.com/charmbracelet/wish/3.2-logging-and-metrics-middleware)
    
*   [Rate Limiting Middleware](https://deepwiki.com/charmbracelet/wish/3.3-rate-limiting-middleware)
    
*   [Extension Modules](https://deepwiki.com/charmbracelet/wish/4-extension-modules)
    
*   [Git Server Integration](https://deepwiki.com/charmbracelet/wish/4.1-git-server-integration)
    
*   [BubbleTea Integration](https://deepwiki.com/charmbracelet/wish/4.2-bubbletea-integration)
    
*   [File Transfer (SCP/SFTP)](https://deepwiki.com/charmbracelet/wish/4.3-file-transfer-(scpsftp))
    
*   [Authentication](https://deepwiki.com/charmbracelet/wish/5-authentication)
    
*   [Public Key Authentication](https://deepwiki.com/charmbracelet/wish/5.1-public-key-authentication)
    
*   [Other Authentication Methods](https://deepwiki.com/charmbracelet/wish/5.2-other-authentication-methods)
    
*   [Usage Examples](https://deepwiki.com/charmbracelet/wish/6-usage-examples)
    
*   [Simple SSH Server](https://deepwiki.com/charmbracelet/wish/6.1-simple-ssh-server)
    
*   [BubbleTea Examples](https://deepwiki.com/charmbracelet/wish/6.2-bubbletea-examples)
    
*   [Git Server Examples](https://deepwiki.com/charmbracelet/wish/6.3-git-server-examples)
    
*   [Development and Contribution](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution)
    
*   [Project Structure](https://deepwiki.com/charmbracelet/wish/7.1-project-structure)
    
*   [CI/CD and Releases](https://deepwiki.com/charmbracelet/wish/7.2-cicd-and-releases)
    

Menu

Development and Contribution
============================

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/build.yml)
    
*   [.github/workflows/examples.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/examples.yml)
    
*   [.github/workflows/lint.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/lint.yml)
    
*   [.gitignore](https://github.com/charmbracelet/wish/blob/19f43236/.gitignore)
    
*   [go.mod](https://github.com/charmbracelet/wish/blob/19f43236/go.mod)
    

This document provides essential information for developers who want to contribute to the Wish project. It covers the development environment setup, contribution workflow, testing practices, and the CI/CD pipeline. For specific details about the codebase organization, refer to [Project Structure](https://deepwiki.com/charmbracelet/wish/7.1-project-structure)
, and for information about release processes, see [CI/CD and Releases](https://deepwiki.com/charmbracelet/wish/7.2-cicd-and-releases)
.

Setting Up Your Development Environment
---------------------------------------

### Prerequisites

To contribute to Wish, you'll need:

*   Go (version 1.23+)
*   Git
*   SSH client for testing

### Local Setup

1.  Fork the repository to your GitHub account
2.  Clone your fork locally
3.  Add the original repository as an upstream remote
4.  Create a new branch for your changes

Sources: [go.mod](https://github.com/charmbracelet/wish/blob/19f43236/go.mod)
 [.gitignore](https://github.com/charmbracelet/wish/blob/19f43236/.gitignore)

Contribution Workflow
---------------------

The Wish project follows a standard GitHub workflow with pull requests and code reviews.

Sources: [.github/workflows/build.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/build.yml)
 [.github/workflows/lint.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/lint.yml)

### Pull Request Guidelines

1.  Ensure all tests pass locally before submitting
2.  Keep pull requests focused on a single concern
3.  Follow the Go style guidelines
4.  Write clear commit messages
5.  Include tests for new functionality
6.  Update documentation as needed

Testing and Linting
-------------------

Wish uses Go's standard testing framework and follows strict code quality standards.

### Running Tests Locally

### Code Linting

The project uses golangci-lint for code quality checks:

Sources: [.github/workflows/build.yml11-22](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/build.yml#L11-L22)
 [.github/workflows/lint.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/lint.yml)

CI/CD Pipeline
--------------

Wish uses GitHub Actions for continuous integration and deployment.

Sources: [.github/workflows/build.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/build.yml)
 [.github/workflows/lint.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/lint.yml)
 [.github/workflows/examples.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/examples.yml)

### Automated Workflows

*   **Build**: Compiles the code, runs tests, and uploads coverage reports to Codecov
*   **Lint**: Checks code quality using golangci-lint
*   **Examples**: Keeps the examples directory's dependencies updated

Project Dependencies
--------------------

Wish relies on several key dependencies, primarily other Charmbracelet projects and Go libraries.

Sources: [go.mod7-25](https://github.com/charmbracelet/wish/blob/19f43236/go.mod#L7-L25)

Development Best Practices
--------------------------

### Code Structure

The Wish project follows idiomatic Go practices for code organization:

| Directory/File Pattern | Purpose |
| --- | --- |
| `*.go` in root | Core server implementation |
| `middleware/` | Middleware implementations |
| `bubbletea/` | BubbleTea TUI integration |
| `git/` | Git server implementation |
| `scp/` | File transfer capabilities |
| `examples/` | Example applications |
| `*_test.go` | Test files |

### Documentation Guidelines

*   All exported functions, types, and variables should have proper documentation comments
*   Examples should be provided where appropriate
*   Comments should follow Go conventions (starting with the name of the element being described)

### Testing Strategy

Wish uses a combination of unit tests and integration tests:

1.  **Unit Tests**: Test individual functions and components in isolation
2.  **Integration Tests**: Test the interaction between components
3.  **Example Tests**: Demonstrate usage through executable examples

Contributing to Examples
------------------------

The examples directory contains sample applications demonstrating Wish capabilities. When contributing to examples:

1.  Ensure they're well-documented with comments
2.  Keep dependencies minimal and up-to-date
3.  Make sure they work correctly with the latest Wish version

An automated workflow keeps the examples' dependencies in sync with the main project.

Sources: [.github/workflows/examples.yml](https://github.com/charmbracelet/wish/blob/19f43236/.github/workflows/examples.yml)
 [.gitignore1-6](https://github.com/charmbracelet/wish/blob/19f43236/.gitignore#L1-L6)

Community and Support
---------------------

*   **Issues**: Use GitHub issues for bug reports and feature requests
*   **Discussions**: GitHub Discussions for general questions and community interaction
*   **Pull Requests**: Submit contributions via pull requests

When reporting issues, include:

*   Go version
*   Wish version
*   Operating system
*   Minimal reproducible example
*   Expected vs. actual behavior

Future Development
------------------

When contributing new features, consider:

1.  Alignment with Wish's goals as a SSH application framework
2.  Backward compatibility
3.  Performance implications
4.  Security considerations

Sources: [go.mod](https://github.com/charmbracelet/wish/blob/19f43236/go.mod)

* * *

By following these guidelines, you'll be able to effectively contribute to the Wish project. Remember that all contributions are subject to code review to maintain quality and consistency throughout the codebase.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development and Contribution](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#development-and-contribution)
    
*   [Setting Up Your Development Environment](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#setting-up-your-development-environment)
    
*   [Prerequisites](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#prerequisites)
    
*   [Local Setup](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#local-setup)
    
*   [Contribution Workflow](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#contribution-workflow)
    
*   [Pull Request Guidelines](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#pull-request-guidelines)
    
*   [Testing and Linting](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#testing-and-linting)
    
*   [Running Tests Locally](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#running-tests-locally)
    
*   [Code Linting](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#code-linting)
    
*   [CI/CD Pipeline](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#cicd-pipeline)
    
*   [Automated Workflows](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#automated-workflows)
    
*   [Project Dependencies](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#project-dependencies)
    
*   [Development Best Practices](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#development-best-practices)
    
*   [Code Structure](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#code-structure)
    
*   [Documentation Guidelines](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#documentation-guidelines)
    
*   [Testing Strategy](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#testing-strategy)
    
*   [Contributing to Examples](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#contributing-to-examples)
    
*   [Community and Support](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#community-and-support)
    
*   [Future Development](https://deepwiki.com/charmbracelet/wish/7-development-and-contribution#future-development)
