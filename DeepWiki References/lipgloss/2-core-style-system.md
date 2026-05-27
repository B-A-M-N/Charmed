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

Core Style System
=================

Relevant source files

*   [borders.go](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go)
    
*   [borders\_test.go](https://github.com/charmbracelet/lipgloss/blob/30441468/borders_test.go)
    
*   [get.go](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go)
    
*   [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go)
    
*   [style.go](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go)
    
*   [unset.go](https://github.com/charmbracelet/lipgloss/blob/30441468/unset.go)
    

The Core Style System provides the fundamental styling primitive for Lipgloss through the `Style` struct. This struct encapsulates all styling properties (colors, dimensions, spacing, borders, text attributes) and provides methods for configuration and rendering. The style system uses an efficient property-tracking mechanism to manage which properties are explicitly set versus inherited or defaulted.

For details on configuring styles with setter methods, see [Style Configuration](https://deepwiki.com/charmbracelet/lipgloss/2.1-style-configuration)
. For text formatting attributes like bold and italic, see [Text Attributes and Styling](https://deepwiki.com/charmbracelet/lipgloss/2.2-text-attributes-and-styling)
. For the multi-stage rendering process, see [Rendering Pipeline](https://deepwiki.com/charmbracelet/lipgloss/2.3-rendering-pipeline)
.

The Style Struct
----------------

The `Style` struct is the central data structure in Lipgloss, defined in [style.go142-195](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L142-L195)
 It contains all styling information needed to render formatted text.

**Sources:** [style.go142-195](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L142-L195)

The struct uses two levels of storage:

1.  **Property tracking bitfield** (`props`): Tracks which properties have been explicitly set
2.  **Value fields**: Store the actual property values

This design allows distinguishing between "not set" and "set to zero/false", which is critical for style inheritance.

Property Key System
-------------------

Properties are identified using the `propKey` type, defined as bit positions in an `int64` bitfield at [style.go18-91](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L18-L91)

**Sources:** [style.go18-91](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L18-L91)

The `props` type at [style.go94-109](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L94-L109)
 provides methods for manipulating this bitfield:

| Method | Purpose | Location |
| --- | --- | --- |
| `set(k propKey)` | Sets a property bit to 1 | [style.go97-99](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L97-L99) |
| `unset(k propKey)` | Clears a property bit to 0 | [style.go102-104](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L102-L104) |
| `has(k propKey)` | Checks if a property bit is set | [style.go107-109](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L107-L109) |

Style Construction and Configuration
------------------------------------

**Sources:** [style.go137-139](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L137-L139)
 [set.go9-108](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L9-L108)
 [set.go195-198](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L195-L198)
 [set.go272-275](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L272-L275)
 [set.go341-352](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L341-L352)

The `NewStyle()` function at [style.go137-139](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L137-L139)
 returns an empty `Style{}` struct. All configuration methods follow a fluent API pattern, returning a modified copy of the style. The internal `set()` method at [set.go9-108](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L9-L108)
 handles storing values and updating the property bitfield.

Property Access Methods
-----------------------

Each property has three associated methods defined across three files:

**Sources:** [set.go9-108](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L9-L108)
 [get.go487-643](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L487-L643)
 [unset.go4-6](https://github.com/charmbracelet/lipgloss/blob/30441468/unset.go#L4-L6)

Getters use internal helper methods like `getAsBool()`, `getAsColor()`, `getAsInt()` at [get.go504-626](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L504-L626)
 which check the property bitfield before returning values. This allows distinguishing between unset properties (return defaults) and explicitly set properties.

Style Inheritance
-----------------

The `Inherit()` method at [style.go238-265](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L238-L265)
 implements style composition by overlaying one style onto another:

**Sources:** [style.go238-265](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L238-L265)
 [set.go110-185](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L110-L185)

Key inheritance rules:

*   **Margins and padding are never inherited** ([style.go245-250](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L245-L250)
    )
*   **Existing values are preserved** ([style.go258-260](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L258-L260)
    )
*   **Background color inheritance affects margins** ([style.go252-256](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L252-L256)
    )

The `setFrom()` method at [set.go110-185](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L110-L185)
 handles copying individual properties from one style to another.

Internal Property Accessors
---------------------------

The getter methods use type-specific internal accessors to safely retrieve property values:

| Accessor Method | Return Type | Checks `isSet()` | Default Value | Location |
| --- | --- | --- | --- | --- |
| `getAsBool(k, default)` | `bool` | Yes | Passed as parameter | [get.go504-509](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L504-L509) |
| `getAsColor(k)` | `color.Color` | Yes | `noColor` | [get.go524-562](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L524-L562) |
| `getAsInt(k)` | `int` | Yes | `0` | [get.go564-599](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L564-L599) |
| `getAsPosition(k)` | `Position` | Yes | `Position(0)` | [get.go601-612](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L601-L612) |
| `getAsRune(k)` | `rune` | Yes | `0` | [get.go491-502](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L491-L502) |
| `getAsColors(k)` | `[]color.Color` | Yes | `nil` | [get.go511-522](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L511-L522) |
| `getAsTransform(k)` | `func(string)string` | Yes | `nil` | [get.go621-626](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L621-L626) |
| `getBorderStyle()` | `Border` | Yes | `noBorder` | [get.go614-619](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L614-L619) |
| `isSet(k)` | `bool` | N/A | N/A | [get.go487-489](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L487-L489) |

**Sources:** [get.go487-626](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L487-L626)

The `isSet()` method at [get.go487-489](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L487-L489)
 checks the property bitfield using `s.props.has(k)` to determine if a property has been explicitly configured.

Value Storage and Rendering
---------------------------

The `Style` struct can store an underlying string value for convenience:

**Sources:** [style.go208-211](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L208-L211)
 [style.go221-223](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L221-L223)
 [style.go268-271](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L268-L271)

*   `SetString()` stores the value in `s.value` for later rendering
*   `String()` implements `fmt.Stringer` by calling `Render()` with the stored value
*   `Render()` can accept strings as arguments OR use the stored `s.value`

When `Render()` is called with arguments at [style.go269-271](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L269-L271)
 any stored value is prepended to the argument list.

Core Rendering Overview
-----------------------

The `Render()` method at [style.go268-526](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L268-L526)
 is the central method that converts a style and text into formatted ANSI output. The rendering process involves multiple stages:

**Sources:** [style.go268-526](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L268-L526)

For detailed information about the rendering pipeline stages, see [Rendering Pipeline](https://deepwiki.com/charmbracelet/lipgloss/2.3-rendering-pipeline)
.

Style Copying
-------------

The `Copy()` method at [style.go229-231](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L229-L231)
 is deprecated because Go structs are copied by value automatically. All style methods return a new `Style` by value, making them immutable from the caller's perspective:

**Sources:** [style.go229-231](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L229-L231)

This value-based copying ensures that modifying a derived style never affects the original style.

Integration Points
------------------

The `Style` system integrates with other Lipgloss subsystems:

**Sources:** [style.go142-195](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L142-L195)
 [borders.go327-495](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L327-L495)
 color system integration visible throughout rendering

*   **Color System** ([#3](https://deepwiki.com/charmbracelet/lipgloss/3-colors-and-terminal-detection)
    ): Style stores `color.Color` fields for foreground, background, and underline colors
*   **Border System** ([#4](https://deepwiki.com/charmbracelet/lipgloss/4-borders)
    ): Style contains border style and rendering calls `applyBorder()` at [borders.go327-495](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L327-L495)
    
*   **Layout System** ([#5](https://deepwiki.com/charmbracelet/lipgloss/5-layout-and-composition)
    ): Layout functions consume and produce styled strings
*   **ANSI Output**: Style rendering creates `ansi.Style` objects from `github.com/charmbracelet/x/ansi` to generate terminal escape sequences

Constants and Special Values
----------------------------

Several special constants affect style behavior:

| Constant | Value | Purpose | Location |
| --- | --- | --- | --- |
| `NBSP` | `'\u00A0'` | Non-breaking space rune | [style.go13](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L13-L13) |
| `tabWidthDefault` | `4` | Default tab width | [style.go14](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L14-L14) |
| `NoTabConversion` | `-1` | Disable tab-to-space conversion | [set.go770](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L770-L770) |

**Sources:** [style.go11-15](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L11-L15)
 [set.go770](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L770-L770)

The `NoTabConversion` constant can be passed to `TabWidth()` to preserve literal tab characters in the output.

Property Categories Summary
---------------------------

The 50+ properties tracked by the Style system fall into these categories:

| Category | Properties | Example Methods |
| --- | --- | --- |
| **Boolean Attributes** | bold, italic, strikethrough, reverse, blink, faint | `Bold()`, `Italic()` |
| **Text Formatting** | underline style, underline color | `UnderlineStyle()`, `UnderlineColor()` |
| **Colors** | foreground, background, margin background | `Foreground()`, `Background()` |
| **Dimensions** | width, height, max width, max height | `Width()`, `Height()`, `MaxWidth()` |
| **Alignment** | horizontal, vertical | `AlignHorizontal()`, `AlignVertical()` |
| **Padding** | top, right, bottom, left, character | `Padding()`, `PaddingChar()` |
| **Margins** | top, right, bottom, left, character, background | `Margin()`, `MarginChar()` |
| **Borders** | style, sides, colors, blend | `Border()`, `BorderForeground()` |
| **Behavior** | inline, tab width, whitespace coloring | `Inline()`, `TabWidth()` |
| **Transformations** | transform function, hyperlink | `Transform()`, `Hyperlink()` |

**Sources:** [style.go21-91](https://github.com/charmbracelet/lipgloss/blob/30441468/style.go#L21-L91)
 [set.go194-826](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go#L194-L826)
 [get.go11-484](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L11-L484)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core Style System](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#core-style-system)
    
*   [The Style Struct](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#the-style-struct)
    
*   [Property Key System](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#property-key-system)
    
*   [Style Construction and Configuration](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#style-construction-and-configuration)
    
*   [Property Access Methods](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#property-access-methods)
    
*   [Style Inheritance](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#style-inheritance)
    
*   [Internal Property Accessors](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#internal-property-accessors)
    
*   [Value Storage and Rendering](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#value-storage-and-rendering)
    
*   [Core Rendering Overview](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#core-rendering-overview)
    
*   [Style Copying](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#style-copying)
    
*   [Integration Points](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#integration-points)
    
*   [Constants and Special Values](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#constants-and-special-values)
    
*   [Property Categories Summary](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system#property-categories-summary)
