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

Example Applications
====================

Relevant source files

*   [examples/bubbletea-options/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea-options/main.go)
    
*   [examples/bubbletea/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go)
    
*   [examples/multiple-groups/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/multiple-groups/main.go)
    

This page showcases complete example applications built with the charmbracelet/huh library. It demonstrates practical implementations of the form system in various contexts, illustrating different integration patterns and usage scenarios. For information about advanced usage techniques like dynamic forms and validation, see [Advanced Usage](https://deepwiki.com/charmbracelet/huh/9-advanced-usage)
.

Integration Patterns
--------------------

The huh library can be integrated into applications in several ways, ranging from simple standalone forms to complex interactive experiences embedded within Bubble Tea applications.

Sources: [examples/multiple-groups/main.go10-79](https://github.com/charmbracelet/huh/blob/151ba059/examples/multiple-groups/main.go#L10-L79)
 [examples/bubbletea/main.go63-108](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L63-L108)
 [examples/bubbletea-options/main.go10-22](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea-options/main.go#L10-L22)

Simple Standalone Form Example
------------------------------

The most basic way to use the huh library is through a standalone form that directly runs in the terminal. The library handles all input processing and rendering.

### Multiple Group Example

This example demonstrates creating a form with multiple groups, each containing selection fields with different heights.

In this example, the application creates a form with three groups, each containing either a `Select` or `MultiSelect` field with various options. The groups are given different heights using `WithHeight()`, allowing for customized display areas. Users can navigate between groups using keyboard controls.

Sources: [examples/multiple-groups/main.go10-79](https://github.com/charmbracelet/huh/blob/151ba059/examples/multiple-groups/main.go#L10-L79)

Bubble Tea Integration Examples
-------------------------------

The huh library can be integrated into Bubble Tea applications in two ways: simple integration with program options or complex integration with custom models.

### Simple Bubble Tea Integration

This example shows how to create a form that uses Bubble Tea's alt screen for a cleaner user experience.

The application creates a simple form with a single input field, enabling the alt screen through Bubble Tea's program options. After the form completes, it accesses the collected value and displays it.

Sources: [examples/bubbletea-options/main.go10-22](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea-options/main.go#L10-L22)

### Advanced Bubble Tea Integration

For more complex applications, the huh form can be embedded within a custom Bubble Tea model, allowing for sophisticated UI composition and state management.

#### Component Structure

The advanced Bubble Tea example demonstrates how to create a rich interactive form experience with custom styling and layout.

This example creates a custom "Charm Employment Application" with:

1.  A custom Model that implements the Bubble Tea Model interface
2.  Embedded huh.Form with multiple fields
3.  Sophisticated styling using lipgloss
4.  Complex view composition with header, form, status, and footer sections
5.  Custom state management and error handling
6.  Dynamic content based on user input

The application demonstrates several advanced techniques:

*   Custom styling and layout with lipgloss
*   Side-by-side display of form and status information
*   Dynamic content generation based on form input
*   Custom error handling and display
*   Integration with Bubble Tea's message system

Sources: [examples/bubbletea/main.go13-289](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L13-L289)

Form State Flow
---------------

All example applications follow a similar state flow for form processing, though with varying degrees of complexity.

In the simplest examples, this flow is handled internally by the huh library. In more complex Bubble Tea integrations, the application manages the form state within the Update method of a custom model.

Sources: [examples/bubbletea/main.go120-148](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L120-L148)
 [examples/multiple-groups/main.go76-78](https://github.com/charmbracelet/huh/blob/151ba059/examples/multiple-groups/main.go#L76-L78)
 [examples/bubbletea-options/main.go16-21](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea-options/main.go#L16-L21)

Application Customization Techniques
------------------------------------

The example applications demonstrate several techniques for customizing huh forms:

| Technique | Description | Example Location |
| --- | --- | --- |
| Custom Styling | Creating a custom style system with lipgloss | [examples/bubbletea/main.go21-54](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L21-L54) |
| Multiple Groups | Organizing fields into multiple navigable groups | [examples/multiple-groups/main.go10-74](https://github.com/charmbracelet/huh/blob/151ba059/examples/multiple-groups/main.go#L10-L74) |
| Form Width | Setting explicit form width | [examples/bubbletea/main.go103](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L103-L103) |
| Custom Validation | Implementing field validation logic | [examples/bubbletea/main.go93-97](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L93-L97) |
| Dynamic Content | Generating content based on form state | [examples/bubbletea/main.go248-281](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L248-L281) |
| Custom Layout | Creating side-by-side layout with form and status | [examples/bubbletea/main.go169-202](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L169-L202) |
| Bubble Tea Options | Applying program options like alt screen | [examples/bubbletea-options/main.go14](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea-options/main.go#L14-L14) |

Sources: [examples/bubbletea/main.go21-281](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L21-L281)
 [examples/multiple-groups/main.go10-74](https://github.com/charmbracelet/huh/blob/151ba059/examples/multiple-groups/main.go#L10-L74)
 [examples/bubbletea-options/main.go10-22](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea-options/main.go#L10-L22)

Practical Implementation Patterns
---------------------------------

The example applications demonstrate several common patterns for implementing forms with the huh library:

1.  **Simple Input Collection**: Gather a single piece of information with minimal UI
2.  **Multi-Option Selection**: Present users with lists of options to choose from
3.  **Multi-Step Forms**: Guide users through a sequence of input sections
4.  **Interactive Feedback**: Provide dynamic feedback based on user input
5.  **Custom Styling**: Create visually distinct interfaces with consistent design

These patterns can be combined and extended to create sophisticated terminal user interfaces for various applications, from simple configuration tools to complex interactive experiences.

Sources: [examples/bubbletea/main.go76-105](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea/main.go#L76-L105)
 [examples/multiple-groups/main.go11-74](https://github.com/charmbracelet/huh/blob/151ba059/examples/multiple-groups/main.go#L11-L74)
 [examples/bubbletea-options/main.go12-14](https://github.com/charmbracelet/huh/blob/151ba059/examples/bubbletea-options/main.go#L12-L14)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Example Applications](https://deepwiki.com/charmbracelet/huh/10-example-applications#example-applications)
    
*   [Integration Patterns](https://deepwiki.com/charmbracelet/huh/10-example-applications#integration-patterns)
    
*   [Simple Standalone Form Example](https://deepwiki.com/charmbracelet/huh/10-example-applications#simple-standalone-form-example)
    
*   [Multiple Group Example](https://deepwiki.com/charmbracelet/huh/10-example-applications#multiple-group-example)
    
*   [Bubble Tea Integration Examples](https://deepwiki.com/charmbracelet/huh/10-example-applications#bubble-tea-integration-examples)
    
*   [Simple Bubble Tea Integration](https://deepwiki.com/charmbracelet/huh/10-example-applications#simple-bubble-tea-integration)
    
*   [Advanced Bubble Tea Integration](https://deepwiki.com/charmbracelet/huh/10-example-applications#advanced-bubble-tea-integration)
    
*   [Component Structure](https://deepwiki.com/charmbracelet/huh/10-example-applications#component-structure)
    
*   [Form State Flow](https://deepwiki.com/charmbracelet/huh/10-example-applications#form-state-flow)
    
*   [Application Customization Techniques](https://deepwiki.com/charmbracelet/huh/10-example-applications#application-customization-techniques)
    
*   [Practical Implementation Patterns](https://deepwiki.com/charmbracelet/huh/10-example-applications#practical-implementation-patterns)
