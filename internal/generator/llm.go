package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/cli-command-assistant/internal/types"
)

// LLMGenerator uses OpenAI API for natural language understanding
type LLMGenerator struct {
	apiKey     string
	model      string
	httpClient *http.Client
	fallback   *Generator // Fallback to pattern-based generator
}

// OpenAI API structures
type openAIRequest struct {
	Model       string    `json:"model"`
	Messages    []message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []choice  `json:"choices"`
	Error   *apiError `json:"error,omitempty"`
}

type choice struct {
	Message message `json:"message"`
}

type apiError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// commandResponse is the expected JSON structure from the LLM
type commandResponse struct {
	Command      string `json:"command"`
	Description  string `json:"description"`
	Flags        []flag `json:"flags"`
	IsDangerous  bool   `json:"is_dangerous"`
	RequiresSudo bool   `json:"requires_sudo"`
	Explanation  string `json:"explanation,omitempty"`
}

type flag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NewLLMGenerator creates a new LLM-powered generator
func NewLLMGenerator(apiKey string) *LLMGenerator {
	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY")
	}

	return &LLMGenerator{
		apiKey:     apiKey,
		model:      "gpt-3.5-turbo",
		httpClient: &http.Client{},
		fallback:   NewGenerator(),
	}
}

// Generate translates natural language to a shell command using LLM
func (g *LLMGenerator) Generate(input string) (*types.Command, error) {
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

	// If no API key, fall back to pattern-based generator
	if g.apiKey == "" {
		return g.fallback.Generate(input)
	}

	// Try LLM generation
	cmd, err := g.generateWithLLM(input)
	if err != nil {
		// Fall back to pattern-based generator on error
		return g.fallback.Generate(input)
	}

	return cmd, nil
}

// generateWithLLM calls OpenAI API to generate command
func (g *LLMGenerator) generateWithLLM(input string) (*types.Command, error) {
	// Construct system prompt
	systemPrompt := `You are a Linux command generator. Given a natural language description, generate the appropriate Linux shell command.

Respond ONLY with valid JSON in this exact format:
{
  "command": "the shell command",
  "description": "brief explanation of what the command does",
  "flags": [
    {"name": "-flag", "description": "what this flag does"}
  ],
  "is_dangerous": false,
  "requires_sudo": false
}

Rules:
- Generate safe, correct Linux commands
- Set is_dangerous to true for commands that delete/modify files or affect system directories
- Set requires_sudo to true if the command needs elevated privileges
- Include all flags used in the command with their descriptions
- Keep descriptions concise and clear
- Do not include any text outside the JSON structure`

	userPrompt := fmt.Sprintf("Generate a Linux command for: %s", input)

	// Prepare request
	reqBody := openAIRequest{
		Model: g.model,
		Messages: []message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		Temperature: 0.3, // Lower temperature for more consistent output
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.apiKey)

	// Send request
	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	var apiResp openAIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check for API errors
	if apiResp.Error != nil {
		return nil, fmt.Errorf("API error: %s", apiResp.Error.Message)
	}

	if len(apiResp.Choices) == 0 {
		return nil, fmt.Errorf("no response from API")
	}

	// Extract command from response
	content := apiResp.Choices[0].Message.Content

	// Parse JSON response
	var cmdResp commandResponse
	if err := json.Unmarshal([]byte(content), &cmdResp); err != nil {
		// Try to extract JSON from markdown code blocks
		content = extractJSON(content)
		if err := json.Unmarshal([]byte(content), &cmdResp); err != nil {
			return nil, fmt.Errorf("failed to parse command response: %w", err)
		}
	}

	// Convert to Command type
	flags := make([]types.Flag, len(cmdResp.Flags))
	for i, f := range cmdResp.Flags {
		flags[i] = types.Flag{
			Name:        f.Name,
			Description: f.Description,
		}
	}

	return &types.Command{
		Raw:          cmdResp.Command,
		Description:  cmdResp.Description,
		Flags:        flags,
		IsDangerous:  cmdResp.IsDangerous,
		RequiresSudo: cmdResp.RequiresSudo,
	}, nil
}

// extractJSON attempts to extract JSON from markdown code blocks
func extractJSON(content string) string {
	// Remove markdown code blocks if present
	content = strings.TrimSpace(content)

	// Check for ```json ... ``` format
	if strings.HasPrefix(content, "```json") {
		content = strings.TrimPrefix(content, "```json")
		content = strings.TrimPrefix(content, "```")
		if idx := strings.LastIndex(content, "```"); idx != -1 {
			content = content[:idx]
		}
	} else if strings.HasPrefix(content, "```") {
		content = strings.TrimPrefix(content, "```")
		if idx := strings.LastIndex(content, "```"); idx != -1 {
			content = content[:idx]
		}
	}

	return strings.TrimSpace(content)
}

// Validate checks if a command is safe to execute
func (g *LLMGenerator) Validate(cmd *types.Command) []types.Warning {
	// Use the fallback generator's validation logic
	return g.fallback.Validate(cmd)
}
