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

Introduction to Lipgloss
========================

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/lipgloss/blob/30441468/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1)
    
*   [go.mod](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/lipgloss/blob/30441468/go.sum)
    
*   [size.go](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go)
    

Purpose and Scope
-----------------

This document provides an overview of the lipgloss library, covering its architecture, installation, module structure, and fundamental concepts. Lipgloss is a Go library for creating styled terminal user interfaces with a declarative, CSS-inspired approach.

For detailed information about the core styling system, see [Core Style System](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system)
. For component-specific documentation, see [Components](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition)
. For practical examples and usage patterns, see [Examples](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns)
.

What is Lipgloss
----------------

Lipgloss is a comprehensive terminal UI styling library that enables developers to create visually appealing command-line interfaces and terminal user interfaces (TUIs). The library takes an expressive, declarative approach to terminal rendering, making it familiar to developers with CSS experience.

**Core Features:**

*   **Declarative Styling**: Fluent API with method chaining similar to CSS
*   **Color System**: ANSI 16-color, 256-color, True Color (24-bit), and adaptive color support
*   **Text Formatting**: Bold, italic, faint, blink, strikethrough, underline, reverse
*   **Layout System**: Padding, margins, borders, alignment, width/height constraints
*   **Positioning**: Text placement in whitespace with horizontal/vertical joining
*   **Style Management**: Style copying, inheritance, and rule unsetting
*   **Components**: Specialized table, list, and tree rendering sub-packages
*   **Terminal Adaptation**: Automatic capability detection and color degradation
*   **Custom Renderers**: Context-specific rendering for different outputs (e.g., SSH sessions)

The library is designed to work seamlessly with terminal user interface frameworks like Bubble Tea while providing standalone utility for command-line applications.

Installation and Module Structure
---------------------------------

### Go Module Information

Lipgloss is distributed as a Go module at `charm.land/lipgloss/v2` requiring Go 1.24.2 or later. The module includes version retractions for problematic beta releases ([go.mod3](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L3-L3)
).

**Installation:**

### Dependencies

The library has 10 direct dependencies and 8 indirect dependencies, primarily from the Charmbracelet ecosystem:

| Dependency | Version | Purpose |
| --- | --- | --- |
| `github.com/charmbracelet/colorprofile` | v0.4.2 | Terminal color capability detection |
| `github.com/charmbracelet/ultraviolet` | v0.0.0-20251205161215 | SSH rendering support via `uv.ScreenBuffer` |
| `github.com/charmbracelet/x/ansi` | v0.11.6 | ANSI escape sequence handling |
| `github.com/charmbracelet/x/term` | v0.2.2 | Terminal control operations |
| `github.com/lucasb-eyer/go-colorful` | v1.3.0 | Color manipulation (darken, lighten, blend) |
| `github.com/rivo/uniseg` | v0.4.7 | Unicode text segmentation |
| `github.com/clipperhouse/displaywidth` | v0.11.0 | Text width calculation for terminal cells |
| `github.com/aymanbagabas/go-udiff` | v0.4.0 | Unified diff operations (testing) |
| `golang.org/x/sys` | v0.41.0 | Platform-specific system calls |

Sources: [go.mod1-31](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L1-L31)
 [go.sum1-36](https://github.com/charmbracelet/lipgloss/blob/30441468/go.sum#L1-L36)

### Package Structure

**Module Hierarchy:**

The main module provides core styling and components. The `examples/` directory is a separate Go module that uses a `replace` directive to depend on the local development version of lipgloss.

Sources: [go.mod1-5](https://github.com/charmbracelet/lipgloss/blob/30441468/go.mod#L1-L5)
 [README.md31-35](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L31-L35)
 [README.md495-797](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L495-L797)

Core Architecture Overview
--------------------------

### System Organization

The lipgloss architecture is organized into eight major subsystems that work together to provide terminal styling capabilities:

**Architecture Diagram:**

This architecture separates concerns into focused subsystems: the **Core Styling Engine** manages style properties and rendering logic; the **Color System** handles terminal capability detection and color adaptation; the **Layout System** provides text measurement and arrangement; the **Border System** generates border characters; **Components** build structured renderers on top of the core; the **Composition** system enables layered UIs; and the **Output** system manages terminal-aware rendering.

Sources: [README.md17-31](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L17-L31)
 [size.go1-41](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go#L1-L41)
 [README.md239-292](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L239-L292)

### Style System Foundation

The core of lipgloss is the `Style` struct defined in `style.go`, which manages visual properties through a fluent API. The style system is organized around three key aspects:

**Property Management:**

Each setter method returns a new `Style` value (immutable design), making styles safe to copy via simple assignment. Unset properties are not inherited or copied, allowing for flexible style composition.

Sources: [README.md294-338](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L294-L338)
 [README.md129-137](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L129-L137)

### Core Rendering Pipeline

The rendering process operates in three distinct phases, transforming style configuration into terminal-ready ANSI output:

**Rendering Pipeline Phases:**

The rendering order is critical: inner elements (padding) are applied before outer elements (borders and margins), ensuring correct visual hierarchy. Each stage transforms the content while preserving ANSI styling information.

Sources: [README.md381-396](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L381-L396)
 [README.md929-947](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L929-L947)

### Key Code Entities Reference

The library provides these primary code entities for user interaction:

| Category | Entity | Purpose |
| --- | --- | --- |
| **Core Creation** | `lipgloss.NewStyle()` | Creates new Style instance with fluent API |
|     | `lipgloss.NewRenderer()` | Creates Renderer for custom terminal contexts |
| **Style Application** | `Style.Render(str)` | Applies style to string, returns ANSI output |
|     | `Style.String()` | Stringer interface, renders preset string |
|     | `Style.Inherit(s)` | Inherits unset properties from another style |
| **Colors** | `lipgloss.Color(str)` | Creates Color from hex, ANSI, or RGB |
|     | `lipgloss.HasDarkBackground(in, out)` | Detects if terminal has dark background |
|     | `lipgloss.LightDark(dark)` | Returns function for light/dark color selection |
|     | `lipgloss.Complete(profile)` | Returns function for profile-based color selection |
|     | `lipgloss.Darken(c, p)` | Returns darkened color by percentage |
|     | `lipgloss.Lighten(c, p)` | Returns lightened color by percentage |
|     | `lipgloss.Complementary(c)` | Returns complementary color |
|     | `lipgloss.Alpha(c, a)` | Returns color with alpha transparency |
| **Measurement** | `lipgloss.Width(str)` | Returns cell width (ANSI-aware) |
|     | `lipgloss.Height(str)` | Returns line count |
|     | `lipgloss.Size(str)` | Returns width and height |
| **Layout** | `lipgloss.JoinHorizontal(pos, strs...)` | Joins strings horizontally |
|     | `lipgloss.JoinVertical(pos, strs...)` | Joins strings vertically |
|     | `lipgloss.PlaceHorizontal(w, pos, str)` | Places string in horizontal space |
|     | `lipgloss.PlaceVertical(h, pos, str)` | Places string in vertical space |
|     | `lipgloss.Place(w, h, hPos, vPos, str)` | Places string in 2D space |
| **Composition** | `lipgloss.NewLayer(str)` | Creates Layer for composition |
|     | `Layer.X(x).Y(y).Z(z)` | Sets layer position and depth |
|     | `lipgloss.NewCompositor()` | Creates Compositor for layering |
|     | `Compositor.Compose(layers...)` | Composes layers into Canvas |
|     | `Canvas.Render()` | Renders composed layers to string |
| **Output** | `lipgloss.Print(a...)` | Prints with color downsampling |
|     | `lipgloss.Println(a...)` | Prints with newline |
|     | `lipgloss.Sprint(a...)` | Returns string with downsampling |
|     | `lipgloss.Fprint(w, a...)` | Writes to io.Writer |
| **Components** | `table.New()` | Creates Table instance |
|     | `list.New(items...)` | Creates List instance |
|     | `tree.Root(str)` | Creates Tree root node |
| **Borders** | `lipgloss.NormalBorder()` | Returns normal border definition |
|     | `lipgloss.RoundedBorder()` | Returns rounded border definition |
|     | `lipgloss.ThickBorder()` | Returns thick border definition |
|     | `lipgloss.DoubleBorder()` | Returns double-line border definition |

Sources: [README.md17-31](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L17-L31)
 [size.go9-41](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go#L9-L41)
 [README.md423-487](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L423-L487)
 [README.md409-417](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L409-L417)
 [README.md107-116](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L107-L116)

Basic Usage Pattern
-------------------

The fundamental usage pattern follows a declarative approach where styles are created, configured, and applied to content:

This pattern demonstrates:

*   Fluent API design for method chaining
*   Separation of style definition from content
*   Immutable style objects that can be reused
*   Terminal-aware rendering with automatic adaptation

Sources: [README.md18-31](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L18-L31)

Color System and Terminal Adaptation
------------------------------------

### Color Profiles and Detection

Lipgloss supports four color profiles with automatic detection and downsampling:

| Profile | Colors | Input Format | Example |
| --- | --- | --- | --- |
| TrueColor | 16,777,216 (24-bit) | Hex strings | `lipgloss.Color("#FF0000")` |
| ANSI256 | 256 colors (8-bit) | Numeric strings | `lipgloss.Color("201")` |
| ANSI | 16 colors (4-bit) | Numeric strings | `lipgloss.Color("9")` |
| ASCII | 2 colors (1-bit) | Monochrome only | Black and white |

**Color Detection and Adaptation:**

The detection happens once at startup. `HasDarkBackground()` queries the terminal using OSC 11 ANSI escape sequences to determine background color. `colorprofile.Detect()` examines environment variables and terminal features to determine capability level.

Sources: [README.md42-103](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L42-L103)
 [README.md799-898](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L799-L898)
 [README.md900-950](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L900-L950)

### Adaptive Color Functions

Lipgloss provides two primary patterns for runtime color adaptation:

**Light/Dark Background Adaptation:**

**Profile-Based Adaptation:**

The `writer.go` functions (`Println`, `Sprint`, `Fprint`, etc.) automatically downsample colors using `colorprofile.Writer` to match terminal capabilities, stripping ANSI codes entirely when output is not a TTY.

Sources: [README.md830-898](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L830-L898)
 [README.md900-917](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L900-L917)
 [README.md919-950](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L919-L950)

Component Ecosystem
-------------------

Lipgloss provides specialized rendering components built on the core styling system:

| Component | Import Path | Purpose |
| --- | --- | --- |
| Tables | `github.com/charmbracelet/lipgloss/table` | Structured data rendering with styling |
| Lists | `github.com/charmbracelet/lipgloss/list` | Enumerated and nested list rendering |
| Trees | `github.com/charmbracelet/lipgloss/tree` | Hierarchical data visualization |

Each component maintains compatibility with the core `Style` system while providing specialized functionality for their respective use cases.

Sources: [README.md405-715](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L405-L715)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Introduction to Lipgloss](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#introduction-to-lipgloss)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#purpose-and-scope)
    
*   [What is Lipgloss](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#what-is-lipgloss)
    
*   [Installation and Module Structure](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#installation-and-module-structure)
    
*   [Go Module Information](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#go-module-information)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#dependencies)
    
*   [Package Structure](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#package-structure)
    
*   [Core Architecture Overview](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#core-architecture-overview)
    
*   [System Organization](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#system-organization)
    
*   [Style System Foundation](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#style-system-foundation)
    
*   [Core Rendering Pipeline](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#core-rendering-pipeline)
    
*   [Key Code Entities Reference](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#key-code-entities-reference)
    
*   [Basic Usage Pattern](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#basic-usage-pattern)
    
*   [Color System and Terminal Adaptation](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#color-system-and-terminal-adaptation)
    
*   [Color Profiles and Detection](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#color-profiles-and-detection)
    
*   [Adaptive Color Functions](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#adaptive-color-functions)
    
*   [Component Ecosystem](https://deepwiki.com/charmbracelet/lipgloss/1-introduction-to-lipgloss#component-ecosystem)
