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

Overview
========

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/pop/blob/cdaa7255/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1)
    
*   [main.go](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go)
    

Pop is a command-line email client developed by Charmbracelet that enables users to send emails directly from their terminal. It provides both a Command Line Interface (CLI) for script integration and a Terminal User Interface (TUI) for interactive email composition.

For detailed installation instructions, see [Installation and Configuration](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration)
. For in-depth architecture details, see [Architecture](https://deepwiki.com/charmbracelet/pop/3-architecture)
.

Core Features
-------------

*   Send emails using either Resend API or SMTP servers
*   Support for Markdown-to-HTML conversion in email body
*   File attachment capability
*   Email previewing before sending
*   Configurable default settings via environment variables
*   Integration with other command-line tools

Sources: [README.md11-31](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L11-L31)
 [main.go72-76](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L72-L76)

System Architecture Overview
----------------------------

Pop's architecture consists of three main layers:

1.  **User Interface Layer**: Handles user input through either CLI or TUI
2.  **Email Processing Layer**: Processes email content, including Markdown to HTML conversion
3.  **Delivery Layer**: Sends emails via either Resend API or SMTP

Sources: [main.go72-158](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L72-L158)
 [README.md15-31](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L15-L31)

Command Line Interface
----------------------

Pop provides a powerful CLI designed for direct email sending and script integration:

The CLI uses Cobra for command parsing and supports numerous flags for configuring email parameters.

Major CLI flags include:

| Flag | Short | Description |
| --- | --- | --- |
| `--from` | `-f` | Email sender |
| `--to` | `-t` | Email recipients |
| `--subject` | `-s` | Email subject |
| `--body` | `-b` | Email body content |
| `--attach` | `-a` | File attachments |
| `--cc` | \-  | Carbon copy recipients |
| `--bcc` | \-  | Blind carbon copy recipients |
| `--preview` | \-  | Preview email before sending |

Sources: [main.go72-76](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L72-L76)
 [main.go197-230](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L197-L230)
 [README.md23-31](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L23-L31)

Terminal User Interface
-----------------------

Pop's interactive TUI provides a form-based interface for composing emails when launched without arguments:

The TUI is built with Bubble Tea and provides an intuitive interface for all email parameters.

Sources: [main.go137-155](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L137-L155)
 [README.md15-21](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L15-L21)

Delivery Methods
----------------

Pop supports two delivery methods for sending emails:

### 1\. Resend API

The Resend API is a modern email delivery service. To use it, you need to:

*   Set the `RESEND_API_KEY` environment variable or use `--resend.key` flag
*   Configure your sender email address

### 2\. SMTP

Standard SMTP protocol for sending via traditional email servers:

*   Set `POP_SMTP_*` environment variables or use `--smtp.*` flags
*   Configure host, port, username, password, and encryption settings

Sources: [main.go19-50](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L19-L50)
 [main.go76-104](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L76-L104)
 [README.md39-64](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L39-L64)

Configuration Options
---------------------

Pop provides extensive configuration via environment variables:

| Environment Variable | Purpose |
| --- | --- |
| `RESEND_API_KEY` | API key for Resend.com service |
| `POP_FROM` | Default sender email address |
| `POP_SIGNATURE` | Default signature for emails |
| `POP_UNSAFE_HTML` | Enable unsafe HTML in email body |
| `POP_SMTP_HOST` | SMTP server hostname |
| `POP_SMTP_PORT` | SMTP server port |
| `POP_SMTP_USERNAME` | SMTP username |
| `POP_SMTP_PASSWORD` | SMTP password |
| `POP_SMTP_ENCRYPTION` | SMTP encryption type |
| `POP_SMTP_INSECURE_SKIP_VERIFY` | Skip TLS verification |

All environment variables have corresponding command-line flags for overriding defaults.

Sources: [main.go19-50](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L19-L50)
 [main.go197-229](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L197-L229)

Integration Examples
--------------------

Pop can be combined with other tools to create powerful email workflows:

1.  **With mods**: Generate email content using AI
    
2.  **With gum**: Interactive recipient selection
    
3.  **With invoice**: Generate and send invoices
    

These integrations demonstrate Pop's flexibility as part of larger command-line workflows.

Sources: [README.md102-152](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L102-L152)

License
-------

Pop is available under the MIT license and is part of the Charmbracelet suite of tools.

Sources: [LICENSE1-22](https://github.com/charmbracelet/pop/blob/cdaa7255/LICENSE#L1-L22)
 [README.md162-178](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L162-L178)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/pop/1-overview#overview)
    
*   [Core Features](https://deepwiki.com/charmbracelet/pop/1-overview#core-features)
    
*   [System Architecture Overview](https://deepwiki.com/charmbracelet/pop/1-overview#system-architecture-overview)
    
*   [Command Line Interface](https://deepwiki.com/charmbracelet/pop/1-overview#command-line-interface)
    
*   [Terminal User Interface](https://deepwiki.com/charmbracelet/pop/1-overview#terminal-user-interface)
    
*   [Delivery Methods](https://deepwiki.com/charmbracelet/pop/1-overview#delivery-methods)
    
*   [1\. Resend API](https://deepwiki.com/charmbracelet/pop/1-overview#1-resend-api)
    
*   [2\. SMTP](https://deepwiki.com/charmbracelet/pop/1-overview#2-smtp)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/pop/1-overview#configuration-options)
    
*   [Integration Examples](https://deepwiki.com/charmbracelet/pop/1-overview#integration-examples)
    
*   [License](https://deepwiki.com/charmbracelet/pop/1-overview#license)
