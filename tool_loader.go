package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

//go:embed tools.json
var toolsJSON []byte

// ToolDefinition represents a tool definition from JSON
type ToolDefinition struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Category    string            `json:"category"`
	Params      []ParamDefinition `json:"params"`
}

// ParamDefinition represents a parameter definition from JSON
type ParamDefinition struct {
	Name          string           `json:"name"`
	Type          string           `json:"type"`
	ItemType      string           `json:"itemType,omitempty"`
	IsRequired    bool             `json:"isRequired"`
	AllowedValues []string         `json:"allowedValues,omitempty"`
	Validation    *ValidationRules `json:"validation,omitempty"`
	Description   string           `json:"description"`
}

// ValidationRules represents validation constraints for a parameter
type ValidationRules struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}

// ToolRegistry represents the collection of all tool definitions
type ToolRegistry struct {
	Tools []ToolDefinition `json:"tools"`
}

// LoadToolDefinitions loads tool definitions from the embedded JSON file
func LoadToolDefinitions() (*ToolRegistry, error) {
	var registry ToolRegistry
	if err := json.Unmarshal(toolsJSON, &registry); err != nil {
		return nil, fmt.Errorf("failed to parse tools JSON: %w", err)
	}

	return &registry, nil
}

// ToMCPTool converts a ToolDefinition to an MCP Tool structure
func (td *ToolDefinition) ToMCPTool() Tool {
	// Build input schema from parameter definitions
	properties := make(map[string]interface{})
	required := []string{}

	for _, param := range td.Params {
		propSchema := map[string]interface{}{
			"description": param.Description,
		}

		switch param.Type {
		case "string":
			propSchema["type"] = "string"
			if len(param.AllowedValues) > 0 {
				propSchema["enum"] = param.AllowedValues
			}

		case "number", "integer":
			propSchema["type"] = param.Type
			if param.Validation != nil {
				if param.Validation.Min != nil {
					propSchema["minimum"] = *param.Validation.Min
				}
				if param.Validation.Max != nil {
					propSchema["maximum"] = *param.Validation.Max
				}
			}

		case "boolean":
			propSchema["type"] = "boolean"

		case "array":
			propSchema["type"] = "array"
			items := map[string]interface{}{}
			if param.ItemType != "" {
				items["type"] = param.ItemType
			}
			if len(param.AllowedValues) > 0 {
				items["enum"] = param.AllowedValues
			}
			propSchema["items"] = items
		}

		properties[param.Name] = propSchema

		if param.IsRequired {
			required = append(required, param.Name)
		}
	}

	inputSchema := map[string]interface{}{
		"type":       "object",
		"properties": properties,
	}

	if len(required) > 0 {
		inputSchema["required"] = required
	}

	return Tool{
		Name:        td.Name,
		Description: td.Description,
		InputSchema: inputSchema,
		Meta:        td.GetToolMeta(),
	}
}

// GetToolMeta returns tool metadata including UI configuration if applicable
func (td *ToolDefinition) GetToolMeta() map[string]interface{} {
	// Add UI metadata for tools with interactive UIs
	// Provide BOTH formats for maximum compatibility:
	// - Flat key "ui/resourceUri" (MCP Apps legacy format)
	// - Nested "ui.resourceUri" (MCP Apps current format)

	log.Printf("GetToolMeta called for tool: %s", td.Name)

	if td.Name == "pipeline-results" {
		meta := map[string]interface{}{
			"ui/resourceUri": "ui://pipeline-results/app.html",
			"ui": map[string]interface{}{
				"resourceUri": "ui://pipeline-results/app.html",
			},
		}
		log.Printf("Returning UI metadata for pipeline-results: %+v", meta)
		return meta
	}
	if td.Name == "static-findings" {
		meta := map[string]interface{}{
			"ui/resourceUri": "ui://static-findings/app.html",
			"ui": map[string]interface{}{
				"resourceUri": "ui://static-findings/app.html",
			},
		}
		log.Printf("Returning UI metadata for static-findings: %+v", meta)
		return meta
	}
	if td.Name == "dynamic-findings" {
		meta := map[string]interface{}{
			"ui/resourceUri": "ui://dynamic-findings/app.html",
			"ui": map[string]interface{}{
				"resourceUri": "ui://dynamic-findings/app.html",
			},
		}
		log.Printf("Returning UI metadata for dynamic-findings: %+v", meta)
		return meta
	}
	return nil
}

// GetToolByName finds a tool definition by name
func (tr *ToolRegistry) GetToolByName(name string) *ToolDefinition {
	for i := range tr.Tools {
		if tr.Tools[i].Name == name {
			return &tr.Tools[i]
		}
	}
	return nil
}

// GetAllMCPTools converts all tool definitions to MCP tools
func (tr *ToolRegistry) GetAllMCPTools() []Tool {
	tools := make([]Tool, len(tr.Tools))
	for i, td := range tr.Tools {
		tools[i] = td.ToMCPTool()
	}
	return tools
}
