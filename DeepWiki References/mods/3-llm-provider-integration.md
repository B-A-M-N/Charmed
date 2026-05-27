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

LLM Provider Integration
========================

Relevant source files

*   [config\_template.yml](https://github.com/charmbracelet/mods/blob/07a05d5b/config_template.yml)
    
*   [mods.go](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go)
    

This document describes the multi-provider LLM architecture in mods, which enables the tool to work with 15+ different AI providers through a unified abstraction layer. The system normalizes interactions with heterogeneous APIs (OpenAI, Anthropic, Cohere, Ollama, Google, Azure, Groq, and others) while providing provider-specific optimizations and fallback mechanisms.

For detailed information about how models and APIs are configured, see [Provider Configuration](https://deepwiki.com/charmbracelet/mods/3.1-provider-configuration)
. For implementation details of client initialization and streaming, see [Client Initialization and Streaming](https://deepwiki.com/charmbracelet/mods/3.2-client-initialization-and-streaming)
. For alias resolution and fallback logic, see [Model Resolution and Fallbacks](https://deepwiki.com/charmbracelet/mods/3.3-model-resolution-and-fallbacks)
.

* * *

Architecture Overview
---------------------

The LLM provider integration layer sits between the Bubble Tea UI ([Bubble Tea UI Model](https://deepwiki.com/charmbracelet/mods/2.3-bubble-tea-ui-model)
) and external AI services. It abstracts provider-specific details behind a common `stream.Client` interface, allowing the application to interact with different AI providers using uniform request and response patterns.

The integration layer handles three primary responsibilities:

1.  **Authentication**: Resolving API keys from configuration, environment variables, or command execution
2.  **Client Initialization**: Creating provider-specific clients with appropriate configurations
3.  **Stream Processing**: Managing streaming responses through a unified interface

Sources: [mods.go1-757](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L1-L757)
 [config\_template.yml1-533](https://github.com/charmbracelet/mods/blob/07a05d5b/config_template.yml#L1-L533)

* * *

Provider Abstraction Pattern
----------------------------

The abstraction pattern uses a two-phase approach:

1.  **Configuration Phase**: The `resolveModel` function matches user input to a specific API and Model configuration
2.  **Client Phase**: Provider-specific clients are initialized and wrapped behind the `stream.Client` interface

All provider clients implement the same interface method `Request(ctx, proto.Request) stream.Stream`, ensuring uniform handling regardless of the underlying API.

Sources: [mods.go276-454](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L276-L454)
 [mods.go704-747](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L704-L747)

* * *

Supported Providers
-------------------

The system supports the following provider categories:

### Primary Hosted Providers

| Provider | API Name | Default Base URL | Authentication Env | Local Support |
| --- | --- | --- | --- | --- |
| OpenAI | `openai` | `https://api.openai.com/v1` | `OPENAI_API_KEY` | No  |
| Anthropic | `anthropic` | `https://api.anthropic.com/v1` | `ANTHROPIC_API_KEY` | No  |
| Google | `google` | (SDK managed) | `GOOGLE_API_KEY` | No  |
| Cohere | `cohere` | `https://api.cohere.com/v1` | `COHERE_API_KEY` | No  |

Sources: [config\_template.yml72-178](https://github.com/charmbracelet/mods/blob/07a05d5b/config_template.yml#L72-L178)

### Specialized Providers

| Provider | API Name | Default Base URL | Authentication Env | Purpose |
| --- | --- | --- | --- | --- |
| Ollama | `ollama` | `http://localhost:11434` | (none) | Local inference |
| Azure | `azure` | Custom resource URL | `AZURE_OPENAI_KEY` | Enterprise OpenAI |
| Groq | `groq` | `https://api.groq.com/openai/v1` | `GROQ_API_KEY` | LPU acceleration |
| Perplexity | `perplexity` | `https://api.perplexity.ai` | `PERPLEXITY_API_KEY` | Search-augmented |
| Mistral | `mistral` | `https://api.mistral.ai/v1` | `MISTRAL_API_KEY` | European provider |
| DeepSeek | `deepseek` | `https://api.deepseek.com/` | `DEEPSEEK_API_KEY` | Reasoning models |

Sources: [config\_template.yml179-425](https://github.com/charmbracelet/mods/blob/07a05d5b/config_template.yml#L179-L425)

### Auxiliary Providers

| Provider | API Name | Notes |
| --- | --- | --- |
| Cerebras | `cerebras` | Fast inference platform |
| SambaNova | `sambanova` | Supports R1 distillations |
| LocalAI | `localai` | Self-hosted alternative |
| RunPod | `runpod` | Serverless GPU |
| GitHub Models | `github-models` | GitHub marketplace integration |

Sources: [config\_template.yml295-533](https://github.com/charmbracelet/mods/blob/07a05d5b/config_template.yml#L295-L533)

* * *

Request Flow
------------

The following diagram shows the complete flow from user request to provider response:

**Provider Selection and Request Flow**

Sources: [mods.go276-454](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L276-L454)
 [mods.go492-528](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L492-L528)
 [mods.go704-747](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L704-L747)

* * *

Authentication Methods
----------------------

The `ensureKey` function [mods.go456-490](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L456-L490)
 implements a cascading authentication resolution strategy:

### Resolution Priority

1.  **Direct API Key** (`api-key` in config): Stored directly in the configuration file
2.  **Environment Variable** (`api-key-env` in config): Read from a specified environment variable
3.  **Command Execution** (`api-key-cmd` in config): Execute a shell command to retrieve the key
4.  **Default Environment Variable**: Fall back to provider-specific default (e.g., `OPENAI_API_KEY`)

### Implementation

    Key Resolution Logic:
    ├─ if API.APIKey != "" → return API.APIKey
    ├─ else if API.APIKeyEnv != "" && API.APIKeyCmd == "" → return os.Getenv(API.APIKeyEnv)
    ├─ else if API.APIKeyCmd != "" → return exec.Command(parsed_cmd).CombinedOutput()
    ├─ else → return os.Getenv(defaultEnv)
    └─ if all fail → return modsError
    

This allows flexibility for different security models:

*   **Development**: Direct key in config or environment variable
*   **CI/CD**: Environment variables injected by the platform
*   **Secure Storage**: Command execution retrieves from password managers (e.g., `rbw get OPENAI_API_KEY`)

Sources: [mods.go456-490](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L456-L490)
 [config\_template.yml74-76](https://github.com/charmbracelet/mods/blob/07a05d5b/config_template.yml#L74-L76)

* * *

Client Initialization by Provider
---------------------------------

The `startCompletionCmd` function [mods.go276-454](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L276-L454)
 contains provider-specific initialization logic:

### Provider-Specific Configuration

| Provider API | Config Type | Special Handling |
| --- | --- | --- |
| `ollama` | `ollama.Config` | No authentication required, BaseURL defaults to localhost |
| `anthropic` | `anthropic.Config` | Requires `ANTHROPIC_API_KEY`, supports custom BaseURL |
| `google` | `google.Config` | Requires `GOOGLE_API_KEY`, model name passed to config, ThinkingBudget support |
| `cohere` | `cohere.Config` | Requires `COHERE_API_KEY`, supports custom BaseURL |
| `azure` | `openai.Config` | Requires `AZURE_OPENAI_KEY`, supports Azure AD authentication (APIType: "azure-ad") |
| `azure-ad` | `openai.Config` | Same as azure but with APIType set to "azure-ad" |
| Default | `openai.Config` | Assumes OpenAI-compatible API, requires `OPENAI_API_KEY` |

### Configuration Code Mapping

**Ollama Initialization** [mods.go314-318](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L314-L318)
:

    case "ollama":
        occfg = ollama.DefaultConfig()
        if api.BaseURL != "" {
            occfg.BaseURL = api.BaseURL
        }
    

**Anthropic Initialization** [mods.go319-327](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L319-L327)
:

    case "anthropic":
        key, err := m.ensureKey(api, "ANTHROPIC_API_KEY", ...)
        accfg = anthropic.DefaultConfig(key)
        if api.BaseURL != "" {
            accfg.BaseURL = api.BaseURL
        }
    

**Google Initialization** [mods.go328-334](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L328-L334)
:

    case "google":
        key, err := m.ensureKey(api, "GOOGLE_API_KEY", ...)
        gccfg = google.DefaultConfig(mod.Name, key)
        gccfg.ThinkingBudget = mod.ThinkingBudget
    

**OpenAI-Compatible Default** [mods.go359-368](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L359-L368)
:

    default:
        key, err := m.ensureKey(api, "OPENAI_API_KEY", ...)
        ccfg = openai.Config{
            AuthToken: key,
            BaseURL:   api.BaseURL,
        }
    

Sources: [mods.go313-368](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L313-L368)

* * *

Client Instantiation
--------------------

After configuration, clients are instantiated based on the model's API type [mods.go427-441](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L427-L441)
:

    switch mod.API {
    case "anthropic":
        client = anthropic.New(accfg)
    case "google":
        client = google.New(gccfg)
    case "cohere":
        client = cohere.New(cccfg)
    case "ollama":
        client, err = ollama.New(occfg)
    default:
        client = openai.New(ccfg)
    }
    

All clients implement the `stream.Client` interface, which has a single method:

    Request(ctx context.Context, request proto.Request) stream.Stream
    

This interface abstraction allows the remainder of the application to be provider-agnostic.

Sources: [mods.go427-444](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L427-L444)

* * *

HTTP Proxy Support
------------------

All provider clients support HTTP proxy configuration [mods.go370-380](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L370-L380)
:

    if cfg.HTTPProxy != "" {
        proxyURL, err := url.Parse(cfg.HTTPProxy)
        httpClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
        ccfg.HTTPClient = httpClient
        accfg.HTTPClient = httpClient
        cccfg.HTTPClient = httpClient
        occfg.HTTPClient = httpClient
    }
    

The proxy URL is set in the configuration and applied uniformly across all provider clients.

Sources: [mods.go370-380](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L370-L380)

* * *

Streaming Response Handling
---------------------------

The `receiveCompletionStreamCmd` function [mods.go492-528](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L492-L528)
 processes streaming responses uniformly across all providers:

**Streaming State Machine**

The streaming loop continues until `stream.Next()` returns false, at which point:

1.  Errors are checked via `stream.Err()`
2.  Tool calls are processed via `stream.CallTools()`
3.  Final messages are extracted via `stream.Messages()`

Sources: [mods.go492-528](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L492-L528)

* * *

Provider Import Structure
-------------------------

The codebase imports provider clients from internal packages [mods.go28-34](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L28-L34)
:

    import (
        "github.com/charmbracelet/mods/internal/anthropic"
        "github.com/charmbracelet/mods/internal/cohere"
        "github.com/charmbracelet/mods/internal/google"
        "github.com/charmbracelet/mods/internal/ollama"
        "github.com/charmbracelet/mods/internal/openai"
        "github.com/charmbracelet/mods/internal/proto"
        "github.com/charmbracelet/mods/internal/stream"
    )
    

Each provider package implements:

*   A `Config` struct for provider-specific configuration
*   A `DefaultConfig()` function for sensible defaults
*   A `New(Config)` function that returns a `stream.Client`

The `proto` package defines the common `Request` and `Message` types used across all providers.

The `stream` package defines the `Client` interface and `Stream` interface that all providers must implement.

Sources: [mods.go28-36](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L28-L36)

* * *

OpenAI-Compatible Providers
---------------------------

Many providers (Azure, Groq, LocalAI, RunPod, Mistral, DeepSeek, GitHub Models, etc.) use OpenAI-compatible APIs. These providers are initialized with the default case in the client instantiation switch, using `openai.New(ccfg)` [mods.go437](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L437-L437)

The only differences are:

*   Different BaseURL (specified in the API configuration)
*   Different authentication environment variables
*   Model naming conventions

This design allows new OpenAI-compatible providers to be added through configuration alone, without code changes.

Sources: [mods.go359-368](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L359-L368)
 [mods.go437](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L437-L437)
 [config\_template.yml206-533](https://github.com/charmbracelet/mods/blob/07a05d5b/config_template.yml#L206-L533)

* * *

Special Model Handling
----------------------

Certain models require special handling during request construction [mods.go386-392](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L386-L392)
:

### O1 Series Models

    if strings.HasPrefix(mod.Name, "o1") {
        cfg.MaxTokens = 0
    }
    

O1 models do not support the `max_tokens` parameter. The system automatically unsets this parameter for models with the "o1" prefix, using `max_completion_tokens` instead.

### JSON Format Mode

    if cfg.Format && config.FormatAs == "json" {
        request.ResponseFormat = &config.FormatAs
    }
    

OpenAI-compatible providers support structured JSON output via the `response_format` parameter. This is set when the `--format` flag is used with `format-as: json` in the configuration.

Sources: [mods.go386-441](https://github.com/charmbracelet/mods/blob/07a05d5b/mods.go#L386-L441)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [LLM Provider Integration](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#llm-provider-integration)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#architecture-overview)
    
*   [Provider Abstraction Pattern](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#provider-abstraction-pattern)
    
*   [Supported Providers](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#supported-providers)
    
*   [Primary Hosted Providers](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#primary-hosted-providers)
    
*   [Specialized Providers](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#specialized-providers)
    
*   [Auxiliary Providers](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#auxiliary-providers)
    
*   [Request Flow](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#request-flow)
    
*   [Authentication Methods](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#authentication-methods)
    
*   [Resolution Priority](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#resolution-priority)
    
*   [Implementation](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#implementation)
    
*   [Client Initialization by Provider](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#client-initialization-by-provider)
    
*   [Provider-Specific Configuration](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#provider-specific-configuration)
    
*   [Configuration Code Mapping](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#configuration-code-mapping)
    
*   [Client Instantiation](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#client-instantiation)
    
*   [HTTP Proxy Support](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#http-proxy-support)
    
*   [Streaming Response Handling](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#streaming-response-handling)
    
*   [Provider Import Structure](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#provider-import-structure)
    
*   [OpenAI-Compatible Providers](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#openai-compatible-providers)
    
*   [Special Model Handling](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#special-model-handling)
    
*   [O1 Series Models](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#o1-series-models)
    
*   [JSON Format Mode](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration#json-format-mode)
