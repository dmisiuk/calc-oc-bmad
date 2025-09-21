package calculation_test

import (
	"math/big"
	"testing"

	"calculator/internal/calculation"
	"calculator/test"
)

func TestValidateExpression(t *testing.T) {
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
			errorMsg:    "first number validation failed",
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
		{
			name:        "too many parts",
			expression:  "2 + 3 + 4",
			expectError: true,
			errorMsg:    "invalid format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := calculation.ValidateExpression(tt.expression)

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

func TestSanitizeExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid expression",
			input:    "2 + 3",
			expected: "2 + 3",
		},
		{
			name:     "remove special characters",
			input:    "2 + 3; DROP TABLE users;",
			expected: "2 + 3   ",
		},
		{
			name:     "remove symbols",
			input:    "2@ + #3",
			expected: "2 + 3",
		},
		{
			name:     "mixed valid and invalid",
			input:    "1.5 * (2 + 3)",
			expected: "1.5 * 2 + 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculation.SanitizeExpression(tt.input)
			if result != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestValidatePrecision(t *testing.T) {
	tests := []struct {
		name              string
		result            string
		expectedPrecision int
		expectError       bool
		errorMsg          string
	}{
		{
			name:              "sufficient precision",
			result:            "1.23456789012345",
			expectedPrecision: 15,
		},
		{
			name:              "simple precision",
			result:            "1.23",
			expectedPrecision: 15,
		},
		{
			name:              "high precision number",
			result:            "1.23456789012345678901234567890",
			expectedPrecision: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := new(big.Float).SetString(tt.result)

			err := calculation.ValidatePrecision(result, tt.expectedPrecision)

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
