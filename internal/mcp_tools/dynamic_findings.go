package mcp_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/dipsylala/veracode-mcp/api"
	"github.com/dipsylala/veracode-mcp/api/rest/generated/applications"
	"github.com/dipsylala/veracode-mcp/workspace"
)

const DynamicFindingsToolName = "dynamic-findings"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(DynamicFindingsToolName, handleGetDynamicFindings)
}

// DynamicFindingsRequest represents the parsed parameters for dynamic-findings
type DynamicFindingsRequest struct {
	ApplicationPath string   `json:"application_path"`
	AppProfile      string   `json:"app_profile,omitempty"`
	Sandbox         string   `json:"sandbox,omitempty"`
	Size            int      `json:"size,omitempty"`
	Page            int      `json:"page,omitempty"`
	Severity        *int32   `json:"severity,omitempty"`
	SeverityGte     *int32   `json:"severity_gte,omitempty"`
	CWEIDs          []string `json:"cwe_ids,omitempty"`
	ViolatesPolicy  *bool    `json:"violates_policy,omitempty"`
}

// parseDynamicFindingsRequest extracts and validates parameters from the raw args map
func parseDynamicFindingsRequest(args map[string]interface{}) (*DynamicFindingsRequest, error) {
	req := &DynamicFindingsRequest{}

	// Extract required fields
	var err error
	req.ApplicationPath, err = extractRequiredString(args, "application_path")
	if err != nil {
		return nil, err
	}

	// Extract optional fields
	req.AppProfile, _ = extractOptionalString(args, "app_profile")
	req.Sandbox, _ = extractOptionalString(args, "sandbox")
	req.Size = extractInt(args, "page_size", 10)
	req.Page = extractInt(args, "page", 0)

	// Extract optional int32 pointers with validation
	req.Severity, _, err = extractOptionalInt32Ptr(args, "severity")
	if err != nil {
		return nil, err
	}
	req.SeverityGte, _, err = extractOptionalInt32Ptr(args, "severity_gte")
	if err != nil {
		return nil, err
	}

	// Extract optional booleans
	// Extract optional booleans — default violates_policy to true if not provided
	violatesPolicy, provided := extractOptionalBool(args, "violates_policy")
	if !provided {
		defaultTrue := true
		violatesPolicy = &defaultTrue
	}
	req.ViolatesPolicy = violatesPolicy

	// Extract CWE IDs
	req.CWEIDs = extractCWEIDs(args)

	// Validate pagination bounds
	if err := validatePaginationParams(req.Size, req.Page); err != nil {
		return nil, err
	}

	// Validate severity bounds
	if err := validateSeverity(req.Severity, "severity"); err != nil {
		return nil, err
	}
	if err := validateSeverity(req.SeverityGte, "severity_gte"); err != nil {
		return nil, err
	}

	return req, nil
}

func handleGetDynamicFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseDynamicFindingsRequest(args)
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
	client, err := api.NewClient()
	if err != nil {
		responseText := fmt.Sprintf(`Dynamic Findings Analysis - Error
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

			responseText := fmt.Sprintf(`Dynamic Findings Analysis - Error
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
		AppProfile:     applicationGUID,
		Sandbox:        req.Sandbox,
		Size:           req.Size,
		Page:           req.Page,
		Severity:       req.Severity,
		SeverityGte:    req.SeverityGte,
		CWEIDs:         req.CWEIDs,
		ViolatesPolicy: req.ViolatesPolicy,
	}

	// Step 5: Call the API to get dynamic findings
	findingsResp, err := client.GetDynamicFindings(ctx, findingsReq)
	if err != nil {
		responseText := fmt.Sprintf(`Dynamic Findings Analysis - Error
========================

Application Path: %s
App Profile: %s
Application GUID: %s
Error: Failed to retrieve dynamic findings

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
	policyFilter := req.ViolatesPolicy != nil && *req.ViolatesPolicy
	return formatDynamicFindingsResponse(ctx, appProfile, applicationGUID, req.Sandbox, findingsResp, policyFilter), nil
}

// formatDynamicFindingsResponse formats the findings API response into an MCP tool response
func formatDynamicFindingsResponse(ctx context.Context, appProfile, applicationGUID, sandbox string, findings *api.FindingsResponse, policyFilter bool) map[string]interface{} {
	// Build MCP response structure
	response := MCPFindingsResponse{
		Application: MCPApplication{
			Name: appProfile,
			ID:   applicationGUID,
		},
		PolicyFilter: policyFilter,
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
			mcpFinding := processDynamicFinding(finding)
			response.Findings = append(response.Findings, mcpFinding)
			updateSummaryCounters(&response.Summary, mcpFinding)
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
		log.Printf("Dynamic findings: Returning %d findings (content: JSON, structuredContent: full data for UI)", len(response.Findings))
		result["structuredContent"] = response
	} else {
		log.Printf("Dynamic findings: Returning %d findings (content: JSON only, no structuredContent - client doesn't support UI)", len(response.Findings))
	}

	return result
}

// processDynamicFinding transforms a single API finding into an MCP finding
func processDynamicFinding(finding api.Finding) MCPFinding {
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
	cleanedDesc, references := TransformDescription(finding.Description, "DAST")

	// Extract numeric CWE ID
	var cweID int32
	if finding.CWE != "" {
		_, _ = fmt.Sscanf(finding.CWE, "CWE-%d", &cweID)
	}

	return MCPFinding{
		FlawID:           finding.ID,
		BuildID:          finding.BuildID,
		ScanType:         "DAST",
		Status:           string(transformedStatus),
		MitigationStatus: mitigationStatus,
		ViolatesPolicy:   violatesPolicy,
		Severity:         string(transformedSeverity),
		SeverityScore:    severityNum,
		CweId:            cweID,
		Description:      cleanedDesc,
		References:       references,
		URL:              finding.URL,
		AttackVector:     finding.AttackVector,
	}
}

// updateSummaryCounters updates the response summary based on a finding
func updateSummaryCounters(summary *MCPFindingsSummary, finding MCPFinding) {
	// Note: TotalFindings is set from pagination.total_elements, not incremented per finding

	if finding.Status == string(StatusOpen) {
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
	summary.ByStatus[statusKey]++

	// Count by mitigation status
	mitigationKey := strings.ToLower(finding.MitigationStatus)
	summary.ByMitigation[mitigationKey]++
}
