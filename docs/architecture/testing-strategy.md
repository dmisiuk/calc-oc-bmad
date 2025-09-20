# Testing Strategy

## Testing Pyramid

```
E2E Tests (10%)
/        \
Integration Tests (30%)
/            \
Frontend Unit Tests (30%)  Backend Unit Tests (30%)
```

## Test Organization

**Frontend Tests:**
```
test/unit/terminal/
├── components/
│   ├── prompt_test.go
│   ├── display_test.go
│   └── history_test.go
├── handlers/
│   ├── input_test.go
│   └── navigation_test.go
└── utils/
    ├── ansi_test.go
    └── screen_test.go
```

**Backend Tests:**
```
test/unit/calculation/
├── engine_test.go
├── operations_test.go
├── validator_test.go
└── parser_test.go

test/integration/
├── end_to_end_test.go
├── history_test.go
└── config_test.go
```

**E2E Tests:**
```
test/e2e/
├── interactive_test.go
├── batch_test.go
├── history_test.go
└── config_test.go
```

## Test Examples

**Frontend Component Test:**
```go
func TestDisplayComponent_ShowResult(t *testing.T) {
    config := &config.Configuration{Precision: 2}
    display := NewDisplayComponent(config)
    
    result := 123.456
    expression := "100 + 23.456"
    
    output := display.FormatResult(result, expression)
    
    expected := "100 + 23.456 = 123.46"
    if output != expected {
        t.Errorf("Expected %s, got %s", expected, output)
    }
}
```

**Backend API Test:**
```go
func TestCalculationEngine_Add(t *testing.T) {
    engine := NewCalculationEngine()
    
    result, err := engine.Calculate("2 + 3")
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    if result != 5 {
        t.Errorf("Expected 5, got %f", result)
    }
}
```

**E2E Test:**
```go
func TestInteractiveMode(t *testing.T) {
    app := NewCalculatorApp()
    
    // Start interactive mode
    go app.RunInteractive()
    
    // Simulate user input
    app.SendInput("2 + 2")
    app.SendInput("exit")
    
    // Check output
    output := app.GetOutput()
    if !strings.Contains(output, "4") {
        t.Errorf("Expected result 4 in output")
    }
}
```
