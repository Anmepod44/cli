package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cli-command-assistant/internal/executor"
	"github.com/cli-command-assistant/internal/generator"
	"github.com/cli-command-assistant/internal/history"
	"github.com/cli-command-assistant/internal/logger"
	"github.com/cli-command-assistant/internal/types"
	"github.com/fatih/color"
)

// CLI implements CLIInterface
type CLI struct {
	generator generator.CommandGenerator
	executor  *executor.Executor
	history   *history.Manager
	logger    *logger.Logger
	clipboard *Clipboard
	config    *types.Config
	reader    *bufio.Reader
}

// NewCLI creates a new CLI interface
func NewCLI(
	gen generator.CommandGenerator,
	exec *executor.Executor,
	hist *history.Manager,
	log *logger.Logger,
	cfg *types.Config,
) *CLI {
	return &CLI{
		generator: gen,
		executor:  exec,
		history:   hist,
		logger:    log,
		clipboard: NewClipboard(),
		config:    cfg,
		reader:    bufio.NewReader(os.Stdin),
	}
}

// Start begins the interactive session
func (c *CLI) Start() error {
	// Display welcome message
	c.displayWelcome()

	// Main loop
	for {
		// Display prompt
		fmt.Print("\n> ")

		// Read input
		input, err := c.reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.TrimSpace(input)

		// Handle empty input
		if input == "" {
			continue
		}

		// Handle special commands
		if c.handleSpecialCommand(input) {
			continue
		}

		// Check for exit commands
		if input == "exit" || input == "quit" {
			fmt.Println("\nGoodbye!")
			return nil
		}

		// Generate command with loading animation
		cmd, err := c.generateWithAnimation(input)
		if err != nil {
			color.Red("✗ Error: %v", err)
			c.logger.LogError("command generation", err)
			continue
		}

		// Validate command
		warnings := c.generator.Validate(cmd)

		// Display command
		c.DisplayCommand(cmd)

		// Display warnings
		for _, warning := range warnings {
			c.displayWarning(warning)
		}

		// Add to history
		if err := c.history.Add(cmd, false); err != nil {
			c.logger.LogError("history add", err)
		}

		// Prompt for action
		action := c.promptAction(cmd)

		switch action {
		case "execute":
			c.executeCommand(cmd)
		case "copy":
			c.copyCommand(cmd)
		case "discard":
			fmt.Println("Command discarded.")
		}
	}
}

// displayWelcome shows the welcome message
func (c *CLI) displayWelcome() {
	color.Cyan("╔════════════════════════════════════════════════════════════╗")
	color.Cyan("║         CLI Command Assistant - Natural Language           ║")
	color.Cyan("║              Linux Command Generator                       ║")
	color.Cyan("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Show generation mode
	if c.config.UseLLM {
		color.Green("🤖 AI Mode: Using LLM-powered command generation")
	} else {
		color.Yellow("📋 Pattern Mode: Using rule-based command generation")
		color.Blue("   Tip: Enable LLM mode in config for better understanding")
	}
	fmt.Println()

	fmt.Println("Welcome! Describe what you want to do in plain English.")
	fmt.Println()
	color.Yellow("Examples:")
	fmt.Println("  • list all files")
	fmt.Println("  • find files larger than 10MB")
	fmt.Println("  • show disk usage")
	fmt.Println("  • find process nginx")
	fmt.Println()
	color.Yellow("Special commands:")
	fmt.Println("  • help or ? - Show this help message")
	fmt.Println("  • history - Show command history")
	fmt.Println("  • exit or quit - Exit the application")
	fmt.Println()
}

// DisplayCommand shows a generated command with formatting
func (c *CLI) DisplayCommand(cmd *types.Command) {
	fmt.Println()
	color.Green("Generated Command:")
	color.Cyan("  %s", cmd.Raw)
	fmt.Println()

	if cmd.Description != "" {
		color.Yellow("Description:")
		fmt.Printf("  %s\n", cmd.Description)
		fmt.Println()
	}

	if len(cmd.Flags) > 0 {
		color.Yellow("Flags:")
		for _, flag := range cmd.Flags {
			fmt.Printf("  %s - %s\n", flag.Name, flag.Description)
		}
		fmt.Println()
	}
}

// displayWarning shows a warning message
func (c *CLI) displayWarning(warning types.Warning) {
	switch warning.Level {
	case types.CriticalLevel:
		color.Red("⚠ CRITICAL: %s", warning.Message)
	case types.WarningLevelWarning:
		color.Yellow("⚠ WARNING: %s", warning.Message)
	case types.InfoLevel:
		color.Blue("ℹ INFO: %s", warning.Message)
	}
}

// Prompt asks the user for input with options
func (c *CLI) Prompt(message string, options []string) (string, error) {
	fmt.Printf("%s\n", message)
	for i, opt := range options {
		fmt.Printf("  %d. %s\n", i+1, opt)
	}
	fmt.Print("Choice: ")

	input, err := c.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// DisplayResult shows command execution results
func (c *CLI) DisplayResult(result types.ExecutionResult) {
	fmt.Println()
	if result.Error != nil {
		color.Red("Command failed (exit code %d):", result.ExitCode)
		fmt.Println(result.Output)
	} else {
		color.Green("Command executed successfully:")
		fmt.Println(result.Output)
	}
	color.Blue("Duration: %v", result.Duration)
}

// handleSpecialCommand handles special commands like help and history
func (c *CLI) handleSpecialCommand(input string) bool {
	switch input {
	case "help", "?":
		c.displayWelcome()
		return true

	case "history":
		c.displayHistory()
		return true

	default:
		return false
	}
}

// displayHistory shows command history
func (c *CLI) displayHistory() {
	entries, err := c.history.List(20)
	if err != nil {
		color.Red("Error retrieving history: %v", err)
		return
	}

	if len(entries) == 0 {
		fmt.Println("No command history yet.")
		return
	}

	fmt.Println()
	color.Yellow("Recent Commands:")
	for i, entry := range entries {
		executed := ""
		if entry.Executed {
			executed = color.GreenString(" [executed]")
		}
		fmt.Printf("  %d. %s%s\n", i+1, entry.Command.Raw, executed)
		fmt.Printf("     %s\n", entry.Timestamp.Format("2006-01-02 15:04:05"))
	}
	fmt.Println()
}

// promptAction asks the user what to do with the command
func (c *CLI) promptAction(cmd *types.Command) string {
	fmt.Println()
	color.Yellow("What would you like to do?")
	fmt.Println("  1. Execute")
	fmt.Println("  2. Copy to clipboard")
	fmt.Println("  3. Discard")
	fmt.Print("Choice (1-3): ")

	input, err := c.reader.ReadString('\n')
	if err != nil {
		return "discard"
	}

	choice := strings.TrimSpace(input)

	switch choice {
	case "1", "execute", "e":
		return "execute"
	case "2", "copy", "c":
		return "copy"
	default:
		return "discard"
	}
}

// executeCommand executes a command
func (c *CLI) executeCommand(cmd *types.Command) {
	// Confirm if dangerous
	if cmd.IsDangerous && c.config.ConfirmDangerous {
		fmt.Print("\n")
		color.Red("⚠ This command is potentially dangerous!")
		fmt.Print("Are you sure you want to execute it? (yes/no): ")

		input, err := c.reader.ReadString('\n')
		if err != nil || strings.ToLower(strings.TrimSpace(input)) != "yes" {
			fmt.Println("Execution cancelled.")
			return
		}
	}

	fmt.Println("\nExecuting command...")

	// Execute
	resultChan := c.executor.Execute(cmd)
	result := <-resultChan

	// Display result
	c.DisplayResult(result)

	// Update history
	if err := c.history.Add(cmd, true); err != nil {
		c.logger.LogError("history update", err)
	}
}

// copyCommand copies a command to clipboard
func (c *CLI) copyCommand(cmd *types.Command) {
	if !c.config.ClipboardEnabled {
		fmt.Println("Clipboard is disabled in configuration.")
		return
	}

	if !c.clipboard.IsAvailable() {
		color.Yellow("Clipboard unavailable. Install xclip with: sudo apt install xclip")
		fmt.Printf("\nCommand: %s\n", cmd.Raw)
		return
	}

	if err := c.clipboard.Copy(cmd.Raw); err != nil {
		color.Red("Failed to copy to clipboard: %v", err)
		c.logger.LogError("clipboard copy", err)
		fmt.Printf("\nCommand: %s\n", cmd.Raw)
		return
	}

	color.Green("✓ Command copied to clipboard!")
}

// generateWithAnimation shows a loading animation while generating command
func (c *CLI) generateWithAnimation(input string) (*types.Command, error) {
	// Channel to receive the result
	resultChan := make(chan struct {
		cmd *types.Command
		err error
	}, 1)

	// Generate command in background
	go func() {
		cmd, err := c.generator.Generate(input)
		resultChan <- struct {
			cmd *types.Command
			err error
		}{cmd, err}
	}()

	// Show loading animation
	spinners := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	messages := []string{
		"Analyzing your request",
		"Generating command",
		"Optimizing solution",
		"Almost there",
	}

	spinnerIdx := 0
	messageIdx := 0
	ticker := time.NewTicker(80 * time.Millisecond)
	defer ticker.Stop()

	messageTicker := time.NewTicker(1500 * time.Millisecond)
	defer messageTicker.Stop()

	fmt.Print("\n")

	for {
		select {
		case result := <-resultChan:
			// Clear the loading line
			fmt.Print("\r\033[K")
			return result.cmd, result.err

		case <-ticker.C:
			// Update spinner
			spinner := spinners[spinnerIdx%len(spinners)]
			message := messages[messageIdx%len(messages)]

			fmt.Printf("\r%s %s...",
				color.CyanString(spinner),
				color.YellowString(message))

			spinnerIdx++

		case <-messageTicker.C:
			// Change message
			messageIdx++
		}
	}
}
