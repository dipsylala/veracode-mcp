package mcp_tools

import (
	"context"
	"testing"
)

// ============================================================================
// parseStaticFindingsRequest tests
// ============================================================================

func TestParseStaticFindingsRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parseStaticFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
	if req.Size != 10 {
		t.Errorf("Expected default size 10, got %d", req.Size)
	}
}

func TestParseStaticFindingsRequest_AllParameters(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"app_profile":      "MyApp",
		"sandbox":          "Dev",
		"size":             float64(50),
		"page":             float64(2),
		"severity":         float64(4),
		"severity_gte":     float64(2),
		"violates_policy":  true,
		"cwe_ids":          []interface{}{float64(79), float64(89)},
	}

	req, err := parseStaticFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.AppProfile != "MyApp" {
		t.Errorf("Expected app_profile 'MyApp', got '%s'", req.AppProfile)
	}
	if req.Size != 50 {
		t.Errorf("Expected size 50, got %d", req.Size)
	}
	if len(req.CWEIDs) != 2 {
		t.Errorf("Expected 2 CWE IDs, got %d", len(req.CWEIDs))
	}
}

func TestParseStaticFindingsRequest_SeverityValidation(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"severity":         float64(6),
	}

	_, err := parseStaticFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for severity=6")
	}
}

func TestParseStaticFindingsRequest_PaginationValidation(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"size":             float64(501),
	}

	_, err := parseStaticFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for size=501")
	}
}

func TestParseStaticFindingsRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parseStaticFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

// ============================================================================
// handleGetStaticFindings tests
// ============================================================================

func TestStaticFindingsTool_Initialize(t *testing.T) {
	tool := NewStaticFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestStaticFindingsTool_RegisterHandlers(t *testing.T) {
	tool := NewStaticFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[StaticFindingsToolName] == nil {
		t.Error("Handler was not registered")
	}
}

func TestStaticFindingsTool_HandleMissingApplicationPath(t *testing.T) {
	ctx := context.Background()

	result, err := handleGetStaticFindings(ctx, map[string]interface{}{})
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
