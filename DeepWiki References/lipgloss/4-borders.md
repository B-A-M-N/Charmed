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

Borders
=======

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/lipgloss/blob/30441468/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1)
    
*   [borders.go](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go)
    
*   [borders\_test.go](https://github.com/charmbracelet/lipgloss/blob/30441468/borders_test.go)
    
*   [get.go](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go)
    
*   [size.go](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go)
    

The border system provides borders around styled content through the `Border` struct and the `Style.applyBorder()` rendering function. The system includes ten predefined border styles, supports custom `Border` creation, per-side enablement (top/right/bottom/left), per-side color configuration, and gradient border effects.

Border rendering occurs in `Style.Render()` after padding and alignment, before margins. The `Style` struct stores border configuration via property keys defined in [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go)
 and retrieves them via methods in [get.go233-408](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L233-L408)
 For the rendering pipeline context, see page 2.3.

Border Structure
----------------

The `Border` struct at [borders.go14-30](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L14-L30)
 defines border characters as string fields. Each field accepts Unicode strings and supports multi-character repeating patterns.

**Border Struct Definition**

**Border Field Usage**

| Field | Rendering Behavior | Line Reference |
| --- | --- | --- |
| `Top`, `Bottom` | Horizontal edges, characters repeat to fill width | [borders.go497-524](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L497-L524) |
| `Left`, `Right` | Vertical edges, characters cycle per line | [borders.go444-477](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L444-L477) |
| `TopLeft`, `TopRight`, `BottomLeft`, `BottomRight` | Corners, limited to first rune via `getFirstRuneAsString()` | [borders.go407-410](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L407-L410) |
| `MiddleLeft`, `MiddleRight`, `Middle`, `MiddleTop`, `MiddleBottom` | Table junctions, used by table package | Unused in `applyBorder()` |

The `Border` methods `GetTopSize()`, `GetRightSize()`, `GetBottomSize()`, and `GetLeftSize()` at [borders.go32-58](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L32-L58)
 call `getBorderEdgeWidth()` at [borders.go60-65](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L60-L65)
 which uses `maxRuneWidth()` at [borders.go564-579](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L564-L579)
 to calculate maximum character display width.

Sources: [borders.go14-30](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L14-L30)
 [borders.go32-65](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L32-L65)
 [borders.go407-410](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L407-L410)
 [borders.go444-524](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L444-L524)
 [borders.go564-579](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L564-L579)

Predefined Border Styles
------------------------

Ten predefined `Border` instances are defined as package-level variables at [borders.go67-219](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L67-L219)
 Each has a corresponding constructor function at [borders.go221-279](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L221-L279)

**Border Constructor Functions**

**Border Style Characters**

| Function | Primary Characters | Line Definition | Table Support |
| --- | --- | --- | --- |
| `NormalBorder()` | `┌─┐│└┘` | [borders.go70-84](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L70-L84) | Yes (includes Middle\* fields) |
| `RoundedBorder()` | `╭─╮│╰╯` | [borders.go86-100](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L86-L100) | Yes (includes Middle\* fields) |
| `BlockBorder()` | `█` (all fields) | [borders.go102-116](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L102-L116) | Yes (all fields `█`) |
| `ThickBorder()` | `┏━┓┃┗┛` | [borders.go140-154](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L140-L154) | Yes (includes Middle\* fields) |
| `DoubleBorder()` | `╔═╗║╚╝` | [borders.go156-170](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L156-L170) | Yes (includes Middle\* fields) |
| `HiddenBorder()` | (all fields) | [borders.go172-186](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L172-L186) | Yes (all fields space) |
| `OuterHalfBlockBorder()` | `▛▀▜▌▙▄▟▐` | [borders.go118-127](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L118-L127) | No (no Middle\* fields) |
| `InnerHalfBlockBorder()` | `▗▄▖▐▝▀▘▌` | [borders.go129-138](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L129-L138) | No (no Middle\* fields) |
| `MarkdownBorder()` | `\|-` | [borders.go188-202](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L188-L202) | Yes (all fields defined) |
| `ASCIIBorder()` | `+-\|` | [borders.go204-218](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L204-L218) | Yes (all fields defined) |

The `noBorder` variable at [borders.go68](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L68-L68)
 is an empty `Border{}` used internally by `getBorderStyle()` at [get.go614-619](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L614-L619)
 to indicate no border is set.

Sources: [borders.go67-219](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L67-L219)
 [borders.go221-279](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L221-L279)
 [get.go614-619](https://github.com/charmbracelet/lipgloss/blob/30441468/get.go#L614-L619)

Border Rendering Pipeline
-------------------------

The `Style.applyBorder()` method at [borders.go327-495](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L327-L495)
 renders borders after padding/alignment in the `Style.Render()` pipeline. The function handles border character rendering, color application, and gradient blending.

**applyBorder() Flow**

**applyBorder() Stages**

| Stage | Lines | Operations |
| --- | --- | --- |
| **Configuration** | [borders.go328-334](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L328-L334) | `getBorderStyle()`, `getAsBool()` for each side |
| **Implicit Border** | [borders.go336-343](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L336-L343) | `isBorderStyleSetWithoutSides()` enables all sides if border set without explicit sides |
| **Early Exit** | [borders.go345-348](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L345-L348) | Returns `str` unchanged if `noBorder` or all sides disabled |
| **Parsing** | [borders.go350](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L350-L350) | `getLines(str)` splits into lines, calculates width |
| **Width Calculation** | [borders.go352-364](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L352-L364) | Adds `maxRuneWidth(border.Left)` and `maxRuneWidth(border.Right)` to width |
| **Default Fill** | [borders.go366-404](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L366-L404) | Fills empty border characters with `" "`, fills/removes corners based on side enablement |
| **Corner Limit** | [borders.go407-410](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L407-L410) | `getFirstRuneAsString()` limits corners to first rune |
| **Color Setup** | [borders.go412-429](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L412-L429) | `getAsColors(borderForegroundBlendKey)` or `getAsColor()` for each side |
| **Top Render** | [borders.go433-442](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L433-L442) | `renderHorizontalEdge()`, then `styleBorderBlend()` or `styleBorder()` |
| **Side Render** | [borders.go444-481](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L444-L481) | Cycles through `border.Left`/`border.Right` runes, applies colors per line |
| **Bottom Render** | [borders.go483-492](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L483-L492) | `renderHorizontalEdge()`, then `styleBorderBlend()` or `styleBorder()` |

**Helper Functions**

| Function | Lines | Purpose |
| --- | --- | --- |
| `renderHorizontalEdge()` | [borders.go497-524](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L497-L524) | Constructs top/bottom by cycling `middle` chars, uses `ansi.StringWidth()` |
| `styleBorder()` | [borders.go526-539](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L526-L539) | Applies single `fg`/`bg` color via `ansi.Style` |
| `styleBorderBlend()` | [borders.go541-562](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L541-L562) | Applies gradient `fg[]` colors using `uniseg.NewGraphemes()` |
| `maxRuneWidth()` | [borders.go564-579](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L564-L579) | Calculates max display width via `displaywidth` |
| `getFirstRuneAsString()` | [borders.go581-587](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L581-L587) | Extracts first rune via `utf8.DecodeRuneInString()` |

**Character Cycling**

*   `border.Top`, `border.Bottom`: Cycle horizontally in `renderHorizontalEdge()` at [borders.go506-520](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L506-L520)
    
*   `border.Left`: Cycles vertically at [borders.go444-458](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L444-L458)
     wraps via `leftIndex = 0`
*   `border.Right`: Cycles vertically at [borders.go466-477](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L466-L477)
     wraps via `rightIndex = 0`

Sources: [borders.go327-587](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L327-L587)

Border Configuration
--------------------

Borders are configured via `Style` setter methods. The `Style` struct stores border settings as property keys checked by `applyBorder()`.

**Border Configuration Methods**

**Border Setter Methods**

| Method Signature | Property Keys Set | Reference |
| --- | --- | --- |
| `BorderStyle(b Border)` | `borderStyleKey` | [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go) |
| `Border(b Border, sides ...bool)` | `borderStyleKey`, side keys if `len(sides) > 0` | [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go) |
| `BorderTop/Right/Bottom/Left(v bool)` | Corresponding `borderTopKey`, etc. | [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go) |
| `BorderForeground(c color.Color)` | All four `border*ForegroundKey` | [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go) |
| `BorderTopForeground(c color.Color)` | `borderTopForegroundKey` | [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go) |
| `BorderBackground(c color.Color)` | All four `border*BackgroundKey` | [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go) |
| `BorderTopBackground(c color.Color)` | `borderTopBackgroundKey` | [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go) |
| `BorderForegroundBlend(colors ...color.Color)` | `borderForegroundBlendKey` | [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go) |
| `BorderForegroundBlendOffset(offset int)` | `borderForegroundBlendOffsetKey` | [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go) |

**Custom Border Creation**

Custom `Border` instances are created via struct literal:

| Field Pattern | Rendering Behavior | Applied By |
| --- | --- | --- |
| Single character | Renders once (corners) or repeats (edges) | `renderHorizontalEdge()`, side loops |
| Multi-character `Top`/`Bottom` | Cycles horizontally at [borders.go506-520](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L506-L520) | `renderHorizontalEdge()` |
| Multi-character `Left`/`Right` | Cycles vertically at [borders.go444-477](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L444-L477) | Side rendering loop |
| Empty string | Filled with `" "` at [borders.go352-379](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L352-L379) | `applyBorder()` |
| Unicode | Width calculated via `maxRuneWidth()` at [borders.go564-579](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L564-L579) | Width calculation stage |

Corners are limited to first rune at [borders.go407-410](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L407-L410)
 via `getFirstRuneAsString()`. Corners are removed if adjacent sides disabled at [borders.go381-404](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L381-L404)

Sources: [borders.go14-30](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L14-L30)
 [borders.go327-587](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L327-L587)
 [set.go](https://github.com/charmbracelet/lipgloss/blob/30441468/set.go)

Unicode and Character Handling
------------------------------

The border system includes sophisticated Unicode handling for international terminal support and variable-width characters.

**Unicode Processing Features**

The border system handles complex Unicode scenarios:

*   **Grapheme Cluster Support**: Uses `uniseg` library to properly handle multi-byte Unicode sequences
*   **Width Calculation**: Determines display width of border characters for proper alignment
*   **Character Cycling**: Supports repeating patterns in border characters for decorative effects
*   **Corner Limitation**: Restricts corners to single runes for consistent terminal rendering

The `maxRuneWidth()` function iterates through grapheme clusters to find the widest character, ensuring proper spacing calculations for international content.

Sources: [borders.go57-65](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L57-L65)
 [borders.go469-490](https://github.com/charmbracelet/lipgloss/blob/30441468/borders.go#L469-L490)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Borders](https://deepwiki.com/charmbracelet/lipgloss/4-borders#borders)
    
*   [Border Structure](https://deepwiki.com/charmbracelet/lipgloss/4-borders#border-structure)
    
*   [Predefined Border Styles](https://deepwiki.com/charmbracelet/lipgloss/4-borders#predefined-border-styles)
    
*   [Border Rendering Pipeline](https://deepwiki.com/charmbracelet/lipgloss/4-borders#border-rendering-pipeline)
    
*   [Border Configuration](https://deepwiki.com/charmbracelet/lipgloss/4-borders#border-configuration)
    
*   [Unicode and Character Handling](https://deepwiki.com/charmbracelet/lipgloss/4-borders#unicode-and-character-handling)
