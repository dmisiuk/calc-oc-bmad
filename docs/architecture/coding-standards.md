# Coding Standards

## Critical Fullstack Rules

- **Type Safety:** Always use Go's type system to prevent runtime errors
- **Error Handling:** Always handle errors explicitly, never ignore them
- **Input Validation:** Validate all user input before processing
- **Memory Management:** Use Go's garbage collection effectively, avoid memory leaks
- **Concurrency:** Use goroutines and channels safely, avoid race conditions
- **Testing:** Write tests for all public functions and complex logic
- **Configuration:** Never hardcode configuration values, use config system
- **Logging:** Use structured logging with appropriate log levels

## Naming Conventions

| Element | Go Convention | Example |
|----------|---------------|---------|
| Functions | camelCase | `calculateResult()` |
| Variables | camelCase | `maxHistory` |
| Constants | SCREAMING_SNAKE_CASE | `MAX_PRECISION` |
| Structs | PascalCase | `CalculationEngine` |
| Interfaces | PascalCase with -er suffix | `Parser`, `Calculator` |
| Files | snake_case | `calculation_engine.go` |
| Directories | snake_case | `terminal_components/` |
| Packages | lowercase | `calculator`, `terminal` |
