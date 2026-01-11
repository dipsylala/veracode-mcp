# API Specifications

This directory contains OpenAPI/Swagger specifications for Veracode APIs.

## Current APIs

### Findings API (v2.1)
- **SwaggerHub**: https://app.swaggerhub.com/apis/Veracode/veracode-findings_api_specification/2.1
- **File**: `veracode-findings.yaml`
- **Description**: SAST/DAST/SCA findings retrieval

### Healthcheck API (v1.0)
- **SwaggerHub**: https://app.swaggerhub.com/apis/Veracode/veracode-healthcheck_api_specification/1.0
- **File**: `veracode-healthcheck.yaml`
- **Description**: API health status checks

### Dynamic Flaw API (v2)
- **SwaggerHub**: https://app.swaggerhub.com/apis/Veracode/veracode-dynamic_flaw_api/v2
- **File**: `veracode-dynamic-flaw.yaml`
- **Description**: Dynamic flaw details and data paths

### Static Finding Data Path API (v2)
- **SwaggerHub**: https://app.swaggerhub.com/apis/Veracode/veracode-static_finding_data_path_api/v2
- **File**: `veracode-static-finding-data-path.yaml`
- **Description**: Static finding data paths and call stacks

## Downloading Spec Files

### Option 1: Manual Download from SwaggerHub

1. Visit the SwaggerHub URL for each API
2. Click the download icon or "Export" → "Download API" → "YAML Resolved"
3. Save to this directory with the corresponding filename

### Option 2: Using the Download Script

```powershell
.\scripts\download-api-specs.ps1
```

This will download all spec files directly from SwaggerHub.

## Regenerating API Clients

After downloading/updating spec files, run:

```powershell
.\scripts\regenerate-api-clients.ps1
```

This will:
1. Clean existing generated code (except READMEs)
2. Generate new clients using OpenAPI Generator
3. Update package names
4. Run `go mod tidy`

## Version Management

When Veracode releases new API versions:
1. Update the SwaggerHub URLs above with new version numbers
2. Download updated spec files
3. Regenerate clients
4. Test compatibility with existing code
5. Commit both spec files and any necessary code changes
