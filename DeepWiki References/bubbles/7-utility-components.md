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

Utility Components
==================

Relevant source files

*   [cursor/cursor.go](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go)
    
*   [cursor/cursor\_test.go](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor_test.go)
    
*   [help/help.go](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go)
    
*   [list/keys.go](https://github.com/charmbracelet/bubbles/blob/42130e89/list/keys.go)
    
*   [spinner/spinner\_test.go](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner_test.go)
    

This page documents the utility components in the Bubbles UI Component Library that enhance the user experience of terminal applications. These components provide supporting functionality and foundational primitives rather than serving as primary UI elements like inputs or displays.

For information on input components like Text Input and Text Area, see [Input Components](https://deepwiki.com/charmbracelet/bubbles/2-input-components)
. For display components like Viewport and List, see [Display Components](https://deepwiki.com/charmbracelet/bubbles/3-display-components)
.

Overview
--------

Utility components in the Bubbles library provide infrastructure and supplementary UI features that work alongside primary UI components. They help with:

*   Displaying help information and key bindings to users.
*   Managing cursor state and animations across different input fields.
*   Supporting consistent interactions and visual feedback across the application.

Sources: [help/help.go1-2](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go#L1-L2)
 [cursor/cursor.go1-3](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L1-L3)

Help Component
--------------

The Help component provides a standardized way to display keyboard shortcuts and their descriptions in Bubble Tea applications. It supports two display modes:

1.  **Short Help**: A concise, single-line view showing the most essential commands, often used in status bars.
2.  **Full Help**: A comprehensive, multi-column view showing all available commands, often displayed in a dedicated overlay or expanded view.

The Help component relies on the `KeyMap` interface [help/help.go18-28](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go#L18-L28)
 to retrieve bindings from other components.

### Key Components and Properties

*   **Model**: The main state container [help/help.go77-90](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go#L77-L90)
     It tracks the `ShowAll` toggle, separators, and the current `width` of the terminal to handle truncation.
*   **KeyMap Interface**: Defines `ShortHelp() []key.Binding` and `FullHelp() [][]key.Binding` [help/help.go19-27](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go#L19-L27)
    
*   **Styles**: Manages the visual appearance of keys and descriptions [help/help.go31-43](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go#L31-L43)
    

For details, see [Help](https://deepwiki.com/charmbracelet/bubbles/7.1-help)
.

Sources: [help/help.go77-90](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go#L77-L90)
 [help/help.go18-28](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go#L18-L28)
 [help/help.go31-43](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go#L31-L43)

Cursor Primitive
----------------

The Cursor package provides a virtual cursor primitive used primarily by `textinput` and `textarea`. It manages the visual state of the cursor, including blinking animations and focus/blur states.

### Core Functionality

*   **Blink Management**: Uses a `BlinkMsg` and `blinkCtx` to handle timed animations [cursor/cursor.go29-41](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L29-L41)
    
*   **Modes**: Supports `CursorBlink`, `CursorStatic`, and `CursorHide` [cursor/cursor.go47-51](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L47-L51)
    
*   **Focus Control**: Responds to `tea.FocusMsg` and `tea.BlurMsg` to start or stop animations [cursor/cursor.go126-131](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L126-L131)
    

### Identity and Multi-Instance Safety

To prevent message leakage between multiple cursors (e.g., two text inputs on one screen), each `cursor.Model` is assigned a unique `id` [cursor/cursor.go102](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L102-L102)
 Blink messages contain this `id` and a `tag` to ensure only the intended cursor responds to a specific tick [cursor/cursor.go143-146](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L143-L146)

For details, see [Cursor](https://deepwiki.com/charmbracelet/bubbles/7.2-cursor)
.

Sources: [cursor/cursor.go64-97](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L64-L97)
 [cursor/cursor.go114-158](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L114-L158)
 [cursor/cursor.go214-223](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L214-L223)

Component Integration
---------------------

Utility components are designed to be composed into larger components.

| Utility | Primary Consumers | Integration Method |
| --- | --- | --- |
| `help.Model` | `list`, `table` | Rendered in the parent's `View()` using the parent's `KeyMap`. |
| `cursor.Model` | `textinput`, `textarea` | Embedded in the parent `Model` struct; updated via parent's `Update()`. |

### Example: List KeyMap Integration

The `list` component defines a `KeyMap` struct [list/keys.go7-31](https://github.com/charmbracelet/bubbles/blob/42130e89/list/keys.go#L7-L31)
 that implements the `help.KeyMap` interface, allowing the list to easily pass its bindings to a `help.Model` for rendering.

Sources: [list/keys.go7-31](https://github.com/charmbracelet/bubbles/blob/42130e89/list/keys.go#L7-L31)
 [help/help.go108-113](https://github.com/charmbracelet/bubbles/blob/42130e89/help/help.go#L108-L113)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Utility Components](https://deepwiki.com/charmbracelet/bubbles/7-utility-components#utility-components)
    
*   [Overview](https://deepwiki.com/charmbracelet/bubbles/7-utility-components#overview)
    
*   [Help Component](https://deepwiki.com/charmbracelet/bubbles/7-utility-components#help-component)
    
*   [Key Components and Properties](https://deepwiki.com/charmbracelet/bubbles/7-utility-components#key-components-and-properties)
    
*   [Cursor Primitive](https://deepwiki.com/charmbracelet/bubbles/7-utility-components#cursor-primitive)
    
*   [Core Functionality](https://deepwiki.com/charmbracelet/bubbles/7-utility-components#core-functionality)
    
*   [Identity and Multi-Instance Safety](https://deepwiki.com/charmbracelet/bubbles/7-utility-components#identity-and-multi-instance-safety)
    
*   [Component Integration](https://deepwiki.com/charmbracelet/bubbles/7-utility-components#component-integration)
    
*   [Example: List KeyMap Integration](https://deepwiki.com/charmbracelet/bubbles/7-utility-components#example-list-keymap-integration)
