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

Architecture
============

Relevant source files

*   [go.mod](https://github.com/charmbracelet/skate/blob/09860417/go.mod)
    
*   [go.sum](https://github.com/charmbracelet/skate/blob/09860417/go.sum)
    
*   [main.go](https://github.com/charmbracelet/skate/blob/09860417/main.go)
    

This document provides a detailed explanation of Skate's system architecture, including its core components, data flow, and design principles. It covers how Skate processes commands, stores data, and interacts with users through the terminal interface. For information about using Skate, see [User Guide](https://deepwiki.com/charmbracelet/skate/3-user-guide)
, and for development details, see [Development](https://deepwiki.com/charmbracelet/skate/5-development)
.

High-Level Architecture Overview
--------------------------------

Skate follows a layered architecture with clear separation of concerns. The application is structured into four primary layers that work together to provide a seamless key-value store experience.

Sources: [main.go1-464](https://github.com/charmbracelet/skate/blob/09860417/main.go#L1-L464)
 [go.mod1-50](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L1-L50)

### Component Interaction

1.  **CLI Layer**: Receives user input through command-line arguments, parses them using the `keyParser` function, and routes commands to the appropriate handlers.
    
2.  **Application Logic**: Contains the core business logic of the application, managing database operations, key-value manipulation, and error handling.
    
3.  **Storage Layer**: Handles data persistence using BadgerDB, manages database paths, and handles multiple databases.
    
4.  **Terminal UI Layer**: Formats output for display to the user, applying styling through the Lipgloss library.
    

Command Processing
------------------

Skate uses the Cobra framework to define and process commands. All commands are registered with the root command and follow a consistent pattern of parsing input, performing operations, and formatting output.

Sources: [main.go37-109](https://github.com/charmbracelet/skate/blob/09860417/main.go#L37-L109)
 [main.go448-454](https://github.com/charmbracelet/skate/blob/09860417/main.go#L448-L454)

### Key Command Functions

Skate's functionality is implemented through several key command handlers:

*   `set`: Stores a value for a key, accepting input from arguments or stdin
*   `get`: Retrieves a value for a key and displays it
*   `del`: Removes a key-value pair from the database
*   `list`: Displays all key-value pairs in a database with formatting options
*   `listDbs`: Shows all available databases
*   `deleteDb`: Deletes an entire database after confirmation

Commands are processed through transaction wrappers that handle BadgerDB operations safely using the `wrap` function.

Sources: [main.go123-186](https://github.com/charmbracelet/skate/blob/09860417/main.go#L123-L186)
 [main.go307-364](https://github.com/charmbracelet/skate/blob/09860417/main.go#L307-L364)
 [main.go456-463](https://github.com/charmbracelet/skate/blob/09860417/main.go#L456-L463)

Data Storage
------------

### BadgerDB Integration

Skate uses BadgerDB (v4) as its underlying storage engine. BadgerDB is an embedded key-value database that provides efficient storage with good performance characteristics.

Sources: [main.go409-418](https://github.com/charmbracelet/skate/blob/09860417/main.go#L409-L418)
 [main.go224-236](https://github.com/charmbracelet/skate/blob/09860417/main.go#L224-L236)

### Database File Structure

Skate stores its databases in a user-specific location determined by the `getFilePath` function using the `go-app-paths` package. Each database is stored in its own directory under the `charm/kv/` path in the user's data directory. The default database is named "default", and custom databases are named according to user specifications.

The database paths are structured as follows:

*   Base directory: `<user data directory>/charm/kv/`
*   Default database: `<base directory>/default/`
*   Custom database: `<base directory>/<database name>/`

Sources: [main.go224-236](https://github.com/charmbracelet/skate/blob/09860417/main.go#L224-L236)
 [main.go409-418](https://github.com/charmbracelet/skate/blob/09860417/main.go#L409-L418)

### Multi-Database Support

One of Skate's key features is its support for multiple databases. Users can specify which database to use by adding `@DB` to key names. If no database is specified, Skate uses the "default" database.

When a database is referenced that doesn't exist, Skate will create it automatically. If a user makes a typo when referencing a database, Skate uses the Levenshtein distance algorithm to suggest similar database names.

Sources: [main.go394-407](https://github.com/charmbracelet/skate/blob/09860417/main.go#L394-L407)
 [main.go276-305](https://github.com/charmbracelet/skate/blob/09860417/main.go#L276-L305)

Key-Value Operations
--------------------

### Key Parsing and Database Selection

Keys in Skate follow the format `KEY@DB`, where:

*   `KEY` is the identifier for the value
*   `@DB` is an optional database specifier

The `keyParser` function splits this input format and returns both the key (as bytes) and the database name (as a string).

Sources: [main.go394-407](https://github.com/charmbracelet/skate/blob/09860417/main.go#L394-L407)
 [main.go409-418](https://github.com/charmbracelet/skate/blob/09860417/main.go#L409-L418)

### Transaction Handling

All database operations in Skate are performed within transactions managed by the `wrap` function. This function:

1.  Creates a new BadgerDB transaction
2.  Executes the provided function within the transaction
3.  Commits or discards the transaction based on the operation's success

This approach ensures data consistency and proper error handling.

Sources: [main.go456-463](https://github.com/charmbracelet/skate/blob/09860417/main.go#L456-L463)

### Binary and Text Data Support

Skate supports both text and binary data, with Unicode compatibility. When displaying data in the terminal, binary data is handled specially:

*   By default, binary data is not shown directly and is replaced with "(omitted binary data)"
*   Users can use the `--show-binary` flag to display binary data
*   Binary detection is performed by checking if the data is valid UTF-8

The `printFromKV` function manages this formatting and output.

Sources: [main.go377-392](https://github.com/charmbracelet/skate/blob/09860417/main.go#L377-L392)

Terminal Interface
------------------

### Output Formatting

Skate uses the Lipgloss library for terminal styling. This provides:

*   Color output for warnings and important information
*   Width adjustments for readability
*   Consistent formatting across the application

The `printFromKV` function handles specialized formatting of key-value data for display, including special handling for binary data and optional delimiters between keys and values.

Sources: [main.go377-392](https://github.com/charmbracelet/skate/blob/09860417/main.go#L377-L392)
 [main.go37](https://github.com/charmbracelet/skate/blob/09860417/main.go#L37-L37)

### List Display Options

The `list` command provides various formatting options:

*   `--reverse` / `-r`: Display items in reverse lexicographic order
*   `--keys-only` / `-k`: Show only keys
*   `--values-only` / `-v`: Show only values
*   `--delimiter` / `-d`: Customize the separator between keys and values
*   `--show-binary` / `-b`: Display binary data as-is

These options allow for flexible display of database contents to suit different use cases.

Sources: [main.go307-364](https://github.com/charmbracelet/skate/blob/09860417/main.go#L307-L364)
 [main.go431-436](https://github.com/charmbracelet/skate/blob/09860417/main.go#L431-L436)

Error Handling and Suggestions
------------------------------

Skate implements sophisticated error handling, particularly for cases where users might mistype database names. The system uses the Levenshtein distance algorithm to suggest similar database names when a specified database isn't found.

The `errDBNotFound` struct and related functions provide this functionality, enhancing user experience by offering helpful suggestions rather than simply reporting errors.

Sources: [main.go112-121](https://github.com/charmbracelet/skate/blob/09860417/main.go#L112-L121)
 [main.go276-305](https://github.com/charmbracelet/skate/blob/09860417/main.go#L276-L305)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Architecture](https://deepwiki.com/charmbracelet/skate/4-architecture#architecture)
    
*   [High-Level Architecture Overview](https://deepwiki.com/charmbracelet/skate/4-architecture#high-level-architecture-overview)
    
*   [Component Interaction](https://deepwiki.com/charmbracelet/skate/4-architecture#component-interaction)
    
*   [Command Processing](https://deepwiki.com/charmbracelet/skate/4-architecture#command-processing)
    
*   [Key Command Functions](https://deepwiki.com/charmbracelet/skate/4-architecture#key-command-functions)
    
*   [Data Storage](https://deepwiki.com/charmbracelet/skate/4-architecture#data-storage)
    
*   [BadgerDB Integration](https://deepwiki.com/charmbracelet/skate/4-architecture#badgerdb-integration)
    
*   [Database File Structure](https://deepwiki.com/charmbracelet/skate/4-architecture#database-file-structure)
    
*   [Multi-Database Support](https://deepwiki.com/charmbracelet/skate/4-architecture#multi-database-support)
    
*   [Key-Value Operations](https://deepwiki.com/charmbracelet/skate/4-architecture#key-value-operations)
    
*   [Key Parsing and Database Selection](https://deepwiki.com/charmbracelet/skate/4-architecture#key-parsing-and-database-selection)
    
*   [Transaction Handling](https://deepwiki.com/charmbracelet/skate/4-architecture#transaction-handling)
    
*   [Binary and Text Data Support](https://deepwiki.com/charmbracelet/skate/4-architecture#binary-and-text-data-support)
    
*   [Terminal Interface](https://deepwiki.com/charmbracelet/skate/4-architecture#terminal-interface)
    
*   [Output Formatting](https://deepwiki.com/charmbracelet/skate/4-architecture#output-formatting)
    
*   [List Display Options](https://deepwiki.com/charmbracelet/skate/4-architecture#list-display-options)
    
*   [Error Handling and Suggestions](https://deepwiki.com/charmbracelet/skate/4-architecture#error-handling-and-suggestions)
