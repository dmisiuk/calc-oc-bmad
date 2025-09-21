package calculation_test

import (
	"math/big"
	"testing"

	"calculator/internal/calculation"
	"calculator/test"
)

func TestPrecisionValidator(t *testing.T) {
	validator := calculation.NewPrecisionValidator(15)

	t.Run("NewPrecisionValidator", func(t *testing.T) {
		if validator == nil {
			t.Error("NewPrecisionValidator returned nil")
		}
		// Test that validator was created successfully by using it
		result, _ := new(big.Float).SetString("123456789.123456789")
		err := validator.ValidateResult(result, "test")
		if err != nil {
			t.Errorf("validator should work after creation: %v", err)
		}
	})

	t.Run("ValidateResult", func(t *testing.T) {
		result, _ := new(big.Float).SetString("123456789.123456789")

		err := validator.ValidateResult(result, "test")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("HasPrecisionLoss", func(t *testing.T) {
		// Test with sufficient precision
		result, _ := new(big.Float).SetString("123456789.123456789")
		if validator.HasPrecisionLoss(result) {
			t.Error("expected no precision loss for high precision number")
		}

		// Test with insufficient precision
		lowPrecision, _ := new(big.Float).SetString("1.23")
		if !validator.HasPrecisionLoss(lowPrecision) {
			t.Error("expected precision loss for low precision number")
		}
	})

	t.Run("ValidateOperationPrecision", func(t *testing.T) {
		// Test successful validation
		a, _ := new(big.Float).SetString("123456789.123456789")
		b, _ := new(big.Float).SetString("987654321.987654321")
		result := new(big.Float).Add(a, b)

		err := validator.ValidateOperationPrecision(a, b, result, "addition")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Test with invalid input precision (should fail)
		lowPrecisionA, _ := new(big.Float).SetString("1.23")
		lowPrecisionB, _ := new(big.Float).SetString("4.56")
		result2 := new(big.Float).Add(lowPrecisionA, lowPrecisionB)

		err = validator.ValidateOperationPrecision(lowPrecisionA, lowPrecisionB, result2, "addition")
		if err == nil {
			t.Error("expected error for low precision inputs")
		}
	})

	t.Run("HasDivisionPrecisionLoss", func(t *testing.T) {
		a, _ := new(big.Float).SetString("10")
		b, _ := new(big.Float).SetString("3")
		result := new(big.Float).Quo(a, b)

		// This should not have precision loss
		if validator.HasDivisionPrecisionLoss(a, b, result) {
			t.Error("expected no division precision loss")
		}
	})

	t.Run("HasMultiplicationPrecisionLoss", func(t *testing.T) {
		a, _ := new(big.Float).SetString("123456789.123456789")
		b, _ := new(big.Float).SetString("2")
		result := new(big.Float).Mul(a, b)

		// This should not have precision loss
		if validator.HasMultiplicationPrecisionLoss(a, b, result) {
			t.Error("expected no multiplication precision loss")
		}
	})

	t.Run("GetPrecisionReport", func(t *testing.T) {
		result, _ := new(big.Float).SetString("123456789.123456789")
		report := validator.GetPrecisionReport(result, "test")

		if report == "" {
			t.Error("expected non-empty precision report")
		}

		if !test.ContainsString(report, "PASSED") {
			t.Error("expected PASSED status in report")
		}
	})

	t.Run("GetPrecisionStatus", func(t *testing.T) {
		// Test passing case
		result, _ := new(big.Float).SetString("123456789.123456789")
		status := validator.GetPrecisionStatus(result, "test")

		if status != "PASSED" {
			t.Errorf("expected PASSED status, got %s", status)
		}

		// Test failing case
		lowPrecision, _ := new(big.Float).SetString("1.23")
		status2 := validator.GetPrecisionStatus(lowPrecision, "test")

		if !test.ContainsString(status2, "FAILED") {
			t.Errorf("expected FAILED status, got %s", status2)
		}
	})

	t.Run("ValidateFloat64Precision", func(t *testing.T) {
		// Test passing case
		bigResult, _ := new(big.Float).SetString("1.23456789012345")
		floatResult, _ := bigResult.Float64()

		err := validator.ValidateFloat64Precision(bigResult, floatResult)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Test case with acceptable precision loss (normal for float64)
		// This should NOT produce an error since the loss is within float64 limits
		veryPrecise, _ := new(big.Float).SetString("1.234567890123456789012345678901234567890")
		floatResult2, _ := veryPrecise.Float64()

		err = validator.ValidateFloat64Precision(veryPrecise, floatResult2)
		if err != nil {
			t.Errorf("unexpected error for acceptable float64 precision loss: %v", err)
		}
	})
}
