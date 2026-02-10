#!/bin/bash

echo "=== Testing CLI Command Assistant Shell Integration ==="
echo ""

# Source the integration
source ./shell-integration.sh

echo ""
echo "The '?' function is now loaded!"
echo ""
echo "=== Usage Examples ==="
echo ""
echo "You can use it in two ways:"
echo ""
echo "1. Using the ? function (with backslash escape):"
echo "   \\? list all files"
echo "   \\? show disk usage"
echo "   \\? find large files"
echo ""
echo "2. Using the ask-cli function (easier to type):"
echo "   ask-cli list all files"
echo "   ask-cli show disk usage"
echo "   ask-cli find large files"
echo ""
echo "3. Or run interactively:"
echo "   cli-assistant"
echo ""
echo "=== Quick Test ==="
echo ""
echo "Testing: ask-cli list all files"
ask-cli list all files
echo ""
echo "Testing: ask-cli show disk usage"
ask-cli show disk usage
echo ""
echo "=== Installation Note ==="
echo ""
echo "After installing the Debian package, you can use:"
echo "  \\? <your query>      (requires backslash)"
echo "  ask-cli <your query> (easier alternative)"
echo "  cli-assistant        (interactive mode)"
echo ""
