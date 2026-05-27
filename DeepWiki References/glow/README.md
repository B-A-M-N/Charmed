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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/glow/blob/752de97c/README.md?plain=1)
    
*   [main.go](https://github.com/charmbracelet/glow/blob/752de97c/main.go)
    

Purpose and Scope
-----------------

This page provides a high-level introduction to Glow, a terminal-based markdown renderer. It covers the application's core purpose, its dual-mode architecture (CLI and TUI), and how the major components interact. For detailed information about specific subsystems, see:

*   Dependency details: [Dependencies](https://deepwiki.com/charmbracelet/glow/1.1-dependencies)
    
*   Installation and usage patterns: [Installation and Usage](https://deepwiki.com/charmbracelet/glow/1.2-installation-and-usage)
    
*   In-depth architecture: [Architecture](https://deepwiki.com/charmbracelet/glow/2-architecture)
    
*   Configuration system: [Configuration System](https://deepwiki.com/charmbracelet/glow/3-configuration-system)
    

What is Glow
------------

Glow is a terminal markdown reader that operates in two distinct modes:

1.  **CLI Mode**: Direct rendering of markdown content to stdout, pipes, or a pager (e.g., `less`)
2.  **TUI Mode**: Interactive file browser and document viewer built with the Bubble Tea framework

The application accepts markdown from multiple sources including local files, standard input, HTTP URLs, and GitHub/GitLab repositories. Content is processed through the [Glamour](https://github.com/charmbracelet/glow/blob/752de97c/Glamour)
 rendering engine, which converts markdown to ANSI-formatted terminal output.

**Sources:** [README.md17-25](https://github.com/charmbracelet/glow/blob/752de97c/README.md?plain=1#L17-L25)
 [README.md123-151](https://github.com/charmbracelet/glow/blob/752de97c/README.md?plain=1#L123-L151)

Application Entry Point
-----------------------

Glow's execution begins in [main.go371-382](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L371-L382)
 with the `main()` function, which delegates to Cobra's command framework via `rootCmd.Execute()`. The `rootCmd` variable ([main.go46-63](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L46-L63)
) defines the CLI structure:

    glow [SOURCE|DIR]
    

The command accepts at most one positional argument ([main.go55](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L55-L55)
) and routes execution through two key functions:

*   `validateOptions()` [main.go165-209](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L165-L209)
    : Loads configuration from Viper and validates settings
*   `execute()` [main.go222-261](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L222-L261)
    : Determines whether to run in CLI or TUI mode

**Sources:** [main.go46-64](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L46-L64)
 [main.go371-382](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L371-L382)

Dual-Mode Architecture
----------------------

**Execution Flow Decision Logic**

The `execute()` function ([main.go222-261](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L222-L261)
) implements the following decision tree:

1.  **Piped stdin detected** ([main.go225-231](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L225-L231)
    ): Route to CLI mode via `executeCLI()`
2.  **Zero arguments** ([main.go235-236](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L235-L236)
    ): Launch TUI in current working directory
3.  **One argument that is a directory** ([main.go239-248](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L239-L248)
    ): Launch TUI in that directory
4.  **One or more file/URL arguments** ([main.go252-257](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L252-L257)
    ): Process each via CLI mode

**Sources:** [main.go222-261](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L222-L261)
 [main.go344-369](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L344-L369)
 [main.go273-342](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L273-L342)

Core Components and Data Flow
-----------------------------

**Component Responsibilities**

| Component | File Location | Purpose |
| --- | --- | --- |
| `sourceFromArg()` | [main.go72-149](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L72-L149) | Resolves argument to readable source (stdin/file/URL/GitHub) |
| `executeCLI()` | [main.go273-342](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L273-L342) | CLI rendering pipeline: read → process → render → output |
| `runTUI()` | [main.go344-369](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L344-L369) | TUI initialization: configure → launch Bubble Tea program |
| `glamour.NewTermRenderer()` | [main.go292-298](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L292-L298) | Creates Glamour renderer with style, width, and color profile |
| `ui.NewProgram()` | [main.go364](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L364-L364) | Initializes Bubble Tea TUI with stash and pager models |
| `utils.RemoveFrontmatter()` | [main.go279](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L279-L279) | Strips YAML/TOML frontmatter from markdown content |

**Sources:** [main.go72-149](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L72-L149)
 [main.go273-342](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L273-L342)
 [main.go344-369](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L344-L369)
 [main.go292-298](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L292-L298)

Configuration System
--------------------

Glow uses a three-tier configuration hierarchy managed by Viper:

1.  **CLI flags** (highest priority): Parsed by Cobra, bound to Viper ([main.go408-417](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L408-L417)
    )
2.  **Environment variables** (middle priority): Prefixed with `GLOW_` ([main.go448-449](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L448-L449)
    )
3.  **Configuration file** (lowest priority): `glow.yml` loaded from standard config paths ([main.go426-468](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L426-L468)
    )

The `validateOptions()` function ([main.go165-209](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L165-L209)
) retrieves values from Viper and validates them:

Key validation logic:

*   **Style validation** ([main.go153-163](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L153-L163)
    ): Checks if style is built-in or custom JSON file exists
*   **Width detection** ([main.go193-207](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L193-L207)
    ): Auto-detects terminal width via `term.GetSize()`, defaults to 80
*   **Terminal detection** ([main.go185-190](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L185-L190)
    ): Switches to `notty` style for non-TTY output

Configuration flows into both execution modes:

*   **CLI mode**: Settings applied directly to `glamour.NewTermRenderer()` ([main.go292-298](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L292-L298)
    )
*   **TUI mode**: Packed into `ui.Config` struct and passed to `ui.NewProgram()` ([main.go346-364](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L346-L364)
    )

For detailed configuration options and the config file format, see [Configuration System](https://deepwiki.com/charmbracelet/glow/3-configuration-system)
.

**Sources:** [main.go165-209](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L165-L209)
 [main.go426-468](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L426-L468)
 [main.go408-417](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L408-L417)
 [README.md191-216](https://github.com/charmbracelet/glow/blob/752de97c/README.md?plain=1#L191-L216)

Source Resolution Pipeline
--------------------------

The `sourceFromArg()` function ([main.go72-149](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L72-L149)
) implements a waterfall resolution strategy to normalize various input types into a unified `source` struct:

**Resolution Priority Order:**

1.  **Stdin** ([main.go75-77](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L75-L77)
    ): Argument `"-"` maps to `os.Stdin`
2.  **GitHub/GitLab URLs** ([main.go79-84](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L79-L84)
    ): Delegates to `readmeURL()` for README resolution
3.  **HTTP(S) URLs** ([main.go87-102](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L87-L102)
    ): Uses `http.Get()` with protocol validation
4.  **Directories** ([main.go104-138](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L104-L138)
    ): Walks directory tree looking for README files (case-insensitive)
5.  **Local files** ([main.go140-148](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L140-L148)
    ): Opens file with `os.Open()`

The `readmeNames` variable ([main.go35](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L35-L35)
) defines the search pattern for READMEs:

This resolution strategy allows users to specify markdown sources in multiple formats:

*   `glow -` (stdin)
*   `glow github.com/owner/repo` (GitHub README)
*   `glow https://example.com/file.md` (HTTP)
*   `glow /path/to/file.md` (local file)
*   `glow /path/to/dir` (directory with README)

For details on GitHub/GitLab API integration, see [Markdown Processing](https://deepwiki.com/charmbracelet/glow/2.3-markdown-processing)
.

**Sources:** [main.go72-149](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L72-L149)
 [main.go35](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L35-L35)
 [README.md139-151](https://github.com/charmbracelet/glow/blob/752de97c/README.md?plain=1#L139-L151)

Key Dependencies
----------------

Glow's architecture is built on the Charmbracelet ecosystem:

| Dependency | Purpose | Usage in Glow |
| --- | --- | --- |
| **bubbletea** | MVU framework for TUI | Powers entire interactive mode via `ui.NewProgram()` |
| **glamour** | Markdown renderer | Converts markdown to ANSI output in both CLI and TUI modes |
| **lipgloss** | Terminal styling | Provides color profile detection and style definitions |
| **bubbles** | Reusable TUI components | Viewport component used in pager for scrolling |
| **cobra** | CLI framework | Parses commands and flags via `rootCmd` |
| **viper** | Configuration management | Loads and merges config from files, env vars, and flags |

The dependency imports are visible in [main.go4-27](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L4-L27)
 For a complete dependency breakdown including version constraints, see [Dependencies](https://deepwiki.com/charmbracelet/glow/1.1-dependencies)
.

**Sources:** [main.go4-27](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L4-L27)
 [README.md17-25](https://github.com/charmbracelet/glow/blob/752de97c/README.md?plain=1#L17-L25)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/glow#overview)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/glow#purpose-and-scope)
    
*   [What is Glow](https://deepwiki.com/charmbracelet/glow#what-is-glow)
    
*   [Application Entry Point](https://deepwiki.com/charmbracelet/glow#application-entry-point)
    
*   [Dual-Mode Architecture](https://deepwiki.com/charmbracelet/glow#dual-mode-architecture)
    
*   [Core Components and Data Flow](https://deepwiki.com/charmbracelet/glow#core-components-and-data-flow)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/glow#configuration-system)
    
*   [Source Resolution Pipeline](https://deepwiki.com/charmbracelet/glow#source-resolution-pipeline)
    
*   [Key Dependencies](https://deepwiki.com/charmbracelet/glow#key-dependencies)
