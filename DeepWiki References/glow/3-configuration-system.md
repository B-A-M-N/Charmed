Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/glow](https://github.com/charmbracelet/glow "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 2 January 2026 ([752de9](https://github.com/charmbracelet/glow/commits/752de97c)
)

*   [Overview](https://deepwiki.com/charmbracelet/glow/1-overview)
    
*   [Dependencies](https://deepwiki.com/charmbracelet/glow/1.1-dependencies)
    
*   [Installation and Usage](https://deepwiki.com/charmbracelet/glow/1.2-installation-and-usage)
    
*   [Architecture](https://deepwiki.com/charmbracelet/glow/2-architecture)
    
*   [Command-Line Interface](https://deepwiki.com/charmbracelet/glow/2.1-command-line-interface)
    
*   [Terminal User Interface](https://deepwiki.com/charmbracelet/glow/2.2-terminal-user-interface)
    
*   [Stash View](https://deepwiki.com/charmbracelet/glow/2.2.1-stash-view)
    
*   [Document View](https://deepwiki.com/charmbracelet/glow/2.2.2-document-view)
    
*   [Markdown Processing](https://deepwiki.com/charmbracelet/glow/2.3-markdown-processing)
    
*   [UI Components and Styling](https://deepwiki.com/charmbracelet/glow/2.4-ui-components-and-styling)
    
*   [Configuration System](https://deepwiki.com/charmbracelet/glow/3-configuration-system)
    
*   [Developer Guide](https://deepwiki.com/charmbracelet/glow/4-developer-guide)
    
*   [Build and Release Process](https://deepwiki.com/charmbracelet/glow/4.1-build-and-release-process)
    
*   [Code Quality and Testing](https://deepwiki.com/charmbracelet/glow/4.2-code-quality-and-testing)
    

Menu

Configuration System
====================

Relevant source files

*   [config\_cmd.go](https://github.com/charmbracelet/glow/blob/752de97c/config_cmd.go)
    
*   [main.go](https://github.com/charmbracelet/glow/blob/752de97c/main.go)
    
*   [ui/config.go](https://github.com/charmbracelet/glow/blob/752de97c/ui/config.go)
    

The Configuration System manages all user-configurable settings in Glow, providing a flexible three-tier hierarchy where CLI flags override environment variables, which in turn override values from the `glow.yml` configuration file. This system determines rendering behavior (style, width, pager usage), TUI behavior (mouse support, file visibility), and content processing options (line numbers, newline preservation).

For information about how these configuration values are used in the TUI, see [Terminal User Interface](https://deepwiki.com/charmbracelet/glow/2.2-terminal-user-interface)
. For CLI rendering behavior, see [Command-Line Interface](https://deepwiki.com/charmbracelet/glow/2.1-command-line-interface)
.

Configuration Architecture
--------------------------

Glow's configuration system uses a layered approach where settings can be specified at multiple levels, with higher-priority sources overriding lower-priority ones.

**Sources:** [main.go165-209](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L165-L209)
 [main.go384-468](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L384-L468)

Configuration File Discovery
----------------------------

The configuration file (`glow.yml`) is searched for in multiple locations following the XDG Base Directory specification. The search order prioritizes user-specific configuration directories.

**Configuration Search Paths (in priority order):**

| Priority | Environment Variable | Default Fallback Path |
| --- | --- | --- |
| 1   | `GLOW_CONFIG_HOME` | N/A |
| 2   | `XDG_CONFIG_HOME` | `$XDG_CONFIG_HOME/glow/` |
| 3   | User scope | Platform-dependent (via go-app-paths) |

**Sources:** [main.go426-468](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L426-L468)

Configuration File Format
-------------------------

The `glow.yml` file uses YAML format and contains rendering and TUI behavior settings. If the file doesn't exist, it is automatically created with default values when accessed via the `glow config` command.

**Default Configuration Template:**

The default configuration is defined in [config\_cmd.go16-26](https://github.com/charmbracelet/glow/blob/752de97c/config_cmd.go#L16-L26)
 as the `defaultConfig` constant.

### Configuration File Creation

**Sources:** [config\_cmd.go56-88](https://github.com/charmbracelet/glow/blob/752de97c/config_cmd.go#L56-L88)

Configuration Options
---------------------

The following table lists all configuration options, their types, default values, and where they are used:

| Option | Type | Default | CLI Flag | Environment Variable | Description | Usage Mode |
| --- | --- | --- | --- | --- | --- | --- |
| `style` | string | `"auto"` | `-s, --style` | `GLOW_STYLE` | Glamour style name or JSON file path | CLI & TUI |
| `width` | uint | `0` (auto-detect) | `-w, --width` | `GLOW_WIDTH` | Word-wrap width (0 = auto) | CLI & TUI |
| `pager` | bool | `false` | `-p, --pager` | `GLOW_PAGER` | Use external pager (e.g., less) | CLI only |
| `tui` | bool | `false` | `-t, --tui` | `GLOW_TUI` | Force TUI mode | CLI only |
| `all` | bool | `true` | `-a, --all` | `GLOW_ALL` | Show hidden/system files | TUI only |
| `mouse` | bool | `false` | `-m, --mouse` | `GLOW_MOUSE` | Enable mouse wheel support | TUI only |
| `preserveNewLines` | bool | `false` | `-n, --preserve-new-lines` | `GLOW_PRESERVENEWLINES` | Preserve newlines in output | CLI & TUI |
| `showLineNumbers` | bool | `false` | `-l, --line-numbers` | `GLOW_SHOWLINENUMBERS` | Display line numbers | TUI only |

**Sources:** [main.go36-44](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L36-L44)
 [main.go396-422](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L396-L422)

Configuration Resolution Flow
-----------------------------

Configuration values flow through multiple stages before reaching the application components. The `validateOptions` function serves as the central resolution point.

**Sources:** [main.go165-209](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L165-L209)

### Style Validation

The `validateStyle` function checks whether a style is a built-in Glamour style or a path to a custom JSON style file.

**Style Validation Logic:**

1.  If `style == "auto"`: Use automatic style selection (no validation needed)
2.  If `style` exists in `styles.DefaultStyles`: Use built-in style
3.  Otherwise: Expand path and verify file exists

**Sources:** [main.go151-163](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L151-L163)

Environment Variables
---------------------

Glow supports two categories of environment variables:

### 1\. Viper-Managed Variables (GLOW\_\* prefix)

Viper automatically reads environment variables with the `GLOW_` prefix after calling `viper.SetEnvPrefix("glow")` and `viper.AutomaticEnv()` at [main.go448-449](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L448-L449)

**Examples:**

*   `GLOW_STYLE="dark"` → maps to `style` configuration
*   `GLOW_WIDTH=100` → maps to `width` configuration
*   `GLOW_PAGER=true` → maps to `pager` configuration

### 2\. TUI-Specific Environment Variables

The `ui.Config` struct uses the `env` tag to parse environment variables directly in [ui/config.go4-20](https://github.com/charmbracelet/glow/blob/752de97c/ui/config.go#L4-L20)
:

| Environment Variable | Field | Default | Purpose |
| --- | --- | --- | --- |
| `GOPATH` | `Gopath` | N/A | Go workspace path (for file discovery) |
| `HOME` | `HomeDir` | N/A | User home directory |
| `GLAMOUR_STYLE` | `GlamourStyle` | N/A | Override Glamour rendering style |
| `GLOW_HIGH_PERFORMANCE_PAGER` | `HighPerformancePager` | `true` | Enable performance optimizations |
| `GLOW_ENABLE_GLAMOUR` | `GlamourEnabled` | `true` | Enable/disable Glamour rendering |

These are parsed using the `env.ParseAs<FileRef file-url="https://github.com/charmbracelet/glow/blob/752de97c/ui.Config" undefined file-path="ui.Config">Hii</FileRef>` call in [main.go346](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L346-L346)

**Sources:** [main.go448-449](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L448-L449)
 [ui/config.go4-20](https://github.com/charmbracelet/glow/blob/752de97c/ui/config.go#L4-L20)
 [main.go346](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L346-L346)

TUI Configuration Propagation
-----------------------------

When launching TUI mode, configuration values are propagated from global variables to the `ui.Config` struct, which is then passed to `ui.NewProgram()`.

**Sources:** [main.go344-369](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L344-L369)
 [ui/config.go4-20](https://github.com/charmbracelet/glow/blob/752de97c/ui/config.go#L4-L20)

Config Command
--------------

The `glow config` command provides an interactive way to edit the configuration file using the user's preferred editor.

**Command Definition:**

The `configCmd` is defined as a Cobra command and registered in [main.go423](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L423-L423)
:

    rootCmd.AddCommand(configCmd, manCmd)
    

**Editor Selection:**

The editor is determined using the `charmbracelet/x/editor` package, which respects:

1.  `EDITOR` environment variable
2.  `VISUAL` environment variable
3.  Platform-specific defaults

**Sources:** [config\_cmd.go28-54](https://github.com/charmbracelet/glow/blob/752de97c/config_cmd.go#L28-L54)
 [main.go423](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L423-L423)

Viper Flag Bindings
-------------------

CLI flags are bound to Viper configuration keys to enable the three-tier configuration hierarchy. This binding is performed in the `init()` function.

**Flag Binding Table:**

| Cobra Flag | Viper Key | Binding Code Location |
| --- | --- | --- |
| `--pager` | `pager` | [main.go409](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L409-L409) |
| `--tui` | `tui` | [main.go410](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L410-L410) |
| `--style` | `style` | [main.go411](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L411-L411) |
| `--width` | `width` | [main.go412](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L412-L412) |
| `--mouse` | `mouse` | [main.go414](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L414-L414) |
| `--preserve-new-lines` | `preserveNewLines` | [main.go415](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L415-L415) |
| `--line-numbers` | `showLineNumbers` | [main.go416](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L416-L416) |
| `--all` | `all` | [main.go417](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L417-L417) |

**Viper Default Values:**

Default values are set for certain configurations in [main.go419-421](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L419-L421)
:

**Resolution Priority:**

When `viper.GetString("key")` or similar methods are called, Viper resolves the value in this order:

1.  Explicitly set via `viper.Set()`
2.  CLI flag value (via `BindPFlag`)
3.  Environment variable (via `AutomaticEnv()` with prefix)
4.  Configuration file value (via `ReadInConfig()`)
5.  Default value (via `SetDefault()`)

**Sources:** [main.go408-421](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L408-L421)

Configuration Override Detection
--------------------------------

Certain configuration behaviors change based on whether a flag was explicitly set by the user. The `cmd.Flags().Changed(flagName)` method detects explicit flag usage.

**Override Detection Examples:**

### 1\. Style Override for Non-TTY Output

When stdout is not a terminal (e.g., piped to a file), Glow automatically uses the `"notty"` style unless the user explicitly set a style via the `--style` flag.

### 2\. Pager/TUI Flag Detection

In [main.go316](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L316-L316)
 and [main.go330](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L330-L330)
 the code checks both the configuration value AND whether the flag was explicitly changed:

    case pager || cmd.Flags().Changed("pager"):
    case tui || cmd.Flags().Changed("tui"):
    

This ensures that even if the config file has `pager: false`, running `glow --pager file.md` will correctly use the pager.

**Sources:** [main.go185-189](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L185-L189)
 [main.go316](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L316-L316)
 [main.go330](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L330-L330)

Configuration Validation and Conflicts
--------------------------------------

The configuration system performs validation to ensure consistency and detect mutually exclusive options.

**Validation Rules:**

1.  **Mutual Exclusion:** The `pager` and `tui` options cannot both be true [main.go175-177](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L175-L177)
    
2.  **Style Existence:** Custom style paths must point to existing files [main.go153-163](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L153-L163)
    
3.  **File Extension:** Config files must use `.yaml` or `.yml` extensions [config\_cmd.go64-66](https://github.com/charmbracelet/glow/blob/752de97c/config_cmd.go#L64-L66)
    

**Validation Errors:**

**Sources:** [main.go165-209](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L165-L209)
 [main.go151-163](https://github.com/charmbracelet/glow/blob/752de97c/main.go#L151-L163)
 [config\_cmd.go64-66](https://github.com/charmbracelet/glow/blob/752de97c/config_cmd.go#L64-L66)

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Configuration System](https://deepwiki.com/charmbracelet/glow/3-configuration-system#configuration-system)
    
*   [Configuration Architecture](https://deepwiki.com/charmbracelet/glow/3-configuration-system#configuration-architecture)
    
*   [Configuration File Discovery](https://deepwiki.com/charmbracelet/glow/3-configuration-system#configuration-file-discovery)
    
*   [Configuration File Format](https://deepwiki.com/charmbracelet/glow/3-configuration-system#configuration-file-format)
    
*   [Configuration File Creation](https://deepwiki.com/charmbracelet/glow/3-configuration-system#configuration-file-creation)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/glow/3-configuration-system#configuration-options)
    
*   [Configuration Resolution Flow](https://deepwiki.com/charmbracelet/glow/3-configuration-system#configuration-resolution-flow)
    
*   [Style Validation](https://deepwiki.com/charmbracelet/glow/3-configuration-system#style-validation)
    
*   [Environment Variables](https://deepwiki.com/charmbracelet/glow/3-configuration-system#environment-variables)
    
*   [1\. Viper-Managed Variables (GLOW\_\* prefix)](https://deepwiki.com/charmbracelet/glow/3-configuration-system#1-viper-managed-variables-glow_-prefix)
    
*   [2\. TUI-Specific Environment Variables](https://deepwiki.com/charmbracelet/glow/3-configuration-system#2-tui-specific-environment-variables)
    
*   [TUI Configuration Propagation](https://deepwiki.com/charmbracelet/glow/3-configuration-system#tui-configuration-propagation)
    
*   [Config Command](https://deepwiki.com/charmbracelet/glow/3-configuration-system#config-command)
    
*   [Viper Flag Bindings](https://deepwiki.com/charmbracelet/glow/3-configuration-system#viper-flag-bindings)
    
*   [Configuration Override Detection](https://deepwiki.com/charmbracelet/glow/3-configuration-system#configuration-override-detection)
    
*   [1\. Style Override for Non-TTY Output](https://deepwiki.com/charmbracelet/glow/3-configuration-system#1-style-override-for-non-tty-output)
    
*   [2\. Pager/TUI Flag Detection](https://deepwiki.com/charmbracelet/glow/3-configuration-system#2-pagertui-flag-detection)
    
*   [Configuration Validation and Conflicts](https://deepwiki.com/charmbracelet/glow/3-configuration-system#configuration-validation-and-conflicts)
