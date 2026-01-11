# Adding New Tools to the MCP Server

This guide shows how to add new tools to the Veracode MCP server using the auto-registration architecture.

## Overview

The tool system uses **auto-registration** - tools register themselves on import using `init()` functions. You only need to:
1. Create a new file in `tools/`
2. Implement the `ToolImplementation` interface
3. Register it in `init()`

That's it! No manual registration needed.

## Step-by-Step Guide

### Step 1: Create Tool File

Create `tools/your_tool.go`:

```go
package tools

import (
    "context"
)

func init() {
    RegisterTool("your-tool-name", func() ToolImplementation {
        return &YourTool{}
    })
}

type YourTool struct{}

// GetName returns the MCP tool name
func (t *YourTool) GetName() string {
    return "your-tool-name"
}

// GetDescription returns the detailed description for LLMs
func (t *YourTool) GetDescription() string {
    return `Detailed description of what your tool does.

When to use:
- Use this tool when...
- Best for situations where...

Returns:
- Success: Description of data structure
- Errors: Common error conditions

Example usage:
{
  "param1": "value",
  "param2": 123
}`
}

// GetInputSchema defines the JSON schema for parameters
func (t *YourTool) GetInputSchema() interface{} {
    return map[string]interface{}{
        "type": "object",
        "properties": map[string]interface{}{
            "param1": map[string]interface{}{
                "type":        "string",
                "description": "Description of param1 with examples",
            },
            "param2": map[string]interface{}{
                "type":        "number",
                "description": "Optional numeric parameter",
                "minimum":     0,
                "maximum":     100,
            },
            "severity": map[string]interface{}{
                "type":        "array",
                "description": "Filter by severity levels",
                "items": map[string]interface{}{
                    "type": "string",
                    "enum": []string{"Critical", "High", "Medium", "Low"},
                },
            },
        },
        "required": []string{"param1"}, // List required parameters
    }
}

// Handle executes the tool logic
func (t *YourTool) Handle(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Extract required parameters
    param1, ok := params["param1"].(string)
    if !ok || param1 == "" {
        return errorResponse("param1 is required"), nil
    }

    // Extract optional parameters with defaults
    param2 := 50.0 // default
    if val, ok := params["param2"].(float64); ok {
        param2 = val
    }

    // Extract array parameters
    var severityFilter []string
    if val, ok := params["severity"].([]interface{}); ok {
        for _, item := range val {
            if str, ok := item.(string); ok {
                severityFilter = append(severityFilter, str)
            }
        }
    }

    // TODO: Implement your tool logic
    // Example: Call API
    // client, err := api.NewVeracodeClient()
    // if err != nil {
    //     return errorResponse(err.Error()), nil
    // }
    // result, err := client.SomeMethod(ctx, ...)

    // Return success response
    return successResponse(fmt.Sprintf("Processed: %s with %v", param1, severityFilter)), nil
}

// Helper functions for consistent responses
func successResponse(message string) map[string]interface{} {
    return map[string]interface{}{
        "success": true,
        "message": message,
    }
}

func errorResponse(message string) map[string]interface{} {
    return map[string]interface{}{
        "success": false,
        "error":   message,
    }
}
```
### Step 2: Supported JSON Schema Types

The `GetInputSchema()` method returns a standard JSON Schema. Common patterns:

**String parameter:**
```go
"param_name": map[string]interface{}{
    "type":        "string",
    "description": "Description here",
}
```

**Enum (limited values):**
```go
"status": map[string]interface{}{
    "type":        "string",
    "description": "Finding status",
    "enum":        []string{"Open", "Closed", "Mitigated"},
}
```

**Number with constraints:**
```go
"page_size": map[string]interface{}{
    "type":        "integer",
    "description": "Number of results per page",
    "minimum":     1,
    "maximum":     100,
    "default":     20,
}
```

**Boolean:**
```go
"include_resolved": map[string]interface{}{
    "type":        "boolean",
    "description": "Include resolved findings",
}
```

**Array of strings:**
```go
"severities": map[string]interface{}{
    "type":        "array",
    "description": "Filter by severity levels",
    "items": map[string]interface{}{
        "type": "string",
        "enum": []string{"Critical", "High", "Medium", "Low"},
    },
}
```

### Step 3: Using API Helpers

To call Veracode APIs, use the `api` package:

```go
import "github.com/dipsylala/veracodemcp-go/api"

func (t *YourTool) Handle(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Create API client
    client, err := api.NewVeracodeClient()
    if err != nil {
        return errorResponse("API not configured: " + err.Error()), nil
    }

    // For simple APIs - use generated client directly
    resp, err := client.healthcheckClient.HealthcheckAPIsApi.HealthcheckStatusGet(ctx)
    
    // For complex APIs - use helpers
    findings, err := client.GetDynamicFindings(ctx, api.FindingsRequest{
        AppProfile: params["app_profile"].(string),
        Severity:   extractStringArray(params, "severity"),
    })

    if err != nil {
        return errorResponse(err.Error()), nil
    }

    return successResponse(fmt.Sprintf("Found %d findings", findings.TotalCount)), nil
}

// Helper to extract string arrays from params
func extractStringArray(params map[string]interface{}, key string) []string {
    var result []string
    if val, ok := params[key].([]interface{}); ok {
        for _, item := range val {
            if str, ok := item.(string); ok {
                result = append(result, str)
            }
        }
    }
    return result
}
```

### Step 4: Test Your Tool

Build and run the server:

```powershell
.\build.ps1 -Quick
.\dist\mcp-server.exe
```

Test using the MCP Inspector or your LLM client.

## Benefits of Auto-Registration

✅ **No manual registration** - Just create the file, it's automatically discovered  
✅ **Type-safe** - Full Go type checking on parameters and responses  
✅ **Self-documenting** - Schema and description in code, not separate JSON  
✅ **Thread-safe** - Registry handles concurrent access automatically  
✅ **Testable** - Each tool is independently testable  

## Advanced Patterns

### Returning Structured Data

```go
func (t *YourTool) Handle(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Return structured data that LLMs can parse
    return map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "findings": []map[string]interface{}{
                {
                    "id":       "123",
                    "severity": "High",
                    "title":    "SQL Injection",
                },
            },
            "summary": map[string]interface{}{
                "total":    10,
                "critical": 2,
                "high":     5,
            },
        },
    }, nil
}
```

### Error Handling

Always return errors as tool results (not Go errors) so LLMs can see them:

```go
// ❌ BAD - LLM won't see this
return nil, fmt.Errorf("failed to fetch data")

// ✅ GOOD - LLM sees the error
return errorResponse("Failed to fetch data: " + err.Error()), nil
```

## Examples

See existing tools for reference:
- [tools/api_health.go](../tools/api_health.go) - Simple tool calling generated client
- [tools/dynamic_findings.go](../tools/dynamic_findings.go) - Using API helpers
- [tools/static_findings.go](../tools/static_findings.go) - Parameter extraction patterns

## Next Steps

1. Create your tool file in `tools/`
2. Implement the `ToolImplementation` interface
3. Register in `init()` function
4. Build and test
5. Integrate with Veracode APIs using `api` package
