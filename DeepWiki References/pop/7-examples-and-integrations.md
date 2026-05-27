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

Examples and Integrations
=========================

Relevant source files

*   [LICENSE](https://github.com/charmbracelet/pop/blob/cdaa7255/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1)
    
*   [examples/cli.tape](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/cli.tape)
    
*   [examples/contacts.txt](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/contacts.txt)
    
*   [examples/demo.tape](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/demo.tape)
    
*   [examples/gum-example.tape](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/gum-example.tape)
    
*   [examples/invoice-example.tape](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/invoice-example.tape)
    
*   [examples/mods-example.tape](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/mods-example.tape)
    

This page provides practical examples of how to use Pop in various workflows and how to integrate it with other command-line tools to create powerful email pipelines. For basic usage instructions, see [Command Line Interface](https://deepwiki.com/charmbracelet/pop/4-command-line-interface)
 and [Terminal User Interface](https://deepwiki.com/charmbracelet/pop/5-terminal-user-interface)
.

Integration Patterns
--------------------

Pop can be easily combined with other command-line tools through standard Unix pipes, command substitution, and file operations. The flexible CLI interface accepts input from stdin, making it ideal for integration with text-generating tools.

Sources: [README.md102-108](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L102-L108)

Integration with Mods
---------------------

[Mods](https://github.com/charmbracelet/pop/blob/cdaa7255/Mods)
 is an AI tool that can generate text content. Combining Mods with Pop allows you to quickly generate email content using AI and send it directly from your terminal.

### Basic Usage Pattern

The `<<<` syntax passes the output of the Mods command as the body of the email.

### Example: AI-Generated Email Content

The `--preview` flag allows you to review and edit the AI-generated content before sending the email.

Sources: [README.md110-123](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L110-L123)
 [examples/mods-example.tape19-21](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/mods-example.tape#L19-L21)

Integration with Gum
--------------------

[Gum](https://github.com/charmbracelet/pop/blob/cdaa7255/Gum)
 is a tool for creating interactive command-line interfaces. When combined with Pop, it allows for interactive selection of email addresses and contacts.

### Common Usage Patterns

### Example: Interactive Email Address Selection

This command opens an interactive interface to:

1.  Choose the sender from a predefined list
2.  Filter and select the recipient from a contacts file

Sources: [README.md127-138](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L127-L138)
 [examples/gum-example.tape16-17](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/gum-example.tape#L16-L17)
 [examples/contacts.txt](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/contacts.txt)

Integration with Invoice
------------------------

[Invoice](https://github.com/charmbracelet/pop/blob/cdaa7255/Invoice)
 is a tool for generating invoices from the command line. When combined with Pop, you can create and send invoices entirely from your terminal.

### Basic Workflow

1.  Generate an invoice PDF using Invoice
2.  Attach the generated PDF to an email using Pop
3.  Send the email with an appropriate message

### Example: Generate and Send an Invoice

Sources: [README.md140-152](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L140-L152)
 [examples/invoice-example.tape16-18](https://github.com/charmbracelet/pop/blob/cdaa7255/examples/invoice-example.tape#L16-L18)

Advanced Integration Examples
-----------------------------

### Combining Multiple Tools

You can create more complex workflows by combining multiple tools. For example:

### Creating Shell Functions

For frequently used workflows, you can create shell functions in your `.bashrc` or `.zshrc`:

Sources: [README.md102-152](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L102-L152)

Integration with Continuous Integration Systems
-----------------------------------------------

Pop can be easily integrated into CI/CD pipelines to send automated emails, such as build notifications, deployment confirmations, or periodic reports.

### Example: GitHub Actions Integration

### Example: Sending Test Reports

Sources: Based on standard CI/CD practices with command-line tools

Tips for Custom Integrations
----------------------------

When creating your own integrations with Pop, consider these guidelines:

1.  **Stdin Integration**: Use `pop <<< "content"` or `echo "content" | pop` to pipe content directly into the email body
2.  **Command Substitution**: Use `$(command)` syntax to dynamically set email parameters
3.  **File Operations**: Use `--attach` to include files generated by other tools
4.  **Preview Mode**: For automated systems, skip the `--preview` flag; for user workflows, include it to allow review before sending
5.  **Exit Codes**: Check Pop's exit code to confirm email was sent successfully in automated scripts

Sources: [README.md102-152](https://github.com/charmbracelet/pop/blob/cdaa7255/README.md?plain=1#L102-L152)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Examples and Integrations](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#examples-and-integrations)
    
*   [Integration Patterns](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#integration-patterns)
    
*   [Integration with Mods](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#integration-with-mods)
    
*   [Basic Usage Pattern](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#basic-usage-pattern)
    
*   [Example: AI-Generated Email Content](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#example-ai-generated-email-content)
    
*   [Integration with Gum](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#integration-with-gum)
    
*   [Common Usage Patterns](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#common-usage-patterns)
    
*   [Example: Interactive Email Address Selection](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#example-interactive-email-address-selection)
    
*   [Integration with Invoice](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#integration-with-invoice)
    
*   [Basic Workflow](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#basic-workflow)
    
*   [Example: Generate and Send an Invoice](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#example-generate-and-send-an-invoice)
    
*   [Advanced Integration Examples](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#advanced-integration-examples)
    
*   [Combining Multiple Tools](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#combining-multiple-tools)
    
*   [Creating Shell Functions](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#creating-shell-functions)
    
*   [Integration with Continuous Integration Systems](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#integration-with-continuous-integration-systems)
    
*   [Example: GitHub Actions Integration](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#example-github-actions-integration)
    
*   [Example: Sending Test Reports](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#example-sending-test-reports)
    
*   [Tips for Custom Integrations](https://deepwiki.com/charmbracelet/pop/7-examples-and-integrations#tips-for-custom-integrations)
