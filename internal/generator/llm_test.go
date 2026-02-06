package generator

import (
	"testing"
)

func TestLLMGeneratorFallback(t *testing.T) {
	// Test that LLM generator falls back to pattern-based when no API key
	gen := NewLLMGenerator("")

	cmd, err := gen.Generate("list all files")
	if err != nil {
		t.Fatalf("Expected fallback to succeed, got error: %v", err)
	}

	if cmd.Raw == "" {
		t.Error("Expected command to be generated")
	}

	if cmd.Description == "" {
		t.Error("Expected description to be present")
	}
}

func TestLLMGeneratorEmptyInput(t *testing.T) {
	gen := NewLLMGenerator("")

	_, err := gen.Generate("")
	if err == nil {
		t.Error("Expected error for empty input")
	}
}

func TestLLMGeneratorWhitespaceInput(t *testing.T) {
	gen := NewLLMGenerator("")

	_, err := gen.Generate("   ")
	if err == nil {
		t.Error("Expected error for whitespace-only input")
	}
}

func TestLLMGeneratorLongInput(t *testing.T) {
	gen := NewLLMGenerator("")

	longInput := ""
	for i := 0; i < 600; i++ {
		longInput += "a"
	}

	_, err := gen.Generate(longInput)
	if err == nil {
		t.Error("Expected error for excessively long input")
	}
}

func TestExtractJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "plain JSON",
			input:    `{"command": "ls"}`,
			expected: `{"command": "ls"}`,
		},
		{
			name:     "JSON in markdown code block",
			input:    "```json\n{\"command\": \"ls\"}\n```",
			expected: `{"command": "ls"}`,
		},
		{
			name:     "JSON in generic code block",
			input:    "```\n{\"command\": \"ls\"}\n```",
			expected: `{"command": "ls"}`,
		},
		{
			name:     "JSON with whitespace",
			input:    "  \n{\"command\": \"ls\"}\n  ",
			expected: `{"command": "ls"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractJSON(tt.input)
			if result != tt.expected {
				t.Errorf("extractJSON() = %q, want %q", result, tt.expected)
			}
		})
	}
}
