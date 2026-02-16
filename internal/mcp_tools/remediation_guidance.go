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
	RegisterMCPTool(RemediationGuidanceToolName, handleGetRemediationGuidance)
}

// RemediationGuidanceRequest represents the parsed parameters for remediation-guidance
type RemediationGuidanceRequest struct {
	ApplicationPath string `json:"application_path"`
	FlawID          *FlawIDComponents
}

// parseRemediationGuidanceRequest extracts and validates parameters from the raw args map
func parseRemediationGuidanceRequest(args map[string]interface{}) (*RemediationGuidanceRequest, error) {
	req := &RemediationGuidanceRequest{}

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

// findAllFlawsInPipelineResults finds all flaws with a specific issue_id
// Use this when you need to handle potential duplicates
func findAllFlawsInPipelineResults(results *PipelineScanResults, flawID int) []PipelineFlaw {
	var matches []PipelineFlaw
	for i := range results.Findings {
		if results.Findings[i].IssueID == flawID {
			matches = append(matches, results.Findings[i])
		}
	}
	return matches
}

// formatOccurrenceList formats a list of occurrences for error messages
func formatOccurrenceList(flaws []PipelineFlaw, issueID int) string {
	var result strings.Builder
	for i := 0; i < len(flaws); i++ {
		flawIDStr := fmt.Sprintf("%d", issueID)
		if i > 0 {
			flawIDStr = fmt.Sprintf("%d-%d", issueID, i+1)
		}
		result.WriteString(fmt.Sprintf("- flaw_id %s: CWE-%s at %s:%d\n",
			flawIDStr, flaws[i].CWEID, flaws[i].Files.SourceFile.File, flaws[i].Files.SourceFile.Line))
	}
	return result.String()
}

// errorResponse creates a standardized error response for MCP
func errorResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": message,
		}},
	}
}

// buildNoResultsError creates error message when pipeline results are not found
func buildNoResultsError(appPath string, issueID int, err error) string {
	return fmt.Sprintf(`Remediation Guidance Lookup
============================

Application Path: %s
Flaw ID: %d

❌ No pipeline results found

%v

Please run a pipeline scan first using the pipeline-scan tool.
`, appPath, issueID, err)
}

// buildFlawNotFoundError creates error message when flaw is not found
func buildFlawNotFoundError(appPath string, issueID int) string {
	return fmt.Sprintf(`Remediation Guidance Lookup
============================

Application Path: %s
Flaw ID: %d

❌ Flaw not found

The specified flaw ID was not found in the pipeline scan results.

Please verify that:
1. The flaw ID is correct
2. The pipeline scan completed successfully
3. The flaw exists in the most recent scan
`, appPath, issueID)
}

// buildOccurrenceNotFoundError creates error message when requested occurrence doesn't exist
func buildOccurrenceNotFoundError(appPath string, issueID, occurrence, matchCount int, matches []PipelineFlaw) string {
	return fmt.Sprintf(`Remediation Guidance Lookup
============================

Application Path: %s
Flaw ID: %d-%d

❌ Occurrence not found

Issue ID %d has %d occurrence(s), but you requested occurrence %d.

Available occurrences:
%s
`, appPath, issueID, occurrence, issueID, matchCount, occurrence, formatOccurrenceList(matches, issueID))
}

// buildDuplicateNote creates a note about duplicate issue_ids when there are multiple occurrences
func buildDuplicateNote(issueID int, matches []PipelineFlaw, selectedOccurrence int) string {
	if len(matches) <= 1 {
		return ""
	}

	log.Printf("Issue ID %d has %d occurrences. Using occurrence %d at %s:%d",
		issueID, len(matches), selectedOccurrence,
		matches[selectedOccurrence-1].Files.SourceFile.File, matches[selectedOccurrence-1].Files.SourceFile.Line)

	note := fmt.Sprintf("\n\n⚠️  Note: Issue ID %d appears %d times in the scan results.\n\nAll occurrences:\n", issueID, len(matches))
	for i := 0; i < len(matches); i++ {
		flawIDStr := fmt.Sprintf("%d", issueID)
		if i > 0 {
			flawIDStr = fmt.Sprintf("%d-%d", issueID, i+1)
		}
		marker := "  "
		if i == selectedOccurrence-1 {
			marker = "→ "
		}
		note += fmt.Sprintf("%s- flaw_id %s: CWE-%s at %s:%d\n",
			marker, flawIDStr, matches[i].CWEID, matches[i].Files.SourceFile.File, matches[i].Files.SourceFile.Line)
	}
	return note
}

// loadAndParsePipelineResults loads and parses the pipeline results file
func loadAndParsePipelineResults(appPath string) (*PipelineScanResults, error) {
	outputDir := filepath.Join(appPath, ".veracode", "pipeline")
	resultsFile, err := findMostRecentResultsFile(outputDir)
	if err != nil {
		return nil, err
	}

	// #nosec G304 -- resultsFile is from findMostRecentResultsFile which validates the directory
	resultsData, err := os.ReadFile(resultsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read results file: %w", err)
	}

	var scanResults PipelineScanResults
	err = json.Unmarshal(resultsData, &scanResults)
	if err != nil {
		return nil, fmt.Errorf("failed to parse results file: %w", err)
	}

	return &scanResults, nil
}

func handleGetRemediationGuidance(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseRemediationGuidanceRequest(args)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}

	// Load and parse pipeline results
	scanResults, err := loadAndParsePipelineResults(req.ApplicationPath)
	if err != nil {
		return errorResponse(buildNoResultsError(req.ApplicationPath, req.FlawID.IssueID, err)), nil
	}

	// Find all flaws with this issue_id (there may be duplicates)
	matches := findAllFlawsInPipelineResults(scanResults, req.FlawID.IssueID)
	if len(matches) == 0 {
		return errorResponse(buildFlawNotFoundError(req.ApplicationPath, req.FlawID.IssueID)), nil
	}

	// Select the correct occurrence
	if req.FlawID.Occurrence > len(matches) {
		return errorResponse(buildOccurrenceNotFoundError(
			req.ApplicationPath, req.FlawID.IssueID, req.FlawID.Occurrence, len(matches), matches)), nil
	}

	// Get the requested occurrence and build duplicate note if needed
	flaw := &matches[req.FlawID.Occurrence-1]
	duplicateNote := buildDuplicateNote(req.FlawID.IssueID, matches, req.FlawID.Occurrence)

	// Extract CWE ID from the flaw (it's a string in the JSON)
	cweID, err := strconv.Atoi(flaw.CWEID)
	if err != nil {
		return errorResponse(fmt.Sprintf(`Remediation Guidance Lookup
============================

Application Path: %s
Flaw ID: %d

❌ Error: Invalid CWE ID format: %s
`, req.ApplicationPath, req.FlawID.IssueID, flaw.CWEID)), nil
	}

	// Extract source file from the flaw
	sourceFile := flaw.Files.SourceFile.File
	if sourceFile == "" {
		log.Printf("Warning: No source file found for flaw %d", req.FlawID.IssueID)
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
		return errorResponse(fmt.Sprintf("Remediation Guidance Lookup\n============================\n\nApplication Path: %s\nFlaw ID: %d\nCWE ID: %d\nSource File: %s\nDetected Language: %s\n\n❌ Error: %v\n", req.ApplicationPath, req.FlawID.IssueID, cweID, sourceFile, language, err)), nil
	}

	// Read the guidance file from embedded FS
	guidanceContent, err := remediationGuidanceFS.ReadFile(guidancePath)
	if err != nil {
		return errorResponse(fmt.Sprintf("Remediation Guidance Lookup\n============================\n\nApplication Path: %s\nFlaw ID: %d\nCWE ID: %d\nSource File: %s\nDetected Language: %s\n\n❌ Error: Failed to read guidance file\n\n%v\n", req.ApplicationPath, req.FlawID.IssueID, cweID, sourceFile, language, err)), nil
	}

	// Format and return the guidance
	return formatRemediationGuidanceResponse(req, cweID, flaw, language, sourceFile, string(guidanceContent), duplicateNote), nil
}

// detectMarkdownSection determines which section a line represents
func detectMarkdownSection(trimmed string) string {
	switch {
	case strings.HasPrefix(trimmed, "## LLM Guidance"):
		return "summary"
	case strings.HasPrefix(trimmed, "## Key Principles"):
		return "principles"
	case strings.HasPrefix(trimmed, "## Remediation Steps"):
		return "steps"
	case strings.HasPrefix(trimmed, "## Safe Pattern"):
		return "code_samples"
	case strings.HasPrefix(trimmed, "##"):
		return ""
	default:
		return "content"
	}
}

// appendToSummary adds content to the summary builder
func appendToSummary(summaryBuilder *strings.Builder, trimmed string) {
	if summaryBuilder.Len() > 0 {
		summaryBuilder.WriteString(" ")
	}
	summaryBuilder.WriteString(trimmed)
}

// appendToListSection adds content to a list section (bullet points or continuation)
func appendToListSection(items *[]string, trimmed string) {
	if strings.HasPrefix(trimmed, "- ") {
		*items = append(*items, strings.TrimPrefix(trimmed, "- "))
	} else if len(*items) > 0 {
		(*items)[len(*items)-1] += " " + trimmed
	}
}

// parseMarkdownGuidance parses markdown guidance into structured sections
func parseMarkdownGuidance(markdown string) (summary string, keyPrinciples, remediationSteps []string, codeSamples string) {
	lines := strings.Split(markdown, "\n")
	var summaryBuilder strings.Builder
	var codeSamplesBuilder strings.Builder
	currentSection := ""

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Skip the main title and empty lines
		if strings.HasPrefix(trimmed, "# CWE-") || trimmed == "" {
			continue
		}

		// Detect section changes
		section := detectMarkdownSection(trimmed)
		if section != "content" {
			currentSection = section
			// For code samples, preserve the section heading
			if section == "code_samples" {
				if codeSamplesBuilder.Len() > 0 {
					codeSamplesBuilder.WriteString("\n\n")
				}
				codeSamplesBuilder.WriteString(trimmed)
				codeSamplesBuilder.WriteString("\n\n")
			}
			continue
		}

		// Process content based on current section
		switch currentSection {
		case "summary":
			appendToSummary(&summaryBuilder, trimmed)
		case "principles":
			appendToListSection(&keyPrinciples, trimmed)
		case "steps":
			appendToListSection(&remediationSteps, trimmed)
		case "code_samples":
			// Preserve original formatting for code samples
			codeSamplesBuilder.WriteString(line)
			codeSamplesBuilder.WriteString("\n")
		}
	}

	return summaryBuilder.String(), keyPrinciples, remediationSteps, strings.TrimSpace(codeSamplesBuilder.String())
}

// extractDataPath extracts source, sink, and stack trace information from the flaw
func extractDataPath(flaw *PipelineFlaw) map[string]interface{} {
	dataPath := map[string]interface{}{}

	// Extract stack trace if available
	if len(flaw.StackDumps.StackDump) > 0 && len(flaw.StackDumps.StackDump[0].Frame) > 0 {
		stackTrace := make([]map[string]interface{}, 0, len(flaw.StackDumps.StackDump[0].Frame))
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

// getSeverityText converts severity score to text
func getSeverityText(severity int) string {
	severityMap := map[int]string{
		0: "Informational",
		1: "Very Low",
		2: "Low",
		3: "Medium",
		4: "High",
		5: "Very High",
	}
	if text, ok := severityMap[severity]; ok {
		return text
	}
	return "Unknown"
}

// formatRemediationGuidanceResponse formats the remediation guidance into a structured JSON response
func formatRemediationGuidanceResponse(req *RemediationGuidanceRequest, cweID int, flaw *PipelineFlaw, language, sourceFile, guidance, duplicateNote string) map[string]interface{} {
	// Parse the markdown guidance into structured sections
	summary, keyPrinciples, remediationSteps, codeSamples := parseMarkdownGuidance(guidance)

	// Build flaw details
	flawDetails := map[string]interface{}{
		"flaw_id":        req.FlawID.IssueID,
		"issue_id":       flaw.IssueID,
		"cwe_id":         cweID,
		"issue_type":     flaw.IssueType,
		"severity":       strings.ToLower(getSeverityText(flaw.Severity)),
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

	// Build remediation guidance section
	remediationGuidance := map[string]interface{}{
		"summary":           summary,
		"key_principles":    keyPrinciples,
		"remediation_steps": remediationSteps,
		"language_specific": language != "",
	}

	// Add code samples if available
	if codeSamples != "" {
		remediationGuidance["code_samples"] = codeSamples
	}

	// Build complete structured response
	structuredResponse := map[string]interface{}{
		"flaw_details":         flawDetails,
		"remediation_guidance": remediationGuidance,
		"data_path":            extractDataPath(flaw),
		"next_steps": map[string]interface{}{
			"instructions_for_llm": buildLLMInstructions(sourceFile, flaw.Files.SourceFile.Line, remediationSteps) + duplicateNote,
		},
	}

	// Marshal to JSON for clean formatting
	responseJSON, err := json.Marshal(structuredResponse)
	if err != nil {
		return errorResponse(fmt.Sprintf("Error formatting JSON response: %v", err))
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

// buildLLMInstructions creates detailed instructions for the LLM
func buildLLMInstructions(sourceFile string, line int, remediationSteps []string) string {
	return fmt.Sprintf(`You must immediately analyze this vulnerability and present a fix summary to the user.

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
		line,
		formatRemediationStepsForLLM(remediationSteps))
}
