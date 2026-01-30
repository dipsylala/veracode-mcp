package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const PipelineStatusToolName = "pipeline-status"

// Auto-register this tool when the package is imported
func init() {
	RegisterSimpleTool(PipelineStatusToolName, handlePipelineStatus)
}

// PipelineStatusRequest represents the parsed parameters for pipeline-status
type PipelineStatusRequest struct {
	ApplicationPath string
}

// parsePipelineStatusRequest extracts and validates parameters from the raw args map
func parsePipelineStatusRequest(args map[string]interface{}) (*PipelineStatusRequest, error) {
	req := &PipelineStatusRequest{}

	// Extract application_path (required)
	appPath, ok := args["application_path"].(string)
	if !ok || appPath == "" {
		return nil, fmt.Errorf("application_path is required and must be a non-empty string")
	}
	req.ApplicationPath = appPath

	return req, nil
}

// handlePipelineStatus checks the status of a running pipeline scan
func handlePipelineStatus(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parsePipelineStatusRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Locate the PID file
	outputDir := filepath.Join(req.ApplicationPath, ".veracode_pipeline")
	pidFile := filepath.Join(outputDir, "pipeline.pid")

	// Check if PID file exists
	// #nosec G304 -- pidFile is constructed from validated outputDir and fixed filename, not user input
	pidData, err := os.ReadFile(pidFile)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]interface{}{
				"content": []map[string]string{{
					"type": "text",
					"text": fmt.Sprintf(`Pipeline Scan Status
==================

Application Path: %s
PID File: %s

❌ No pipeline scan found

No active or recent pipeline scan detected. The PID file does not exist.

To start a new scan, use the pipeline-static-scan tool.
`, req.ApplicationPath, pidFile),
				}},
			}, nil
		}
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to read PID file: %v", err),
		}, nil
	}

	// Parse PID info (JSON format with pid, results_file, log_file)
	var pidInfo struct {
		PID         int    `json:"pid"`
		ResultsFile string `json:"results_file"`
		LogFile     string `json:"log_file"`
	}

	// Try to parse as JSON first
	err = json.Unmarshal(pidData, &pidInfo)
	if err != nil {
		// Fallback: try parsing as plain PID number for backward compatibility
		pidStr := strings.TrimSpace(string(pidData))
		pid, parseErr := strconv.Atoi(pidStr)
		if parseErr != nil {
			return map[string]interface{}{
				"error": fmt.Sprintf("Invalid PID file format: %v", err),
			}, nil
		}
		pidInfo.PID = pid
		pidInfo.ResultsFile = "unknown"
		pidInfo.LogFile = "unknown"
	}

	pid := pidInfo.PID

	// Check if process is running
	isRunning, _ := checkProcessStatus(pid)

	// Build response
	if isRunning {
		responseText := fmt.Sprintf(`Pipeline Scan Status
==================

Application Path: %s
PID: %d
Status: RUNNING ⏳
Results File: %s
Log File: %s

The pipeline scan is currently in progress.

Check again later to see if the scan has completed.
Results will be available in: %s
Log output is being written to: %s
`, req.ApplicationPath, pid, pidInfo.ResultsFile, pidInfo.LogFile, pidInfo.ResultsFile, pidInfo.LogFile)

		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": responseText,
			}},
		}, nil
	}

	// Process has completed
	responseText := fmt.Sprintf(`Pipeline Scan Status
==================

Application Path: %s
PID: %d
Status: COMPLETED ✓

The pipeline scan has finished.

Results File: %s
Log File: %s

Check the log file for detailed output from the scan.
`, req.ApplicationPath, pid, pidInfo.ResultsFile, pidInfo.LogFile)

	// Clean up PID file now that scan is complete
	_ = os.Remove(pidFile)

	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": responseText,
		}},
	}, nil
}

// checkProcessStatus checks if a process is running and returns its exit code if available
func checkProcessStatus(pid int) (isRunning bool, exitCode int) {
	if runtime.GOOS == "windows" {
		return checkProcessStatusWindows(pid)
	}
	return checkProcessStatusUnix(pid)
}

// checkProcessStatusWindows checks process status on Windows
func checkProcessStatusWindows(pid int) (isRunning bool, exitCode int) {
	// On Windows, we use FindProcess and then try Wait with timeout
	// FindProcess always succeeds, so we need another method
	process, err := os.FindProcess(pid)
	if err != nil {
		return false, -1
	}

	// Try to wait with a short timeout to check process state
	// If Wait returns quickly, process has exited
	// If it times out, process is still running
	done := make(chan error, 1)
	go func() {
		_, err := process.Wait()
		done <- err
	}()

	// Give the Wait call a brief moment to complete if process has already exited
	select {
	case <-done:
		// Process has exited
		return false, 0
	case <-time.After(100 * time.Millisecond):
		// Process is still running (Wait is still blocking after timeout)
		return true, 0
	}
}

// checkProcessStatusUnix checks process status on Unix-like systems
func checkProcessStatusUnix(pid int) (isRunning bool, exitCode int) {
	// On Unix, we can send signal 0 to check if process exists
	process, err := os.FindProcess(pid)
	if err != nil {
		return false, -1
	}

	// Signal 0 doesn't actually send a signal, just checks if we can access the process
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		// Process doesn't exist
		return false, -1
	}

	// Process is running
	return true, 0
}
