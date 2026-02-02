package tools

import (
	"github.com/dipsylala/veracodemcp-go/internal/mcp_tools"
)

// ToolHandlerRegistry manages the mapping of tool names to their handler functions
// It implements the mcp_tools.HandlerRegistry interface
type ToolHandlerRegistry struct {
	handlers map[string]mcp_tools.ToolHandler
}

// NewToolHandlerRegistry creates a new handler registry
func NewToolHandlerRegistry() *ToolHandlerRegistry {
	return &ToolHandlerRegistry{
		handlers: make(map[string]mcp_tools.ToolHandler),
	}
}

// RegisterHandler adds a handler for a specific tool (implements mcp_tools.HandlerRegistry)
func (r *ToolHandlerRegistry) RegisterHandler(toolName string, handler mcp_tools.ToolHandler) {
	r.handlers[toolName] = handler
}

// GetHandler retrieves the handler for a specific tool
func (r *ToolHandlerRegistry) GetHandler(toolName string) (mcp_tools.ToolHandler, bool) {
	handler, exists := r.handlers[toolName]
	return handler, exists
}
