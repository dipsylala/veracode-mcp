# API Package

This package provides a clean interface to Veracode APIs, organized in three layers:

```text
api/
├── client.go           # Client orchestrator (manages all API clients)
├── helpers/            # Business logic wrappers
│   ├── findings.go     # SAST/DAST findings with filtering
│   └── health.go       # Authenticated principal health check
└── generated/          # Auto-generated Swagger clients
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

- `NewClient()` - Creates the authenticated REST and XML API clients
- `GetAuthContext()` - Adds Veracode HMAC authentication to requests
- `IsConfigured()` - Checks if credentials are set
- Holds generated clients for findings, flaw details, applications, and policy APIs

### helpers/ (Business Logic)

#### helpers/health.go

Checks API connectivity and credentials with an authenticated request to
`GET /api/authn/v2/principal`:

- `CheckHealth()` - Returns health status struct
- `CheckHealthSimple()` - Returns simple boolean
- Used by `api-health` tool

A `200 OK` response means the regional Veracode API is reachable and the HMAC
credentials were accepted. Non-200 responses and transport errors are returned
as unavailable health status values. The response body is not used or exposed.

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

- `findings/` - Findings API client (SAST/DAST/SCA)
- `dynamic_flaw/` - Dynamic flaw details API client
- `static_finding_data_path/` - Static finding data path API client

See [generated/README.md](generated/README.md) for regeneration instructions.

## Usage in Tools

Tools import and use the API package:

```go
package tools

import (
    "github.com/dipsylala/veracode-mcp/api"
)

func (t *APIHealthTool) handleAPIHealth(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Create API client
    client, err := api.NewClient()
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
   import new_api "github.com/dipsylala/veracode-mcp/api/generated/new_api"
   
   type Client struct {
       // ... existing clients ...
       newApiClient *new_api.APIClient
   }
   
   func NewClient() (*Client, error) {
       // ... existing clients ...
       newApiCfg := new_api.NewConfiguration()
       
       return &Client{
           // ... existing clients ...
           newApiClient: new_api.NewAPIClient(newApiCfg),
       }
   }
   ```

4. **Create helper (if needed)**:

   ```go
   // api/helpers/new_feature.go
   package api
   
   func (c *Client) DoSomething(ctx context.Context, params Request) (*Response, error) {
       authCtx := c.GetAuthContext(ctx)
       resp, _, err := c.newApiClient.SomeApi.SomeMethod(authCtx, ...)
       // Add filtering, type conversion, business logic
       return result, nil
   }
   ```

5. **Use in tools**:

   ```go
   import "github.com/dipsylala/veracode-mcp/api"
   
   client, _ := api.NewClient()
   result, err := client.DoSomething(ctx, params)
   ```

## When to Create Helpers

Create helpers in `helpers/` when you need:

- **Type conversion**: Complex optional.Interface → simple structs
- **Client-side filtering**: API doesn't support all filter combinations
- **Pagination logic**: Handling page/size parameters
- **Severity mapping**: Converting numeric codes to human-readable strings
- **Business logic**: Combining multiple API calls or adding domain logic

Skip helpers only when direct generated-client usage is clean enough. The
health check is implemented in `rest/health.go` because it uses the Identity API
principal endpoint rather than a generated API client.
