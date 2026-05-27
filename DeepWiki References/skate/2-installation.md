Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/skate](https://github.com/charmbracelet/skate "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 21 April 2025 ([098604](https://github.com/charmbracelet/skate/commits/09860417)
)

*   [Overview](https://deepwiki.com/charmbracelet/skate/1-overview)
    
*   [Installation](https://deepwiki.com/charmbracelet/skate/2-installation)
    
*   [User Guide](https://deepwiki.com/charmbracelet/skate/3-user-guide)
    
*   [Basic Commands](https://deepwiki.com/charmbracelet/skate/3.1-basic-commands)
    
*   [Working with Multiple Databases](https://deepwiki.com/charmbracelet/skate/3.2-working-with-multiple-databases)
    
*   [Advanced Features](https://deepwiki.com/charmbracelet/skate/3.3-advanced-features)
    
*   [Architecture](https://deepwiki.com/charmbracelet/skate/4-architecture)
    
*   [Command Processing](https://deepwiki.com/charmbracelet/skate/4.1-command-processing)
    
*   [Data Storage](https://deepwiki.com/charmbracelet/skate/4.2-data-storage)
    
*   [Terminal Interface](https://deepwiki.com/charmbracelet/skate/4.3-terminal-interface)
    
*   [Development](https://deepwiki.com/charmbracelet/skate/5-development)
    
*   [Build Process](https://deepwiki.com/charmbracelet/skate/5.1-build-process)
    
*   [Release Process](https://deepwiki.com/charmbracelet/skate/5.2-release-process)
    
*   [Code Quality](https://deepwiki.com/charmbracelet/skate/5.3-code-quality)
    
*   [Dependency Management](https://deepwiki.com/charmbracelet/skate/5.4-dependency-management)
    
*   [Reference](https://deepwiki.com/charmbracelet/skate/6-reference)
    
*   [Command Line Interface](https://deepwiki.com/charmbracelet/skate/6.1-command-line-interface)
    
*   [API Functions](https://deepwiki.com/charmbracelet/skate/6.2-api-functions)
    
*   [License](https://deepwiki.com/charmbracelet/skate/7-license)
    

Menu

Installation
============

Relevant source files

*   [.goreleaser.yml](https://github.com/charmbracelet/skate/blob/09860417/.goreleaser.yml)
    
*   [README.md](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1)
    
*   [go.mod](https://github.com/charmbracelet/skate/blob/09860417/go.mod)
    

This page provides comprehensive instructions for installing Skate, a personal key-value store, on various operating systems and through different installation methods. For information about using Skate after installation, see the [User Guide](https://deepwiki.com/charmbracelet/skate/3-user-guide)
.

Prerequisites
-------------

Before installing Skate, ensure you have the following:

*   For package manager installations: The appropriate package manager for your system
*   For Go installation: Go 1.23.0 or later
*   Administrative/sudo privileges (for system-wide installations)

Sources: [go.mod1-5](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L1-L5)

Installation Methods Overview
-----------------------------

Skate can be installed using several methods depending on your operating system and preferences.

Sources: [README.md52-91](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L52-L91)

Package Manager Installation
----------------------------

### macOS and Linux (via Homebrew)

### Arch Linux

### NixOS/Nix

### Debian/Ubuntu

### Fedora/RHEL

Sources: [README.md54-79](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L54-L79)

Installing from Pre-built Binaries
----------------------------------

Skate provides pre-built binaries for various platforms. You can download these from the GitHub releases page.

### Steps for Binary Installation

1.  Visit the [Skate releases page](https://github.com/charmbracelet/skate/blob/09860417/Skate%20releases%20page)
     on GitHub
2.  Download the appropriate package for your system:
    *   `.deb` for Debian-based systems
    *   `.rpm` for RHEL-based systems
    *   OS-specific binary for direct use
3.  For package files (`.deb`, `.rpm`), install using your system's package manager
4.  For binary files:
    *   Extract the file (if compressed)
    *   Move to a directory in your PATH (e.g., `/usr/local/bin` on Unix-like systems)
    *   Make executable with `chmod +x /path/to/skate`

Sources: [README.md82-86](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L82-L86)
 [.goreleaser.yml1-16](https://github.com/charmbracelet/skate/blob/09860417/.goreleaser.yml#L1-L16)

Building from Source with Go
----------------------------

You can install Skate directly using Go:

This will download the source code, build it, and install the binary in your Go bin directory (usually `~/go/bin/`).

### Dependencies

When building from source, Skate has the following key dependencies:

| Dependency | Purpose |
| --- | --- |
| github.com/dgraph-io/badger/v4 | Underlying key-value database |
| github.com/spf13/cobra | Command-line interface framework |
| github.com/charmbracelet/lipgloss | Terminal styling and formatting |
| github.com/muesli/go-app-paths | Application paths management |
| github.com/agnivade/levenshtein | Fuzzy matching for suggestions |

Sources: [go.mod7-16](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L7-L16)
 [README.md87-91](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L87-L91)

Installation Verification
-------------------------

After installing Skate, verify the installation by running:

This should display the installed version of Skate.

Installation Architecture
-------------------------

The following diagram illustrates how Skate is installed and its relationship with system components:

Sources: [go.mod7-16](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L7-L16)

Platform-specific Considerations
--------------------------------

### Windows

For Windows users, the recommended installation method is to download the pre-built binary from the GitHub releases page. After downloading, add the binary location to your PATH environment variable to make it available from the command line.

### Docker

Skate can also be run in a Docker container, which is useful for testing or for environments where installation is not preferred.

Note that data stored in a container will be lost when the container is removed unless volumes are used to persist the data.

Troubleshooting Installation Issues
-----------------------------------

If you encounter issues during installation, consider the following troubleshooting steps:

1.  Ensure you have the required system permissions
2.  Check that your package manager is up to date
3.  Verify that you're using a compatible Go version (1.23.0+) if installing with Go
4.  If using pre-built binaries, ensure you've downloaded the correct version for your platform

For persistent issues, consult the GitHub repository for known issues or to report a new one.

Sources: [go.mod3-5](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L3-L5)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Installation](https://deepwiki.com/charmbracelet/skate/2-installation#installation)
    
*   [Prerequisites](https://deepwiki.com/charmbracelet/skate/2-installation#prerequisites)
    
*   [Installation Methods Overview](https://deepwiki.com/charmbracelet/skate/2-installation#installation-methods-overview)
    
*   [Package Manager Installation](https://deepwiki.com/charmbracelet/skate/2-installation#package-manager-installation)
    
*   [macOS and Linux (via Homebrew)](https://deepwiki.com/charmbracelet/skate/2-installation#macos-and-linux-via-homebrew)
    
*   [Arch Linux](https://deepwiki.com/charmbracelet/skate/2-installation#arch-linux)
    
*   [NixOS/Nix](https://deepwiki.com/charmbracelet/skate/2-installation#nixosnix)
    
*   [Debian/Ubuntu](https://deepwiki.com/charmbracelet/skate/2-installation#debianubuntu)
    
*   [Fedora/RHEL](https://deepwiki.com/charmbracelet/skate/2-installation#fedorarhel)
    
*   [Installing from Pre-built Binaries](https://deepwiki.com/charmbracelet/skate/2-installation#installing-from-pre-built-binaries)
    
*   [Steps for Binary Installation](https://deepwiki.com/charmbracelet/skate/2-installation#steps-for-binary-installation)
    
*   [Building from Source with Go](https://deepwiki.com/charmbracelet/skate/2-installation#building-from-source-with-go)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/skate/2-installation#dependencies)
    
*   [Installation Verification](https://deepwiki.com/charmbracelet/skate/2-installation#installation-verification)
    
*   [Installation Architecture](https://deepwiki.com/charmbracelet/skate/2-installation#installation-architecture)
    
*   [Platform-specific Considerations](https://deepwiki.com/charmbracelet/skate/2-installation#platform-specific-considerations)
    
*   [Windows](https://deepwiki.com/charmbracelet/skate/2-installation#windows)
    
*   [Docker](https://deepwiki.com/charmbracelet/skate/2-installation#docker)
    
*   [Troubleshooting Installation Issues](https://deepwiki.com/charmbracelet/skate/2-installation#troubleshooting-installation-issues)
