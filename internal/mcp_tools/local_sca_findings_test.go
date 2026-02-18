package mcp_tools

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestParseGetLocalSCAFindingsRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parseGetLocalSCAFindingsRequest(args)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
}

func TestParseGetLocalSCAFindingsRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parseGetLocalSCAFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

func TestParseGetLocalSCAFindingsRequest_EmptyApplicationPath(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "",
	}

	_, err := parseGetLocalSCAFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for empty application_path")
	}
}

func TestGetLocalSCAFindingsTool_HandleMissingResultsFile(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory without results file
	tempDir := t.TempDir()

	result, err := handleGetLocalSCAFindings(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should return content with "No results found" message
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["content"] == nil {
		t.Error("Expected content in response")
	}
}

func TestGetLocalSCAFindingsTool_HandleValidResultsFile(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory with results file
	tempDir := t.TempDir()
	scaDir := filepath.Join(tempDir, ".veracode", "sca")
	if err := os.MkdirAll(scaDir, 0750); err != nil {
		t.Fatalf("Failed to create SCA directory: %v", err)
	}

	// Create a sample results file with Grype structure
	resultsFile := filepath.Join(scaDir, "veracode.json")
	sampleResults := SCAFindings{
		Vulnerabilities: SCAVulnerabilities{
			Matches: []SCAMatch{
				{
					Artifact: SCAArtifact{
						Name:    "test-package",
						Version: "1.0.0",
						Type:    "npm",
					},
					Vulnerability: SCAVulnerability{
						ID:          "CVE-2023-12345",
						Severity:    "Critical",
						Description: "A critical vulnerability in test-package",
					},
				},
			},
		},
	}

	resultsJSON, err := json.Marshal(sampleResults)
	if err != nil {
		t.Fatalf("Failed to marshal sample results: %v", err)
	}

	if err := os.WriteFile(resultsFile, resultsJSON, 0644); err != nil {
		t.Fatalf("Failed to write results file: %v", err)
	}

	result, err := handleGetLocalSCAFindings(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should return content with results
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["content"] == nil {
		t.Error("Expected content in response")
	}

	// Verify content structure
	content, ok := resultMap["content"].([]map[string]interface{})
	if !ok {
		t.Fatal("Expected content to be array of maps")
	}

	if len(content) < 1 {
		t.Error("Expected at least one content item (combined header and JSON)")
	}

	// Verify structuredContent is present
	if resultMap["structuredContent"] == nil {
		t.Error("Expected structuredContent in response")
	}
}

func TestFormatSCAFindingsResponse(t *testing.T) {
	results := &SCAFindings{
		Vulnerabilities: SCAVulnerabilities{
			Matches: []SCAMatch{
				{
					Artifact: SCAArtifact{
						Name:    "test-lib",
						Version: "1.0.0",
						Type:    "npm",
					},
					Vulnerability: SCAVulnerability{
						ID:          "CVE-2023-12345",
						Severity:    "Critical",
						Description: "Test vulnerability",
					},
				},
				{
					Artifact: SCAArtifact{
						Name:    "another-lib",
						Version: "2.0.0",
						Type:    "pip",
					},
					Vulnerability: SCAVulnerability{
						ID:          "CVE-2023-67890",
						Severity:    "High",
						Description: "SQL injection",
					},
				},
			},
		},
	}

	req := &GetLocalSCAFindingsRequest{
		ApplicationPath: "/test/path",
		Size:            50,
		Page:            0,
	}

	response := formatSCAFindingsResponse("/test/path", "/test/path/.veracode/sca/veracode.json", results, req)

	// Verify response structure
	if response["content"] == nil {
		t.Fatal("Expected content in response")
	}

	content, ok := response["content"].([]map[string]interface{})
	if !ok {
		t.Fatal("Expected content to be array of maps")
	}

	if len(content) != 1 {
		t.Errorf("Expected 1 content item (JSON only), got %d", len(content))
	}

	// Verify text is valid JSON
	jsonText, ok := content[0]["text"].(string)
	if !ok {
		t.Fatal("Expected text to be string")
	}

	// Parse the JSON to verify it's valid
	var parsedData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonText), &parsedData); err != nil {
		t.Fatalf("Failed to parse JSON text: %v", err)
	}

	// Verify parsed data has expected fields
	if parsedData["application"] == nil {
		t.Error("Expected application in JSON")
	}

	if parsedData["summary"] == nil {
		t.Error("Expected summary in JSON")
	}

	if parsedData["findings"] == nil {
		t.Error("Expected findings in JSON")
	}

	if parsedData["pagination"] == nil {
		t.Error("Expected pagination in JSON")
	}

	// Verify structuredContent is present
	if response["structuredContent"] == nil {
		t.Fatal("Expected structuredContent in response")
	}

	// Verify structuredContent structure
	structuredContent, ok := response["structuredContent"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected structuredContent to be a map")
	}

	if structuredContent["summary"] == nil {
		t.Error("Expected summary in structuredContent")
	}

	if structuredContent["findings"] == nil {
		t.Error("Expected findings in structuredContent")
	}
}

func TestTransformSCASeverity(t *testing.T) {
	tests := []struct {
		severity string
		expected string
	}{
		{"Critical", "critical"},
		{"High", "high"},
		{"Medium", "medium"},
		{"Low", "low"},
		{"Negligible", "negligible"},
		{"Unknown", "unknown"},
		{"", "unknown"},
		{"InvalidValue", "unknown"},
	}

	for _, test := range tests {
		result := transformSCASeverity(test.severity)
		if result != test.expected {
			t.Errorf("transformSCASeverity(%q) = %q, expected %q", test.severity, result, test.expected)
		}
	}
}
