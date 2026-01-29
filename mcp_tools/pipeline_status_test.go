package mcp_tools

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestPipelineStatusTool_Initialize(t *testing.T) {
	tool := NewPipelineStatusTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestPipelineStatusTool_RegisterHandlers(t *testing.T) {
	tool := NewPipelineStatusTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[PipelineStatusToolName] == nil {
		t.Fatal("Handler not registered")
	}
}

func TestParsePipelineStatusRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parsePipelineStatusRequest(args)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
}

func TestParsePipelineStatusRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parsePipelineStatusRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

func TestPipelineStatusTool_NoPIDFile(t *testing.T) {
	tool := NewPipelineStatusTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	handler := registry.handlers[PipelineStatusToolName]
	ctx := context.Background()

	// Create a temporary directory
	tempDir := t.TempDir()

	result, err := handler(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should return content indicating no scan found
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["content"] == nil {
		t.Error("Expected content in result")
	}
}

func TestPipelineStatusTool_WithPIDFile(t *testing.T) {
	tool := NewPipelineStatusTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	handler := registry.handlers[PipelineStatusToolName]
	ctx := context.Background()

	// Create a temporary directory
	tempDir := t.TempDir()
	outputDir := filepath.Join(tempDir, ".veracode_pipeline")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Fatalf("Failed to create output directory: %v", err)
	}

	// Write PID file with current process PID (which is running)
	pidFile := filepath.Join(outputDir, "pipeline.pid")
	if err := os.WriteFile(pidFile, []byte("999999"), 0644); err != nil {
		t.Fatalf("Failed to write PID file: %v", err)
	}

	result, err := handler(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should return content with status
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["content"] == nil {
		t.Error("Expected content in result")
	}
}

func TestCheckProcessStatus_NonexistentPID(t *testing.T) {
	// Use a very high PID that likely doesn't exist
	isRunning, exitCode := checkProcessStatus(999999)

	if isRunning {
		t.Error("Expected process to not be running for nonexistent PID")
	}

	if exitCode != -1 {
		t.Errorf("Expected exit code -1 for nonexistent process, got %d", exitCode)
	}
}

func TestCheckProcessStatus_CurrentProcess(t *testing.T) {
	// Check the current process (which is definitely running)
	currentPID := os.Getpid()
	isRunning, _ := checkProcessStatus(currentPID)

	// On Windows, process status checking is less reliable
	// so we'll just ensure it doesn't crash
	if runtime.GOOS != "windows" {
		if !isRunning {
			t.Error("Expected current process to be running")
		}
	}
}
