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

// Auto-register this tool when the package is imported
func init() {
	RegisterTool("dynamic-findings", func() ToolImplementation {
		return NewDynamicFindingsTool()
	})
}

// DynamicFindingsTool provides the get-dynamic-findings tool
type DynamicFindingsTool struct {
	name        string
	description string
}

// NewDynamicFindingsTool creates a new dynamic findings tool
func NewDynamicFindingsTool() *DynamicFindingsTool {
	return &DynamicFindingsTool{
		name:        "dynamic-findings",
		description: "Provides get-dynamic-findings tool for retrieving runtime vulnerabilities from DAST scans",
	}
}

// Name returns the tool name
func (t *DynamicFindingsTool) Name() string {
	return t.name
}

// Description returns the tool description
func (t *DynamicFindingsTool) Description() string {
	return t.description
}

// Initialize sets up the tool
func (t *DynamicFindingsTool) Initialize() error {
	log.Printf("Initializing tool: %s", t.name)
	// TODO: Initialize Veracode API client, load credentials, etc.
	return nil
}

// RegisterHandlers registers the dynamic findings handler
func (t *DynamicFindingsTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", t.name)
	registry.RegisterHandler("get-dynamic-findings", t.handleGetDynamicFindings)
	return nil
}

// Shutdown cleans up tool resources
func (t *DynamicFindingsTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", t.name)
	// TODO: Close API connections, cleanup resources
	return nil
}

// DynamicFindingsRequest represents the parsed parameters for get-dynamic-findings
type DynamicFindingsRequest struct {
	ApplicationPath string `json:"application_path"`
	AppProfile      string `json:"app_profile,omitempty"`
	Sandbox         string `json:"sandbox,omitempty"`
	Size            int    `json:"size,omitempty"`
	Page            int    `json:"page,omitempty"`
	Severity        *int32 `json:"severity,omitempty"`
	SeverityGte     *int32 `json:"severity_gte,omitempty"`
}

// parseDynamicFindingsRequest extracts and validates parameters from the raw args map
func parseDynamicFindingsRequest(args map[string]interface{}) (*DynamicFindingsRequest, error) {
	// Set defaults
	req := &DynamicFindingsRequest{
		Size: 50,
		Page: 0,
	}

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
	return formatDynamicFindingsResponse(req.ApplicationPath, appProfile, applicationGUID, req.Sandbox, findingsResp), nil
}

// formatDynamicFindingsResponse formats the findings API response into an MCP tool response
func formatDynamicFindingsResponse(appPath, appProfile, applicationGUID, sandbox string, findings *api.FindingsResponse) map[string]interface{} {
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
			ByScanType: map[string]int{
				"static": 0,
				"sca":    0,
				"dast":   0,
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
	}

	// Process each finding
	if findings != nil && len(findings.Findings) > 0 {
		for _, finding := range findings.Findings {
			// Transform severity
			var severityNum int32
			if finding.Severity != "" {
				_, _ = fmt.Sscanf(finding.Severity, "%d", &severityNum)
			}
			transformedSeverity := TransformSeverity(&severityNum)

			// Transform status (OPEN/CLOSED/UNKNOWN)
			// Create a minimal FindingStatus for transformation
			// Note: The API simplified Finding doesn't include full FindingStatus,
			// so we work with the status string directly
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

			// Extract mitigation status from mitigations if available
			mitigationStatus := MitigationNone
			if len(finding.Mitigations) > 0 {
				// Get the latest mitigation status
				latestStatus := strings.ToUpper(finding.Mitigations[len(finding.Mitigations)-1].Status)
				switch latestStatus {
				case "PROPOSED":
					mitigationStatus = MitigationProposed
				case "APPROVED":
					mitigationStatus = MitigationApproved
				case "REJECTED":
					mitigationStatus = MitigationRejected
				default:
					mitigationStatus = MitigationNone
				}
			}

			// Determine policy violation using business rule:
			// Only CLOSED + APPROVED findings don't violate policy
			violatesPolicyValue := finding.ViolatesPolicy
			violatesPolicy := DeterminesPolicyViolation(
				transformedStatus,
				mitigationStatus,
				&violatesPolicyValue,
			)

			// Clean and extract references from description (DAST may have base64)
			cleanedDesc, references := TransformDescription(finding.Description, "DAST")

			// Parse CWE for weakness type/name
			var cweID int32
			weaknessType := finding.CWE
			weaknessName := finding.CWE
			if finding.CWE != "" {
				// Try to extract numeric CWE ID
				if strings.HasPrefix(strings.ToUpper(finding.CWE), "CWE-") {
					_, _ = fmt.Sscanf(finding.CWE, "CWE-%d", &cweID)
					weaknessType = TransformWeaknessType(&cweID)
				}
				// TODO: Map CWE ID to name using a lookup table
			}

			mcpFinding := MCPFinding{
				FlawID:           finding.ID,
				ScanType:         "DAST",
				Status:           string(transformedStatus),
				MitigationStatus: string(mitigationStatus),
				ViolatesPolicy:   violatesPolicy,
				Severity:         string(transformedSeverity),
				SeverityScore:    severityNum,
				WeaknessType:     weaknessType,
				WeaknessName:     weaknessName,
				Description:      cleanedDesc,
				References:       references,
				URL:              finding.URL,
			}

			response.Findings = append(response.Findings, mcpFinding)

			// Update summary counts
			response.Summary.TotalFindings++
			if transformedStatus == StatusOpen {
				response.Summary.OpenFindings++
			}
			if violatesPolicy {
				response.Summary.PolicyViolations++
			}

			// Count by severity
			severityKey := strings.ToLower(string(transformedSeverity))
			response.Summary.BySeverity[severityKey]++

			// Count by scan type
			response.Summary.ByScanType["dast"]++

			// Count by status
			statusKey := strings.ToLower(string(transformedStatus))
			response.Summary.ByStatus[statusKey]++

			// Count by mitigation status
			mitigationKey := strings.ToLower(string(mitigationStatus))
			response.Summary.ByMitigation[mitigationKey]++
		}
	}

	// Marshal response to JSON string
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf("Error formatting response: %v", err),
			}},
			"isError": true,
		}
	}

	// Return as MCP tool response with JSON content
	return map[string]interface{}{
		"content": []map[string]interface{}{{
			"type": "resource",
			"resource": map[string]interface{}{
				"uri":      fmt.Sprintf("veracode://findings/dynamic/%s", applicationGUID),
				"mimeType": "application/json",
				"text":     string(responseJSON),
			},
		}},
	}
}
