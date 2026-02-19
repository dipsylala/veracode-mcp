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
    },
    @{
        Name = "Policy API"
        File = "specs/veracode-policy.json"
        Url = "https://api.swaggerhub.com/apis/Veracode/veracode-policy_api_specification/1.0"
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
        
        # Remove _links properties from all specs (not needed for client generation)
        if ($Verbose) {
            Write-Host "    Removing _links properties..." -ForegroundColor Gray
        }
        
        $specContent = Get-Content -Path $spec.File -Raw
        
        # Remove "_links" property and its value, handling various patterns:
        # 1. "_links": { ... } with nested objects
        # 2. "_links": [ ... ] with arrays
        # Need to handle nested braces/brackets carefully
        
        # Match "_links": followed by either array or object, with proper nesting
        # This regex handles nested structures up to 3 levels deep
        $pattern = '"_links"\s*:\s*(?:\{(?:[^{}]|\{(?:[^{}]|\{[^{}]*\})*\})*\}|\[(?:[^\[\]]|\[(?:[^\[\]]|\[[^\[\]]*\])*\])*\])\s*,?'
        $specContent = $specContent -replace $pattern, ''
        
        # Clean up trailing commas before closing braces
        $specContent = $specContent -replace ',(\s*[}\]])', '$1'
        
        $specContent | Set-Content -Path $spec.File -Encoding UTF8 -NoNewline
        
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
            
            # Add Page and Size parameters to findings endpoint (missing from SwaggerHub spec)
            $specContent = $specContent -replace '(\{\s*"\$ref":\s*"#/components/parameters/PolicyViolation"\s*\}\s*)\](\s*,\s*"responses")', ('$1,{' + "`n" + '            "$ref": "#/components/parameters/Page"' + "`n" + '          },{' + "`n" + '            "$ref": "#/components/parameters/Size"' + "`n" + '          }]$2')
            
            # Fix discovered_by_vsa: API returns number (0/1), not string or boolean
            $specContent = $specContent -replace '("discovered_by_vsa":\s*\{[^}]*)"type":\s*"string"', '$1"type": "integer"'
            
            # Fix metadata: API returns object, not string
            $specContent = $specContent -replace '("metadata":\s*\{[^}]*)"type":\s*"string"([^}]*"description":[^}]*\})', '$1"type": "object"$2'
            
            # Fix CVE severity: API returns string, not integer
            $specContent = $specContent -replace '("ScaFindingCve"[\s\S]*?"severity":\s*\{[^}]*)"type":\s*"integer"', '$1"type": "string"'
            
            # Fix CVSS3 severity: API also returns string like "Very High" for CVSS3 severity
            $specContent = $specContent -replace '("cvss3"[\s\S]{0,500}?"severity":\s*\{[^}]*)"type":\s*"integer"', '$1"type": "string"'
            
            $specContent | Set-Content -Path $spec.File -Encoding UTF8 -NoNewline
            
            if ($Verbose) {
                Write-Host "    Applied spec fixes to StaticFinding and DynamicFinding schemas" -ForegroundColor Gray
            }
        }
        
        # Apply fixes to Static Finding Data Path API spec
        if ($spec.File -eq "specs/veracode-static-finding-data-path.json") {
            if ($Verbose) {
                Write-Host "    Applying fixes to Static Finding Data Path API spec..." -ForegroundColor Gray
            }
            
            $specContent = Get-Content -Path $spec.File -Raw
            
            # Fix IssueSummary.name: API returns string (application name), not int32
            $specContent = $specContent -replace '("IssueSummary"[\s\S]{0,500}?"name":\s*\{[^}]*)"type":\s*"integer"[^}]*"format":\s*"int32"', '$1"type": "string"'
            
            $specContent | Set-Content -Path $spec.File -Encoding UTF8 -NoNewline
            
            if ($Verbose) {
                Write-Host "    Applied spec fixes to Static Finding Data Path API" -ForegroundColor Gray
            }
        }
        
        # Apply fixes to Dynamic Flaw API spec
        if ($spec.File -eq "specs/veracode-dynamic-flaw.json") {
            if ($Verbose) {
                Write-Host "    Applying fixes to Dynamic Flaw API spec..." -ForegroundColor Gray
            }
            
            $specContent = Get-Content -Path $spec.File -Raw
            
            # Fix Request.port: API returns string (like "443"), not integer
            $specContent = $specContent -replace '("port":\s*\{[^}]*)"type":\s*"integer"', '$1"type": "string"'
            
            $specContent | Set-Content -Path $spec.File -Encoding UTF8 -NoNewline
            
            if ($Verbose) {
                Write-Host "    Applied spec fixes to Dynamic Flaw API" -ForegroundColor Gray
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
