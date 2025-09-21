package calculation_test

import (
	"testing"

	"calculator/internal/calculation"
	"calculator/test"
)

func TestCalculationEngine_Calculate(t *testing.T) {
	engine := calculation.NewCalculationEngine()

	tests := []struct {
		name        string
		expression  string
		expected    float64
		expectError bool
		errorMsg    string
	}{
		// Addition tests
		{
			name:       "simple addition",
			expression: "2 + 3",
			expected:   5.0,
		},
		{
			name:       "decimal addition",
			expression: "1.5 + 2.25",
			expected:   3.75,
		},
		{
			name:       "negative addition",
			expression: "-5 + 3",
			expected:   -2.0,
		},
		{
			name:       "large number addition",
			expression: "123456789.123456789 + 987654321.987654321",
			expected:   1111111111.11111111,
		},

		// Subtraction tests
		{
			name:       "simple subtraction",
			expression: "10 - 4",
			expected:   6.0,
		},
		{
			name:       "decimal subtraction",
			expression: "5.5 - 2.25",
			expected:   3.25,
		},
		{
			name:       "negative result subtraction",
			expression: "3 - 5",
			expected:   -2.0,
		},
		{
			name:       "negative subtraction",
			expression: "-5 - 3",
			expected:   -8.0,
		},

		// Multiplication tests
		{
			name:       "simple multiplication",
			expression: "6 * 7",
			expected:   42.0,
		},
		{
			name:       "decimal multiplication",
			expression: "3.5 * 2.0",
			expected:   7.0,
		},
		{
			name:       "negative multiplication",
			expression: "-4 * 5",
			expected:   -20.0,
		},
		{
			name:       "zero multiplication",
			expression: "0 * 100",
			expected:   0.0,
		},

		// Division tests
		{
			name:       "simple division",
			expression: "15 / 3",
			expected:   5.0,
		},
		{
			name:       "decimal division",
			expression: "7.5 / 2.5",
			expected:   3.0,
		},
		{
			name:       "negative division",
			expression: "-10 / 2",
			expected:   -5.0,
		},
		{
			name:       "fractional division",
			expression: "1 / 2",
			expected:   0.5,
		},

		// Error cases
		{
			name:        "division by zero",
			expression:  "10 / 0",
			expectError: true,
			errorMsg:    "division by zero",
		},
		{
			name:        "invalid operator",
			expression:  "5 % 2",
			expectError: true,
			errorMsg:    "unsupported operator",
		},
		{
			name:        "invalid number",
			expression:  "abc + 2",
			expectError: true,
			errorMsg:    "invalid first number",
		},
		{
			name:        "empty expression",
			expression:  "",
			expectError: true,
			errorMsg:    "expression cannot be empty",
		},
		{
			name:        "invalid format",
			expression:  "2 +",
			expectError: true,
			errorMsg:    "invalid format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := engine.Calculate(tt.expression)

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
				if !test.AlmostEqual(result, tt.expected, 1e-10) {
					t.Errorf("expected %f, got %f", tt.expected, result)
				}
			}
		})
	}
}

func TestCalculationEngine_Validate(t *testing.T) {
	engine := calculation.NewCalculationEngine()

	tests := []struct {
		name        string
		expression  string
		expectError bool
		errorMsg    string
	}{
		{
			name:       "valid addition",
			expression: "2 + 3",
		},
		{
			name:       "valid subtraction",
			expression: "10 - 4",
		},
		{
			name:       "valid multiplication",
			expression: "6 * 7",
		},
		{
			name:       "valid division",
			expression: "15 / 3",
		},
		{
			name:        "division by zero",
			expression:  "10 / 0",
			expectError: true,
			errorMsg:    "division by zero detected",
		},
		{
			name:        "invalid operator",
			expression:  "5 % 2",
			expectError: true,
			errorMsg:    "unsupported operator",
		},
		{
			name:        "invalid number",
			expression:  "abc + 2",
			expectError: true,
			errorMsg:    "invalid first number",
		},
		{
			name:        "empty expression",
			expression:  "",
			expectError: true,
			errorMsg:    "expression cannot be empty",
		},
		{
			name:        "invalid format",
			expression:  "2 +",
			expectError: true,
			errorMsg:    "invalid format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := engine.Validate(tt.expression)

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
			}
		})
	}
}

func TestCalculationEngine_GetSupportedOperations(t *testing.T) {
	engine := calculation.NewCalculationEngine()

	operations := engine.GetSupportedOperations()

	expected := []string{"+", "-", "*", "/"}

	if len(operations) != len(expected) {
		t.Errorf("expected %d operations, got %d", len(expected), len(operations))
	}

	for i, op := range expected {
		if i >= len(operations) || operations[i] != op {
			t.Errorf("expected operation %d to be '%s', got '%s'", i, op, operations[i])
		}
	}
}
