package mcp_tools

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

// ============================================================================
// parsePipelineFindingsRequest tests
// ============================================================================

func TestParsePipelineFindingsRequest_Success(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parsePipelineFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.ApplicationPath != "/path/to/app" {
		t.Errorf("Expected application_path '/path/to/app', got '%s'", req.ApplicationPath)
	}
	if req.Size != 10 {
		t.Errorf("Expected default size 10, got %d", req.Size)
	}
}

func TestParsePipelineFindingsRequest_WithPagination(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"page_size":        float64(50),
		"page":             float64(2),
	}

	req, err := parsePipelineFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.Size != 50 {
		t.Errorf("Expected size 50, got %d", req.Size)
	}
	if req.Page != 2 {
		t.Errorf("Expected page 2, got %d", req.Page)
	}
}

func TestParsePipelineFindingsRequest_SizeValidation(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"page_size":        float64(501),
	}

	_, err := parsePipelineFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for size=501")
	}
}

func TestParsePipelineFindingsRequest_PageValidation(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"page":             float64(-1),
	}

	_, err := parsePipelineFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for page=-1")
	}
}

func TestParsePipelineFindingsRequest_MissingApplicationPath(t *testing.T) {
	args := map[string]interface{}{}

	_, err := parsePipelineFindingsRequest(args)
	if err == nil {
		t.Fatal("Expected error for missing application_path")
	}
}

func TestParsePipelineFindingsRequest_CWEIDs(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
		"cwe_ids":          []interface{}{float64(502), float64(89)},
	}

	req, err := parsePipelineFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(req.CWEIDs) != 2 {
		t.Fatalf("Expected 2 CWE IDs, got %d", len(req.CWEIDs))
	}
	if req.CWEIDs[0] != "502" {
		t.Errorf("Expected CWE ID '502', got '%s'", req.CWEIDs[0])
	}
	if req.CWEIDs[1] != "89" {
		t.Errorf("Expected CWE ID '89', got '%s'", req.CWEIDs[1])
	}
}

func TestParsePipelineFindingsRequest_NoCWEIDs(t *testing.T) {
	args := map[string]interface{}{
		"application_path": "/path/to/app",
	}

	req, err := parsePipelineFindingsRequest(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if req.CWEIDs != nil {
		t.Errorf("Expected nil CWE IDs, got %v", req.CWEIDs)
	}
}

func TestFormatPipelineFindingsResponse_CWEFilter(t *testing.T) {
	ctx := context.WithValue(context.Background(), UICapabilityKey, true)
	findings := []PipelineFlaw{
		{IssueID: 1, CWEID: "502", Severity: 4, Title: "Deserialization"},
		{IssueID: 2, CWEID: "89", Severity: 3, Title: "SQL Injection"},
		{IssueID: 3, CWEID: "79", Severity: 3, Title: "XSS"},
	}
	results := &PipelineScanResults{Findings: findings}

	req := &PipelineFindingsRequest{
		ApplicationPath: "/app",
		Size:            10,
		Page:            0,
		CWEIDs:          []string{"502"},
	}

	result := formatPipelineFindingsResponse(ctx, "/app", "results.json", results, nil, req)

	structured, ok := result["structuredContent"].(MCPFindingsResponse)
	if !ok {
		t.Fatal("Expected structuredContent to be MCPFindingsResponse")
	}

	if len(structured.Findings) != 1 {
		t.Fatalf("Expected 1 finding after CWE filter, got %d", len(structured.Findings))
	}
	if structured.Findings[0].CweId != 502 {
		t.Errorf("Expected CWE ID 502, got %d", structured.Findings[0].CweId)
	}
}

func TestFormatPipelineFindingsResponse_NoCWEFilter(t *testing.T) {
	ctx := context.WithValue(context.Background(), UICapabilityKey, true)
	findings := []PipelineFlaw{
		{IssueID: 1, CWEID: "502", Severity: 4, Title: "Deserialization"},
		{IssueID: 2, CWEID: "89", Severity: 3, Title: "SQL Injection"},
	}
	results := &PipelineScanResults{Findings: findings}

	req := &PipelineFindingsRequest{
		ApplicationPath: "/app",
		Size:            10,
		Page:            0,
	}

	result := formatPipelineFindingsResponse(ctx, "/app", "results.json", results, nil, req)

	structured, ok := result["structuredContent"].(MCPFindingsResponse)
	if !ok {
		t.Fatal("Expected structuredContent to be MCPFindingsResponse")
	}

	if len(structured.Findings) != 2 {
		t.Errorf("Expected 2 findings with no CWE filter, got %d", len(structured.Findings))
	}
}

// ============================================================================
// handleGetPipelineFindings tests
// ============================================================================

func TestPipelineFindingsTool_Initialize(t *testing.T) {
	tool := NewPipelineFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()
}

func TestPipelineFindingsTool_RegisterHandlers(t *testing.T) {
	tool := NewPipelineFindingsTool()
	if err := tool.Initialize(); err != nil {
		t.Fatalf("Failed to initialize tool: %v", err)
	}
	defer tool.Shutdown()

	registry := newMockHandlerRegistry()
	if err := tool.RegisterHandlers(registry); err != nil {
		t.Fatalf("Failed to register handlers: %v", err)
	}

	if registry.handlers[PipelineFindingsToolName] == nil {
		t.Error("Handler was not registered")
	}
}

func TestPipelineFindingsTool_HandleMissingResultsFile(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory without results file
	tempDir := t.TempDir()

	result, err := handlePipelineFindings(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should return structured content with empty findings for UI
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	if resultMap["content"] == nil {
		t.Error("Expected content in response")
	}

	// Verify structuredContent is present with proper structure
	if resultMap["structuredContent"] == nil {
		t.Error("Expected structuredContent in response for UI rendering")
	}

	structuredContent, ok := resultMap["structuredContent"].(MCPFindingsResponse)
	if !ok {
		t.Fatal("Expected structuredContent to be MCPFindingsResponse")
	}

	// Verify empty findings structure
	if len(structuredContent.Findings) != 0 {
		t.Errorf("Expected 0 findings, got %d", len(structuredContent.Findings))
	}

	if structuredContent.Summary.TotalFindings != 0 {
		t.Errorf("Expected 0 total findings, got %d", structuredContent.Summary.TotalFindings)
	}

	if structuredContent.Application.Name == "" {
		t.Error("Expected application name to be set")
	}
}

func TestPipelineFindingsTool_HandleValidPath(t *testing.T) {
	ctx := context.Background()

	// Create a temporary directory with pipeline structure
	tempDir := t.TempDir()
	pipelineDir := filepath.Join(tempDir, ".veracode", "pipeline")
	if err := os.MkdirAll(pipelineDir, 0755); err != nil {
		t.Fatalf("Failed to create pipeline directory: %v", err)
	}

	result, err := handlePipelineFindings(ctx, map[string]interface{}{
		"application_path": tempDir,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("Expected map result")
	}

	// Should always have content (even if no results file exists)
	if resultMap["content"] == nil {
		t.Error("Expected content in response")
	}
}
