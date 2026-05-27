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

Theme System
============

Relevant source files

*   [examples/filepicker/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/filepicker/main.go)
    
*   [examples/gh/create.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/gh/create.go)
    
*   [examples/theme/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/theme/main.go)
    
*   [form.go](https://github.com/charmbracelet/huh/blob/151ba059/form.go)
    
*   [group.go](https://github.com/charmbracelet/huh/blob/151ba059/group.go)
    
*   [theme.go](https://github.com/charmbracelet/huh/blob/151ba059/theme.go)
    

This document details the theme system in the Charmbracelet/Huh library, explaining how styling is applied to forms and their components. The theme system provides consistent visual styling across form elements while allowing for customization.

Overview
--------

The Huh theme system uses [Lip Gloss](https://github.com/charmbracelet/huh/blob/151ba059/Lip%20Gloss)
 to style terminal UI components. It provides a structured approach to styling different elements of forms, groups, and fields with consistent visual language.

Sources: [theme.go9-78](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L9-L78)
 [form.go636-644](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L636-L644)
 [group.go90-102](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L90-L102)

Theme Structure
---------------

The theme system is built around the `Theme` struct, which contains styles for all form components. These styles cascade from the form level down to individual fields.

### Core Theme Components

Sources: [theme.go9-78](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L9-L78)

Built-in Themes
---------------

The Huh library provides several built-in themes to get started quickly. Each theme provides a different visual style tailored to specific use cases.

| Theme Function | Description | Base Style |
| --- | --- | --- |
| `ThemeBase()` | Provides a foundational theme with minimal styling | Simple black and white |
| `ThemeCharm()` | Default theme with Charm branding colors | Indigo, fuchsia, and green accents |
| `ThemeDracula()` | Dark theme based on the Dracula color scheme | Purple, green, and yellow on dark background |
| `ThemeBase16()` | Theme using Base16 color palette | 16-color terminal palette |
| `ThemeCatppuccin()` | Theme based on Catppuccin color scheme | Pastel colors with soft contrast |

To use a built-in theme:

Sources: [theme.go84-329](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L84-L329)
 [examples/theme/main.go10-19](https://github.com/charmbracelet/huh/blob/151ba059/examples/theme/main.go#L10-L19)

Theme Inheritance
-----------------

Themes are organized in an inheritance hierarchy, with `ThemeBase` providing the foundation for all other themes. Custom themes can extend built-in themes by modifying specific styles.

Sources: [theme.go84-123](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L84-L123)
 [theme.go125-172](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L125-L172)
 [theme.go174-223](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L174-L223)
 [theme.go225-268](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L225-L268)
 [theme.go270-329](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L270-L329)

Theme Application
-----------------

Themes can be applied at different levels, from the entire form down to individual fields. The theme system uses a cascading approach, where styles specified at lower levels override those at higher levels.

### Theme Cascade Flow

Sources: [form.go263-278](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L263-L278)
 [group.go90-102](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L90-L102)

Theme Processing Flow
---------------------

When a theme is applied to a form, the theme cascades down through the form hierarchy. Each component (`Form`, `Group`, and `Field`) processes the theme according to its needs.

Sources: [form.go635-644](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L635-L644)
 [group.go288-295](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L288-L295)
 [form.go647-653](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L647-L653)

Creating Custom Themes
----------------------

Custom themes can be created by extending an existing theme or creating a new one from scratch.

### Extending an Existing Theme

The recommended approach is to start with an existing theme like `ThemeBase()` or `ThemeCharm()` and modify specific aspects:

### Creating a Theme From Scratch

For complete control, you can create a theme from scratch by initializing a new `Theme` struct and setting all required styles:

Sources: [examples/theme/main.go11-16](https://github.com/charmbracelet/huh/blob/151ba059/examples/theme/main.go#L11-L16)
 [examples/gh/create.go29-31](https://github.com/charmbracelet/huh/blob/151ba059/examples/gh/create.go#L29-L31)

Theme Components in Detail
--------------------------

### FormStyles

FormStyles apply to the entire form and provide the base container styling:

### GroupStyles

GroupStyles apply to each group within a form:

### FieldStyles

FieldStyles are the most extensive, as they apply to the various field types. They are divided into "Focused" and "Blurred" variants to style fields differently based on their focus state:

### TextInputStyles

TextInputStyles provide detailed styling for text input fields:

Sources: [theme.go20-77](https://github.com/charmbracelet/huh/blob/151ba059/theme.go#L20-L77)

Theme System Integration
------------------------

The theme system integrates with the form system through the `WithTheme()` methods on Form, Group, and Field. These methods cascade the theme down through the component hierarchy.

Sources: [form.go263-278](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L263-L278)
 [form.go635-644](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L635-L644)
 [group.go90-102](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L90-L102)
 [group.go288-295](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L288-L295)

Theme System Default Fallbacks
------------------------------

The theme system implements a fallback mechanism to ensure that components always have styles, even if not explicitly provided:

Sources: [form.go635-640](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L635-L640)
 [group.go288-293](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L288-L293)

Example Use Cases
-----------------

### Basic Theme Application

### Mixed Theme Application

### Custom Theme Modification

Sources: [examples/gh/create.go29-31](https://github.com/charmbracelet/huh/blob/151ba059/examples/gh/create.go#L29-L31)
 [examples/theme/main.go11-19](https://github.com/charmbracelet/huh/blob/151ba059/examples/theme/main.go#L11-L19)

Conclusion
----------

The theme system in Charmbracelet/Huh provides a flexible and powerful way to style terminal UI forms. By using a hierarchical approach with cascading styles, it allows for consistent styling across components while enabling customization at any level. The built-in themes provide ready-to-use styling options, and the ability to create custom themes offers unlimited flexibility for application-specific designs.

For information about integrating forms with Bubble Tea applications, see [Bubble Tea Integration](https://deepwiki.com/charmbracelet/huh/8-bubble-tea-integration)
.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Theme System](https://deepwiki.com/charmbracelet/huh/5-theme-system#theme-system)
    
*   [Overview](https://deepwiki.com/charmbracelet/huh/5-theme-system#overview)
    
*   [Theme Structure](https://deepwiki.com/charmbracelet/huh/5-theme-system#theme-structure)
    
*   [Core Theme Components](https://deepwiki.com/charmbracelet/huh/5-theme-system#core-theme-components)
    
*   [Built-in Themes](https://deepwiki.com/charmbracelet/huh/5-theme-system#built-in-themes)
    
*   [Theme Inheritance](https://deepwiki.com/charmbracelet/huh/5-theme-system#theme-inheritance)
    
*   [Theme Application](https://deepwiki.com/charmbracelet/huh/5-theme-system#theme-application)
    
*   [Theme Cascade Flow](https://deepwiki.com/charmbracelet/huh/5-theme-system#theme-cascade-flow)
    
*   [Theme Processing Flow](https://deepwiki.com/charmbracelet/huh/5-theme-system#theme-processing-flow)
    
*   [Creating Custom Themes](https://deepwiki.com/charmbracelet/huh/5-theme-system#creating-custom-themes)
    
*   [Extending an Existing Theme](https://deepwiki.com/charmbracelet/huh/5-theme-system#extending-an-existing-theme)
    
*   [Creating a Theme From Scratch](https://deepwiki.com/charmbracelet/huh/5-theme-system#creating-a-theme-from-scratch)
    
*   [Theme Components in Detail](https://deepwiki.com/charmbracelet/huh/5-theme-system#theme-components-in-detail)
    
*   [FormStyles](https://deepwiki.com/charmbracelet/huh/5-theme-system#formstyles)
    
*   [GroupStyles](https://deepwiki.com/charmbracelet/huh/5-theme-system#groupstyles)
    
*   [FieldStyles](https://deepwiki.com/charmbracelet/huh/5-theme-system#fieldstyles)
    
*   [TextInputStyles](https://deepwiki.com/charmbracelet/huh/5-theme-system#textinputstyles)
    
*   [Theme System Integration](https://deepwiki.com/charmbracelet/huh/5-theme-system#theme-system-integration)
    
*   [Theme System Default Fallbacks](https://deepwiki.com/charmbracelet/huh/5-theme-system#theme-system-default-fallbacks)
    
*   [Example Use Cases](https://deepwiki.com/charmbracelet/huh/5-theme-system#example-use-cases)
    
*   [Basic Theme Application](https://deepwiki.com/charmbracelet/huh/5-theme-system#basic-theme-application)
    
*   [Mixed Theme Application](https://deepwiki.com/charmbracelet/huh/5-theme-system#mixed-theme-application)
    
*   [Custom Theme Modification](https://deepwiki.com/charmbracelet/huh/5-theme-system#custom-theme-modification)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/huh/5-theme-system#conclusion)
