Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/x](https://github.com/charmbracelet/x "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 7 February 2026 ([764291](https://github.com/charmbracelet/x/commits/7642919e)
)

*   [Overview](https://deepwiki.com/charmbracelet/x/1-overview)
    
*   [Core Terminal Rendering](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering)
    
*   [cellbuf Package](https://deepwiki.com/charmbracelet/x/2.1-cellbuf-package)
    
*   [Virtual Terminal Emulator (vt)](https://deepwiki.com/charmbracelet/x/2.2-virtual-terminal-emulator-(vt))
    
*   [Input Processing System](https://deepwiki.com/charmbracelet/x/3-input-processing-system)
    
*   [Input Package](https://deepwiki.com/charmbracelet/x/3.1-input-package)
    
*   [Windows Console Input](https://deepwiki.com/charmbracelet/x/3.2-windows-console-input)
    
*   [ANSI Processing and Control](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control)
    
*   [Parser and Text Manipulation](https://deepwiki.com/charmbracelet/x/4.1-parser-and-text-manipulation)
    
*   [Terminal Control Sequences](https://deepwiki.com/charmbracelet/x/4.2-terminal-control-sequences)
    
*   [Styling and Colors](https://deepwiki.com/charmbracelet/x/4.3-styling-and-colors)
    
*   [Platform Abstraction Layer](https://deepwiki.com/charmbracelet/x/5-platform-abstraction-layer)
    
*   [Pseudo-Terminal Interface (xpty)](https://deepwiki.com/charmbracelet/x/5.1-pseudo-terminal-interface-(xpty))
    
*   [Windows ConPTY](https://deepwiki.com/charmbracelet/x/5.2-windows-conpty)
    
*   [Terminal I/O (termios and wcwidth)](https://deepwiki.com/charmbracelet/x/5.3-terminal-io-(termios-and-wcwidth))
    
*   [Utility Packages](https://deepwiki.com/charmbracelet/x/6-utility-packages)
    
*   [SSH Keys and Colors](https://deepwiki.com/charmbracelet/x/6.1-ssh-keys-and-colors)
    
*   [Experimental Packages](https://deepwiki.com/charmbracelet/x/6.2-experimental-packages)
    
*   [Development and Testing](https://deepwiki.com/charmbracelet/x/7-development-and-testing)
    
*   [Examples and Integration](https://deepwiki.com/charmbracelet/x/7.1-examples-and-integration)
    
*   [Testing Framework (teatest)](https://deepwiki.com/charmbracelet/x/7.2-testing-framework-(teatest))
    
*   [Repository Structure and CI/CD](https://deepwiki.com/charmbracelet/x/7.3-repository-structure-and-cicd)
    

Menu

Utility Packages
================

Relevant source files

*   [.github/dependabot.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/dependabot.yml)
    
*   [.github/workflows/gitignore.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/gitignore.yml)
    
*   [.github/workflows/vttest.yml](https://github.com/charmbracelet/x/blob/7642919e/.github/workflows/vttest.yml)
    
*   [.gitignore](https://github.com/charmbracelet/x/blob/7642919e/.gitignore)
    
*   [LICENSE](https://github.com/charmbracelet/x/blob/7642919e/LICENSE)
    
*   [README.md](https://github.com/charmbracelet/x/blob/7642919e/README.md?plain=1)
    
*   [colors/colors.go](https://github.com/charmbracelet/x/blob/7642919e/colors/colors.go)
    
*   [colors/go.mod](https://github.com/charmbracelet/x/blob/7642919e/colors/go.mod)
    
*   [colors/go.sum](https://github.com/charmbracelet/x/blob/7642919e/colors/go.sum)
    
*   [exp/golden/testdata/TestTypes/SliceOfBytes.golden](https://github.com/charmbracelet/x/blob/7642919e/exp/golden/testdata/TestTypes/SliceOfBytes.golden)
    
*   [exp/golden/testdata/TestTypes/String.golden](https://github.com/charmbracelet/x/blob/7642919e/exp/golden/testdata/TestTypes/String.golden)
    
*   [scripts/dependabot](https://github.com/charmbracelet/x/blob/7642919e/scripts/dependabot)
    
*   [sshkey/\_examples/key](https://github.com/charmbracelet/x/blob/7642919e/sshkey/_examples/key)
    
*   [sshkey/\_examples/key.pub](https://github.com/charmbracelet/x/blob/7642919e/sshkey/_examples/key.pub)
    
*   [sshkey/\_examples/main.go](https://github.com/charmbracelet/x/blob/7642919e/sshkey/_examples/main.go)
    
*   [sshkey/go.mod](https://github.com/charmbracelet/x/blob/7642919e/sshkey/go.mod)
    
*   [sshkey/go.sum](https://github.com/charmbracelet/x/blob/7642919e/sshkey/go.sum)
    
*   [sshkey/sshkey.go](https://github.com/charmbracelet/x/blob/7642919e/sshkey/sshkey.go)
    

This document covers the supporting utility packages in the charmbracelet/x ecosystem that provide specialized functionality for terminal applications. These packages handle SSH key operations, color management, and experimental color palettes. For core terminal rendering functionality, see [Core Terminal Rendering](https://deepwiki.com/charmbracelet/x/2-core-terminal-rendering)
. For terminal control and ANSI processing, see [ANSI Processing and Control](https://deepwiki.com/charmbracelet/x/4-ansi-processing-and-control)
.

Package Overview
----------------

The charmbracelet/x repository includes several utility packages that complement the core terminal functionality:

*   **`sshkey`** - SSH key parsing with interactive passphrase prompts
*   **`colors`** - Predefined color constants for consistent styling
*   **`exp/charmtone`** - Experimental CharmTone color palette API

Utility Package Architecture
----------------------------

Sources: [sshkey/sshkey.go1-79](https://github.com/charmbracelet/x/blob/7642919e/sshkey/sshkey.go#L1-L79)
 [colors/colors.go1-37](https://github.com/charmbracelet/x/blob/7642919e/colors/colors.go#L1-L37)
 [exp/charmtone/charmtone.go1-359](https://github.com/charmbracelet/x/blob/7642919e/exp/charmtone/charmtone.go#L1-L359)

SSH Key Utilities
-----------------

The `sshkey` package provides utilities for parsing SSH private keys with interactive passphrase handling.

### Core Functions

| Function | Purpose | Parameters | Return Type |
| --- | --- | --- | --- |
| `Open` | Read and parse key from file path | `keyPath string` | `ssh.Signer, error` |
| `Parse` | Parse PEM bytes into ssh.Signer | `identifier string, pemBytes []byte` | `ssh.Signer, error` |
| `ParseRaw` | Parse PEM bytes into raw private key | `identifier string, pemBytes []byte` | `any, error` |

### SSH Key Processing Flow

Sources: [sshkey/sshkey.go12-79](https://github.com/charmbracelet/x/blob/7642919e/sshkey/sshkey.go#L12-L79)

The `doParse` function is a generic helper that handles both `ssh.Signer` and raw private key parsing through the same passphrase prompt flow [sshkey/sshkey.go37-59](https://github.com/charmbracelet/x/blob/7642919e/sshkey/sshkey.go#L37-L59)
 The `ask` function uses the `huh` library to create an interactive password input prompt [sshkey/sshkey.go66-78](https://github.com/charmbracelet/x/blob/7642919e/sshkey/sshkey.go#L66-L78)

Color Definitions
-----------------

The `colors` package provides predefined color constants using `lipgloss.AdaptiveColor` for light/dark theme compatibility.

### Color Categories

| Category | Colors | Type |
| --- | --- | --- |
| **Basic** | `Normal`, `NormalDim` | `lipgloss.AdaptiveColor` |
| **Gray Scale** | `Gray`, `GrayMid`, `GrayDark`, `GrayBright`, `GrayBrightDim` | `lipgloss.AdaptiveColor` |
| **Indigo** | `Indigo`, `IndigoDim`, `IndigoSubtle`, `IndigoSubtleDim` | `lipgloss.AdaptiveColor` |
| **Green** | `YellowGreen`, `YellowGreenDull`, `Green`, `GreenDim` | Mixed types |
| **Pink/Red** | `Fuschia`, `FuchsiaDim`, `FuchsiaDull`, `Red`, `RedDull` | `lipgloss.AdaptiveColor` |

### Adaptive Color Structure

Sources: [colors/colors.go6-36](https://github.com/charmbracelet/x/blob/7642919e/colors/colors.go#L6-L36)

Most colors are defined as `lipgloss.AdaptiveColor` with separate light and dark variants [colors/colors.go7-35](https://github.com/charmbracelet/x/blob/7642919e/colors/colors.go#L7-L35)
 while some like `Green` use a fixed `lipgloss.Color` [colors/colors.go31](https://github.com/charmbracelet/x/blob/7642919e/colors/colors.go#L31-L31)

CharmTone Color Palette
-----------------------

The `exp/charmtone` package provides an experimental API for the CharmTone color palette with 48 named colors and color blending capabilities.

### Key Type and Constants

The `Key` type is an integer-based enumeration representing color keys [exp/charmtone/charmtone.go14-76](https://github.com/charmbracelet/x/blob/7642919e/exp/charmtone/charmtone.go#L14-L76)
:

Sources: [exp/charmtone/charmtone.go14-287](https://github.com/charmbracelet/x/blob/7642919e/exp/charmtone/charmtone.go#L14-L287)
 [exp/charmtone/charmtone.go310-358](https://github.com/charmbracelet/x/blob/7642919e/exp/charmtone/charmtone.go#L310-L358)

### Color Classification

The palette includes classification methods for organizing colors into primary, secondary, and tertiary groups:

| Method | Colors | Count |
| --- | --- | --- |
| `IsPrimary()` | `Charple`, `Dolly`, `Julep`, `Zest`, `Butter` | 5   |
| `IsSecondary()` | `Hazy`, `Blush`, `Bok` | 3   |
| `IsTertiary()` | `Turtle`, `Malibu`, `Violet`, `Tuna`, `Coral`, `Uni` | 6   |

### Color Blending

The `BlendColors` function creates smooth color transitions between multiple color stops using HCL color space blending [exp/charmtone/charmtone.go310-358](https://github.com/charmbracelet/x/blob/7642919e/exp/charmtone/charmtone.go#L310-L358)
:

Sources: [exp/charmtone/charmtone.go310-358](https://github.com/charmbracelet/x/blob/7642919e/exp/charmtone/charmtone.go#L310-L358)

The function distributes colors evenly across segments, handling remainder distribution to ensure the exact requested number of colors [exp/charmtone/charmtone.go328-339](https://github.com/charmbracelet/x/blob/7642919e/exp/charmtone/charmtone.go#L328-L339)

Integration with Core Systems
-----------------------------

These utility packages integrate with the broader charmbracelet/x ecosystem through well-defined interfaces and dependencies. The `sshkey` package uses the `huh` TUI library for user interaction, `colors` provides styling constants for `lipgloss` integration, and `charmtone` offers advanced color manipulation for sophisticated terminal applications.

Sources: [sshkey/go.mod5-8](https://github.com/charmbracelet/x/blob/7642919e/sshkey/go.mod#L5-L8)
 [colors/go.mod5](https://github.com/charmbracelet/x/blob/7642919e/colors/go.mod#L5-L5)
 [exp/charmtone/go.mod5](https://github.com/charmbracelet/x/blob/7642919e/exp/charmtone/go.mod#L5-L5)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Utility Packages](https://deepwiki.com/charmbracelet/x/6-utility-packages#utility-packages)
    
*   [Package Overview](https://deepwiki.com/charmbracelet/x/6-utility-packages#package-overview)
    
*   [Utility Package Architecture](https://deepwiki.com/charmbracelet/x/6-utility-packages#utility-package-architecture)
    
*   [SSH Key Utilities](https://deepwiki.com/charmbracelet/x/6-utility-packages#ssh-key-utilities)
    
*   [Core Functions](https://deepwiki.com/charmbracelet/x/6-utility-packages#core-functions)
    
*   [SSH Key Processing Flow](https://deepwiki.com/charmbracelet/x/6-utility-packages#ssh-key-processing-flow)
    
*   [Color Definitions](https://deepwiki.com/charmbracelet/x/6-utility-packages#color-definitions)
    
*   [Color Categories](https://deepwiki.com/charmbracelet/x/6-utility-packages#color-categories)
    
*   [Adaptive Color Structure](https://deepwiki.com/charmbracelet/x/6-utility-packages#adaptive-color-structure)
    
*   [CharmTone Color Palette](https://deepwiki.com/charmbracelet/x/6-utility-packages#charmtone-color-palette)
    
*   [Key Type and Constants](https://deepwiki.com/charmbracelet/x/6-utility-packages#key-type-and-constants)
    
*   [Color Classification](https://deepwiki.com/charmbracelet/x/6-utility-packages#color-classification)
    
*   [Color Blending](https://deepwiki.com/charmbracelet/x/6-utility-packages#color-blending)
    
*   [Integration with Core Systems](https://deepwiki.com/charmbracelet/x/6-utility-packages#integration-with-core-systems)
