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
		// Display modern prompt
		color.New(color.FgHiMagenta, color.Bold).Print("→ ")

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
			color.New(color.FgHiCyan).Println("\nGoodbye!")
			return nil
		}

		// Generate command with loading animation
		cmd, err := c.generateWithAnimation(input)
		if err != nil {
			color.New(color.FgHiRed, color.Bold).Printf("✗ Error: %v\n", err)
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
	// Modern gradient-style header with box drawing
	cyan := color.New(color.FgCyan, color.Bold)
	magenta := color.New(color.FgMagenta, color.Bold)

	cyan.Println("┌─────────────────────────────────────────────────────────────┐")
	cyan.Print("│ ")
	magenta.Print("CLI Command Assistant")
	cyan.Println("                                │")
	cyan.Print("│ ")
	color.New(color.FgHiWhite).Print("Natural Language → Linux Commands")
	cyan.Println("                     │")
	cyan.Println("└─────────────────────────────────────────────────────────────┘")
	fmt.Println()

	// Show generation mode with modern icons
	if c.config.UseLLM {
		color.New(color.FgHiGreen, color.Bold).Print("● ")
		color.HiWhite("AI Mode: LLM-powered command generation")
	} else {
		color.New(color.FgHiYellow, color.Bold).Print("● ")
		color.HiWhite("Pattern Mode: Rule-based command generation")
		color.New(color.FgHiBlack).Println("  → Enable LLM mode in config for better understanding")
	}
	fmt.Println()

	color.New(color.FgHiWhite).Println("Describe what you want to do in plain English.")
	fmt.Println()

	color.New(color.FgHiCyan, color.Bold).Println("Examples:")
	color.New(color.FgHiBlack).Println("  ▸ list all files")
	color.New(color.FgHiBlack).Println("  ▸ find files larger than 10MB")
	color.New(color.FgHiBlack).Println("  ▸ show disk usage")
	color.New(color.FgHiBlack).Println("  ▸ find process nginx")
	fmt.Println()

	color.New(color.FgHiMagenta, color.Bold).Println("Commands:")
	color.New(color.FgHiBlack).Println("  ▸ help or ? - Show this help")
	color.New(color.FgHiBlack).Println("  ▸ history - Show command history")
	color.New(color.FgHiBlack).Println("  ▸ exit or quit - Exit")
	fmt.Println()
}

// DisplayCommand shows a generated command with formatting
func (c *CLI) DisplayCommand(cmd *types.Command) {
	fmt.Println()

	// Modern command display with gradient-style colors
	color.New(color.FgHiGreen, color.Bold).Print("✓ ")
	color.New(color.FgHiWhite, color.Bold).Println("Generated Command")

	// Command in a subtle box
	color.New(color.FgHiBlack).Print("  ┌─ ")
	color.New(color.FgHiCyan, color.Bold).Println(cmd.Raw)
	color.New(color.FgHiBlack).Println("  └─")
	fmt.Println()

	if cmd.Description != "" {
		color.New(color.FgHiMagenta, color.Bold).Print("● ")
		color.New(color.FgHiWhite).Println("Description")
		color.New(color.FgHiBlack).Printf("  %s\n", cmd.Description)
		fmt.Println()
	}

	if len(cmd.Flags) > 0 {
		color.New(color.FgHiYellow, color.Bold).Print("● ")
		color.New(color.FgHiWhite).Println("Flags")
		for _, flag := range cmd.Flags {
			color.New(color.FgHiBlack).Printf("  ▸ ")
			color.New(color.FgHiCyan).Printf("%s", flag.Name)
			color.New(color.FgHiBlack).Printf(" - %s\n", flag.Description)
		}
		fmt.Println()
	}
}

// displayWarning shows a warning message
func (c *CLI) displayWarning(warning types.Warning) {
	switch warning.Level {
	case types.CriticalLevel:
		color.New(color.FgHiRed, color.Bold).Print("⚠ CRITICAL: ")
		color.HiWhite(warning.Message)
	case types.WarningLevelWarning:
		color.New(color.FgHiYellow, color.Bold).Print("⚠ WARNING: ")
		color.HiWhite(warning.Message)
	case types.InfoLevel:
		color.New(color.FgHiCyan, color.Bold).Print("ℹ INFO: ")
		color.HiWhite(warning.Message)
	}
}

// Prompt asks the user for input with options
func (c *CLI) Prompt(message string, options []string) (string, error) {
	color.New(color.FgHiWhite, color.Bold).Println(message)
	for i, opt := range options {
		color.New(color.FgHiBlack).Print("  ")
		color.New(color.FgHiCyan, color.Bold).Printf("%d", i+1)
		color.New(color.FgHiBlack).Print(". ")
		color.New(color.FgHiWhite).Println(opt)
	}
	color.New(color.FgHiMagenta, color.Bold).Print("→ ")

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
		color.New(color.FgHiRed, color.Bold).Printf("✗ Command failed (exit code %d)\n", result.ExitCode)
		color.New(color.FgHiBlack).Println(result.Output)
	} else {
		color.New(color.FgHiGreen, color.Bold).Println("✓ Command executed successfully")
		fmt.Println(result.Output)
	}
	color.New(color.FgHiBlack).Printf("Duration: %v\n", result.Duration)
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
		color.New(color.FgHiRed, color.Bold).Printf("✗ Error retrieving history: %v\n", err)
		return
	}

	if len(entries) == 0 {
		color.New(color.FgHiBlack).Println("No command history yet.")
		return
	}

	fmt.Println()
	color.New(color.FgHiCyan, color.Bold).Println("Recent Commands")
	color.New(color.FgHiBlack).Println("───────────────")

	for i, entry := range entries {
		color.New(color.FgHiBlack).Printf("%2d. ", i+1)
		color.New(color.FgHiWhite).Print(entry.Command.Raw)

		if entry.Executed {
			color.New(color.FgHiGreen, color.Bold).Print(" ✓")
		}
		fmt.Println()

		color.New(color.FgHiBlack).Printf("    %s\n", entry.Timestamp.Format("2006-01-02 15:04:05"))
	}
	fmt.Println()
}

// promptAction asks the user what to do with the command
func (c *CLI) promptAction(cmd *types.Command) string {
	fmt.Println()
	color.New(color.FgHiWhite, color.Bold).Println("What would you like to do?")

	color.New(color.FgHiBlack).Print("  ")
	color.New(color.FgHiGreen, color.Bold).Print("1")
	color.New(color.FgHiBlack).Print(". ")
	color.New(color.FgHiWhite).Println("Execute")

	color.New(color.FgHiBlack).Print("  ")
	color.New(color.FgHiCyan, color.Bold).Print("2")
	color.New(color.FgHiBlack).Print(". ")
	color.New(color.FgHiWhite).Println("Copy to clipboard")

	color.New(color.FgHiBlack).Print("  ")
	color.New(color.FgHiRed, color.Bold).Print("3")
	color.New(color.FgHiBlack).Print(". ")
	color.New(color.FgHiWhite).Println("Discard")

	color.New(color.FgHiMagenta, color.Bold).Print("→ ")

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
		color.New(color.FgHiRed, color.Bold).Println("⚠ This command is potentially dangerous!")
		color.New(color.FgHiYellow).Print("Are you sure you want to execute it? (yes/no): ")

		input, err := c.reader.ReadString('\n')
		if err != nil || strings.ToLower(strings.TrimSpace(input)) != "yes" {
			color.New(color.FgHiBlack).Println("Execution cancelled.")
			return
		}
	}

	color.New(color.FgHiCyan).Println("\nExecuting command...")

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
		color.New(color.FgHiBlack).Println("Clipboard is disabled in configuration.")
		return
	}

	if !c.clipboard.IsAvailable() {
		color.New(color.FgHiYellow).Println("Clipboard unavailable. Install xclip with: sudo apt install xclip")
		fmt.Printf("\nCommand: %s\n", cmd.Raw)
		return
	}

	if err := c.clipboard.Copy(cmd.Raw); err != nil {
		color.New(color.FgHiRed, color.Bold).Printf("✗ Failed to copy to clipboard: %v\n", err)
		c.logger.LogError("clipboard copy", err)
		fmt.Printf("\nCommand: %s\n", cmd.Raw)
		return
	}

	color.New(color.FgHiGreen, color.Bold).Println("✓ Command copied to clipboard!")
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

	// Modern loading animation with dots and gradient colors
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
			// Update spinner with modern colors
			spinner := spinners[spinnerIdx%len(spinners)]
			message := messages[messageIdx%len(messages)]

			// Create gradient effect with cyan to magenta
			fmt.Printf("\r%s %s...",
				color.New(color.FgHiCyan, color.Bold).Sprint(spinner),
				color.New(color.FgHiWhite).Sprint(message))

			spinnerIdx++

		case <-messageTicker.C:
			// Change message
			messageIdx++
		}
	}
}
