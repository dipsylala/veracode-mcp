package tools

import (
	"github.com/dipsylala/veracodemcp-go/tools"
)

// ToolHandlerRegistry manages the mapping of tool names to their handler functions
// It implements the tools.HandlerRegistry interface
type ToolHandlerRegistry struct {
	handlers map[string]tools.ToolHandler
}

// NewToolHandlerRegistry creates a new handler registry
func NewToolHandlerRegistry() *ToolHandlerRegistry {
	return &ToolHandlerRegistry{
		handlers: make(map[string]tools.ToolHandler),
	}
}

// RegisterHandler adds a handler for a specific tool (implements tools.HandlerRegistry)
func (r *ToolHandlerRegistry) RegisterHandler(toolName string, handler tools.ToolHandler) {
	r.handlers[toolName] = handler
}

// GetHandler retrieves the handler for a specific tool
func (r *ToolHandlerRegistry) GetHandler(toolName string) (tools.ToolHandler, bool) {
	handler, exists := r.handlers[toolName]
	return handler, exists
}
