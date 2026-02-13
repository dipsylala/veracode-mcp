# REST API Integration Architecture

## Overview

This document describes how REST API calls are integrated with the MCP server tools, addressing the many:many relationship between tools and API endpoints.

## Package Structure

```text
VeracodeMCP-Go/
├── api/                        # API package
│   ├── client.go              # Client orchestrator (manages all API clients)
│   ├── findings.go            # Findings endpoints (SAST/DAST) with filtering and business logic
│   ├── health.go              # Health check endpoints with business logic
│   ├── applications.go        # Application lookup and management
│   ├── generated/             # Swagger-generated API clients (DO NOT EDIT)
│   │   ├── healthcheck/       # Healthcheck API client
│   │   ├── findings/          # Findings API client
│   │   ├── dynamic_flaw/      # Dynamic flaw API client
│   │   ├── static_finding_data_path/  # Static finding data path API client
│   │   ├── applications/      # Applications API client
│   │   └── README.md
│   └── README.md              # API package documentation
│
└── tools/                      # Tools package (USES api/)
    ├── api_health.go          # Uses api.CheckHealth()
    ├── dynamic_findings.go    # Uses api.GetDynamicFindings()
    └── static_findings.go     # Uses api.GetStaticFindings()
```

## Architecture Principles

### 1. Separation of Concerns

**Generated Code** (`api/generated/`):

- ✅ Swagger-generated HTTP client
- ✅ Never manually edited
- ✅ Regenerate when API spec changes
- ❌ Don't import directly in tools

**API Client Orchestrator** (`api/client.go`):

- ✅ Manages all 4 generated API clients
- ✅ Handles authentication (VERACODE_API_ID, VERACODE_API_KEY)
- ✅ Provides GetAuthContext() for HMAC signing
- ✅ Single point of initialization

**API Business Logic** (`api/`):

- ✅ Clean Go interfaces over raw API clients
- ✅ Type conversions (optional.Interface → simple structs)
- ✅ Business logic (filtering, pagination)
- ✅ Reusable across multiple tools

**Tools** (`mcp_tools/`):

- ✅ MCP tool implementations
- ✅ Import api package only (not generated)
- ✅ Focus on tool-specific logic
- ✅ Return MCP-formatted responses

### 2. Many-to-Many Relationships

**One Tool → Multiple API Calls:**

```text
get-dynamic-findings tool
├── api.GetDynamicFindings()
├── api.GetFindingByID()
└── api.GetMitigations()
```

**One API Call → Multiple Tools:**

```text
api.GetDynamicFindings()
├── get-dynamic-findings tool
├── compare-scans tool (future)
└── generate-report tool (future)
```

### 3. Maintainability

✅ **Adding new API endpoint**: Create file in `api/`, no tool changes  
✅ **Adding new tool**: Import `api/`, use existing methods  
✅ **Regenerating client**: Only affects `api/generated/`, api wrapper shields tools  
✅ **Testing**: Mock api package, not generated client  

## Usage Pattern

### Step 1: Define API Method (api/)

```go
// api/findings.go
package api

import (
    "github.com/dipsylala/veracodemcp-go/api/generated/findings"
)

func (c *Client) GetDynamicFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error) {
    authCtx := c.GetAuthContext(ctx)
    
    // Build options for the generated client
    opts := &findings.ApplicationFindingsInformation_ApiGetFindingsUsingGETOpts{
        ScanType: optional.NewInterface([]string{"DYNAMIC"}),
    }
    
    // Call generated client
    resp, _, err := c.findingsClient.ApplicationFindingsInformation_Api.GetFindingsUsingGET(
        authCtx, req.AppProfile, opts)
    if err != nil {
        return nil, fmt.Errorf("API call failed: %w", err)
    }
    
    // Transform complex API response to simple Finding structs
    convertedFindings := convertFindings(resp.Embedded.Findings, "DYNAMIC")
    
    // Apply client-side filtering
    filteredFindings := applyFilters(convertedFindings, req)
    
    return &FindingsResponse{
        Findings: filteredFindings,
        TotalCount: len(filteredFindings),
    }, nil
}
```

### Step 2: Use in Tool (tools/)

```go
// tools/dynamic_findings.go
import "github.com/dipsylala/veracodemcp-go/api"

func (t *DynamicFindingsTool) handleGetDynamicFindings(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Create API client
    client, err := api.NewClient()
    if err != nil {
        return errorResponse("API not configured: " + err.Error()), nil
    }
    
    // Build request from tool parameters
    req := api.FindingsRequest{
        AppProfile: params["app_profile"].(string),
        Severity:   params["severity"].([]string),
        // ... other params
    }
    
    // Call API
    findings, err := client.GetDynamicFindings(ctx, req)
    if err != nil {
        return errorResponse("Failed to get findings: " + err.Error()), nil
    }
    
    // Format for MCP response
    responseText := formatFindings(findings)
    
    return successResponse(responseText), nil
}
```

## Authentication

Environment variables required:

- `VERACODE_API_ID` - Your Veracode API ID
- `VERACODE_API_KEY` - Your Veracode API Secret Key

The `api.Client` handles authentication:

```go
// Client creation checks credentials
client, err := api.NewClient()
if err != nil {
    // Credentials not set or invalid
}

// All API calls use authenticated context
status, err := client.CheckHealth(ctx)
```

## Adding New Endpoints

### 1. Obtain OpenAPI/Swagger Spec

Get the API specification from Veracode documentation or SwaggerHub and save it to `specs/`:

```bash
# Example: Download or save spec file
curl https://api.veracode.com/swagger/new-api.yaml > specs/veracode-new-api.yaml
```

### 2. Update Regeneration Script

Edit `scripts/regenerate-api-clients.ps1` and add your API to the `$apis` array:

```powershell
$apis = @(
    # ... existing APIs ...
    @{
        Name = "new_api"
        Spec = "specs/veracode-new-api.yaml"
        Package = "new_api"
    }
)
```

### 3. Generate Client

Run the regeneration script:

```powershell
.\scripts\regenerate-api-clients.ps1
```

This will:

- Generate the client in `api/generated/new_api/`
- Set the package name to `new_api`
- Run `go mod tidy`

```go
import new_api "github.com/dipsylala/veracodemcp-go/api/generated/new_api"

type Client struct {
    healthcheckClient           *healthcheck.APIClient
    findingsClient              *findings.APIClient
    dynamicFlawClient           *dynamic_flaw.APIClient
    staticFindingDataPathClient *static_finding_data_path.APIClient
    newApiClient                *new_api.APIClient  // Add this
    apiID                       string
    apiKey                      string
}

func NewClient() (*Client, error) {
    // ... existing code ...
    newApiCfg := new_api.NewConfiguration()
    
    return &Client{
        // ... existing clients ...
        newApiClient: new_api.NewAPIClient(newApiCfg),
        // ...
    }, nil
}
```

### 4. Integrate into api/client.go

Add the new client to `api/client.go`:

Only create business logic in `api/` if you need type conversion, filtering, or business logic:

```go
// api/new_feature.go
package api

type NewFeatureRequest struct {
    Param1 string `json:"param1"`
}

type NewFeatureResponse struct {
    Result string `json:"result"`
}

func (c *Client) NewFeature(ctx context.Context, req NewFeatureRequest) (*NewFeatureResponse, error) {
    authCtx := c.GetAuthContext(ctx)
    
    resp, _, err := c.newApiClient.SomeApi.SomeMethod(authCtx, req.Param1)
    if err != nil {
        return nil, err
    }
    
    // Add any type conversion, filtering, business logic here
    
    return &NewFeatureResponse{
        Result: resp.Data,
    }, nil
}
```

### 5. Use in Tools

For simple APIs, skip the business logic layer and call the generated client directly from tools.

## Benefits

✅ **Clean Architecture**: Generated code isolated from business logic  
✅ **Reusability**: API methods shared across tools  
✅ **Type Safety**: Strong typing in api package  
✅ **Testability**: Easy to mock api package  
✅ **Maintainability**: Changes localized to specific layers  
✅ **2-Step Workflow**: Adding tools still only requires tools.json + tools/file.go  

## Example: Real API Integration

### Current State (Placeholder)

```go
// tools/api_health.go
func (t *APIHealthTool) handleAPIHealth(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // TEMPORARY: Returns hardcoded response
    return placeholderResponse(), nil
}
```

### After Integration

```go
// tools/api_health.go
import "github.com/dipsylala/veracodemcp-go/api"

func (t *APIHealthTool) handleAPIHealth(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    client, err := api.NewClient()
    if err != nil {
        return errorResponse(err.Error()), nil
    }
    
    status, err := client.CheckHealth(ctx)
    if err != nil {
        return errorResponse(err.Error()), nil
    }
    
    return successResponse(fmt.Sprintf("API is %s (status: %d)", 
        status.Message, status.StatusCode)), nil
}
```

## Error Handling

API errors should be returned as tool results, not Go errors:

```go
// ❌ BAD: Error not visible to LLM
findings, err := client.GetFindings(ctx, req)
if err != nil {
    return nil, err  // User/LLM won't see this!
}

// ✅ GOOD: Error visible to LLM
findings, err := client.GetFindings(ctx, req)
if err != nil {
    return map[string]interface{}{
        "error": fmt.Sprintf("Failed to get findings: %v", err),
    }, nil
}
```

## Testing

### Mock API Package

```go
// tools/dynamic_findings_test.go
type mockAPIClient struct{}

func (m *mockAPIClient) GetDynamicFindings(ctx context.Context, req api.FindingsRequest) (*api.FindingsResponse, error) {
    return &api.FindingsResponse{
        Findings: []api.Finding{
            {ID: "1", Severity: "High"},
        },
    }, nil
}

func TestGetDynamicFindings(t *testing.T) {
    tool := NewDynamicFindingsTool()
    // Inject mock client
    // Test tool behavior
}
```

## Next Steps

1. ✅ **Structure created**: api/ package with client.go, business logic files, generated/
2. ✅ **5 API clients integrated**: healthcheck, findings, dynamic_flaw, static_finding_data_path, applications
3. ✅ **Business logic created**: findings.go with filtering/conversion, health.go, applications.go
4. ⏳ **Implement HMAC auth**: Add Veracode HMAC signature to GetAuthContext()
5. ⏳ **Update tools**: Replace placeholders with actual API calls
6. ⏳ **Add tests**: Unit tests for api package, integration tests for tools
7. ⏳ **Add more endpoints**: SCA, SBOM, Upload, etc.

## Summary

**Where to put REST API code?**

- **Generated client**: `api/generated/` (auto-generated by Swagger, don't edit)
- **Client orchestrator**: `api/client.go` (manages all API clients, authentication)
- **API business logic**: `api/` (business logic, filtering, type conversion)
- **Tool usage**: `tools/` package (imports api, uses business logic or client directly)

**Many:many relationships handled by:**

- API business logic shared across multiple tools
- Tools can call multiple API business logic functions
- Clean separation via 3-layer architecture: tools → business logic → generated clients
