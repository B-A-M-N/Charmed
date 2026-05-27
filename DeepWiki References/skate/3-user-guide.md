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

User Guide
==========

Relevant source files

*   [README.md](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1)
    
*   [main.go](https://github.com/charmbracelet/skate/blob/09860417/main.go)
    

This guide provides comprehensive information on how to use Skate, a personal key-value store application. You'll learn how to get started with Skate, understand its core concepts, and master its features for managing your data effectively. For detailed information about specific command usage, see [Basic Commands](https://deepwiki.com/charmbracelet/skate/3.1-basic-commands)
. For information about working with multiple databases, see [Working with Multiple Databases](https://deepwiki.com/charmbracelet/skate/3.2-working-with-multiple-databases)
.

Core Concepts
-------------

Skate allows you to store and retrieve key-value pairs through a simple command-line interface. Each key-value pair consists of:

*   **Key**: A unique identifier for your data (can contain Unicode characters)
*   **Value**: The data associated with the key (can be text or binary)
*   **Database**: An optional namespace to organize your key-value pairs (prefixed with `@`)

Sources: [main.go38-109](https://github.com/charmbracelet/skate/blob/09860417/main.go#L38-L109)
 [README.md17-50](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L17-L50)

Command Structure
-----------------

Skate uses a simple verb-based command structure. Most commands follow this pattern:

    skate <command> [key[@database]] [value]
    

Where:

*   `command` is the operation (set, get, delete, list, etc.)
*   `key` is your data identifier
*   `@database` is an optional database specifier
*   `value` is the data to store (for set commands)

### Core Commands

| Command | Purpose | Example |
| --- | --- | --- |
| `set` | Store a key-value pair | `skate set name value` |
| `get` | Retrieve a value by key | `skate get name` |
| `delete` | Remove a key-value pair | `skate delete name` |
| `list` | Show all key-value pairs | `skate list` |
| `list-dbs` | Show all databases | `skate list-dbs` |
| `delete-db` | Delete a database | `skate delete-db @dbname` |

Sources: [main.go39-92](https://github.com/charmbracelet/skate/blob/09860417/main.go#L39-L92)
 [README.md17-50](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L17-L50)

Data Flow and Storage
---------------------

When you use Skate, your data flows through the system as follows:

Sources: [main.go123-145](https://github.com/charmbracelet/skate/blob/09860417/main.go#L123-L145)
 [main.go409-418](https://github.com/charmbracelet/skate/blob/09860417/main.go#L409-L418)
 [main.go225-236](https://github.com/charmbracelet/skate/blob/09860417/main.go#L225-L236)

Database Organization
---------------------

Skate organizes data into databases that are stored locally on your system:

Sources: [main.go225-236](https://github.com/charmbracelet/skate/blob/09860417/main.go#L225-L236)
 [main.go409-418](https://github.com/charmbracelet/skate/blob/09860417/main.go#L409-L418)

Working with Keys and Values
----------------------------

### Key Format

Keys in Skate:

*   Are converted to lowercase automatically
*   Can include Unicode characters
*   Can be any text string
*   Are separated from the database with `@`

The key parsing is handled by the `keyParser()` function which separates the key from any database specification:

Sources: [main.go394-407](https://github.com/charmbracelet/skate/blob/09860417/main.go#L394-L407)

### Value Handling

Skate can store:

*   Plain text values
*   Binary data
*   Unicode text

When retrieving or listing binary data, Skate will omit it by default in terminal output unless the `--show-binary` flag is used.

Sources: [main.go377-392](https://github.com/charmbracelet/skate/blob/09860417/main.go#L377-L392)
 [README.md33-36](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L33-L36)

Common Usage Patterns
---------------------

### Storing and Retrieving Text Data

    skate set mykey "This is my value"
    skate get mykey
    

### Working with Binary Data

    # Store binary data from a file
    skate set myimage < image.jpg
    
    # Retrieve binary data to a file
    skate get myimage > retrieved_image.jpg
    

### Using Multiple Databases

    # Store in a custom database
    skate set secretkey@passwords mysecretvalue
    
    # List all databases
    skate list-dbs
    
    # List all entries in a specific database
    skate list @passwords
    

### Filtering and Formatting List Output

    # List only keys
    skate list -k
    
    # List only values
    skate list -v
    
    # Reverse order listing
    skate list -r
    
    # Custom delimiter
    skate list -d " = "
    

Sources: [README.md98-114](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L98-L114)
 [README.md118-131](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L118-L131)

Storage Location
----------------

Skate stores all data in the user's data directory:

| Platform | Path |
| --- | --- |
| macOS | `~/Library/Application Support/charm/kv/` |
| Linux | `~/.local/share/charm/kv/` |
| Windows | `%AppData%\charm\kv\` |

Inside this directory, each database is stored as a separate folder containing BadgerDB files.

Sources: [main.go225-236](https://github.com/charmbracelet/skate/blob/09860417/main.go#L225-L236)

Integration with Scripts
------------------------

Skate is designed to be easily integrated with shell scripts and other tools:

Sources: [README.md137-161](https://github.com/charmbracelet/skate/blob/09860417/README.md?plain=1#L137-L161)

Command Line Flags
------------------

Many Skate commands support additional flags for customizing behavior:

| Command | Flag | Description |
| --- | --- | --- |
| `list` | `-r, --reverse` | List in reverse lexicographic order |
| `list` | `-k, --keys-only` | Only print keys |
| `list` | `-v, --values-only` | Only print values |
| `list` | `-d, --delimiter` | Set custom delimiter between keys and values |
| `list` | `-b, --show-binary` | Show binary values instead of omitting them |
| `get` | `-b, --show-binary` | Show binary values instead of omitting them |

Sources: [main.go431-436](https://github.com/charmbracelet/skate/blob/09860417/main.go#L431-L436)

For detailed information about specific commands and their options, see [Basic Commands](https://deepwiki.com/charmbracelet/skate/3.1-basic-commands)
. For advanced usage scenarios, see [Advanced Features](https://deepwiki.com/charmbracelet/skate/3.3-advanced-features)
.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [User Guide](https://deepwiki.com/charmbracelet/skate/3-user-guide#user-guide)
    
*   [Core Concepts](https://deepwiki.com/charmbracelet/skate/3-user-guide#core-concepts)
    
*   [Command Structure](https://deepwiki.com/charmbracelet/skate/3-user-guide#command-structure)
    
*   [Core Commands](https://deepwiki.com/charmbracelet/skate/3-user-guide#core-commands)
    
*   [Data Flow and Storage](https://deepwiki.com/charmbracelet/skate/3-user-guide#data-flow-and-storage)
    
*   [Database Organization](https://deepwiki.com/charmbracelet/skate/3-user-guide#database-organization)
    
*   [Working with Keys and Values](https://deepwiki.com/charmbracelet/skate/3-user-guide#working-with-keys-and-values)
    
*   [Key Format](https://deepwiki.com/charmbracelet/skate/3-user-guide#key-format)
    
*   [Value Handling](https://deepwiki.com/charmbracelet/skate/3-user-guide#value-handling)
    
*   [Common Usage Patterns](https://deepwiki.com/charmbracelet/skate/3-user-guide#common-usage-patterns)
    
*   [Storing and Retrieving Text Data](https://deepwiki.com/charmbracelet/skate/3-user-guide#storing-and-retrieving-text-data)
    
*   [Working with Binary Data](https://deepwiki.com/charmbracelet/skate/3-user-guide#working-with-binary-data)
    
*   [Using Multiple Databases](https://deepwiki.com/charmbracelet/skate/3-user-guide#using-multiple-databases)
    
*   [Filtering and Formatting List Output](https://deepwiki.com/charmbracelet/skate/3-user-guide#filtering-and-formatting-list-output)
    
*   [Storage Location](https://deepwiki.com/charmbracelet/skate/3-user-guide#storage-location)
    
*   [Integration with Scripts](https://deepwiki.com/charmbracelet/skate/3-user-guide#integration-with-scripts)
    
*   [Command Line Flags](https://deepwiki.com/charmbracelet/skate/3-user-guide#command-line-flags)
