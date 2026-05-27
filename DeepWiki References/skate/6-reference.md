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

Reference
=========

Relevant source files

*   [go.mod](https://github.com/charmbracelet/skate/blob/09860417/go.mod)
    
*   [main.go](https://github.com/charmbracelet/skate/blob/09860417/main.go)
    

This reference document provides comprehensive technical information about Skate's internals, APIs, and interfaces. It serves as a detailed technical guide for developers who need to understand Skate's implementation or build upon its functionality. For user-oriented documentation on how to use Skate, see [User Guide](https://deepwiki.com/charmbracelet/skate/3-user-guide)
.

Command Line Interface Reference
--------------------------------

Skate's CLI is built using the Cobra framework and provides several commands for interacting with the key-value store.

### Command Structure

Sources: [main.go37-110](https://github.com/charmbracelet/skate/blob/09860417/main.go#L37-L110)
 [main.go430-446](https://github.com/charmbracelet/skate/blob/09860417/main.go#L430-L446)

### Command Reference Table

| Command | Syntax | Description | Examples |
| --- | --- | --- | --- |
| `set` | `set KEY[@DB] [VALUE]` | Set a value for a key with an optional @DB. If VALUE is omitted, reads from stdin. | `skate set foo bar`  <br>`skate set foo@mydb bar`  <br>`skate set foo <./bar.txt` |
| `get` | `get KEY[@DB]` | Retrieve a value for a key with an optional @DB. | `skate get foo`  <br>`skate get foo@mydb` |
| `delete` | `delete KEY[@DB]` | Delete a key with an optional @DB. | `skate delete foo`  <br>`skate delete foo@mydb` |
| `list` | `list [@DB]` | List all key-value pairs in a database. | `skate list`  <br>`skate list @mydb` |
| `list-dbs` | `list-dbs` | List all available databases. | `skate list-dbs` |
| `delete-db` | `delete-db [@DB]` | Delete an entire database. | `skate delete-db @mydb` |
| `man` | `man` | Generate man pages. | `skate man` |

Sources: [main.go48-109](https://github.com/charmbracelet/skate/blob/09860417/main.go#L48-L109)

### Command Flags

#### List Command Flags

| Flag | Short | Default | Description |
| --- | --- | --- | --- |
| `--reverse` | `-r` | `false` | List in reverse lexicographic order |
| `--keys-only` | `-k` | `false` | Only print keys and don't fetch values from the database |
| `--values-only` | `-v` | `false` | Only print values |
| `--delimiter` | `-d` | `\t` | Delimiter to separate keys and values |
| `--show-binary` | `-b` | `false` | Print binary values |

#### Get Command Flags

| Flag | Short | Default | Description |
| --- | --- | --- | --- |
| `--show-binary` | `-b` | `false` | Print binary values |

Sources: [main.go431-436](https://github.com/charmbracelet/skate/blob/09860417/main.go#L431-L436)

Core API Reference
------------------

### Key Functions

Sources: [main.go123-464](https://github.com/charmbracelet/skate/blob/09860417/main.go#L123-L464)

### Core Function Reference Table

| Function | Purpose | Parameters | Return Values |
| --- | --- | --- | --- |
| `keyParser` | Parses a key string to extract the key and optional DB name | `k string` | `[]byte, string, error` (key bytes, db name, error) |
| `openKV` | Opens a BadgerDB database instance | `name string` | `*badger.DB, error` |
| `getFilePath` | Returns the file path for the skate database(s) | `args ...string` | `string, error` (path, error) |
| `wrap` | Utility to wrap BadgerDB transactions | `db *badger.DB, readonly bool, fn func(tx *badger.Txn) error` | `error` |
| `findDb` | Finds a database and provides suggestions if not found | `name string` | `string, error` (path, error) |
| `nameFromArgs` | Extracts database name from command args | `args []string` | `string, error` (db name, error) |
| `printFromKV` | Formats and prints key-value data | `pf string, vs ...[]byte` | none |
| `getDbs` | Gets list of available databases | none | `[]string, error` |

Sources: [main.go198-463](https://github.com/charmbracelet/skate/blob/09860417/main.go#L198-L463)

Data Storage Reference
----------------------

### Storage Structure

Sources: [main.go225-236](https://github.com/charmbracelet/skate/blob/09860417/main.go#L225-L236)

### Storage Path Resolution

Skate uses the `go-app-paths` library to determine where to store databases. The path resolution follows this pattern:

1.  User scope data directory is obtained via `gap.NewScope(gap.User, "charm")`
2.  The path is extended with `kv` to create the main Skate data directory
3.  Each database is stored as a subdirectory within this path

Example (Unix-like systems):

*   Default database: `~/.local/share/charm/kv/default/`
*   Custom database: `~/.local/share/charm/kv/mydb/`

Sources: [main.go225-236](https://github.com/charmbracelet/skate/blob/09860417/main.go#L225-L236)
 [main.go409-418](https://github.com/charmbracelet/skate/blob/09860417/main.go#L409-L418)

### Database Operations Flow

Sources: [main.go123-186](https://github.com/charmbracelet/skate/blob/09860417/main.go#L123-L186)
 [main.go456-463](https://github.com/charmbracelet/skate/blob/09860417/main.go#L456-L463)

Error Handling Reference
------------------------

### Custom Error Types

Sources: [main.go112-121](https://github.com/charmbracelet/skate/blob/09860417/main.go#L112-L121)

### Database Error Handling

Skate implements a fuzzy matching system for database names. When a database is not found, it attempts to find similar database names using the Levenshtein distance algorithm and provides suggestions to the user.

The key components of this system are:

1.  Custom error type `errDBNotFound` which contains suggested database names
2.  `findDb()` function which uses Levenshtein distance to find similar database names
3.  Error handling in commands that display suggestions to users

Sources: [main.go112-121](https://github.com/charmbracelet/skate/blob/09860417/main.go#L112-L121)
 [main.go276-305](https://github.com/charmbracelet/skate/blob/09860417/main.go#L276-L305)

### Error Handling Flow

Sources: [main.go276-305](https://github.com/charmbracelet/skate/blob/09860417/main.go#L276-L305)
 [main.go239-274](https://github.com/charmbracelet/skate/blob/09860417/main.go#L239-L274)

Key-Value Data Handling
-----------------------

### Key Format and Parsing

Keys in Skate can include an optional database specification using the `@` symbol. The format is:

    KEY[@DB]
    

Where:

*   `KEY` is the actual key (automatically converted to lowercase)
*   `@DB` is an optional database specification (automatically converted to lowercase)

The `keyParser()` function handles this parsing:

Sources: [main.go394-407](https://github.com/charmbracelet/skate/blob/09860417/main.go#L394-L407)

### Binary Data Handling

Skate can store and retrieve both text and binary data. When binary data is encountered during output operations:

1.  By default, binary data is displayed as "(omitted binary data)" when output is to a terminal
2.  The `--show-binary` flag forces display of binary data (useful for redirecting to files)
3.  Binary detection uses `utf8.Valid()` to check if data is valid UTF-8

The `printFromKV()` function handles this functionality.

Sources: [main.go377-392](https://github.com/charmbracelet/skate/blob/09860417/main.go#L377-L392)

Database Management
-------------------

### Database Operations

Sources: [main.go189-223](https://github.com/charmbracelet/skate/blob/09860417/main.go#L189-L223)
 [main.go239-305](https://github.com/charmbracelet/skate/blob/09860417/main.go#L239-L305)

### Database Path Structure

Skate databases are organized in the following structure:

| Component | Purpose | Default Location |
| --- | --- | --- |
| User Data Directory | Root directory for user data | OS-specific (e.g., `~/.local/share` on Linux) |
| `charm/` | Charm applications data folder | `<User Data Directory>/charm/` |
| `kv/` | Skate data folder | `<User Data Directory>/charm/kv/` |
| `<db-name>/` | Individual database folder | `<User Data Directory>/charm/kv/<db-name>/` |
| BadgerDB files | BadgerDB storage files | Inside each `<db-name>/` folder |

Sources: [main.go225-236](https://github.com/charmbracelet/skate/blob/09860417/main.go#L225-L236)
 [main.go409-418](https://github.com/charmbracelet/skate/blob/09860417/main.go#L409-L418)

Dependencies
------------

Skate relies on several key dependencies:

### Core Dependencies

| Dependency | Purpose |
| --- | --- |
| `github.com/dgraph-io/badger/v4` | Embedded key-value database engine |
| `github.com/spf13/cobra` | Command-line interface framework |
| `github.com/charmbracelet/lipgloss` | Terminal styling and formatting |
| `github.com/muesli/go-app-paths` | Cross-platform application paths |
| `github.com/agnivade/levenshtein` | Levenshtein distance algorithm for fuzzy matching |
| `github.com/muesli/mango-cobra` | Man page generation for Cobra commands |
| `github.com/muesli/roff` | ROFF document format support |
| `golang.org/x/term` | Terminal utilities |

Sources: [go.mod7-16](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L7-L16)

### Dependency Flow

Sources: [go.mod7-16](https://github.com/charmbracelet/skate/blob/09860417/go.mod#L7-L16)
 [main.go3-22](https://github.com/charmbracelet/skate/blob/09860417/main.go#L3-L22)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Reference](https://deepwiki.com/charmbracelet/skate/6-reference#reference)
    
*   [Command Line Interface Reference](https://deepwiki.com/charmbracelet/skate/6-reference#command-line-interface-reference)
    
*   [Command Structure](https://deepwiki.com/charmbracelet/skate/6-reference#command-structure)
    
*   [Command Reference Table](https://deepwiki.com/charmbracelet/skate/6-reference#command-reference-table)
    
*   [Command Flags](https://deepwiki.com/charmbracelet/skate/6-reference#command-flags)
    
*   [List Command Flags](https://deepwiki.com/charmbracelet/skate/6-reference#list-command-flags)
    
*   [Get Command Flags](https://deepwiki.com/charmbracelet/skate/6-reference#get-command-flags)
    
*   [Core API Reference](https://deepwiki.com/charmbracelet/skate/6-reference#core-api-reference)
    
*   [Key Functions](https://deepwiki.com/charmbracelet/skate/6-reference#key-functions)
    
*   [Core Function Reference Table](https://deepwiki.com/charmbracelet/skate/6-reference#core-function-reference-table)
    
*   [Data Storage Reference](https://deepwiki.com/charmbracelet/skate/6-reference#data-storage-reference)
    
*   [Storage Structure](https://deepwiki.com/charmbracelet/skate/6-reference#storage-structure)
    
*   [Storage Path Resolution](https://deepwiki.com/charmbracelet/skate/6-reference#storage-path-resolution)
    
*   [Database Operations Flow](https://deepwiki.com/charmbracelet/skate/6-reference#database-operations-flow)
    
*   [Error Handling Reference](https://deepwiki.com/charmbracelet/skate/6-reference#error-handling-reference)
    
*   [Custom Error Types](https://deepwiki.com/charmbracelet/skate/6-reference#custom-error-types)
    
*   [Database Error Handling](https://deepwiki.com/charmbracelet/skate/6-reference#database-error-handling)
    
*   [Error Handling Flow](https://deepwiki.com/charmbracelet/skate/6-reference#error-handling-flow)
    
*   [Key-Value Data Handling](https://deepwiki.com/charmbracelet/skate/6-reference#key-value-data-handling)
    
*   [Key Format and Parsing](https://deepwiki.com/charmbracelet/skate/6-reference#key-format-and-parsing)
    
*   [Binary Data Handling](https://deepwiki.com/charmbracelet/skate/6-reference#binary-data-handling)
    
*   [Database Management](https://deepwiki.com/charmbracelet/skate/6-reference#database-management)
    
*   [Database Operations](https://deepwiki.com/charmbracelet/skate/6-reference#database-operations)
    
*   [Database Path Structure](https://deepwiki.com/charmbracelet/skate/6-reference#database-path-structure)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/skate/6-reference#dependencies)
    
*   [Core Dependencies](https://deepwiki.com/charmbracelet/skate/6-reference#core-dependencies)
    
*   [Dependency Flow](https://deepwiki.com/charmbracelet/skate/6-reference#dependency-flow)
