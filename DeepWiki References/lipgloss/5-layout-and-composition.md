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

Layout and Composition
======================

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/lipgloss/blob/30441468/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1)
    
*   [align.go](https://github.com/charmbracelet/lipgloss/blob/30441468/align.go)
    
*   [join.go](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go)
    
*   [lipgloss.go](https://github.com/charmbracelet/lipgloss/blob/30441468/lipgloss.go)
    
*   [position.go](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go)
    
*   [size.go](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go)
    

Purpose and Scope
-----------------

This document provides an overview of Lipgloss's layout and composition capabilities, which enable arranging styled text blocks into complex terminal user interfaces. The system operates at two levels: **simple string-based layout functions** for basic arrangements, and an **advanced layer-based composition system** for complex overlapping UIs.

For detailed information about specific subsystems:

*   Text measurement primitives: see [Measurement Utilities](https://deepwiki.com/charmbracelet/lipgloss/5.1-measurement-utilities)
    
*   Horizontal and vertical joining: see [Joining and Alignment](https://deepwiki.com/charmbracelet/lipgloss/5.2-joining-and-alignment)
    
*   Positioning text in whitespace: see [Placement](https://deepwiki.com/charmbracelet/lipgloss/5.3-placement)
    
*   Complex layered UIs: see [Canvas and Layer System](https://deepwiki.com/charmbracelet/lipgloss/5.4-canvas-and-layer-system)
    

This page covers the architectural design, key concepts, and relationships between these subsystems. Information about the core styling system can be found in [Core Style System](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system)
, while border rendering is covered in [Borders](https://deepwiki.com/charmbracelet/lipgloss/4-borders)
.

**Sources:** [README.md1-997](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L1-L997)

Two-Level Architecture
----------------------

Lipgloss provides two distinct approaches to layout and composition, each designed for different use cases:

### Simple Layout Functions

String-based utilities that work directly with text blocks. These functions are stateless, returning new strings without modifying inputs. They handle multi-line strings, preserve ANSI styling, and account for wide Unicode characters.

| Function | Purpose | Key Files |
| --- | --- | --- |
| `Width`, `Height`, `Size` | Measure text dimensions | [size.go1-41](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go#L1-L41) |
| `JoinHorizontal`, `JoinVertical` | Combine text blocks | [join.go1-176](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L1-L176) |
| `PlaceHorizontal`, `PlaceVertical`, `Place` | Position in whitespace | [position.go1-135](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L1-L135) |

### Advanced Composition System

Cell-based compositor for rendering layered, overlapping content with Z-ordering. This system uses a buffer approach where layers are flattened into absolute coordinates and rendered based on depth.

| Component | Purpose | Key Files |
| --- | --- | --- |
| `Layer` | Content with position and Z-index | layer.go |
| `Canvas` | Cell buffer for drawing | canvas.go |
| `Compositor` | Combines and renders layers | layer.go |

The simple layout functions are sufficient for most terminal UIs (status bars, dialogs, columns). The advanced system is necessary when elements need to overlap or when interactive positioning is required.

**Sources:** [README.md397-489](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L397-L489)
 [join.go1-176](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L1-L176)
 [position.go1-135](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L1-L135)
 [size.go1-41](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go#L1-L41)

Layout System Architecture
--------------------------

**Diagram: Layout System Component Relationships**

This diagram shows how measurement utilities form the foundation for both simple layout functions and how they integrate with the style rendering pipeline. The `Position` type is shared across multiple functions to provide consistent alignment behavior.

**Sources:** [size.go1-41](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go#L1-L41)
 [join.go1-176](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L1-L176)
 [position.go1-135](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L1-L135)
 [align.go1-83](https://github.com/charmbracelet/lipgloss/blob/30441468/align.go#L1-L83)

Position System
---------------

The `Position` type provides a unified way to specify alignment and placement across all layout functions:

**Diagram: Position Type Usage Across Layout Functions**

The `Position` type represents a fractional value along an axis, where 0.0 is the start and 1.0 is the end. The [position.go19-32](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L19-L32)
 implementation clamps values to this range using `Position.value()`. Constants like `Top`, `Bottom`, `Center`, `Left`, and `Right` are provided for readability, but any float64 value between 0 and 1 can be used for precise positioning.

**Sources:** [position.go19-32](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L19-L32)
 [README.md427-437](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L427-L437)

Simple Layout Workflow
----------------------

**Diagram: Simple Layout Function Processing Flow**

This diagram shows the internal processing steps for the three main categories of simple layout functions. All functions preserve ANSI escape sequences and properly measure wide Unicode characters using `ansi.StringWidth()`.

**Sources:** [join.go28-175](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L28-L175)
 [position.go36-134](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L36-L134)

Key Implementation Details
--------------------------

### ANSI-Aware Measurement

All layout functions use `ansi.StringWidth()` from the `x/ansi` package to measure text width, which correctly handles:

*   ANSI escape sequences (ignored in width calculation)
*   Wide Unicode characters (counted as 2 cells)
*   Emoji and other multi-cell glyphs

The `getLines()` helper function [join.go48](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L48-L48)
 is used throughout to split multi-line strings and calculate maximum width while respecting these rules.

### Whitespace Rendering

The placement functions (`PlaceHorizontal`, `PlaceVertical`, `Place`) support styling the whitespace used for positioning through `WhitespaceOption`:

The `newWhitespace()` function creates a whitespace renderer that can apply styles to padding characters.

**Sources:** [position.go36-134](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L36-L134)
 [README.md471-489](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L471-L489)

### Alignment Within Styles

The alignment functions in [align.go1-83](https://github.com/charmbracelet/lipgloss/blob/30441468/align.go#L1-L83)
 are used internally by the `Style.Render()` pipeline when the `Align()` method is set on a style. These functions differ from placement functions in that they:

*   Operate within the context of a styled block (can apply style to padding)
*   Are called during the rendering pipeline, not as standalone utilities
*   Work with the `width` and `height` properties of a style

**Diagram: Alignment in Rendering Pipeline**

**Sources:** [align.go1-83](https://github.com/charmbracelet/lipgloss/blob/30441468/align.go#L1-L83)
 [README.md215-225](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L215-L225)

Usage Patterns
--------------

### Horizontal Layout Example

The `JoinHorizontal` function [join.go28-96](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L28-L96)
 aligns blocks along their vertical axis. When blocks have different heights, the shorter ones are padded with empty lines according to the position parameter.

### Vertical Stacking Example

The `JoinVertical` function [join.go116-175](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L116-L175)
 stacks blocks vertically, aligning them along their horizontal axis. Lines are padded with spaces to match the widest block.

### Placement in Whitespace Example

Placement functions [position.go36-134](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L36-L134)
 position text within a defined area, filling the remaining space with whitespace (which can be styled).

**Sources:** [README.md427-489](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L427-L489)
 [join.go1-176](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L1-L176)
 [position.go1-135](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L1-L135)

Advanced Composition Overview
-----------------------------

For complex UIs requiring overlapping elements or interactive positioning, Lipgloss provides a layer-based composition system:

This system uses:

*   **Layer**: Content with X, Y, Z coordinates and optional child layers
*   **Canvas**: Cell-based buffer using `uv.ScreenBuffer` for efficient drawing
*   **Compositor**: Flattens layer hierarchy, sorts by Z-index, and renders to canvas

Detailed information about this system can be found in [Canvas and Layer System](https://deepwiki.com/charmbracelet/lipgloss/5.4-canvas-and-layer-system)
.

**Sources:** [README.md402-421](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L402-L421)

Relationship to Style System
----------------------------

The layout functions integrate with the core style system in several ways:

| Integration Point | Description | Reference |
| --- | --- | --- |
| Alignment | `alignTextHorizontal()` and `alignTextVertical()` are called by `Style.Render()` | [align.go1-83](https://github.com/charmbracelet/lipgloss/blob/30441468/align.go#L1-L83) |
| Measurement | `Width()`, `Height()`, `Size()` measure rendered style output | [size.go1-41](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go#L1-L41) |
| Composition | Layout functions work with styled strings (ANSI codes preserved) | [join.go1-176](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L1-L176) |
| Rendering | Style padding and margins can be combined with placement | [position.go1-135](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L1-L135) |

The style rendering pipeline (see [Rendering Pipeline](https://deepwiki.com/charmbracelet/lipgloss/2.3-rendering-pipeline)
) applies alignment internally, while layout functions are typically used to compose multiple styled blocks together.

**Sources:** [align.go1-83](https://github.com/charmbracelet/lipgloss/blob/30441468/align.go#L1-L83)
 [size.go1-41](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go#L1-L41)
 [README.md1-997](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L1-L997)

Performance Considerations
--------------------------

The simple layout functions are designed for efficiency:

1.  **Single-pass processing**: Most functions iterate through content once
2.  **String builder usage**: All functions use `strings.Builder` for efficient string concatenation [join.go82](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L82-L82)
     [position.go53](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L53-L53)
    
3.  **No allocation for measurement**: The `getLines()` helper returns slices of the original string
4.  **ANSI sequence preservation**: No re-parsing or re-application of ANSI codes

For very large or frequently updated layouts, consider:

*   Caching measurement results when content doesn't change
*   Using the Canvas system for absolute positioning instead of repeated joins
*   Pre-calculating positions for static elements

**Sources:** [join.go1-176](https://github.com/charmbracelet/lipgloss/blob/30441468/join.go#L1-L176)
 [position.go1-135](https://github.com/charmbracelet/lipgloss/blob/30441468/position.go#L1-L135)
 [size.go1-41](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go#L1-L41)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Layout and Composition](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#layout-and-composition)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#purpose-and-scope)
    
*   [Two-Level Architecture](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#two-level-architecture)
    
*   [Simple Layout Functions](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#simple-layout-functions)
    
*   [Advanced Composition System](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#advanced-composition-system)
    
*   [Layout System Architecture](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#layout-system-architecture)
    
*   [Position System](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#position-system)
    
*   [Simple Layout Workflow](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#simple-layout-workflow)
    
*   [Key Implementation Details](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#key-implementation-details)
    
*   [ANSI-Aware Measurement](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#ansi-aware-measurement)
    
*   [Whitespace Rendering](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#whitespace-rendering)
    
*   [Alignment Within Styles](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#alignment-within-styles)
    
*   [Usage Patterns](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#usage-patterns)
    
*   [Horizontal Layout Example](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#horizontal-layout-example)
    
*   [Vertical Stacking Example](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#vertical-stacking-example)
    
*   [Placement in Whitespace Example](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#placement-in-whitespace-example)
    
*   [Advanced Composition Overview](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#advanced-composition-overview)
    
*   [Relationship to Style System](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#relationship-to-style-system)
    
*   [Performance Considerations](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition#performance-considerations)
