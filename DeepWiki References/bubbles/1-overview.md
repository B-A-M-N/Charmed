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

Overview
========

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/bubbles/blob/42130e89/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1)
    
*   [bubbles.go](https://github.com/charmbracelet/bubbles/blob/42130e89/bubbles.go)
    
*   [go.mod](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/bubbles/blob/42130e89/go.sum)
    
*   [key/key.go](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key.go)
    
*   [key/key\_test.go](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key_test.go)
    
*   [list/README.md](https://github.com/charmbracelet/bubbles/blob/42130e89/list/README.md?plain=1)
    

Bubbles is a collection of interactive UI components (or "bubbles") for [Bubble Tea](https://github.com/charmbracelet/bubbles/blob/42130e89/Bubble%20Tea)
 applications [bubbles.go1-4](https://github.com/charmbracelet/bubbles/blob/42130e89/bubbles.go#L1-L4)
 These components provide professionally crafted primitives for common terminal UI patterns—such as text inputs, viewports, and progress bars—allowing developers to build sophisticated applications without reinventing low-level terminal logic [README.md10-11](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L10-L11)

For a detailed explanation of the architectural concepts and how components integrate with the Bubble Tea framework, see [Architecture and Concepts](https://deepwiki.com/charmbracelet/bubbles/1.1-architecture-and-concepts)
. If you're looking to get started quickly or are upgrading to v2, see [Getting Started](https://deepwiki.com/charmbracelet/bubbles/1.2-getting-started)
.

Component Ecosystem
-------------------

Bubbles provides a variety of UI components organized by functionality:

Sources: [README.md21-153](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L21-L153)
 [go.mod1-17](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod#L1-L17)
 [list/README.md43-55](https://github.com/charmbracelet/bubbles/blob/42130e89/list/README.md?plain=1#L43-L55)

Key Features
------------

The Bubbles library offers several key features for terminal UI development:

1.  **Model-View-Update (MVU) Compliance**: Every component implements the standard `tea.Model` interface (`Init`, `Update`, `View`), making them easy to embed in parent models [key/key.go21-33](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key.go#L21-L33)
    
2.  **Customizable Styling**: Components use `lipgloss` for layout and colors, often exposing a `Styles` struct for deep customization [list/README.md45-59](https://github.com/charmbracelet/bubbles/blob/42130e89/list/README.md?plain=1#L45-L59)
    
3.  **Keyboard Driven**: Built-in support for customizable keybindings via the `key` package [key/key.go1-37](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key.go#L1-L37)
    
4.  **Production Ready**: Used in production tools like `Glow` and `Crush` [bubbles.go1-3](https://github.com/charmbracelet/bubbles/blob/42130e89/bubbles.go#L1-L3)
     [README.md11](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L11-L11)
    

Sources: [bubbles.go1-4](https://github.com/charmbracelet/bubbles/blob/42130e89/bubbles.go#L1-L4)
 [README.md11](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L11-L11)
 [key/key.go5-19](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key.go#L5-L19)
 [list/README.md45-59](https://github.com/charmbracelet/bubbles/blob/42130e89/list/README.md?plain=1#L45-L59)

Component Architecture Pattern
------------------------------

Bubbles components are designed to be "embedded" within a main Bubble Tea model. The parent model delegates messages to the component's `Update` method and includes the component's `View` output in its own rendering.

Sources: [key/key.go21-33](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key.go#L21-L33)
 [key/key.go43-47](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key.go#L43-L47)
 [list/README.md45-59](https://github.com/charmbracelet/bubbles/blob/42130e89/list/README.md?plain=1#L45-L59)

User Interaction Flow
---------------------

The following diagram illustrates how user input (via `tea.KeyPressMsg`) is matched against component bindings:

Sources: [key/key.go21-30](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key.go#L21-L30)
 [key/key.go129-140](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key.go#L129-L140)

Key Categories
--------------

| Category | Key Components | Description |
| --- | --- | --- |
| **Input** | `textinput`, `textarea` | Single and multi-line text entry with cursor management [README.md31-52](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L31-L52) |
| **Display** | `viewport`, `list`, `table` | Scrollable regions, item lists with filtering, and data tables [README.md54-114](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L54-L114) |
| **Progress** | `progress`, `spinner` | Visual feedback for background tasks or loading states [README.md21-75](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L21-L75) |
| **Navigation** | `paginator`, `filepicker` | Page logic and filesystem navigation [README.md77-124](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L77-L124) |
| **Utility** | `key`, `help`, `cursor` | Primitives for keybindings, help menus, and terminal cursors [README.md144-159](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L144-L159) |

Sources: [README.md21-159](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L21-L159)

Core Dependencies
-----------------

Bubbles leverages several specialized libraries within the Charm ecosystem and beyond:

*   **`charm.land/bubbletea/v2`**: The foundational runtime framework [go.mod6](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod#L6-L6)
    
*   **`charm.land/lipgloss/v2`**: Used for all component styling and layout [go.mod7](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod#L7-L7)
    
*   **`github.com/charmbracelet/harmonica`**: Provides spring-based animation physics for components like the progress bar [go.mod10](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod#L10-L10)
    
*   **`github.com/atotto/clipboard`**: Enables cross-platform copy/paste support for input components [go.mod9](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod#L9-L9)
    
*   **`github.com/sahilm/fuzzy`**: Powers the filtering logic in the `list` component [go.mod16](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod#L16-L16)
    

Sources: [go.mod5-17](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod#L5-L17)

License
-------

Bubbles is released under the **MIT License** [LICENSE1-21](https://github.com/charmbracelet/bubbles/blob/42130e89/LICENSE#L1-L21)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/bubbles/1-overview#overview)
    
*   [Component Ecosystem](https://deepwiki.com/charmbracelet/bubbles/1-overview#component-ecosystem)
    
*   [Key Features](https://deepwiki.com/charmbracelet/bubbles/1-overview#key-features)
    
*   [Component Architecture Pattern](https://deepwiki.com/charmbracelet/bubbles/1-overview#component-architecture-pattern)
    
*   [User Interaction Flow](https://deepwiki.com/charmbracelet/bubbles/1-overview#user-interaction-flow)
    
*   [Key Categories](https://deepwiki.com/charmbracelet/bubbles/1-overview#key-categories)
    
*   [Core Dependencies](https://deepwiki.com/charmbracelet/bubbles/1-overview#core-dependencies)
    
*   [License](https://deepwiki.com/charmbracelet/bubbles/1-overview#license)
