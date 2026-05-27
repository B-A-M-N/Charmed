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

Examples and Usage Patterns
===========================

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/lipgloss/blob/30441468/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1)
    
*   [examples/layout/main.go](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/layout/main.go)
    
*   [examples/list/sublist/main.go](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/list/sublist/main.go)
    
*   [examples/ssh/main.go](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/ssh/main.go)
    
*   [size.go](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go)
    

This page provides practical examples of using lipgloss for terminal UI development, from basic text styling to complex layouts and components. The examples demonstrate common patterns for building styled terminal interfaces, including table rendering, list formatting, tree visualization, and SSH server integration.

The lipgloss repository includes a dedicated examples module at [examples/](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/)
 that showcases real-world integration patterns. For basic styling concepts, see [Core Style System](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system)
. For component-specific details, see [Structured Components](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition)
.

Basic Styling Patterns
----------------------

### Creating and Configuring Styles

Styles are created using `lipgloss.NewStyle()` and configured through a fluent API. All setter methods return a new `Style` instance, allowing method chaining:

Sources: [README.md22-35](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L22-L35)

### Color Application Patterns

Lipgloss supports multiple color specification methods that automatically degrade based on terminal capabilities:

| Color Type | Example | Use Case |
| --- | --- | --- |
| ANSI 16 colors | `lipgloss.Color("5")` | Basic terminal compatibility |
| ANSI 256 colors | `lipgloss.Color("201")` | Extended color palette |
| True Color (hex) | `lipgloss.Color("#0000FF")` | Full RGB color specification |
| Adaptive colors | `lipgloss.AdaptiveColor{Light: "236", Dark: "248"}` | Light/dark background adaptation |
| Complete colors | `lipgloss.CompleteColor{TrueColor: "#0000FF", ANSI256: "86", ANSI: "5"}` | Explicit per-profile colors |

The color profile is automatically detected via the `Renderer` and colors outside the terminal's gamut are coerced to the closest available value.

Sources: [README.md37-104](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L37-L104)

### Text Formatting Attributes

ANSI text formatting is applied through boolean setter methods:

Each attribute method sets the corresponding text property in the style's internal fields. See [Text Attributes and Styling](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#2.4)
 for implementation details.

Sources: [README.md106-119](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L106-L119)

### Block-Level Layout Patterns

#### Padding and Margins

Padding and margin methods follow CSS-style shorthand conventions:

Individual side methods are also available: `PaddingTop()`, `PaddingRight()`, `PaddingBottom()`, `PaddingLeft()`, and corresponding margin methods.

Sources: [README.md121-157](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L121-L157)

#### Width, Height, and Alignment

Minimum dimensions and text alignment are specified through setter methods:

Alignment values include `lipgloss.Left`, `lipgloss.Right`, and `lipgloss.Center`. The `Width` and `Height` methods set minimum dimensions that are enforced during rendering.

Sources: [README.md159-181](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L159-L181)

### Border Application Patterns

Borders are applied using predefined border styles or custom `Border` structs:

Predefined border styles include `NormalBorder()`, `RoundedBorder()`, `ThickBorder()`, `DoubleBorder()`, `HiddenBorder()`, and `BlockBorder()`. See [Borders](https://deepwiki.com/charmbracelet/lipgloss/2.3-rendering-pipeline)
 for border rendering implementation.

Sources: [README.md183-226](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L183-L226)

### Style Composition Patterns

#### Copying Styles

`Style` structs contain only primitive types, so assignment creates a true copy without mutation:

Sources: [README.md230-245](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L230-L245)

#### Style Inheritance

The `Inherit()` method copies unset properties from a source style:

Inheritance respects the `props` bitmask to determine which properties are set. See [Style Configuration](https://deepwiki.com/charmbracelet/lipgloss/2.1-style-configuration)
 for property management details.

Sources: [README.md247-262](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L247-L262)

#### Unsetting Properties

All style properties can be unset using `Unset*()` methods:

Unsetting clears both the property value and its corresponding bit in the `props` bitmask, preventing inheritance and copying of that property.

Sources: [README.md264-277](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L264-L277)

### Rendering Enforcement Patterns

The `Inline()`, `MaxWidth()`, and `MaxHeight()` methods enforce rendering constraints:

These methods are useful for component development where styles must respect specific UI constraints.

Sources: [README.md278-293](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L278-L293)

### Tab Width Configuration

Tab rendering is normalized to prevent terminal inconsistencies:

The default behavior converts tabs to 4 spaces during rendering. Use `TabWidth()` to customize or `NoTabConversion` to preserve original tab characters.

Sources: [README.md295-307](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L295-L307)

Layout Composition Patterns
---------------------------

### Measurement Functions

The `Width()`, `Height()`, and `Size()` functions measure rendered text accounting for ANSI escape sequences and multi-cell characters:

These functions use `ansi.StringWidth()` for accurate width calculation that handles Unicode grapheme clusters and ANSI escape codes. See [Measurement Utilities](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#4.1)
 for implementation details.

Sources: [README.md370-388](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L370-L388)
 [size.go1-42](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go#L1-L42)

### Joining Text Blocks

#### Diagram: JoinHorizontal and JoinVertical Flow

Sources: [README.md354-368](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L354-L368)

`JoinHorizontal()` and `JoinVertical()` combine multiple text blocks with position-based alignment:

The `Position` type accepts constants (`Left`, `Right`, `Top`, `Bottom`, `Center`) or float64 values (0.0-1.0) for fine-grained alignment control. See [Joining and Alignment](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#4.2)
 for implementation details.

Sources: [README.md354-368](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L354-L368)

### Text Placement in Whitespace

The `Place()`, `PlaceHorizontal()`, and `PlaceVertical()` functions position text within defined dimensions:

Whitespace can be styled using the optional trailing `Style` parameter. The placement functions use the internal `whitespace` renderer to generate styled padding. See [Placement](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#4.3)
 for details.

Sources: [README.md390-407](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L390-L407)

Dependency Ecosystem
--------------------

The examples module integrates lipgloss with specialized terminal application frameworks through the following dependency stack:

### Integration Categories by Dependencies

Sources: [examples/go.mod9-18](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L9-L18)
 [examples/go.mod20-50](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L20-L50)

Application Integration Categories
----------------------------------

The dependency configuration indicates three primary integration categories:

### SSH Server Integration

Dependencies: `charmbracelet/ssh` v0.0.0-20240401141849-854cddfa2917, `charmbracelet/wish` v1.4.0

*   SSH server framework for styled terminal interfaces
*   Authentication middleware with session management
*   Remote terminal rendering validation

Sources: [examples/go.mod11-12](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L11-L12)

### Color Processing Integration

Dependencies: `lucasb-eyer/go-colorful` v1.2.0, `muesli/gamut` v0.3.1

*   Color space conversion operations
*   Palette generation with clustering algorithms (`muesli/clusters`, `muesli/kmeans`)
*   Color theory applications in terminal interfaces

Sources: [examples/go.mod14-16](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L14-L16)

### Terminal Control Integration

Dependencies: `creack/pty` v1.1.21, `muesli/termenv` v0.16.0, `golang.org/x/term` v0.29.0

*   Pseudo-terminal creation and management
*   Terminal capability detection and adaptive rendering
*   Cross-platform terminal I/O operations

Sources: [examples/go.mod13-17](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L13-L17)

Development and Testing Integration
-----------------------------------

The examples module serves dual purposes as both demonstration and integration testing infrastructure:

### Dependency Version Management

The module maintains specific version locks for external dependencies while using local replacements for lipgloss components. This approach ensures:

*   **Development Synchronization**: Examples always use the latest local lipgloss development version
*   **Stable External Dependencies**: Locked versions prevent external dependency changes from breaking examples
*   **Integration Testing**: Examples serve as comprehensive integration tests for lipgloss functionality

### Module Isolation

The separate module structure provides:

*   **Independent Builds**: Examples can be built and tested independently of the main library
*   **Dependency Isolation**: Example-specific dependencies don't affect the main lipgloss module
*   **Version Compatibility Testing**: Examples validate lipgloss compatibility across different dependency versions

Sources: [examples/go.mod1-7](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.mod#L1-L7)
 [examples/go.sum1-89](https://github.com/charmbracelet/lipgloss/blob/30441468/examples/go.sum#L1-L89)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Examples and Usage Patterns](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#examples-and-usage-patterns)
    
*   [Basic Styling Patterns](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#basic-styling-patterns)
    
*   [Creating and Configuring Styles](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#creating-and-configuring-styles)
    
*   [Color Application Patterns](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#color-application-patterns)
    
*   [Text Formatting Attributes](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#text-formatting-attributes)
    
*   [Block-Level Layout Patterns](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#block-level-layout-patterns)
    
*   [Padding and Margins](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#padding-and-margins)
    
*   [Width, Height, and Alignment](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#width-height-and-alignment)
    
*   [Border Application Patterns](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#border-application-patterns)
    
*   [Style Composition Patterns](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#style-composition-patterns)
    
*   [Copying Styles](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#copying-styles)
    
*   [Style Inheritance](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#style-inheritance)
    
*   [Unsetting Properties](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#unsetting-properties)
    
*   [Rendering Enforcement Patterns](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#rendering-enforcement-patterns)
    
*   [Tab Width Configuration](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#tab-width-configuration)
    
*   [Layout Composition Patterns](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#layout-composition-patterns)
    
*   [Measurement Functions](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#measurement-functions)
    
*   [Joining Text Blocks](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#joining-text-blocks)
    
*   [Diagram: JoinHorizontal and JoinVertical Flow](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#diagram-joinhorizontal-and-joinvertical-flow)
    
*   [Text Placement in Whitespace](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#text-placement-in-whitespace)
    
*   [Dependency Ecosystem](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#dependency-ecosystem)
    
*   [Integration Categories by Dependencies](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#integration-categories-by-dependencies)
    
*   [Application Integration Categories](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#application-integration-categories)
    
*   [SSH Server Integration](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#ssh-server-integration)
    
*   [Color Processing Integration](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#color-processing-integration)
    
*   [Terminal Control Integration](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#terminal-control-integration)
    
*   [Development and Testing Integration](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#development-and-testing-integration)
    
*   [Dependency Version Management](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#dependency-version-management)
    
*   [Module Isolation](https://deepwiki.com/charmbracelet/lipgloss/7-examples-and-usage-patterns#module-isolation)
