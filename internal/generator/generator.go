package generator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cli-command-assistant/internal/types"
)

// Generator implements CommandGenerator interface
type Generator struct {
	patterns []pattern
}

// pattern represents a natural language pattern and its command template
type pattern struct {
	regex       *regexp.Regexp
	template    string
	description string
	flags       []types.Flag
}

// NewGenerator creates a new command generator
func NewGenerator() *Generator {
	g := &Generator{}
	g.initializePatterns()
	return g
}

// initializePatterns sets up the pattern matching rules
func (g *Generator) initializePatterns() {
	g.patterns = []pattern{
		// List operations
		{
			regex:       regexp.MustCompile(`(?i)^list\s+(all\s+)?files?$`),
			template:    "ls -la",
			description: "List all files in the current directory with details",
			flags: []types.Flag{
				{Name: "-l", Description: "Use long listing format"},
				{Name: "-a", Description: "Show hidden files"},
			},
		},
		{
			regex:       regexp.MustCompile(`(?i)^list\s+(all\s+)?director(y|ies)$`),
			template:    "ls -ld */",
			description: "List all directories in the current directory",
			flags: []types.Flag{
				{Name: "-l", Description: "Use long listing format"},
				{Name: "-d", Description: "List directories themselves, not their contents"},
			},
		},

		// Find operations
		{
			regex:       regexp.MustCompile(`(?i)^find\s+files?\s+named?\s+["\']?(.+?)["\']?$`),
			template:    "find . -name \"%s\"",
			description: "Find files by name in current directory and subdirectories",
			flags: []types.Flag{
				{Name: "-name", Description: "Search by filename pattern"},
			},
		},
		{
			regex:       regexp.MustCompile(`(?i)^find\s+files?\s+larger\s+than\s+(\d+)(mb|gb|kb)?$`),
			template:    "find . -type f -size +%s%s",
			description: "Find files larger than specified size",
			flags: []types.Flag{
				{Name: "-type f", Description: "Search for files only"},
				{Name: "-size", Description: "Filter by file size"},
			},
		},
		{
			regex:       regexp.MustCompile(`(?i)^find\s+files?\s+modified\s+in\s+last\s+(\d+)\s+days?$`),
			template:    "find . -type f -mtime -%s",
			description: "Find files modified in the last N days",
			flags: []types.Flag{
				{Name: "-type f", Description: "Search for files only"},
				{Name: "-mtime", Description: "Filter by modification time"},
			},
		},

		// Process management
		{
			regex:       regexp.MustCompile(`(?i)^(list|show)\s+all\s+processes?$`),
			template:    "ps aux",
			description: "Show all running processes",
			flags: []types.Flag{
				{Name: "a", Description: "Show processes for all users"},
				{Name: "u", Description: "Display user-oriented format"},
				{Name: "x", Description: "Include processes without controlling terminals"},
			},
		},
		{
			regex:       regexp.MustCompile(`(?i)^find\s+process(es)?\s+["\']?(.+?)["\']?$`),
			template:    "ps aux | grep \"%s\"",
			description: "Find processes matching a pattern",
			flags: []types.Flag{
				{Name: "aux", Description: "Show all processes in user format"},
			},
		},
		{
			regex:       regexp.MustCompile(`(?i)^kill\s+process\s+(\d+)$`),
			template:    "kill %s",
			description: "Terminate a process by PID",
			flags:       []types.Flag{},
		},
		{
			regex:       regexp.MustCompile(`(?i)^force\s+kill\s+process\s+(\d+)$`),
			template:    "kill -9 %s",
			description: "Forcefully terminate a process by PID",
			flags: []types.Flag{
				{Name: "-9", Description: "Send SIGKILL signal (force kill)"},
			},
		},

		// Disk usage
		{
			regex:       regexp.MustCompile(`(?i)^(show|check)\s+disk\s+(usage|space)$`),
			template:    "df -h",
			description: "Show disk space usage for all mounted filesystems",
			flags: []types.Flag{
				{Name: "-h", Description: "Human-readable sizes (KB, MB, GB)"},
			},
		},
		{
			regex:       regexp.MustCompile(`(?i)^(show|check)\s+directory\s+size$`),
			template:    "du -sh *",
			description: "Show size of each item in current directory",
			flags: []types.Flag{
				{Name: "-s", Description: "Display only total for each argument"},
				{Name: "-h", Description: "Human-readable sizes"},
			},
		},

		// Network operations
		{
			regex:       regexp.MustCompile(`(?i)^(show|check)\s+network\s+connections?$`),
			template:    "netstat -tuln",
			description: "Show active network connections and listening ports",
			flags: []types.Flag{
				{Name: "-t", Description: "Show TCP connections"},
				{Name: "-u", Description: "Show UDP connections"},
				{Name: "-l", Description: "Show listening sockets"},
				{Name: "-n", Description: "Show numerical addresses"},
			},
		},
		{
			regex:       regexp.MustCompile(`(?i)^ping\s+(.+)$`),
			template:    "ping -c 4 %s",
			description: "Send 4 ping packets to test connectivity",
			flags: []types.Flag{
				{Name: "-c", Description: "Number of packets to send"},
			},
		},

		// File operations
		{
			regex:       regexp.MustCompile(`(?i)^copy\s+["\']?(.+?)["\']?\s+to\s+["\']?(.+?)["\']?$`),
			template:    "cp \"%s\" \"%s\"",
			description: "Copy file or directory to destination",
			flags:       []types.Flag{},
		},
		{
			regex:       regexp.MustCompile(`(?i)^move\s+["\']?(.+?)["\']?\s+to\s+["\']?(.+?)["\']?$`),
			template:    "mv \"%s\" \"%s\"",
			description: "Move or rename file or directory",
			flags:       []types.Flag{},
		},
		{
			regex:       regexp.MustCompile(`(?i)^delete\s+(file|directory)\s+["\']?(.+?)["\']?$`),
			template:    "rm -rf \"%s\"",
			description: "Delete file or directory (use with caution)",
			flags: []types.Flag{
				{Name: "-r", Description: "Remove directories recursively"},
				{Name: "-f", Description: "Force removal without confirmation"},
			},
		},

		// System information
		{
			regex:       regexp.MustCompile(`(?i)^(show|check)\s+memory\s+usage$`),
			template:    "free -h",
			description: "Show memory usage statistics",
			flags: []types.Flag{
				{Name: "-h", Description: "Human-readable sizes"},
			},
		},
		{
			regex:       regexp.MustCompile(`(?i)^(show|check)\s+system\s+info(rmation)?$`),
			template:    "uname -a",
			description: "Show system information",
			flags: []types.Flag{
				{Name: "-a", Description: "Show all system information"},
			},
		},
		{
			regex:       regexp.MustCompile(`(?i)^(show|check)\s+uptime$`),
			template:    "uptime",
			description: "Show how long the system has been running",
			flags:       []types.Flag{},
		},

		// Search in files
		{
			regex:       regexp.MustCompile(`(?i)^search\s+for\s+["\']?(.+?)["\']?\s+in\s+files?$`),
			template:    "grep -r \"%s\" .",
			description: "Search for text pattern in all files recursively",
			flags: []types.Flag{
				{Name: "-r", Description: "Search recursively in directories"},
			},
		},
	}
}

// Generate translates natural language to a shell command
func (g *Generator) Generate(input string) (*types.Command, error) {
	// Handle empty input
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	input = strings.TrimSpace(input)

	// Handle whitespace-only input
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty or whitespace only")
	}

	// Handle excessively long input
	if len(input) > 500 {
		return nil, fmt.Errorf("input is too long (max 500 characters)")
	}

	// Try to match against patterns
	for _, p := range g.patterns {
		if matches := p.regex.FindStringSubmatch(input); matches != nil {
			// Extract captured groups (skip first match which is the full string)
			params := matches[1:]

			// Build command from template
			var cmdStr string
			if len(params) > 0 {
				// Handle size units for find commands
				if strings.Contains(p.template, "find") && strings.Contains(p.template, "size") {
					size := params[0]
					unit := "M" // default to MB
					if len(params) > 1 && params[1] != "" {
						switch strings.ToLower(params[1]) {
						case "kb":
							unit = "k"
						case "mb":
							unit = "M"
						case "gb":
							unit = "G"
						}
					}
					cmdStr = fmt.Sprintf(p.template, size, unit)
				} else {
					// Standard parameter substitution
					args := make([]interface{}, len(params))
					for i, param := range params {
						args[i] = param
					}
					cmdStr = fmt.Sprintf(p.template, args...)
				}
			} else {
				cmdStr = p.template
			}

			return &types.Command{
				Raw:         cmdStr,
				Description: p.description,
				Flags:       p.flags,
			}, nil
		}
	}

	return nil, fmt.Errorf("could not translate input to a command. Try rephrasing or use 'help' for examples")
}

// Validate checks if a command is safe to execute
func (g *Generator) Validate(cmd *types.Command) []types.Warning {
	var warnings []types.Warning

	// Check for dangerous commands
	dangerousCommands := []string{"rm", "dd", "mkfs", "fdisk", "parted", "shred", "wipefs", ">"}
	for _, dangerous := range dangerousCommands {
		if strings.Contains(cmd.Raw, dangerous) {
			cmd.IsDangerous = true
			warnings = append(warnings, types.Warning{
				Level:   types.CriticalLevel,
				Message: fmt.Sprintf("Command contains potentially destructive operation: %s", dangerous),
			})
			break
		}
	}

	// Check for file modification/deletion patterns
	if strings.Contains(cmd.Raw, "rm -rf") || strings.Contains(cmd.Raw, "rm -fr") {
		cmd.IsDangerous = true
		warnings = append(warnings, types.Warning{
			Level:   types.CriticalLevel,
			Message: "Command will recursively delete files without confirmation",
		})
	}

	// Check for system directory targets
	systemDirs := []string{"/etc", "/sys", "/proc", "/boot", "/dev", "/root"}
	for _, sysDir := range systemDirs {
		if strings.Contains(cmd.Raw, sysDir) {
			cmd.IsDangerous = true
			warnings = append(warnings, types.Warning{
				Level:   types.CriticalLevel,
				Message: fmt.Sprintf("Command targets system directory: %s", sysDir),
			})
			break
		}
	}

	// Check for sudo requirement
	if strings.HasPrefix(strings.TrimSpace(cmd.Raw), "sudo") {
		cmd.RequiresSudo = true
		warnings = append(warnings, types.Warning{
			Level:   types.WarningLevelWarning,
			Message: "Command requires elevated privileges (sudo)",
		})
	}

	// Check for commands that typically need sudo
	sudoCommands := []string{"apt", "apt-get", "yum", "dnf", "systemctl", "service", "mount", "umount"}
	for _, sudoCmd := range sudoCommands {
		if strings.Contains(cmd.Raw, sudoCmd) && !cmd.RequiresSudo {
			warnings = append(warnings, types.Warning{
				Level:   types.InfoLevel,
				Message: fmt.Sprintf("Command may require sudo: %s", sudoCmd),
			})
			break
		}
	}

	// Check for file modification operations
	modifyOps := []string{"mv", "cp", "chmod", "chown", "touch"}
	for _, op := range modifyOps {
		if strings.Contains(cmd.Raw, op) {
			warnings = append(warnings, types.Warning{
				Level:   types.WarningLevelWarning,
				Message: fmt.Sprintf("Command will modify files: %s", op),
			})
			break
		}
	}

	return warnings
}
