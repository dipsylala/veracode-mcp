package tools

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

// GetAllTools returns instances of all registered tools
// This is called by the main package during server initialization
func GetAllTools() []ToolImplementation {
	registryMu.RLock()
	defer registryMu.RUnlock()

	tools := make([]ToolImplementation, 0, len(toolRegistry))
	for _, constructor := range toolRegistry {
		tools = append(tools, constructor())
	}
	return tools
}
