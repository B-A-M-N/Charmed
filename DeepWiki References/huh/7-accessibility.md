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

Accessibility
=============

Relevant source files

*   [accessibility/accessibility.go](https://github.com/charmbracelet/huh/blob/151ba059/accessibility/accessibility.go)
    
*   [field\_input.go](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go)
    
*   [field\_text.go](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go)
    
*   [form.go](https://github.com/charmbracelet/huh/blob/151ba059/form.go)
    
*   [group.go](https://github.com/charmbracelet/huh/blob/151ba059/group.go)
    
*   [internal/accessibility/accessibility.go](https://github.com/charmbracelet/huh/blob/151ba059/internal/accessibility/accessibility.go)
    
*   [spinner/examples/context-and-action-and-error/main.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/context-and-action-and-error/main.go)
    
*   [spinner/examples/context-and-action/main.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/context-and-action/main.go)
    
*   [spinner/examples/loading/main.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/loading/main.go)
    
*   [spinner/spinner.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go)
    
*   [spinner/spinner\_test.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner_test.go)
    
*   [theme.go](https://github.com/charmbracelet/huh/blob/151ba059/theme.go)
    

This document explains the accessibility features in the charmbracelet/huh library and how to build terminal forms that work well with screen readers and other assistive technologies.

Introduction
------------

The huh library includes built-in support for making terminal forms accessible to users who rely on screen readers or who are working in environments where standard terminal rendering isn't optimal. The library can automatically detect certain environments that may require accessibility features and adjust accordingly.

Sources: [form.go75-78](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L75-L78)
 [form.go234-237](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L234-L237)

How Accessible Mode Works
-------------------------

When accessible mode is enabled, the library switches from an interactive Bubble Tea TUI to a simpler sequential interface that's more compatible with screen readers. Instead of redrawing the screen for each interaction, the form displays each field one at a time and waits for user input before proceeding to the next field.

### Automatic Detection

The library automatically enables accessible mode in the following situations:

1.  When the terminal environment is detected as "dumb" (`TERM=dumb`)
2.  When explicitly enabled through the API

Sources: [form.go124-127](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L124-L127)

### Manual Configuration

You can explicitly enable or disable accessible mode on a form:

Sources: [form.go234-237](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L234-L237)

Form Accessibility
------------------

When a form runs in accessible mode, it processes each field sequentially:

1.  Each field is displayed with its prompt
2.  The user provides input
3.  The input is validated
4.  The form moves to the next field

This approach is more compatible with screen readers as it avoids the dynamic content updates that can be difficult for screen readers to interpret.

### Limitations in Accessible Mode

Some features are not available in accessible mode:

*   **Timeouts**: Form timeouts are not supported in accessible mode
*   **Dynamic navigation**: Users cannot freely navigate between fields
*   **Visual styling**: Most styling is ignored in accessible mode

Sources: [form.go702-720](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L702-L720)
 [form.go57-58](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L57-L58)

Field-Specific Accessibility
----------------------------

Each field type implements its own accessible version with appropriate prompts and handling:

### Text Input Fields

Text input fields use simple prompts in accessible mode:

    Title: [user types here]
    

For password fields, the library uses terminal features to avoid echoing characters, providing security while remaining accessible.

Sources: [field\_input.go432-464](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go#L432-L464)
 [field\_text.go417-440](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go#L417-L440)

### Selection Fields

For selection fields like `Select` and `MultiSelect`, the accessible mode presents options in a list format that's easier to navigate with screen readers.

### Confirmation Fields

Confirmation fields in accessible mode accept various formats of "yes"/"no" answers, with clear prompting:

    Do you agree? [y/N]: 
    

Spinner Accessibility
---------------------

The spinner component also supports accessible mode, which transforms the animated spinner into a static indicator:

    Loading... [static indicator]
    

Instead of showing animation frames that might confuse screen readers, the accessible spinner shows a consistent message.

Sources: [spinner/spinner.go106-110](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L106-L110)
 [spinner/spinner.go194-222](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L194-L222)

Implementation Details
----------------------

The accessibility implementation relies on several key components:

Sources: [form.go63-96](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L63-L96)
 [form.go138](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L138-L138)
 [internal/accessibility/accessibility.go](https://github.com/charmbracelet/huh/blob/151ba059/internal/accessibility/accessibility.go)

### Accessible Field Implementation

Each field implements `runAccessible` to handle its specific input requirements:

1.  **Input Fields**: Use `PromptString` for basic text and `PromptPassword` for passwords
2.  **Text Fields**: Handle multi-line text input in an accessible way
3.  **Select Fields**: Present options as a numbered list
4.  **Confirm Fields**: Use `PromptBool` to handle yes/no input

Sources: [field\_input.go432-464](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go#L432-L464)
 [field\_text.go417-440](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go#L417-L440)

Best Practices for Building Accessible Forms
--------------------------------------------

When building forms that need to be accessible:

1.  **Provide clear titles and descriptions**: These become the prompts in accessible mode
2.  **Add validation messages**: Make sure error messages are descriptive
3.  **Test with screen readers**: Verify that your forms work with common screen readers
4.  **Be mindful of timeouts**: Remember that timeouts aren't supported in accessible mode

Example: Creating an Accessible Form
------------------------------------

Here's a conceptual example of an accessible form:

Terminal Compatibility
----------------------

The accessible mode works best in:

1.  Screen reader environments
2.  SSH sessions with limited terminal capabilities
3.  Terminals marked as "dumb"
4.  Terminals without support for interactive features

Sources: [form.go124-127](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L124-L127)

Related Resources
-----------------

For more information about specific field types and their accessibility features, see [Field Types](https://deepwiki.com/charmbracelet/huh/3-field-types)
.

For information about integration with Bubble Tea (which is bypassed in accessible mode), see [Bubble Tea Integration](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration)
.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Accessibility](https://deepwiki.com/charmbracelet/huh/7-accessibility#accessibility)
    
*   [Introduction](https://deepwiki.com/charmbracelet/huh/7-accessibility#introduction)
    
*   [How Accessible Mode Works](https://deepwiki.com/charmbracelet/huh/7-accessibility#how-accessible-mode-works)
    
*   [Automatic Detection](https://deepwiki.com/charmbracelet/huh/7-accessibility#automatic-detection)
    
*   [Manual Configuration](https://deepwiki.com/charmbracelet/huh/7-accessibility#manual-configuration)
    
*   [Form Accessibility](https://deepwiki.com/charmbracelet/huh/7-accessibility#form-accessibility)
    
*   [Limitations in Accessible Mode](https://deepwiki.com/charmbracelet/huh/7-accessibility#limitations-in-accessible-mode)
    
*   [Field-Specific Accessibility](https://deepwiki.com/charmbracelet/huh/7-accessibility#field-specific-accessibility)
    
*   [Text Input Fields](https://deepwiki.com/charmbracelet/huh/7-accessibility#text-input-fields)
    
*   [Selection Fields](https://deepwiki.com/charmbracelet/huh/7-accessibility#selection-fields)
    
*   [Confirmation Fields](https://deepwiki.com/charmbracelet/huh/7-accessibility#confirmation-fields)
    
*   [Spinner Accessibility](https://deepwiki.com/charmbracelet/huh/7-accessibility#spinner-accessibility)
    
*   [Implementation Details](https://deepwiki.com/charmbracelet/huh/7-accessibility#implementation-details)
    
*   [Accessible Field Implementation](https://deepwiki.com/charmbracelet/huh/7-accessibility#accessible-field-implementation)
    
*   [Best Practices for Building Accessible Forms](https://deepwiki.com/charmbracelet/huh/7-accessibility#best-practices-for-building-accessible-forms)
    
*   [Example: Creating an Accessible Form](https://deepwiki.com/charmbracelet/huh/7-accessibility#example-creating-an-accessible-form)
    
*   [Terminal Compatibility](https://deepwiki.com/charmbracelet/huh/7-accessibility#terminal-compatibility)
    
*   [Related Resources](https://deepwiki.com/charmbracelet/huh/7-accessibility#related-resources)
