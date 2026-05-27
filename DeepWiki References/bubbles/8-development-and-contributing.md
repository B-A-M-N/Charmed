Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/bubbles](https://github.com/charmbracelet/bubbles "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 19 April 2026 ([42130e](https://github.com/charmbracelet/bubbles/commits/42130e89)
)

*   [Overview](https://deepwiki.com/charmbracelet/bubbles/1-overview)
    
*   [Architecture and Concepts](https://deepwiki.com/charmbracelet/bubbles/1.1-architecture-and-concepts)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/bubbles/1.2-getting-started)
    
*   [Input Components](https://deepwiki.com/charmbracelet/bubbles/2-input-components)
    
*   [Text Input](https://deepwiki.com/charmbracelet/bubbles/2.1-text-input)
    
*   [Text Area](https://deepwiki.com/charmbracelet/bubbles/2.2-text-area)
    
*   [Display Components](https://deepwiki.com/charmbracelet/bubbles/3-display-components)
    
*   [Viewport](https://deepwiki.com/charmbracelet/bubbles/3.1-viewport)
    
*   [List](https://deepwiki.com/charmbracelet/bubbles/3.2-list)
    
*   [Table](https://deepwiki.com/charmbracelet/bubbles/3.3-table)
    
*   [Progress Indicators](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators)
    
*   [Progress Bar](https://deepwiki.com/charmbracelet/bubbles/4.1-progress-bar)
    
*   [Spinner](https://deepwiki.com/charmbracelet/bubbles/4.2-spinner)
    
*   [Navigation Components](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components)
    
*   [Paginator](https://deepwiki.com/charmbracelet/bubbles/5.1-paginator)
    
*   [File Picker](https://deepwiki.com/charmbracelet/bubbles/5.2-file-picker)
    
*   [Time Components](https://deepwiki.com/charmbracelet/bubbles/6-time-components)
    
*   [Stopwatch and Timer](https://deepwiki.com/charmbracelet/bubbles/6.1-stopwatch-and-timer)
    
*   [Utility Components](https://deepwiki.com/charmbracelet/bubbles/7-utility-components)
    
*   [Help](https://deepwiki.com/charmbracelet/bubbles/7.1-help)
    
*   [Cursor](https://deepwiki.com/charmbracelet/bubbles/7.2-cursor)
    
*   [Development and Contributing](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing)
    
*   [Testing and Linting](https://deepwiki.com/charmbracelet/bubbles/8.1-testing-and-linting)
    
*   [CI/CD and Releases](https://deepwiki.com/charmbracelet/bubbles/8.2-cicd-and-releases)
    
*   [Migrating from v1 to v2](https://deepwiki.com/charmbracelet/bubbles/8.3-migrating-from-v1-to-v2)
    
*   [Glossary](https://deepwiki.com/charmbracelet/bubbles/9-glossary)
    

Menu

Development and Contributing
============================

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/build.yml)
    
*   [.github/workflows/coverage.yml](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/coverage.yml)
    
*   [.github/workflows/lint-sync.yml](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/lint-sync.yml)
    
*   [.github/workflows/lint.yml](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/lint.yml)
    
*   [.github/workflows/release.yml](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/release.yml)
    
*   [.gitignore](https://github.com/charmbracelet/bubbles/blob/42130e89/.gitignore)
    
*   [.golangci.yml](https://github.com/charmbracelet/bubbles/blob/42130e89/.golangci.yml)
    
*   [.goreleaser.yml](https://github.com/charmbracelet/bubbles/blob/42130e89/.goreleaser.yml)
    

This page provides a high-level overview for developers who want to contribute to the `bubbles` UI component library. It covers the development environment, coding standards, and the automated workflows that govern the codebase.

The library is currently transitioning to **v2**. For developers looking to migrate existing projects or understand the differences between versions, see [Migrating from v1 to v2](https://deepwiki.com/charmbracelet/bubbles/8.3-migrating-from-v1-to-v2)
.

Development Environment and Tools
---------------------------------

To contribute, you must have a Go environment configured (supporting `stable` and `oldstable` versions). The project uses several tools to maintain code quality and automate repetitive tasks.

| Tool | Purpose | Configuration |
| --- | --- | --- |
| **Go** | Primary runtime and build system | [go.mod](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod) |
| **golangci-lint** | Aggregated linting engine | [.golangci.yml1-50](https://github.com/charmbracelet/bubbles/blob/42130e89/.golangci.yml#L1-L50) |
| **GoReleaser** | Release automation | [.goreleaser.yml1-6](https://github.com/charmbracelet/bubbles/blob/42130e89/.goreleaser.yml#L1-L6) |
| **GitHub Actions** | CI/CD pipeline | [.github/workflows/](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/) |

### Core Commands

*   **Build**: `go build -v ./...` [.github/workflows/build.yml25](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.github/workflows/build.yml#L25-L25)
    
*   **Test**: `go test ./...` [.github/workflows/build.yml28](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.github/workflows/build.yml#L28-L28)
    
*   **Lint**: Handled via GitHub Actions using the `.golangci.yml` configuration. [.github/workflows/lint.yml8-12](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.github/workflows/lint.yml#L8-L12)
    

Sources: [.github/workflows/build.yml7-28](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/build.yml#L7-L28)
 [.golangci.yml1-50](https://github.com/charmbracelet/bubbles/blob/42130e89/.golangci.yml#L1-L50)
 [.goreleaser.yml1-6](https://github.com/charmbracelet/bubbles/blob/42130e89/.goreleaser.yml#L1-L6)

Coding Standards and Linting
----------------------------

The project enforces strict coding standards to ensure consistency across the various UI components. This is managed through a comprehensive `golangci-lint` configuration.

For a deep dive into the specific linting rules and the testing patterns used (such as golden files and internal utilities), see **[Testing and Linting](https://deepwiki.com/charmbracelet/bubbles/8.1-testing-and-linting)
**.

### Linter Pipeline

The following diagram maps the linting requirements to the specific linters enabled in the codebase:

Sources: [.golangci.yml5-28](https://github.com/charmbracelet/bubbles/blob/42130e89/.golangci.yml#L5-L28)

CI/CD and Release Process
-------------------------

The `bubbles` library utilizes GitHub Actions for a robust CI/CD pipeline. Every push and pull request triggers a battery of tests across different operating systems and Go versions.

For detailed information on the workflow triggers, matrix strategies, and how releases are tagged and published, see **[CI/CD and Releases](https://deepwiki.com/charmbracelet/bubbles/8.2-cicd-and-releases)
**.

### CI/CD Workflow Architecture

Sources: [.github/workflows/build.yml1-10](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/build.yml#L1-L10)
 [.github/workflows/coverage.yml1-2](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/coverage.yml#L1-L2)
 [.github/workflows/lint.yml1-4](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/lint.yml#L1-L4)
 [.github/workflows/release.yml1-6](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/release.yml#L1-L6)
 [.github/workflows/lint-sync.yml1-5](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/lint-sync.yml#L1-L5)

Contribution Workflow
---------------------

1.  **Issue/Feature**: Identify a bug or feature request.
2.  **Development**: Implement changes locally. Ensure `go test ./...` passes on your machine. [.github/workflows/build.yml28](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.github/workflows/build.yml#L28-L28)
    
3.  **Linting**: Verify changes against the rules in `.golangci.yml`. [.golangci.yml4-28](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.golangci.yml#L4-L28)
    
4.  **Pull Request**: Submit a PR. This triggers the `build`, `lint`, and `coverage` workflows. [.github/workflows/build.yml3](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.github/workflows/build.yml#L3-L3)
     [.github/workflows/lint.yml4](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.github/workflows/lint.yml#L4-L4)
     [.github/workflows/coverage.yml2](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.github/workflows/coverage.yml#L2-L2)
    
5.  **Review**: Codeowners review the PR. Automated dependency updates are handled by `dependabot`. [.github/workflows/build.yml30-47](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.github/workflows/build.yml#L30-L47)
    
6.  **Merge & Release**: Once merged, releases are handled by `goreleaser` via Git tags. [.github/workflows/release.yml13-14](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.github/workflows/release.yml#L13-L14)
     [.goreleaser.yml1-6](https://github.com/charmbracelet/bubbles/blob/42130e89/%20.goreleaser.yml#L1-L6)
    

Sources: [.github/workflows/build.yml30-47](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/build.yml#L30-L47)
 [.github/workflows/release.yml1-14](https://github.com/charmbracelet/bubbles/blob/42130e89/.github/workflows/release.yml#L1-L14)
 [.goreleaser.yml1-6](https://github.com/charmbracelet/bubbles/blob/42130e89/.goreleaser.yml#L1-L6)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development and Contributing](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing#development-and-contributing)
    
*   [Development Environment and Tools](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing#development-environment-and-tools)
    
*   [Core Commands](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing#core-commands)
    
*   [Coding Standards and Linting](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing#coding-standards-and-linting)
    
*   [Linter Pipeline](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing#linter-pipeline)
    
*   [CI/CD and Release Process](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing#cicd-and-release-process)
    
*   [CI/CD Workflow Architecture](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing#cicd-workflow-architecture)
    
*   [Contribution Workflow](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing#contribution-workflow)
