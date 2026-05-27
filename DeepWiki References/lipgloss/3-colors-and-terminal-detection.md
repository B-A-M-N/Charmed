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

Colors and Terminal Detection
=============================

Relevant source files

*   [ansi\_unix.go](https://github.com/charmbracelet/lipgloss/blob/30441468/ansi_unix.go)
    
*   [ansi\_windows.go](https://github.com/charmbracelet/lipgloss/blob/30441468/ansi_windows.go)
    
*   [color.go](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go)
    
*   [compat/color.go](https://github.com/charmbracelet/lipgloss/blob/30441468/compat/color.go)
    
*   [compat/doc.go](https://github.com/charmbracelet/lipgloss/blob/30441468/compat/doc.go)
    
*   [query.go](https://github.com/charmbracelet/lipgloss/blob/30441468/query.go)
    
*   [terminal.go](https://github.com/charmbracelet/lipgloss/blob/30441468/terminal.go)
    
*   [whitespace.go](https://github.com/charmbracelet/lipgloss/blob/30441468/whitespace.go)
    
*   [whitespace\_test.go](https://github.com/charmbracelet/lipgloss/blob/30441468/whitespace_test.go)
    
*   [writer.go](https://github.com/charmbracelet/lipgloss/blob/30441468/writer.go)
    

Purpose and Scope
-----------------

This document provides an overview of Lipgloss's color management system, terminal capability detection, and adaptive color selection patterns. The system enables developers to specify colors that automatically adapt to different terminal capabilities (ranging from monochrome to 16M true colors) and background colors (light or dark).

For detailed information about color definitions and manipulation functions, see page 3.1. For adaptive color selection patterns, see page 3.2. For terminal adaptation and output writers, see page 3.3.

* * *

System Architecture
-------------------

The color and terminal detection system consists of three main subsystems:

**Sources:** [color.go1-360](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L1-L360)
 [query.go1-90](https://github.com/charmbracelet/lipgloss/blob/30441468/query.go#L1-L90)
 [terminal.go1-125](https://github.com/charmbracelet/lipgloss/blob/30441468/terminal.go#L1-L125)
 [writer.go1-161](https://github.com/charmbracelet/lipgloss/blob/30441468/writer.go#L1-L161)

* * *

Color Input System
------------------

Lipgloss provides multiple ways to specify colors, all of which implement Go's `color.Color` interface:

| Type | Description | Example |
| --- | --- | --- |
| `Color()` function | Parses hex strings (#RRGGBB, #RGB) or ANSI codes (0-255) or RGB integers | `Color("#0000ff")`, `Color("21")`, `Color(0xFF0000)` |
| `NoColor` | Represents absence of color (uses terminal default) | `NoColor{}` |
| `RGBColor` | Direct RGB specification with R, G, B uint8 fields | `RGBColor{R: 255, G: 0, B: 0}` |
| `ANSIColor` | Type alias for `ansi.IndexedColor` (0-255) | `ANSIColor(21)` |

All colors can be manipulated using functions like `Alpha()`, `Darken()`, `Lighten()`, and `Complementary()`.

**Sources:** [color.go15-162](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L15-L162)
 [color.go290-359](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L290-L359)

* * *

Terminal Capability Detection
-----------------------------

The terminal detection system queries the terminal for two critical pieces of information:

### Color Profile Detection

The color profile determines which color rendering capabilities the terminal supports:

| Profile | Capability | Colors Available | Detection Method |
| --- | --- | --- | --- |
| `colorprofile.TrueColor` | 24-bit RGB | 16,777,216 colors | `COLORTERM=truecolor` or feature detection |
| `colorprofile.ANSI256` | 8-bit indexed | 256 colors | `TERM=*256*` or feature detection |
| `colorprofile.ANSI` | 4-bit basic | 16 colors | `TERM` variable parsing |
| `colorprofile.ASCII` | No color | 0 colors (text only) | Fallback or explicit detection |

Detection is performed by `colorprofile.Detect()` from the external colorprofile package, which examines environment variables (`TERM`, `COLORTERM`) and terminal capabilities.

**Sources:** [writer.go1-14](https://github.com/charmbracelet/lipgloss/blob/30441468/writer.go#L1-L14)
 [compat/color.go15-16](https://github.com/charmbracelet/lipgloss/blob/30441468/compat/color.go#L15-L16)

### Background Color Detection

The background detection system queries whether the terminal has a light or dark background using ANSI escape sequences:

The detection process:

1.  **Put terminal in raw mode** - Prevents line buffering and echo ([query.go13-18](https://github.com/charmbracelet/lipgloss/blob/30441468/query.go#L13-L18)
    )
2.  **Send OSC 11 query** - Request background color ([terminal.go45](https://github.com/charmbracelet/lipgloss/blob/30441468/terminal.go#L45-L45)
    )
3.  **Parse response** - Extract RGB values ([terminal.go31-36](https://github.com/charmbracelet/lipgloss/blob/30441468/terminal.go#L31-L36)
    )
4.  **Calculate luminance** - Determine if color is dark ([color.go225-233](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L225-L233)
    )
5.  **Restore terminal state** - Return to normal mode ([query.go18](https://github.com/charmbracelet/lipgloss/blob/30441468/query.go#L18-L18)
    )

The `HasDarkBackground()` function defaults to `true` on error, assuming a dark background is the safer fallback for visibility.

**Sources:** [query.go12-89](https://github.com/charmbracelet/lipgloss/blob/30441468/query.go#L12-L89)
 [terminal.go14-124](https://github.com/charmbracelet/lipgloss/blob/30441468/terminal.go#L14-L124)
 [color.go214-233](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L214-L233)

* * *

Adaptive Color Selection Patterns
---------------------------------

Lipgloss provides two primary patterns for adaptive color selection:

### LightDark Pattern

The `LightDarkFunc` pattern allows runtime selection between two colors based on terminal background:

Usage example:

**Sources:** [color.go163-212](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L163-L212)

### Complete Pattern

The `CompleteFunc` pattern allows runtime selection based on terminal color profile:

Usage example:

**Sources:** [color.go235-277](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L235-L277)

### Compatibility Layer

For backward compatibility, the `compat/` package provides the legacy type-based approach with `AdaptiveColor`, `CompleteColor`, and `CompleteAdaptiveColor` types. These types use global variables for detection and implement the `color.Color` interface, automatically selecting colors based on their `RGBA()` method.

**Sources:** [compat/color.go1-78](https://github.com/charmbracelet/lipgloss/blob/30441468/compat/color.go#L1-L78)

* * *

Output Adaptation System
------------------------

The output adaptation system automatically downsamples colors to match terminal capabilities using `colorprofile.Writer`:

### Automatic Downsampling

When styled text is written to output, the `colorprofile.Writer` examines each ANSI escape sequence and converts it to match the detected terminal profile:

| Original Color | TrueColor Terminal | ANSI256 Terminal | ANSI Terminal | ASCII Terminal |
| --- | --- | --- | --- | --- |
| `#5A56E0` (true color) | `#5A56E0` | Closest 256-color | Closest 16-color | No color |
| `Color(62)` (ANSI256) | Converted to RGB | `Color(62)` | Closest 16-color | No color |
| `Color(12)` (ANSI) | Converted to RGB | Converted to 256 | `Color(12)` | No color |

The default writer is initialized at [writer.go14](https://github.com/charmbracelet/lipgloss/blob/30441468/writer.go#L14-L14)
:

### Output Functions

Lipgloss provides multiple output functions that automatically use the downsampling writer:

*   `Print()`, `Println()`, `Printf()` - Write to `os.Stdout` via default `Writer`
*   `Fprint()`, `Fprintln()`, `Fprintf()` - Write to any `io.Writer` with downsampling
*   `Sprint()`, `Sprintln()`, `Sprintf()` - Return downsampled strings

**Sources:** [writer.go1-161](https://github.com/charmbracelet/lipgloss/blob/30441468/writer.go#L1-L161)

* * *

Color Rendering Pipeline
------------------------

The complete flow from color specification to terminal output:

This pipeline ensures that colors specified once in application code automatically adapt to any terminal environment, from modern terminals with true color support to legacy terminals with limited color capabilities.

**Sources:** [color.go63-95](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L63-L95)
 [writer.go26-161](https://github.com/charmbracelet/lipgloss/blob/30441468/writer.go#L26-L161)

* * *

RGBA Interface Compatibility
----------------------------

All color types implement Go's `color.Color` interface through the `RGBA()` method, allowing lipgloss colors to interoperate with the standard `image/color` package. However, this method is marked as deprecated because:

1.  It requires converting to RGB, which may not be meaningful for ANSI color codes
2.  The conversion depends on the global renderer instance
3.  Error cases default to black (0x0, 0x0, 0x0, 0xFFFF)

The `RGBA()` implementations at [color.go37-39](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L37-L39)
 [color.go57-59](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L57-L59)
 [color.go82-85](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L82-L85)
 [color.go112-114](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L112-L114)
 [color.go145-147](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L145-L147)
 and [color.go170-172](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L170-L172)
 all follow the same pattern: they convert via `termenv.ConvertToRGB()` and return black on error.

**Sources:** [color.go30-172](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L30-L172)
 [color\_test.go90-237](https://github.com/charmbracelet/lipgloss/blob/30441468/color_test.go#L90-L237)

* * *

Usage Guidelines
----------------

### When to Use Each Color Type

**`NoColor`**: Use when you want to explicitly remove color styling, letting the terminal's default colors show through.

**`Color`**: Use for most cases where you want a specific color and don't need special adaptation. Supports both hex and ANSI formats. Automatically degrades across terminal profiles.

**`ANSIColor`**: Use when working with numeric ANSI codes for type safety, particularly when color indices are computed programmatically.

**`AdaptiveColor`**: Use when your design requires different colors for light and dark backgrounds (e.g., ensuring text remains readable). Both variants still auto-degrade across profiles.

**`CompleteColor`**: Use when you need explicit control over color degradation and want to specify exactly how the color appears on each terminal type. Prevents automatic approximations.

**`CompleteAdaptiveColor`**: Use for maximum control when you need both background awareness and profile-specific colors. Ideal for design systems with strict color requirements across all terminal types.

**Sources:** [color.go15-172](https://github.com/charmbracelet/lipgloss/blob/30441468/color.go#L15-L172)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Colors and Terminal Detection](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#colors-and-terminal-detection)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#purpose-and-scope)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#system-architecture)
    
*   [Color Input System](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#color-input-system)
    
*   [Terminal Capability Detection](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#terminal-capability-detection)
    
*   [Color Profile Detection](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#color-profile-detection)
    
*   [Background Color Detection](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#background-color-detection)
    
*   [Adaptive Color Selection Patterns](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#adaptive-color-selection-patterns)
    
*   [LightDark Pattern](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#lightdark-pattern)
    
*   [Complete Pattern](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#complete-pattern)
    
*   [Compatibility Layer](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#compatibility-layer)
    
*   [Output Adaptation System](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#output-adaptation-system)
    
*   [Automatic Downsampling](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#automatic-downsampling)
    
*   [Output Functions](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#output-functions)
    
*   [Color Rendering Pipeline](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#color-rendering-pipeline)
    
*   [RGBA Interface Compatibility](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#rgba-interface-compatibility)
    
*   [Usage Guidelines](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#usage-guidelines)
    
*   [When to Use Each Color Type](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection#when-to-use-each-color-type)
