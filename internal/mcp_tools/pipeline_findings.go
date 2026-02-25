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

const PipelineFindingsToolName = "pipeline-findings"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(PipelineFindingsToolName, handlePipelineFindings)
}

// PipelineFindingsRequest represents the parsed parameters for pipeline-findings
type PipelineFindingsRequest struct {
	ApplicationPath string
	Size            int   `json:"size,omitempty"`
	Page            int   `json:"page,omitempty"`
	ViolatesPolicy  *bool `json:"violates_policy,omitempty"`
}

// parsePipelineFindingsRequest extracts and validates parameters from the raw args map
func parsePipelineFindingsRequest(args map[string]interface{}) (*PipelineFindingsRequest, error) {
	req := &PipelineFindingsRequest{}

	// Extract required fields
	var err error
	req.ApplicationPath, err = extractRequiredString(args, "application_path")
	if err != nil {
		return nil, err
	}

	// Extract optional fields with defaults — default violates_policy to true if not provided
	req.Size = extractInt(args, "page_size", 10)
	req.Page = extractInt(args, "page", 0)
	violatesPolicy, provided := extractOptionalBool(args, "violates_policy")
	if !provided {
		defaultTrue := true
		violatesPolicy = &defaultTrue
	}
	req.ViolatesPolicy = violatesPolicy

	// Validate pagination bounds
	if err := validatePaginationParams(req.Size, req.Page); err != nil {
		return nil, err
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
	FlawMatch       FlawMatch  `json:"flaw_match,omitempty"`
	StackDumps      StackDumps `json:"stack_dumps,omitempty"`
}

// FlawMatch contains the hash data used to uniquely identify a flaw across scans
type FlawMatch struct {
	FlawHash string `json:"flaw_hash"`
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

// formatEmptyPipelineFindingsResponse returns a properly structured empty response when no results exist
func formatEmptyPipelineFindingsResponse(appPath string) map[string]interface{} {
	response := MCPFindingsResponse{
		Application: MCPApplication{
			Name: filepath.Base(appPath),
			ID:   appPath,
		},
		PolicyFilter: false,
		Findings:     []MCPFinding{},
		Summary: MCPFindingsSummary{
			TotalFindings:    0,
			OpenFindings:     0,
			PolicyViolations: 0,
			BySeverity: map[string]int{
				"very high": 0,
				"high":      0,
				"medium":    0,
				"low":       0,
				"very low":  0,
				"info":      0,
			},
			ByStatus:     map[string]int{"open": 0},
			ByMitigation: map[string]int{"none": 0},
		},
	}

	responseJSON, _ := json.Marshal(response)

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(responseJSON),
			},
		},
		"structuredContent": response,
	}
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

// handlePipelineFindings retrieves and formats pipeline scan results
func handlePipelineFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parsePipelineFindingsRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Locate the results directory
	outputDir := veracodeWorkDir(req.ApplicationPath, "pipeline")

	// Find the most recent full results file (used for totals)
	resultsFile, err := findMostRecentFile(outputDir, "results-", ".json")
	if err != nil {
		// Return empty results with proper structure for UI
		return formatEmptyPipelineFindingsResponse(req.ApplicationPath), nil
	}

	// Read and parse the full results file
	// #nosec G304 -- resultsFile is from findMostRecentFile which validates the directory
	resultsData, err := os.ReadFile(resultsFile)
	if err != nil {
		return pipelineErrorResponse(fmt.Sprintf("Failed to read results file: %v", err)), nil
	}

	var scanResults PipelineScanResults
	if err = json.Unmarshal(resultsData, &scanResults); err != nil {
		return pipelineErrorResponse(fmt.Sprintf("Failed to parse results file: %v", err)), nil
	}

	// Always load filtered results to know which flaws violate policy
	var filteredResults *PipelineScanResults
	if filteredFile, ferr := findMostRecentFile(outputDir, "filtered-results-", ".json"); ferr == nil {
		// #nosec G304 -- filteredFile is from findMostRecentFile which validates the directory
		if filteredData, ferr := os.ReadFile(filteredFile); ferr == nil {
			var fr PipelineScanResults
			if ferr = json.Unmarshal(filteredData, &fr); ferr == nil {
				filteredResults = &fr
			}
		}
	}

	// Format and return the response
	return formatPipelineFindingsResponse(ctx, req.ApplicationPath, resultsFile, &scanResults, filteredResults, req), nil
}

// buildPipelineSummaryTotals computes summary counts from the full (unfiltered) results.
func buildPipelineSummaryTotals(findings []PipelineFlaw) MCPFindingsSummary {
	summary := MCPFindingsSummary{
		TotalFindings: len(findings),
		OpenFindings:  len(findings), // all pipeline findings are "open"
		BySeverity: map[string]int{
			"very high": 0,
			"high":      0,
			"medium":    0,
			"low":       0,
			"very low":  0,
			"info":      0,
		},
		ByStatus:     map[string]int{"open": 0},
		ByMitigation: map[string]int{"none": 0},
	}
	for _, f := range findings {
		sev := transformPipelineSeverity(f.Severity)
		if _, ok := summary.BySeverity[sev]; ok {
			summary.BySeverity[sev]++
		}
		summary.ByMitigation["none"]++
	}
	return summary
}

// buildPolicyViolatingIDs returns a set of composite keys (issue_id:flaw_hash) present in the filtered (policy-failing) results.
func buildPolicyViolatingIDs(filteredResults *PipelineScanResults) map[string]bool {
	ids := make(map[string]bool)
	if filteredResults != nil {
		for _, f := range filteredResults.Findings {
			ids[flawKey(f)] = true
		}
	}
	return ids
}

// flawKey returns a composite key for a flaw using issue_id and flaw_hash.
// flaw_hash alone is not unique across scans; combined with issue_id it identifies a specific occurrence.
func flawKey(f PipelineFlaw) string {
	return fmt.Sprintf("%d:%s", f.IssueID, f.FlawMatch.FlawHash)
}

// findMostRecentFile finds the most recent file with the given prefix and suffix in the directory
func findMostRecentFile(dir, prefix, suffix string) (string, error) {
	// Check if directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return "", fmt.Errorf("pipeline directory does not exist: %s", dir)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}

	var matches []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, suffix) {
			matches = append(matches, filepath.Join(dir, name))
		}
	}

	if len(matches) == 0 {
		return "", fmt.Errorf("no %s*%s files found in directory: %s", prefix, suffix, dir)
	}

	sort.Strings(matches)
	return matches[len(matches)-1], nil
}

// formatPipelineFindingsResponse formats the pipeline findings into an MCP response
func formatPipelineFindingsResponse(ctx context.Context, appPath, resultsFile string, results *PipelineScanResults, filteredResults *PipelineScanResults, req *PipelineFindingsRequest) map[string]interface{} {
	// Determine which findings to display (paginate)
	// Totals are always calculated from the full results.
	violatesPolicy := req.ViolatesPolicy != nil && *req.ViolatesPolicy
	displaySource := results.Findings
	if violatesPolicy && filteredResults != nil {
		displaySource = filteredResults.Findings
	}
	// Build MCP response structure similar to static findings
	response := MCPFindingsResponse{
		Application: MCPApplication{
			Name: filepath.Base(appPath),
			ID:   appPath,
		},
		PolicyFilter: violatesPolicy,
		Findings:     []MCPFinding{},
	}

	// Compute summary totals directly from the full results (always results-*.json)
	response.Summary = buildPipelineSummaryTotals(results.Findings)

	// Build set of issue IDs that violate policy (from filtered-results)
	policyViolatingIDs := buildPolicyViolatingIDs(filteredResults)
	if filteredResults != nil {
		response.Summary.PolicyViolations = len(filteredResults.Findings)
	}

	// Build display findings from the appropriate source:
	//   violates_policy=true  → filtered-results-*.json (policy-failing flaws only)
	//   violates_policy=false → results-*.json (all flaws)
	displayFindings := make([]MCPFinding, 0, len(displaySource))
	issueIDCounts := make(map[int]int)
	for _, finding := range displaySource {
		issueIDCounts[finding.IssueID]++
		mcpFinding := processPipelineFinding(finding, issueIDCounts[finding.IssueID])
		mcpFinding.ViolatesPolicy = policyViolatingIDs[flawKey(finding)]
		displayFindings = append(displayFindings, mcpFinding)
	}

	// Sort by severity descending, then flaw ID for stable ordering
	sort.Slice(displayFindings, func(i, j int) bool {
		if displayFindings[i].SeverityScore != displayFindings[j].SeverityScore {
			return displayFindings[i].SeverityScore > displayFindings[j].SeverityScore
		}
		return displayFindings[i].FlawID < displayFindings[j].FlawID
	})

	// Paginate
	displayCount := len(displayFindings)
	response.Pagination = &MCPPagination{
		CurrentPage:   req.Page,
		PageSize:      req.Size,
		TotalElements: displayCount,
		TotalPages:    (displayCount + req.Size - 1) / req.Size,
		HasNext:       (req.Page+1)*req.Size < displayCount,
		HasPrevious:   req.Page > 0,
	}

	startIdx := req.Page * req.Size
	endIdx := min(startIdx+req.Size, displayCount)
	if startIdx > displayCount {
		startIdx = displayCount
	}
	for i := startIdx; i < endIdx; i++ {
		response.Findings = append(response.Findings, displayFindings[i])
	}

	// Marshal response to JSON for non-UI clients
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("Warning: Failed to marshal response to JSON: %v", err)
		responseJSON, _ = json.Marshal(response) // Fall back to compact JSON
	}

	// Build pagination summary for LLM
	findingsLabel := "findings"
	if violatesPolicy {
		findingsLabel = "policy-relevant findings"
	}
	var paginationSummary string
	if response.Pagination != nil {
		paginationSummary = fmt.Sprintf("Showing %d findings on page %d of %d (Total: %d %s across all pages)\n\n",
			len(response.Findings),
			response.Pagination.CurrentPage+1, // Display as 1-based
			response.Pagination.TotalPages,
			response.Pagination.TotalElements,
			findingsLabel,
		)
	} else {
		paginationSummary = fmt.Sprintf("Showing %d %s\n\n", len(response.Findings), findingsLabel)
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
	log.Printf("[PIPELINE-FINDINGS] ClientSupportsUIFromContext(ctx) returned: %v", clientSupportsUI)

	if clientSupportsUI {
		log.Printf("Pipeline findings: Returning %d findings from %s (content: JSON, structuredContent: full data for UI)", len(response.Findings), resultsFile)
		result["structuredContent"] = response
	} else {
		log.Printf("Pipeline findings: Returning %d findings from %s (content: JSON only, no structuredContent - client doesn't support UI)", len(response.Findings), resultsFile)
	}

	return result
}

// processPipelineFinding converts a pipeline flaw to MCP finding format
// occurrence is used to make flaw_id unique when issue_id is duplicated
func processPipelineFinding(flaw PipelineFlaw, occurrence int) MCPFinding {
	// Parse CWE ID from string to int
	var cweID int32
	_, _ = fmt.Sscanf(flaw.CWEID, "%d", &cweID) // Ignore error, default to 0 if parse fails

	// Validate severity to prevent overflow
	severityScore := flaw.Severity
	if severityScore > math.MaxInt32 {
		severityScore = math.MaxInt32
	}

	// Create a unique flaw_id with occurrence suffix (always includes suffix, even for first occurrence)
	// This ensures flaw_id is always a string format like "1000-1", "1000-2", etc.
	flawID := fmt.Sprintf("%d-%d", flaw.IssueID, occurrence)

	finding := MCPFinding{
		FlawID:         flawID,
		ScanType:       "STATIC",
		Severity:       transformPipelineSeverity(flaw.Severity),
		SeverityScore:  int32(severityScore), // #nosec G115 - validated above
		CweId:          cweID,
		Description:    CleanDescription(flaw.DisplayText),
		Status:         "open",                          // Pipeline findings are always open
		ViolatesPolicy: false,                           // may be overridden by caller when sourced from filtered-results
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
