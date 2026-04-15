package mcp_tools

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestParseGetLocalSCASummaryRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}
	req, err := parseGetLocalSCASummaryRequest(args)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
	if req.SeverityGTE != nil {
		t.Errorf("Expected nil SeverityGTE, got %v", req.SeverityGTE)
	}
}

func TestParseGetLocalSCASummaryRequest_WithSeverityGTE(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"severity_gte":     float64(4),
	}
	req, err := parseGetLocalSCASummaryRequest(args)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if req.SeverityGTE == nil || *req.SeverityGTE != 4 {
		t.Errorf("Expected SeverityGTE=4, got %v", req.SeverityGTE)
	}
}

func TestParseGetLocalSCASummaryRequest_MissingApplicationPath(t *testing.T) {
	_, err := parseGetLocalSCASummaryRequest(map[string]interface{}{})
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

func TestGetLocalSCASummary_MissingResultsFile(t *testing.T) {
	ctx := context.Background()
	tempDir := t.TempDir()

	result, err := handleGetLocalSCASummary(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}
	if resultMap["content"] == nil {
		t.Error("Expected content in response")
	}
}

// sampleSCAResults returns a multi-component SCAFindings fixture for testing
func sampleSCAResults() SCAFindings {
	return SCAFindings{
		Vulnerabilities: SCAVulnerabilities{
			Matches: []SCAMatch{
				// Component A: two CVEs, with suggested upgrade versions
				{
					Artifact: SCAArtifact{
						Name:     "lodash",
						Version:  "4.17.15",
						Type:     "npm",
						Language: "javascript",
					},
					Vulnerability: SCAVulnerability{
						ID:       "CVE-2021-23337",
						Severity: "High",
						EPSS:     []SCAEPSS{{CVE: "CVE-2021-23337", EPSS: 0.05}},
					},
					MatchDetails: []SCAMatchDetail{{Fix: SCAMatchFix{SuggestedVersion: "4.17.21"}}},
				},
				{
					Artifact: SCAArtifact{
						Name:     "lodash",
						Version:  "4.17.15",
						Type:     "npm",
						Language: "javascript",
					},
					Vulnerability: SCAVulnerability{
						ID:       "CVE-2020-8203",
						Severity: "Critical",
						EPSS:     []SCAEPSS{{CVE: "CVE-2020-8203", EPSS: 0.12}},
					},
					MatchDetails: []SCAMatchDetail{{Fix: SCAMatchFix{SuggestedVersion: "4.17.19"}}},
				},
				// Component B: one CVE, no fix version
				{
					Artifact: SCAArtifact{
						Name:     "some-lib",
						Version:  "1.0.0",
						Type:     "npm",
						Language: "javascript",
					},
					Vulnerability: SCAVulnerability{
						ID:       "CVE-2023-99999",
						Severity: "Medium",
					},
					MatchDetails: []SCAMatchDetail{},
				},
			},
		},
	}
}

func TestBuildComponentSummaries_GroupsByComponent(t *testing.T) {
	results := sampleSCAResults()
	req := &GetLocalSCASummaryRequest{ApplicationPath: "/tmp/app"}

	summaries := buildComponentSummaries(&results, req)

	if len(summaries) != 2 {
		t.Fatalf("Expected 2 component summaries, got %d", len(summaries))
	}

	// lodash should appear first (has Critical CVE)
	if summaries[0].Name != "lodash" {
		t.Errorf("Expected lodash first (highest severity), got %s", summaries[0].Name)
	}
	if len(summaries[0].CVEs) != 2 {
		t.Errorf("Expected 2 CVEs for lodash, got %d", len(summaries[0].CVEs))
	}
}

func TestBuildComponentSummaries_PicksHighestUpgradeVersion(t *testing.T) {
	results := sampleSCAResults()
	req := &GetLocalSCASummaryRequest{ApplicationPath: "/tmp/app"}

	summaries := buildComponentSummaries(&results, req)

	var lodash componentSummary
	found := false
	for _, s := range summaries {
		if s.Name == "lodash" {
			lodash = s
			found = true
			break
		}
	}
	if !found {
		t.Fatal("lodash not found in summaries")
	}

	// Should pick 4.17.21 over 4.17.19
	if lodash.RecommendedUpgrade != "4.17.21" {
		t.Errorf("Expected recommended upgrade 4.17.21, got %s", lodash.RecommendedUpgrade)
	}
}

func TestBuildComponentSummaries_NoFixVersion(t *testing.T) {
	results := sampleSCAResults()
	req := &GetLocalSCASummaryRequest{ApplicationPath: "/tmp/app"}

	summaries := buildComponentSummaries(&results, req)

	var someLib componentSummary
	found := false
	for _, s := range summaries {
		if s.Name == "some-lib" {
			someLib = s
			found = true
			break
		}
	}
	if !found {
		t.Fatal("some-lib not found in summaries")
	}
	if someLib.RecommendedUpgrade != "" {
		t.Errorf("Expected empty recommended upgrade, got %s", someLib.RecommendedUpgrade)
	}
}

func TestBuildComponentSummaries_SeverityFilter(t *testing.T) {
	results := sampleSCAResults()
	minSev := 4 // High and above
	req := &GetLocalSCASummaryRequest{ApplicationPath: "/tmp/app", SeverityGTE: &minSev}

	summaries := buildComponentSummaries(&results, req)

	// Only lodash has High/Critical — some-lib (Medium) should be excluded
	for _, s := range summaries {
		if s.Name == "some-lib" {
			t.Error("Expected some-lib to be filtered out by severity_gte=4")
		}
	}
	if len(summaries) != 1 || summaries[0].Name != "lodash" {
		t.Errorf("Expected only lodash, got %v", summaries)
	}
}

func TestPickHighestVersion(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{[]string{"4.17.21", "4.17.19"}, "4.17.21"},
		{[]string{"1.0.0", "2.0.0", "1.5.0"}, "2.0.0"},
		{[]string{"v1.2.3", "v1.10.0"}, "v1.10.0"},
		{[]string{"3.0.0"}, "3.0.0"},
		{[]string{}, ""},
	}
	for _, c := range cases {
		got := pickHighestVersion(c.input)
		if got != c.expected {
			t.Errorf("pickHighestVersion(%v) = %q, want %q", c.input, got, c.expected)
		}
	}
}

func TestCompareVersionStrings(t *testing.T) {
	cases := []struct {
		a, b string
		want int // sign only: negative, zero, positive
	}{
		{"4.17.21", "4.17.19", 1},
		{"4.17.19", "4.17.21", -1},
		{"1.0.0", "1.0.0", 0},
		{"2.0.0", "1.9.9", 1},
		{"1.10.0", "1.9.0", 1},
	}
	for _, c := range cases {
		got := compareVersionStrings(c.a, c.b)
		switch {
		case c.want > 0 && got <= 0:
			t.Errorf("compareVersionStrings(%q, %q) = %d, want positive", c.a, c.b, got)
		case c.want < 0 && got >= 0:
			t.Errorf("compareVersionStrings(%q, %q) = %d, want negative", c.a, c.b, got)
		case c.want == 0 && got != 0:
			t.Errorf("compareVersionStrings(%q, %q) = %d, want 0", c.a, c.b, got)
		}
	}
}

func TestGetLocalSCASummary_ValidResultsFile(t *testing.T) {
	ctx := context.Background()
	tempDir := t.TempDir()
	scaDir := veracodeWorkDir(tempDir, "sca")
	_ = os.RemoveAll(filepath.Dir(scaDir))
	t.Cleanup(func() { _ = os.RemoveAll(filepath.Dir(scaDir)) })
	if err := os.MkdirAll(scaDir, 0750); err != nil {
		t.Fatalf("Failed to create SCA directory: %v", err)
	}

	results := sampleSCAResults()
	data, err := json.Marshal(results)
	if err != nil {
		t.Fatalf("Failed to marshal results: %v", err)
	}
	if err := os.WriteFile(filepath.Join(scaDir, "veracode.json"), data, 0644); err != nil {
		t.Fatalf("Failed to write results file: %v", err)
	}

	result, err := handleGetLocalSCASummary(ctx, map[string]interface{}{
		"application_path": tempDir,
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}
	if resultMap["content"] == nil {
		t.Error("Expected content in response")
	}
	if resultMap["structuredContent"] == nil {
		t.Error("Expected structuredContent in response")
	}

	// Validate structured content has components
	sc, ok := resultMap["structuredContent"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected structuredContent to be a map")
	}
	components, ok := sc["components"].([]map[string]interface{})
	if !ok {
		t.Fatal("Expected components in structuredContent")
	}
	if len(components) != 2 {
		t.Errorf("Expected 2 components, got %d", len(components))
	}
}
