#!/usr/bin/env pwsh
# Download Veracode API specifications from SwaggerHub

param(
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

function Write-Error {
    param([string]$Message)
    Write-Host "  ✗ $Message" -ForegroundColor Red
}

Write-Host "`n╔════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║   Download Veracode API Specs          ║" -ForegroundColor Cyan
Write-Host "╚════════════════════════════════════════╝" -ForegroundColor Cyan

# Define API specifications with SwaggerHub URLs
$specs = @(
    @{
        Name = "Findings API"
        File = "specs/veracode-findings.json"
        Url = "https://api.swaggerhub.com/apis/Veracode/veracode-findings_api_specification/2.1"
    },
    @{
        Name = "Healthcheck API"
        File = "specs/veracode-healthcheck.json"
        Url = "https://api.swaggerhub.com/apis/Veracode/veracode-healthcheck_api_specification/1.0"
    },
    @{
        Name = "Dynamic Flaw API"
        File = "specs/veracode-dynamic-flaw.json"
        Url = "https://api.swaggerhub.com/apis/Veracode/veracode-dynamic_flaw_api/v2"
    },
    @{
        Name = "Static Finding Data Path API"
        File = "specs/veracode-static-finding-data-path.json"
        Url = "https://api.swaggerhub.com/apis/Veracode/veracode-static_finding_data_path_api/v2"
    },
    @{
        Name = "Applications API"
        File = "specs/veracode-applications.json"
        Url = "https://api.swaggerhub.com/apis/Veracode/veracode-applications_api_specification/1.0"
    }
)

# Ensure specs directory exists
if (-not (Test-Path "specs")) {
    New-Item -ItemType Directory -Path "specs" | Out-Null
}

# Download each spec
Write-Step "Downloading API specifications..."
$downloaded = 0
$failed = 0

foreach ($spec in $specs) {
    Write-Host "`n  Downloading $($spec.Name)..." -ForegroundColor Yellow
    
    try {
        if ($Verbose) {
            Write-Host "    URL: $($spec.Url)" -ForegroundColor Gray
            Write-Host "    File: $($spec.File)" -ForegroundColor Gray
        }
        
        # Download using Invoke-WebRequest
        $response = Invoke-WebRequest -Uri $spec.Url -UseBasicParsing
        
        # Parse and pretty-print if it's JSON
        $content = $response.Content
        if ($spec.Url -match "\.json" -or $content.TrimStart() -match "^[\{\[]") {
            try {
                # Parse JSON and convert back with indentation
                $jsonObject = $content | ConvertFrom-Json
                $content = $jsonObject | ConvertTo-Json -Depth 100
                if ($Verbose) {
                    Write-Host "    Pretty-printed JSON" -ForegroundColor Gray
                }
            }
            catch {
                # If JSON parsing fails, save as-is
                if ($Verbose) {
                    Write-Host "    Saved as plain text (not valid JSON)" -ForegroundColor Gray
                }
            }
        }
        
        # Save to file
        $content | Set-Content -Path $spec.File -Encoding UTF8
        
        # Apply spec-specific fixes
        if ($spec.File -eq "specs/veracode-findings.json") {
            if ($Verbose) {
                Write-Host "    Applying fixes to Findings API spec..." -ForegroundColor Gray
            }
            
            # Fix StaticFinding and DynamicFinding schemas to match actual API behavior
            # The spec is in JSON format, so we'll load it, fix it as text, and save
            $specContent = Get-Content -Path $spec.File -Raw
            
            # Fix severity: change "type": "string" to "type": "integer" for severity property
            $specContent = $specContent -replace '("severity":\s*\{[^}]*)"type":\s*"string"', '$1"type": "integer"'
            
            # Remove $ref for severity and replace with type: integer
            $specContent = $specContent -replace '("severity":\s*\{[^}]*)\"\$ref\":\s*"#/components/schemas/Severity"', '$1"type": "integer"'
            
            # Fix file_line_number: change type from string to integer
            $specContent = $specContent -replace '("file_line_number":\s*\{[^}]*)"type":\s*"string"', '$1"type": "integer"'
            
            # Fix relative_location: change type from string to integer  
            $specContent = $specContent -replace '("relative_location":\s*\{[^}]*)"type":\s*"string"', '$1"type": "integer"'
            
            # Fix exploitability: change type from number to integer
            $specContent = $specContent -replace '("exploitability":\s*\{[^}]*)"type":\s*"number"', '$1"type": "integer"'
            
            # Fix cwe: replace $ref with object type
            $cweObject = '"type": "object", "properties": { "id": { "type": "integer", "format": "int32" }, "name": { "type": "string" }, "href": { "type": "string" } }'
            $specContent = $specContent -replace '("cwe":\s*\{[^}]*)\"\$ref\":\s*"#/components/schemas/Cwe"', ('$1' + $cweObject)
            
            # Fix finding_category: change from string to object
            $categoryObject = '"type": "object", "properties": { "id": { "type": "integer", "format": "int32" }, "name": { "type": "string" }, "href": { "type": "string" } }'
            $specContent = $specContent -replace '("finding_category":\s*\{[^}]*)"type":\s*"string"', ('$1' + $categoryObject)
            
            $specContent | Set-Content -Path $spec.File -Encoding UTF8 -NoNewline
            
            if ($Verbose) {
                Write-Host "    Applied spec fixes to StaticFinding and DynamicFinding schemas" -ForegroundColor Gray
            }
        }
        
        # Verify file was created
        if (Test-Path $spec.File) {
            $fileSize = (Get-Item $spec.File).Length
            Write-Success "$($spec.Name) downloaded ($fileSize bytes)"
            $downloaded++
        } else {
            Write-Error "Failed to save $($spec.Name)"
            $failed++
        }
    }
    catch {
        Write-Error "Failed to download $($spec.Name): $($_.Exception.Message)"
        $failed++
    }
}

# Summary
Write-Host "`n╔════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║          Download Complete!            ║" -ForegroundColor Cyan
Write-Host "╚════════════════════════════════════════╝" -ForegroundColor Cyan
Write-Host "`n  Downloaded: $downloaded spec file(s)" -ForegroundColor Green

if ($failed -gt 0) {
    Write-Host "  Failed: $failed spec file(s)" -ForegroundColor Red
    Write-Host "`n  Check your internet connection or SwaggerHub URLs" -ForegroundColor Yellow
}

if ($downloaded -gt 0) {
    Write-Host "`nNext steps:" -ForegroundColor Cyan
    Write-Host "  1. Review downloaded specs in specs/" -ForegroundColor White
    Write-Host "  2. Run: .\scripts\regenerate-api-clients.ps1" -ForegroundColor White
    Write-Host "  3. Commit spec files to git" -ForegroundColor White
}

Write-Host ""

# Exit with error code if any downloads failed
if ($failed -gt 0) {
    exit 1
}
