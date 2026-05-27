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

Installation and Configuration
==============================

Relevant source files

*   [.goreleaser.yml](https://github.com/charmbracelet/pop/blob/cdaa7255/.goreleaser.yml)
    
*   [LICENSE](https://github.com/charmbracelet/pop/blob/cdaa7255/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1)
    
*   [go.mod](https://github.com/charmbracelet/pop/blob/cdaa7255/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/pop/blob/cdaa7255/go.sum)
    

This document provides detailed instructions on how to install the Pop email client and configure it for use. It covers installation methods, configuration options for email delivery, and available environment variables and command-line flags. For information about using the command-line interface, see [Command Line Interface](https://deepwiki.com/charmbracelet/pop/4-command-line-interface)
, and for the terminal user interface, see [Terminal User Interface](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface)
.

Installation
------------

Pop can be installed through various methods depending on your operating system and preferences.

### Package Managers

For macOS and Linux users, you can use Homebrew:

For Nix users:

For Arch Linux users:

### Go Install

If you have Go installed, you can install Pop directly:

### Binary Download

Alternatively, you can download pre-compiled binaries from the [releases page](https://github.com/charmbracelet/pop/blob/cdaa7255/releases%20page)

Sources: [README.md79-100](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L79-L100)

Configuration Overview
----------------------

Pop supports two primary email delivery methods: the Resend API and SMTP. At least one of these methods must be configured for Pop to send emails.

Sources: [README.md36-44](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L36-L44)

Delivery Method Configuration
-----------------------------

### Resend API Configuration

To use the Resend API delivery method, you need an API key from [Resend](https://resend.com/api-keys)
. Set the key using the `RESEND_API_KEY` environment variable:

You can also specify the API key directly with the `--resend.key` or `-r` flag:

> **Note**: If you're using a Resend account without a custom domain, you can use `onboarding@resend.dev` as your "From" email address.

Sources: [README.md44-52](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L44-L52)

### SMTP Configuration

To use SMTP for email delivery, configure the following environment variables:

You can also configure SMTP settings using command-line flags:

Sources: [README.md54-64](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L54-L64)

Environment Variables and Command-line Flags
--------------------------------------------

The following diagram shows the relationship between environment variables and command-line flags in Pop:

Sources: [README.md36-77](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L36-L77)

### Additional Environment Variables

These environment variables provide default values that are used when not explicitly specified in the command:

| Environment Variable | Description | Default |
| --- | --- | --- |
| `POP_FROM` | Default sender email address | None |
| `POP_SIGNATURE` | Email signature (supports Markdown) | None |
| `POP_UNSAFE_HTML` | Allow unsafe HTML in email body | `false` |

Example:

Sources: [README.md66-77](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L66-L77)

### Command-line Flag Reference

This table provides a comprehensive list of available command-line flags for Pop:

| Flag | Short | Description |
| --- | --- | --- |
| `--from` | `-f` | Sender email address |
| `--to` | `-t` | Recipient email address(es) |
| `--subject` | `-s` | Email subject |
| `--body` | `-b` | Email body (supports Markdown) |
| `--attach` | `-a` | Attach file(s) to the email |
| `--cc` |     | Carbon copy recipient(s) |
| `--bcc` |     | Blind carbon copy recipient(s) |
| `--signature` | `-x` | Email signature (supports Markdown) |
| `--unsafe` | `-u` | Allow unsafe HTML in email body |
| `--preview` |     | Preview the email before sending |
| `--resend.key` | `-r` | Resend API key |
| `--smtp.host` | `-H` | SMTP server hostname |
| `--smtp.port` | `-P` | SMTP server port |
| `--smtp.username` | `-U` | SMTP username |
| `--smtp.password` | `-p` | SMTP password |
| `--smtp.encryption` | `-e` | SMTP encryption (`tls`, `ssl`, or `none`) |
| `--smtp.insecure` | `-i` | Skip TLS verification for SMTP |

Sources: [README.md22-31](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L22-L31)

Configuration Precedence
------------------------

Pop applies configuration in the following order of precedence (highest to lowest):

1.  Command-line flags
2.  Environment variables
3.  Default values

This means that values specified via command-line flags will override environment variables, which in turn override default values.

Sources: [README.md36-77](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L36-L77)

Example Configurations
----------------------

### Resend API Example

### SMTP Example

Sources: [README.md22-77](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L22-L77)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Installation and Configuration](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#installation-and-configuration)
    
*   [Installation](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#installation)
    
*   [Package Managers](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#package-managers)
    
*   [Go Install](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#go-install)
    
*   [Binary Download](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#binary-download)
    
*   [Configuration Overview](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#configuration-overview)
    
*   [Delivery Method Configuration](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#delivery-method-configuration)
    
*   [Resend API Configuration](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#resend-api-configuration)
    
*   [SMTP Configuration](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#smtp-configuration)
    
*   [Environment Variables and Command-line Flags](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#environment-variables-and-command-line-flags)
    
*   [Additional Environment Variables](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#additional-environment-variables)
    
*   [Command-line Flag Reference](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#command-line-flag-reference)
    
*   [Configuration Precedence](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#configuration-precedence)
    
*   [Example Configurations](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#example-configurations)
    
*   [Resend API Example](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#resend-api-example)
    
*   [SMTP Example](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration#smtp-example)
