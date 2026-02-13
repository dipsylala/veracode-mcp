package rest

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/dipsylala/veracodemcp-go/credentials"
	"github.com/dipsylala/veracodemcp-go/hmac"
)

// TestScaFindings_DebugRawResponse captures the raw API response for SCA findings
func TestScaFindings_DebugRawResponse(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping debug test in short mode")
	}

	apiID, apiSecret, baseURL, err := credentials.GetCredentials()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	// Make authenticated API request for SCA findings
	bodyBytes := makeAuthenticatedRequestForSca(t, apiID, apiSecret, baseURL)

	// Save and log responses
	saveAndLogResponses(t, bodyBytes, "sca")

	// Parse and analyze findings structure
	analyzeEmbeddedFindings(t, bodyBytes)
}

func makeAuthenticatedRequestForSca(t *testing.T, apiID, apiSecret, baseURL string) []byte {
	// Use MCPVerademo which has SCA findings
	appID := "f4e74197-1e26-42c4-ab4b-245870c93280"

	// Construct the URL for SCA findings
	url := baseURL + "/appsec/v2/applications/" + appID + "/findings?scan_type=SCA&size=5"

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
	rawOutputFile := "findings_api_response_raw_sca.json"
	if err := os.WriteFile(rawOutputFile, bodyBytes, 0644); err != nil {
		t.Logf("Warning: Failed to write raw response to file: %v", err)
	} else {
		t.Logf("Raw SCA response saved to: %s", rawOutputFile)
	}

	return bodyBytes
}
