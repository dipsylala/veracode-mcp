package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/dipsylala/veracodemcp-go/credentials"
	"github.com/dipsylala/veracodemcp-go/hmac"
)

// findingsAPIResponse represents the structure of the Findings API response
type findingsAPIResponse struct {
	Page *struct {
		TotalElements *int64 `json:"total_elements"`
	} `json:"page"`
	Embedded *struct {
		Findings []map[string]interface{} `json:"findings"`
	} `json:"_embedded"`
}

// fetchFindingsWithStatus fetches findings for a given status filter
func fetchFindingsWithStatus(baseURL, appGUID, statusValue, apiID, apiKey string) (*findingsAPIResponse, int, error) {
	testURL := baseURL + "/appsec/v2/applications/" + appGUID + "/findings?scan_type=STATIC&status=" + statusValue + "&size=10"

	parsedURL, _ := url.Parse(testURL)
	authHeader, _ := hmac.CalculateAuthorizationHeader(parsedURL, "GET", apiID, apiKey)

	req, _ := http.NewRequestWithContext(context.Background(), "GET", testURL, nil)
	req.Header.Set("Authorization", authHeader)

	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer httpResp.Body.Close()

	body, _ := io.ReadAll(httpResp.Body)

	if httpResp.StatusCode != 200 {
		return nil, httpResp.StatusCode, nil
	}

	var resp findingsAPIResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, httpResp.StatusCode, err
	}

	return &resp, httpResp.StatusCode, nil
}

// getStatusDistribution extracts status distribution from findings
func getStatusDistribution(findings []map[string]interface{}) map[string]int {
	statuses := make(map[string]int)
	for _, finding := range findings {
		if findingStatus, ok := finding["finding_status"].(map[string]interface{}); ok {
			if status, ok := findingStatus["status"].(string); ok {
				statuses[status]++
			}
		}
	}
	return statuses
}

// testStatusFilter tests a single status filter value
func testStatusFilter(t *testing.T, baseURL, appGUID, statusValue, apiID, apiKey string, baselineTotal int64) {
	t.Logf("\n===== Testing status=%s =====", statusValue)

	resp, statusCode, err := fetchFindingsWithStatus(baseURL, appGUID, statusValue, apiID, apiKey)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}

	if statusCode != 200 {
		t.Logf("❌ API rejected status parameter (HTTP %d)", statusCode)
		return
	}

	totalElements := int64(0)
	if resp.Page != nil && resp.Page.TotalElements != nil {
		totalElements = *resp.Page.TotalElements
	}

	findingsCount := 0
	if resp.Embedded != nil {
		findingsCount = len(resp.Embedded.Findings)
	}

	t.Logf("Response with status=%s:", statusValue)
	t.Logf("  - Total elements: %d", totalElements)
	t.Logf("  - Findings returned: %d", findingsCount)

	// Verify all returned findings have the requested status
	if resp.Embedded != nil && len(resp.Embedded.Findings) > 0 {
		statuses := getStatusDistribution(resp.Embedded.Findings)
		t.Logf("  - Status distribution: %v", statuses)

		if len(statuses) == 1 && statuses[statusValue] == findingsCount {
			t.Logf("  ✅ All findings have status '%s'", statusValue)
		} else {
			t.Logf("  ❌ Not all findings match status filter (API may not support this parameter)")
		}
	}

	// Compare to baseline
	if totalElements < baselineTotal {
		t.Logf("  ✅ FILTERING WORKS! Total reduced from %d to %d", baselineTotal, totalElements)
	} else if totalElements == baselineTotal {
		t.Logf("  ❌ FILTERING NOT SUPPORTED: Same total as baseline (%d)", baselineTotal)
	}
}

// TestStatusFilterSupport tests if the Findings API supports undocumented status parameter
func TestStatusFilterSupport(t *testing.T) {
	client, err := NewVeracodeClient()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	if !client.IsConfigured() {
		t.Skip("Skipping test: API credentials not configured")
	}

	// MCPVerademo app GUID (known to have mix of open/closed findings)
	appGUID := "f4e74197-1e26-42c4-ab4b-245870c93280"

	// Get credentials for manual request
	apiID, apiKey, baseURL, err := credentials.GetCredentials()
	if err != nil {
		t.Fatalf("Failed to get credentials: %v", err)
	}

	// First, get baseline without status filter
	t.Log("\n===== BASELINE: No status filter =====")
	baselineURL := baseURL + "/appsec/v2/applications/" + appGUID + "/findings?scan_type=STATIC&size=5"

	parsedURL, _ := url.Parse(baselineURL)
	authHeader, _ := hmac.CalculateAuthorizationHeader(parsedURL, "GET", apiID, apiKey)

	req, _ := http.NewRequestWithContext(context.Background(), "GET", baselineURL, nil)
	req.Header.Set("Authorization", authHeader)

	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	defer httpResp.Body.Close()

	bodyBaseline, _ := io.ReadAll(httpResp.Body)

	var respBaseline findingsAPIResponse
	if err := json.Unmarshal(bodyBaseline, &respBaseline); err != nil {
		t.Fatalf("Failed to decode baseline response: %v", err)
	}

	baselineTotal := int64(0)
	if respBaseline.Page != nil && respBaseline.Page.TotalElements != nil {
		baselineTotal = *respBaseline.Page.TotalElements
	}
	baselineCount := 0
	if respBaseline.Embedded != nil {
		baselineCount = len(respBaseline.Embedded.Findings)
	}

	t.Logf("Baseline (no status filter):")
	t.Logf("  - Total elements: %d", baselineTotal)
	t.Logf("  - Findings returned: %d", baselineCount)

	// Check statuses in baseline
	if respBaseline.Embedded != nil && len(respBaseline.Embedded.Findings) > 0 {
		statuses := getStatusDistribution(respBaseline.Embedded.Findings)
		t.Logf("  - Status distribution in sample: %v", statuses)
	}

	// Test each status value
	statusValues := []string{"OPEN", "NEW", "CLOSED"}
	for _, status := range statusValues {
		t.Run("Status_"+status, func(t *testing.T) {
			testStatusFilter(t, baseURL, appGUID, status, apiID, apiKey, baselineTotal)
		})
	}

	// Summary
	t.Log("\n===== SUMMARY =====")
	t.Log("The test checked if the Findings API supports a 'status' parameter.")
	t.Log("If filtering worked, you should see different totals for OPEN/NEW/CLOSED.")
}
