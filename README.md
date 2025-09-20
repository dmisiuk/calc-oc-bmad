# Calculator

A terminal-based calculator application built with Go, featuring a TUI (Text User Interface) for intuitive calculations.

## Overview

This calculator provides a visual interface for performing mathematical calculations directly in the terminal. It supports basic arithmetic operations, advanced functions, and maintains calculation history.

## Installation

### Prerequisites

- Go 1.21 or higher
- Terminal with ANSI support

### Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd calculator
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build ./cmd/calculator
   ```

## Usage

### Basic Usage

Run the calculator:
```bash
./calculator
```

### Examples

- Addition: `2 + 3`
- Subtraction: `10 - 4`
- Multiplication: `5 * 6`
- Division: `15 / 3`
- Complex expression: `(2 + 3) * 4`

## Development

### Setup Development Environment

1. Ensure Go 1.21+ is installed
2. Clone and setup as above
3. Run tests:
   ```bash
   go test ./...
   ```

### Project Structure

- `cmd/calculator/` - Main application entry point
- `internal/` - Internal packages
  - `calculation/` - Calculation engine
  - `terminal/` - TUI components
  - `config/` - Configuration management
  - `history/` - Calculation history
  - `parser/` - Expression parsing
  - `models/` - Data models
- `pkg/` - Public packages
- `test/` - Test files

### Testing

Run all tests:
```bash
go test ./...
```

Run specific test types:
```bash
go test ./test/unit/      # Unit tests
go test ./test/integration/  # Integration tests
go test ./test/e2e/       # End-to-end tests
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

[Add license information]