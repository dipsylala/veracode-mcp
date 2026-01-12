package tools

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

const PipelineScanToolName = "pipeline-static-scan"

// Auto-register this tool when the package is imported
func init() {
	RegisterTool(PipelineScanToolName, func() ToolImplementation {
		return NewPipelineScanTool()
	})
}

// PipelineScanTool provides the pipeline-static-scan tool
type PipelineScanTool struct{}

// NewPipelineScanTool creates a new pipeline scan tool
func NewPipelineScanTool() *PipelineScanTool {
	return &PipelineScanTool{}
}

// Initialize sets up the tool
func (t *PipelineScanTool) Initialize() error {
	log.Printf("Initializing tool: %s", PipelineScanToolName)
	return nil
}

// RegisterHandlers registers the pipeline scan handler
func (t *PipelineScanTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", PipelineScanToolName)
	registry.RegisterHandler(PipelineScanToolName, t.handlePipelineScan)
	return nil
}

// Shutdown cleans up tool resources
func (t *PipelineScanTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", PipelineScanToolName)
	return nil
}

// PipelineScanRequest represents the parsed parameters for pipeline-static-scan
type PipelineScanRequest struct {
	ApplicationPath string
	Filename        string
	Verbose         bool
}

// parsePipelineScanRequest extracts and validates parameters from the raw args map
func parsePipelineScanRequest(args map[string]interface{}) (*PipelineScanRequest, error) {
	req := &PipelineScanRequest{}

	// Extract application_path (required)
	appPath, ok := args["application_path"].(string)
	if !ok || appPath == "" {
		return nil, fmt.Errorf("application_path is required and must be a non-empty string")
	}
	req.ApplicationPath = appPath

	// Extract filename (optional)
	if filename, ok := args["filename"].(string); ok {
		req.Filename = filename
	}

	// Extract verbose (optional)
	if verbose, ok := args["verbose"].(bool); ok {
		req.Verbose = verbose
	}

	return req, nil
}

// handlePipelineScan performs a local static scan using Veracode Pipeline Scanner
func (t *PipelineScanTool) handlePipelineScan(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parsePipelineScanRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Validate application path exists
	_, err = os.Stat(req.ApplicationPath)
	if os.IsNotExist(err) {
		return map[string]interface{}{
			"error": fmt.Sprintf("Application path does not exist: %s", req.ApplicationPath),
		}, nil
	}
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to access application path: %v", err),
		}, nil
	}

	// Create output directory for scan results
	outputDir := filepath.Join(req.ApplicationPath, ".veracode_pipeline")
	err = os.MkdirAll(outputDir, 0750)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to create output directory: %v", err),
		}, nil
	}

	// Determine which file to scan
	var scanTarget string
	if req.Filename != "" {
		// Use the provided filename
		scanTarget = req.Filename
		// If it's not an absolute path, assume it's in the application path
		if !filepath.IsAbs(scanTarget) {
			scanTarget = filepath.Join(req.ApplicationPath, scanTarget)
		}

		// Validate the file exists
		var fileInfo os.FileInfo
		fileInfo, err = os.Stat(scanTarget)
		if os.IsNotExist(err) {
			return map[string]interface{}{
				"error": fmt.Sprintf("Specified file does not exist: %s", scanTarget),
			}, nil
		}
		_ = fileInfo // Unused but needed for os.Stat call
	} else {
		// Find the largest file in .veracode_packaging
		packagingDir := filepath.Join(req.ApplicationPath, ".veracode_packaging")
		scanTarget, err = findLargestFile(packagingDir)
		if err != nil {
			return map[string]interface{}{
				"error": fmt.Sprintf("Failed to find file to scan: %v. Either specify a filename parameter or ensure .veracode_packaging contains packaged files.", err),
			}, nil
		}
	}

	// Build pipeline scan command
	resultsFile := filepath.Join(outputDir, fmt.Sprintf("results-%s.json", time.Now().Format("20060102-150405")))

	cmdArgs := []string{
		"static",
		"scan",
		scanTarget,
		"--results-file", resultsFile,
	}

	// Add verbose flag if requested
	if req.Verbose {
		cmdArgs = append(cmdArgs, "-v")
	}

	// #nosec G204 -- veracode command is hardcoded, only arguments are user-controlled and validated
	cmd := exec.CommandContext(ctx, "veracode", cmdArgs...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

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

	// Build response
	return buildPipelineScanResponse(req, scanTarget, resultsFile, exitCode, duration, stdout, stderr), nil
}

// findLargestFile finds the largest file in the specified directory
func findLargestFile(dir string) (string, error) {
	// Check if directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return "", fmt.Errorf("directory does not exist: %s", dir)
	}

	var largestFile string
	var largestSize int64

	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(dir, entry.Name())
		info, err := entry.Info()
		if err != nil {
			continue
		}

		if info.Size() > largestSize {
			largestSize = info.Size()
			largestFile = filePath
		}
	}

	if largestFile == "" {
		return "", fmt.Errorf("no files found in directory: %s", dir)
	}

	return largestFile, nil
}

// buildPipelineScanResponse constructs the response based on command execution results
func buildPipelineScanResponse(req *PipelineScanRequest, scanTarget string, resultsFile string, exitCode int, duration time.Duration, stdout, stderr bytes.Buffer) map[string]interface{} {
	// Interpret the exit code
	cliInfo := InterpretVeracodeExitCode(exitCode)

	// Build base response information
	baseInfo := fmt.Sprintf(`Veracode Pipeline Static Scan
============================

Application Path: %s
Scan Target: %s
Results File: %s
Duration: %v
Exit Code: %d

Command executed:
veracode static scan %s --results-file %s

`, req.ApplicationPath, scanTarget, resultsFile, duration, exitCode, scanTarget, resultsFile)

	// Customize next steps for pipeline scan context
	nextSteps := cliInfo.NextSteps
	switch exitCode {
	case 0:
		nextSteps = fmt.Sprintf("Next steps:\n- Review scan results in: %s\n- Analyze identified vulnerabilities\n- Integrate findings into your development workflow", resultsFile)
	case 3:
		nextSteps = fmt.Sprintf("Next steps:\n- Review scan results in: %s\n- Check policy violations\n- Address critical vulnerabilities before deployment", resultsFile)
	case 4:
		nextSteps = "Next steps:\n- Review warnings in scan output\n- Verify build artifacts are available\n- Check if source files were properly compiled"
	}

	responseText := baseInfo + fmt.Sprintf("%s %s\n\n", cliInfo.Icon, cliInfo.Message)

	if stdout.Len() > 0 {
		responseText += fmt.Sprintf("\nOutput:\n%s\n", stdout.String())
	}

	if stderr.Len() > 0 {
		responseText += fmt.Sprintf("\nError output:\n%s\n", stderr.String())
	}

	responseText += fmt.Sprintf("\n%s", nextSteps)

	// Check if results file was created
	if _, err := os.Stat(resultsFile); err == nil {
		responseText += fmt.Sprintf("\n\n✓ Results file created successfully: %s", resultsFile)
	}

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
