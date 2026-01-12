# Workspace Package

The `workspace` package provides functionality for finding and parsing `.veracode-workspace.json` configuration files in project directories.

## Overview

This package allows tools to automatically discover the Veracode application profile name associated with a project by reading a simple JSON configuration file.

## Usage

### Basic Usage

```go
import "github.com/dipsylala/veracodemcp-go/workspace"

// Find workspace config in a specific directory
appName, err := workspace.FindWorkspaceConfig("/path/to/project")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Application: %s\n", appName)
```

### Current Directory

```go
// Find workspace config in current working directory
appName, err := workspace.FindWorkspaceConfigInCurrentDir()
if err != nil {
    log.Fatal(err)
}
```

## Workspace File Format

Create a file named `.veracode-workspace.json` in your project root:

```json
{
  "name": "YourApplicationProfileName"
}
```

### Example

```json
{
  "name": "MyApp-Production"
}
```

The `name` field must match the exact name of your Veracode application profile.

## Error Handling

The package provides helpful error messages for common issues:

### File Not Found

```
workspace configuration not found

The directory '/path/to/project' does not contain a .veracode-workspace.json file.

To create a workspace configuration:

1. Create a file named '.veracode-workspace.json' in your project directory
2. Add the following JSON content:

{
  "name": "YourApplicationProfileName"
}
...
```

### Invalid JSON

```
invalid JSON in .veracode-workspace.json: ...

The workspace configuration file must be valid JSON with this format:

{
  "name": "YourApplicationProfileName"
}
...
```

### Missing Name Field

```
missing or empty "name" field in .veracode-workspace.json

The workspace configuration must include an application profile name:

{
  "name": "YourApplicationProfileName"
}
...
```

## Features

- ✅ Searches for `.veracode-workspace.json` in specified directory
- ✅ Supports both absolute and relative paths
- ✅ Validates JSON format
- ✅ Validates required `name` field
- ✅ Provides helpful error messages with examples
- ✅ Thread-safe (no global state)

## API Reference

### Functions

#### `FindWorkspaceConfig(directory string) (string, error)`

Searches for `.veracode-workspace.json` in the given directory and returns the application profile name.

**Parameters:**
- `directory`: Path to the directory to search (absolute or relative)

**Returns:**
- `string`: The application profile name from the workspace file
- `error`: Error with guidance if file not found or invalid

**Example:**
```go
name, err := workspace.FindWorkspaceConfig("./my-project")
```

#### `FindWorkspaceConfigInCurrentDir() (string, error)`

Convenience function that searches for `.veracode-workspace.json` in the current working directory.

**Returns:**
- `string`: The application profile name from the workspace file
- `error`: Error with guidance if file not found or invalid

**Example:**
```go
name, err := workspace.FindWorkspaceConfigInCurrentDir()
```

### Constants

#### `WorkspaceFileName`

```go
const WorkspaceFileName = ".veracode-workspace.json"
```

The standard name for workspace configuration files.

### Types

#### `WorkspaceConfig`

```go
type WorkspaceConfig struct {
    Name string `json:"name"`
}
```

Represents the structure of `.veracode-workspace.json`.

## Integration Examples

### With MCP Tools

```go
// In a tool handler
func handleGetFindings(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Get directory from params or use current dir
    directory := params["directory"].(string)
    
    // Find workspace config
    appName, err := workspace.FindWorkspaceConfig(directory)
    if err != nil {
        return nil, err
    }
    
    // Use appName to call Veracode API
    findings, err := client.GetFindingsByAppName(ctx, appName)
    // ...
}
```

### With CLI Tools

```go
func main() {
    // Try to find workspace config
    appName, err := workspace.FindWorkspaceConfigInCurrentDir()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Printf("Scanning application: %s\n", appName)
    // ... continue with scan
}
```

## Testing

Run the test suite:

```bash
go test -v ./workspace
```

The test suite includes:
- ✅ Success case with valid config
- ✅ File not found error handling
- ✅ Invalid JSON error handling
- ✅ Empty name validation
- ✅ Missing name field validation
- ✅ Non-existent directory handling
- ✅ Relative path support
- ✅ Special characters in names
- ✅ Current directory convenience function

All tests use temporary directories and clean up automatically.

## Best Practices

1. **Place at project root**: Put `.veracode-workspace.json` at the root of your project
2. **Exact name match**: Use the exact application profile name from Veracode Platform
3. **Version control**: Commit the workspace file to your repository
4. **Case sensitive**: Application names are case-sensitive
5. **No comments**: JSON doesn't support comments; keep the file simple

## Example Project Structure

```
my-project/
├── .veracode-workspace.json   ← Workspace config
├── src/
├── tests/
└── README.md
```

## Error Recovery

If you see an error about missing workspace file:

1. Create `.veracode-workspace.json` in your project root
2. Add the application name from Veracode Platform
3. Verify the JSON is valid (use a JSON validator)
4. Ensure the name field is not empty

## Future Enhancements

Potential additions for future versions:
- Support for additional configuration (region, credentials path, etc.)
- Hierarchical search (search parent directories)
- Multiple workspace file formats (YAML, TOML)
- Workspace file validation command
