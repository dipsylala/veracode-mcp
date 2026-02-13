package rest

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestEndToEndDynamicFindingsExtraction verifies the complete data extraction pipeline
func TestEndToEndDynamicFindingsExtraction(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Test dynamic findings with MCPVerademo
	req := FindingsRequest{
		AppProfile: testAppDynamicID, // MCPVerademo
		Page:       0,
		Size:       5,
	}

	resp, err := client.GetDynamicFindings(ctx, req)
	if err != nil {
		t.Fatalf("Failed: %v", err)
	}

	if len(resp.Findings) == 0 {
		t.Fatal("No findings returned")
	}

	t.Logf("Testing dynamic findings extraction with %d findings\n", len(resp.Findings))

	// Verify each finding has the expected fields populated
	for i, finding := range resp.Findings {
		t.Logf("\nFinding %d:", i+1)
		t.Logf("  ID: %s", finding.ID)
		t.Logf("  CWE: %s", finding.CWE)
		t.Logf("  Severity: %s", finding.Severity)
		t.Logf("  SeverityScore: %d", finding.SeverityScore)
		t.Logf("  Status: %s", finding.Status)
		t.Logf("  ResolutionStatus: %s", finding.ResolutionStatus)
		t.Logf("  ViolatesPolicy: %v", finding.ViolatesPolicy)

		// Extract numeric CWE ID to simulate what MCP tool does
		var cweID int32
		if finding.CWE != "" {
			_, _ = fmt.Sscanf(finding.CWE, "CWE-%d", &cweID)
		}
		t.Logf("  Extracted CWE ID: %d", cweID)

		// Verify critical fields are not empty/zero
		if finding.CWE == "" {
			t.Errorf("Finding %s: CWE is empty", finding.ID)
		}
		if finding.SeverityScore == 0 && finding.ViolatesPolicy {
			t.Errorf("Finding %s: SeverityScore is 0 but violates policy", finding.ID)
		}
		if cweID == 0 && finding.CWE != "" {
			t.Errorf("Finding %s: CWE extraction failed (CWE=%s)", finding.ID, finding.CWE)
		}
	}
}
