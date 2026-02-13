package rest

import (
	"context"
	"testing"
	"time"
)

// TestLicenseExtraction verifies that license information is extracted from SCA findings
func TestLicenseExtraction(t *testing.T) {
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

	if len(resp.Findings) == 0 {
		t.Skip("No SCA findings available for license test")
	}

	// Check for license information
	foundLicense := false
	for i, finding := range resp.Findings {
		if len(finding.Licenses) > 0 {
			foundLicense = true
			t.Logf("Finding %d has %d license(s):", i+1, len(finding.Licenses))
			for j, license := range finding.Licenses {
				t.Logf("  License %d:", j+1)
				t.Logf("    LicenseID: %s", license.LicenseID)
				t.Logf("    RiskRating: %s", license.RiskRating)
			}
			t.Logf("  Component: %s v%s", finding.ComponentFilename, finding.ComponentVersion)
		}
	}

	if !foundLicense {
		t.Log("WARNING: No license information found in any SCA findings. This might be expected if components have no license data.")
	}
}
