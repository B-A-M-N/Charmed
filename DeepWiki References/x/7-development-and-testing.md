Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/x](https://github.com/charmbracelet/x "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 7 February 2026 ([764291](https://github.com/charmbracelet/x/commits/7642919e)
)

*   [Overview](https://deepwiki.com/charmbracelet/x/1-overview)
    
*   [Core Terminal Rendering](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering)
    
*   [cellbuf Package](https://deepwiki.com/charmbracelet/x/2.1-cellbuf-package)
    
*   [Virtual Terminal Emulator (vt)](https://deepwiki.com/charmbracelet/x/2.2-virtual-terminal-emulator-(vt))
    
*   [Input Processing System](https://deepwiki.com/charmbracelet/x/3-input-processing-system)
    
*   [Input Package](https://deepwiki.com/charmbracelet/x/3.1-input-package)
    
*   [Windows Console Input](https://deepwiki.com/charmbracelet/x/3.2-windows-console-input)
    
*   [ANSI Processing and Control](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control)
    
*   [Parser and Text Manipulation](https://deepwiki.com/charmbracelet/x/4.1-parser-and-text-manipulation)
    
*   [Terminal Control Sequences](https://deepwiki.com/charmbracelet/x/4.2-terminal-control-sequences)
    
*   [Styling and Colors](https://deepwiki.com/charmbracelet/x/4.3-styling-and-colors)
    
*   [Platform Abstraction Layer](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer)
    
*   [Pseudo-Terminal Interface (xpty)](https://deepwiki.com/charmbracelet/x/5.1-pseudo-terminal-interface-(xpty))
    
*   [Windows ConPTY](https://deepwiki.com/charmbracelet/x/5.2-windows-conpty)
    
*   [Terminal I/O (termios and wcwidth)](https://deepwiki.com/charmbracelet/x/5.3-terminal-io-(termios-and-wcwidth))
    
*   [Utility Packages](https://deepwiki.com/charmbracelet/x/6-utility-packages)
    
*   [SSH Keys and Colors](https://deepwiki.com/charmbracelet/x/6.1-ssh-keys-and-colors)
    
*   [Experimental Packages](https://deepwiki.com/charmbracelet/x/6.2-experimental-packages)
    
*   [Development and Testing](https://deepwiki.com/charmbracelet/x/7-development-and-testing)
    
*   [Examples and Integration](https://deepwiki.com/charmbracelet/x/7.1-examples-and-integration)
    
*   [Testing Framework (teatest)](https://deepwiki.com/charmbracelet/x/7.2-testing-framework-(teatest))
    
*   [Repository Structure and CI/CD](https://deepwiki.com/charmbracelet/x/7.3-repository-structure-and-cicd)
    

Menu

Development and Testing
=======================

Relevant source files

*   [.github/dependabot.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/dependabot.yml)
    
*   [.github/workflows/gitignore.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/gitignore.yml)
    
*   [.github/workflows/vttest.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/vttest.yml)
    
*   [.gitignore](https://github.com/charmbracelet/x/blob/7642919e/.gitignore)
    
*   [LICENSE](https://github.com/charmbracelet/x/blob/7642919e/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1)
    
*   [examples/cellbuf/main.go](https://github.com/charmbracelet/x/blob/7642919e/examples/cellbuf/main.go)
    
*   [examples/charmtone/guide.go](https://github.com/charmbracelet/x/blob/7642919e/examples/charmtone/guide.go)
    
*   [examples/go.mod](https://github.com/charmbracelet/x/blob/7642919e/examples/go.mod)
    
*   [examples/go.sum](https://github.com/charmbracelet/x/blob/7642919e/examples/go.sum)
    
*   [examples/layout/main.go](https://github.com/charmbracelet/x/blob/7642919e/examples/layout/main.go)
    
*   [examples/parserlog/main.go](https://github.com/charmbracelet/x/blob/7642919e/examples/parserlog/main.go)
    
*   [examples/parserlog2/main.go](https://github.com/charmbracelet/x/blob/7642919e/examples/parserlog2/main.go)
    
*   [examples/toner/main.go](https://github.com/charmbracelet/x/blob/7642919e/examples/toner/main.go)
    
*   [exp/golden/testdata/TestTypes/SliceOfBytes.golden](https://github.com/charmbracelet/x/blob/7642919e/exp/golden/testdata/TestTypes/SliceOfBytes.golden)
    
*   [exp/golden/testdata/TestTypes/String.golden](https://github.com/charmbracelet/x/blob/7642919e/exp/golden/testdata/TestTypes/String.golden)
    
*   [exp/teatest/app\_test.go](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/app_test.go)
    
*   [exp/teatest/go.mod](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/go.mod)
    
*   [exp/teatest/go.sum](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/go.sum)
    
*   [exp/teatest/send\_test.go](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/send_test.go)
    
*   [exp/teatest/teatest.go](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/teatest.go)
    
*   [exp/teatest/teatest\_test.go](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/teatest_test.go)
    
*   [exp/teatest/testdata/TestApp.golden](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/testdata/TestApp.golden)
    
*   [exp/teatest/testdata/TestAppSendToOtherProgram.golden](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/testdata/TestAppSendToOtherProgram.golden)
    
*   [exp/teatest/testdata/TestRequireEqualOutputUpdate.golden](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/testdata/TestRequireEqualOutputUpdate.golden)
    
*   [scripts/dependabot](https://github.com/charmbracelet/x/blob/7642919e/scripts/dependabot)
    

The charmbracelet/x repository provides comprehensive development infrastructure including example applications, specialized testing frameworks, and automated CI/CD systems. This monorepo manages 33+ independent Go modules with coordinated dependency management and cross-platform testing.

This page provides an overview of the development and testing ecosystem. For detailed information, see:

*   **Examples and Integration** ([page 7.1](https://deepwiki.com/charmbracelet/x/7.1-examples-and-integration)
    ) - Example applications demonstrating package usage
*   **Testing Framework (teatest)** ([page 7.2](https://deepwiki.com/charmbracelet/x/7.2-testing-framework-(teatest))
    ) - Framework for testing Bubble Tea applications
*   **Repository Structure and CI/CD** ([page 7.3](https://deepwiki.com/charmbracelet/x/7.3-repository-structure-and-cicd)
    ) - Monorepo organization and automation </old\_str>

<old\_str>

Testing Framework
-----------------

The `exp/teatest` package provides specialized testing utilities for Bubble Tea applications, enabling programmatic testing of interactive terminal programs with output verification.

**teatest Testing Flow**

The testing framework provides three core capabilities:

| Capability | Function | Usage |
| --- | --- | --- |
| **Test Model Creation** | `teatest.NewTestModel(tb, model, opts...)` | Creates testable Bubble Tea program with `tea.WithInput()`, `tea.WithOutput()`, `tea.WithoutSignals()` |
| **Programmatic Interaction** | `tm.Send()`, `tm.Type()` | Sends messages and simulates keyboard input to test program |
| **Output Verification** | `teatest.WaitFor()`, `tm.FinalOutput()`, `teatest.RequireEqualOutput()` | Validates output with condition checking and golden file comparison |

The framework includes synchronization utilities like `WaitFor()` with configurable duration and check intervals for testing asynchronous behavior.

See [Testing Framework (teatest)](https://deepwiki.com/charmbracelet/x/7.2-testing-framework-(teatest))
 for complete API documentation and usage patterns.

Sources: [exp/teatest/teatest.go1-306](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/teatest.go#L1-L306)
 [exp/teatest/app\_test.go15-52](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/app_test.go#L15-L52)
 [exp/teatest/send\_test.go13-50](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/send_test.go#L13-L50)
 [exp/teatest/testdata/TestApp.golden1-3](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/testdata/TestApp.golden#L1-L3)
 </old\_str> <new\_str>

Development Infrastructure Overview
-----------------------------------

**Development Infrastructure Components**

The development infrastructure consists of three main components:

| Component | Purpose | Key Entities |
| --- | --- | --- |
| **Examples** | Integration demonstrations and manual testing | `examples/cellbuf`, `examples/layout`, `examples/parserlog` |
| **Testing Framework** | Automated testing for Bubble Tea applications | `exp/teatest.TestModel`, `exp/golden.RequireEqual()` |
| **CI/CD System** | Automated builds, tests, and dependency updates | `.github/dependabot.yml`, `.github/workflows/`, `scripts/dependabot` |

Sources: [examples/go.mod1-53](https://github.com/charmbracelet/x/blob/7642919e/examples/go.mod#L1-L53)
 [exp/teatest/go.mod1-33](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/go.mod#L1-L33)
 [.github/dependabot.yml1-566](https://github.com/charmbracelet/x/blob/7642919e/.github/dependabot.yml#L1-L566)
 [scripts/dependabot1-44](https://github.com/charmbracelet/x/blob/7642919e/scripts/dependabot#L1-L44)

Example Applications
--------------------

The `examples/` directory contains demonstration applications that serve as both usage documentation and integration tests. Each example is a standalone Go program demonstrating real-world usage patterns.

**Example Application Architecture**

Key examples include:

*   **`examples/cellbuf/main.go`** - Basic screen rendering with input handling using `cellbuf.NewScreen()`, `input.NewReader()`, and a render loop
*   **`examples/layout/main.go`** - Complex layouts with Lip Gloss integration demonstrating tabs, dialogs, and color grids
*   **`examples/parserlog/main.go`** - ANSI parser demonstration using `ansi.NewParser()` with callback handlers
*   **`examples/parserlog2/main.go`** - Alternative parser approach using `ansi.DecodeSequence()`
*   **`examples/toner/main.go`** - Color toning demonstration using `toner.Writer`
*   **`examples/charmtone/main.go`** - Charm color palette guide using `charmtone` package

See [Examples and Integration](https://deepwiki.com/charmbracelet/x/7.1-examples-and-integration)
 for detailed documentation of example applications.

Sources: [examples/cellbuf/main.go15-118](https://github.com/charmbracelet/x/blob/7642919e/examples/cellbuf/main.go#L15-L118)
 [examples/layout/main.go52-454](https://github.com/charmbracelet/x/blob/7642919e/examples/layout/main.go#L52-L454)
 [examples/parserlog/main.go14-75](https://github.com/charmbracelet/x/blob/7642919e/examples/parserlog/main.go#L14-L75)
 [examples/parserlog2/main.go14-137](https://github.com/charmbracelet/x/blob/7642919e/examples/parserlog2/main.go#L14-L137)
 [examples/toner/main.go12-22](https://github.com/charmbracelet/x/blob/7642919e/examples/toner/main.go#L12-L22)
 [examples/charmtone/guide.go15-199](https://github.com/charmbracelet/x/blob/7642919e/examples/charmtone/guide.go#L15-L199)

Examples and Integration System
-------------------------------

The repository includes an extensive examples system that demonstrates integration patterns and usage across multiple packages. These examples serve both as documentation and integration tests for the various terminal components.

### Core Integration Examples

The examples demonstrate complete integration patterns between multiple charmbracelet/x modules:

#### Cellbuf Screen Management Example

The cellbuf example shows comprehensive terminal screen management with input handling:

*   **Screen Initialization**: Uses `cellbuf.NewScreen()` with configurable options including terminal type detection, alternate screen mode, and relative cursor positioning
*   **Input Integration**: Integrates `input.NewReader()` for cross-platform event handling including keyboard, mouse, and resize events
*   **Render Loop**: Implements a complete render loop with `scr.Fill()`, `scrw.PrintCropAt()`, `scr.Render()`, and `scr.Flush()`
*   **Event Processing**: Handles `input.WindowSizeEvent`, `input.MouseClickEvent`, and `input.KeyPressEvent` types

#### Layout and Styling Integration

The layout example demonstrates advanced styling with Lip Gloss integration:

*   **Color Detection**: Uses `lipgloss.HasDarkBackground()` and `lipgloss.LightDark()` for adaptive theming
*   **Complex Layouts**: Implements tabs, dialogs, color grids, and status bars using `lipgloss.JoinHorizontal()` and `lipgloss.JoinVertical()`
*   **Screen Writer Integration**: Uses `cellbuf.NewScreenWriter()` and `scrw.SetContentRect()` for precise content placement
*   **Interactive Elements**: Provides moveable dialog boxes with mouse and keyboard interaction

#### ANSI Parser Demonstrations

Two parser examples show different approaches to ANSI sequence processing:

*   **Handler-Based Parsing**: Uses `ansi.NewParser()` with `ansi.Handler` callbacks for `Print`, `Execute`, `HandleCsi`, `HandleEsc`, `HandleDcs`, `HandleOsc` functions
*   **Sequence Decoding**: Demonstrates `ansi.DecodeSequence()` for direct sequence processing with state management

Sources: [examples/cellbuf/main.go1-126](https://github.com/charmbracelet/x/blob/7642919e/examples/cellbuf/main.go#L1-L126)
 [examples/layout/main.go1-525](https://github.com/charmbracelet/x/blob/7642919e/examples/layout/main.go#L1-L525)
 [examples/parserlog/main.go1-100](https://github.com/charmbracelet/x/blob/7642919e/examples/parserlog/main.go#L1-L100)
 [examples/parserlog2/main.go1-137](https://github.com/charmbracelet/x/blob/7642919e/examples/parserlog2/main.go#L1-L137)
 [examples/go.mod1-36](https://github.com/charmbracelet/x/blob/7642919e/examples/go.mod#L1-L36)

Repository Structure and CI/CD
------------------------------

The repository is organized as a Go monorepo with 33+ independent modules managed through automated dependency management and continuous integration.

**Monorepo Module Organization**

### Automated Dependency Management

The repository uses a custom script to manage Dependabot configuration across all modules:

**`scripts/dependabot` Configuration Generator**

    #!/usr/bin/env bash
    # Generates .github/dependabot.yml by scanning for go.mod files
    find . -type f -name go.mod | while read -r mod; do
      echo '  - package-ecosystem: "gomod"'
      echo '    directory: "/'"$(dirname "$mod")"'"'
      # ... schedule and configuration ...
    done
    

The generated `.github/dependabot.yml` includes:

*   **GitHub Actions Updates**: Weekly dependency updates for workflow actions
*   **Per-Module Go Updates**: Separate configuration for each of 33+ modules
*   **Automated Schedule**: Weekly updates on Monday at 05:00 America/New\_York
*   **Grouped Updates**: All dependencies updated together with `groups.all.patterns: ["*"]`

### CI/CD Workflow Structure

Each module has dedicated GitHub Actions workflows with consistent structure:

| Workflow Component | Configuration |
| --- | --- |
| **Build Matrix** | Tests on `ubuntu-latest`, `macos-latest`, `windows-latest` |
| **Test Execution** | `go build -v ./...` and `go test -race -v ./...` |
| **Coverage** | Uploads coverage reports to Codecov |
| **Dependabot Integration** | Auto-approves and merges dependency PRs |
| **Linting** | Uses `charmbracelet/meta/.github/workflows/lint.yml@main` |

Example workflow structure from `vttest.yml`:

*   Build job with platform matrix testing
*   Dependabot job with auto-approval for dependency PRs
*   Lint job using shared Charmbracelet configuration
*   Coverage job with race detection and Codecov upload

See [Repository Structure and CI/CD](https://deepwiki.com/charmbracelet/x/7.3-repository-structure-and-cicd)
 for complete documentation of monorepo organization and automation.

Sources: [.github/dependabot.yml1-566](https://github.com/charmbracelet/x/blob/7642919e/.github/dependabot.yml#L1-L566)
 [scripts/dependabot1-44](https://github.com/charmbracelet/x/blob/7642919e/scripts/dependabot#L1-L44)
 [.github/workflows/vttest.yml1-86](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/vttest.yml#L1-L86)
 [.github/workflows/gitignore.yml1-64](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/gitignore.yml#L1-L64)

Workspace Organization and Build System
---------------------------------------

The repository uses Go workspaces to manage multiple related modules with coordinated development and testing across the entire toolkit.

### Module Categories and Dependencies

The workspace organizes modules into logical categories:

#### Core Terminal Modules

*   **ansi**: ANSI escape sequence processing and generation
*   **cellbuf**: Screen buffer management and rendering
*   **input**: Cross-platform input event handling
*   **term**: Terminal control and capability detection
*   **vt**: Virtual terminal emulation
*   **windows**: Windows-specific console API bindings

#### Platform Abstraction Modules

*   **conpty**: Windows ConPTY (Console Pseudo-Terminal) support
*   **termios**: Unix terminal I/O control interface
*   **xpty**: Cross-platform pseudo-terminal interface

#### Utility and Support Modules

*   **colors**: Predefined color constants and utilities
*   **sshkey**: SSH key parsing and handling
*   **wcwidth**: Unicode character width calculation
*   **json**: JSON processing utilities

#### Experimental Modules

*   **exp/teatest**: Testing framework for Bubble Tea applications
*   **exp/golden**: Golden file testing utilities
*   **exp/slice**: Generic slice operation utilities
*   **exp/open**: Cross-platform file/URL opening
*   **exp/toner**: Terminal output enhancement
*   **exp/charmtone**: Color tone manipulation

### Dependency Management

The workspace uses centralized dependency management:

*   **`go.work.sum`**: Contains checksums for all workspace dependencies
*   **Cross-Module References**: Modules reference each other using workspace-relative paths
*   **External Dependencies**: Common dependencies like `golang.org/x/sys` are shared across modules
*   **Version Coordination**: The workspace ensures compatible versions across all modules

### Build and Development Workflow

The workspace enables coordinated development across modules:

*   **Unified Testing**: Tests can be run across all modules simultaneously
*   **Cross-Module Changes**: Changes in one module are immediately visible to dependent modules
*   **Dependency Updates**: Dependency updates can be coordinated across the entire toolkit
*   **Release Management**: Releases can be coordinated to ensure compatibility

Sources: [go.work.sum1-34](https://github.com/charmbracelet/x/blob/7642919e/go.work.sum#L1-L34)
 [examples/go.mod1-36](https://github.com/charmbracelet/x/blob/7642919e/examples/go.mod#L1-L36)
 [exp/teatest/go.mod1-29](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/go.mod#L1-L29)
 [mosaic/go.mod1-15](https://github.com/charmbracelet/x/blob/7642919e/mosaic/go.mod#L1-L15)
 [json/go.mod1-4](https://github.com/charmbracelet/x/blob/7642919e/json/go.mod#L1-L4)
 [exp/open/go.mod1-4](https://github.com/charmbracelet/x/blob/7642919e/exp/open/go.mod#L1-L4)

Quality Assurance and Testing Patterns
--------------------------------------

The repository implements comprehensive quality assurance through multiple testing strategies and development tools.

### Testing Infrastructure Components

#### Golden File Testing System

The repository uses golden file testing for output regression testing:

*   **`exp/teatest/testdata/*.golden`**: Contains expected output files for teatest framework tests
*   **Update Mechanism**: Golden files can be updated using test flags for maintenance
*   **Diff Integration**: Uses system `diff` tool for comparing actual vs expected output
*   **Version Control**: Golden files are committed to ensure consistent test behavior

#### Cross-Platform Testing

Testing is designed to work across different platforms:

*   **Build Constraints**: Uses `//go:build` directives for platform-specific code
*   **Platform-Specific Examples**: `examples/faketty/main.go` includes `//go:build !windows` constraint
*   **Conditional Logic**: Examples include runtime checks like `runtime.GOOS != "windows"`

#### Development and Debugging Tools

Several utilities assist with development and debugging:

*   **FakeTTY Utility**: `examples/faketty/main.go` creates pseudo-terminals for isolated testing
*   **Parser Log Tools**: Multiple parser logging examples for debugging ANSI sequence processing
*   **Interactive Examples**: Full interactive applications for manual testing and demonstration

### Test Coverage Patterns

The testing approach covers multiple layers:

#### Unit Test Coverage

*   **Framework Testing**: `exp/teatest/teatest_test.go` tests the testing framework itself
*   **Error Handling**: Tests include error conditions like `TestWaitForErrorReader` and `TestWaitForTimeout`
*   **Timeout Behavior**: Tests verify timeout handling with `TestWaitFinishedWithTimeoutFn`

#### Integration Test Coverage

*   **Application Testing**: `exp/teatest/app_test.go` demonstrates full application testing
*   **Inter-Program Communication**: `exp/teatest/send_test.go` tests message passing between programs
*   **Real-World Scenarios**: Tests simulate actual user interactions and program lifecycles

#### Example Application Coverage

*   **Comprehensive Demos**: Examples demonstrate real-world usage patterns
*   **Error Handling**: Examples include proper error handling and cleanup
*   **Resource Management**: Examples show proper resource cleanup with `defer` statements

Sources: [exp/teatest/testdata/TestApp.golden1-5](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/testdata/TestApp.golden#L1-L5)
 [exp/teatest/testdata/TestAppSendToOtherProgram.golden1-9](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/testdata/TestAppSendToOtherProgram.golden#L1-L9)
 [exp/teatest/testdata/TestRequireEqualOutputUpdate.golden1](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/testdata/TestRequireEqualOutputUpdate.golden#L1-L1)
 [examples/faketty/main.go1-91](https://github.com/charmbracelet/x/blob/7642919e/examples/faketty/main.go#L1-L91)
 [exp/teatest/teatest\_test.go1-53](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/teatest_test.go#L1-L53)
 [exp/teatest/app\_test.go1-131](https://github.com/charmbracelet/x/blob/7642919e/exp/teatest/app_test.go#L1-L131)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development and Testing](https://deepwiki.com/charmbracelet/x/7-development-and-testing#development-and-testing)
    
*   [Development Infrastructure Overview](https://deepwiki.com/charmbracelet/x/7-development-and-testing#development-infrastructure-overview)
    
*   [Example Applications](https://deepwiki.com/charmbracelet/x/7-development-and-testing#example-applications)
    
*   [Examples and Integration System](https://deepwiki.com/charmbracelet/x/7-development-and-testing#examples-and-integration-system)
    
*   [Core Integration Examples](https://deepwiki.com/charmbracelet/x/7-development-and-testing#core-integration-examples)
    
*   [Cellbuf Screen Management Example](https://deepwiki.com/charmbracelet/x/7-development-and-testing#cellbuf-screen-management-example)
    
*   [Layout and Styling Integration](https://deepwiki.com/charmbracelet/x/7-development-and-testing#layout-and-styling-integration)
    
*   [ANSI Parser Demonstrations](https://deepwiki.com/charmbracelet/x/7-development-and-testing#ansi-parser-demonstrations)
    
*   [Repository Structure and CI/CD](https://deepwiki.com/charmbracelet/x/7-development-and-testing#repository-structure-and-cicd)
    
*   [Automated Dependency Management](https://deepwiki.com/charmbracelet/x/7-development-and-testing#automated-dependency-management)
    
*   [CI/CD Workflow Structure](https://deepwiki.com/charmbracelet/x/7-development-and-testing#cicd-workflow-structure)
    
*   [Workspace Organization and Build System](https://deepwiki.com/charmbracelet/x/7-development-and-testing#workspace-organization-and-build-system)
    
*   [Module Categories and Dependencies](https://deepwiki.com/charmbracelet/x/7-development-and-testing#module-categories-and-dependencies)
    
*   [Core Terminal Modules](https://deepwiki.com/charmbracelet/x/7-development-and-testing#core-terminal-modules)
    
*   [Platform Abstraction Modules](https://deepwiki.com/charmbracelet/x/7-development-and-testing#platform-abstraction-modules)
    
*   [Utility and Support Modules](https://deepwiki.com/charmbracelet/x/7-development-and-testing#utility-and-support-modules)
    
*   [Experimental Modules](https://deepwiki.com/charmbracelet/x/7-development-and-testing#experimental-modules)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/x/7-development-and-testing#dependency-management)
    
*   [Build and Development Workflow](https://deepwiki.com/charmbracelet/x/7-development-and-testing#build-and-development-workflow)
    
*   [Quality Assurance and Testing Patterns](https://deepwiki.com/charmbracelet/x/7-development-and-testing#quality-assurance-and-testing-patterns)
    
*   [Testing Infrastructure Components](https://deepwiki.com/charmbracelet/x/7-development-and-testing#testing-infrastructure-components)
    
*   [Golden File Testing System](https://deepwiki.com/charmbracelet/x/7-development-and-testing#golden-file-testing-system)
    
*   [Cross-Platform Testing](https://deepwiki.com/charmbracelet/x/7-development-and-testing#cross-platform-testing)
    
*   [Development and Debugging Tools](https://deepwiki.com/charmbracelet/x/7-development-and-testing#development-and-debugging-tools)
    
*   [Test Coverage Patterns](https://deepwiki.com/charmbracelet/x/7-development-and-testing#test-coverage-patterns)
    
*   [Unit Test Coverage](https://deepwiki.com/charmbracelet/x/7-development-and-testing#unit-test-coverage)
    
*   [Integration Test Coverage](https://deepwiki.com/charmbracelet/x/7-development-and-testing#integration-test-coverage)
    
*   [Example Application Coverage](https://deepwiki.com/charmbracelet/x/7-development-and-testing#example-application-coverage)
