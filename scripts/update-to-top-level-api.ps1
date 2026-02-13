#!/usr/bin/env pwsh
# Update tools to use top-level api package instead of api/rest

$ErrorActionPreference = "Stop"

Write-Host "Updating tool imports to use api package (not api/rest)..."

$toolFiles = Get-ChildItem -Path "internal/mcp_tools" -Filter "*.go" -File

foreach ($file in $toolFiles) {
    $content = Get-Content -Path $file.FullName -Raw
    
    # Change import from api/rest to api
    $newContent = $content -replace '"github\.com/dipsylala/veracodemcp-go/api/rest"', '"github.com/dipsylala/veracodemcp-go/api"'
    
    # Change generated package imports back to not include rest/
    # (they're re-exported through the api package)
    $newContent = $newContent -replace 'github\.com/dipsylala/veracodemcp-go/api/rest/generated/', 'github.com/dipsylala/veracodemcp-go/api/rest/generated/'
    
    # Change rest.TypeName back to api.TypeName
    $newContent = $newContent -replace '\*rest\.([A-Z][a-zA-Z0-9]*)', '*api.$1'
    $newContent = $newContent -replace '(\s)rest\.([A-Z][a-zA-Z0-9]*)', '$1api.$2'
    $newContent = $newContent -replace '([^"\/])rest\.([A-Z][a-zA-Z0-9]*)', '$1api.$2'
    
    if ($content -ne $newContent) {
        Set-Content -Path $file.FullName -Value $newContent -NoNewline
        Write-Host "  Updated $($file.Name)"
    }
}

Write-Host "`nComplete!"
