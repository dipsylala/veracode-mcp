package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const PipelineStatusToolName = "pipeline-status"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(PipelineStatusToolName, handlePipelineStatus)
}

// PipelineStatusRequest represents the parsed parameters for pipeline-status
type PipelineStatusRequest struct {
	ApplicationPath string
}

// parsePipelineStatusRequest extracts and validates parameters from the raw args map
func parsePipelineStatusRequest(args map[string]interface{}) (*PipelineStatusRequest, error) {
	req := &PipelineStatusRequest{}

	// Extract required fields
	var err error
	req.ApplicationPath, err = extractRequiredString(args, "application_path")
	if err != nil {
		return nil, err
	}

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
	outputDir := veracodeWorkDir(req.ApplicationPath, "pipeline")
	pidFile := filepath.Join(outputDir, "pipeline.pid")

	// Check if PID file exists
	// #nosec G304 -- pidFile is constructed from validated outputDir and fixed filename, not user input
	pidData, err := os.ReadFile(pidFile)
	if err != nil {
		if os.IsNotExist(err) {
			// No PID file — check whether a completed results file exists
			matches, globErr := filepath.Glob(filepath.Join(outputDir, "results-*.json"))
			if globErr == nil && len(matches) > 0 {
				// Use the most recent results file (Glob returns lexicographic order,
				// and the timestamp format sorts correctly)
				latest := matches[len(matches)-1]
				return map[string]interface{}{
					"content": []map[string]string{{
						"type": "text",
						"text": fmt.Sprintf(`Pipeline Scan Status
==================

Application Path: %s
Status: COMPLETED ✓

No active scan in progress. A previous scan completed successfully.

Results File: %s

Use pipeline-findings to retrieve the results.
`, req.ApplicationPath, latest),
					}},
				}, nil
			}

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
