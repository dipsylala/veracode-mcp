# Plugin Architecture

The Veracode MCP Server uses a plugin architecture to extend functionality while following the **open/closed principle** - open for extension, closed for modification.

## Overview

Plugins allow you to add new MCP tools without modifying the core server code. Each plugin:
- Implements the `Plugin` interface
- Registers handlers for one or more tools
- Has lifecycle management (Initialize, Shutdown)
- Is auto-discovered from the `plugins/` directory

## Creating a Plugin

### 1. Define Tool Schemas

Add your tool definitions to `tools.json`:

```json
{
  "name": "my-tool",
  "description": "What this tool does",
  "params": [
    {
      "name": "param1",
      "description": "Parameter description",
      "type": "string",
      "is_required": true
    }
  ],
  "category": "analysis"
}
```

### 2. Implement the Plugin Interface

Create a new file in `plugins/`:

```go
package plugins

import (
    "fmt"
)

// MyPlugin implements the Plugin interface
type MyPlugin struct {
    // Plugin-specific fields
    config map[string]interface{}
}

// Name returns the plugin name
func (p *MyPlugin) Name() string {
    return "my-plugin"
}

// Description returns what the plugin does
func (p *MyPlugin) Description() string {
    return "My plugin provides awesome functionality"
}

// Initialize performs plugin setup
func (p *MyPlugin) Initialize() error {
    p.config = make(map[string]interface{})
    // Load config, connect to APIs, etc.
    return nil
}

// RegisterHandlers registers tool handlers
func (p *MyPlugin) RegisterHandlers(registry *ToolHandlerRegistry) {
    registry.Register("my-tool", p.handleMyTool)
}

// Shutdown cleans up plugin resources
func (p *MyPlugin) Shutdown() error {
    // Close connections, save state, etc.
    return nil
}

// Handler implementation
func (p *MyPlugin) handleMyTool(args map[string]interface{}) (*CallToolResult, error) {
    // Extract and validate parameters
    param1, ok := args["param1"].(string)
    if !ok || param1 == "" {
        return &CallToolResult{
            Content: []Content{{
                Type: "text",
                Text: "Error: param1 is required",
            }},
            IsError: true,
        }, nil
    }

    // Implement tool logic
    result := fmt.Sprintf("Processed: %s", param1)

    return &CallToolResult{
        Content: []Content{{
            Type: "text",
            Text: result,
        }},
    }, nil
}
```

### 3. Register in LoadAllPlugins

Update `plugin.go` to include your plugin:

```go
func (r *PluginRegistry) LoadAllPlugins() error {
    // Register built-in plugins
    r.Register(&VeracodeFindingsPlugin{})
    r.Register(&MyPlugin{})  // Add your plugin here
    
    return nil
}
```

## Plugin Lifecycle

1. **Registration**: Plugin is registered with `PluginRegistry.Register()`
2. **Initialization**: Plugin's `Initialize()` method is called
3. **Handler Registration**: Plugin registers tool handlers via `RegisterHandlers()`
4. **Active**: Handlers respond to tool calls
5. **Shutdown**: Plugin's `Shutdown()` method is called on server shutdown

## Best Practices

### Parameter Validation

Always validate required parameters:

```go
func (p *MyPlugin) handleTool(args map[string]interface{}) (*CallToolResult, error) {
    // Check required parameters
    param, ok := args["param"].(string)
    if !ok || param == "" {
        return &CallToolResult{
            Content: []Content{{
                Type: "text",
                Text: "Error: param is required and must be a string",
            }},
            IsError: true,
        }, nil
    }
    
    // Process...
}
```

### Error Handling

Return errors as tool results, not Go errors:

```go
// Good - Error visible to LLM
return &CallToolResult{
    Content: []Content{{Type: "text", Text: "API connection failed"}},
    IsError: true,
}, nil

// Bad - Error not visible to LLM
return nil, fmt.Errorf("API connection failed")
```

### Resource Management

Use Initialize/Shutdown for resource management:

```go
type MyPlugin struct {
    client *APIClient
}

func (p *MyPlugin) Initialize() error {
    client, err := NewAPIClient()
    if err != nil {
        return err
    }
    p.client = client
    return nil
}

func (p *MyPlugin) Shutdown() error {
    if p.client != nil {
        return p.client.Close()
    }
    return nil
}
```

### Structured Output

Return formatted, readable results:

```go
result := fmt.Sprintf(`Analysis Complete
================

Status: Success
Items Found: %d
Time: %s

Details:
%s`, count, duration, details)

return &CallToolResult{
    Content: []Content{{Type: "text", Text: result}},
}, nil
```

## Example Plugins

### Veracode Findings Plugin

Located in `plugins/veracode_findings.go`, this plugin:
- Handles `get-dynamic-findings` and `get-static-findings` tools
- Validates `application_path` parameter
- Connects to Veracode API (TODO)
- Returns formatted vulnerability findings

### Future Plugin Ideas

- **Veracode SCA Plugin**: Software Composition Analysis tools
- **Veracode SBOM Plugin**: Software Bill of Materials generation
- **Workspace Plugin**: Manage `.veracode-workspace.json` files
- **Report Plugin**: Generate PDF/HTML reports
- **Authentication Plugin**: Centralized credential management

## Architecture Benefits

1. **Separation of Concerns**: Each plugin focuses on one domain
2. **Independent Development**: Add features without touching core code
3. **Easy Testing**: Test plugins in isolation
4. **Versioning**: Plugins can be versioned independently
5. **Optional Features**: Enable/disable plugins as needed
6. **Clean Dependencies**: Plugins have their own dependencies

## Plugin Discovery

Currently, plugins are statically registered in `LoadAllPlugins()`. Future enhancements could include:
- Dynamic loading from compiled `.so` files (requires CGO)
- Configuration-based plugin enabling/disabling
- Plugin marketplace/registry
- Hot-reload during development
