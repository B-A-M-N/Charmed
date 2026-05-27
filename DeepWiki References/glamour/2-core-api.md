Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/glamour](https://github.com/charmbracelet/glamour "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 14 February 2026 ([69661f](https://github.com/charmbracelet/glamour/commits/69661fd5)
)

*   [Overview](https://deepwiki.com/charmbracelet/glamour/1-overview)
    
*   [Installation](https://deepwiki.com/charmbracelet/glamour/1.1-installation)
    
*   [Quick Start](https://deepwiki.com/charmbracelet/glamour/1.2-quick-start)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/glamour/1.3-architecture-overview)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/glamour/1.4-dependencies)
    
*   [Core API](https://deepwiki.com/charmbracelet/glamour/2-core-api)
    
*   [TermRenderer](https://deepwiki.com/charmbracelet/glamour/2.1-termrenderer)
    
*   [Render Functions](https://deepwiki.com/charmbracelet/glamour/2.2-render-functions)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/glamour/2.3-configuration-options)
    
*   [Styling System](https://deepwiki.com/charmbracelet/glamour/3-styling-system)
    
*   [Style Configuration](https://deepwiki.com/charmbracelet/glamour/3.1-style-configuration)
    
*   [Built-in Themes](https://deepwiki.com/charmbracelet/glamour/3.2-built-in-themes)
    
*   [Custom Styles](https://deepwiki.com/charmbracelet/glamour/3.3-custom-styles)
    
*   [Chroma Syntax Highlighting](https://deepwiki.com/charmbracelet/glamour/3.4-chroma-syntax-highlighting)
    
*   [Element Renderers](https://deepwiki.com/charmbracelet/glamour/4-element-renderers)
    
*   [Rendering Lifecycle](https://deepwiki.com/charmbracelet/glamour/4.1-rendering-lifecycle)
    
*   [Block Elements](https://deepwiki.com/charmbracelet/glamour/4.2-block-elements)
    
*   [Tables](https://deepwiki.com/charmbracelet/glamour/4.3-tables)
    
*   [Code Blocks](https://deepwiki.com/charmbracelet/glamour/4.4-code-blocks)
    
*   [Links and Images](https://deepwiki.com/charmbracelet/glamour/4.5-links-and-images)
    
*   [Lists and Tasks](https://deepwiki.com/charmbracelet/glamour/4.6-lists-and-tasks)
    
*   [Layout and Text Flow](https://deepwiki.com/charmbracelet/glamour/5-layout-and-text-flow)
    
*   [BlockStack](https://deepwiki.com/charmbracelet/glamour/5.1-blockstack)
    
*   [Text Wrapping and Margins](https://deepwiki.com/charmbracelet/glamour/5.2-text-wrapping-and-margins)
    
*   [Advanced Usage](https://deepwiki.com/charmbracelet/glamour/6-advanced-usage)
    
*   [Custom Renderers](https://deepwiki.com/charmbracelet/glamour/6.1-custom-renderers)
    
*   [Environment-Based Configuration](https://deepwiki.com/charmbracelet/glamour/6.2-environment-based-configuration)
    
*   [Developer Guide](https://deepwiki.com/charmbracelet/glamour/7-developer-guide)
    
*   [Testing](https://deepwiki.com/charmbracelet/glamour/7.1-testing)
    
*   [Code Quality](https://deepwiki.com/charmbracelet/glamour/7.2-code-quality)
    
*   [Contributing](https://deepwiki.com/charmbracelet/glamour/7.3-contributing)
    

Menu

Core API
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1)
    
*   [ansi/renderer.go](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer.go)
    
*   [ansi/renderer\_test.go](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer_test.go)
    
*   [examples/custom\_renderer/main.go](https://github.com/charmbracelet/glamour/blob/69661fd5/examples/custom_renderer/main.go)
    
*   [examples/helloworld/main.go](https://github.com/charmbracelet/glamour/blob/69661fd5/examples/helloworld/main.go)
    
*   [glamour.go](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go)
    
*   [glamour\_test.go](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour_test.go)
    
*   [testdata/example.md](https://github.com/charmbracelet/glamour/blob/69661fd5/testdata/example.md?plain=1)
    

This document describes the primary public API for using Glamour to render Markdown content to ANSI-styled terminal output. It covers the `TermRenderer` type, simple render functions, and configuration options. For information about customizing styles and themes, see [Styling System](https://deepwiki.com/charmbracelet/glamour/3-styling-system)
. For details on how individual Markdown elements are rendered, see [Element Renderers](https://deepwiki.com/charmbracelet/glamour/4-element-renderers)
.

Scope
-----

The Core API provides two approaches for rendering Markdown:

1.  **Simple Functions** - One-line render functions for quick conversions: `Render()`, `RenderBytes()`, `RenderWithEnvironmentConfig()`
2.  **TermRenderer** - Reusable renderer instances with full configuration control via `NewTermRenderer()` and `TermRendererOption` functions

API Architecture
----------------

**Purpose**: This diagram shows the relationship between the public API functions, the `TermRenderer` type, configuration options, and internal components. Simple render functions internally create a `TermRenderer`, while direct instantiation allows reuse and streaming.

**Sources**: [glamour.go1-323](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L1-L323)

Core Components
---------------

### TermRenderer Type

The `TermRenderer` struct is the central component for rendering Markdown content. It encapsulates a Goldmark parser with an ANSI renderer and maintains internal buffers for conversion.

**Structure** ([glamour.go34-41](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L34-L41)
):

The type supports both immediate rendering via `Render()` and `RenderBytes()` methods, and streaming I/O via `Write()`, `Close()`, and `Read()` methods. See [TermRenderer](https://deepwiki.com/charmbracelet/glamour/2.1-termrenderer)
 for detailed documentation.

**Sources**: [glamour.go34-41](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L34-L41)

### Simple Render Functions

Glamour provides three convenience functions that create a one-time `TermRenderer` and render content immediately:

| Function | Signature | Purpose |
| --- | --- | --- |
| `Render` | `(string, string) -> (string, error)` | Render Markdown string with named style |
| `RenderBytes` | `([]byte, string) -> ([]byte, error)` | Render Markdown bytes with named style |
| `RenderWithEnvironmentConfig` | `(string) -> (string, error)` | Render using `GLAMOUR_STYLE` environment variable |

These functions are ideal for one-off rendering where configuration reuse is not needed. See [Render Functions](https://deepwiki.com/charmbracelet/glamour/2.2-render-functions)
 for details.

**Sources**: [glamour.go45-67](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L45-L67)
 [README.md23-36](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1#L23-L36)

### Configuration Options

The `TermRendererOption` type is a function that configures a `TermRenderer` during construction:

Options are passed to `NewTermRenderer()` and applied sequentially. Available options include:

*   **Style Options**: `WithStyles()`, `WithStandardStyle()`, `WithAutoStyle()`, `WithStylePath()`, `WithStylesFromJSONFile()`, `WithStylesFromJSONBytes()`
*   **Layout Options**: `WithWordWrap()`, `WithTableWrap()`, `WithInlineTableLinks()`
*   **Output Options**: `WithColorProfile()`, `WithEmoji()`, `WithPreservedNewLines()`, `WithChromaFormatter()`
*   **Context Options**: `WithBaseURL()`, `WithEnvironmentConfig()`
*   **Meta Options**: `WithOptions()` (apply multiple options)

See [Configuration Options](https://deepwiki.com/charmbracelet/glamour/2.3-configuration-options)
 for complete documentation.

**Sources**: [glamour.go31-251](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L31-L251)

API Usage Flow
--------------

**Purpose**: This sequence diagram illustrates the initialization and rendering flow. During initialization, options configure the `TermRenderer`, which creates and configures a Goldmark parser with an `ANSIRenderer`. During rendering, Goldmark parses Markdown to an AST and the `ANSIRenderer` converts it to ANSI output.

**Sources**: [glamour.go69-100](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L69-L100)
 [glamour.go284-295](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L284-L295)
 [ansi/renderer.go34-39](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer.go#L34-L39)

Common Usage Patterns
---------------------

### Pattern 1: Quick One-Off Rendering

For simple cases where you render once and don't need configuration reuse:

**Sources**: [README.md34](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1#L34-L34)
 [examples/helloworld/main.go18](https://github.com/charmbracelet/glamour/blob/69661fd5/examples/helloworld/main.go#L18-L18)

### Pattern 2: Reusable Configured Renderer

For applications that render multiple documents with consistent styling:

**Sources**: [README.md45-54](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1#L45-L54)
 [examples/custom\_renderer/main.go15-21](https://github.com/charmbracelet/glamour/blob/69661fd5/examples/custom_renderer/main.go#L15-L21)

### Pattern 3: Streaming I/O

For processing Markdown from streams or when output needs to be streamed:

The streaming pattern uses internal buffers ([glamour.go39-40](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L39-L40)
) to accumulate input via `Write()`, processes it when `Close()` is called, and makes output available via `Read()`.

**Sources**: [glamour.go253-282](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L253-L282)
 [glamour\_test.go19-47](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour_test.go#L19-L47)

### Pattern 4: Environment-Based Configuration

For applications that respect user preferences via environment variables:

The `GLAMOUR_STYLE` environment variable can be set to a built-in style name (e.g., "dark", "light") or a file path to a custom style JSON file.

**Sources**: [glamour.go52-55](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L52-L55)
 [glamour.go138-142](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L138-L142)
 [glamour.go297-304](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L297-L304)
 [README.md64-65](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1#L64-L65)

API Component Relationships
---------------------------

**Purpose**: This diagram shows how different API entry points connect through the configuration layer to the internal rendering engine. The `TermRenderer` maintains references to the Goldmark parser, ANSI options, and I/O buffers, bridging the public API to the rendering implementation.

**Sources**: [glamour.go34-41](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L34-L41)
 [glamour.go69-100](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L69-L100)
 [glamour.go306-322](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L306-L322)

Error Handling
--------------

All rendering functions return errors that may occur during:

1.  **Option Application** - Invalid configuration values (e.g., missing style files)
2.  **Markdown Parsing** - Malformed input (rare, as Goldmark is permissive)
3.  **I/O Operations** - Buffer read/write failures
4.  **Style Loading** - JSON parsing errors or file access issues

Example error handling:

**Sources**: [glamour.go86-89](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L86-L89)
 [glamour.go147-161](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L147-L161)
 [glamour.go275-278](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L275-L278)

Thread Safety
-------------

The `TermRenderer` type is **not thread-safe**. Each goroutine should maintain its own instance:

The internal buffers ([glamour.go39-40](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L39-L40)
) and Goldmark state are not protected by locks, so concurrent access would cause data races.

**Sources**: [glamour.go34-41](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L34-L41)

Internal Component Reference
----------------------------

The `TermRenderer` wraps several internal components:

| Component | Type | Purpose | Reference |
| --- | --- | --- | --- |
| `md` | `goldmark.Markdown` | Markdown parser with extensions | [glamour.go37](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L37-L37) |
| `ansiOptions` | `ansi.Options` | Configuration for ANSI renderer | [glamour.go38](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L38-L38) |
| `buf` | `bytes.Buffer` | Input accumulator for streaming | [glamour.go39](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L39-L39) |
| `renderBuf` | `bytes.Buffer` | Output buffer for streaming | [glamour.go40](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L40-L40) |

The `goldmark.Markdown` instance is configured with GFM (GitHub Flavored Markdown) and Definition List extensions ([glamour.go72-76](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L72-L76)
). The `ansi.Options` struct contains all style and layout configuration passed to the `ANSIRenderer` ([glamour.go81-84](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L81-L84)
).

**Sources**: [glamour.go34-41](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L34-L41)
 [glamour.go72-84](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L72-L84)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Core API](https://deepwiki.com/charmbracelet/glamour/2-core-api#core-api)
    
*   [Scope](https://deepwiki.com/charmbracelet/glamour/2-core-api#scope)
    
*   [API Architecture](https://deepwiki.com/charmbracelet/glamour/2-core-api#api-architecture)
    
*   [Core Components](https://deepwiki.com/charmbracelet/glamour/2-core-api#core-components)
    
*   [TermRenderer Type](https://deepwiki.com/charmbracelet/glamour/2-core-api#termrenderer-type)
    
*   [Simple Render Functions](https://deepwiki.com/charmbracelet/glamour/2-core-api#simple-render-functions)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/glamour/2-core-api#configuration-options)
    
*   [API Usage Flow](https://deepwiki.com/charmbracelet/glamour/2-core-api#api-usage-flow)
    
*   [Common Usage Patterns](https://deepwiki.com/charmbracelet/glamour/2-core-api#common-usage-patterns)
    
*   [Pattern 1: Quick One-Off Rendering](https://deepwiki.com/charmbracelet/glamour/2-core-api#pattern-1-quick-one-off-rendering)
    
*   [Pattern 2: Reusable Configured Renderer](https://deepwiki.com/charmbracelet/glamour/2-core-api#pattern-2-reusable-configured-renderer)
    
*   [Pattern 3: Streaming I/O](https://deepwiki.com/charmbracelet/glamour/2-core-api#pattern-3-streaming-io)
    
*   [Pattern 4: Environment-Based Configuration](https://deepwiki.com/charmbracelet/glamour/2-core-api#pattern-4-environment-based-configuration)
    
*   [API Component Relationships](https://deepwiki.com/charmbracelet/glamour/2-core-api#api-component-relationships)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/glamour/2-core-api#error-handling)
    
*   [Thread Safety](https://deepwiki.com/charmbracelet/glamour/2-core-api#thread-safety)
    
*   [Internal Component Reference](https://deepwiki.com/charmbracelet/glamour/2-core-api#internal-component-reference)
