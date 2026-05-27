Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/charm](https://github.com/charmbracelet/charm "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 21 April 2025 ([aff307](https://github.com/charmbracelet/charm/commits/aff3071d)
)

*   [Overview](https://deepwiki.com/charmbracelet/charm/1-overview)
    
*   [Core Concepts](https://deepwiki.com/charmbracelet/charm/1.1-core-concepts)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/charm/1.2-architecture-overview)
    
*   [Charm Server](https://deepwiki.com/charmbracelet/charm/2-charm-server)
    
*   [HTTP Server](https://deepwiki.com/charmbracelet/charm/2.1-http-server)
    
*   [SSH Server](https://deepwiki.com/charmbracelet/charm/2.2-ssh-server)
    
*   [Server Configuration](https://deepwiki.com/charmbracelet/charm/2.3-server-configuration)
    
*   [Database and Storage](https://deepwiki.com/charmbracelet/charm/2.4-database-and-storage)
    
*   [Metrics and Monitoring](https://deepwiki.com/charmbracelet/charm/2.5-metrics-and-monitoring)
    
*   [Charm Client](https://deepwiki.com/charmbracelet/charm/3-charm-client)
    
*   [Client Configuration](https://deepwiki.com/charmbracelet/charm/3.1-client-configuration)
    
*   [Authentication](https://deepwiki.com/charmbracelet/charm/3.2-authentication)
    
*   [Key Management](https://deepwiki.com/charmbracelet/charm/3.3-key-management)
    
*   [Terminal UI](https://deepwiki.com/charmbracelet/charm/4-terminal-ui)
    
*   [UI Components](https://deepwiki.com/charmbracelet/charm/4.1-ui-components)
    
*   [Key Management UI](https://deepwiki.com/charmbracelet/charm/4.2-key-management-ui)
    
*   [User Profile UI](https://deepwiki.com/charmbracelet/charm/4.3-user-profile-ui)
    
*   [Data Storage](https://deepwiki.com/charmbracelet/charm/5-data-storage)
    
*   [File System (FS)](https://deepwiki.com/charmbracelet/charm/5.1-file-system-(fs))
    
*   [Key-Value Store (KV)](https://deepwiki.com/charmbracelet/charm/5.2-key-value-store-(kv))
    
*   [Encryption](https://deepwiki.com/charmbracelet/charm/5.3-encryption)
    
*   [Deployment](https://deepwiki.com/charmbracelet/charm/6-deployment)
    
*   [Self-Hosting Guide](https://deepwiki.com/charmbracelet/charm/6.1-self-hosting-guide)
    
*   [Docker Deployment](https://deepwiki.com/charmbracelet/charm/6.2-docker-deployment)
    
*   [Backup and Restore](https://deepwiki.com/charmbracelet/charm/6.3-backup-and-restore)
    
*   [Development](https://deepwiki.com/charmbracelet/charm/7-development)
    
*   [Building and Testing](https://deepwiki.com/charmbracelet/charm/7.1-building-and-testing)
    
*   [CI/CD Workflows](https://deepwiki.com/charmbracelet/charm/7.2-cicd-workflows)
    
*   [Command-Line Interface](https://deepwiki.com/charmbracelet/charm/7.3-command-line-interface)
    

Menu

Terminal UI
===========

Relevant source files

*   [ui/common/common.go](https://github.com/charmbracelet/charm/blob/aff3071d/ui/common/common.go)
    
*   [ui/common/views.go](https://github.com/charmbracelet/charm/blob/aff3071d/ui/common/views.go)
    
*   [ui/keys/keyview.go](https://github.com/charmbracelet/charm/blob/aff3071d/ui/keys/keyview.go)
    
*   [ui/ui.go](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go)
    

The Terminal UI (TUI) system in Charm provides a rich interactive interface that runs directly in the user's terminal. It serves as the primary graphical interface for account management, SSH key operations, and other Charm features. This document explains the architecture, components, and functionality of the TUI system.

For specific UI components implementation details, see [UI Components](https://deepwiki.com/charmbracelet/charm/4.1-ui-components)
. For information about the Key Management interface, see [Key Management UI](https://deepwiki.com/charmbracelet/charm/4.2-key-management-ui)
.

Overview
--------

The Charm Terminal UI is built using the [Bubble Tea](https://github.com/charmbracelet/charm/blob/aff3071d/Bubble%20Tea)
 framework, which implements the Model-View-Update (MVU) pattern for terminal applications. The TUI provides users with a menu-driven interface for:

*   Managing SSH keys
*   Linking multiple devices to a Charm account
*   Setting/changing username
*   Viewing account information
*   Managing backups

Sources: [ui/ui.go1-32](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L1-L32)
 [ui/ui.go90-117](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L90-L117)

Architecture
------------

The Terminal UI uses a hierarchical component structure where a root model manages the overall application state and sub-models handle specific features.

Sources: [ui/ui.go34-65](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L34-L65)
 [ui/ui.go68-87](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L68-L87)
 [ui/ui.go90-117](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L90-L117)

### State Management

The Terminal UI follows the Bubble Tea implementation of the Model-View-Update (MVU) pattern with three main functions:

1.  **Init**: Initializes the model and returns commands to be executed
2.  **Update**: Processes messages and updates the model state
3.  **View**: Renders the current state as a string to be displayed in the terminal

State transitions are handled by the `Update` function, which processes various message types including:

*   Window size changes
*   Keyboard input
*   Client events
*   Child component messages

Sources: [ui/ui.go119-223](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L119-L223)
 [ui/ui.go225-330](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L225-L330)

Core Components
---------------

### Root Model

The root model (`model` struct) serves as the main container for all UI state and sub-components. It tracks:

*   Application status (e.g., initializing, ready, linking)
*   Menu selection
*   Error states
*   Client connection
*   Child UI models

    type model struct {
        cfg        *client.Config
        cc         *client.Client
        styles     common.Styles
        user       *charm.User
        err        error
        status     status
        menuIndex  int
        menuChoice menuChoice
        
        spinner  spinner.Model
        info     info.Model
        linkgen  linkgen.Model
        username username.Model
        keys     keys.Model
        
        terminalWidth int
    }
    

Sources: [ui/ui.go90-107](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L90-L107)

### Status States

The UI uses a status type to track its high-level state:

| Status Value | Description |
| --- | --- |
| `statusInit` | Initial application state |
| `statusKeygen` | Generating SSH keys |
| `statusKeygenComplete` | Key generation finished |
| `statusFetching` | Fetching data from server |
| `statusReady` | Ready for user interaction |
| `statusLinking` | Linking a new device |
| `statusBrowsingKeys` | Viewing/managing SSH keys |
| `statusSettingUsername` | Setting username |
| `statusShowBackupInfo` | Showing backup information |
| `statusQuitting` | Application is shutting down |
| `statusError` | Error state |

Sources: [ui/ui.go34-65](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L34-L65)

### Menu System

The TUI provides a main menu with the following options:

*   Link a machine
*   Manage linked keys
*   Set Username
*   Backup
*   Exit

The menu is rendered by the `menuView` function and responds to keyboard navigation (up/down arrows, j/k keys) and selection (enter key).

Sources: [ui/ui.go68-87](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L68-L87)
 [ui/ui.go365-383](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L365-L383)

UI Flow and Navigation
----------------------

Sources: [ui/ui.go225-330](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L225-L330)
 [ui/ui.go126-175](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L126-L175)

View Rendering
--------------

The main `View` function renders different content based on the current UI status. When in the `statusReady` state, it shows:

1.  The Charm logo
2.  User information (from `info.View()`)
3.  The menu options
4.  Footer with help text

In other states, it renders the appropriate UI component for that state:

*   In `statusLinking`, it renders the link generator UI
*   In `statusBrowsingKeys`, it renders the key browser UI
*   In `statusSettingUsername`, it renders the username UI
*   In `statusShowBackupInfo`, it renders backup information

Sources: [ui/ui.go332-423](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L332-L423)

Common UI Components
--------------------

The TUI uses common UI components from the `common` package to provide consistent visual elements across different screens.

### UI States

The UI defines several states that affect styling:

*   `StateNormal`: Default state
*   `StateSelected`: Element is selected
*   `StateActive`: Element is active
*   `StateSpecial`: Special highlighting
*   `StateDeleting`: Element is being deleted

Sources: [ui/common/views.go12-21](https://github.com/charmbracelet/charm/blob/aff3071d/ui/common/views.go#L12-L21)

### Helper Components

Common UI elements include:

*   **Vertical Line**: Used for visual separation with state-appropriate colors
*   **Spinner**: Animated indicator for loading states
*   **Key-Value View**: Consistently formatted key-value pairs
*   **Help View**: Standardized help text, typically shown at the bottom
*   **Button Views**: Various button styles (OK, Cancel, Yes, No)

Sources: [ui/common/views.go30-182](https://github.com/charmbracelet/charm/blob/aff3071d/ui/common/views.go#L30-L182)

Integration with Charm Client
-----------------------------

The Terminal UI integrates with the Charm client to perform operations like:

*   Fetching user information
*   Managing SSH keys
*   Setting username
*   Generating linking tokens

Communication with the client happens through the `charmclient` package and message types like:

*   `charmclient.NewClientMsg`
*   `charmclient.ErrMsg`
*   `charmclient.SSHAuthErrorMsg`

Sources: [ui/ui.go176-215](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L176-L215)

Customization
-------------

The TUI uses the `lipgloss` styling library for consistent visual presentation. Styles are defined in the `common.Styles` struct and initialized with defaults in the `initialModel` function.

Sources: [ui/ui.go111-116](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L111-L116)

Summary
-------

The Charm Terminal UI provides a rich, interactive experience for managing Charm accounts directly in the terminal. It uses the Bubble Tea framework to implement a clean, component-based architecture that follows the Model-View-Update pattern. The UI is divided into specialized components for different functions while sharing common UI elements and styling.

The system handles various states and transitions between different screens based on user input, providing a seamless experience for managing SSH keys, linking devices, and configuring account settings.

Sources: [ui/ui.go1-32](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L1-L32)
 [ui/ui.go90-107](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L90-L107)
 [ui/ui.go119-223](https://github.com/charmbracelet/charm/blob/aff3071d/ui/ui.go#L119-L223)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Terminal UI](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#terminal-ui)
    
*   [Overview](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#overview)
    
*   [Architecture](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#architecture)
    
*   [State Management](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#state-management)
    
*   [Core Components](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#core-components)
    
*   [Root Model](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#root-model)
    
*   [Status States](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#status-states)
    
*   [Menu System](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#menu-system)
    
*   [UI Flow and Navigation](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#ui-flow-and-navigation)
    
*   [View Rendering](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#view-rendering)
    
*   [Common UI Components](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#common-ui-components)
    
*   [UI States](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#ui-states)
    
*   [Helper Components](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#helper-components)
    
*   [Integration with Charm Client](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#integration-with-charm-client)
    
*   [Customization](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#customization)
    
*   [Summary](https://deepwiki.com/charmbracelet/charm/4-terminal-ui#summary)
