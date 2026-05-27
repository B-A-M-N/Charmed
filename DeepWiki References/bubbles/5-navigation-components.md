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

Navigation Components
=====================

Relevant source files

*   [filepicker/filepicker.go](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go)
    
*   [paginator/paginator.go](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go)
    

This page documents the navigation components provided by the `bubbles` library, which help users navigate through content in terminal-based applications. These components include tools for pagination and file system navigation.

Overview
--------

Navigation components in the `bubbles` library serve to help users traverse different kinds of content:

1.  The **Paginator** component - handles pagination state and rendering pagination information. For details, see [Paginator](https://deepwiki.com/charmbracelet/bubbles/5.1-paginator)
    .
2.  The **File Picker** component - enables file system navigation and file selection. For details, see [File Picker](https://deepwiki.com/charmbracelet/bubbles/5.2-file-picker)
    .

### Component Relationships

The following diagram illustrates how navigation components interface with other library entities and user applications.

Title: Navigation Component Integration

Sources: [paginator/paginator.go1-5](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L1-L5)
 [filepicker/filepicker.go1-3](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L1-L3)

Paginator Component
-------------------

The Paginator component (package `paginator`) provides mechanisms for calculating pagination and rendering pagination indicators. It is a logic-heavy component that handles pagination state, key bindings for navigation, and rendering pagination status strings.

Title: Paginator Model Structure

Sources: [paginator/paginator.go39-57](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L39-L57)
 [paginator/paginator.go24-27](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L24-L27)

### Key Concepts

*   **Pagination Types**: Supports `Arabic` (numeric e.g., "1/5") and `Dots` (visual e.g., "•○") rendering [paginator/paginator.go18-21](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L18-L21)
    
*   **Slice Management**: The `GetSliceBounds` method [paginator/paginator.go92-96](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L92-L96)
     is used to determine which subset of a data array should be rendered based on the current `Page` and `PerPage` settings.
*   **State Helpers**: Methods like `OnLastPage()` [paginator/paginator.go116-118](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L116-L118)
     and `SetTotalPages()` [paginator/paginator.go63-73](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L63-L73)
     simplify integration with dynamic data sets.

For more in-depth technical details, see [Paginator](https://deepwiki.com/charmbracelet/bubbles/5.1-paginator)
.

Sources: [paginator/paginator.go39-57](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L39-L57)
 [paginator/paginator.go92-96](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L92-L96)

File Picker Component
---------------------

The File Picker component (package `filepicker`) allows users to navigate file systems and select files or directories within a terminal-based interface.

Title: File Picker Selection Lifecycle

Sources: [filepicker/filepicker.go187-196](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L187-L196)
 [filepicker/filepicker.go197-225](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L197-L225)
 [filepicker/filepicker.go447-465](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L447-L465)

### Key Features

*   **Stack-based History**: Uses internal stacks (`selectedStack`, `minStack`, `maxStack`) to restore cursor positions and scroll offsets when navigating back up the directory tree [filepicker/filepicker.go150-155](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L150-L155)
    
*   **Asynchronous I/O**: Directory reading is performed via `readDirMsg`, ensuring the UI remains responsive [filepicker/filepicker.go55-58](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L55-L58)
    
*   **Filtering & Permissions**: Supports `AllowedTypes` for extension filtering and toggles for `ShowPermissions` or `ShowSize` [filepicker/filepicker.go136-144](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L136-L144)
    
*   **Styling**: Highly customizable via the `Styles` struct, including specific colors for symlinks, directories, and file sizes [filepicker/filepicker.go95-107](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L95-L107)
    

For more in-depth technical details, see [File Picker](https://deepwiki.com/charmbracelet/bubbles/5.2-file-picker)
.

Sources: [filepicker/filepicker.go127-162](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L127-L162)
 [filepicker/filepicker.go95-107](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L95-L107)

Shared Navigation Patterns
--------------------------

While distinct in purpose, these components share several Bubbles-standard patterns:

1.  **Key-Driven Movement**: Both utilize the `charm.land/bubbles/v2/key` package to define flexible `KeyMap` structures [paginator/paginator.go24-27](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L24-L27)
     [filepicker/filepicker.go67-77](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L67-L77)
    
2.  **ID-Based Message Routing**: The File Picker uses a `nextID()` generator and `id` field in messages like `readDirMsg` to ensure that directory results are routed to the correct component instance [filepicker/filepicker.go22-30](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L22-L30)
     [filepicker/filepicker.go248-251](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L248-L251)
    
3.  **Viewport Logic**: Both components manage internal offsets (`minIdx`/`maxIdx` in File Picker, `Page` in Paginator) to simulate a viewport over a larger collection of data [filepicker/filepicker.go152-153](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L152-L153)
     [paginator/paginator.go43](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L43-L43)
    

Sources: [paginator/paginator.go24-27](https://github.com/charmbracelet/bubbles/blob/42130e89/paginator/paginator.go#L24-L27)
 [filepicker/filepicker.go22-30](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L22-L30)
 [filepicker/filepicker.go67-77](https://github.com/charmbracelet/bubbles/blob/42130e89/filepicker/filepicker.go#L67-L77)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Navigation Components](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components#navigation-components)
    
*   [Overview](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components#overview)
    
*   [Component Relationships](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components#component-relationships)
    
*   [Paginator Component](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components#paginator-component)
    
*   [Key Concepts](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components#key-concepts)
    
*   [File Picker Component](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components#file-picker-component)
    
*   [Key Features](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components#key-features)
    
*   [Shared Navigation Patterns](https://deepwiki.com/charmbracelet/bubbles/5-navigation-components#shared-navigation-patterns)
