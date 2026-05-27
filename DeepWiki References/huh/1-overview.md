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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1)
    
*   [examples/burger/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go)
    
*   [examples/go.mod](https://github.com/charmbracelet/huh/blob/151ba059/examples/go.mod)
    
*   [examples/go.sum](https://github.com/charmbracelet/huh/blob/151ba059/examples/go.sum)
    
*   [go.mod](https://github.com/charmbracelet/huh/blob/151ba059/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/huh/blob/151ba059/go.sum)
    
*   [huh\_test.go](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go)
    

Charmbracelet/Huh is a Go library for building interactive forms and prompts in terminal user interfaces. It provides a simple yet powerful API for creating rich, interactive terminal forms with various field types, validation, and styling. This page provides a high-level overview of the library's purpose, architecture, and core components.

Sources: [README.md11-17](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L11-L17)
 [README.md1-11](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L1-L11)

Purpose and Scope
-----------------

Huh enables developers to easily create forms for terminal applications with features including:

*   Creating standalone interactive forms
*   Building multi-page form workflows
*   Collecting user input through various field types
*   Supporting accessible mode for screen readers
*   Integrating with Bubble Tea applications

Rather than building complex terminal UI logic from scratch, Huh provides ready-to-use components that handle user interaction, validation, and rendering.

Sources: [README.md11-17](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L11-L17)
 [README.md156-180](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L156-L180)

High-Level Architecture
-----------------------

The following diagram illustrates the core components of Huh and their relationships:

Sources: [README.md134-151](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L134-L151)
 [huh\_test.go26-278](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go#L26-L278)

Huh is built around a hierarchical structure where:

*   Forms contain one or more Groups
*   Groups contain one or more Fields
*   Fields implement specific input types (Input, Select, etc.)
*   The system integrates with Bubble Tea for terminal UI rendering and Lip Gloss for styling

Core Components
---------------

### Form

The Form is the top-level container that orchestrates the entire form experience. It manages form state, navigation between groups, validation, and rendering.

Key capabilities:

*   Managing multiple groups of fields
*   Handling form navigation
*   Collecting and validating user input
*   Applying themes and layouts
*   Supporting accessibility mode

Sources: [README.md46-109](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L46-L109)
 [huh\_test.go42-109](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go#L42-L109)

### Group

Groups act as containers for fields and represent pages or sections within a form. A form with multiple groups creates a multi-page experience where users navigate between groups.

Key capabilities:

*   Organizing related fields
*   Adding titles and descriptions to sections
*   Controlling visibility with conditional logic

Sources: [README.md42-45](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L42-L45)
 [README.md86-108](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L86-L108)
 [huh\_test.go64-86](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go#L64-L86)

### Field Interface

The Field interface defines the contract that all form field types must implement. This allows different field types to be used interchangeably within groups.

Sources: [huh\_test.go26-278](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go#L26-L278)
 [README.md134-151](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L134-L151)

Field Types
-----------

The library provides several field types to handle different types of user input:

| Field Type | Purpose | Example Use Case |
| --- | --- | --- |
| Input | Single-line text input | Names, email addresses |
| Text | Multi-line text input | Comments, descriptions |
| Select | Single selection from options | Choosing one item from a list |
| MultiSelect | Multiple selection from options | Selecting multiple items |
| Confirm | Yes/No choices | Confirmations, agreements |
| Note | Display information | Instructions, headings |
| FilePicker | File system selection | Selecting files or directories |

Sources: [README.md134-151](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L134-L151)
 [README.md157-236](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L157-L236)
 [README.md237-268](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L237-L268)
 [README.md269-306](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L269-L306)
 [README.md307-350](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L307-L350)

Usage Patterns
--------------

The following diagram illustrates the typical flow of form creation and usage:

Sources: [README.md26-125](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L26-L125)
 [examples/burger/main.go59-152](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L59-L152)
 [README.md155-401](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L155-L401)

Form Creation and Execution
---------------------------

Creating a form in Huh follows a builder pattern:

1.  Define variables to store form responses
2.  Create fields with appropriate field types
3.  Group related fields together
4.  Create a form from the groups
5.  Run the form to collect user input

The form's `Run()` method blocks until the user completes the form or the form is otherwise terminated.

Sources: [README.md26-125](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L26-L125)
 [examples/burger/main.go153-158](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L153-L158)

Integration with Bubble Tea
---------------------------

Huh can be used standalone or integrated into Bubble Tea applications. For standalone use, the `Run()` method handles the terminal lifecycle. For Bubble Tea integration, Huh implements the Bubble Tea Model interface, allowing it to be composed within larger Bubble Tea applications.

Sources: [README.md402-466](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L402-L466)
 [README.md416-463](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L416-L463)

Themes and Styling
------------------

Huh provides built-in themes and styling options through the Lip Gloss library:

*   Default Theme
*   Charm Theme
*   Dracula Theme
*   Catppuccin Theme
*   Base 16 Theme

Themes control colors, borders, padding, and other visual aspects of the form.

Sources: [README.md258-281](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L258-L281)
 [go.mod7-18](https://github.com/charmbracelet/huh/blob/151ba059/go.mod#L7-L18)

Accessibility Support
---------------------

Huh includes a dedicated accessibility mode designed for screen readers. When enabled, the library switches from TUI-based interfaces to standard prompts that are more compatible with screen readers and other assistive technologies.

Sources: [README.md238-256](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L238-L256)
 [huh\_test.go1129-1147](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go#L1129-L1147)
 [huh\_test.go1149-1378](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go#L1149-L1378)

Additional Features
-------------------

### Dynamic Forms

Forms can respond dynamically to user input by using functions to compute properties on the fly. This enables creating conditional form flows where fields, options, or content changes based on previous input.

Sources: [README.md283-351](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L283-L351)
 [README.md323-342](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L323-L342)

### Spinner Component

Huh includes a standalone spinner component for indicating background activity, particularly useful for showing progress after form submission.

Sources: [README.md355-396](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L355-L396)
 [examples/burger/main.go164](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L164-L164)

* * *

This overview provides a foundation for understanding the Huh library. For more detailed information about specific components, see the following pages:

*   Core Components: [Core Components](https://deepwiki.com/charmbracelet/huh/2-core-components)
    
*   Field Types: [Field Types](https://deepwiki.com/charmbracelet/huh/3-field-types)
    
*   Layout System: [Layout System](https://deepwiki.com/charmbracelet/huh/4-layout-system)
    
*   Theme System: [Theme System](https://deepwiki.com/charmbracelet/huh/5-theme-system)
    
*   Bubble Tea Integration: [Bubble Tea Integration](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration)
    

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/huh/1-overview#overview)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/huh/1-overview#purpose-and-scope)
    
*   [High-Level Architecture](https://deepwiki.com/charmbracelet/huh/1-overview#high-level-architecture)
    
*   [Core Components](https://deepwiki.com/charmbracelet/huh/1-overview#core-components)
    
*   [Form](https://deepwiki.com/charmbracelet/huh/1-overview#form)
    
*   [Group](https://deepwiki.com/charmbracelet/huh/1-overview#group)
    
*   [Field Interface](https://deepwiki.com/charmbracelet/huh/1-overview#field-interface)
    
*   [Field Types](https://deepwiki.com/charmbracelet/huh/1-overview#field-types)
    
*   [Usage Patterns](https://deepwiki.com/charmbracelet/huh/1-overview#usage-patterns)
    
*   [Form Creation and Execution](https://deepwiki.com/charmbracelet/huh/1-overview#form-creation-and-execution)
    
*   [Integration with Bubble Tea](https://deepwiki.com/charmbracelet/huh/1-overview#integration-with-bubble-tea)
    
*   [Themes and Styling](https://deepwiki.com/charmbracelet/huh/1-overview#themes-and-styling)
    
*   [Accessibility Support](https://deepwiki.com/charmbracelet/huh/1-overview#accessibility-support)
    
*   [Additional Features](https://deepwiki.com/charmbracelet/huh/1-overview#additional-features)
    
*   [Dynamic Forms](https://deepwiki.com/charmbracelet/huh/1-overview#dynamic-forms)
    
*   [Spinner Component](https://deepwiki.com/charmbracelet/huh/1-overview#spinner-component)
