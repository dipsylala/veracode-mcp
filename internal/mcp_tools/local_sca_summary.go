package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const GetLocalSCASummaryToolName = "local-sca-summary"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(GetLocalSCASummaryToolName, handleGetLocalSCASummary)
}

// GetLocalSCASummaryRequest represents the parsed parameters for local-sca-summary
type GetLocalSCASummaryRequest struct {
	ApplicationPath string
	SeverityGTE     *int `json:"severity_gte,omitempty"`
	Size            int  `json:"page_size,omitempty"`
	Page            int  `json:"page,omitempty"`
}

// parseGetLocalSCASummaryRequest extracts and validates parameters from the raw args map
func parseGetLocalSCASummaryRequest(args map[string]interface{}) (*GetLocalSCASummaryRequest, error) {
	req := &GetLocalSCASummaryRequest{}

	var err error
	req.ApplicationPath, err = extractRequiredString(args, "application_path")
	if err != nil {
		return nil, err
	}

	if sevGTE, ok := args["severity_gte"].(float64); ok {
		sevInt := int(sevGTE)
		req.SeverityGTE = &sevInt
	}

	// Extract optional pagination fields with defaults
	req.Size = extractInt(args, "page_size", 10)
	req.Page = extractInt(args, "page", 0)

	if err := validatePaginationParams(req.Size, req.Page); err != nil {
		return nil, err
	}

	return req, nil
}

// componentKey is a unique identifier for a component (name + version)
type componentKey struct {
	Name    string
	Version string
}

// componentSummary holds aggregated vulnerability info for a single component version
type componentSummary struct {
	Name               string
	Version            string
	Type               string
	Language           string
	PURL               string
	RecommendedUpgrade string
	CVEs               []cveSummaryEntry
	BySeverity         map[string]int
	MaxEPSS            float64
}

// cveSummaryEntry holds per-CVE info within a component summary
type cveSummaryEntry struct {
	ID               string  `json:"id"`
	Severity         string  `json:"severity"`
	EPSS             float64 `json:"epss,omitempty"`
	SuggestedVersion string  `json:"suggested_version,omitempty"`
}

// handleGetLocalSCASummary groups SCA findings by component and computes upgrade recommendations
func handleGetLocalSCASummary(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	req, err := parseGetLocalSCASummaryRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	outputDir := veracodeWorkDir(req.ApplicationPath, "sca")
	resultsFile := filepath.Join(outputDir, "veracode.json")

	if _, statErr := os.Stat(resultsFile); os.IsNotExist(statErr) {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Local SCA Scan Results
===================

Application Path: %s
Results File: %s

❌ No results found

The results file does not exist. Run a local SCA scan using the local-sca-scan tool first.
`, req.ApplicationPath, resultsFile),
			}},
		}, nil
	}

	// #nosec G304 -- resultsFile is constructed from validated application path
	resultsData, err := os.ReadFile(resultsFile)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to read results file: %v", err),
		}, nil
	}

	var scaResults SCAFindings
	if err := json.Unmarshal(resultsData, &scaResults); err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to parse results file: %v", err),
		}, nil
	}

	return formatSCASummaryResponse(req.ApplicationPath, resultsFile, &scaResults, req), nil
}

// getOrCreateComponent returns the existing component summary for key, creating it if absent.
func getOrCreateComponent(summaryMap map[componentKey]*componentSummary, key componentKey, artifact SCAArtifact) *componentSummary {
	if comp, exists := summaryMap[key]; exists {
		return comp
	}
	comp := &componentSummary{
		Name:       artifact.Name,
		Version:    artifact.Version,
		Type:       artifact.Type,
		Language:   artifact.Language,
		PURL:       artifact.PURL,
		BySeverity: make(map[string]int),
	}
	summaryMap[key] = comp
	return comp
}

// matchSuggestedVersion returns the suggested fix version from match details, or empty string.
func matchSuggestedVersion(match SCAMatch) string {
	if len(match.MatchDetails) > 0 {
		return match.MatchDetails[0].Fix.SuggestedVersion
	}
	return ""
}

// collectMatchCVEEntries returns CVE entries for a match's primary vulnerability and related CVEs.
func collectMatchCVEEntries(match SCAMatch, suggestedVer string) []cveSummaryEntry {
	epssScore := 0.0
	if len(match.Vulnerability.EPSS) > 0 {
		epssScore = match.Vulnerability.EPSS[0].EPSS
	}
	entries := []cveSummaryEntry{{
		ID:               match.Vulnerability.ID,
		Severity:         strings.ToLower(match.Vulnerability.Severity),
		EPSS:             epssScore,
		SuggestedVersion: suggestedVer,
	}}
	for _, rv := range match.RelatedVulns {
		if !strings.HasPrefix(rv.ID, "CVE-") {
			continue
		}
		rvEPSS := 0.0
		if len(rv.EPSS) > 0 {
			rvEPSS = rv.EPSS[0].EPSS
		}
		entries = append(entries, cveSummaryEntry{
			ID:               rv.ID,
			Severity:         strings.ToLower(rv.Severity),
			EPSS:             rvEPSS,
			SuggestedVersion: suggestedVer, // related CVEs share the same fix
		})
	}
	return entries
}

// mergeComponentCVEs deduplicates and merges CVE entries into the component.
func mergeComponentCVEs(comp *componentSummary, entries []cveSummaryEntry) {
	existing := make(map[string]bool, len(comp.CVEs))
	for _, c := range comp.CVEs {
		existing[c.ID] = true
	}
	for _, c := range entries {
		if existing[c.ID] {
			continue
		}
		existing[c.ID] = true
		comp.CVEs = append(comp.CVEs, c)
		comp.BySeverity[c.Severity]++
		if c.EPSS > comp.MaxEPSS {
			comp.MaxEPSS = c.EPSS
		}
	}
}

// sortComponentSummaries sorts the component slice and the CVEs within each component.
func sortComponentSummaries(result []componentSummary) {
	sort.Slice(result, func(i, j int) bool {
		maxSevI := maxSeverityLevel(result[i].BySeverity)
		maxSevJ := maxSeverityLevel(result[j].BySeverity)
		if maxSevI != maxSevJ {
			return maxSevI > maxSevJ
		}
		if len(result[i].CVEs) != len(result[j].CVEs) {
			return len(result[i].CVEs) > len(result[j].CVEs)
		}
		return strings.ToLower(result[i].Name) < strings.ToLower(result[j].Name)
	})
	for i := range result {
		sort.Slice(result[i].CVEs, func(a, b int) bool {
			sevA := severityToInt(result[i].CVEs[a].Severity)
			sevB := severityToInt(result[i].CVEs[b].Severity)
			if sevA != sevB {
				return sevA > sevB
			}
			return result[i].CVEs[a].ID < result[i].CVEs[b].ID
		})
	}
}

// buildComponentSummaries groups SCA matches by component and computes per-component upgrade recommendations
func buildComponentSummaries(results *SCAFindings, req *GetLocalSCASummaryRequest) []componentSummary {
	summaryMap := make(map[componentKey]*componentSummary)
	upgradeCandidates := make(map[componentKey][]string)

	for _, match := range results.Vulnerabilities.Matches {
		if req.SeverityGTE != nil && severityToInt(match.Vulnerability.Severity) < *req.SeverityGTE {
			continue
		}

		key := componentKey{Name: match.Artifact.Name, Version: match.Artifact.Version}
		comp := getOrCreateComponent(summaryMap, key, match.Artifact)

		suggestedVer := matchSuggestedVersion(match)
		if suggestedVer != "" {
			upgradeCandidates[key] = append(upgradeCandidates[key], suggestedVer)
		}

		entries := collectMatchCVEEntries(match, suggestedVer)
		mergeComponentCVEs(comp, entries)
	}

	// Compute recommended upgrade: highest suggested version across all CVEs fixes all of them.
	for key, comp := range summaryMap {
		comp.RecommendedUpgrade = pickHighestVersion(upgradeCandidates[key])
	}

	result := make([]componentSummary, 0, len(summaryMap))
	for _, comp := range summaryMap {
		result = append(result, *comp)
	}
	sortComponentSummaries(result)
	return result
}

// maxSeverityLevel returns the highest numeric severity in the map
func maxSeverityLevel(bySeverity map[string]int) int {
	max := 0
	for sev, count := range bySeverity {
		if count > 0 {
			if v := severityToInt(sev); v > max {
				max = v
			}
		}
	}
	return max
}

// pickHighestVersion returns the highest version string from candidates.
// It attempts semver-style comparison (splitting on ".") and falls back to lexicographic.
func pickHighestVersion(versions []string) string {
	if len(versions) == 0 {
		return ""
	}
	best := versions[0]
	for _, v := range versions[1:] {
		if compareVersionStrings(v, best) > 0 {
			best = v
		}
	}
	return best
}

// compareVersionStrings compares two version strings.
// Returns positive if a > b, negative if a < b, 0 if equal.
// Strips leading "v" or "=" prefixes before comparing.
func compareVersionStrings(a, b string) int {
	a = strings.TrimLeft(a, "v=")
	b = strings.TrimLeft(b, "v=")

	partsA := strings.Split(a, ".")
	partsB := strings.Split(b, ".")

	maxLen := len(partsA)
	if len(partsB) > maxLen {
		maxLen = len(partsB)
	}

	for i := 0; i < maxLen; i++ {
		var pA, pB int
		if i < len(partsA) {
			pA, _ = strconv.Atoi(partsA[i])
		}
		if i < len(partsB) {
			pB, _ = strconv.Atoi(partsB[i])
		}
		if pA != pB {
			return pA - pB
		}
	}
	return 0
}

// formatSCASummaryResponse formats the component-grouped summary into an MCP response
func formatSCASummaryResponse(appPath, resultsFile string, results *SCAFindings, req *GetLocalSCASummaryRequest) map[string]interface{} {
	components := buildComponentSummaries(results, req)

	// Build overall summary counts
	totalCVEs := 0
	bySeverity := map[string]int{"critical": 0, "high": 0, "medium": 0, "low": 0}
	componentsWithFix := 0
	componentsNoFix := 0

	for _, comp := range components {
		totalCVEs += len(comp.CVEs)
		for sev, count := range comp.BySeverity {
			bySeverity[sev] += count
		}
		if comp.RecommendedUpgrade != "" {
			componentsWithFix++
		} else {
			componentsNoFix++
		}
	}

	// Apply pagination to the component list
	totalComponents := len(components)
	if req.Size <= 0 {
		req.Size = 10
	}
	startIdx := req.Page * req.Size
	endIdx := startIdx + req.Size
	if startIdx > totalComponents {
		startIdx = totalComponents
	}
	if endIdx > totalComponents {
		endIdx = totalComponents
	}
	pagedComponents := components[startIdx:endIdx]

	totalPages := 0
	if req.Size > 0 {
		totalPages = (totalComponents + req.Size - 1) / req.Size
	}
	pagination := map[string]interface{}{
		"current_page":   req.Page,
		"page_size":      req.Size,
		"total_elements": totalComponents,
		"total_pages":    totalPages,
		"has_next":       endIdx < totalComponents,
		"has_previous":   req.Page > 0,
	}

	// Build the component list for the response
	componentList := make([]map[string]interface{}, 0, len(pagedComponents))
	for _, comp := range pagedComponents {
		entry := map[string]interface{}{
			"name":            comp.Name,
			"current_version": comp.Version,
			"type":            comp.Type,
			"language":        comp.Language,
			"cve_count":       len(comp.CVEs),
			"by_severity":     comp.BySeverity,
			"max_epss":        comp.MaxEPSS,
			"cves":            comp.CVEs,
		}
		if comp.PURL != "" {
			entry["purl"] = comp.PURL
		}
		if comp.RecommendedUpgrade != "" {
			entry["recommended_upgrade"] = comp.RecommendedUpgrade
			entry["upgrade_fixes_all_cves"] = true
		} else {
			entry["recommended_upgrade"] = nil
			entry["upgrade_fixes_all_cves"] = false
		}
		componentList = append(componentList, entry)
	}

	filters := map[string]interface{}{}
	if req.SeverityGTE != nil {
		severityNames := map[int]string{2: "low", 3: "medium", 4: "high", 5: "critical"}
		if name, ok := severityNames[*req.SeverityGTE]; ok {
			filters["severity_gte"] = name
		}
	}

	responseData := map[string]interface{}{
		"application": map[string]string{
			"name": filepath.Base(appPath),
			"path": appPath,
		},
		"summary": map[string]interface{}{
			"vulnerable_components":  totalComponents,
			"components_with_fix":    componentsWithFix,
			"components_without_fix": componentsNoFix,
			"total_cves":             totalCVEs,
			"by_severity":            bySeverity,
		},
		"pagination": pagination,
		"components": componentList,
	}

	if len(filters) > 0 {
		responseData["filters"] = filters
	}

	responseJSON, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf("Error formatting response: %v", err),
			}},
		}
	}

	result := map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(responseJSON),
			},
		},
		"structuredContent": responseData,
	}

	return result
}
