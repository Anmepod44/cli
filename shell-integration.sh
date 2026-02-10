#!/bin/bash

# CLI Command Assistant - Shell Integration
# This file is automatically sourced by /etc/profile.d/

# Prevent multiple loading
if [ -n "$CLI_ASSISTANT_LOADED" ]; then
    return 0
fi
export CLI_ASSISTANT_LOADED=1

# Main function for command generation
function __cli_assistant_ask {
    # Get the query
    local query="$*"
    
    if [ -z "$query" ]; then
        echo "Usage: ask <describe what you want to do>"
        echo "Example: ask search for text in files"
        return 1
    fi
    
    # Try to find cli-assistant in PATH or current directory
    local cli_cmd="cli-assistant"
    if ! command -v cli-assistant &> /dev/null; then
        # Try current directory
        if [ -f "./cli-assistant" ]; then
            cli_cmd="./cli-assistant"
        else
            echo "Error: cli-assistant not found. Please install it or add to PATH."
            return 1
        fi
    fi
    
    # Call cli-assistant with the query and get the command
    local result=$(echo "$query" | $cli_cmd --quick 2>/dev/null)
    
    if [ $? -eq 0 ] && [ -n "$result" ]; then
        echo -e "\033[1;36m→\033[0m $result"
        echo -en "\033[1;33mRun this command? [Y/n/e(dit)]:\033[0m "
        read -r response
        
        case "$response" in
            [Yy]|"")
                eval "$result"
                ;;
            [Ee])
                # Allow editing
                read -e -i "$result" -p "$(echo -e '\033[1;32mEdit:\033[0m ')" edited_cmd
                eval "$edited_cmd"
                ;;
            *)
                echo "Command cancelled."
                ;;
        esac
    else
        echo "Could not generate command. Try: cli-assistant"
    fi
}

# Create convenient aliases
alias ask='__cli_assistant_ask'
alias cmd='cli-assistant'
