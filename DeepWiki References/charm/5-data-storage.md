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

Data Storage
============

Relevant source files

*   [client/auth.go](https://github.com/charmbracelet/charm/blob/aff3071d/client/auth.go)
    
*   [client/http.go](https://github.com/charmbracelet/charm/blob/aff3071d/client/http.go)
    
*   [cmd/cmd.go](https://github.com/charmbracelet/charm/blob/aff3071d/cmd/cmd.go)
    
*   [fs/fs.go](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go)
    
*   [kv/kv.go](https://github.com/charmbracelet/charm/blob/aff3071d/kv/kv.go)
    

This page provides a technical overview of the data storage mechanisms in Charm. Charm offers two complementary data storage systems: a Key-Value Store (KV) for structured data and a File System (FS) for hierarchical file storage. Both systems implement client-side encryption to ensure data security.

For information about encryption implementation details, see [Encryption](https://deepwiki.com/charmbracelet/charm/5.3-encryption)
.

Storage Overview
----------------

Charm provides two primary storage mechanisms:

1.  **Key-Value Store (KV)**: A BadgerDB-backed key-value store that synchronizes with Charm Cloud
2.  **File System (FS)**: A cloud-based file system implementing Go's standard `fs.FS` interfaces

Both systems share several key characteristics:

*   Client-side encryption and decryption
*   Transparent cloud synchronization
*   Multi-device support through linked accounts
*   Offline-first design with eventual consistency

Sources: [kv/kv.go17-30](https://github.com/charmbracelet/charm/blob/aff3071d/kv/kv.go#L17-L30)
 [fs/fs.go24-30](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go#L24-L30)

Key-Value Store
---------------

The Key-Value Store (KV) is built on BadgerDB, a performant embedded key-value database. It provides a persistent, synced storage mechanism with transaction support.

### KV Architecture

Sources: [kv/kv.go1-30](https://github.com/charmbracelet/charm/blob/aff3071d/kv/kv.go#L1-L30)
 [kv/kv.go102-214](https://github.com/charmbracelet/charm/blob/aff3071d/kv/kv.go#L102-L214)

### KV Operations

The KV store provides both transaction-based operations and convenience methods:

| Operation | Description | Method |
| --- | --- | --- |
| Create transaction | Creates a transaction with cloud-managed timestamp | `NewTransaction(update bool)` |
| View | Read-only transaction | `View(func(txn *badger.Txn) error)` |
| Commit | Commits transaction and syncs diff to cloud | `Commit(txn *badger.Txn, callback func(error))` |
| Set | Convenience method to set a key-value pair | `Set(key []byte, value []byte)` |
| Get | Convenience method to retrieve a value | `Get(key []byte)` |
| Delete | Convenience method to delete a key | `Delete(key []byte)` |
| Keys | Returns all keys in the store | `Keys()` |
| Sync | Synchronize with cloud changes | `Sync()` |
| Reset | Delete local data and resync from cloud | `Reset()` |

Sources: [kv/kv.go102-245](https://github.com/charmbracelet/charm/blob/aff3071d/kv/kv.go#L102-L245)

### KV Data Flow

Sources: [kv/kv.go133-178](https://github.com/charmbracelet/charm/blob/aff3071d/kv/kv.go#L133-L178)

File System (FS)
----------------

The File System (FS) provides a cloud-based file system that implements Go's standard `fs.FS`, `fs.ReadFileFS`, and `fs.ReadDirFS` interfaces, with additional write capabilities. All file data and paths are encrypted client-side.

### FS Architecture

Sources: [fs/fs.go24-30](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go#L24-L30)
 [fs/fs.go91-285](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go#L91-L285)

### FS Components

| Component | Description | Implementation |
| --- | --- | --- |
| FS  | Main file system implementation | `FS` struct |
| File | Implements `fs.File` interface | `File` struct |
| FileInfo | Implements `fs.FileInfo` interface | `FileInfo` struct |
| DirFile | Represents a directory entry | `DirFile` struct |

Sources: [fs/fs.go24-71](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go#L24-L71)

### FS Operations

The FS provides standard file system operations:

| Operation | Description | Method |
| --- | --- | --- |
| Open | Opens a file or directory | `Open(name string)` |
| ReadFile | Reads an entire file | `ReadFile(name string)` |
| WriteFile | Writes data to a file | `WriteFile(name string, src fs.File)` |
| Remove | Deletes a file | `Remove(name string)` |
| ReadDir | Lists directory contents | `ReadDir(name string)` |
| EncryptPath | Encrypts a file path | `EncryptPath(path string)` |
| DecryptPath | Decrypts a file path | `DecryptPath(path string)` |

Sources: [fs/fs.go91-425](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go#L91-L425)

### FS Data Flow

Sources: [fs/fs.go91-179](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go#L91-L179)
 [fs/fs.go184-270](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go#L184-L270)

Client-Server Data Transport
----------------------------

Both storage systems rely on the Charm client's HTTP communication layer to securely transmit data to and from the Charm server.

Sources: [client/http.go25-91](https://github.com/charmbracelet/charm/blob/aff3071d/client/http.go#L25-L91)
 [client/auth.go10-45](https://github.com/charmbracelet/charm/blob/aff3071d/client/auth.go#L10-L45)

Storage Security Model
----------------------

Charm's storage systems implement a robust security model with these key features:

1.  **Client-side encryption**: All data is encrypted before leaving the client
2.  **Path encryption**: File paths and directory structures are encrypted
3.  **JWT authentication**: All requests use JWT tokens for authentication
4.  **SSH-based key exchange**: Initial authentication leverages SSH for secure key exchange
5.  **Local caching**: Data is cached locally for offline use

Sources: [fs/fs.go145-152](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go#L145-L152)
 [kv/kv.go72-79](https://github.com/charmbracelet/charm/blob/aff3071d/kv/kv.go#L72-L79)
 [client/auth.go10-45](https://github.com/charmbracelet/charm/blob/aff3071d/client/auth.go#L10-L45)

Integration with Applications
-----------------------------

Application developers can integrate with Charm's storage systems using standard Go interfaces:

Sources: [kv/kv.go46-70](https://github.com/charmbracelet/charm/blob/aff3071d/kv/kv.go#L46-L70)
 [fs/fs.go73-89](https://github.com/charmbracelet/charm/blob/aff3071d/fs/fs.go#L73-L89)

Summary
-------

Charm's data storage systems provide a secure, synchronized approach to data persistence:

*   The Key-Value Store offers a transactional database with cloud synchronization
*   The File System provides a hierarchical storage system compatible with Go's standard interfaces
*   Both systems implement client-side encryption for data security
*   The underlying transport layer handles authentication and secure communication

Together, these systems enable developers to build applications with persistent storage that works across multiple devices through Charm Cloud.

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Data Storage](https://deepwiki.com/charmbracelet/charm/5-data-storage#data-storage)
    
*   [Storage Overview](https://deepwiki.com/charmbracelet/charm/5-data-storage#storage-overview)
    
*   [Key-Value Store](https://deepwiki.com/charmbracelet/charm/5-data-storage#key-value-store)
    
*   [KV Architecture](https://deepwiki.com/charmbracelet/charm/5-data-storage#kv-architecture)
    
*   [KV Operations](https://deepwiki.com/charmbracelet/charm/5-data-storage#kv-operations)
    
*   [KV Data Flow](https://deepwiki.com/charmbracelet/charm/5-data-storage#kv-data-flow)
    
*   [File System (FS)](https://deepwiki.com/charmbracelet/charm/5-data-storage#file-system-fs)
    
*   [FS Architecture](https://deepwiki.com/charmbracelet/charm/5-data-storage#fs-architecture)
    
*   [FS Components](https://deepwiki.com/charmbracelet/charm/5-data-storage#fs-components)
    
*   [FS Operations](https://deepwiki.com/charmbracelet/charm/5-data-storage#fs-operations)
    
*   [FS Data Flow](https://deepwiki.com/charmbracelet/charm/5-data-storage#fs-data-flow)
    
*   [Client-Server Data Transport](https://deepwiki.com/charmbracelet/charm/5-data-storage#client-server-data-transport)
    
*   [Storage Security Model](https://deepwiki.com/charmbracelet/charm/5-data-storage#storage-security-model)
    
*   [Integration with Applications](https://deepwiki.com/charmbracelet/charm/5-data-storage#integration-with-applications)
    
*   [Summary](https://deepwiki.com/charmbracelet/charm/5-data-storage#summary)
