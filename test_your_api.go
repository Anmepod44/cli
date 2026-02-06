package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cli-command-assistant/internal/config"
	"github.com/cli-command-assistant/internal/generator"
)

func main() {
	// Load your actual config
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".cli-assistant", "config.yaml")

	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║          Testing Your OpenAI API Integration              ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Load config
	configMgr := config.NewManager(configPath)
	cfg, err := configMgr.Load()
	if err != nil {
		fmt.Printf("❌ Error loading config: %v\n", err)
		return
	}

	fmt.Printf("Config loaded from: %s\n", configPath)
	fmt.Printf("  LLM Mode: %v\n", cfg.UseLLM)
	if cfg.OpenAIAPIKey != "" {
		fmt.Printf("  API Key: %s...%s\n", cfg.OpenAIAPIKey[:10], cfg.OpenAIAPIKey[len(cfg.OpenAIAPIKey)-4:])
	}
	fmt.Println()

	// Create generator
	var gen generator.CommandGenerator
	if cfg.UseLLM {
		gen = generator.NewLLMGenerator(cfg.OpenAIAPIKey)
		fmt.Println("🤖 Using LLM-powered generation")
	} else {
		gen = generator.NewGenerator()
		fmt.Println("📋 Using pattern-based generation")
	}
	fmt.Println()

	// Test with a complex query that LLM should handle better
	testQuery := "show me all log files modified in the last week"
	fmt.Printf("Testing query: '%s'\n", testQuery)
	fmt.Println("Calling OpenAI API...")
	fmt.Println()

	cmd, err := gen.Generate(testQuery)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		fmt.Println()
		fmt.Println("This might mean:")
		fmt.Println("  - API key is invalid")
		fmt.Println("  - No internet connection")
		fmt.Println("  - API quota exceeded")
		fmt.Println("  - Falling back to pattern mode")
		return
	}

	fmt.Println("✅ SUCCESS! LLM generated command:")
	fmt.Println()
	fmt.Printf("  Command: %s\n", cmd.Raw)
	fmt.Printf("  Description: %s\n", cmd.Description)
	fmt.Printf("  Dangerous: %v\n", cmd.IsDangerous)
	fmt.Printf("  Requires Sudo: %v\n", cmd.RequiresSudo)
	fmt.Println()
	fmt.Println("🎉 Your OpenAI integration is working!")
}
