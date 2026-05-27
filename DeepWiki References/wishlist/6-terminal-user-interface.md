Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/wishlist](https://github.com/charmbracelet/wishlist "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 23 April 2025 ([5fe34a](https://github.com/charmbracelet/wishlist/commits/5fe34a29)
)

*   [Introduction to Wishlist](https://deepwiki.com/charmbracelet/wishlist/1-introduction-to-wishlist)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/wishlist/1.1-getting-started)
    
*   [Architecture Overview](https://deepwiki.com/charmbracelet/wishlist/2-architecture-overview)
    
*   [Operation Modes](https://deepwiki.com/charmbracelet/wishlist/2.1-operation-modes)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/wishlist/3-configuration-system)
    
*   [SSH Config Integration](https://deepwiki.com/charmbracelet/wishlist/3.1-ssh-config-integration)
    
*   [Endpoint Configuration](https://deepwiki.com/charmbracelet/wishlist/3.2-endpoint-configuration)
    
*   [SSH Client Implementation](https://deepwiki.com/charmbracelet/wishlist/4-ssh-client-implementation)
    
*   [Authentication Methods](https://deepwiki.com/charmbracelet/wishlist/4.1-authentication-methods)
    
*   [Proxy Jump Support](https://deepwiki.com/charmbracelet/wishlist/4.2-proxy-jump-support)
    
*   [SSH Server Implementation](https://deepwiki.com/charmbracelet/wishlist/5-ssh-server-implementation)
    
*   [Middleware System](https://deepwiki.com/charmbracelet/wishlist/5.1-middleware-system)
    
*   [Terminal User Interface](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface)
    
*   [Endpoint Discovery](https://deepwiki.com/charmbracelet/wishlist/7-endpoint-discovery)
    
*   [Utility Components](https://deepwiki.com/charmbracelet/wishlist/8-utility-components)
    
*   [Deployment](https://deepwiki.com/charmbracelet/wishlist/9-deployment)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/wishlist/10-development-guide)
    

Menu

Terminal User Interface
=======================

Relevant source files

*   [listitem.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/listitem.go)
    
*   [listitem\_test.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/listitem_test.go)
    
*   [styles\_test.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/styles_test.go)
    
*   [wishlist.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go)
    
*   [wishlist\_test.go](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist_test.go)
    

This document describes the Terminal User Interface (TUI) implementation in the Wishlist project. The TUI provides a list-based interface for browsing and connecting to SSH endpoints. It's built on the Charm Bracelet suite of libraries, particularly the Bubble Tea framework for terminal applications.

Overview
--------

The Wishlist TUI displays a filterable, navigable list of SSH endpoints with their descriptions and other metadata. Users can select an endpoint from the list to establish an SSH connection, or perform other actions like copying an endpoint's IP address.

Sources: [wishlist.go49-58](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L49-L58)
 [wishlist.go27-47](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L27-L47)
 [listitem.go12-16](https://github.com/charmbracelet/wishlist/blob/5fe34a29/listitem.go#L12-L16)

Core Components
---------------

### ListModel

The `ListModel` is the central component of the TUI system, responsible for:

1.  Maintaining the list of endpoints
2.  Handling user interactions
3.  Executing SSH commands when an endpoint is selected
4.  Managing the UI state and error handling

Sources: [wishlist.go49-58](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L49-L58)
 [wishlist.go115-231](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L115-L231)

### ItemWrapper

The `ItemWrapper` struct implements the `list.Item` interface from Bubble Tea's list component. It's responsible for formatting endpoint data for display in the list:

*   `Title()`: Returns the endpoint's name
*   `FilterValue()`: Provides the value used for filtering (the endpoint name)
*   `Description()`: Generates a formatted description using descriptor functions

Sources: [listitem.go12-31](https://github.com/charmbracelet/wishlist/blob/5fe34a29/listitem.go#L12-L31)

### Descriptor Functions

Descriptor functions determine how different aspects of an endpoint are formatted for display:

1.  `withDescription`: Displays the endpoint's description or "no description"
2.  `withLink`: Shows the endpoint's link URL and optional name, or "no link"
3.  `withSSHURL`: Formats the endpoint's address as an SSH URL

The system automatically determines which descriptors to use based on the available data in the endpoints.

Sources: [listitem.go33-51](https://github.com/charmbracelet/wishlist/blob/5fe34a29/listitem.go#L33-L51)
 [wishlist.go71-97](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L71-L97)

User Interaction Flow
---------------------

The TUI implements the Model-View-Update (MVU) pattern from the Bubble Tea framework:

Sources: [wishlist.go128-183](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L128-L183)

### Key Bindings

The TUI defines the following key bindings:

| Key | Action | Description |
| --- | --- | --- |
| `y` | Copy IP Address | Copies the endpoint's IP to clipboard |
| `enter`/`o` | Connect | Connects to the selected endpoint |
| `q` | Quit | Exits the TUI |
| `/` | Filter | Activates list filtering |
| Arrow keys | Navigation | Navigate through the list |

These bindings are defined as key.Binding objects and integrated with the list's help system.

Sources: [wishlist.go16-25](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L16-L25)
 [wishlist.go32-37](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L32-L37)

Rendering System
----------------

The TUI's rendering is handled through several key components:

Sources: [wishlist.go60-69](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L60-L69)
 [wishlist.go71-97](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L71-L97)
 [wishlist.go99-112](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L99-L112)
 [wishlist.go198-231](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L198-L231)

### List Item Generation

The process for generating list items from endpoints involves:

1.  Determining which features (descriptions, links) are available in the endpoints
2.  Selecting the appropriate descriptor functions based on those features
3.  Creating `ItemWrapper` instances for each valid endpoint
4.  Setting these wrappers as items in the Bubble Tea list

The list delegate's height is automatically adjusted based on the number of descriptor lines needed.

Sources: [wishlist.go60-112](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L60-L112)

### Error Handling Display

When errors occur (e.g., SSH connection failures), the TUI displays:

1.  An error header
2.  The root cause of the error (extracted via `rootCause()`)
3.  A footer instructing the user how to return to the list

This provides clear feedback when operations fail without crashing the application.

Sources: [wishlist.go198-231](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L198-L231)
 [wishlist.go222-231](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L222-L231)

Integration with SSH Client
---------------------------

The TUI integrates with the SSH client through the `SSHClient` interface:

When a user selects an endpoint and presses Enter, the TUI:

1.  Gets the selected `ItemWrapper` via the `selected()` method
2.  Uses the `SSHClient` interface to create a command for the selected endpoint
3.  Executes this command using Tea's `tea.Exec` function
4.  Handles any errors that occur during the connection process

The connection itself is handled by the SSH client implementation, which is outside the scope of the TUI component.

Sources: [wishlist.go128-160](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L128-L160)
 [wishlist.go185-196](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L185-L196)

Usage in Different Operation Modes
----------------------------------

The Wishlist TUI can be used in two different operational contexts:

1.  **Local Mode**: Running as a standalone application on the user's machine
2.  **Server Mode**: Running as an SSH server that clients can connect to

In both cases, the same TUI code is used, but with different SSH client implementations provided to the `NewListing` function.

For more information about these operation modes, see [Operation Modes](https://deepwiki.com/charmbracelet/wishlist/2.1-operation-modes)
.

Sources: [wishlist.go27-29](https://github.com/charmbracelet/wishlist/blob/5fe34a29/wishlist.go#L27-L29)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Terminal User Interface](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#terminal-user-interface)
    
*   [Overview](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#overview)
    
*   [Core Components](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#core-components)
    
*   [ListModel](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#listmodel)
    
*   [ItemWrapper](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#itemwrapper)
    
*   [Descriptor Functions](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#descriptor-functions)
    
*   [User Interaction Flow](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#user-interaction-flow)
    
*   [Key Bindings](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#key-bindings)
    
*   [Rendering System](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#rendering-system)
    
*   [List Item Generation](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#list-item-generation)
    
*   [Error Handling Display](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#error-handling-display)
    
*   [Integration with SSH Client](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#integration-with-ssh-client)
    
*   [Usage in Different Operation Modes](https://deepwiki.com/charmbracelet/wishlist/6-terminal-user-interface#usage-in-different-operation-modes)
