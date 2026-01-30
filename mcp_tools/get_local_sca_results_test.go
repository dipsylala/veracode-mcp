package mcp_tools

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestParseGetLocalSCAResultsRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parseGetLocalSCAResultsRequest(args)
	if err != nil {
		t.Fatalf("Failed to parse request: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
}

func TestParseGetLocalSCAResultsRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parseGetLocalSCAResultsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

func TestParseGetLocalSCAResultsRequest_EmptyApplicationPath(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "",
	}

	_, err := parseGetLocalSCAResultsRequest(args)
	if err == nil {
		t.Fatal("Expected error for empty application_path")
	}
}

func TestGetLocalSCAResultsTool_HandleMissingResultsFile(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory without results file
	tempDir := t.TempDir()

	result, err := handleGetLocalSCAResults(ctx, map[string]interface{}{
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

func TestGetLocalSCAResultsTool_HandleValidResultsFile(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory with results file
	tempDir := t.TempDir()
	scaDir := filepath.Join(tempDir, ".veracode_sca")
	if err := os.MkdirAll(scaDir, 0750); err != nil {
		t.Fatalf("Failed to create SCA directory: %v", err)
	}

	// Create a sample results file with Grype structure
	resultsFile := filepath.Join(scaDir, "veracode.json")
	sampleResults := SCAResults{
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

	result, err := handleGetLocalSCAResults(ctx, map[string]interface{}{
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

	if len(content) < 2 {
		t.Error("Expected at least header and JSON in content")
	}
}

func TestFormatSCAResultsResponse(t *testing.T) {
	results := &SCAResults{
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

	response := formatSCAResultsResponse("/test/path", "/test/path/.veracode_sca/veracode.json", results)

	// Verify response structure
	if response["content"] == nil {
		t.Fatal("Expected content in response")
	}

	content, ok := response["content"].([]map[string]interface{})
	if !ok {
		t.Fatal("Expected content to be array of maps")
	}

	if len(content) != 2 {
		t.Errorf("Expected 2 content items (header and JSON), got %d", len(content))
	}

	// Verify header contains expected information
	header, ok := content[0]["text"].(string)
	if !ok {
		t.Fatal("Expected header to be string")
	}

	if !contains(header, "Total Vulnerability Matches: 2") {
		t.Errorf("Header should contain total matches count, got: %s", header)
	}

	if !contains(header, "Critical: 1") {
		t.Errorf("Header should contain critical count, got: %s", header)
	}

	if !contains(header, "High: 1") {
		t.Errorf("Header should contain high count, got: %s", header)
	}

	// Verify JSON data can be parsed
	jsonData, ok := content[1]["text"].(string)
	if !ok {
		t.Fatal("Expected JSON data to be string")
	}

	var parsedData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &parsedData); err != nil {
		t.Fatalf("Failed to parse JSON data: %v", err)
	}

	// Verify parsed data structure
	if parsedData["summary"] == nil {
		t.Error("Expected summary in parsed data")
	}

	if parsedData["findings"] == nil {
		t.Error("Expected findings in parsed data")
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

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
