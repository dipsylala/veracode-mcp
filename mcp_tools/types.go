package mcp_tools

import (
	"context"
	"log"
)

// Context key for UI capability (exported for use by server package)
type contextKey string

// UICapabilityKey is the context key for UI capability information
const UICapabilityKey contextKey = "veracode-mcp:ui-capability"

// ClientSupportsUIFromContext retrieves UI capability from context
// Returns true if the client supports MCP Apps UI (text/html;profile=mcp-app)
func ClientSupportsUIFromContext(ctx context.Context) bool {
	val := ctx.Value(UICapabilityKey)
	log.Printf("[UI-CONTEXT] Context value for UICapabilityKey: %v (type: %T)", val, val)

	if val != nil {
		if supportsUI, ok := val.(bool); ok {
			log.Printf("[UI-CONTEXT] Returning: %v", supportsUI)
			return supportsUI
		}
		log.Printf("[UI-CONTEXT] Value was not a bool, returning false")
	} else {
		log.Printf("[UI-CONTEXT] Value was nil, returning false")
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
