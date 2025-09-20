# Data Storage

## File-based History Storage

History will be stored in JSON format for simplicity and portability:

**History File Format (JSON):**
```json
{
  "session_id": "session-123",
  "created_at": "2025-09-19T22:30:00Z",
  "last_updated": "2025-09-19T22:35:00Z",
  "calculations": [
    {
      "id": "calc-1",
      "expression": "2 + 2",
      "result": 4.0,
      "timestamp": "2025-09-19T22:30:15Z",
      "operation": "add",
      "operands": [2.0, 2.0]
    },
    {
      "id": "calc-2", 
      "expression": "10 * 5",
      "result": 50.0,
      "timestamp": "2025-09-19T22:31:00Z",
      "operation": "multiply",
      "operands": [10.0, 5.0]
    }
  ]
}
```

**Configuration File Format (YAML):**
```yaml
precision: 4
max_history: 100
auto_save: true
theme: default
debug_mode: false
batch_mode: false
output_format: text
scientific_mode: false  # For future scientific calculations
```

**Storage Structure:**
```
~/.calculator/
├── config.yaml          # User configuration
├── history/             # History files directory
│   ├── session-123.json
│   └── session-456.json
└── logs/                # Optional log files
```
