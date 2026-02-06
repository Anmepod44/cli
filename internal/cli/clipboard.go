package cli

import (
	"fmt"
	"os/exec"
)

// Clipboard handles clipboard operations
type Clipboard struct{}

// NewClipboard creates a new clipboard handler
func NewClipboard() *Clipboard {
	return &Clipboard{}
}

// Copy copies text to the system clipboard
func (c *Clipboard) Copy(text string) error {
	// Try xclip first
	if err := c.copyWithXclip(text); err == nil {
		return nil
	}

	// Try xsel as fallback
	if err := c.copyWithXsel(text); err == nil {
		return nil
	}

	return fmt.Errorf("clipboard unavailable: neither xclip nor xsel found. Install with: sudo apt install xclip")
}

// copyWithXclip uses xclip to copy to clipboard
func (c *Clipboard) copyWithXclip(text string) error {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if _, err := stdin.Write([]byte(text)); err != nil {
		return err
	}

	if err := stdin.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}

// copyWithXsel uses xsel to copy to clipboard
func (c *Clipboard) copyWithXsel(text string) error {
	cmd := exec.Command("xsel", "--clipboard", "--input")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if _, err := stdin.Write([]byte(text)); err != nil {
		return err
	}

	if err := stdin.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}

// IsAvailable checks if clipboard functionality is available
func (c *Clipboard) IsAvailable() bool {
	// Check for xclip
	if _, err := exec.LookPath("xclip"); err == nil {
		return true
	}

	// Check for xsel
	if _, err := exec.LookPath("xsel"); err == nil {
		return true
	}

	return false
}
