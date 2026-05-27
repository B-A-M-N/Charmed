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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1)
    
*   [ansi/renderer.go](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer.go)
    
*   [ansi/renderer\_test.go](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer_test.go)
    
*   [examples/custom\_renderer/main.go](https://github.com/charmbracelet/glamour/blob/69661fd5/examples/custom_renderer/main.go)
    
*   [examples/helloworld/main.go](https://github.com/charmbracelet/glamour/blob/69661fd5/examples/helloworld/main.go)
    
*   [glamour.go](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go)
    
*   [glamour\_test.go](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour_test.go)
    
*   [go.mod](https://github.com/charmbracelet/glamour/blob/69661fd5/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/glamour/blob/69661fd5/go.sum)
    
*   [testdata/example.md](https://github.com/charmbracelet/glamour/blob/69661fd5/testdata/example.md?plain=1)
    

Glamour is a Go library that renders Markdown documents as ANSI-styled output for terminal applications. It transforms Markdown text into beautifully formatted, color-coded terminal output with customizable themes and styles. This page provides a high-level understanding of Glamour's architecture, core components, and data processing pipeline.

For detailed information on specific subsystems, see:

*   Installation and setup: [Installation](https://deepwiki.com/charmbracelet/glamour/1.1-installation)
    
*   Basic usage examples: [Quick Start](https://deepwiki.com/charmbracelet/glamour/1.2-quick-start)
    
*   Detailed architecture explanation: [Architecture Overview](https://deepwiki.com/charmbracelet/glamour/1.3-architecture-overview)
    
*   Public API reference: [Core API](https://deepwiki.com/charmbracelet/glamour/2-core-api)
    
*   Theme customization: [Styling System](https://deepwiki.com/charmbracelet/glamour/3-styling-system)
    

**Sources:** [README.md1-20](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1#L1-L20)
 [glamour.go1-4](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L1-L4)

What Glamour Does
-----------------

Glamour converts Markdown text into ANSI escape sequences that produce styled terminal output. It operates as a pipeline that:

1.  Parses Markdown using the Goldmark parser into an Abstract Syntax Tree (AST)
2.  Traverses the AST and applies style rules from a theme configuration
3.  Renders each element to ANSI escape codes with proper layout, indentation, and text wrapping
4.  Outputs styled text suitable for display in terminal applications

The library is designed for CLI applications that need to display formatted documentation, help text, or other Markdown content directly in the terminal.

**Sources:** [glamour.go1-4](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L1-L4)
 [README.md12-19](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1#L12-L19)

Core Architecture Layers
------------------------

The architecture consists of four primary layers:

| Layer | Components | Responsibility |
| --- | --- | --- |
| **Public API** | `TermRenderer`, `Render()`, `NewTermRenderer()` | Provides user-facing functions and configuration options |
| **Rendering Engine** | `goldmark.Markdown`, `ansi.ANSIRenderer`, `ansi.Options` | Orchestrates parsing and conversion to ANSI output |
| **Styling** | `ansi.StyleConfig`, `styles.DefaultStyles`, JSON themes | Manages themes, colors, and formatting rules |
| **Element System** | `NewElement()`, `ElementRenderer`, `ElementFinisher` | Renders individual Markdown elements (headings, tables, etc.) |

**Sources:** [glamour.go34-41](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L34-L41)
 [glamour.go69-100](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L69-L100)
 [ansi/renderer.go17-32](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer.go#L17-L32)

Markdown Processing Pipeline
----------------------------

The processing pipeline follows these stages:

1.  **Input**: Markdown text is provided as `string` or `[]byte`
2.  **Parsing**: `goldmark.Markdown` converts text to an AST (Abstract Syntax Tree)
3.  **Traversal**: `ANSIRenderer.renderNode()` walks the AST nodes
4.  **Element Creation**: `NewElement()` creates appropriate element renderers based on node type
5.  **Style Application**: `StyleConfig` provides colors, formatting, and layout rules
6.  **Layout Management**: `BlockStack` tracks indentation and margins for nested structures
7.  **Buffering**: Output accumulates in `bytes.Buffer` instances
8.  **ANSI Generation**: Final output contains ANSI escape sequences for styling

**Sources:** [glamour.go285-295](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L285-L295)
 [ansi/renderer.go93-139](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer.go#L93-L139)

Key Component Mapping
---------------------

This diagram maps high-level concepts to actual code entities:

### Primary Types

| Code Entity | File Location | Purpose |
| --- | --- | --- |
| `TermRenderer` | [glamour.go36-41](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L36-L41) | Main public type for rendering Markdown |
| `ANSIRenderer` | [ansi/renderer.go30-32](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer.go#L30-L32) | Core renderer implementing goldmark's `renderer.NodeRenderer` |
| `Options` | [ansi/renderer.go18-27](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer.go#L18-L27) | Configuration for rendering (word wrap, colors, styles) |
| `StyleConfig` | `ansi/style.go` | Complete theme definition with element-specific styles |
| `RenderContext` | `ansi/context.go` | State management during rendering (block stack, buffers) |

**Sources:** [glamour.go34-41](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L34-L41)
 [ansi/renderer.go17-39](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer.go#L17-L39)

Main Entry Points
-----------------

Glamour provides two usage patterns:

**Simple Functions** (for one-off rendering):

*   `glamour.Render(markdown, stylePath)` - render with a named style
*   `glamour.RenderWithEnvironmentConfig(markdown)` - use `GLAMOUR_STYLE` environment variable

**Configurable Renderer** (for repeated use with custom options):

*   Create with `glamour.NewTermRenderer(options...)`
*   Configure using `TermRendererOption` functions like `WithWordWrap()`, `WithAutoStyle()`
*   Render with `(*TermRenderer).Render(markdown)` or `(*TermRenderer).RenderBytes(markdown)`

**Sources:** [glamour.go45-67](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L45-L67)
 [glamour.go70-100](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L70-L100)
 [glamour.go102-251](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L102-L251)

Rendering Lifecycle
-------------------

The two-phase rendering approach (`entering=true` then `entering=false`) allows elements to:

1.  Set up state before children render (push indent, open styling)
2.  Clean up after children complete (pop indent, close styling, flush buffers)

**Sources:** [ansi/renderer.go93-139](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/renderer.go#L93-L139)
 [glamour.go285-295](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L285-L295)

External Dependencies
---------------------

Glamour relies on carefully selected external libraries:

| Dependency | Version | Purpose |
| --- | --- | --- |
| `yuin/goldmark` | v1.7.13 | Markdown parsing to AST, extensibility framework |
| `charmbracelet/lipgloss` | v1.1.1+ | Terminal styling primitives (colors, borders, layout) |
| `alecthomas/chroma/v2` | v2.20.0 | Syntax highlighting for code blocks (31 token types) |
| `muesli/termenv` | v0.16.0 | Terminal capability detection (color support, dark mode) |
| `muesli/reflow` | v0.3.0 | Text wrapping and indentation utilities |
| `yuin/goldmark-emoji` | v1.0.6 | Emoji shortcode parsing (e.g., `:+1:` → 👍) |
| `microcosm-cc/bluemonday` | v1.0.27 | HTML sanitization for safe rendering |

**Sources:** [go.mod5-18](https://github.com/charmbracelet/glamour/blob/69661fd5/go.mod#L5-L18)

Style System Overview
---------------------

Styles can be specified in multiple ways:

*   Named built-in themes: `"dark"`, `"light"`, `"dracula"`, `"tokyo-night"`, `"pink"`, `"notty"`, `"ascii"`
*   Auto-detection: `styles.AutoStyle` selects dark/light based on terminal background
*   JSON files: Load custom styles from `.json` files
*   Programmatic: Pass `ansi.StyleConfig` directly to `WithStyles()`

For detailed style customization, see [Styling System](https://deepwiki.com/charmbracelet/glamour/3-styling-system)
.

**Sources:** [glamour.go121-161](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L121-L161)
 [glamour.go306-322](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L306-L322)

Usage Patterns
--------------

### Simple Rendering

**Sources:** [glamour.go45-55](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L45-L55)
 [README.md23-36](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1#L23-L36)

### Custom Configuration

**Sources:** [glamour.go70-100](https://github.com/charmbracelet/glamour/blob/69661fd5/glamour.go#L70-L100)
 [README.md42-54](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1#L42-L54)

Next Steps
----------

*   For installation instructions: [Installation](https://deepwiki.com/charmbracelet/glamour/1.1-installation)
    
*   For practical usage examples: [Quick Start](https://deepwiki.com/charmbracelet/glamour/1.2-quick-start)
    
*   For in-depth architecture details: [Architecture Overview](https://deepwiki.com/charmbracelet/glamour/1.3-architecture-overview)
    
*   For API reference: [Core API](https://deepwiki.com/charmbracelet/glamour/2-core-api)
    
*   For theme customization: [Styling System](https://deepwiki.com/charmbracelet/glamour/3-styling-system)
    
*   For element rendering details: [Element Renderers](https://deepwiki.com/charmbracelet/glamour/4-element-renderers)
    

**Sources:** [README.md1-103](https://github.com/charmbracelet/glamour/blob/69661fd5/README.md?plain=1#L1-L103)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/glamour#overview)
    
*   [What Glamour Does](https://deepwiki.com/charmbracelet/glamour#what-glamour-does)
    
*   [Core Architecture Layers](https://deepwiki.com/charmbracelet/glamour#core-architecture-layers)
    
*   [Markdown Processing Pipeline](https://deepwiki.com/charmbracelet/glamour#markdown-processing-pipeline)
    
*   [Key Component Mapping](https://deepwiki.com/charmbracelet/glamour#key-component-mapping)
    
*   [Primary Types](https://deepwiki.com/charmbracelet/glamour#primary-types)
    
*   [Main Entry Points](https://deepwiki.com/charmbracelet/glamour#main-entry-points)
    
*   [Rendering Lifecycle](https://deepwiki.com/charmbracelet/glamour#rendering-lifecycle)
    
*   [External Dependencies](https://deepwiki.com/charmbracelet/glamour#external-dependencies)
    
*   [Style System Overview](https://deepwiki.com/charmbracelet/glamour#style-system-overview)
    
*   [Usage Patterns](https://deepwiki.com/charmbracelet/glamour#usage-patterns)
    
*   [Simple Rendering](https://deepwiki.com/charmbracelet/glamour#simple-rendering)
    
*   [Custom Configuration](https://deepwiki.com/charmbracelet/glamour#custom-configuration)
    
*   [Next Steps](https://deepwiki.com/charmbracelet/glamour#next-steps)
