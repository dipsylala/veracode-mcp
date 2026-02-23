package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const GetLocalSCAFindingsToolName = "local-sca-findings"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(GetLocalSCAFindingsToolName, handleGetLocalSCAFindings)
}

// GetLocalSCAFindingsRequest represents the parsed parameters for local-sca-findings
type GetLocalSCAFindingsRequest struct {
	ApplicationPath string
	CVE             string `json:"cve,omitempty"`
	ComponentName   string `json:"component_name,omitempty"`
	SeverityGTE     *int   `json:"severity_gte,omitempty"`
	Size            int    `json:"size,omitempty"`
	Page            int    `json:"page,omitempty"`
}

// parseGetLocalSCAFindingsRequest extracts and validates parameters from the raw args map
func parseGetLocalSCAFindingsRequest(args map[string]interface{}) (*GetLocalSCAFindingsRequest, error) {
	req := &GetLocalSCAFindingsRequest{}

	// Extract required fields
	var err error
	req.ApplicationPath, err = extractRequiredString(args, "application_path")
	if err != nil {
		return nil, err
	}

	// Extract filter fields
	if cve, ok := args["cve"].(string); ok && cve != "" {
		req.CVE = cve
	}
	if component, ok := args["component_name"].(string); ok && component != "" {
		req.ComponentName = component
	}
	if sevGTE, ok := args["severity_gte"].(float64); ok {
		sevInt := int(sevGTE)
		req.SeverityGTE = &sevInt
	}

	// Extract optional fields with defaults
	req.Size = extractInt(args, "page_size", 10)
	req.Page = extractInt(args, "page", 0)

	// Validate pagination bounds
	if err := validatePaginationParams(req.Size, req.Page); err != nil {
		return nil, err
	}

	return req, nil
}

// SCAFindings represents the JSON structure from Veracode SCA scan
type SCAFindings struct {
	Vulnerabilities SCAVulnerabilities `json:"vulnerabilities"`
}

// SCAVulnerabilities contains the vulnerability matches
type SCAVulnerabilities struct {
	Matches []SCAMatch `json:"matches"`
}

// SCAMatch represents a single vulnerability match
type SCAMatch struct {
	Artifact      SCAArtifact      `json:"artifact"`
	Vulnerability SCAVulnerability `json:"vulnerability"`
	MatchDetails  []SCAMatchDetail `json:"matchDetails"`
	RelatedVulns  []SCARelatedVuln `json:"relatedVulnerabilities"`
}

// SCAArtifact represents the vulnerable artifact/component
type SCAArtifact struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Version   string          `json:"version"`
	Type      string          `json:"type"`
	Language  string          `json:"language"`
	Licenses  []string        `json:"licenses"`
	PURL      string          `json:"purl"`
	Locations []SCALocation   `json:"locations"`
	Metadata  SCAArtifactMeta `json:"metadata"`
}

// SCALocation represents where the artifact was found
type SCALocation struct {
	Path       string `json:"path"`
	AccessPath string `json:"accessPath"`
}

// SCAArtifactMeta contains artifact metadata
type SCAArtifactMeta struct {
	PomGroupID    string `json:"pomGroupID"`
	PomArtifactID string `json:"pomArtifactID"`
}

// SCAVulnerability represents the vulnerability details
type SCAVulnerability struct {
	ID          string    `json:"id"`
	Severity    string    `json:"severity"`
	Description string    `json:"description"`
	DataSource  string    `json:"dataSource"`
	CVSS        []SCACVSS `json:"cvss"`
	CWEs        []SCACWE  `json:"cwes"`
	EPSS        []SCAEPSS `json:"epss"`
	Fix         SCAFix    `json:"fix"`
	Risk        float64   `json:"risk"`
}

// SCACVSS represents CVSS scoring information
type SCACVSS struct {
	Version string         `json:"version"`
	Vector  string         `json:"vector"`
	Metrics SCACVSSMetrics `json:"metrics"`
}

// SCACVSSMetrics contains CVSS metric scores
type SCACVSSMetrics struct {
	BaseScore           float64 `json:"baseScore"`
	ExploitabilityScore float64 `json:"exploitabilityScore"`
	ImpactScore         float64 `json:"impactScore"`
}

// SCACWE represents CWE information
type SCACWE struct {
	CVE    string `json:"cve"`
	CWE    string `json:"cwe"`
	Source string `json:"source"`
}

// SCAEPSS represents EPSS scoring information
type SCAEPSS struct {
	CVE        string  `json:"cve"`
	Date       string  `json:"date"`
	EPSS       float64 `json:"epss"`
	Percentile float64 `json:"percentile"`
}

// SCAFix represents fix information
type SCAFix struct {
	State    string   `json:"state"`
	Versions []string `json:"versions"`
}

// SCAMatchDetail contains match details
type SCAMatchDetail struct {
	Fix SCAMatchFix `json:"fix"`
}

// SCAMatchFix contains suggested fix version
type SCAMatchFix struct {
	SuggestedVersion string `json:"suggestedVersion"`
}

// SCARelatedVuln represents related vulnerabilities
type SCARelatedVuln struct {
	ID          string    `json:"id"`
	Severity    string    `json:"severity"`
	Description string    `json:"description"`
	CVSS        []SCACVSS `json:"cvss"`
	CWEs        []SCACWE  `json:"cwes"`
	EPSS        []SCAEPSS `json:"epss"`
}

// handleGetLocalSCAFindings retrieves and formats local SCA scan findings
func handleGetLocalSCAFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseGetLocalSCAFindingsRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Locate the results directory
	outputDir := veracodeWorkDir(req.ApplicationPath, "sca")
	resultsFile := filepath.Join(outputDir, "veracode.json")

	// Check if results file exists
	if _, statErr := os.Stat(resultsFile); os.IsNotExist(statErr) {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Local SCA Scan Results
===================

Application Path: %s
Results File: %s

❌ No results found

The results file does not exist. Run a local SCA scan using the run-sca-scan tool first.
`, req.ApplicationPath, resultsFile),
			}},
		}, nil
	}

	// Read and parse the results file
	// #nosec G304 -- resultsFile is constructed from validated application path
	resultsData, err := os.ReadFile(resultsFile)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to read results file: %v", err),
		}, nil
	}

	var scaResults SCAFindings
	err = json.Unmarshal(resultsData, &scaResults)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to parse results file: %v", err),
		}, nil
	}

	// Format and return the response
	return formatSCAFindingsResponse(req.ApplicationPath, resultsFile, &scaResults, req), nil
}

// scaSummary holds summary statistics for SCA findings
type scaSummary struct {
	Total          int
	Critical       int
	High           int
	Medium         int
	Low            int
	WithEPSS       int
	TotalArtifacts int
	TotalVulns     int
}

// buildSCASummary builds summary statistics from SCA findings
func buildSCASummary(results *SCAFindings) scaSummary {
	summary := scaSummary{}
	artifactMap := make(map[string]bool)
	vulnMap := make(map[string]bool)

	summary.Total = len(results.Vulnerabilities.Matches)

	for _, match := range results.Vulnerabilities.Matches {
		artifactMap[match.Artifact.ID] = true
		vulnMap[match.Vulnerability.ID] = true

		sevNorm := strings.ToLower(strings.ReplaceAll(match.Vulnerability.Severity, " ", ""))
		switch sevNorm {
		case "critical":
			summary.Critical++
		case "high":
			summary.High++
		case "medium":
			summary.Medium++
		case "low":
			summary.Low++
		}

		if len(match.Vulnerability.EPSS) > 0 && match.Vulnerability.EPSS[0].EPSS > 0 {
			summary.WithEPSS++
		}
	}
	summary.TotalArtifacts = len(artifactMap)
	summary.TotalVulns = len(vulnMap)
	return summary
}

// severityToInt converts severity string to numeric value for filtering
func severityToInt(severity string) int {
	sevNorm := strings.ToLower(strings.TrimSpace(severity))
	switch sevNorm {
	case "critical":
		return 5
	case "high":
		return 4
	case "medium":
		return 3
	case "low":
		return 2
	default:
		return 0
	}
}

// matchesFilters checks if an SCA match passes all active filters
func matchesFilters(match SCAMatch, req *GetLocalSCAFindingsRequest) bool {
	// Filter by CVE
	if req.CVE != "" {
		found := false
		// Check primary vulnerability ID
		if strings.EqualFold(match.Vulnerability.ID, req.CVE) {
			found = true
		}
		// Check EPSS CVE
		if !found && len(match.Vulnerability.EPSS) > 0 {
			for _, epss := range match.Vulnerability.EPSS {
				if strings.EqualFold(epss.CVE, req.CVE) {
					found = true
					break
				}
			}
		}
		// Check related vulnerabilities
		if !found {
			for _, rv := range match.RelatedVulns {
				if strings.EqualFold(rv.ID, req.CVE) {
					found = true
					break
				}
			}
		}
		if !found {
			return false
		}
	}

	// Filter by component name (partial match, case-insensitive)
	if req.ComponentName != "" {
		if !strings.Contains(strings.ToLower(match.Artifact.Name), strings.ToLower(req.ComponentName)) {
			return false
		}
	}

	// Filter by minimum severity
	if req.SeverityGTE != nil {
		matchSev := severityToInt(match.Vulnerability.Severity)
		if matchSev < *req.SeverityGTE {
			return false
		}
	}

	return true
}

// convertMatchToFinding converts an SCA match to LLM-friendly format
func convertMatchToFinding(match SCAMatch) map[string]interface{} {
	finding := map[string]interface{}{
		"vulnerability_id": match.Vulnerability.ID,
		"severity":         strings.ToLower(match.Vulnerability.Severity),
		"description":      match.Vulnerability.Description,
		"risk_score":       match.Vulnerability.Risk,
		"data_source":      match.Vulnerability.DataSource,
		"component": map[string]interface{}{
			"name":      match.Artifact.Name,
			"version":   match.Artifact.Version,
			"type":      match.Artifact.Type,
			"language":  match.Artifact.Language,
			"purl":      match.Artifact.PURL,
			"licenses":  match.Artifact.Licenses,
			"locations": match.Artifact.Locations,
		},
		"fix": map[string]interface{}{
			"state":               match.Vulnerability.Fix.State,
			"versions":            match.Vulnerability.Fix.Versions,
			"recommended_version": "",
		},
	}

	if len(match.MatchDetails) > 0 && match.MatchDetails[0].Fix.SuggestedVersion != "" {
		finding["fix"].(map[string]interface{})["recommended_version"] = match.MatchDetails[0].Fix.SuggestedVersion
	}

	addCVSSData(finding, match.Vulnerability.CVSS)
	addCWEData(finding, match.Vulnerability.CWEs)
	addEPSSData(finding, match.Vulnerability.EPSS)
	addRelatedCVEs(finding, match.RelatedVulns)

	return finding
}

// addCVSSData adds CVSS data to a finding if available
func addCVSSData(finding map[string]interface{}, cvssData []SCACVSS) {
	if len(cvssData) == 0 {
		return
	}
	cvss := make([]map[string]interface{}, 0, len(cvssData))
	for _, c := range cvssData {
		cvss = append(cvss, map[string]interface{}{
			"version":              c.Version,
			"vector":               c.Vector,
			"base_score":           c.Metrics.BaseScore,
			"exploitability_score": c.Metrics.ExploitabilityScore,
			"impact_score":         c.Metrics.ImpactScore,
		})
	}
	finding["cvss"] = cvss
}

// addCWEData adds CWE data to a finding if available (deduplicates)
func addCWEData(finding map[string]interface{}, cweData []SCACWE) {
	if len(cweData) == 0 {
		return
	}
	// Use a map to track unique CWEs
	cweMap := make(map[string]bool)
	cweList := make([]string, 0)
	for _, cwe := range cweData {
		if cwe.CWE != "" && !cweMap[cwe.CWE] {
			cweMap[cwe.CWE] = true
			cweList = append(cweList, cwe.CWE)
		}
	}
	if len(cweList) > 0 {
		finding["cwes"] = cweList
	}
}

// addEPSSData adds EPSS data to a finding if available
func addEPSSData(finding map[string]interface{}, epssData []SCAEPSS) {
	if len(epssData) == 0 {
		return
	}
	epss := epssData[0]
	finding["epss"] = map[string]interface{}{
		"cve":        epss.CVE,
		"score":      epss.EPSS,
		"percentile": epss.Percentile,
		"date":       epss.Date,
	}
}

// addRelatedCVEs adds related CVEs to a finding if available (deduplicates)
func addRelatedCVEs(finding map[string]interface{}, relatedVulns []SCARelatedVuln) {
	if len(relatedVulns) == 0 {
		return
	}
	// Use a map to track unique CVEs
	cveMap := make(map[string]bool)
	relatedCVEs := make([]string, 0)
	for _, rv := range relatedVulns {
		if strings.HasPrefix(rv.ID, "CVE-") && !cveMap[rv.ID] {
			cveMap[rv.ID] = true
			relatedCVEs = append(relatedCVEs, rv.ID)
		}
	}
	if len(relatedCVEs) > 0 {
		finding["related_cves"] = relatedCVEs
	}
}

// formatSCAFindingsResponse formats the SCA findings into an MCP response
func formatSCAFindingsResponse(appPath, resultsFile string, results *SCAFindings, req *GetLocalSCAFindingsRequest) map[string]interface{} {
	// Apply filters and convert matches to LLM-friendly format
	allFindings := make([]map[string]interface{}, 0, len(results.Vulnerabilities.Matches))
	filteredMatches := make([]SCAMatch, 0, len(results.Vulnerabilities.Matches))
	for _, match := range results.Vulnerabilities.Matches {
		if matchesFilters(match, req) {
			allFindings = append(allFindings, convertMatchToFinding(match))
			filteredMatches = append(filteredMatches, match)
		}
	}

	// Build summary from filtered results
	filteredResults := &SCAFindings{
		Vulnerabilities: SCAVulnerabilities{
			Matches: filteredMatches,
		},
	}
	summary := buildSCASummary(filteredResults)

	// Sort findings by severity (descending), then component name (ascending)
	sort.Slice(allFindings, func(i, j int) bool {
		sevI := allFindings[i]["severity"].(string)
		sevJ := allFindings[j]["severity"].(string)
		sevIntI := severityToInt(sevI)
		sevIntJ := severityToInt(sevJ)

		// Sort by severity descending (higher severity first)
		if sevIntI != sevIntJ {
			return sevIntI > sevIntJ
		}

		// If same severity, sort by component name ascending
		compI := allFindings[i]["component"].(map[string]interface{})["name"].(string)
		compJ := allFindings[j]["component"].(map[string]interface{})["name"].(string)
		return strings.ToLower(compI) < strings.ToLower(compJ)
	})

	// Apply pagination
	totalFindings := len(allFindings)

	// Ensure we have a valid page size
	if req.Size <= 0 {
		req.Size = 10 // Fallback default
	}

	startIdx := req.Page * req.Size
	endIdx := startIdx + req.Size
	if startIdx > totalFindings {
		startIdx = totalFindings
	}
	if endIdx > totalFindings {
		endIdx = totalFindings
	}

	// Get paginated findings
	llmFriendlyFindings := allFindings[startIdx:endIdx]

	// Build pagination info
	totalPages := 0
	if req.Size > 0 {
		totalPages = (totalFindings + req.Size - 1) / req.Size
	}
	pagination := map[string]interface{}{
		"current_page":   req.Page,
		"page_size":      req.Size,
		"total_elements": totalFindings,
		"total_pages":    totalPages,
		"has_next":       endIdx < totalFindings,
		"has_previous":   req.Page > 0,
	}

	// Build filters info to show what's active
	filters := map[string]interface{}{}
	if req.CVE != "" {
		filters["cve"] = req.CVE
	}
	if req.ComponentName != "" {
		filters["component_name"] = req.ComponentName
	}
	if req.SeverityGTE != nil {
		severityNames := map[int]string{
			2: "low",
			3: "medium",
			4: "high",
			5: "critical",
		}
		if name, ok := severityNames[*req.SeverityGTE]; ok {
			filters["severity_gte"] = name
		}
	}

	// Build the response structure
	responseData := map[string]interface{}{
		"application": map[string]string{
			"name": filepath.Base(appPath),
			"path": appPath,
		},
		"summary": map[string]interface{}{
			"total_matches":          summary.Total,
			"unique_vulnerabilities": summary.TotalVulns,
			"vulnerable_components":  summary.TotalArtifacts,
			"by_severity": map[string]int{
				"critical": summary.Critical,
				"high":     summary.High,
				"medium":   summary.Medium,
				"low":      summary.Low,
			},
			"epss_data_available": summary.WithEPSS,
		},
		"pagination": pagination,
		"findings":   llmFriendlyFindings,
	}

	// Include filters only if any are active
	if len(filters) > 0 {
		responseData["filters"] = filters
	}

	// Marshal response to JSON string
	responseJSON, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf("Error formatting response: %v", err),
			}},
		}
	}

	// Build result with JSON content and structuredContent
	result := map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(responseJSON),
			},
		},
	}

	// Include structuredContent for MCP Apps UI
	// The converter will handle checking if client supports UI
	result["structuredContent"] = responseData

	return result
}

// transformSCASeverity converts string severity to normalized lowercase
func transformSCASeverity(severity string) string {
	normalized := strings.ToLower(strings.TrimSpace(severity))
	switch normalized {
	case "critical", "high", "medium", "low", "negligible":
		return normalized
	default:
		return "unknown"
	}
}
