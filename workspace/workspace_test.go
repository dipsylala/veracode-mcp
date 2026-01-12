package workspace

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFindWorkspaceConfig_Success(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create a valid workspace file
	workspaceFile := filepath.Join(tempDir, WorkspaceFileName)
	content := `{
  "name": "TestApplication"
}`
	if err := os.WriteFile(workspaceFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test workspace file: %v", err)
	}

	// Test finding the config
	name, err := FindWorkspaceConfig(tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if name != "TestApplication" {
		t.Errorf("Expected name 'TestApplication', got '%s'", name)
	}
}

func TestFindWorkspaceConfig_FileNotFound(t *testing.T) {
	// Create a temporary directory without a workspace file
	tempDir := t.TempDir()

	// Test finding the config
	name, err := FindWorkspaceConfig(tempDir)
	if err == nil {
		t.Fatal("Expected error for missing workspace file, got nil")
	}

	if name != "" {
		t.Errorf("Expected empty name, got '%s'", name)
	}

	// Verify error message contains helpful guidance
	errMsg := err.Error()
	expectedPhrases := []string{
		"workspace configuration not found",
		WorkspaceFileName,
		"Create a file named",
		`"name": "YourApplicationProfileName"`,
	}

	for _, phrase := range expectedPhrases {
		if !strings.Contains(errMsg, phrase) {
			t.Errorf("Error message should contain '%s'\nGot: %s", phrase, errMsg)
		}
	}
}

func TestFindWorkspaceConfig_InvalidJSON(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create an invalid workspace file
	workspaceFile := filepath.Join(tempDir, WorkspaceFileName)
	content := `{
  "name": "TestApp"
  "invalid": json
}`
	if err := os.WriteFile(workspaceFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test workspace file: %v", err)
	}

	// Test finding the config
	name, err := FindWorkspaceConfig(tempDir)
	if err == nil {
		t.Fatal("Expected error for invalid JSON, got nil")
	}

	if name != "" {
		t.Errorf("Expected empty name, got '%s'", name)
	}

	// Verify error message contains helpful guidance
	errMsg := err.Error()
	if !strings.Contains(errMsg, "invalid JSON") {
		t.Errorf("Error message should mention 'invalid JSON'\nGot: %s", errMsg)
	}
}

func TestFindWorkspaceConfig_EmptyName(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create a workspace file with empty name
	workspaceFile := filepath.Join(tempDir, WorkspaceFileName)
	content := `{
  "name": ""
}`
	if err := os.WriteFile(workspaceFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test workspace file: %v", err)
	}

	// Test finding the config
	name, err := FindWorkspaceConfig(tempDir)
	if err == nil {
		t.Fatal("Expected error for empty name, got nil")
	}

	if name != "" {
		t.Errorf("Expected empty name, got '%s'", name)
	}

	// Verify error message contains helpful guidance
	errMsg := err.Error()
	expectedPhrases := []string{
		"missing or empty",
		`"name"`,
	}

	for _, phrase := range expectedPhrases {
		if !strings.Contains(errMsg, phrase) {
			t.Errorf("Error message should contain '%s'\nGot: %s", phrase, errMsg)
		}
	}
}

func TestFindWorkspaceConfig_MissingNameField(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create a workspace file without name field
	workspaceFile := filepath.Join(tempDir, WorkspaceFileName)
	content := `{
  "other": "field"
}`
	if err := os.WriteFile(workspaceFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test workspace file: %v", err)
	}

	// Test finding the config
	name, err := FindWorkspaceConfig(tempDir)
	if err == nil {
		t.Fatal("Expected error for missing name field, got nil")
	}

	if name != "" {
		t.Errorf("Expected empty name, got '%s'", name)
	}
}

func TestFindWorkspaceConfig_DirectoryDoesNotExist(t *testing.T) {
	// Use a non-existent directory
	nonExistentDir := filepath.Join(t.TempDir(), "does-not-exist")

	// Test finding the config
	name, err := FindWorkspaceConfig(nonExistentDir)
	if err == nil {
		t.Fatal("Expected error for non-existent directory, got nil")
	}

	if name != "" {
		t.Errorf("Expected empty name, got '%s'", name)
	}

	// Verify error message mentions directory doesn't exist
	errMsg := err.Error()
	if !strings.Contains(errMsg, "does not exist") {
		t.Errorf("Error message should mention directory doesn't exist\nGot: %s", errMsg)
	}
}

func TestFindWorkspaceConfig_RelativePath(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create a valid workspace file
	workspaceFile := filepath.Join(tempDir, WorkspaceFileName)
	content := `{
  "name": "RelativePathTest"
}`
	if err := os.WriteFile(workspaceFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test workspace file: %v", err)
	}

	// Save current directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	// Change to parent of temp directory
	if err := os.Chdir(filepath.Dir(tempDir)); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Use relative path (just the directory name)
	relativePath := filepath.Base(tempDir)

	// Test finding the config with relative path
	name, err := FindWorkspaceConfig(relativePath)
	if err != nil {
		t.Fatalf("Expected no error with relative path, got: %v", err)
	}

	if name != "RelativePathTest" {
		t.Errorf("Expected name 'RelativePathTest', got '%s'", name)
	}
}

func TestFindWorkspaceConfig_WithSpecialCharacters(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create a workspace file with special characters in the name
	workspaceFile := filepath.Join(tempDir, WorkspaceFileName)
	content := `{
  "name": "My-App_v2.0 (Production)"
}`
	if err := os.WriteFile(workspaceFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test workspace file: %v", err)
	}

	// Test finding the config
	name, err := FindWorkspaceConfig(tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expected := "My-App_v2.0 (Production)"
	if name != expected {
		t.Errorf("Expected name '%s', got '%s'", expected, name)
	}
}

func TestFindWorkspaceConfigInCurrentDir(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create a valid workspace file
	workspaceFile := filepath.Join(tempDir, WorkspaceFileName)
	content := `{
  "name": "CurrentDirTest"
}`
	if err := os.WriteFile(workspaceFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test workspace file: %v", err)
	}

	// Save current directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	// Change to temp directory
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Test finding the config in current directory
	name, err := FindWorkspaceConfigInCurrentDir()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if name != "CurrentDirTest" {
		t.Errorf("Expected name 'CurrentDirTest', got '%s'", name)
	}
}
