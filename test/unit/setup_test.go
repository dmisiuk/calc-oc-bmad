package setup

import (
	"os"
	"strings"
	"testing"
)

func TestGoModExists(t *testing.T) {
	if _, err := os.Stat("../../go.mod"); os.IsNotExist(err) {
		t.Fatal("go.mod file does not exist")
	}
}

func TestGoModContent(t *testing.T) {
	content, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatalf("Failed to read go.mod: %v", err)
	}
	str := string(content)
	if !strings.Contains(str, "module calculator") {
		t.Error("go.mod does not contain 'module calculator'")
	}
	if !strings.Contains(str, "go 1.21") {
		t.Error("go.mod does not contain Go version 1.21+")
	}
}
