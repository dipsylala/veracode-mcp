package mcp_tools

import (
	"context"
	"testing"
)

// ============================================================================
// parseScaFindingsRequest tests
// ============================================================================

func TestParseScaFindingsRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parseScaFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
	// SCA default size is 10 (same as other findings)
	if req.Size != 10 {
		t.Errorf("Expected default size 10 for SCA, got %d", req.Size)
	}
}

func TestParseScaFindingsRequest_AllParameters(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"app_profile":      "MyApp",
		"page_size":        float64(100),
		"page":             float64(2),
		"severity":         float64(4),
		"severity_gte":     float64(2),
	}

	req, err := parseScaFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.AppProfile != "MyApp" {
		t.Errorf("Expected app_profile 'MyApp', got '%s'", req.AppProfile)
	}
	if req.Size != 100 {
		t.Errorf("Expected size 100, got %d", req.Size)
	}
	if req.Page != 2 {
		t.Errorf("Expected page 2, got %d", req.Page)
	}
}

func TestParseScaFindingsRequest_SeverityValidation(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"severity":         float64(6),
	}

	_, err := parseScaFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for severity=6")
	}
}

func TestParseScaFindingsRequest_PaginationValidation(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"page_size":        float64(0),
	}

	_, err := parseScaFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for size=0")
	}
}

func TestParseScaFindingsRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parseScaFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

// ============================================================================
// handleGetScaFindings tests
// ============================================================================

func TestScaFindingsTool_Initialize(t *testing.T) {
	tool := NewScaFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestScaFindingsTool_RegisterHandlers(t *testing.T) {
	tool := NewScaFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[ScaFindingsToolName] == nil {
		t.Error("Handler was not registered")
	}
}

func TestScaFindingsTool_HandleMissingApplicationPath(t *testing.T) {
	ctx := context.Background()

	result, err := handleGetScaFindings(ctx, map[string]interface{}{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["error"] == nil {
		t.Error("Expected error for missing application_path")
	}
}
