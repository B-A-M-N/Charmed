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

Bubble Tea Integration
======================

Relevant source files

*   [README.md](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1)
    
*   [examples/bubbletea-options/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea-options/main.go)
    
*   [examples/bubbletea/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go)
    
*   [examples/burger/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go)
    
*   [examples/multiple-groups/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/multiple-groups/main.go)
    
*   [form.go](https://github.com/charmbracelet/huh/blob/151ba059/form.go)
    
*   [group.go](https://github.com/charmbracelet/huh/blob/151ba059/group.go)
    
*   [theme.go](https://github.com/charmbracelet/huh/blob/151ba059/theme.go)
    

This document explains how to integrate the `huh` form library with Bubble Tea applications. `huh` offers first-class support for [Bubble Tea](https://github.com/charmbracelet/huh/blob/151ba059/Bubble%20Tea)
 allowing you to embed interactive forms within larger terminal user interfaces. For standalone form usage, see [Core Components](https://deepwiki.com/charmbracelet/huh/2-core-components)
.

Overview
--------

Bubble Tea is a Go framework for building terminal applications using the Elm Architecture. The `huh` library is designed to work seamlessly with Bubble Tea, providing a robust solution for form-based interactions within Bubble Tea applications.

Sources: [README.md403-410](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L403-L410)
 [form.go60-96](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L60-L96)

Integration Architecture
------------------------

The key to `huh` integration with Bubble Tea is that both `Form` and `Field` types implement the Bubble Tea `Model` interface, which requires three methods:

Sources: [form.go138-189](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L138-L189)
 [form.go500-524](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L500-L524)
 [form.go520-657](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L520-L657)

Basic Integration Pattern
-------------------------

To integrate a `huh` form into your Bubble Tea application, include it as a field in your application's model:

1.  Create the form and include it in your application model
2.  Initialize the form in your model's `Init()` method
3.  Update the form in your model's `Update()` method
4.  Include the form's view in your model's `View()` method
5.  Check the form's state to determine when it's completed

### Model Structure

Sources: [examples/bubbletea/main.go63-107](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L63-L107)

### Initialization

Sources: [examples/bubbletea/main.go109-111](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L109-L111)

### Update Handling

Sources: [examples/bubbletea/main.go120-148](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L120-L148)

### View Rendering

Sources: [examples/bubbletea/main.go150-218](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L150-L218)

Message Flow
------------

The following diagram illustrates how messages flow between your Bubble Tea application and the `huh` form:

Sources: [form.go520-625](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L520-L625)
 [group.go250-286](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L250-L286)

Form State Management
---------------------

The `Form` type has a `State` field that indicates the current state of the form:

You can check this state in your application's `Update` method to determine when the form has been completed or aborted:

Sources: [form.go36-48](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L36-L48)
 [examples/bubbletea/main.go142-145](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L142-L145)

Accessing Form Values
---------------------

Once the form is completed, you can access the values using the following methods:

*   `Get(key string) any`: Get a value by its key
*   `GetString(key string) string`: Get a string value by its key
*   `GetInt(key string) int`: Get an integer value by its key
*   `GetBool(key string) bool`: Get a boolean value by its key

Example:

Sources: [form.go440-470](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L440-L470)
 [examples/bubbletea/main.go155-156](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L155-L156)

Advanced Integration Options
----------------------------

### Custom Bubble Tea Program Options

You can pass custom Bubble Tea program options to your form:

Sources: [examples/bubbletea-options/main.go10-14](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea-options/main.go#L10-L14)
 [form.go348-353](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L348-L353)

### Custom Theming and Styling

Forms can be integrated with your application's theme:

See [Theme System](https://deepwiki.com/charmbracelet/huh/5-theme-system)
 for more detailed information about theming.

Sources: [form.go263-278](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L263-L278)
 [theme.go125-172](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L125-L172)

### Layout and Size Control

Control the width and height of the form to match your application layout:

Sources: [form.go296-325](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L296-L325)
 [examples/bubbletea/main.go102-104](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L102-L104)

Complete Integration Example
----------------------------

Here's a walkthrough of the key components from the example Bubble Tea application that integrates a `huh` form:

1.  **Application Model Structure**:
    
    *   Contains the form as a field
    *   Includes additional state for the application
2.  **Form Initialization**:
    
    *   Creates a form with fields for character creation
    *   Configures form options like width and help display
3.  **Update Logic**:
    
    *   Delegates form-related messages to the form
    *   Checks form state for completion
    *   Handles application-specific keys (e.g., quit, interrupt)
4.  **View Rendering**:
    
    *   Renders different UI based on form state
    *   Shows form when active, results when completed
    *   Combines form view with additional application UI
5.  **Form State Detection**:
    
    *   Uses form.State to detect completion
    *   Accesses form values using GetString() and other getters

The complete example demonstrates how to build a character creation form for a game, with job role information updated dynamically based on selected class and level.

Sources: [examples/bubbletea/main.go63-289](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L63-L289)

Best Practices
--------------

1.  **State Handling**
    
    *   Always check for proper type assertion when updating the form
    *   Handle all form states (normal, completed, aborted)
2.  **Message Passing**
    
    *   Let the form handle form-specific messages
    *   Handle application-specific messages in your app's Update method
3.  **Error Handling**
    
    *   Check for validation errors in the form
    *   Display errors appropriately in your UI
4.  **Form Completion**
    
    *   Use the form's State field to detect completion
    *   Access form values only after completion
5.  **Layout**
    
    *   Set appropriate width/height based on your application's layout
    *   Consider window size changes

Sources: [examples/bubbletea/main.go120-148](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L120-L148)
 [form.go36-48](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L36-L48)
 [form.go424-432](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L424-L432)

Conclusion
----------

The `huh` library provides a clean, elegant way to integrate forms into Bubble Tea applications. By implementing the Bubble Tea Model interface, forms can be seamlessly embedded within larger terminal UI applications while maintaining the component-based architecture that makes Bubble Tea powerful.

With proper integration, you can create complex interactive forms within your Bubble Tea applications, improving the user experience for data collection and configuration tasks.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Bubble Tea Integration](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#bubble-tea-integration)
    
*   [Overview](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#overview)
    
*   [Integration Architecture](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#integration-architecture)
    
*   [Basic Integration Pattern](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#basic-integration-pattern)
    
*   [Model Structure](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#model-structure)
    
*   [Initialization](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#initialization)
    
*   [Update Handling](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#update-handling)
    
*   [View Rendering](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#view-rendering)
    
*   [Message Flow](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#message-flow)
    
*   [Form State Management](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#form-state-management)
    
*   [Accessing Form Values](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#accessing-form-values)
    
*   [Advanced Integration Options](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#advanced-integration-options)
    
*   [Custom Bubble Tea Program Options](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#custom-bubble-tea-program-options)
    
*   [Custom Theming and Styling](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#custom-theming-and-styling)
    
*   [Layout and Size Control](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#layout-and-size-control)
    
*   [Complete Integration Example](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#complete-integration-example)
    
*   [Best Practices](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#best-practices)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration#conclusion)
