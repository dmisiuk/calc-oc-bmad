# Unified Project Structure

```
calculator/
├── .github/                    # CI/CD workflows
│   └── workflows/
│       ├── ci.yaml
│       └── release.yaml
├── cmd/
│   └── calculator/              # Main application
│       └── main.go
├── internal/                   # Internal application code
│   ├── calculation/            # Calculation engine
│   │   ├── engine.go
│   │   ├── operations.go
│   │   └── validator.go
│   ├── terminal/               # Terminal UI components
│   │   ├── components/
│   │   ├── handlers/
│   │   ├── themes/
│   │   └── utils/
│   ├── config/                  # Configuration management
│   │   ├── config.go
│   │   └── loader.go
│   ├── history/                 # History management
│   │   ├── manager.go
│   │   └── storage.go
│   ├── parser/                  # Input parsing
│   │   ├── parser.go
│   │   └── validator.go
│   └── models/                  # Data models
│       ├── calculation.go
│       ├── config.go
│       └── history.go
├── pkg/                        # Public packages
│   ├── calculator/             # Calculator library
│   └── terminal/               # Terminal utilities
├── test/                       # Test files
│   ├── integration/
│   ├── unit/
│   └── e2e/
├── docs/                       # Documentation
│   ├── prd.md
│   ├── front-end-spec.md
│   └── fullstack-architecture.md
├── scripts/                    # Build and deployment scripts
│   ├── build.sh
│   ├── test.sh
│   └── release.sh
├── configs/                    # Configuration files
│   ├── default.yaml
│   └── themes/
├── go.mod                      # Go modules
├── go.sum                      # Go modules checksum
├── Makefile                    # Build automation
└── README.md                   # Project documentation
```
