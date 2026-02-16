package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const PipelineDetailedResultsToolName = "pipeline-detailed-results"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(PipelineDetailedResultsToolName, handlePipelineDetailedResults)
}

// PipelineDetailedResultsRequest represents the parsed parameters for pipeline-detailed-results
type PipelineDetailedResultsRequest struct {
	ApplicationPath string
	FlawID          *FlawIDComponents
}

// parsePipelineDetailedResultsRequest extracts and validates parameters from the raw args map
func parsePipelineDetailedResultsRequest(args map[string]interface{}) (*PipelineDetailedResultsRequest, error) {
	req := &PipelineDetailedResultsRequest{}

	// Extract required fields
	var err error
	req.ApplicationPath, err = extractRequiredString(args, "application_path")
	if err != nil {
		return nil, err
	}

	req.FlawID, err = extractFlawIDString(args)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// PipelineDetailedFlaw represents a detailed finding with data paths
type PipelineDetailedFlaw struct {
	Title           string     `json:"title"`
	IssueID         int        `json:"issue_id"`
	CWEID           string     `json:"cwe_id"`
	IssueType       string     `json:"issue_type"`
	IssueTypeID     string     `json:"issue_type_id"`
	Severity        int        `json:"severity"`
	DisplayText     string     `json:"display_text"`
	Files           FileInfo   `json:"files"`
	FlawDetailsLink string     `json:"flaw_details_link"`
	DataPaths       []DataPath `json:"data_paths"`
}

// DataPath represents a data flow path (transformed from stack_dump)
type DataPath struct {
	Steps []Step `json:"steps"`
}

// Step represents a single step in the data path
type Step struct {
	FrameID               string `json:"frame_id"`
	FunctionName          string `json:"function_name"`
	SourceFile            string `json:"source_file"`
	SourceLine            string `json:"source_line"`
	SourceFileID          string `json:"source_file_id,omitempty"`
	VarNames              string `json:"var_names,omitempty"`
	QualifiedFunctionName string `json:"qualified_function_name,omitempty"`
	FunctionPrototype     string `json:"function_prototype,omitempty"`
	RelativeLocation      string `json:"relative_location,omitempty"`
}

// cleanVeracodeAnnotations removes Veracode's internal annotation markers from strings
// Removes patterns like /**X-VC ... */ (may appear multiple times)
func cleanVeracodeAnnotations(input string) string {
	// Pattern matches /**X-VC followed by anything until the first */
	re := regexp.MustCompile(`/\*\*X-VC\s[^*]*\*/`)
	return re.ReplaceAllString(input, "")
}

// StackDumps represents the stack_dumps structure in the JSON
type StackDumps struct {
	StackDump []StackDump `json:"stack_dump"`
}

// StackDump represents a single stack dump
type StackDump struct {
	Frame []RawFrame `json:"Frame"`
}

// RawFrame represents the raw frame structure from JSON
type RawFrame struct {
	FrameID               string      `json:"FrameId"`
	FunctionName          string      `json:"FunctionName"`
	SourceFile            string      `json:"SourceFile"`
	SourceLine            string      `json:"SourceLine"`
	SourceFileID          string      `json:"SourceFileId"`
	StatementText         interface{} `json:"StatementText"`
	VarNames              string      `json:"VarNames"`
	QualifiedFunctionName string      `json:"QualifiedFunctionName"`
	FunctionPrototype     string      `json:"FunctionPrototype"`
	RelativeLocation      string      `json:"RelativeLocation"`
}

// PipelineFlawWithStackDumps extends PipelineFlaw with stack dumps
type PipelineFlawWithStackDumps struct {
	Title           string     `json:"title"`
	IssueID         int        `json:"issue_id"`
	CWEID           string     `json:"cwe_id"`
	IssueType       string     `json:"issue_type"`
	IssueTypeID     string     `json:"issue_type_id"`
	Severity        int        `json:"severity"`
	DisplayText     string     `json:"display_text"`
	Files           FileInfo   `json:"files"`
	FlawDetailsLink string     `json:"flaw_details_link"`
	StackDumps      StackDumps `json:"stack_dumps"`
}

// handlePipelineDetailedResults retrieves detailed information about a specific flaw
func handlePipelineDetailedResults(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parsePipelineDetailedResultsRequest(args)
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
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Pipeline Detailed Results
========================

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
		return pipelineErrorResponse(fmt.Sprintf("Failed to read results file: %v", err)), nil
	}

	// Parse the full results to extract findings
	var scanResults struct {
		Findings []PipelineFlawWithStackDumps `json:"findings"`
	}
	if err = json.Unmarshal(resultsData, &scanResults); err != nil {
		return pipelineErrorResponse(fmt.Sprintf("Failed to parse results file: %v", err)), nil
	}

	// Find all flaws with the matching issue_id (there may be duplicates)
	var matches []PipelineFlawWithStackDumps
	for i := range scanResults.Findings {
		if scanResults.Findings[i].IssueID == req.FlawID.IssueID {
			matches = append(matches, scanResults.Findings[i])
		}
	}

	if len(matches) == 0 {
		return pipelineErrorResponse(fmt.Sprintf(`Pipeline Detailed Results
========================

Application Path: %s
Flaw ID: %d

❌ Flaw not found

The specified flaw ID was not found in the pipeline scan results.
`, req.ApplicationPath, req.FlawID.IssueID)), nil
	}

	// Select the correct occurrence
	if req.FlawID.Occurrence > len(matches) {
		occurrenceList := ""
		for i := 0; i < len(matches); i++ {
			flawIDStr := fmt.Sprintf("%d", req.FlawID.IssueID)
			if i > 0 {
				flawIDStr = fmt.Sprintf("%d-%d", req.FlawID.IssueID, i+1)
			}
			occurrenceList += fmt.Sprintf("- flaw_id %s: CWE-%s at %s:%d\n",
				flawIDStr, matches[i].CWEID, matches[i].Files.SourceFile.File, matches[i].Files.SourceFile.Line)
		}

		return pipelineErrorResponse(fmt.Sprintf(`Pipeline Detailed Results
========================

Application Path: %s
Flaw ID: %d-%d

❌ Occurrence not found

Issue ID %d has %d occurrence(s), but you requested occurrence %d.

Available occurrences:
%s
`, req.ApplicationPath, req.FlawID.IssueID, req.FlawID.Occurrence,
			req.FlawID.IssueID, len(matches), req.FlawID.Occurrence, occurrenceList)), nil
	}

	// Get the requested occurrence (occurrence is 1-based)
	targetFlaw := &matches[req.FlawID.Occurrence-1]

	// Transform stack dumps to data paths
	detailedFlaw := transformToDetailedFlaw(targetFlaw)

	// Format and return the response
	return formatPipelineDetailedResultsResponse(req.ApplicationPath, resultsFile, &detailedFlaw), nil
}

// transformToDetailedFlaw converts a flaw with stack dumps to a detailed flaw with data paths
func transformToDetailedFlaw(flaw *PipelineFlawWithStackDumps) PipelineDetailedFlaw {
	detailed := PipelineDetailedFlaw{
		Title:           flaw.Title,
		IssueID:         flaw.IssueID,
		CWEID:           flaw.CWEID,
		IssueType:       flaw.IssueType,
		IssueTypeID:     flaw.IssueTypeID,
		Severity:        flaw.Severity,
		DisplayText:     flaw.DisplayText,
		Files:           flaw.Files,
		FlawDetailsLink: flaw.FlawDetailsLink,
		DataPaths:       []DataPath{},
	}

	// Transform stack_dumps to data_paths
	for _, stackDump := range flaw.StackDumps.StackDump {
		dataPath := DataPath{
			Steps: make([]Step, 0, len(stackDump.Frame)),
		}

		for _, rawFrame := range stackDump.Frame {
			step := Step{
				FrameID:               rawFrame.FrameID,
				FunctionName:          rawFrame.FunctionName,
				SourceFile:            rawFrame.SourceFile,
				SourceLine:            rawFrame.SourceLine,
				SourceFileID:          rawFrame.SourceFileID,
				VarNames:              cleanVeracodeAnnotations(rawFrame.VarNames),
				QualifiedFunctionName: rawFrame.QualifiedFunctionName,
				FunctionPrototype:     rawFrame.FunctionPrototype,
				RelativeLocation:      rawFrame.RelativeLocation,
			}
			dataPath.Steps = append(dataPath.Steps, step)
		}

		detailed.DataPaths = append(detailed.DataPaths, dataPath)
	}

	return detailed
}

// formatPipelineDetailedResultsResponse formats the detailed results into an MCP response
func formatPipelineDetailedResultsResponse(appPath, resultsFile string, flaw *PipelineDetailedFlaw) map[string]interface{} {
	// Parse CWE ID
	var cweID int32
	_, _ = fmt.Sscanf(flaw.CWEID, "%d", &cweID)

	// Extract references and clean the display text
	references := ExtractReferences(flaw.DisplayText)
	cleanDesc := CleanDescription(flaw.DisplayText)

	// Build references section
	referencesSection := ""
	if len(references) > 0 {
		referencesSection = "\nReferences:\n"
		for _, ref := range references {
			if ref.Name != ref.URL {
				referencesSection += fmt.Sprintf("- [%s](%s)\n", ref.Name, ref.URL)
			} else {
				referencesSection += fmt.Sprintf("- %s\n", ref.URL)
			}
		}
	}

	// Build header with flaw info
	header := fmt.Sprintf(`Pipeline Detailed Results
========================

Application Path: %s
Results File: %s

Flaw ID: %d
Title: %s
CWE: CWE-%d
Issue Type: %s
Severity: %s

Description:
%s
%s
Source File: %s
Line: %d
Function: %s

Data Paths: %d

IMPORTANT: When presenting this data to the user, create clickable markdown links for all file references.
For each source file and line number in the data paths below, format as: [filename:line](filepath#Lline)
Example: [UserController.java:165](com/veracode/verademo/controller/UserController.java#L165)

Offer to explain the datapath and flaw in more detail, or to provide remediation guidance.
`,
		appPath,
		filepath.Base(resultsFile),
		flaw.IssueID,
		flaw.Title,
		cweID,
		flaw.IssueType,
		transformPipelineSeverity(flaw.Severity),
		cleanDesc,
		referencesSection,
		flaw.Files.SourceFile.File,
		flaw.Files.SourceFile.Line,
		flaw.Files.SourceFile.FunctionName,
		len(flaw.DataPaths),
	)

	// Marshal the detailed flaw to JSON
	responseJSON, err := json.Marshal(flaw)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf("Error formatting response: %v", err),
			}},
		}
	}

	// Return MCP response with both text header and JSON data
	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": header,
			},
			{
				"type": "text",
				"text": string(responseJSON),
			},
		},
	}
}
