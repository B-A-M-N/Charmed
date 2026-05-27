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

Spinner Component
=================

Relevant source files

*   [examples/burger/burger.gif](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/burger.gif)
    
*   [examples/burger/demo.tape](https://github.com/charmbracelet/huh/blob/151ba059/examples/burger/demo.tape)
    
*   [spinner/examples/context-and-action-and-error/main.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/context-and-action-and-error/main.go)
    
*   [spinner/examples/context-and-action/main.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/context-and-action/main.go)
    
*   [spinner/examples/loading/main.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/loading/main.go)
    
*   [spinner/examples/loading/spinner.gif](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/loading/spinner.gif)
    
*   [spinner/go.mod](https://github.com/charmbracelet/huh/blob/151ba059/spinner/go.mod)
    
*   [spinner/go.sum](https://github.com/charmbracelet/huh/blob/151ba059/spinner/go.sum)
    
*   [spinner/spinner.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go)
    
*   [spinner/spinner\_test.go](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner_test.go)
    

The Spinner Component provides a customizable loading indicator for terminal applications built with the Charmbracelet/Huh library. It's designed to show activity while operations are being performed, with support for various spinner styles, titles, and accessibility features.

Sources: [spinner/spinner.go16-32](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L16-L32)

Overview
--------

The Spinner Component is implemented as a standalone module within the Huh library. It can be used both within forms and as an independent element in terminal applications. Unlike other form fields, spinners don't collect user input but instead provide visual feedback during background operations.

Sources: [spinner/spinner.go16-32](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L16-L32)
 [spinner/spinner.go126-162](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L126-L162)

Basic Usage
-----------

Creating and running a spinner is straightforward:

Here's how to use the spinner in its simplest form:

1.  Create a new spinner using `spinner.New()`
2.  Optionally configure it with methods like `Title()` and `Action()`
3.  Call the `Run()` method to start the spinner
4.  The spinner runs until the action completes or context is canceled

Sources: [spinner/spinner.go112-124](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L112-L124)
 [spinner/spinner.go164-192](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L164-L192)
 [spinner/examples/loading/main.go10-19](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/loading/main.go#L10-L19)

Spinner Types
-------------

The Spinner Component supports multiple visual styles for the animation:

| Type Name | Description |
| --- | --- |
| Line | Line segments animation |
| Dots | Braille dots pattern |
| MiniDot | Smaller dots animation |
| Jump | Jumping animation |
| Points | Moving points |
| Pulse | Pulsing animation |
| Globe | Rotating globe animation |
| Moon | Moon phases animation |
| Monkey | Playful monkey animation |
| Meter | Progress meter animation |
| Hamburger | Hamburger-style animation |
| Ellipsis | Classic ellipsis animation |

The default spinner type is `Dots`. To change the spinner type:

    spinner.New().Type(spinner.Line)
    

Sources: [spinner/spinner.go34-48](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L34-L48)
 [spinner/spinner.go51-55](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L51-L55)

Configuration Options
---------------------

The Spinner Component can be customized through a fluent API:

| Method | Description |
| --- | --- |
| `Type(t Type)` | Sets the spinner animation style |
| `Title(title string)` | Sets the text displayed next to the spinner |
| `Output(w io.Writer)` | Sets the output writer (defaults to stdout/stderr) |
| `Action(action func())` | Sets a function to execute while the spinner is running |
| `ActionWithErr(action func(context.Context) error)` | Sets a function that can return an error and takes a context |
| `Context(ctx context.Context)` | Sets a context that can cancel the spinner |
| `Style(style lipgloss.Style)` | Sets the style for the spinner animation |
| `TitleStyle(style lipgloss.Style)` | Sets the style for the title text |
| `Accessible(accessible bool)` | Enables static mode for accessibility |

Sources: [spinner/spinner.go51-110](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L51-L110)

Working with Contexts and Actions
---------------------------------

The spinner can be configured to run actions and respect context cancellation:

### Action Execution

The spinner can execute an action while displaying the animation:

    spinner.New().
        Title("Processing...").
        Action(func() {
            // Perform a task here
        }).
        Run()
    

The spinner will stop automatically when the action completes.

### Context Integration

You can provide a context to control the spinner's lifecycle:

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    spinner.New().
        Title("Loading...").
        Context(ctx).
        Run()
    

The spinner will stop if the context is canceled or times out.

### Error Handling

The `ActionWithErr` method allows actions to return errors:

    err := spinner.New().
        Title("Fetching data...").
        ActionWithErr(func(ctx context.Context) error {
            // Perform a task that might return an error
            return nil
        }).
        Run()
    
    if err != nil {
        // Handle error
    }
    

Sources: [spinner/spinner.go70-86](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L70-L86)
 [spinner/spinner.go88-92](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L88-L92)
 [spinner/examples/context-and-action/main.go13-28](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/context-and-action/main.go#L13-L28)
 [spinner/examples/context-and-action-and-error/main.go12-28](https://github.com/charmbracelet/huh/blob/151ba059/spinner/examples/context-and-action-and-error/main.go#L12-L28)

Accessibility Support
---------------------

The Spinner Component includes an accessibility mode that renders the spinner as static text instead of an animation, making it compatible with screen readers and other assistive technologies.

Enable accessibility mode with:

    spinner.New().Accessible(true)
    

When accessibility is enabled:

1.  The spinner animation is replaced with static dots ("...")
2.  Output defaults to stdout instead of stderr
3.  Animation frames aren't rendered, improving compatibility with screen readers

Sources: [spinner/spinner.go106-110](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L106-L110)
 [spinner/spinner.go194-222](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L194-L222)

Implementation Details
----------------------

The Spinner Component implements the Bubble Tea Model interface, allowing it to be used within Bubble Tea applications:

### Internal Architecture

1.  The Spinner uses the Bubble Tea framework to manage its lifecycle and rendering
2.  It wraps the `bubbles/spinner.Model` to provide the actual animation
3.  When an action is provided, it's executed in a separate goroutine
4.  The spinner listens for completion signals or context cancellation
5.  Styling is handled through Lip Gloss, providing consistent visual appearance

### Termination Conditions

The spinner terminates under the following conditions:

*   The provided action completes
*   The context is canceled
*   The user presses Ctrl+C

Sources: [spinner/spinner.go126-162](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L126-L162)
 [spinner/spinner.go164-222](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L164-L222)
 [spinner/spinner.go224-226](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L224-L226)

Relationship to Other Components
--------------------------------

The Spinner Component can be used independently of other Huh components, but it complements them in complex form applications where some operations may require loading time.

Unlike form fields like Input, Select, or Confirm, the Spinner doesn't collect user input. Instead, it provides visual feedback during potentially slow operations like network requests, file operations, or complex calculations.

For advanced form validation or dynamic content generation based on API calls, consider using spinners between form field interactions. For more information on integrating spinners with forms, see [Form](https://deepwiki.com/charmbracelet/huh/2.1-form)
 and [Dynamic Forms](https://deepwiki.com/charmbracelet/huh/9.1-dynamic-forms)
.

Sources: [spinner/spinner.go16-32](https://github.com/charmbracelet/huh/blob/151ba059/spinner/spinner.go#L16-L32)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Spinner Component](https://deepwiki.com/charmbracelet/huh/6-spinner-component#spinner-component)
    
*   [Overview](https://deepwiki.com/charmbracelet/huh/6-spinner-component#overview)
    
*   [Basic Usage](https://deepwiki.com/charmbracelet/huh/6-spinner-component#basic-usage)
    
*   [Spinner Types](https://deepwiki.com/charmbracelet/huh/6-spinner-component#spinner-types)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/huh/6-spinner-component#configuration-options)
    
*   [Working with Contexts and Actions](https://deepwiki.com/charmbracelet/huh/6-spinner-component#working-with-contexts-and-actions)
    
*   [Action Execution](https://deepwiki.com/charmbracelet/huh/6-spinner-component#action-execution)
    
*   [Context Integration](https://deepwiki.com/charmbracelet/huh/6-spinner-component#context-integration)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/huh/6-spinner-component#error-handling)
    
*   [Accessibility Support](https://deepwiki.com/charmbracelet/huh/6-spinner-component#accessibility-support)
    
*   [Implementation Details](https://deepwiki.com/charmbracelet/huh/6-spinner-component#implementation-details)
    
*   [Internal Architecture](https://deepwiki.com/charmbracelet/huh/6-spinner-component#internal-architecture)
    
*   [Termination Conditions](https://deepwiki.com/charmbracelet/huh/6-spinner-component#termination-conditions)
    
*   [Relationship to Other Components](https://deepwiki.com/charmbracelet/huh/6-spinner-component#relationship-to-other-components)
