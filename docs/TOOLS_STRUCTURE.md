# Tools Package Structure

## Overview

The MCP server now uses a proper Go package structure with all tool implementations in the `tools/` directory as a separate package.

## Structure

```
VeracodeMCP-Go/
├── go.mod                       # Module definition
├── main.go                      # Entry point
├── server.go                    # MCP server core
├── types.go                     # MCP protocol types
├── stdio.go, http.go           # Transport layers
├── tool_loader.go              # JSON tool definition loader
├── tool_handlers.go            # Handler registry (implements tools.HandlerRegistry)
├── tool_implementations.go     # Tool registry and LoadAllTools()
├── tools.json                  # Tool definitions
└── tools/                      # Tools package
    ├── types.go                # ToolImplementation interface
    ├── dynamic_findings.go     # DAST tool
    ├── static_findings.go      # SAST tool
    └── README.md              # Tool development guide
```

## How It Works

1. **tools/ package** (`package mcp_tools`):
   - Defines `ToolImplementation` interface
   - Defines `HandlerRegistry` interface
   - Defines `ToolHandler` function type
   - Contains all tool implementations

2. **main package**:
   - Imports `github.com/dipsylala/veracodemcp-go/mcp_tools`
   - Implements `ToolHandlerRegistry` (satisfies `tools.HandlerRegistry`)
   - Calls `LoadAllTools()` to initialize and register tools

3. **Build process**:
   - Standard `go build` works - no scripts needed!
   - Go automatically compiles the tools package
   - Clean, idiomatic Go structure

## Benefits

✅ **Idiomatic Go** - follows standard package organization  
✅ **No build scripts** - standard `go build` works  
✅ **Clean separation** - tools isolated from server code  
✅ **Easy to extend** - just add files to tools/  
✅ **Type safety** - interface enforcement at compile time  
✅ **Testable** - tools can be tested independently  

## Adding Tools

**Only 2 steps - tools auto-register!**

1. Add JSON definition to `tools.json`
2. Create `mcp_tools/my_tool.go` with `package mcp_tools` and `init()` function

**No other changes needed!**

### Example

```go
// tools/my_tool.go
package mcp_tools

// Auto-register on import
func init() {
	RegisterTool("my-tool", func() ToolImplementation {
		return NewMyTool()
	})
}

func NewMyTool() *MyTool { ... }

func (t *MyTool) RegisterHandlers(registry HandlerRegistry) error {
	registry.RegisterHandler("my-action", t.handleAction)
	return nil
}

func (t *MyTool) handleAction(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	// Implementation
}
```

Then just run `go build` - the tool is automatically discovered and loaded!
