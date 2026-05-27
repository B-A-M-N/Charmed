Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/glow](https://github.com/charmbracelet/glow "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 2 January 2026 ([752de9](https://github.com/charmbracelet/glow/commits/752de97c)
)

*   [Overview](https://deepwiki.com/charmbracelet/glow/1-overview)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/glow/1.1-dependencies)
    
*   [Installation and Usage](https://deepwiki.com/charmbracelet/glow/1.2-installation-and-usage)
    
*   [Architecture](https://deepwiki.com/charmbracelet/glow/2-architecture)
    
*   [Command-Line Interface](https://deepwiki.com/charmbracelet/glow/2.1-command-line-interface)
    
*   [Terminal User Interface](https://deepwiki.com/charmbracelet/glow/2.2-terminal-user-interface)
    
*   [Stash View](https://deepwiki.com/charmbracelet/glow/2.2.1-stash-view)
    
*   [Document View](https://deepwiki.com/charmbracelet/glow/2.2.2-document-view)
    
*   [Markdown Processing](https://deepwiki.com/charmbracelet/glow/2.3-markdown-processing)
    
*   [UI Components and Styling](https://deepwiki.com/charmbracelet/glow/2.4-ui-components-and-styling)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/glow/3-configuration-system)
    
*   [Developer Guide](https://deepwiki.com/charmbracelet/glow/4-developer-guide)
    
*   [Build and Release Process](https://deepwiki.com/charmbracelet/glow/4.1-build-and-release-process)
    
*   [Code Quality and Testing](https://deepwiki.com/charmbracelet/glow/4.2-code-quality-and-testing)
    

Menu

Developer Guide
===============

Relevant source files

*   [.github/workflows/goreleaser.yml](https://github.com/charmbracelet/glow/blob/752de97c/.github/workflows/goreleaser.yml)
    
*   [.golangci.yml](https://github.com/charmbracelet/glow/blob/752de97c/.golangci.yml)
    
*   [.goreleaser.yml](https://github.com/charmbracelet/glow/blob/752de97c/.goreleaser.yml)
    
*   [console\_windows.go](https://github.com/charmbracelet/glow/blob/752de97c/console_windows.go)
    
*   [go.mod](https://github.com/charmbracelet/glow/blob/752de97c/go.mod)
    

This page provides an overview for developers who want to contribute to Glow, understand the codebase structure, and extend its functionality. It covers the development environment setup, project organization, key entry points for modifications, and common development workflows.

For detailed information about the automated build and release process, see [Build and Release Process](https://deepwiki.com/charmbracelet/glow/4.1-build-and-release-process)
. For code quality tooling and testing practices, see [Code Quality and Testing](https://deepwiki.com/charmbracelet/glow/4.2-code-quality-and-testing)
.

Development Environment Setup
-----------------------------

### Go Version Requirements

Glow requires Go 1.24.0 or higher, with the project using toolchain go1.24.1. This is specified in [go.mod3-5](https://github.com/charmbracelet/glow/blob/752de97c/go.mod#L3-L5)
 and ensures compatibility with the latest Go features and module system improvements.

### Repository Structure

    github.com/charmbracelet/glow/
    ├── main.go              # CLI entry point, cobra command setup
    ├── ui/                  # TUI implementation (bubbletea models)
    │   ├── ui.go           # NewProgram, model initialization
    │   ├── stash.go        # stashModel (file browser)
    │   ├── pager.go        # pagerModel (document viewer)
    │   ├── markdown.go     # markdown struct, content handling
    │   └── styles.go       # lipgloss styling definitions
    ├── go.mod              # Module definition, dependencies
    ├── .goreleaser.yml     # Release automation config
    ├── .golangci.yml       # Linter configuration
    └── .github/workflows/  # CI/CD workflows
    

### Dependency Installation

All dependencies are managed through Go modules. After cloning the repository:

The project relies heavily on the Charmbracelet ecosystem as defined in [go.mod7-32](https://github.com/charmbracelet/glow/blob/752de97c/go.mod#L7-L32)
:

| Dependency | Purpose | Version |
| --- | --- | --- |
| `github.com/charmbracelet/bubbletea` | TUI framework (MVU architecture) | v1.3.10 |
| `github.com/charmbracelet/bubbles` | Reusable UI components | v0.21.0 |
| `github.com/charmbracelet/lipgloss` | Terminal styling | v1.1.1 |
| `github.com/charmbracelet/glamour` | Markdown rendering | v0.10.0 |
| `github.com/spf13/cobra` | CLI framework | v1.10.2 |
| `github.com/spf13/viper` | Configuration management | v1.21.0 |
| `github.com/fsnotify/fsnotify` | File watching | v1.9.0 |
| `github.com/sahilm/fuzzy` | Fuzzy search | v0.1.1 |

Sources: [go.mod1-76](https://github.com/charmbracelet/glow/blob/752de97c/go.mod#L1-L76)

Project Structure and Code Organization
---------------------------------------

The following diagram maps the high-level systems to their concrete code locations:

Sources: [go.mod1-76](https://github.com/charmbracelet/glow/blob/752de97c/go.mod#L1-L76)
 based on architecture diagrams

Development Workflow
--------------------

### Building Locally

Build the binary for your platform:

The resulting `glow` binary can be run directly:

### Running in Development Mode

For rapid iteration during development:

### Common Development Tasks

| Task | Command | Description |
| --- | --- | --- |
| Build | `go build` | Compile binary |
| Test | `go test ./...` | Run all tests |
| Lint | `golangci-lint run` | Run linters (requires golangci-lint) |
| Format | `gofumpt -w .` | Format code |
| Mod tidy | `go mod tidy` | Clean up dependencies |
| Generate man page | `./glow man > glow.1` | Create man page |

### Windows-Specific Considerations

The project includes Windows-specific code for ANSI color support in [console\_windows.go1-24](https://github.com/charmbracelet/glow/blob/752de97c/console_windows.go#L1-L24)
 This file uses build tags (`// +build windows`) and is automatically included only on Windows builds. It enables virtual terminal processing to support ANSI escape sequences in the Windows console [console\_windows.go13-19](https://github.com/charmbracelet/glow/blob/752de97c/console_windows.go#L13-L19)

Sources: [console\_windows.go1-24](https://github.com/charmbracelet/glow/blob/752de97c/console_windows.go#L1-L24)

Key Extension Points
--------------------

The following diagram identifies the primary code locations where developers typically add new features or modify behavior:

Sources: Based on architecture overview and code structure

### Extension Point Details

#### Adding New CLI Commands

New commands follow the cobra pattern. Add a new `cobra.Command` instance and register it with `rootCmd.AddCommand()` in `main.go`. The existing `manCmd` and `configCmd` serve as examples of subcommand implementation.

#### Modifying TUI State Transitions

The TUI uses the Elm Architecture (MVU pattern). State transitions occur in:

*   `model.Update()` in `ui/ui.go` - top-level message routing
*   `stashModel.update()` in `ui/stash.go` - file browser interactions
*   `pagerModel.update()` in `ui/pager.go` - document viewer interactions

To add new keybindings or state transitions, modify these `Update` methods to handle new `tea.Msg` types.

#### Extending Source Resolution

The `sourceFromArg()` function in `main.go` implements a waterfall resolution strategy:

1.  Check if stdin has content
2.  Check if argument is a GitHub/GitLab URL
3.  Check if argument is an HTTP URL
4.  Check if argument is a local file path

To support new source types (e.g., Bitbucket, Gitea), add new conditional branches in this function and implement corresponding resolver functions similar to `readmeURL()`.

#### Customizing Rendering

Markdown rendering configuration happens in:

*   `glamourRender()` in `ui/pager.go` - creates `glamour.TermRenderer` with style and width settings
*   Style selection is based on `Config.Style` which comes from CLI flags or `glow.yml`

To add new rendering options, modify the `glamour.NewTermRenderer()` call and extend `Config` with new fields.

#### Styling Modifications

All lipgloss style definitions live in `ui/styles.go`. These styles are used throughout:

*   `stashItemView()` - renders each item in the file browser
*   `pagerModel.View()` - renders the document viewer
*   Various helper functions for status messages and UI chrome

Modify these `lipgloss.Style` definitions to change colors, borders, padding, and other visual aspects.

Sources: Based on architecture diagrams and code structure

Code Organization Principles
----------------------------

### Separation of Concerns

Glow maintains clear boundaries between:

1.  **CLI Layer** (`main.go`) - Argument parsing, configuration validation, direct rendering
2.  **TUI Layer** (`ui/` package) - Interactive file browsing and viewing
3.  **Processing Layer** (functions in `main.go` and `ui/`) - Content fetching, frontmatter removal, markdown rendering

### State Management

The TUI follows bubbletea's MVU (Model-View-Update) architecture:

*   **Model**: Application state (`model`, `stashModel`, `pagerModel`)
*   **View**: Rendering functions (`View()` methods)
*   **Update**: State transitions (`Update()` methods handling `tea.Msg`)

State is immutable - `Update()` methods return new model instances rather than modifying in-place.

### Configuration Hierarchy

Configuration follows a three-tier priority system implemented in `validateOptions()`:

1.  CLI flags (highest priority)
2.  Environment variables
3.  `glow.yml` file (lowest priority)

This is managed through viper's automatic binding and override system.

### Platform Abstraction

Platform-specific code uses build tags:

*   `console_windows.go` has `// +build windows` for Windows ANSI support
*   The Go build system automatically includes/excludes these files

When adding platform-specific code, use the same pattern with appropriate build tags.

Sources: Based on architecture overview, [console\_windows.go1-24](https://github.com/charmbracelet/glow/blob/752de97c/console_windows.go#L1-L24)

Development Best Practices
--------------------------

### Dependency Management

*   Always run `go mod tidy` after adding or removing dependencies
*   Pin major versions for stability (note the `v2` suffix in the module path: `github.com/charmbracelet/glow/v2`)
*   The `go.mod` file uses `toolchain` directive [go.mod5](https://github.com/charmbracelet/glow/blob/752de97c/go.mod#L5-L5)
     to ensure consistent builds across developers

### Linting and Formatting

The project uses golangci-lint with a comprehensive ruleset defined in [.golangci.yml1-45](https://github.com/charmbracelet/glow/blob/752de97c/.golangci.yml#L1-L45)
 Key enabled linters include:

*   `goconst` - Detect repeated constants
*   `gosec` - Security issues
*   `misspell` - Spelling errors
*   `revive` - Style mistakes
*   `unparam` - Unused parameters

Formatters are automatically applied:

*   `gofumpt` - Stricter gofmt
*   `goimports` - Import organization

Run these before committing:

### Release Process

Releases are automated using goreleaser, configured in [.goreleaser.yml1-15](https://github.com/charmbracelet/glow/blob/752de97c/.goreleaser.yml#L1-L15)
 The configuration includes shared settings from the Charmbracelet meta repository:

This provides multi-platform builds, package generation, and artifact publishing. See [Build and Release Process](https://deepwiki.com/charmbracelet/glow/4.1-build-and-release-process)
 for details.

Sources: [.golangci.yml1-45](https://github.com/charmbracelet/glow/blob/752de97c/.golangci.yml#L1-L45)
 [.goreleaser.yml1-15](https://github.com/charmbracelet/glow/blob/752de97c/.goreleaser.yml#L1-L15)

Testing and Debugging
---------------------

### Running Tests

### Debugging TUI Applications

Debugging TUI apps requires special handling since they control the terminal:

1.  **Log to file**: Use `charmbracelet/log` package to write to a file
2.  **Dual terminal**: Run the app in one terminal, tail logs in another
3.  **Printf debugging**: Temporarily disable TUI mode and add print statements

### Manual Testing Checklist

When making changes, manually test:

*   [ ]  CLI mode: `./glow README.md`
*   [ ]  TUI mode: `./glow`
*   [ ]  Stdin: `cat README.md | ./glow -`
*   [ ]  URLs: `./glow https://github.com/charmbracelet/glow/blob/master/README.md`
*   [ ]  Config command: `./glow config`
*   [ ]  Man command: `./glow man`
*   [ ]  Different terminal sizes (resize during operation)
*   [ ]  File watching (edit file while viewing in TUI)

Contributing Guidelines
-----------------------

When preparing contributions:

1.  **Fork and branch**: Create a feature branch from `master`
2.  **Make atomic commits**: Each commit should be a logical unit
3.  **Follow conventions**: Match existing code style and structure
4.  **Test thoroughly**: Include both automated and manual testing
5.  **Document changes**: Update relevant wiki pages if architecture changes
6.  **Run linters**: Ensure `golangci-lint run` passes
7.  **Submit PR**: Reference any related issues

The project uses GitHub Actions for CI/CD, which automatically runs linters and tests on pull requests. See [Code Quality and Testing](https://deepwiki.com/charmbracelet/glow/4.2-code-quality-and-testing)
 for details on the CI pipeline.

Sources: [.github/workflows/goreleaser.yml1-26](https://github.com/charmbracelet/glow/blob/752de97c/.github/workflows/goreleaser.yml#L1-L26)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Developer Guide](https://deepwiki.com/charmbracelet/glow/4-developer-guide#developer-guide)
    
*   [Development Environment Setup](https://deepwiki.com/charmbracelet/glow/4-developer-guide#development-environment-setup)
    
*   [Go Version Requirements](https://deepwiki.com/charmbracelet/glow/4-developer-guide#go-version-requirements)
    
*   [Repository Structure](https://deepwiki.com/charmbracelet/glow/4-developer-guide#repository-structure)
    
*   [Dependency Installation](https://deepwiki.com/charmbracelet/glow/4-developer-guide#dependency-installation)
    
*   [Project Structure and Code Organization](https://deepwiki.com/charmbracelet/glow/4-developer-guide#project-structure-and-code-organization)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/glow/4-developer-guide#development-workflow)
    
*   [Building Locally](https://deepwiki.com/charmbracelet/glow/4-developer-guide#building-locally)
    
*   [Running in Development Mode](https://deepwiki.com/charmbracelet/glow/4-developer-guide#running-in-development-mode)
    
*   [Common Development Tasks](https://deepwiki.com/charmbracelet/glow/4-developer-guide#common-development-tasks)
    
*   [Windows-Specific Considerations](https://deepwiki.com/charmbracelet/glow/4-developer-guide#windows-specific-considerations)
    
*   [Key Extension Points](https://deepwiki.com/charmbracelet/glow/4-developer-guide#key-extension-points)
    
*   [Extension Point Details](https://deepwiki.com/charmbracelet/glow/4-developer-guide#extension-point-details)
    
*   [Adding New CLI Commands](https://deepwiki.com/charmbracelet/glow/4-developer-guide#adding-new-cli-commands)
    
*   [Modifying TUI State Transitions](https://deepwiki.com/charmbracelet/glow/4-developer-guide#modifying-tui-state-transitions)
    
*   [Extending Source Resolution](https://deepwiki.com/charmbracelet/glow/4-developer-guide#extending-source-resolution)
    
*   [Customizing Rendering](https://deepwiki.com/charmbracelet/glow/4-developer-guide#customizing-rendering)
    
*   [Styling Modifications](https://deepwiki.com/charmbracelet/glow/4-developer-guide#styling-modifications)
    
*   [Code Organization Principles](https://deepwiki.com/charmbracelet/glow/4-developer-guide#code-organization-principles)
    
*   [Separation of Concerns](https://deepwiki.com/charmbracelet/glow/4-developer-guide#separation-of-concerns)
    
*   [State Management](https://deepwiki.com/charmbracelet/glow/4-developer-guide#state-management)
    
*   [Configuration Hierarchy](https://deepwiki.com/charmbracelet/glow/4-developer-guide#configuration-hierarchy)
    
*   [Platform Abstraction](https://deepwiki.com/charmbracelet/glow/4-developer-guide#platform-abstraction)
    
*   [Development Best Practices](https://deepwiki.com/charmbracelet/glow/4-developer-guide#development-best-practices)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/glow/4-developer-guide#dependency-management)
    
*   [Linting and Formatting](https://deepwiki.com/charmbracelet/glow/4-developer-guide#linting-and-formatting)
    
*   [Release Process](https://deepwiki.com/charmbracelet/glow/4-developer-guide#release-process)
    
*   [Testing and Debugging](https://deepwiki.com/charmbracelet/glow/4-developer-guide#testing-and-debugging)
    
*   [Running Tests](https://deepwiki.com/charmbracelet/glow/4-developer-guide#running-tests)
    
*   [Debugging TUI Applications](https://deepwiki.com/charmbracelet/glow/4-developer-guide#debugging-tui-applications)
    
*   [Manual Testing Checklist](https://deepwiki.com/charmbracelet/glow/4-developer-guide#manual-testing-checklist)
    
*   [Contributing Guidelines](https://deepwiki.com/charmbracelet/glow/4-developer-guide#contributing-guidelines)
