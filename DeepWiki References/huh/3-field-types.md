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

Field Types
===========

Relevant source files

*   [README.md](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1)
    
*   [accessibility/accessibility.go](https://github.com/charmbracelet/huh/blob/151ba059/accessibility/accessibility.go)
    
*   [examples/burger/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go)
    
*   [field\_input.go](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go)
    
*   [field\_multiselect.go](https://github.com/charmbracelet/huh/blob/151ba059/field_multiselect.go)
    
*   [field\_select.go](https://github.com/charmbracelet/huh/blob/151ba059/field_select.go)
    
*   [field\_text.go](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go)
    
*   [huh\_test.go](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go)
    
*   [internal/accessibility/accessibility.go](https://github.com/charmbracelet/huh/blob/151ba059/internal/accessibility/accessibility.go)
    

This page provides an overview of the various field types available in the `huh` library. Each field type implements the `Field` interface but provides specialized functionality for different types of user input or display needs.

For information about the core `Field` interface itself, see [Field Interface](https://deepwiki.com/charmbracelet/huh/2.3-field-interface)
.

Field Type Overview
-------------------

Sources:

*   [README.md136-141](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L136-L141)
    
*   [field\_input.go](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go)
    
*   [field\_text.go](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go)
    
*   [field\_select.go](https://github.com/charmbracelet/huh/blob/151ba059/field_select.go)
    
*   [field\_multiselect.go](https://github.com/charmbracelet/huh/blob/151ba059/field_multiselect.go)
    

Common Structure and Features
-----------------------------

All field types in `huh` follow a consistent pattern and share several common features:

### Constructor Pattern

Each field type has a `New*` constructor function (e.g., `NewInput()`, `NewText()`) that returns a new instance of the field.

### Builder Pattern

Fields use the builder pattern, where each method returns the field itself, allowing for method chaining:

### Common Methods

All field types implement these common methods from the `Field` interface:

| Method | Purpose |
| --- | --- |
| `Focus()` | Focuses the field and returns any tea commands |
| `Blur()` | Blurs the field and returns any tea commands |
| `Update()` | Handles updates to the field's state |
| `View()` | Renders the field to a string |
| `GetKey()` | Returns the field's key |
| `GetValue()` | Returns the field's current value |
| `Error()` | Returns any validation error |

### Common Configuration Methods

All field types also provide these common configuration methods:

| Method | Purpose |
| --- | --- |
| `Key(string)` | Sets a unique identifier for retrieving the field's value |
| `Title(string)` | Sets the field's title |
| `TitleFunc(func() string, bindings)` | Sets a dynamic title function |
| `Description(string)` | Sets the field's description |
| `DescriptionFunc(func() string, bindings)` | Sets a dynamic description function |
| `Value(pointer)` | Sets the pointer where the field's value will be stored |
| `Validate(func(T) error)` | Sets a validation function for the field's value |

Sources:

*   [field\_input.go72-236](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go#L72-L236)
    
*   [field\_text.go50-177](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go#L50-L177)
    
*   [field\_select.go60-161](https://github.com/charmbracelet/huh/blob/151ba059/field_select.go#L60-L161)
    
*   [field\_multiselect.go55-206](https://github.com/charmbracelet/huh/blob/151ba059/field_multiselect.go#L55-L206)
    

Field Types in Detail
---------------------

### Input Field

The `Input` field provides a single-line text input for capturing short text values.

Key features:

*   Single-line text input
*   Password input with `EchoMode` options: `EchoModeNormal`, `EchoModePassword`, `EchoModeNone`
*   Character limits via `CharLimit(int)`
*   Placeholder text with `Placeholder(string)`
*   Inline mode with `Inline(bool)` (title and input on same line)
*   Autocomplete suggestions with `Suggestions([]string)` or `SuggestionsFunc()`

Example:

Sources:

*   [README.md157-169](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L157-L169)
    
*   [field\_input.go](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go)
    
*   [examples/burger/main.go126-136](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L126-L136)
    

### Text Field

The `Text` field provides multi-line text input for longer text content.

Key features:

*   Multi-line text input
*   Line limits via `Lines(int)`
*   Character limits via `CharLimit(int)`
*   Line numbers via `ShowLineNumbers(bool)`
*   External editor integration via `ExternalEditor(bool)` and `Editor(...string)`
*   Custom file extension for external editor via `EditorExtension(string)`

Example:

Sources:

*   [README.md172-182](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L172-L182)
    
*   [field\_text.go](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go)
    
*   [examples/burger/main.go138-144](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L138-L144)
    

### Select Field

The `Select` field allows users to choose a single option from a list of choices.

Key features:

*   Single option selection from a list
*   Generic type support via `Select<T>`
*   Dynamic option loading via `OptionsFunc(func() []Option[T], bindings)`
*   Filtering with the `/` key
*   Customizable height via `Height(int)`
*   Inline mode (compact display) via `Inline(bool)`

Example:

Sources:

*   [README.md185-200](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L185-L200)
    
*   [field\_select.go](https://github.com/charmbracelet/huh/blob/151ba059/field_select.go)
    
*   [examples/burger/main.go70-80](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L70-L80)
    

### MultiSelect Field

The `MultiSelect` field enables users to choose multiple options from a list of choices.

Key features:

*   Multiple option selection from a list
*   Generic type support via `MultiSelect<T>`
*   Selection limits via `Limit(int)`
*   Dynamic option loading via `OptionsFunc(func() []Option[T], bindings)`
*   Filtering via `Filterable(bool)` and the `/` key
*   Customizable height via `Height(int)`

Example:

Sources:

*   [README.md203-222](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L203-L222)
    
*   [field\_multiselect.go](https://github.com/charmbracelet/huh/blob/151ba059/field_multiselect.go)
    
*   [examples/burger/main.go82-102](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L82-L102)
    

### Confirm Field

The `Confirm` field presents a simple yes/no choice for binary decisions.

Key features:

*   Simple Yes/No selection
*   Custom labels for affirmative and negative options via `Affirmative(string)` and `Negative(string)`

Example:

Sources:

*   [README.md225-236](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L225-L236)
    
*   [examples/burger/main.go146-150](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L146-L150)
    

### Note Field

The `Note` field displays non-interactive text for informational purposes.

Key features:

*   Display of non-interactive text
*   Optional "Next" button via `Next(bool)`
*   Custom label for the Next button via `NextLabel(string)`
*   Markdown-like formatting in the description

Example:

Sources:

*   [examples/burger/main.go60-65](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L60-L65)
    

### FilePicker Field

The `FilePicker` field provides file system navigation and selection capabilities.

Key features:

*   File system navigation and selection
*   Starting directory specification via `Path(string)`
*   File extension filtering via `Extension(...string)`
*   Directory selection mode via `Directory(bool)`

Example:

Sources:

*   [huh\_test.go728-744](https://github.com/charmbracelet/huh/blob/151ba059/huh_test.go#L728-L744)
    

Comparison of Field Types
-------------------------

Here's a comparison of key capabilities across field types:

| Feature | Input | Text | Select | MultiSelect | Confirm | Note | FilePicker |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Value Type | string | string | T (generic) | \[\]T (generic) | bool | none | string |
| Multiple Lines | No  | Yes | No  | No  | No  | Yes | No  |
| Selection | No  | No  | Single | Multiple | Binary | No  | File/Dir |
| Filtering | No\* | No  | Yes | Yes | No  | No  | Yes (extension) |
| Character Limit | Yes | Yes | No  | No  | No  | No  | No  |
| Dynamic Options | No\*\* | No  | Yes | Yes | No  | No  | No  |

\* Input has suggestion support rather than filtering  
\*\* Input has dynamic placeholder and suggestions

Sources:

*   [field\_input.go](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go)
    
*   [field\_text.go](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go)
    
*   [field\_select.go](https://github.com/charmbracelet/huh/blob/151ba059/field_select.go)
    
*   [field\_multiselect.go](https://github.com/charmbracelet/huh/blob/151ba059/field_multiselect.go)
    
*   [README.md136-236](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L136-L236)
    

Accessibility Support
---------------------

All field types in `huh` include built-in accessibility support for screen readers. When a field is rendered with accessibility mode enabled (via `WithAccessible(true)`), the field will use a simpler, more screen reader-friendly rendering approach.

The accessible mode implementations are defined in each field type's `runAccessible` method, which typically:

1.  Displays a simple text prompt
2.  Reads input directly from stdin
3.  Handles validation inline
4.  Provides clear error messages

The accessibility functionality is implemented through helper functions in the `internal/accessibility` package:

*   `PromptString` - Prompts for text input
*   `PromptInt` - Prompts for numeric input within a range
*   `PromptBool` - Prompts for yes/no input
*   `PromptPassword` - Prompts for password input with masked display

Sources:

*   [internal/accessibility/accessibility.go](https://github.com/charmbracelet/huh/blob/151ba059/internal/accessibility/accessibility.go)
    
*   [accessibility/accessibility.go](https://github.com/charmbracelet/huh/blob/151ba059/accessibility/accessibility.go)
    
*   [field\_input.go417-464](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go#L417-L464)
    
*   [field\_text.go401-440](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go#L401-L440)
    
*   [field\_select.go702-734](https://github.com/charmbracelet/huh/blob/151ba059/field_select.go#L702-L734)
    
*   [field\_multiselect.go701-747](https://github.com/charmbracelet/huh/blob/151ba059/field_multiselect.go#L701-L747)
    

Option Type for Selection Fields
--------------------------------

Both `Select` and `MultiSelect` fields use the `Option[T]` type to define selectable options:

You can create options using the `NewOption` function:

For convenience, you can create multiple options at once using `NewOptions`:

Sources:

*   [field\_select.go20-24](https://github.com/charmbracelet/huh/blob/151ba059/field_select.go#L20-L24)
    
*   [field\_multiselect.go20-24](https://github.com/charmbracelet/huh/blob/151ba059/field_multiselect.go#L20-L24)
    
*   [README.md191-200](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L191-L200)
    

Dynamic Field Content
---------------------

Many field properties support dynamic updates based on other field values, using the `*Func` variant of methods:

*   `TitleFunc(func() string, bindings)`
*   `DescriptionFunc(func() string, bindings)`
*   `PlaceholderFunc(func() string, bindings)`
*   `OptionsFunc(func() []Option[T], bindings)`
*   `SuggestionsFunc(func() []string, bindings)`

The `bindings` parameter indicates which values, when changed, should trigger a re-evaluation of the function. This enables creating dynamic forms that adapt based on user input.

Example:

Sources:

*   [README.md283-351](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L283-L351)
    
*   [field\_input.go108-132](https://github.com/charmbracelet/huh/blob/151ba059/field_input.go#L108-L132)
    
*   [field\_text.go107-139](https://github.com/charmbracelet/huh/blob/151ba059/field_text.go#L107-L139)
    
*   [field\_select.go128-161](https://github.com/charmbracelet/huh/blob/151ba059/field_select.go#L128-L161)
    
*   [field\_multiselect.go106-171](https://github.com/charmbracelet/huh/blob/151ba059/field_multiselect.go#L106-L171)
    

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Field Types](https://deepwiki.com/charmbracelet/huh/3-field-types#field-types)
    
*   [Field Type Overview](https://deepwiki.com/charmbracelet/huh/3-field-types#field-type-overview)
    
*   [Common Structure and Features](https://deepwiki.com/charmbracelet/huh/3-field-types#common-structure-and-features)
    
*   [Constructor Pattern](https://deepwiki.com/charmbracelet/huh/3-field-types#constructor-pattern)
    
*   [Builder Pattern](https://deepwiki.com/charmbracelet/huh/3-field-types#builder-pattern)
    
*   [Common Methods](https://deepwiki.com/charmbracelet/huh/3-field-types#common-methods)
    
*   [Common Configuration Methods](https://deepwiki.com/charmbracelet/huh/3-field-types#common-configuration-methods)
    
*   [Field Types in Detail](https://deepwiki.com/charmbracelet/huh/3-field-types#field-types-in-detail)
    
*   [Input Field](https://deepwiki.com/charmbracelet/huh/3-field-types#input-field)
    
*   [Text Field](https://deepwiki.com/charmbracelet/huh/3-field-types#text-field)
    
*   [Select Field](https://deepwiki.com/charmbracelet/huh/3-field-types#select-field)
    
*   [MultiSelect Field](https://deepwiki.com/charmbracelet/huh/3-field-types#multiselect-field)
    
*   [Confirm Field](https://deepwiki.com/charmbracelet/huh/3-field-types#confirm-field)
    
*   [Note Field](https://deepwiki.com/charmbracelet/huh/3-field-types#note-field)
    
*   [FilePicker Field](https://deepwiki.com/charmbracelet/huh/3-field-types#filepicker-field)
    
*   [Comparison of Field Types](https://deepwiki.com/charmbracelet/huh/3-field-types#comparison-of-field-types)
    
*   [Accessibility Support](https://deepwiki.com/charmbracelet/huh/3-field-types#accessibility-support)
    
*   [Option Type for Selection Fields](https://deepwiki.com/charmbracelet/huh/3-field-types#option-type-for-selection-fields)
    
*   [Dynamic Field Content](https://deepwiki.com/charmbracelet/huh/3-field-types#dynamic-field-content)
