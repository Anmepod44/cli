#!/bin/bash

# Enable alias expansion in script
shopt -s expand_aliases

# Source the integration
source ./shell-integration.sh

echo "=== CLI Command Assistant - Testing 'ask' command ==="
echo ""

# Test the ask command
echo "Test 1: ask list all files"
ask list all files

echo ""
echo "Test 2: ask show disk usage"
ask show disk usage

echo ""
echo "Test 3: ask find large files"
ask find large files

echo ""
echo "=== All tests complete! ==="
echo ""
echo "After installing, use these commands:"
echo "  ask list all files"
echo "  ask show disk usage"
echo "  ask find process nginx"
echo "  cmd (opens interactive mode)"
echo ""
