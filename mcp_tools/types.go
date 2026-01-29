package tools

import "context"

// Context key for UI capability
type contextKey string

const uiCapabilityKey contextKey = "uiCapability"

// ClientSupportsUIFromContext retrieves UI capability from context
// Returns true if the client supports MCP Apps UI (text/html;profile=mcp-app)
func ClientSupportsUIFromContext(ctx context.Context) bool {
	if val := ctx.Value(uiCapabilityKey); val != nil {
		if supportsUI, ok := val.(bool); ok {
			return supportsUI
		}
	}
	return false
}

// ToolImplementation defines the interface that all MCP tools must implement
type ToolImplementation interface {
	// Initialize is called when the tool is loaded
	Initialize() error

	// RegisterHandlers registers the tool's handler functions
	RegisterHandlers(registry HandlerRegistry) error

	// Shutdown is called when the server is shutting down
	Shutdown() error
}

// HandlerRegistry allows tools to register their handler functions
type HandlerRegistry interface {
	RegisterHandler(toolName string, handler ToolHandler)
}

// ToolHandler is the function signature for tool execution handlers
type ToolHandler func(ctx context.Context, params map[string]interface{}) (interface{}, error)
