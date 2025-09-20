# Deployment Architecture

## Deployment Strategy

**Frontend Deployment:**
- **Platform:** Native terminal application
- **Build Command:** `go build -o calculator cmd/calculator/main.go`
- **Output Directory:** ./calculator (executable)
- **CDN/Edge:** Not applicable for terminal application

**Backend Deployment:**
- **Platform:** Cross-platform compiled executable
- **Build Command:** `go build -o calculator cmd/calculator/main.go`
- **Deployment Method:** GitHub releases with pre-compiled binaries

## CI/CD Pipeline

```yaml
name: CI/CD Pipeline

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    - uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Run tests
      run: go test -v ./...
    
    - name: Run integration tests
      run: go test -tags=integration ./...
    
    - name: Build application
      run: go build -o calculator cmd/calculator/main.go

  build:
    needs: test
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest, windows-latest, macos-latest]
        arch: [amd64, arm64]
    
    steps:
    - uses: actions/checkout@v3
    - uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Build for ${{ matrix.os }}-${{ matrix.arch }}
      run: |
        GOOS=${{ matrix.os }} GOARCH=${{ matrix.arch }} go build -o calculator-${{ matrix.os }}-${{ matrix.arch }} cmd/calculator/main.go
    
    - name: Upload artifacts
      uses: actions/upload-artifact@v3
      with:
        name: calculator-${{ matrix.os }}-${{ matrix.arch }}
        path: calculator-${{ matrix.os }}-${{ matrix.arch }}
```

## Environments

| Environment | Frontend URL | Backend URL | Purpose |
|-------------|--------------|-------------|---------|
| Development | Local terminal | Local executable | Local development |
| Testing | CI/CD pipeline | Test runners | Automated testing |
| Production | User's terminal | Released binaries | End-user deployment |
