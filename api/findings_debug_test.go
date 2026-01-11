package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/dipsylala/veracodemcp-go/credentials"
	"github.com/dipsylala/veracodemcp-go/hmac"
)

// TestStaticFindings_DebugRawResponse captures the raw API response to diagnose deserialization issues
func TestStaticFindings_DebugRawResponse(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping debug test in short mode")
	}

	apiID, apiSecret, baseURL, err := credentials.GetCredentials()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	// Construct the URL manually
	url := baseURL + "/appsec/v2/applications/" + testAppID + "/findings?scan_type=STATIC&size=1"

	// Create HTTP client
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	// Create request
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Calculate and add HMAC auth header
	reqURL := req.URL
	authHeader, err := hmac.CalculateAuthorizationHeader(reqURL, req.Method, apiID, apiSecret)
	if err != nil {
		t.Fatalf("Failed to calculate HMAC: %v", err)
	}
	req.Header.Set("Authorization", authHeader)

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	t.Logf("Response status: %d %s", resp.StatusCode, resp.Status)
	t.Logf("Response headers:")
	for k, v := range resp.Header {
		t.Logf("  %s: %v", k, v)
	}

	// Read the raw response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Write raw response to file
	rawOutputFile := "findings_api_response_raw.json"
	if err := os.WriteFile(rawOutputFile, bodyBytes, 0644); err != nil {
		t.Logf("Warning: Failed to write raw response to file: %v", err)
	} else {
		t.Logf("Raw response saved to: %s", rawOutputFile)
	}

	t.Logf("\n=== RAW RESPONSE BODY ===")
	t.Logf("%s", string(bodyBytes))
	t.Logf("=== END RAW RESPONSE ===\n")

	// Try to pretty-print the JSON
	var prettyJSON interface{}
	if err := json.Unmarshal(bodyBytes, &prettyJSON); err == nil {
		prettyBytes, _ := json.MarshalIndent(prettyJSON, "", "  ")

		// Write pretty-printed response to file
		prettyOutputFile := "findings_api_response_pretty.json"
		if err := os.WriteFile(prettyOutputFile, prettyBytes, 0644); err != nil {
			t.Logf("Warning: Failed to write pretty response to file: %v", err)
		} else {
			t.Logf("Pretty-printed response saved to: %s", prettyOutputFile)
		}

		t.Logf("\n=== PRETTY-PRINTED JSON ===")
		t.Logf("%s", string(prettyBytes))
		t.Logf("=== END PRETTY JSON ===\n")
	}

	// Parse to see the structure
	var response map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Look for embedded findings
	if embedded, ok := response["_embedded"].(map[string]interface{}); ok {
		if findings, ok := embedded["findings"].([]interface{}); ok {
			t.Logf("Found %d findings", len(findings))
			if len(findings) > 0 {
				firstFinding := findings[0].(map[string]interface{})
				t.Logf("\n=== FIRST FINDING ===")
				firstFindingBytes, _ := json.MarshalIndent(firstFinding, "", "  ")
				t.Logf("%s", string(firstFindingBytes))
				t.Logf("=== END FIRST FINDING ===\n")

				// Check for finding_details field
				if details, ok := firstFinding["finding_details"].(map[string]interface{}); ok {
					t.Logf("\n=== FINDING_DETAILS ===")
					detailsBytes, _ := json.MarshalIndent(details, "", "  ")
					t.Logf("%s", string(detailsBytes))
					t.Logf("=== END FINDING_DETAILS ===\n")

					// Log all field names in finding_details
					t.Logf("Fields in finding_details:")
					for k := range details {
						t.Logf("  - %s", k)
					}
				}
			}
		}
	}
}
