package mcp_tools

import (
	"context"
	"testing"
)

// ============================================================================
// parseDynamicFindingsRequest tests
// ============================================================================

func TestParseDynamicFindingsRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"app_profile":      "MyApp",
		"sandbox":          "Development",
		"size":             float64(25),
		"page":             float64(1),
	}

	req, err := parseDynamicFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
	if req.AppProfile != "MyApp" {
		t.Errorf("Expected app_profile 'MyApp', got '%s'", req.AppProfile)
	}
	if req.Sandbox != "Development" {
		t.Errorf("Expected sandbox 'Development', got '%s'", req.Sandbox)
	}
	if req.Size != 25 {
		t.Errorf("Expected size 25, got %d", req.Size)
	}
	if req.Page != 1 {
		t.Errorf("Expected page 1, got %d", req.Page)
	}
}

func TestParseDynamicFindingsRequest_DefaultValues(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parseDynamicFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.Size != 10 {
		t.Errorf("Expected default size 10, got %d", req.Size)
	}
	if req.Page != 0 {
		t.Errorf("Expected default page 0, got %d", req.Page)
	}
	if req.AppProfile != "" {
		t.Errorf("Expected empty app_profile, got '%s'", req.AppProfile)
	}
}

func TestParseDynamicFindingsRequest_WithSeverity(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"severity":         float64(4),
	}

	req, err := parseDynamicFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.Severity == nil {
		t.Fatal("Expected severity to be set")
	}
	if *req.Severity != 4 {
		t.Errorf("Expected severity 4, got %d", *req.Severity)
	}
}

func TestParseDynamicFindingsRequest_SeverityOutOfRange(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"severity":         float64(6),
	}

	_, err := parseDynamicFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for severity=6")
	}
}

func TestParseDynamicFindingsRequest_SizeTooSmall(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"size":             float64(0),
	}

	_, err := parseDynamicFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for size=0")
	}
}

func TestParseDynamicFindingsRequest_PageNegative(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"page":             float64(-1),
	}

	_, err := parseDynamicFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for page=-1")
	}
}

func TestParseDynamicFindingsRequest_WithCWEIDs(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"cwe_ids":          []interface{}{float64(79), float64(89)},
	}

	req, err := parseDynamicFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(req.CWEIDs) != 2 {
		t.Fatalf("Expected 2 CWE IDs, got %d", len(req.CWEIDs))
	}
	if req.CWEIDs[0] != "79" || req.CWEIDs[1] != "89" {
		t.Errorf("Expected ['79', '89'], got %v", req.CWEIDs)
	}
}

func TestParseDynamicFindingsRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parseDynamicFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

// ============================================================================
// handleGetDynamicFindings tests
// ============================================================================

func TestDynamicFindingsTool_Initialize(t *testing.T) {
	tool := NewDynamicFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestDynamicFindingsTool_RegisterHandlers(t *testing.T) {
	tool := NewDynamicFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[DynamicFindingsToolName] == nil {
		t.Error("Handler was not registered")
	}
}

func TestDynamicFindingsTool_HandleMissingApplicationPath(t *testing.T) {
	ctx := context.Background()

	result, err := handleGetDynamicFindings(ctx, map[string]interface{}{})
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
