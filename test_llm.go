package main

import (
	"fmt"
	"os"

	"github.com/cli-command-assistant/internal/generator"
)

func main() {
	// Test with environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		fmt.Println("⚠ No OPENAI_API_KEY environment variable set")
		fmt.Println("Testing fallback to pattern mode...")
	} else {
		fmt.Println("✓ Found OPENAI_API_KEY")
		fmt.Printf("  Key starts with: %s...\n", apiKey[:7])
	}

	// Create LLM generator
	gen := generator.NewLLMGenerator(apiKey)

	// Test command
	testInput := "list all files"
	fmt.Printf("\nTesting: '%s'\n", testInput)

	cmd, err := gen.Generate(testInput)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	fmt.Println("\n✓ Command generated successfully!")
	fmt.Printf("  Command: %s\n", cmd.Raw)
	fmt.Printf("  Description: %s\n", cmd.Description)
	fmt.Printf("  Dangerous: %v\n", cmd.IsDangerous)
	fmt.Printf("  Requires Sudo: %v\n", cmd.RequiresSudo)
}
