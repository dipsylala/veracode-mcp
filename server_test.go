package main

import (
	"encoding/json"
	"testing"
)

func TestLoadToolDefinitions(t *testing.T) {
	registry, err := LoadToolDefinitions("tools.json")
	if err != nil {
		t.Fatalf("Failed to load tools.json: %v", err)
	}

	if len(registry.Tools) != 2 {
		t.Errorf("Expected 2 tools, got %d", len(registry.Tools))
	}

	// Check dynamic findings tool
	dynamicTool := registry.GetToolByName("get-dynamic-findings")
	if dynamicTool == nil {
		t.Fatal("get-dynamic-findings tool not found")
	}

	if dynamicTool.Category != "findings" {
		t.Errorf("Expected category 'findings', got '%s'", dynamicTool.Category)
	}

	if len(dynamicTool.Params) != 10 {
		t.Errorf("Expected 10 params for get-dynamic-findings, got %d", len(dynamicTool.Params))
	}

	// Check that application_path is first and required
	firstParam := dynamicTool.Params[0]
	if firstParam.Name != "application_path" {
		t.Errorf("Expected first param to be 'application_path', got '%s'", firstParam.Name)
	}
	if !firstParam.IsRequired {
		t.Error("application_path should be required")
	}

	// Check static findings tool
	staticTool := registry.GetToolByName("get-static-findings")
	if staticTool == nil {
		t.Fatal("get-static-findings tool not found")
	}
}

func TestToMCPTool(t *testing.T) {
	registry, err := LoadToolDefinitions("tools.json")
	if err != nil {
		t.Fatalf("Failed to load tools.json: %v", err)
	}

	dynamicTool := registry.GetToolByName("get-dynamic-findings")
	if dynamicTool == nil {
		t.Fatal("get-dynamic-findings tool not found")
	}

	mcpTool := dynamicTool.ToMCPTool()

	if mcpTool.Name != "get-dynamic-findings" {
		t.Errorf("Expected name 'get-dynamic-findings', got '%s'", mcpTool.Name)
	}

	// Check that input schema was generated
	if mcpTool.InputSchema == nil {
		t.Fatal("InputSchema is nil")
	}

	properties, ok := mcpTool.InputSchema["properties"].(map[string]interface{})
	if !ok {
		t.Fatal("properties is not a map")
	}

	// Check app_profile parameter
	appProfile, ok := properties["app_profile"].(map[string]interface{})
	if !ok {
		t.Fatal("app_profile property not found")
	}

	if appProfile["type"] != "string" {
		t.Errorf("Expected app_profile type 'string', got '%v'", appProfile["type"])
	}

	// Check severity parameter (array with enum)
	severity, ok := properties["severity"].(map[string]interface{})
	if !ok {
		t.Fatal("severity property not found")
	}

	if severity["type"] != "array" {
		t.Errorf("Expected severity type 'array', got '%v'", severity["type"])
	}

	items, ok := severity["items"].(map[string]interface{})
	if !ok {
		t.Fatal("severity items not found")
	}

	enum, ok := items["enum"].([]string)
	if !ok {
		t.Fatal("severity enum not found")
	}

	if len(enum) != 6 {
		t.Errorf("Expected 6 severity levels, got %d", len(enum))
	}
}

func TestServerInitialization(t *testing.T) {
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	if server == nil {
		t.Fatal("Server is nil")
	}

	if server.toolRegistry == nil {
		t.Fatal("Tool registry is nil")
	}

	if server.handlerRegistry == nil {
		t.Fatal("Handler registry is nil")
	}

	// Should have at least the JSON-defined tools plus echo
	if len(server.tools) < 3 {
		t.Errorf("Expected at least 3 tools, got %d", len(server.tools))
	}

	// Verify tools/list works
	result := server.handleListTools()
	if result == nil {
		t.Fatal("handleListTools returned nil")
	}

	if len(result.Tools) < 3 {
		t.Errorf("Expected at least 3 tools in list, got %d", len(result.Tools))
	}

	// Check that our JSON-defined tools are present
	foundDynamic := false
	foundStatic := false

	for _, tool := range result.Tools {
		if tool.Name == "get-dynamic-findings" {
			foundDynamic = true
		}
		if tool.Name == "get-static-findings" {
			foundStatic = true
		}
	}

	if !foundDynamic {
		t.Error("get-dynamic-findings not found in tools list")
	}

	if !foundStatic {
		t.Error("get-static-findings not found in tools list")
	}
}

func TestToolCallHandling(t *testing.T) {
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Test dynamic findings call with required parameter
	params := CallToolParams{
		Name: "get-dynamic-findings",
		Arguments: map[string]interface{}{
			"application_path": "/home/user/myapp",
			"app_profile":      "test-app",
			"severity":         []interface{}{"High", "Critical"},
		},
	}

	paramsJSON, _ := json.Marshal(params)
	result, err := server.handleCallTool(paramsJSON)

	if err != nil {
		t.Fatalf("handleCallTool failed: %v", err)
	}

	if result.IsError {
		t.Errorf("Tool call returned error: %s", result.Content[0].Text)
	}

	if len(result.Content) == 0 {
		t.Error("No content returned from tool call")
	}

	// Test missing required parameter
	paramsNoPath := CallToolParams{
		Name: "get-dynamic-findings",
		Arguments: map[string]interface{}{
			"app_profile": "test-app",
		},
	}

	paramsNoPathJSON, _ := json.Marshal(paramsNoPath)
	resultNoPath, err := server.handleCallTool(paramsNoPathJSON)

	if err != nil {
		t.Fatalf("handleCallTool failed: %v", err)
	}

	if !resultNoPath.IsError {
		t.Error("Expected error for missing application_path, but got success")
	}
}

// TODO: Add tests for new auto-registration tool system
// Tests should verify:
// - Tools are auto-discovered from init() functions
// - Tool schemas are correctly generated
// - Tool handlers are properly invoked
// - Tool responses follow MCP format
