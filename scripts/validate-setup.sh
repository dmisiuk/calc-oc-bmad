#!/bin/bash

# Setup Validation Script for Calculator Project
# Validates Go version and project directory structure

set -e

echo "🔍 Validating Calculator Project Setup..."
echo

# Check Go version
echo "📦 Checking Go version..."
GO_VERSION=$(go version | grep -oE 'go[0-9]+\.[0-9]+\.[0-9]+' | sed 's/go//')
REQUIRED_VERSION="1.21"

if ! [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" = "$REQUIRED_VERSION" ]; then
    echo "❌ Go version $GO_VERSION is below required $REQUIRED_VERSION"
    exit 1
fi

echo "✅ Go version $GO_VERSION meets requirement (>= $REQUIRED_VERSION)"
echo

# Check go.mod
echo "📄 Checking go.mod..."
if [ ! -f "go.mod" ]; then
    echo "❌ go.mod not found"
    exit 1
fi

MODULE_NAME=$(grep "^module" go.mod | awk '{print $2}')
if [ "$MODULE_NAME" != "calculator" ]; then
    echo "❌ Module name is '$MODULE_NAME', expected 'calculator'"
    exit 1
fi

echo "✅ go.mod exists with correct module name: $MODULE_NAME"
echo

# Check directory structure
echo "📁 Checking directory structure..."

REQUIRED_DIRS=(
    "cmd/calculator"
    "internal/calculation"
    "internal/terminal"
    "internal/config"
    "internal/history"
    "internal/parser"
    "internal/models"
    "pkg/calculator"
    "pkg/terminal"
    "test/unit"
    "test/integration"
    "test/e2e"
    "docs"
    "scripts"
    "configs"
)

MISSING_DIRS=()

for dir in "${REQUIRED_DIRS[@]}"; do
    if [ ! -d "$dir" ]; then
        MISSING_DIRS+=("$dir")
    fi
done

if [ ${#MISSING_DIRS[@]} -ne 0 ]; then
    echo "❌ Missing directories:"
    printf '  - %s\n' "${MISSING_DIRS[@]}"
    exit 1
fi

echo "✅ All required directories exist"
echo

# Check key files
echo "📄 Checking key files..."

REQUIRED_FILES=(
    "cmd/calculator/main.go"
    "README.md"
    ".gitignore"
    "go.mod"
)

MISSING_FILES=()

for file in "${REQUIRED_FILES[@]}"; do
    if [ ! -f "$file" ]; then
        MISSING_FILES+=("$file")
    fi
done

if [ ${#MISSING_FILES[@]} -ne 0 ]; then
    echo "❌ Missing files:"
    printf '  - %s\n' "${MISSING_FILES[@]}"
    exit 1
fi

echo "✅ All required files exist"
echo

# Check Git repository
echo "🔧 Checking Git repository..."
if [ ! -d ".git" ]; then
    echo "❌ Git repository not initialized"
    exit 1
fi

if ! git remote get-url origin > /dev/null 2>&1; then
    echo "❌ Git remote 'origin' not configured"
    exit 1
fi

echo "✅ Git repository initialized and remote configured"
echo

# Check build
echo "🔨 Checking build..."
if ! go build ./cmd/calculator > /dev/null 2>&1; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Project builds successfully"
echo

# Check tests
echo "🧪 Checking tests..."
if ! go test ./... > /dev/null 2>&1; then
    echo "❌ Tests failed"
    exit 1
fi

echo "✅ All tests pass"
echo

echo "🎉 Setup validation complete! All checks passed."
echo
echo "Project is ready for development."