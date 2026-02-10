package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/cli-command-assistant/internal/cli"
	"github.com/cli-command-assistant/internal/config"
	"github.com/cli-command-assistant/internal/executor"
	"github.com/cli-command-assistant/internal/generator"
	"github.com/cli-command-assistant/internal/history"
	"github.com/cli-command-assistant/internal/logger"
)

func main() {
	// Parse command line flags
	quickMode := flag.Bool("quick", false, "Quick mode: read from stdin and output command only")
	flag.Parse()

	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to get home directory: %v\n", err)
		os.Exit(1)
	}

	// Setup paths
	appDir := filepath.Join(homeDir, ".cli-assistant")
	configPath := filepath.Join(appDir, "config.yaml")
	historyPath := filepath.Join(appDir, "history.json")
	logPath := filepath.Join(appDir, "errors.log")

	// Initialize logger
	log := logger.NewLogger(logPath)

	// Load configuration
	configMgr := config.NewManager(configPath)
	cfg, err := configMgr.Load()
	if err != nil {
		log.LogError("config load", err)
		fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
	}

	// Initialize components
	var gen generator.CommandGenerator
	if cfg.UseLLM {
		gen = generator.NewLLMGenerator(cfg.OpenAIAPIKey)
		log.LogInfo("initialization", "Using LLM-powered command generation")
	} else {
		gen = generator.NewGenerator()
		log.LogInfo("initialization", "Using pattern-based command generation")
	}
	exec := executor.NewExecutor(cfg.MaxConcurrent)
	hist := history.NewManager(cfg.HistorySize, historyPath)

	// Load history
	if err := hist.Load(); err != nil {
		log.LogError("history load", err)
		fmt.Fprintf(os.Stderr, "Warning: failed to load history: %v\n", err)
	}

	// Create CLI
	cliApp := cli.NewCLI(gen, exec, hist, log, cfg)

	// Handle quick mode
	if *quickMode {
		handleQuickMode(gen, log)
		return
	}

	// Setup graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\n\nShutting down gracefully...")

		// Save history
		if err := hist.Save(); err != nil {
			log.LogError("history save on shutdown", err)
			fmt.Fprintf(os.Stderr, "Warning: failed to save history: %v\n", err)
		}

		os.Exit(0)
	}()

	// Start CLI
	if err := cliApp.Start(); err != nil {
		log.LogError("cli start", err)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Save history on normal exit
	if err := hist.Save(); err != nil {
		log.LogError("history save on exit", err)
		fmt.Fprintf(os.Stderr, "Warning: failed to save history: %v\n", err)
	}
}

// handleQuickMode processes input from stdin and outputs command only
func handleQuickMode(gen generator.CommandGenerator, log *logger.Logger) {
	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder

	// Read all input from stdin
	for scanner.Scan() {
		input.WriteString(scanner.Text())
		input.WriteString(" ")
	}

	if err := scanner.Err(); err != nil {
		log.LogError("stdin read", err)
		os.Exit(1)
	}

	query := strings.TrimSpace(input.String())
	if query == "" {
		os.Exit(1)
	}

	// Generate command
	cmd, err := gen.Generate(query)
	if err != nil {
		log.LogError("quick mode generation", err)
		os.Exit(1)
	}

	// Output only the command
	fmt.Println(cmd.Raw)
}
