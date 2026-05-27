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

Structured Components
=====================

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/lipgloss/blob/30441468/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1)
    
*   [size.go](https://github.com/charmbracelet/lipgloss/blob/30441468/size.go)
    
*   [table/resizing.go](https://github.com/charmbracelet/lipgloss/blob/30441468/table/resizing.go)
    
*   [table/rows.go](https://github.com/charmbracelet/lipgloss/blob/30441468/table/rows.go)
    
*   [table/table.go](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go)
    
*   [table/table\_test.go](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go)
    
*   [table/testdata/TestBorderedCells.golden](https://github.com/charmbracelet/lipgloss/blob/30441468/table/testdata/TestBorderedCells.golden)
    
*   [table/testdata/TestTableNoStyleFunc.golden](https://github.com/charmbracelet/lipgloss/blob/30441468/table/testdata/TestTableNoStyleFunc.golden)
    
*   [table/util.go](https://github.com/charmbracelet/lipgloss/blob/30441468/table/util.go)
    

Purpose and Scope
-----------------

Structured components are specialized rendering packages within lipgloss that provide high-level abstractions for rendering complex data structures in terminal interfaces. This page provides an overview of the three structured component packages: **table**, **tree**, and **list**. Each component builds upon the core lipgloss Style system to provide domain-specific rendering capabilities with configurable styling, borders, and layout options.

For details on the core Style system that underpins these components, see [Core Style System](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system)
. For specific implementation details of each component, see [Tables](https://deepwiki.com/charmbracelet/lipgloss/5.1-measurement-utilities)
, [Trees](https://deepwiki.com/charmbracelet/lipgloss/5.2-joining-and-alignment)
, and [Lists](https://deepwiki.com/charmbracelet/lipgloss/5.3-placement)
.

Sources: [README.md409-719](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L409-L719)

Component Architecture Overview
-------------------------------

The three structured components share a common architectural pattern: they accept data in structured formats, apply styles through function-based callbacks, and render the results to strings using the core lipgloss rendering primitives.

**Architecture Diagram: Structured Components Integration**

All three components follow a similar flow: they accept structured data and style configuration from user code, process it through component-specific rendering logic that leverages core lipgloss primitives (Style, Join, Size, Border), and output ANSI-formatted strings.

Sources: [table/table.go1-504](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L1-L504)
 [tree/tree.go1-370](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L1-L370)
 [tree/renderer.go1-141](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L1-L141)
 [README.md409-719](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L409-L719)

Common Patterns
---------------

### Fluent API Configuration

All structured components expose a fluent API that allows chaining configuration methods. This pattern is consistent with the core lipgloss Style API.

| Component | Configuration Methods | Example |
| --- | --- | --- |
| Table | `Border()`, `StyleFunc()`, `Headers()`, `Row()`, `Width()`, `Height()` | `table.New().Border(lipgloss.NormalBorder()).Headers("A", "B")` |
| Tree | `Root()`, `Child()`, `Enumerator()`, `ItemStyle()`, `RootStyle()` | `tree.Root("root").Child("child1", "child2")` |
| List | `New()`, `Item()`, `Enumerator()`, `ItemStyle()` | `list.New("A", "B", "C").Enumerator(list.Roman)` |

Sources: [table/table.go74-227](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L74-L227)
 [tree/tree.go335-370](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L335-L370)
 [README.md409-719](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L409-L719)

### Style Function Callbacks

Each component uses style functions to enable conditional styling based on position or context. This allows users to implement alternating row colors, highlighting, or other dynamic styling patterns.

**Table StyleFunc**: Takes row and column indices, returns a `lipgloss.Style`.

    type StyleFunc func(row, col int) lipgloss.Style
    

[table/table.go15-37](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L15-L37)

**Tree StyleFunc**: Takes children collection and index, returns a `lipgloss.Style`.

    type StyleFunc func(children Children, i int) lipgloss.Style
    

[tree/renderer.go9-10](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L9-L10)

Sources: [table/table.go15-37](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L15-L37)
 [tree/renderer.go9-17](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L9-L17)
 [table/table\_test.go13-22](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go#L13-L22)

### Rendering Pipeline

**Rendering Pipeline: Common Flow Across Components**

Sources: [table/table.go229-294](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L229-L294)
 [tree/renderer.go41-140](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L41-L140)

Component Summary
-----------------

### Tables

The `table` package provides rendering for tabular data with support for headers, borders, cell styling, and dynamic sizing. Tables can adapt their column widths to fit specified dimensions and support filtering through a `Data` interface abstraction.

**Key Types:**

*   `Table` struct: Main configuration and rendering struct [table/table.go44-72](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L44-L72)
    
*   `StyleFunc`: Callback for cell-specific styling [table/table.go37](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L37-L37)
    
*   `Data` interface: Abstraction for table data sources
*   `HeaderRow` constant: Sentinel value (-1) for header row identification [table/table.go13](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L13-L13)
    

**Core Methods:**

*   `New()`: Creates a new table with defaults [table/table.go78-91](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L78-L91)
    
*   `String()`: Renders the table to a string [table/table.go229-294](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L229-L294)
    
*   `Headers()`, `Row()`, `Rows()`: Data configuration [table/table.go139-143](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L139-L143)
     [table/table.go130-137](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L130-L137)
     [table/table.go119-128](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L119-L128)
    
*   `Border()`, `BorderTop()`, `BorderBottom()`, etc.: Border configuration [table/table.go145-191](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L145-L191)
    
*   `Width()`, `Height()`: Dimension control [table/table.go202-212](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L202-L212)
    

Sources: [table/table.go1-504](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L1-L504)
 [README.md409-510](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L409-L510)

### Trees

The `tree` package renders hierarchical data structures with configurable enumerators (branch indicators) and indenters. Trees support nested structures, conditional styling, and custom rendering for each node.

**Key Types:**

*   `Node` interface: Abstraction for tree nodes [tree/tree.go35-43](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L35-L43)
    
*   `Tree` struct: Implements `Node` with children support [tree/tree.go94-103](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L94-L103)
    
*   `Leaf` struct: Implements `Node` without children [tree/tree.go45-49](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L45-L49)
    
*   `Children` interface: Collection of nodes
*   `Enumerator`: Function type for branch indicators [tree/tree.go](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go)
    
*   `Indenter`: Function type for indentation [tree/tree.go](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go)
    

**Core Methods:**

*   `New()`, `Root()`: Tree construction [tree/tree.go364-369](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L364-L369)
     [tree/tree.go342-345](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L342-L345)
    
*   `Child()`: Add children [tree/tree.go167-203](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L167-L203)
    
*   `String()`: Render the tree [tree/tree.go151-153](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L151-L153)
    
*   `Enumerator()`, `Indenter()`: Customize rendering [tree/tree.go295-324](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L295-L324)
    
*   `ItemStyle()`, `RootStyle()`, `EnumeratorStyle()`: Apply styles [tree/tree.go263-270](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L263-L270)
     [tree/tree.go257-261](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L257-L261)
     [tree/tree.go232-237](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L232-L237)
    

Sources: [tree/tree.go1-370](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L1-L370)
 [tree/renderer.go1-141](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L1-L141)
 [README.md619-719](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L619-L719)

### Lists

The `list` package provides enumerated list rendering with support for nesting, multiple enumeration styles (bullets, numbers, roman numerals), and custom enumerators. Lists can be built incrementally or from existing data.

**Key Features:**

*   Multiple predefined enumerators: Arabic, Alphabet, Roman, Bullet, Tree
*   Nesting support for hierarchical lists
*   Custom enumerator functions for domain-specific numbering
*   Style configuration for enumerators and items

Note: List source files are not included in this documentation set, but the component is described in [README.md512-617](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L512-L617)

Sources: [README.md512-617](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L512-L617)

Integration with Core Style System
----------------------------------

**Diagram: Component Integration with Core Types**

Each structured component maintains references to `lipgloss.Style` instances and `lipgloss.Border` definitions, which are used during the rendering phase. The components call core lipgloss functions like `JoinHorizontal`, `JoinVertical`, and `Style.Render()` to assemble the final output.

### Table Integration

Tables store a `borderStyle` field [table/table.go57](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L57-L57)
 and accept a `border` field of type `lipgloss.Border` [table/table.go47](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L47-L47)
 The `StyleFunc` callback returns `lipgloss.Style` instances that are applied to individual cells during rendering [table/table.go106-111](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L106-L111)

During rendering, tables use:

*   `lipgloss.JoinHorizontal()` to assemble cells into rows [table/table.go479](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L479-L479)
    
*   `Style.Render()` to apply cell styles [table/table.go363-368](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L363-L368)
     [table/table.go457-463](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L457-L463)
    
*   `lipgloss.NewStyle().MaxHeight().MaxWidth().Render()` for final output [table/table.go290-293](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L290-L293)
    
*   `lipgloss.Width()` and `lipgloss.Height()` for dimension calculations [table/table.go](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go)
    

Sources: [table/table.go44-72](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L44-L72)
 [table/table.go229-294](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L229-L294)
 [table/table.go432-493](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L432-L493)

### Tree Integration

Trees use the `renderer` struct [tree/tree.go101-102](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L101-L102)
 which contains a `Style` struct with style functions [tree/renderer.go12-17](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L12-L17)
 and references to `Enumerator` and `Indenter` functions. The rendering process [tree/renderer.go41-140](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L41-L140)
 extensively uses:

*   `lipgloss.JoinHorizontal()` to combine prefixes and items [tree/renderer.go110-116](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L110-L116)
    
*   `lipgloss.JoinVertical()` to handle multiline nodes [tree/renderer.go94-106](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L94-L106)
    
*   `Style.Render()` to apply styles to node values, enumerators, and indentation [tree/renderer.go54](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L54-L54)
     [tree/renderer.go68](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L68-L68)
     [tree/renderer.go82-87](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L82-L87)
    
*   `lipgloss.Width()` and `lipgloss.Height()` for alignment calculations [tree/renderer.go69](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L69-L69)
     [tree/renderer.go93-106](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L93-L106)
    

Sources: [tree/tree.go94-103](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L94-L103)
 [tree/renderer.go1-141](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L1-L141)

Data Abstraction and Filtering
------------------------------

### Table Data Interface

Tables support dynamic data sources through the `Data` interface pattern. While the interface definition is not shown in the provided sources, usage in tests demonstrates a `StringData` type [table/table\_test.go556-561](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go#L556-L561)
 and a `Filter` type [table/table\_test.go563-571](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go#L563-L571)
 that wraps data with filtering logic.

The `Table.Data()` method accepts any type implementing the `Data` interface [table/table.go113-117](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L113-L117)
 allowing for:

*   Static row storage with `StringData`
*   Filtered views with `NewFilter()`
*   Custom data sources with specific row/column access patterns

Example filtering pattern:

    filter := NewFilter(data).Filter(func(row int) bool {
        return data.At(row, 0) != "French"
    })
    table := New().Data(filter).Headers("LANGUAGE", "FORMAL", "INFORMAL")
    

Sources: [table/table.go113-117](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L113-L117)
 [table/table\_test.go555-574](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go#L555-L574)

### Tree Children Interface

Trees use a `Children` interface [tree/tree.go39](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L39-L39)
 to represent collections of nodes. The `NodeChildren` type implements this interface and provides methods like `At()`, `Length()`, `Remove()`, and `Append()` [tree/tree.go171-178](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L171-L178)

This abstraction allows:

*   Uniform access to child nodes regardless of type (Tree or Leaf)
*   Filtering by hiding nodes [tree/tree.go40-41](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L40-L41)
    
*   Dynamic modification of tree structure [tree/tree.go167-203](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L167-L203)
    

Sources: [tree/tree.go35-43](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L35-L43)
 [tree/tree.go167-203](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L167-L203)

Rendering Features Comparison
-----------------------------

| Feature | Table | Tree | List |
| --- | --- | --- | --- |
| Borders | Full control (top, bottom, left, right, header, column, row) | No explicit borders | No explicit borders |
| Styling | Cell-level via StyleFunc(row, col) | Node-level via ItemStyleFunc, RootStyle, EnumeratorStyle | Item-level via ItemStyle, EnumeratorStyle |
| Nesting | No  | Yes (recursive children) | Yes (nested lists) |
| Dynamic Sizing | Width and Height with auto-fit | No explicit sizing | No explicit sizing |
| Data Filtering | Via Data interface | Via Hidden() nodes | Not shown in sources |
| Wrapping | Optional text wrapping [table/table.go223-227](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L223-L227) | Handles multiline values | Handles multiline values |
| Offset/Pagination | Offset() method [table/table.go214-221](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L214-L221) | Offset() method [tree/tree.go120-138](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L120-L138) | Not shown |
| Custom Enumerators | N/A | Enumerator function | Enumerator function |

Sources: [table/table.go44-72](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L44-L72)
 [tree/tree.go94-103](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L94-L103)
 [README.md409-719](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L409-L719)

Usage Patterns
--------------

### Creating and Configuring Components

All components follow a builder pattern with fluent APIs:

**Table Example:**

    table.New().
        Border(lipgloss.NormalBorder()).
        BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
        StyleFunc(func(row, col int) lipgloss.Style {
            if row == table.HeaderRow {
                return headerStyle
            }
            return cellStyle
        }).
        Headers("LANGUAGE", "FORMAL", "INFORMAL").
        Row("Chinese", "Nǐn hǎo", "Nǐ hǎo")
    

[README.md432-461](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L432-L461)

**Tree Example:**

    tree.Root(".").
        Child(
            "macOS",
            tree.New().Root("Linux").Child("NixOS", "Arch Linux"),
            tree.New().Root("BSD").Child("FreeBSD", "OpenBSD"),
        ).
        Enumerator(tree.RoundedEnumerator).
        ItemStyle(itemStyle)
    

[README.md647-663](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L647-L663)

### Conditional Styling

Both tables and trees support conditional styling through style functions that receive context about the item being rendered:

**Table Conditional Styling:**

    StyleFunc(func(row, col int) lipgloss.Style {
        switch {
        case row == table.HeaderRow:
            return headerStyle
        case row%2 == 0:
            return evenRowStyle
        default:
            return oddRowStyle
        }
    })
    

[table/table\_test.go53-61](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go#L53-L61)

**Tree Conditional Styling:**

    EnumeratorStyleFunc(func(_ tree.Children, i int) lipgloss.Style {
        if selected == i {
            return lipgloss.NewStyle().Foreground(highlightColor)
        }
        return lipgloss.NewStyle().Foreground(dimColor)
    })
    

[tree/tree.go242-248](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L242-L248)

Sources: [table/table\_test.go13-22](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go#L13-L22)
 [table/table\_test.go38-70](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go#L38-L70)
 [tree/tree.go229-287](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/tree.go#L229-L287)

### Incremental Construction

Components can be built incrementally for dynamic scenarios:

**Table Incremental:**

    t := table.New()
    t.Headers("A", "B", "C")
    for _, row := range dataSource {
        t.Row(row...)
    }
    

[table/table.go130-137](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L130-L137)

**Tree Incremental:**

    t := tree.New()
    for i := 0; i < count; i++ {
        t.Child("Item " + strconv.Itoa(i))
    }
    

[README.md710-717](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L710-L717)

Sources: [table/table.go119-137](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L119-L137)
 [README.md609-617](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L609-L617)
 [README.md710-717](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L710-L717)

Dimension and Layout Control
----------------------------

### Table Dimensions

Tables provide explicit width and height control [table/table.go202-212](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L202-L212)
:

*   `Width(w int)`: Sets table width, triggering automatic column resizing to fit [table/table.go199-205](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L199-L205)
    
*   `Height(h int)`: Limits visible rows, enabling pagination [table/table.go207-212](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L207-L212)
    
*   Internal `resize()` method: Calculates column widths and row heights based on content and constraints

The width calculation supports:

*   Expansion: Distributing extra space across columns
*   Contraction: Smart cropping to fit smaller widths [table/table\_test.go435-452](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go#L435-L452)
    
*   Content wrapping: Optional text wrapping within cells [table/table.go223-227](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L223-L227)
    

Sources: [table/table.go199-227](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table.go#L199-L227)
 [table/table\_test.go412-553](https://github.com/charmbracelet/lipgloss/blob/30441468/table/table_test.go#L412-L553)

### Tree and List Dimensions

Trees and lists do not impose explicit dimension constraints. They calculate their dimensions based on content and styling:

*   Trees use `lipgloss.Width()` and `lipgloss.Height()` to measure rendered items [tree/renderer.go69](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L69-L69)
     [tree/renderer.ro93-106](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.ro#L93-L106)
    
*   Layout is controlled through enumerators and indenters rather than explicit sizing

Sources: [tree/renderer.go41-140](https://github.com/charmbracelet/lipgloss/blob/30441468/tree/renderer.go#L41-L140)

Next Steps
----------

For detailed information on each component:

*   **Tables**: See [Tables](https://deepwiki.com/charmbracelet/lipgloss/5.1-measurement-utilities)
     for table-specific features including Data interface, filtering, border customization, and dimension management
*   **Trees**: See [Trees](https://deepwiki.com/charmbracelet/lipgloss/5.2-joining-and-alignment)
     for tree-specific features including Node interface, enumerators, indenters, and hierarchical rendering
*   **Lists**: See [Lists](https://deepwiki.com/charmbracelet/lipgloss/5.3-placement)
     for list-specific features including enumeration styles, nesting, and custom enumerators

For information on the underlying style system, see [Core Style System](https://deepwiki.com/charmbracelet/lipgloss/2-core-style-system)
.

Sources: [README.md409-719](https://github.com/charmbracelet/lipgloss/blob/30441468/README.md?plain=1#L409-L719)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Structured Components](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#structured-components)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#purpose-and-scope)
    
*   [Component Architecture Overview](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#component-architecture-overview)
    
*   [Common Patterns](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#common-patterns)
    
*   [Fluent API Configuration](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#fluent-api-configuration)
    
*   [Style Function Callbacks](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#style-function-callbacks)
    
*   [Rendering Pipeline](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#rendering-pipeline)
    
*   [Component Summary](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#component-summary)
    
*   [Tables](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#tables)
    
*   [Trees](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#trees)
    
*   [Lists](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#lists)
    
*   [Integration with Core Style System](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#integration-with-core-style-system)
    
*   [Table Integration](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#table-integration)
    
*   [Tree Integration](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#tree-integration)
    
*   [Data Abstraction and Filtering](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#data-abstraction-and-filtering)
    
*   [Table Data Interface](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#table-data-interface)
    
*   [Tree Children Interface](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#tree-children-interface)
    
*   [Rendering Features Comparison](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#rendering-features-comparison)
    
*   [Usage Patterns](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#usage-patterns)
    
*   [Creating and Configuring Components](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#creating-and-configuring-components)
    
*   [Conditional Styling](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#conditional-styling)
    
*   [Incremental Construction](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#incremental-construction)
    
*   [Dimension and Layout Control](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#dimension-and-layout-control)
    
*   [Table Dimensions](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#table-dimensions)
    
*   [Tree and List Dimensions](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#tree-and-list-dimensions)
    
*   [Next Steps](https://deepwiki.com/charmbracelet/lipgloss/6-structured-components#next-steps)
