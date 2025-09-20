package setup

import (
	"os"
	"strings"
	"testing"
)

// TestGoModExists verifies that the go.mod file exists in the project root
func TestGoModExists(t *testing.T) {
	goModPath := "../../go.mod"

	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		t.Fatalf("go.mod file does not exist at %s", goModPath)
	}

	// Verify it's a regular file, not a directory
	info, err := os.Stat(goModPath)
	if err != nil {
		t.Fatalf("Failed to stat go.mod: %v", err)
	}
	if !info.Mode().IsRegular() {
		t.Fatalf("go.mod exists but is not a regular file")
	}
}

// TestGoModContent validates the content of the go.mod file using table-driven tests
func TestGoModContent(t *testing.T) {
	content, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatalf("Failed to read go.mod: %v", err)
	}

	contentStr := string(content)

	tests := []struct {
		name        string
		required    string
		description string
	}{
		{
			name:        "module_declaration",
			required:    "module calculator",
			description: "Go module name should be 'calculator'",
		},
		{
			name:        "go_version",
			required:    "go 1.21",
			description: "Go version should be 1.21 or higher",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !strings.Contains(contentStr, tt.required) {
				t.Errorf("%s: %s", tt.description, tt.required)
			}
		})
	}
}

// TestGoModStructure validates the overall structure and format of go.mod
func TestGoModStructure(t *testing.T) {
	content, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatalf("Failed to read go.mod: %v", err)
	}

	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	// Verify basic structure
	if len(lines) < 2 {
		t.Fatal("go.mod file appears to be malformed (too few lines)")
	}

	// Check for module line at the beginning
	if !strings.HasPrefix(strings.TrimSpace(lines[0]), "module ") {
		t.Error("go.mod should start with module declaration")
	}

	// Check for go version line
	goVersionFound := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "go ") {
			goVersionFound = true
			break
		}
	}
	if !goVersionFound {
		t.Error("go.mod should contain a go version directive")
	}
}
