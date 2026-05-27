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

Input Components
================

Relevant source files

*   [textarea/textarea.go](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go)
    
*   [textinput/textinput.go](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go)
    

This page documents the input components available in the Bubbles UI library. Input components allow users to enter and edit text in terminal-based applications built with the Bubble Tea framework. These components handle keyboard input, cursor management, text manipulation, and provide various styling options.

The Bubbles library provides two main input components:

1.  **[Text Input](https://deepwiki.com/charmbracelet/bubbles/2.1-text-input)
    ** — A single-line text input field.
2.  **[Text Area](https://deepwiki.com/charmbracelet/bubbles/2.2-text-area)
    ** — A multi-line text input field with scrolling and wrapping capabilities.

Component Architecture
----------------------

Input components in the Bubbles library follow the Model-View-Update (MVU) pattern. They implement the `tea.Model` interface and maintain their own internal state, including text values stored as rune slices to handle Unicode correctly.

### Entity Mapping: Interface to Implementation

The following diagram bridges the high-level `tea.Model` requirements with the specific implementations in `textinput` and `textarea`.

Sources: [textinput/textinput.go89-154](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L89-L154)
 [textarea/textarea.go187-280](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L187-L280)

Shared Concepts
---------------

While serving different layout needs, both components share a core set of primitives for cursor management and input sanitization.

### Virtual vs. Real Cursor

Both components utilize the `cursor` package primitive. They can operate in two modes:

*   **Virtual Cursor**: A software-rendered cursor managed within the component's `View()` output. This allows for custom blinking and styling via `cursor.Model` [textinput/textinput.go104-105](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L104-L105)
    
*   **Real Cursor**: Leveraging the terminal's native cursor via `tea.Cursor` instructions [textarea/textarea.go162-163](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L162-L163)
    

### Input Sanitization (runeutil)

To prevent control characters or incompatible whitespace (like tabs) from breaking terminal layouts, both components use `runeutil.Sanitizer`. This ensures that text entered or pasted is "cleaned" before being stored in the internal `[]rune` buffers [textinput/textinput.go144](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L144-L144)
 [textarea/textarea.go275](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L275-L275)

### Focus and Blinking

Focus management determines whether a component processes `tea.KeyPressMsg`. When focused, the cursor begins its blinking cycle, often initiated by the component's `Focus()` method returning a `cursor.Blink` command.

Sources: [textinput/textinput.go242-252](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L242-L252)
 [textarea/textarea.go581-592](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L581-L592)

Text Input Component
--------------------

The `textinput` package provides a single-line field. It is optimized for short entries and includes a built-in suggestion engine for autocompletion.

### Key Features

*   **EchoMode**: Supports `EchoNormal`, `EchoPassword` (masked), and `EchoNone` (hidden) for sensitive data [textinput/textinput.go28-41](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L28-L41)
    
*   **Suggestions**: Can display a list of `matchedSuggestions` that the user can cycle through and accept via `Tab` [textinput/textinput.go147-154](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L147-L154)
    
*   **Validation**: Supports a `ValidateFunc` that runs on input change, setting the `Err` field if the input fails custom logic [textinput/textinput.go43-44](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L43-L44)
    

For details, see [Text Input](https://deepwiki.com/charmbracelet/bubbles/2.1-text-input)
.

Sources: [textinput/textinput.go27-44](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L27-L44)
 [textinput/textinput.go147-154](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L147-L154)

Text Area Component
-------------------

The `textarea` package provides a multi-line input. Unlike the single-line version, it manages text as a 2D slice (`[][]rune`) and integrates a `viewport` for vertical and horizontal scrolling.

### Key Features

*   **Soft Wrapping**: Automatically wraps logical lines into visual lines based on the component's width [textarea/textarea.go111-138](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L111-L138)
    
*   **Viewport Integration**: Uses an internal `viewport.Model` to handle rendering of large text blocks that exceed the available height [textarea/textarea.go246](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L246-L246)
    
*   **Case Transformation**: Includes specialized keybindings for `UppercaseWordForward`, `LowercaseWordForward`, and `CapitalizeWordForward` [textarea/textarea.go69-71](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L69-L71)
    

For details, see [Text Area](https://deepwiki.com/charmbracelet/bubbles/2.2-text-area)
.

Sources: [textarea/textarea.go47-74](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L47-L74)
 [textarea/textarea.go246](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L246-L246)

Input Processing Flow
---------------------

The following diagram illustrates how a `tea.KeyPressMsg` is routed through the component's `Update` loop, comparing it against the `KeyMap` before falling back to character insertion.

Sources: [textinput/textinput.go553-647](https://github.com/charmbracelet/bubbles/blob/42130e89/textinput/textinput.go#L553-L647)
 [textarea/textarea.go957-1088](https://github.com/charmbracelet/bubbles/blob/42130e89/textarea/textarea.go#L957-L1088)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Input Components](https://deepwiki.com/charmbracelet/bubbles/2-input-components#input-components)
    
*   [Component Architecture](https://deepwiki.com/charmbracelet/bubbles/2-input-components#component-architecture)
    
*   [Entity Mapping: Interface to Implementation](https://deepwiki.com/charmbracelet/bubbles/2-input-components#entity-mapping-interface-to-implementation)
    
*   [Shared Concepts](https://deepwiki.com/charmbracelet/bubbles/2-input-components#shared-concepts)
    
*   [Virtual vs. Real Cursor](https://deepwiki.com/charmbracelet/bubbles/2-input-components#virtual-vs-real-cursor)
    
*   [Input Sanitization (runeutil)](https://deepwiki.com/charmbracelet/bubbles/2-input-components#input-sanitization-runeutil)
    
*   [Focus and Blinking](https://deepwiki.com/charmbracelet/bubbles/2-input-components#focus-and-blinking)
    
*   [Text Input Component](https://deepwiki.com/charmbracelet/bubbles/2-input-components#text-input-component)
    
*   [Key Features](https://deepwiki.com/charmbracelet/bubbles/2-input-components#key-features)
    
*   [Text Area Component](https://deepwiki.com/charmbracelet/bubbles/2-input-components#text-area-component)
    
*   [Key Features](https://deepwiki.com/charmbracelet/bubbles/2-input-components#key-features-1)
    
*   [Input Processing Flow](https://deepwiki.com/charmbracelet/bubbles/2-input-components#input-processing-flow)
