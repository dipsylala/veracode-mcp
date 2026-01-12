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
	RegisterTool("static-findings", func() ToolImplementation {
		return NewStaticFindingsTool()
	})
}

// StaticFindingsTool provides the get-static-findings tool
type StaticFindingsTool struct {
	name        string
	description string
}

// NewStaticFindingsTool creates a new static findings tool
func NewStaticFindingsTool() *StaticFindingsTool {
	return &StaticFindingsTool{
		name:        "static-findings",
		description: "Provides get-static-findings tool for retrieving source code vulnerabilities from SAST scans",
	}
}

// Name returns the tool name
func (t *StaticFindingsTool) Name() string {
	return t.name
}

// Description returns the tool description
func (t *StaticFindingsTool) Description() string {
	return t.description
}

// Initialize sets up the tool
func (t *StaticFindingsTool) Initialize() error {
	log.Printf("Initializing tool: %s", t.name)
	// TODO: Initialize Veracode API client, load credentials, etc.
	return nil
}

// RegisterHandlers registers the static findings handler
func (t *StaticFindingsTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", t.name)
	registry.RegisterHandler("get-static-findings", t.handleGetStaticFindings)
	return nil
}

// Shutdown cleans up tool resources
func (t *StaticFindingsTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", t.name)
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

	// Step 3: Get application GUID using the profile name
	application, err := client.GetApplicationByName(ctx, appProfile)
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
	if findings == nil || len(findings.Findings) == 0 {
		responseText := fmt.Sprintf(`Static Findings Analysis
========================

Application Path: %s
App Profile: %s
Application GUID: %s
Sandbox: %s

Status: ✓ No static findings found

The application has been scanned and no security vulnerabilities were detected in the source code.`,
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
	responseText := fmt.Sprintf(`Static Findings Analysis
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

		if finding.FilePath != "" {
			responseText += fmt.Sprintf("File: %s", finding.FilePath)
			if finding.LineNumber > 0 {
				responseText += fmt.Sprintf(":%d", finding.LineNumber)
			}
			responseText += "\n"
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
