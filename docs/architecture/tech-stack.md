# Tech Stack

This is the DEFINITIVE technology selection for the entire project. All development must use these exact versions.

## Technology Stack Table

| Category | Technology | Version | Purpose | Rationale |
|----------|------------|---------|---------|-----------|
| Frontend Language | Go | 1.21+ | Terminal UI development | Cross-platform compilation, excellent performance, robust standard library |
| Frontend Framework | Tview | Latest | Terminal UI components | Popular TUI framework with rich widgets, mouse support, and cross-platform compatibility |
| UI Component Library | Custom ANSI components | v1.0 | Terminal UI components | Minimal dependencies, terminal-native appearance |
| State Management | Go channels & structs | Native | Application state management | Simple, performant, idiomatic Go patterns |
| Backend Language | Go | 1.21+ | Calculation engine and business logic | Type safety, excellent performance, cross-platform |
| Backend Framework | Standard library | Native | Core application framework | Minimal dependencies, robust and well-tested |
| API Style | TUI Events | v1.0 | User interface protocol | Visual calculator interface with mouse and keyboard events, intuitive interaction |
| Database | Not required | - | Local history storage | In-memory storage with optional file persistence for session history |
| Cache | In-memory cache | Native | Runtime performance optimization | Fast access, minimal overhead, simple implementation |
| File Storage | Local filesystem | Native | Configuration and history files | Cross-platform, no external dependencies |
| Authentication | Not required | - | Terminal application security | Single-user local application, no external access |
| Frontend Testing | Go testing | Native | Terminal UI testing | Integrated with Go toolchain, mock-friendly |
| Backend Testing | Go testing | Native | Calculation engine testing | Built-in testing framework, benchmarking support |
| E2E Testing | CLI automation | v1.0 | Full application testing | Real terminal interaction testing, cross-platform |
| Build Tool | Go build | 1.21+ | Application compilation | Simple, fast, cross-platform compilation |
| Bundler | Go modules | 1.21+ | Dependency management | Official dependency management, version pinning |
| IaC Tool | Not required | - | Infrastructure management | Local deployment, no cloud infrastructure |
| CI/CD | GitHub Actions | Latest | Automated testing and building | Free for public repos, cross-platform testing |
| Monitoring | Structured logging | Native | Application monitoring | Built-in Go logging, configurable output levels |
| Logging | Go log package | Native | Application logging | Standard library integration, flexible output |
| CSS Framework | ANSI escape codes | v1.0 | Terminal styling | Cross-platform terminal formatting, minimal overhead |
