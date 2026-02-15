package mcp_tools

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/dipsylala/veracode-mcp/api"
	dynamicflaw "github.com/dipsylala/veracode-mcp/api/rest/generated/dynamic_flaw"
	staticfindingdatapath "github.com/dipsylala/veracode-mcp/api/rest/generated/static_finding_data_path"
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
func lookupApplicationGUID(ctx context.Context, client api.Client, appProfile string) (string, error) {
	authCtx := client.GetAuthContext(ctx)
	app, err := client.GetApplicationByName(authCtx, appProfile)
	if err != nil {
		return "", err
	}
	return *app.Guid, nil
}

// tryGetStaticFlaw attempts to retrieve static flaw details
func tryGetStaticFlaw(ctx context.Context, client api.Client, appGUID string, flawID int) (interface{}, *int64, error) {
	authCtx := client.GetAuthContext(ctx)
	issueIDStr := strconv.Itoa(flawID)

	staticFlawReq := client.StaticFindingDataPathClient().StaticFlawDataPathsInformationAPI.
		AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet(authCtx, appGUID, issueIDStr)

	staticFlaw, staticResp, staticErr := staticFlawReq.Execute()

	if staticErr == nil && staticResp != nil && staticResp.StatusCode == 200 && staticFlaw != nil {
		log.Printf("Found static flaw details for flaw ID %d", flawID)
		// Extract build_id from IssueSummary
		var buildID *int64
		if staticFlaw.IssueSummary != nil && staticFlaw.IssueSummary.BuildId != nil {
			//nolint:gosec // G115: Safe conversion from int32 to int64
			buildIDInt64 := int64(*staticFlaw.IssueSummary.BuildId)
			buildID = &buildIDInt64
			log.Printf("Extracted build_id %d from static flaw", buildIDInt64)
		}
		return staticFlaw, buildID, nil
	}

	if staticErr != nil {
		log.Printf("Static flaw lookup failed for flaw ID %d: %v", flawID, staticErr)
		return nil, nil, staticErr
	}
	if staticResp != nil {
		log.Printf("Static flaw lookup returned status %d for flaw ID %d", staticResp.StatusCode, flawID)
	}
	return nil, nil, fmt.Errorf("static flaw not found")
}

// tryGetDynamicFlaw attempts to retrieve dynamic flaw details
func tryGetDynamicFlaw(ctx context.Context, client api.Client, appGUID string, flawID int) (interface{}, *int64, error) {
	authCtx := client.GetAuthContext(ctx)
	issueIDStr := strconv.Itoa(flawID)

	dynamicFlawReq := client.DynamicFlawClient().DefaultAPI.
		AppsecV2ApplicationsAppGuidFindingsIssueIdDynamicFlawInfoGet(authCtx, appGUID, issueIDStr)

	dynamicFlaw, dynamicResp, dynamicErr := dynamicFlawReq.Execute()

	if dynamicErr == nil && dynamicResp != nil && dynamicResp.StatusCode == 200 && dynamicFlaw != nil {
		log.Printf("Found dynamic flaw details for flaw ID %d", flawID)
		// Extract build_id from IssueSummary
		var buildID *int64
		if dynamicFlaw.IssueSummary != nil && dynamicFlaw.IssueSummary.BuildId != nil {
			//nolint:gosec // G115: Safe conversion from int32 to int64
			buildIDInt64 := int64(*dynamicFlaw.IssueSummary.BuildId)
			buildID = &buildIDInt64
			log.Printf("Extracted build_id %d from dynamic flaw", buildIDInt64)
		}
		return dynamicFlaw, buildID, nil
	}

	if dynamicErr != nil {
		log.Printf("Dynamic flaw lookup failed for flaw ID %d: %v", flawID, dynamicErr)
		return nil, nil, dynamicErr
	}
	if dynamicResp != nil {
		log.Printf("Dynamic flaw lookup returned status %d for flaw ID %d", dynamicResp.StatusCode, flawID)
	}
	return nil, nil, fmt.Errorf("dynamic flaw not found")
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
	client, err := api.NewClient()
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
	staticFlaw, buildID, staticErr := tryGetStaticFlaw(ctx, client, appGUID, req.FlawID)
	if staticErr == nil {
		// Get mitigation info if we have a build_id
		var mitigationInfo *api.MitigationIssue
		if buildID != nil {
			log.Printf("Fetching mitigation info for build_id=%d, flaw_id=%d", *buildID, req.FlawID)
			//nolint:gosec // G115: Safe conversion from int to int64
			mitigationInfo, _ = client.GetMitigationInfoForSingleFlaw(ctx, *buildID, int64(req.FlawID))
			if mitigationInfo != nil {
				log.Printf("Retrieved %d mitigation actions", len(mitigationInfo.MitigationActions))
			}
		}
		return formatStaticFlawDetailsResponse(req.ApplicationPath, appProfile, req.FlawID, staticFlaw, mitigationInfo), nil
	}

	// Try dynamic flaw
	dynamicFlaw, buildID, dynamicErr := tryGetDynamicFlaw(ctx, client, appGUID, req.FlawID)
	if dynamicErr == nil {
		// Get mitigation info if we have a build_id
		var mitigationInfo *api.MitigationIssue
		if buildID != nil {
			log.Printf("Fetching mitigation info for build_id=%d, flaw_id=%d", *buildID, req.FlawID)
			//nolint:gosec // G115: Safe conversion from int to int64
			mitigationInfo, _ = client.GetMitigationInfoForSingleFlaw(ctx, *buildID, int64(req.FlawID))
			if mitigationInfo != nil {
				log.Printf("Retrieved %d mitigation actions", len(mitigationInfo.MitigationActions))
			}
		}
		return formatDynamicFlawDetailsResponse(req.ApplicationPath, appProfile, req.FlawID, dynamicFlaw, mitigationInfo), nil
	}

	// Neither found - return error
	return formatErrorResponse(req.ApplicationPath, appProfile, appGUID, req.FlawID, staticErr, dynamicErr), nil
}

// formatStaticFlawDetailsResponse formats static flaw details into an MCP response
func formatStaticFlawDetailsResponse(appPath, appProfile string, flawID int, flaw interface{}, mitigationInfo *api.MitigationIssue) map[string]interface{} {
	// Build the LLM-optimized JSON structure
	result := buildLLMOptimizedResponse(appPath, appProfile, flawID, "STATIC", flaw, mitigationInfo)

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting flaw details: %v", err),
				},
			},
		}
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(jsonBytes),
			},
		},
	}
}

// sortMitigationActionsByDate sorts mitigation actions by date (oldest to newest)
func sortMitigationActionsByDate(actions []api.MitigationAction) {
	sort.Slice(actions, func(i, j int) bool {
		// Parse dates - format appears to be "2025-10-13 04:28:33"
		timeI, errI := time.Parse("2006-01-02 15:04:05", actions[i].Date)
		timeJ, errJ := time.Parse("2006-01-02 15:04:05", actions[j].Date)

		// If parsing fails, fall back to string comparison
		if errI != nil || errJ != nil {
			return actions[i].Date < actions[j].Date
		}

		return timeI.Before(timeJ)
	})
}

// buildLLMOptimizedResponse creates an LLM-optimized JSON structure for flaw details
func buildLLMOptimizedResponse(appPath, appProfile string, flawID int, scanType string, flaw interface{}, mitigationInfo *api.MitigationIssue) map[string]interface{} {
	response := map[string]interface{}{
		"flaw_id": flawID,
		"application": map[string]interface{}{
			"name": appProfile,
			"path": appPath,
		},
		"scan_type": scanType,
	}

	// Extract vulnerability information based on scan type
	if scanType == "STATIC" {
		if staticFlaw, ok := flaw.(*staticfindingdatapath.StaticFlaws); ok {
			extractStaticFlawInfo(response, staticFlaw)
		}
	} else if scanType == "DYNAMIC" {
		if dynamicFlaw, ok := flaw.(*dynamicflaw.DynamicFlaw); ok {
			extractDynamicFlawInfo(response, dynamicFlaw)
		}
	}

	// Add mitigation information if available
	if mitigationInfo != nil {
		response["mitigation"] = buildMitigationInfo(mitigationInfo)
	}

	return response
}

// extractStaticFlawInfo extracts static flaw information into the response
func extractStaticFlawInfo(response map[string]interface{}, staticFlaw *staticfindingdatapath.StaticFlaws) {
	// Note: The detailed flaw API doesn't return vulnerability metadata (CWE, severity, description)
	// Those are only available in the findings list API

	// Extract location from first data path if available
	if len(staticFlaw.DataPaths) > 0 && len(staticFlaw.DataPaths[0].Calls) > 0 {
		firstPath := staticFlaw.DataPaths[0]
		// The sink is typically the last call in the path
		sinkCall := firstPath.Calls[len(firstPath.Calls)-1]

		response["location"] = map[string]interface{}{
			"file_path":   safeString(sinkCall.FilePath),
			"file_name":   safeString(sinkCall.FileName),
			"function":    safeString(sinkCall.FunctionName),
			"line_number": safeInt32(sinkCall.LineNumber),
		}

		if firstPath.ModuleName != nil {
			response["location"].(map[string]interface{})["module"] = *firstPath.ModuleName
		}
	}

	// Extract data paths
	if len(staticFlaw.DataPaths) > 0 {
		dataPaths := make([]map[string]interface{}, 0, len(staticFlaw.DataPaths))
		for idx, dataPath := range staticFlaw.DataPaths {
			pathInfo := map[string]interface{}{
				"path_id":     idx + 1,
				"total_steps": len(dataPath.Calls),
			}

			// Build steps array
			steps := make([]map[string]interface{}, 0, len(dataPath.Calls))
			for stepIdx, call := range dataPath.Calls {
				// Determine step type based on position
				stepType := "propagation"
				if stepIdx == 0 {
					stepType = "source"
				} else if stepIdx == len(dataPath.Calls)-1 {
					stepType = "sink"
				}

				step := map[string]interface{}{
					"step_number": stepIdx + 1,
					"type":        stepType,
					"file_path":   safeString(call.FilePath),
					"file_name":   safeString(call.FileName),
					"function":    safeString(call.FunctionName),
					"line_number": safeInt32(call.LineNumber),
				}
				steps = append(steps, step)
			}
			pathInfo["steps"] = steps
			dataPaths = append(dataPaths, pathInfo)
		}
		response["data_paths"] = dataPaths
	}
}

// extractDynamicFlawInfo extracts dynamic flaw information into the response
func extractDynamicFlawInfo(response map[string]interface{}, dynamicFlaw *dynamicflaw.DynamicFlaw) {
	// Note: The detailed flaw API doesn't return vulnerability metadata (CWE, severity, description)
	// Those are only available in the findings list API
	// For dynamic flaws, there are no data paths like static
	// But we can include the URL and request/response info

	if dynamicFlaw.DynamicFlawInfo != nil {
		if dynamicFlaw.DynamicFlawInfo.Request != nil && dynamicFlaw.DynamicFlawInfo.Request.Url != nil {
			response["location"] = map[string]interface{}{
				"url": *dynamicFlaw.DynamicFlawInfo.Request.Url,
			}
		}

		// Add HTTP request/response details if available
		reqResp := make(map[string]interface{})
		if dynamicFlaw.DynamicFlawInfo.Request != nil && dynamicFlaw.DynamicFlawInfo.Request.RawBytes != nil {
			if decoded, err := base64.StdEncoding.DecodeString(*dynamicFlaw.DynamicFlawInfo.Request.RawBytes); err == nil {
				reqResp["http_request"] = string(decoded)
			}
		}
		if dynamicFlaw.DynamicFlawInfo.Response != nil && dynamicFlaw.DynamicFlawInfo.Response.RawBytes != nil {
			if decoded, err := base64.StdEncoding.DecodeString(*dynamicFlaw.DynamicFlawInfo.Response.RawBytes); err == nil {
				reqResp["http_response"] = string(decoded)
			}
		}
		if len(reqResp) > 0 {
			response["http_details"] = reqResp
		}
	}
}

// buildMitigationInfo creates mitigation information structure
func buildMitigationInfo(mitigationInfo *api.MitigationIssue) map[string]interface{} {
	if mitigationInfo == nil {
		return nil
	}

	// Sort actions by date (oldest to newest)
	sortMitigationActionsByDate(mitigationInfo.MitigationActions)

	mitigation := map[string]interface{}{
		"category":     mitigationInfo.Category,
		"action_count": len(mitigationInfo.MitigationActions),
	}

	// Determine current status from the last action
	if len(mitigationInfo.MitigationActions) > 0 {
		lastAction := mitigationInfo.MitigationActions[len(mitigationInfo.MitigationActions)-1]
		mitigation["status"] = lastAction.Action

		// If last action is accepted/rejected, include approval details
		if lastAction.Action == "accepted" || lastAction.Action == "rejected" {
			mitigation["approved_by"] = lastAction.Reviewer
			mitigation["approved_date"] = lastAction.Date
			if lastAction.Comment != "" {
				mitigation["approval_comment"] = lastAction.Comment
			}
		}
	}

	// Add all actions
	if len(mitigationInfo.MitigationActions) > 0 {
		actions := make([]map[string]interface{}, 0, len(mitigationInfo.MitigationActions))
		for _, action := range mitigationInfo.MitigationActions {
			actionInfo := map[string]interface{}{
				"action_type": action.Action,
				"reviewer":    action.Reviewer,
				"date":        action.Date,
			}
			if action.Desc != "" {
				actionInfo["description"] = action.Desc
			}
			if action.Comment != "" {
				actionInfo["comment"] = action.Comment
			}
			actions = append(actions, actionInfo)
		}
		mitigation["actions"] = actions
	}

	return mitigation
}

// safeString safely dereferences a string pointer
func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// safeInt32 safely dereferences an int32 pointer
func safeInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// formatDynamicFlawDetailsResponse formats dynamic flaw details into an MCP response
func formatDynamicFlawDetailsResponse(appPath, appProfile string, flawID int, flaw interface{}, mitigationInfo *api.MitigationIssue) map[string]interface{} {
	// Build the LLM-optimized JSON structure
	result := buildLLMOptimizedResponse(appPath, appProfile, flawID, "DYNAMIC", flaw, mitigationInfo)

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Error formatting flaw details: %v", err),
				},
			},
		}
	}

	return map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": string(jsonBytes),
			},
		},
	}
}
