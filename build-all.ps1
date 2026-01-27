# Build script for Veracode MCP Server with UI

param(
    [switch]$SkipUI,
    [switch]$UIOnly
)

Write-Host "Building Veracode MCP Server..." -ForegroundColor Cyan

# Build UI if not skipped
if (-not $SkipUI) {
    # Build Pipeline Results UI
    Write-Host "`nBuilding Pipeline Results UI..." -ForegroundColor Yellow
    
    $pipelineUIPath = Join-Path $PSScriptRoot "ui\pipeline-results-app"
    
    if (-not (Test-Path $pipelineUIPath)) {
        Write-Host "Error: Pipeline UI directory not found at $pipelineUIPath" -ForegroundColor Red
        exit 1
    }
    
    Push-Location $pipelineUIPath
    
    try {
        # Install dependencies if node_modules doesn't exist
        if (-not (Test-Path "node_modules")) {
            Write-Host "Installing UI dependencies..." -ForegroundColor Yellow
            npm install
            if ($LASTEXITCODE -ne 0) {
                throw "npm install failed"
            }
        }
        
        # Build UI
        Write-Host "Building Pipeline Results UI..." -ForegroundColor Yellow
        npm run build
        if ($LASTEXITCODE -ne 0) {
            throw "Pipeline UI build failed"
        }
        
        Write-Host "Pipeline Results UI built successfully!" -ForegroundColor Green
    }
    catch {
        Write-Host "Error building Pipeline UI: $_" -ForegroundColor Red
        Pop-Location
        exit 1
    }
    
    Pop-Location
    
    # Build Static Findings UI
    Write-Host "`nBuilding Static Findings UI..." -ForegroundColor Yellow
    
    $staticUIPath = Join-Path $PSScriptRoot "ui\static-findings-app"
    
    if (-not (Test-Path $staticUIPath)) {
        Write-Host "Error: Static Findings UI directory not found at $staticUIPath" -ForegroundColor Red
        exit 1
    }
    
    Push-Location $staticUIPath
    
    try {
        # Install dependencies if node_modules doesn't exist
        if (-not (Test-Path "node_modules")) {
            Write-Host "Installing UI dependencies..." -ForegroundColor Yellow
            npm install
            if ($LASTEXITCODE -ne 0) {
                throw "npm install failed"
            }
        }
        
        # Build UI
        Write-Host "Building Static Findings UI..." -ForegroundColor Yellow
        npm run build
        if ($LASTEXITCODE -ne 0) {
            throw "Static Findings UI build failed"
        }
        
        Write-Host "Static Findings UI built successfully!" -ForegroundColor Green
    }
    catch {
        Write-Host "Error building Static Findings UI: $_" -ForegroundColor Red
        Pop-Location
        exit 1
    }
    
    Pop-Location
    
    # Build Dynamic Findings UI
    Write-Host "`nBuilding Dynamic Findings UI..." -ForegroundColor Yellow
    
    $dynamicUIPath = Join-Path $PSScriptRoot "ui\dynamic-findings-app"
    
    if (-not (Test-Path $dynamicUIPath)) {
        Write-Host "Error: Dynamic Findings UI directory not found at $dynamicUIPath" -ForegroundColor Red
        exit 1
    }
    
    Push-Location $dynamicUIPath
    
    try {
        # Install dependencies if node_modules doesn't exist
        if (-not (Test-Path "node_modules")) {
            Write-Host "Installing UI dependencies..." -ForegroundColor Yellow
            npm install
            if ($LASTEXITCODE -ne 0) {
                throw "npm install failed"
            }
        }
        
        # Build UI
        Write-Host "Building Dynamic Findings UI..." -ForegroundColor Yellow
        npm run build
        if ($LASTEXITCODE -ne 0) {
            throw "Dynamic Findings UI build failed"
        }
        
        Write-Host "Dynamic Findings UI built successfully!" -ForegroundColor Green
    }
    catch {
        Write-Host "Error building Dynamic Findings UI: $_" -ForegroundColor Red
        Pop-Location
        exit 1
    }
    
    Pop-Location
}

# Exit if UI-only build
if ($UIOnly) {
    Write-Host "`nBuild complete!" -ForegroundColor Green
    exit 0
}

# Build Go server
Write-Host "`nBuilding Go server..." -ForegroundColor Yellow

# Create dist directory if it doesn't exist
if (-not (Test-Path "dist")) {
    New-Item -ItemType Directory -Path "dist" | Out-Null
}

try {
    go build -o dist/veracode-mcp.exe
    if ($LASTEXITCODE -ne 0) {
        throw "Go build failed"
    }
    
    $fileInfo = Get-Item dist/veracode-mcp.exe
    $sizeMB = [math]::Round($fileInfo.Length / 1MB, 2)
    Write-Host "Go server built successfully! (${sizeMB} MB)" -ForegroundColor Green
}
catch {
    Write-Host "Error building Go server: $_" -ForegroundColor Red
    exit 1
}

Write-Host "`nBuild complete! Run .\dist\veracode-mcp.exe -mode stdio to start the server." -ForegroundColor Green
