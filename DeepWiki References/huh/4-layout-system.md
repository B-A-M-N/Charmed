Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/huh](https://github.com/charmbracelet/huh "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 18 April 2025 ([151ba0](https://github.com/charmbracelet/huh/commits/151ba059)
)

*   [Overview](https://deepwiki.com/charmbracelet/huh/1-overview)
    
*   [Architecture](https://deepwiki.com/charmbracelet/huh/1.1-architecture)
    
*   [Installation](https://deepwiki.com/charmbracelet/huh/1.2-installation)
    
*   [Core Components](https://deepwiki.com/charmbracelet/huh/2-core-components)
    
*   [Form](https://deepwiki.com/charmbracelet/huh/2.1-form)
    
*   [Group](https://deepwiki.com/charmbracelet/huh/2.2-group)
    
*   [Field Interface](https://deepwiki.com/charmbracelet/huh/2.3-field-interface)
    
*   [Field Types](https://deepwiki.com/charmbracelet/huh/3-field-types)
    
*   [Text Input Fields](https://deepwiki.com/charmbracelet/huh/3.1-text-input-fields)
    
*   [Selection Fields](https://deepwiki.com/charmbracelet/huh/3.2-selection-fields)
    
*   [Confirmation Field](https://deepwiki.com/charmbracelet/huh/3.3-confirmation-field)
    
*   [Note Field](https://deepwiki.com/charmbracelet/huh/3.4-note-field)
    
*   [FilePicker](https://deepwiki.com/charmbracelet/huh/3.5-filepicker)
    
*   [Layout System](https://deepwiki.com/charmbracelet/huh/4-layout-system)
    
*   [Theme System](https://deepwiki.com/charmbracelet/huh/5-theme-system)
    
*   [Spinner Component](https://deepwiki.com/charmbracelet/huh/6-spinner-component)
    
*   [Accessibility](https://deepwiki.com/charmbracelet/huh/7-accessibility)
    
*   [Bubble Tea Integration](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration)
    
*   [Advanced Usage](https://deepwiki.com/charmbracelet/huh/9-advanced-usage)
    
*   [Dynamic Forms](https://deepwiki.com/charmbracelet/huh/9.1-dynamic-forms)
    
*   [Validation](https://deepwiki.com/charmbracelet/huh/9.2-validation)
    
*   [SSH Integration](https://deepwiki.com/charmbracelet/huh/9.3-ssh-integration)
    
*   [Example Applications](https://deepwiki.com/charmbracelet/huh/10-example-applications)
    

Menu

Layout System
=============

Relevant source files

*   [examples/layout/columns/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/columns/main.go)
    
*   [examples/layout/default/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/default/main.go)
    
*   [examples/layout/grid/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/grid/main.go)
    
*   [examples/layout/stack/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/stack/main.go)
    
*   [examples/scroll/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/scroll/main.go)
    
*   [huh.go](https://github.com/charmbracelet/huh/blob/151ba059/huh.go)
    
*   [internal/selector/selector.go](https://github.com/charmbracelet/huh/blob/151ba059/internal/selector/selector.go)
    
*   [layout.go](https://github.com/charmbracelet/huh/blob/151ba059/layout.go)
    

The Layout System in charmbracelet/huh defines how groups are arranged and displayed within a form. It provides a flexible way to organize form elements, allowing developers to create various visual structures ranging from single-group views to complex multi-column grids.

For information about styling the elements within a layout, see [Theme System](https://deepwiki.com/charmbracelet/huh/5-theme-system)
.

Layout Interface
----------------

The Layout System is centered around the `Layout` interface, which defines the contract that all layout implementations must fulfill.

Sources: [layout.go7-11](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L7-L11)
 [layout.go29-37](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L29-L37)
 [layout.go90-107](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L90-L107)
 [layout.go39-88](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L39-L88)
 [layout.go109-170](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L109-L170)

The `Layout` interface requires two methods:

1.  `View(f *Form) string`: Renders the form's groups according to the layout's rules.
2.  `GroupWidth(f *Form, g *Group, w int) int`: Calculates the appropriate width for a group within the layout, given the total available width.

Built-in Layouts
----------------

The library provides four built-in layout implementations:

### Default Layout

The Default Layout (`LayoutDefault`) displays a single group at a time. This is the simplest layout and is used by default when no specific layout is specified.

Sources: [layout.go14-37](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L14-L37)
 [examples/layout/default/main.go1-25](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/default/main.go#L1-L25)

### Stack Layout

The Stack Layout (`LayoutStack`) displays all groups stacked vertically, one above the other.

Sources: [layout.go17-107](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L17-L107)
 [examples/layout/stack/main.go1-25](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/stack/main.go#L1-L25)

### Columns Layout

The Columns Layout (`LayoutColumns`) arranges groups horizontally in columns. It shows a subset of groups at once, based on the specified number of columns and the currently selected group.

Sources: [layout.go20-88](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L20-L88)
 [examples/layout/columns/main.go1-25](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/columns/main.go#L1-L25)

### Grid Layout

The Grid Layout (`LayoutGrid`) arranges groups in a two-dimensional grid with specified rows and columns. Like the Columns layout, it shows a subset of groups at a time based on the selected group.

Sources: [layout.go25-170](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L25-L170)
 [examples/layout/grid/main.go1-30](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/grid/main.go#L1-L30)

Layout Selection and Form Integration
-------------------------------------

Layouts are integrated with forms using the `WithLayout` method. This integration is demonstrated in the following diagram:

Sources: [layout.go7-170](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L7-L170)

Form and Layout Data Flow
-------------------------

The integration between Forms, the Selector system, and Layouts is crucial for understanding how layouts operate. The following diagram shows this relationship:

Sources: [layout.go7-170](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L7-L170)
 [internal/selector/selector.go4-97](https://github.com/charmbracelet/huh/blob/151ba059/internal/selector/selector.go#L4-L97)

Using Layouts in Applications
-----------------------------

To apply a specific layout to a form, use the `WithLayout` method when creating the form. Here's a table summarizing the available layouts and how to use them:

| Layout | Description | Usage Example |
| --- | --- | --- |
| Default | Shows one group at a time | `form.WithLayout(huh.LayoutDefault)` or omit (default) |
| Stack | Shows all groups vertically stacked | `form.WithLayout(huh.LayoutStack)` |
| Columns | Arranges groups in columns | `form.WithLayout(huh.LayoutColumns(2))` |
| Grid | Arranges groups in a grid | `form.WithLayout(huh.LayoutGrid(2, 3))` |

Sources: [examples/layout/default/main.go1-25](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/default/main.go#L1-L25)
 [examples/layout/stack/main.go1-25](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/stack/main.go#L1-L25)
 [examples/layout/columns/main.go1-25](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/columns/main.go#L1-L25)
 [examples/layout/grid/main.go1-30](https://github.com/charmbracelet/huh/blob/151ba059/examples/layout/grid/main.go#L1-L30)

Layout Implementation Details
-----------------------------

### Selector Integration

All layouts interact with the form's `selector` to determine which groups are visible and which group is currently selected. The Selector provides methods for navigating through items and managing selection state.

Sources: [internal/selector/selector.go4-97](https://github.com/charmbracelet/huh/blob/151ba059/internal/selector/selector.go#L4-L97)
 [layout.go43-63](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L43-L63)
 [layout.go92-97](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L92-L97)
 [layout.go113-144](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L113-L144)

### Layout Rendering Process

The layout rendering process involves several steps:

1.  The form calls the layout's `View` method
2.  The layout determines which groups should be visible
3.  For each visible group, the layout calls `Group.Content()` to get its rendered content
4.  The layout arranges these group renderings according to its layout logic
5.  The layout combines the arranged groups with header and footer content if applicable
6.  The final combined string is returned to the form for display

This process varies slightly between different layout implementations, but the general flow remains consistent.

Sources: [layout.go31-33](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L31-L33)
 [layout.go65-84](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L65-L84)
 [layout.go92-102](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L92-L102)
 [layout.go146-166](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L146-L166)

Custom Layouts
--------------

While the library provides several built-in layouts, developers can create custom layouts by implementing the `Layout` interface. This requires implementing both the `View` and `GroupWidth` methods.

Sources: [layout.go7-11](https://github.com/charmbracelet/huh/blob/151ba059/layout.go#L7-L11)

Conclusion
----------

The Layout System in charmbracelet/huh provides a flexible foundation for organizing form elements in terminal user interfaces. By offering multiple pre-built layouts and allowing for custom implementations, it enables developers to create visually appealing and structurally organized forms that work well in terminal environments.

Through the `Layout` interface, layout implementations encapsulate the logic for determining which groups are visible and how they should be arranged, freeing the rest of the system from these concerns. This separation of responsibilities contributes to the library's overall modularity and extensibility.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Layout System](https://deepwiki.com/charmbracelet/huh/4-layout-system#layout-system)
    
*   [Layout Interface](https://deepwiki.com/charmbracelet/huh/4-layout-system#layout-interface)
    
*   [Built-in Layouts](https://deepwiki.com/charmbracelet/huh/4-layout-system#built-in-layouts)
    
*   [Default Layout](https://deepwiki.com/charmbracelet/huh/4-layout-system#default-layout)
    
*   [Stack Layout](https://deepwiki.com/charmbracelet/huh/4-layout-system#stack-layout)
    
*   [Columns Layout](https://deepwiki.com/charmbracelet/huh/4-layout-system#columns-layout)
    
*   [Grid Layout](https://deepwiki.com/charmbracelet/huh/4-layout-system#grid-layout)
    
*   [Layout Selection and Form Integration](https://deepwiki.com/charmbracelet/huh/4-layout-system#layout-selection-and-form-integration)
    
*   [Form and Layout Data Flow](https://deepwiki.com/charmbracelet/huh/4-layout-system#form-and-layout-data-flow)
    
*   [Using Layouts in Applications](https://deepwiki.com/charmbracelet/huh/4-layout-system#using-layouts-in-applications)
    
*   [Layout Implementation Details](https://deepwiki.com/charmbracelet/huh/4-layout-system#layout-implementation-details)
    
*   [Selector Integration](https://deepwiki.com/charmbracelet/huh/4-layout-system#selector-integration)
    
*   [Layout Rendering Process](https://deepwiki.com/charmbracelet/huh/4-layout-system#layout-rendering-process)
    
*   [Custom Layouts](https://deepwiki.com/charmbracelet/huh/4-layout-system#custom-layouts)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/huh/4-layout-system#conclusion)
