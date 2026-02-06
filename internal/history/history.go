package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/cli-command-assistant/internal/types"
	"github.com/google/uuid"
)

// Manager implements HistoryManager interface
type Manager struct {
	entries  []*types.HistoryEntry
	maxSize  int
	mu       sync.RWMutex
	filePath string
}

// NewManager creates a new history manager
func NewManager(maxSize int, filePath string) *Manager {
	if maxSize <= 0 {
		maxSize = 1000 // default
	}

	return &Manager{
		entries:  make([]*types.HistoryEntry, 0),
		maxSize:  maxSize,
		filePath: filePath,
	}
}

// Add stores a command in history
func (m *Manager) Add(cmd *types.Command, executed bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Create history entry
	entry := &types.HistoryEntry{
		ID:        uuid.New().String(),
		Command:   cmd,
		Timestamp: time.Now(),
		Executed:  executed,
	}

	// Add to beginning (most recent first)
	m.entries = append([]*types.HistoryEntry{entry}, m.entries...)

	// Enforce size limit (circular buffer behavior)
	if len(m.entries) > m.maxSize {
		m.entries = m.entries[:m.maxSize]
	}

	return nil
}

// List retrieves recent commands with optional limit
func (m *Manager) List(limit int) ([]*types.HistoryEntry, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// If limit is 0 or negative, return all
	if limit <= 0 {
		limit = len(m.entries)
	}

	// Return up to limit entries
	if limit > len(m.entries) {
		limit = len(m.entries)
	}

	// Create a copy to avoid external modification
	result := make([]*types.HistoryEntry, limit)
	copy(result, m.entries[:limit])

	return result, nil
}

// Search finds commands matching a query
func (m *Manager) Search(query string) ([]*types.HistoryEntry, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	query = strings.ToLower(query)
	var results []*types.HistoryEntry

	for _, entry := range m.entries {
		// Search in command raw string
		if strings.Contains(strings.ToLower(entry.Command.Raw), query) {
			results = append(results, entry)
			continue
		}

		// Search in command description
		if strings.Contains(strings.ToLower(entry.Command.Description), query) {
			results = append(results, entry)
		}
	}

	return results, nil
}

// Clear removes all history
func (m *Manager) Clear() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.entries = make([]*types.HistoryEntry, 0)

	return nil
}

// Save persists history to disk
func (m *Manager) Save() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.filePath == "" {
		return fmt.Errorf("no file path configured for history persistence")
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(m.entries, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal history: %w", err)
	}

	// Ensure directory exists
	dir := filepath.Dir(m.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create history directory: %w", err)
	}

	// Write to file
	if err := os.WriteFile(m.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write history file: %w", err)
	}

	return nil
}

// Load reads history from disk
func (m *Manager) Load() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.filePath == "" {
		return fmt.Errorf("no file path configured for history persistence")
	}

	// Check if file exists
	if _, err := os.Stat(m.filePath); os.IsNotExist(err) {
		// File doesn't exist yet, that's okay
		return nil
	}

	// Read file
	data, err := os.ReadFile(m.filePath)
	if err != nil {
		return fmt.Errorf("failed to read history file: %w", err)
	}

	// Unmarshal JSON
	var entries []*types.HistoryEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		// History file is corrupted, backup and create new
		backupPath := m.filePath + ".backup." + time.Now().Format("20060102-150405")
		if copyErr := os.WriteFile(backupPath, data, 0644); copyErr == nil {
			// Successfully backed up corrupted file
			m.entries = make([]*types.HistoryEntry, 0)
			return fmt.Errorf("history file corrupted, backed up to %s: %w", backupPath, err)
		}
		return fmt.Errorf("failed to unmarshal history (backup failed): %w", err)
	}

	m.entries = entries

	// Enforce size limit
	if len(m.entries) > m.maxSize {
		m.entries = m.entries[:m.maxSize]
	}

	return nil
}
