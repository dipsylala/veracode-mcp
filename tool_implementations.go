package main

import (
	"log"

	"github.com/dipsylala/veracodemcp-go/tools"
)

// ToolImplRegistry manages all loaded tool implementations
type ToolImplRegistry struct {
	tools map[string]tools.ToolImplementation
}

// NewToolImplRegistry creates a new tool implementation registry
func NewToolImplRegistry() *ToolImplRegistry {
	return &ToolImplRegistry{
		tools: make(map[string]tools.ToolImplementation),
	}
}

// Register adds a tool implementation to the registry
func (r *ToolImplRegistry) Register(tool tools.ToolImplementation) error {
	name := tool.Name()
	if _, exists := r.tools[name]; exists {
		return nil // Already registered, skip
	}

	// Initialize the tool
	if err := tool.Initialize(); err != nil {
		return err
	}

	r.tools[name] = tool
	return nil
}

// Get retrieves a tool implementation by name
func (r *ToolImplRegistry) Get(name string) (tools.ToolImplementation, bool) {
	tool, exists := r.tools[name]
	return tool, exists
}

// GetAll returns all registered tool implementations
func (r *ToolImplRegistry) GetAll() []tools.ToolImplementation {
	result := make([]tools.ToolImplementation, 0, len(r.tools))
	for _, tool := range r.tools {
		result = append(result, tool)
	}
	return result
}

// ShutdownAll calls Shutdown on all tool implementations
func (r *ToolImplRegistry) ShutdownAll() {
	for _, tool := range r.tools {
		if err := tool.Shutdown(); err != nil {
			log.Printf("Error shutting down tool: %v", err)
		}
	}
}

// LoadAllTools initializes and registers all available tool implementations
// Tools are automatically discovered via init() functions in the tools package
func LoadAllTools(registry *ToolImplRegistry, handlerRegistry *ToolHandlerRegistry) error {
	// Get all auto-registered tools from the tools package
	allTools := tools.GetAllTools()

	for _, tool := range allTools {
		// Initialize the tool
		if err := tool.Initialize(); err != nil {
			log.Printf("Failed to initialize tool %s: %v", tool.Name(), err)
			continue
		}

		// Register the tool in the implementation registry
		if err := registry.Register(tool); err != nil {
			log.Printf("Failed to register tool %s: %v", tool.Name(), err)
			continue
		}

		// Register the tool's handlers
		if err := tool.RegisterHandlers(handlerRegistry); err != nil {
			log.Printf("Failed to register handlers for tool %s: %v", tool.Name(), err)
			continue
		}

		log.Printf("Successfully loaded tool: %s", tool.Name())
	}

	return nil
}
