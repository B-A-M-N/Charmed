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

Email Delivery Methods
======================

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/pop/blob/cdaa7255/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1)
    
*   [email.go](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go)
    
*   [main.go](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go)
    

Purpose and Scope
-----------------

This document provides an overview of the available email delivery methods in Pop, detailing how Pop determines which method to use, how each method is configured, and the internal mechanisms for sending emails. For specific configuration details, see [Resend API Configuration](https://deepwiki.com/charmbracelet/pop/6.1-resend-api-configuration)
 and [SMTP Configuration](https://deepwiki.com/charmbracelet/pop/6.2-smtp-configuration)
.

Overview of Delivery Methods
----------------------------

Pop supports two primary methods for sending emails:

1.  **Resend API** - A cloud-based email delivery service
2.  **SMTP** - Standard protocol for sending emails through mail servers

### Delivery Method Selection Process

Pop uses a priority-based system to determine which delivery method to use. This happens at runtime based on your configuration.

Sources: [main.go77-104](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L77-L104)

Resend API Delivery Method
--------------------------

The Resend API delivery method uses the Resend service to send emails. This is the recommended approach for most users due to its simplicity and reliability.

### Resend Configuration

To use the Resend delivery method, you need to set the `RESEND_API_KEY` environment variable or provide it through the `--resend.key` flag.

    export RESEND_API_KEY=your_api_key
    

Sources: [main.go24-25](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L24-L25)
 [main.go227-228](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L227-L228)

### How Resend Delivery Works

When using the Resend API, Pop uses the `resend-go` client library to send emails. The process follows these steps:

1.  Create a Resend client using your API key
2.  Convert Markdown body to HTML (if applicable)
3.  Create email request with recipients, subject, body, and attachments
4.  Send the email using the Resend API
5.  Process response or handle errors

Sources: [email.go134-172](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L134-L172)

SMTP Delivery Method
--------------------

The SMTP delivery method allows you to send emails through standard SMTP servers, including services like Gmail, Outlook, or your own mail servers.

### SMTP Configuration

To use the SMTP delivery method, you need to set the following environment variables or provide them through command-line flags:

| Environment Variable | Command-line Flag | Description | Default |
| --- | --- | --- | --- |
| `POP_SMTP_HOST` | `--smtp.host` | SMTP server hostname | \-  |
| `POP_SMTP_PORT` | `--smtp.port` | SMTP server port | 587 |
| `POP_SMTP_USERNAME` | `--smtp.username` | Username for authentication | \-  |
| `POP_SMTP_PASSWORD` | `--smtp.password` | Password for authentication | \-  |
| `POP_SMTP_ENCRYPTION` | `--smtp.encryption` | Encryption type (starttls, ssl, none) | "starttls" |
| `POP_SMTP_INSECURE_SKIP_VERIFY` | `--smtp.insecure` | Skip TLS verification | false |

Sources: [main.go33-50](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L33-L50)
 [main.go212-226](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L212-L226)

### Gmail Special Case

Pop provides special handling for Gmail accounts. If your SMTP username ends with "@gmail.com" and you haven't specified a host or port, Pop will automatically use Gmail's SMTP settings:

*   Host: smtp.gmail.com
*   Port: 587

Sources: [email.go62-83](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L62-L83)

### How SMTP Delivery Works

When using SMTP, Pop uses the `go-simple-mail` library to send emails. The process follows these steps:

1.  Create an SMTP client with your server configuration
2.  Establish connection to the SMTP server
3.  Create a new email message with recipients, subject, and body
4.  Convert Markdown body to HTML (if applicable)
5.  Add attachments (if any)
6.  Send the email
7.  Process response or handle errors

Sources: [email.go66-132](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L66-L132)

Delivery Method Components and Structure
----------------------------------------

The following diagram shows the key components involved in email delivery and their relationships:

Sources: [main.go77-104](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L77-L104)
 [email.go43-48](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L43-L48)
 [email.go66-172](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L66-L172)

Comparison of Delivery Methods
------------------------------

| Feature | Resend API | SMTP |
| --- | --- | --- |
| **Setup Complexity** | Low (only API key needed) | Medium (server details and credentials required) |
| **Dependency** | External service (Resend.com) | Any SMTP-compatible mail server |
| **Features** | HTML emails, attachments | HTML emails, attachments |
| **Configuration** | `RESEND_API_KEY` or `--resend.key` | Multiple SMTP-related environment variables or flags |
| **Best For** | Quick setup, modern features | Self-hosted solutions, specific mail servers |

Error Handling
--------------

Both delivery methods include error handling to help users recover from failures:

1.  If an email fails to send, Pop will save the email body to a temporary file
2.  The error message displayed to the user includes the path to this temporary file
3.  This allows users to recover their email content even if sending fails

Sources: [email.go52-56](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L52-L56)
 [email.go194-209](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L194-L209)

Selecting a Delivery Method
---------------------------

*   **For simplicity**: Resend API is recommended for most users
*   **For specific mail servers**: SMTP is more flexible
*   **For self-hosted solutions**: SMTP allows connecting to your own mail server

Sources: [main.go77-104](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L77-L104)
 [email.go43-59](https://github.com/charmbracelet/pop/blob/cdaa7255/email.go#L43-L59)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Email Delivery Methods](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#email-delivery-methods)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#purpose-and-scope)
    
*   [Overview of Delivery Methods](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#overview-of-delivery-methods)
    
*   [Delivery Method Selection Process](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#delivery-method-selection-process)
    
*   [Resend API Delivery Method](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#resend-api-delivery-method)
    
*   [Resend Configuration](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#resend-configuration)
    
*   [How Resend Delivery Works](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#how-resend-delivery-works)
    
*   [SMTP Delivery Method](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#smtp-delivery-method)
    
*   [SMTP Configuration](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#smtp-configuration)
    
*   [Gmail Special Case](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#gmail-special-case)
    
*   [How SMTP Delivery Works](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#how-smtp-delivery-works)
    
*   [Delivery Method Components and Structure](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#delivery-method-components-and-structure)
    
*   [Comparison of Delivery Methods](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#comparison-of-delivery-methods)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#error-handling)
    
*   [Selecting a Delivery Method](https://deepwiki.com/charmbracelet/pop/6-email-delivery-methods#selecting-a-delivery-method)
