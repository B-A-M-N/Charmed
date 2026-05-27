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

Architecture
============

Relevant source files

*   [email.go](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go)
    
*   [main.go](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go)
    
*   [model.go](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go)
    

This document provides a comprehensive overview of the Pop email client's architecture, including its components, their interactions, and the system's workflow. The architecture section focuses on the technical design and implementation of Pop, explaining how different parts of the system collaborate to provide email capabilities from the terminal.

For specific information about using the command-line interface, see [Command Line Interface](https://deepwiki.com/charmbracelet/pop/4-command-line-interface)
. For details on the interactive terminal UI, see [Terminal User Interface](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface)
.

System Architecture Overview
----------------------------

Pop is designed with a layered architecture that separates user interfaces, core logic, and email delivery mechanisms. This separation allows for flexibility in how emails are composed and sent.

Sources: [main.go72-158](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L72-L158)
 [model.go1-86](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L1-L86)
 [email.go32-60](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L32-L60)

Component Interactions and Flow
-------------------------------

The following diagram illustrates how the different components interact during the email sending process:

Sources: [main.go76-156](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L76-L156)
 [email.go32-60](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L32-L60)
 [model.go229-362](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L229-L362)

Core Components
---------------

### User Interface Layer

Pop provides two different interfaces for users:

1.  **Command Line Interface (CLI)**: Used when all required arguments are provided directly through command-line flags
2.  **Terminal User Interface (TUI)**: An interactive form-based interface activated when arguments are missing

#### CLI Implementation

The CLI is implemented using the Cobra library, which parses command-line arguments and flags to configure the email sending process.

    ┌─────────────────────┐
    │ CLI Component       │
    ├─────────────────────┤
    │ - Parse arguments   │
    │ - Set email config  │
    │ - Handle direct     │
    │   email sending     │
    └─────────────────────┘
    

Sources: [main.go72-158](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L72-L158)
 [main.go196-228](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L196-L228)

#### TUI Implementation

The TUI is implemented using the Bubble Tea library (and related Bubble components), providing an interactive form for composing emails.

Sources: [model.go20-86](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L20-L86)
 [model.go88-212](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L88-L212)

### Core Logic Layer

The core logic layer consists of:

1.  **Model**: Maintains the application state and UI components
2.  **EmailProcessor**: Handles email composition and sending
3.  **KeyMap**: Defines key bindings for user interaction

#### Model Structure

The `Model` struct is the central component that coordinates the UI and email processing:

    ┌─────────────────────────────────┐
    │ Model                           │
    ├─────────────────────────────────┤
    │ - state: State                  │
    │ - DeliveryMethod: DeliveryMethod│
    │ - From: textinput.Model         │
    │ - To: textinput.Model           │
    │ - Subject: textinput.Model      │
    │ - Body: textarea.Model          │
    │ - Attachments: list.Model       │
    │ - filepicker: filepicker.Model  │
    │ - keymap: KeyMap                │
    └─────────────────────────────────┘
    

Sources: [model.go51-85](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L51-L85)

### Delivery Layer

Pop supports two email delivery methods:

1.  **Resend API**: Uses the Resend.com API for email delivery
2.  **SMTP**: Uses standard SMTP protocols for email delivery

Sources: [main.go77-104](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L77-L104)
 [email.go66-132](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L66-L132)
 [email.go134-172](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L134-L172)

Email Processing Flow
---------------------

The email processing flow is handled by the `sendEmailCmd()` function, which processes email requests and sends them through the selected delivery method:

Sources: [email.go32-60](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L32-L60)
 [email.go194-209](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L194-L209)

### Markdown to HTML Conversion

Both delivery methods convert Markdown content to HTML:

    ┌────────────────────────────────┐
    │ Markdown Processing            │
    ├────────────────────────────────┤
    │ - Uses goldmark library        │
    │ - Optional unsafe HTML mode    │
    │ - Extensions: strikethrough,   │
    │   table, linkify (if unsafe)   │
    └────────────────────────────────┘
    

Sources: [email.go115-122](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L115-L122)
 [email.go137-153](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L137-L153)

Configuration Management
------------------------

Pop's architecture includes a configuration system that supports environment variables and command-line flags:

Sources: [main.go19-50](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L19-L50)
 [main.go196-228](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L196-L228)

### Environment Variables

Pop uses environment variables for default configuration:

| Environment Variable | Purpose |
| --- | --- |
| `RESEND_API_KEY` | Enables Resend as delivery method |
| `POP_FROM` | Default sender email address |
| `POP_SIGNATURE` | Default email signature |
| `POP_UNSAFE_HTML` | Enable unsafe HTML in email body |
| `POP_SMTP_HOST` | SMTP server host |
| `POP_SMTP_PORT` | SMTP server port |
| `POP_SMTP_USERNAME` | SMTP username |
| `POP_SMTP_PASSWORD` | SMTP password |
| `POP_SMTP_ENCRYPTION` | SMTP encryption type |
| `POP_SMTP_INSECURE_SKIP_VERIFY` | Skip TLS verification for SMTP |

Sources: [main.go19-50](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L19-L50)
 [main.go196-228](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L196-L228)

Application State Management
----------------------------

The TUI interface manages application state through a state machine pattern:

Sources: [model.go20-34](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L20-L34)
 [model.go229-362](https://github.com/charmbracelet/pop/blob/cdaa7255/model.go#L229-L362)

The state management system handles input focus and navigation between different form fields in the TUI, creating a smooth user experience while composing emails.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Architecture](https://deepwiki.com/charmbracelet/pop/3-architecture#architecture)
    
*   [System Architecture Overview](https://deepwiki.com/charmbracelet/pop/3-architecture#system-architecture-overview)
    
*   [Component Interactions and Flow](https://deepwiki.com/charmbracelet/pop/3-architecture#component-interactions-and-flow)
    
*   [Core Components](https://deepwiki.com/charmbracelet/pop/3-architecture#core-components)
    
*   [User Interface Layer](https://deepwiki.com/charmbracelet/pop/3-architecture#user-interface-layer)
    
*   [CLI Implementation](https://deepwiki.com/charmbracelet/pop/3-architecture#cli-implementation)
    
*   [TUI Implementation](https://deepwiki.com/charmbracelet/pop/3-architecture#tui-implementation)
    
*   [Core Logic Layer](https://deepwiki.com/charmbracelet/pop/3-architecture#core-logic-layer)
    
*   [Model Structure](https://deepwiki.com/charmbracelet/pop/3-architecture#model-structure)
    
*   [Delivery Layer](https://deepwiki.com/charmbracelet/pop/3-architecture#delivery-layer)
    
*   [Email Processing Flow](https://deepwiki.com/charmbracelet/pop/3-architecture#email-processing-flow)
    
*   [Markdown to HTML Conversion](https://deepwiki.com/charmbracelet/pop/3-architecture#markdown-to-html-conversion)
    
*   [Configuration Management](https://deepwiki.com/charmbracelet/pop/3-architecture#configuration-management)
    
*   [Environment Variables](https://deepwiki.com/charmbracelet/pop/3-architecture#environment-variables)
    
*   [Application State Management](https://deepwiki.com/charmbracelet/pop/3-architecture#application-state-management)
