package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dipsylala/veracodemcp-go/api"
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

	// Step 3: Get application GUID using the profile name
	application, err := client.GetApplicationByName(ctx, appProfile)
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

	applicationGUID := "unknown"
	if application.Guid != nil {
		applicationGUID = *application.Guid
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
	if findings == nil || len(findings.Findings) == 0 {
		responseText := fmt.Sprintf(`Dynamic Findings Analysis
========================

Application Path: %s
App Profile: %s
Application GUID: %s
Sandbox: %s

Status: ✓ No dynamic findings found

The application has been scanned and no runtime security vulnerabilities were detected.`,
			appPath,
			appProfile,
			applicationGUID,
			valueOrDefault(sandbox, "policy scan (production)"),
		)

		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": responseText,
			}},
		}
	}

	// Build findings summary
	responseText := fmt.Sprintf(`Dynamic Findings Analysis
========================

Application Path: %s
App Profile: %s
Application GUID: %s
Sandbox: %s

Total Findings: %d (showing %d)

`,
		appPath,
		appProfile,
		applicationGUID,
		valueOrDefault(sandbox, "policy scan (production)"),
		findings.TotalCount,
		len(findings.Findings),
	)

	// Add individual findings
	for i, finding := range findings.Findings {
		responseText += fmt.Sprintf("Finding #%d\n", i+1)
		responseText += "-----------\n"
		responseText += fmt.Sprintf("ID: %s\n", finding.ID)
		responseText += fmt.Sprintf("Severity: %s\n", finding.Severity)
		responseText += fmt.Sprintf("CWE: %s\n", finding.CWE)
		responseText += fmt.Sprintf("Status: %s\n", finding.Status)
		responseText += fmt.Sprintf("Description: %s\n", finding.Description)

		if finding.URL != "" {
			responseText += fmt.Sprintf("URL: %s\n", finding.URL)
		}

		if finding.ViolatesPolicy {
			responseText += "⚠️  VIOLATES POLICY\n"
		}

		responseText += "\n"
	}

	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": responseText,
		}},
	}
}

// Helper function to return value or default
func valueOrDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
