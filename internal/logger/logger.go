package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Logger handles error logging to file
type Logger struct {
	filePath string
	mu       sync.Mutex
}

// NewLogger creates a new error logger
func NewLogger(filePath string) *Logger {
	return &Logger{
		filePath: filePath,
	}
}

// LogError writes an error to the log file
func (l *Logger) LogError(context string, err error) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Ensure directory exists
	dir := filepath.Dir(l.filePath)
	if mkdirErr := os.MkdirAll(dir, 0755); mkdirErr != nil {
		return fmt.Errorf("failed to create log directory: %w", mkdirErr)
	}

	// Open file in append mode
	file, openErr := os.OpenFile(l.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if openErr != nil {
		return fmt.Errorf("failed to open log file: %w", openErr)
	}
	defer file.Close()

	// Format log entry
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s: %v\n", timestamp, context, err)

	// Write to file
	if _, writeErr := file.WriteString(logEntry); writeErr != nil {
		return fmt.Errorf("failed to write to log file: %w", writeErr)
	}

	return nil
}

// LogInfo writes an informational message to the log file
func (l *Logger) LogInfo(context string, message string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Ensure directory exists
	dir := filepath.Dir(l.filePath)
	if mkdirErr := os.MkdirAll(dir, 0755); mkdirErr != nil {
		return fmt.Errorf("failed to create log directory: %w", mkdirErr)
	}

	// Open file in append mode
	file, openErr := os.OpenFile(l.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if openErr != nil {
		return fmt.Errorf("failed to open log file: %w", openErr)
	}
	defer file.Close()

	// Format log entry
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] INFO - %s: %s\n", timestamp, context, message)

	// Write to file
	if _, writeErr := file.WriteString(logEntry); writeErr != nil {
		return fmt.Errorf("failed to write to log file: %w", writeErr)
	}

	return nil
}
