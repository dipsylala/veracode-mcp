package mcp_tools

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestParseLocalSCAScanRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parseLocalSCAScanRequest(args)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
}

func TestParseLocalSCAScanRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parseLocalSCAScanRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

func TestParseLocalSCAScanRequest_EmptyApplicationPath(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "",
	}

	_, err := parseLocalSCAScanRequest(args)
	if err == nil {
		t.Fatal("Expected error for empty application_path")
	}
}

func TestValidateAndPrepareSCADirectories_Success(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	outputDir, outputFile, err := validateAndPrepareSCADirectories(tempDir)
	if err != nil {
		t.Fatalf("Failed to validate and prepare directories: %v", err)
	}

	expectedOutputDir := filepath.Join(tempDir, ".veracode", "sca")
	if outputDir != expectedOutputDir {
		t.Errorf("Expected output directory '%s', got '%s'", expectedOutputDir, outputDir)
	}

	expectedOutputFile := filepath.Join(expectedOutputDir, "veracode.json")
	if outputFile != expectedOutputFile {
		t.Errorf("Expected output file '%s', got '%s'", expectedOutputFile, outputFile)
	}

	// Verify the directory was created
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		t.Errorf("Expected directory '%s' to be created", outputDir)
	}
}

func TestValidateAndPrepareSCADirectories_NonexistentPath(t *testing.T) {
	_, _, err := validateAndPrepareSCADirectories("/nonexistent/path/to/app")
	if err == nil {
		t.Fatal("Expected error for nonexistent path")
	}
}

func TestValidateAndPrepareSCADirectories_CreatesDirectoryIfNeeded(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Call twice to ensure it handles existing directory
	_, _, err := validateAndPrepareSCADirectories(tempDir)
	if err != nil {
		t.Fatalf("First call failed: %v", err)
	}

	outputDir, outputFile, err := validateAndPrepareSCADirectories(tempDir)
	if err != nil {
		t.Fatalf("Second call failed: %v", err)
	}

	expectedOutputDir := filepath.Join(tempDir, ".veracode", "sca")
	if outputDir != expectedOutputDir {
		t.Errorf("Expected output directory '%s', got '%s'", expectedOutputDir, outputDir)
	}

	expectedOutputFile := filepath.Join(expectedOutputDir, "veracode.json")
	if outputFile != expectedOutputFile {
		t.Errorf("Expected output file '%s', got '%s'", expectedOutputFile, outputFile)
	}

	// Verify the directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		t.Errorf("Expected directory '%s' to exist", outputDir)
	}
}

func TestLocalSCAScanTool_HandleInvalidPath(t *testing.T) {
	ctx := context.Background()

	result, err := handleLocalSCAScan(ctx, map[string]interface{}{
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

func TestLocalSCAScanTool_HandleValidPath(t *testing.T) {
	tool := NewLocalSCAScanTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	handler := registry.handlers[LocalSCAScanToolName]
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

func TestLocalSCAScanTool_OutputFileLocation(t *testing.T) {
	tempDir := t.TempDir()

	// Create the expected output structure
	outputDir, outputFile, err := validateAndPrepareSCADirectories(tempDir)
	if err != nil {
		t.Fatalf("Failed to prepare directories: %v", err)
	}

	expectedDir := filepath.Join(tempDir, ".veracode", "sca")
	if outputDir != expectedDir {
		t.Errorf("Expected output dir '%s', got '%s'", expectedDir, outputDir)
	}

	expectedFile := filepath.Join(expectedDir, "veracode.json")
	if outputFile != expectedFile {
		t.Errorf("Expected output file '%s', got '%s'", expectedFile, outputFile)
	}

	// Verify directory was created
	info, err := os.Stat(outputDir)
	if err != nil {
		t.Fatalf("Output directory was not created: %v", err)
	}

	if !info.IsDir() {
		t.Error("Output path is not a directory")
	}
}

func TestBuildSCAScanResponse_Success(t *testing.T) {
	req := &LocalSCAScanRequest{
		ApplicationPath: "/test/path",
	}

	outputDir := "/test/path/.veracode/sca"
	outputFile := "/test/path/.veracode/sca/veracode.json"

	var stdout, stderr bytes.Buffer
	response := buildSCAScanResponse(req, outputDir, outputFile, 0, 1000000, stdout, stderr)

	// Should have content for successful exit code
	if response["content"] == nil {
		t.Error("Expected content in response")
	}
}

func TestBuildSCAScanResponse_Error(t *testing.T) {
	req := &LocalSCAScanRequest{
		ApplicationPath: "/test/path",
	}

	outputDir := "/test/path/.veracode/sca"
	outputFile := "/test/path/.veracode/sca/veracode.json"

	var stdout, stderr bytes.Buffer
	// Exit code 1 is a real error (not a warning)
	response := buildSCAScanResponse(req, outputDir, outputFile, 1, 1000000, stdout, stderr)

	// Should have error for failing exit code
	if response["error"] == nil {
		t.Error("Expected error in response for non-zero exit code")
	}
}

func TestBuildSCAScanResponse_Warning(t *testing.T) {
	req := &LocalSCAScanRequest{
		ApplicationPath: "/test/path",
	}

	outputDir := "/test/path/.veracode/sca"
	outputFile := "/test/path/.veracode/sca/veracode.json"

	var stdout, stderr bytes.Buffer
	// Exit code 3 is a warning (policy failure but scan succeeded)
	response := buildSCAScanResponse(req, outputDir, outputFile, 3, 1000000, stdout, stderr)

	// Should have content (not error) for warning exit codes
	if response["content"] == nil {
		t.Error("Expected content in response for warning exit code")
	}
}
