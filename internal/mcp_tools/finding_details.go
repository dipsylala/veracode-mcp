package mcp_tools

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
	RegisterMCPTool(FindingDetailsToolName, handleGetFindingDetails)
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

	// Extract required fields
	var err error
	req.ApplicationPath, err = extractRequiredString(args, "application_path")
	if err != nil {
		return nil, err
	}

	req.FlawID, err = extractFlawID(args)
	if err != nil {
		return nil, err
	}

	// Extract optional fields
	req.AppProfile, _ = extractOptionalString(args, "app_profile")

	return req, nil
}

// getAppProfileName retrieves the application profile name from request or workspace config
func getAppProfileName(req *FindingDetailsRequest) (string, error) {
	if req.AppProfile != "" {
		return req.AppProfile, nil
	}
	return workspace.FindWorkspaceConfig(req.ApplicationPath)
}

// lookupApplicationGUID looks up the application GUID by profile name
func lookupApplicationGUID(ctx context.Context, client api.VeracodeClient, appProfile string) (string, error) {
	authCtx := client.GetAuthContext(ctx)
	app, err := client.GetApplicationByName(authCtx, appProfile)
	if err != nil {
		return "", err
	}
	return *app.Guid, nil
}

// tryGetStaticFlaw attempts to retrieve static flaw details
func tryGetStaticFlaw(ctx context.Context, client api.VeracodeClient, appGUID string, flawID int) (interface{}, error) {
	authCtx := client.GetAuthContext(ctx)
	issueIDStr := strconv.Itoa(flawID)

	staticFlawReq := client.StaticFindingDataPathClient().StaticFlawDataPathsInformationAPI.
		AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet(authCtx, appGUID, issueIDStr)

	staticFlaw, staticResp, staticErr := staticFlawReq.Execute()

	if staticErr == nil && staticResp != nil && staticResp.StatusCode == 200 && staticFlaw != nil {
		log.Printf("Found static flaw details for flaw ID %d", flawID)
		return staticFlaw, nil
	}

	if staticErr != nil {
		log.Printf("Static flaw lookup failed for flaw ID %d: %v", flawID, staticErr)
		return nil, staticErr
	}
	if staticResp != nil {
		log.Printf("Static flaw lookup returned status %d for flaw ID %d", staticResp.StatusCode, flawID)
	}
	return nil, fmt.Errorf("static flaw not found")
}

// tryGetDynamicFlaw attempts to retrieve dynamic flaw details
func tryGetDynamicFlaw(ctx context.Context, client api.VeracodeClient, appGUID string, flawID int) (interface{}, error) {
	authCtx := client.GetAuthContext(ctx)
	issueIDStr := strconv.Itoa(flawID)

	dynamicFlawReq := client.DynamicFlawClient().DefaultAPI.
		AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet(authCtx, appGUID, issueIDStr)

	dynamicFlaw, dynamicResp, dynamicErr := dynamicFlawReq.Execute()

	if dynamicErr == nil && dynamicResp != nil && dynamicResp.StatusCode == 200 && dynamicFlaw != nil {
		log.Printf("Found dynamic flaw details for flaw ID %d", flawID)
		return dynamicFlaw, nil
	}

	if dynamicErr != nil {
		log.Printf("Dynamic flaw lookup failed for flaw ID %d: %v", flawID, dynamicErr)
		return nil, dynamicErr
	}
	if dynamicResp != nil {
		log.Printf("Dynamic flaw lookup returned status %d for flaw ID %d", dynamicResp.StatusCode, flawID)
	}
	return nil, fmt.Errorf("dynamic flaw not found")
}

// formatErrorResponse creates a formatted error response
func formatErrorResponse(appPath, appProfile, appGUID string, flawID int, staticErr, dynamicErr error) map[string]interface{} {
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
`, appPath, appProfile, appGUID, flawID, errorMsg),
		}},
	}
}

func handleGetFindingDetails(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseFindingDetailsRequest(args)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}

	// Get application profile name
	appProfile, err := getAppProfileName(req)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf("Finding Details Lookup\n======================\n\nApplication Path: %s\n\n❌ Error: %v\n", req.ApplicationPath, err),
			}},
		}, nil
	}

	// Create Veracode API client
	client, err := api.NewVeracodeClient()
	if err != nil {
		return map[string]interface{}{"error": fmt.Sprintf("Failed to create Veracode API client: %v", err)}, nil
	}

	// Look up application GUID
	appGUID, err := lookupApplicationGUID(ctx, client, appProfile)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf("Finding Details Lookup\n======================\n\nApplication Path: %s\nApplication Profile: %s\n\n❌ Error: Failed to lookup application\n\n%v\n\nPlease verify that the application profile name is correct and that you have access to this application in Veracode.\n", req.ApplicationPath, appProfile, err),
			}},
		}, nil
	}

	// Try to get static flaw details first
	staticFlaw, staticErr := tryGetStaticFlaw(ctx, client, appGUID, req.FlawID)
	if staticErr == nil {
		return formatStaticFlawDetailsResponse(req.ApplicationPath, appProfile, req.FlawID, staticFlaw), nil
	}

	// Try dynamic flaw
	dynamicFlaw, dynamicErr := tryGetDynamicFlaw(ctx, client, appGUID, req.FlawID)
	if dynamicErr == nil {
		return formatDynamicFlawDetailsResponse(req.ApplicationPath, appProfile, req.FlawID, dynamicFlaw), nil
	}

	// Neither found - return error
	return formatErrorResponse(req.ApplicationPath, appProfile, appGUID, req.FlawID, staticErr, dynamicErr), nil
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
