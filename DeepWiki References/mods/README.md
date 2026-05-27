Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/mods](https://github.com/charmbracelet/mods "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 7 November 2025 ([07a05d](https://github.com/charmbracelet/mods/commits/07a05d5b)
)

*   [Introduction to Mods](https://deepwiki.com/charmbracelet/mods/1-introduction-to-mods)
    
*   [Installation and Setup](https://deepwiki.com/charmbracelet/mods/1.1-installation-and-setup)
    
*   [Quick Start Guide](https://deepwiki.com/charmbracelet/mods/1.2-quick-start-guide)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/mods/2-core-architecture)
    
*   [Application Entry Point](https://deepwiki.com/charmbracelet/mods/2.1-application-entry-point)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/mods/2.2-configuration-system)
    
*   [Bubble Tea UI Model](https://deepwiki.com/charmbracelet/mods/2.3-bubble-tea-ui-model)
    
*   [Message Stream Context](https://deepwiki.com/charmbracelet/mods/2.4-message-stream-context)
    
*   [Conversation Management](https://deepwiki.com/charmbracelet/mods/2.5-conversation-management)
    
*   [LLM Provider Integration](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration)
    
*   [Provider Configuration](https://deepwiki.com/charmbracelet/mods/3.1-provider-configuration)
    
*   [Client Initialization and Streaming](https://deepwiki.com/charmbracelet/mods/3.2-client-initialization-and-streaming)
    
*   [Model Resolution and Fallbacks](https://deepwiki.com/charmbracelet/mods/3.3-model-resolution-and-fallbacks)
    
*   [User Interface and Experience](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience)
    
*   [Terminal Capabilities and Styling](https://deepwiki.com/charmbracelet/mods/4.1-terminal-capabilities-and-styling)
    
*   [Animation System](https://deepwiki.com/charmbracelet/mods/4.2-animation-system)
    
*   [Output Rendering and Formatting](https://deepwiki.com/charmbracelet/mods/4.3-output-rendering-and-formatting)
    
*   [Data Persistence](https://deepwiki.com/charmbracelet/mods/5-data-persistence)
    
*   [Conversation Database](https://deepwiki.com/charmbracelet/mods/5.1-conversation-database)
    
*   [Cache System](https://deepwiki.com/charmbracelet/mods/5.2-cache-system)
    
*   [Supporting Systems](https://deepwiki.com/charmbracelet/mods/6-supporting-systems)
    
*   [Flag Parsing and Error Handling](https://deepwiki.com/charmbracelet/mods/6.1-flag-parsing-and-error-handling)
    
*   [Content Loading](https://deepwiki.com/charmbracelet/mods/6.2-content-loading)
    
*   [Message Utilities](https://deepwiki.com/charmbracelet/mods/6.3-message-utilities)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/mods/7-development-guide)
    
*   [Dependencies and Build System](https://deepwiki.com/charmbracelet/mods/7.1-dependencies-and-build-system)
    
*   [Testing Strategy](https://deepwiki.com/charmbracelet/mods/7.2-testing-strategy)
    
*   [Code Quality and CI/CD](https://deepwiki.com/charmbracelet/mods/7.3-code-quality-and-cicd)
    

Menu

Introduction to Mods
====================

Relevant source files

*   [README.md](https://github.com/charmbracelet/mods/blob/07a05d5b/README.md?plain=1)
    
*   [features.md](https://github.com/charmbracelet/mods/blob/07a05d5b/features.md?plain=1)
    
*   [go.mod](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod)
    

Purpose and Scope
-----------------

This document provides a high-level introduction to the `mods` CLI tool, explaining its core purpose, architecture, and key components. It serves as an entry point for understanding how the system works and how different parts integrate together.

For detailed information about specific subsystems, see the Configuration System ([2.2](https://deepwiki.com/charmbracelet/mods/2.2-configuration-system)
), LLM Provider Integration ([3](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration)
), User Interface components ([4](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience)
), and External Tool Integration ([5](https://deepwiki.com/charmbracelet/mods/5-data-persistence)
).

What is Mods
------------

Mods is a command-line interface tool designed to integrate Large Language Models (LLMs) into terminal workflows and Unix pipelines. Built with Go, it provides a sophisticated terminal user interface for interacting with multiple AI providers while maintaining the simplicity and composability expected from command-line tools.

The tool supports both direct command execution and interactive terminal sessions, allowing users to:

*   Send prompts directly to various LLM providers (OpenAI, Anthropic, Google, Cohere, Ollama)
*   Stream responses in real-time with rich terminal formatting
*   Manage conversation history with SQLite-based storage
*   Integrate with external tools through the MCP (Multi-Cloud Provider) protocol
*   Pipe input and output seamlessly with other Unix commands

System Architecture Overview
----------------------------

The following diagram shows how the main components of mods are organized, using actual code entities from the codebase:

Sources: [go.mod1-98](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L1-L98)
 [README.md1-262](https://github.com/charmbracelet/mods/blob/07a05d5b/README.md?plain=1#L1-L262)

User Interaction Flow
---------------------

This diagram illustrates how user commands flow through the system, highlighting the key decision points and code components involved:

Sources: [README.md108-177](https://github.com/charmbracelet/mods/blob/07a05d5b/README.md?plain=1#L108-L177)
 [features.md1-121](https://github.com/charmbracelet/mods/blob/07a05d5b/features.md?plain=1#L1-L121)

Core Components
---------------

### CLI and Configuration

The application entry point is [main.go](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go)
 which uses the Cobra library to parse commands and flags. Configuration is managed by [config.go](https://github.com/charmbracelet/mods/blob/07a05d5b/config.go)
 which loads settings from multiple sources with the following precedence: CLI flags > environment variables > user config file > embedded defaults.

| Component | File | Purpose |
| --- | --- | --- |
| CLI Entry Point | `main.go` | Application bootstrap, command parsing |
| Configuration | `config.go` | Multi-source config loading, validation |
| Default Settings | `config_template.yml` | Embedded default configuration |

### Terminal User Interface

The main application logic is implemented as a Bubble Tea model in [mods.go](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go)
 providing rich terminal interactions with animations ([anim.go](https://github.com/charmbracelet/mods/blob/07a05d5b/anim.go)
), styling ([styles.go](https://github.com/charmbracelet/mods/blob/07a05d5b/styles.go)
), and markdown rendering via Glamour.

### LLM Integration

Multiple LLM providers are supported through a unified streaming interface defined in [stream.go](https://github.com/charmbracelet/mods/blob/07a05d5b/stream.go)
 Each provider has its own implementation in the `internal/` directory:

*   `internal/anthropic/` - Anthropic Claude models
*   `internal/openai/` - OpenAI GPT models
*   `internal/google/` - Google Gemini models
*   Ollama and Cohere are integrated via their respective Go SDK packages

### Data Management

Conversation history is managed through a dual storage approach:

*   [db.go](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go)
     - SQLite database for conversation metadata (titles, timestamps, IDs)
*   File-based cache for full conversation content stored in XDG directories

### External Tool Integration

The [mcp.go](https://github.com/charmbracelet/mods/blob/07a05d5b/mcp.go)
 file implements the Multi-Cloud Provider protocol, allowing integration with external tools and services. The system also provides direct integration with system clipboard, external editors, and web browsers.

Key Design Principles
---------------------

1.  **Pipeline Compatibility**: Designed to work seamlessly with Unix pipes and command composition
2.  **Provider Abstraction**: Unified interface across different LLM providers while preserving provider-specific features
3.  **Rich Terminal Experience**: Advanced TUI capabilities while maintaining command-line simplicity
4.  **Conversation Persistence**: Intelligent conversation management with branching and continuation support
5.  **Extensibility**: MCP protocol support for integrating external tools and services

This architecture enables mods to function both as a simple command-line tool for automation and as a sophisticated interactive interface for complex AI-assisted workflows.

Sources: [go.mod1-98](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L1-L98)
 [README.md1-262](https://github.com/charmbracelet/mods/blob/07a05d5b/README.md?plain=1#L1-L262)
 [features.md1-121](https://github.com/charmbracelet/mods/blob/07a05d5b/features.md?plain=1#L1-L121)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Introduction to Mods](https://deepwiki.com/charmbracelet/mods#introduction-to-mods)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/mods#purpose-and-scope)
    
*   [What is Mods](https://deepwiki.com/charmbracelet/mods#what-is-mods)
    
*   [System Architecture Overview](https://deepwiki.com/charmbracelet/mods#system-architecture-overview)
    
*   [User Interaction Flow](https://deepwiki.com/charmbracelet/mods#user-interaction-flow)
    
*   [Core Components](https://deepwiki.com/charmbracelet/mods#core-components)
    
*   [CLI and Configuration](https://deepwiki.com/charmbracelet/mods#cli-and-configuration)
    
*   [Terminal User Interface](https://deepwiki.com/charmbracelet/mods#terminal-user-interface)
    
*   [LLM Integration](https://deepwiki.com/charmbracelet/mods#llm-integration)
    
*   [Data Management](https://deepwiki.com/charmbracelet/mods#data-management)
    
*   [External Tool Integration](https://deepwiki.com/charmbracelet/mods#external-tool-integration)
    
*   [Key Design Principles](https://deepwiki.com/charmbracelet/mods#key-design-principles)
