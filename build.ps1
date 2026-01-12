#!/usr/bin/env pwsh
# Enhanced build script with code quality checks

param(
    [switch]$Quick,      # Skip quality checks
    [switch]$NoTest,     # Skip tests
    [switch]$Verbose     # Verbose output
)

$ErrorActionPreference = "Stop"

function Write-Step {
    param([string]$Message)
    Write-Host "`n$Message" -ForegroundColor Cyan
}

function Write-Success {
    param([string]$Message)
    Write-Host "  ✓ $Message" -ForegroundColor Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "  ⚠ $Message" -ForegroundColor Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "  ✗ $Message" -ForegroundColor Red
}

Write-Host "`n╔═══════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║   Veracode MCP Server Build Script    ║" -ForegroundColor Cyan
Write-Host "╚═══════════════════════════════════════╝" -ForegroundColor Cyan

# 1. Format code
if (-not $Quick) {
    Write-Step "1. Formatting code..."
    go fmt ./...
    Write-Success "Code formatted"
}

# 2. Go vet
if (-not $Quick) {
    Write-Step "2. Running go vet..."
    # Exclude generated code from vet checks
    $vetOutput = go vet $(go list ./... | Where-Object { $_ -notmatch '/api/generated/' }) 2>&1
    if ($LASTEXITCODE -eq 0) {
        Write-Success "No issues found"
    } else {
        Write-Error "Go vet found issues:"
        Write-Host $vetOutput -ForegroundColor Red
        exit 1
    }
}

# 3. golangci-lint (if available)
if (-not $Quick) {
    Write-Step "3. Running golangci-lint..."
    if (Get-Command golangci-lint -ErrorAction SilentlyContinue) {
        golangci-lint run
        if ($LASTEXITCODE -eq 0) {
            Write-Success "All linters passed"
        } else {
            Write-Error "Linter checks failed"
            # exit 1
        }
    } else {
        Write-Warning "golangci-lint not installed (skipping)"
        Write-Host "    Install: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" -ForegroundColor Gray
    }
}

# 4. Run tests
if (-not $NoTest -and -not $Quick) {
    Write-Step "4. Running tests..."
    if ($Verbose) {
        go test ./... -v
    } else {
        go test ./...
    }
    
    if ($LASTEXITCODE -eq 0) {
        Write-Success "All tests passed"
    } else {
        Write-Error "Tests failed"
        exit 1
    }
}

# 5. Build
Write-Step "5. Building binary..."

# Create dist directory if it doesn't exist
if (-not (Test-Path "dist")) {
    New-Item -ItemType Directory -Path "dist" | Out-Null
}

go build -o dist/veracode-mcp.exe .
if ($LASTEXITCODE -eq 0) {
    $fileInfo = Get-Item dist/veracode-mcp.exe
    $sizeKB = [math]::Round($fileInfo.Length / 1KB, 2)
    Write-Success "Build successful (${sizeKB} KB)"
    Write-Host "    Output: dist/veracode-mcp.exe" -ForegroundColor Gray
} else {
    Write-Error "Build failed"
    exit 1
}

Write-Host "`n╔════════════════════════════════════════╗" -ForegroundColor Green
Write-Host "║          Build Successful! ✓           ║" -ForegroundColor Green
Write-Host "╚════════════════════════════════════════╝" -ForegroundColor Green

if ($Quick) {
    Write-Warning "Quick build - skipped quality checks"
}

Write-Host "`nRun with:" -ForegroundColor Cyan
Write-Host "  .\dist\veracode-mcp.exe -mode stdio" -ForegroundColor White
Write-Host "`nBuild options:" -ForegroundColor Gray
Write-Host "  .\build.ps1          # Full build with all checks" -ForegroundColor Gray
Write-Host "  .\build.ps1 -Quick   # Fast build, skip checks" -ForegroundColor Gray
Write-Host "  .\build.ps1 -NoTest  # Build without running tests" -ForegroundColor Gray
Write-Host "  .\build.ps1 -Verbose # Show detailed test output" -ForegroundColor Gray
