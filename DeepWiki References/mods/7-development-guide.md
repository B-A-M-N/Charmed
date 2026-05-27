Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/mods](https://github.com/charmbracelet/mods "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 7 November 2025 ([07a05d](https://github.com/charmbracelet/mods/commits/07a05d5b)
)

*   [Introduction to Mods](https://deepwiki.com/charmbracelet/mods/1-introduction-to-mods)
    
*   [Installation and Setup](https://deepwiki.com/charmbracelet/mods/1.1-installation-and-setup)
    
*   [Quick Start Guide](https://deepwiki.com/charmbracelet/mods/1.2-quick-start-guide)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/mods/2-core-architecture)
    
*   [Application Entry Point](https://deepwiki.com/charmbracelet/mods/2.1-application-entry-point)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/mods/2.2-configuration-system)
    
*   [Bubble Tea UI Model](https://deepwiki.com/charmbracelet/mods/2.3-bubble-tea-ui-model)
    
*   [Message Stream Context](https://deepwiki.com/charmbracelet/mods/2.4-message-stream-context)
    
*   [Conversation Management](https://deepwiki.com/charmbracelet/mods/2.5-conversation-management)
    
*   [LLM Provider Integration](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration)
    
*   [Provider Configuration](https://deepwiki.com/charmbracelet/mods/3.1-provider-configuration)
    
*   [Client Initialization and Streaming](https://deepwiki.com/charmbracelet/mods/3.2-client-initialization-and-streaming)
    
*   [Model Resolution and Fallbacks](https://deepwiki.com/charmbracelet/mods/3.3-model-resolution-and-fallbacks)
    
*   [User Interface and Experience](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience)
    
*   [Terminal Capabilities and Styling](https://deepwiki.com/charmbracelet/mods/4.1-terminal-capabilities-and-styling)
    
*   [Animation System](https://deepwiki.com/charmbracelet/mods/4.2-animation-system)
    
*   [Output Rendering and Formatting](https://deepwiki.com/charmbracelet/mods/4.3-output-rendering-and-formatting)
    
*   [Data Persistence](https://deepwiki.com/charmbracelet/mods/5-data-persistence)
    
*   [Conversation Database](https://deepwiki.com/charmbracelet/mods/5.1-conversation-database)
    
*   [Cache System](https://deepwiki.com/charmbracelet/mods/5.2-cache-system)
    
*   [Supporting Systems](https://deepwiki.com/charmbracelet/mods/6-supporting-systems)
    
*   [Flag Parsing and Error Handling](https://deepwiki.com/charmbracelet/mods/6.1-flag-parsing-and-error-handling)
    
*   [Content Loading](https://deepwiki.com/charmbracelet/mods/6.2-content-loading)
    
*   [Message Utilities](https://deepwiki.com/charmbracelet/mods/6.3-message-utilities)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/mods/7-development-guide)
    
*   [Dependencies and Build System](https://deepwiki.com/charmbracelet/mods/7.1-dependencies-and-build-system)
    
*   [Testing Strategy](https://deepwiki.com/charmbracelet/mods/7.2-testing-strategy)
    
*   [Code Quality and CI/CD](https://deepwiki.com/charmbracelet/mods/7.3-code-quality-and-cicd)
    

Menu

Development Guide
=================

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml)
    
*   [go.mod](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod)
    

This guide provides technical information for developers contributing to the mods codebase. It covers the build system, dependency management, testing practices, and CI/CD infrastructure. For information about the application architecture and runtime behavior, see [Core Architecture](https://deepwiki.com/charmbracelet/mods/2-core-architecture)
. For details on specific subsystems, refer to the child pages below.

Overview
--------

The mods project is written in Go 1.24.0 and uses standard Go tooling for building and testing. The codebase leverages the Charm ecosystem heavily, particularly [bubbletea](https://github.com/charmbracelet/mods/blob/07a05d5b/bubbletea)
 for the TUI framework, [lipgloss](https://github.com/charmbracelet/mods/blob/07a05d5b/lipgloss)
 for styling, and [glamour](https://github.com/charmbracelet/mods/blob/07a05d5b/glamour)
 for markdown rendering. Development workflow is automated through GitHub Actions with cross-platform builds and automated dependency management.

The project follows a modular architecture with clear separation of concerns:

*   **Entry point and CLI**: [main.go](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go)
     with cobra commands
*   **Configuration**: [config.go](https://github.com/charmbracelet/mods/blob/07a05d5b/config.go)
     with YAML/env/flag parsing
*   **UI layer**: [mods.go](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go)
     implementing the Bubble Tea model
*   **AI integration**: Provider-specific logic in [mods.go](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#LNaN-LNaN)
    
*   **Data persistence**: [db.go](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go)
     for SQLite and [cache.go](https://github.com/charmbracelet/mods/blob/07a05d5b/cache.go)
     for message storage

Build System and Dependencies
-----------------------------

### Go Module Configuration

The project uses Go modules for dependency management, defined in [go.mod1-98](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L1-L98)
 The module path is `github.com/charmbracelet/mods` and requires Go 1.24.0 as specified in [go.mod3](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L3-L3)

### Dependency Categories

Dependencies are organized into direct and indirect requirements. Direct dependencies fall into several functional categories:

**Core Framework Dependencies:**

| Package | Version | Purpose |
| --- | --- | --- |
| `github.com/charmbracelet/bubbletea` | v1.3.10 | TUI framework for state machine and event handling |
| `github.com/charmbracelet/lipgloss` | v1.1.1 | Terminal styling and layout |
| `github.com/charmbracelet/glamour` | v0.10.0 | Markdown rendering with syntax highlighting |
| `github.com/charmbracelet/bubbles` | v0.21.1 | Reusable TUI components (viewport, spinner) |

**CLI and Configuration:**

| Package | Version | Purpose |
| --- | --- | --- |
| `github.com/spf13/cobra` | v1.10.1 | Command-line interface framework |
| `github.com/spf13/pflag` | v1.0.10 | POSIX-style flag parsing |
| `github.com/caarlos0/env/v9` | v9.0.0 | Environment variable unmarshalling |
| `gopkg.in/yaml.v3` | v3.0.1 | YAML configuration parsing |
| `github.com/adrg/xdg` | v0.5.3 | XDG base directory specification |

**AI Provider SDKs:**

| Package | Version | Purpose |
| --- | --- | --- |
| `github.com/openai/openai-go` | v1.12.0 | OpenAI GPT models |
| `github.com/anthropics/anthropic-sdk-go` | v1.14.0 | Anthropic Claude models |
| `github.com/cohere-ai/cohere-go/v2` | v2.15.3 | Cohere Command models |
| `github.com/ollama/ollama` | v0.12.6 | Local LLM inference |
| `github.com/mark3labs/mcp-go` | v0.42.0 | Model Context Protocol |

**Data Persistence:**

| Package | Version | Purpose |
| --- | --- | --- |
| `modernc.org/sqlite` | v1.39.1 | Pure Go SQLite implementation |
| `github.com/jmoiron/sqlx` | v1.4.0 | SQL extensions for database operations |

Sources: [go.mod5-38](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L5-L38)

### Dependency Graph

Sources: [go.mod5-38](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L5-L38)

### Build Process

The build process is straightforward using standard Go tooling:

The GitHub Actions workflow in [.github/workflows/build.yml24-30](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L24-L30)
 automates this process on every push and pull request.

### Build Workflow Structure

Sources: [.github/workflows/build.yml1-30](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L1-L30)

Testing Strategy
----------------

### Test File Organization

The mods codebase includes unit tests for core functionality. Test files follow the Go convention of `*_test.go` naming:

*   **mods\_test.go**: Tests for the main Bubble Tea model and message handling
*   **db\_test.go**: Tests for the SQLite conversation database operations
*   **flag\_test.go**: Tests for custom flag parsing logic

### Test Execution

Tests are executed on every CI run with the following parameters:

The `-v` flag provides verbose output, `-cover` generates coverage reports, and `-timeout=30s` ensures tests don't hang indefinitely. This is configured in [.github/workflows/build.yml30](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L30-L30)

### Cross-Platform Testing

The build matrix in [.github/workflows/build.yml8-9](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L8-L9)
 ensures tests run on three operating systems:

This verifies compatibility with different terminal capabilities, file system behaviors, and system APIs that the application depends on (e.g., TTY detection in `term.go`, clipboard operations via `atotto/clipboard`).

Sources: [.github/workflows/build.yml8-30](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L8-L30)

CI/CD Infrastructure
--------------------

### GitHub Actions Workflows

The project uses multiple GitHub Actions workflows for automated quality assurance:

1.  **build.yml**: Main build and test workflow
2.  **lint.yml**: Code quality checks
3.  **lint-sync.yml**: Periodic linting configuration updates
4.  **dependabot.yml**: Automated dependency updates
5.  **dependabot-sync.yml**: Periodic dependency configuration updates

### Build Workflow Jobs

The [build.yml](https://github.com/charmbracelet/mods/blob/07a05d5b/build.yml)
 workflow contains three jobs:

Sources: [.github/workflows/build.yml1-55](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L1-L55)

### Environment Configuration

The build job sets the `GO111MODULE` environment variable to enable Go module support explicitly:

This is defined in [.github/workflows/build.yml11-12](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L11-L12)

### Snapshot Release Job

The `snapshot` job at [.github/workflows/build.yml32-35](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L32-L35)
 leverages a reusable workflow from the Charm meta repository:

This creates development builds for testing without publishing full releases.

### Dependabot Auto-Merge

The `dependabot` job at [.github/workflows/build.yml37-54](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L37-L54)
 automatically approves and merges dependency update PRs when builds pass:

The job only runs when:

*   The actor is `dependabot[bot]`
*   The event is a pull request
*   The build job succeeds

This is enforced by the conditions in [.github/workflows/build.yml43](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L43-L43)
:

Sources: [.github/workflows/build.yml37-54](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L37-L54)

### Permissions Model

The dependabot job requires specific permissions defined in [.github/workflows/build.yml40-42](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L40-L42)
:

These permissions allow the workflow to:

*   Approve pull requests (`pull-requests: write`)
*   Merge changes into the repository (`contents: write`)

### Integration with Charm Ecosystem

The project integrates with shared workflows from the `charmbracelet/meta` repository, promoting consistency across the Charm ecosystem. The reusable workflow pattern allows centralized maintenance of common CI/CD tasks like snapshot releases and linting configuration synchronization.

Sources: [.github/workflows/build.yml32-35](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L32-L35)

Development Workflow
--------------------

### Local Development

For local development, developers should:

1.  Clone the repository
2.  Run `go mod download` to fetch dependencies
3.  Build with `go build` to create the `mods` binary
4.  Run tests with `go test ./...` before submitting changes
5.  Use `go fmt` to format code according to Go conventions

### Contribution Guidelines

The automated CI/CD pipeline ensures that:

*   All changes build successfully on Linux, macOS, and Windows
*   Tests pass with adequate coverage
*   Code meets linting standards (enforced via lint.yml, not shown in provided files)
*   Dependencies are kept up to date via Dependabot

Pull requests must pass all CI checks before merging. The dependabot auto-merge mechanism reduces maintenance burden for routine dependency updates while maintaining quality through the full test suite.

Sources: [.github/workflows/build.yml1-55](https://github.com/charmbracelet/mods/blob/07a05d5b/.github/workflows/build.yml#L1-L55)
 [go.mod1-98](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L1-L98)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development Guide](https://deepwiki.com/charmbracelet/mods/7-development-guide#development-guide)
    
*   [Overview](https://deepwiki.com/charmbracelet/mods/7-development-guide#overview)
    
*   [Build System and Dependencies](https://deepwiki.com/charmbracelet/mods/7-development-guide#build-system-and-dependencies)
    
*   [Go Module Configuration](https://deepwiki.com/charmbracelet/mods/7-development-guide#go-module-configuration)
    
*   [Dependency Categories](https://deepwiki.com/charmbracelet/mods/7-development-guide#dependency-categories)
    
*   [Dependency Graph](https://deepwiki.com/charmbracelet/mods/7-development-guide#dependency-graph)
    
*   [Build Process](https://deepwiki.com/charmbracelet/mods/7-development-guide#build-process)
    
*   [Build Workflow Structure](https://deepwiki.com/charmbracelet/mods/7-development-guide#build-workflow-structure)
    
*   [Testing Strategy](https://deepwiki.com/charmbracelet/mods/7-development-guide#testing-strategy)
    
*   [Test File Organization](https://deepwiki.com/charmbracelet/mods/7-development-guide#test-file-organization)
    
*   [Test Execution](https://deepwiki.com/charmbracelet/mods/7-development-guide#test-execution)
    
*   [Cross-Platform Testing](https://deepwiki.com/charmbracelet/mods/7-development-guide#cross-platform-testing)
    
*   [CI/CD Infrastructure](https://deepwiki.com/charmbracelet/mods/7-development-guide#cicd-infrastructure)
    
*   [GitHub Actions Workflows](https://deepwiki.com/charmbracelet/mods/7-development-guide#github-actions-workflows)
    
*   [Build Workflow Jobs](https://deepwiki.com/charmbracelet/mods/7-development-guide#build-workflow-jobs)
    
*   [Environment Configuration](https://deepwiki.com/charmbracelet/mods/7-development-guide#environment-configuration)
    
*   [Snapshot Release Job](https://deepwiki.com/charmbracelet/mods/7-development-guide#snapshot-release-job)
    
*   [Dependabot Auto-Merge](https://deepwiki.com/charmbracelet/mods/7-development-guide#dependabot-auto-merge)
    
*   [Permissions Model](https://deepwiki.com/charmbracelet/mods/7-development-guide#permissions-model)
    
*   [Integration with Charm Ecosystem](https://deepwiki.com/charmbracelet/mods/7-development-guide#integration-with-charm-ecosystem)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/mods/7-development-guide#development-workflow)
    
*   [Local Development](https://deepwiki.com/charmbracelet/mods/7-development-guide#local-development)
    
*   [Contribution Guidelines](https://deepwiki.com/charmbracelet/mods/7-development-guide#contribution-guidelines)
