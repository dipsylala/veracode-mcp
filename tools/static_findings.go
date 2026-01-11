package tools

import (
	"context"
	"fmt"
	"log"
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

func (t *StaticFindingsTool) handleGetStaticFindings(ctx context.Context, args map[string]interface{}) (interface{}, error) {
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
	responseText := fmt.Sprintf(`Static Findings Analysis
========================

Application Path: %s
App Profile: %s
Sandbox: %s

Status: Ready to scan for source code vulnerabilities

This tool would:
1. Validate the application path exists
2. Check for .veracode-workspace.json in the directory
3. Connect to Veracode API using credentials
4. Fetch STATIC scan findings for the application
5. Filter by: %v
6. Return findings with file paths and line numbers
7. Generate clickable VS Code links for each finding

Finding types detected:
- XSS (Cross-site Scripting)
- SQL Injection
- Cryptographic Issues
- Race Conditions
- Buffer Overflows
- Unsafe Function Usage

Next steps:
- Implement Veracode API client
- Add authentication handling
- Process and transform findings
- Generate workspace file links`,
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
