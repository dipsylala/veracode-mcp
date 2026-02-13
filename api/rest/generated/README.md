# Generated API Clients

This directory contains **auto-generated** API client code from OpenAPI/Swagger specifications.

## ⚠️ DO NOT EDIT

All code in this directory is automatically generated. Any manual changes will be lost when the clients are regenerated.

## Structure

Each subdirectory contains a separate API client:

- **`healthcheck/`** - Veracode Healthcheck API
  - Generated from: specs/veracode-healthcheck.yaml
  - Package: `healthcheck`
  - Endpoints: Health status checks

- **`findings/`** - Veracode Findings API
  - Generated from: specs/veracode-findings.yaml
  - Package: `findings`
  - Endpoints: SAST/DAST/SCA findings retrieval

- **`dynamic_flaw/`** - Veracode Dynamic Flaw API
  - Generated from: specs/veracode-dynamic-flaw.yaml
  - Package: `dynamic_flaw`
  - Endpoints: Dynamic flaw details

- **`static_finding_data_path/`** - Veracode Static Finding Data Path API
  - Generated from: specs/veracode-static-finding-data-path.yaml
  - Package: `static_finding_data_path`
  - Endpoints: Static finding data paths

## Regenerating Clients

When API specs are updated, regenerate all clients:

```bash
.\scripts\regenerate-api-clients.ps1
```

Or regenerate manually using OpenAPI Generator:

```bash
# Install OpenAPI Generator
npm install -g @openapitools/openapi-generator-cli
# or
brew install openapi-generator

# Generate a client
openapi-generator generate \
  -i specs/veracode-findings.yaml \
  -g go \
  -o api/generated/findings \
  --additional-properties=packageName=findings,enumClassPrefix=true
```

## Usage

**Don't import these directly in tools!**

Instead, use the `api/` package which wraps these clients:

```go
// ❌ BAD - Don't do this in tools
import healthcheck "github.com/dipsylala/veracodemcp-go/api/generated/healthcheck"

// ✅ GOOD - Use the api wrapper
import "github.com/dipsylala/veracodemcp-go/api"

client, err := api.NewClient()
status, err := client.CheckHealth(ctx)
```

## Package Names

Each generated client uses its own package name to avoid conflicts:
- `healthcheck` - Healthcheck API
- `findings` - Findings API
- `upload` - Upload API

This allows the `api/` package to import multiple clients simultaneously.
