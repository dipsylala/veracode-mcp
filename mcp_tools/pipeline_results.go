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
	RegisterTool(PipelineResultsToolName, func() ToolImplementation {
		return NewPipelineResultsTool()
	})
}

// PipelineResultsTool provides the pipeline-results tool
type PipelineResultsTool struct{}

// NewPipelineResultsTool creates a new pipeline results tool
func NewPipelineResultsTool() *PipelineResultsTool {
	return &PipelineResultsTool{}
}

// Initialize sets up the tool
func (t *PipelineResultsTool) Initialize() error {
	log.Printf("Initializing tool: %s", PipelineResultsToolName)
	return nil
}

// RegisterHandlers registers the pipeline results handler
func (t *PipelineResultsTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", PipelineResultsToolName)
	registry.RegisterHandler(PipelineResultsToolName, t.handlePipelineResults)
	return nil
}

// Shutdown cleans up tool resources
func (t *PipelineResultsTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", PipelineResultsToolName)
	return nil
}

// PipelineResultsRequest represents the parsed parameters for pipeline-results
type PipelineResultsRequest struct {
	ApplicationPath string
}

// parsePipelineResultsRequest extracts and validates parameters from the raw args map
func parsePipelineResultsRequest(args map[string]interface{}) (*PipelineResultsRequest, error) {
	req := &PipelineResultsRequest{}

	// Extract application_path (required)
	appPath, ok := args["application_path"].(string)
	if !ok || appPath == "" {
		return nil, fmt.Errorf("application_path is required and must be a non-empty string")
	}
	req.ApplicationPath = appPath

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
	Title           string   `json:"title"`
	IssueID         int      `json:"issue_id"`
	CWEID           string   `json:"cwe_id"` // Note: This is a string in the actual JSON
	IssueType       string   `json:"issue_type"`
	IssueTypeID     string   `json:"issue_type_id"`
	Severity        int      `json:"severity"`
	DisplayText     string   `json:"display_text"`
	Files           FileInfo `json:"files"`
	FlawDetailsLink string   `json:"flaw_details_link"`
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

// handlePipelineResults retrieves and formats pipeline scan results
func (t *PipelineResultsTool) handlePipelineResults(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parsePipelineResultsRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Locate the results directory
	outputDir := filepath.Join(req.ApplicationPath, ".veracode_pipeline")

	// Find the most recent results file
	resultsFile, err := findMostRecentResultsFile(outputDir)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Pipeline Scan Results
==================

Application Path: %s
Results Directory: %s

❌ No results found

%v

To generate results, run a pipeline scan using the pipeline-static-scan tool.
`, req.ApplicationPath, outputDir, err),
			}},
		}, nil
	}

	// Read and parse the results file
	// #nosec G304 -- resultsFile is from findMostRecentResultsFile which validates the directory
	resultsData, err := os.ReadFile(resultsFile)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to read results file: %v", err),
		}, nil
	}

	var scanResults PipelineScanResults
	err = json.Unmarshal(resultsData, &scanResults)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to parse results file: %v", err),
		}, nil
	}

	// Format and return the response
	return formatPipelineResultsResponse(ctx, req.ApplicationPath, resultsFile, &scanResults), nil
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
func formatPipelineResultsResponse(ctx context.Context, appPath, resultsFile string, results *PipelineScanResults) map[string]interface{} {
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
				"open": len(results.Findings), // All pipeline findings are "open"
			},
			ByMitigation: map[string]int{
				"none": len(results.Findings), // Pipeline findings have no mitigation
			},
		},
		Findings: []MCPFinding{},
	}

	// Set total findings - all pipeline findings are "open" by default
	response.Summary.TotalFindings = len(results.Findings)
	response.Summary.OpenFindings = len(results.Findings)
	response.Summary.PolicyViolations = 0 // Pipeline scanner doesn't provide policy violation info per finding

	// Process each finding
	for _, finding := range results.Findings {
		mcpFinding := processPipelineFinding(finding)
		response.Findings = append(response.Findings, mcpFinding)

		// Update summary counters
		severity := strings.ToLower(mcpFinding.Severity)
		if count, ok := response.Summary.BySeverity[severity]; ok {
			response.Summary.BySeverity[severity] = count + 1
		}
	}

	// Sort findings by severity (most severe first), then by Flaw ID for consistent ordering
	sort.Slice(response.Findings, func(i, j int) bool {
		if response.Findings[i].SeverityScore != response.Findings[j].SeverityScore {
			return response.Findings[i].SeverityScore > response.Findings[j].SeverityScore
		}
		// Secondary sort by Flaw ID for consistent ordering when severities are equal
		return response.Findings[i].FlawID < response.Findings[j].FlawID
	})

	// Marshal response to JSON for non-UI clients
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("Warning: Failed to marshal response to JSON: %v", err)
		responseJSON, _ = json.Marshal(response) // Fall back to compact JSON
	}

	// Build result with content
	result := map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(responseJSON),
			},
		},
	}

	// Only include structuredContent if client supports MCP Apps UI
	// This can be forced via --force-mcp-app flag
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

	// Clean up the display text (remove HTML tags)
	cleanDesc := CleanDescription(flaw.DisplayText)

	if flaw.Severity > math.MaxInt32 {
		flaw.Severity = math.MaxInt32
	}

	finding := MCPFinding{
		FlawID:         fmt.Sprintf("%d", flaw.IssueID),
		ScanType:       "STATIC",
		Severity:       transformPipelineSeverity(flaw.Severity),
		SeverityScore:  int32(flaw.Severity), // #nosec G115 - validated above
		CweId:          cweID,
		Description:    cleanDesc,
		Status:         "open",                          // Pipeline findings are always open
		ViolatesPolicy: false,                           // Pipeline scanner doesn't provide this per-finding
		FirstFound:     time.Now().Format(time.RFC3339), // Pipeline scans don't track this
		FilePath:       flaw.Files.SourceFile.File,
		LineNumber:     flaw.Files.SourceFile.Line,
	}

	// Add function information if available
	if flaw.Files.SourceFile.FunctionName != "" {
		finding.Procedure = flaw.Files.SourceFile.FunctionName
	}

	return finding
}

// transformPipelineSeverity converts numeric severity to text representation
func transformPipelineSeverity(severity int) string {
	switch severity {
	case 5:
		return "very high"
	case 4:
		return "high"
	case 3:
		return "medium"
	case 2:
		return "low"
	case 1:
		return "very low"
	case 0:
		return "info"
	default:
		return "info"
	}
}
