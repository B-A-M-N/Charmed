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

Development Guide
=================

Relevant source files

*   [.github/dependabot.yml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/dependabot.yml)
    
*   [.github/workflows/build.yml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/build.yml)
    
*   [.github/workflows/dependabot-sync.yml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/dependabot-sync.yml)
    
*   [.github/workflows/lint-soft.yml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/lint-soft.yml)
    
*   [.github/workflows/lint.yml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/lint.yml)
    

This guide provides instructions for developers who want to contribute to the Wishlist project. It covers setting up your development environment, understanding the codebase structure, development workflow, and contribution guidelines. For information about using Wishlist, see [Introduction to Wishlist](https://deepwiki.com/charmbracelet/wishlist/1-introduction-to-wishlist)
 and [Getting Started](https://deepwiki.com/charmbracelet/wishlist/1.1-getting-started)
.

Setting Up Development Environment
----------------------------------

### Prerequisites

*   Go (latest stable version)
*   Git
*   golangci-lint (for local code linting)

### Initial Setup

1.  Fork the repository
2.  Clone your fork locally:
    
3.  Add the upstream remote:
    
4.  Install dependencies:
    
5.  Build the project:
    

Sources: [.github/workflows/build.yml23-27](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/build.yml#L23-L27)

Codebase Structure
------------------

The Wishlist codebase is organized around several key components. Understanding this structure will help you locate where changes should be made.

Sources: [cmd/wishlist/main.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/cmd/wishlist/main.go)
 [client\_auth.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/client_auth.go)
 [server.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/server.go)
 [middleware.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/middleware.go)
 [config.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/config.go)
 [wishlist.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go)

Development Workflow
--------------------

### Code Contribution Process

Sources: [.github/workflows/build.yml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/build.yml)
 [.github/workflows/lint.yml](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/lint.yml)

### Testing

Always run tests locally before submitting a pull request:

This is the same command that runs in CI, so running it locally will help catch issues early.

Sources: [.github/workflows/build.yml29-30](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/build.yml#L29-L30)

### Linting

The project uses golangci-lint for code quality checks. There are two linting configurations:

1.  **Strict linting** - Must pass for CI to succeed:
    
2.  **Soft linting** - Provides additional suggestions but doesn't fail CI:
    

Sources: [.github/workflows/lint.yml22-28](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/lint.yml#L22-L28)
 [.github/workflows/lint-soft.yml22-28](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/lint-soft.yml#L22-L28)

Continuous Integration
----------------------

The project uses GitHub Actions for continuous integration. The following workflows run automatically:

Sources: [.github/workflows/build.yml5-54](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/build.yml#L5-L54)
 [.github/workflows/lint.yml1-28](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/lint.yml#L1-L28)
 [.github/workflows/lint-soft.yml1-28](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/lint-soft.yml#L1-L28)

Contributing Guidelines
-----------------------

### Commit Messages

Based on the Dependabot configuration, the project follows these commit message conventions:

*   `chore`: For dependency updates and maintenance tasks
*   `feat`: For new features
*   Include scope where appropriate (e.g., `chore(deps): update module`)

Sources: [.github/dependabot.yml11-41](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/dependabot.yml#L11-L41)

### Pull Request Workflow

1.  Create a branch for your changes
2.  Make your changes
3.  Ensure tests pass locally
4.  Ensure linting passes
5.  Push changes to your fork
6.  Create a pull request
7.  Respond to review feedback

Dependency Management
---------------------

The project uses Dependabot for automated dependency updates:

*   Go modules are checked weekly (Mondays)
*   GitHub Actions dependencies are checked weekly
*   Docker dependencies are checked weekly

Dependabot PRs are automatically approved and merged if CI checks pass.

Sources: [.github/dependabot.yml3-41](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/dependabot.yml#L3-L41)
 [.github/workflows/dependabot-sync.yml1-18](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/dependabot-sync.yml#L1-L18)
 [.github/workflows/build.yml37-54](https://github.com/charmbracelet/wishlist/blob/5fe34a29/.github/workflows/build.yml#L37-L54)

Component Development Guide
---------------------------

When working on specific components of Wishlist, understand how they interact with the rest of the system:

### Key Components

1.  **SSH Client**: Manages connections to endpoints, authentication methods, and session handling
2.  **SSH Server**: Provides the SSH server functionality when running in server mode
3.  **Configuration**: Handles loading and parsing configuration from various sources
4.  **Discovery**: Implements endpoint discovery methods (Tailscale, Zeroconf, SRV)
5.  **Terminal UI**: Provides the interactive terminal interface for endpoint selection

When making changes, consider how they affect other components and the overall architecture described in [Architecture Overview](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview)
.

Testing Strategy
----------------

When adding new features or fixing bugs, follow these testing principles:

1.  Write unit tests for individual functions
2.  Add integration tests for component interactions
3.  Test edge cases and error conditions
4.  Maintain or improve overall test coverage

Documentation
-------------

Update documentation when making changes:

1.  Update code comments in godoc format
2.  Update README.md for user-facing changes
3.  Consider updating wiki pages for architectural changes

Sources: All files

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development Guide](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#development-guide)
    
*   [Setting Up Development Environment](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#setting-up-development-environment)
    
*   [Prerequisites](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#prerequisites)
    
*   [Initial Setup](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#initial-setup)
    
*   [Codebase Structure](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#codebase-structure)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#development-workflow)
    
*   [Code Contribution Process](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#code-contribution-process)
    
*   [Testing](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#testing)
    
*   [Linting](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#linting)
    
*   [Continuous Integration](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#continuous-integration)
    
*   [Contributing Guidelines](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#contributing-guidelines)
    
*   [Commit Messages](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#commit-messages)
    
*   [Pull Request Workflow](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#pull-request-workflow)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#dependency-management)
    
*   [Component Development Guide](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#component-development-guide)
    
*   [Key Components](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#key-components)
    
*   [Testing Strategy](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#testing-strategy)
    
*   [Documentation](https://deepwiki.com/charmbracelet/wishlist/10-development-guide#documentation)
