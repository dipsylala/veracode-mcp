package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tools "github.com/dipsylala/veracode-mcp/internal/tool_registry"
	"github.com/dipsylala/veracode-mcp/internal/types"
)

// MCP protocol method handlers that process specific JSON-RPC requests
// and coordinate with the appropriate business logic.

// handleInitializeRequest processes the MCP initialize handshake.
// This is critical for detecting client UI capabilities, which affects
// whether we return full JSON data or brief summaries for UI-enabled clients.
func (s *MCPServer) handleInitializeRequest(req *types.JSONRPCRequest, resp *types.JSONRPCResponse) {
	// Convert params to json.RawMessage for processing
	var paramsRaw json.RawMessage
	if req.Params != nil {
		if paramsBytes, err := json.Marshal(req.Params); err == nil {
			paramsRaw = paramsBytes
		}
	}

	result, err := s.handleInitialize(paramsRaw)
	if err != nil {
		resp.Error = &types.RPCError{
			Code:    -32603,
			Message: err.Error(),
		}
	} else {
		resp.Result = result
	}
}

// handleToolsCallRequest processes tool execution requests.
// It validates parameters, looks up handlers, and coordinates tool execution
// with proper error handling and logging.
func (s *MCPServer) handleToolsCallRequest(req *types.JSONRPCRequest, resp *types.JSONRPCResponse) {
	log.Printf(">>> tools/call invoked (UI support: %v)", s.clientSupportsUI)

	// Convert params to json.RawMessage for processing
	var paramsRaw json.RawMessage
	if req.Params != nil {
		if paramsBytes, err := json.Marshal(req.Params); err == nil {
			paramsRaw = paramsBytes

			// Parse the tool name for better logging
			var callParams types.CallToolParams
			if err := json.Unmarshal(paramsRaw, &callParams); err == nil {
				log.Printf(">>> Calling tool: %s", callParams.Name)
			}
		}
	}

	result, err := s.handleCallTool(paramsRaw)
	if err != nil {
		resp.Error = &types.RPCError{
			Code:    -32603,
			Message: err.Error(),
		}
		log.Printf(">>> tools/call ERROR: %v", err)
	} else {
		s.logToolCallResult(result)
		resp.Result = result
	}
}

// handleResourcesReadRequest processes resource read requests.
// This is used for serving UI applications and other static resources.
func (s *MCPServer) handleResourcesReadRequest(req *types.JSONRPCRequest, resp *types.JSONRPCResponse) {
	log.Printf("\n╔════════════════════════════════════════╗")
	log.Printf("║    🎨 RESOURCES/READ CALLED! 🎨       ║")
	log.Printf("╚════════════════════════════════════════╝")

	// Convert params to json.RawMessage for processing
	var paramsRaw json.RawMessage
	if req.Params != nil {
		if paramsBytes, err := json.Marshal(req.Params); err == nil {
			paramsRaw = paramsBytes
			log.Printf("[RESOURCES/READ] Request params: %s", string(paramsBytes))
		}
	}

	result, err := s.handleReadResource(paramsRaw)
	if err != nil {
		log.Printf("[RESOURCES/READ] ❌ ERROR: %v", err)
		resp.Error = &types.RPCError{
			Code:    -32603,
			Message: err.Error(),
		}
	} else {
		if result != nil && len(result.Contents) > 0 {
			log.Printf("[RESOURCES/READ] ✅ SUCCESS: Serving %d content items", len(result.Contents))
			for _, content := range result.Contents {
				log.Printf("[RESOURCES/READ]   - URI: %s, MimeType: %s, Size: %d bytes",
					content.URI, content.MimeType, len(content.Text))
			}
		}
		resp.Result = result
	}
	log.Printf("╚════════════════════════════════════════╝\n")
}

// handleInitialize processes the initialize request and establishes
// the protocol version and capabilities between client and server.
func (s *MCPServer) handleInitialize(params json.RawMessage) (*InitializeResult, error) {
	initParams, err := s.parseInitParams(params)
	if err != nil {
		return nil, err
	}

	s.initialized = true
	s.logClientInfo(initParams)
	s.clientSupportsUI = s.detectUICapability(initParams.Capabilities)

	return s.buildInitializeResult(initParams), nil
}

// parseInitParams validates and parses the initialization parameters from the client.
func (s *MCPServer) parseInitParams(params json.RawMessage) (*InitializeParams, error) {
	var initParams InitializeParams
	if err := json.Unmarshal(params, &initParams); err != nil {
		return nil, fmt.Errorf("invalid initialize params: %w", err)
	}
	return &initParams, nil
}

// logClientInfo logs detailed information about the connecting client for debugging.
func (s *MCPServer) logClientInfo(initParams *InitializeParams) {
	log.Printf("\n╔════════════════════════════════════════╗")
	log.Printf("║      CLIENT INITIALIZATION              ║")
	log.Printf("╚════════════════════════════════════════╝")
	log.Printf("Client Name: %s", initParams.ClientInfo.Name)
	log.Printf("Client Version: %s", initParams.ClientInfo.Version)
	log.Printf("Protocol Version: %s", initParams.ProtocolVersion)

	// Log full capabilities structure for debugging
	capsJSON, _ := json.MarshalIndent(initParams.Capabilities, "", "  ")
	log.Printf("Client Capabilities (raw):\n%s", string(capsJSON))
	log.Printf("╚════════════════════════════════════════╝\n")
}

// buildInitializeResult constructs the response for the initialize request
// with negotiated protocol version and server capabilities.
func (s *MCPServer) buildInitializeResult(initParams *InitializeParams) *InitializeResult {
	// Use the protocol version requested by the client if we support it
	protocolVersion := MCPProtocolVersion
	if initParams.ProtocolVersion == MCPProtocolVersion || initParams.ProtocolVersion >= MCPProtocolVersion {
		protocolVersion = initParams.ProtocolVersion
	}

	// Note: version should be passed from main or configured elsewhere
	serverVersion := "dev" // Default version, should be set by build process

	return &InitializeResult{
		ProtocolVersion: protocolVersion,
		Capabilities:    s.capabilities,
		ServerInfo: Implementation{
			Name:    "veracode-mcp-server",
			Version: serverVersion,
		},
		Instructions: embeddedInstructions,
	}
}

// handleListTools returns the list of available tools with their schemas.
// Conditionally adds UI metadata based on client capabilities.
func (s *MCPServer) handleListTools() *types.ListToolsResult {
	// Create a copy of tools with conditional UI metadata
	toolsList := make([]types.Tool, len(s.tools))
	copy(toolsList, s.tools)

	// Add UI metadata only if client supports it
	log.Printf("[TOOLS] s.clientSupportsUI=%v", s.clientSupportsUI)
	if s.clientSupportsUI {
		log.Printf("[TOOLS] Client supports UI - adding UI metadata to tool definitions")
		for i := range toolsList {
			if uiMeta := tools.GetUIMetaForTool(toolsList[i].Name); uiMeta != nil {
				toolsList[i].Meta = uiMeta
				log.Printf("[TOOLS] Added UI metadata to %s: %+v", toolsList[i].Name, uiMeta)
			}
		}
	} else {
		log.Printf("[TOOLS] Client does NOT support UI - omitting UI metadata from tool definitions")
	}

	result := &types.ListToolsResult{
		Tools: toolsList,
	}

	// Debug: log metadata for UI tools
	for _, tool := range result.Tools {
		if tool.Name == "pipeline-findings" || tool.Name == "static-findings" || tool.Name == "dynamic-findings" {
			if len(tool.Meta) > 0 {
				flatUri, _ := tool.Meta["ui/resourceUri"].(string)
				nestedUI, _ := tool.Meta["ui"].(map[string]interface{})
				var nestedUri string
				if nestedUI != nil {
					nestedUri, _ = nestedUI["resourceUri"].(string)
				}
				log.Printf("[TOOLS/LIST] %s tool has UI metadata: flat='%s', nested='%s'",
					tool.Name, flatUri, nestedUri)

				// Log the full tool as JSON to see what's actually being sent
				toolJSON, _ := json.MarshalIndent(tool, "  ", "  ")
				log.Printf("[TOOLS/LIST] Full %s tool JSON:\n%s", tool.Name, string(toolJSON))
			} else {
				log.Printf("[TOOLS/LIST] WARNING: %s tool has NO UI metadata!", tool.Name)
			}
		}
	}

	return result
}

// handleCallTool executes a tool by name with the provided arguments.
// It performs validation, looks up handlers, and manages the execution context.
func (s *MCPServer) handleCallTool(params json.RawMessage) (*types.CallToolResult, error) {
	callParams, err := s.parseToolCallParams(params)
	if err != nil {
		return nil, err
	}

	handler, err := s.lookupToolHandler(callParams.Name)
	if err != nil {
		return s.createToolError(err.Error()), nil
	}

	if err := s.validateToolCall(callParams); err != nil {
		return s.createToolError(err.Error()), nil
	}

	return s.executeToolCall(handler, callParams.Arguments)
}

// parseToolCallParams validates and parses the tool call parameters.
func (s *MCPServer) parseToolCallParams(params json.RawMessage) (*types.CallToolParams, error) {
	var callParams types.CallToolParams
	if err := json.Unmarshal(params, &callParams); err != nil {
		return nil, fmt.Errorf("invalid tool call params: %w", err)
	}
	return &callParams, nil
}

// lookupToolHandler finds the handler function for the specified tool name.
func (s *MCPServer) lookupToolHandler(toolName string) (func(context.Context, map[string]interface{}) (interface{}, error), error) {
	log.Printf("Looking for handler for tool: %s", toolName)
	handler, exists := s.toolManager.GetToolHandler(toolName)
	if !exists {
		log.Printf("Handler not found! Available tools: %v", s.toolManager.GetAvailableToolNames())
		return nil, fmt.Errorf("Unknown tool: %s. Available tools: %s", toolName, s.getAvailableToolNames())
	}
	return handler, nil
}

// validateToolCall validates the tool call arguments against the tool definition.
func (s *MCPServer) validateToolCall(callParams *types.CallToolParams) error {
	return s.toolManager.ValidateToolArguments(callParams.Name, callParams.Arguments)
}

// executeToolCall runs the tool handler with proper context and error handling.
func (s *MCPServer) executeToolCall(handler func(context.Context, map[string]interface{}) (interface{}, error), arguments map[string]interface{}) (*types.CallToolResult, error) {
	// Call the handler with context containing UI capability
	ctx := WithUICapability(context.Background(), s.clientSupportsUI)
	result, err := handler(ctx, arguments)
	if err != nil {
		return &types.CallToolResult{
			Content: []types.Content{{Type: "text", Text: fmt.Sprintf("Tool execution error: %v", err)}},
			IsError: true,
		}, nil
	}

	// Convert the result to CallToolResult format
	converted := tools.ConvertToCallToolResult(result)
	log.Printf("[HANDLER] After conversion: hasStructuredContent=%v", converted.StructuredContent != nil)
	if converted.StructuredContent != nil {
		log.Printf("[HANDLER] StructuredContent type=%T", converted.StructuredContent)
	}
	return converted, nil
}

// createToolError creates a standardized error response for tool calls.
func (s *MCPServer) createToolError(message string) *types.CallToolResult {
	return &types.CallToolResult{
		Content: []types.Content{{Type: "text", Text: message}},
		IsError: true,
	}
}

// handleListResources returns available resources including UI applications.
// Resources provide static content like HTML UIs for interactive tools.
func (s *MCPServer) handleListResources() *ListResourcesResult {
	result := &ListResourcesResult{
		Resources: []Resource{
			{
				URI:         "ui://pipeline-findings/app.html",
				Name:        "Pipeline Findings UI",
				Description: "Interactive UI for pipeline scan findings",
				MimeType:    UICapabilityMimeType,
			},
			{
				URI:         "ui://static-findings/app.html",
				Name:        "Static Findings UI",
				Description: "Interactive UI for static analysis findings",
				MimeType:    UICapabilityMimeType,
			},
			{
				URI:         "ui://dynamic-findings/app.html",
				Name:        "Dynamic Findings UI",
				Description: "Interactive UI for dynamic analysis findings",
				MimeType:    UICapabilityMimeType,
			},
			{
				URI:         "ui://local-sca-findings/app.html",
				Name:        "Local SCA Findings UI",
				Description: "Interactive UI for local SCA scan findings",
				MimeType:    UICapabilityMimeType,
			},
		},
	}

	log.Printf("[RESOURCES/LIST] Returning %d resources", len(result.Resources))
	for _, res := range result.Resources {
		if strings.HasPrefix(res.URI, "ui://") {
			log.Printf("[RESOURCES/LIST] UI Resource: %s (%s)", res.URI, res.MimeType)
		}
	}

	return result
}

// handleReadResource serves the content for a specific resource URI.
// This includes both static content and embedded UI applications.
func (s *MCPServer) handleReadResource(params json.RawMessage) (*ReadResourceResult, error) {
	var readParams ReadResourceParams
	if err := json.Unmarshal(params, &readParams); err != nil {
		return nil, fmt.Errorf("invalid read resource params: %w", err)
	}

	log.Printf("[RESOURCES/READ] Requested URI: %s", readParams.URI)

	switch readParams.URI {
	case "ui://pipeline-findings/app.html":
		log.Printf("[RESOURCES/READ] 🎯 Serving Pipeline Results UI - HTML length: %d bytes", len(embeddedPipelineFindingsHTML))
		if len(embeddedPipelineFindingsHTML) < 1000 {
			log.Printf("[RESOURCES/READ] ⚠️  WARNING: HTML is suspiciously small! May not be built correctly.")
		}
		return s.serveUIResource(readParams.URI, embeddedPipelineFindingsHTML)

	case "ui://static-findings/app.html":
		log.Printf("[RESOURCES/READ] 🎯 Serving Static Findings UI - HTML length: %d bytes", len(embeddedStaticFindingsHTML))
		if len(embeddedStaticFindingsHTML) < 1000 {
			log.Printf("[RESOURCES/READ] ⚠️  WARNING: HTML is suspiciously small! May not be built correctly.")
		}
		return s.serveUIResource(readParams.URI, embeddedStaticFindingsHTML)

	case "ui://dynamic-findings/app.html":
		log.Printf("[RESOURCES/READ] 🎯 Serving Dynamic Findings UI - HTML length: %d bytes", len(embeddedDynamicFindingsHTML))
		if len(embeddedDynamicFindingsHTML) < 1000 {
			log.Printf("[RESOURCES/READ] ⚠️  WARNING: HTML is suspiciously small! May not be built correctly.")
		}
		return s.serveUIResource(readParams.URI, embeddedDynamicFindingsHTML)

	case "ui://local-sca-findings/app.html":
		log.Printf("[RESOURCES/READ] 🎯 Serving Local SCA Findings UI - HTML length: %d bytes", len(embeddedLocalSCAFindingsHTML))
		if len(embeddedLocalSCAFindingsHTML) < 1000 {
			log.Printf("[RESOURCES/READ] ⚠️  WARNING: HTML is suspiciously small! May not be built correctly.")
		}
		return s.serveUIResource(readParams.URI, embeddedLocalSCAFindingsHTML)

	default:
		log.Printf("[RESOURCES/READ] ❌ Resource not found: %s", readParams.URI)
		log.Printf("[RESOURCES/READ] Available URIs: ui://pipeline-findings/app.html, ui://static-findings/app.html, ui://dynamic-findings/app.html, ui://local-sca-findings/app.html")
		return nil, fmt.Errorf("resource not found: %s", readParams.URI)
	}
}

// serveUIResource creates a UI resource response with proper MCP Apps metadata.
// This helper ensures consistent UI resource serving across different applications.
func (s *MCPServer) serveUIResource(uri, htmlContent string) (*ReadResourceResult, error) {
	return &ReadResourceResult{
		Contents: []ResourceContents{
			{
				URI:      uri,
				MimeType: UICapabilityMimeType,
				Text:     htmlContent,
				Meta: map[string]interface{}{
					"ui": map[string]interface{}{
						"permissions": map[string]interface{}{},
					},
				},
			},
		},
	}, nil
}

// logToolCallResult provides detailed logging for tool execution results
// including content analysis and debugging information.
func (s *MCPServer) logToolCallResult(result *types.CallToolResult) {
	if result != nil {
		contentCount := len(result.Content)
		hasStructured := result.Meta != nil
		contentLen := 0
		if contentCount > 0 && result.Content[0].Text != "" {
			contentLen = len(result.Content[0].Text)
		}
		log.Printf(">>> tools/call completed: content items=%d, content[0] length=%d chars, hasMeta=%v",
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
