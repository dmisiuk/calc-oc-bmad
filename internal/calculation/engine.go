package calculation

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

// CalculationEngine provides high-precision arithmetic operations
// Source: docs/architecture/components.md - CalculationEngine component
type CalculationEngine struct{}

// NewCalculationEngine creates a new instance of the calculation engine
func NewCalculationEngine() *CalculationEngine {
	return &CalculationEngine{}
}

// Calculate parses and evaluates a mathematical expression with 15-digit precision
// Supports: addition (+), subtraction (-), multiplication (*), division (/)
// Source: docs/architecture/components.md - Calculate interface
func (ce *CalculationEngine) Calculate(expression string) (float64, error) {
	// Validate input first
	if err := ce.Validate(expression); err != nil {
		return 0, err
	}

	// Parse the expression using simplified parser
	num1, op, num2, err := ce.parseSimpleExpression(expression)
	if err != nil {
		return 0, fmt.Errorf("expression parsing failed: %w", err)
	}

	result := new(big.Float)

	switch op {
	case "+":
		result.Add(num1, num2)
	case "-":
		result.Sub(num1, num2)
	case "*":
		result.Mul(num1, num2)
	case "/":
		if num2.Sign() == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		result.Quo(num1, num2)
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}

	// Convert back to float64 with precision handling
	floatResult, _ := result.Float64()

	// Validate precision - ensure result has reasonable precision for the operation
	// Source: docs/architecture/tech-stack.md - math/big for precision
	if result.Prec() < 50 {
		return floatResult, fmt.Errorf("insufficient precision in calculation result")
	}

	return floatResult, nil
}

// parseSimpleExpression provides a simplified, robust expression parser
// Supports format: "number operator number" with flexible whitespace
func (ce *CalculationEngine) parseSimpleExpression(expression string) (*big.Float, string, *big.Float, error) {
	// Trim whitespace
	expr := strings.TrimSpace(expression)

	// Split by whitespace to get parts
	parts := strings.Fields(expr)
	if len(parts) != 3 {
		return nil, "", nil, fmt.Errorf("invalid format: expected 'number operator number'")
	}

	num1Str, op, num2Str := parts[0], parts[1], parts[2]

	// Validate operator
	validOps := []string{"+", "-", "*", "/"}
	isValidOp := false
	for _, validOp := range validOps {
		if op == validOp {
			isValidOp = true
			break
		}
	}
	if !isValidOp {
		return nil, "", nil, fmt.Errorf("unsupported operator: %s", op)
	}

	// Parse numbers
	num1, err := ce.parseBigFloat(num1Str)
	if err != nil {
		return nil, "", nil, fmt.Errorf("invalid first number: %w", err)
	}

	num2, err := ce.parseBigFloat(num2Str)
	if err != nil {
		return nil, "", nil, fmt.Errorf("invalid second number: %w", err)
	}

	return num1, op, num2, nil
}

// Validate checks if the expression is syntactically valid
// Source: docs/architecture/components.md - Validate interface
func (ce *CalculationEngine) Validate(expression string) error {
	if strings.TrimSpace(expression) == "" {
		return fmt.Errorf("expression cannot be empty")
	}

	// Use the same parser as Calculate for consistency
	_, op, num2, err := ce.parseSimpleExpression(expression)
	if err != nil {
		return err
	}

	// Check for division by zero
	if op == "/" && num2.Sign() == 0 {
		return fmt.Errorf("division by zero detected")
	}

	return nil
}

// GetSupportedOperations returns the list of supported arithmetic operations
// Source: docs/architecture/components.md - GetSupportedOperations interface
func (ce *CalculationEngine) GetSupportedOperations() []string {
	return []string{"+", "-", "*", "/"}
}

// parseBigFloat converts a string to big.Float with error handling
func (ce *CalculationEngine) parseBigFloat(s string) (*big.Float, error) {
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
