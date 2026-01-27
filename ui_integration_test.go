package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestPipelineResultsUIIntegration verifies the complete pipeline-results UI flow
func TestPipelineResultsUIIntegration(t *testing.T) {
	// Setup server
	server, err := NewMCPServer()
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Initialize
	initParams := InitializeParams{
		ProtocolVersion: "2024-11-05",
		Capabilities:    ClientCapabilities{},
		ClientInfo: Implementation{
			Name:    "test-client",
			Version: "1.0.0",
		},
	}

	initParamsJSON, _ := json.Marshal(initParams)
	_, err = server.handleInitialize(initParamsJSON)
	if err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}

	t.Run("tools_list_includes_ui_metadata", func(t *testing.T) {
		result := server.handleListTools()

		var pipelineTool *Tool
		for i := range result.Tools {
			if result.Tools[i].Name == "pipeline-results" {
				pipelineTool = &result.Tools[i]
				break
			}
		}

		if pipelineTool == nil {
			t.Fatal("pipeline-results tool not found")
		}

		if pipelineTool.Meta == nil {
			t.Fatal("Tool metadata is nil")
		}

		// Check for BOTH formats (flat key and nested) for compatibility
		flatUri, okFlat := pipelineTool.Meta["ui/resourceUri"].(string)
		if !okFlat {
			t.Errorf("Tool should have _meta['ui/resourceUri'], got: %+v", pipelineTool.Meta)
		}

		nestedUI, okNested := pipelineTool.Meta["ui"].(map[string]interface{})
		if !okNested {
			t.Errorf("Tool should have _meta['ui'], got: %+v", pipelineTool.Meta)
		} else {
			nestedUri, okUri := nestedUI["resourceUri"].(string)
			if !okUri {
				t.Errorf("Tool should have _meta.ui.resourceUri, got: %+v", nestedUI)
			}
			expectedURI := "ui://pipeline-results/app.html"
			if nestedUri != expectedURI {
				t.Errorf("Expected nested resourceUri=%s, got %s", expectedURI, nestedUri)
			}
		}

		expectedURI := "ui://pipeline-results/app.html"
		if flatUri != expectedURI {
			t.Errorf("Expected flat resourceUri=%s, got %s", expectedURI, flatUri)
		}

		// Print full tool as JSON to verify serialization
		toolJSON, _ := json.MarshalIndent(pipelineTool, "", "  ")
		t.Logf("Full tool JSON:\n%s", string(toolJSON))
		t.Logf("✓ Tool metadata: %+v", pipelineTool.Meta)
	})

	t.Run("resources_list_includes_ui_resource", func(t *testing.T) {
		result := server.handleListResources()

		var uiResource *Resource
		for i := range result.Resources {
			if result.Resources[i].URI == "ui://pipeline-results/app.html" {
				uiResource = &result.Resources[i]
				break
			}
		}

		if uiResource == nil {
			t.Fatal("UI resource not found in resources list")
		}

		if uiResource.MimeType != "text/html;profile=mcp-app" {
			t.Errorf("Expected MimeType 'text/html;profile=mcp-app', got '%s'", uiResource.MimeType)
		}

		t.Logf("✓ UI Resource: %+v", uiResource)
	})

	t.Run("ui_resource_is_embedded_and_has_csp", func(t *testing.T) {
		params := ReadResourceParams{
			URI: "ui://pipeline-results/app.html",
		}
		paramsJSON, _ := json.Marshal(params)

		result, err := server.handleReadResource(paramsJSON)
		if err != nil {
			t.Fatalf("Failed to read UI resource: %v", err)
		}

		if len(result.Contents) == 0 {
			t.Fatal("No contents returned")
		}

		htmlContent := result.Contents[0].Text

		if len(htmlContent) == 0 {
			t.Fatal("HTML content is empty")
		}

		// Check for CSP header
		if !strings.Contains(htmlContent, "Content-Security-Policy") {
			t.Error("HTML does not contain CSP header")
		}

		// Check for expected CSP directives
		expectedDirectives := []string{
			"default-src 'none'",
			"script-src 'unsafe-inline'",
			"style-src 'unsafe-inline'",
		}

		for _, directive := range expectedDirectives {
			if !strings.Contains(htmlContent, directive) {
				t.Errorf("HTML CSP missing directive: %s", directive)
			}
		}

		// Verify it's a single-file bundle (React code should be inline)
		if !strings.Contains(htmlContent, "React") && !strings.Contains(htmlContent, "react") {
			t.Error("HTML does not appear to contain inlined React code")
		}

		// Check for _meta.ui.permissions in the resource response
		if result.Contents[0].Meta == nil {
			t.Error("Resource should have _meta field")
		} else {
			if ui, ok := result.Contents[0].Meta["ui"].(map[string]interface{}); ok {
				if _, ok := ui["permissions"]; !ok {
					t.Error("Resource _meta.ui should have permissions field")
				}
			} else {
				t.Error("Resource _meta should have ui field")
			}
		}

		t.Logf("✓ UI HTML size: %d bytes", len(htmlContent))
		t.Logf("✓ CSP headers present")
		t.Logf("✓ React code inlined")
		t.Logf("✓ Resource metadata present")
	})

	t.Run("pipeline_results_tool_returns_json_content", func(t *testing.T) {
		// Find a results file
		resultsDir := ".veracode_pipeline"
		if _, err := os.Stat(resultsDir); os.IsNotExist(err) {
			t.Skip("No .veracode_pipeline directory found")
		}

		entries, err := os.ReadDir(resultsDir)
		if err != nil {
			t.Fatalf("Failed to read results directory: %v", err)
		}

		var resultsFile string
		for _, entry := range entries {
			if strings.HasPrefix(entry.Name(), "results-") && strings.HasSuffix(entry.Name(), ".json") {
				resultsFile = filepath.Join(resultsDir, entry.Name())
				break
			}
		}

		if resultsFile == "" {
			t.Skip("No pipeline results file found")
		}

		t.Logf("Using results file: %s", resultsFile)

		// Get absolute path to current directory
		absPath, err := filepath.Abs(".")
		if err != nil {
			t.Fatalf("Failed to get absolute path: %v", err)
		}

		// Call the tool
		callParams := CallToolParams{
			Name: "pipeline-results",
			Arguments: map[string]interface{}{
				"application_path": absPath,
				"results_file":     resultsFile,
			},
		}
		callParamsJSON, _ := json.Marshal(callParams)

		result, err := server.handleCallTool(callParamsJSON)
		if err != nil {
			t.Fatalf("Failed to call tool: %v", err)
		}

		if result.IsError {
			if len(result.Content) > 0 {
				t.Fatalf("Tool returned error: %s", result.Content[0].Text)
			}
			t.Fatal("Tool returned error (no error message)")
		}

		// Should have 1 content item (JSON data only, no header per VS Code requirement)
		if len(result.Content) != 1 {
			t.Fatalf("Expected 1 content item, got %d", len(result.Content))
		}

		// Item should be JSON text
		if result.Content[0].Type != "text" {
			t.Errorf("Expected content type 'text', got '%s'", result.Content[0].Type)
		}

		// Verify it's valid JSON
		var jsonData interface{}
		if err := json.Unmarshal([]byte(result.Content[0].Text), &jsonData); err != nil {
			t.Fatalf("Content item is not valid JSON: %v", err)
		}

		// Verify structuredContent is present (required for MCP Apps)
		if result.StructuredContent == nil {
			t.Error("Result should have structuredContent for MCP App UI")
		} else {
			// Verify it has the expected fields
			if app, ok := result.StructuredContent["application"].(map[string]interface{}); !ok || app == nil {
				t.Error("structuredContent should have 'application' field")
			}
			if summary, ok := result.StructuredContent["summary"].(map[string]interface{}); !ok || summary == nil {
				t.Error("structuredContent should have 'summary' field")
			}
			if findings, ok := result.StructuredContent["findings"].([]interface{}); !ok || findings == nil {
				t.Error("structuredContent should have 'findings' field")
			}
		}

		t.Logf("✓ Tool returns single JSON content item")
		t.Logf("✓ Content is valid JSON")
		t.Logf("✓ structuredContent present for MCP App")
	})
}
