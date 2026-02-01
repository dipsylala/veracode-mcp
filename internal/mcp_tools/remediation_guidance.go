package mcp_tools

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

//go:embed remediation_guidance
var remediationGuidanceFS embed.FS

const RemediationGuidanceToolName = "remediation-guidance"

// Auto-register this tool when the package is imported
func init() {
	RegisterSimpleTool(RemediationGuidanceToolName, handleGetRemediationGuidance)
}

// RemediationGuidanceRequest represents the parsed parameters for remediation-guidance
type RemediationGuidanceRequest struct {
	ApplicationPath string `json:"application_path"`
	FlawID          int    `json:"flaw_id"`
}

// parseRemediationGuidanceRequest extracts and validates parameters from the raw args map
func parseRemediationGuidanceRequest(args map[string]interface{}) (*RemediationGuidanceRequest, error) {
	req := &RemediationGuidanceRequest{}

	// Use JSON marshaling to automatically map args to struct
	jsonData, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal arguments: %w", err)
	}

	if err := json.Unmarshal(jsonData, req); err != nil {
		return nil, fmt.Errorf("failed to unmarshal arguments: %w", err)
	}

	// Validate required fields
	if req.ApplicationPath == "" {
		return nil, fmt.Errorf("application_path is required and must be an absolute path")
	}

	if req.FlawID == 0 {
		return nil, fmt.Errorf("flaw_id is required and must be a non-zero integer")
	}

	return req, nil
}

// detectLanguageFromFilename detects the programming language from a filename
func detectLanguageFromFilename(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))

	languageMap := map[string]string{
		".java":  "java",
		".jsp":   "java",
		".py":    "python",
		".aspx":  "csharp",
		".cs":    "csharp",
		".js":    "javascript",
		".jsx":   "javascript",
		".ts":    "javascript",
		".tsx":   "javascript",
		".php":   "php",
		".go":    "go",
		".rb":    "ruby",
		".cpp":   "cpp",
		".c":     "c",
		".h":     "c",
		".hpp":   "cpp",
		".swift": "swift",
		".kt":    "kotlin",
		".kts":   "kotlin",
		".rs":    "rust",
		".scala": "scala",
		".m":     "objectivec",
		".mm":    "objectivec",
	}

	if lang, ok := languageMap[ext]; ok {
		return lang
	}

	return ""
}

// getRemediationGuidancePath constructs the path to the appropriate remediation guidance file in embedded FS
func getRemediationGuidancePath(cweID int, language string) (string, error) {
	cweIDStr := strconv.Itoa(cweID)
	// Use path.Join for embedded FS (always forward slashes)
	cweDir := path.Join("remediation_guidance", cweIDStr)

	// Check if CWE directory exists in embedded FS
	_, err := remediationGuidanceFS.ReadDir(cweDir)
	if err != nil {
		return "", fmt.Errorf("no remediation guidance found for CWE-%d", cweID)
	}

	// If language is detected and language-specific guidance exists, use it
	if language != "" {
		langPath := path.Join(cweDir, language, "INDEX.md")
		if _, err := remediationGuidanceFS.ReadFile(langPath); err == nil {
			log.Printf("Found language-specific guidance: %s", langPath)
			return langPath, nil
		}
		log.Printf("No language-specific guidance found for %s, falling back to generic", language)
	}

	// Fall back to generic guidance
	genericPath := path.Join(cweDir, "INDEX.md")
	if _, err := remediationGuidanceFS.ReadFile(genericPath); err != nil {
		return "", fmt.Errorf("no remediation guidance INDEX.md found for CWE-%d", cweID)
	}

	return genericPath, nil
}

// findFlawInPipelineResults finds a specific flaw by issue_id in pipeline results
func findFlawInPipelineResults(results *PipelineScanResults, flawID int) *PipelineFlaw {
	for i := range results.Findings {
		if results.Findings[i].IssueID == flawID {
			return &results.Findings[i]
		}
	}
	return nil
}

func handleGetRemediationGuidance(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseRemediationGuidanceRequest(args)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}

	// Locate the pipeline results directory
	outputDir := filepath.Join(req.ApplicationPath, ".veracode", "pipeline")

	// Find the most recent results file
	resultsFile, err := findMostRecentResultsFile(outputDir)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Remediation Guidance Lookup
============================

Application Path: %s
Flaw ID: %d

❌ No pipeline results found

%v

Please run a pipeline scan first using the pipeline-scan tool.
`, req.ApplicationPath, req.FlawID, err),
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

	// Find the specific flaw by issue_id
	flaw := findFlawInPipelineResults(&scanResults, req.FlawID)
	if flaw == nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Remediation Guidance Lookup
============================

Application Path: %s
Flaw ID: %d

❌ Flaw not found

The specified flaw ID was not found in the pipeline scan results.

Please verify that:
1. The flaw ID is correct
2. The pipeline scan completed successfully
3. The flaw exists in the most recent scan
`, req.ApplicationPath, req.FlawID),
			}},
		}, nil
	}

	// Extract CWE ID from the flaw (it's a string in the JSON)
	cweID, err := strconv.Atoi(flaw.CWEID)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Remediation Guidance Lookup
============================

Application Path: %s
Flaw ID: %d

❌ Error: Invalid CWE ID format: %s
`, req.ApplicationPath, req.FlawID, flaw.CWEID),
			}},
		}, nil
	}

	// Extract source file from the flaw
	sourceFile := flaw.Files.SourceFile.File
	if sourceFile == "" {
		log.Printf("Warning: No source file found for flaw %d", req.FlawID)
	}

	// Detect language from filename
	language := ""
	if sourceFile != "" {
		language = detectLanguageFromFilename(sourceFile)
		log.Printf("Detected language '%s' from filename '%s'", language, sourceFile)
	}

	// Get the appropriate remediation guidance file
	guidancePath, err := getRemediationGuidancePath(cweID, language)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf("Remediation Guidance Lookup\n============================\n\nApplication Path: %s\nFlaw ID: %d\nCWE ID: %d\nSource File: %s\nDetected Language: %s\n\n❌ Error: %v\n", req.ApplicationPath, req.FlawID, cweID, sourceFile, language, err),
			}},
		}, nil
	}

	// Read the guidance file from embedded FS
	guidanceContent, err := remediationGuidanceFS.ReadFile(guidancePath)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf("Remediation Guidance Lookup\n============================\n\nApplication Path: %s\nFlaw ID: %d\nCWE ID: %d\nSource File: %s\nDetected Language: %s\n\n❌ Error: Failed to read guidance file\n\n%v\n", req.ApplicationPath, req.FlawID, cweID, sourceFile, language, err),
			}},
		}, nil
	}

	// Format and return the guidance
	return formatRemediationGuidanceResponse(req, cweID, flaw, language, sourceFile, string(guidanceContent)), nil
}

// parseMarkdownGuidance parses markdown guidance into structured sections
// nolint:gocyclo // Parsing markdown requires checking multiple conditions
func parseMarkdownGuidance(markdown string) map[string]interface{} {
	lines := strings.Split(markdown, "\n")

	var summary strings.Builder
	var keyPrinciples []string
	var remediationSteps []string

	currentSection := ""
	inList := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Skip the main title
		if strings.HasPrefix(trimmed, "# CWE-") {
			continue
		}

		// Detect sections
		if strings.HasPrefix(trimmed, "## LLM Guidance") {
			currentSection = "summary"
			inList = false
			continue
		} else if strings.HasPrefix(trimmed, "## Key Principles") {
			currentSection = "principles"
			inList = false
			continue
		} else if strings.HasPrefix(trimmed, "## Remediation Steps") {
			currentSection = "steps"
			inList = false
			continue
		} else if strings.HasPrefix(trimmed, "##") {
			currentSection = ""
			inList = false
			continue
		}

		// Skip empty lines between sections
		if trimmed == "" && !inList {
			continue
		}

		// Process content based on current section
		switch currentSection {
		case "summary":
			if trimmed != "" {
				if summary.Len() > 0 {
					summary.WriteString(" ")
				}
				summary.WriteString(trimmed)
			}
		case "principles":
			if strings.HasPrefix(trimmed, "- ") {
				inList = true
				keyPrinciples = append(keyPrinciples, strings.TrimPrefix(trimmed, "- "))
			} else if inList && trimmed != "" {
				// Continuation of previous item
				if len(keyPrinciples) > 0 {
					keyPrinciples[len(keyPrinciples)-1] += " " + trimmed
				}
			}
		case "steps":
			if strings.HasPrefix(trimmed, "- ") {
				inList = true
				remediationSteps = append(remediationSteps, strings.TrimPrefix(trimmed, "- "))
			} else if inList && trimmed != "" {
				// Continuation of previous item
				if len(remediationSteps) > 0 {
					remediationSteps[len(remediationSteps)-1] += " " + trimmed
				}
			}
		}
	}

	return map[string]interface{}{
		"summary":           summary.String(),
		"key_principles":    keyPrinciples,
		"remediation_steps": remediationSteps,
	}
}

// extractDataPath extracts source, sink, and stack trace information from the flaw
func extractDataPath(flaw *PipelineFlaw) map[string]interface{} {
	dataPath := map[string]interface{}{}

	// Extract stack trace if available
	stackTrace := []map[string]interface{}{}
	if len(flaw.StackDumps.StackDump) > 0 && len(flaw.StackDumps.StackDump[0].Frame) > 0 {
		for _, frame := range flaw.StackDumps.StackDump[0].Frame {
			stackFrame := map[string]interface{}{
				"frame_id": frame.FrameID,
				"function": frame.FunctionName,
				"file":     frame.SourceFile,
				"line":     frame.SourceLine,
			}
			if frame.QualifiedFunctionName != "" {
				stackFrame["qualified_function"] = frame.QualifiedFunctionName
			}
			stackTrace = append(stackTrace, stackFrame)
		}
	}

	if len(stackTrace) > 0 {
		dataPath["stack_trace"] = stackTrace
	}

	// Add sink details (where the vulnerability occurs)
	if flaw.Files.SourceFile.File != "" {
		dataPath["sink"] = map[string]interface{}{
			"file":     flaw.Files.SourceFile.File,
			"line":     flaw.Files.SourceFile.Line,
			"function": flaw.Files.SourceFile.FunctionName,
		}
	}

	return dataPath
}

// formatRemediationStepsForLLM formats remediation steps into a structured string for LLM consumption
func formatRemediationStepsForLLM(steps []string) string {
	if len(steps) == 0 {
		return "- Apply security best practices to fix this vulnerability"
	}

	var formatted strings.Builder
	for i, step := range steps {
		formatted.WriteString(fmt.Sprintf("%d. %s\n", i+1, step))
	}
	return formatted.String()
}

// formatRemediationGuidanceResponse formats the remediation guidance into a structured JSON response
func formatRemediationGuidanceResponse(req *RemediationGuidanceRequest, cweID int, flaw *PipelineFlaw, language, sourceFile, guidance string) map[string]interface{} {
	severityText := "Unknown"
	severityMap := map[int]string{
		0: "Informational",
		1: "Very Low",
		2: "Low",
		3: "Medium",
		4: "High",
		5: "Very High",
	}
	if sev, ok := severityMap[flaw.Severity]; ok {
		severityText = sev
	}

	// Parse the markdown guidance into structured sections
	parsedGuidance := parseMarkdownGuidance(guidance)

	// Build flaw details
	flawDetails := map[string]interface{}{
		"flaw_id":        req.FlawID,
		"issue_id":       flaw.IssueID,
		"cwe_id":         cweID,
		"issue_type":     flaw.IssueType,
		"severity":       strings.ToLower(severityText),
		"severity_score": flaw.Severity,
	}

	if sourceFile != "" {
		flawDetails["file_path"] = sourceFile
		flawDetails["line"] = flaw.Files.SourceFile.Line
	}

	if flaw.Files.SourceFile.FunctionName != "" {
		flawDetails["function_name"] = flaw.Files.SourceFile.FunctionName
	}

	if language != "" {
		flawDetails["language"] = language
	}

	// Build remediation guidance section (for user)
	remediationGuidance := map[string]interface{}{
		"summary":           parsedGuidance["summary"],
		"key_principles":    parsedGuidance["key_principles"],
		"language_specific": language != "",
	}

	// Build data path section
	dataPath := extractDataPath(flaw)

	// Build detailed instructions for the LLM using the remediation steps
	// Extract remediation steps to use as LLM guidance
	remediationSteps, _ := parsedGuidance["remediation_steps"].([]string)

	llmInstructions := fmt.Sprintf(`You must immediately analyze this vulnerability and present a fix summary to the user.

Read the vulnerable code at %s line %d and apply the following remediation steps:

%s

Then present your recommendation in this format:

VULNERABILITY SUMMARY:
[1-2 sentences explaining what's wrong and why it's dangerous]

PROPOSED SOLUTION:
[Brief description of the approach to fix this - e.g., "Add input validation before file operations"]

REQUIRED CHANGES:
• [List specific changes needed, e.g., "Add validatePath() function to sanitize user input"]
• [e.g., "Modify ConfigureLogging() to call validation before os.OpenFile()"]
• [e.g., "Add filepath import package"]

IMPACT:
- Files affected: [number and names]
- Estimated lines added: [rough estimate]
- Breaking changes: [Yes/No and explanation]

Do NOT include large code blocks or complete implementations. Focus on describing WHAT needs to change, not showing all the code. The developer will implement based on your guidance.`,
		sourceFile,
		flaw.Files.SourceFile.Line,
		formatRemediationStepsForLLM(remediationSteps))

	nextSteps := map[string]interface{}{
		"instructions_for_llm": llmInstructions,
	}

	// Build the complete structured response
	structuredResponse := map[string]interface{}{
		"flaw_details":         flawDetails,
		"remediation_guidance": remediationGuidance,
		"data_path":            dataPath,
		"next_steps":           nextSteps,
	}

	// Marshal to JSON for clean formatting
	responseJSON, err := json.Marshal(structuredResponse)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf("Error formatting JSON response: %v", err),
			}},
		}
	}

	// Return MCP-formatted response with JSON content
	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(responseJSON),
			},
		},
	}
}
