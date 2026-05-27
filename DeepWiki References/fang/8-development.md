Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/fang](https://github.com/charmbracelet/fang "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 15 March 2026 ([d89b30](https://github.com/charmbracelet/fang/commits/d89b30af)
)

*   [Overview](https://deepwiki.com/charmbracelet/fang/1-overview)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/fang/2-getting-started)
    
*   [Core Components](https://deepwiki.com/charmbracelet/fang/3-core-components)
    
*   [fang.Execute Function](https://deepwiki.com/charmbracelet/fang/3.1-fang.execute-function)
    
*   [Help Rendering System](https://deepwiki.com/charmbracelet/fang/3.2-help-rendering-system)
    
*   [Styling and Themes](https://deepwiki.com/charmbracelet/fang/3.3-styling-and-themes)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/fang/3.4-error-handling)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/fang/4-configuration-options)
    
*   [Features](https://deepwiki.com/charmbracelet/fang/5-features)
    
*   [Version Management](https://deepwiki.com/charmbracelet/fang/5.1-version-management)
    
*   [Manpage Generation](https://deepwiki.com/charmbracelet/fang/5.2-manpage-generation)
    
*   [Shell Completions](https://deepwiki.com/charmbracelet/fang/5.3-shell-completions)
    
*   [Signal Handling](https://deepwiki.com/charmbracelet/fang/5.4-signal-handling)
    
*   [Dependencies and Architecture](https://deepwiki.com/charmbracelet/fang/6-dependencies-and-architecture)
    
*   [Testing and Examples](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples)
    
*   [Golden File Testing](https://deepwiki.com/charmbracelet/fang/7.1-golden-file-testing)
    
*   [Example Application](https://deepwiki.com/charmbracelet/fang/7.2-example-application)
    
*   [Development](https://deepwiki.com/charmbracelet/fang/8-development)
    
*   [Code Quality](https://deepwiki.com/charmbracelet/fang/8.1-code-quality)
    
*   [Platform Support](https://deepwiki.com/charmbracelet/fang/8.2-platform-support)
    
*   [License](https://deepwiki.com/charmbracelet/fang/8.3-license)
    

Menu

Development
===========

Relevant source files

*   [.golangci.yml](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml)
    
*   [LICENSE.md](https://github.com/charmbracelet/fang/blob/d89b30af/LICENSE.md?plain=1)
    
*   [README.md](https://github.com/charmbracelet/fang/blob/d89b30af/README.md?plain=1)
    

This page provides information for developers who want to contribute to the fang library, including development setup, code quality standards, linting configuration, and contribution guidelines. For information about the testing methodology and test scenarios, see [Testing and Examples](https://deepwiki.com/charmbracelet/fang/6-dependencies-and-architecture)
.

Development Environment
-----------------------

The fang library is a standard Go module that requires Go 1.21 or later for development. The project follows conventional Go development practices with enhanced code quality enforcement through automated linting and formatting.

### Development Workflow

Sources: [.golangci.yml1-41](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L1-L41)

Code Quality Standards
----------------------

The project enforces strict code quality standards through `golangci-lint` with 24 enabled linters and automatic code formatting. The configuration prioritizes security, performance, and maintainability.

### Linter Configuration

The linting configuration is defined in [.golangci.yml1-41](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L1-L41)
 and includes several categories of checks:

| Category | Linters | Purpose |
| --- | --- | --- |
| Security | `gosec`, `noctx` | Identify security vulnerabilities and context usage |
| Code Quality | `revive`, `goconst`, `nestif` | Enforce coding standards and complexity limits |
| Error Handling | `nilerr`, `rowserrcheck`, `wrapcheck` | Ensure proper error handling patterns |
| Performance | `prealloc`, `bodyclose`, `sqlclosecheck` | Optimize resource usage |
| Style | `godot`, `misspell`, `whitespace` | Maintain consistent code style |

### Enabled Linters

The complete list of enabled linters from [.golangci.yml5-27](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L5-L27)
:

Sources: [.golangci.yml5-27](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L5-L27)

### Formatter Configuration

The project uses two automatic formatters configured in [.golangci.yml35-40](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L35-L40)
:

*   **gofumpt**: Stricter version of `gofmt` for consistent formatting
*   **goimports**: Automatic import management and organization

### Quality Enforcement

The linting configuration includes strict enforcement settings from [.golangci.yml32-34](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L32-L34)
:

*   `max-issues-per-linter: 0` - No tolerance for linter violations
*   `max-same-issues: 0` - Each issue must be addressed individually
*   Tests are excluded from linting with `tests: false` in [.golangci.yml3](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L3-L3)
    

Development Commands
--------------------

### Essential Commands

| Command | Purpose |
| --- | --- |
| `go mod tidy` | Clean up module dependencies |
| `golangci-lint run` | Run all quality checks |
| `go test ./...` | Execute test suite |
| `go build ./...` | Build all packages |

### Code Quality Checks

Sources: [.golangci.yml1-41](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L1-L41)

Contributing Guidelines
-----------------------

### Code Standards

All contributions must meet the quality standards enforced by the linting configuration. The project uses:

*   **Zero tolerance policy**: All linter violations must be resolved
*   **Automatic formatting**: Code is automatically formatted on commit
*   **Comprehensive testing**: Changes must include appropriate test coverage

### Exclusions and Exceptions

The linting configuration includes specific exclusions from [.golangci.yml28-31](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L28-L31)
 and [.golangci.yml39-40](https://github.com/charmbracelet/fang/blob/d89b30af/.golangci.yml#L39-L40)
:

*   Generated code is handled with "lax" rules
*   Common false positives are filtered out using preset exclusions
*   Both linters and formatters exclude generated files

License
-------

The fang library is released under the MIT License. The full license text is available in [LICENSE.md1-22](https://github.com/charmbracelet/fang/blob/d89b30af/LICENSE.md?plain=1#L1-L22)
 Key points:

*   **Copyright**: 2024-2025 Charmbracelet, Inc
*   **Permission**: Unrestricted use, modification, and distribution
*   **Conditions**: Copyright notice and license must be included in distributions
*   **Warranty**: Software provided "as is" without warranty

Sources: [LICENSE.md1-22](https://github.com/charmbracelet/fang/blob/d89b30af/LICENSE.md?plain=1#L1-L22)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development](https://deepwiki.com/charmbracelet/fang/8-development#development)
    
*   [Development Environment](https://deepwiki.com/charmbracelet/fang/8-development#development-environment)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/fang/8-development#development-workflow)
    
*   [Code Quality Standards](https://deepwiki.com/charmbracelet/fang/8-development#code-quality-standards)
    
*   [Linter Configuration](https://deepwiki.com/charmbracelet/fang/8-development#linter-configuration)
    
*   [Enabled Linters](https://deepwiki.com/charmbracelet/fang/8-development#enabled-linters)
    
*   [Formatter Configuration](https://deepwiki.com/charmbracelet/fang/8-development#formatter-configuration)
    
*   [Quality Enforcement](https://deepwiki.com/charmbracelet/fang/8-development#quality-enforcement)
    
*   [Development Commands](https://deepwiki.com/charmbracelet/fang/8-development#development-commands)
    
*   [Essential Commands](https://deepwiki.com/charmbracelet/fang/8-development#essential-commands)
    
*   [Code Quality Checks](https://deepwiki.com/charmbracelet/fang/8-development#code-quality-checks)
    
*   [Contributing Guidelines](https://deepwiki.com/charmbracelet/fang/8-development#contributing-guidelines)
    
*   [Code Standards](https://deepwiki.com/charmbracelet/fang/8-development#code-standards)
    
*   [Exclusions and Exceptions](https://deepwiki.com/charmbracelet/fang/8-development#exclusions-and-exceptions)
    
*   [License](https://deepwiki.com/charmbracelet/fang/8-development#license)
