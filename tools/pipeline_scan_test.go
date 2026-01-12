package tools

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestPipelineScanTool_Initialize(t *testing.T) {
	tool := NewPipelineScanTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestPipelineScanTool_RegisterHandlers(t *testing.T) {
	tool := NewPipelineScanTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[PipelineScanToolName] == nil {
		t.Fatal("Handler not registered")
	}
}

func TestParsePipelineScanRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parsePipelineScanRequest(args)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}

	if req.Filename != "" {
		t.Errorf("Expected empty filename, got '%s'", req.Filename)
	}
}

func TestParsePipelineScanRequest_WithFilename(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"filename":         "myapp.zip",
	}

	req, err := parsePipelineScanRequest(args)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}

	if req.Filename != "myapp.zip" {
		t.Errorf("Expected filename 'myapp.zip', got '%s'", req.Filename)
	}

	if req.Verbose {
		t.Errorf("Expected verbose to be false, got true")
	}
}

func TestParsePipelineScanRequest_WithVerbose(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"verbose":          true,
	}

	req, err := parsePipelineScanRequest(args)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}

	if !req.Verbose {
		t.Errorf("Expected verbose to be true, got false")
	}
}

func TestParsePipelineScanRequest_AllParameters(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"filename":         "myapp.zip",
		"verbose":          true,
	}

	req, err := parsePipelineScanRequest(args)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}

	if req.Filename != "myapp.zip" {
		t.Errorf("Expected filename 'myapp.zip', got '%s'", req.Filename)
	}

	if !req.Verbose {
		t.Errorf("Expected verbose to be true, got false")
	}
}

func TestParsePipelineScanRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parsePipelineScanRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

func TestParsePipelineScanRequest_EmptyApplicationPath(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "",
	}

	_, err := parsePipelineScanRequest(args)
	if err == nil {
		t.Fatal("Expected error for empty application_path")
	}
}

func TestPipelineScanTool_HandleInvalidPath(t *testing.T) {
	tool := NewPipelineScanTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	handler := registry.handlers[PipelineScanToolName]
	ctx := context.Background()

	result, err := handler(ctx, map[string]interface{}{
		"application_path": "/nonexistent/path/to/app",
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should return an error response
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["error"] == nil {
		t.Error("Expected error for nonexistent path")
	}
}

func TestPipelineScanTool_HandleValidPath(t *testing.T) {
	tool := NewPipelineScanTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	handler := registry.handlers[PipelineScanToolName]
	ctx := context.Background()

	// Create a temporary directory
	tempDir := t.TempDir()

	result, err := handler(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Result should be a map
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	// Should have either content or error (depending on whether veracode CLI is available)
	if resultMap["content"] == nil && resultMap["error"] == nil {
		t.Error("Expected either content or error in result")
	}
}

func TestFindLargestFile_Success(t *testing.T) {
	// Create a temporary directory with multiple files
	tempDir := t.TempDir()

	// Create some test files with different sizes
	smallFile := filepath.Join(tempDir, "small.txt")
	if err := os.WriteFile(smallFile, []byte("small"), 0644); err != nil {
		t.Fatalf("Failed to create small file: %v", err)
	}

	largeFile := filepath.Join(tempDir, "large.txt")
	if err := os.WriteFile(largeFile, []byte("this is a much larger file with more content"), 0644); err != nil {
		t.Fatalf("Failed to create large file: %v", err)
	}

	mediumFile := filepath.Join(tempDir, "medium.txt")
	if err := os.WriteFile(mediumFile, []byte("medium size"), 0644); err != nil {
		t.Fatalf("Failed to create medium file: %v", err)
	}

	// Find the largest file
	result, err := findLargestFile(tempDir)
	if err != nil {
		t.Fatalf("Failed to find largest file: %v", err)
	}

	if result != largeFile {
		t.Errorf("Expected largest file to be '%s', got '%s'", largeFile, result)
	}
}

func TestFindLargestFile_NonexistentDir(t *testing.T) {
	_, err := findLargestFile("/nonexistent/directory")
	if err == nil {
		t.Error("Expected error for nonexistent directory")
	}
}

func TestFindLargestFile_EmptyDir(t *testing.T) {
	tempDir := t.TempDir()

	_, err := findLargestFile(tempDir)
	if err == nil {
		t.Error("Expected error for empty directory")
	}
}

func TestFindLargestFile_IgnoresDirectories(t *testing.T) {
	tempDir := t.TempDir()

	// Create a subdirectory
	subDir := filepath.Join(tempDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	// Create a file
	file := filepath.Join(tempDir, "file.txt")
	if err := os.WriteFile(file, []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Find the largest file (should skip the directory)
	result, err := findLargestFile(tempDir)
	if err != nil {
		t.Fatalf("Failed to find largest file: %v", err)
	}

	if result != file {
		t.Errorf("Expected file to be '%s', got '%s'", file, result)
	}
}
