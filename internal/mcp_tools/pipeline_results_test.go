package mcp_tools

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

// ============================================================================
// parsePipelineResultsRequest tests
// ============================================================================

func TestParsePipelineResultsRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parsePipelineResultsRequest(args)
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

func TestParsePipelineResultsRequest_WithPagination(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"size":             float64(50),
		"page":             float64(2),
	}

	req, err := parsePipelineResultsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.Size != 50 {
		t.Errorf("Expected size 50, got %d", req.Size)
	}
	if req.Page != 2 {
		t.Errorf("Expected page 2, got %d", req.Page)
	}
}

func TestParsePipelineResultsRequest_SizeValidation(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"size":             float64(501),
	}

	_, err := parsePipelineResultsRequest(args)
	if err == nil {
		t.Fatal("Expected error for size=501")
	}
}

func TestParsePipelineResultsRequest_PageValidation(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"page":             float64(-1),
	}

	_, err := parsePipelineResultsRequest(args)
	if err == nil {
		t.Fatal("Expected error for page=-1")
	}
}

func TestParsePipelineResultsRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parsePipelineResultsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

// ============================================================================
// handleGetPipelineResults tests
// ============================================================================

func TestPipelineResultsTool_Initialize(t *testing.T) {
	tool := NewPipelineResultsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestPipelineResultsTool_RegisterHandlers(t *testing.T) {
	tool := NewPipelineResultsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[PipelineResultsToolName] == nil {
		t.Error("Handler was not registered")
	}
}

func TestPipelineResultsTool_HandleMissingResultsFile(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory without results file
	tempDir := t.TempDir()

	result, err := handlePipelineResults(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should return content with "No results found" message
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["content"] == nil {
		t.Error("Expected content in response")
	}
}

func TestPipelineResultsTool_HandleValidPath(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory with pipeline structure
	tempDir := t.TempDir()
	pipelineDir := filepath.Join(tempDir, ".veracode", "pipeline")
	if err := os.MkdirAll(pipelineDir, 0755); err != nil {
		t.Fatalf("Failed to create pipeline directory: %v", err)
	}

	result, err := handlePipelineResults(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	// Should always have content (even if no results file exists)
	if resultMap["content"] == nil {
		t.Error("Expected content in response")
	}
}
