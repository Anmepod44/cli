package executor

import "github.com/cli-command-assistant/internal/types"

// CommandExecutor executes shell commands
type CommandExecutor interface {
	// Execute runs a command asynchronously and returns a result channel
	Execute(cmd *types.Command) <-chan types.ExecutionResult

	// ExecuteBatch runs multiple commands concurrently
	ExecuteBatch(cmds []*types.Command) <-chan types.ExecutionResult

	// Cancel stops a running command by ID
	Cancel(id string) error
}
