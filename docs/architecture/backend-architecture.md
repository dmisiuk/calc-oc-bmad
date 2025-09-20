# Backend Architecture

## Service Architecture

**Service Organization:**
```
internal/
├── calculation/
│   ├── engine.go           # Core calculation engine
│   ├── operations.go       # Mathematical operations
│   ├── validator.go        # Expression validation
│   └── precision.go        # Number formatting
├── history/
│   ├── manager.go          # History management
│   ├── storage.go          # File-based storage
│   ├── session.go          # Session management
│   └── export.go           # History export functionality
├── config/
│   ├── manager.go          # Configuration management
│   ├── loader.go           # Configuration loading
│   ├── saver.go            # Configuration saving
│   └── validator.go        # Configuration validation
├── parser/
│   ├── expression.go       # Expression parsing
│   ├── tokenizer.go        # Token parsing
│   ├── builder.go          # Expression building
│   └── evaluator.go        # Expression evaluation
└── utils/
    ├── math.go             # Mathematical utilities
    ├── formatting.go       # Number formatting
    ├── error.go            # Error handling
    └── validation.go       # Input validation
```

**Calculation Engine Template:**
```go
package calculation

import (
    "math/big"
    "strconv"
    "strings"
)

type CalculationEngine struct {
    precision int
}

func NewCalculationEngine(precision int) *CalculationEngine {
    return &CalculationEngine{
        precision: precision,
    }
}

func (ce *CalculationEngine) Calculate(expression string) (*big.Float, error) {
    // Parse expression into operands and operator
    operands, operator, err := ce.parseExpression(expression)
    if err != nil {
        return nil, err
    }
    
    // Perform calculation with high precision
    result, err := ce.performOperation(operands[0], operands[1], operator)
    if err != nil {
        return nil, err
    }
    
    // Round to specified precision
    return result.SetMode(big.ToNearestEven, uint(ce.precision)), nil
}
```

## File Storage Architecture

**File Storage Design:**
```go
package storage

import (
    "encoding/json"
    "os"
    "path/filepath"
    "time"
)

type HistoryStorage struct {
    basePath string
}

func NewHistoryStorage(basePath string) *HistoryStorage {
    return &HistoryStorage{
        basePath: basePath,
    }
}

func (hs *HistoryStorage) SaveSession(session *models.History) error {
    // Ensure directory exists
    if err := os.MkdirAll(hs.basePath, 0755); err != nil {
        return err
    }
    
    // Create filename
    filename := filepath.Join(hs.basePath, session.SessionID+".json")
    
    // Marshal to JSON
    data, err := json.MarshalIndent(session, "", "  ")
    if err != nil {
        return err
    }
    
    // Write to file
    return os.WriteFile(filename, data, 0644)
}

func (hs *HistoryStorage) LoadSession(sessionID string) (*models.History, error) {
    filename := filepath.Join(hs.basePath, sessionID+".json")
    
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    
    var session models.History
    if err := json.Unmarshal(data, &session); err != nil {
        return nil, err
    }
    
    return &session, nil
}
```

**Data Access Layer:**
```go
package storage

type HistoryStorage interface {
    SaveSession(session *models.History) error
    LoadSession(sessionID string) (*models.History, error)
    ListSessions() ([]string, error)
    DeleteSession(sessionID string) error
    ExportSession(sessionID string, format string) (string, error)
}

type FileHistoryStorage struct {
    basePath string
}

func NewFileHistoryStorage(basePath string) *FileHistoryStorage {
    return &FileHistoryStorage{
        basePath: basePath,
    }
}
```

## Application Security

**Security Considerations:**
Since this is a local terminal application, traditional authentication is not required. However, we should consider:

**File Security:**
```go
package security

import (
    "os"
    "path/filepath"
)

type FileSecurity struct {
    configPath string
    historyPath string
}

func (fs *FileSecurity) EnsureSecurePermissions() error {
    // Ensure config file has appropriate permissions
    if err := os.Chmod(fs.configPath, 0600); err != nil {
        return err
    }
    
    // Ensure history directory has appropriate permissions
    if err := os.Chmod(fs.historyPath, 0700); err != nil {
        return err
    }
    
    return nil
}

func (fs *FileSecurity) ValidateFilePath(path string) bool {
    // Prevent directory traversal attacks
    cleanPath := filepath.Clean(path)
    return filepath.HasPrefix(cleanPath, filepath.Dir(fs.configPath))
}
```

**Input Validation:**
```go
package validation

import (
    "regexp"
    "strings"
)

type InputValidator struct {
    // Regular expression for valid mathematical expressions
    expressionRegex *regexp.Regexp
}

func NewInputValidator() *InputValidator {
    return &InputValidator{
        expressionRegex: regexp.MustCompile(`^[0-9\s\+\-\*\/\(\)\.]+$`),
    }
}

func (iv *InputValidator) ValidateExpression(expression string) error {
    // Check for potentially dangerous characters
    if !iv.expressionRegex.MatchString(expression) {
        return errors.New("invalid characters in expression")
    }
    
    // Check for division by zero
    if strings.Contains(expression, "/ 0") || strings.Contains(expression, "/0") {
        return errors.New("division by zero")
    }
    
    return nil
}
```
