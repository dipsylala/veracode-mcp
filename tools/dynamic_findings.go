package tools

import (
	"context"
	"fmt"
	"log"
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

func (t *DynamicFindingsTool) handleGetDynamicFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Extract application_path (required)
	appPath, ok := args["application_path"].(string)
	if !ok || appPath == "" {
		return map[string]interface{}{
			"error": "application_path is required and must be an absolute path",
		}, nil
	}

	// Extract optional parameters
	appProfile, _ := args["app_profile"].(string)
	sandbox, _ := args["sandbox"].(string)

	// Build response
	responseText := fmt.Sprintf(`Dynamic Findings Analysis
========================

Application Path: %s
App Profile: %s
Sandbox: %s

Status: Ready to scan for runtime vulnerabilities

This tool would:
1. Validate the application path exists
2. Check for .veracode-workspace.json in the directory
3. Connect to Veracode API using credentials
4. Fetch DYNAMIC scan findings for the application
5. Filter by: %v
6. Return findings in structured format

Next steps:
- Implement Veracode API client
- Add authentication handling
- Process and transform findings
- Generate clickable file links for VS Code`,
		appPath,
		valueOrDefault(appProfile, "auto-detect from workspace"),
		valueOrDefault(sandbox, "policy scan (production)"),
		args,
	)

	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": responseText,
		}},
	}, nil
}

// Helper function to return value or default
func valueOrDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
