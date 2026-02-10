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
                # Execute and capture exit code
                eval "$result"
                local exit_code=$?
                
                # Provide feedback based on exit code
                if [ $exit_code -eq 0 ]; then
                    echo -e "\033[1;32m✓ Command completed successfully\033[0m"
                elif [ $exit_code -eq 127 ]; then
                    echo -e "\033[1;31m✗ Command failed: Command not found\033[0m"
                    echo -e "\033[1;33mTip: The command may need to be installed or isn't in your PATH\033[0m"
                    echo -e "\033[1;36mTry asking: ask how to install ${result%% *}\033[0m"
                elif [ $exit_code -eq 126 ]; then
                    echo -e "\033[1;31m✗ Command failed: Permission denied\033[0m"
                    echo -e "\033[1;33mTip: You may need to run with sudo or check file permissions\033[0m"
                elif [ $exit_code -eq 130 ]; then
                    echo -e "\033[1;33m⚠ Command interrupted (Ctrl+C)\033[0m"
                elif [ $exit_code -gt 128 ]; then
                    echo -e "\033[1;31m✗ Command terminated by signal (exit code: $exit_code)\033[0m"
                else
                    echo -e "\033[1;31m✗ Command failed with exit code: $exit_code\033[0m"
                    echo -e "\033[1;33mTip: Check the error message above for details\033[0m"
                fi
                
                return $exit_code
                ;;
            [Ee])
                # Allow editing
                read -e -i "$result" -p "$(echo -e '\033[1;32mEdit:\033[0m ')" edited_cmd
                
                # Execute edited command and capture exit code
                eval "$edited_cmd"
                local exit_code=$?
                
                # Provide feedback
                if [ $exit_code -eq 0 ]; then
                    echo -e "\033[1;32m✓ Command completed successfully\033[0m"
                else
                    echo -e "\033[1;31m✗ Command failed with exit code: $exit_code\033[0m"
                fi
                
                return $exit_code
                ;;
            *)
                echo "Command cancelled."
                return 0
                ;;
        esac
    else
        echo "Could not generate command. Try: cli-assistant"
        return 1
    fi
}

# Create convenient aliases
alias ask='__cli_assistant_ask'
alias cmd='cli-assistant'
