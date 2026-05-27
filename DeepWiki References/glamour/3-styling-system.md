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

Styling System
==============

Relevant source files

*   [CONTRIBUTING.md](https://github.com/charmbracelet/glamour/blob/69661fd5/CONTRIBUTING.md?plain=1)
    
*   [ansi/ansi.go](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/ansi.go)
    
*   [ansi/style.go](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go)
    
*   [ansi/templatehelper.go](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/templatehelper.go)
    
*   [examples/stdin-stdout-custom-styles/main.go](https://github.com/charmbracelet/glamour/blob/69661fd5/examples/stdin-stdout-custom-styles/main.go)
    
*   [internal/generate-style-json/main.go](https://github.com/charmbracelet/glamour/blob/69661fd5/internal/generate-style-json/main.go)
    
*   [styles/README.md](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/README.md?plain=1)
    
*   [styles/dracula.go](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/dracula.go)
    
*   [styles/dracula.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/dracula.json)
    
*   [styles/examples/block\_quote.style](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/examples/block_quote.style)
    
*   [styles/gallery/README.md](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/gallery/README.md?plain=1)
    
*   [styles/pink.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/pink.json)
    
*   [styles/styles.go](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/styles.go)
    
*   [styles/tokyo-night.go](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/tokyo-night.go)
    

The Glamour Styling System provides a flexible and comprehensive way to customize how Markdown content is rendered in terminal environments. This document explains the styling architecture, how styles are defined, the hierarchy of style elements, and methods for creating custom styles.

For information about how to apply styles within the Rendering API, see [Core API](https://deepwiki.com/charmbracelet/glamour/2-core-api)
. For details about specific element rendering, see [Element Renderers](https://deepwiki.com/charmbracelet/glamour/4-element-renderers)
.

Overview of the Styling System
------------------------------

Glamour's styling system allows for fine-grained control over how Markdown elements appear in the terminal. Styles control aspects like colors, text formatting (bold, italic), indentation, margins, and specialized rendering for lists, code blocks, tables and other elements.

Sources: [ansi/style.go97-139](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go#L97-L139)
 [styles/README.md1-9](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/README.md?plain=1#L1-L9)

Style Structure Hierarchy
-------------------------

The styling system is built on a hierarchy of style structures, from basic primitives to specialized style blocks.

Sources: [ansi/style.go38-95](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go#L38-L95)
 [ansi/style.go97-139](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go#L97-L139)

### Style Primitive

The `StylePrimitive` struct is the foundation of the styling system, defining basic styling attributes that can be applied to any element:

| Attribute | Type | Description |
| --- | --- | --- |
| BlockPrefix | string | Text to display before the element (in parent style) |
| BlockSuffix | string | Text to display after the element (in parent style) |
| Prefix | string | Text to display before the element |
| Suffix | string | Text to display after the element |
| Color | \*string | Text color (ANSI color code or hex) |
| BackgroundColor | \*string | Background color |
| Bold | \*bool | Makes text bold |
| Italic | \*bool | Makes text italic |
| CrossedOut | \*bool | Adds strikethrough to text |
| Underline | \*bool | Underlines text |
| Overlined | \*bool | Adds overline to text |
| Format | string | Format template for the element |

Sources: [ansi/style.go38-59](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go#L38-L59)
 [styles/README.md26-40](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/README.md?plain=1#L26-L40)

### Style Block

`StyleBlock` extends `StylePrimitive` with block-level formatting capabilities:

| Attribute | Type | Description |
| --- | --- | --- |
| Indent | \*uint | Indentation amount for the block |
| IndentToken | \*string | String used for indentation (e.g., "│ " for quotes) |
| Margin | \*uint | Margin around the block |

Sources: [ansi/style.go68-74](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go#L68-L74)
 [styles/README.md13-24](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/README.md?plain=1#L13-L24)

### Specialized Style Blocks

Glamour provides specialized style blocks for complex elements:

1.  `StyleCodeBlock`: For code blocks with syntax highlighting
    
    *   Additional attributes: Theme, Chroma settings
2.  `StyleList`: For rendering ordered and unordered lists
    
    *   Additional attribute: LevelIndent (indentation for nested lists)
3.  `StyleTable`: For table formatting
    
    *   Additional attributes: CenterSeparator, ColumnSeparator, RowSeparator
4.  `StyleTask`: For task/checkbox lists
    
    *   Additional attributes: Ticked, Unticked formatting

Sources: [ansi/style.go61-95](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go#L61-L95)
 [styles/README.md152-181](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/README.md?plain=1#L152-L181)

StyleConfig: The Complete Style Definition
------------------------------------------

The `StyleConfig` struct brings everything together, containing style definitions for all Markdown elements:

    StyleConfig
    ├── Document (StyleBlock)
    ├── BlockQuote (StyleBlock)
    ├── Paragraph (StyleBlock)
    ├── List (StyleList)
    ├── Heading (StyleBlock)
    ├── H1-H6 (StyleBlock)
    ├── Text, Emph, Strong, etc. (StylePrimitive)
    ├── Code (StyleBlock)
    ├── CodeBlock (StyleCodeBlock)
    ├── Table (StyleTable)
    └── Other specialized elements...
    

Each element inherits styles from its parent elements, with specific overrides applied as defined.

Sources: [ansi/style.go97-139](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go#L97-L139)

Built-in Themes
---------------

Glamour provides several built-in themes that offer comprehensive styling for all Markdown elements:

| Theme | Description |
| --- | --- |
| Dark | Dark background with light text, colorful syntax highlighting |
| Light | Light background with dark text |
| NoTTY | Plain text with no colors (for terminals without color support) |
| Dracula | Dracula color scheme (purple, pink, green accents) |
| Tokyo Night | Dark theme inspired by Tokyo night colors |
| Simple | Minimalist theme with basic formatting |
| Pink | Pink-accent theme |

Sources: [styles/dark.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/dark.json)
 [styles/light.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/light.json)
 [styles/notty.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/notty.json)
 [styles/dracula.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/dracula.json)
 [styles/gallery/README.md](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/gallery/README.md?plain=1)

### Dark Theme

The Dark theme is optimized for dark terminal backgrounds with light text. It features colorful headers, distinguishable code blocks, and a rich color palette for syntax highlighting.

Key style features:

*   Document color: Light gray (252)
*   Headings: Blue (39) with yellow H1 (228 on background 63)
*   Code: Salmon (203) on dark gray (236)
*   Links: Green (30) with underline

Sources: [styles/dark.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/dark.json)

### Light Theme

The Light theme is designed for light terminal backgrounds with dark text. It maintains good readability with distinct but not overly bright colors.

Key style features:

*   Document color: Dark gray (234)
*   Headings: Blue (27)
*   Code: Salmon (203) on very light gray (254)
*   Links: Cyan (36) with underline

Sources: [styles/light.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/light.json)

### NoTTY Theme

The NoTTY theme (pronounced "naughty") is designed for environments without ANSI color support. It uses only basic formatting and ASCII characters to represent styling.

Key style features:

*   No colors used
*   Strikethrough represented as `~~text~~`
*   Emphasis represented as `*text*`
*   Strong represented as `**text**`

Sources: [styles/notty.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/notty.json)
 [styles/gallery/README.md11-15](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/gallery/README.md?plain=1#L11-L15)

Style Inheritance and Cascading
-------------------------------

Glamour implements a cascading style system similar to CSS, where styles inherit from parent elements and can be overridden by more specific style definitions.

The cascading functions in the code handle this inheritance:

*   `cascadeStyles`: Merges multiple `StyleBlock` instances
*   `cascadeStylePrimitives`: Merges multiple `StylePrimitive` instances
*   `cascadeStylePrimitive`: Merges a parent and child `StylePrimitive`
*   `cascadeStyle`: Merges a parent and child `StyleBlock`

Sources: [ansi/style.go141-257](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go#L141-L257)

Custom Styles
-------------

You can define custom styles in JSON format. These follow the same structure as the built-in themes but allow for complete customization.

Here's an example of a custom style JSON structure:

To use a custom style, you would load it from a file and pass it to the renderer. The JSON file needs to follow the structure of the `StyleConfig` struct.

Sources: [styles/README.md](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/README.md?plain=1)
 [styles/dark.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/dark.json)
 [styles/light.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/light.json)
 [styles/notty.json](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/notty.json)

Style Elements Reference
------------------------

Here's a quick reference of the main elements that can be styled:

### Block Elements

*   `document`: The entire document
*   `block_quote`: Block quotes
*   `paragraph`: Paragraphs
*   `heading`, `h1`\-`h6`: Headings
*   `list`: Lists (ordered and unordered)
*   `code_block`: Code blocks
*   `table`: Tables
*   `definition_list`: Definition lists

### Inline Elements

*   `text`: Regular text
*   `emph`: Emphasized text (_italic_)
*   `strong`: Strong text (**bold**)
*   `strikethrough`: Strikethrough text
*   `code`: Inline code
*   `link`, `link_text`: Links and link text
*   `image`, `image_text`: Images and image text
*   `task`: Task list items

Each element supports the appropriate subset of styling attributes depending on whether it's a block or inline element.

Sources: [styles/README.md8-565](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/README.md?plain=1#L8-L565)

Advanced: Code Block Syntax Highlighting
----------------------------------------

The `StyleCodeBlock` includes a `Chroma` field that defines styles for syntax highlighting. Glamour uses the Chroma library for syntax highlighting in code blocks.

The `Chroma` struct contains style definitions for different syntax elements:

    Chroma
    ├── Text
    ├── Comment, CommentPreproc
    ├── Keyword, KeywordReserved, KeywordNamespace, KeywordType
    ├── Operator
    ├── Punctuation
    ├── Name, NameBuiltin, NameTag, NameAttribute, NameClass, etc.
    ├── Literal, LiteralNumber, LiteralString, etc.
    ├── Generic styles
    └── Background
    

Each of these elements can have its own styling defined using `StylePrimitive`.

Sources: [ansi/style.go4-36](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/style.go#L4-L36)
 [styles/dark.json92-181](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/dark.json#L92-L181)
 [styles/light.json92-181](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/light.json#L92-L181)

Template Formatting
-------------------

The styling system supports template-based formatting through the `Format` field in `StylePrimitive`. This is especially useful for elements like image text.

For example, the image text format in the dark theme is:

    "format": "Image: {{.text}} →"
    

Glamour provides a set of template helper functions in `TemplateFuncMap` for advanced formatting needs.

Sources: [ansi/templatehelper.go](https://github.com/charmbracelet/glamour/blob/69661fd5/ansi/templatehelper.go)
 [styles/dark.json81-84](https://github.com/charmbracelet/glamour/blob/69661fd5/styles/dark.json#L81-L84)

Conclusion
----------

The Glamour Styling System provides a powerful and flexible way to customize Markdown rendering in terminal environments. With its hierarchical style structure, cascading inheritance, and comprehensive element support, you can create highly customized terminal Markdown viewers tailored to your specific needs.

For more information on applying these styles through the API, see [Configuration Options](https://deepwiki.com/charmbracelet/glamour/2.3-configuration-options)
. For examples of custom styles, see [Custom Styles](https://deepwiki.com/charmbracelet/glamour/3.3-custom-styles)
.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Styling System](https://deepwiki.com/charmbracelet/glamour/3-styling-system#styling-system)
    
*   [Overview of the Styling System](https://deepwiki.com/charmbracelet/glamour/3-styling-system#overview-of-the-styling-system)
    
*   [Style Structure Hierarchy](https://deepwiki.com/charmbracelet/glamour/3-styling-system#style-structure-hierarchy)
    
*   [Style Primitive](https://deepwiki.com/charmbracelet/glamour/3-styling-system#style-primitive)
    
*   [Style Block](https://deepwiki.com/charmbracelet/glamour/3-styling-system#style-block)
    
*   [Specialized Style Blocks](https://deepwiki.com/charmbracelet/glamour/3-styling-system#specialized-style-blocks)
    
*   [StyleConfig: The Complete Style Definition](https://deepwiki.com/charmbracelet/glamour/3-styling-system#styleconfig-the-complete-style-definition)
    
*   [Built-in Themes](https://deepwiki.com/charmbracelet/glamour/3-styling-system#built-in-themes)
    
*   [Dark Theme](https://deepwiki.com/charmbracelet/glamour/3-styling-system#dark-theme)
    
*   [Light Theme](https://deepwiki.com/charmbracelet/glamour/3-styling-system#light-theme)
    
*   [NoTTY Theme](https://deepwiki.com/charmbracelet/glamour/3-styling-system#notty-theme)
    
*   [Style Inheritance and Cascading](https://deepwiki.com/charmbracelet/glamour/3-styling-system#style-inheritance-and-cascading)
    
*   [Custom Styles](https://deepwiki.com/charmbracelet/glamour/3-styling-system#custom-styles)
    
*   [Style Elements Reference](https://deepwiki.com/charmbracelet/glamour/3-styling-system#style-elements-reference)
    
*   [Block Elements](https://deepwiki.com/charmbracelet/glamour/3-styling-system#block-elements)
    
*   [Inline Elements](https://deepwiki.com/charmbracelet/glamour/3-styling-system#inline-elements)
    
*   [Advanced: Code Block Syntax Highlighting](https://deepwiki.com/charmbracelet/glamour/3-styling-system#advanced-code-block-syntax-highlighting)
    
*   [Template Formatting](https://deepwiki.com/charmbracelet/glamour/3-styling-system#template-formatting)
    
*   [Conclusion](https://deepwiki.com/charmbracelet/glamour/3-styling-system#conclusion)
