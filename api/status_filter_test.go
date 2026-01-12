package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/dipsylala/veracodemcp-go/credentials"
	"github.com/dipsylala/veracodemcp-go/hmac"
)

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

	type testCase struct {
		name        string
		statusValue string
		expectedMsg string
	}

	testCases := []testCase{
		{
			name:        "Status OPEN",
			statusValue: "OPEN",
			expectedMsg: "open findings",
		},
		{
			name:        "Status NEW",
			statusValue: "NEW",
			expectedMsg: "new findings",
		},
		{
			name:        "Status CLOSED",
			statusValue: "CLOSED",
			expectedMsg: "closed findings",
		},
	}

	// First, get baseline without status filter
	t.Log("\n===== BASELINE: No status filter =====")
	baselineURL := baseURL + "/appsec/v2/applications/" + appGUID + "/findings?scan_type=STATIC&size=5"

	parsedURL, _ := url.Parse(baselineURL)
	authHeader, _ := hmac.CalculateAuthorizationHeader(parsedURL, "GET", apiID, apiKey)

	req, _ := http.NewRequest("GET", baselineURL, nil)
	req.Header.Set("Authorization", authHeader)

	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	defer httpResp.Body.Close()

	bodyBaseline, _ := io.ReadAll(httpResp.Body)

	var respBaseline struct {
		Page *struct {
			TotalElements *int64 `json:"total_elements"`
		} `json:"page"`
		Embedded *struct {
			Findings []map[string]interface{} `json:"findings"`
		} `json:"_embedded"`
	}

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
		statuses := make(map[string]int)
		for _, finding := range respBaseline.Embedded.Findings {
			if findingStatus, ok := finding["finding_status"].(map[string]interface{}); ok {
				if status, ok := findingStatus["status"].(string); ok {
					statuses[status]++
				}
			}
		}
		t.Logf("  - Status distribution in sample: %v", statuses)
	}

	// Now test each status value
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("\n===== Testing status=%s =====", tc.statusValue)

			testURL := baseURL + "/appsec/v2/applications/" + appGUID + "/findings?scan_type=STATIC&status=" + tc.statusValue + "&size=10"
			t.Logf("Testing URL: %s", testURL)

			parsedURL, _ := url.Parse(testURL)
			authHeader, _ := hmac.CalculateAuthorizationHeader(parsedURL, "GET", apiID, apiKey)

			req, _ := http.NewRequest("GET", testURL, nil)
			req.Header.Set("Authorization", authHeader)

			httpResp, err := httpClient.Do(req)
			if err != nil {
				t.Fatalf("HTTP request failed: %v", err)
			}
			defer httpResp.Body.Close()

			t.Logf("Response status: %d", httpResp.StatusCode)

			body, _ := io.ReadAll(httpResp.Body)

			// Check if we got an error response
			if httpResp.StatusCode != 200 {
				t.Logf("Response body: %s", string(body))
				t.Logf("❌ API rejected status parameter (HTTP %d)", httpResp.StatusCode)
				return
			}

			var resp struct {
				Page *struct {
					TotalElements *int64 `json:"total_elements"`
				} `json:"page"`
				Embedded *struct {
					Findings []map[string]interface{} `json:"findings"`
				} `json:"_embedded"`
			}

			if err := json.Unmarshal(body, &resp); err != nil {
				t.Logf("Response body: %s", string(body))
				t.Fatalf("Failed to decode response: %v", err)
			}

			totalElements := int64(0)
			if resp.Page != nil && resp.Page.TotalElements != nil {
				totalElements = *resp.Page.TotalElements
			}

			findingsCount := 0
			if resp.Embedded != nil {
				findingsCount = len(resp.Embedded.Findings)
			}

			t.Logf("Response with status=%s:", tc.statusValue)
			t.Logf("  - Total elements: %d", totalElements)
			t.Logf("  - Findings returned: %d", findingsCount)

			// Verify all returned findings have the requested status
			if resp.Embedded != nil && len(resp.Embedded.Findings) > 0 {
				allMatch := true
				statuses := make(map[string]int)

				for i, finding := range resp.Embedded.Findings {
					if findingStatus, ok := finding["finding_status"].(map[string]interface{}); ok {
						if status, ok := findingStatus["status"].(string); ok {
							statuses[status]++
							if status != tc.statusValue {
								t.Logf("  ⚠️  Finding %d has status '%s' (expected '%s')", i+1, status, tc.statusValue)
								allMatch = false
							}
						}
					}
				}

				t.Logf("  - Status distribution: %v", statuses)

				if allMatch && len(statuses) == 1 && statuses[tc.statusValue] == findingsCount {
					t.Logf("  ✅ All findings have status '%s'", tc.statusValue)
				} else if !allMatch {
					t.Logf("  ❌ Not all findings match status filter (API may not support this parameter)")
				}
			}

			// Compare to baseline
			if totalElements < baselineTotal {
				t.Logf("  ✅ FILTERING WORKS! Total reduced from %d to %d", baselineTotal, totalElements)
			} else if totalElements == baselineTotal {
				t.Logf("  ❌ FILTERING NOT SUPPORTED: Same total as baseline (%d)", baselineTotal)
			}
		})
	}

	// Summary
	t.Log("\n===== SUMMARY =====")
	t.Log("The test checked if the Findings API supports a 'status' parameter.")
	t.Log("If filtering worked, you should see different totals for OPEN/NEW/CLOSED.")
}
