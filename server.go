package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// MCPServer represents the core MCP server
type MCPServer struct {
	initialized     bool
	capabilities    ServerCapabilities
	tools           []Tool
	toolRegistry    *ToolRegistry
	handlerRegistry *ToolHandlerRegistry
	implRegistry    *ToolImplRegistry
}

func NewMCPServer() (*MCPServer, error) {
	// Create registries
	toolRegistry, err := LoadToolDefinitions("tools.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load tool definitions: %w", err)
	}

	handlerRegistry := NewToolHandlerRegistry()
	implRegistry := NewToolImplRegistry()

	s := &MCPServer{
		capabilities: ServerCapabilities{
			Tools: &ToolsCapability{
				ListChanged: false,
			},
			Resources: &ResourcesCapability{
				Subscribe:   false,
				ListChanged: false,
			},
			Prompts: &PromptsCapability{
				ListChanged: false,
			},
		},
		toolRegistry:    toolRegistry,
		handlerRegistry: handlerRegistry,
		implRegistry:    implRegistry,
	}

	// Initialize tool implementations and register their handlers
	if err := LoadAllTools(s.implRegistry, s.handlerRegistry); err != nil {
		return nil, fmt.Errorf("failed to load tool implementations: %w", err)
	}

	// Convert tool definitions to MCP tools for tools/list
	s.tools = toolRegistry.GetAllMCPTools()

	return s, nil
}

func (s *MCPServer) ServeStdio() error {
	transport := NewStdioTransport(s)
	return transport.Start()
}

func (s *MCPServer) ServeHTTP(addr string) error {
	transport := NewHTTPTransport(s)
	return transport.Start(addr)
}

func (s *MCPServer) HandleRequest(req *JSONRPCRequest) *JSONRPCResponse {
	log.Printf("Handling request: %s (id: %v)", req.Method, req.ID)

	resp := &JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
	}

	switch req.Method {
	case "initialize":
		result, err := s.handleInitialize(req.Params)
		if err != nil {
			resp.Error = &RPCError{
				Code:    -32603,
				Message: err.Error(),
			}
		} else {
			resp.Result = result
		}

	case "tools/list":
		resp.Result = s.handleListTools()

	case "tools/call":
		result, err := s.handleCallTool(req.Params)
		if err != nil {
			resp.Error = &RPCError{
				Code:    -32603,
				Message: err.Error(),
			}
		} else {
			resp.Result = result
		}

	case "resources/list":
		resp.Result = s.handleListResources()

	case "resources/read":
		result, err := s.handleReadResource(req.Params)
		if err != nil {
			resp.Error = &RPCError{
				Code:    -32603,
				Message: err.Error(),
			}
		} else {
			resp.Result = result
		}

	case "notifications/initialized":
		// Client confirms initialization - no response needed for notifications
		return &JSONRPCResponse{JSONRPC: "2.0"}

	default:
		resp.Error = &RPCError{
			Code:    -32601,
			Message: fmt.Sprintf("Method not found: %s", req.Method),
		}
	}

	return resp
}

func (s *MCPServer) handleInitialize(params json.RawMessage) (*InitializeResult, error) {
	var initParams InitializeParams
	if err := json.Unmarshal(params, &initParams); err != nil {
		return nil, fmt.Errorf("invalid initialize params: %w", err)
	}

	s.initialized = true
	log.Printf("Initialized by client: %s %s", initParams.ClientInfo.Name, initParams.ClientInfo.Version)

	return &InitializeResult{
		ProtocolVersion: "2024-11-05",
		Capabilities:    s.capabilities,
		ServerInfo: Implementation{
			Name:    "veracode-mcp-server",
			Version: "0.1.0",
		},
	}, nil
}

func (s *MCPServer) handleListTools() *ListToolsResult {
	return &ListToolsResult{
		Tools: s.tools,
	}
}

func (s *MCPServer) handleCallTool(params json.RawMessage) (*CallToolResult, error) {
	var callParams CallToolParams
	if err := json.Unmarshal(params, &callParams); err != nil {
		return nil, fmt.Errorf("invalid tool call params: %w", err)
	}

	// Look up the handler in the registry
	handler, exists := s.handlerRegistry.GetHandler(callParams.Name)
	if !exists {
		return &CallToolResult{
			Content: []Content{{
				Type: "text",
				Text: fmt.Sprintf("Unknown tool: %s. Available tools: %s",
					callParams.Name,
					s.getAvailableToolNames()),
			}},
			IsError: true,
		}, nil
	}

	// Validate required parameters against tool definition
	if s.toolRegistry != nil {
		toolDef := s.toolRegistry.GetToolByName(callParams.Name)
		if toolDef != nil {
			if err := s.validateToolArguments(toolDef, callParams.Arguments); err != nil {
				return &CallToolResult{
					Content: []Content{{Type: "text", Text: err.Error()}},
					IsError: true,
				}, nil
			}
		}
	}

	// Call the handler with context
	ctx := context.Background()
	result, err := handler(ctx, callParams.Arguments)
	if err != nil {
		return &CallToolResult{
			Content: []Content{{Type: "text", Text: fmt.Sprintf("Tool execution error: %v", err)}},
			IsError: true,
		}, nil
	}

	// Convert the result to CallToolResult format
	return convertToCallToolResult(result), nil
}

// convertToCallToolResult converts the generic tool result to MCP CallToolResult format
func convertToCallToolResult(result interface{}) *CallToolResult {
	// If it's already a CallToolResult, return it
	if ctr, ok := result.(*CallToolResult); ok {
		return ctr
	}

	// If it's a map, look for common response patterns
	if resultMap, ok := result.(map[string]interface{}); ok {
		// Check for error field
		if errMsg, hasErr := resultMap["error"]; hasErr {
			return &CallToolResult{
				Content: []Content{{Type: "text", Text: fmt.Sprintf("%v", errMsg)}},
				IsError: true,
			}
		}

		// Check for content field
		if content, hasContent := resultMap["content"]; hasContent {
			if contentList, ok := content.([]map[string]string); ok {
				contents := make([]Content, len(contentList))
				for i, c := range contentList {
					contents[i] = Content{
						Type: c["type"],
						Text: c["text"],
					}
				}
				return &CallToolResult{Content: contents}
			}
		}

		// If we have text field, use it
		if text, ok := resultMap["text"].(string); ok {
			return &CallToolResult{
				Content: []Content{{Type: "text", Text: text}},
			}
		}
	}

	// Default: convert result to JSON string
	jsonBytes, _ := json.Marshal(result)
	return &CallToolResult{
		Content: []Content{{Type: "text", Text: string(jsonBytes)}},
	}
}

// validateToolArguments checks that required parameters are present
func (s *MCPServer) validateToolArguments(toolDef *ToolDefinition, args map[string]interface{}) error {
	for _, param := range toolDef.Params {
		if param.IsRequired {
			value, exists := args[param.Name]
			if !exists || value == nil {
				return fmt.Errorf("missing required parameter: %s - %s", param.Name, param.Description)
			}

			// Additional validation for string parameters
			if param.Type == "string" {
				strVal, ok := value.(string)
				if !ok {
					return fmt.Errorf("parameter %s must be a string", param.Name)
				}
				if strVal == "" {
					return fmt.Errorf("parameter %s cannot be empty", param.Name)
				}
			}
		}
	}
	return nil
}

// getAvailableToolNames returns a comma-separated list of available tools
func (s *MCPServer) getAvailableToolNames() string {
	names := make([]string, len(s.tools))
	for i, tool := range s.tools {
		names[i] = tool.Name
	}
	return fmt.Sprintf("[%s]", fmt.Sprint(names))
}

func (s *MCPServer) handleListResources() *ListResourcesResult {
	return &ListResourcesResult{
		Resources: []Resource{
			{
				URI:         "resource://example/hello",
				Name:        "Hello Resource",
				Description: "A simple example resource",
				MimeType:    "text/plain",
			},
		},
	}
}

func (s *MCPServer) handleReadResource(params json.RawMessage) (*ReadResourceResult, error) {
	var readParams ReadResourceParams
	if err := json.Unmarshal(params, &readParams); err != nil {
		return nil, fmt.Errorf("invalid read resource params: %w", err)
	}

	switch readParams.URI {
	case "resource://example/hello":
		return &ReadResourceResult{
			Contents: []ResourceContents{
				{
					URI:      readParams.URI,
					MimeType: "text/plain",
					Text:     "Hello from MCP resource!",
				},
			},
		}, nil

	default:
		return nil, fmt.Errorf("resource not found: %s", readParams.URI)
	}
}
