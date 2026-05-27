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

Glossary
========

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/bubbles/blob/42130e89/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1)
    
*   [cursor/cursor.go](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go)
    
*   [cursor/cursor\_test.go](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor_test.go)
    
*   [go.mod](https://github.com/charmbracelet/bubbles/blob/42130e89/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/bubbles/blob/42130e89/go.sum)
    
*   [key/key.go](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key.go)
    
*   [key/key\_test.go](https://github.com/charmbracelet/bubbles/blob/42130e89/key/key_test.go)
    
*   [list/README.md](https://github.com/charmbracelet/bubbles/blob/42130e89/list/README.md?plain=1)
    
*   [list/list.go](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go)
    
*   [progress/progress.go](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go)
    
*   [spinner/spinner.go](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go)
    
*   [textarea/textarea.go](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go)
    
*   [textinput/textinput.go](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go)
    
*   [viewport/viewport.go](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go)
    

This page provides definitions for codebase-specific terms, jargon, and architectural patterns used throughout the `bubbles` library. It is intended to help onboarding engineers navigate the interaction between the Bubble Tea framework and these UI components.

Core Framework Terms
--------------------

The `bubbles` library is built on the **Model-View-Update (MVU)** pattern provided by the [Bubble Tea](https://github.com/charmbracelet/bubbles/blob/42130e89/Bubble%20Tea)
 framework.

| Term | Definition |
| --- | --- |
| **Model** | A `struct` that holds the internal state of a component. In `bubbles`, every component provides a `Model` (e.g., `textinput.Model`, `viewport.Model`) [textinput/textinput.go89-154](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L89-L154) |
| **Update** | A function that handles incoming `tea.Msg` events and returns a new `Model` and an optional `tea.Cmd`. This is where all state transitions occur [spinner/spinner.go131-157](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L131-L157) |
| **View** | A function that returns a `string` representation of the component's state, usually styled with [Lip Gloss](https://github.com/charmbracelet/bubbles/blob/42130e89/Lip%20Gloss)<br> [viewport/viewport.go375-380](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L375-L380) |
| **Cmd** | A `tea.Cmd` is a function that performs an I/O operation and returns a `tea.Msg`. In `bubbles`, these are often used for timers or clipboard access [textinput/textinput.go616-620](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L616-L620) |
| **Msg** | A `tea.Msg` is an interface representing an event (keypress, timer tick, network response) that triggers an `Update` [spinner/spinner.go124-128](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L124-L128) |

**Sources:** [textinput/textinput.go89-154](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L89-L154)
 [spinner/spinner.go131-157](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L131-L157)
 [viewport/viewport.go375-380](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L375-L380)

Component Specific Terminology
------------------------------

### Navigation and Display

*   **Viewport**: A window into a larger piece of content. It manages vertical (`yOffset`) and horizontal (`xOffset`) scrolling [viewport/viewport.go71-75](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L71-L75)
    
*   **Gutter**: A vertical column on the left side of a `viewport` or `textarea` used for line numbers or status indicators, defined via a `GutterFunc` [viewport/viewport.go89-95](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L89-L95)
    
*   **Delegate**: An interface (`ItemDelegate`) used by the `list` component to offload the rendering and update logic of individual items, allowing the list to be agnostic of the data it displays [list/list.go46-61](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L46-L61)
    
*   **Fuzzy Filtering**: A search technique used in the `list` component to find items based on approximate string matching, typically using the `sahilm/fuzzy` library [list/list.go98-109](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L98-L109)
    

### Input and Cursors

*   **Virtual Cursor**: A cursor rendered manually by the component (usually a styled block or character) rather than relying on the terminal's hardware cursor. Managed by the `cursor` package [cursor/cursor.go64-97](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L64-L97)
    
*   **Echo Mode**: A setting in `textinput` that determines how characters are displayed (e.g., `EchoPassword` hides input with a mask) [textinput/textinput.go28-41](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L28-L41)
    
*   **Soft Wrap**: A wrapping mode where long lines are visually broken into multiple lines to fit the width of the component without inserting actual newline characters [viewport/viewport.go57-59](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L57-L59)
    

### Animation and Feedback

*   **Spring**: A physics-based animation primitive from the `harmonica` library used by the `progress` bar to create smooth transitions [progress/progress.go157-161](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L157-L161)
    
*   **Tick**: A recurring `tea.Msg` (e.g., `spinner.TickMsg`) used to drive animations by triggering the next frame in the `Update` loop [spinner/spinner.go124-128](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L124-L128)
    

**Sources:** [viewport/viewport.go57-95](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L57-L95)
 [list/list.go46-61](https://github.com/charmbracelet/bubbles/blob/42130e89/list/list.go#L46-L61)
 [textinput/textinput.go28-41](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L28-L41)
 [cursor/cursor.go64-97](https://github.com/charmbracelet/bubbles/blob/42130e89/cursor/cursor.go#L64-L97)
 [progress/progress.go157-161](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L157-L161)

Internal Implementation Concepts
--------------------------------

### Message Routing (ID & Tag Pattern)

To prevent "message leakage" where multiple instances of the same component (e.g., two spinners) receive each other's messages, `bubbles` uses an ID and Tagging system.

**The ID/Tag Flow**

1.  **ID**: A unique integer assigned to each `Model` instance upon creation via `nextID()` [spinner/spinner.go16-18](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L16-L18)
    
2.  **Tag**: A sequence number that increments with every tick.
3.  **Verification**: In the `Update` function, the component checks if the incoming `Msg.ID` matches its own `Model.id`. If not, it ignores the message [spinner/spinner.go133-145](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L133-L145)
    

### Data Flow: Message Routing

The following diagram illustrates how a `TickMsg` is routed to the correct `spinner.Model` instance using the ID pattern.

**Spinner Message Routing**

**Sources:** [spinner/spinner.go100-107](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L100-L107)
 [spinner/spinner.go131-157](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L131-L157)

### Coordinate Mapping (Text Area)

The `textarea` component manages complex mappings between the raw data (`[][]rune`) and the visual display.

| Entity | Code Symbol | Role |
| --- | --- | --- |
| **Raw Value** | `Model.value` | The underlying slice of runes representing the text [textinput/textinput.go123](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L123-L123) |
| **Line Metadata** | `LineInfo` | Tracks width, height, and offsets for soft-wrapped lines [textarea/textarea.go111-138](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L111-L138) |
| **Visual Viewport** | `viewport.Model` | Handles the actual scrolling and rendering of the visible window [textarea/textarea.go19](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L19-L19) |

**Coordinate Space Mapping**

**Sources:** [textarea/textarea.go111-138](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L111-L138)
 [textinput/textinput.go123](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L123-L123)
 [viewport/viewport.go71-75](https://github.com/charmbracelet/bubbles/blob/42130e89/viewport/viewport.go#L71-L75)

Key Abbreviations
-----------------

*   **MVU**: Model-View-Update architecture.
*   **TUI**: Terminal User Interface.
*   **ANSI**: American National Standards Institute (refers to terminal escape codes for color and cursor movement).
*   **FPS**: Frames Per Second (used in `progress` and `spinner` for animation timing) [progress/progress.go48](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L48-L48)
    

**Sources:** [README.md10-11](https://github.com/charmbracelet/bubbles/blob/42130e89/README.md?plain=1#L10-L11)
 [progress/progress.go48](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L48-L48)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Glossary](https://deepwiki.com/charmbracelet/bubbles/9-glossary#glossary)
    
*   [Core Framework Terms](https://deepwiki.com/charmbracelet/bubbles/9-glossary#core-framework-terms)
    
*   [Component Specific Terminology](https://deepwiki.com/charmbracelet/bubbles/9-glossary#component-specific-terminology)
    
*   [Navigation and Display](https://deepwiki.com/charmbracelet/bubbles/9-glossary#navigation-and-display)
    
*   [Input and Cursors](https://deepwiki.com/charmbracelet/bubbles/9-glossary#input-and-cursors)
    
*   [Animation and Feedback](https://deepwiki.com/charmbracelet/bubbles/9-glossary#animation-and-feedback)
    
*   [Internal Implementation Concepts](https://deepwiki.com/charmbracelet/bubbles/9-glossary#internal-implementation-concepts)
    
*   [Message Routing (ID & Tag Pattern)](https://deepwiki.com/charmbracelet/bubbles/9-glossary#message-routing-id-tag-pattern)
    
*   [Data Flow: Message Routing](https://deepwiki.com/charmbracelet/bubbles/9-glossary#data-flow-message-routing)
    
*   [Coordinate Mapping (Text Area)](https://deepwiki.com/charmbracelet/bubbles/9-glossary#coordinate-mapping-text-area)
    
*   [Key Abbreviations](https://deepwiki.com/charmbracelet/bubbles/9-glossary#key-abbreviations)
