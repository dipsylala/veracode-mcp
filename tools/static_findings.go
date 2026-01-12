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

const StaticFindingsToolName = "get-static-findings"

// Auto-register this tool when the package is imported
func init() {
	RegisterTool(StaticFindingsToolName, func() ToolImplementation {
		return NewStaticFindingsTool()
	})
}

// StaticFindingsTool provides the get-static-findings tool
type StaticFindingsTool struct{}

// NewStaticFindingsTool creates a new static findings tool
func NewStaticFindingsTool() *StaticFindingsTool {
	return &StaticFindingsTool{}
}

// Initialize sets up the tool
func (t *StaticFindingsTool) Initialize() error {
	log.Printf("Initializing tool: %s", StaticFindingsToolName)
	// TODO: Initialize Veracode API client, load credentials, etc.
	return nil
}

// RegisterHandlers registers the static findings handler
func (t *StaticFindingsTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", StaticFindingsToolName)
	registry.RegisterHandler(StaticFindingsToolName, t.handleGetStaticFindings)
	return nil
}

// Shutdown cleans up tool resources
func (t *StaticFindingsTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", StaticFindingsToolName)
	return nil
}

// StaticFindingsRequest represents the parsed parameters for get-static-findings
type StaticFindingsRequest struct {
	ApplicationPath string `json:"application_path"`
	AppProfile      string `json:"app_profile,omitempty"`
	Sandbox         string `json:"sandbox,omitempty"`
	Size            int    `json:"size,omitempty"`
	Page            int    `json:"page,omitempty"`
	Severity        *int32 `json:"severity,omitempty"`
	SeverityGte     *int32 `json:"severity_gte,omitempty"`
}

// parseStaticFindingsRequest extracts and validates parameters from the raw args map
func parseStaticFindingsRequest(args map[string]interface{}) (*StaticFindingsRequest, error) {
	// Set defaults
	req := &StaticFindingsRequest{
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

func (t *StaticFindingsTool) handleGetStaticFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
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
		AppProfile:  applicationGUID,
		Sandbox:     req.Sandbox,
		Size:        req.Size,
		Page:        req.Page,
		Severity:    req.Severity,
		SeverityGte: req.SeverityGte,
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
	return formatStaticFindingsResponse(req.ApplicationPath, appProfile, applicationGUID, req.Sandbox, findingsResp), nil
}

// formatStaticFindingsResponse formats the findings API response into an MCP tool response
func formatStaticFindingsResponse(appPath, appProfile, applicationGUID, sandbox string, findings *api.FindingsResponse) map[string]interface{} {
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

			// Clean and extract references from description
			cleanedDesc, references := TransformDescription(finding.Description, "STATIC")

			mcpFinding := MCPFinding{
				FlawID:           finding.ID,
				ScanType:         "STATIC",
				Status:           finding.Status,
				MitigationStatus: "NONE", // TODO: Extract from API when available
				ViolatesPolicy:   finding.ViolatesPolicy,
				Severity:         string(transformedSeverity),
				SeverityScore:    severityNum,
				WeaknessType:     finding.CWE,
				WeaknessName:     finding.CWE, // TODO: Map CWE to name
				Description:      cleanedDesc,
				References:       references,
				FilePath:         finding.FilePath,
				LineNumber:       finding.LineNumber,
			}

			response.Findings = append(response.Findings, mcpFinding)

			// Update summary counts
			response.Summary.TotalFindings++
			if finding.Status == "OPEN" || finding.Status == "NEW" {
				response.Summary.OpenFindings++
			}
			if finding.ViolatesPolicy {
				response.Summary.PolicyViolations++
			}

			// Count by severity
			severityKey := strings.ToLower(string(transformedSeverity))
			response.Summary.BySeverity[severityKey]++

			// Count by scan type
			response.Summary.ByScanType["static"]++

			// Count by status
			statusKey := strings.ToLower(finding.Status)
			if statusKey == "new" {
				statusKey = "open"
			}
			response.Summary.ByStatus[statusKey]++
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
				"uri":      fmt.Sprintf("veracode://findings/static/%s", applicationGUID),
				"mimeType": "application/json",
				"text":     string(responseJSON),
			},
		}},
	}
}
