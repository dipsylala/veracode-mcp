package mcp_tools

import (
	"context"
	"testing"
)

// ============================================================================
// parseFindingDetailsRequest tests
// ============================================================================

func TestParseFindingDetailsRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          "12345",
	}

	req, err := parseFindingDetailsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
	if req.FlawID != "12345" {
		t.Errorf("Expected flaw_id '12345', got '%s'", req.FlawID)
	}
}

func TestParseFindingDetailsRequest_WithAppProfile(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"app_profile":      "MyApp",
		"flaw_id":          "999",
	}

	req, err := parseFindingDetailsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.AppProfile != "MyApp" {
		t.Errorf("Expected app_profile 'MyApp', got '%s'", req.AppProfile)
	}
	if req.FlawID != "999" {
		t.Errorf("Expected flaw_id '999', got '%s'", req.FlawID)
	}
}

func TestParseFindingDetailsRequest_MissingFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	_, err := parseFindingDetailsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing flaw_id")
	}
}

func TestParseFindingDetailsRequest_ZeroFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          "",
	}

	_, err := parseFindingDetailsRequest(args)
	if err == nil {
		t.Fatal("Expected error for empty flaw_id")
	}
}

func TestParseFindingDetailsRequest_NegativeFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          "-1",
	}

	req, err := parseFindingDetailsRequest(args)
	// Parsing succeeds for string "-1", validation happens later in handler
	if err != nil {
		t.Fatalf("Expected no error during parsing, got: %v", err)
	}
	if req.FlawID != "-1" {
		t.Errorf("Expected flaw_id '-1', got '%s'", req.FlawID)
	}
}

func TestParseFindingDetailsRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{
		"flaw_id": "12345",
	}

	_, err := parseFindingDetailsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

// ============================================================================
// handleGetFindingDetails tests
// ============================================================================

func TestFindingDetailsTool_Initialize(t *testing.T) {
	tool := NewFindingDetailsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestFindingDetailsTool_RegisterHandlers(t *testing.T) {
	tool := NewFindingDetailsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[FindingDetailsToolName] == nil {
		t.Error("Handler was not registered")
	}
}

func TestFindingDetailsTool_HandleMissingFlawID(t *testing.T) {
	ctx := context.Background()

	result, err := handleGetFindingDetails(ctx, map[string]interface{}{
		"application_path": "/path/to/app",
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["error"] == nil {
		t.Error("Expected error for missing flaw_id")
	}
}

func TestFindingDetailsTool_HandleInvalidFlawID(t *testing.T) {
	ctx := context.Background()

	result, err := handleGetFindingDetails(ctx, map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          "",
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["error"] == nil {
		t.Error("Expected error for empty flaw_id")
	}
}
