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

Time Components
===============

Relevant source files

*   [stopwatch/stopwatch.go](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go)
    
*   [timer/timer.go](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go)
    

Overview
--------

The Bubbles library provides two time-related components designed for tracking and measuring time in terminal applications: `stopwatch` and `timer`. These components implement the Bubble Tea Model-View-Update (MVU) pattern and can be easily integrated into Bubble Tea applications to provide time tracking functionality.

For details, see [Stopwatch and Timer](https://deepwiki.com/charmbracelet/bubbles/6.1-stopwatch-and-timer)
.

Sources: [stopwatch/stopwatch.go1-4](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go#L1-L4)
 [timer/timer.go1-2](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L1-L2)

Component Architecture
----------------------

The time components in the Bubbles library follow a consistent architecture pattern, leveraging internal IDs and tags to ensure message routing safety in multi-instance environments.

### Natural Language to Code Entity Mapping: Architecture

Both components implement unique ID management via a package-level `lastID` counter and an atomic `nextID()` function to allow multiple instances to run simultaneously without message collision.

Sources: [stopwatch/stopwatch.go11-15](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go#L11-L15)
 [stopwatch/stopwatch.go55-63](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go#L55-L63)
 [timer/timer.go11-15](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L11-L15)
 [timer/timer.go90-100](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L90-L100)

Shared Patterns
---------------

### Message Flow and Routing

Time components use a `TickMsg` for interval-based updates and a `StartStopMsg` for state control. To prevent "message leakage" where one component's tick triggers another's update, both models check the `ID` and `tag` fields within their `Update` methods.

### Component Comparison

| Feature | Stopwatch | Timer |
| --- | --- | --- |
| **Direction** | Counts up (Elapsed) | Counts down (Timeout) |
| **Primary State** | `d time.Duration` [stopwatch/stopwatch.go56](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go#L56-L56) | `Timeout time.Duration` [timer/timer.go92](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L92-L92) |
| **Messages** | `TickMsg`, `StartStopMsg`, `ResetMsg` | `TickMsg`, `StartStopMsg`, `TimeoutMsg` |
| **End State** | Continuous until stopped | Sends `TimeoutMsg` at 0 [timer/timer.go81-87](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L81-L87) |
| **Default Interval** | 1 second | 1 second [timer/timer.go106](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L106-L106) |

Sources: [stopwatch/stopwatch.go31-52](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go#L31-L52)
 [timer/timer.go57-87](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L57-L87)
 [timer/timer.go142-167](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L142-L167)
 [stopwatch/stopwatch.go122-152](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go#L122-L152)

Usage and Integration
---------------------

### Message Propagation

All `stopwatch` and `timer` messages must be passed to the component's `Update` method. Both components provide helper methods like `Start()`, `Stop()`, and `Toggle()` that return `tea.Cmd` wrapping a `StartStopMsg`.

### Configuration

Components are initialized using functional options. The most common option is `WithInterval`, which defines the granularity of the `TickMsg` loop.

*   **Stopwatch**: `stopwatch.New(stopwatch.WithInterval(time.Millisecond))` [stopwatch/stopwatch.go24-28](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go#L24-L28)
    
*   **Timer**: `timer.New(duration, timer.WithInterval(time.Second))` [timer/timer.go24-28](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L24-L28)
    

### Natural Language to Code Entity Mapping: Logic Flow

For detailed implementation specifics on resetting stopwatches or handling timer timeouts, see [Stopwatch and Timer](https://deepwiki.com/charmbracelet/bubbles/6.1-stopwatch-and-timer)
.

Sources: [stopwatch/stopwatch.go134-149](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go#L134-L149)
 [timer/timer.go150-164](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L150-L164)
 [stopwatch/stopwatch.go160-162](https://github.com/charmbracelet/bubbles/blob/42130e89/stopwatch/stopwatch.go#L160-L162)
 [timer/timer.go170-172](https://github.com/charmbracelet/bubbles/blob/42130e89/timer/timer.go#L170-L172)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Time Components](https://deepwiki.com/charmbracelet/bubbles/6-time-components#time-components)
    
*   [Overview](https://deepwiki.com/charmbracelet/bubbles/6-time-components#overview)
    
*   [Component Architecture](https://deepwiki.com/charmbracelet/bubbles/6-time-components#component-architecture)
    
*   [Natural Language to Code Entity Mapping: Architecture](https://deepwiki.com/charmbracelet/bubbles/6-time-components#natural-language-to-code-entity-mapping-architecture)
    
*   [Shared Patterns](https://deepwiki.com/charmbracelet/bubbles/6-time-components#shared-patterns)
    
*   [Message Flow and Routing](https://deepwiki.com/charmbracelet/bubbles/6-time-components#message-flow-and-routing)
    
*   [Component Comparison](https://deepwiki.com/charmbracelet/bubbles/6-time-components#component-comparison)
    
*   [Usage and Integration](https://deepwiki.com/charmbracelet/bubbles/6-time-components#usage-and-integration)
    
*   [Message Propagation](https://deepwiki.com/charmbracelet/bubbles/6-time-components#message-propagation)
    
*   [Configuration](https://deepwiki.com/charmbracelet/bubbles/6-time-components#configuration)
    
*   [Natural Language to Code Entity Mapping: Logic Flow](https://deepwiki.com/charmbracelet/bubbles/6-time-components#natural-language-to-code-entity-mapping-logic-flow)
