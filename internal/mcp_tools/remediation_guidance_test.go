package mcp_tools

import (
	"context"
	"testing"
)

// ============================================================================
// parseRemediationGuidanceRequest tests
// ============================================================================

func TestParseRemediationGuidanceRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          "555",
	}

	req, err := parseRemediationGuidanceRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
	if req.FlawID.IssueID != 555 {
		t.Errorf("Expected flaw_id 555, got %d", req.FlawID.IssueID)
	}
}

func TestParseRemediationGuidanceRequest_LargeFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          "987654321",
	}

	req, err := parseRemediationGuidanceRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.FlawID.IssueID != 987654321 {
		t.Errorf("Expected flaw_id 987654321, got %d", req.FlawID.IssueID)
	}
}

func TestParseRemediationGuidanceRequest_MissingFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	_, err := parseRemediationGuidanceRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing flaw_id")
	}
}

func TestParseRemediationGuidanceRequest_ZeroFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          "0",
	}

	_, err := parseRemediationGuidanceRequest(args)
	if err == nil {
		t.Fatal("Expected error for flaw_id=0")
	}
}

func TestParseRemediationGuidanceRequest_NegativeFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          "-123",
	}

	_, err := parseRemediationGuidanceRequest(args)
	if err == nil {
		t.Fatal("Expected error for negative flaw_id")
	}
}

func TestParseRemediationGuidanceRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{
		"flaw_id": "555",
	}

	_, err := parseRemediationGuidanceRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

// ============================================================================
// handleRemediationGuidance tests
// ============================================================================

func TestRemediationGuidanceTool_Initialize(t *testing.T) {
	tool := NewRemediationGuidanceTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestRemediationGuidanceTool_RegisterHandlers(t *testing.T) {
	tool := NewRemediationGuidanceTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[RemediationGuidanceToolName] == nil {
		t.Error("Handler was not registered")
	}
}

func TestRemediationGuidanceTool_HandleMissingFlawID(t *testing.T) {
	ctx := context.Background()

	result, err := handleGetRemediationGuidance(ctx, map[string]interface{}{
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

func TestRemediationGuidanceTool_HandleInvalidFlawID(t *testing.T) {
	ctx := context.Background()

	result, err := handleGetRemediationGuidance(ctx, map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          "0",
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["error"] == nil {
		t.Error("Expected error for flaw_id=0")
	}
}

func TestRemediationGuidanceTool_HandleMissingResultsFile(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory without results file
	tempDir := t.TempDir()

	result, err := handleGetRemediationGuidance(ctx, map[string]interface{}{
		"application_path": tempDir,
		"flaw_id":          "12345",
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	// Should have content or error (likely error about missing results file)
	if resultMap["content"] == nil && resultMap["error"] == nil {
		t.Error("Expected either content or error in response")
	}
}
