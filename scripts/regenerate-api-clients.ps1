#!/usr/bin/env pwsh
# Regenerate all API clients from OpenAPI/Swagger specs

param(
    [switch]$SkipClean,  # Don't clean existing generated code
    [switch]$Verbose     # Show verbose output
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

Write-Host "`n╔════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║   Regenerate Veracode API Clients      ║" -ForegroundColor Cyan
Write-Host "╚════════════════════════════════════════╝" -ForegroundColor Cyan

# Check for OpenAPI Generator
Write-Step "Checking for OpenAPI Generator..."
$generator = $null

# Try openapi-generator-cli (npm package)
if (Get-Command openapi-generator-cli -ErrorAction SilentlyContinue) {
    $generator = "openapi-generator-cli"
    Write-Success "Found openapi-generator-cli"
}
# Try openapi-generator (brew/direct install)
elseif (Get-Command openapi-generator -ErrorAction SilentlyContinue) {
    $generator = "openapi-generator"
    Write-Success "Found openapi-generator"
}
# Try swagger-codegen as fallback
elseif (Get-Command swagger-codegen -ErrorAction SilentlyContinue) {
    $generator = "swagger-codegen"
    Write-Warning "Using swagger-codegen (OpenAPI Generator recommended)"
}
else {
    Write-Error "No code generator found!"
    Write-Host "`nInstall OpenAPI Generator:" -ForegroundColor Yellow
    Write-Host "  npm install -g @openapitools/openapi-generator-cli" -ForegroundColor White
    Write-Host "  # or" -ForegroundColor Gray
    Write-Host "  brew install openapi-generator" -ForegroundColor White
    Write-Host "`nOr download from: https://github.com/OpenAPITools/openapi-generator/releases" -ForegroundColor Gray
    exit 1
}

# Define API specifications
$apis = @(
    @{
        Name = "healthcheck"
        Spec = "specs/veracode-healthcheck.json"
        Package = "healthcheck"
    },
    @{
        Name = "findings"
        Spec = "specs/veracode-findings.json"
        Package = "findings"
    },
    @{
        Name = "dynamic_flaw"
        Spec = "specs/veracode-dynamic-flaw.json"
        Package = "dynamic_flaw"
    },
    @{
        Name = "static_finding_data_path"
        Spec = "specs/veracode-static-finding-data-path.json"
        Package = "static_finding_data_path"
    }
    ,
    @{
        Name = "applications"
        Spec = "specs/veracode-applications.json"
        Package = "applications"
    }
)

# Clean existing generated code
if (-not $SkipClean) {
    Write-Step "Cleaning existing generated code..."
    foreach ($api in $apis) {
        $outputDir = "api/generated/$($api.Name)"
        if (Test-Path $outputDir) {
            # Keep README.md if it exists
            if (Test-Path "$outputDir/README.md") {
                $readme = Get-Content "$outputDir/README.md" -Raw
            }
            
            Remove-Item -Recurse -Force $outputDir
            New-Item -ItemType Directory -Path $outputDir | Out-Null
            
            # Restore README.md
            if ($readme) {
                Set-Content -Path "$outputDir/README.md" -Value $readme
            }
            
            Write-Success "Cleaned $($api.Name)"
        }
    }
}

# Generate API clients
Write-Step "Generating API clients..."
$generated = 0

foreach ($api in $apis) {
    if (-not (Test-Path $api.Spec)) {
        Write-Warning "Spec file not found: $($api.Spec) - Skipping $($api.Name)"
        continue
    }
    
    Write-Host "`n  Generating $($api.Name)..." -ForegroundColor Yellow
    
    $outputDir = "api/generated/$($api.Name)"
    
    # Generate based on which tool is available
    if ($generator -eq "openapi-generator-cli" -or $generator -eq "openapi-generator") {
        $cmd = "$generator generate -i $($api.Spec) -g go -o $outputDir --additional-properties=packageName=$($api.Package)"
    } else {
        # swagger-codegen
        $cmd = "$generator generate -i $($api.Spec) -l go -o $outputDir --additional-properties packageName=$($api.Package)"
    }
    
    if ($Verbose) {
        Write-Host "    Command: $cmd" -ForegroundColor Gray
        Invoke-Expression $cmd
    } else {
        Invoke-Expression "$cmd 2>&1" | Out-Null
    }
    
    if ($LASTEXITCODE -eq 0) {
        Write-Success "Generated $($api.Name)"
        $generated++
        
        # Clean up generated module files and test directories
        Remove-Item -Path "$outputDir/go.mod" -Force -ErrorAction SilentlyContinue
        Remove-Item -Path "$outputDir/go.sum" -Force -ErrorAction SilentlyContinue
        Remove-Item -Path "$outputDir/test" -Recurse -Force -ErrorAction SilentlyContinue
        Remove-Item -Path "$outputDir/.openapi-generator" -Recurse -Force -ErrorAction SilentlyContinue
    } else {
        Write-Error "Failed to generate $($api.Name)"
    }
}

# Update package names if using swagger-codegen (it defaults to 'swagger')
if ($generator -eq "swagger-codegen") {
    Write-Step "Updating package names..."
    foreach ($api in $apis) {
        $outputDir = "api/generated/$($api.Name)"
        if (Test-Path $outputDir) {
            Get-ChildItem -Path $outputDir -Recurse -Include "*.go" | ForEach-Object {
                (Get-Content $_.FullName) -replace 'package swagger', "package $($api.Package)" | Set-Content $_.FullName
            }
            Write-Success "Updated package name for $($api.Name)"
        }
    }
}

# Run go mod tidy
Write-Step "Updating Go modules..."
go mod tidy
if ($LASTEXITCODE -eq 0) {
    Write-Success "Dependencies updated"
} else {
    Write-Warning "go mod tidy reported warnings (this is usually okay)"
}

# Summary
Write-Host "`n╔════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║          Generation Complete!          ║" -ForegroundColor Cyan
Write-Host "╚════════════════════════════════════════╝" -ForegroundColor Cyan
Write-Host "`n  Generated $generated API client(s)" -ForegroundColor Green

if ($generated -lt $apis.Count) {
    Write-Host "`n  Note: $($apis.Count - $generated) API(s) skipped (spec file not found)" -ForegroundColor Yellow
    Write-Host "  Place spec files in specs/ directory to generate all clients" -ForegroundColor Gray
}

Write-Host "`nNext steps:" -ForegroundColor Cyan
Write-Host "  1. Review generated code in api/generated/" -ForegroundColor White
Write-Host "  2. Run: .\build.ps1 -Quick" -ForegroundColor White
Write-Host "  3. Update api/client.go if new clients were added" -ForegroundColor White
Write-Host ""
