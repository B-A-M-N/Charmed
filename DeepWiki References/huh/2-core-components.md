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

Core Components
===============

Relevant source files

*   [form.go](https://github.com/charmbracelet/huh/blob/151ba059/form.go)
    
*   [group.go](https://github.com/charmbracelet/huh/blob/151ba059/group.go)
    
*   [huh\_test.go](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go)
    
*   [theme.go](https://github.com/charmbracelet/huh/blob/151ba059/theme.go)
    

This page documents the primary building blocks of the Charmbracelet/Huh form system: Form, Group, and Field. These core components work together to create interactive forms in terminal user interfaces.

For information about specific field types like Input, Select, etc., see [Field Types](https://deepwiki.com/charmbracelet/huh/3-field-types)
. For details on customizing appearance, see [Layout System](https://deepwiki.com/charmbracelet/huh/4-layout-system)
 and [Theme System](https://deepwiki.com/charmbracelet/huh/5-theme-system)
.

Core Components Overview
------------------------

Huh is organized around three primary components that form a hierarchical relationship:

**Component Hierarchy Diagram**

Sources: [form.go63-96](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L63-L96)
 [group.go14-44](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L14-L44)
 [form.go132-189](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L132-L189)

*   **Form**: The top-level container that manages a collection of groups, displayed one at a time as "pages"
*   **Group**: A container for related fields that are displayed together on a single page of the form
*   **Field**: An interface representing a single input control (text input, select, etc.)

Form Component
--------------

The Form component is the main container that manages navigation between groups and the overall state of the form.

**Form Structure Diagram**

Sources: [form.go37-47](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L37-L47)
 [form.go63-96](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L63-L96)
 [form.go440-494](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L440-L494)

Forms provide several key functions:

1.  **Group Management**: Forms organize and navigate between groups of fields
2.  **State Tracking**: Forms track completion state (normal, completed, or aborted)
3.  **Data Collection**: Forms store field values and provide methods to retrieve them
4.  **Configuration**: Forms can be customized with themes, layouts, and key mappings

### Creating and Running Forms

Forms are created using the `NewForm` function:

Sources: [form.go98-130](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L98-L130)
 [form.go654-699](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L654-L699)
 [form.go440-470](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L440-L470)

Group Component
---------------

Groups organize related fields and display them together on a single page of the form.

**Group Structure Diagram**

Sources: [group.go14-44](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L14-L44)
 [group.go46-64](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L46-L64)
 [group.go153-164](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L153-L164)

Key features of the Group component:

1.  **Field Collection**: Groups manage a collection of fields that are navigated together
2.  **Metadata**: Groups can have titles and descriptions that describe their purpose
3.  **Error Management**: Groups collect and display errors from their fields
4.  **Conditional Display**: Groups can be conditionally shown or hidden
5.  **Navigation**: Groups handle navigation between their fields

### Group Creation and Configuration

Groups can be conditionally shown or hidden:

Sources: [group.go46-64](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L46-L64)
 [group.go66-75](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L66-L75)
 [group.go141-150](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L141-L150)

Field Interface
---------------

The Field interface defines the contract that all form input controls must implement.

**Field Interface Diagram**

Sources: [form.go132-189](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L132-L189)
 [form.go192-210](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L192-L210)

The Field interface includes methods for:

1.  **Bubble Tea Integration**: Fields implement the Bubble Tea Model interface
2.  **Focus Management**: Fields can be focused and blurred as the user navigates
3.  **Validation**: Fields can validate input and report errors
4.  **Customization**: Fields can be customized with themes and other options
5.  **Data Access**: Fields provide access to their values

Various field implementations provide different types of input controls like text inputs, select lists, confirmation prompts, etc. For details on specific field types, see [Field Types](https://deepwiki.com/charmbracelet/huh/3-field-types)
.

Component Lifecycle and Interaction
-----------------------------------

The core components interact through a well-defined lifecycle that leverages the Bubble Tea architecture:

**Component Interaction Diagram**

Sources: [form.go502-518](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L502-L518)
 [form.go521-625](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L521-L625)
 [form.go647-653](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L647-L653)
 [group.go195-215](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L195-L215)
 [group.go252-286](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L252-L286)
 [group.go361-372](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L361-L372)

This lifecycle can be summarized as:

1.  **Initialization**: The form initializes groups, which initialize fields
2.  **Updates**: Messages flow down from form to active group to focused field
3.  **Views**: Rendering builds up from field to group to form
4.  **Navigation**: Special messages trigger navigation between fields and groups
5.  **Completion**: When all groups are completed, the form is submitted

Customization Options
---------------------

Core components provide numerous customization options:

| Component | Method | Description |
| --- | --- | --- |
| Form | `WithTheme(*Theme)` | Set the visual theme for the form |
| Form | `WithLayout(Layout)` | Define how groups are arranged |
| Form | `WithKeyMap(*KeyMap)` | Set keyboard shortcuts |
| Form | `WithWidth(int)` | Set form width |
| Form | `WithHeight(int)` | Set form height |
| Form | `WithAccessible(bool)` | Enable screen reader support |
| Form | `WithTimeout(time.Duration)` | Set form timeout |
| Group | `Title(string)` | Set group title |
| Group | `Description(string)` | Set group description |
| Group | `WithHide(bool)` | Hide or show the group |
| Group | `WithHideFunc(func() bool)` | Conditionally hide the group |
| Field | `WithTheme(*Theme)` | Set field-specific theme |
| Field | `WithWidth(int)` | Set field width |
| Field | `WithHeight(int)` | Set field height |

Sources: [form.go230-362](https://github.com/charmbracelet/huh/blob/151ba059/form.go#L230-L362)
 [group.go66-151](https://github.com/charmbracelet/huh/blob/151ba059/group.go#L66-L151)

### Form Layout Example

Example: Complete Form
----------------------

Here's a complete example showing how the core components work together:

Sources: [huh\_test.go26-277](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go#L26-L277)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core Components](https://deepwiki.com/charmbracelet/huh/2-core-components#core-components)
    
*   [Core Components Overview](https://deepwiki.com/charmbracelet/huh/2-core-components#core-components-overview)
    
*   [Form Component](https://deepwiki.com/charmbracelet/huh/2-core-components#form-component)
    
*   [Creating and Running Forms](https://deepwiki.com/charmbracelet/huh/2-core-components#creating-and-running-forms)
    
*   [Group Component](https://deepwiki.com/charmbracelet/huh/2-core-components#group-component)
    
*   [Group Creation and Configuration](https://deepwiki.com/charmbracelet/huh/2-core-components#group-creation-and-configuration)
    
*   [Field Interface](https://deepwiki.com/charmbracelet/huh/2-core-components#field-interface)
    
*   [Component Lifecycle and Interaction](https://deepwiki.com/charmbracelet/huh/2-core-components#component-lifecycle-and-interaction)
    
*   [Customization Options](https://deepwiki.com/charmbracelet/huh/2-core-components#customization-options)
    
*   [Form Layout Example](https://deepwiki.com/charmbracelet/huh/2-core-components#form-layout-example)
    
*   [Example: Complete Form](https://deepwiki.com/charmbracelet/huh/2-core-components#example-complete-form)
