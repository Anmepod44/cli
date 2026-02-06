package history

import "github.com/cli-command-assistant/internal/types"

// HistoryManager manages command history
type HistoryManager interface {
	// Add stores a command in history
	Add(cmd *types.Command, executed bool) error

	// List retrieves recent commands with optional limit
	List(limit int) ([]*types.HistoryEntry, error)

	// Search finds commands matching a query
	Search(query string) ([]*types.HistoryEntry, error)

	// Clear removes all history
	Clear() error
}
