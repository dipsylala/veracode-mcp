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

const GetLocalIACFindingsToolName = "local-iac-findings"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(GetLocalIACFindingsToolName, handleGetLocalIACFindings)
}

// GetLocalIACFindingsRequest represents the parsed parameters for local-iac-findings
type GetLocalIACFindingsRequest struct {
	ApplicationPath string
	Target          string `json:"target,omitempty"`
	CheckID         string `json:"check_id,omitempty"`
	SeverityGTE     *int   `json:"severity_gte,omitempty"`
	Status          string `json:"status,omitempty"` // "fail", "pass", "all" — default "fail"
	Size            int    `json:"size,omitempty"`
	Page            int    `json:"page,omitempty"`
}

// parseGetLocalIACFindingsRequest extracts and validates parameters from the raw args map
func parseGetLocalIACFindingsRequest(args map[string]interface{}) (*GetLocalIACFindingsRequest, error) {
	req := &GetLocalIACFindingsRequest{}

	var err error
	req.ApplicationPath, err = extractRequiredString(args, "application_path")
	if err != nil {
		return nil, err
	}

	if target, ok := args["target"].(string); ok && target != "" {
		req.Target = target
	}
	if checkID, ok := args["check_id"].(string); ok && checkID != "" {
		req.CheckID = strings.ToUpper(checkID)
	}
	if sevGTE, ok := args["severity_gte"].(float64); ok {
		sevInt := int(sevGTE)
		req.SeverityGTE = &sevInt
	}

	// Status filter: default to "fail"
	req.Status = "fail"
	if status, ok := args["status"].(string); ok && status != "" {
		lower := strings.ToLower(status)
		switch lower {
		case "fail", "pass", "all":
			req.Status = lower
		}
	}

	req.Size = extractInt(args, "page_size", 10)
	req.Page = extractInt(args, "page", 0)

	if err := validatePaginationParams(req.Size, req.Page); err != nil {
		return nil, err
	}

	return req, nil
}

// IACFindings represents the top-level JSON structure for IaC results
type IACFindings struct {
	Configs []IACConfig `json:"configs"`
}

// IACConfig represents a single IaC misconfiguration check result
type IACConfig struct {
	ID             string            `json:"ID"`
	Title          string            `json:"Title"`
	Description    string            `json:"Description"`
	Message        string            `json:"Message"`
	Severity       string            `json:"Severity"`
	Status         string            `json:"Status"`
	Target         string            `json:"Target"`
	Type           string            `json:"Type"`
	Resolution     string            `json:"Resolution"`
	PrimaryURL     string            `json:"PrimaryURL"`
	References     []string          `json:"References"`
	Namespace      string            `json:"Namespace"`
	CauseMetadata  IACCauseMetadata  `json:"CauseMetadata"`
	CustomerPolicy IACCustomerPolicy `json:"customerPolicyResult"`
}

// IACCauseMetadata contains metadata about what caused the finding
type IACCauseMetadata struct {
	Provider string `json:"Provider"`
	Service  string `json:"Service"`
}

// IACCustomerPolicy holds the customer policy evaluation result
type IACCustomerPolicy struct {
	Status string `json:"Status"`
}

// iacSeverityToInt converts IaC severity string to numeric value for filtering
func iacSeverityToInt(severity string) int {
	switch strings.ToUpper(strings.TrimSpace(severity)) {
	case "CRITICAL":
		return 5
	case "HIGH":
		return 4
	case "MEDIUM":
		return 3
	case "LOW":
		return 2
	default:
		return 0
	}
}

// iacSummary holds summary statistics for IaC findings
type iacSummary struct {
	Total         int
	Fail          int
	Pass          int
	Critical      int
	High          int
	Medium        int
	Low           int
	UniqueChecks  int
	UniqueTargets int
}

// buildIACSummary builds summary statistics from IaC findings
func buildIACSummary(configs []IACConfig) iacSummary {
	summary := iacSummary{Total: len(configs)}
	checkMap := make(map[string]bool)
	targetMap := make(map[string]bool)

	for _, c := range configs {
		checkMap[c.ID] = true
		targetMap[c.Target] = true

		switch strings.ToUpper(c.Status) {
		case "FAIL":
			summary.Fail++
		case "PASS":
			summary.Pass++
		}

		switch strings.ToUpper(strings.TrimSpace(c.Severity)) {
		case "CRITICAL":
			summary.Critical++
		case "HIGH":
			summary.High++
		case "MEDIUM":
			summary.Medium++
		case "LOW":
			summary.Low++
		}
	}
	summary.UniqueChecks = len(checkMap)
	summary.UniqueTargets = len(targetMap)
	return summary
}

// iacMatchesFilters checks if an IaC config passes all active filters
func iacMatchesFilters(config IACConfig, req *GetLocalIACFindingsRequest) bool {
	// Status filter
	if req.Status != "all" {
		if !strings.EqualFold(config.Status, req.Status) {
			return false
		}
	}

	// Target filter (partial match, case-insensitive)
	if req.Target != "" {
		if !strings.Contains(strings.ToLower(config.Target), strings.ToLower(req.Target)) {
			return false
		}
	}

	// Check ID filter (exact, case-insensitive)
	if req.CheckID != "" {
		if !strings.EqualFold(config.ID, req.CheckID) {
			return false
		}
	}

	// Minimum severity filter
	if req.SeverityGTE != nil {
		if iacSeverityToInt(config.Severity) < *req.SeverityGTE {
			return false
		}
	}

	return true
}

// convertConfigToFinding converts an IACConfig to LLM-friendly format
func convertConfigToFinding(config IACConfig) map[string]interface{} {
	finding := map[string]interface{}{
		"check_id":    config.ID,
		"title":       config.Title,
		"description": config.Description,
		"message":     config.Message,
		"severity":    strings.ToLower(config.Severity),
		"status":      strings.ToLower(config.Status),
		"target":      config.Target,
		"type":        config.Type,
		"resolution":  config.Resolution,
		"primary_url": config.PrimaryURL,
		"references":  config.References,
		"provider":    config.CauseMetadata.Provider,
	}
	if config.CustomerPolicy.Status != "" {
		finding["policy_status"] = config.CustomerPolicy.Status
	}
	return finding
}

// handleGetLocalIACFindings retrieves and formats local IaC scan findings
func handleGetLocalIACFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	req, err := parseGetLocalIACFindingsRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// IaC results are stored alongside SCA results in the same veracode.json
	outputDir := veracodeWorkDir(req.ApplicationPath, "sca")
	resultsFile := filepath.Join(outputDir, "veracode.json")

	if _, statErr := os.Stat(resultsFile); os.IsNotExist(statErr) {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Local IaC Scan Results
======================

Application Path: %s
Results File: %s

❌ No results found

The results file does not exist. Run a local SCA scan using the run-sca-scan tool first.
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

	var fullResults IACFindings
	if err := json.Unmarshal(resultsData, &fullResults); err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to parse results file: %v", err),
		}, nil
	}

	return formatIACFindingsResponse(req.ApplicationPath, &fullResults, req), nil
}

// sortIACFindings sorts findings: FAIL before PASS, then severity descending, then target ascending
func sortIACFindings(findings []map[string]interface{}) {
	sort.Slice(findings, func(i, j int) bool {
		statusI := findings[i]["status"].(string)
		statusJ := findings[j]["status"].(string)
		if statusI != statusJ {
			return statusI == "fail"
		}
		sevI := iacSeverityToInt(findings[i]["severity"].(string))
		sevJ := iacSeverityToInt(findings[j]["severity"].(string))
		if sevI != sevJ {
			return sevI > sevJ
		}
		targetI := findings[i]["target"].(string)
		targetJ := findings[j]["target"].(string)
		return strings.ToLower(targetI) < strings.ToLower(targetJ)
	})
}

// paginateIACFindings applies pagination and returns the page slice plus pagination metadata
func paginateIACFindings(findings []map[string]interface{}, req *GetLocalIACFindingsRequest) ([]map[string]interface{}, map[string]interface{}) {
	if req.Size <= 0 {
		req.Size = 10
	}
	total := len(findings)
	startIdx := req.Page * req.Size
	endIdx := startIdx + req.Size
	if startIdx > total {
		startIdx = total
	}
	if endIdx > total {
		endIdx = total
	}
	totalPages := 0
	if req.Size > 0 {
		totalPages = (total + req.Size - 1) / req.Size
	}
	pagination := map[string]interface{}{
		"current_page":   req.Page,
		"page_size":      req.Size,
		"total_elements": total,
		"total_pages":    totalPages,
		"has_next":       endIdx < total,
		"has_previous":   req.Page > 0,
	}
	return findings[startIdx:endIdx], pagination
}

// buildIACFilters constructs the active-filters map from the request
func buildIACFilters(req *GetLocalIACFindingsRequest) map[string]interface{} {
	filters := map[string]interface{}{}
	if req.Target != "" {
		filters["target"] = req.Target
	}
	if req.CheckID != "" {
		filters["check_id"] = req.CheckID
	}
	if req.SeverityGTE != nil {
		severityNames := map[int]string{2: "low", 3: "medium", 4: "high", 5: "critical"}
		if name, ok := severityNames[*req.SeverityGTE]; ok {
			filters["severity_gte"] = name
		}
	}
	if req.Status != "fail" {
		filters["status"] = req.Status
	}
	return filters
}

// formatIACFindingsResponse formats the IaC findings into an MCP response
func formatIACFindingsResponse(appPath string, results *IACFindings, req *GetLocalIACFindingsRequest) map[string]interface{} {
	// Apply filters and convert
	allFindings := make([]map[string]interface{}, 0, len(results.Configs))
	filteredConfigs := make([]IACConfig, 0, len(results.Configs))
	for _, config := range results.Configs {
		if iacMatchesFilters(config, req) {
			allFindings = append(allFindings, convertConfigToFinding(config))
			filteredConfigs = append(filteredConfigs, config)
		}
	}

	summary := buildIACSummary(filteredConfigs)
	sortIACFindings(allFindings)
	page, pagination := paginateIACFindings(allFindings, req)
	filters := buildIACFilters(req)

	responseData := map[string]interface{}{
		"application": map[string]string{
			"name": filepath.Base(appPath),
			"path": appPath,
		},
		"summary": map[string]interface{}{
			"total_checks":   summary.Total,
			"fail":           summary.Fail,
			"pass":           summary.Pass,
			"unique_checks":  summary.UniqueChecks,
			"unique_targets": summary.UniqueTargets,
			"by_severity": map[string]int{
				"critical": summary.Critical,
				"high":     summary.High,
				"medium":   summary.Medium,
				"low":      summary.Low,
			},
		},
		"pagination": pagination,
		"findings":   page,
	}

	if len(filters) > 0 {
		responseData["filters"] = filters
	}

	responseJSON, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{{
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
