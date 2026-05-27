Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/harmonica](https://github.com/charmbracelet/harmonica "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 24 April 2025 ([6f6967](https://github.com/charmbracelet/harmonica/commits/6f6967c9)
)

*   [Overview](https://deepwiki.com/charmbracelet/harmonica/1-overview)
    
*   [Core Concepts](https://deepwiki.com/charmbracelet/harmonica/2-core-concepts)
    
*   [Spring Animation](https://deepwiki.com/charmbracelet/harmonica/2.1-spring-animation)
    
*   [Damping Types](https://deepwiki.com/charmbracelet/harmonica/2.2-damping-types)
    
*   [Projectile Motion](https://deepwiki.com/charmbracelet/harmonica/2.3-projectile-motion)
    
*   [Usage Guide](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide)
    
*   [Spring Implementation](https://deepwiki.com/charmbracelet/harmonica/3.1-spring-implementation)
    
*   [Projectile Implementation](https://deepwiki.com/charmbracelet/harmonica/3.2-projectile-implementation)
    
*   [Example Applications](https://deepwiki.com/charmbracelet/harmonica/4-example-applications)
    
*   [Terminal UI Example](https://deepwiki.com/charmbracelet/harmonica/4.1-terminal-ui-example)
    
*   [OpenGL Example](https://deepwiki.com/charmbracelet/harmonica/4.2-opengl-example)
    
*   [Particle Example](https://deepwiki.com/charmbracelet/harmonica/4.3-particle-example)
    
*   [Contributing](https://deepwiki.com/charmbracelet/harmonica/5-contributing)
    

Menu

Contributing
============

Relevant source files

*   [.github/workflows/build.yml](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/build.yml)
    
*   [.github/workflows/lint.yml](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/lint.yml)
    

This document outlines the process for contributing to the Harmonica physics-based animation library. It covers the development setup, workflow, testing procedures, and guidelines for submitting changes. For information about using the library in your applications, see [Usage Guide](https://deepwiki.com/charmbracelet/harmonica/3-usage-guide)
.

Development Setup
-----------------

### Prerequisites

To contribute to Harmonica, you'll need:

*   Go 1.16 or newer
*   Git
*   For building the OpenGL examples:
    *   On Ubuntu/Linux: libgl1-mesa-dev and xorg-dev packages
    *   Equivalent OpenGL development libraries for other platforms

### Getting Started

1.  Fork the repository on GitHub
2.  Clone your fork locally:
    
        git clone https://github.com/yourusername/harmonica.git
        cd harmonica
        
    
3.  Add the upstream repository as a remote:
    
        git remote add upstream https://github.com/charmbracelet/harmonica.git
        
    
4.  Create a new branch for your changes:
    
        git checkout -b feature/your-feature-name
        
    

Development Workflow
--------------------

Sources: [.github/workflows/build.yml1-38](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/build.yml#L1-L38)
 [.github/workflows/lint.yml1-20](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/lint.yml#L1-L20)

### Building the Project

To build the main library:

    go build -v ./...
    

To build the examples (Ubuntu/Linux):

    sudo apt-get install libgl1-mesa-dev xorg-dev
    cd examples
    go build -v ./...
    

Sources: [.github/workflows/build.yml26-34](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/build.yml#L26-L34)

### Running Tests

Run the test suite with:

    go test ./...
    

This will execute all tests in the project. Make sure all tests pass before submitting your changes.

Sources: [.github/workflows/build.yml36-37](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/build.yml#L36-L37)

Code Quality Guidelines
-----------------------

### Linting

The project uses golangci-lint for code quality checks. You can install it following the instructions on the [golangci-lint installation page](https://golangci-lint.run/usage/install/)
.

Run the linter locally with:

    golangci-lint run
    

Fix any issues before submitting your pull request.

Sources: [.github/workflows/lint.yml1-20](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/lint.yml#L1-L20)

### Coding Style

*   Follow standard Go code conventions
*   Use meaningful variable and function names
*   Add comments for complex logic
*   Write tests for new functionality
*   Keep performance in mind (Harmonica is used in animation contexts where performance matters)

Pull Request Process
--------------------

Sources: [.github/workflows/build.yml1-38](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/build.yml#L1-L38)
 [.github/workflows/lint.yml1-20](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/lint.yml#L1-L20)

1.  Ensure your code passes all tests and lint checks
2.  Create a pull request to the main repository
3.  Fill in the pull request template (if available)
4.  Wait for CI checks to complete
5.  Address any feedback from maintainers
6.  Once approved, a maintainer will merge your changes

Continuous Integration
----------------------

Harmonica uses GitHub Actions for continuous integration. Every pull request and push triggers the following workflows:

### Build Workflow

Sources: [.github/workflows/build.yml1-38](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/build.yml#L1-L38)

The build workflow:

*   Builds on multiple platforms (Ubuntu, macOS, Windows)
*   Tests with multiple Go versions (1.16 and latest)
*   Builds examples on Ubuntu with required dependencies
*   Runs all tests

### Lint Workflow

Sources: [.github/workflows/lint.yml1-20](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/lint.yml#L1-L20)

The lint workflow:

*   Runs golangci-lint on the codebase
*   For pull requests, only reports new issues
*   Does not fail the build on linting issues (--issues-exit-code=0)

Project Structure Map
---------------------

The following diagram shows the key components of the Harmonica library and how contribution activities relate to them:

Sources: [.github/workflows/build.yml1-38](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/build.yml#L1-L38)
 [.github/workflows/lint.yml1-20](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/lint.yml#L1-L20)

Common Contribution Tasks
-------------------------

| Task | Command/Action | Notes |
| --- | --- | --- |
| Build library | `go build -v ./...` | Run from project root |
| Run tests | `go test ./...` | Ensure all tests pass |
| Build examples | Install OpenGL deps, then `cd examples && go build -v ./...` | Required for UI examples |
| Run linter | `golangci-lint run` | Fix issues before submitting PR |
| Update documentation | Edit comments and README | Go docs follow standard format |
| Submit changes | Create a GitHub PR | Wait for CI and review |

Sources: [.github/workflows/build.yml1-38](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/build.yml#L1-L38)
 [.github/workflows/lint.yml1-20](https://github.com/charmbracelet/harmonica/blob/6f6967c9/.github/workflows/lint.yml#L1-L20)

Best Practices
--------------

*   Keep pull requests focused on a single feature or fix
*   Add tests for new functionality
*   Update documentation for significant changes
*   Follow existing code style patterns
*   Optimize for both readability and performance
*   Be responsive to feedback on your pull request

By following these guidelines, you'll help maintain the quality and performance of the Harmonica library while making the contribution process smooth for everyone involved.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Contributing](https://deepwiki.com/charmbracelet/harmonica/5-contributing#contributing)
    
*   [Development Setup](https://deepwiki.com/charmbracelet/harmonica/5-contributing#development-setup)
    
*   [Prerequisites](https://deepwiki.com/charmbracelet/harmonica/5-contributing#prerequisites)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/harmonica/5-contributing#getting-started)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/harmonica/5-contributing#development-workflow)
    
*   [Building the Project](https://deepwiki.com/charmbracelet/harmonica/5-contributing#building-the-project)
    
*   [Running Tests](https://deepwiki.com/charmbracelet/harmonica/5-contributing#running-tests)
    
*   [Code Quality Guidelines](https://deepwiki.com/charmbracelet/harmonica/5-contributing#code-quality-guidelines)
    
*   [Linting](https://deepwiki.com/charmbracelet/harmonica/5-contributing#linting)
    
*   [Coding Style](https://deepwiki.com/charmbracelet/harmonica/5-contributing#coding-style)
    
*   [Pull Request Process](https://deepwiki.com/charmbracelet/harmonica/5-contributing#pull-request-process)
    
*   [Continuous Integration](https://deepwiki.com/charmbracelet/harmonica/5-contributing#continuous-integration)
    
*   [Build Workflow](https://deepwiki.com/charmbracelet/harmonica/5-contributing#build-workflow)
    
*   [Lint Workflow](https://deepwiki.com/charmbracelet/harmonica/5-contributing#lint-workflow)
    
*   [Project Structure Map](https://deepwiki.com/charmbracelet/harmonica/5-contributing#project-structure-map)
    
*   [Common Contribution Tasks](https://deepwiki.com/charmbracelet/harmonica/5-contributing#common-contribution-tasks)
    
*   [Best Practices](https://deepwiki.com/charmbracelet/harmonica/5-contributing#best-practices)
