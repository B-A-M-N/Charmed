Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/mods](https://github.com/charmbracelet/mods "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 7 November 2025 ([07a05d](https://github.com/charmbracelet/mods/commits/07a05d5b)
)

*   [Introduction to Mods](https://deepwiki.com/charmbracelet/mods/1-introduction-to-mods)
    
*   [Installation and Setup](https://deepwiki.com/charmbracelet/mods/1.1-installation-and-setup)
    
*   [Quick Start Guide](https://deepwiki.com/charmbracelet/mods/1.2-quick-start-guide)
    
*   [Core Architecture](https://deepwiki.com/charmbracelet/mods/2-core-architecture)
    
*   [Application Entry Point](https://deepwiki.com/charmbracelet/mods/2.1-application-entry-point)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/mods/2.2-configuration-system)
    
*   [Bubble Tea UI Model](https://deepwiki.com/charmbracelet/mods/2.3-bubble-tea-ui-model)
    
*   [Message Stream Context](https://deepwiki.com/charmbracelet/mods/2.4-message-stream-context)
    
*   [Conversation Management](https://deepwiki.com/charmbracelet/mods/2.5-conversation-management)
    
*   [LLM Provider Integration](https://deepwiki.com/charmbracelet/mods/3-llm-provider-integration)
    
*   [Provider Configuration](https://deepwiki.com/charmbracelet/mods/3.1-provider-configuration)
    
*   [Client Initialization and Streaming](https://deepwiki.com/charmbracelet/mods/3.2-client-initialization-and-streaming)
    
*   [Model Resolution and Fallbacks](https://deepwiki.com/charmbracelet/mods/3.3-model-resolution-and-fallbacks)
    
*   [User Interface and Experience](https://deepwiki.com/charmbracelet/mods/4-user-interface-and-experience)
    
*   [Terminal Capabilities and Styling](https://deepwiki.com/charmbracelet/mods/4.1-terminal-capabilities-and-styling)
    
*   [Animation System](https://deepwiki.com/charmbracelet/mods/4.2-animation-system)
    
*   [Output Rendering and Formatting](https://deepwiki.com/charmbracelet/mods/4.3-output-rendering-and-formatting)
    
*   [Data Persistence](https://deepwiki.com/charmbracelet/mods/5-data-persistence)
    
*   [Conversation Database](https://deepwiki.com/charmbracelet/mods/5.1-conversation-database)
    
*   [Cache System](https://deepwiki.com/charmbracelet/mods/5.2-cache-system)
    
*   [Supporting Systems](https://deepwiki.com/charmbracelet/mods/6-supporting-systems)
    
*   [Flag Parsing and Error Handling](https://deepwiki.com/charmbracelet/mods/6.1-flag-parsing-and-error-handling)
    
*   [Content Loading](https://deepwiki.com/charmbracelet/mods/6.2-content-loading)
    
*   [Message Utilities](https://deepwiki.com/charmbracelet/mods/6.3-message-utilities)
    
*   [Development Guide](https://deepwiki.com/charmbracelet/mods/7-development-guide)
    
*   [Dependencies and Build System](https://deepwiki.com/charmbracelet/mods/7.1-dependencies-and-build-system)
    
*   [Testing Strategy](https://deepwiki.com/charmbracelet/mods/7.2-testing-strategy)
    
*   [Code Quality and CI/CD](https://deepwiki.com/charmbracelet/mods/7.3-code-quality-and-cicd)
    

Menu

Data Persistence
================

Relevant source files

*   [db.go](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go)
    
*   [db\_test.go](https://github.com/charmbracelet/mods/blob/07a05d5b/db_test.go)
    
*   [messages\_test.go](https://github.com/charmbracelet/mods/blob/07a05d5b/messages_test.go)
    

Purpose and Scope
-----------------

This document describes the data persistence architecture in mods, which stores conversation metadata and message histories to enable conversation continuation, branching, and retrieval. The system uses a dual-storage pattern: a SQLite database for lightweight metadata operations and a cache directory for complete message arrays.

For details on the database implementation and schema, see [Conversation Database](https://deepwiki.com/charmbracelet/mods/5.1-conversation-database)
. For the cache system that stores full message histories, see [Cache System](https://deepwiki.com/charmbracelet/mods/5.2-cache-system)
.

Dual-Storage Architecture
-------------------------

The persistence layer separates concerns between **indexable metadata** and **full conversation content**. This design enables fast listing and searching operations while maintaining complete message history for context reconstruction.

**Sources:** [db.go96-107](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L96-L107)
 [db.go1-287](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L1-L287)

### Storage Responsibilities

| Storage Layer | Data Stored | Primary Operations | Performance Characteristics |
| --- | --- | --- | --- |
| **SQLite Database** | Conversation metadata (ID, title, timestamps, model, API) | List, Find, Search, Sort | Fast queries, indexed lookups, SQL aggregations |
| **Cache Directory** | Complete message arrays (system prompts, user messages, assistant responses) | Read full context, Write conversation state | Simple file I/O, content addressable by SHA-1 ID |

**Sources:** [db.go43-79](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L43-L79)
 [db.go100-107](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L100-L107)

Conversation Identification
---------------------------

Conversations are identified using **SHA-1 hashes** generated at creation time, similar to Git commits. Users can reference conversations by:

*   **Full SHA-1 ID** (40 characters): `df31ae23ab8b75b5643c2f846c570997edc71333`
*   **Short SHA-1 prefix** (minimum 4 characters): `df31` or `df31ae`
*   **Exact title match**: User-assigned human-readable titles

**Sources:** [db.go251-271](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L251-L271)
 [db.go192-219](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L192-L219)

Database Schema and Operations
------------------------------

### Conversations Table

The `conversations` table stores metadata with the following schema:

**Indexes:**

*   `idx_conv_id` on `id` column
*   `idx_conv_title` on `title` column

**Sources:** [db.go43-64](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L43-L64)

### Migration System

The database uses a **column-existence check** pattern for migrations, adding columns if they don't exist:

The `hasColumn` function queries `pragma_table_info` to check for column existence before attempting to add it.

**Sources:** [db.go29-82](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L29-L82)
 [db.go84-94](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L84-L94)

### Core Operations Mapping

| Operation | Method | SQL Pattern | Use Case |
| --- | --- | --- | --- |
| **Save** | `Save(id, title, api, model string)` | UPDATE then INSERT | Create or update conversation metadata |
| **Find** | `Find(in string)` | WHERE id glob ? OR title = ? | Retrieve by ID prefix or exact title |
| **FindHEAD** | `FindHEAD()` | ORDER BY updated\_at DESC LIMIT 1 | Get most recent conversation |
| **List** | `List()` | ORDER BY updated\_at DESC | Display all conversations |
| **Delete** | `Delete(id string)` | DELETE WHERE id = ? | Remove conversation metadata |
| **ListOlderThan** | `ListOlderThan(t time.Duration)` | WHERE updated\_at < ? | Find conversations for cleanup |
| **Completions** | `Completions(in string)` | UNION of id/title matches | Shell tab-completion support |

**Sources:** [db.go113-147](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L113-L147)
 [db.go149-158](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L149-L158)
 [db.go175-190](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L175-L190)
 [db.go251-271](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L251-L271)
 [db.go273-286](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L273-L286)
 [db.go160-173](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L160-L173)
 [db.go221-249](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L221-L249)

Data Flow for Conversation Operations
-------------------------------------

### Creating a New Conversation

**Sources:** [db.go113-147](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L113-L147)

### Continuing an Existing Conversation

**Sources:** [db.go251-271](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L251-L271)
 [db.go113-147](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L113-L147)

### Branching a Conversation

When using `--continue` with a new `--title`, the conversation is **branched** (forked):

The original conversation's messages are loaded from cache, a new ID is generated, and both the cache and database entries are created for the new conversation without modifying the original.

**Sources:** [db.go113-147](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L113-L147)

Error Handling
--------------

### SQLite-Specific Errors

The `handleSqliteErr` function wraps SQLite errors with human-readable error codes:

**Sources:** [db.go17-27](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L17-L27)

### Conversation Query Errors

The `Find` method returns specific errors for common failure cases:

| Error | Condition | Error Value |
| --- | --- | --- |
| `errNoMatches` | No conversations match the input | "no conversations found" |
| `errManyMatches` | Multiple conversations match the input | "multiple conversations matched the input" |

These errors allow the UI layer to provide specific user feedback, such as suggesting the user provide more characters for disambiguation.

**Sources:** [db.go12-15](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L12-L15)
 [db.go263-270](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L263-L270)

Integration with Application Flow
---------------------------------

The database is initialized once at application startup and shared across the Bubble Tea model lifecycle:

The `convoDB` instance is created in [db.go29-82](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L29-L82)
 and passed to the Bubble Tea model, where it's used throughout the conversation lifecycle to persist metadata while the cache system handles message content.

**Sources:** [db.go29-82](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L29-L82)
 [db.go96-98](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L96-L98)
 [db.go113-147](https://github.com/charmbracelet/mods/blob/07a05d5b/db.go#L113-L147)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Data Persistence](https://deepwiki.com/charmbracelet/mods/5-data-persistence#data-persistence)
    
*   [Purpose and Scope](https://deepwiki.com/charmbracelet/mods/5-data-persistence#purpose-and-scope)
    
*   [Dual-Storage Architecture](https://deepwiki.com/charmbracelet/mods/5-data-persistence#dual-storage-architecture)
    
*   [Storage Responsibilities](https://deepwiki.com/charmbracelet/mods/5-data-persistence#storage-responsibilities)
    
*   [Conversation Identification](https://deepwiki.com/charmbracelet/mods/5-data-persistence#conversation-identification)
    
*   [Database Schema and Operations](https://deepwiki.com/charmbracelet/mods/5-data-persistence#database-schema-and-operations)
    
*   [Conversations Table](https://deepwiki.com/charmbracelet/mods/5-data-persistence#conversations-table)
    
*   [Migration System](https://deepwiki.com/charmbracelet/mods/5-data-persistence#migration-system)
    
*   [Core Operations Mapping](https://deepwiki.com/charmbracelet/mods/5-data-persistence#core-operations-mapping)
    
*   [Data Flow for Conversation Operations](https://deepwiki.com/charmbracelet/mods/5-data-persistence#data-flow-for-conversation-operations)
    
*   [Creating a New Conversation](https://deepwiki.com/charmbracelet/mods/5-data-persistence#creating-a-new-conversation)
    
*   [Continuing an Existing Conversation](https://deepwiki.com/charmbracelet/mods/5-data-persistence#continuing-an-existing-conversation)
    
*   [Branching a Conversation](https://deepwiki.com/charmbracelet/mods/5-data-persistence#branching-a-conversation)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/mods/5-data-persistence#error-handling)
    
*   [SQLite-Specific Errors](https://deepwiki.com/charmbracelet/mods/5-data-persistence#sqlite-specific-errors)
    
*   [Conversation Query Errors](https://deepwiki.com/charmbracelet/mods/5-data-persistence#conversation-query-errors)
    
*   [Integration with Application Flow](https://deepwiki.com/charmbracelet/mods/5-data-persistence#integration-with-application-flow)
