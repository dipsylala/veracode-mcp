# API Package

This package provides a clean interface to Veracode APIs, organized in three layers:

```
api/
├── client.go           # Client orchestrator (manages all API clients)
├── helpers/            # Business logic wrappers
│   ├── findings.go     # SAST/DAST findings with filtering
│   └── health.go       # Health check wrapper
└── generated/          # Auto-generated Swagger clients
    ├── healthcheck/
    ├── findings/
    ├── dynamic_flaw/
    └── static_finding_data_path/
```

## Purpose

- **Clean abstraction** over generated Swagger code
- **Authentication handling** (VERACODE_API_ID, VERACODE_API_KEY)
- **Business logic helpers** with filtering, pagination, type conversion
- **Reusable API methods** shared across multiple tools
- **Type-safe interfaces** with simple Go structs

## Structure

### client.go (Orchestrator)

Manages all API clients and authentication:
- `NewVeracodeClient()` - Creates client with all 4 API clients initialized
- `GetAuthContext()` - Adds Veracode HMAC authentication to requests
- `IsConfigured()` - Checks if credentials are set
- Holds: healthcheckClient, findingsClient, dynamicFlawClient, staticFindingDataPathClient

### helpers/ (Business Logic)

#### helpers/health.go

Wraps healthcheck API:

- `CheckHealth()` - Returns health status struct
- `CheckHealthSimple()` - Returns simple boolean
- Used by `api-health` tool

#### helpers/findings.go

Wraps findings API with advanced features:

- `GetDynamicFindings()` - DAST results with client-side filtering
- `GetStaticFindings()` - SAST results with client-side filtering
- `convertFindings()` - Transforms complex API types to simple Finding struct
- `applyFilters()` - Client-side filtering for severity/status
- Handles severity mapping (0-5 integers → "High", "Medium", etc.)
- Used by `dynamic-findings` and `static-findings` tools

### generated/ (Auto-Generated)

Swagger-generated API clients - **DO NOT EDIT**:

- `healthcheck/` - Health check API client
- `findings/` - Findings API client (SAST/DAST/SCA)
- `dynamic_flaw/` - Dynamic flaw details API client
- `static_finding_data_path/` - Static finding data path API client

See [generated/README.md](generated/README.md) for regeneration instructions.

## Usage in Tools

Tools import and use the API package:

```go
package tools

import (
    "github.com/dipsylala/veracodemcp-go/api"
)

func (t *APIHealthTool) handleAPIHealth(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Create API client
    client, err := api.NewVeracodeClient()
    if err != nil {
        return errorResponse(err.Error()), nil
    }

    // Call API
    status, err := client.CheckHealth(ctx)
    if err != nil {
        return errorResponse(err.Error()), nil
    }

    // Return result
    return successResponse(status.Message), nil
}
```

## Environment Variables

Required for authentication:

- `VERACODE_API_ID` - Your Veracode API ID
- `VERACODE_API_KEY` - Your Veracode API Secret Key

## Many-to-Many Relationships

The API package enables clean many:many relationships:

**One tool → Multiple API calls:**

- `get-dynamic-findings` tool calls:
  - `GetDynamicFindings()` (findings.go)
  - `GetFindingByID()` (findings.go)

**One API method → Multiple tools:**

- `GetDynamicFindings()` method used by:
  - `get-dynamic-findings` tool
  - Future: `compare-scans` tool
  - Future: `generate-report` tool

## Extending

To add a new Swagger-generated API:

1. **Generate client**:

   ```bash
   swagger-codegen generate \
     -i specs/veracode-new-api.yaml \
     -l go \
     -o go-client-generated \
     --additional-properties packageName=new_api
   ```

2. **Move and rename package**:

   ```powershell
   Move-Item go-client-generated api/generated/new_api
   Get-ChildItem api/generated/new_api -Recurse -File | ForEach-Object {
       (Get-Content $_.FullName) -replace 'package swagger', 'package new_api' | Set-Content $_.FullName
   }
   ```

3. **Integrate into client.go**:

   ```go
   import new_api "github.com/dipsylala/veracodemcp-go/api/generated/new_api"
   
   type VeracodeClient struct {
       // ... existing clients ...
       newApiClient *new_api.APIClient
   }
   
   func NewVeracodeClient() (*VeracodeClient, error) {
       // ... existing clients ...
       newApiCfg := new_api.NewConfiguration()
       
       return &VeracodeClient{
           // ... existing clients ...
           newApiClient: new_api.NewAPIClient(newApiCfg),
       }
   }
   ```

4. **Create helper (if needed)**:

   ```go
   // api/helpers/new_feature.go
   package api
   
   func (c *VeracodeClient) DoSomething(ctx context.Context, params Request) (*Response, error) {
       authCtx := c.GetAuthContext(ctx)
       resp, _, err := c.newApiClient.SomeApi.SomeMethod(authCtx, ...)
       // Add filtering, type conversion, business logic
       return result, nil
   }
   ```

5. **Use in tools**:

   ```go
   import "github.com/dipsylala/veracodemcp-go/api"
   
   client, _ := api.NewVeracodeClient()
   result, err := client.DoSomething(ctx, params)
   ```

## When to Create Helpers

Create helpers in `helpers/` when you need:

- **Type conversion**: Complex optional.Interface → simple structs
- **Client-side filtering**: API doesn't support all filter combinations
- **Pagination logic**: Handling page/size parameters
- **Severity mapping**: Converting numeric codes to human-readable strings
- **Business logic**: Combining multiple API calls or adding domain logic

Skip helpers for simple APIs where direct client usage is clean enough (like healthcheck).
