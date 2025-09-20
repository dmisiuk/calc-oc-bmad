package e2e

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestGitRepositoryInitialized(t *testing.T) {
	if _, err := os.Stat("../../.git"); os.IsNotExist(err) {
		t.Fatal("Git repository not initialized")
	}
}

func TestInitialCommitExists(t *testing.T) {
	cmd := exec.Command("git", "log", "--oneline")
	cmd.Dir = "../../"
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get git log: %v", err)
	}
	if !strings.Contains(string(output), "feat: initial project setup") {
		t.Error("Initial commit not found or incorrect message")
	}
}
