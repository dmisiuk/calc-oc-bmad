# Data Models

## Calculation

**Purpose:** Represents a mathematical calculation with its inputs, operations, and results

**Key Attributes:**
- id: string - Unique identifier for the calculation
- expression: string - The mathematical expression entered by the user
- result: float64 - The calculated result
- timestamp: time.Time - When the calculation was performed
- operation: string - The type of operation (add, subtract, multiply, divide)
- operands: []float64 - The numbers used in the calculation
- error: string - Error message if calculation failed

**Go Struct:**
```go
type Calculation struct {
    ID        string      `json:"id"`
    Expression string      `json:"expression"`
    Result    float64     `json:"result"`
    Timestamp time.Time   `json:"timestamp"`
    Operation string      `json:"operation"`
    Operands  []float64   `json:"operands"`
    Error     string      `json:"error,omitempty"`
}
```

**Relationships:**
- A Calculation belongs to a History session
- Multiple calculations can be part of a batch operation
- Calculations can be exported to different formats

## History

**Purpose:** Manages the session history of calculations performed by the user

**Key Attributes:**
- id: string - Unique identifier for the history session
- session_id: string - Identifier for the current terminal session
- calculations: []Calculation - List of calculations in this session
- created_at: time.Time - When the session started
- last_updated: time.Time - When the session was last modified
- size: int - Number of calculations in the session

**Go Struct:**
```go
type History struct {
    ID           string         `json:"id"`
    SessionID    string         `json:"session_id"`
    Calculations []Calculation  `json:"calculations"`
    CreatedAt    time.Time      `json:"created_at"`
    LastUpdated time.Time      `json:"last_updated"`
    Size         int            `json:"size"`
}
```

**Relationships:**
- A History contains multiple Calculations
- A History belongs to a user session
- History can be persisted to file or database

## Configuration

**Purpose:** Stores application configuration and user preferences

**Key Attributes:**
- precision: int - Number of decimal places for output
- max_history: int - Maximum number of calculations to keep in memory
- auto_save: bool - Whether to automatically save history
- theme: string - Color theme for terminal output
- debug_mode: bool - Enable debug logging
- batch_mode: bool - Default to batch processing mode
- output_format: string - Format for batch output (json, csv, text)
- scientific_mode: bool - Enable scientific calculation features (future)

**Go Struct:**
```go
type Configuration struct {
    Precision       int     `yaml:"precision" json:"precision"`
    MaxHistory      int     `yaml:"max_history" json:"max_history"`
    AutoSave        bool    `yaml:"auto_save" json:"auto_save"`
    Theme           string  `yaml:"theme" json:"theme"`
    DebugMode       bool    `yaml:"debug_mode" json:"debug_mode"`
    BatchMode       bool    `yaml:"batch_mode" json:"batch_mode"`
    OutputFormat    string  `yaml:"output_format" json:"output_format"`
    ScientificMode  bool    `yaml:"scientific_mode" json:"scientific_mode"`
}
```

**Relationships:**
- Configuration is used by the Calculation Service for formatting
- Configuration is loaded/saved by the Configuration Manager
- Configuration affects UI display preferences
