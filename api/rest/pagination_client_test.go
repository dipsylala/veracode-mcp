package rest

import (
	"context"
	"testing"
)

// TestPaginationWithGeneratedClient tests pagination using the updated generated client
func TestPaginationWithGeneratedClient(t *testing.T) {
	client, err := NewVeracodeClient()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	if !client.IsConfigured() {
		t.Skip("Skipping test: API credentials not configured")
	}

	ctx := context.Background()
	authCtx := client.GetAuthContext(ctx)

	// MCPVerademo app GUID (known to have 204 static findings)
	appGUID := "f4e74197-1e26-42c4-ab4b-245870c93280"

	t.Log("===== TEST 1: Using Page() and Size() methods =====")

	// Test with page=0, size=5
	resp1, httpResp1, err := client.findingsClient.ApplicationFindingsInformationAPI.
		GetFindingsUsingGET(authCtx, appGUID).
		ScanType([]string{"STATIC"}).
		Page(0).
		Size(5).
		Execute()

	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}
	defer httpResp1.Body.Close()

	t.Logf("Page 0, Size 5:")
	t.Logf("  - HTTP Status: %d", httpResp1.StatusCode)
	if resp1.Page != nil {
		t.Logf("  - Page Number: %d", *resp1.Page.Number)
		t.Logf("  - Page Size: %d", *resp1.Page.Size)
		t.Logf("  - Total Elements: %d", *resp1.Page.TotalElements)
		t.Logf("  - Total Pages: %d", *resp1.Page.TotalPages)
	}
	if resp1.Embedded != nil {
		t.Logf("  - Findings Returned: %d", len(resp1.Embedded.Findings))

		if len(resp1.Embedded.Findings) != 5 {
			t.Errorf("Expected 5 findings, got %d", len(resp1.Embedded.Findings))
		}
	}

	t.Log("\n===== TEST 2: Page 1, Size 10 =====")

	resp2, httpResp2, err := client.findingsClient.ApplicationFindingsInformationAPI.
		GetFindingsUsingGET(authCtx, appGUID).
		ScanType([]string{"STATIC"}).
		Page(1).
		Size(10).
		Execute()

	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}
	defer httpResp2.Body.Close()

	t.Logf("Page 1, Size 10:")
	t.Logf("  - HTTP Status: %d", httpResp2.StatusCode)
	if resp2.Page != nil {
		t.Logf("  - Page Number: %d", *resp2.Page.Number)
		t.Logf("  - Page Size: %d", *resp2.Page.Size)
		t.Logf("  - Total Elements: %d", *resp2.Page.TotalElements)
		t.Logf("  - Total Pages: %d", *resp2.Page.TotalPages)
	}
	if resp2.Embedded != nil {
		t.Logf("  - Findings Returned: %d", len(resp2.Embedded.Findings))

		if len(resp2.Embedded.Findings) != 10 {
			t.Errorf("Expected 10 findings, got %d", len(resp2.Embedded.Findings))
		}
	}

	// Verify pagination metadata
	if resp1.Page != nil && resp2.Page != nil {
		if *resp1.Page.TotalElements != *resp2.Page.TotalElements {
			t.Errorf("Total elements should be same across pages: %d vs %d",
				*resp1.Page.TotalElements, *resp2.Page.TotalElements)
		}

		t.Logf("\n✅ Pagination working correctly!")
		t.Logf("  - Total findings: %d", *resp1.Page.TotalElements)
		t.Logf("  - Total pages (size=5): %d", *resp1.Page.TotalPages)
		t.Logf("  - Total pages (size=10): %d", *resp2.Page.TotalPages)
	}
}
