package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// handlePipelineFindingDetails handles pipeline scan flaws from local results files
func handlePipelineFindingDetails(ctx context.Context, req *FindingDetailsRequest) (interface{}, error) {
	// Parse the pipeline flaw ID (format: "1234-1")
	flawIDComponents, err := parsePipelineFlawIDString(req.FlawID)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Invalid pipeline flaw ID '%s': %v", req.FlawID, err),
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
				"text": fmt.Sprintf(`Pipeline Finding Details
========================

Application Path: %s
Results Directory: %s

❌ No results found

%v

To generate results, run a pipeline scan using the pipeline-scan tool.
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
		if scanResults.Findings[i].IssueID == flawIDComponents.IssueID {
			matches = append(matches, scanResults.Findings[i])
		}
	}

	if len(matches) == 0 {
		return pipelineErrorResponse(fmt.Sprintf(`Pipeline Finding Details
========================

Application Path: %s
Flaw ID: %s

❌ Flaw not found

The specified flaw ID was not found in the pipeline scan results.
`, req.ApplicationPath, req.FlawID)), nil
	}

	// Check if the requested occurrence exists
	if flawIDComponents.Occurrence > len(matches) {
		occurrenceList := ""
		for i := 0; i < len(matches); i++ {
			flawIDStr := fmt.Sprintf("%d-%d", flawIDComponents.IssueID, i+1)
			occurrenceList += fmt.Sprintf("- flaw_id %s: CWE-%s at %s:%d\n",
				flawIDStr, matches[i].CWEID, matches[i].Files.SourceFile.File, matches[i].Files.SourceFile.Line)
		}

		return pipelineErrorResponse(fmt.Sprintf(`Pipeline Finding Details
========================

Application Path: %s
Flaw ID: %s

❌ Occurrence not found

Issue ID %d has %d occurrence(s), but you requested occurrence %d.

Available occurrences:
%s
`, req.ApplicationPath, req.FlawID, flawIDComponents.IssueID, len(matches), flawIDComponents.Occurrence, occurrenceList)), nil
	}

	// Get the requested occurrence (occurrence is 1-based)
	targetFlaw := &matches[flawIDComponents.Occurrence-1]

	// Transform stack dumps to data paths
	detailedFlaw := transformToDetailedFlaw(targetFlaw)

	// Format and return the response
	return formatPipelineDetailedFindingsResponse(req.ApplicationPath, resultsFile, &detailedFlaw, flawIDComponents.IssueID, flawIDComponents.Occurrence), nil
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

// cleanVeracodeAnnotations removes Veracode's internal annotation markers from strings
// Removes patterns like /**X-VC ... */ (may appear multiple times)
func cleanVeracodeAnnotations(input string) string {
	// Pattern matches /**X-VC followed by anything until the first */
	re := regexp.MustCompile(`/\*\*X-VC\s[^*]*\*/`)
	return re.ReplaceAllString(input, "")
}

// formatPipelineDetailedFindingsResponse formats the detailed findings into an MCP response
func formatPipelineDetailedFindingsResponse(appPath, resultsFile string, flaw *PipelineDetailedFlaw, issueID, occurrence int) map[string]interface{} {
	// Build LLM-optimized JSON structure
	result := buildPipelineLLMOptimizedResponse(appPath, resultsFile, flaw, issueID, occurrence)

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting flaw details: %v", err),
				},
			},
		}
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(jsonBytes),
			},
		},
	}
}

// buildPipelineLLMOptimizedResponse creates an LLM-optimized JSON structure for pipeline flaw details
func buildPipelineLLMOptimizedResponse(appPath, resultsFile string, flaw *PipelineDetailedFlaw, issueID, occurrence int) map[string]interface{} {
	// Parse CWE ID
	var cweID int
	_, _ = fmt.Sscanf(flaw.CWEID, "%d", &cweID)

	// Extract references and clean the display text
	references := ExtractReferences(flaw.DisplayText)
	cleanDesc := CleanDescription(flaw.DisplayText)

	// Build the response structure
	response := map[string]interface{}{
		"flaw_id": fmt.Sprintf("%d-%d", issueID, occurrence),
		"application": map[string]interface{}{
			"path":         appPath,
			"results_file": filepath.Base(resultsFile),
		},
		"scan_type":   "PIPELINE",
		"title":       flaw.Title,
		"cwe_id":      cweID,
		"issue_type":  flaw.IssueType,
		"severity":    transformPipelineSeverity(flaw.Severity),
		"description": cleanDesc,
		"location": map[string]interface{}{
			"source_file": flaw.Files.SourceFile.File,
			"line":        flaw.Files.SourceFile.Line,
			"function":    flaw.Files.SourceFile.FunctionName,
		},
	}

	// Add references if present
	if len(references) > 0 {
		refList := make([]map[string]string, 0, len(references))
		for _, ref := range references {
			refList = append(refList, map[string]string{
				"name": ref.Name,
				"url":  ref.URL,
			})
		}
		response["references"] = refList
	}

	// Add data paths
	if len(flaw.DataPaths) > 0 {
		dataPaths := make([]map[string]interface{}, 0, len(flaw.DataPaths))
		for _, path := range flaw.DataPaths {
			steps := make([]map[string]interface{}, 0, len(path.Steps))
			for _, step := range path.Steps {
				stepData := map[string]interface{}{
					"source_file":   step.SourceFile,
					"source_line":   step.SourceLine,
					"function_name": step.FunctionName,
				}
				if step.VarNames != "" {
					stepData["var_names"] = step.VarNames
				}
				if step.QualifiedFunctionName != "" {
					stepData["qualified_function_name"] = step.QualifiedFunctionName
				}
				if step.RelativeLocation != "" {
					stepData["relative_location"] = step.RelativeLocation
				}
				steps = append(steps, stepData)
			}
			dataPaths = append(dataPaths, map[string]interface{}{
				"steps": steps,
			})
		}
		response["data_paths"] = dataPaths
		response["data_path_count"] = len(flaw.DataPaths)
	}

	// Add helpful instructions for the LLM
	response["_llm_instructions"] = "When presenting data paths to the user, create clickable markdown links for file references. Format as: [filename:line](filepath#Lline). Example: [UserController.java:165](com/veracode/verademo/controller/UserController.java#L165). Offer to explain the data flow or provide remediation guidance."

	return response
}
