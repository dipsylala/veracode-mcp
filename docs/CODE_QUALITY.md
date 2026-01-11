# Code Quality & Build Tools

## Essential Go Tools

### 1. Built-in Tools (Already Available)

**`go fmt`** - Formats code to Go standards

```bash
go fmt ./...
```

**`go vet`** - Detects suspicious code

```bash
go vet ./...
```

**`goimports`** - Manages imports + formatting

```bash
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

### 2. golangci-lint (Recommended - All-in-One)

The industry standard meta-linter that runs 50+ linters.

**Install:**

```bash
# Windows (PowerShell)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Or download binary from: https://github.com/golangci/golangci-lint/releases
```

**Run:**

```bash
golangci-lint run
```

**What it checks:**

- Code style (gofmt, goimports)
- Bugs (govet, staticcheck, errcheck)
- Complexity (gocyclo, gocognit)
- Unused code (unused, deadcode)
- Security issues (gosec)
- Performance (prealloc)
- Best practices (ineffassign, misspell)

### 3. staticcheck (Advanced Analysis)

Deep static analysis for bugs and performance.

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...
```

## Recommended Setup

### Option 1: PowerShell Build Script

Create `build.ps1` with quality checks:

```powershell
#!/usr/bin/env pwsh
# Build script with quality checks

Write-Host "`n=== Code Quality Checks ===" -ForegroundColor Cyan

# Format code
Write-Host "`n1. Formatting code..." -ForegroundColor Yellow
go fmt ./...

# Check for common mistakes
Write-Host "`n2. Running go vet..." -ForegroundColor Yellow
go vet ./...
if ($LASTEXITCODE -ne 0) { exit 1 }

# Run linter (if installed)
if (Get-Command golangci-lint -ErrorAction SilentlyContinue) {
    Write-Host "`n3. Running golangci-lint..." -ForegroundColor Yellow
    golangci-lint run
    if ($LASTEXITCODE -ne 0) { exit 1 }
} else {
    Write-Host "`n3. golangci-lint not found (skipping)" -ForegroundColor Gray
}

# Run tests
Write-Host "`n4. Running tests..." -ForegroundColor Yellow
go test ./... -v
if ($LASTEXITCODE -ne 0) { exit 1 }

# Build
Write-Host "`n5. Building..." -ForegroundColor Yellow
go build -o mcp-server.exe .
if ($LASTEXITCODE -ne 0) { exit 1 }

Write-Host "`n=== Build Successful! ===" -ForegroundColor Green
```

### Option 2: Makefile (Cross-platform)

```makefile
.PHONY: all fmt vet lint test build clean

all: fmt vet lint test build

fmt:
  @echo "Formatting code..."
  @go fmt ./...

vet:
  @echo "Running go vet..."
  @go vet ./...

lint:
  @echo "Running golangci-lint..."
  @golangci-lint run || echo "golangci-lint not installed"

test:
  @echo "Running tests..."
  @go test ./... -v

build:
  @echo "Building..."
  @go build -o mcp-server.exe .

clean:
  @echo "Cleaning..."
  @rm -f mcp-server.exe
  @go clean

# Quick build without all checks
quick:
  @go build -o mcp-server.exe .
```

### Option 3: golangci-lint Configuration

Create `.golangci.yml` for customization:

```yaml
run:
  timeout: 5m
  tests: true

linters:
  enable:
    - errcheck      # Check for unchecked errors
    - gosimple      # Simplify code
    - govet         # Go vet
    - ineffassign   # Detect ineffectual assignments
    - staticcheck   # Advanced static analysis
    - unused        # Find unused code
    - gofmt         # Check formatting
    - goimports     # Check imports
    - misspell      # Check spelling
    - unconvert     # Remove unnecessary conversions
    - unparam       # Find unused function parameters
    - gosec         # Security issues
    
linters-settings:
  gofmt:
    simplify: true
  goimports:
    local-prefixes: github.com/dipsylala/veracodemcp-go
  govet:
    check-shadowing: true
  errcheck:
    check-blank: true

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0
```

## Quick Start

1. **Install golangci-lint:**

  ```bash
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

2. **Run quality checks:**

  ```bash
  # Format
  go fmt ./...
   
  # Check for issues
  go vet ./...

  # Run all linters
  golangci-lint run

  # Run tests
  go test ./...
  ```

3. **Or use the build script:**

  ```bash
  .\build.ps1
  ```

## CI/CD Integration

For GitHub Actions:

```yaml
- name: Run golangci-lint
  uses: golangci/golangci-lint-action@v3
  with:
    version: latest
```

## Benefits

✅ **Catch bugs early** - Before they reach production  
✅ **Consistent style** - Entire codebase looks uniform  
✅ **Better performance** - Identifies inefficiencies  
✅ **Security** - Detects common vulnerabilities  
✅ **Maintainability** - Cleaner, more readable code  
