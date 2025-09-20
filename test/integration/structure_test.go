package integration

import (
	"os"
	"path/filepath"
	"testing"
)

// TestDirectoryStructure validates the complete project directory structure
// against the unified project structure specification
func TestDirectoryStructure(t *testing.T) {
	projectRoot := "../../"

	// Required directories based on unified-project-structure.md
	requiredDirs := []struct {
		path        string
		description string
		category    string
	}{
		// Application structure
		{"cmd/calculator", "Main application entry point", "application"},
		{"internal/calculation", "Core calculation logic", "internal"},
		{"internal/terminal", "Terminal UI components", "internal"},
		{"internal/config", "Configuration management", "internal"},
		{"internal/history", "Calculation history", "internal"},
		{"internal/parser", "Expression parsing", "internal"},
		{"internal/models", "Data models", "internal"},

		// Public packages
		{"pkg/calculator", "Public calculator interfaces", "public"},
		{"pkg/terminal", "Public terminal utilities", "public"},

		// Test structure
		{"test/unit", "Unit test files", "testing"},
		{"test/integration", "Integration test files", "testing"},
		{"test/e2e", "End-to-end test files", "testing"},

		// Documentation and scripts
		{"docs", "Project documentation", "documentation"},
		{"scripts", "Build and utility scripts", "scripts"},
		{"configs", "Configuration files", "configuration"},
	}

	missingDirs := []string{}

	for _, dir := range requiredDirs {
		fullPath := filepath.Join(projectRoot, dir.path)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			missingDirs = append(missingDirs, dir.path)
			t.Errorf("Missing required directory: %s (%s)", dir.path, dir.description)
		} else {
			// Verify it's actually a directory
			info, err := os.Stat(fullPath)
			if err != nil {
				t.Errorf("Failed to stat directory %s: %v", dir.path, err)
			} else if !info.IsDir() {
				t.Errorf("Path %s exists but is not a directory", dir.path)
			}
		}
	}

	if len(missingDirs) > 0 {
		t.Fatalf("Project structure incomplete. Missing %d required directories: %v", len(missingDirs), missingDirs)
	}
}

// TestKeyFilesExistence validates that critical files exist in their expected locations
func TestKeyFilesExistence(t *testing.T) {
	projectRoot := "../../"

	keyFiles := []struct {
		path        string
		description string
	}{
		{"go.mod", "Go module definition"},
		{"README.md", "Project documentation"},
		{".gitignore", "Git ignore rules"},
		{"cmd/calculator/main.go", "Application entry point"},
		{"scripts/validate-setup.sh", "Setup validation script"},
	}

	for _, file := range keyFiles {
		fullPath := filepath.Join(projectRoot, file.path)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Errorf("Missing critical file: %s (%s)", file.path, file.description)
		} else {
			// Verify it's a regular file
			info, err := os.Stat(fullPath)
			if err != nil {
				t.Errorf("Failed to stat file %s: %v", file.path, err)
			} else if !info.Mode().IsRegular() {
				t.Errorf("Path %s exists but is not a regular file", file.path)
			}
		}
	}
}

// TestDirectoryPermissions validates that directories have appropriate permissions
func TestDirectoryPermissions(t *testing.T) {
	projectRoot := "../../"

	testDirs := []string{
		"cmd/calculator",
		"internal/calculation",
		"test/unit",
		"docs",
	}

	for _, dir := range testDirs {
		fullPath := filepath.Join(projectRoot, dir)
		info, err := os.Stat(fullPath)
		if err != nil {
			t.Errorf("Failed to stat directory %s: %v", dir, err)
			continue
		}

		// Check if directory is readable and writable
		if info.Mode().Perm()&0400 == 0 {
			t.Errorf("Directory %s is not readable", dir)
		}
		if info.Mode().Perm()&0200 == 0 {
			t.Errorf("Directory %s is not writable", dir)
		}
	}
}
