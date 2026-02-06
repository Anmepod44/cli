package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/cli-command-assistant/internal/types"
	"github.com/fatih/color"
)

// displayWelcomeImproved shows an enhanced welcome message
func (c *CLI) displayWelcomeImproved() {
	// Clear screen for clean start
	fmt.Print("\033[H\033[2J")

	// Animated banner
	banner := []string{
		"╔════════════════════════════════════════════════════════════╗",
		"║                                                            ║",
		"║         🚀 CLI Command Assistant 🚀                        ║",
		"║                                                            ║",
		"║         Transform Natural Language → Linux Commands       ║",
		"║                                                            ║",
		"╚════════════════════════════════════════════════════════════╝",
	}

	for _, line := range banner {
		color.Cyan(line)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println()

	// Show generation mode with icon
	if c.config.UseLLM {
		color.HiGreen("  🤖 AI Mode Active")
		color.Green("     Powered by ChatGPT for advanced understanding")
	} else {
		color.HiYellow("  📋 Pattern Mode Active")
		color.Yellow("     Fast offline command generation")
		fmt.Println()
		color.HiBlue("     💡 Tip: Enable AI mode for smarter results")
		color.Blue("        Run: ./setup.sh")
	}

	fmt.Println()
	fmt.Println(strings.Repeat("─", 62))
	fmt.Println()

	color.HiWhite("  ✨ Just describe what you want to do in plain English")
	fmt.Println()

	color.HiYellow("  📚 Examples:")
	examples := []string{
		"list all files",
		"find files larger than 10MB",
		"show disk usage",
		"find process nginx",
		"show me all log files from last week",
	}
	for _, ex := range examples {
		fmt.Printf("     • %s\n", color.HiWhiteString(ex))
		time.Sleep(30 * time.Millisecond)
	}

	fmt.Println()
	color.HiYellow("  🔧 Special Commands:")
	fmt.Printf("     • %s - Show help\n", color.CyanString("help"))
	fmt.Printf("     • %s - View command history\n", color.CyanString("history"))
	fmt.Printf("     • %s - Exit application\n", color.CyanString("exit"))

	fmt.Println()
	fmt.Println(strings.Repeat("─", 62))
	fmt.Println()
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

// DisplayCommandImproved shows a generated command with enhanced formatting
func (c *CLI) DisplayCommandImproved(cmd *types.Command) {
	fmt.Println()
	fmt.Println(strings.Repeat("─", 62))
	fmt.Println()

	color.HiGreen("  ✓ Command Generated Successfully!")
	fmt.Println()

	// Command box
	color.HiCyan("  ┌─ Command ─────────────────────────────────────────────────┐")
	color.HiWhite("  │ %s", cmd.Raw)
	color.HiCyan("  └───────────────────────────────────────────────────────────┘")
	fmt.Println()

	if cmd.Description != "" {
		color.Yellow("  📝 What it does:")
		fmt.Printf("     %s\n", cmd.Description)
		fmt.Println()
	}

	if len(cmd.Flags) > 0 {
		color.Yellow("  🏴 Flags explained:")
		for _, flag := range cmd.Flags {
			fmt.Printf("     %s → %s\n",
				color.CyanString(flag.Name),
				flag.Description)
		}
		fmt.Println()
	}

	fmt.Println(strings.Repeat("─", 62))
}
