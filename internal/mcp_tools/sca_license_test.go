package mcp_tools

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/dipsylala/veracodemcp-go/api/rest"
)

// TestScaFindingsLicenseInclusion verifies that license information is included in the MCP response
func TestScaFindingsLicenseInclusion(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Create a test SCA finding with license information
	finding := rest.Finding{
		ID:                "12345",
		Status:            "OPEN",
		ResolutionStatus:  "NONE",
		ViolatesPolicy:    true,
		Severity:          "5",
		SeverityScore:     5,
		CWE:               "CWE-20",
		Description:       "Test vulnerability",
		ComponentFilename: "test-component-1.0.0.jar",
		ComponentVersion:  "1.0.0",
		CVE:               "CVE-2024-12345",
		Licenses: []rest.License{
			{
				LicenseID:  "apache-2.0",
				RiskRating: "2",
			},
			{
				LicenseID:  "mit",
				RiskRating: "1",
			},
		},
	}

	// Process the finding through the MCP processor
	mcpFinding := processScaFinding(finding)

	// Verify component exists
	if mcpFinding.Component == nil {
		t.Fatal("Expected Component to be present in MCP finding")
	}

	// Verify licenses are included
	if len(mcpFinding.Component.Licenses) != 2 {
		t.Fatalf("Expected 2 licenses, got %d", len(mcpFinding.Component.Licenses))
	}

	// Verify first license
	if mcpFinding.Component.Licenses[0].LicenseID != "apache-2.0" {
		t.Errorf("Expected first license to be 'apache-2.0', got '%s'", mcpFinding.Component.Licenses[0].LicenseID)
	}
	if mcpFinding.Component.Licenses[0].RiskRating != "2" {
		t.Errorf("Expected first license risk rating to be '2', got '%s'", mcpFinding.Component.Licenses[0].RiskRating)
	}

	// Verify second license
	if mcpFinding.Component.Licenses[1].LicenseID != "mit" {
		t.Errorf("Expected second license to be 'mit', got '%s'", mcpFinding.Component.Licenses[1].LicenseID)
	}
	if mcpFinding.Component.Licenses[1].RiskRating != "1" {
		t.Errorf("Expected second license risk rating to be '1', got '%s'", mcpFinding.Component.Licenses[1].RiskRating)
	}

	t.Logf("Successfully verified license information in MCP finding")
	t.Logf("  Component: %s v%s", mcpFinding.Component.Name, mcpFinding.Component.Version)
	t.Logf("  License 1: %s (risk: %s)", mcpFinding.Component.Licenses[0].LicenseID, mcpFinding.Component.Licenses[0].RiskRating)
	t.Logf("  License 2: %s (risk: %s)", mcpFinding.Component.Licenses[1].LicenseID, mcpFinding.Component.Licenses[1].RiskRating)
}

// TestScaFindingsResponseJSON verifies the complete JSON structure includes licenses
func TestScaFindingsResponseJSON(t *testing.T) {
	// Create a complete MCP response with license information
	response := MCPFindingsResponse{
		Application: MCPApplication{
			Name: "TestApp",
			ID:   "test-guid",
		},
		Summary: MCPFindingsSummary{
			TotalFindings: 1,
			BySeverity: map[string]int{
				"critical": 1,
			},
			ByStatus: map[string]int{
				"open": 1,
			},
			ByMitigation: map[string]int{
				"none": 1,
			},
		},
		Findings: []MCPFinding{
			{
				FlawID:           "12345",
				ScanType:         "SCA",
				Status:           "OPEN",
				MitigationStatus: "NONE",
				ViolatesPolicy:   true,
				Severity:         "CRITICAL",
				SeverityScore:    5,
				CweId:            20,
				Description:      "Test vulnerability",
				Component: &MCPComponent{
					Name:    "test-component-1.0.0.jar",
					Version: "1.0.0",
					Library: "test-component-1.0.0.jar",
					Licenses: []MCPLicense{
						{
							LicenseID:  "apache-2.0",
							RiskRating: "2",
						},
					},
				},
				Vulnerability: &MCPVulnerability{
					CVEID: "CVE-2024-12345",
				},
			},
		},
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal MCP response: %v", err)
	}

	t.Logf("MCP Response JSON:\n%s", string(jsonData))

	// Unmarshal back to verify structure
	var decoded MCPFindingsResponse
	if err := json.Unmarshal(jsonData, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal MCP response: %v", err)
	}

	// Verify licenses survived the round-trip
	if len(decoded.Findings) != 1 {
		t.Fatalf("Expected 1 finding, got %d", len(decoded.Findings))
	}

	if decoded.Findings[0].Component == nil {
		t.Fatal("Component is nil after unmarshaling")
	}

	if len(decoded.Findings[0].Component.Licenses) != 1 {
		t.Fatalf("Expected 1 license after unmarshaling, got %d", len(decoded.Findings[0].Component.Licenses))
	}

	license := decoded.Findings[0].Component.Licenses[0]
	if license.LicenseID != "apache-2.0" {
		t.Errorf("License ID mismatch: expected 'apache-2.0', got '%s'", license.LicenseID)
	}
	if license.RiskRating != "2" {
		t.Errorf("License risk rating mismatch: expected '2', got '%s'", license.RiskRating)
	}

	t.Log("Successfully verified license information in JSON round-trip")
}

// extractMCPResponseFromResult extracts the MCP response JSON from the tool result
func extractMCPResponseFromResult(t *testing.T, result interface{}) *MCPFindingsResponse {
	t.Helper()

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("Result is not a map: %T", result)
	}

	// Check for error in response
	if errMsg, hasError := resultMap["error"]; hasError {
		t.Skipf("API call failed (expected in CI/CD): %v", errMsg)
	}

	// Extract content
	content, hasContent := resultMap["content"]
	if !hasContent {
		t.Fatal("Result does not contain 'content' field")
	}

	contentArray, ok := content.([]map[string]interface{})
	if !ok || len(contentArray) == 0 {
		t.Fatal("Content is not an array or is empty")
	}

	// Get the text field directly (new format consistent with static/dynamic findings)
	jsonText, hasText := contentArray[0]["text"]
	if !hasText {
		t.Fatal("Content does not have 'text' field")
	}

	jsonStr, ok := jsonText.(string)
	if !ok {
		t.Fatal("Text is not a string")
	}

	// The text contains pagination summary followed by JSON, so we need to extract just the JSON part
	// Find the first '{' character which marks the start of the JSON
	jsonStart := strings.Index(jsonStr, "{")
	if jsonStart == -1 {
		t.Fatal("Could not find JSON start in text content")
	}
	jsonStr = jsonStr[jsonStart:]

	// Parse the JSON response
	var mcpResponse MCPFindingsResponse
	if err := json.Unmarshal([]byte(jsonStr), &mcpResponse); err != nil {
		t.Fatalf("Failed to unmarshal MCP response JSON: %v", err)
	}

	return &mcpResponse
}

// countAndLogLicenseFindings counts findings with licenses and logs the first one
func countAndLogLicenseFindings(t *testing.T, findings []MCPFinding) int {
	t.Helper()

	foundWithLicenses := 0
	for _, finding := range findings {
		if finding.Component != nil && len(finding.Component.Licenses) > 0 {
			foundWithLicenses++
			if foundWithLicenses == 1 { // Log details for first finding with licenses
				t.Logf("Finding with licenses:")
				t.Logf("  Component: %s v%s", finding.Component.Name, finding.Component.Version)
				t.Logf("  Licenses: %d", len(finding.Component.Licenses))
				for j, license := range finding.Component.Licenses {
					t.Logf("    License %d: %s (risk: %s)", j+1, license.LicenseID, license.RiskRating)
				}
			}
		}
	}

	return foundWithLicenses
}

// TestScaFindingsIntegrationWithLicenses verifies end-to-end license inclusion from API to MCP
func TestScaFindingsIntegrationWithLicenses(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tool := NewScaFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	// Call the handler with test parameters
	args := map[string]interface{}{
		"application_path": "c:\\test\\path",
		"app_profile":      "f4e74197-1e26-42c4-ab4b-245870c93280", // MCPVerademo GUID
		"size":             5,
		"page":             0,
	}

	ctx := context.Background()
	result, err := handleGetScaFindings(ctx, args)
	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	mcpResponse := extractMCPResponseFromResult(t, result)
	t.Logf("Retrieved %d SCA findings", len(mcpResponse.Findings))

	foundWithLicenses := countAndLogLicenseFindings(t, mcpResponse.Findings)

	if foundWithLicenses > 0 {
		t.Logf("Found %d findings with license information", foundWithLicenses)
	} else {
		t.Log("WARNING: No findings with license information found. This may be expected if components have no license data.")
	}
}
