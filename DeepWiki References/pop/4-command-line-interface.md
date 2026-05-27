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

Command Line Interface
======================

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/pop/blob/cdaa7255/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1)
    
*   [main.go](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go)
    

This document details how to use Pop's command line interface (CLI), including all available flags, options, and common usage patterns. For information about the interactive terminal user interface, see [Terminal User Interface](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface)
.

The Pop command line interface allows you to send emails directly from your terminal without entering an interactive mode, making it ideal for scripting and automation tasks.

Overview
--------

Pop's CLI enables you to send emails by providing all necessary information through command-line arguments. When all required parameters are specified, the email is sent immediately without launching the interactive interface.

Sources: [main.go72-158](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L72-L158)

Basic Usage
-----------

The simplest way to use Pop's CLI is to pipe or redirect your email content to the command while specifying the sender, recipient, and subject:

If any required parameter is missing (from, to, subject, or body), Pop will automatically launch the interactive TUI to collect the missing information.

Sources: [main.go118-156](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L118-L156)
 [README.md25-31](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L25-L31)

Command Line Arguments
----------------------

Pop accepts the following command line arguments:

| Flag | Short | Description | Environment Variable |
| --- | --- | --- | --- |
| `--from` | `-f` | Email sender address | `POP_FROM` |
| `--to` | `-t` | Recipients (can be specified multiple times) | \-  |
| `--cc` | \-  | Carbon copy recipients | \-  |
| `--bcc` | \-  | Blind carbon copy recipients | \-  |
| `--subject` | `-s` | Email subject | \-  |
| `--body` | `-b` | Email content | \-  |
| `--attach` | `-a` | Email attachments (can be specified multiple times) | \-  |
| `--preview` | \-  | Preview the email before sending | \-  |
| `--unsafe` | `-u` | Allow unsafe HTML in the email body | `POP_UNSAFE_HTML` |
| `--signature` | `-x` | Signature to display at the end of the email | `POP_SIGNATURE` |

Sources: [main.go199-211](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L199-L211)

Delivery Method Configuration
-----------------------------

Pop supports two delivery methods: Resend API and SMTP. You can configure either method through command line arguments or environment variables.

### Resend API Configuration

| Flag | Short | Description | Environment Variable |
| --- | --- | --- | --- |
| `--resend.key` | `-r` | API key for Resend.com | `RESEND_API_KEY` |

Sources: [main.go228](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L228-L228)
 [README.md44-51](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L44-L51)

### SMTP Configuration

| Flag | Short | Description | Environment Variable |
| --- | --- | --- | --- |
| `--smtp.host` | `-H` | SMTP server host | `POP_SMTP_HOST` |
| `--smtp.port` | `-P` | SMTP server port | `POP_SMTP_PORT` |
| `--smtp.username` | `-U` | SMTP username | `POP_SMTP_USERNAME` |
| `--smtp.password` | `-p` | SMTP password | `POP_SMTP_PASSWORD` |
| `--smtp.encryption` | `-e` | Encryption type (starttls, ssl, or none) | `POP_SMTP_ENCRYPTION` |
| `--smtp.insecure` | `-i` | Skip TLS verification | `POP_SMTP_INSECURE_SKIP_VERIFY` |

Sources: [main.go213-226](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L213-L226)
 [README.md54-64](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L54-L64)

CLI Process Flow
----------------

Sources: [main.go72-158](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L72-L158)

Environment Variables and Flag Relationships
--------------------------------------------

Sources: [main.go204-228](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L204-L228)

Delivery Method Selection Logic
-------------------------------

Pop automatically selects the delivery method based on the provided configuration:

Sources: [main.go77-104](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L77-L104)

Input Handling
--------------

Pop checks for input from several sources in the following order:

1.  Command line arguments (`--body` flag)
2.  Standard input (piped content)
3.  Interactive TUI (if standard input and `--body` are not provided)

Sources: [main.go106-116](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L106-L116)

Examples
--------

### Basic Usage

Send an email with content from a file:

### Multiple Recipients

Send to multiple recipients:

### Attachments

Send an email with attachments:

### Preview Mode

Preview an email before sending:

Sources: [README.md25-31](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L25-L31)
 [README.md117-121](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L117-L121)

Integration Examples
--------------------

Pop can be combined with other tools to create powerful email workflows:

### Using with Mods (AI)

### Using with Gum (User Input)

### Using with Invoice (Document Generation)

Sources: [README.md102-152](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L102-L152)

Error Handling
--------------

Pop provides clear error messages for common issues:

*   Missing delivery method configuration (no Resend API key or SMTP credentials)
*   Conflicting delivery method configuration (both Resend API key and SMTP credentials provided)
*   Email sending failures

If an email fails to send, Pop will provide an error message with details about the failure.

Sources: [main.go90-104](https://github.com/charmbracelet/pop/blob/cdaa7255/main.go#L90-L104)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Command Line Interface](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#command-line-interface)
    
*   [Overview](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#overview)
    
*   [Basic Usage](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#basic-usage)
    
*   [Command Line Arguments](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#command-line-arguments)
    
*   [Delivery Method Configuration](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#delivery-method-configuration)
    
*   [Resend API Configuration](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#resend-api-configuration)
    
*   [SMTP Configuration](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#smtp-configuration)
    
*   [CLI Process Flow](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#cli-process-flow)
    
*   [Environment Variables and Flag Relationships](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#environment-variables-and-flag-relationships)
    
*   [Delivery Method Selection Logic](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#delivery-method-selection-logic)
    
*   [Input Handling](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#input-handling)
    
*   [Examples](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#examples)
    
*   [Basic Usage](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#basic-usage-1)
    
*   [Multiple Recipients](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#multiple-recipients)
    
*   [Attachments](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#attachments)
    
*   [Preview Mode](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#preview-mode)
    
*   [Integration Examples](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#integration-examples)
    
*   [Using with Mods (AI)](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#using-with-mods-ai)
    
*   [Using with Gum (User Input)](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#using-with-gum-user-input)
    
*   [Using with Invoice (Document Generation)](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#using-with-invoice-document-generation)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/pop/4-command-line-interface#error-handling)
