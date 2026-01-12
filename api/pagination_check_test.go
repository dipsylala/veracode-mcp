package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/dipsylala/veracodemcp-go/credentials"
	"github.com/dipsylala/veracodemcp-go/hmac"
)

// TestCheckPaginationSupport tests if the Findings API supports undocumented pagination parameters
func TestCheckPaginationSupport(t *testing.T) {
	client, err := NewVeracodeClient()
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

	// Now let's try to manually add size parameter to the URL
	t.Log("\n===== TEST 2: Try to add size=1 to URL query manually =====")

	// Get credentials for manual request
	apiID, apiKey, baseURL, err := credentials.GetCredentials()
	if err != nil {
		t.Fatalf("Failed to get credentials: %v", err)
	}

	testURL := baseURL + "/appsec/v2/applications/" + appGUID + "/findings?scan_type=STATIC&size=1"
	t.Logf("Testing URL: %s", testURL)

	parsedURL, _ := url.Parse(testURL)
	authHeader, err := hmac.CalculateAuthorizationHeader(parsedURL, "GET", apiID, apiKey)
	if err != nil {
		t.Fatalf("Failed to generate auth header: %v", err)
	}

	req, _ := http.NewRequest("GET", testURL, nil)
	req.Header.Set("Authorization", authHeader)

	httpClient := &http.Client{}
	httpResp2, err := httpClient.Do(req)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	defer httpResp2.Body.Close()

	t.Logf("Response 2 (size=1 in URL):")
	t.Logf("  - HTTP Status: %d", httpResp2.StatusCode)

	body2, _ := io.ReadAll(httpResp2.Body)

	// Parse the response
	var resp2 struct {
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

	if err := json.Unmarshal(body2, &resp2); err != nil {
		t.Logf("  - Response body: %s", string(body2))
		t.Fatalf("Failed to decode response: %v", err)
	}

	if resp2.Page != nil {
		t.Logf("  - Page Number: %v", resp2.Page.Number)
		t.Logf("  - Page Size: %v", resp2.Page.Size)
		t.Logf("  - Total Elements: %v", resp2.Page.TotalElements)
		t.Logf("  - Total Pages: %v", resp2.Page.TotalPages)
	}
	if resp2.Embedded != nil {
		t.Logf("  - Findings Returned: %d", len(resp2.Embedded.Findings))
	}

	// Now test with page=1&size=10
	t.Log("\n===== TEST 3: Try page=1&size=10 =====")

	testURL3 := baseURL + "/appsec/v2/applications/" + appGUID + "/findings?scan_type=STATIC&page=1&size=10"
	t.Logf("Testing URL: %s", testURL3)

	parsedURL3, _ := url.Parse(testURL3)
	authHeader3, _ := hmac.CalculateAuthorizationHeader(parsedURL3, "GET", apiID, apiKey)

	req3, _ := http.NewRequest("GET", testURL3, nil)
	req3.Header.Set("Authorization", authHeader3)

	httpResp3, err := httpClient.Do(req3)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	defer httpResp3.Body.Close()

	t.Logf("Response 3 (page=1&size=10):")
	t.Logf("  - HTTP Status: %d", httpResp3.StatusCode)

	body3, _ := io.ReadAll(httpResp3.Body)

	var resp3 struct {
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

	if err := json.Unmarshal(body3, &resp3); err != nil {
		t.Logf("  - Response body: %s", string(body3))
		t.Fatalf("Failed to decode response: %v", err)
	}

	if resp3.Page != nil {
		t.Logf("  - Page Number: %v", resp3.Page.Number)
		t.Logf("  - Page Size: %v", resp3.Page.Size)
		t.Logf("  - Total Elements: %v", resp3.Page.TotalElements)
		t.Logf("  - Total Pages: %v", resp3.Page.TotalPages)
	}
	if resp3.Embedded != nil {
		t.Logf("  - Findings Returned: %d", len(resp3.Embedded.Findings))
	}

	t.Log("\n===== CONCLUSION =====")

	// Compare results
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
		t.Log(fmt.Sprintf("All requests returned %d findings regardless of page/size parameters.", findings1))
	}
}
