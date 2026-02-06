package generator

import "github.com/cli-command-assistant/internal/types"

// CommandGenerator translates natural language to shell commands
type CommandGenerator interface {
	// Generate translates natural language to a shell command
	Generate(input string) (*types.Command, error)

	// Validate checks if a command is safe to execute
	Validate(cmd *types.Command) []types.Warning
}
