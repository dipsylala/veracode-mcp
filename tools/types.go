package tools

import "context"

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
