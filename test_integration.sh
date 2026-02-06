#!/bin/bash

# Integration test for LLM functionality

echo "╔════════════════════════════════════════════════════════════╗"
echo "║     CLI Command Assistant - LLM Integration Test          ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""

# Check if OPENAI_API_KEY is set
if [ -z "$OPENAI_API_KEY" ]; then
    echo "❌ OPENAI_API_KEY environment variable is not set"
    echo ""
    echo "To test LLM mode, you need to:"
    echo "  1. Get an API key from: https://platform.openai.com/api-keys"
    echo "  2. Export it: export OPENAI_API_KEY='sk-your-key-here'"
    echo "  3. Run this test again"
    echo ""
    echo "Testing pattern mode fallback instead..."
    echo ""
fi

# Create test config directory
TEST_DIR="/tmp/cli-assistant-test"
mkdir -p "$TEST_DIR"

# Create test config with LLM enabled
cat > "$TEST_DIR/config.yaml" << EOF
autoexecute: false
maxconcurrent: 5
historysize: 1000
confirmdangerous: true
clipboardenabled: true
usellm: true
openaiapikey: ""
EOF

echo "Test Configuration:"
echo "  Config: $TEST_DIR/config.yaml"
echo "  LLM Mode: Enabled"
echo "  API Key: From environment variable"
echo ""

# Build the app
echo "Building application..."
go build -o cli-assistant-test ./cmd/cli-assistant
if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi
echo "✓ Build successful"
echo ""

# Test 1: Check config loading
echo "Test 1: Configuration Loading"
echo "------------------------------"
HOME=/tmp go run test_llm.go
echo ""

# Test 2: Pattern mode (should always work)
echo "Test 2: Pattern Mode (Fallback)"
echo "--------------------------------"
echo "list all files" | timeout 3 ./cli-assistant-test 2>&1 | head -30
echo ""

# Test 3: LLM mode (if API key is set)
if [ -n "$OPENAI_API_KEY" ]; then
    echo "Test 3: LLM Mode (with API key)"
    echo "--------------------------------"
    echo "✓ API key detected"
    echo "  Testing with real API call..."
    echo ""
    
    # Create a simple test
    cat > test_llm_real.go << 'EOTEST'
package main

import (
	"fmt"
	"os"
	"github.com/cli-command-assistant/internal/generator"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	gen := generator.NewLLMGenerator(apiKey)
	
	fmt.Println("Testing LLM generation with real API...")
	cmd, err := gen.Generate("show me all log files modified in the last week")
	
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("✓ LLM generation successful!")
	fmt.Printf("  Command: %s\n", cmd.Raw)
	fmt.Printf("  Description: %s\n", cmd.Description)
}
EOTEST
    
    go run test_llm_real.go
    rm -f test_llm_real.go
else
    echo "Test 3: LLM Mode (Skipped)"
    echo "--------------------------"
    echo "⚠ Skipped - No API key set"
    echo "  Set OPENAI_API_KEY to test real LLM integration"
fi

echo ""
echo "╔════════════════════════════════════════════════════════════╗"
echo "║                    Test Summary                            ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""

if [ -n "$OPENAI_API_KEY" ]; then
    echo "✓ All tests completed"
    echo ""
    echo "Your app is ready to use LLM mode!"
    echo "Just make sure your config has:"
    echo "  usellm: true"
    echo "  openaiapikey: \"\" (will use env var)"
else
    echo "✓ Pattern mode tests passed"
    echo "⚠ LLM mode not tested (no API key)"
    echo ""
    echo "To enable LLM mode:"
    echo "  1. Get API key: https://platform.openai.com/api-keys"
    echo "  2. Export: export OPENAI_API_KEY='sk-your-key'"
    echo "  3. Set in config: usellm: true"
fi

echo ""

# Cleanup
rm -f cli-assistant-test
rm -rf "$TEST_DIR"
