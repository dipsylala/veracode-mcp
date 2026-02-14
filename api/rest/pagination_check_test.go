package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/dipsylala/veracode-mcp/credentials"
	"github.com/dipsylala/veracode-mcp/hmac"
)

// paginationTestResponse represents the API response structure for pagination tests
type paginationTestResponse struct {
	Page *struct {
		Number        *int64 `json:"number"`
		Size          *int64 `json:"size"`
		TotalElements *int64 `json:"total_elements"`
		TotalPages    *int64 `json:"total_pages"`
	} `json:"page"`
	Embedded *struct {
		Findings []interface{} `json:"findings"`
	} `json:"_embedded"`
}

// fetchFindingsWithParams fetches findings with the given URL parameters
func fetchFindingsWithParams(ctx context.Context, testURL, apiID, apiKey string) (*paginationTestResponse, int, error) {
	parsedURL, _ := url.Parse(testURL)
	authHeader, err := hmac.CalculateAuthorizationHeader(parsedURL, "GET", apiID, apiKey)
	if err != nil {
		return nil, 0, err
	}

	req, _ := http.NewRequestWithContext(ctx, "GET", testURL, nil)
	req.Header.Set("Authorization", authHeader)

	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer httpResp.Body.Close()

	body, _ := io.ReadAll(httpResp.Body)

	var resp paginationTestResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, httpResp.StatusCode, err
	}

	return &resp, httpResp.StatusCode, nil
}

// logPaginationResponse logs the pagination response details
func logPaginationResponse(t *testing.T, label string, resp *paginationTestResponse, statusCode int) {
	t.Logf("Response %s:", label)
	t.Logf("  - HTTP Status: %d", statusCode)
	if resp.Page != nil {
		t.Logf("  - Page Number: %v", resp.Page.Number)
		t.Logf("  - Page Size: %v", resp.Page.Size)
		t.Logf("  - Total Elements: %v", resp.Page.TotalElements)
		t.Logf("  - Total Pages: %v", resp.Page.TotalPages)
	}
	if resp.Embedded != nil {
		t.Logf("  - Findings Returned: %d", len(resp.Embedded.Findings))
	}
}

// comparePaginationResults compares findings counts and logs the results
func comparePaginationResults(t *testing.T, findings1, findings2, findings3 int) {
	t.Logf("Findings count comparison:")
	t.Logf("  - Default (no params): %d", findings1)
	t.Logf("  - With size=1: %d", findings2)
	t.Logf("  - With page=1&size=10: %d", findings3)

	if findings2 == 1 && findings2 < findings1 {
		t.Logf("✅ PAGINATION WORKS! size=1 returned only 1 finding vs %d without size param", findings1)
	} else if findings2 == findings1 {
		t.Logf("❌ PAGINATION NOT SUPPORTED: size=1 returned same %d findings as default", findings1)
	} else {
		t.Logf("⚠️  UNEXPECTED: size=1 returned %d findings (expected 1 or %d)", findings2, findings1)
	}

	if findings3 == 10 && findings3 < findings1 {
		t.Logf("✅ PAGINATION WORKS! page=1&size=10 returned 10 findings")
	} else if findings3 == findings1 {
		t.Logf("❌ PAGINATION NOT SUPPORTED: page=1&size=10 returned all %d findings", findings1)
	} else {
		t.Logf("⚠️  UNEXPECTED: page=1&size=10 returned %d findings (expected 10 or %d)", findings3, findings1)
	}

	// Log a summary
	if findings2 < findings1 || findings3 < findings1 {
		t.Log("\n🎉 RESULT: The Findings API DOES support pagination parameters!")
		t.Log("This means the OpenAPI spec is incomplete and we should add Page/Size support to the client.")
	} else {
		t.Log("\n📝 RESULT: The Findings API does NOT support pagination parameters.")
		t.Logf("All requests returned %d findings regardless of page/size parameters.", findings1)
	}
}

// TestCheckPaginationSupport tests if the Findings API supports undocumented pagination parameters
func TestCheckPaginationSupport(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	if !client.IsConfigured() {
		t.Skip("Skipping test: API credentials not configured")
	}

	ctx := context.Background()

	// MCPVerademo app GUID (known to have 204 static findings)
	appGUID := "f4e74197-1e26-42c4-ab4b-245870c93280"

	t.Log("===== TEST 1: Call API without any custom parameters =====")

	authCtx := client.GetAuthContext(ctx)

	// First call: No size parameter (default behavior)
	apiReq1 := client.findingsClient.ApplicationFindingsInformationAPI.GetFindingsUsingGET(authCtx, appGUID).
		ScanType([]string{"STATIC"})

	resp1, httpResp1, err1 := apiReq1.Execute()
	if err1 != nil {
		t.Fatalf("API call 1 failed: %v", err1)
	}
	defer httpResp1.Body.Close()

	t.Logf("Response 1 (no size param):")
	t.Logf("  - HTTP Status: %d", httpResp1.StatusCode)
	if resp1.Page != nil {
		t.Logf("  - Page Number: %v", resp1.Page.Number)
		t.Logf("  - Page Size: %v", resp1.Page.Size)
		t.Logf("  - Total Elements: %v", resp1.Page.TotalElements)
		t.Logf("  - Total Pages: %v", resp1.Page.TotalPages)
	}
	if resp1.Embedded != nil {
		t.Logf("  - Findings Returned: %d", len(resp1.Embedded.Findings))
	}

	// Get credentials for manual requests
	apiID, apiKey, baseURL, err := credentials.GetCredentials()
	if err != nil {
		t.Fatalf("Failed to get credentials: %v", err)
	}

	// Test 2: size=1
	t.Log("\n===== TEST 2: Try to add size=1 to URL query manually =====")
	testURL2 := baseURL + "/appsec/v2/applications/" + appGUID + "/findings?scan_type=STATIC&size=1"
	t.Logf("Testing URL: %s", testURL2)

	resp2, statusCode2, err := fetchFindingsWithParams(ctx, testURL2, apiID, apiKey)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	logPaginationResponse(t, "2 (size=1 in URL)", resp2, statusCode2)

	// Test 3: page=1&size=10
	t.Log("\n===== TEST 3: Try page=1&size=10 =====")
	testURL3 := baseURL + "/appsec/v2/applications/" + appGUID + "/findings?scan_type=STATIC&page=1&size=10"
	t.Logf("Testing URL: %s", testURL3)

	resp3, statusCode3, err := fetchFindingsWithParams(ctx, testURL3, apiID, apiKey)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	logPaginationResponse(t, "3 (page=1&size=10)", resp3, statusCode3)

	// Compare results
	t.Log("\n===== CONCLUSION =====")

	var findings1, findings2, findings3 int
	if resp1.Embedded != nil {
		findings1 = len(resp1.Embedded.Findings)
	}
	if resp2.Embedded != nil {
		findings2 = len(resp2.Embedded.Findings)
	}
	if resp3.Embedded != nil {
		findings3 = len(resp3.Embedded.Findings)
	}

	comparePaginationResults(t, findings1, findings2, findings3)
}
