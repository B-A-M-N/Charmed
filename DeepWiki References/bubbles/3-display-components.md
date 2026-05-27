Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/bubbles](https://github.com/charmbracelet/bubbles "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 19 April 2026 ([42130e](https://github.com/charmbracelet/bubbles/commits/42130e89)
)

*   [Overview](https://deepwiki.com/charmbracelet/bubbles/1-overview)
    
*   [Architecture and Concepts](https://deepwiki.com/charmbracelet/bubbles/1.1-architecture-and-concepts)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/bubbles/1.2-getting-started)
    
*   [Input Components](https://deepwiki.com/charmbracelet/bubbles/2-input-components)
    
*   [Text Input](https://deepwiki.com/charmbracelet/bubbles/2.1-text-input)
    
*   [Text Area](https://deepwiki.com/charmbracelet/bubbles/2.2-text-area)
    
*   [Display Components](https://deepwiki.com/charmbracelet/bubbles/3-display-components)
    
*   [Viewport](https://deepwiki.com/charmbracelet/bubbles/3.1-viewport)
    
*   [List](https://deepwiki.com/charmbracelet/bubbles/3.2-list)
    
*   [Table](https://deepwiki.com/charmbracelet/bubbles/3.3-table)
    
*   [Progress Indicators](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators)
    
*   [Progress Bar](https://deepwiki.com/charmbracelet/bubbles/4.1-progress-bar)
    
*   [Spinner](https://deepwiki.com/charmbracelet/bubbles/4.2-spinner)
    
*   [Navigation Components](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components)
    
*   [Paginator](https://deepwiki.com/charmbracelet/bubbles/5.1-paginator)
    
*   [File Picker](https://deepwiki.com/charmbracelet/bubbles/5.2-file-picker)
    
*   [Time Components](https://deepwiki.com/charmbracelet/bubbles/6-time-components)
    
*   [Stopwatch and Timer](https://deepwiki.com/charmbracelet/bubbles/6.1-stopwatch-and-timer)
    
*   [Utility Components](https://deepwiki.com/charmbracelet/bubbles/7-utility-components)
    
*   [Help](https://deepwiki.com/charmbracelet/bubbles/7.1-help)
    
*   [Cursor](https://deepwiki.com/charmbracelet/bubbles/7.2-cursor)
    
*   [Development and Contributing](https://deepwiki.com/charmbracelet/bubbles/8-development-and-contributing)
    
*   [Testing and Linting](https://deepwiki.com/charmbracelet/bubbles/8.1-testing-and-linting)
    
*   [CI/CD and Releases](https://deepwiki.com/charmbracelet/bubbles/8.2-cicd-and-releases)
    
*   [Migrating from v1 to v2](https://deepwiki.com/charmbracelet/bubbles/8.3-migrating-from-v1-to-v2)
    
*   [Glossary](https://deepwiki.com/charmbracelet/bubbles/9-glossary)
    

Menu

Display Components
==================

Relevant source files

*   [list/list.go](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go)
    
*   [table/table.go](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table.go)
    
*   [table/table\_test.go](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table_test.go)
    
*   [viewport/viewport.go](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go)
    

Display components in the Bubbles library provide a way to present and navigate through content in terminal-based user interfaces. These components handle common display patterns such as scrollable content, selectable lists, and tabular data, allowing you to build rich interactive UIs with minimal effort.

This page covers the primary display components available in the Bubbles library: Viewport, List, and Table.

Core Display Components Overview
--------------------------------

Sources: [viewport/viewport.go51-116](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L51-L116)
 [list/list.go147-204](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L147-L204)
 [table/table.go16-29](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table.go#L16-L29)

Viewport Component
------------------

The Viewport component provides a scrollable window that can display content larger than the visible area. It's ideal for showing large text blocks, logs, or any content that requires vertical and horizontal scrolling. It serves as a foundational building block for other components like the `table` and `textarea`.

For in-depth documentation, see [Viewport](https://deepwiki.com/charmbracelet/bubbles/3.1-viewport)
.

Sources: [viewport/viewport.go51-116](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L51-L116)
 [viewport/viewport.go137-146](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L137-L146)

### Key Features

*   **Scroll Management**: Tracks `yOffset` and `xOffset` for 2D navigation [viewport/viewport.go71-75](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L71-L75)
    
*   **Gutter System**: Supports custom left gutters (e.g., line numbers) via `LeftGutterFunc` [viewport/viewport.go89-95](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L89-L95)
    
*   **Dynamic Styling**: Provides `StyleLineFunc` to style individual lines based on index [viewport/viewport.go110-112](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L110-L112)
    
*   **High Performance**: Optimized rendering logic for large content buffers.

List Component
--------------

The List component provides a flexible, interactive list with built-in pagination, filtering, and keyboard navigation. It uses a delegate-based system to allow custom rendering of items without changing the underlying data structure.

For in-depth documentation, see [List](https://deepwiki.com/charmbracelet/bubbles/3.2-list)
.

Sources: [list/list.go34-38](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L34-L38)
 [list/list.go46-61](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L46-L61)
 [list/list.go147-204](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L147-L204)

### Key Features

*   **Fuzzy Filtering**: Built-in support for searching items using `sahilm/fuzzy` [list/list.go96-109](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L96-L109)
    
*   **Pagination**: Integrates `paginator.Model` for handling large datasets across multiple views [list/list.go182](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L182-L182)
    
*   **Delegate Pattern**: Decouples item data from view logic via the `ItemDelegate` interface [list/list.go40-45](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L40-L45)
    
*   **State Management**: Tracks `FilterState` (Unfiltered, Filtering, FilterApplied) to manage UI modes [list/list.go131-135](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L131-L135)
    

Table Component
---------------

The Table component presents data in a tabular format with columns and rows. It leverages an internal `viewport.Model` to handle scrolling of large datasets while maintaining a fixed header.

For in-depth documentation, see [Table](https://deepwiki.com/charmbracelet/bubbles/3.3-table)
.

Sources: [table/table.go16-29](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table.go#L16-L29)
 [table/table.go31-38](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table.go#L31-L38)

### Key Features

*   **Viewport Integration**: Uses a `viewport.Model` internally to handle vertical scrolling of rows [table/table.go26](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table.go#L26-L26)
    
*   **Cursor Management**: Tracks a `cursor` index to manage row selection and highlighting [table/table.go22](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table.go#L22-L22)
    
*   **Data Utilities**: Includes helpers like `FromValues` to populate tables from CSV-like strings [table/table\_test.go210-214](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table_test.go#L210-L214)
    
*   **Visual Truncation**: Automatically handles cell overflow using `ansi.Truncate` to fit defined column widths [table/table.go12](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table.go#L12-L12)
    

Component Patterns
------------------

The display components share several architectural patterns:

| Pattern | Description | Implementation |
| --- | --- | --- |
| **Viewport Integration** | Using a scrollable sub-component for overflow. | Found in `table` and `textarea`. |
| **Delegate Rendering** | Offloading visual representation to a separate interface. | Primary pattern for `list.ItemDelegate`. |
| **Help Integration** | Implementing `help.KeyMap` to automatically populate help menus. | Used by `list.KeyMap` and `table.KeyMap`. |
| **Functional Options** | Using `WithWidth`, `WithHeight` etc., for clean initialization. | Standard across `viewport.New` and `table.New`. |

Sources: [viewport/viewport.go27-39](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L27-L39)
 [list/list.go46-61](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L46-L61)
 [table/table.go153-199](https://github.com/charmbracelet/bubbles/blob/42130e89/table/table.go#L153-L199)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Display Components](https://deepwiki.com/charmbracelet/bubbles/3-display-components#display-components)
    
*   [Core Display Components Overview](https://deepwiki.com/charmbracelet/bubbles/3-display-components#core-display-components-overview)
    
*   [Viewport Component](https://deepwiki.com/charmbracelet/bubbles/3-display-components#viewport-component)
    
*   [Key Features](https://deepwiki.com/charmbracelet/bubbles/3-display-components#key-features)
    
*   [List Component](https://deepwiki.com/charmbracelet/bubbles/3-display-components#list-component)
    
*   [Key Features](https://deepwiki.com/charmbracelet/bubbles/3-display-components#key-features-1)
    
*   [Table Component](https://deepwiki.com/charmbracelet/bubbles/3-display-components#table-component)
    
*   [Key Features](https://deepwiki.com/charmbracelet/bubbles/3-display-components#key-features-2)
    
*   [Component Patterns](https://deepwiki.com/charmbracelet/bubbles/3-display-components#component-patterns)
