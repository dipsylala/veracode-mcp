package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/dipsylala/veracodemcp-go/api"
	"github.com/dipsylala/veracodemcp-go/api/generated/applications"
	"github.com/dipsylala/veracodemcp-go/workspace"
)

const StaticFindingsToolName = "static-findings"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(StaticFindingsToolName, handleGetStaticFindings)
}

// StaticFindingsRequest represents the parsed parameters for get-static-findings
type StaticFindingsRequest struct {
	ApplicationPath    string   `json:"application_path"`
	AppProfile         string   `json:"app_profile,omitempty"`
	Sandbox            string   `json:"sandbox,omitempty"`
	Size               int      `json:"size,omitempty"`
	Page               int      `json:"page,omitempty"`
	Severity           *int32   `json:"severity,omitempty"`
	SeverityGte        *int32   `json:"severity_gte,omitempty"`
	CWEIDs             []string `json:"cwe_ids,omitempty"`
	ViolatesPolicy     *bool    `json:"violates_policy,omitempty"`
	IncludeMitigations bool     `json:"include_mitigations,omitempty"`
}

// convertCWEIDsToStrings converts CWE IDs from various types to string slice
func convertCWEIDsToStrings(cweIdsRaw interface{}) []string {
	cweArray, ok := cweIdsRaw.([]interface{})
	if !ok {
		return nil
	}

	cweIDs := make([]string, len(cweArray))
	for i, cwe := range cweArray {
		switch v := cwe.(type) {
		case float64:
			cweIDs[i] = fmt.Sprintf("%.0f", v)
		case int:
			cweIDs[i] = fmt.Sprintf("%d", v)
		case string:
			cweIDs[i] = v
		default:
			cweIDs[i] = fmt.Sprintf("%v", v)
		}
	}
	return cweIDs
}

// parseStaticFindingsRequest extracts and validates parameters from the raw args map
func parseStaticFindingsRequest(args map[string]interface{}) (*StaticFindingsRequest, error) {
	req := &StaticFindingsRequest{
		Size: 10,
		Page: 0,
	}

	// Extract fields directly
	if appPath, ok := args["application_path"].(string); ok {
		req.ApplicationPath = appPath
	}
	if appProfile, ok := args["app_profile"].(string); ok {
		req.AppProfile = appProfile
	}
	if sandbox, ok := args["sandbox"].(string); ok {
		req.Sandbox = sandbox
	}
	if size, ok := args["size"].(float64); ok {
		req.Size = int(size)
	}
	if page, ok := args["page"].(float64); ok {
		req.Page = int(page)
	}
	if severity, ok := args["severity"].(float64); ok {
		sev := int32(severity)
		req.Severity = &sev
	}
	if severityGte, ok := args["severity_gte"].(float64); ok {
		sev := int32(severityGte)
		req.SeverityGte = &sev
	}
	if violatesPolicy, ok := args["violates_policy"].(bool); ok {
		req.ViolatesPolicy = &violatesPolicy
	}
	if includeMitigations, ok := args["include_mitigations"].(bool); ok {
		req.IncludeMitigations = includeMitigations
	}
	if cweIdsRaw, ok := args["cwe_ids"]; ok {
		req.CWEIDs = convertCWEIDsToStrings(cweIdsRaw)
	}

	if req.ApplicationPath == "" {
		return nil, fmt.Errorf("application_path is required and must be an absolute path")
	}

	return req, nil
}

func handleGetStaticFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseStaticFindingsRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Step 1: Retrieve application profile name
	appProfile := req.AppProfile
	hasAppProfile := appProfile != ""
	if !hasAppProfile {
		// Fall back to workspace config if app_profile not provided
		appProfile, err = workspace.FindWorkspaceConfig(req.ApplicationPath)
		if err != nil {
			return map[string]interface{}{
				"error": fmt.Sprintf("Failed to find workspace configuration: %v", err),
			}, nil
		}
	}

	// Step 2: Create API client
	client, err := api.NewVeracodeClient()
	if err != nil {
		responseText := fmt.Sprintf(`Static Findings Analysis - Error
========================

Application Path: %s
App Profile: %s
Error: Failed to create API client

%v

Please ensure VERACODE_API_ID and VERACODE_API_KEY environment variables are set with valid credentials.`,
			req.ApplicationPath,
			appProfile,
			err,
		)

		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": responseText,
			}},
		}, nil
	}

	// Step 3: Get application GUID using the profile name (or use directly if it's a GUID)
	var applicationGUID string

	// Check if appProfile is already a GUID (format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
	if len(appProfile) == 36 && appProfile[8] == '-' && appProfile[13] == '-' && appProfile[18] == '-' && appProfile[23] == '-' {
		// It's a GUID, use it directly
		applicationGUID = appProfile
	} else {
		// It's an app name, look it up
		var application *applications.Application
		application, err = client.GetApplicationByName(ctx, appProfile)
		if err != nil {
			appProfileSource := "from workspace config"
			if hasAppProfile {
				appProfileSource = "from parameter (overriding workspace)"
			}

			responseText := fmt.Sprintf(`Static Findings Analysis - Error
========================

Application Path: %s
App Profile: %s (%s)
Error: Failed to lookup application

%v

Please verify:
- The application name exists in Veracode platform
- API credentials have sufficient permissions
- The application name matches exactly (case-insensitive)`,
				req.ApplicationPath,
				appProfile,
				appProfileSource,
				err,
			)

			return map[string]interface{}{
				"content": []map[string]string{{
					"type": "text",
					"text": responseText,
				}},
			}, nil
		}

		// Extract GUID from application
		if application.Guid != nil {
			applicationGUID = *application.Guid
		} else {
			applicationGUID = "unknown"
		}
	}

	// Step 4: Build the findings request
	findingsReq := api.FindingsRequest{
		AppProfile:         applicationGUID,
		Sandbox:            req.Sandbox,
		Size:               req.Size,
		Page:               req.Page,
		Severity:           req.Severity,
		SeverityGte:        req.SeverityGte,
		CWEIDs:             req.CWEIDs,
		ViolatesPolicy:     req.ViolatesPolicy,
		IncludeMitigations: req.IncludeMitigations,
	}

	// Step 5: Call the API to get static findings
	findingsResp, err := client.GetStaticFindings(ctx, findingsReq)
	if err != nil {
		responseText := fmt.Sprintf(`Static Findings Analysis - Error
========================

Application Path: %s
App Profile: %s
Application GUID: %s
Error: Failed to retrieve static findings

%v

Please verify:
- The application has been scanned
- API credentials have access to view findings
- The sandbox name is correct (if specified)`,
			req.ApplicationPath,
			appProfile,
			applicationGUID,
			err,
		)

		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": responseText,
			}},
		}, nil
	}

	// Step 6: Format and return the response
	return formatStaticFindingsResponse(ctx, appProfile, applicationGUID, req.Sandbox, findingsResp), nil
}

// formatStaticFindingsResponse formats the findings API response into an MCP tool response
func formatStaticFindingsResponse(ctx context.Context, appProfile, applicationGUID, sandbox string, findings *api.FindingsResponse) map[string]interface{} {
	// Build MCP response structure
	response := MCPFindingsResponse{
		Application: MCPApplication{
			Name: appProfile,
			ID:   applicationGUID,
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
				"open":   0,
				"closed": 0,
			},
			ByMitigation: map[string]int{
				"none":     0,
				"proposed": 0,
				"approved": 0,
				"rejected": 0,
			},
		},
		Findings: []MCPFinding{},
	}

	// Add sandbox if specified
	if sandbox != "" {
		response.Sandbox = &MCPSandbox{
			Name: sandbox,
			ID:   sandbox,
		}
	}

	// Add pagination info
	if findings != nil {
		response.Pagination = &MCPPagination{
			CurrentPage:   findings.Page,
			PageSize:      findings.Size,
			TotalElements: findings.TotalCount,
			TotalPages:    (findings.TotalCount + findings.Size - 1) / findings.Size,
			HasNext:       (findings.Page+1)*findings.Size < findings.TotalCount,
			HasPrevious:   findings.Page > 0,
		}
		// Set total_findings to the total across all pages
		response.Summary.TotalFindings = findings.TotalCount
	}

	// Process each finding
	if findings != nil && len(findings.Findings) > 0 {
		for _, finding := range findings.Findings {
			mcpFinding := processStaticFinding(finding)
			response.Findings = append(response.Findings, mcpFinding)
			updateStaticSummaryCounters(&response.Summary, mcpFinding)
		}
	}

	// Sort findings by severity (Very High first)
	sort.Slice(response.Findings, func(i, j int) bool {
		return response.Findings[i].SeverityScore > response.Findings[j].SeverityScore
	})

	// Marshal response to JSON for non-UI clients
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("Warning: Failed to marshal response to JSON: %v", err)
		responseJSON, _ = json.Marshal(response) // Fall back to compact JSON
	}

	// Build pagination summary for LLM
	var paginationSummary string
	if response.Pagination != nil {
		paginationSummary = fmt.Sprintf("Showing %d findings on page %d of %d (Total: %d findings across all pages)\n\n",
			len(response.Findings),
			response.Pagination.CurrentPage+1, // Display as 1-based
			response.Pagination.TotalPages,
			response.Pagination.TotalElements,
		)
	} else {
		paginationSummary = fmt.Sprintf("Showing %d findings\n\n", len(response.Findings))
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
	if ClientSupportsUIFromContext(ctx) {
		log.Printf("Static findings: Returning %d findings (content: JSON, structuredContent: full data for UI)", len(response.Findings))
		result["structuredContent"] = response
	} else {
		log.Printf("Static findings: Returning %d findings (content: JSON only, no structuredContent - client doesn't support UI)", len(response.Findings))
	}

	return result
}

// processStaticFinding transforms a single API finding into an MCP finding
func processStaticFinding(finding api.Finding) MCPFinding {
	// Use severity score from API extraction
	severityNum := finding.SeverityScore
	transformedSeverity := TransformSeverity(&severityNum)

	// Clean and extract references from description
	cleanedDesc, references := TransformDescription(finding.Description, "STATIC")

	// Transform mitigation status
	mitigationStatus := "NONE"
	if finding.ResolutionStatus != "" {
		mitigationStatus = finding.ResolutionStatus
	}

	// Extract numeric CWE ID
	var cweID int32
	if finding.CWE != "" {
		_, _ = fmt.Sscanf(finding.CWE, "CWE-%d", &cweID)
	}

	return MCPFinding{
		FlawID:           finding.ID,
		ScanType:         "STATIC",
		Status:           finding.Status,
		MitigationStatus: mitigationStatus,
		ViolatesPolicy:   finding.ViolatesPolicy,
		Severity:         string(transformedSeverity),
		SeverityScore:    severityNum,
		CweId:            cweID,
		Description:      cleanedDesc,
		References:       references,
		FilePath:         finding.FilePath,
		LineNumber:       finding.LineNumber,
		Module:           finding.Module,
		Procedure:        finding.Procedure,
		AttackVector:     finding.AttackVector,
		Mitigations:      convertMitigations(finding.Mitigations),
	}
}

// convertMitigations converts API mitigations to MCP mitigations
func convertMitigations(apiMitigations []api.Mitigation) []MCPMitigation {
	if len(apiMitigations) == 0 {
		return nil
	}

	mitigations := make([]MCPMitigation, len(apiMitigations))
	for i, m := range apiMitigations {
		mitigations[i] = MCPMitigation{
			Action:    m.Action,
			Comment:   m.Comment,
			Submitter: m.Submitter,
			Date:      m.Date,
		}
	}
	return mitigations
}

// updateStaticSummaryCounters updates the response summary based on a static finding
func updateStaticSummaryCounters(summary *MCPFindingsSummary, finding MCPFinding) {
	// Note: TotalFindings is set from pagination.total_elements, not incremented per finding

	if finding.Status == "OPEN" || finding.Status == "NEW" {
		summary.OpenFindings++
	}

	if finding.ViolatesPolicy {
		summary.PolicyViolations++
	}

	// Count by severity
	severityKey := strings.ToLower(finding.Severity)
	summary.BySeverity[severityKey]++

	// Count by status
	statusKey := strings.ToLower(finding.Status)
	if statusKey == "new" {
		statusKey = "open"
	}
	summary.ByStatus[statusKey]++
}
