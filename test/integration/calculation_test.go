package integration_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"calculator/internal/calculation"
	"calculator/test"
)

func TestCalculationEngine_EndToEndWorkflow(t *testing.T) {
	engine := calculation.NewCalculationEngine()

	tests := []struct {
		name        string
		expression  string
		expected    float64
		description string
	}{
		{
			name:        "complete addition workflow",
			expression:  "123.456789 + 987.654321",
			expected:    1111.11111,
			description: "Test full addition workflow with decimal precision",
		},
		{
			name:        "complete subtraction workflow",
			expression:  "1000.000000 - 123.456789",
			expected:    876.543211,
			description: "Test full subtraction workflow with precision",
		},
		{
			name:        "complete multiplication workflow",
			expression:  "12.345678 * 9.876543",
			expected:    121.9326196312,
			description: "Test full multiplication workflow",
		},
		{
			name:        "complete division workflow",
			expression:  "144.000000 / 12.000000",
			expected:    12.0,
			description: "Test full division workflow",
		},
		{
			name:        "negative numbers workflow",
			expression:  "-25.5 + 10.25",
			expected:    -15.25,
			description: "Test negative number handling",
		},
		{
			name:        "fractional division workflow",
			expression:  "1.000000 / 3.000000",
			expected:    0.3333333333,
			description: "Test fractional division precision",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate expression first
			err := engine.Validate(tt.expression)
			if err != nil {
				t.Fatalf("validation failed for %s: %v", tt.description, err)
			}

			// Perform calculation
			start := time.Now()
			result, err := engine.Calculate(tt.expression)
			duration := time.Since(start)

			if err != nil {
				t.Fatalf("calculation failed for %s: %v", tt.description, err)
			}

			// Check result accuracy
			if !test.AlmostEqual(result, tt.expected, 1e-10) {
				t.Errorf("%s: expected %.10f, got %.10f", tt.description, tt.expected, result)
			}

			// Check performance (should be under 50ms)
			if duration > 50*time.Millisecond {
				t.Errorf("%s: calculation took %v, exceeds 50ms limit", tt.description, duration)
			}

			// Verify supported operations
			operations := engine.GetSupportedOperations()
			if len(operations) == 0 {
				t.Error("no supported operations returned")
			}
		})
	}
}

func TestCalculationEngine_ErrorHandlingWorkflow(t *testing.T) {
	engine := calculation.NewCalculationEngine()

	errorTests := []struct {
		name        string
		expression  string
		expectedErr string
		description string
	}{
		{
			name:        "division by zero workflow",
			expression:  "100 / 0",
			expectedErr: "division by zero",
			description: "Test division by zero error handling",
		},
		{
			name:        "invalid operator workflow",
			expression:  "10 ^ 2",
			expectedErr: "unsupported operator",
			description: "Test invalid operator error handling",
		},
		{
			name:        "invalid number workflow",
			expression:  "abc + 5",
			expectedErr: "invalid first number",
			description: "Test invalid number error handling",
		},
		{
			name:        "malformed expression workflow",
			expression:  "2 +",
			expectedErr: "invalid format",
			description: "Test malformed expression error handling",
		},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate should catch the error
			err := engine.Validate(tt.expression)
			if err == nil {
				t.Errorf("%s: expected validation to fail", tt.description)
			} else if !test.ContainsString(err.Error(), tt.expectedErr) {
				t.Errorf("%s: expected error containing '%s', got '%s'", tt.description, tt.expectedErr, err.Error())
			}

			// Calculate should also fail with same error
			_, err = engine.Calculate(tt.expression)
			if err == nil {
				t.Errorf("%s: expected calculation to fail", tt.description)
			} else if !test.ContainsString(err.Error(), tt.expectedErr) {
				t.Errorf("%s: expected error containing '%s', got '%s'", tt.description, tt.expectedErr, err.Error())
			}
		})
	}
}

func TestCalculationEngine_PrecisionWorkflow(t *testing.T) {
	engine := calculation.NewCalculationEngine()

	precisionTests := []struct {
		name        string
		expression  string
		minDigits   int
		description string
	}{
		{
			name:        "high precision addition",
			expression:  "1.23456789012345 + 9.87654321098765",
			minDigits:   14,
			description: "Test 15-digit precision in addition",
		},
		{
			name:        "high precision multiplication",
			expression:  "1.11111111111111 * 2.22222222222222",
			minDigits:   14,
			description: "Test 15-digit precision in multiplication",
		},
		{
			name:        "high precision division",
			expression:  "10.00000000000000 / 3.00000000000000",
			minDigits:   14,
			description: "Test 15-digit precision in division",
		},
	}

	for _, tt := range precisionTests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := engine.Calculate(tt.expression)
			if err != nil {
				t.Fatalf("%s: calculation failed: %v", tt.description, err)
			}

			// Check if result has sufficient precision
			resultStr := fmt.Sprintf("%.15f", result)
			// Remove trailing zeros and decimal point if needed
			resultStr = strings.TrimRight(resultStr, "0")
			resultStr = strings.TrimRight(resultStr, ".")

			// Count significant digits
			significantDigits := test.CountSignificantDigits(resultStr)
			if significantDigits < tt.minDigits {
				t.Errorf("%s: expected at least %d significant digits, got %d (result: %s)",
					tt.description, tt.minDigits, significantDigits, resultStr)
			}
		})
	}
}

func TestCalculationEngine_OperationsCoverage(t *testing.T) {
	engine := calculation.NewCalculationEngine()

	operations := engine.GetSupportedOperations()
	expectedOps := []string{"+", "-", "*", "/"}

	if len(operations) != len(expectedOps) {
		t.Errorf("expected %d operations, got %d", len(expectedOps), len(operations))
	}

	// Test each operation
	for _, op := range expectedOps {
		t.Run("operation_"+op, func(t *testing.T) {
			expression := "10 " + op + " 5"
			_, err := engine.Calculate(expression)
			if err != nil {
				t.Errorf("operation %s failed: %v", op, err)
			}
		})
	}
}
