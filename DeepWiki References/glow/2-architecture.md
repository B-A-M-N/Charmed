Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/glow](https://github.com/charmbracelet/glow "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 2 January 2026 ([752de9](https://github.com/charmbracelet/glow/commits/752de97c)
)

*   [Overview](https://deepwiki.com/charmbracelet/glow/1-overview)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/glow/1.1-dependencies)
    
*   [Installation and Usage](https://deepwiki.com/charmbracelet/glow/1.2-installation-and-usage)
    
*   [Architecture](https://deepwiki.com/charmbracelet/glow/2-architecture)
    
*   [Command-Line Interface](https://deepwiki.com/charmbracelet/glow/2.1-command-line-interface)
    
*   [Terminal User Interface](https://deepwiki.com/charmbracelet/glow/2.2-terminal-user-interface)
    
*   [Stash View](https://deepwiki.com/charmbracelet/glow/2.2.1-stash-view)
    
*   [Document View](https://deepwiki.com/charmbracelet/glow/2.2.2-document-view)
    
*   [Markdown Processing](https://deepwiki.com/charmbracelet/glow/2.3-markdown-processing)
    
*   [UI Components and Styling](https://deepwiki.com/charmbracelet/glow/2.4-ui-components-and-styling)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/glow/3-configuration-system)
    
*   [Developer Guide](https://deepwiki.com/charmbracelet/glow/4-developer-guide)
    
*   [Build and Release Process](https://deepwiki.com/charmbracelet/glow/4.1-build-and-release-process)
    
*   [Code Quality and Testing](https://deepwiki.com/charmbracelet/glow/4.2-code-quality-and-testing)
    

Menu

Architecture
============

Relevant source files

*   [console\_windows.go](https://github.com/charmbracelet/glow/blob/752de97c/console_windows.go)
    
*   [go.mod](https://github.com/charmbracelet/glow/blob/752de97c/go.mod)
    
*   [main.go](https://github.com/charmbracelet/glow/blob/752de97c/main.go)
    
*   [ui/ui.go](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go)
    

This document provides a comprehensive overview of Glow's architectural design, focusing on the separation between CLI and TUI operational modes, the state management system, and how major components interact. The architecture is built around a dual-mode design that supports both direct terminal output for scripting workflows and an interactive file browsing experience.

For detailed information about CLI execution flow, see [Command-Line Interface](https://deepwiki.com/charmbracelet/glow/2.1-command-line-interface)
. For TUI implementation details, see [Terminal User Interface](https://deepwiki.com/charmbracelet/glow/2.2-terminal-user-interface)
. For content processing specifics, see [Markdown Processing](https://deepwiki.com/charmbracelet/glow/2.3-markdown-processing)
.

Dual-Mode Architecture
----------------------

Glow operates in two distinct modes, each with different initialization paths and execution flows. The routing decision happens early in [main.go222-261](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L222-L261)
 within the `execute` function.

**CLI Mode** renders markdown directly to the terminal or pager, ideal for:

*   Piped content workflows (`cat file.md | glow`)
*   Quick file rendering (`glow README.md`)
*   Integration with scripts and automation
*   Output redirection and processing

**TUI Mode** provides an interactive file browser and document viewer, designed for:

*   Exploring directories of markdown files
*   Live editing with file watching
*   Interactive navigation and search
*   Sustained reading sessions

### Mode Selection Logic

**Sources:** [main.go222-261](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L222-L261)
 [main.go344-369](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L344-L369)

Entry Point and Command Framework
---------------------------------

The application uses Cobra for command-line parsing and Viper for configuration management. The root command structure is defined in [main.go46-64](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L46-L64)

| Component | Purpose | Location |
| --- | --- | --- |
| `rootCmd` | Main command definition with argument validation | [main.go46-64](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L46-L64) |
| `validateOptions()` | Pre-run validation and configuration loading | [main.go165-209](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L165-L209) |
| `execute()` | Primary execution router | [main.go222-261](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L222-L261) |
| `configCmd` | Configuration file editor command | Referenced at [main.go423](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L423-L423) |
| `manCmd` | Man page generation command | Referenced at [main.go423](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L423-L423) |

### Configuration Hierarchy

Configuration is loaded through a three-tier system, with CLI flags having the highest precedence:

1.  **CLI Flags** (highest priority) - parsed by Cobra
2.  **Environment Variables** - with `GLOW_` prefix via Viper
3.  **Configuration File** (`glow.yml`) - loaded from standard locations

The configuration loading happens in [main.go426-468](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L426-L468)
 within `tryLoadConfigFromDefaultPlaces()`, which searches multiple directories in order:

*   `$GLOW_CONFIG_HOME`
*   `$XDG_CONFIG_HOME/glow`
*   Platform-specific user config directories via `go-app-paths`

**Sources:** [main.go46-64](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L46-L64)
 [main.go165-209](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L165-L209)
 [main.go384-424](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L384-L424)
 [main.go426-468](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L426-L468)

TUI System Architecture
-----------------------

The TUI is built on the Bubbletea framework, which implements the Model-View-Update (MVU) pattern. The architecture uses a hierarchical state machine with delegation.

### Model Hierarchy

**Sources:** [ui/ui.go78-112](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L78-L112)

### MVU Pattern Implementation

The core MVU functions are defined in the `model` type:

| Function | Purpose | Location |
| --- | --- | --- |
| `newModel()` | Initializes model with configuration and determines initial state | [ui/ui.go133-184](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L133-L184) |
| `Init()` | Returns initial commands, starts file search or loads document | [ui/ui.go186-203](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L186-L203) |
| `Update()` | Processes messages and updates model state | [ui/ui.go205-325](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L205-L325) |
| `View()` | Renders current view based on state | [ui/ui.go327-338](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L327-L338) |

The `Update()` function implements a dispatch pattern based on `m.state`, delegating messages to either `stashModel.update()` or `pagerModel.update()` as shown in [ui/ui.go312-322](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L312-L322)

### State Transitions

**Sources:** [ui/ui.go78-90](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L78-L90)
 [ui/ui.go116-131](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L116-L131)
 [ui/ui.go205-325](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L205-L325)

Message Flow and Commands
-------------------------

Glow uses Bubbletea's message-passing system for asynchronous operations and state updates. Key message types are defined in [ui/ui.go51-75](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L51-L75)

### Core Message Types

**Sources:** [ui/ui.go51-75](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L51-L75)

### Command Pattern

Commands are functions that return `tea.Cmd`, which is a function returning `tea.Msg`. Key commands in [ui/ui.go357-419](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L357-L419)
:

| Command | Purpose | Location |
| --- | --- | --- |
| `findLocalFiles()` | Initializes file discovery using gitcha | [ui/ui.go357-398](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L357-L398) |
| `findNextLocalFile()` | Reads next result from discovery channel | [ui/ui.go400-412](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L400-L412) |
| `waitForStatusMessageTimeout()` | Schedules status message timeout | [ui/ui.go414-419](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L414-L419) |

Commands execute asynchronously and send messages back to the update loop. For example, `findLocalFiles()` spawns a goroutine via gitcha and returns a message with a channel, then `findNextLocalFile()` reads from that channel in a loop.

**Sources:** [ui/ui.go357-419](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L357-L419)

Content Processing Pipeline
---------------------------

The content pipeline transforms markdown from various sources into rendered ANSI output. The pipeline has two paths depending on operational mode.

### CLI Mode Pipeline

**Sources:** [main.go72-149](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L72-L149)
 [main.go273-342](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L273-L342)

### TUI Mode Pipeline

In TUI mode, content flows through the model system:

1.  **Initial Load** - `newModel()` determines if starting with stash or document view [ui/ui.go155-183](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L155-L183)
    
2.  **File Selection** - User selects file in stash, triggers `fetchedMarkdownMsg`
3.  **Content Load** - `os.ReadFile()` reads content, frontmatter removed [ui/ui.go275-279](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L275-L279)
    
4.  **Rendering** - `renderWithGlamour()` command processes markdown asynchronously
5.  **Display** - `contentRenderedMsg` updates pager viewport

**Sources:** [ui/ui.go133-184](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L133-L184)
 [ui/ui.go186-203](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L186-L203)
 [ui/ui.go275-282](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L275-L282)

Dependency Integration
----------------------

Glow relies heavily on the Charmbracelet ecosystem and supporting libraries defined in [go.mod7-32](https://github.com/charmbracelet/glow/blob/752de97c/go.mod#L7-L32)

### Core Framework Dependencies

**Sources:** [go.mod7-32](https://github.com/charmbracelet/glow/blob/752de97c/go.mod#L7-L32)

### Key Integration Points

| Dependency | Integration Point | Purpose |
| --- | --- | --- |
| `bubbletea` | [ui/ui.go11](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L11-L11) | Provides `tea.Program`, `tea.Model`, `tea.Cmd`, `tea.Msg` types |
| `cobra` | [main.go24](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L24-L24)<br> [main.go46-64](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L46-L64) | Command definition, flag parsing, subcommands |
| `viper` | [main.go25](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L25-L25)<br> [main.go165-209](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L165-L209) | Configuration loading and flag binding |
| `glamour` | [main.go17](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L17-L17)<br> [main.go292-298](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L292-L298) | Markdown to ANSI rendering with style support |
| `gitcha` | [ui/ui.go15](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L15-L15)<br> [ui/ui.go357-398](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L357-L398) | Asynchronous file discovery respecting `.gitignore` |
| `fsnotify` | Used in pager model | Live file watching for auto-reload |

**Sources:** [go.mod7-32](https://github.com/charmbracelet/glow/blob/752de97c/go.mod#L7-L32)
 [main.go1-27](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L1-L27)
 [ui/ui.go1-17](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L1-L17)

Platform-Specific Considerations
--------------------------------

Glow includes platform-specific initialization for Windows ANSI support:

The [console\_windows.go1-24](https://github.com/charmbracelet/glow/blob/752de97c/console_windows.go#L1-L24)
 file enables ANSI color sequences on Windows 10+ by calling `windows.SetConsoleMode()` with the `ENABLE_VIRTUAL_TERMINAL_PROCESSING` flag. This runs automatically via an `init()` function, ensuring color rendering works correctly before the main application starts.

**Sources:** [console\_windows.go1-24](https://github.com/charmbracelet/glow/blob/752de97c/console_windows.go#L1-L24)

Error Handling
--------------

Error handling follows different patterns in CLI vs TUI mode:

**CLI Mode** - Errors are returned directly from functions and propagated up to `main()`, which exits with status 1 on error [main.go377-382](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L377-L382)

**TUI Mode** - Errors are handled through the message system:

*   Fatal errors set `model.fatalErr` and display an error view [ui/ui.go103](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L103-L103)
     [ui/ui.go206-211](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L206-L211)
    
*   Recoverable errors send `errMsg` messages [ui/ui.go51-53](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L51-L53)
    
*   The `errorView()` function renders errors with styling [ui/ui.go340-353](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L340-L353)
    

**Sources:** [main.go371-382](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L371-L382)
 [ui/ui.go51-53](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L51-L53)
 [ui/ui.go103](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L103-L103)
 [ui/ui.go206-211](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L206-L211)
 [ui/ui.go327-353](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L327-L353)

Performance Optimizations
-------------------------

The architecture includes several performance optimizations:

1.  **Asynchronous File Discovery** - Files are discovered in a background goroutine via gitcha channels [ui/ui.go357-398](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L357-L398)
     [ui/ui.go400-412](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L400-L412)
    
2.  **Lazy Loading** - Documents are only loaded and rendered when accessed, not during initial file scan
3.  **High-Performance Pager Mode** - Optional viewport rendering optimization for large documents via `Config.HighPerformancePager` [ui/ui.go36-38](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L36-L38)
    
4.  **Filter Debouncing** - Filtering updates are triggered conditionally via `shouldUpdateFilter()` to avoid excessive re-rendering [ui/ui.go298-300](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L298-L300)
    

**Sources:** [ui/ui.go33-49](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L33-L49)
 [ui/ui.go357-412](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L357-L412)
 [ui/ui.go298-300](https://github.com/charmbracelet/glow/blob/752de97c/ui/ui.go#L298-L300)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Architecture](https://deepwiki.com/charmbracelet/glow/2-architecture#architecture)
    
*   [Dual-Mode Architecture](https://deepwiki.com/charmbracelet/glow/2-architecture#dual-mode-architecture)
    
*   [Mode Selection Logic](https://deepwiki.com/charmbracelet/glow/2-architecture#mode-selection-logic)
    
*   [Entry Point and Command Framework](https://deepwiki.com/charmbracelet/glow/2-architecture#entry-point-and-command-framework)
    
*   [Configuration Hierarchy](https://deepwiki.com/charmbracelet/glow/2-architecture#configuration-hierarchy)
    
*   [TUI System Architecture](https://deepwiki.com/charmbracelet/glow/2-architecture#tui-system-architecture)
    
*   [Model Hierarchy](https://deepwiki.com/charmbracelet/glow/2-architecture#model-hierarchy)
    
*   [MVU Pattern Implementation](https://deepwiki.com/charmbracelet/glow/2-architecture#mvu-pattern-implementation)
    
*   [State Transitions](https://deepwiki.com/charmbracelet/glow/2-architecture#state-transitions)
    
*   [Message Flow and Commands](https://deepwiki.com/charmbracelet/glow/2-architecture#message-flow-and-commands)
    
*   [Core Message Types](https://deepwiki.com/charmbracelet/glow/2-architecture#core-message-types)
    
*   [Command Pattern](https://deepwiki.com/charmbracelet/glow/2-architecture#command-pattern)
    
*   [Content Processing Pipeline](https://deepwiki.com/charmbracelet/glow/2-architecture#content-processing-pipeline)
    
*   [CLI Mode Pipeline](https://deepwiki.com/charmbracelet/glow/2-architecture#cli-mode-pipeline)
    
*   [TUI Mode Pipeline](https://deepwiki.com/charmbracelet/glow/2-architecture#tui-mode-pipeline)
    
*   [Dependency Integration](https://deepwiki.com/charmbracelet/glow/2-architecture#dependency-integration)
    
*   [Core Framework Dependencies](https://deepwiki.com/charmbracelet/glow/2-architecture#core-framework-dependencies)
    
*   [Key Integration Points](https://deepwiki.com/charmbracelet/glow/2-architecture#key-integration-points)
    
*   [Platform-Specific Considerations](https://deepwiki.com/charmbracelet/glow/2-architecture#platform-specific-considerations)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/glow/2-architecture#error-handling)
    
*   [Performance Optimizations](https://deepwiki.com/charmbracelet/glow/2-architecture#performance-optimizations)
