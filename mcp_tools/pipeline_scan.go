package mcp_tools

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const PipelineScanToolName = "pipeline-scan"

// Auto-register this tool when the package is imported
func init() {
	RegisterSimpleTool(PipelineScanToolName, handlePipelineScan)
}

// PipelineScanRequest represents the parsed parameters for pipeline-scan
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

	// Extract filename (optional) - if just a filename (no path), look in .veracode/packaging
	if filename, ok := args["filename"].(string); ok && filename != "" {
		// Check if filename is just a name (no directory separators)
		if filepath.Base(filename) == filename {
			// Just a filename - prepend .veracode/packaging directory
			req.Filename = filepath.Join(appPath, ".veracode", "packaging", filename)
		} else {
			// Full or relative path - use as-is
			req.Filename = filename
		}
	}

	// Extract verbose (optional)
	if verbose, ok := args["verbose"].(bool); ok {
		req.Verbose = verbose
	}

	return req, nil
}

// handlePipelineScan performs a local static scan using Veracode Pipeline Scanner
func handlePipelineScan(ctx context.Context, args map[string]interface{}) (interface{}, error) {
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
	outputDir := filepath.Join(req.ApplicationPath, ".veracode", "pipeline")
	err = os.MkdirAll(outputDir, 0750)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to create output directory: %v", err),
		}, nil
	}

	// Clean up any existing files in the pipeline directory before starting a new scan
	err = cleanupPipelineDirectory(outputDir)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to cleanup pipeline directory: %v", err),
		}, nil
	}

	// Determine which file to scan
	var scanTarget string
	if req.Filename != "" {
		// Use the provided filename (already resolved in parsePipelineScanRequest)
		scanTarget = req.Filename

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
		// Find the largest file in .veracode/packaging
		packagingDir := filepath.Join(req.ApplicationPath, ".veracode", "packaging")
		scanTarget, err = findLargestFile(packagingDir)
		if err != nil {
			return map[string]interface{}{
				"error": fmt.Sprintf("Failed to find file to scan: %v. Either specify a filename parameter or ensure .veracode/packaging contains packaged files.", err),
			}, nil
		}
	}

	// Build pipeline scan command
	resultsFile := filepath.Join(outputDir, fmt.Sprintf("results-%s.json", time.Now().Format("20060102-150405")))
	logFile := filepath.Join(outputDir, fmt.Sprintf("scan-%s.log", time.Now().Format("20060102-150405")))

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

	// Create log file for stdout/stderr
	// #nosec G304 -- logFile is constructed from validated outputDir and timestamp, not user input
	logFileHandle, err := os.Create(logFile)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to create log file: %v", err),
		}, nil
	}
	defer logFileHandle.Close()

	// Redirect stdout and stderr to log file
	cmd.Stdout = logFileHandle
	cmd.Stderr = logFileHandle

	// Start the command asynchronously
	err = cmd.Start()
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to start pipeline scan: %v", err),
		}, nil
	}

	// Write PID info to file for status checking (includes PID, results file, log file)
	pidFile := filepath.Join(outputDir, "pipeline.pid")
	pidInfo := fmt.Sprintf("{\"pid\":%d,\"results_file\":\"%s\",\"log_file\":\"%s\"}",
		cmd.Process.Pid,
		filepath.ToSlash(resultsFile),
		filepath.ToSlash(logFile))
	// #nosec G306 -- PID file needs to be readable by status checker, contains process info
	err = os.WriteFile(pidFile, []byte(pidInfo), 0644)
	if err != nil {
		// Kill the process if we can't write PID
		_ = cmd.Process.Kill()
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to write PID file: %v", err),
		}, nil
	}

	// Return immediately with scan started status
	responseText := fmt.Sprintf(`Veracode Pipeline Static Scan - Started
============================

Application Path: %s
Scan Target: %s
Results File: %s
Log File: %s

Command executed:
veracode static scan %s --results-file %s

✓ Pipeline scan started successfully in background

Next steps:
- Use pipeline-status tool to check scan progress
- Results will be available in: %s
- Log output will be in: %s
`, req.ApplicationPath, scanTarget, resultsFile, logFile, scanTarget, resultsFile, resultsFile, logFile)

	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": responseText,
		}},
	}, nil
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

// cleanupPipelineDirectory removes all files from the pipeline directory before starting a new scan
func cleanupPipelineDirectory(outputDir string) error {
	// Read directory contents
	entries, err := os.ReadDir(outputDir)
	if err != nil {
		// If directory doesn't exist or can't be read, that's fine
		return nil
	}

	// Remove all files in the directory
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(outputDir, entry.Name())
		if err := os.Remove(filePath); err != nil {
			// Log the error but continue with other files
			log.Printf("Warning: Failed to remove file %s: %v", filePath, err)
		}
	}

	return nil
}
