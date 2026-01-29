package tools

import (
	"context"
	"fmt"
	"log"

	"github.com/dipsylala/veracodemcp-go/internal/types"
	"github.com/dipsylala/veracodemcp-go/tools"
)

// ToolManager consolidates all tool-related registries and provides a unified
// interface for tool management. It coordinates between tool definitions,
// handler functions, and implementation instances.
type ToolManager struct {
	definitions     *ToolRegistry        // Tool schemas from tools.json
	handlers        *ToolHandlerRegistry // Handler function mappings
	implementations *ToolImplRegistry    // Tool implementation instances
}

// NewToolManager creates a new tool manager with all necessary registries.
// It loads tool definitions, initializes registries, and prepares for tool registration.
func NewToolManager() (*ToolManager, error) {
	// Load tool definitions from tools.json
	definitions, err := LoadToolDefinitions()
	if err != nil {
		return nil, fmt.Errorf("failed to load tool definitions: %w", err)
	}

	// Create handler and implementation registries
	handlers := NewToolHandlerRegistry()
	implementations := NewToolImplRegistry()

	manager := &ToolManager{
		definitions:     definitions,
		handlers:        handlers,
		implementations: implementations,
	}

	return manager, nil
}

// LoadAllTools initializes and registers all available tool implementations.
// This discovers tools automatically and registers their handlers.
func (tm *ToolManager) LoadAllTools() error {
	// Get all auto-registered tools from the tools package
	allTools := tools.GetAllTools()

	for _, regTool := range allTools {
		// Register the tool in the implementation registry (this also initializes it)
		if err := tm.implementations.Register(regTool.Name, regTool.Impl); err != nil {
			log.Printf("Failed to register tool %s: %v", regTool.Name, err)
			continue
		}

		// Register the tool's handlers
		if err := regTool.Impl.RegisterHandlers(tm.handlers); err != nil {
			log.Printf("Failed to register handlers for tool %s: %v", regTool.Name, err)
			continue
		}

		log.Printf("Successfully loaded tool: %s", regTool.Name)
	}

	return nil
}

// GetAllMCPTools returns all tool definitions converted to MCP Tool format.
// This is used for the tools/list MCP method response.
func (tm *ToolManager) GetAllMCPTools() []types.Tool {
	return tm.definitions.GetAllMCPTools()
}

// GetToolDefinition retrieves a tool definition by name.
// Returns nil if the tool is not found.
func (tm *ToolManager) GetToolDefinition(name string) *ToolDefinition {
	return tm.definitions.GetToolByName(name)
}

// GetToolHandler retrieves a tool handler function by name.
// Returns the handler and a boolean indicating whether it was found.
func (tm *ToolManager) GetToolHandler(name string) (func(context.Context, map[string]interface{}) (interface{}, error), bool) {
	return tm.handlers.GetHandler(name)
}

// GetToolImplementation retrieves a tool implementation by name.
// Returns the implementation and a boolean indicating whether it was found.
func (tm *ToolManager) GetToolImplementation(name string) (tools.ToolImplementation, bool) {
	return tm.implementations.Get(name)
}

// ValidateToolArguments checks that required parameters are present and valid
// for the specified tool using its definition schema.
func (tm *ToolManager) ValidateToolArguments(toolName string, args map[string]interface{}) error {
	toolDef := tm.GetToolDefinition(toolName)
	if toolDef == nil {
		return nil // No definition available, skip validation
	}

	for _, param := range toolDef.Params {
		if param.IsRequired {
			value, exists := args[param.Name]
			if !exists || value == nil {
				return fmt.Errorf("missing required parameter: %s - %s", param.Name, param.Description)
			}

			// Additional validation for string parameters
			if param.Type == "string" {
				strVal, ok := value.(string)
				if !ok {
					return fmt.Errorf("parameter %s must be a string", param.Name)
				}
				if strVal == "" {
					return fmt.Errorf("parameter %s cannot be empty", param.Name)
				}
			}
		}
	}
	return nil
}

// GetAvailableToolNames returns a formatted list of all available tool names.
// This is useful for error messages and debugging.
func (tm *ToolManager) GetAvailableToolNames() []string {
	tools := tm.GetAllMCPTools()
	names := make([]string, len(tools))
	for i, tool := range tools {
		names[i] = tool.Name
	}
	return names
}

// Shutdown gracefully shuts down all tool implementations.
// This should be called when the server is shutting down.
func (tm *ToolManager) Shutdown() {
	tm.implementations.ShutdownAll()
}

// GetStats returns statistics about the loaded tools for monitoring and debugging.
func (tm *ToolManager) GetStats() ToolManagerStats {
	allTools := tm.GetAllMCPTools()

	stats := ToolManagerStats{
		TotalTools:           len(allTools),
		ToolsWithUI:          0,
		DefinitionsCount:     len(tm.definitions.Tools),
		HandlersCount:        tm.getHandlerCount(),
		ImplementationsCount: tm.getImplementationCount(),
	}

	// Count UI-enabled tools
	for _, tool := range allTools {
		if tool.Meta != nil {
			// Type assert to map[string]interface{} before indexing
			if metaMap, ok := tool.Meta.(map[string]interface{}); ok {
				if _, hasUI := metaMap["ui"]; hasUI {
					stats.ToolsWithUI++
				}
			}
		}
	}

	return stats
}

// getHandlerCount returns the number of registered handlers
// This is a helper to avoid accessing private fields directly
func (tm *ToolManager) getHandlerCount() int {
	// We'll count by trying to get handlers for all known tools
	count := 0
	for _, tool := range tm.definitions.Tools {
		if _, exists := tm.handlers.GetHandler(tool.Name); exists {
			count++
		}
	}
	return count
}

// getImplementationCount returns the number of registered implementations
// This is a helper to avoid accessing private fields directly
func (tm *ToolManager) getImplementationCount() int {
	// We'll count by trying to get implementations for all known tools
	count := 0
	for _, tool := range tm.definitions.Tools {
		if _, exists := tm.implementations.Get(tool.Name); exists {
			count++
		}
	}
	return count
}

// ToolManagerStats provides statistics about the tool manager state.
type ToolManagerStats struct {
	TotalTools           int `json:"total_tools"`
	ToolsWithUI          int `json:"tools_with_ui"`
	DefinitionsCount     int `json:"definitions_count"`
	HandlersCount        int `json:"handlers_count"`
	ImplementationsCount int `json:"implementations_count"`
}
