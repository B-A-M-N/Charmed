Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/gum](https://github.com/charmbracelet/gum "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 2 May 2026 ([30cc51](https://github.com/charmbracelet/gum/commits/30cc5169)
)

*   [Overview](https://deepwiki.com/charmbracelet/gum/1-overview)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/gum/2-getting-started)
    
*   [Architecture](https://deepwiki.com/charmbracelet/gum/3-architecture)
    
*   [Application Entry Point](https://deepwiki.com/charmbracelet/gum/3.1-application-entry-point)
    
*   [Framework Components](https://deepwiki.com/charmbracelet/gum/3.2-framework-components)
    
*   [Command Patterns](https://deepwiki.com/charmbracelet/gum/3.3-command-patterns)
    
*   [Interactive Commands](https://deepwiki.com/charmbracelet/gum/4-interactive-commands)
    
*   [Choose Command](https://deepwiki.com/charmbracelet/gum/4.1-choose-command)
    
*   [Confirm Command](https://deepwiki.com/charmbracelet/gum/4.2-confirm-command)
    
*   [Filter Command](https://deepwiki.com/charmbracelet/gum/4.3-filter-command)
    
*   [Input Command](https://deepwiki.com/charmbracelet/gum/4.4-input-command)
    
*   [Write Command](https://deepwiki.com/charmbracelet/gum/4.5-write-command)
    
*   [File Command](https://deepwiki.com/charmbracelet/gum/4.6-file-command)
    
*   [Table Command](https://deepwiki.com/charmbracelet/gum/4.7-table-command)
    
*   [Pager Command](https://deepwiki.com/charmbracelet/gum/4.8-pager-command)
    
*   [Progress and Background Commands](https://deepwiki.com/charmbracelet/gum/5-progress-and-background-commands)
    
*   [Spin Command](https://deepwiki.com/charmbracelet/gum/5.1-spin-command)
    
*   [Formatting Commands](https://deepwiki.com/charmbracelet/gum/6-formatting-commands)
    
*   [Style Command](https://deepwiki.com/charmbracelet/gum/6.1-style-command)
    
*   [Format Command](https://deepwiki.com/charmbracelet/gum/6.2-format-command)
    
*   [Join Command](https://deepwiki.com/charmbracelet/gum/6.3-join-command)
    
*   [Log Command](https://deepwiki.com/charmbracelet/gum/6.4-log-command)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/gum/7-development-guide)
    
*   [Building and Testing](https://deepwiki.com/charmbracelet/gum/7.1-building-and-testing)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/gum/7.2-dependency-management)
    
*   [CI/CD Pipeline](https://deepwiki.com/charmbracelet/gum/7.3-cicd-pipeline)
    
*   [Shell Completion](https://deepwiki.com/charmbracelet/gum/7.4-shell-completion)
    
*   [Internal Utilities](https://deepwiki.com/charmbracelet/gum/8-internal-utilities)
    
*   [Glossary](https://deepwiki.com/charmbracelet/gum/9-glossary)
    

Menu

Development Guide
=================

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/build.yml)
    
*   [.github/workflows/goreleaser.yml](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/goreleaser.yml)
    
*   [.github/workflows/lint-sync.yml](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/lint-sync.yml)
    
*   [.github/workflows/lint.yml](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/lint.yml)
    
*   [.github/workflows/nightly.yml](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/nightly.yml)
    
*   [.golangci.yml](https://github.com/charmbracelet/gum/blob/30cc5169/.golangci.yml)
    
*   [.goreleaser.yml](https://github.com/charmbracelet/gum/blob/30cc5169/.goreleaser.yml)
    

This document serves as the high-level entry point for developers looking to contribute to `gum`. It outlines the development environment, the lifecycle of a contribution, and the automation tools used to maintain code quality.

High-Level Workflow
-------------------

The `gum` development workflow is centered around the Go toolchain and GitHub Actions. The project relies heavily on the Charmbracelet ecosystem to provide TUI components.

Sources: [.github/workflows/lint.yml1-9](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/lint.yml#L1-L9)
 [.github/workflows/build.yml1-14](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/build.yml#L1-L14)
 [.github/workflows/goreleaser.yml1-16](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/goreleaser.yml#L1-L16)

Building and Testing
--------------------

`gum` is built using the standard Go toolchain. Developers should ensure they are using a recent version of Go as defined in the project's module configuration. The test suite covers both logic and TUI interactions.

For detailed instructions on local builds and organizing tests, see **[Building and Testing](https://deepwiki.com/charmbracelet/gum/7.1-building-and-testing)
**.

Sources: [.github/workflows/build.yml10-11](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/build.yml#L10-L11)

Dependency Management
---------------------

The project is built on the **Bubble Tea** framework and uses several sibling libraries from the Charmbracelet ecosystem (e.g., `bubbles`, `lipgloss`, `log`). Dependencies are managed via Go modules.

For a full list of core dependencies and how they are updated, see **[Dependency Management](https://deepwiki.com/charmbracelet/gum/7.2-dependency-management)
**.

CI/CD Pipeline
--------------

`gum` utilizes reusable workflows from `charmbracelet/meta` to standardize linting, building, and releasing across the organization.

### Automated Checks and Releases

| Workflow | Trigger | Purpose |
| --- | --- | --- |
| `lint.yml` | Push/PR | Runs `golangci-lint` with settings from `.golangci.yml`. |
| `build.yml` | Push/PR | Verifies the project compiles across platforms. |
| `goreleaser.yml` | Tag (`v*`) | Orchestrates the full release process using `goreleaser`. |
| `nightly.yml` | Push to `main` | Generates nightly builds for early testing. |
| `lint-sync.yml` | Schedule | Ensures linting rules stay synchronized. |

For details on how GitHub Actions and GoReleaser are configured, see **[CI/CD Pipeline](https://deepwiki.com/charmbracelet/gum/7.3-cicd-pipeline)
**.

Sources: [.github/workflows/lint.yml1-9](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/lint.yml#L1-L9)
 [.github/workflows/build.yml1-14](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/build.yml#L1-L14)
 [.github/workflows/goreleaser.yml1-16](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/goreleaser.yml#L1-L16)
 [.github/workflows/nightly.yml1-10](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/nightly.yml#L1-L10)
 [.github/workflows/lint-sync.yml1-15](https://github.com/charmbracelet/gum/blob/30cc5169/.github/workflows/lint-sync.yml#L1-L15)

Code Quality Standards
----------------------

The project enforces strict code quality through `.golangci.yml`. This configuration enables a wide array of linters including `gosec` for security, `revive` for style, and `wrapcheck` to ensure error wrapping.

Sources: [.golangci.yml4-27](https://github.com/charmbracelet/gum/blob/30cc5169/.golangci.yml#L4-L27)
 [.golangci.yml42-45](https://github.com/charmbracelet/gum/blob/30cc5169/.golangci.yml#L42-L45)

Shell Completion
----------------

`gum` supports shell completion for `bash`, `zsh`, `fish`, and `powershell`. These completions are generated dynamically based on the CLI structure defined via the `kong` library.

For details on generating and installing these scripts, see **[Shell Completion](https://deepwiki.com/charmbracelet/gum/7.4-shell-completion)
**.

Release Configuration
---------------------

The release process is governed by `.goreleaser.yml`. It defines how binaries are packaged, including metadata like the maintainer and project description. It also handles integration with package managers like Homebrew and Scoop.

| Variable | Value |
| --- | --- |
| `scoop_name` | `charm-gum` |
| `maintainer` | `Maas Lalani <maas@charm.sh>` |
| `github_url` | `https://github.com/charmbracelet/gum` |

Sources: [.goreleaser.yml9-17](https://github.com/charmbracelet/gum/blob/30cc5169/.goreleaser.yml#L9-L17)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development Guide](https://deepwiki.com/charmbracelet/gum/7-development-guide#development-guide)
    
*   [High-Level Workflow](https://deepwiki.com/charmbracelet/gum/7-development-guide#high-level-workflow)
    
*   [Building and Testing](https://deepwiki.com/charmbracelet/gum/7-development-guide#building-and-testing)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/gum/7-development-guide#dependency-management)
    
*   [CI/CD Pipeline](https://deepwiki.com/charmbracelet/gum/7-development-guide#cicd-pipeline)
    
*   [Automated Checks and Releases](https://deepwiki.com/charmbracelet/gum/7-development-guide#automated-checks-and-releases)
    
*   [Code Quality Standards](https://deepwiki.com/charmbracelet/gum/7-development-guide#code-quality-standards)
    
*   [Shell Completion](https://deepwiki.com/charmbracelet/gum/7-development-guide#shell-completion)
    
*   [Release Configuration](https://deepwiki.com/charmbracelet/gum/7-development-guide#release-configuration)
