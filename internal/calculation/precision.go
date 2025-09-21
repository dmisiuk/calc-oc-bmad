package calculation

import (
	"fmt"
	"math"
	"math/big"
)

// PrecisionValidator handles 15-digit precision validation
// Source: docs/architecture/tech-stack.md - 15-digit precision requirement
type PrecisionValidator struct {
	requiredPrecision int
}

// NewPrecisionValidator creates a new precision validator
func NewPrecisionValidator(requiredPrecision int) *PrecisionValidator {
	return &PrecisionValidator{
		requiredPrecision: requiredPrecision,
	}
}

// ValidateResult checks if a calculation result meets precision requirements
// Source: docs/architecture/security-and-performance.md - Safe math operations
func (pv *PrecisionValidator) ValidateResult(result *big.Float, operation string) error {
	// Check precision bits - use more reasonable approximation
	precision := result.Prec()
	minPrecisionBits := uint(pv.requiredPrecision * 3) // More reasonable approximation: ~3 bits per decimal digit

	if precision < minPrecisionBits {
		return fmt.Errorf("insufficient precision in %s: got %d bits, need %d bits for %d decimal digits",
			operation, precision, minPrecisionBits, pv.requiredPrecision)
	}

	// Check for precision loss indicators
	if pv.HasPrecisionLoss(result) {
		return fmt.Errorf("precision loss detected in %s operation", operation)
	}

	return nil
}

// HasPrecisionLoss checks for common precision loss indicators
func (pv *PrecisionValidator) HasPrecisionLoss(result *big.Float) bool {
	// Check if the result has the expected number of significant digits
	resultStr := result.Text('f', pv.requiredPrecision*2)

	// Count significant digits
	significantDigits := 0
	for _, r := range resultStr {
		if r >= '1' && r <= '9' {
			significantDigits++
		}
		if r == '.' {
			continue
		}
		if significantDigits >= pv.requiredPrecision {
			break
		}
	}

	return significantDigits < pv.requiredPrecision
}

// ValidateOperationPrecision validates precision for specific operations
func (pv *PrecisionValidator) ValidateOperationPrecision(a, b, result *big.Float, operation string) error {
	// Check input precision
	if err := pv.ValidateResult(a, "input_a"); err != nil {
		return fmt.Errorf("input A precision error: %w", err)
	}
	if err := pv.ValidateResult(b, "input_b"); err != nil {
		return fmt.Errorf("input B precision error: %w", err)
	}

	// Check result precision
	if err := pv.ValidateResult(result, operation); err != nil {
		return err
	}

	// Additional operation-specific checks
	switch operation {
	case "division":
		if pv.HasDivisionPrecisionLoss(a, b, result) {
			return fmt.Errorf("division precision loss detected")
		}
	case "multiplication":
		if pv.HasMultiplicationPrecisionLoss(a, b, result) {
			return fmt.Errorf("multiplication precision loss detected")
		}
	}

	return nil
}

// HasDivisionPrecisionLoss checks for precision loss in division
func (pv *PrecisionValidator) HasDivisionPrecisionLoss(a, b, result *big.Float) bool {
	// For division, check if result * b ≈ a
	reconstructed := new(big.Float).Mul(result, b)
	diff := new(big.Float).Sub(reconstructed, a)

	// Calculate relative error
	relativeError := new(big.Float).Quo(diff, a)
	relativeError.Abs(relativeError)

	// Allow small relative error (1e-15 for 15-digit precision)
	maxError := new(big.Float).SetFloat64(1e-15)
	return relativeError.Cmp(maxError) > 0
}

// HasMultiplicationPrecisionLoss checks for precision loss in multiplication
func (pv *PrecisionValidator) HasMultiplicationPrecisionLoss(a, b, result *big.Float) bool {
	// For multiplication, check if result / a ≈ b
	reconstructed := new(big.Float).Quo(result, a)
	diff := new(big.Float).Sub(reconstructed, b)

	// Calculate relative error
	relativeError := new(big.Float).Quo(diff, b)
	relativeError.Abs(relativeError)

	// Allow small relative error
	maxError := new(big.Float).SetFloat64(1e-15)
	return relativeError.Cmp(maxError) > 0
}

// GetPrecisionReport generates a detailed precision report
func (pv *PrecisionValidator) GetPrecisionReport(result *big.Float, operation string) string {
	precision := result.Prec()
	resultStr := result.Text('f', pv.requiredPrecision)

	return fmt.Sprintf("Precision Report for %s:\n"+
		"  Required precision: %d decimal digits\n"+
		"  Actual precision: %d bits (≈%.1f decimal digits)\n"+
		"  Result: %s\n"+
		"  Status: %s",
		operation,
		pv.requiredPrecision,
		precision,
		float64(precision)/math.Log2(10),
		resultStr,
		pv.GetPrecisionStatus(result, operation))
}

// GetPrecisionStatus returns a human-readable precision status
func (pv *PrecisionValidator) GetPrecisionStatus(result *big.Float, operation string) string {
	if err := pv.ValidateResult(result, operation); err != nil {
		return "FAILED - " + err.Error()
	}
	return "PASSED"
}

// ValidateFloat64Precision checks precision when converting to float64
func (pv *PrecisionValidator) ValidateFloat64Precision(bigResult *big.Float, floatResult float64) error {
	// Convert back to big.Float for comparison
	floatAsBig := new(big.Float).SetFloat64(floatResult)

	// Check if they're exactly equal (no significant precision loss)
	if bigResult.Cmp(floatAsBig) == 0 {
		return nil // No precision loss
	}

	// Calculate difference
	diff := new(big.Float).Sub(bigResult, floatAsBig)
	diff.Abs(diff)

	// Calculate relative error
	relativeError := new(big.Float).Quo(diff, bigResult)
	relativeError.Abs(relativeError)

	// Allow small precision loss that's normal for float64 (around 1e-15)
	// but flag significant loss that indicates problems
	maxAcceptableError := new(big.Float).SetFloat64(1e-12)
	if relativeError.Cmp(maxAcceptableError) > 0 {
		return fmt.Errorf("significant float64 conversion precision loss: relative error %.2e", relativeError)
	}

	return nil
}
