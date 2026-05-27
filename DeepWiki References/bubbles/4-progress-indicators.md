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

Progress Indicators
===================

Relevant source files

*   [progress/progress.go](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go)
    
*   [spinner/spinner.go](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go)
    

Progress indicators in the bubbles library provide visual feedback about ongoing processes in terminal applications. This document covers the two primary progress indicator components: the Progress Bar and the Spinner. Both components integrate with Bubble Tea applications and share common patterns for animation and message routing.

For details on the progress bar, see [Progress Bar](https://deepwiki.com/charmbracelet/bubbles/4.1-progress-bar)
. For details on the spinner, see [Spinner](https://deepwiki.com/charmbracelet/bubbles/4.2-spinner)
.

Overview
--------

Progress indicators serve crucial roles in terminal user interfaces based on the predictability of the task duration:

1.  **Progress Bar**: Displays completion percentage of a task with a definite end.
2.  **Spinner**: Indicates activity for tasks with indeterminate duration.

Sources: [progress/progress.go1-2](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L1-L2)
 [spinner/spinner.go1-2](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L1-L2)

Shared Patterns
---------------

### Internal ID and Tag Management

To support multiple instances of progress indicators within a single Bubble Tea application, both components use an internal ID system. A global counter is incremented atomically to assign a unique `id` to every new model.

*   **ID**: Ensures messages are routed to the correct component instance [progress/progress.go26-30](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L26-L30)
     [spinner/spinner.go14-18](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L14-L18)
    
*   **Tag**: A secondary identifier used to prevent "message leakage" where a component might process stale animation messages, leading to animations running too fast [progress/progress.go183-193](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L183-L193)
     [spinner/spinner.go101-102](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L101-L102)
    

### Animation Tick Loop

Animations are driven by the Bubble Tea `Update` loop. Components emit a command that triggers a message after a certain duration (e.g., 60 FPS for progress bars, or variable FPS for spinners), which then advances the internal state.

Sources: [progress/progress.go181-184](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L181-L184)
 [spinner/spinner.go124-128](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L124-L128)

* * *

Progress Bar
------------

The progress bar component visualizes task completion as a horizontal bar. It features smooth spring-based animations and high-resolution color blending.

### Key Features

*   **Spring Animation**: Uses the `harmonica` library for fluid transitions [progress/progress.go212-217](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L212-L217)
    
*   **Half-Block Rendering**: Uses `DefaultFullCharHalfBlock` ('▌') to double color resolution by manipulating foreground and background colors [progress/progress.go32-37](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L32-L37)
    
*   **Color Blending**: Supports solid colors, gradients, or dynamic `ColorFunc` callbacks [progress/progress.go18-22](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L18-L22)
    

For in-depth documentation, see [Progress Bar](https://deepwiki.com/charmbracelet/bubbles/4.1-progress-bar)
.

Sources: [progress/progress.go187-229](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L187-L229)
 [progress/progress.go61-70](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L61-L70)

* * *

Spinner
-------

The spinner component provides animated loading indicators for tasks with an indeterminate duration. It cycles through a sequence of frames defined in a `Spinner` struct.

### Key Features

*   **Predefined Types**: Includes numerous styles such as `Line`, `Dot`, `Globe`, and `Monkey` [spinner/spinner.go27-84](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L27-L84)
    
*   **Customizable FPS**: Each spinner type defines its own animation speed [spinner/spinner.go21-24](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L21-L24)
    
*   **Manual Trigger**: The `Tick()` method allows manual control over starting the animation loop [spinner/spinner.go170-182](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L170-L182)
    

For in-depth documentation, see [Spinner](https://deepwiki.com/charmbracelet/bubbles/4.2-spinner)
.

### Predefined Spinner Varieties

| Type | Frames |
| --- | --- |
| `Line` | \`  |
| `Dot` | Braille patterns (⣾, ⣽, etc.) |
| `Pulse` | Block gradients (█, ▓, ▒, ░) |
| `Globe` | 🌍, 🌎, 🌏 |

Sources: [spinner/spinner.go28-83](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L28-L83)

* * *

Choosing a Component
--------------------

| Requirement | Use Case | Recommended Component |
| --- | --- | --- |
| Known completion state | File download, installation | **Progress Bar** |
| Unknown duration | Database query, API request | **Spinner** |
| High-precision UI | Dashboards with gradients | **Progress Bar** |
| Minimalist activity indicator | CLI tools, sidebars | **Spinner** |

Sources: [progress/progress.go1-2](https://github.com/charmbracelet/bubbles/blob/42130e89/progress/progress.go#L1-L2)
 [spinner/spinner.go1-2](https://github.com/charmbracelet/bubbles/blob/42130e89/spinner/spinner.go#L1-L2)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Progress Indicators](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#progress-indicators)
    
*   [Overview](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#overview)
    
*   [Shared Patterns](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#shared-patterns)
    
*   [Internal ID and Tag Management](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#internal-id-and-tag-management)
    
*   [Animation Tick Loop](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#animation-tick-loop)
    
*   [Progress Bar](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#progress-bar)
    
*   [Key Features](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#key-features)
    
*   [Spinner](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#spinner)
    
*   [Key Features](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#key-features-1)
    
*   [Predefined Spinner Varieties](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#predefined-spinner-varieties)
    
*   [Choosing a Component](https://deepwiki.com/charmbracelet/bubbles/4-progress-indicators#choosing-a-component)
