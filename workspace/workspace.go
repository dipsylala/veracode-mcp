package workspace

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const WorkspaceFileName = ".veracode-workspace.json"

// WorkspaceConfig represents the structure of .veracode-workspace.json
type WorkspaceConfig struct {
	Name string `json:"name"`
}

// FindWorkspaceConfig searches for .veracode-workspace.json in the given directory
// and returns the application profile name.
//
// Returns an error with guidance if:
// - The file doesn't exist
// - The file cannot be read
// - The JSON is invalid
// - The name field is missing or empty
func FindWorkspaceConfig(directory string) (string, error) {
	// Ensure directory path is absolute
	absDir, err := filepath.Abs(directory)
	if err != nil {
		return "", fmt.Errorf("failed to resolve directory path: %w", err)
	}

	// Check if directory exists
	if _, statErr := os.Stat(absDir); os.IsNotExist(statErr) {
		return "", fmt.Errorf("directory does not exist: %s", absDir)
	}

	// Construct the workspace file path
	workspaceFile := filepath.Join(absDir, WorkspaceFileName)

	// Check if workspace file exists
	if _, statErr := os.Stat(workspaceFile); os.IsNotExist(statErr) {
		return "", fmt.Errorf(`workspace configuration not found

The directory '%s' does not contain a %s file.

To create a workspace configuration:

1. Create a file named '%s' in your project directory
2. Add the following JSON content:

{
  "name": "YourApplicationProfileName"
}

Replace "YourApplicationProfileName" with the exact name of your Veracode application profile.

Example:
{
  "name": "MyApp-Production"
}

Note: The application profile name must match exactly as it appears in Veracode Platform.`,
			absDir, WorkspaceFileName, WorkspaceFileName)
	}

	// Read the workspace file
	data, err := os.ReadFile(workspaceFile) // nolint:gosec // Intentional workspace config file read
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w\n\nEnsure the file has read permissions.", WorkspaceFileName, err)
	}

	// Parse JSON
	var config WorkspaceConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return "", fmt.Errorf(`invalid JSON in %s: %w

The workspace configuration file must be valid JSON with this format:

{
  "name": "YourApplicationProfileName"
}

Current file location: %s`, WorkspaceFileName, err, workspaceFile)
	}

	// Validate that name is not empty
	if config.Name == "" {
		return "", fmt.Errorf(`missing or empty "name" field in %s

The workspace configuration must include an application profile name:

{
  "name": "YourApplicationProfileName"
}

Current file location: %s`, WorkspaceFileName, workspaceFile)
	}

	return config.Name, nil
}

// FindWorkspaceConfigInCurrentDir is a convenience function that searches for
// .veracode-workspace.json in the current working directory
func FindWorkspaceConfigInCurrentDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}
	return FindWorkspaceConfig(cwd)
}
