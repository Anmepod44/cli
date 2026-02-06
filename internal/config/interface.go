package config

import "github.com/cli-command-assistant/internal/types"

// ConfigManager manages application configuration
type ConfigManager interface {
	// Load reads configuration from file
	Load() (*types.Config, error)

	// Save writes configuration to file
	Save(cfg *types.Config) error
}
