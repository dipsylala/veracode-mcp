# Quick Start: Adding a New Tool

This guide shows you exactly what to do when adding a new MCP tool.

## Step-by-Step Workflow

### 1. Define in `tools.json`

Add your tool definition to the `tools.json` file in the root directory:

```json
{
  "tools": [
    {
      "name": "your-tool-name",
      "description": "Clear description of what this tool does and when to use it",
      "category": "analysis",
      "params": [
        {
          "name": "param_name",
          "type": "string",
          "is_required": true,
          "description": "What this parameter does"
        }
      ]
    }
  ]
}
```

### 2. Create implementation in `tools/` directory

Create `tools/your_tool.go`:

```go
package main

import (
	"log"
)

type YourTool struct {
	name        string
	description string
}

func NewYourTool() *YourTool {
	return &YourTool{
		name:        "your-tool-name",
		description: "Brief description",
	}
}

func (t *YourTool) Name() string { return t.name }
func (t *YourTool) Description() string { return t.description }
func (t *YourTool) Initialize() error {
	log.Printf("Initializing tool: %s", t.name)
	return nil
}
func (t *YourTool) Shutdown() error { return nil }

func (t *YourTool) RegisterHandlers(registry *ToolHandlerRegistry) error {
	registry.Register("your-tool-name", t.handleYourTool)
	return nil
}

func (t *YourTool) handleYourTool(args map[string]interface{}) (*CallToolResult, error) {
	// Your implementation here
	return &CallToolResult{
		Content: []Content{{Type: "text", Text: "Success"}},
	}, nil
}
```

### 3. Register in `tool_implementations.go`

Add one line to the `LoadAllTools()` function:

```go
func (r *ToolImplRegistry) LoadAllTools() error {
	if err := r.Register(NewDynamicFindingsTool()); err != nil {
		log.Printf("Warning: Failed to register dynamic findings tool: %v", err)
	}
	if err := r.Register(NewStaticFindingsTool()); err != nil {
		log.Printf("Warning: Failed to register static findings tool: %v", err)
	}
	// ADD THIS LINE:
	if err := r.Register(NewYourTool()); err != nil {
		log.Printf("Warning: Failed to register your tool: %v", err)
	}
	return nil
}
```

### 4. Build and test

```powershell
# Build
.\build.ps1

# Test
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | .\mcp-server.exe -mode stdio
```

## That's It!

Three files to touch:
1. `tools.json` - Define parameters and description
2. `tools/your_tool.go` - Implement the logic
3. `tool_implementations.go` - Register it (one line)

Then build and run. Your tool is now available via MCP!

## Tips

- Copy `tools/dynamic_findings.go` as a template
- Tool name in code MUST match name in tools.json
- Validate all required parameters in your handler
- Return structured, readable text output
- Use Initialize() for setup, Shutdown() for cleanup
