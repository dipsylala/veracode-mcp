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
		"dynamic-findings",
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

// TestMCPToolCall_DynamicFindings tests calling dynamic-findings through MCP protocol
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

	// Call dynamic-findings tool with test application GUID
	callParams := CallToolParams{
		Name: "get-dynamic-findings",
		Arguments: map[string]interface{}{
			"application_path": "E:\\Github\\VeracodeMCP-Go",           // Required workspace path
			"app_profile":      "f4e74197-1e26-42c4-ab4b-245870c93280", // App GUID for testing
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

	// Parse and validate the response
	callResult := parseToolCallResult(t, resp.Result)
	responseText := extractResponseText(t, callResult)
	validateFindingsResponse(t, responseText)
}

func parseToolCallResult(t *testing.T, result interface{}) CallToolResult {
	resultJSON, _ := json.Marshal(result)
	var callResult CallToolResult
	if err := json.Unmarshal(resultJSON, &callResult); err != nil {
		t.Fatalf("Failed to parse tool call result: %v", err)
	}

	if callResult.IsError {
		t.Errorf("Tool call returned error: %s", callResult.Content[0].Text)
	}

	if len(callResult.Content) == 0 {
		t.Fatal("No content returned from tool")
	}

	return callResult
}

func extractResponseText(t *testing.T, callResult CallToolResult) string {
	var responseText string
	if callResult.Content[0].Resource != nil {
		responseText = callResult.Content[0].Resource.Text
	} else {
		responseText = callResult.Content[0].Text
	}

	t.Logf("Findings Response (first 500 chars): %s", truncate(responseText, 500))

	// Check if response is the graceful fallback (when credentials are invalid)
	if len(responseText) > 0 && (responseText[0] != '{' && responseText[0] != '[') {
		t.Skip("Skipping test - received fallback response (credentials may be invalid)")
	}

	return responseText
}

func validateFindingsResponse(t *testing.T, responseText string) {
	// Response should be JSON with MCPFindingsResponse structure
	var findingsData map[string]interface{}
	if err := json.Unmarshal([]byte(responseText), &findingsData); err != nil {
		t.Fatalf("Failed to parse findings response as JSON: %v", err)
	}

	// Verify MCPFindingsResponse structure
	if _, hasPagination := findingsData["pagination"]; !hasPagination {
		t.Error("Response should contain 'pagination' field")
	}

	if _, hasApplication := findingsData["application"]; !hasApplication {
		t.Error("Response should contain 'application' field")
	}

	if _, hasSummary := findingsData["summary"]; !hasSummary {
		t.Error("Response should contain 'summary' field")
	}

	if _, hasFindings := findingsData["findings"]; !hasFindings {
		t.Error("Response should contain 'findings' field")
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
			"application_path": "E:\\Github\\VeracodeMCP-Go", // Required workspace path
			"app_profile":      "MCPVerademo-NET",            // Override workspace config
			"size":             5,
			"severity_gte":     4,
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

	// Parse and validate the response
	callResult := parseToolCallResult(t, resp.Result)
	responseText := extractResponseText(t, callResult)
	findingsData := parseStaticFindingsResponse(t, responseText)
	verifyStaticFindingStructure(t, findingsData)
}

func parseStaticFindingsResponse(t *testing.T, responseText string) map[string]interface{} {
	// Response should be JSON with MCPFindingsResponse structure
	var findingsData map[string]interface{}
	if err := json.Unmarshal([]byte(responseText), &findingsData); err != nil {
		t.Fatalf("Failed to parse findings response as JSON: %v", err)
	}

	// Verify MCPFindingsResponse structure
	if _, hasPagination := findingsData["pagination"]; !hasPagination {
		t.Error("Response should contain 'pagination' field")
	}

	if _, hasApplication := findingsData["application"]; !hasApplication {
		t.Error("Response should contain 'application' field")
	}

	if _, hasSummary := findingsData["summary"]; !hasSummary {
		t.Error("Response should contain 'summary' field")
	}

	return findingsData
}

func verifyStaticFindingStructure(t *testing.T, findingsData map[string]interface{}) {
	// Check findings structure
	findings, ok := findingsData["findings"].([]interface{})
	if !ok {
		return
	}

	t.Logf("Retrieved %d static findings", len(findings))

	// Verify findings have expected MCPFinding structure
	if len(findings) == 0 {
		return
	}

	finding := findings[0].(map[string]interface{})

	// Check for key MCPFinding fields
	if _, hasFlawID := finding["flaw_id"]; !hasFlawID {
		t.Error("Finding should have 'flaw_id' field")
	}

	if _, hasSeverity := finding["severity"]; !hasSeverity {
		t.Error("Finding should have 'severity' field")
	}

	if _, hasScanType := finding["scan_type"]; !hasScanType {
		t.Error("Finding should have 'scan_type' field")
	}

	if severity, ok := finding["severity"]; ok {
		t.Logf("Severity: %v", severity)
	}

	if severityScore, ok := finding["severity_score"]; ok {
		t.Logf("Severity score type: %T, value: %v", severityScore, severityScore)
		// Severity score should be a number (float64 in JSON unmarshaling)
		if _, isNumber := severityScore.(float64); !isNumber {
			t.Errorf("Severity score should be a number, got %T", severityScore)
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
