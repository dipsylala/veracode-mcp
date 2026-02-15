package xml

import (
	"context"
	"testing"
)

func TestGetMitigationInfo(t *testing.T) {
	// Create client
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping test: credentials not available: %v", err)
	}

	if !client.IsConfigured() {
		t.Skip("Skipping test: API credentials not configured")
	}

	// Test with provided example: build_id=53691725, flaw_id=50
	ctx := context.Background()
	buildID := int64(53691725)
	flawIDs := []int64{50}

	mitigationInfo, err := client.GetMitigationInfo(ctx, buildID, flawIDs)
	if err != nil {
		t.Fatalf("Failed to get mitigation info: %v", err)
	}

	t.Logf("MitigationInfo Version: %s", mitigationInfo.Version)
	t.Logf("Build ID: %d", mitigationInfo.BuildID)
	t.Logf("Number of issues: %d", len(mitigationInfo.Issues))
	t.Logf("Number of errors: %d", len(mitigationInfo.Errors))

	// Verify build ID matches
	if mitigationInfo.BuildID != buildID {
		t.Errorf("Expected BuildID %d, got %d", buildID, mitigationInfo.BuildID)
	}

	// Log issues
	for i, issue := range mitigationInfo.Issues {
		t.Logf("Issue %d: FlawID=%d, Category=%s", i, issue.FlawID, issue.Category)
		for j, action := range issue.MitigationActions {
			t.Logf("  Action %d: Type=%s, Reviewer=%s, Date=%s", j, action.Action, action.Reviewer, action.Date)
			if action.Comment != "" {
				t.Logf("    Comment: %s", action.Comment)
			}
		}
	}

	// Log errors
	for i, err := range mitigationInfo.Errors {
		t.Logf("Error %d: Type=%s, FlawIDList=%s", i, err.Type, err.FlawIDList)
	}
}

func TestGetMitigationInfoMultipleFlaws(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping test: credentials not available: %v", err)
	}

	if !client.IsConfigured() {
		t.Skip("Skipping test: API credentials not configured")
	}

	ctx := context.Background()
	buildID := int64(53691725)
	flawIDs := []int64{50, 51, 52} // Multiple flaw IDs

	mitigationInfo, err := client.GetMitigationInfo(ctx, buildID, flawIDs)
	if err != nil {
		t.Fatalf("Failed to get mitigation info: %v", err)
	}

	t.Logf("Retrieved mitigation info for %d flaws", len(flawIDs))
	t.Logf("Number of issues returned: %d", len(mitigationInfo.Issues))
	t.Logf("Number of errors returned: %d", len(mitigationInfo.Errors))
}

func TestGetMitigationInfoNoFlaws(t *testing.T) {
	client := NewClientUnconfigured()
	ctx := context.Background()

	_, err := client.GetMitigationInfo(ctx, 12345, []int64{})
	if err == nil {
		t.Error("Expected error when no flaw IDs provided, got nil")
	}
}
