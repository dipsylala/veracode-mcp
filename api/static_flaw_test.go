package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/dipsylala/veracodemcp-go/api"
)

// TestGetStaticFlawDetails_Integration tests retrieving static flaw details from the API
func TestGetStaticFlawDetails_Integration(t *testing.T) {
	// Skip if no credentials
	if os.Getenv("VERACODE_API_KEY_ID") == "" || os.Getenv("VERACODE_API_KEY_SECRET") == "" {
		t.Skip("Skipping integration test - no Veracode credentials")
	}

	client, err := api.NewVeracodeClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	authCtx := client.GetAuthContext(ctx)

	// MCPVerademo app GUID
	appGUID := "f4e74197-1e26-42c4-ab4b-245870c93280"
	
	// Known static flaw ID from MCPVerademo
	flawID := "6"

	// Get static flaw details
	req := client.StaticFindingDataPathClient().StaticFlawDataPathsInformationAPI.
		AppsecV2ApplicationsAppGuidFindingsIssueIdStaticFlawInfoGet(authCtx, appGUID, flawID)

	staticFlaw, resp, err := req.Execute()
	
	if err != nil {
		t.Fatalf("Failed to get static flaw details: %v", err)
	}

	if resp == nil || resp.StatusCode != 200 {
		t.Fatalf("Expected 200 status code, got: %v", resp)
	}

	if staticFlaw == nil {
		t.Fatal("Static flaw response is nil")
	}

	// Verify the response structure
	if staticFlaw.StaticFlaws == nil {
		t.Fatal("StaticFlaws is nil")
	}

	if len(staticFlaw.StaticFlaws) == 0 {
		t.Fatal("StaticFlaws array is empty")
	}

	flaw := staticFlaw.StaticFlaws[0]
	
	// Verify some basic fields exist
	if flaw.IssueSummary == nil {
		t.Fatal("IssueSummary is nil")
	}

	if flaw.IssueSummary.Name == nil || *flaw.IssueSummary.Name == "" {
		t.Fatal("IssueSummary.Name is empty")
	}

	t.Logf("✓ Successfully retrieved static flaw details for ID %s", flawID)
	t.Logf("  Application: %s", *flaw.IssueSummary.Name)
	if flaw.IssueSummary.Description != nil {
		t.Logf("  Description: %.100s...", *flaw.IssueSummary.Description)
	}
}
