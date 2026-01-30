package mcp_tools

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const RunSCAScanToolName = "run-sca-scan"

// Auto-register this tool when the package is imported
func init() {
	RegisterSimpleTool(RunSCAScanToolName, handleRunSCAScan)
}

// RunSCAScanRequest represents the parsed parameters for run-sca-scan
type RunSCAScanRequest struct {
	ApplicationPath string
}

// parseRunSCAScanRequest extracts and validates parameters from the raw args map
func parseRunSCAScanRequest(args map[string]interface{}) (*RunSCAScanRequest, error) {
	req := &RunSCAScanRequest{}

	// Extract application_path (required)
	appPath, ok := args["application_path"].(string)
	if !ok || appPath == "" {
		return nil, fmt.Errorf("application_path is required and must be a non-empty string")
	}
	req.ApplicationPath = appPath

	return req, nil
}

// handleRunSCAScan runs a Software Composition Analysis scan on the workspace
func handleRunSCAScan(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseRunSCAScanRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Validate and prepare directories
	outputDir, outputFile, err := validateAndPrepareSCADirectories(req.ApplicationPath)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Build and execute SCA scan command
	exitCode, duration, stdout, stderr, err := executeSCAScanCommand(ctx, req, outputFile)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Build and return response
	return buildSCAScanResponse(req, outputDir, outputFile, exitCode, duration, stdout, stderr), nil
}

// validateAndPrepareSCADirectories validates the application path and creates the output directory
func validateAndPrepareSCADirectories(applicationPath string) (string, string, error) {
	// Validate application path exists
	_, err := os.Stat(applicationPath)
	if os.IsNotExist(err) {
		return "", "", fmt.Errorf("Application path does not exist: %s", applicationPath)
	}
	if err != nil {
		return "", "", fmt.Errorf("Failed to access application path: %v", err)
	}

	// Create output directory (.veracode/sca)
	// MkdirAll creates the directory if it doesn't exist, or does nothing if it already exists
	outputDir := filepath.Join(applicationPath, ".veracode", "sca")
	err = os.MkdirAll(outputDir, 0750)
	if err != nil {
		return "", "", fmt.Errorf("Failed to create output directory: %v", err)
	}

	// Define output file path
	outputFile := filepath.Join(outputDir, "veracode.json")

	return outputDir, outputFile, nil
}

// executeSCAScanCommand builds and executes the Veracode SCA scan command
func executeSCAScanCommand(ctx context.Context, req *RunSCAScanRequest, outputFile string) (int, time.Duration, bytes.Buffer, bytes.Buffer, error) {
	// Build command arguments
	cmdArgs := []string{
		"scan",
		"--format", "json",
		"-s", req.ApplicationPath,
		"-o", outputFile,
		"--type", "directory",
	}

	// #nosec G204 -- veracode command is hardcoded, only arguments are user-controlled and validated
	cmd := exec.CommandContext(ctx, "veracode", cmdArgs...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	startTime := time.Now()
	err := cmd.Run()
	duration := time.Since(startTime)

	// Extract exit code
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		}
	}

	return exitCode, duration, stdout, stderr, nil
}

// buildSCAScanResponse constructs the response based on command execution results
func buildSCAScanResponse(req *RunSCAScanRequest, outputDir, outputFile string, exitCode int, duration time.Duration, stdout, stderr bytes.Buffer) map[string]interface{} {
	// Interpret the exit code
	cliInfo := InterpretVeracodeExitCode(exitCode)

	// Build command arguments for display
	cmdArgs := []string{
		"scan",
		"--format", "json",
		"-s", req.ApplicationPath,
		"-o", outputFile,
		"--type", "directory",
	}

	// Build base response information
	baseInfo := fmt.Sprintf(`Veracode SCA Scan
=================

Application Path: %s
Output Directory: %s
Output File: %s
Duration: %v
Exit Code: %d

Command executed:
veracode %s

`, req.ApplicationPath, outputDir, outputFile, duration, exitCode, joinArgs(cmdArgs))

	// Customize next steps for SCA scan context
	nextSteps := cliInfo.NextSteps
	switch exitCode {
	case 0:
		nextSteps = fmt.Sprintf("Next steps:\n- Review SCA scan results in: %s\n- Check for vulnerable dependencies and license issues\n- Review remediation recommendations", outputFile)
	case 3:
		nextSteps = fmt.Sprintf("Next steps:\n- Review SCA scan results in: %s\n- Check policy violations in dependencies\n- Address vulnerable components before deployment", outputFile)
	}

	responseText := baseInfo + fmt.Sprintf("%s %s\n\n", cliInfo.Icon, cliInfo.Message)

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
