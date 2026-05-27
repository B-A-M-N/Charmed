Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/soft-serve](https://github.com/charmbracelet/soft-serve "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 18 April 2025 ([767bdd](https://github.com/charmbracelet/soft-serve/commits/767bdd5a)
)

*   [Overview](https://deepwiki.com/charmbracelet/soft-serve/1-overview)
    
*   [What is Soft Serve?](https://deepwiki.com/charmbracelet/soft-serve/1.1-what-is-soft-serve)
    
*   [Key Features](https://deepwiki.com/charmbracelet/soft-serve/1.2-key-features)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/soft-serve/1.3-system-architecture)
    
*   [Installation and Configuration](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration)
    
*   [Installation Methods](https://deepwiki.com/charmbracelet/soft-serve/2.1-installation-methods)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/soft-serve/2.2-configuration-options)
    
*   [System Service Setup](https://deepwiki.com/charmbracelet/soft-serve/2.3-system-service-setup)
    
*   [Core Components](https://deepwiki.com/charmbracelet/soft-serve/3-core-components)
    
*   [CLI Tool Structure](https://deepwiki.com/charmbracelet/soft-serve/3.1-cli-tool-structure)
    
*   [Server Components](https://deepwiki.com/charmbracelet/soft-serve/3.2-server-components)
    
*   [SSH Server](https://deepwiki.com/charmbracelet/soft-serve/3.2.1-ssh-server)
    
*   [Git Daemon](https://deepwiki.com/charmbracelet/soft-serve/3.2.2-git-daemon)
    
*   [HTTP Server](https://deepwiki.com/charmbracelet/soft-serve/3.2.3-http-server)
    
*   [Repository Management](https://deepwiki.com/charmbracelet/soft-serve/3.3-repository-management)
    
*   [Git Hooks and Webhooks](https://deepwiki.com/charmbracelet/soft-serve/3.4-git-hooks-and-webhooks)
    
*   [User Interfaces](https://deepwiki.com/charmbracelet/soft-serve/4-user-interfaces)
    
*   [Terminal UI Architecture](https://deepwiki.com/charmbracelet/soft-serve/4.1-terminal-ui-architecture)
    
*   [UI Components](https://deepwiki.com/charmbracelet/soft-serve/4.2-ui-components)
    
*   [Command Line Interface](https://deepwiki.com/charmbracelet/soft-serve/4.3-command-line-interface)
    
*   [Developer Guide](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide)
    
*   [Project Dependencies](https://deepwiki.com/charmbracelet/soft-serve/5.1-project-dependencies)
    
*   [Build System](https://deepwiki.com/charmbracelet/soft-serve/5.2-build-system)
    
*   [Testing Framework](https://deepwiki.com/charmbracelet/soft-serve/5.3-testing-framework)
    
*   [Code Quality Tools](https://deepwiki.com/charmbracelet/soft-serve/5.4-code-quality-tools)
    
*   [Release Process](https://deepwiki.com/charmbracelet/soft-serve/5.5-release-process)
    

Menu

Developer Guide
===============

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/build.yml)
    
*   [.github/workflows/coverage.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/coverage.yml)
    
*   [.github/workflows/lint-sync.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/lint-sync.yml)
    
*   [.github/workflows/lint.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/lint.yml)
    
*   [.golangci.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.golangci.yml)
    
*   [cmd/cmd.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/cmd.go)
    
*   [pkg/access/context\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/access/context_test.go)
    
*   [pkg/backend/backend.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/backend/backend.go)
    
*   [pkg/config/context\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/context_test.go)
    
*   [pkg/config/file\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/file_test.go)
    
*   [testscript/script\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/script_test.go)
    
*   [testscript/testdata/http.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/http.txtar)
    
*   [testscript/testdata/mirror.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/mirror.txtar)
    
*   [testscript/testdata/repo-create.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/repo-create.txtar)
    
*   [testscript/testdata/repo-import.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/repo-import.txtar)
    
*   [testscript/testdata/ssh-lfs.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/ssh-lfs.txtar)
    

This guide provides essential information for developers who want to contribute to or extend Soft Serve. It covers the project's dependencies, build system, testing framework, code quality tools, and release process. For information on installing and using Soft Serve, see [Installation and Configuration](https://deepwiki.com/charmbracelet/soft-serve/2-installation-and-configuration)
.

Project Dependencies
--------------------

Soft Serve relies on several key dependencies that developers should be familiar with:

| Dependency | Purpose |
| --- | --- |
| github.com/charmbracelet/keygen | SSH key generation and management |
| github.com/charmbracelet/log | Structured logging system |
| github.com/spf13/cobra | Command-line interface framework |
| golang.org/x/crypto/ssh | SSH server functionality |
| github.com/rogpeppe/go-internal/testscript | Integration testing infrastructure |
| Database drivers | Support for SQLite and PostgreSQL |

Sources: [testscript/script\_test.go3-30](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/script_test.go#L3-L30)

Development Architecture
------------------------

The following diagram illustrates the key components of the Soft Serve development architecture:

Sources: [.github/workflows/build.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/build.yml)
 [.github/workflows/lint.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/lint.yml)
 [.github/workflows/coverage.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/coverage.yml)

Build System
------------

### Local Development Build

For local development, you can build Soft Serve using Go's standard build tools:

For development builds with race detection and code coverage:

This is implemented in the test framework as seen in [testscript/script\_test.go37-44](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/script_test.go#L37-L44)
:

### CI/CD Pipeline

The CI/CD pipeline uses GitHub Actions workflows for automated building and testing:

*   The main build workflow ([.github/workflows/build.yml10-16](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/build.yml#L10-L16)
    ) handles multi-platform builds
*   The snapshot workflow ([.github/workflows/build.yml13-16](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/build.yml#L13-L16)
    ) creates development snapshots
*   The PostgreSQL test workflow ([.github/workflows/build.yml18-46](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/build.yml#L18-L46)
    ) validates database functionality

Sources: [.github/workflows/build.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/build.yml)

Testing Framework
-----------------

Soft Serve employs a comprehensive testing strategy that combines unit tests and integration tests.

Sources: [testscript/script\_test.go70-190](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/script_test.go#L70-L190)
 [.github/workflows/coverage.yml27-58](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/coverage.yml#L27-L58)

### Unit Tests

Standard Go unit tests are used throughout the codebase to test individual components:

Examples of unit tests can be found in [pkg/config/context\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/context_test.go)
 and [pkg/access/context\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/access/context_test.go)

### Integration Tests with TestScript

The integration tests use the `testscript` package to create comprehensive integration tests. The main setup is in [testscript/script\_test.go70-190](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/script_test.go#L70-L190)
 and works as follows:

1.  TestMain builds a temporary Soft Serve binary with coverage instrumentation
2.  The tests set up random ports for different services (SSH, Git, HTTP, Stats)
3.  Temporary directories and SSH keys are created for test data
4.  TXTAR-format test scripts run commands against the running server

Key test commands include:

*   `soft` and `usoft` - Test the CLI with admin and regular user permissions respectively
*   `git` and `ugit` - Test Git operations with different user permissions
*   `curl` - Test HTTP endpoints
*   `ui` and `uui` - Test the Terminal UI with different user permissions

Sources: [testscript/script\_test.go192-215](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/script_test.go#L192-L215)
 [testscript/script\_test.go309-332](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/script_test.go#L309-L332)
 [testscript/script\_test.go397-482](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/script_test.go#L397-L482)

### TXTAR Test Format

The TXTAR format allows defining test scripts with commands and expected outputs. For example, from [testscript/testdata/repo-create.txtar8-19](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/repo-create.txtar#L8-L19)
:

    # start soft serve
    exec soft serve &
    # wait for SSH server to start
    ensureserverrunning SSH_PORT
    
    # create a repo
    soft repo create repo1 -d 'description' -H -p -n 'repo11'
    stderr 'Created repository repo1.*'
    stdout ssh://localhost:$SSH_PORT/repo1.git
    

Other test files include [testscript/testdata/http.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/http.txtar)
 [testscript/testdata/ssh-lfs.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/ssh-lfs.txtar)
 and [testscript/testdata/mirror.txtar](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/testdata/mirror.txtar)

### Test Coverage

Test coverage is tracked using Go's coverage tools and reported to Codecov. The workflow in [.github/workflows/coverage.yml35-57](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/coverage.yml#L35-L57)
 collects coverage from both unit and integration tests:

Code Quality Tools
------------------

### Linting

Soft Serve uses golangci-lint for static code analysis with configuration in [.golangci.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.golangci.yml)

Key linters enabled include:

| Linter | Purpose |
| --- | --- |
| bodyclose | Checks for unclosed HTTP response bodies |
| exhaustive | Checks exhaustiveness of enum switch statements |
| goconst | Finds repeated strings that could be constants |
| godot | Checks if comments end with a period |
| gosec | Inspects source code for security problems |
| misspell | Finds commonly misspelled English words |

The linting workflow is defined in [.github/workflows/lint.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/lint.yml)
 and is synchronized weekly via [.github/workflows/lint-sync.yml](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/lint-sync.yml)

### Formatting

Code formatting is enforced with:

*   gofumpt - A stricter gofmt
*   goimports - Formats imports and adds missing ones

Context System
--------------

Soft Serve makes extensive use of Go's context package for dependency injection:

Sources: [cmd/cmd.go20-42](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/cmd.go#L20-L42)
 [pkg/backend/backend.go14-42](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/backend/backend.go#L14-L42)
 [pkg/config/context\_test.go](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/config/context_test.go)

The backend initialization in [cmd/cmd.go20-42](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/cmd.go#L20-L42)
 shows how these contexts are wired together:

Release Process
---------------

Soft Serve uses GoReleaser for creating and publishing releases. The release workflow is triggered when Git tags are created.

The snapshot build workflow defined in [.github/workflows/build.yml13-16](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/build.yml#L13-L16)
 is used for testing pre-release builds.

Contributing Guidelines
-----------------------

### Setting Up Development Environment

1.  Clone the repository:
    
2.  Install dependencies:
    
3.  Build the binary:
    
4.  Run tests:
    

### Testing with Different Databases

Soft Serve supports both SQLite (default) and PostgreSQL. To test with PostgreSQL:

As shown in [.github/workflows/build.yml45-46](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/.github/workflows/build.yml#L45-L46)
 and [testscript/script\_test.go544-618](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/testscript/script_test.go#L544-L618)

### Pull Request Process

1.  Fork the repository
2.  Create a feature branch
3.  Make your changes
4.  Ensure tests pass and linting is clean
5.  Submit a pull request

Pull requests should include test coverage for new functionality and follow the project's code style.

Extending Soft Serve
--------------------

### Core Components Architecture

Sources: [pkg/backend/backend.go14-42](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/backend/backend.go#L14-L42)
 [cmd/cmd.go58-71](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/cmd.go#L58-L71)

### Adding New Commands

New commands can be added by extending the CLI using Cobra. The command system is initialized in the `cmd` package.

### Adding New Backend Features

The backend is the core of Soft Serve, defined in [pkg/backend/backend.go14-42](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/pkg/backend/backend.go#L14-L42)
 New features can be added by extending this package:

### Git Hooks

Git hooks can be initialized using the function in [cmd/cmd.go58-71](https://github.com/charmbracelet/soft-serve/blob/767bdd5a/cmd/cmd.go#L58-L71)
:

By understanding these core components and patterns, you can effectively contribute to and extend the Soft Serve codebase.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Developer Guide](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#developer-guide)
    
*   [Project Dependencies](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#project-dependencies)
    
*   [Development Architecture](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#development-architecture)
    
*   [Build System](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#build-system)
    
*   [Local Development Build](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#local-development-build)
    
*   [CI/CD Pipeline](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#cicd-pipeline)
    
*   [Testing Framework](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#testing-framework)
    
*   [Unit Tests](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#unit-tests)
    
*   [Integration Tests with TestScript](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#integration-tests-with-testscript)
    
*   [TXTAR Test Format](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#txtar-test-format)
    
*   [Test Coverage](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#test-coverage)
    
*   [Code Quality Tools](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#code-quality-tools)
    
*   [Linting](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#linting)
    
*   [Formatting](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#formatting)
    
*   [Context System](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#context-system)
    
*   [Release Process](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#release-process)
    
*   [Contributing Guidelines](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#contributing-guidelines)
    
*   [Setting Up Development Environment](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#setting-up-development-environment)
    
*   [Testing with Different Databases](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#testing-with-different-databases)
    
*   [Pull Request Process](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#pull-request-process)
    
*   [Extending Soft Serve](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#extending-soft-serve)
    
*   [Core Components Architecture](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#core-components-architecture)
    
*   [Adding New Commands](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#adding-new-commands)
    
*   [Adding New Backend Features](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#adding-new-backend-features)
    
*   [Git Hooks](https://deepwiki.com/charmbracelet/soft-serve/5-developer-guide#git-hooks)
