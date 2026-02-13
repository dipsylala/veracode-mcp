package xml

import (
	"context"
	"encoding/xml"
	"testing"
)

func TestGetMitigationInfoDetailed(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping test: credentials not available: %v", err)
	}

	if !client.IsConfigured() {
		t.Skip("Skipping test: API credentials not configured")
	}

	ctx := context.Background()
	buildID := int64(53691725)
	flawIDs := []int64{50}

	mitigationInfo, err := client.GetMitigationInfo(ctx, buildID, flawIDs)
	if err != nil {
		t.Fatalf("Failed to get mitigation info: %v", err)
	}

	// Marshal back to XML to see the full response
	xmlBytes, err := xml.MarshalIndent(mitigationInfo, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %v", err)
	}

	t.Logf("Full XML Response:\n%s", string(xmlBytes))

	// Verify structure
	t.Logf("\nParsed Structure:")
	t.Logf("Version: %s", mitigationInfo.Version)
	t.Logf("BuildID: %d", mitigationInfo.BuildID)
	t.Logf("Issues: %d", len(mitigationInfo.Issues))
	t.Logf("Errors: %d", len(mitigationInfo.Errors))

	for i, issue := range mitigationInfo.Issues {
		t.Logf("\nIssue %d:", i)
		t.Logf("  FlawID: %d", issue.FlawID)
		t.Logf("  Category: %s", issue.Category)
		t.Logf("  Mitigation Actions: %d", len(issue.MitigationActions))

		for j, action := range issue.MitigationActions {
			t.Logf("    Action %d:", j)
			t.Logf("      Type: %s", action.Action)
			t.Logf("      Desc: %s", action.Desc)
			t.Logf("      Reviewer: %s", action.Reviewer)
			t.Logf("      Date: %s", action.Date)
			t.Logf("      Comment: %s", action.Comment)
		}
	}

	for i, errItem := range mitigationInfo.Errors {
		t.Logf("\nError %d:", i)
		t.Logf("  Type: %s", errItem.Type)
		t.Logf("  FlawIDList: %s", errItem.FlawIDList)
	}
}
