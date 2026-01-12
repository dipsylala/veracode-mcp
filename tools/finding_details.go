package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/dipsylala/veracodemcp-go/api"
	"github.com/dipsylala/veracodemcp-go/workspace"
)

const FindingDetailsToolName = "get-finding-details"

// Auto-register this tool when the package is imported
func init() {
	RegisterTool(FindingDetailsToolName, func() ToolImplementation {
		return NewFindingDetailsTool()
	})
}

// FindingDetailsTool provides the get-finding-details tool
type FindingDetailsTool struct{}

// NewFindingDetailsTool creates a new finding details tool
func NewFindingDetailsTool() *FindingDetailsTool {
	return &FindingDetailsTool{}
}

// Initialize sets up the tool
func (t *FindingDetailsTool) Initialize() error {
	log.Printf("Initializing tool: %s", FindingDetailsToolName)
	return nil
}

// RegisterHandlers registers the finding details handler
func (t *FindingDetailsTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", FindingDetailsToolName)
	registry.RegisterHandler(FindingDetailsToolName, t.handleGetFindingDetails)
	return nil
}

// Shutdown cleans up tool resources
func (t *FindingDetailsTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", FindingDetailsToolName)
	return nil
}

// FindingDetailsRequest represents the parsed parameters for get-finding-details
type FindingDetailsRequest struct {
	ApplicationPath string `json:"application_path"`
	AppProfile      string `json:"app_profile,omitempty"`
	FlawID          int    `json:"flaw_id"`
}

// parseFindingDetailsRequest extracts and validates parameters from the raw args map
func parseFindingDetailsRequest(args map[string]interface{}) (*FindingDetailsRequest, error) {
	req := &FindingDetailsRequest{}

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

	if req.FlawID == 0 {
		return nil, fmt.Errorf("flaw_id is required and must be a non-zero integer")
	}

	return req, nil
}

func (t *FindingDetailsTool) handleGetFindingDetails(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseFindingDetailsRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}

	// Step 1: Retrieve application profile name
	appProfile := req.AppProfile
	hasAppProfile := appProfile != ""
	if !hasAppProfile {
		// Load workspace configuration
		appProfile, err = workspace.FindWorkspaceConfig(req.ApplicationPath)
		if err != nil {
			return map[string]interface{}{
				"content": []map[string]string{{
					"type": "text",
					"text": fmt.Sprintf(`Finding Details Lookup
======================

Application Path: %s

❌ Error: %v
`, req.ApplicationPath, err),
				}},
			}, nil
		}
	}

	// Step 2: Create Veracode API client
	client, err := api.NewVeracodeClient()
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to create Veracode API client: %v", err),
		}, nil
	}

	authCtx := client.GetAuthContext(ctx)

	// Step 3: Look up application GUID by profile name
	app, err := client.GetApplicationByName(authCtx, appProfile)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Finding Details Lookup
======================

Application Path: %s
Application Profile: %s

❌ Error: Failed to lookup application

%v

Please verify that the application profile name is correct and that you have access to this application in Veracode.
`, req.ApplicationPath, appProfile, err),
			}},
		}, nil
	}

	appGUID := *app.Guid

	// Step 4: Try to get static flaw details first
	issueIDStr := strconv.Itoa(req.FlawID)
	staticFlawReq := client.StaticFindingDataPathClient().StaticFlawDataPathsInformationAPI.
		AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet(authCtx, appGUID, issueIDStr)

	staticFlaw, staticResp, staticErr := staticFlawReq.Execute()

	// Check if static flaw was found
	if staticErr == nil && staticResp != nil && staticResp.StatusCode == 200 && staticFlaw != nil {
		// Successfully retrieved static flaw details
		log.Printf("Found static flaw details for flaw ID %d", req.FlawID)
		return formatStaticFlawDetailsResponse(req.ApplicationPath, appProfile, req.FlawID, staticFlaw), nil
	}

	// Log why static flaw wasn't found
	if staticErr != nil {
		log.Printf("Static flaw lookup failed for flaw ID %d: %v", req.FlawID, staticErr)
	} else if staticResp != nil {
		log.Printf("Static flaw lookup returned status %d for flaw ID %d", staticResp.StatusCode, req.FlawID)
	}

	// Step 5: Static flaw not found, try dynamic flaw
	dynamicFlawReq := client.DynamicFlawClient().DefaultAPI.
		AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet(authCtx, appGUID, issueIDStr)

	dynamicFlaw, dynamicResp, dynamicErr := dynamicFlawReq.Execute()

	// Check if dynamic flaw was found
	if dynamicErr == nil && dynamicResp != nil && dynamicResp.StatusCode == 200 && dynamicFlaw != nil {
		// Successfully retrieved dynamic flaw details
		log.Printf("Found dynamic flaw details for flaw ID %d", req.FlawID)
		return formatDynamicFlawDetailsResponse(req.ApplicationPath, appProfile, req.FlawID, dynamicFlaw), nil
	}

	// Log why dynamic flaw wasn't found
	if dynamicErr != nil {
		log.Printf("Dynamic flaw lookup failed for flaw ID %d: %v", req.FlawID, dynamicErr)
	} else if dynamicResp != nil {
		log.Printf("Dynamic flaw lookup returned status %d for flaw ID %d", dynamicResp.StatusCode, req.FlawID)
	}

	// Step 6: Neither found - return error
	errorMsg := "Not found in either static or dynamic scan results"
	if staticErr != nil {
		errorMsg = fmt.Sprintf("Static API error: %v", staticErr)
		if dynamicErr != nil {
			errorMsg = fmt.Sprintf("Static API error: %v; Dynamic API error: %v", staticErr, dynamicErr)
		}
	} else if dynamicErr != nil {
		errorMsg = fmt.Sprintf("Dynamic API error: %v", dynamicErr)
	}

	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": fmt.Sprintf(`Finding Details Lookup
======================

Application Path: %s
Application Profile: %s
Application GUID: %s
Flaw ID: %d

❌ Finding not found

The specified flaw ID was not found in either static or dynamic scan results for this application.

Error details: %s

Please verify that:
1. The flaw ID is correct
2. The flaw belongs to this application
3. You have access to view findings for this application
`, req.ApplicationPath, appProfile, appGUID, req.FlawID, errorMsg),
		}},
	}, nil
}

// formatStaticFlawDetailsResponse formats static flaw details into an MCP response
func formatStaticFlawDetailsResponse(appPath, appProfile string, flawID int, flaw interface{}) map[string]interface{} {
	// Marshal the flaw data to JSON for display
	flawJSON, err := json.MarshalIndent(flaw, "", "  ")
	if err != nil {
		flawJSON = []byte(fmt.Sprintf("Error formatting flaw data: %v", err))
	}

	header := fmt.Sprintf(`Finding Details (Static Analysis)
==================================

Application Path: %s
Application Profile: %s
Flaw ID: %d
Scan Type: STATIC

`, appPath, appProfile, flawID)

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": header,
			},
			{
				"type": "text",
				"text": string(flawJSON),
			},
		},
	}
}

// formatDynamicFlawDetailsResponse formats dynamic flaw details into an MCP response
func formatDynamicFlawDetailsResponse(appPath, appProfile string, flawID int, flaw interface{}) map[string]interface{} {
	// Marshal the flaw data to JSON for display
	flawJSON, err := json.MarshalIndent(flaw, "", "  ")
	if err != nil {
		flawJSON = []byte(fmt.Sprintf("Error formatting flaw data: %v", err))
	}

	header := fmt.Sprintf(`Finding Details (Dynamic Analysis)
===================================

Application Path: %s
Application Profile: %s
Flaw ID: %d
Scan Type: DYNAMIC

`, appPath, appProfile, flawID)

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": header,
			},
			{
				"type": "text",
				"text": string(flawJSON),
			},
		},
	}
}
