package mcp_tools

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestStaticFindingsTool_AppProfileParameterOverride tests that the optional
// app_profile parameter overrides the workspace config
func TestStaticFindingsTool_AppProfileParameterOverride(t *testing.T) {
	// Create a temporary directory with a workspace config
	tempDir := t.TempDir()

	// Create workspace config file with one app name
	workspaceFile := filepath.Join(tempDir, ".veracode-workspace.json")
	workspaceContent := `{
  "name": "WorkspaceApp"
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
	handler := registry.handlers[StaticFindingsToolName]
	if handler == nil {
		t.Fatal("Handler not registered")
	}

	// Call the handler with app_profile parameter (should override workspace)
	ctx := context.Background()
	result, err := handler(ctx, map[string]interface{}{
		"application_path": tempDir,
		"app_profile":      "OverrideApp", // This should override WorkspaceApp
	})

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	// Verify the result
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

	// Verify the override app name appears in the response (not workspace name)
	if !strings.Contains(text, "OverrideApp") {
		t.Errorf("Response should contain override app name 'OverrideApp'\nGot: %s", text)
	}

	// Verify it mentions parameter override
	if !strings.Contains(text, "parameter") || !strings.Contains(text, "overriding") {
		t.Errorf("Response should mention parameter override\nGot: %s", text)
	}
}

// TestDynamicFindingsTool_AppProfileParameterOverride tests that the optional
// app_profile parameter overrides the workspace config
func TestDynamicFindingsTool_AppProfileParameterOverride(t *testing.T) {
	// Create a temporary directory with a workspace config
	tempDir := t.TempDir()

	// Create workspace config file with one app name
	workspaceFile := filepath.Join(tempDir, ".veracode-workspace.json")
	workspaceContent := `{
  "name": "WorkspaceApp"
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

	// Call the handler with app_profile parameter (should override workspace)
	ctx := context.Background()
	result, err := handler(ctx, map[string]interface{}{
		"application_path": tempDir,
		"app_profile":      "OverrideApp", // This should override WorkspaceApp
	})

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	// Verify the result
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

	// Verify the override app name appears in the response (not workspace name)
	if !strings.Contains(text, "OverrideApp") {
		t.Errorf("Response should contain override app name 'OverrideApp'\nGot: %s", text)
	}

	// Verify it mentions parameter override
	if !strings.Contains(text, "parameter") || !strings.Contains(text, "overriding") {
		t.Errorf("Response should mention parameter override\nGot: %s", text)
	}
}
