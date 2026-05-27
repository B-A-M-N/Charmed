Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 8 March 2026 ([304414](https://github.com/charmbracelet/lipgloss/commits/30441468)
)

*   [Introduction to Lipgloss](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss)
    
*   [Core Style System](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system)
    
*   [Style Configuration](https://deepwiki.com/charmbracelet/lipgloss/2.1-style-configuration)
    
*   [Text Attributes and Styling](https://deepwiki.com/charmbracelet/lipgloss/2.2-text-attributes-and-styling)
    
*   [Rendering Pipeline](https://deepwiki.com/charmbracelet/lipgloss/2.3-rendering-pipeline)
    
*   [Colors and Terminal Detection](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection)
    
*   [Color System](https://deepwiki.com/charmbracelet/lipgloss/3.1-color-system)
    
*   [Adaptive Colors and Background Detection](https://deepwiki.com/charmbracelet/lipgloss/3.2-adaptive-colors-and-background-detection)
    
*   [Terminal Adaptation and Output](https://deepwiki.com/charmbracelet/lipgloss/3.3-terminal-adaptation-and-output)
    
*   [Borders](https://deepwiki.com/charmbracelet/lipgloss/4-borders)
    
*   [Layout and Composition](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition)
    
*   [Measurement Utilities](https://deepwiki.com/charmbracelet/lipgloss/5.1-measurement-utilities)
    
*   [Joining and Alignment](https://deepwiki.com/charmbracelet/lipgloss/5.2-joining-and-alignment)
    
*   [Placement](https://deepwiki.com/charmbracelet/lipgloss/5.3-placement)
    
*   [Canvas and Layer System](https://deepwiki.com/charmbracelet/lipgloss/5.4-canvas-and-layer-system)
    
*   [Structured Components](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components)
    
*   [Tables](https://deepwiki.com/charmbracelet/lipgloss/6.1-tables)
    
*   [Lists](https://deepwiki.com/charmbracelet/lipgloss/6.2-lists)
    
*   [Trees](https://deepwiki.com/charmbracelet/lipgloss/6.3-trees)
    
*   [Examples and Usage Patterns](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns)
    
*   [Basic Styling Examples](https://deepwiki.com/charmbracelet/lipgloss/7.1-basic-styling-examples)
    
*   [Layout and Document Examples](https://deepwiki.com/charmbracelet/lipgloss/7.2-layout-and-document-examples)
    
*   [Component Examples](https://deepwiki.com/charmbracelet/lipgloss/7.3-component-examples)
    
*   [Advanced Integration Patterns](https://deepwiki.com/charmbracelet/lipgloss/7.4-advanced-integration-patterns)
    
*   [Development and Testing](https://deepwiki.com/charmbracelet/lipgloss/8-development-and-testing)
    
*   [Testing Infrastructure](https://deepwiki.com/charmbracelet/lipgloss/8.1-testing-infrastructure)
    
*   [Code Quality and CI/CD](https://deepwiki.com/charmbracelet/lipgloss/8.2-code-quality-and-cicd)
    
*   [Module Architecture and Dependencies](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies)
    

Menu

Module Architecture and Dependencies
====================================

Relevant source files

*   [examples/go.mod](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod)
    
*   [examples/go.sum](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.sum)
    
*   [go.mod](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/lipgloss/blob/30441468/go.sum)
    

This document describes the Go module structure of Lipgloss, including the organization of the main library module and the examples module, their dependencies, and how Lipgloss integrates with the Charmbracelet ecosystem. For information about the internal package structure and code organization, see [Introduction to Lipgloss](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss)
. For details on testing infrastructure and CI/CD workflows, see [Code Quality and CI/CD](https://deepwiki.com/charmbracelet/lipgloss/8.2-code-quality-and-cicd)
.

Module Structure Overview
-------------------------

Lipgloss is organized as two separate Go modules: the main library module (`charm.land/lipgloss/v2`) and an examples module (`examples/`). This separation allows the examples to depend on the library while maintaining clean module boundaries during development.

### Main Module Configuration

The main Lipgloss module is defined at [go.mod1-5](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L1-L5)
:

    module charm.land/lipgloss/v2
    go 1.24.2
    toolchain go1.24.4
    

The module uses Go 1.24.2 as the minimum required version and Go 1.24.4 as the toolchain version. A retraction directive at [go.mod3](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L3-L3)
 marks `v2.0.0-beta1` as retracted due to version numbering conventions.

### Examples Module Configuration

The examples module is defined at [examples/go.mod1-7](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L1-L7)
:

    module examples
    go 1.24.3
    toolchain go1.24.4
    replace charm.land/lipgloss/v2 => ../
    

The `replace` directive at [examples/go.mod7](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L7-L7)
 enables the examples to depend on the local development version of Lipgloss, allowing simultaneous development and testing without publishing intermediate versions.

**Sources:** [go.mod1-31](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L1-L31)
 [examples/go.mod1-47](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L1-L47)

Dependency Categories
---------------------

Lipgloss dependencies are organized into three categories: Charmbracelet ecosystem packages, external third-party libraries, and platform-specific indirect dependencies.

### Charmbracelet Ecosystem Dependencies

The following table lists Charmbracelet packages used by Lipgloss:

| Package | Version | Purpose |
| --- | --- | --- |
| `github.com/charmbracelet/colorprofile` | v0.4.2 | Terminal color capability detection |
| `github.com/charmbracelet/ultraviolet` | v0.0.0-20251205161215 | SSH screen buffer rendering |
| `github.com/charmbracelet/x/ansi` | v0.11.6 | ANSI escape sequence utilities |
| `github.com/charmbracelet/x/term` | v0.2.2 | Terminal control and querying |
| `github.com/charmbracelet/x/termios` | v0.1.1 (indirect) | Terminal I/O settings |
| `github.com/charmbracelet/x/windows` | v0.2.2 (indirect) | Windows terminal support |
| `github.com/charmbracelet/x/exp/golden` | v0.0.0-20250806222409 | Golden file testing utilities |

**Sources:** [go.mod9-20](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L9-L20)

### External Third-Party Dependencies

| Package | Version | Purpose |
| --- | --- | --- |
| `github.com/lucasb-eyer/go-colorful` | v1.3.0 | Color manipulation and conversion |
| `github.com/rivo/uniseg` | v0.4.7 | Unicode text segmentation |
| `github.com/clipperhouse/displaywidth` | v0.11.0 | Visual display width calculation |
| `github.com/aymanbagabas/go-udiff` | v0.4.0 | Unified diff generation for testing |

**Sources:** [go.mod9-20](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L9-L20)

### Platform-Specific Indirect Dependencies

| Package | Version | Purpose |
| --- | --- | --- |
| `github.com/clipperhouse/uax29/v2` | v2.7.0 | Unicode text segmentation algorithms |
| `github.com/mattn/go-runewidth` | v0.0.19 | East Asian character width calculation |
| `github.com/muesli/cancelreader` | v0.2.2 | Cancellable reader for input |
| `github.com/xo/terminfo` | v0.0.0-20220910002029 | Terminal capability database |
| `golang.org/x/exp` | v0.0.0-20231006140011 | Experimental Go features |
| `golang.org/x/sync` | v0.18.0 | Synchronization primitives |
| `golang.org/x/sys` | v0.41.0 | Low-level system calls |

**Sources:** [go.mod22-30](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L22-L30)

Direct Dependencies Architecture
--------------------------------

**Sources:** [go.mod9-20](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L9-L20)
 [go.sum1-36](https://github.com/charmbracelet/lipgloss/blob/30441468/go.sum#L1-L36)

Main Module Direct Dependencies
-------------------------------

The main module declares 10 direct dependencies at [go.mod9-20](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L9-L20)
:

### Color and Terminal Detection

*   **`github.com/charmbracelet/colorprofile` v0.4.2**: Provides `Detect()` function for determining terminal color capabilities (TrueColor, ANSI256, ANSI, ASCII). Used throughout the color system for adaptive rendering.
    
*   **`github.com/charmbracelet/x/term` v0.2.2**: Supplies terminal querying capabilities, including background color detection through OSC sequences. The `QueryBackground()` function enables light/dark theme detection.
    
*   **`github.com/lucasb-eyer/go-colorful` v1.3.0**: Offers color space conversions and manipulation functions (RGBA, HSV, LAB). Used by `color.go` for color transformations like `Darken()`, `Lighten()`, and `Complementary()`.
    

### ANSI and Text Processing

*   **`github.com/charmbracelet/x/ansi` v0.11.6**: Provides ANSI escape sequence parsing and stripping utilities. Used extensively in the rendering pipeline to handle styled text, measure widths, and manage escape codes.
    
*   **`github.com/rivo/uniseg` v0.4.7**: Implements Unicode text segmentation according to UAX #29. Used for accurate grapheme cluster counting in text measurement functions.
    
*   **`github.com/clipperhouse/displaywidth` v0.11.0**: Calculates visual display width of strings, accounting for wide characters and combining marks. Essential for proper text alignment and wrapping.
    

### Canvas and Composition

*   **`github.com/charmbracelet/ultraviolet` v0.0.0-20251205161215-1948445e3318**: Provides `ScreenBuffer` type for character-level rendering in the Canvas system. Enables efficient composition of layered UIs.

### Testing Utilities

*   **`github.com/aymanbagabas/go-udiff` v0.4.0**: Generates unified diffs for test output comparison. Used in golden file tests to show differences between expected and actual output.
    
*   **`github.com/charmbracelet/x/exp/golden` v0.0.0-20250806222409-83e3a29d542f**: Framework for golden file testing. The `RequireEqual()` function compares rendered output against reference files.
    

### System-Level Access

*   **`golang.org/x/sys` v0.41.0**: Low-level system calls for platform-specific terminal operations. Required for Windows ANSI support and Unix terminal control.

**Sources:** [go.mod9-20](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L9-L20)
 [go.sum1-36](https://github.com/charmbracelet/lipgloss/blob/30441468/go.sum#L1-L36)

Indirect Dependencies
---------------------

The main module has 8 indirect dependencies at [go.mod22-30](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L22-L30)
 automatically managed by Go's module system:

| Package | Brought in by | Purpose |
| --- | --- | --- |
| `charmbracelet/x/termios` v0.1.1 | x/term | Terminal I/O control structures |
| `charmbracelet/x/windows` v0.2.2 | x/term, colorprofile | Windows console API wrappers |
| `clipperhouse/uax29/v2` v2.7.0 | displaywidth | UAX #29 segmentation implementation |
| `mattn/go-runewidth` v0.0.19 | displaywidth | East Asian Width property handling |
| `muesli/cancelreader` v0.2.2 | x/term | Cancellable I/O operations |
| `xo/terminfo` v0.0.0-20220910 | colorprofile | Terminal capability database access |
| `golang.org/x/exp` v0.0.0-20231006 | golden | Experimental standard library features |
| `golang.org/x/sync` v0.18.0 | x/term | Advanced synchronization primitives |

**Sources:** [go.mod22-30](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L22-L30)
 [go.sum13-36](https://github.com/charmbracelet/lipgloss/blob/30441468/go.sum#L13-L36)

Examples Module Dependencies
----------------------------

The examples module at [examples/go.mod1-47](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L1-L47)
 includes additional dependencies beyond the main library:

### Additional Charmbracelet Dependencies

| Package | Version | Purpose |
| --- | --- | --- |
| `charm.land/bubbletea/v2` | v2.0.0-rc.2.0.20251201 | TUI framework for interactive examples |
| `charmbracelet/ssh` | v0.0.0-20241211182756 | SSH server for remote rendering demos |
| `charmbracelet/wish/v2` | v2.0.0-20251106193208 | SSH middleware for TUI applications |
| `charmbracelet/log/v2` | v2.0.0-20251106192421 | Structured logging |
| `charmbracelet/x/exp/charmtone` | v0.0.0-20250627134340 | Terminal theme color palette |

The examples module uses a `replace` directive at [examples/go.mod46](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L46-L46)
 for the log package to use a specific development version.

**Sources:** [examples/go.mod9-18](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L9-L18)
 [examples/go.mod20-43](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L20-L43)

Dependency Flow and Integration
-------------------------------

**Sources:** [go.mod9-20](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L9-L20)

Module Version Constraints
--------------------------

### Minimum Go Version

Both modules require Go 1.24.2+ at [go.mod5](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L5-L5)
 and [examples/go.mod3](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L3-L3)
 with toolchain Go 1.24.4 specified at [go.mod7](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L7-L7)
 and [examples/go.mod5](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L5-L5)
 This ensures access to the latest Go features and performance improvements.

### Version Compatibility

The module uses semantic versioning with a v2 major version, indicated by the module path `charm.land/lipgloss/v2` at [go.mod1](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L1-L1)
 This allows breaking changes from v1 while maintaining clear upgrade paths for users.

The retraction at [go.mod3](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L3-L3)
 demonstrates version management:

    retract v2.0.0-beta1 // We add a "." after the "beta" in the version number.
    

This prevents users from accidentally depending on an incorrectly tagged beta version.

**Sources:** [go.mod1-7](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L1-L7)
 [examples/go.mod1-7](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L1-L7)

Charmbracelet Ecosystem Integration
-----------------------------------

Lipgloss is tightly integrated with the Charmbracelet ecosystem, relying on several complementary libraries:

### Terminal Detection Stack

1.  **`colorprofile`** - Top-level capability detection
2.  **`x/term`** - Terminal querying and control
3.  **`x/termios`** - Low-level terminal I/O (Unix)
4.  **`x/windows`** - Console API access (Windows)

This stack enables cross-platform terminal capability detection without requiring platform-specific code in Lipgloss itself.

### Text Processing Stack

1.  **`x/ansi`** - ANSI escape sequence handling
2.  **`displaywidth`** - Visual width calculation
3.  **`uniseg`** - Unicode segmentation
4.  **`go-colorful`** - Color space operations

These libraries provide the foundation for accurate text measurement and rendering across different Unicode scenarios and terminal types.

### Composition Stack

1.  **`ultraviolet`** - Screen buffer abstraction
2.  **Lipgloss Canvas** - High-level composition API
3.  **Lipgloss Layer** - Hierarchical positioning

The Canvas system builds on `ultraviolet`'s `ScreenBuffer` type to provide efficient character-level rendering for complex layered UIs.

**Sources:** [go.mod9-20](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L9-L20)

Development Workflow
--------------------

The module structure supports several development workflows:

### Local Development

The examples module's `replace` directive at [examples/go.mod7](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L7-L7)
 enables:

1.  Modify core Lipgloss code
2.  Run examples immediately without publishing
3.  Test changes in realistic scenarios
4.  Iterate rapidly on API design

### Dependency Updates

Dependencies are managed through standard Go module commands:

The `go.sum` file at [go.sum1-36](https://github.com/charmbracelet/lipgloss/blob/30441468/go.sum#L1-L36)
 ensures reproducible builds by recording cryptographic checksums of all dependencies.

### Version Testing

The module structure allows testing against specific dependency versions by editing `go.mod` require statements at [go.mod9-20](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L9-L20)
 or using `go get` with version specifiers.

**Sources:** [go.mod1-31](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L1-L31)
 [examples/go.mod1-47](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L1-L47)
 [go.sum1-36](https://github.com/charmbracelet/lipgloss/blob/30441468/go.sum#L1-L36)
 [examples/go.sum1-73](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.sum#L1-L73)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Module Architecture and Dependencies](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#module-architecture-and-dependencies)
    
*   [Module Structure Overview](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#module-structure-overview)
    
*   [Main Module Configuration](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#main-module-configuration)
    
*   [Examples Module Configuration](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#examples-module-configuration)
    
*   [Dependency Categories](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#dependency-categories)
    
*   [Charmbracelet Ecosystem Dependencies](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#charmbracelet-ecosystem-dependencies)
    
*   [External Third-Party Dependencies](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#external-third-party-dependencies)
    
*   [Platform-Specific Indirect Dependencies](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#platform-specific-indirect-dependencies)
    
*   [Direct Dependencies Architecture](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#direct-dependencies-architecture)
    
*   [Main Module Direct Dependencies](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#main-module-direct-dependencies)
    
*   [Color and Terminal Detection](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#color-and-terminal-detection)
    
*   [ANSI and Text Processing](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#ansi-and-text-processing)
    
*   [Canvas and Composition](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#canvas-and-composition)
    
*   [Testing Utilities](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#testing-utilities)
    
*   [System-Level Access](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#system-level-access)
    
*   [Indirect Dependencies](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#indirect-dependencies)
    
*   [Examples Module Dependencies](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#examples-module-dependencies)
    
*   [Additional Charmbracelet Dependencies](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#additional-charmbracelet-dependencies)
    
*   [Dependency Flow and Integration](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#dependency-flow-and-integration)
    
*   [Module Version Constraints](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#module-version-constraints)
    
*   [Minimum Go Version](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#minimum-go-version)
    
*   [Version Compatibility](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#version-compatibility)
    
*   [Charmbracelet Ecosystem Integration](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#charmbracelet-ecosystem-integration)
    
*   [Terminal Detection Stack](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#terminal-detection-stack)
    
*   [Text Processing Stack](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#text-processing-stack)
    
*   [Composition Stack](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#composition-stack)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#development-workflow)
    
*   [Local Development](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#local-development)
    
*   [Dependency Updates](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#dependency-updates)
    
*   [Version Testing](https://deepwiki.com/charmbracelet/lipgloss/9-module-architecture-and-dependencies#version-testing)
