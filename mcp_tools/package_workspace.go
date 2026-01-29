package mcp_tools

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const PackageWorkspaceToolName = "package-workspace"

// Auto-register this tool when the package is imported
func init() {
	RegisterTool(PackageWorkspaceToolName, func() ToolImplementation {
		return NewPackageWorkspaceTool()
	})
}

// PackageWorkspaceTool provides the package-workspace tool
type PackageWorkspaceTool struct{}

// NewPackageWorkspaceTool creates a new package workspace tool
func NewPackageWorkspaceTool() *PackageWorkspaceTool {
	return &PackageWorkspaceTool{}
}

// Initialize sets up the tool
func (t *PackageWorkspaceTool) Initialize() error {
	log.Printf("Initializing tool: %s", PackageWorkspaceToolName)
	return nil
}

// RegisterHandlers registers the package workspace handler
func (t *PackageWorkspaceTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", PackageWorkspaceToolName)
	registry.RegisterHandler(PackageWorkspaceToolName, t.handlePackageWorkspace)
	return nil
}

// Shutdown cleans up tool resources
func (t *PackageWorkspaceTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", PackageWorkspaceToolName)
	return nil
}

// PackageWorkspaceRequest represents the parsed parameters for package-workspace
type PackageWorkspaceRequest struct {
	ApplicationPath string
	Verbose         bool
	LogToFile       bool
}

// parsePackageWorkspaceRequest extracts and validates parameters from the raw args map
func parsePackageWorkspaceRequest(args map[string]interface{}) (*PackageWorkspaceRequest, error) {
	req := &PackageWorkspaceRequest{}

	// Extract application_path (required)
	appPath, ok := args["application_path"].(string)
	if !ok || appPath == "" {
		return nil, fmt.Errorf("application_path is required and must be a non-empty string")
	}
	req.ApplicationPath = appPath

	// Extract verbose (optional)
	if verbose, ok := args["verbose"].(bool); ok {
		req.Verbose = verbose
	}

	// Extract logToFile (optional)
	if logToFile, ok := args["logToFile"].(bool); ok {
		req.LogToFile = logToFile
	}

	return req, nil
}

// handlePackageWorkspace packages the workspace for Veracode scanning
func (t *PackageWorkspaceTool) handlePackageWorkspace(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parsePackageWorkspaceRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Validate and prepare directories
	outputDir, err := validateAndPrepareDirectories(req.ApplicationPath)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Clean up non-log files from output directory
	err = cleanupOutputDirectory(outputDir)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to clean output directory: %v", err),
		}, nil
	}

	// Build and execute packaging command
	exitCode, duration, stdout, stderr, logFile, err := executePackagingCommand(ctx, req, outputDir)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Build and return response
	return buildPackagingResponse(req, outputDir, exitCode, duration, stdout, stderr, logFile), nil
}

// validateAndPrepareDirectories validates the application path and creates the output directory
func validateAndPrepareDirectories(applicationPath string) (string, error) {
	// Validate application path exists
	_, err := os.Stat(applicationPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("Application path does not exist: %s", applicationPath)
	}
	if err != nil {
		return "", fmt.Errorf("Failed to access application path: %v", err)
	}

	// Create output directory (.veracode_packaging)
	outputDir := filepath.Join(applicationPath, ".veracode_packaging")
	err = os.MkdirAll(outputDir, 0750)
	if err != nil {
		return "", fmt.Errorf("Failed to create output directory: %v", err)
	}

	return outputDir, nil
}

// cleanupOutputDirectory removes all non-log files from the output directory
func cleanupOutputDirectory(outputDir string) error {
	// Read directory contents
	entries, err := os.ReadDir(outputDir)
	if err != nil {
		// If directory doesn't exist or can't be read, that's fine
		return nil
	}

	// Remove all files that are not log files
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		// Keep log files (those ending with .log)
		if filepath.Ext(filename) == ".log" {
			continue
		}

		// Remove non-log files
		filePath := filepath.Join(outputDir, filename)
		if err := os.Remove(filePath); err != nil {
			// Log the error but continue with other files
			log.Printf("Warning: Failed to remove file %s: %v", filePath, err)
		}
	}

	return nil
}

// executePackagingCommand builds and executes the Veracode packaging command
func executePackagingCommand(ctx context.Context, req *PackageWorkspaceRequest, outputDir string) (int, time.Duration, bytes.Buffer, bytes.Buffer, *os.File, error) {
	// Build command arguments
	cmdArgs := []string{
		"package",
		"-s",              // Include source files
		"-t", "directory", // Target type is directory
		"-a",                      // Trust the source directory
		"-s", req.ApplicationPath, // Source directory
		"-o", outputDir, // Output directory
	}

	// Add verbose flag if requested
	if req.Verbose {
		cmdArgs = append(cmdArgs, "-v")
	}

	// #nosec G204 -- veracode command is hardcoded, only arguments are user-controlled and validated
	cmd := exec.CommandContext(ctx, "veracode", cmdArgs...)

	var stdout, stderr bytes.Buffer
	var logFile *os.File
	var err error

	// Set up logging
	if req.LogToFile {
		timestamp := time.Now().Format("20060102-150405")
		logFilePath := filepath.Join(outputDir, fmt.Sprintf("veracode-packaging-%s.log", timestamp))
		// #nosec G304 -- logFilePath is constructed from validated outputDir and timestamp, not user input
		logFile, err = os.Create(logFilePath)
		if err != nil {
			return 0, 0, stdout, stderr, nil, fmt.Errorf("Failed to create log file: %v", err)
		}

		// Write to log file
		cmd.Stdout = logFile
		cmd.Stderr = logFile
	} else {
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
	}

	// Execute the command
	startTime := time.Now()
	err = cmd.Run()
	duration := time.Since(startTime)

	// Extract exit code
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		}
	}

	return exitCode, duration, stdout, stderr, logFile, nil
}

// buildPackagingResponse constructs the response based on command execution results
func buildPackagingResponse(req *PackageWorkspaceRequest, outputDir string, exitCode int, duration time.Duration, stdout, stderr bytes.Buffer, logFile *os.File) map[string]interface{} {
	// Interpret the exit code
	cliInfo := InterpretVeracodeExitCode(exitCode)

	// Build command arguments for display
	cmdArgs := []string{
		"package",
		"-s", "-t", "directory", "-a",
		"-s", req.ApplicationPath,
		"-o", outputDir,
	}
	if req.Verbose {
		cmdArgs = append(cmdArgs, "-v")
	}

	// Build base response information
	baseInfo := fmt.Sprintf(`Veracode Workspace Packaging
============================

Application Path: %s
Output Directory: %s
Duration: %v
Exit Code: %d

Command executed:
veracode %s

`, req.ApplicationPath, outputDir, duration, exitCode, joinArgs(cmdArgs))

	// Customize next steps for packaging context
	nextSteps := cliInfo.NextSteps
	switch exitCode {
	case 0:
		nextSteps = fmt.Sprintf("Next steps:\n- Review packaging results in: %s\n- Upload the packaged artifact to Veracode\n- Submit for security scanning", outputDir)
	case 3:
		nextSteps = fmt.Sprintf("Next steps:\n- Review packaging results in: %s\n- Check policy violations\n- Address policy issues before submission", outputDir)
	}

	responseText := baseInfo + fmt.Sprintf("%s %s\n\n", cliInfo.Icon, cliInfo.Message)

	if req.LogToFile && logFile != nil {
		responseText += fmt.Sprintf("Log file: %s\n", logFile.Name())
	}

	if stdout.Len() > 0 {
		responseText += fmt.Sprintf("\nOutput:\n%s\n", stdout.String())
	}

	if stderr.Len() > 0 {
		responseText += fmt.Sprintf("\nError output:\n%s\n", stderr.String())
	}

	responseText += fmt.Sprintf("\n%s", nextSteps)

	// Return as error for truly failing exit codes, but as content for warnings
	if exitCode > 0 && !cliInfo.IsWarning {
		return map[string]interface{}{
			"error": responseText,
		}
	}

	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": responseText,
		}},
	}
}

// joinArgs joins command arguments for display purposes
func joinArgs(args []string) string {
	result := ""
	for i, arg := range args {
		if i > 0 {
			result += " "
		}
		// Quote arguments that contain spaces
		if len(arg) > 0 && (arg[0] == '-' || filepath.IsAbs(arg)) {
			result += arg
		} else {
			result += fmt.Sprintf("%q", arg)
		}
	}
	return result
}
