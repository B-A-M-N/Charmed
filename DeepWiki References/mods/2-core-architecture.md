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

Core Architecture
=================

Relevant source files

*   [examples.go](https://github.com/charmbracelet/mods/blob/07a05d5b/examples.go)
    
*   [go.mod](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod)
    
*   [main.go](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go)
    
*   [mods.go](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go)
    

Purpose and Scope
-----------------

This document explains the overall system architecture and key components that make up the mods application. It covers the high-level design patterns, component relationships, and architectural decisions that enable mods to function as a sophisticated CLI tool for interacting with multiple Large Language Models.

For specific implementation details of individual components, see [CLI Interface and Entry Point](https://deepwiki.com/charmbracelet/mods/2.1-application-entry-point)
, [Configuration System](https://deepwiki.com/charmbracelet/mods/2.2-configuration-system)
, [LLM Interaction Model](https://deepwiki.com/charmbracelet/mods/2.3-bubble-tea-ui-model)
, and [Conversation Management](https://deepwiki.com/charmbracelet/mods/2.4-message-stream-context)
. For LLM provider integrations, see [LLM Provider Integration](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration)
.

System Overview
---------------

Mods is architected as a layered system that transforms command-line inputs into rich, interactive terminal experiences for LLM interactions. The architecture follows a clear separation of concerns across four primary layers: User Interface, Core Logic, AI Integration, and Data Persistence.

### System Architecture with Code Entities

Sources: [main.go70-241](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L70-L241)
 [mods.go52-757](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L52-L757)
 [config.go](https://github.com/charmbracelet/mods/blob/07a05d5b/config.go)
 [stream.go](https://github.com/charmbracelet/mods/blob/07a05d5b/stream.go)
 [db.go](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go)
 [internal/cache](https://github.com/charmbracelet/mods/blob/07a05d5b/internal/cache)

Core Components
---------------

### Entry Point and Command Processing

The application entry point is defined in `main()` [main.go336-399](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L336-L399)
 which initializes the configuration, database, and executes `rootCmd`. The `rootCmd.RunE` function [main.go80-241](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L80-L241)
 contains the primary execution logic.

#### Execution Flow from main() to Bubble Tea

The system determines whether to use TTY-aware rendering based on `isInputTTY()` and `isOutputTTY()` checks [main.go85-92](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L85-L92)
 enabling both interactive and pipeline modes.

Sources: [main.go336-399](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L336-L399)
 [main.go80-241](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L80-L241)
 [mods.go80-105](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L80-L105)
 [mods.go120-122](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L120-L122)

### Configuration System

The configuration system implements a four-tier cascade with precedence from lowest to highest. The `Config` struct in `config.go` is populated by `ensureConfig()` [config.go](https://github.com/charmbracelet/mods/blob/07a05d5b/config.go)
 which loads and merges settings from multiple sources.

#### Configuration Cascade and Precedence

| Tier | Source | Loader | Precedence |
| --- | --- | --- | --- |
| 1   | Embedded Template | `config_template.yml` | Lowest |
| 2   | User Config File | `~/.config/mods/mods.yml` (XDG) | Low |
| 3   | Environment Variables | `MODS_*`, `OPENAI_API_KEY`, etc. | High |
| 4   | CLI Flags | `--model`, `--api`, `--temperature`, etc. | Highest |

The `initFlags()` function [main.go246-334](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L246-L334)
 binds command-line flags to `Config` struct fields using `pflag`. Key configuration fields include:

*   `Config.Model` [main.go248](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L248-L248)
    : Model identifier (e.g., "gpt-4o", "claude-3.5-sonnet")
*   `Config.API` [main.go250](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L250-L250)
    : API provider name
*   `Config.Temperature`, `Config.TopP`, `Config.TopK` [main.go272-275](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L272-L275)
    : Sampling parameters
*   `Config.MaxTokens` [main.go270](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L270-L270)
    : Token limit
*   `Config.Continue`, `Config.Title` [main.go257-260](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L257-L260)
    : Conversation management

Sources: [config.go](https://github.com/charmbracelet/mods/blob/07a05d5b/config.go)
 [config\_template.yml](https://github.com/charmbracelet/mods/blob/07a05d5b/config_template.yml)
 [main.go246-334](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L246-L334)
 [main.go71-72](https://github.com/charmbracelet/mods/blob/07a05d5b/main.go#L71-L72)
 [go.mod10](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L10-L10)
 [go.mod36](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L36-L36)

### LLM Request and Streaming Architecture

The LLM interaction begins in `startCompletionCmd()` [mods.go276-454](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L276-L454)
 which resolves the model, sets up the stream context, and initializes the appropriate provider client. The streaming system abstracts provider APIs behind a unified interface.

#### Request Initialization and Provider Selection

The `resolveModel()` function [mods.go704-747](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L704-L747)
 matches the user's model string against configured APIs and their models, handling aliases [mods.go710](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L710-L710)
 and validating model availability [mods.go715-720](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L715-L720)

#### Streaming Response Processing

The `receiveCompletionStreamCmd()` function [mods.go492-528](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L492-L528)
 processes streaming chunks iteratively, appending content via `appendToOutput()` [mods.go635-661](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L635-L661)
 which handles both raw output to stdout and glamour-rendered markdown for TTY displays [mods.go637-654](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L637-L654)

Sources: [mods.go276-454](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L276-L454)
 [mods.go456-490](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L456-L490)
 [mods.go492-528](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L492-L528)
 [mods.go635-661](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L635-L661)
 [mods.go704-747](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L704-L747)
 [stream.go](https://github.com/charmbracelet/mods/blob/07a05d5b/stream.go)
 [config.go](https://github.com/charmbracelet/mods/blob/07a05d5b/config.go)
 [internal/anthropic](https://github.com/charmbracelet/mods/blob/07a05d5b/internal/anthropic)
 [internal/openai](https://github.com/charmbracelet/mods/blob/07a05d5b/internal/openai)
 [internal/google](https://github.com/charmbracelet/mods/blob/07a05d5b/internal/google)
 [internal/ollama](https://github.com/charmbracelet/mods/blob/07a05d5b/internal/ollama)
 [internal/cohere](https://github.com/charmbracelet/mods/blob/07a05d5b/internal/cohere)

### Bubble Tea State Machine

The `Mods` struct [mods.go52-78](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L52-L78)
 implements the `tea.Model` interface with a state-driven architecture. The `state` type [mods.go39-48](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L39-L48)
 defines six distinct states that govern the application lifecycle.

#### Mods State Machine

The `Mods.Update()` method [mods.go125-215](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L125-L215)
 handles state transitions by processing tea.Msg messages. The `Mods.View()` method [mods.go222-257](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L222-L257)
 renders UI based on the current state, displaying animations during `requestState` [mods.go228-230](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L228-L230)
 and streaming content during `responseState` [mods.go232-249](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L232-L249)

Sources: [mods.go39-48](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L39-L48)
 [mods.go52-78](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L52-L78)
 [mods.go120-122](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L120-L122)
 [mods.go125-215](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L125-L215)
 [mods.go222-257](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L222-L257)
 [mods.go276-454](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L276-L454)
 [mods.go492-528](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L492-L528)
 [mods.go604-615](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L604-L615)
 [mods.go635-661](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L635-L661)

Key Architectural Patterns
--------------------------

### Provider Abstraction with stream.Client Interface

The system abstracts LLM provider differences behind the `stream.Client` interface. Each provider package (anthropic, openai, google, ollama, cohere) implements this interface, enabling runtime provider switching in `startCompletionCmd()` [mods.go427-441](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L427-L441)
 without changing core logic.

The `resolveModel()` function [mods.go704-747](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L704-L747)
 maps user input to provider-specific configurations by iterating through `config.APIs` [mods.go705](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L705-L705)
 and matching against `api.Models` [mods.go709-720](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L709-L720)

### Dual Storage Architecture for Conversations

Conversation persistence uses two complementary storage mechanisms:

| Storage | Location | Content | Implementation |
| --- | --- | --- | --- |
| SQLite Database | `~/.cache/mods/conversations/mods.db` | Metadata: ID, Title, UpdatedAt, API, Model | `convoDB` struct in `db.go` |
| Cache Directory | `~/.cache/mods/conversations/<SHA1>.json` | Full message arrays: `[]proto.Message` | `cache.Conversations` in `internal/cache` |

The `convoDB.Save()` method [db.go](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go)
 stores lightweight metadata for fast listing and searching, while `cache.Write()` [internal/cache](https://github.com/charmbracelet/mods/blob/07a05d5b/internal/cache)
 persists complete message histories. This separation enables efficient conversation browsing without loading full contents.

The `findCacheOpsDetails()` function [mods.go534-587](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L534-L587)
 coordinates both storage systems, resolving conversation IDs from partial SHA-1 hashes or titles via `db.Find()` [db.go](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go)
 and loading message arrays via `cache.Read()` [internal/cache](https://github.com/charmbracelet/mods/blob/07a05d5b/internal/cache)

### Event-Driven Message Passing with Bubble Tea

The Bubble Tea framework uses message passing for asynchronous updates. Key message types:

*   `completionInput` [mods.go108-110](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L108-L110)
    : Wraps stdin content
*   `completionOutput` [mods.go113-117](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L113-L117)
    : Wraps streaming chunks and error handlers
*   `cacheDetailsMsg` [mods.go530-532](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L530-L532)
    : Carries conversation metadata
*   `modsError` [error.go](https://github.com/charmbracelet/mods/blob/07a05d5b/error.go)
    : Signals error states

The `Mods.Update()` method [mods.go125-215](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L125-L215)
 processes messages and dispatches commands that return new messages, creating an event loop. Streaming is handled by returning `receiveCompletionStreamCmd()` [mods.go182-185](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L182-L185)
 which recursively processes chunks until the stream completes [mods.go494-527](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L494-L527)

Sources: [mods.go427-441](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L427-L441)
 [mods.go492-528](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L492-L528)
 [mods.go534-587](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L534-L587)
 [mods.go704-747](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L704-L747)
 [db.go](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go)
 [internal/cache](https://github.com/charmbracelet/mods/blob/07a05d5b/internal/cache)
 [stream.go](https://github.com/charmbracelet/mods/blob/07a05d5b/stream.go)
 [go.mod23](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L23-L23)
 [go.mod6](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L6-L6)

Technology Stack
----------------

### Core Dependencies

| Component | Package | Purpose |
| --- | --- | --- |
| CLI Framework | `spf13/cobra` + `spf13/pflag` | Command-line interface |
| TUI Framework | `charmbracelet/bubbletea` | Interactive terminal UI |
| Styling | `charmbracelet/lipgloss` | Terminal styling and layout |
| Configuration | `caarlos0/env/v9` + `gopkg.in/yaml.v3` | Multi-source configuration |
| Database | `modernc.org/sqlite` + `jmoiron/sqlx` | Conversation storage |
| LLM Providers | Multiple SDKs | AI model integrations |

### External Tool Integration

The MCP (Model Context Protocol) system in `mcp.go` enables integration with external tools and services, extending the application's capabilities beyond basic LLM interaction.

Sources: [mcp.go](https://github.com/charmbracelet/mods/blob/07a05d5b/mcp.go)
 [go.mod25](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L25-L25)
 [go.mod37](https://github.com/charmbracelet/mods/blob/07a05d5b/go.mod#L37-L37)

The architecture successfully balances simplicity with extensibility, providing a robust foundation for sophisticated LLM interactions while maintaining clean separation of concerns across all system layers.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core Architecture](https://deepwiki.com/charmbracelet/mods/2-core-architecture#core-architecture)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/mods/2-core-architecture#purpose-and-scope)
    
*   [System Overview](https://deepwiki.com/charmbracelet/mods/2-core-architecture#system-overview)
    
*   [System Architecture with Code Entities](https://deepwiki.com/charmbracelet/mods/2-core-architecture#system-architecture-with-code-entities)
    
*   [Core Components](https://deepwiki.com/charmbracelet/mods/2-core-architecture#core-components)
    
*   [Entry Point and Command Processing](https://deepwiki.com/charmbracelet/mods/2-core-architecture#entry-point-and-command-processing)
    
*   [Execution Flow from main() to Bubble Tea](https://deepwiki.com/charmbracelet/mods/2-core-architecture#execution-flow-from-main-to-bubble-tea)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/mods/2-core-architecture#configuration-system)
    
*   [Configuration Cascade and Precedence](https://deepwiki.com/charmbracelet/mods/2-core-architecture#configuration-cascade-and-precedence)
    
*   [LLM Request and Streaming Architecture](https://deepwiki.com/charmbracelet/mods/2-core-architecture#llm-request-and-streaming-architecture)
    
*   [Request Initialization and Provider Selection](https://deepwiki.com/charmbracelet/mods/2-core-architecture#request-initialization-and-provider-selection)
    
*   [Streaming Response Processing](https://deepwiki.com/charmbracelet/mods/2-core-architecture#streaming-response-processing)
    
*   [Bubble Tea State Machine](https://deepwiki.com/charmbracelet/mods/2-core-architecture#bubble-tea-state-machine)
    
*   [Mods State Machine](https://deepwiki.com/charmbracelet/mods/2-core-architecture#mods-state-machine)
    
*   [Key Architectural Patterns](https://deepwiki.com/charmbracelet/mods/2-core-architecture#key-architectural-patterns)
    
*   [Provider Abstraction with stream.Client Interface](https://deepwiki.com/charmbracelet/mods/2-core-architecture#provider-abstraction-with-streamclient-interface)
    
*   [Dual Storage Architecture for Conversations](https://deepwiki.com/charmbracelet/mods/2-core-architecture#dual-storage-architecture-for-conversations)
    
*   [Event-Driven Message Passing with Bubble Tea](https://deepwiki.com/charmbracelet/mods/2-core-architecture#event-driven-message-passing-with-bubble-tea)
    
*   [Technology Stack](https://deepwiki.com/charmbracelet/mods/2-core-architecture#technology-stack)
    
*   [Core Dependencies](https://deepwiki.com/charmbracelet/mods/2-core-architecture#core-dependencies)
    
*   [External Tool Integration](https://deepwiki.com/charmbracelet/mods/2-core-architecture#external-tool-integration)
