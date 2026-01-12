package api

import (
	"context"
	"testing"
	"time"
)

const testAppMCPVerademo = "f4e74197-1e26-42c4-ab4b-245870c93280" // MCPVerademo

// TestGetScaFindings_Integration tests SCA findings retrieval
func TestGetScaFindings_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client, err := NewVeracodeClient()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req := FindingsRequest{
		AppProfile: testAppMCPVerademo,
		Page:       0,
		Size:       10,
	}

	resp, err := client.GetScaFindings(ctx, req)
	if err != nil {
		t.Fatalf("Failed to get SCA findings: %v", err)
	}

	t.Logf("SCA findings result:")
	t.Logf("  Total count: %d", resp.TotalCount)
	t.Logf("  Page: %d", resp.Page)
	t.Logf("  Size: %d", resp.Size)
	t.Logf("  Findings returned: %d", len(resp.Findings))

	if len(resp.Findings) == 0 {
		t.Log("WARNING: No SCA findings found for MCPVerademo app")
		return
	}

	// Display first few findings
	displayCount := 3
	if len(resp.Findings) < displayCount {
		displayCount = len(resp.Findings)
	}

	for i := 0; i < displayCount; i++ {
		finding := resp.Findings[i]
		t.Logf("  Finding %d:", i+1)
		t.Logf("    ID: %s", finding.ID)
		t.Logf("    Status: %s", finding.Status)
		t.Logf("    ResolutionStatus: %s", finding.ResolutionStatus)
		t.Logf("    ViolatesPolicy: %v", finding.ViolatesPolicy)
		t.Logf("    Severity: %s", finding.Severity)
		t.Logf("    SeverityScore: %d", finding.SeverityScore)
		t.Logf("    CWE: %s", finding.CWE)
		t.Logf("    CVE: %s", finding.CVE)
		t.Logf("    ComponentFilename: %s", finding.ComponentFilename)
		t.Logf("    ComponentVersion: %s", finding.ComponentVersion)
		t.Logf("    Description: %.100s...", finding.Description)
	}

	// Verify SCA-specific fields are populated
	hasComponent := false
	hasCVE := false
	for _, finding := range resp.Findings {
		if finding.ComponentFilename != "" {
			hasComponent = true
		}
		if finding.CVE != "" {
			hasCVE = true
		}
	}

	if !hasComponent {
		t.Error("Expected at least one finding to have ComponentFilename populated")
	}

	t.Logf("SCA-specific fields found: hasComponent=%v, hasCVE=%v", hasComponent, hasCVE)
}
