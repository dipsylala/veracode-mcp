package mcp_tools

import (
	"context"
	"testing"
)

// ============================================================================
// parsePipelineDetailedResultsRequest tests
// ============================================================================

func TestParsePipelineDetailedResultsRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          float64(999),
	}

	req, err := parsePipelineDetailedResultsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
	if req.FlawID != 999 {
		t.Errorf("Expected flaw_id 999, got %d", req.FlawID)
	}
}

func TestParsePipelineDetailedResultsRequest_LargeFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          float64(123456789),
	}

	req, err := parsePipelineDetailedResultsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.FlawID != 123456789 {
		t.Errorf("Expected flaw_id 123456789, got %d", req.FlawID)
	}
}

func TestParsePipelineDetailedResultsRequest_MissingFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	_, err := parsePipelineDetailedResultsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing flaw_id")
	}
}

func TestParsePipelineDetailedResultsRequest_ZeroFlawID(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          float64(0),
	}

	_, err := parsePipelineDetailedResultsRequest(args)
	if err == nil {
		t.Fatal("Expected error for flaw_id=0")
	}
}

func TestParsePipelineDetailedResultsRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{
		"flaw_id": float64(12345),
	}

	_, err := parsePipelineDetailedResultsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

// ============================================================================
// handleGetPipelineDetailedResults tests
// ============================================================================

func TestPipelineDetailedResultsTool_Initialize(t *testing.T) {
	tool := NewPipelineDetailedResultsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestPipelineDetailedResultsTool_RegisterHandlers(t *testing.T) {
	tool := NewPipelineDetailedResultsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[PipelineDetailedResultsToolName] == nil {
		t.Error("Handler was not registered")
	}
}

func TestPipelineDetailedResultsTool_HandleMissingFlawID(t *testing.T) {
	ctx := context.Background()

	result, err := handlePipelineDetailedResults(ctx, map[string]interface{}{
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

func TestPipelineDetailedResultsTool_HandleInvalidFlawID(t *testing.T) {
	ctx := context.Background()

	result, err := handlePipelineDetailedResults(ctx, map[string]interface{}{
		"application_path": "/path/to/app",
		"flaw_id":          float64(0),
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

func TestPipelineDetailedResultsTool_HandleMissingResultsFile(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory without results file
	tempDir := t.TempDir()

	result, err := handlePipelineDetailedResults(ctx, map[string]interface{}{
		"application_path": tempDir,
		"flaw_id":          float64(12345),
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	// Should have content (likely an error message about missing results)
	if resultMap["content"] == nil && resultMap["error"] == nil {
		t.Error("Expected either content or error in response")
	}
}
