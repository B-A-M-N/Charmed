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

Overview
========

Relevant source files

*   [README.md](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1)
    
*   [go.mod](https://github.com/charmbracelet/skate/blob/09860417/go.mod)
    
*   [main.go](https://github.com/charmbracelet/skate/blob/09860417/main.go)
    

This document provides a comprehensive overview of Skate, a personal key-value store application with a command-line interface. Skate allows users to store and retrieve arbitrary data, including text and binary content, through a simple CLI experience.

What is Skate?
--------------

Skate is a Go-based command-line application that functions as a personal key-value store. It enables users to save, retrieve, and manage arbitrary data through a straightforward interface. Skate stores data locally using BadgerDB as its underlying storage engine.

Sources: [main.go1-22](https://github.com/charmbracelet/skate/blob/09860417/main.go#L1-L22)
 [README.md8-19](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L8-L19)

Core Features
-------------

Skate provides several key capabilities:

1.  **Simple Key-Value Operations**: Store, retrieve, and delete data using intuitive commands
2.  **Multiple Database Support**: Organize data across separate databases with the `@DB` syntax
3.  **Binary Data Support**: Handle both text and binary data seamlessly
4.  **Unicode Support**: Full compatibility with Unicode characters
5.  **Terminal-Friendly Output**: Formatted output for terminal display with binary data handling

Sources: [README.md17-43](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L17-L43)
 [main.go24-46](https://github.com/charmbracelet/skate/blob/09860417/main.go#L24-L46)

System Architecture
-------------------

Skate follows a clean, modular architecture with distinct layers handling different concerns:

Sources: [main.go24-110](https://github.com/charmbracelet/skate/blob/09860417/main.go#L24-L110)
 [main.go393-463](https://github.com/charmbracelet/skate/blob/09860417/main.go#L393-L463)

### Command Layer

The command layer is implemented using the Cobra library, defining the CLI commands and their arguments. Each command is registered with the root command and has its own run function.

| Command | Description | Function |
| --- | --- | --- |
| set | Set a value for a key with optional @DB | set() |
| get | Get a value for a key with optional @DB | get() |
| delete | Delete a key with optional @DB | del() |
| list | List key-value pairs with optional @DB | list() |
| list-dbs | List available databases | listDbs() |
| delete-db | Delete a database | deleteDb() |

Sources: [main.go38-109](https://github.com/charmbracelet/skate/blob/09860417/main.go#L38-L109)
 [main.go437-446](https://github.com/charmbracelet/skate/blob/09860417/main.go#L437-L446)

### Core Data Flow

The following diagram illustrates how data flows through the system when executing commands:

Sources: [main.go123-190](https://github.com/charmbracelet/skate/blob/09860417/main.go#L123-L190)
 [main.go393-407](https://github.com/charmbracelet/skate/blob/09860417/main.go#L393-L407)

Storage Organization
--------------------

Skate organizes data into separate databases stored in the user's application data directory. Each database is a separate directory containing BadgerDB files.

Sources: [main.go225-236](https://github.com/charmbracelet/skate/blob/09860417/main.go#L225-L236)
 [main.go409-418](https://github.com/charmbracelet/skate/blob/09860417/main.go#L409-L418)

### Database Management

Databases are created automatically when used for the first time. The system provides commands to list available databases and delete databases:

*   `list-dbs`: Shows all available databases
*   `delete-db @DB`: Deletes a specified database and all its contents

When a database doesn't exist, Skate uses the Levenshtein algorithm to suggest similar database names, improving user experience.

Sources: [main.go189-195](https://github.com/charmbracelet/skate/blob/09860417/main.go#L189-L195)
 [main.go238-305](https://github.com/charmbracelet/skate/blob/09860417/main.go#L238-L305)

Key and Value Processing
------------------------

Skate parses keys using the following format:

    key[@database]
    

If no database is specified, the "default" database is used. Keys and values are processed as byte arrays, allowing both text and binary data to be stored.

Sources: [main.go393-407](https://github.com/charmbracelet/skate/blob/09860417/main.go#L393-L407)
 [main.go409-418](https://github.com/charmbracelet/skate/blob/09860417/main.go#L409-L418)

### Binary Data Handling

Skate can store binary data and handles its display in terminal environments. By default, binary data is detected and marked as "(omitted binary data)" when displayed in a terminal, unless the `--show-binary` flag is used.

Sources: [main.go377-392](https://github.com/charmbracelet/skate/blob/09860417/main.go#L377-L392)
 [main.go435-436](https://github.com/charmbracelet/skate/blob/09860417/main.go#L435-L436)

Dependencies
------------

Skate relies on several key libraries:

*   **BadgerDB** (`github.com/dgraph-io/badger/v4`): Embedded key-value database for storage
*   **Cobra** (`github.com/spf13/cobra`): Command-line interface framework
*   **Lipgloss** (`github.com/charmbracelet/lipgloss`): Terminal styling library
*   **go-app-paths** (`github.com/muesli/go-app-paths`): Platform-independent application paths
*   **Levenshtein** (`github.com/agnivade/levenshtein`): String distance calculation for suggestions

Sources: [go.mod7-16](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L7-L16)
 [main.go1-22](https://github.com/charmbracelet/skate/blob/09860417/main.go#L1-L22)

Related Documentation
---------------------

For more detailed information about specific aspects of Skate, refer to:

*   Installation instructions: [Installation](https://deepwiki.com/charmbracelet/skate/2-installation)
    
*   Detailed usage guide: [User Guide](https://deepwiki.com/charmbracelet/skate/3-user-guide)
    
*   Architecture details: [Architecture](https://deepwiki.com/charmbracelet/skate/4-architecture)
    
*   Development information: [Development](https://deepwiki.com/charmbracelet/skate/5-development)
    

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Overview](https://deepwiki.com/charmbracelet/skate/1-overview#overview)
    
*   [What is Skate?](https://deepwiki.com/charmbracelet/skate/1-overview#what-is-skate)
    
*   [Core Features](https://deepwiki.com/charmbracelet/skate/1-overview#core-features)
    
*   [System Architecture](https://deepwiki.com/charmbracelet/skate/1-overview#system-architecture)
    
*   [Command Layer](https://deepwiki.com/charmbracelet/skate/1-overview#command-layer)
    
*   [Core Data Flow](https://deepwiki.com/charmbracelet/skate/1-overview#core-data-flow)
    
*   [Storage Organization](https://deepwiki.com/charmbracelet/skate/1-overview#storage-organization)
    
*   [Database Management](https://deepwiki.com/charmbracelet/skate/1-overview#database-management)
    
*   [Key and Value Processing](https://deepwiki.com/charmbracelet/skate/1-overview#key-and-value-processing)
    
*   [Binary Data Handling](https://deepwiki.com/charmbracelet/skate/1-overview#binary-data-handling)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/skate/1-overview#dependencies)
    
*   [Related Documentation](https://deepwiki.com/charmbracelet/skate/1-overview#related-documentation)
