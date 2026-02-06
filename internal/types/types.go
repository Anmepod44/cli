package types

import "time"

// WarningLevel represents the severity of a warning
type WarningLevel int

const (
	InfoLevel WarningLevel = iota
	WarningLevelWarning
	CriticalLevel
)

// Command represents a shell command with metadata
type Command struct {
	Raw          string // The shell command string
	Description  string // Human-readable explanation
	Flags        []Flag // Parsed flags and options
	IsDangerous  bool   // Whether command is potentially destructive
	RequiresSudo bool   // Whether command needs elevated privileges
}

// Flag represents a command flag or option
type Flag struct {
	Name        string
	Description string
}

// Warning represents a safety warning about a command
type Warning struct {
	Level   WarningLevel
	Message string
}

// ExecutionResult represents the result of command execution
type ExecutionResult struct {
	CommandID string
	Output    string
	Error     error
	ExitCode  int
	Duration  time.Duration
}

// HistoryEntry represents a command in history
type HistoryEntry struct {
	ID        string
	Command   *Command
	Timestamp time.Time
	Executed  bool
}

// Config represents application configuration
type Config struct {
	AutoExecute      bool
	MaxConcurrent    int
	HistorySize      int
	ConfirmDangerous bool
	ClipboardEnabled bool
	UseLLM           bool   // Enable LLM-powered generation
	OpenAIAPIKey     string // OpenAI API key (can also use OPENAI_API_KEY env var)
}
