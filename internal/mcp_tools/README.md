# MCP Tools Directory

This directory contains all MCP tool implementations. **Tools automatically register themselves** - just add the file!

## Quick Start: Adding a New Tool

### Step 1: Add to `tools.json`

```json
{
  "name": "my-action",
  "description": "What this tool does",
  "category": "veracode",
  "params": [
    {
      "name": "application_path",
      "type": "string",
      "isRequired": true,
      "description": "Absolute path to application"
    }
  ]
}
```

### Step 2: Create `mcp_tools/my_tool.go`

```go
package mcp_tools

import (
 "context"
 "fmt"
 "log"
)

// Auto-register when package loads
func init() {
 RegisterTool("my-tool", func() ToolImplementation {
  return NewMyTool()
 })
}

type MyTool struct {
 name        string
 description string
}

func NewMyTool() *MyTool {
 return &MyTool{
  name:        "my-tool",
  description: "My tool description",
 }
}

func (t *MyTool) Name() string { return t.name }
func (t *MyTool) Description() string { return t.description }

func (t *MyTool) Initialize() error {
 log.Printf("Initializing: %s", t.name)
 return nil
}

func (t *MyTool) RegisterHandlers(registry HandlerRegistry) error {
 registry.RegisterHandler("my-action", t.handleMyAction)
 return nil
}

func (t *MyTool) Shutdown() error { return nil }

func (t *MyTool) handleMyAction(ctx context.Context, params map[string]interface{}) (interface{}, error) {
 appPath, _ := params["application_path"].(string)
 
 return map[string]interface{}{
  "content": []map[string]string{{
   "type": "text",
   "text": fmt.Sprintf("Result: %s", appPath),
  }},
 }, nil
}
```

### Step 3: Build

```powershell
go build -o mcp-server.exe .
```

**That's it!** No manual registration needed. The `init()` function auto-registers your tool.

## How Auto-Registration Works

1. Each tool file includes an `init()` function
2. When Go imports the `tools` package, all `init()` functions run
3. Tools call `RegisterTool()` to add themselves to the registry
4. Server calls `tools.GetAllTools()` to get all registered tools

## Required Components

### 1. init() Function (Required)

```go
func init() {
    RegisterTool("tool-name", func() ToolImplementation {
        return NewMyTool()
    })
}
```

### 2. Implement ToolImplementation Interface

```go
type ToolImplementation interface {
 Name() string
 Description() string
 Initialize() error
 RegisterHandlers(registry HandlerRegistry) error
 Shutdown() error
}
```

### 3. Handler Function Signature

```go
func (t *MyTool) handleAction(ctx context.Context, params map[string]interface{}) (interface{}, error)
```

## Return Formats

### Success

```go
return map[string]interface{}{
 "content": []map[string]string{{
  "type": "text",
  "text": "Your result here",
 }},
}, nil
```

### User Error

```go
return map[string]interface{}{
 "error": "Error message for user",
}, nil
```

### System Error

```go
return nil, fmt.Errorf("critical error: %w", err)
```

## Best Practices

✅ **Always include init()** - Required for auto-registration  
✅ **Package name: `tools`** - Must be in tools package  
✅ **Export types** - Tool struct and New function capitalized  
✅ **Validate early** - Check required parameters first  
✅ **Log initialization** - Use `log.Printf()` for debugging  
✅ **Clean up** - Implement Shutdown for resource cleanup  
✅ **Use context** - Pass through for cancellation support  

## Files in This Directory

- `types.go` - Interface definitions
- `registry.go` - Auto-registration system
- `dynamic_findings.go` - DAST tool example
- `static_findings.go` - SAST tool example
- `package_workspace.go` - Workspace packaging tool
- `run_sca_scan.go` - SCA scan tool
- `get_local_sca_results.go` - SCA results parser
- `pipeline_scan.go` - Pipeline scan tool
- `README.md` - This file

## Checklist for New Tool

- [ ] Add definition to `tools.json`
- [ ] Create `mcp_tools/my_tool.go`
- [ ] Add `package mcp_tools` at top
- [ ] Include `init()` with `RegisterTool()`
- [ ] Implement all interface methods
- [ ] Run `go build`
- [ ] Test with `tools/list`

No changes to any other files needed!
