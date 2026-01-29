package mcp_tools

import "sync"

// Global registry for auto-registration of tools
var (
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu   sync.RWMutex
)

// RegisterTool registers a tool constructor for auto-discovery
// This is called by init() functions in each tool file
func RegisterTool(name string, constructor func() ToolImplementation) {
	registryMu.Lock()
	defer registryMu.Unlock()
	toolRegistry[name] = constructor
}

// RegisteredTool wraps a tool implementation with its registered name
type RegisteredTool struct {
	Name string
	Impl ToolImplementation
}

// GetAllTools returns instances of all registered tools with their names
// This is called by the main package during server initialization
func GetAllTools() []RegisteredTool {
	registryMu.RLock()
	defer registryMu.RUnlock()

	tools := make([]RegisteredTool, 0, len(toolRegistry))
	for name, constructor := range toolRegistry {
		tools = append(tools, RegisteredTool{
			Name: name,
			Impl: constructor(),
		})
	}
	return tools
}
