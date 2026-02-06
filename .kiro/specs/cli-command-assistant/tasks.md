# Implementation Plan: CLI Command Assistant

## Overview

This implementation plan breaks down the CLI Command Assistant into discrete, incremental coding tasks. Each task builds on previous work, starting with core data structures and interfaces, then implementing each major component, and finally integrating everything into a working CLI application. The plan includes property-based tests and unit tests as sub-tasks to validate correctness early and often.

## Tasks

- [x] 1. Set up project structure and core data models
  - Create Go module with `go mod init`
  - Define directory structure: `cmd/`, `internal/generator/`, `internal/executor/`, `internal/history/`, `internal/config/`, `internal/cli/`
  - Create core data structures: `Command`, `Flag`, `Warning`, `ExecutionResult`, `HistoryEntry`, `Config`
  - Define interfaces: `CommandGenerator`, `CommandExecutor`, `HistoryManager`, `ConfigManager`, `CLIInterface`
  - _Requirements: 1.1, 3.2, 5.1, 9.1_

- [ ] 2. Implement Command Generator component
  - [x] 2.1 Create pattern matching system for natural language translation
    - Implement regex-based pattern matcher
    - Create command template system with parameter substitution
    - Build mapping of common phrases to command templates (ls, find, grep, ps, df, etc.)
    - _Requirements: 1.1, 1.4, 1.5_

  - [ ]* 2.2 Write property test for valid command generation
    - **Property 1: Valid Command Generation**
    - **Validates: Requirements 1.1, 2.2**

  - [x] 2.3 Implement command safety classifier
    - Create rules for identifying dangerous operations (rm, dd, mkfs, etc.)
    - Detect file modification and deletion commands
    - Identify system directory targets (/etc, /sys, /proc, /boot)
    - Detect sudo requirements
    - _Requirements: 7.1, 7.2, 7.3, 7.4_

  - [ ]* 2.4 Write property tests for safety classification
    - **Property 11: Command Safety Classification**
    - **Validates: Requirements 7.1, 7.3, 7.4**
    - **Property 12: Privilege Detection**
    - **Validates: Requirements 7.2**

  - [x] 2.5 Implement error handling for invalid inputs
    - Return descriptive errors for unmappable inputs
    - Handle empty and malformed inputs
    - _Requirements: 1.3_

  - [ ]* 2.6 Write property test for error handling
    - **Property 2: Error Handling for Invalid Input**
    - **Validates: Requirements 1.3**

  - [ ]* 2.7 Write unit tests for command generator
    - Test specific natural language patterns (list files, find large files, kill process)
    - Test dangerous command detection (rm -rf, dd if=/dev/zero)
    - Test sudo detection edge cases
    - _Requirements: 1.1, 1.4, 1.5, 7.1, 7.2_

- [ ] 3. Implement Command Executor component
  - [x] 3.1 Create command execution with os/exec
    - Implement Execute method using os/exec.Command
    - Capture stdout and stderr streams
    - Record exit codes and execution duration
    - Return results via channel
    - _Requirements: 3.2, 3.3, 3.4_

  - [ ]* 3.2 Write property test for execution result capture
    - **Property 5: Complete Execution Result Capture**
    - **Validates: Requirements 3.3, 3.4**

  - [x] 3.3 Implement worker pool for concurrent execution
    - Create worker pool with configurable concurrency limit
    - Use goroutines and channels for async execution
    - Implement command cancellation via context
    - _Requirements: 6.1, 6.2, 6.4_

  - [ ]* 3.4 Write property test for output isolation
    - **Property 10: Concurrent Output Isolation**
    - **Validates: Requirements 6.2**

  - [ ]* 3.5 Write unit tests for command executor
    - Test successful command execution (echo, ls)
    - Test failed command execution (invalid command)
    - Test command cancellation
    - Test concurrent execution of multiple commands
    - _Requirements: 3.2, 3.3, 3.4, 6.2_

- [ ] 4. Checkpoint - Ensure all tests pass
  - Ensure all tests pass, ask the user if questions arise.

- [ ] 5. Implement History Manager component
  - [x] 5.1 Create in-memory history storage
    - Implement Add method with timestamp generation
    - Implement List method with limit support
    - Implement Search method with substring matching
    - Implement circular buffer for size limits
    - _Requirements: 5.1, 5.2, 5.5, 9.5_

  - [ ]* 5.2 Write property tests for history operations
    - **Property 8: History Chronological Ordering**
    - **Validates: Requirements 5.2**
    - **Property 9: History Search Filtering**
    - **Validates: Requirements 5.5**
    - **Property 14: History Size Limit Enforcement**
    - **Validates: Requirements 9.5**

  - [x] 5.3 Implement history persistence to JSON file
    - Save history to ~/.cli-assistant/history.json
    - Load history from file on startup
    - Handle corrupted history files gracefully
    - _Requirements: 5.4, 10.4_

  - [ ]* 5.4 Write property test for history persistence
    - **Property 7: History Persistence Round-Trip**
    - **Validates: Requirements 5.1, 5.4**

  - [ ]* 5.5 Write unit tests for history manager
    - Test adding and retrieving entries
    - Test empty history state
    - Test search with no results
    - Test history file corruption recovery
    - _Requirements: 5.1, 5.2, 5.4, 5.5, 10.4_

- [ ] 6. Implement Configuration Manager component
  - [x] 6.1 Create configuration loading and saving
    - Define default configuration values
    - Implement Load method to read from ~/.cli-assistant/config.yaml
    - Implement Save method to write configuration
    - Handle missing configuration files (use defaults)
    - Validate configuration values and clamp to acceptable ranges
    - _Requirements: 9.1, 9.2, 9.3, 9.4, 9.5_

  - [ ]* 6.2 Write property test for configuration loading
    - **Property 13: Configuration Loading**
    - **Validates: Requirements 9.1**

  - [ ]* 6.3 Write unit tests for configuration manager
    - Test loading valid configuration
    - Test missing configuration (defaults)
    - Test invalid configuration values
    - _Requirements: 9.1, 9.2_

- [ ] 7. Implement error logging system
  - [x] 7.1 Create error logger
    - Implement logging to ~/.cli-assistant/errors.log
    - Include timestamps and error details
    - Ensure all error paths log appropriately
    - _Requirements: 10.5_

  - [ ]* 7.2 Write property test for error logging
    - **Property 15: Error Logging Completeness**
    - **Validates: Requirements 10.5**

- [ ] 8. Implement clipboard integration
  - [x] 8.1 Create clipboard operations
    - Implement copy to clipboard using xclip or xsel for Linux
    - Handle clipboard unavailability gracefully
    - Return confirmation on successful copy
    - _Requirements: 4.1, 4.2, 10.3_

  - [ ]* 8.2 Write property test for clipboard operations
    - **Property 6: Clipboard Operation Success**
    - **Validates: Requirements 4.1**

  - [ ]* 8.3 Write unit tests for clipboard
    - Test successful copy operation
    - Test clipboard unavailable error handling
    - _Requirements: 4.1, 4.2, 10.3_

- [ ] 9. Checkpoint - Ensure all tests pass
  - Ensure all tests pass, ask the user if questions arise.

- [ ] 10. Implement CLI Interface component
  - [x] 10.1 Create interactive CLI using promptui or bubbletea
    - Implement Start method with main loop
    - Display welcome message and instructions on startup
    - Implement input prompt for natural language commands
    - _Requirements: 8.1, 8.2_

  - [x] 10.2 Implement command display and user options
    - Display generated commands with formatting
    - Show command description and flag explanations
    - Present options: execute, copy, discard
    - Implement confirmation prompt for execution
    - _Requirements: 2.1, 2.3, 2.4, 3.1_

  - [ ]* 10.3 Write property test for execution confirmation
    - **Property 4: Execution Confirmation Required**
    - **Validates: Requirements 3.1**

  - [x] 10.4 Implement special commands (help, history, exit)
    - Handle "help" and "?" to display usage instructions
    - Handle "history" to display command history
    - Handle "exit" and "quit" to gracefully terminate
    - _Requirements: 8.5, 8.6_

  - [ ]* 10.5 Write unit tests for CLI interface
    - Test help command display
    - Test exit command handling
    - Test startup welcome message
    - _Requirements: 8.1, 8.5, 8.6_

- [ ] 11. Implement main application orchestration
  - [x] 11.1 Create main application entry point
    - Initialize all components (generator, executor, history, config, CLI)
    - Wire components together with dependency injection
    - Load configuration on startup
    - Load history on startup
    - Start CLI interface
    - _Requirements: 9.1, 9.2, 5.4_

  - [x] 11.2 Implement graceful shutdown
    - Handle interrupt signals (SIGINT, SIGTERM)
    - Wait for running commands to complete or offer cancellation
    - Save history before exit
    - _Requirements: 6.5_

  - [ ]* 11.3 Write integration tests
    - Test end-to-end workflow: input → generate → execute → history
    - Test configuration loading and application behavior
    - Test graceful shutdown with running commands
    - _Requirements: 1.1, 3.2, 5.1, 6.5, 9.1_

- [ ] 12. Add command template library
  - [ ] 12.1 Expand natural language pattern coverage
    - Add more file operation patterns (copy, move, delete, search)
    - Add more system operation patterns (network, disk, memory, processes)
    - Add parameter extraction for sizes, paths, names, etc.
    - _Requirements: 1.4, 1.5_

  - [ ]* 12.2 Write unit tests for expanded patterns
    - Test file operations (copy files, move directories, delete patterns)
    - Test system operations (check disk space, list processes, network status)
    - _Requirements: 1.4, 1.5_

- [ ] 13. Final checkpoint - Ensure all tests pass
  - Ensure all tests pass, ask the user if questions arise.

- [x] 14. Create README and usage documentation
  - Write README.md with project overview, installation, and usage examples
  - Document configuration options
  - Document supported natural language patterns
  - Add examples of common use cases
  - _Requirements: 8.1, 9.1_

## Notes

- Tasks marked with `*` are optional and can be skipped for faster MVP
- Each task references specific requirements for traceability
- Checkpoints ensure incremental validation at reasonable breaks
- Property tests validate universal correctness properties across randomized inputs
- Unit tests validate specific examples, edge cases, and integration points
- All components should be implemented with error handling from the start
- Use Go's built-in `testing` package and `gopter` or `rapid` for property-based testing
