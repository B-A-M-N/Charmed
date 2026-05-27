Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/pop](https://github.com/charmbracelet/pop "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 20 April 2025 ([cdaa72](https://github.com/charmbracelet/pop/commits/cdaa7255)
)

*   [Overview](https://deepwiki.com/charmbracelet/pop/1-overview)
    
*   [Installation and Configuration](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration)
    
*   [Architecture](https://deepwiki.com/charmbracelet/pop/3-architecture)
    
*   [Email Delivery System](https://deepwiki.com/charmbracelet/pop/3.1-email-delivery-system)
    
*   [User Interface Components](https://deepwiki.com/charmbracelet/pop/3.2-user-interface-components)
    
*   [Command Line Interface](https://deepwiki.com/charmbracelet/pop/4-command-line-interface)
    
*   [Terminal User Interface](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface)
    
*   [Email Delivery Methods](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods)
    
*   [Resend API Configuration](https://deepwiki.com/charmbracelet/pop/6.1-resend-api-configuration)
    
*   [SMTP Configuration](https://deepwiki.com/charmbracelet/pop/6.2-smtp-configuration)
    
*   [Examples and Integrations](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations)
    
*   [Development](https://deepwiki.com/charmbracelet/pop/8-development)
    
*   [Building and Testing](https://deepwiki.com/charmbracelet/pop/8.1-building-and-testing)
    
*   [Contributing Guidelines](https://deepwiki.com/charmbracelet/pop/8.2-contributing-guidelines)
    

Menu

Terminal User Interface
=======================

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/pop/blob/cdaa7255/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1)
    
*   [keymap.go](https://github.com/charmbracelet/pop/blob/cdaa7255/keymap.go)
    
*   [model.go](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go)
    

This document explains how to use Pop's interactive Terminal User Interface (TUI) for composing and sending emails directly from your terminal. For information about the Command Line Interface, see [Command Line Interface](https://deepwiki.com/charmbracelet/pop/4-command-line-interface)
.

Overview
--------

Pop's TUI provides a simple form-based interface that allows you to compose emails with rich formatting, add attachments, and send them via either Resend API or SMTP. The TUI is launched by simply running the `pop` command without arguments:

Sources: [README.md15-21](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L15-L21)

UI Components and Layout
------------------------

The TUI presents a structured email composition form with the following fields:

*   **From**: Sender email address
*   **To**: Recipient email address(es)
*   **Cc/Bcc**: Carbon copy and blind carbon copy recipients (optional)
*   **Subject**: Email subject line
*   **Body**: Email content (supports Markdown formatting)
*   **Attachments**: List of attached files
*   **Send button**: To dispatch the email when ready

Sources: [model.go52-85](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L52-L85)
 [model.go425-473](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L425-L473)

UI States
---------

The TUI operates in different states that determine which component is active and receives input. These states are defined in the `State` type:

Sources: [model.go20-33](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L20-L33)
 [model.go241-319](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L241-L319)

Navigation
----------

The TUI uses standard keyboard navigation to move between fields:

| Key | Action |
| --- | --- |
| Tab | Move to the next field |
| Shift+Tab | Move to the previous field |
| Enter | When on the "Send" button, send the email |
|     | When in attachments, open the file picker |
| Esc | Return from file picker to the main form |
| Ctrl+C | Quit the application |

The active field is highlighted with a different color to indicate focus.

Sources: [keymap.go5-14](https://github.com/charmbracelet/pop/blob/cdaa7255/keymap.go#L5-L14)
 [keymap.go16-51](https://github.com/charmbracelet/pop/blob/cdaa7255/keymap.go#L16-L51)
 [model.go241-319](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L241-L319)

Key Mapping
-----------

Pop uses a defined keymap to handle user interactions. These key bindings are defined in the `KeyMap` struct and are context-sensitive, meaning they change based on the current state of the UI.

The keymap is dynamically updated based on the current state:

*   The `Attach` key is only enabled when in the attachments editing state
*   The `Send` key is only enabled when all required fields are filled and the Send button is selected
*   The `Unattach` key is only enabled when in the attachments editing state and there are attachments
*   The `Back` key is only enabled when in the file picker state

Sources: [keymap.go1-91](https://github.com/charmbracelet/pop/blob/cdaa7255/keymap.go#L1-L91)
 [model.go71-75](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L71-L75)

Email Composition Flow
----------------------

Composing and sending an email in the TUI follows this process:

1.  Fill in the "From" field with your email address
2.  Fill in the "To" field with recipient addresses (comma-separated)
3.  Optionally fill in "Cc" and "Bcc" fields (if visible)
4.  Enter a subject
5.  Compose the email body (supports Markdown formatting)
6.  Optionally add file attachments
7.  Navigate to the Send button and press Enter

If any required field is missing (From, To, Subject, Body), the Send button will be disabled.

Sources: [model.go87-212](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L87-L212)
 [model.go88-90](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L88-L90)

Handling Attachments
--------------------

The attachments section allows you to add files to your email:

1.  Tab to the "Attachments" section
2.  Press Enter to open the file picker
3.  Navigate through directories using arrow keys
4.  Press Enter to select a file or open a directory
5.  After selecting a file, you return to the attachments list
6.  To remove an attachment, highlight it and press `x`

The file picker displays your current directory and allows navigation through your file system.

Sources: [model.go339-348](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L339-L348)
 [model.go432-434](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L432-L434)

Error Handling
--------------

If an error occurs during email sending:

1.  The error is displayed at the bottom of the TUI
2.  The cursor returns to the "From" field
3.  The error message automatically clears after 10 seconds
4.  The email content is preserved so you can fix any issues and try again

Sources: [model.go221-240](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L221-L240)
 [model.go467-470](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L467-L470)

Advanced Features
-----------------

### Optional Cc and Bcc Fields

The Cc and Bcc fields are hidden by default but are automatically shown when:

*   Values are provided for them in a previous session
*   Environment variables set them
*   Command-line flags set them

This keeps the interface clean for simple emails while supporting more complex use cases when needed.

Sources: [model.go73-74](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L73-L74)
 [model.go196](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L196-L196)

### Default Values

The TUI can be pre-populated with default values from:

*   Environment variables (like `POP_FROM`)
*   Command-line flags (partial arguments)
*   Previous email attempts that failed

This allows for faster email composition and consistent defaults.

Sources: [model.go87-212](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L87-L212)

### Email Preview

The TUI provides a live preview of your email as you compose it. The Markdown in the body will be converted to HTML when the email is sent.

Sources: [model.go52-85](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L52-L85)

Technical Implementation
------------------------

The TUI is built using the Bubble Tea library, a Go framework for building terminal user interfaces. The main components used include:

*   `textinput.Model`: For single-line text inputs (From, To, Subject, etc.)
*   `textarea.Model`: For multi-line body input
*   `list.Model`: For displaying attachments
*   `filepicker.Model`: For selecting files
*   `spinner.Model`: For displaying a loading indicator during email sending
*   `help.Model`: For displaying key bindings

The TUI follows the Model-View-Update (MVU) architecture, where:

*   The `Model` struct holds the application state
*   The `Update` method handles messages and updates the model
*   The `View` method renders the current state as a string

Sources: [model.go1-19](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L1-L19)
 [model.go214-362](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L214-L362)
 [model.go425-473](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L425-L473)

TUI to Email Delivery Integration
---------------------------------

When an email is ready to be sent, the TUI calls the `sendEmailCmd()` method, which:

1.  Constructs an email request from the form fields
2.  Converts the Markdown body to HTML
3.  Sends the email using the configured delivery method (Resend API or SMTP)
4.  Returns a success or failure message to the TUI

The delivery method is determined based on configuration and credentials available in the environment.

Sources: [model.go304-307](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L304-L307)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Terminal User Interface](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#terminal-user-interface)
    
*   [Overview](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#overview)
    
*   [UI Components and Layout](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#ui-components-and-layout)
    
*   [UI States](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#ui-states)
    
*   [Navigation](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#navigation)
    
*   [Key Mapping](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#key-mapping)
    
*   [Email Composition Flow](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#email-composition-flow)
    
*   [Handling Attachments](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#handling-attachments)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#error-handling)
    
*   [Advanced Features](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#advanced-features)
    
*   [Optional Cc and Bcc Fields](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#optional-cc-and-bcc-fields)
    
*   [Default Values](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#default-values)
    
*   [Email Preview](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#email-preview)
    
*   [Technical Implementation](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#technical-implementation)
    
*   [TUI to Email Delivery Integration](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface#tui-to-email-delivery-integration)
