package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dipsylala/veracodemcp-go/api"
	"github.com/dipsylala/veracodemcp-go/api/generated/applications"
	"github.com/dipsylala/veracodemcp-go/workspace"
)

const ScaFindingsToolName = "get-sca-findings"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(ScaFindingsToolName, handleGetScaFindings)
}

// ScaFindingsRequest represents the parsed parameters for get-sca-findings
type ScaFindingsRequest struct {
	ApplicationPath string `json:"application_path"`
	AppProfile      string `json:"app_profile,omitempty"`
	Sandbox         string `json:"sandbox,omitempty"`
	Size            int    `json:"size,omitempty"`
	Page            int    `json:"page,omitempty"`
	Severity        *int32 `json:"severity,omitempty"`
	SeverityGte     *int32 `json:"severity_gte,omitempty"`
}

// parseScaFindingsRequest extracts and validates parameters from the raw args map
func parseScaFindingsRequest(args map[string]interface{}) (*ScaFindingsRequest, error) {
	req := &ScaFindingsRequest{
		Size: 200,
		Page: 0,
	}

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

	if req.ApplicationPath == "" {
		return nil, fmt.Errorf("application_path is required and must be an absolute path")
	}

	return req, nil
}

func handleGetScaFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseScaFindingsRequest(args)
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
		responseText := fmt.Sprintf(`SCA Findings Analysis - Error
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

			responseText := fmt.Sprintf(`SCA Findings Analysis - Error
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
		AppProfile:  applicationGUID,
		Sandbox:     req.Sandbox,
		Size:        req.Size,
		Page:        req.Page,
		Severity:    req.Severity,
		SeverityGte: req.SeverityGte,
	}

	// Step 5: Call the API to get SCA findings
	findingsResp, err := client.GetScaFindings(ctx, findingsReq)
	if err != nil {
		responseText := fmt.Sprintf(`SCA Findings Analysis - Error
========================

Application Path: %s
App Profile: %s
Application GUID: %s
Error: Failed to retrieve SCA findings

%v

Please verify:
- The application has been scanned with SCA
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
	return formatScaFindingsResponse(appProfile, applicationGUID, req.Sandbox, findingsResp), nil
}

// formatScaFindingsResponse formats the findings API response into an MCP tool response
func formatScaFindingsResponse(appProfile, applicationGUID, sandbox string, findings *api.FindingsResponse) map[string]interface{} {
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
			mcpFinding := processScaFinding(finding)
			response.Findings = append(response.Findings, mcpFinding)
			updateSummaryCounters(&response.Summary, mcpFinding)
		}
	}

	// Marshal response to JSON for non-UI clients
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		// Fall back to compact JSON
		responseJSON, _ = json.Marshal(response)
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

	// Return as MCP tool response with text content (consistent with static/dynamic findings)
	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": paginationSummary + string(responseJSON),
			},
		},
	}
}

// processScaFinding transforms a single API finding into an MCP finding
func processScaFinding(finding api.Finding) MCPFinding {
	// Use severity score from API extraction
	severityNum := finding.SeverityScore
	transformedSeverity := TransformSeverity(&severityNum)

	// Transform status (OPEN/CLOSED/UNKNOWN)
	status := finding.Status
	if status == "" {
		status = "OPEN" // Default
	}

	// Normalize status
	var transformedStatus FindingStatus
	switch strings.ToUpper(status) {
	case "OPEN", "NEW":
		transformedStatus = StatusOpen
	case "CLOSED":
		transformedStatus = StatusClosed
	default:
		transformedStatus = StatusUnknown
	}

	// Map ResolutionStatus from FindingStatus API directly to mitigation status
	mitigationStatus := finding.ResolutionStatus
	if mitigationStatus == "" {
		mitigationStatus = "NONE"
	}

	// Determine policy violation using business rule
	violatesPolicyValue := finding.ViolatesPolicy
	violatesPolicy := DeterminesPolicyViolation(
		transformedStatus,
		mitigationStatus,
		&violatesPolicyValue,
	)

	// Clean and extract references from description
	cleanedDesc, references := TransformDescription(finding.Description, "SCA")

	// Extract numeric CWE ID
	var cweID int32
	if finding.CWE != "" {
		_, _ = fmt.Sscanf(finding.CWE, "CWE-%d", &cweID)
	}

	mcpFinding := MCPFinding{
		FlawID:           finding.ID,
		ScanType:         "SCA",
		Status:           string(transformedStatus),
		MitigationStatus: mitigationStatus,
		ViolatesPolicy:   violatesPolicy,
		Severity:         string(transformedSeverity),
		SeverityScore:    severityNum,
		CweId:            cweID,
		Description:      cleanedDesc,
		References:       references,
	}

	// Add SCA-specific component information if available
	if finding.ComponentFilename != "" || finding.ComponentVersion != "" {
		component := &MCPComponent{
			Name:    finding.ComponentFilename,
			Version: finding.ComponentVersion,
			Library: finding.ComponentFilename,
		}

		// Add license information if available
		if len(finding.Licenses) > 0 {
			component.Licenses = make([]MCPLicense, 0, len(finding.Licenses))
			for _, lic := range finding.Licenses {
				component.Licenses = append(component.Licenses, MCPLicense{
					LicenseID:  lic.LicenseID,
					RiskRating: lic.RiskRating,
				})
			}
		}

		mcpFinding.Component = component
	}

	// Add vulnerability information if available
	if finding.CVE != "" {
		mcpFinding.Vulnerability = &MCPVulnerability{
			CVEID: finding.CVE,
		}
	}

	return mcpFinding
}
