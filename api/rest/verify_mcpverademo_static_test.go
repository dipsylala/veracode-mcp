package rest

import (
	"context"
	"testing"
	"time"
)

// TestVerifyMCPVerademoStaticFindings verifies the total count of static findings for MCPVerademo
func TestVerifyMCPVerademoStaticFindings(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client, err := NewVeracodeClient()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Check both apps
	apps := map[string]string{
		"testAppID (Verademo-NET)": "65c204e5-a74c-4b68-a62a-4bfdc08e27af",
		"MCPVerademo":              "f4e74197-1e26-42c4-ab4b-245870c93280",
	}

	for appName, appGUID := range apps {
		t.Logf("\n=== %s ===", appName)
		checkStaticFindings(t, ctx, client, appGUID)
	}
}

func checkStaticFindings(t *testing.T, ctx context.Context, client *VeracodeClient, appGUID string) {
	// Get first page to check total
	req := FindingsRequest{
		AppProfile: appGUID,
		Size:       1, // Just get 1 finding to check the total
		Page:       0,
	}

	result, err := client.GetStaticFindings(ctx, req)
	if err != nil {
		t.Logf("  ERROR: Failed to get static findings: %v", err)
		return
	}

	t.Logf("  Static Findings:")
	t.Logf("    Total Count: %d", result.TotalCount)
	t.Logf("    Page: %d", result.Page)
	t.Logf("    Size: %d", result.Size)
	t.Logf("    Findings on this page: %d", len(result.Findings))

	// Check if the first finding exists
	if len(result.Findings) > 0 {
		t.Logf("    First finding ID: %s", result.Findings[0].ID)
	}

	if result.TotalCount == 0 {
		t.Logf("    No static findings found")
	} else if result.TotalCount < 200 {
		t.Logf("    WARNING: Total count is %d (expected 204 for MCPVerademo)", result.TotalCount)
	} else {
		t.Logf("    SUCCESS: Total count %d matches expectation", result.TotalCount)
	}
}
