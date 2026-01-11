package tools

import (
	"testing"
)

// TestActualToolsRegistration verifies that the actual tool implementations
// (dynamic_findings.go, static_findings.go) register themselves correctly
func TestActualToolsRegistration(t *testing.T) {
	// Get all registered tools (should include actual tools via their init() functions)
	tools := GetAllTools()

	if len(tools) == 0 {
		t.Fatal("No tools registered - init() functions may not be running")
	}

	// Build a map of registered tool names
	registeredTools := make(map[string]ToolImplementation)
	for _, tool := range tools {
		registeredTools[tool.Name()] = tool
	}

	// Verify expected tools are registered
	expectedTools := []string{
		"api-health",
		"dynamic-findings",
		"static-findings",
	}

	for _, expected := range expectedTools {
		if tool, exists := registeredTools[expected]; !exists {
			t.Errorf("Expected tool '%s' is not registered", expected)
		} else {
			// Verify the tool has valid name and description
			if tool.Name() == "" {
				t.Errorf("Tool '%s' has empty name", expected)
			}
			if tool.Description() == "" {
				t.Errorf("Tool '%s' has empty description", expected)
			}

			// Verify tool can be initialized
			if err := tool.Initialize(); err != nil {
				t.Errorf("Tool '%s' failed to initialize: %v", expected, err)
			}

			// Verify tool can register handlers
			mockRegistry := newMockHandlerRegistry()
			if err := tool.RegisterHandlers(mockRegistry); err != nil {
				t.Errorf("Tool '%s' failed to register handlers: %v", expected, err)
			}

			// Verify at least one handler was registered
			if len(mockRegistry.handlers) == 0 {
				t.Errorf("Tool '%s' did not register any handlers", expected)
			}

			// Verify tool can shutdown
			if err := tool.Shutdown(); err != nil {
				t.Errorf("Tool '%s' failed to shutdown: %v", expected, err)
			}
		}
	}
}

// TestToolNamesUnique verifies that all registered tools have unique names
func TestToolNamesUnique(t *testing.T) {
	tools := GetAllTools()

	names := make(map[string]bool)
	for _, tool := range tools {
		name := tool.Name()
		if names[name] {
			t.Errorf("Duplicate tool name found: %s", name)
		}
		names[name] = true
	}
}

// TestToolsImplementInterface verifies all tools implement the interface correctly
func TestToolsImplementInterface(t *testing.T) {
	tools := GetAllTools()

	for _, tool := range tools {
		// Verify the tool implements all required methods
		var _ ToolImplementation = tool

		// Test each method returns valid values
		if tool.Name() == "" {
			t.Errorf("Tool has empty name")
		}

		if tool.Description() == "" {
			t.Errorf("Tool '%s' has empty description", tool.Name())
		}

		// Initialize should not error
		if err := tool.Initialize(); err != nil {
			t.Errorf("Tool '%s' Initialize() returned error: %v", tool.Name(), err)
		}

		// RegisterHandlers should work with valid registry
		mockReg := newMockHandlerRegistry()
		if err := tool.RegisterHandlers(mockReg); err != nil {
			t.Errorf("Tool '%s' RegisterHandlers() returned error: %v", tool.Name(), err)
		}

		// Shutdown should not error
		if err := tool.Shutdown(); err != nil {
			t.Errorf("Tool '%s' Shutdown() returned error: %v", tool.Name(), err)
		}
	}
}
