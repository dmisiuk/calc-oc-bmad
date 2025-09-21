package performance_test

import (
	"math/big"
	"testing"
	"time"

	"calculator/internal/calculation"
)

// BenchmarkCalculationEngine_Add benchmarks addition performance
func BenchmarkCalculationEngine_Add(b *testing.B) {
	engine := calculation.NewCalculationEngine()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := engine.Calculate("123456789.123456789 + 987654321.987654321")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkCalculationEngine_Subtract benchmarks subtraction performance
func BenchmarkCalculationEngine_Subtract(b *testing.B) {
	engine := calculation.NewCalculationEngine()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := engine.Calculate("1000000000.000000000 - 123456789.123456789")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkCalculationEngine_Multiply benchmarks multiplication performance
func BenchmarkCalculationEngine_Multiply(b *testing.B) {
	engine := calculation.NewCalculationEngine()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := engine.Calculate("12345.678901234567890 * 98765.432109876543210")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkCalculationEngine_Divide benchmarks division performance
func BenchmarkCalculationEngine_Divide(b *testing.B) {
	engine := calculation.NewCalculationEngine()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := engine.Calculate("1000000000.000000000 / 3.141592653589793")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkOperations_Add benchmarks the Add function directly
func BenchmarkOperations_Add(b *testing.B) {
	a, _ := new(big.Float).SetString("123456789.123456789")
	c, _ := new(big.Float).SetString("987654321.987654321")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := calculation.Add(a, c)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkOperations_Subtract benchmarks the Subtract function directly
func BenchmarkOperations_Subtract(b *testing.B) {
	a, _ := new(big.Float).SetString("1000000000.000000000")
	c, _ := new(big.Float).SetString("123456789.123456789")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := calculation.Subtract(a, c)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkOperations_Multiply benchmarks the Multiply function directly
func BenchmarkOperations_Multiply(b *testing.B) {
	a, _ := new(big.Float).SetString("12345.678901234567890")
	c, _ := new(big.Float).SetString("98765.432109876543210")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := calculation.Multiply(a, c)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkOperations_Divide benchmarks the Divide function directly
func BenchmarkOperations_Divide(b *testing.B) {
	a, _ := new(big.Float).SetString("1000000000.000000000")
	c, _ := new(big.Float).SetString("3.141592653589793")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := calculation.Divide(a, c)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// TestPerformanceRequirements validates the 50ms performance target
func TestPerformanceRequirements(t *testing.T) {
	engine := calculation.NewCalculationEngine()

	testCases := []string{
		"123456789.123456789 + 987654321.987654321",
		"1000000000.000000000 - 123456789.123456789",
		"12345.678901234567890 * 98765.432109876543210",
		"1000000000.000000000 / 3.141592653589793",
	}

	maxDuration := 50 * time.Millisecond

	for _, expr := range testCases {
		t.Run("perf_"+expr, func(t *testing.T) {
			start := time.Now()
			_, err := engine.Calculate(expr)
			duration := time.Since(start)

			if err != nil {
				t.Fatalf("calculation failed: %v", err)
			}

			if duration > maxDuration {
				t.Errorf("calculation took %v, exceeds 50ms limit", duration)
			}

			t.Logf("Expression: %s, Duration: %v", expr, duration)
		})
	}
}

// BenchmarkPrecisionValidation benchmarks precision validation
func BenchmarkPrecisionValidation(b *testing.B) {
	validator := calculation.NewPrecisionValidator(15)
	result, _ := new(big.Float).SetString("123456789.123456789")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := validator.ValidateResult(result, "test")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkExpressionParsing benchmarks expression parsing
func BenchmarkExpressionParsing(b *testing.B) {
	engine := calculation.NewCalculationEngine()
	expression := "123456789.123456789 + 987654321.987654321"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := engine.Validate(expression)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkMemoryUsage benchmarks memory usage patterns
func BenchmarkMemoryUsage(b *testing.B) {
	engine := calculation.NewCalculationEngine()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := engine.Calculate("123456789.123456789 + 987654321.987654321")
		if err != nil {
			b.Fatal(err)
		}
	}
}
