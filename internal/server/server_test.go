package server

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	tools "github.com/dipsylala/veracodemcp-go/internal/tool_registry"
	"github.com/dipsylala/veracodemcp-go/internal/types"
)

// TestMain sets up test fixtures before running tests
func TestMain(m *testing.M) {
	// Load tools.json for tests
	// Navigate up from internal/server to project root
	toolsJSONPath := filepath.Join("..", "..", "tools.json")
	// #nosec G304 -- toolsJSONPath is a fixed test fixture path, not user input
	toolsJSONData, err := os.ReadFile(toolsJSONPath)
	if err != nil {
		panic("Failed to read tools.json for tests: " + err.Error())
	}

	// Set the tools JSON data
	tools.SetToolsJSON(toolsJSONData)

	// Run tests
	os.Exit(m.Run())
}

func TestLoadToolDefinitions(t *testing.T) {
	// Create a server which initializes the ToolManager internally
	server, err := NewMCPServer(false)
	if err != nil {
		t.Fatalf("Failed to create server with tool manager: %v", err)
	}

	// Get all MCP tools from the tool manager
	mcpTools := server.toolManager.GetAllMCPTools()

	// tools.json now has 12 tools: api-health, dynamic-findings, static-findings, get-finding-details, get-sca-findings, package-workspace, pipeline-scan, pipeline-status, pipeline-results, pipeline-detailed-results, run-sca-scan, get-local-sca-results
	if len(mcpTools) != 12 {
		t.Errorf("Expected 12 tools, got %d", len(mcpTools))
	}

	// Check dynamic findings tool
	dynamicTool := server.toolManager.GetToolDefinition("dynamic-findings")
	if dynamicTool == nil {
		t.Fatal("dynamic-findings tool not found")
	}

	// Note: Category is optional and may not be set
	if dynamicTool.Category != "" && dynamicTool.Category != "findings" {
		t.Errorf("Expected category 'findings' or empty, got '%s'", dynamicTool.Category)
	}

	if len(dynamicTool.Params) != 9 {
		t.Errorf("Expected 9 params for dynamic-findings, got %d", len(dynamicTool.Params))
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
	staticTool := server.toolManager.GetToolDefinition("static-findings")
	if staticTool == nil {
		t.Fatal("static-findings tool not found")
	}
}

func TestToMCPTool(t *testing.T) {
	// Create a server which initializes the ToolManager internally
	server, err := NewMCPServer(false)
	if err != nil {
		t.Fatalf("Failed to create server with tool manager: %v", err)
	}

	dynamicTool := server.toolManager.GetToolDefinition("dynamic-findings")
	if dynamicTool == nil {
		t.Fatal("dynamic-findings tool not found")
	}

	mcpTool := dynamicTool.ToMCPTool()

	if mcpTool.Name != "dynamic-findings" {
		t.Errorf("Expected name 'dynamic-findings', got '%s'", mcpTool.Name)
	}

	// Check that input schema was generated
	if mcpTool.InputSchema == nil {
		t.Fatal("InputSchema is nil")
	}

	schema, ok := mcpTool.InputSchema.(map[string]interface{})
	if !ok {
		t.Fatal("InputSchema is not a map[string]interface{}")
	}
	properties, ok := schema["properties"].(map[string]interface{})
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

	// Check severity parameter (number with validation)
	severity, ok := properties["severity"].(map[string]interface{})
	if !ok {
		t.Fatal("severity property not found")
	}

	if severity["type"] != "number" {
		t.Errorf("Expected severity type 'number', got '%v'", severity["type"])
	}

	minimum, hasMin := severity["minimum"]
	maximum, hasMax := severity["maximum"]
	if !hasMin || !hasMax {
		t.Fatal("severity minimum/maximum validation not found")
	}

	if minimum != float64(0) || maximum != float64(5) {
		t.Errorf("Expected severity min=0 max=5, got min=%v max=%v", minimum, maximum)
	}
}

func TestServerInitialization(t *testing.T) {
	server := createTestServer(t)
	validateServerStructure(t, server)
	validateToolsLoaded(t, server)
	validateToolsListWorks(t, server)
	validateSpecificToolsPresent(t, server)
	validateToolManagerStats(t, server)
}

// Helper functions for TestServerInitialization

func createTestServer(t *testing.T) *MCPServer {
	server, err := NewMCPServer(false)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	if server == nil {
		t.Fatal("Server is nil")
	}
	return server
}

func validateServerStructure(t *testing.T, server *MCPServer) {
	if server.toolManager == nil {
		t.Fatal("Tool manager is nil")
	}
}

func validateToolsLoaded(t *testing.T, server *MCPServer) {
	if len(server.tools) < 5 {
		t.Errorf("Expected at least 5 tools, got %d", len(server.tools))
	}

	// Verify tool manager has definitions loaded
	dynamicTool := server.toolManager.GetToolDefinition("dynamic-findings")
	if dynamicTool == nil {
		t.Error("dynamic-findings tool definition not found in tool manager")
	}

	// Verify handler lookup works
	_, exists := server.toolManager.GetToolHandler("dynamic-findings")
	if !exists {
		t.Error("dynamic-findings handler not found in tool manager")
	}
}

func validateToolsListWorks(t *testing.T, server *MCPServer) {
	result := server.handleListTools()
	if result == nil {
		t.Fatal("handleListTools returned nil")
	}

	if len(result.Tools) < 5 {
		t.Errorf("Expected at least 5 tools in list, got %d", len(result.Tools))
	}
}

func validateSpecificToolsPresent(t *testing.T, server *MCPServer) {
	result := server.handleListTools()
	foundDynamic := false
	foundStatic := false

	for _, tool := range result.Tools {
		switch tool.Name {
		case "dynamic-findings":
			foundDynamic = true
		case "static-findings":
			foundStatic = true
		}
	}

	if !foundDynamic {
		t.Error("dynamic-findings not found in tools list")
	}
	if !foundStatic {
		t.Error("static-findings not found in tools list")
	}
}

func validateToolManagerStats(t *testing.T, server *MCPServer) {
	stats := server.GetToolStats()
	if stats.TotalTools == 0 {
		t.Error("Tool manager reports zero total tools")
	}
	if stats.DefinitionsCount == 0 {
		t.Error("Tool manager reports zero definitions")
	}
}

func TestToolCallHandling(t *testing.T) {
	server, err := NewMCPServer(false)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Create a temporary directory with workspace config for testing
	tempDir := t.TempDir()
	workspaceFile := tempDir + "/.veracode-workspace.json"
	workspaceContent := `{"name": "test-app"}`
	if err := os.WriteFile(workspaceFile, []byte(workspaceContent), 0644); err != nil {
		t.Fatalf("Failed to create workspace file: %v", err)
	}

	// Test dynamic findings call with required parameter
	params := types.CallToolParams{
		Name: "dynamic-findings",
		Arguments: map[string]interface{}{
			"application_path": tempDir,
			"severity_gte":     4,
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
	paramsNoPath := types.CallToolParams{
		Name:      "dynamic-findings",
		Arguments: map[string]interface{}{},
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

func TestToolManagerCreation(t *testing.T) {
	manager, err := tools.NewToolManager()
	if err != nil {
		t.Fatalf("Failed to create tool manager: %v", err)
	}

	if manager == nil {
		t.Fatal("Tool manager is nil")
	}

	// Test that definitions were loaded
	mcpTools := manager.GetAllMCPTools()
	if len(mcpTools) < 5 {
		t.Errorf("Expected at least 5 tools, got %d", len(mcpTools))
	}
}

func TestToolManagerDefinitionLookup(t *testing.T) {
	manager, err := tools.NewToolManager()
	if err != nil {
		t.Fatalf("Failed to create tool manager: %v", err)
	}

	// Test existing tool
	dynamicTool := manager.GetToolDefinition("dynamic-findings")
	if dynamicTool == nil {
		t.Error("dynamic-findings definition not found")
	} else {
		if dynamicTool.Name != "dynamic-findings" {
			t.Errorf("Expected name 'dynamic-findings', got '%s'", dynamicTool.Name)
		}
	}

	// Test non-existent tool
	nonExistent := manager.GetToolDefinition("non-existent-tool")
	if nonExistent != nil {
		t.Error("Expected nil for non-existent tool, got definition")
	}
}

func TestToolManagerValidation(t *testing.T) {
	manager, err := tools.NewToolManager()
	if err != nil {
		t.Fatalf("Failed to create tool manager: %v", err)
	}

	// Test validation with valid arguments
	validArgs := map[string]interface{}{
		"application_path": "/tmp/test",
		"severity_gte":     4,
	}
	err = manager.ValidateToolArguments("dynamic-findings", validArgs)
	if err != nil {
		t.Errorf("Validation failed for valid args: %v", err)
	}

	// Test validation with missing required parameter
	invalidArgs := map[string]interface{}{
		"severity_gte": 4, // missing application_path
	}
	err = manager.ValidateToolArguments("dynamic-findings", invalidArgs)
	if err == nil {
		t.Error("Expected validation error for missing required parameter")
	}

	// Test validation for non-existent tool (should not error)
	err = manager.ValidateToolArguments("non-existent-tool", validArgs)
	if err != nil {
		t.Errorf("Validation should not error for non-existent tool: %v", err)
	}
}

func TestToolManagerAvailableToolNames(t *testing.T) {
	manager, err := tools.NewToolManager()
	if err != nil {
		t.Fatalf("Failed to create tool manager: %v", err)
	}

	names := manager.GetAvailableToolNames()
	if len(names) < 5 {
		t.Errorf("Expected at least 5 tool names, got %d", len(names))
	}

	// Check for specific known tools
	foundDynamic := false
	foundStatic := false
	for _, name := range names {
		switch name {
		case "dynamic-findings":
			foundDynamic = true
		case "static-findings":
			foundStatic = true
		}
	}

	if !foundDynamic {
		t.Error("dynamic-findings not found in available tool names")
	}
	if !foundStatic {
		t.Error("static-findings not found in available tool names")
	}
}

func TestToolManagerStatistics(t *testing.T) {
	manager, err := tools.NewToolManager()
	if err != nil {
		t.Fatalf("Failed to create tool manager: %v", err)
	}

	// Load tools to get handlers registered
	err = manager.LoadAllTools()
	if err != nil {
		t.Fatalf("Failed to load tools: %v", err)
	}

	stats := manager.GetStats()

	if stats.TotalTools == 0 {
		t.Error("Expected non-zero total tools")
	}

	if stats.DefinitionsCount == 0 {
		t.Error("Expected non-zero definitions count")
	}

	// After loading tools, we should have some handlers
	if stats.HandlersCount == 0 {
		t.Error("Expected non-zero handlers count after loading tools")
	}

	// Check that total tools matches definitions (they should be the same)
	if stats.TotalTools != stats.DefinitionsCount {
		t.Errorf("Total tools (%d) should match definitions count (%d)",
			stats.TotalTools, stats.DefinitionsCount)
	}
}

func TestServerRefactoredHandlers(t *testing.T) {
	server, err := NewMCPServer(false)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	t.Run("InitializeHandling", func(t *testing.T) {
		// Test the refactored initialize handling
		initParams := InitializeParams{
			ProtocolVersion: "2024-11-05",
			ClientInfo: Implementation{
				Name:    "test-client",
				Version: "1.0.0",
			},
			Capabilities: ClientCapabilities{},
		}

		paramsJSON, err := json.Marshal(initParams)
		if err != nil {
			t.Fatalf("Failed to marshal init params: %v", err)
		}

		result, err := server.handleInitialize(paramsJSON)
		if err != nil {
			t.Fatalf("handleInitialize failed: %v", err)
		}

		if result == nil {
			t.Fatal("Initialize result is nil")
		}

		if result.ProtocolVersion != "2024-11-05" {
			t.Errorf("Expected protocol version '2024-11-05', got '%s'", result.ProtocolVersion)
		}

		if result.ServerInfo.Name != "veracode-mcp-server" {
			t.Errorf("Expected server name 'veracode-mcp-server', got '%s'", result.ServerInfo.Name)
		}
	})

	t.Run("ToolCallValidation", func(t *testing.T) {
		// Test the new validation methods
		callParams := &types.CallToolParams{
			Name: "dynamic-findings",
			Arguments: map[string]interface{}{
				"application_path": "/tmp/test",
			},
		}

		err := server.validateToolCall(callParams)
		if err != nil {
			t.Errorf("Validation failed for valid call params: %v", err)
		}

		// Test with missing required parameter
		invalidParams := &types.CallToolParams{
			Name:      "dynamic-findings",
			Arguments: map[string]interface{}{}, // missing application_path
		}

		err = server.validateToolCall(invalidParams)
		if err == nil {
			t.Error("Expected validation error for missing required parameter")
		}
	})

	t.Run("ErrorHandling", func(t *testing.T) {
		// Test the new error handling methods
		errorResult := server.createToolError("Test error message")

		if errorResult == nil {
			t.Fatal("Error result is nil")
		}

		if !errorResult.IsError {
			t.Error("Expected IsError to be true")
		}

		if len(errorResult.Content) == 0 {
			t.Error("Expected error content")
		} else if errorResult.Content[0].Text != "Test error message" {
			t.Errorf("Expected error message 'Test error message', got '%s'", errorResult.Content[0].Text)
		}
	})
}
