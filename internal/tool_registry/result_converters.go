package tools

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dipsylala/veracode-mcp/internal/types"
)

// Tool result conversion utilities for transforming various result formats
// into the standardized MCP CallToolResult format.

// ConvertToCallToolResult converts the generic tool result to MCP CallToolResult format.
// This handles different result types that tools may return and normalizes them
// for consistent MCP protocol compliance.
func ConvertToCallToolResult(result interface{}) *types.CallToolResult {
	// If it's already a CallToolResult, return it as-is
	if ctr, ok := result.(*types.CallToolResult); ok {
		return ctr
	}

	// If it's a map, try to convert it using structured conversion
	if resultMap, ok := result.(map[string]interface{}); ok {
		return convertMapToCallToolResult(resultMap)
	}

	// Default: convert result to JSON string format
	return marshalResultAsJSON(result)
}

// convertMapToCallToolResult handles map[string]interface{} results
// and transforms them into proper MCP CallToolResult structures.
func convertMapToCallToolResult(resultMap map[string]interface{}) *types.CallToolResult {
	// Check for error field first
	if errMsg, hasErr := resultMap["error"]; hasErr {
		return &types.CallToolResult{
			Content: []types.Content{{Type: "text", Text: fmt.Sprintf("%v", errMsg)}},
			IsError: true,
		}
	}

	result := &types.CallToolResult{}

	// Check for content field (standard MCP content)
	if content, hasContent := resultMap["content"]; hasContent {
		if contents := convertContentField(content); contents != nil {
			result.Content = contents
		}
	}

	// Check for meta field (for MCP Apps UI or additional metadata)
	if meta, hasMeta := resultMap["meta"]; hasMeta {
		result.Meta = meta
	}

	// Check for structuredContent field (for MCP Apps UI data)
	// Always include this if present - UI capability is checked elsewhere
	if structuredContent, hasStructured := resultMap["structuredContent"]; hasStructured {
		log.Printf("[CONVERTER] Found structuredContent in map, type=%T, nil=%v", structuredContent, structuredContent == nil)
		// Try direct assignment first (if already map[string]interface{})
		if sc, ok := structuredContent.(map[string]interface{}); ok {
			result.StructuredContent = sc
		} else {
			// Convert struct to map[string]interface{} via JSON
			// This is necessary because tool results often return structs
			// but MCP protocol requires map[string]interface{}
			scJSON, err := json.Marshal(structuredContent)
			if err == nil {
				var scMap map[string]interface{}
				if unmarshalErr := json.Unmarshal(scJSON, &scMap); unmarshalErr == nil {
					result.StructuredContent = scMap
					log.Printf("[CONVERTER] Converted structuredContent from %T to map[string]interface{}", structuredContent)
				} else {
					log.Printf("[CONVERTER] Failed to unmarshal structuredContent: %v", unmarshalErr)
				}
			} else {
				log.Printf("[CONVERTER] Failed to marshal structuredContent: %v", err)
			}
		}
		log.Printf("[CONVERTER] Set result.StructuredContent, type=%T, nil=%v", result.StructuredContent, result.StructuredContent == nil)
	}

	// If we have content, meta, or structuredContent, return the result
	if len(result.Content) > 0 || result.Meta != nil || result.StructuredContent != nil {
		return result
	}

	// If we have a text field, use it directly
	if text, ok := resultMap["text"].(string); ok {
		return &types.CallToolResult{
			Content: []types.Content{{Type: "text", Text: text}},
		}
	}

	// Fallback to JSON marshaling
	return marshalResultAsJSON(resultMap)
}

// convertContentField converts various content field formats to []Content.
// This handles both simple text content and more complex resource content.
func convertContentField(content interface{}) []types.Content {
	// Try []map[string]interface{} first (for resources and complex content)
	if contentList, ok := content.([]map[string]interface{}); ok {
		return convertDetailedContentList(contentList)
	}

	// Fallback to []map[string]string (for simple text content)
	if contentList, ok := content.([]map[string]string); ok {
		return convertSimpleContentList(contentList)
	}

	return nil
}

// convertDetailedContentList converts []map[string]interface{} to []Content.
// This handles complex content with additional metadata.
func convertDetailedContentList(contentList []map[string]interface{}) []types.Content {
	contents := make([]types.Content, len(contentList))
	for i, c := range contentList {
		cont := types.Content{}
		if typ, ok := c["type"].(string); ok {
			cont.Type = typ
		}
		if text, ok := c["text"].(string); ok {
			cont.Text = text
		}
		if mimeType, ok := c["mimeType"].(string); ok {
			cont.MimeType = mimeType
		}
		if data, ok := c["data"].(string); ok {
			cont.Data = data
		}
		contents[i] = cont
	}
	return contents
}

// convertSimpleContentList converts []map[string]string to []Content.
// This handles simple text-based content without additional metadata.
func convertSimpleContentList(contentList []map[string]string) []types.Content {
	contents := make([]types.Content, len(contentList))
	for i, c := range contentList {
		contents[i] = types.Content{
			Type: c["type"],
			Text: c["text"],
		}
	}
	return contents
}

// marshalResultAsJSON converts any result to JSON string format.
// This is the fallback conversion method when other structured
// conversions are not applicable.
func marshalResultAsJSON(result interface{}) *types.CallToolResult {
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		log.Printf("Failed to marshal result: %v", err)
		return &types.CallToolResult{
			Content: []types.Content{{Type: "text", Text: fmt.Sprintf("Error: %v", err)}},
			IsError: true,
		}
	}
	return &types.CallToolResult{
		Content: []types.Content{{Type: "text", Text: string(jsonBytes)}},
	}
}
