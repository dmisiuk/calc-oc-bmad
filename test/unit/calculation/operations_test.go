package calculation_test

import (
	"math/big"
	"testing"

	"calculator/internal/calculation"
	"calculator/test"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "simple addition",
			a:        "2",
			b:        "3",
			expected: "5",
		},
		{
			name:     "decimal addition",
			a:        "1.5",
			b:        "2.25",
			expected: "3.75",
		},
		{
			name:     "negative addition",
			a:        "-5",
			b:        "3",
			expected: "-2",
		},
		{
			name:     "large numbers",
			a:        "123456789.123456789",
			b:        "987654321.987654321",
			expected: "1111111111.11111111",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := big.NewFloat(0).SetString(tt.a)
			b, _ := big.NewFloat(0).SetString(tt.b)
			expected, _ := big.NewFloat(0).SetString(tt.expected)

			result, err := calculation.Add(a, b)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result.Cmp(expected) != 0 {
				t.Errorf("expected %s, got %s", expected.String(), result.String())
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "simple subtraction",
			a:        "10",
			b:        "4",
			expected: "6",
		},
		{
			name:     "decimal subtraction",
			a:        "5.5",
			b:        "2.25",
			expected: "3.25",
		},
		{
			name:     "negative result",
			a:        "3",
			b:        "5",
			expected: "-2",
		},
		{
			name:     "negative subtraction",
			a:        "-5",
			b:        "3",
			expected: "-8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := big.NewFloat(0).SetString(tt.a)
			b, _ := big.NewFloat(0).SetString(tt.b)
			expected, _ := big.NewFloat(0).SetString(tt.expected)

			result, err := calculation.Subtract(a, b)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result.Cmp(expected) != 0 {
				t.Errorf("expected %s, got %s", expected.String(), result.String())
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "simple multiplication",
			a:        "6",
			b:        "7",
			expected: "42",
		},
		{
			name:     "decimal multiplication",
			a:        "3.5",
			b:        "2.0",
			expected: "7",
		},
		{
			name:     "negative multiplication",
			a:        "-4",
			b:        "5",
			expected: "-20",
		},
		{
			name:     "zero multiplication",
			a:        "0",
			b:        "100",
			expected: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := big.NewFloat(0).SetString(tt.a)
			b, _ := big.NewFloat(0).SetString(tt.b)
			expected, _ := big.NewFloat(0).SetString(tt.expected)

			result, err := calculation.Multiply(a, b)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result.Cmp(expected) != 0 {
				t.Errorf("expected %s, got %s", expected.String(), result.String())
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           string
		b           string
		expected    string
		expectError bool
		errorMsg    string
	}{
		{
			name:     "simple division",
			a:        "15",
			b:        "3",
			expected: "5",
		},
		{
			name:     "decimal division",
			a:        "7.5",
			b:        "2.5",
			expected: "3",
		},
		{
			name:     "negative division",
			a:        "-10",
			b:        "2",
			expected: "-5",
		},
		{
			name:     "fractional division",
			a:        "1",
			b:        "2",
			expected: "0.5",
		},
		{
			name:        "division by zero",
			a:           "10",
			b:           "0",
			expectError: true,
			errorMsg:    "division by zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := big.NewFloat(0).SetString(tt.a)
			b, _ := big.NewFloat(0).SetString(tt.b)

			result, err := calculation.Divide(a, b)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error containing '%s', got nil", tt.errorMsg)
				} else if !test.ContainsString(err.Error(), tt.errorMsg) {
					t.Errorf("expected error containing '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				expected, _ := big.NewFloat(0).SetString(tt.expected)
				if result.Cmp(expected) != 0 {
					t.Errorf("expected %s, got %s", expected.String(), result.String())
				}
			}
		})
	}
}
