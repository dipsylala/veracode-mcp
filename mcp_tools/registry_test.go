package mcp_tools

import (
	"context"
	"testing"
)

// Mock tool for testing
type mockTool struct {
	name        string
	description string
	initialized bool
	handlers    map[string]ToolHandler
}

func newMockTool(name, description string) *mockTool {
	return &mockTool{
		name:        name,
		description: description,
		handlers:    make(map[string]ToolHandler),
	}
}

func (m *mockTool) Initialize() error {
	m.initialized = true
	return nil
}

func (m *mockTool) RegisterHandlers(registry HandlerRegistry) error {
	// Register a simple test handler
	registry.RegisterHandler(m.name+"-action", func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
		return map[string]interface{}{"success": true}, nil
	})
	return nil
}

func (m *mockTool) Shutdown() error {
	m.initialized = false
	return nil
}

func TestRegisterTool(t *testing.T) {
	// Clear the global registry before test
	registryMu.Lock()
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu.Unlock()

	// Test registering a tool
	toolName := "test-tool"
	constructor := func() ToolImplementation {
		return newMockTool(toolName, "A test tool")
	}

	RegisterTool(toolName, constructor)

	// Verify tool was registered
	registryMu.RLock()
	_, exists := toolRegistry[toolName]
	registryMu.RUnlock()

	if !exists {
		t.Errorf("Tool %s was not registered", toolName)
	}

	// Verify we can retrieve it
	tools := GetAllTools()
	if len(tools) != 1 {
		t.Fatalf("Expected 1 tool, got %d", len(tools))
	}

	if tools[0].Name != toolName {
		t.Errorf("Expected tool name %s, got %s", toolName, tools[0].Name)
	}
}

func TestRegisterMultipleTools(t *testing.T) {
	// Clear the global registry
	registryMu.Lock()
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu.Unlock()

	// Register multiple tools
	tools := []struct {
		name        string
		description string
	}{
		{"tool-1", "First tool"},
		{"tool-2", "Second tool"},
		{"tool-3", "Third tool"},
	}

	for _, tool := range tools {
		name := tool.name
		desc := tool.description
		RegisterTool(name, func() ToolImplementation {
			return newMockTool(name, desc)
		})
	}

	// Verify all tools are registered
	allTools := GetAllTools()
	if len(allTools) != len(tools) {
		t.Fatalf("Expected %d tools, got %d", len(tools), len(allTools))
	}

	// Verify each tool exists
	toolNames := make(map[string]bool)
	for _, tool := range allTools {
		toolNames[tool.Name] = true
	}

	for _, expected := range tools {
		if !toolNames[expected.name] {
			t.Errorf("Tool %s was not found in registered tools", expected.name)
		}
	}
}

func TestDuplicateRegistration(t *testing.T) {
	// Clear the global registry
	registryMu.Lock()
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu.Unlock()

	toolName := "duplicate-tool"

	// Register first time
	RegisterTool(toolName, func() ToolImplementation {
		return newMockTool(toolName, "First registration")
	})

	// Register again with same name
	RegisterTool(toolName, func() ToolImplementation {
		return newMockTool(toolName, "Second registration")
	})

	// Should still only have one tool (last one wins)
	tools := GetAllTools()
	if len(tools) != 1 {
		t.Fatalf("Expected 1 tool after duplicate registration, got %d", len(tools))
	}

	// The second registration should have replaced the first
	mock := tools[0].Impl.(*mockTool)
	if mock.description != "Second registration" {
		t.Errorf("Expected description 'Second registration', got '%s'", mock.description)
	}
}

func TestGetAllToolsReturnsNewInstances(t *testing.T) {
	// Clear the global registry
	registryMu.Lock()
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu.Unlock()

	// Register a tool
	RegisterTool("instance-test", func() ToolImplementation {
		return newMockTool("instance-test", "Instance test")
	})

	// Get tools twice
	tools1 := GetAllTools()
	tools2 := GetAllTools()

	if len(tools1) != 1 || len(tools2) != 1 {
		t.Fatal("Expected 1 tool in each call")
	}

	// Modify first instance
	mock1 := tools1[0].Impl.(*mockTool)
	mock1.initialized = true

	// Second instance should be independent
	mock2 := tools2[0].Impl.(*mockTool)
	if mock2.initialized {
		t.Error("Second instance should not be affected by changes to first instance")
	}
}

func TestToolInitialization(t *testing.T) {
	// Clear the global registry
	registryMu.Lock()
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu.Unlock()

	// Register a tool
	RegisterTool("init-test", func() ToolImplementation {
		return newMockTool("init-test", "Initialization test")
	})

	// Get the tool
	tools := GetAllTools()
	if len(tools) != 1 {
		t.Fatal("Expected 1 tool")
	}

	mock := tools[0].Impl.(*mockTool)

	// Verify not initialized yet
	if mock.initialized {
		t.Error("Tool should not be initialized yet")
	}

	// Initialize the tool
	if err := mock.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}

	// Verify initialized
	if !mock.initialized {
		t.Error("Tool should be initialized")
	}

	// Shutdown
	if err := mock.Shutdown(); err != nil {
		t.Fatalf("Failed to shutdown tool: %v", err)
	}

	// Verify shutdown
	if mock.initialized {
		t.Error("Tool should be shutdown")
	}
}

func TestHandlerRegistration(t *testing.T) {
	// Clear the global registry
	registryMu.Lock()
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu.Unlock()

	// Register a tool
	RegisterTool("handler-test", func() ToolImplementation {
		return newMockTool("handler-test", "Handler test")
	})

	// Get the tool
	tools := GetAllTools()
	if len(tools) != 1 {
		t.Fatal("Expected 1 tool")
	}

	// Create mock handler registry
	handlerReg := newMockHandlerRegistry()

	// Register handlers
	if err := tools[0].Impl.RegisterHandlers(handlerReg); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	// Verify handler was registered
	expectedHandlerName := "handler-test-action"
	if _, exists := handlerReg.handlers[expectedHandlerName]; !exists {
		t.Errorf("Expected handler %s was not registered", expectedHandlerName)
	}

	// Test calling the handler
	handler := handlerReg.handlers[expectedHandlerName]
	result, err := handler(context.Background(), map[string]interface{}{})
	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Handler result is not a map")
	}

	if success, ok := resultMap["success"].(bool); !ok || !success {
		t.Error("Handler did not return success: true")
	}
}

func TestEmptyRegistry(t *testing.T) {
	// Clear the global registry
	registryMu.Lock()
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu.Unlock()

	// Get tools from empty registry
	tools := GetAllTools()

	if len(tools) != 0 {
		t.Errorf("Expected 0 tools from empty registry, got %d", len(tools))
	}
}

func TestConcurrentRegistration(t *testing.T) {
	// Clear the global registry
	registryMu.Lock()
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu.Unlock()

	// Register tools concurrently
	done := make(chan bool)
	numTools := 10

	for i := 0; i < numTools; i++ {
		go func(idx int) {
			toolName := "concurrent-tool-" + string(rune('A'+idx))
			RegisterTool(toolName, func() ToolImplementation {
				return newMockTool(toolName, "Concurrent test")
			})
			done <- true
		}(i)
	}

	// Wait for all registrations
	for i := 0; i < numTools; i++ {
		<-done
	}

	// Verify all tools were registered
	tools := GetAllTools()
	if len(tools) != numTools {
		t.Errorf("Expected %d tools after concurrent registration, got %d", numTools, len(tools))
	}
}

func TestConcurrentGetAllTools(t *testing.T) {
	// Clear and populate registry
	registryMu.Lock()
	toolRegistry = make(map[string]func() ToolImplementation)
	registryMu.Unlock()

	RegisterTool("concurrent-get-test", func() ToolImplementation {
		return newMockTool("concurrent-get-test", "Test")
	})

	// Call GetAllTools concurrently
	done := make(chan bool)
	numCalls := 10

	for i := 0; i < numCalls; i++ {
		go func() {
			tools := GetAllTools()
			if len(tools) != 1 {
				t.Errorf("Expected 1 tool, got %d", len(tools))
			}
			done <- true
		}()
	}

	// Wait for all calls
	for i := 0; i < numCalls; i++ {
		<-done
	}
}
