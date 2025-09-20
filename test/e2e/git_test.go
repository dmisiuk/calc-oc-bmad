package e2e

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestGitRepositoryInitialized validates that Git repository is properly initialized
func TestGitRepositoryInitialized(t *testing.T) {
	gitDir := "../../.git"

	// Check if .git directory exists
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		t.Fatalf("Git repository not initialized - .git directory missing at %s", gitDir)
	}

	// Verify .git is a directory
	info, err := os.Stat(gitDir)
	if err != nil {
		t.Fatalf("Failed to stat .git directory: %v", err)
	}
	if !info.IsDir() {
		t.Fatalf(".git exists but is not a directory")
	}

	// Check for essential Git files
	essentialFiles := []string{
		"../../.git/HEAD",
		"../../.git/config",
		"../../.git/refs/heads",
	}

	for _, file := range essentialFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Essential Git file missing: %s", file)
		}
	}
}

// TestInitialCommitExists validates that repository has commits and follows conventional format
func TestInitialCommitExists(t *testing.T) {
	cmd := exec.Command("git", "log", "--oneline", "--reverse")
	cmd.Dir = "../../"

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get git log: %v", err)
	}

	outputStr := string(output)
	lines := strings.Split(strings.TrimSpace(outputStr), "\n")

	if len(lines) == 0 {
		t.Fatal("No commits found in repository")
	}

	// Check that we have commits and they follow conventional commit format
	foundConventionalCommit := false
	for _, line := range lines {
		if strings.Contains(line, "feat:") || strings.Contains(line, "fix:") || strings.Contains(line, "docs:") {
			foundConventionalCommit = true
			break
		}
	}

	if !foundConventionalCommit {
		t.Error("No conventional commits found in repository history")
	}

	// Verify we have at least one commit
	if len(lines) < 1 {
		t.Error("Repository should have at least one commit")
	}

	t.Logf("Found %d commits in repository", len(lines))
}

// TestGitRemoteConfiguration validates GitHub remote repository setup
func TestGitRemoteConfiguration(t *testing.T) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = "../../"

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get git remotes: %v", err)
	}

	outputStr := string(output)

	// Check for origin remote
	if !strings.Contains(outputStr, "origin") {
		t.Error("Git remote 'origin' not configured")
	}

	// Check for GitHub URL pattern
	if !strings.Contains(outputStr, "github.com") {
		t.Error("GitHub remote URL not found in remotes")
	}

	// Verify remote is accessible (optional, but good practice)
	cmd = exec.Command("git", "ls-remote", "--heads", "origin")
	cmd.Dir = "../../"

	if err := cmd.Run(); err != nil {
		t.Errorf("GitHub remote not accessible: %v", err)
	}
}

// TestGitIgnoreConfiguration validates .gitignore file exists and contains Go-specific patterns
func TestGitIgnoreConfiguration(t *testing.T) {
	gitignorePath := "../../.gitignore"

	// Check if .gitignore exists
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		t.Fatalf(".gitignore file not found at %s", gitignorePath)
	}

	// Read .gitignore content
	content, err := os.ReadFile(gitignorePath)
	if err != nil {
		t.Fatalf("Failed to read .gitignore: %v", err)
	}

	contentStr := string(content)

	// Check for essential Go-specific ignore patterns
	requiredPatterns := []string{
		"*.exe",   // Windows executables
		"*.exe~",  // Windows executable backups
		"*.dll",   // Windows DLLs
		"*.so",    // Linux shared objects
		"*.dylib", // macOS dynamic libraries
		"bin/",    // Binary directory
	}

	for _, pattern := range requiredPatterns {
		if !strings.Contains(contentStr, pattern) {
			t.Errorf(".gitignore missing Go-specific pattern: %s", pattern)
		}
	}
}

// TestGitRepositoryStatus validates Git repository status and provides informative feedback
func TestGitRepositoryStatus(t *testing.T) {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = "../../"

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get git status: %v", err)
	}

	outputStr := strings.TrimSpace(string(output))

	// During development, working directory may not be clean
	// Just verify git status command works and repository is functional
	if len(outputStr) > 0 {
		t.Logf("Working directory has uncommitted changes: %s", outputStr)
		// This is not a failure - just informational
	}

	// Verify we can get branch information
	cmd = exec.Command("git", "branch", "--show-current")
	cmd.Dir = "../../"

	branchOutput, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get current branch: %v", err)
	}

	currentBranch := strings.TrimSpace(string(branchOutput))
	if currentBranch == "" {
		t.Error("Unable to determine current git branch")
	} else {
		t.Logf("Current branch: %s", currentBranch)
	}
}
