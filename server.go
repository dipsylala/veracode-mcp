package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

// Context key for UI capability
type contextKey string

const uiCapabilityKey contextKey = "uiCapability"

// WithUICapability adds UI capability information to the context
func WithUICapability(ctx context.Context, supportsUI bool) context.Context {
	return context.WithValue(ctx, uiCapabilityKey, supportsUI)
}

// ClientSupportsUIFromContext retrieves UI capability from context
func ClientSupportsUIFromContext(ctx context.Context) bool {
	if val := ctx.Value(uiCapabilityKey); val != nil {
		if supportsUI, ok := val.(bool); ok {
			return supportsUI
		}
	}
	return false
}

//go:embed ui/pipeline-results-app/dist/mcp-app.html
var pipelineResultsHTML string

//go:embed ui/static-findings-app/dist/mcp-app.html
var staticFindingsHTML string

//go:embed ui/dynamic-findings-app/dist/mcp-app.html
var dynamicFindingsHTML string

// MCPServer represents the core MCP server
type MCPServer struct {
	initialized      bool
	clientSupportsUI bool
	capabilities     ServerCapabilities
	tools            []Tool
	toolRegistry     *ToolRegistry
	handlerRegistry  *ToolHandlerRegistry
	implRegistry     *ToolImplRegistry
}

func NewMCPServer() (*MCPServer, error) {
	// Create registries
	toolRegistry, err := LoadToolDefinitions()
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
		s.handleInitializeRequest(req, resp)
	case "tools/list":
		resp.Result = s.handleListTools()
	case "tools/call":
		s.handleToolsCallRequest(req, resp)
	case "resources/list":
		log.Printf(">>> resources/list called - returning UI resource")
		resp.Result = s.handleListResources()
	case "resources/read":
		s.handleResourcesReadRequest(req, resp)
	case "notifications/initialized":
		// Client confirms initialization - no response needed for notifications
		log.Println("Client sent initialized notification")
		return nil // Don't send a response for notifications
	default:
		resp.Error = &RPCError{
			Code:    -32601,
			Message: fmt.Sprintf("Method not found: %s", req.Method),
		}
	}

	return resp
}

// handleInitializeRequest processes the initialize request
func (s *MCPServer) handleInitializeRequest(req *JSONRPCRequest, resp *JSONRPCResponse) {
	result, err := s.handleInitialize(req.Params)
	if err != nil {
		resp.Error = &RPCError{
			Code:    -32603,
			Message: err.Error(),
		}
	} else {
		resp.Result = result
	}
}

// handleToolsCallRequest processes the tools/call request
func (s *MCPServer) handleToolsCallRequest(req *JSONRPCRequest, resp *JSONRPCResponse) {
	log.Printf(">>> tools/call invoked (UI support: %v)", s.clientSupportsUI)

	// Parse the tool name for better logging
	var callParams CallToolParams
	if err := json.Unmarshal(req.Params, &callParams); err == nil {
		log.Printf(">>> Calling tool: %s", callParams.Name)
	}

	result, err := s.handleCallTool(req.Params)
	if err != nil {
		resp.Error = &RPCError{
			Code:    -32603,
			Message: err.Error(),
		}
		log.Printf(">>> tools/call ERROR: %v", err)
	} else {
		s.logToolCallResult(result)
		resp.Result = result
	}
}

// logToolCallResult logs details about the tool call result
func (s *MCPServer) logToolCallResult(result *CallToolResult) {
	if result != nil {
		contentCount := len(result.Content)
		hasStructured := result.StructuredContent != nil
		contentLen := 0
		if contentCount > 0 && result.Content[0].Text != "" {
			contentLen = len(result.Content[0].Text)
		}
		log.Printf(">>> tools/call completed: content items=%d, content[0] length=%d chars, hasStructuredContent=%v",
			contentCount, contentLen, hasStructured)

		// Log a preview of the content for debugging
		if contentCount > 0 && contentLen > 0 {
			previewLen := 200
			if contentLen < previewLen {
				previewLen = contentLen
			}
			log.Printf(">>> Content preview (first %d chars): %s...", previewLen, result.Content[0].Text[:previewLen])
		}
	} else {
		log.Printf(">>> tools/call completed: result is nil")
	}
}

// handleResourcesReadRequest processes the resources/read request
func (s *MCPServer) handleResourcesReadRequest(req *JSONRPCRequest, resp *JSONRPCResponse) {
	log.Printf(">>> resources/read called - this should load the UI HTML!")
	result, err := s.handleReadResource(req.Params)
	if err != nil {
		resp.Error = &RPCError{
			Code:    -32603,
			Message: err.Error(),
		}
	} else {
		resp.Result = result
	}
}

func (s *MCPServer) handleInitialize(params json.RawMessage) (*InitializeResult, error) {
	var initParams InitializeParams
	if err := json.Unmarshal(params, &initParams); err != nil {
		return nil, fmt.Errorf("invalid initialize params: %w", err)
	}

	s.initialized = true
	log.Printf("=== CLIENT INITIALIZATION ===")
	log.Printf("Client Name: %s", initParams.ClientInfo.Name)
	log.Printf("Client Version: %s", initParams.ClientInfo.Version)
	log.Printf("Protocol Version: %s", initParams.ProtocolVersion)

	// Log full capabilities structure
	capsJSON, _ := json.MarshalIndent(initParams.Capabilities, "", "  ")
	log.Printf("Client Capabilities (raw):\n%s", string(capsJSON))

	// Check for MCP Apps UI support
	log.Printf("\n=== CHECKING FOR UI SUPPORT ===")
	log.Printf("Extensions field present: %v", initParams.Capabilities.Extensions != nil)

	if extensions := initParams.Capabilities.Extensions; extensions != nil {
		extensionsJSON, _ := json.MarshalIndent(extensions, "", "  ")
		log.Printf("Extensions object:\n%s", string(extensionsJSON))

		// Check all keys in extensions
		log.Printf("Extension keys found:")
		for key := range extensions {
			log.Printf("  - %s", key)
		}

		if uiExt, ok := extensions["io.modelcontextprotocol/ui"].(map[string]interface{}); ok {
			log.Printf("✓ Found 'io.modelcontextprotocol/ui' extension")
			uiExtJSON, _ := json.MarshalIndent(uiExt, "", "  ")
			log.Printf("UI Extension content:\n%s", string(uiExtJSON))

			if mimeTypes, ok := uiExt["mimeTypes"].([]interface{}); ok {
				log.Printf("MimeTypes found: %v", mimeTypes)
				for i, mt := range mimeTypes {
					log.Printf("  [%d] %v (type: %T)", i, mt, mt)
					if mtStr, ok := mt.(string); ok && mtStr == "text/html;profile=mcp-app" {
						s.clientSupportsUI = true
						log.Printf("✓✓✓ Client supports MCP Apps UI! ✓✓✓")
						break
					}
				}
			} else {
				log.Printf("⚠ 'mimeTypes' field not found or not an array in UI extension")
			}
		} else {
			log.Printf("⚠ 'io.modelcontextprotocol/ui' key not found in extensions")
		}
	} else {
		log.Printf("⚠ No 'extensions' field in client capabilities")
	}

	if !s.clientSupportsUI {
		log.Printf("\n❌ Client does NOT support MCP Apps UI (will use text-only mode)")
		log.Printf("   This means full JSON will be sent to the LLM instead of brief summaries")
	} else {
		log.Printf("\n✓ UI mode enabled - will send brief summaries to LLM")
	}
	log.Printf("=== END UI SUPPORT CHECK ===\n")

	// Use the protocol version requested by the client if we support it
	protocolVersion := "2024-11-05"
	if initParams.ProtocolVersion == "2024-11-05" || initParams.ProtocolVersion >= "2024-11-05" {
		protocolVersion = initParams.ProtocolVersion
	}

	return &InitializeResult{
		ProtocolVersion: protocolVersion,
		Capabilities:    s.capabilities,
		ServerInfo: Implementation{
			Name:    "veracode-mcp-server",
			Version: version,
		},
	}, nil
}

func (s *MCPServer) handleListTools() *ListToolsResult {
	result := &ListToolsResult{
		Tools: s.tools,
	}

	// Debug: log metadata for UI tools
	for _, tool := range result.Tools {
		if tool.Name == "pipeline-results" || tool.Name == "get-static-findings" {
			if tool.Meta != nil {
				flatUri, _ := tool.Meta["ui/resourceUri"].(string)
				nestedUI, _ := tool.Meta["ui"].(map[string]interface{})
				var nestedUri string
				if nestedUI != nil {
					nestedUri, _ = nestedUI["resourceUri"].(string)
				}
				log.Printf("%s tool has UI metadata: flat='%s', nested='%s', full=%+v",
					tool.Name, flatUri, nestedUri, tool.Meta)
			} else {
				log.Printf("WARNING: %s tool has NO UI metadata!", tool.Name)
			}
		}
	}

	return result
}

// ClientSupportsUI returns whether the current client supports MCP Apps UI
func (s *MCPServer) ClientSupportsUI() bool {
	return s.clientSupportsUI
}

func (s *MCPServer) handleCallTool(params json.RawMessage) (*CallToolResult, error) {
	var callParams CallToolParams
	if err := json.Unmarshal(params, &callParams); err != nil {
		return nil, fmt.Errorf("invalid tool call params: %w", err)
	}

	// Look up the handler in the registry
	log.Printf("Looking for handler for tool: %s", callParams.Name)
	handler, exists := s.handlerRegistry.GetHandler(callParams.Name)
	if !exists {
		log.Printf("Handler not found! Registered handlers: %+v", s.handlerRegistry)
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

	// Call the handler with context containing UI capability
	ctx := WithUICapability(context.Background(), s.clientSupportsUI)
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

	// If it's a map, try to convert it
	if resultMap, ok := result.(map[string]interface{}); ok {
		return convertMapToCallToolResult(resultMap)
	}

	// Default: convert result to JSON string
	return marshalResultAsJSON(result)
}

// convertMapToCallToolResult handles map[string]interface{} results
func convertMapToCallToolResult(resultMap map[string]interface{}) *CallToolResult {
	// Check for error field
	if errMsg, hasErr := resultMap["error"]; hasErr {
		return &CallToolResult{
			Content: []Content{{Type: "text", Text: fmt.Sprintf("%v", errMsg)}},
			IsError: true,
		}
	}

	result := &CallToolResult{}

	// Check for content field
	if content, hasContent := resultMap["content"]; hasContent {
		if contents := convertContentField(content); contents != nil {
			result.Content = contents
		}
	}

	// Check for structuredContent field (for MCP Apps)
	if structuredContent, hasStructured := resultMap["structuredContent"]; hasStructured {
		// Try direct assignment first (if already map[string]interface{})
		if sc, ok := structuredContent.(map[string]interface{}); ok {
			result.StructuredContent = sc
		} else {
			// Convert struct to map[string]interface{} via JSON
			scJSON, err := json.Marshal(structuredContent)
			if err == nil {
				var scMap map[string]interface{}
				if err := json.Unmarshal(scJSON, &scMap); err == nil {
					result.StructuredContent = scMap
				}
			}
		}
	}

	// If we have content or structuredContent, return the result
	if len(result.Content) > 0 || result.StructuredContent != nil {
		return result
	}

	// If we have text field, use it
	if text, ok := resultMap["text"].(string); ok {
		return &CallToolResult{
			Content: []Content{{Type: "text", Text: text}},
		}
	}

	// Fallback to JSON
	return marshalResultAsJSON(resultMap)
}

// convertContentField converts various content field formats to []Content
func convertContentField(content interface{}) []Content {
	// Try []map[string]interface{} first (for resources)
	if contentList, ok := content.([]map[string]interface{}); ok {
		return convertDetailedContentList(contentList)
	}

	// Fallback to []map[string]string (for simple text content)
	if contentList, ok := content.([]map[string]string); ok {
		return convertSimpleContentList(contentList)
	}

	return nil
}

// convertDetailedContentList converts []map[string]interface{} to []Content
func convertDetailedContentList(contentList []map[string]interface{}) []Content {
	contents := make([]Content, len(contentList))
	for i, c := range contentList {
		cont := Content{}
		if typ, ok := c["type"].(string); ok {
			cont.Type = typ
		}
		if text, ok := c["text"].(string); ok {
			cont.Text = text
		}
		// Handle resource field
		if resource, ok := c["resource"].(map[string]interface{}); ok {
			cont.Resource = convertResourceField(resource)
		}
		contents[i] = cont
	}
	return contents
}

// convertResourceField converts a resource map to ResourceContents
func convertResourceField(resource map[string]interface{}) *ResourceContents {
	rc := &ResourceContents{}
	if uri, ok := resource["uri"].(string); ok {
		rc.URI = uri
	}
	if mimeType, ok := resource["mimeType"].(string); ok {
		rc.MimeType = mimeType
	}
	if text, ok := resource["text"].(string); ok {
		rc.Text = text
	}
	return rc
}

// convertSimpleContentList converts []map[string]string to []Content
func convertSimpleContentList(contentList []map[string]string) []Content {
	contents := make([]Content, len(contentList))
	for i, c := range contentList {
		contents[i] = Content{
			Type: c["type"],
			Text: c["text"],
		}
	}
	return contents
}

// marshalResultAsJSON converts any result to JSON string format
func marshalResultAsJSON(result interface{}) *CallToolResult {
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		log.Printf("Failed to marshal result: %v", err)
		return &CallToolResult{
			Content: []Content{{Type: "text", Text: fmt.Sprintf("Error: %v", err)}},
			IsError: true,
		}
	}
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
			{
				URI:         "ui://pipeline-results/app.html",
				Name:        "Pipeline Results UI",
				Description: "Interactive UI for pipeline scan results",
				MimeType:    "text/html;profile=mcp-app",
			},
			{
				URI:         "ui://static-findings/app.html",
				Name:        "Static Findings UI",
				Description: "Interactive UI for static analysis findings",
				MimeType:    "text/html;profile=mcp-app",
			},
			{
				URI:         "ui://dynamic-findings/app.html",
				Name:        "Dynamic Findings UI",
				Description: "Interactive UI for dynamic analysis findings",
				MimeType:    "text/html;profile=mcp-app",
			},
		},
	}
}

func (s *MCPServer) handleReadResource(params json.RawMessage) (*ReadResourceResult, error) {
	var readParams ReadResourceParams
	if err := json.Unmarshal(params, &readParams); err != nil {
		return nil, fmt.Errorf("invalid read resource params: %w", err)
	}

	log.Printf("Reading resource: %s", readParams.URI)

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

	case "ui://pipeline-results/app.html":
		// Return the embedded UI HTML with metadata
		return &ReadResourceResult{
			Contents: []ResourceContents{
				{
					URI:      readParams.URI,
					MimeType: "text/html;profile=mcp-app",
					Text:     pipelineResultsHTML,
					Meta: map[string]interface{}{
						"ui": map[string]interface{}{
							"permissions": map[string]interface{}{},
						},
					},
				},
			},
		}, nil

	case "ui://static-findings/app.html":
		// Return the embedded UI HTML with metadata
		log.Printf("Serving static findings UI - HTML length: %d bytes", len(staticFindingsHTML))
		return &ReadResourceResult{
			Contents: []ResourceContents{
				{
					URI:      readParams.URI,
					MimeType: "text/html;profile=mcp-app",
					Text:     staticFindingsHTML,
					Meta: map[string]interface{}{
						"ui": map[string]interface{}{
							"permissions": map[string]interface{}{},
						},
					},
				},
			},
		}, nil

	case "ui://dynamic-findings/app.html":
		// Return the embedded UI HTML with metadata
		log.Printf("Serving dynamic findings UI - HTML length: %d bytes", len(dynamicFindingsHTML))
		return &ReadResourceResult{
			Contents: []ResourceContents{
				{
					URI:      readParams.URI,
					MimeType: "text/html;profile=mcp-app",
					Text:     dynamicFindingsHTML,
					Meta: map[string]interface{}{
						"ui": map[string]interface{}{
							"permissions": map[string]interface{}{},
						},
					},
				},
			},
		}, nil

	default:
		return nil, fmt.Errorf("resource not found: %s", readParams.URI)
	}
}
