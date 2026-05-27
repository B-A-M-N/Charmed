Loading...

Index your code with Devin

[DeepWiki](https://deepwiki.com/)

[DeepWiki](https://deepwiki.com/)

[charmbracelet/fang](https://github.com/charmbracelet/fang "Open repository")

Index your code with

Devin

Edit WikiShare

Loading...

Last indexed: 15 March 2026 ([d89b30](https://github.com/charmbracelet/fang/commits/d89b30af)
)

*   [Overview](https://deepwiki.com/charmbracelet/fang/1-overview)
    
*   [Getting Started](https://deepwiki.com/charmbracelet/fang/2-getting-started)
    
*   [Core Components](https://deepwiki.com/charmbracelet/fang/3-core-components)
    
*   [fang.Execute Function](https://deepwiki.com/charmbracelet/fang/3.1-fang.execute-function)
    
*   [Help Rendering System](https://deepwiki.com/charmbracelet/fang/3.2-help-rendering-system)
    
*   [Styling and Themes](https://deepwiki.com/charmbracelet/fang/3.3-styling-and-themes)
    
*   [Error Handling](https://deepwiki.com/charmbracelet/fang/3.4-error-handling)
    
*   [Configuration Options](https://deepwiki.com/charmbracelet/fang/4-configuration-options)
    
*   [Features](https://deepwiki.com/charmbracelet/fang/5-features)
    
*   [Version Management](https://deepwiki.com/charmbracelet/fang/5.1-version-management)
    
*   [Manpage Generation](https://deepwiki.com/charmbracelet/fang/5.2-manpage-generation)
    
*   [Shell Completions](https://deepwiki.com/charmbracelet/fang/5.3-shell-completions)
    
*   [Signal Handling](https://deepwiki.com/charmbracelet/fang/5.4-signal-handling)
    
*   [Dependencies and Architecture](https://deepwiki.com/charmbracelet/fang/6-dependencies-and-architecture)
    
*   [Testing and Examples](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples)
    
*   [Golden File Testing](https://deepwiki.com/charmbracelet/fang/7.1-golden-file-testing)
    
*   [Example Application](https://deepwiki.com/charmbracelet/fang/7.2-example-application)
    
*   [Development](https://deepwiki.com/charmbracelet/fang/8-development)
    
*   [Code Quality](https://deepwiki.com/charmbracelet/fang/8.1-code-quality)
    
*   [Platform Support](https://deepwiki.com/charmbracelet/fang/8.2-platform-support)
    
*   [License](https://deepwiki.com/charmbracelet/fang/8.3-license)
    

Menu

Testing and Examples
====================

Relevant source files

*   [example/main.go](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go)
    
*   [fang\_test.go](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go)
    
*   [testdata/TestSetup/without\_version.golden](https://github.com/charmbracelet/fang/blob/d89b30af/testdata/TestSetup/without_version.golden)
    

Fang uses golden file testing to validate CLI output consistency and includes a comprehensive example application demonstrating all major features. The test suite exercises help rendering, error handling, version management, and feature configuration through automated comparison against expected outputs.

For detailed information on golden file testing methodology, see page 7.1. For a complete walkthrough of the example application, see page 7.2.

Testing Architecture
--------------------

Fang's testing strategy combines unit tests with golden file validation to ensure consistent terminal output across different scenarios and configurations.

### Testing Infrastructure Components

The test infrastructure uses factory functions (`mkroot`) to create fresh `cobra.Command` instances for each test, ensuring test isolation. The `exercise()` helper runs three standard scenarios (help, version, error) while `doExercise()` runs custom scenarios with specific arguments and options.

_Sources: [fang\_test.go302-369](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L302-L369)
 [example/main.go14-153](https://github.com/charmbracelet/fang/blob/d89b30af/example/main.go#L14-L153)
_

Test Helper Functions
---------------------

The test infrastructure provides reusable patterns for validating CLI behavior through golden file comparison. Helper functions create fresh command instances, execute scenarios, and validate outputs.

### Test Execution Flow

The `__FANG_TEST_WIDTH` environment variable (set to `"45"` in tests) ensures consistent terminal width for reproducible output formatting across different development environments.

_Sources: [fang\_test.go333-353](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L333-L353)
_

### Helper Function Reference

| Function | Signature | Purpose |
| --- | --- | --- |
| `toMkroot()` | `func(c *cobra.Command) func() *cobra.Command` | Wraps a command in factory function for test isolation |
| `exercise()` | `func(t, mkroot, options...)` | Runs three standard subtests: error, version, help |
| `doExercise()` | `func(t, mkroot, args, assert, options...)` | Executes single scenario with custom arguments |
| `assertNoError()` | `func(t, err, stdout, stderr)` | Validates successful execution, compares stdout to golden file |
| `assertError()` | `func(t, err, stderr, stderr)` | Validates error case, compares stderr to golden file |

The `exercise()` function automatically runs three subtests for each command configuration:

1.  `error` subtest: passes `[]string{"--nope-nope-nope"}` to trigger unknown flag error
2.  `version` subtest: passes `[]string{"--version"}` to validate version output
3.  `help` subtest: passes `[]string{"--help"}` to validate help rendering

_Sources: [fang\_test.go302-369](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L302-L369)
_

Test Scenarios
--------------

The `TestSetup()` function in [fang\_test.go](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go)
 contains 13 test cases validating different command configurations, feature toggles, and output scenarios.

### Complete Test Case Catalog

| Test Case | Location | Configuration | Validated Outputs |
| --- | --- | --- | --- |
| `simple` | [fang\_test.go16-20](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L16-L20) | Minimal command with `Use` only | help, version, error |
| `custom error handler` | [fang\_test.go22-32](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L22-L32) | `WithErrorHandler()` option | Custom error formatting |
| `complete` | [fang\_test.go34-50](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L34-L50) | Command with Short, Long, includes man subtest | help, version, error, man -h |
| `use with args` | [fang\_test.go52-58](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L52-L58) | `Use: "simple [args] [something-else]"` | Help with argument placeholders |
| `without completions` | [fang\_test.go60-78](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L60-L78) | `WithoutCompletions()` option | Help without completion command |
| `without manpage` | [fang\_test.go80-97](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L80-L97) | `WithoutManpage()` option | Help without man command |
| `without version` | [fang\_test.go99-107](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L99-L107) | `WithoutVersion()` option | Error on --version flag |
| `with version and hash` | [fang\_test.go109-116](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L109-L116) | `WithVersion("v1.2.3")`, `WithCommit("aaabbb")` | Version output with commit hash |
| `with flags` | [fang\_test.go118-143](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L118-L143) | 13 flags: string, int, float64, bool variations | Flag rendering with defaults |
| `with subcommands` | [fang\_test.go145-184](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L145-L184) | Nested subcommands (sub1, sub1/sub2) | Subcommand help, sub --help, sub sub --help |
| `with command groups` | [fang\_test.go186-217](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L186-L217) | Two groups with categorized commands | Grouped command display |
| `with examples` | [fang\_test.go219-281](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L219-L281) | Complex example block with syntax highlighting | Example formatting, comment styling |
| `with multiline flag descriptions` | [fang\_test.go283-299](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L283-L299) | Multi-line `--format` flag description | Flag description wrapping |

### Test Case Structure Pattern

Each test case defines a `mkroot` factory function that creates a fresh `cobra.Command` instance, then calls `exercise()` which automatically runs error, version, and help subtests. Some test cases add additional custom subtests (e.g., `complete` adds a `man` subtest, `with subcommands` adds `help-sub` and `help-sub-sub` subtests).

_Sources: [fang\_test.go15-300](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L15-L300)
_

### Flag Type Test Coverage

The `with flags` test case validates rendering of all flag types supported by Cobra:

The test also validates hidden flag behavior (`MarkHidden()`) and flags with empty descriptions.

_Sources: [fang\_test.go118-143](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L118-L143)
_

Testing Infrastructure
----------------------

### Test Environment Setup

The testing infrastructure establishes controlled conditions to ensure reproducible results across different development environments.

The test infrastructure controls terminal width using the `__FANG_TEST_WIDTH` environment variable and captures both stdout and stderr for comprehensive output validation.

_Sources: [fang\_test.go243-251](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L243-L251)
_

### Test Pattern Implementation

| Pattern | Purpose | Implementation |
| --- | --- | --- |
| Standard Exercise | Tests help, version, and error scenarios | [fang\_test.go204-233](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L204-L233) |
| Custom Exercise | Tests specific scenarios with custom options | [fang\_test.go235-252](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L235-L252) |
| Assertion Functions | Validates success/error states with golden files | [fang\_test.go254-262](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L254-L262) |

The test patterns provide consistent structure for validating different aspects of fang's functionality while maintaining code reuse and readability.

_Sources: [fang\_test.go204-262](https://github.com/charmbracelet/fang/blob/d89b30af/fang_test.go#L204-L262)
_

Dismiss

Refresh this wiki

Enter email to refresh

### On this page

*   [Testing and Examples](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#testing-and-examples)
    
*   [Testing Architecture](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#testing-architecture)
    
*   [Testing Infrastructure Components](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#testing-infrastructure-components)
    
*   [Test Helper Functions](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#test-helper-functions)
    
*   [Test Execution Flow](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#test-execution-flow)
    
*   [Helper Function Reference](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#helper-function-reference)
    
*   [Test Scenarios](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#test-scenarios)
    
*   [Complete Test Case Catalog](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#complete-test-case-catalog)
    
*   [Test Case Structure Pattern](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#test-case-structure-pattern)
    
*   [Flag Type Test Coverage](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#flag-type-test-coverage)
    
*   [Testing Infrastructure](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#testing-infrastructure)
    
*   [Test Environment Setup](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#test-environment-setup)
    
*   [Test Pattern Implementation](https://deepwiki.com/charmbracelet/fang/7-testing-and-examples#test-pattern-implementation)
