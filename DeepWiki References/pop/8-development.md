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

Development
===========

Relevant source files

*   [.github/workflows/lint.yml](https://github.com/charmbracelet/pop/blob/cdaa7255/.github/workflows/lint.yml)
    
*   [go.mod](https://github.com/charmbracelet/pop/blob/cdaa7255/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/pop/blob/cdaa7255/go.sum)
    

This page provides information for developers who want to contribute to or modify the Pop email client. For installation and usage instructions, see [Overview](https://deepwiki.com/charmbracelet/pop/1-overview)
 and [Installation and Configuration](https://deepwiki.com/charmbracelet/pop/2-installation-and-configuration)
. For detailed build and test procedures, see [Building and Testing](https://deepwiki.com/charmbracelet/pop/8.1-building-and-testing)
.

Development Environment Setup
-----------------------------

To set up a development environment for Pop, you need:

1.  **Go 1.21+**: Required as specified in [go.mod3-4](https://github.com/charmbracelet/pop/blob/cdaa7255/go.mod#L3-L4)
    
2.  **Git**: For source code management
3.  **golangci-lint**: For code quality checks

Clone the repository and install dependencies:

Sources: [go.mod1-4](https://github.com/charmbracelet/pop/blob/cdaa7255/go.mod#L1-L4)
 [.github/workflows/lint.yml7-16](https://github.com/charmbracelet/pop/blob/cdaa7255/.github/workflows/lint.yml#L7-L16)

Project Architecture
--------------------

### Component Structure Diagram

Pop follows a modular architecture that separates user interfaces, email processing logic, and delivery mechanisms. The CLI is built with Cobra while the TUI uses Bubble Tea. Both interfaces leverage the same core email processing logic, which handles Markdown conversion via Goldmark and email delivery through either Resend API or SMTP.

Sources: [go.mod6-17](https://github.com/charmbracelet/pop/blob/cdaa7255/go.mod#L6-L17)

Key Dependencies
----------------

Pop relies on these core libraries:

| Library | Purpose | Version |
| --- | --- | --- |
| github.com/charmbracelet/bubbles | UI components | v0.21.0 |
| github.com/charmbracelet/bubbletea | TUI framework | v1.3.4 |
| github.com/charmbracelet/lipgloss | UI styling | v1.1.0 |
| github.com/spf13/cobra | CLI framework | v1.9.1 |
| github.com/resendlabs/resend-go | Resend API client | v1.7.0 |
| github.com/xhit/go-simple-mail/v2 | SMTP client | v2.16.0 |
| github.com/yuin/goldmark | Markdown processing | v1.7.8 |

Sources: [go.mod6-17](https://github.com/charmbracelet/pop/blob/cdaa7255/go.mod#L6-L17)

Development Workflow
--------------------

### Contribution Process Diagram

The project uses GitHub Actions for continuous integration, automatically running golangci-lint on each push and pull request. This workflow helps maintain code quality and consistency across contributions.

Sources: [.github/workflows/lint.yml1-21](https://github.com/charmbracelet/pop/blob/cdaa7255/.github/workflows/lint.yml#L1-L21)

Extending Pop
-------------

### Adding a New Email Provider

To add support for a new email delivery method:

1.  Create a client implementation that handles the delivery logic
2.  Add configuration options for the new provider
3.  Integrate with the existing email processing logic

### UI Customization

The terminal UI uses Bubble Tea's Model-View-Update pattern:

1.  **Model**: Defines the application state
2.  **View**: Renders the current state as text
3.  **Update**: Processes messages and updates state

When modifying the UI, maintain this pattern for consistency with the existing codebase.

Sources: [go.mod7-9](https://github.com/charmbracelet/pop/blob/cdaa7255/go.mod#L7-L9)

Dependency Management
---------------------

Pop uses Go modules for dependency management. Key operations include:

When adding new dependencies, ensure they are compatible with the project's license and maintain backward compatibility when possible.

Sources: [go.mod1-47](https://github.com/charmbracelet/pop/blob/cdaa7255/go.mod#L1-L47)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Development](https://deepwiki.com/charmbracelet/pop/8-development#development)
    
*   [Development Environment Setup](https://deepwiki.com/charmbracelet/pop/8-development#development-environment-setup)
    
*   [Project Architecture](https://deepwiki.com/charmbracelet/pop/8-development#project-architecture)
    
*   [Component Structure Diagram](https://deepwiki.com/charmbracelet/pop/8-development#component-structure-diagram)
    
*   [Key Dependencies](https://deepwiki.com/charmbracelet/pop/8-development#key-dependencies)
    
*   [Development Workflow](https://deepwiki.com/charmbracelet/pop/8-development#development-workflow)
    
*   [Contribution Process Diagram](https://deepwiki.com/charmbracelet/pop/8-development#contribution-process-diagram)
    
*   [Extending Pop](https://deepwiki.com/charmbracelet/pop/8-development#extending-pop)
    
*   [Adding a New Email Provider](https://deepwiki.com/charmbracelet/pop/8-development#adding-a-new-email-provider)
    
*   [UI Customization](https://deepwiki.com/charmbracelet/pop/8-development#ui-customization)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/pop/8-development#dependency-management)
