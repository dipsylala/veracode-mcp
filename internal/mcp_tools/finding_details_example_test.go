package mcp_tools

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dipsylala/veracode-mcp/api"
	dynamicflaw "github.com/dipsylala/veracode-mcp/api/rest/generated/dynamic_flaw"
)

// This test demonstrates what the LLM-optimized output looks like for a dynamic flaw
func TestDynamicFlawOutputFormat(t *testing.T) {
	// Create a sample dynamic flaw response (simulating API data)
	urlStr := "https://verademo.mwatson.vuln.sa.veracode.io:443/"
	rawReq := "R0VUIC8gSFRUUC8xLjEKSG9zdDogdmVyYWRlbW8ubXdhdHNvbi52dWxuLnNhLnZlcmFjb2RlLmlvOjQ0Mw==" // Base64: "GET / HTTP/1.1\nHost: verademo.mwatson.vuln.sa.veracode.io:443"
	rawResp := "SFRUUC8xLjEgMjAwIE9LCkNvbnRlbnQtVHlwZTogdGV4dC9odG1s"                                // Base64: "HTTP/1.1 200 OK\nContent-Type: text/html"

	dynamicFlaw := &dynamicflaw.DynamicFlaw{
		DynamicFlawInfo: &dynamicflaw.DynamicSpecificFlawInfo{
			Request: &dynamicflaw.Request{
				Url:      &urlStr,
				RawBytes: &rawReq,
			},
			Response: &dynamicflaw.Response{
				RawBytes: &rawResp,
			},
		},
	}

	// Create sample mitigation info
	mitigationInfo := &api.MitigationIssue{
		FlawID:   221,
		Category: "Server Configuration - Weak TLS Cipher",
		MitigationActions: []api.MitigationAction{
			{
				Action:   "appdesign",
				Desc:     "Mitigate by Design",
				Reviewer: "Security Team",
				Date:     "2025-08-07 10:30:00",
				Comment:  "This is acceptable for internal testing environments only",
			},
		},
	}

	// Build the response using our LLM-optimized format
	result := buildLLMOptimizedResponse(
		"e:\\Github\\verademo",
		"Verademo",
		221,
		"DYNAMIC",
		dynamicFlaw,
		mitigationInfo,
	)

	// Pretty print the result
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal result: %v", err)
	}

	fmt.Println("\n=== DYNAMIC FINDING OUTPUT (LLM-Optimized Format) ===")
	fmt.Println(string(jsonBytes))
	fmt.Println("=== END ===")

	// Verify the structure
	if result["flaw_id"] != 221 {
		t.Errorf("Expected flaw_id 221, got %v", result["flaw_id"])
	}

	if result["scan_type"] != "DYNAMIC" {
		t.Errorf("Expected scan_type DYNAMIC, got %v", result["scan_type"])
	}

	location, ok := result["location"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected location to be a map")
	}

	if location["url"] != urlStr {
		t.Errorf("Expected URL %s, got %v", urlStr, location["url"])
	}

	httpDetails, ok := result["http_details"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected http_details to be present")
	}

	if httpDetails["http_request"] == nil {
		t.Error("Expected http_request to be present")
	}

	if httpDetails["http_response"] == nil {
		t.Error("Expected http_response to be present")
	}

	mitigation, ok := result["mitigation"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected mitigation to be present")
	}

	if mitigation["status"] != "appdesign" {
		t.Errorf("Expected mitigation status 'appdesign', got %v", mitigation["status"])
	}
}
