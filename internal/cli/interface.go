package cli

import "github.com/cli-command-assistant/internal/types"

// CLIInterface provides an interactive terminal interface
type CLIInterface interface {
	// Start begins the interactive session
	Start() error

	// DisplayCommand shows a generated command with formatting
	DisplayCommand(cmd *types.Command)

	// Prompt asks the user for input with options
	Prompt(message string, options []string) (string, error)

	// DisplayResult shows command execution results
	DisplayResult(result types.ExecutionResult)
}
