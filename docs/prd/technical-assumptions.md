# Technical Assumptions

## Repository Structure: Monorepo
Using a monorepo structure to contain both the calculator application and BMAD framework documentation, allowing for unified versioning and easier maintenance of the complete project ecosystem.

## Service Architecture
Single compiled executable with modular internal structure:
- Core calculation engine
- Command-line interface handler
- Input/output formatting module
- History management system
- Error handling subsystem

## Testing Requirements: Full Testing Pyramid
Comprehensive testing approach including:
- Unit tests for individual calculation functions
- Integration tests for CLI interface and user flows
- End-to-end tests for complete user scenarios
- Manual testing for edge cases and usability

## Additional Technical Assumptions and Requests
- **Language**: Go (Golang) 1.21+ for compiled executables, excellent performance, and cross-platform support
- **Dependencies**: Minimal external dependencies, leveraging Go's robust standard library
- **Package Management**: Go modules for dependency management
- **Build System**: Simple `go build` for creating platform-specific executables
- **Documentation**: Integrated with BMAD framework documentation standards
- **Version Control**: Git with conventional commit messages
- **CI/CD**: GitHub Actions for cross-platform compilation and testing
- **Deployment**: Compiled executables for Linux, macOS, and Windows
- **Logging**: Structured logging using Go's standard log package
- **Configuration**: Environment variables and optional config file support
- **Distribution**: GitHub releases with pre-compiled binaries for major platforms
