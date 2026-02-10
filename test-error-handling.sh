#!/bin/bash

# Enable alias expansion
shopt -s expand_aliases

# Source the integration
source ./shell-integration.sh

echo "=== Testing Error Handling in CLI Command Assistant ==="
echo ""

echo "Test 1: Successful command"
echo "Command: ask list all files"
echo ""
echo "y" | ask list all files
echo ""
echo "---"
echo ""

echo "Test 2: Command not found (exit code 127)"
echo "Command: ask play music"
echo ""
echo "y" | ask play music
echo ""
echo "---"
echo ""

echo "Test 3: Permission denied simulation"
echo "Command: ask show system logs"
echo ""
echo "y" | ask show system logs
echo ""
echo "---"
echo ""

echo "=== Error Handling Features ==="
echo ""
echo "The enhanced 'ask' command now:"
echo "  ✓ Detects command execution failures"
echo "  ✓ Shows colored success/error messages"
echo "  ✓ Provides contextual tips based on error type"
echo "  ✓ Suggests next steps when commands fail"
echo ""
echo "Exit codes handled:"
echo "  0   - Success (green checkmark)"
echo "  127 - Command not found (suggests installation)"
echo "  126 - Permission denied (suggests sudo)"
echo "  130 - Interrupted (Ctrl+C)"
echo "  >128 - Terminated by signal"
echo "  Other - Generic failure with exit code"
echo ""
