#!/bin/bash

# Demo script to show the CLI Command Assistant in action

echo "╔════════════════════════════════════════════════════════════╗"
echo "║     CLI Command Assistant - Interactive Demo              ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""
echo "Starting the application..."
echo ""

# Test commands to demonstrate both pattern and LLM modes
cat << 'EOF' | ./cli-assistant
list all files
3
show me all log files modified in the last week
3
find the largest files in my home directory
3
exit
EOF

echo ""
echo "Demo complete!"
