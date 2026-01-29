package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const GetLocalSCAResultsToolName = "get-local-sca-results"

// Auto-register this tool when the package is imported
func init() {
	RegisterTool(GetLocalSCAResultsToolName, func() ToolImplementation {
		return NewGetLocalSCAResultsTool()
	})
}

// GetLocalSCAResultsTool provides the get-local-sca-results tool
type GetLocalSCAResultsTool struct{}

// NewGetLocalSCAResultsTool creates a new get local SCA results tool
func NewGetLocalSCAResultsTool() *GetLocalSCAResultsTool {
	return &GetLocalSCAResultsTool{}
}

// Initialize sets up the tool
func (t *GetLocalSCAResultsTool) Initialize() error {
	log.Printf("Initializing tool: %s", GetLocalSCAResultsToolName)
	return nil
}

// RegisterHandlers registers the get local SCA results handler
func (t *GetLocalSCAResultsTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", GetLocalSCAResultsToolName)
	registry.RegisterHandler(GetLocalSCAResultsToolName, t.handleGetLocalSCAResults)
	return nil
}

// Shutdown cleans up tool resources
func (t *GetLocalSCAResultsTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", GetLocalSCAResultsToolName)
	return nil
}

// GetLocalSCAResultsRequest represents the parsed parameters for get-local-sca-results
type GetLocalSCAResultsRequest struct {
	ApplicationPath string
}

// parseGetLocalSCAResultsRequest extracts and validates parameters from the raw args map
func parseGetLocalSCAResultsRequest(args map[string]interface{}) (*GetLocalSCAResultsRequest, error) {
	req := &GetLocalSCAResultsRequest{}

	// Extract application_path (required)
	appPath, ok := args["application_path"].(string)
	if !ok || appPath == "" {
		return nil, fmt.Errorf("application_path is required and must be a non-empty string")
	}
	req.ApplicationPath = appPath

	return req, nil
}

// SCAResults represents the JSON structure from Veracode SCA scan
type SCAResults struct {
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

// handleGetLocalSCAResults retrieves and formats local SCA scan results
func (t *GetLocalSCAResultsTool) handleGetLocalSCAResults(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseGetLocalSCAResultsRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Locate the results directory
	outputDir := filepath.Join(req.ApplicationPath, ".veracode_sca")
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

	var scaResults SCAResults
	err = json.Unmarshal(resultsData, &scaResults)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to parse results file: %v", err),
		}, nil
	}

	// Format and return the response
	return formatSCAResultsResponse(req.ApplicationPath, resultsFile, &scaResults), nil
}

// scaSummary holds summary statistics for SCA results
type scaSummary struct {
	Total          int
	Critical       int
	High           int
	Medium         int
	Low            int
	Informational  int
	WithEPSS       int
	TotalArtifacts int
	TotalVulns     int
}

// buildSCASummary builds summary statistics from SCA results
func buildSCASummary(results *SCAResults) scaSummary {
	summary := scaSummary{}
	artifactMap := make(map[string]bool)
	vulnMap := make(map[string]bool)

	summary.Total = len(results.Vulnerabilities.Matches)

	for _, match := range results.Vulnerabilities.Matches {
		artifactMap[match.Artifact.ID] = true
		vulnMap[match.Vulnerability.ID] = true

		switch strings.ToLower(match.Vulnerability.Severity) {
		case "critical":
			summary.Critical++
		case "high":
			summary.High++
		case "medium":
			summary.Medium++
		case "low":
			summary.Low++
		default:
			summary.Informational++
		}

		if len(match.Vulnerability.EPSS) > 0 && match.Vulnerability.EPSS[0].EPSS > 0 {
			summary.WithEPSS++
		}
	}
	summary.TotalArtifacts = len(artifactMap)
	summary.TotalVulns = len(vulnMap)
	return summary
}

// buildSCAHeader builds the text header for SCA results
func buildSCAHeader(appPath, resultsFile string, summary scaSummary) string {
	return fmt.Sprintf(`Local SCA Scan Results
===================

Application Path: %s
Results File: %s

Total Vulnerability Matches: %d
Unique Vulnerabilities: %d
Vulnerable Components: %d

Severity Breakdown:
- Critical: %d
- High: %d
- Medium: %d
- Low: %d
- Informational: %d

EPSS Data Available: %d vulnerabilities

`,
		appPath,
		filepath.Base(resultsFile),
		summary.Total,
		summary.TotalVulns,
		summary.TotalArtifacts,
		summary.Critical,
		summary.High,
		summary.Medium,
		summary.Low,
		summary.Informational,
		summary.WithEPSS,
	)
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
			"state":             match.Vulnerability.Fix.State,
			"versions":          match.Vulnerability.Fix.Versions,
			"suggested_version": "",
		},
	}

	if len(match.MatchDetails) > 0 && match.MatchDetails[0].Fix.SuggestedVersion != "" {
		finding["fix"].(map[string]interface{})["suggested_version"] = match.MatchDetails[0].Fix.SuggestedVersion
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

// addCWEData adds CWE data to a finding if available
func addCWEData(finding map[string]interface{}, cweData []SCACWE) {
	if len(cweData) == 0 {
		return
	}
	cweList := make([]string, 0)
	for _, cwe := range cweData {
		if cwe.CWE != "" {
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

// addRelatedCVEs adds related CVEs to a finding if available
func addRelatedCVEs(finding map[string]interface{}, relatedVulns []SCARelatedVuln) {
	if len(relatedVulns) == 0 {
		return
	}
	relatedCVEs := make([]string, 0)
	for _, rv := range relatedVulns {
		if strings.HasPrefix(rv.ID, "CVE-") {
			relatedCVEs = append(relatedCVEs, rv.ID)
		}
	}
	if len(relatedCVEs) > 0 {
		finding["related_cves"] = relatedCVEs
	}
}

// formatSCAResultsResponse formats the SCA results into an MCP response
func formatSCAResultsResponse(appPath, resultsFile string, results *SCAResults) map[string]interface{} {
	summary := buildSCASummary(results)
	header := buildSCAHeader(appPath, resultsFile, summary)

	// Convert matches to LLM-friendly format
	llmFriendlyFindings := make([]map[string]interface{}, 0, len(results.Vulnerabilities.Matches))
	for _, match := range results.Vulnerabilities.Matches {
		llmFriendlyFindings = append(llmFriendlyFindings, convertMatchToFinding(match))
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
				"critical":      summary.Critical,
				"high":          summary.High,
				"medium":        summary.Medium,
				"low":           summary.Low,
				"informational": summary.Informational,
			},
			"epss_data_available": summary.WithEPSS,
		},
		"findings": llmFriendlyFindings,
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
