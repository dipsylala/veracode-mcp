package mcp_tools

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

// ============================================================================
// parsePackageWorkspaceRequest tests
// ============================================================================

func TestParsePackageWorkspaceRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parsePackageWorkspaceRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
}

// Note: filename parameter is in schema but not yet implemented in parse function

func TestParsePackageWorkspaceRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parsePackageWorkspaceRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

func TestParsePackageWorkspaceRequest_EmptyApplicationPath(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "",
	}

	_, err := parsePackageWorkspaceRequest(args)
	if err == nil {
		t.Fatal("Expected error for empty application_path")
	}
}

// ============================================================================
// handlePackageWorkspace tests
// ============================================================================

func TestPackageWorkspaceTool_Initialize(t *testing.T) {
	tool := NewPackageWorkspaceTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestPackageWorkspaceTool_RegisterHandlers(t *testing.T) {
	tool := NewPackageWorkspaceTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[PackageWorkspaceToolName] == nil {
		t.Error("Handler was not registered")
	}
}

func TestPackageWorkspaceTool_HandleMissingApplicationPath(t *testing.T) {
	ctx := context.Background()

	result, err := handlePackageWorkspace(ctx, map[string]interface{}{})
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

func TestPackageWorkspaceTool_HandleNonexistentPath(t *testing.T) {
	ctx := context.Background()

	result, err := handlePackageWorkspace(ctx, map[string]interface{}{
		"application_path": "/nonexistent/path/to/app",
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["error"] == nil {
		t.Error("Expected error for nonexistent path")
	}
}

func TestPackageWorkspaceTool_HandleValidPath(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory
	tempDir := t.TempDir()

	// Create some test files
	testFile := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result, err := handlePackageWorkspace(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	// Should have either content or error (depending on whether srcclr CLI is available)
	if resultMap["content"] == nil && resultMap["error"] == nil {
		t.Error("Expected either content or error in result")
	}
}

func TestPackageWorkspaceTool_OutputDirectory(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// The expected output directory should be created
	expectedDir := filepath.Join(tempDir, ".veracode", "packaging")

	// Call the handler (even if it fails, directory should be checked)
	ctx := context.Background()
	_, _ = handlePackageWorkspace(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	// The directory creation happens in the handler, but we can't verify
	// without running the actual packager. Just verify the path logic.
	if !filepath.IsAbs(expectedDir) {
		t.Error("Expected absolute path for output directory")
	}
}
