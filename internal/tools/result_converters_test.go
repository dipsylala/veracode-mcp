package tools

import (
	"encoding/json"
	"testing"
)

func TestConvertToCallToolResult_WithStructuredContent(t *testing.T) {
	// Create a result map with structuredContent (like the tools return)
	resultMap := map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": "Summary text for LLM",
			},
		},
		"structuredContent": map[string]interface{}{
			"findings": []map[string]interface{}{
				{
					"id":       "123",
					"severity": "HIGH",
					"title":    "SQL Injection",
				},
			},
			"summary": map[string]interface{}{
				"total": 1,
				"high":  1,
			},
		},
	}

	// Convert to CallToolResult
	result := ConvertToCallToolResult(resultMap)

	// Verify structuredContent was set
	if result.StructuredContent == nil {
		t.Fatal("StructuredContent was not set")
	}

	// Verify it can be marshaled to JSON
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Failed to marshal CallToolResult: %v", err)
	}

	// Verify the JSON contains structuredContent
	jsonStr := string(jsonBytes)
	if !contains(jsonStr, "structuredContent") {
		t.Errorf("Marshaled JSON does not contain 'structuredContent': %s", jsonStr)
	}

	t.Logf("SUCCESS: Marshaled JSON contains structuredContent")
	t.Logf("JSON length: %d bytes", len(jsonBytes))
	t.Logf("JSON (first 500 chars): %s", jsonStr[:min(500, len(jsonStr))])
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
