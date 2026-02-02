package mcp_tools

import (
	"context"
	"testing"
)

// TestSimpleTool tests the SimpleTool wrapper for stateless tools
func TestSimpleTool_Initialize(t *testing.T) {
	handler := func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
		return map[string]string{"status": "ok"}, nil
	}

	tool := NewSimpleTool("test-tool", handler)

	// Should succeed (no-op)
	if err := tool.Initialize(); err != nil {
		t.Errorf("Initialize() should succeed, got error: %v", err)
	}
}

func TestSimpleTool_RegisterHandlers(t *testing.T) {
	handler := func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
		return map[string]string{"status": "ok"}, nil
	}

	tool := NewSimpleTool("test-tool", handler)
	registry := newMockHandlerRegistry()

	// Should register the handler
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Errorf("RegisterHandlers() should succeed, got error: %v", err)
	}

	// Verify handler was registered
	if registry.handlers["test-tool"] == nil {
		t.Error("Handler was not registered")
	}

	// Verify registered handler works
	result, err := registry.handlers["test-tool"](context.Background(), nil)
	if err != nil {
		t.Errorf("Handler execution failed: %v", err)
	}

	resultMap, ok := result.(map[string]string)
	if !ok || resultMap["status"] != "ok" {
		t.Errorf("Handler returned unexpected result: %v", result)
	}
}

func TestSimpleTool_Shutdown(t *testing.T) {
	handler := func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
		return nil, nil
	}

	tool := NewSimpleTool("test-tool", handler)

	// Should succeed (no-op)
	if err := tool.Shutdown(); err != nil {
		t.Errorf("Shutdown() should succeed, got error: %v", err)
	}
}

func TestRegisterMCPTool(t *testing.T) {
	handler := func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
		return "test-result", nil
	}

	// Register the tool
	RegisterMCPTool("simple-test-tool", handler)

	// Verify it was registered in the global registry
	allTools := GetAllTools()
	found := false
	for _, regTool := range allTools {
		if regTool.Name == "simple-test-tool" {
			found = true

			// Verify the handler works
			registry := newMockHandlerRegistry()
			if err := regTool.Impl.RegisterHandlers(registry); err != nil {
				t.Fatalf("Failed to register handlers: %v", err)
			}

			result, err := registry.handlers["simple-test-tool"](context.Background(), nil)
			if err != nil {
				t.Fatalf("Handler execution failed: %v", err)
			}

			if result != "test-result" {
				t.Errorf("Expected 'test-result', got: %v", result)
			}
			break
		}
	}

	if !found {
		t.Error("Tool was not registered in global registry")
	}
}

// Mock handler registry for testing
type mockHandlerRegistry struct {
	handlers map[string]ToolHandler
}

func newMockHandlerRegistry() *mockHandlerRegistry {
	return &mockHandlerRegistry{
		handlers: make(map[string]ToolHandler),
	}
}

func (r *mockHandlerRegistry) RegisterHandler(toolName string, handler ToolHandler) {
	r.handlers[toolName] = handler
}
