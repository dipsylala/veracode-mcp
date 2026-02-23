package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/dipsylala/veracode-mcp/api"
	"github.com/dipsylala/veracode-mcp/workspace"
)

const PipelineScanToolName = "pipeline-scan"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(PipelineScanToolName, handlePipelineScan)
}

// PipelineScanRequest represents the parsed parameters for pipeline-scan
type PipelineScanRequest struct {
	ApplicationPath string
	AppProfile      string
	Filename        string
}

// appInfo holds resolved application details used during the scan
type appInfo struct {
	NumericID  string
	PolicyName string
	Warnings   []string
}

// parsePipelineScanRequest extracts and validates parameters from the raw args map
func parsePipelineScanRequest(args map[string]interface{}) (*PipelineScanRequest, error) {
	req := &PipelineScanRequest{}

	// Extract required fields
	var err error
	req.ApplicationPath, err = extractRequiredString(args, "application_path")
	if err != nil {
		return nil, err
	}

	// Extract optional app profile
	req.AppProfile, _ = extractOptionalString(args, "app_profile")

	// Extract optional filename
	if filename, ok := extractOptionalString(args, "filename"); ok && filename != "" {
		// Check if filename is just a name (no directory separators)
		if filepath.Base(filename) == filename {
			// Just a filename - prepend temp .veracode/packaging directory
			req.Filename = filepath.Join(veracodeWorkDir(req.ApplicationPath, "packaging"), filename)
		} else {
			// Full or relative path - use as-is
			req.Filename = filename
		}
	}

	return req, nil
}

// handlePipelineScan performs a local static scan using Veracode Pipeline Scanner
func handlePipelineScan(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parsePipelineScanRequest(args)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}

	outputDir, err := prepareOutputDir(req.ApplicationPath)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}

	scanTarget, err := resolveScanTarget(req)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}

	// Resolve application info (numeric ID for --app-id, policy GUID for policy fetch)
	info := resolveAppInfo(ctx, req.ApplicationPath, req.AppProfile)

	// Fetch and save policy if we have a name (best-effort; does not block the scan)
	var savedPolicyPath string
	if info.PolicyName != "" {
		savedPolicyPath = fetchAndSavePolicy(ctx, outputDir, info.PolicyName)
	}

	// Build pipeline scan command
	timestamp := time.Now().Format("20060102-150405")
	resultsFile := filepath.Join(outputDir, fmt.Sprintf("results-%s.json", timestamp))
	filteredResultsFile := filepath.Join(outputDir, fmt.Sprintf("filtered-results-%s.json", timestamp))
	logFile := filepath.Join(outputDir, fmt.Sprintf("scan-%s.log", timestamp))

	cmdArgs := []string{
		"static", "scan", scanTarget,
		"--results-file", resultsFile,
		"--filtered-json-output-file", filteredResultsFile,
		"-v",
	}
	if info.NumericID != "" {
		cmdArgs = append(cmdArgs, "--app-id", info.NumericID)
	}
	if savedPolicyPath != "" {
		cmdArgs = append(cmdArgs, "--policy-file", savedPolicyPath)
	}

	pid, err := launchScan(ctx, cmdArgs, outputDir, resultsFile, logFile)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}

	// Build optional warnings section
	var warningsSection string
	if len(info.Warnings) > 0 {
		warningsSection = "\n⚠ Warnings:\n"
		for _, w := range info.Warnings {
			warningsSection += fmt.Sprintf("- %s\n", w)
		}
	}

	responseText := fmt.Sprintf(`Veracode Pipeline Static Scan - Started
============================

Application Path: %s
Scan Target: %s
PID: %d
Results File: %s
Filtered Results File: %s
Log File: %s
%s
Command executed:
veracode %s

✓ Pipeline scan started successfully in background

Next steps:
- Use pipeline-status tool to check scan progress
- Results will be available in: %s
- Filtered results will be available in: %s
- Log output will be in: %s
`, req.ApplicationPath, scanTarget, pid, resultsFile, filteredResultsFile, logFile, warningsSection, strings.Join(cmdArgs, " "), resultsFile, filteredResultsFile, logFile)

	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": responseText,
		}},
	}, nil
}

// prepareOutputDir creates the pipeline output directory, cleans up prior run artifacts,
// and validates that the application path exists.
func prepareOutputDir(applicationPath string) (string, error) {
	if _, err := os.Stat(applicationPath); os.IsNotExist(err) {
		return "", fmt.Errorf("application path does not exist: %s", applicationPath)
	} else if err != nil {
		return "", fmt.Errorf("failed to access application path: %v", err)
	}

	outputDir := veracodeWorkDir(applicationPath, "pipeline")
	if err := os.MkdirAll(outputDir, 0750); err != nil {
		return "", fmt.Errorf("failed to create output directory: %v", err)
	}
	if err := cleanupPipelineDirectory(outputDir); err != nil {
		return "", fmt.Errorf("failed to cleanup pipeline directory: %v", err)
	}
	return outputDir, nil
}

// resolveScanTarget returns the absolute path of the file to scan.
// It uses req.Filename if set, otherwise picks the largest file in .veracode/packaging.
func resolveScanTarget(req *PipelineScanRequest) (string, error) {
	if req.Filename != "" {
		if _, err := os.Stat(req.Filename); os.IsNotExist(err) {
			return "", fmt.Errorf("specified file does not exist: %s", req.Filename)
		}
		return req.Filename, nil
	}
	packagingDir := veracodeWorkDir(req.ApplicationPath, "packaging")
	target, err := findLargestFile(packagingDir)
	if err != nil {
		return "", fmt.Errorf("failed to find file to scan: %v. Either specify a filename parameter or ensure .veracode/packaging contains packaged files", err)
	}
	return target, nil
}

// launchScan starts the veracode static scan process, writes the log header and PID file,
// and returns the process PID on success.
func launchScan(ctx context.Context, cmdArgs []string, outputDir, resultsFile, logFile string) (int, error) {
	// #nosec G204 -- veracode command is hardcoded, only arguments are user-controlled and validated
	cmd := exec.CommandContext(ctx, "veracode", cmdArgs...)

	// #nosec G304 -- logFile is constructed from validated outputDir and timestamp, not user input
	logFileHandle, err := os.Create(logFile)
	if err != nil {
		return 0, fmt.Errorf("failed to create log file: %v", err)
	}
	// Intentionally not closing logFileHandle here: the async child process continues
	// writing to it after this function returns. The OS will close the handle on process exit.

	fmt.Fprintf(logFileHandle, "Command: veracode %s\n\n", strings.Join(cmdArgs, " "))
	cmd.Stdout = logFileHandle
	cmd.Stderr = logFileHandle

	if err := cmd.Start(); err != nil {
		return 0, fmt.Errorf("failed to start pipeline scan: %v", err)
	}

	pid := cmd.Process.Pid
	pidInfo := fmt.Sprintf(`{"pid":%d,"results_file":"%s","log_file":"%s"}`,
		pid, filepath.ToSlash(resultsFile), filepath.ToSlash(logFile))
	pidFile := filepath.Join(outputDir, "pipeline.pid")
	// #nosec G306 -- PID file needs to be readable by status checker, contains process info
	if err := os.WriteFile(pidFile, []byte(pidInfo), 0644); err != nil {
		_ = cmd.Process.Kill()
		return 0, fmt.Errorf("failed to write PID file: %v", err)
	}
	return pid, nil
}

// resolveAppInfo looks up the numeric Veracode application ID and policy details.
// It uses app_profile if provided, otherwise falls back to .veracode-workspace.json.
// All fields are best-effort: warnings are returned for informational failures.
func resolveAppInfo(ctx context.Context, applicationPath, appProfile string) appInfo {
	info := appInfo{}

	name := appProfile
	if name == "" {
		var err error
		name, err = workspace.FindWorkspaceConfig(applicationPath)
		if err != nil {
			info.Warnings = append(info.Warnings, "No .veracode-workspace.json found. This scan did not use an application profile or a non-standard policy.")
			return info
		}
	}

	client, err := api.NewClient()
	if err != nil {
		log.Printf("pipeline-scan: skipping app lookup, failed to create API client: %v", err)
		info.Warnings = append(info.Warnings, fmt.Sprintf("Could not connect to Veracode API: %v", err))
		return info
	}

	application, err := client.GetApplicationByName(ctx, name)
	if err != nil || application == nil {
		log.Printf("pipeline-scan: skipping app lookup, could not resolve app '%s': %v", name, err)
		info.Warnings = append(info.Warnings, fmt.Sprintf("Application '%s' not found in Veracode. This scan did not use an application profile or a non-standard policy.", name))
		return info
	}

	if application.Id != nil {
		info.NumericID = fmt.Sprintf("%d", *application.Id)
	}
	if application.Profile != nil {
		for _, p := range application.Profile.GetPolicies() {
			if p.Name != nil {
				info.PolicyName = *p.Name
				break // use the first policy
			}
		}
	}
	return info
}

// fetchAndSavePolicy retrieves the policy by name and saves it to outputDir/policy.json.
// Returns the saved file path on success, or empty string on failure.
func fetchAndSavePolicy(ctx context.Context, outputDir, policyName string) string {
	client, err := api.NewClient()
	if err != nil {
		log.Printf("pipeline-scan: skipping policy fetch, failed to create API client: %v", err)
		return ""
	}

	policyVersion, err := client.GetPolicy(ctx, policyName)
	if err != nil {
		log.Printf("pipeline-scan: skipping policy save, failed to fetch policy '%s': %v", policyName, err)
		return ""
	}

	data, err := json.MarshalIndent(policyVersion, "", "  ")
	if err != nil {
		log.Printf("pipeline-scan: failed to marshal policy: %v", err)
		return ""
	}

	policyFile := filepath.Join(outputDir, "policy.json")
	// #nosec G306 -- policy file is in a controlled output directory
	if err := os.WriteFile(policyFile, data, 0644); err != nil {
		log.Printf("pipeline-scan: failed to save policy to %s: %v", policyFile, err)
		return ""
	}
	log.Printf("pipeline-scan: saved policy '%s' to %s", policyName, policyFile)
	return policyFile
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

		info, err := entry.Info()
		if err != nil {
			continue
		}

		if info.Size() > largestSize {
			largestSize = info.Size()
			largestFile = filepath.Join(dir, entry.Name())
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
