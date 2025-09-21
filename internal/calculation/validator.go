package calculation

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

// ValidateExpression performs comprehensive validation of mathematical expressions
// Source: docs/architecture/security-and-performance.md - Input validation
func ValidateExpression(expression string) error {
	// Basic validation
	if err := validateBasicFormat(expression); err != nil {
		return err
	}

	// Parse and validate components
	num1, op, num2, err := parseExpressionComponents(expression)
	if err != nil {
		return err
	}

	// Validate numbers
	if err := validateNumber(num1); err != nil {
		return fmt.Errorf("first number validation failed: %w", err)
	}
	if err := validateNumber(num2); err != nil {
		return fmt.Errorf("second number validation failed: %w", err)
	}

	// Validate operator
	if err := validateOperator(op); err != nil {
		return err
	}

	// Check for division by zero
	if op == "/" {
		num2Float, err := parseBigFloat(num2)
		if err != nil {
			return fmt.Errorf("failed to parse second number for division check: %w", err)
		}
		if isZero(num2Float) {
			return fmt.Errorf("division by zero detected")
		}
	}

	return nil
}

// validateBasicFormat checks the basic structure of the expression
func validateBasicFormat(expression string) error {
	if strings.TrimSpace(expression) == "" {
		return fmt.Errorf("expression cannot be empty")
	}

	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return fmt.Errorf("invalid format: expected 'number operator number', got %d parts", len(parts))
	}

	return nil
}

// parseExpressionComponents breaks down the expression into components
func parseExpressionComponents(expression string) (string, string, string, error) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("invalid expression format")
	}
	return parts[0], parts[1], parts[2], nil
}

// validateNumber checks if a string represents a valid number
// Source: docs/architecture/security-and-performance.md - Safe math operations
func validateNumber(numStr string) error {
	numStr = strings.TrimSpace(numStr)

	// Check for basic numeric format
	if matched, _ := regexp.MatchString(`^-?\d+(\.\d+)?$`, numStr); !matched {
		return fmt.Errorf("invalid number format: %s", numStr)
	}

	// Try to parse as big.Float to ensure it's a valid number
	_, _, err := big.ParseFloat(numStr, 10, 100, big.ToNearestEven)
	if err != nil {
		return fmt.Errorf("failed to parse number: %w", err)
	}

	// Check for reasonable size limits (prevent overflow attacks)
	if len(numStr) > 1000 {
		return fmt.Errorf("number too large: maximum 1000 characters allowed")
	}

	return nil
}

// validateOperator checks if the operator is supported
func validateOperator(op string) error {
	validOps := []string{"+", "-", "*", "/"}
	for _, validOp := range validOps {
		if op == validOp {
			return nil
		}
	}
	return fmt.Errorf("unsupported operator: %s", op)
}

// isZero checks if a big.Float represents zero
func isZero(num *big.Float) bool {
	return num.Sign() == 0
}

// parseBigFloat converts a string to big.Float with error handling
func parseBigFloat(s string) (*big.Float, error) {
	// Remove any whitespace
	s = strings.TrimSpace(s)

	// Basic validation for numeric format
	if matched, _ := regexp.MatchString(`^-?\d+(\.\d+)?$`, s); !matched {
		return nil, fmt.Errorf("invalid number format: %s", s)
	}

	// Use high precision (more than 15 digits to ensure accuracy)
	precision := uint(100) // Higher than required 15 digits
	result, _, err := big.ParseFloat(s, 10, precision, big.ToNearestEven)
	if err != nil {
		return nil, fmt.Errorf("failed to parse number: %w", err)
	}

	return result, nil
}

// SanitizeExpression removes potentially harmful characters
// Source: docs/architecture/security-and-performance.md - Input sanitization
func SanitizeExpression(expression string) string {
	// Remove any non-numeric, non-operator, non-whitespace characters
	reg := regexp.MustCompile(`[^0-9+\-*/.\s]`)
	return reg.ReplaceAllString(expression, "")
}

// ValidatePrecision checks if a calculation result meets precision requirements
// Source: docs/architecture/tech-stack.md - 15-digit precision
func ValidatePrecision(result *big.Float, expectedPrecision int) error {
	precision := result.Prec()
	minPrecision := uint(expectedPrecision * 4) // Approximate bits needed for decimal digits

	if precision < minPrecision {
		return fmt.Errorf("insufficient precision: got %d bits, need at least %d bits for %d decimal digits",
			precision, minPrecision, expectedPrecision)
	}

	return nil
}
