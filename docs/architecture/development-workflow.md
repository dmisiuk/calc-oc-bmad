# Development Workflow

## Local Development Setup

**Prerequisites:**
```bash
# Install Go 1.21+
curl -OL https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Set up environment
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verify installation
go version
```

**Initial Setup:**
```bash
# Clone repository
git clone https://github.com/username/calculator.git
cd calculator

# Install dependencies
go mod tidy

# Build for testing
go build -o calculator cmd/calculator/main.go

# Run tests
go test ./...

# Run application
./calculator
```

**Development Commands:**
```bash
# Start calculator in interactive mode
go run cmd/calculator/main.go

# Run tests
go test ./...

# Run with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. ./...

# Build for current platform
go build -o calculator cmd/calculator/main.go

# Build for multiple platforms
./scripts/build.sh

# Run linting
golangci-lint run

# Format code
go fmt ./...
```

## Environment Configuration

**Required Environment Variables:**
```bash
# Application settings
export CALCULATOR_PRECISION=4
export CALCULATOR_MAX_HISTORY=100
export CALCULATOR_AUTO_SAVE=true
export CALCULATOR_THEME=default
export CALCULATOR_DEBUG=false
export CALCULATOR_BATCH_MODE=false
export CALCULATOR_OUTPUT_FORMAT=text

# File paths
export CALCULATOR_CONFIG_PATH=~/.calculator/config.yaml
export CALCULATOR_HISTORY_PATH=~/.calculator/history.db
export CALCULATOR_LOG_PATH=~/.calculator/logs/
```
