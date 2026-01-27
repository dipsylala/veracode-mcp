package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dipsylala/veracodemcp-go/api"
	"github.com/dipsylala/veracodemcp-go/api/generated/applications"
	"github.com/dipsylala/veracodemcp-go/workspace"
)

const DynamicFindingsToolName = "dynamic-findings"

// Auto-register this tool when the package is imported
func init() {
	RegisterTool(DynamicFindingsToolName, func() ToolImplementation {
		return NewDynamicFindingsTool()
	})
}

// DynamicFindingsTool provides the dynamic-findings tool
type DynamicFindingsTool struct{}

// NewDynamicFindingsTool creates a new dynamic findings tool
func NewDynamicFindingsTool() *DynamicFindingsTool {
	return &DynamicFindingsTool{}
}

// Initialize sets up the tool
func (t *DynamicFindingsTool) Initialize() error {
	log.Printf("Initializing tool: %s", DynamicFindingsToolName)
	// TODO: Initialize Veracode API client, load credentials, etc.
	return nil
}

// RegisterHandlers registers the dynamic findings handler
func (t *DynamicFindingsTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", DynamicFindingsToolName)
	registry.RegisterHandler(DynamicFindingsToolName, t.handleGetDynamicFindings)
	return nil
}

// Shutdown cleans up tool resources
func (t *DynamicFindingsTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", DynamicFindingsToolName)
	// TODO: Close API connections, cleanup resources
	return nil
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
}

// parseDynamicFindingsRequest extracts and validates parameters from the raw args map
func parseDynamicFindingsRequest(args map[string]interface{}) (*DynamicFindingsRequest, error) {
	// Set defaults
	req := &DynamicFindingsRequest{
		Size: 200,
		Page: 0,
	}

	// Extract cwe_ids separately since they may come as numbers
	if cweIdsRaw, ok := args["cwe_ids"]; ok {
		if cweArray, ok := cweIdsRaw.([]interface{}); ok {
			req.CWEIDs = make([]string, len(cweArray))
			for i, cwe := range cweArray {
				switch v := cwe.(type) {
				case float64:
					req.CWEIDs[i] = fmt.Sprintf("%.0f", v)
				case int:
					req.CWEIDs[i] = fmt.Sprintf("%d", v)
				case string:
					req.CWEIDs[i] = v
				default:
					req.CWEIDs[i] = fmt.Sprintf("%v", v)
				}
			}
		}
		delete(args, "cwe_ids") // Remove to avoid unmarshaling conflict
	}

	// Use JSON marshaling to automatically map remaining args to struct
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

	return req, nil
}

func (t *DynamicFindingsTool) handleGetDynamicFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
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
	client, err := api.NewVeracodeClient()
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
		AppProfile:  applicationGUID,
		Sandbox:     req.Sandbox,
		Size:        req.Size,
		Page:        req.Page,
		Severity:    req.Severity,
		SeverityGte: req.SeverityGte,
		CWEIDs:      req.CWEIDs,
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
	return formatDynamicFindingsResponse(ctx, appProfile, applicationGUID, req.Sandbox, findingsResp), nil
}

// formatDynamicFindingsResponse formats the findings API response into an MCP tool response
func formatDynamicFindingsResponse(ctx context.Context, appProfile, applicationGUID, sandbox string, findings *api.FindingsResponse) map[string]interface{} {
	// Build MCP response structure
	response := MCPFindingsResponse{
		Application: MCPApplication{
			Name: appProfile,
			ID:   applicationGUID,
		},
		Summary: MCPFindingsSummary{
			BySeverity: map[string]int{
				"critical":      0,
				"high":          0,
				"medium":        0,
				"low":           0,
				"informational": 0,
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

	// Marshal response to JSON for non-UI clients
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("Warning: Failed to marshal response to JSON: %v", err)
		responseJSON, _ = json.Marshal(response) // Fall back to compact JSON
	}

	log.Printf("Dynamic findings: Returning %d findings (content: JSON, structuredContent: full data for UI)", len(response.Findings))

	// Return MCP response with content and structured data
	// - content: Full JSON data (for LLM and non-UI clients)
	// - structuredContent: Full structured data (for MCP Apps UI rendering)
	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(responseJSON),
			},
		},
		"structuredContent": response,
	}
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
		Mitigations:      convertDynamicMitigations(finding.Mitigations),
	}
}

// convertDynamicMitigations converts API mitigations to MCP mitigations
func convertDynamicMitigations(apiMitigations []api.Mitigation) []MCPMitigation {
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
