package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const PipelineResultsToolName = "pipeline-results"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(PipelineResultsToolName, handlePipelineResults)
}

// PipelineResultsRequest represents the parsed parameters for pipeline-results
type PipelineResultsRequest struct {
	ApplicationPath string
	Size            int `json:"size,omitempty"`
	Page            int `json:"page,omitempty"`
}

// parsePipelineResultsRequest extracts and validates parameters from the raw args map
func parsePipelineResultsRequest(args map[string]interface{}) (*PipelineResultsRequest, error) {
	req := &PipelineResultsRequest{
		Size: 10,
		Page: 0,
	}

	// Extract application path
	appPath, ok := args["application_path"].(string)
	if !ok || appPath == "" {
		return nil, fmt.Errorf("application_path is required and must be a non-empty string")
	}
	req.ApplicationPath = appPath

	// Extract size if provided
	if size, ok := args["size"].(float64); ok {
		req.Size = int(size)
	}

	// Extract page if provided
	if page, ok := args["page"].(float64); ok {
		req.Page = int(page)
	}

	return req, nil
}

// PipelineScanResults represents the JSON structure from Veracode Pipeline Scanner
type PipelineScanResults struct {
	ScanID     string         `json:"scan_id"`
	ScanStatus string         `json:"scan_status"`
	Message    string         `json:"message"`
	Modules    []string       `json:"modules"`
	Findings   []PipelineFlaw `json:"findings"`
}

// PipelineFlaw represents a finding from the pipeline scanner
type PipelineFlaw struct {
	Title           string     `json:"title"`
	IssueID         int        `json:"issue_id"`
	CWEID           string     `json:"cwe_id"` // Note: This is a string in the actual JSON
	IssueType       string     `json:"issue_type"`
	IssueTypeID     string     `json:"issue_type_id"`
	Severity        int        `json:"severity"`
	DisplayText     string     `json:"display_text"`
	Files           FileInfo   `json:"files"`
	FlawDetailsLink string     `json:"flaw_details_link"`
	StackDumps      StackDumps `json:"stack_dumps,omitempty"`
}

type FileInfo struct {
	SourceFile SourceFileInfo `json:"source_file"`
}

type SourceFileInfo struct {
	File              string `json:"file"`
	Line              int    `json:"line"`
	FunctionName      string `json:"function_name,omitempty"`
	QualifiedName     string `json:"qualified_function_name,omitempty"`
	FunctionPrototype string `json:"function_prototype,omitempty"`
	Scope             string `json:"scope,omitempty"`
}

// pipelineErrorResponse creates a standardized error response for pipeline results
func pipelineErrorResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": message,
		}},
	}
}

// handlePipelineResults retrieves and formats pipeline scan results
func handlePipelineResults(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parsePipelineResultsRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Locate the results directory
	outputDir := filepath.Join(req.ApplicationPath, ".veracode", "pipeline")

	// Find the most recent results file
	resultsFile, err := findMostRecentResultsFile(outputDir)
	if err != nil {
		return pipelineErrorResponse(fmt.Sprintf(`Pipeline Scan Results
==================

Application Path: %s
Results Directory: %s

❌ No results found

%v

To generate results, run a pipeline scan using the pipeline-static-scan tool.
`, req.ApplicationPath, outputDir, err)), nil
	}

	// Read and parse the results file
	// #nosec G304 -- resultsFile is from findMostRecentResultsFile which validates the directory
	resultsData, err := os.ReadFile(resultsFile)
	if err != nil {
		return pipelineErrorResponse(fmt.Sprintf("Failed to read results file: %v", err)), nil
	}

	var scanResults PipelineScanResults
	if err = json.Unmarshal(resultsData, &scanResults); err != nil {
		return pipelineErrorResponse(fmt.Sprintf("Failed to parse results file: %v", err)), nil
	}

	// Format and return the response
	return formatPipelineResultsResponse(ctx, req.ApplicationPath, resultsFile, &scanResults, req), nil
}

// findMostRecentResultsFile finds the most recent results-*.json file in the directory
func findMostRecentResultsFile(dir string) (string, error) {
	// Check if directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return "", fmt.Errorf("pipeline directory does not exist: %s", dir)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}

	var resultsFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if strings.HasPrefix(name, "results-") && strings.HasSuffix(name, ".json") {
			resultsFiles = append(resultsFiles, filepath.Join(dir, name))
		}
	}

	if len(resultsFiles) == 0 {
		return "", fmt.Errorf("no results files found in directory: %s", dir)
	}

	// Sort by filename (which includes timestamp) to get the most recent
	sort.Strings(resultsFiles)
	return resultsFiles[len(resultsFiles)-1], nil
}

// formatPipelineResultsResponse formats the pipeline results into an MCP response
func formatPipelineResultsResponse(ctx context.Context, appPath, resultsFile string, results *PipelineScanResults, req *PipelineResultsRequest) map[string]interface{} {
	// Build MCP response structure similar to static findings
	response := MCPFindingsResponse{
		Application: MCPApplication{
			Name: filepath.Base(appPath),
			ID:   appPath,
		},
		Summary: MCPFindingsSummary{
			BySeverity: map[string]int{
				"very high": 0,
				"high":      0,
				"medium":    0,
				"low":       0,
				"very low":  0,
				"info":      0,
			},
			ByStatus: map[string]int{
				"open": 0,
			},
			ByMitigation: map[string]int{
				"none": 0,
			},
		},
		Findings: []MCPFinding{},
	}

	// Set total findings - all pipeline findings are "open" by default
	totalFindings := len(results.Findings)
	response.Summary.TotalFindings = totalFindings
	response.Summary.PolicyViolations = 0 // Pipeline scanner doesn't provide policy violation info per finding

	// Add pagination info
	response.Pagination = &MCPPagination{
		CurrentPage:   req.Page,
		PageSize:      req.Size,
		TotalElements: totalFindings,
		TotalPages:    (totalFindings + req.Size - 1) / req.Size,
		HasNext:       (req.Page+1)*req.Size < totalFindings,
		HasPrevious:   req.Page > 0,
	}

	// First, convert all findings and sort them by severity
	allFindings := make([]MCPFinding, 0, len(results.Findings))
	for _, finding := range results.Findings {
		mcpFinding := processPipelineFinding(finding)
		allFindings = append(allFindings, mcpFinding)
	}

	// Sort findings by severity (most severe first), then by Flaw ID for consistent ordering
	sort.Slice(allFindings, func(i, j int) bool {
		if allFindings[i].SeverityScore != allFindings[j].SeverityScore {
			return allFindings[i].SeverityScore > allFindings[j].SeverityScore
		}
		// Secondary sort by Flaw ID for consistent ordering when severities are equal
		return allFindings[i].FlawID < allFindings[j].FlawID
	})

	// Apply pagination
	startIdx := req.Page * req.Size
	endIdx := startIdx + req.Size
	if startIdx > len(allFindings) {
		startIdx = len(allFindings)
	}
	if endIdx > len(allFindings) {
		endIdx = len(allFindings)
	}

	// Process paginated findings and update counters
	for i := startIdx; i < endIdx; i++ {
		mcpFinding := allFindings[i]
		response.Findings = append(response.Findings, mcpFinding)

		// Update summary counters for the current page
		severity := strings.ToLower(mcpFinding.Severity)
		if count, ok := response.Summary.BySeverity[severity]; ok {
			response.Summary.BySeverity[severity] = count + 1
		}

		if mcpFinding.Status == "open" || mcpFinding.Status == "OPEN" {
			response.Summary.OpenFindings++
		}

		mitigationStatus := strings.ToLower(mcpFinding.MitigationStatus)
		if mitigationStatus == "" {
			mitigationStatus = "none"
		}
		if count, ok := response.Summary.ByMitigation[mitigationStatus]; ok {
			response.Summary.ByMitigation[mitigationStatus] = count + 1
		}
	}

	// Marshal response to JSON for non-UI clients
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("Warning: Failed to marshal response to JSON: %v", err)
		responseJSON, _ = json.Marshal(response) // Fall back to compact JSON
	}

	// Build pagination summary for LLM
	var paginationSummary string
	if response.Pagination != nil {
		paginationSummary = fmt.Sprintf("Showing %d findings on page %d of %d (Total: %d findings across all pages)\n\n",
			len(response.Findings),
			response.Pagination.CurrentPage+1, // Display as 1-based
			response.Pagination.TotalPages,
			response.Pagination.TotalElements,
		)
	} else {
		paginationSummary = fmt.Sprintf("Showing %d findings\n\n", len(response.Findings))
	}

	// Build result with content
	result := map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": paginationSummary + string(responseJSON),
			},
		},
	}

	// Only include structuredContent if client supports MCP Apps UI
	clientSupportsUI := ClientSupportsUIFromContext(ctx)
	log.Printf("[PIPELINE-RESULTS] ClientSupportsUIFromContext(ctx) returned: %v", clientSupportsUI)

	if clientSupportsUI {
		log.Printf("Pipeline results: Returning %d findings from %s (content: JSON, structuredContent: full data for UI)", len(response.Findings), resultsFile)
		result["structuredContent"] = response
	} else {
		log.Printf("Pipeline results: Returning %d findings from %s (content: JSON only, no structuredContent - client doesn't support UI)", len(response.Findings), resultsFile)
	}

	return result
}

// processPipelineFinding converts a pipeline flaw to MCP finding format
func processPipelineFinding(flaw PipelineFlaw) MCPFinding {
	// Parse CWE ID from string to int
	var cweID int32
	_, _ = fmt.Sscanf(flaw.CWEID, "%d", &cweID) // Ignore error, default to 0 if parse fails

	// Validate severity to prevent overflow
	severityScore := flaw.Severity
	if severityScore > math.MaxInt32 {
		severityScore = math.MaxInt32
	}

	finding := MCPFinding{
		FlawID:         fmt.Sprintf("%d", flaw.IssueID),
		ScanType:       "STATIC",
		Severity:       transformPipelineSeverity(flaw.Severity),
		SeverityScore:  int32(severityScore), // #nosec G115 - validated above
		CweId:          cweID,
		Description:    CleanDescription(flaw.DisplayText),
		Status:         "open",                          // Pipeline findings are always open
		ViolatesPolicy: false,                           // Pipeline scanner doesn't provide this per-finding
		FirstFound:     time.Now().Format(time.RFC3339), // Pipeline scans don't track this
		FilePath:       flaw.Files.SourceFile.File,
		LineNumber:     flaw.Files.SourceFile.Line,
		Procedure:      flaw.Files.SourceFile.FunctionName,
	}

	return finding
}

// transformPipelineSeverity converts numeric severity to text representation
func transformPipelineSeverity(severity int) string {
	severityMap := map[int]string{
		5: "very high",
		4: "high",
		3: "medium",
		2: "low",
		1: "very low",
		0: "info",
	}
	if text, ok := severityMap[severity]; ok {
		return text
	}
	return "info"
}
