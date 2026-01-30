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

**⚠️ IMPORTANT: Build UI First**

This project embeds UI HTML files in the binary using `go:embed`. Before running golangci-lint or building, you **must** build the UI apps first:

```bash
# Build UI apps only
.\build.ps1 -UIOnly

# Or run full build (recommended)
.\build.ps1
```

If you run `golangci-lint` directly without building the UI first, you'll get typecheck errors like:
```
Error: pattern ui/dynamic-findings-app/dist/mcp-app.html: no matching files found (typecheck)
```

**Proper workflow:**

```bash
# 1. Build UI apps (creates embedded HTML files)
.\build.ps1 -UIOnly

# 2. Now you can run linting
golangci-lint run

# 3. Or just use the full build script
.\build.ps1   # Builds UI, runs quality checks, tests, and compiles binary
```

---

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

For GitHub Actions, ensure UI is built before linting:

```yaml
- name: Set up Node.js
  uses: actions/setup-node@v4
  with:
    node-version: '20'

- name: Build UI Apps
  run: |
    # Build all UI apps (required for go:embed to work)
    cd ui/pipeline-results-app && npm install && npm run build && cd ../..
    cd ui/static-findings-app && npm install && npm run build && cd ../..
    cd ui/dynamic-findings-app && npm install && npm run build && cd ../..

- name: Run golangci-lint
  uses: golangci/golangci-lint-action@v4
  with:
    version: latest
```

**Why this is needed:** The `main.go` file uses `//go:embed` directives to bundle UI HTML files into the binary. These files must exist when Go's type checker runs, or you'll get errors like `pattern ui/.../dist/mcp-app.html: no matching files found`.

## Benefits

- **Catch bugs early** - Before they reach production  
- **Consistent style** - Entire codebase looks uniform  
- **Better performance** - Identifies inefficiencies  
- **Security** - Detects common vulnerabilities  
- **Maintainability** - Cleaner, more readable code  
