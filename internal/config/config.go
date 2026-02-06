package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cli-command-assistant/internal/types"
	"gopkg.in/yaml.v3"
)

// Manager implements ConfigManager interface
type Manager struct {
	filePath string
}

// NewManager creates a new configuration manager
func NewManager(filePath string) *Manager {
	return &Manager{
		filePath: filePath,
	}
}

// DefaultConfig returns the default configuration
func DefaultConfig() *types.Config {
	return &types.Config{
		AutoExecute:      false,
		MaxConcurrent:    5,
		HistorySize:      1000,
		ConfirmDangerous: true,
		ClipboardEnabled: true,
		UseLLM:           false,
		OpenAIAPIKey:     "",
	}
}

// Load reads configuration from file
func (m *Manager) Load() (*types.Config, error) {
	// Start with defaults
	cfg := DefaultConfig()

	// Check if file exists
	if _, err := os.Stat(m.filePath); os.IsNotExist(err) {
		// File doesn't exist, return defaults
		return cfg, nil
	}

	// Read file
	data, err := os.ReadFile(m.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal YAML
	if err := yaml.Unmarshal(data, cfg); err != nil {
		// Config file is invalid, return defaults and log error
		return DefaultConfig(), fmt.Errorf("failed to parse config file (using defaults): %w", err)
	}

	// Validate and clamp values
	if cfg.MaxConcurrent <= 0 {
		cfg.MaxConcurrent = 5
	}
	if cfg.MaxConcurrent > 20 {
		cfg.MaxConcurrent = 20
	}

	if cfg.HistorySize <= 0 {
		cfg.HistorySize = 1000
	}
	if cfg.HistorySize > 10000 {
		cfg.HistorySize = 10000
	}

	return cfg, nil
}

// Save writes configuration to file
func (m *Manager) Save(cfg *types.Config) error {
	if cfg == nil {
		return fmt.Errorf("config cannot be nil")
	}

	// Validate and clamp values before saving
	if cfg.MaxConcurrent <= 0 {
		cfg.MaxConcurrent = 5
	}
	if cfg.MaxConcurrent > 20 {
		cfg.MaxConcurrent = 20
	}

	if cfg.HistorySize <= 0 {
		cfg.HistorySize = 1000
	}
	if cfg.HistorySize > 10000 {
		cfg.HistorySize = 10000
	}

	// Marshal to YAML
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Ensure directory exists
	dir := filepath.Dir(m.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Write to file
	if err := os.WriteFile(m.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
