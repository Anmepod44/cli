#!/bin/bash

# CLI Command Assistant Setup Script

set -e

echo "╔════════════════════════════════════════════════════════════╗"
echo "║     CLI Command Assistant - Setup                          ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""

# Create config directory
CONFIG_DIR="$HOME/.cli-assistant"
CONFIG_FILE="$CONFIG_DIR/config.yaml"

echo "Creating configuration directory..."
mkdir -p "$CONFIG_DIR"

# Check if config already exists
if [ -f "$CONFIG_FILE" ]; then
    echo "⚠ Configuration file already exists at: $CONFIG_FILE"
    read -p "Do you want to overwrite it? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Setup cancelled. Existing configuration preserved."
        exit 0
    fi
fi

# Ask about LLM mode
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "LLM Mode Configuration"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "The CLI Assistant can work in two modes:"
echo "  1. Pattern-based (default) - Fast, offline, rule-based"
echo "  2. LLM-powered - AI understanding using OpenAI GPT"
echo ""
read -p "Do you want to enable LLM mode? (y/N): " -n 1 -r
echo

USE_LLM=false
API_KEY=""

if [[ $REPLY =~ ^[Yy]$ ]]; then
    USE_LLM=true
    echo ""
    echo "To use LLM mode, you need an OpenAI API key."
    echo "Get one at: https://platform.openai.com/api-keys"
    echo ""
    read -p "Enter your OpenAI API key (or press Enter to skip): " API_KEY
    
    if [ -z "$API_KEY" ]; then
        echo ""
        echo "ℹ No API key provided. You can add it later by:"
        echo "  1. Editing $CONFIG_FILE"
        echo "  2. Or setting OPENAI_API_KEY environment variable"
        echo ""
    fi
fi

# Create configuration file
echo ""
echo "Creating configuration file..."

cat > "$CONFIG_FILE" << EOF
# CLI Command Assistant Configuration
# Generated on $(date)

# Skip confirmation prompts before executing commands
autoexecute: false

# Maximum number of commands that can run concurrently (1-20)
maxconcurrent: 5

# Maximum number of commands to store in history (1-10000)
historysize: 1000

# Require explicit confirmation before executing dangerous commands
confirmdangerous: true

# Enable clipboard functionality (requires xclip or xsel)
clipboardenabled: true

# Enable LLM-powered command generation using OpenAI's GPT
usellm: $USE_LLM

# OpenAI API key for LLM mode
openaiapikey: "$API_KEY"
EOF

echo "✓ Configuration file created at: $CONFIG_FILE"

# Check for clipboard tools
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "Clipboard Support"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

if command -v xclip &> /dev/null; then
    echo "✓ xclip is installed"
elif command -v xsel &> /dev/null; then
    echo "✓ xsel is installed"
else
    echo "⚠ No clipboard tool found (xclip or xsel)"
    echo ""
    read -p "Do you want to install xclip? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "Installing xclip..."
        sudo apt-get update && sudo apt-get install -y xclip
        echo "✓ xclip installed"
    else
        echo "ℹ You can install it later with: sudo apt install xclip"
    fi
fi

# Summary
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "Setup Complete!"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Configuration:"
echo "  Location: $CONFIG_FILE"
echo "  LLM Mode: $USE_LLM"
if [ "$USE_LLM" = true ] && [ -n "$API_KEY" ]; then
    echo "  API Key: Configured ✓"
elif [ "$USE_LLM" = true ]; then
    echo "  API Key: Not configured (will use OPENAI_API_KEY env var)"
fi
echo ""
echo "Next steps:"
echo "  1. Build the application: go build ./cmd/cli-assistant"
echo "  2. Run it: ./cli-assistant"
echo ""
echo "To modify configuration later, edit: $CONFIG_FILE"
echo ""
