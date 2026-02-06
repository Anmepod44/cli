# Design Document: CLI Command Assistant

## Overview

The CLI Command Assistant is a Go-based terminal application that bridges natural language and Linux shell commands. The system architecture follows a modular design with clear separation of concerns: command parsing, execution, history management, and user interface. The application leverages Go's concurrency primitives (goroutines and channels) to handle command execution asynchronously while maintaining a responsive CLI interface.

The core workflow involves: (1) accepting natural language input from the user, (2) translating it to shell commands using pattern matching and rule-based translation, (3) presenting the command for review, and (4) optionally executing it with proper safety checks. All operations are designed to be non-blocking, with concurrent command execution managed through a worker pool pattern.

## Architecture

The application follows a layered architecture with the following components:

```
┌─────────────────────────────────────────────────────────┐
│                    CLI Interface Layer                   │
│  (Input handling, display, prompts, user interaction)   │
└─────────────────────────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────┐
│                   Application Layer                      │
│     (Command orchestration, workflow management)        │
└─────────────────────────────────────────────────────────┘
                           │
        ┌──────────────────┼──────────────────┐
        ▼                  ▼                  ▼
┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│   Command    │  │   Command    │  │   History    │
│  Generator   │  │  Executor    │  │   Manager    │
└──────────────┘  └──────────────┘  └──────────────┘
        │                  │                  │
        └──────────────────┼──────────────────┘
                           ▼
                  ┌──────────────┐
                  │   Storage    │
                  │   (Config,   │
                  │   History)   │
                  └──────────────┘
```

**Component Responsibilities:**

- **CLI Interface Layer**: Handles all user interaction, input parsing, output formatting, and display
- **Application Layer**: Orchestrates the workflow between components and manages application state
- **Command Generator**: Translates natural language to shell commands using pattern matching
- **Command Executor**: Executes shell commands asynchronously using goroutines and channels
- **History Manager**: Persists and retrieves command history with search capabilities
- **Storage**: Manages configuration files and persistent data

## Components and Interfaces

### 1. Command Generator

**Purpose**: Translate natural language input to Linux shell commands.

**Interface**:
```go
type CommandGenerator interface {
    // Generate translates natural language to a shell command
    Generate(input string) (*Command, error)
    
    // Validate checks if a command is safe to execute
    Validate(cmd *Command) []Warning
}

type Command struct {
    Raw         string   // The shell command string
    Description string   // Human-readable explanation
    Flags       []Flag   // Parsed flags and options
    IsDangerous bool     // Whether command is potentially destructive
    RequiresSudo bool    // Whether command needs elevated privileges
}

type Flag struct {
    Name        string
    Description string
}

type Warning struct {
    Level   WarningLevel // Info, Warning, Critical
    Message string
}
```

**Implementation Strategy**:
- Use a rule-based pattern matching system with regex patterns
- Maintain a mapping of common natural language phrases to command templates
- Support parameterized templates (e.g., "find files larger than {size}")
- Implement a safety classifier to identify dangerous operations

### 2. Command Executor

**Purpose**: Execute shell commands concurrently with proper output handling.

**Interface**:
```go
type CommandExecutor interface {
    // Execute runs a command asynchronously and returns a result channel
    Execute(cmd *Command) <-chan ExecutionResult
    
    // ExecuteBatch runs multiple commands concurrently
    ExecuteBatch(cmds []*Command) <-chan ExecutionResult
    
    // Cancel stops a running command by ID
    Cancel(id string) error
}

type ExecutionResult struct {
    CommandID string
    Output    string
    Error     error
    ExitCode  int
    Duration  time.Duration
}
```

**Implementation Strategy**:
- Use `os/exec` package to run shell commands
- Implement a worker pool with configurable concurrency limit
- Use channels to communicate results back to the main goroutine
- Capture both stdout and stderr streams
- Support command cancellation via context

### 3. History Manager

**Purpose**: Store and retrieve command history with search capabilities.

**Interface**:
```go
type HistoryManager interface {
    // Add stores a command in history
    Add(cmd *Command, executed bool) error
    
    // List retrieves recent commands with optional limit
    List(limit int) ([]*HistoryEntry, error)
    
    // Search finds commands matching a query
    Search(query string) ([]*HistoryEntry, error)
    
    // Clear removes all history
    Clear() error
}

type HistoryEntry struct {
    ID        string
    Command   *Command
    Timestamp time.Time
    Executed  bool
}
```

**Implementation Strategy**:
- Store history in a JSON file in the user's home directory (~/.cli-assistant/history.json)
- Implement in-memory caching for fast access
- Use a circular buffer to limit history size
- Support full-text search across command strings and descriptions

### 4. CLI Interface

**Purpose**: Provide an interactive terminal interface for user interaction.

**Interface**:
```go
type CLIInterface interface {
    // Start begins the interactive session
    Start() error
    
    // DisplayCommand shows a generated command with formatting
    DisplayCommand(cmd *Command)
    
    // Prompt asks the user for input with options
    Prompt(message string, options []string) (string, error)
    
    // DisplayResult shows command execution results
    DisplayResult(result ExecutionResult)
}
```

**Implementation Strategy**:
- Use a CLI library like `github.com/charmbracelet/bubbletea` or `github.com/manifoldco/promptui` for rich interactions
- Implement syntax highlighting for shell commands
- Support keyboard shortcuts for common actions
- Display progress indicators for long-running commands

### 5. Configuration Manager

**Purpose**: Load and manage application configuration.

**Interface**:
```go
type Config struct {
    AutoExecute      bool
    MaxConcurrent    int
    HistorySize      int
    ConfirmDangerous bool
    ClipboardEnabled bool
}

type ConfigManager interface {
    // Load reads configuration from file
    Load() (*Config, error)
    
    // Save writes configuration to file
    Save(cfg *Config) error
}
```

**Implementation Strategy**:
- Store configuration in YAML or JSON format at ~/.cli-assistant/config.yaml
- Provide sensible defaults (AutoExecute: false, MaxConcurrent: 5, HistorySize: 1000)
- Support environment variable overrides

## Data Models

### Command Structure
```go
type Command struct {
    Raw         string
    Description string
    Flags       []Flag
    IsDangerous bool
    RequiresSudo bool
}
```

### History Entry
```go
type HistoryEntry struct {
    ID        string
    Command   *Command
    Timestamp time.Time
    Executed  bool
}
```

### Execution Result
```go
type ExecutionResult struct {
    CommandID string
    Output    string
    Error     error
    ExitCode  int
    Duration  time.Duration
}
```

### Configuration
```go
type Config struct {
    AutoExecute      bool
    MaxConcurrent    int
    HistorySize      int
    ConfirmDangerous bool
    ClipboardEnabled bool
}
```


## Correctness Properties

A property is a characteristic or behavior that should hold true across all valid executions of a system—essentially, a formal statement about what the system should do. Properties serve as the bridge between human-readable specifications and machine-verifiable correctness guarantees.

### Property 1: Valid Command Generation

*For any* valid natural language input describing a file or system operation, the Command Generator should produce a non-empty shell command string with a description.

**Validates: Requirements 1.1, 2.2**

### Property 2: Error Handling for Invalid Input

*For any* natural language input that cannot be mapped to a valid command, the Command Generator should return an error with an explanatory message.

**Validates: Requirements 1.3**

### Property 3: Command Metadata Completeness

*For any* generated command that contains flags or options, the Command structure should include descriptions for all flags present in the command string.

**Validates: Requirements 2.2, 2.3**

### Property 4: Execution Confirmation Required

*For any* command execution request when auto-execute mode is disabled, the system should require explicit user confirmation before proceeding with execution.

**Validates: Requirements 3.1**

### Property 5: Complete Execution Result Capture

*For any* executed command, the execution result should contain either output (for success) or error information with exit code (for failure), along with execution duration.

**Validates: Requirements 3.3, 3.4**

### Property 6: Clipboard Operation Success

*For any* generated command, when the copy operation is invoked, the system clipboard should contain the exact command string.

**Validates: Requirements 4.1**

### Property 7: History Persistence Round-Trip

*For any* command added to history, after saving and reloading the history from disk, the command should be retrievable with its original content and timestamp.

**Validates: Requirements 5.1, 5.4**

### Property 8: History Chronological Ordering

*For any* history state with multiple entries, retrieving the history list should return entries in reverse chronological order (newest first).

**Validates: Requirements 5.2**

### Property 9: History Search Filtering

*For any* search query and history state, all returned results should contain the search query as a substring in either the command string or description.

**Validates: Requirements 5.5**

### Property 10: Concurrent Output Isolation

*For any* set of commands executed concurrently, each execution result should contain output only from its corresponding command, with no mixing of output streams.

**Validates: Requirements 6.2**

### Property 11: Command Safety Classification

*For any* command that performs file deletion, file modification, or targets system directories (/etc, /sys, /proc, /boot), the Command structure should have IsDangerous set to true.

**Validates: Requirements 7.1, 7.3, 7.4**

### Property 12: Privilege Detection

*For any* command string that begins with "sudo" or requires root access, the Command structure should have RequiresSudo set to true.

**Validates: Requirements 7.2**

### Property 13: Configuration Loading

*For any* valid configuration file, loading the configuration should produce a Config structure with all fields matching the file contents.

**Validates: Requirements 9.1**

### Property 14: History Size Limit Enforcement

*For any* history state where the number of entries exceeds the configured HistorySize limit, the history should contain at most HistorySize entries, keeping the most recent ones.

**Validates: Requirements 9.5**

### Property 15: Error Logging Completeness

*For any* error that occurs during command generation, execution, or history operations, an entry should be written to the error log file with timestamp and error details.

**Validates: Requirements 10.5**

## Error Handling

The application implements comprehensive error handling at multiple levels:

### Command Generation Errors
- **Invalid Input**: Return descriptive error explaining why translation failed
- **Ambiguous Input**: Request clarification with suggested interpretations
- **Unsupported Operations**: Inform user of limitations and suggest alternatives

### Command Execution Errors
- **Non-zero Exit Codes**: Capture and display stderr output with exit code
- **Timeout**: Cancel long-running commands and notify user
- **Permission Denied**: Suggest using sudo or checking file permissions
- **Command Not Found**: Suggest installing required packages

### History Errors
- **Corrupted History File**: Create new history file, backup corrupted one
- **Disk Full**: Warn user and continue with in-memory history only
- **Permission Issues**: Fall back to temporary directory

### Configuration Errors
- **Missing Config**: Use default configuration values
- **Invalid Config Format**: Log error, use defaults, inform user
- **Invalid Values**: Validate and clamp to acceptable ranges

### Clipboard Errors
- **Clipboard Unavailable**: Inform user, offer to display command for manual copy
- **Clipboard Access Denied**: Log error, continue without clipboard functionality

### Concurrency Errors
- **Goroutine Panic**: Recover from panics, log error, continue operation
- **Channel Deadlock**: Implement timeouts on all channel operations
- **Resource Exhaustion**: Limit concurrent operations via worker pool

## Testing Strategy

The CLI Command Assistant will be tested using a dual approach combining unit tests and property-based tests to ensure comprehensive coverage and correctness.

### Unit Testing

Unit tests will verify specific examples, edge cases, and integration points:

**Command Generator Tests**:
- Test common natural language patterns (list files, find files, process management)
- Test edge cases (empty input, special characters, very long input)
- Test specific dangerous command detection (rm -rf, dd, mkfs)
- Test sudo detection for various command formats

**Command Executor Tests**:
- Test successful command execution with output capture
- Test failed command execution with error capture
- Test command cancellation
- Test timeout handling

**History Manager Tests**:
- Test adding and retrieving single entries
- Test empty history state
- Test history file corruption recovery
- Test search with no results

**Configuration Tests**:
- Test loading valid configuration files
- Test missing configuration (defaults)
- Test invalid configuration values

**CLI Interface Tests**:
- Test help command display
- Test exit command handling
- Test startup welcome message

### Property-Based Testing

Property-based tests will verify universal properties across randomized inputs using a Go property testing library such as `gopter` or `rapid`:

**Configuration**: Each property test will run a minimum of 100 iterations to ensure thorough coverage through randomization.

**Test Tagging**: Each property test will include a comment tag referencing its design document property:
```go
// Feature: cli-command-assistant, Property 1: Valid Command Generation
```

**Property Test Coverage**:

1. **Property 1 - Valid Command Generation**: Generate random valid natural language inputs, verify non-empty command and description
2. **Property 2 - Error Handling**: Generate random invalid inputs, verify error is returned
3. **Property 3 - Metadata Completeness**: Generate random commands with flags, verify all flags have descriptions
4. **Property 5 - Execution Result Capture**: Execute random commands, verify results contain output or error info
5. **Property 6 - Clipboard Operation**: Generate random commands, copy to clipboard, verify clipboard contents
6. **Property 7 - History Round-Trip**: Generate random commands, save and reload history, verify persistence
7. **Property 8 - Chronological Ordering**: Generate random history entries with timestamps, verify ordering
8. **Property 9 - Search Filtering**: Generate random history and search queries, verify all results match query
9. **Property 10 - Output Isolation**: Execute random concurrent commands, verify no output mixing
10. **Property 11 - Safety Classification**: Generate random dangerous commands, verify IsDangerous flag
11. **Property 12 - Privilege Detection**: Generate random sudo commands, verify RequiresSudo flag
12. **Property 13 - Config Loading**: Generate random valid configs, verify correct loading
13. **Property 14 - History Size Limits**: Generate random history exceeding limit, verify size enforcement
14. **Property 15 - Error Logging**: Generate random errors, verify all are logged

### Testing Balance

- **Unit tests** focus on specific examples, edge cases, and integration between components
- **Property tests** handle comprehensive input coverage through randomization
- Together, they provide confidence in both specific behaviors and general correctness
- Unit tests catch concrete bugs in specific scenarios
- Property tests verify universal invariants hold across all inputs

### Test Infrastructure

- Use Go's built-in `testing` package for unit tests
- Use `gopter` or `rapid` for property-based testing
- Mock external dependencies (filesystem, clipboard, command execution) for isolated testing
- Use table-driven tests for unit test organization
- Implement custom generators for property tests (natural language inputs, commands, history entries)
