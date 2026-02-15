package rest

import (
	"encoding/json"
	"testing"

	"github.com/dipsylala/veracode-mcp/api/rest/generated/findings"
)

func TestUnmarshalFindingDetails(t *testing.T) {
	// Real finding_details JSON from the API
	jsonData := `{
		"attack_vector": "system_diagnostics_process_dll.System.Diagnostics.Process.Start",
		"cwe": {
			"href": "https://api.veracode.com/appsec/v1/cwes/78",
			"id": 78,
			"name": "Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')"
		},
		"exploitability": 0,
		"file_line_number": 68,
		"file_name": "toolscontroller.cs",
		"file_path": "users/mwatson/documents/github/verademo-net/app/controllers/toolscontroller.cs",
		"finding_category": {
			"href": "https://api.veracode.com/appsec/v1/categories/18",
			"id": 18,
			"name": "Command or Argument Injection"
		},
		"module": "app.dll",
		"procedure": "app_dll.Verademo.Controllers.ToolsController.Ping",
		"relative_location": 62,
		"severity": 5
	}`

	var details findings.FindingFindingDetails
	err := json.Unmarshal([]byte(jsonData), &details)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	t.Logf("Details unmarshaled successfully")
	t.Logf("StaticFinding: %v", details.StaticFinding)
	t.Logf("DynamicFinding: %v", details.DynamicFinding)

	if details.StaticFinding == nil {
		t.Fatal("StaticFinding is nil!")
	}

	if details.StaticFinding.Severity == nil {
		t.Fatal("Severity is nil!")
	}

	t.Logf("Severity: %d", *details.StaticFinding.Severity)
	t.Logf("CWE ID: %d", *details.StaticFinding.Cwe.Id)
	t.Logf("FilePath: %s", *details.StaticFinding.FilePath)
	t.Logf("FileLineNumber: %d", *details.StaticFinding.FileLineNumber)
}
