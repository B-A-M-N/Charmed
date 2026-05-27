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

Advanced Usage
==============

Relevant source files

*   [.github/workflows/lint-sync.yml](https://github.com/charmbracelet/huh/blob/151ba059/.github/workflows/lint-sync.yml)
    
*   [.golangci.yml](https://github.com/charmbracelet/huh/blob/151ba059/.golangci.yml)
    
*   [README.md](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1)
    
*   [clamp.go](https://github.com/charmbracelet/huh/blob/151ba059/clamp.go)
    
*   [examples/burger/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go)
    
*   [examples/dynamic/dynamic-bubbletea/main.go](https://github.com/charmbracelet/huh/blob/151ba059/examples/dynamic/dynamic-bubbletea/main.go)
    
*   [validate.go](https://github.com/charmbracelet/huh/blob/151ba059/validate.go)
    

This page covers advanced features and techniques for the `charmbracelet/huh` library. Here we'll explore creating dynamic forms that respond to user input, implementing robust validation, and integrating with SSH applications. If you're looking for information about core components, refer to [Core Components](https://deepwiki.com/charmbracelet/huh/2-core-components)
 or for field types, see [Field Types](https://deepwiki.com/charmbracelet/huh/3-field-types)
.

Dynamic Forms
-------------

Dynamic forms allow you to create interactive experiences where form content changes based on user input. This is particularly useful for creating conditional flows, wizard-like interfaces, or forms where certain options depend on previous selections.

### How Dynamic Forms Work

In `huh`, you can make a form dynamic by using callback functions for properties that should be computed at runtime. These callback functions are reevaluated whenever their bound values change.

Sources: [README.md126-355](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L126-L355)

### Dynamic Property Methods

The following field properties can be made dynamic through their `Func` variants:

| Standard Method | Dynamic Equivalent | Purpose |
| --- | --- | --- |
| `Title(string)` | `TitleFunc(func() string, binding any)` | Dynamic field title |
| `Description(string)` | `DescriptionFunc(func() string, binding any)` | Dynamic field description |
| `Options([]Option)` | `OptionsFunc(func() []Option, binding any)` | Dynamic available options |
| `Validate(func() error)` | N/A (Already dynamic) | Dynamic validation |

Sources: [README.md310-342](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L310-L342)

### Dynamic Options Example

A common use case is to change available options based on a user's previous selection, such as showing states/provinces based on the selected country:

The implementation looks like this:

    var country string
    var state string
    
    // Define the country select field normally
    huh.NewSelect<FileRef file-url="https://github.com/charmbracelet/huh/blob/151ba059/string" undefined  file-path="string">Hii</FileRef>
        Options(huh.NewOptions("United States", "Canada", "Mexico")...).
        Value(&country).
        Title("Country")
    
    // Define the state field with dynamic properties
    huh.NewSelect<FileRef file-url="https://github.com/charmbracelet/huh/blob/151ba059/string" undefined  file-path="string">Hii</FileRef>
        Value(&state).
        Height(8).
        TitleFunc(func() string {
            switch country {
            case "United States":
                return "State"
            case "Canada":
                return "Province"
            default:
                return "Territory"
            }
        }, &country).  // country variable triggers re-evaluation
        OptionsFunc(func() []huh.Option[string] {
            opts := fetchStatesForCountry(country)
            return huh.NewOptions(opts...)
        }, &country)   // country variable triggers re-evaluation
    

Sources: [README.md310-342](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L310-L342)
 [examples/dynamic/dynamic-bubbletea/main.go78-111](https://github.com/charmbracelet/huh/blob/151ba059/examples/dynamic/dynamic-bubbletea/main.go#L78-L111)

### Advanced Dynamic Forms in Bubble Tea

For more complex scenarios, you can integrate dynamic forms with Bubble Tea applications:

In a Bubble Tea application, dynamic forms are integrated by:

1.  Including the form in your application model
2.  Forwarding messages to the form in your Update method
3.  Checking the form's state to determine when it's completed
4.  Accessing form data through the form's `Get` methods

Sources: [README.md403-462](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L403-L462)
 [examples/dynamic/dynamic-bubbletea/main.go1-298](https://github.com/charmbracelet/huh/blob/151ba059/examples/dynamic/dynamic-bubbletea/main.go#L1-L298)

Validation
----------

Validation ensures that user input meets specific requirements. The `huh` library provides both built-in and custom validation capabilities.

### Validation Flow

Sources: [validate.go1-61](https://github.com/charmbracelet/huh/blob/151ba059/validate.go#L1-L61)
 [README.md91-98](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L91-L98)

### Built-in Validation Functions

The library provides utility functions for common validation needs:

| Function | Purpose |
| --- | --- |
| `ValidateNotEmpty()` | Ensures input is not empty |
| `ValidateMinLength(min int)` | Ensures input has minimum length |
| `ValidateMaxLength(max int)` | Ensures input doesn't exceed maximum length |
| `ValidateLength(min, max int)` | Ensures input length is within range |
| `ValidateOneOf(options ...string)` | Ensures input is one of specified options |

These functions return a validation function with signature `func(string) error` that can be passed to a field's `Validate` method.

Sources: [validate.go8-61](https://github.com/charmbracelet/huh/blob/151ba059/validate.go#L8-L61)

### Type-Specific Validation

Different field types require different validation function signatures:

| Field Type | Validation Function Signature |
| --- | --- |
| Input/Text | `func(string) error` |
| Select<T> | `func(T) error` |
| MultiSelect<T> | `func([]T) error` |
| Confirm | `func(bool) error` |

Examples:

Validating an input field:

    huh.NewInput().
        Value(&name).
        Title("What's your name?").
        Validate(func(s string) error {
            if s == "Frank" {
                return errors.New("no franks, sorry")
            }
            return nil
        })
    

Validating a multi-select field:

    huh.NewMultiSelect<FileRef file-url="https://github.com/charmbracelet/huh/blob/151ba059/string" undefined  file-path="string">Hii</FileRef>
        Title("Toppings").
        Options(...).
        Validate(func(t []string) error {
            if len(t) <= 0 {
                return fmt.Errorf("at least one topping is required")
            }
            return nil
        }).
        Value(&toppings)
    

Sources: [README.md91-98](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L91-L98)
 [examples/burger/main.go74-100](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L74-L100)

### Combining Validation Functions

For complex validation requirements, you can combine multiple validation functions:

    func combineValidators(validators ...func(string) error) func(string) error {
        return func(s string) error {
            for _, validator := range validators {
                if err := validator(s); err != nil {
                    return err
                }
            }
            return nil
        }
    }
    
    // Usage
    huh.NewInput().
        Value(&email).
        Validate(combineValidators(
            ValidateNotEmpty(),
            ValidateMinLength(3),
            ValidateMaxLength(100),
            customEmailValidator,
        ))
    

Sources: [validate.go38-46](https://github.com/charmbracelet/huh/blob/151ba059/validate.go#L38-L46)

SSH Integration
---------------

The `huh` library works seamlessly in SSH applications, making it an excellent choice for creating interactive terminal applications accessible remotely.

### SSH Integration Architecture

Sources: [README.md403-462](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L403-L462)

### Considerations for SSH Applications

When implementing `huh` forms in SSH applications:

1.  **Terminal Capabilities**: SSH clients may have different terminal capabilities. Test with various SSH clients.
    
2.  **Window Size Handling**: Handle window size changes to ensure your forms adapt properly:
    
        case tea.WindowSizeMsg:
            // Update form width/height based on terminal size
            form.WithWidth(min(msg.Width, maxWidth))
        
    
3.  **Accessible Mode**: Consider enabling accessible mode for users with screen readers:
    
        // Read from environment or command flag
        accessibleMode := os.Getenv("ACCESSIBLE") != ""
        form.WithAccessible(accessibleMode)
        
    
4.  **Network Latency**: Be mindful of network latency. Prefer designs that minimize round-trips and provide good user feedback.
    

Sources: [README.md238-255](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L238-L255)
 [examples/dynamic/dynamic-bubbletea/main.go131-140](https://github.com/charmbracelet/huh/blob/151ba059/examples/dynamic/dynamic-bubbletea/main.go#L131-L140)

### Supporting Multiple Input Methods

For SSH applications that might be used by various clients, consider:

1.  Providing keyboard shortcuts that work across different terminals
2.  Using a common theme that displays well in different terminal color schemes
3.  Keeping forms simple and intuitive for users accessing remotely
4.  Using spinners or other indicators for operations that take time

Sources: [README.md355-393](https://github.com/charmbracelet/huh/blob/151ba059/README.md?plain=1#L355-L393)
 [examples/burger/main.go160-164](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/main.go#L160-L164)

Performance Optimization
------------------------

When working with complex forms or in resource-constrained environments, consider these optimization techniques:

1.  **Caching Computed Values**: For expensive `OptionsFunc` or other dynamic computations, implement caching to avoid unnecessary recalculations.
    
2.  **Throttling Updates**: For rapid input changes, consider throttling updates to prevent excessive re-renders.
    
3.  **Pagination for Large Data Sets**: When displaying large option sets, implement pagination to improve performance.
    
4.  **Progressive Loading**: Load form sections progressively rather than all at once for complex multi-page forms.
    
5.  **Limit Validator Complexity**: Keep validation functions lightweight, as they run on every input change.
    

Sources: [clamp.go1-9](https://github.com/charmbracelet/huh/blob/151ba059/clamp.go#L1-L9)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Advanced Usage](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#advanced-usage)
    
*   [Dynamic Forms](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#dynamic-forms)
    
*   [How Dynamic Forms Work](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#how-dynamic-forms-work)
    
*   [Dynamic Property Methods](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#dynamic-property-methods)
    
*   [Dynamic Options Example](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#dynamic-options-example)
    
*   [Advanced Dynamic Forms in Bubble Tea](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#advanced-dynamic-forms-in-bubble-tea)
    
*   [Validation](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#validation)
    
*   [Validation Flow](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#validation-flow)
    
*   [Built-in Validation Functions](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#built-in-validation-functions)
    
*   [Type-Specific Validation](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#type-specific-validation)
    
*   [Combining Validation Functions](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#combining-validation-functions)
    
*   [SSH Integration](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#ssh-integration)
    
*   [SSH Integration Architecture](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#ssh-integration-architecture)
    
*   [Considerations for SSH Applications](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#considerations-for-ssh-applications)
    
*   [Supporting Multiple Input Methods](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#supporting-multiple-input-methods)
    
*   [Performance Optimization](https://deepwiki.com/charmbracelet/huh/9-advanced-usage#performance-optimization)
