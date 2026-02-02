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

### 2. Create implementation in `mcp_tools/` directory

Create `mcp_tools/your_tool.go`:

```go
package mcp_tools

import (
  "context"
  "log"
)

// Auto-register this tool when the package is imported
func init() {
  RegisterTool("your-tool-name", func() ToolImplementation {
    return NewYourTool()
  })
}

type YourTool struct{}

func NewYourTool() *YourTool {
  return &YourTool{}
}

func (t *YourTool) Initialize() error {
  log.Printf("Initializing tool: your-tool-name")
  return nil
}

func (t *YourTool) Shutdown() error { return nil }

func (t *YourTool) RegisterHandlers(registry HandlerRegistry) error {
  registry.RegisterHandler("your-tool-name", t.handleYourTool)
  return nil
}

func (t *YourTool) handleYourTool(ctx context.Context, args map[string]interface{}) (interface{}, error) {
  // Your implementation here
  return map[string]interface{}{
    "content": []map[string]string{{
      "type": "text",
      "text": "Success",
    }},
  }, nil
}
```

### 3. Build and test - that's it!

Tools automatically register themselves when you import the tools package! No manual registration needed.

### 4. Build and test

```powershell
# Build
.\build.ps1

# Test
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | .\mcp-server.exe -mode stdio
```

## That's It!

Two files to touch:
1. `tools.json` - Define parameters and description
2. `mcp_tools/your_tool.go` - Implement the logic with auto-registration

Then build and run. Your tool is now available via MCP!

## Tips

- Copy `mcp_tools/dynamic_findings.go` as a template
- Tool name in code MUST match name in tools.json
- Validate all required parameters in your handler
- Return structured, readable text output
- Use Initialize() for setup, Shutdown() for cleanup
