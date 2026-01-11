package main

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"
)

// TestMCPServerInitialization verifies the MCP server initializes correctly
func TestMCPServerInitialization(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping MCP integration test in short mode")
	}

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create MCP server: %v", err)
	}

	if server == nil {
		t.Fatal("Server is nil")
	}

	// Verify server has required components
	if server.toolRegistry == nil {
		t.Error("Tool registry is nil")
	}
	if server.handlerRegistry == nil {
		t.Error("Handler registry is nil")
	}
	if server.implRegistry == nil {
		t.Error("Implementation registry is nil")
	}

	// Verify tools are loaded
	if len(server.tools) < 3 {
		t.Errorf("Expected at least 3 tools (api-health, dynamic-findings, static-findings), got %d", len(server.tools))
	}

	// Verify specific tools exist
	toolNames := make(map[string]bool)
	for _, tool := range server.tools {
		toolNames[tool.Name] = true
	}

	expectedTools := []string{
		"api-health",
		"get-dynamic-findings",
		"get-static-findings",
	}

	for _, expected := range expectedTools {
		if !toolNames[expected] {
			t.Errorf("Expected tool '%s' not found in server tools", expected)
		}
	}
}

// TestMCPProtocolInitialize tests the initialize handshake
func TestMCPProtocolInitialize(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping MCP integration test in short mode")
	}

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create MCP server: %v", err)
	}

	initParams := InitializeParams{
		ProtocolVersion: "2024-11-05",
		Capabilities: ClientCapabilities{
			Sampling: &SamplingCapability{},
		},
		ClientInfo: Implementation{
			Name:    "test-client",
			Version: "1.0.0",
		},
	}

	paramsJSON, _ := json.Marshal(initParams)

	req := &JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "initialize",
		Params:  paramsJSON,
	}

	resp := server.HandleRequest(req)

	if resp.Error != nil {
		t.Fatalf("Initialize returned error: %v", resp.Error)
	}

	// Parse the result
	resultJSON, _ := json.Marshal(resp.Result)
	var initResult InitializeResult
	if err := json.Unmarshal(resultJSON, &initResult); err != nil {
		t.Fatalf("Failed to parse initialize result: %v", err)
	}

	// Verify response
	if initResult.ProtocolVersion != "2024-11-05" {
		t.Errorf("Expected protocol version 2024-11-05, got %s", initResult.ProtocolVersion)
	}

	if initResult.ServerInfo.Name != "veracode-mcp-server" {
		t.Errorf("Expected server name 'veracode-mcp-server', got %s", initResult.ServerInfo.Name)
	}

	// Verify capabilities
	if initResult.Capabilities.Tools == nil {
		t.Error("Server should advertise tools capability")
	}
}

// TestMCPProtocolListTools tests the tools/list method
func TestMCPProtocolListTools(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping MCP integration test in short mode")
	}

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create MCP server: %v", err)
	}

	req := &JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      2,
		Method:  "tools/list",
	}

	resp := server.HandleRequest(req)

	if resp.Error != nil {
		t.Fatalf("tools/list returned error: %v", resp.Error)
	}

	// Parse the result
	resultJSON, _ := json.Marshal(resp.Result)
	var listResult ListToolsResult
	if err := json.Unmarshal(resultJSON, &listResult); err != nil {
		t.Fatalf("Failed to parse tools/list result: %v", err)
	}

	// Verify we have tools
	if len(listResult.Tools) < 3 {
		t.Errorf("Expected at least 3 tools, got %d", len(listResult.Tools))
	}

	// Verify tool structure
	for _, tool := range listResult.Tools {
		if tool.Name == "" {
			t.Error("Tool has empty name")
		}
		if tool.Description == "" {
			t.Errorf("Tool '%s' has empty description", tool.Name)
		}
		if tool.InputSchema == nil {
			t.Errorf("Tool '%s' has nil InputSchema", tool.Name)
		}
	}
}

// TestMCPToolCall_APIHealth tests calling the api-health tool through MCP protocol
func TestMCPToolCall_APIHealth(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping MCP integration test in short mode")
	}

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create MCP server: %v", err)
	}

	// Call api-health tool (no parameters required)
	callParams := CallToolParams{
		Name:      "api-health",
		Arguments: map[string]interface{}{},
	}

	paramsJSON, _ := json.Marshal(callParams)

	req := &JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      3,
		Method:  "tools/call",
		Params:  paramsJSON,
	}

	resp := server.HandleRequest(req)

	if resp.Error != nil {
		t.Fatalf("api-health tool call returned error: %v", resp.Error)
	}

	// Parse the result
	resultJSON, _ := json.Marshal(resp.Result)
	var callResult CallToolResult
	if err := json.Unmarshal(resultJSON, &callResult); err != nil {
		t.Fatalf("Failed to parse tool call result: %v", err)
	}

	// Verify response
	if callResult.IsError {
		t.Errorf("Tool call returned error: %s", callResult.Content[0].Text)
	}

	if len(callResult.Content) == 0 {
		t.Fatal("No content returned from api-health")
	}

	// Log the response for debugging
	t.Logf("API Health Response: %s", callResult.Content[0].Text)

	// Verify the response contains health information
	if callResult.Content[0].Type != "text" {
		t.Errorf("Expected content type 'text', got '%s'", callResult.Content[0].Text)
	}

	// Response should contain key health check information
	responseText := callResult.Content[0].Text
	if responseText == "" {
		t.Error("Response text is empty")
	}

	// Just verify we got some response - the format may vary
	t.Logf("Health check returned %d characters", len(responseText))
}

// TestMCPToolCall_DynamicFindings tests calling get-dynamic-findings through MCP protocol
func TestMCPToolCall_DynamicFindings(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping MCP integration test in short mode")
	}

	// Verify credentials are set
	if os.Getenv("VERACODE_API_ID") == "" || os.Getenv("VERACODE_API_KEY") == "" {
		t.Skip("VERACODE_API_ID and VERACODE_API_KEY must be set for integration tests")
	}

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create MCP server: %v", err)
	}

	// Call get-dynamic-findings tool with test application
	callParams := CallToolParams{
		Name: "get-dynamic-findings",
		Arguments: map[string]interface{}{
			"application_guid": "65c204e5-a74c-4b68-a62a-4bfdc08e27af", // MCPVerademo-NET
			"size":             5,
		},
	}

	paramsJSON, _ := json.Marshal(callParams)

	req := &JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      4,
		Method:  "tools/call",
		Params:  paramsJSON,
	}

	// Add timeout for API call
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Note: HandleRequest doesn't take context, but the tool implementation should use it
	resp := server.HandleRequest(req)

	if ctx.Err() != nil {
		t.Fatalf("Request timed out: %v", ctx.Err())
	}

	if resp.Error != nil {
		t.Fatalf("get-dynamic-findings tool call returned error: %v", resp.Error)
	}

	// Parse the result
	resultJSON, _ := json.Marshal(resp.Result)
	var callResult CallToolResult
	if err := json.Unmarshal(resultJSON, &callResult); err != nil {
		t.Fatalf("Failed to parse tool call result: %v", err)
	}

	// Verify response
	if callResult.IsError {
		t.Errorf("Tool call returned error: %s", callResult.Content[0].Text)
	}

	if len(callResult.Content) == 0 {
		t.Fatal("No content returned from get-dynamic-findings")
	}

	t.Logf("Dynamic Findings Response (first 500 chars): %s", truncate(callResult.Content[0].Text, 500))

	// Response should be JSON with findings
	var findingsData map[string]interface{}
	if err := json.Unmarshal([]byte(callResult.Content[0].Text), &findingsData); err != nil {
		t.Fatalf("Failed to parse findings response as JSON: %v", err)
	}

	// Verify response structure
	if _, hasPage := findingsData["page"]; !hasPage {
		t.Error("Response should contain 'page' field")
	}

	t.Logf("Findings data keys: %v", getKeys(findingsData))
}

// TestMCPToolCall_StaticFindings tests calling get-static-findings through MCP protocol
func TestMCPToolCall_StaticFindings(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping MCP integration test in short mode")
	}

	// Verify credentials are set
	if os.Getenv("VERACODE_API_ID") == "" || os.Getenv("VERACODE_API_KEY") == "" {
		t.Skip("VERACODE_API_ID and VERACODE_API_KEY must be set for integration tests")
	}

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create MCP server: %v", err)
	}

	// Call get-static-findings tool with test application
	callParams := CallToolParams{
		Name: "get-static-findings",
		Arguments: map[string]interface{}{
			"application_guid": "65c204e5-a74c-4b68-a62a-4bfdc08e27af", // MCPVerademo-NET
			"size":             5,
			"severity":         []interface{}{"High", "Very High"},
		},
	}

	paramsJSON, _ := json.Marshal(callParams)

	req := &JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      5,
		Method:  "tools/call",
		Params:  paramsJSON,
	}

	resp := server.HandleRequest(req)

	if resp.Error != nil {
		t.Fatalf("get-static-findings tool call returned error: %v", resp.Error)
	}

	// Parse the result
	resultJSON, _ := json.Marshal(resp.Result)
	var callResult CallToolResult
	if err := json.Unmarshal(resultJSON, &callResult); err != nil {
		t.Fatalf("Failed to parse tool call result: %v", err)
	}

	// Verify response
	if callResult.IsError {
		t.Errorf("Tool call returned error: %s", callResult.Content[0].Text)
	}

	if len(callResult.Content) == 0 {
		t.Fatal("No content returned from get-static-findings")
	}

	t.Logf("Static Findings Response (first 500 chars): %s", truncate(callResult.Content[0].Text, 500))

	// Response should be JSON with findings
	var findingsData map[string]interface{}
	if err := json.Unmarshal([]byte(callResult.Content[0].Text), &findingsData); err != nil {
		t.Fatalf("Failed to parse findings response as JSON: %v", err)
	}

	// Verify response structure
	if _, hasPage := findingsData["page"]; !hasPage {
		t.Error("Response should contain 'page' field")
	}

	if embedded, ok := findingsData["_embedded"].(map[string]interface{}); ok {
		if findings, ok := embedded["findings"].([]interface{}); ok {
			t.Logf("Retrieved %d static findings", len(findings))

			// Verify findings have expected structure
			if len(findings) > 0 {
				finding := findings[0].(map[string]interface{})

				// Check for key fields
				if _, hasDetails := finding["finding_details"]; !hasDetails {
					t.Error("Finding should have 'finding_details' field")
				}

				if details, ok := finding["finding_details"].(map[string]interface{}); ok {
					// Verify severity is an integer (from our spec fixes)
					if severity, ok := details["severity"]; ok {
						t.Logf("Severity type: %T, value: %v", severity, severity)
						// Severity should be a number (float64 in JSON unmarshaling)
						if _, isNumber := severity.(float64); !isNumber {
							t.Errorf("Severity should be a number, got %T", severity)
						}
					}

					// Verify CWE is an object (from our spec fixes)
					if cwe, ok := details["cwe"]; ok {
						t.Logf("CWE type: %T, value: %v", cwe, cwe)
						if cweObj, ok := cwe.(map[string]interface{}); ok {
							if _, hasID := cweObj["id"]; !hasID {
								t.Error("CWE object should have 'id' field")
							}
							if _, hasName := cweObj["name"]; !hasName {
								t.Error("CWE object should have 'name' field")
							}
						} else {
							t.Errorf("CWE should be an object, got %T", cwe)
						}
					}
				}
			}
		}
	}

	t.Logf("Findings data keys: %v", getKeys(findingsData))
}

// TestMCPToolCall_MissingRequiredParam tests parameter validation
func TestMCPToolCall_MissingRequiredParam(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping MCP integration test in short mode")
	}

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create MCP server: %v", err)
	}

	// Call get-dynamic-findings without required application_guid parameter
	callParams := CallToolParams{
		Name: "get-dynamic-findings",
		Arguments: map[string]interface{}{
			"size": 5,
		},
	}

	paramsJSON, _ := json.Marshal(callParams)

	req := &JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      6,
		Method:  "tools/call",
		Params:  paramsJSON,
	}

	resp := server.HandleRequest(req)

	if resp.Error != nil {
		t.Fatalf("Expected validation error in result, got RPC error: %v", resp.Error)
	}

	// Parse the result
	resultJSON, _ := json.Marshal(resp.Result)
	var callResult CallToolResult
	if err := json.Unmarshal(resultJSON, &callResult); err != nil {
		t.Fatalf("Failed to parse tool call result: %v", err)
	}

	// Should return an error result
	if !callResult.IsError {
		t.Error("Expected IsError=true for missing required parameter")
	}

	// Error message should mention the missing parameter
	if len(callResult.Content) > 0 {
		t.Logf("Validation error message: %s", callResult.Content[0].Text)
	} else {
		t.Error("Expected error content")
	}
}

// TestMCPToolCall_UnknownTool tests calling a non-existent tool
func TestMCPToolCall_UnknownTool(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping MCP integration test in short mode")
	}

	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create MCP server: %v", err)
	}

	callParams := CallToolParams{
		Name:      "non-existent-tool",
		Arguments: map[string]interface{}{},
	}

	paramsJSON, _ := json.Marshal(callParams)

	req := &JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      7,
		Method:  "tools/call",
		Params:  paramsJSON,
	}

	resp := server.HandleRequest(req)

	if resp.Error != nil {
		t.Fatalf("Expected error in result, got RPC error: %v", resp.Error)
	}

	// Parse the result
	resultJSON, _ := json.Marshal(resp.Result)
	var callResult CallToolResult
	if err := json.Unmarshal(resultJSON, &callResult); err != nil {
		t.Fatalf("Failed to parse tool call result: %v", err)
	}

	// Should return an error result
	if !callResult.IsError {
		t.Error("Expected IsError=true for unknown tool")
	}

	t.Logf("Unknown tool error message: %s", callResult.Content[0].Text)
}

// Helper functions

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
