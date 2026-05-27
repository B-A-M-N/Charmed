Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/sequin](https://github.com/charmbracelet/sequin "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 20 April 2025 ([025641](https://github.com/charmbracelet/sequin/commits/025641c1)
)

*   [Introduction to Sequin](https://deepwiki.com/charmbracelet/sequin/1-introduction-to-sequin)
    
*   [Installation](https://deepwiki.com/charmbracelet/sequin/1.1-installation)
    
*   [Basic Usage](https://deepwiki.com/charmbracelet/sequin/1.2-basic-usage)
    
*   [Architecture](https://deepwiki.com/charmbracelet/sequin/2-architecture)
    
*   [ANSI Sequence Processing](https://deepwiki.com/charmbracelet/sequin/2.1-ansi-sequence-processing)
    
*   [Command Execution](https://deepwiki.com/charmbracelet/sequin/2.2-command-execution)
    
*   [ANSI Sequence Handlers](https://deepwiki.com/charmbracelet/sequin/3-ansi-sequence-handlers)
    
*   [SGR (Select Graphic Rendition)](https://deepwiki.com/charmbracelet/sequin/3.1-sgr-(select-graphic-rendition))
    
*   [Cursor Commands](https://deepwiki.com/charmbracelet/sequin/3.2-cursor-commands)
    
*   [Screen and Line Commands](https://deepwiki.com/charmbracelet/sequin/3.3-screen-and-line-commands)
    
*   [Window Title and Colors](https://deepwiki.com/charmbracelet/sequin/3.4-window-title-and-colors)
    
*   [Terminal Modes](https://deepwiki.com/charmbracelet/sequin/3.5-terminal-modes)
    
*   [Pointer, Clipboard, and Links](https://deepwiki.com/charmbracelet/sequin/3.6-pointer-clipboard-and-links)
    
*   [FinalTerm Commands](https://deepwiki.com/charmbracelet/sequin/3.7-finalterm-commands)
    
*   [Theme System](https://deepwiki.com/charmbracelet/sequin/4-theme-system)
    
*   [Testing](https://deepwiki.com/charmbracelet/sequin/5-testing)
    
*   [Build and Development](https://deepwiki.com/charmbracelet/sequin/6-build-and-development)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/sequin/6.1-dependency-management)
    
*   [Release Process](https://deepwiki.com/charmbracelet/sequin/6.2-release-process)
    

Menu

Theme System
============

Relevant source files

*   [exec.go](https://github.com/charmbracelet/sequin/blob/025641c1/exec.go)
    
*   [main.go](https://github.com/charmbracelet/sequin/blob/025641c1/main.go)
    
*   [testdata/TestSequences/others/pm.golden](https://github.com/charmbracelet/sequin/blob/025641c1/testdata/TestSequences/others/pm.golden)
    
*   [testdata/TestSequences/others/sos.golden](https://github.com/charmbracelet/sequin/blob/025641c1/testdata/TestSequences/others/sos.golden)
    
*   [theme.go](https://github.com/charmbracelet/sequin/blob/025641c1/theme.go)
    

Purpose and Overview
--------------------

The Theme System in Sequin is responsible for styling and formatting the visual representation of ANSI sequences when they are displayed to the user. It provides consistent and readable styling for different sequence types, explanations, and text content. The Theme System ensures that ANSI sequences are presented in a visually distinct and comprehensible manner, helping users quickly identify and understand different elements in the output.

The Theme System works closely with the ANSI Sequence Processor and Handlers to apply appropriate styling to processed sequences. For information about how sequences are processed, see [ANSI Sequence Processing](https://deepwiki.com/charmbracelet/sequin/2.1-ansi-sequence-processing)
.

Sources: [theme.go1-131](https://github.com/charmbracelet/sequin/blob/025641c1/theme.go#L1-L131)
 [main.go68-79](https://github.com/charmbracelet/sequin/blob/025641c1/main.go#L68-L79)

Theme Structure
---------------

### Core Components

The Theme System is built around the `theme` struct, which encapsulates all styling information used throughout Sequin:

Each theme has several style components that serve different purposes:

| Style Component | Purpose |
| --- | --- |
| `raw` | Base style used in raw mode output |
| `kind` | Base style for sequence type labels (CSI, OSC, etc.) |
| `sequence` | Style for the sequence string representation |
| `separator` | Style for the separator between sequence and explanation |
| `text` | Style for regular text content |
| `error` | Style for error messages |
| `explanation` | Style for explanatory text describing what a sequence does |
| `kindColors` | Color definitions for different sequence types |

Sources: [theme.go10-24](https://github.com/charmbracelet/sequin/blob/025641c1/theme.go#L10-L24)

### The kindStyle Method

The `kindStyle` method is a central part of the Theme System. It returns an appropriate style for a given sequence type (kind):

The method:

1.  Determines the base style (raw or kind)
2.  Applies the appropriate color based on the sequence type
3.  In non-raw mode, adds the appropriate label text (CSI, OSC, etc.)

Sources: [theme.go26-71](https://github.com/charmbracelet/sequin/blob/025641c1/theme.go#L26-L71)

Theme Types and Selection
-------------------------

Sequin provides two built-in theme types:

1.  **Charm Theme (`charmTheme`)**: The default theme with a rich color palette that adapts to light/dark terminal backgrounds
2.  **Base16 Theme (`base16Theme`)**: An alternative theme using Base16 color scheme

The theme selection process occurs in the `process` function:

In the `main.go` file, theme selection occurs at the beginning of processing:

Sources: [main.go68-79](https://github.com/charmbracelet/sequin/blob/025641c1/main.go#L68-L79)
 [theme.go73-130](https://github.com/charmbracelet/sequin/blob/025641c1/theme.go#L73-L130)

Theme Application
-----------------

The Theme System is applied throughout the sequence processing pipeline. Here's how different styles are applied in various contexts:

### Raw Mode vs. Normal Mode

The Theme System behaves differently based on whether raw mode is enabled:

| Mode | Behavior |
| --- | --- |
| Normal Mode | Full styling with sequence type labels, separators, and explanations |
| Raw Mode (`--raw` flag) | Minimal styling with just sequence strings, no explanations |

In raw mode, only the sequence is displayed with minimal styling. In normal mode, sequences are displayed with their type label, the sequence itself, and an explanation of what the sequence does.

Sources: [main.go79-78](https://github.com/charmbracelet/sequin/blob/025641c1/main.go#L79-L78)
 [main.go80-185](https://github.com/charmbracelet/sequin/blob/025641c1/main.go#L80-L185)

### Style Application Examples

Here are examples of how the Theme System is applied to different types of sequences:

1.  **CSI Sequences**:
    
        CSI: <FileRef file-url="https://github.com/charmbracelet/sequin/blob/025641c1/31m Foreground color#LNaN-LNaN" NaN  file-path="31m Foreground color">Hii</FileRef> <FileRef file-url="https://github.com/charmbracelet/sequin/blob/025641c1/testdata/TestSequences/others/pm.golden#L1-L2" min=1 max=2 file-path="testdata/TestSequences/others/pm.golden">Hii</FileRef> <FileRef file-url="https://github.com/charmbracelet/sequin/blob/025641c1/testdata/TestSequences/others/sos.golden#L1-L2" min=1 max=2 file-path="testdata/TestSequences/others/sos.golden">Hii</FileRef>
        
    

Theme Implementation Details
----------------------------

### Charm Theme

The Charm theme is the default and adapts to the terminal's background color. The `charmTheme` function:

1.  Creates a function to select colors based on background brightness
2.  Initializes all style components with appropriate colors and formatting
3.  Sets up colors for different sequence types

Key features:

*   Bold labels with right alignment
*   Separator formatted as ": "
*   Different colors for each sequence type
*   Explanations colored for optimal readability

Sources: [theme.go73-108](https://github.com/charmbracelet/sequin/blob/025641c1/theme.go#L73-L108)

### Base16 Theme

The Base16 theme uses a more standard color palette. The `base16Theme` function:

1.  Starts with a Charm theme
2.  Overrides the colors with Base16 colors
3.  Uses simpler ANSI colors (Red, Blue, Green, etc.)

This theme can be activated by setting the `SEQUIN_THEME` environment variable to "ansi", "carlos", "secret\_carlos", or "matchy".

Sources: [theme.go110-130](https://github.com/charmbracelet/sequin/blob/025641c1/theme.go#L110-L130)

Integration with the Processing Pipeline
----------------------------------------

The Theme System is integrated into the main processing pipeline at several points:

1.  **Initialization**: At the start of the `process` function
2.  **Sequence Printing**: In the `seqPrint` function
3.  **Text Printing**: In the `flushPrint` function
4.  **Handler Output**: After sequence handlers run

This integration ensures that all output is consistently styled based on the selected theme.

Sources: [main.go68-79](https://github.com/charmbracelet/sequin/blob/025641c1/main.go#L68-L79)
 [main.go80-185](https://github.com/charmbracelet/sequin/blob/025641c1/main.go#L80-L185)
 [main.go187-202](https://github.com/charmbracelet/sequin/blob/025641c1/main.go#L187-L202)

Customization
-------------

The Theme System can be customized through:

1.  **The `--raw` flag**: Enables raw mode with minimal styling
2.  **The `SEQUIN_THEME` environment variable**: Changes the theme type
    

For developers looking to add new themes, the structure allows for easy extension by:

1.  Creating a new theme function similar to `charmTheme` or `base16Theme`
2.  Adding the theme to the selection logic in the `process` function

Sources: [main.go68-78](https://github.com/charmbracelet/sequin/blob/025641c1/main.go#L68-L78)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Theme System](https://deepwiki.com/charmbracelet/sequin/4-theme-system#theme-system)
    
*   [Purpose and Overview](https://deepwiki.com/charmbracelet/sequin/4-theme-system#purpose-and-overview)
    
*   [Theme Structure](https://deepwiki.com/charmbracelet/sequin/4-theme-system#theme-structure)
    
*   [Core Components](https://deepwiki.com/charmbracelet/sequin/4-theme-system#core-components)
    
*   [The kindStyle Method](https://deepwiki.com/charmbracelet/sequin/4-theme-system#the-kindstyle-method)
    
*   [Theme Types and Selection](https://deepwiki.com/charmbracelet/sequin/4-theme-system#theme-types-and-selection)
    
*   [Theme Application](https://deepwiki.com/charmbracelet/sequin/4-theme-system#theme-application)
    
*   [Raw Mode vs. Normal Mode](https://deepwiki.com/charmbracelet/sequin/4-theme-system#raw-mode-vs-normal-mode)
    
*   [Style Application Examples](https://deepwiki.com/charmbracelet/sequin/4-theme-system#style-application-examples)
    
*   [Theme Implementation Details](https://deepwiki.com/charmbracelet/sequin/4-theme-system#theme-implementation-details)
    
*   [Charm Theme](https://deepwiki.com/charmbracelet/sequin/4-theme-system#charm-theme)
    
*   [Base16 Theme](https://deepwiki.com/charmbracelet/sequin/4-theme-system#base16-theme)
    
*   [Integration with the Processing Pipeline](https://deepwiki.com/charmbracelet/sequin/4-theme-system#integration-with-the-processing-pipeline)
    
*   [Customization](https://deepwiki.com/charmbracelet/sequin/4-theme-system#customization)
