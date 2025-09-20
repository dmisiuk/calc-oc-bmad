package integration

import (
	"os"
	"testing"
)

func TestDirectoryStructure(t *testing.T) {
	dirs := []string{
		"../../cmd/calculator",
		"../../internal/calculation",
		"../../internal/terminal",
		"../../internal/config",
		"../../internal/history",
		"../../internal/parser",
		"../../internal/models",
		"../../pkg/calculator",
		"../../pkg/terminal",
		"../../test/unit",
		"../../test/integration",
		"../../test/e2e",
		"../../docs",
		"../../scripts",
		"../../configs",
	}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Directory %s does not exist", dir)
		}
	}
}
