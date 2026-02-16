package mcp_tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/dipsylala/veracode-mcp/api"
	"github.com/dipsylala/veracode-mcp/workspace"
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
	FlawID          string `json:"flaw_id"` // String to support both platform "12345" and pipeline "1234-1" formats
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

	req.FlawID, err = extractRequiredString(args, "flaw_id")
	if err != nil {
		return nil, err
	}

	// Validate flaw_id is not empty
	if req.FlawID == "" {
		return nil, fmt.Errorf("flaw_id cannot be empty")
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
func lookupApplicationGUID(ctx context.Context, client api.Client, appProfile string) (string, error) {
	authCtx := client.GetAuthContext(ctx)
	app, err := client.GetApplicationByName(authCtx, appProfile)
	if err != nil {
		return "", err
	}
	return *app.Guid, nil
}

// handleGetFindingDetails routes the request to the appropriate handler based on flaw ID format
func handleGetFindingDetails(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	req, err := parseFindingDetailsRequest(args)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}

	// Detect scan type based on flaw ID format
	// Pipeline flaws have format "1234-1", platform flaws are pure numeric "12345"
	if strings.Contains(req.FlawID, "-") {
		// Pipeline scan - route to pipeline logic
		return handlePipelineFindingDetails(ctx, req)
	}

	// Platform scan (static/dynamic) - route to platform logic
	return handlePlatformFindingDetails(ctx, req)
}
