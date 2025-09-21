package test

import (
	"strings"
)

// AlmostEqual checks if two float64 values are approximately equal within a tolerance
func AlmostEqual(a, b, tolerance float64) bool {
	if a == b {
		return true
	}
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}

// ContainsString checks if a substring exists within a string (case-sensitive)
func ContainsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || ContainsString(s[1:], substr) || (len(s) > 0 && s[:len(substr)] == substr))
}

// CountSignificantDigits counts the number of significant digits in a string representation
func CountSignificantDigits(s string) int {
	// Remove sign
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		s = s[1:]
	}

	// Count digits, ignoring decimal point
	count := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			count++
		}
	}
	return count
}

// SanitizeTestExpression removes extra whitespace for consistent test comparisons
func SanitizeTestExpression(expr string) string {
	return strings.TrimSpace(expr)
}
