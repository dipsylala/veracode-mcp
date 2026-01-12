package tools

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestStaticFindingsTool_WithWorkspaceConfig tests that the static findings tool
// correctly reads the application name from .veracode-workspace.json
func TestStaticFindingsTool_WithWorkspaceConfig(t *testing.T) {
	// Create a temporary directory with a workspace config
	tempDir := t.TempDir()

	// Create workspace config file
	workspaceFile := filepath.Join(tempDir, ".veracode-workspace.json")
	workspaceContent := `{
  "name": "TestApp-StaticScan"
}`
	if err := os.WriteFile(workspaceFile, []byte(workspaceContent), 0644); err != nil {
		t.Fatalf("Failed to create workspace file: %v", err)
	}

	// Create the tool
	tool := NewStaticFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	// Create a mock registry
	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	// Get the handler
	handler := registry.handlers["get-static-findings"]
	if handler == nil {
		t.Fatal("Handler not registered")
	}

	// Call the handler with application_path
	ctx := context.Background()
	result, err := handler(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	// Verify the result contains the application name from workspace
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("Result is not a map: %T", result)
	}

	// Check for content
	content, hasContent := resultMap["content"]
	if !hasContent {
		t.Fatal("Result should have 'content' field")
	}

	contentList, ok := content.([]map[string]string)
	if !ok {
		t.Fatalf("Content is not a list of maps: %T", content)
	}

	if len(contentList) == 0 {
		t.Fatal("Content list is empty")
	}

	text := contentList[0]["text"]

	// Verify the application name appears in the response
	if !strings.Contains(text, "TestApp-StaticScan") {
		t.Errorf("Response should contain application name 'TestApp-StaticScan'\nGot: %s", text)
	}

	// Verify it mentions workspace config
	if !strings.Contains(text, "workspace config") {
		t.Errorf("Response should mention 'workspace config'\nGot: %s", text)
	}
}

// TestStaticFindingsTool_MissingWorkspaceConfig tests error handling when
// .veracode-workspace.json is not found
func TestStaticFindingsTool_MissingWorkspaceConfig(t *testing.T) {
	// Create a temporary directory WITHOUT a workspace config
	tempDir := t.TempDir()

	// Create the tool
	tool := NewStaticFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	// Create a mock registry
	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	// Get the handler
	handler := registry.handlers[StaticFindingsToolName]
	if handler == nil {
		t.Fatal("Handler not registered")
	}

	// Call the handler with application_path (no workspace file)
	ctx := context.Background()
	result, err := handler(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	// Verify the result contains an error message
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("Result is not a map: %T", result)
	}

	// Check for error field
	errorMsg, hasError := resultMap["error"]
	if !hasError {
		t.Fatal("Result should have 'error' field when workspace config is missing")
	}

	errorStr, ok := errorMsg.(string)
	if !ok {
		t.Fatalf("Error is not a string: %T", errorMsg)
	}

	// Verify error message is helpful
	expectedPhrases := []string{
		"workspace configuration",
		".veracode-workspace.json",
	}

	for _, phrase := range expectedPhrases {
		if !strings.Contains(errorStr, phrase) {
			t.Errorf("Error message should contain '%s'\nGot: %s", phrase, errorStr)
		}
	}
}

// TestDynamicFindingsTool_WithWorkspaceConfig tests that the dynamic findings tool
// correctly reads the application name from .veracode-workspace.json
func TestDynamicFindingsTool_WithWorkspaceConfig(t *testing.T) {
	// Create a temporary directory with a workspace config
	tempDir := t.TempDir()

	// Create workspace config file
	workspaceFile := filepath.Join(tempDir, ".veracode-workspace.json")
	workspaceContent := `{
  "name": "TestApp-DynamicScan"
}`
	if err := os.WriteFile(workspaceFile, []byte(workspaceContent), 0644); err != nil {
		t.Fatalf("Failed to create workspace file: %v", err)
	}

	// Create the tool
	tool := NewDynamicFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	// Create a mock registry
	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	// Get the handler
	handler := registry.handlers[DynamicFindingsToolName]
	if handler == nil {
		t.Fatal("Handler not registered")
	}

	// Call the handler with application_path
	ctx := context.Background()
	result, err := handler(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	// Verify the result contains the application name from workspace
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("Result is not a map: %T", result)
	}

	// Check for content
	content, hasContent := resultMap["content"]
	if !hasContent {
		t.Fatal("Result should have 'content' field")
	}

	contentList, ok := content.([]map[string]string)
	if !ok {
		t.Fatalf("Content is not a list of maps: %T", content)
	}

	if len(contentList) == 0 {
		t.Fatal("Content list is empty")
	}

	text := contentList[0]["text"]

	// Verify the application name appears in the response
	if !strings.Contains(text, "TestApp-DynamicScan") {
		t.Errorf("Response should contain application name 'TestApp-DynamicScan'\nGot: %s", text)
	}

	// Verify it mentions workspace config
	if !strings.Contains(text, "workspace config") {
		t.Errorf("Response should mention 'workspace config'\nGot: %s", text)
	}
}

// TestDynamicFindingsTool_MissingWorkspaceConfig tests error handling when
// .veracode-workspace.json is not found
func TestDynamicFindingsTool_MissingWorkspaceConfig(t *testing.T) {
	// Create a temporary directory WITHOUT a workspace config
	tempDir := t.TempDir()

	// Create the tool
	tool := NewDynamicFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	// Create a mock registry
	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	// Get the handler
	handler := registry.handlers[DynamicFindingsToolName]
	if handler == nil {
		t.Fatal("Handler not registered")
	}

	// Call the handler with application_path (no workspace file)
	ctx := context.Background()
	result, err := handler(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	// Verify the result contains an error message
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("Result is not a map: %T", result)
	}

	// Check for error field
	errorMsg, hasError := resultMap["error"]
	if !hasError {
		t.Fatal("Result should have 'error' field when workspace config is missing")
	}

	errorStr, ok := errorMsg.(string)
	if !ok {
		t.Fatalf("Error is not a string: %T", errorMsg)
	}

	// Verify error message is helpful
	expectedPhrases := []string{
		"workspace configuration",
		".veracode-workspace.json",
	}

	for _, phrase := range expectedPhrases {
		if !strings.Contains(errorStr, phrase) {
			t.Errorf("Error message should contain '%s'\nGot: %s", phrase, errorStr)
		}
	}
}
