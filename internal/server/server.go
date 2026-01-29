// Package main implements a Model Context Protocol (MCP) server for Veracode security tools.
// It supports both stdio and HTTP transports and provides auto-registered tool discovery.
package server

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/dipsylala/veracodemcp-go/internal/tools"
	"github.com/dipsylala/veracodemcp-go/internal/transport"
	"github.com/dipsylala/veracodemcp-go/internal/types"
)

// Context key for UI capability
type contextKey string

const uiCapabilityKey contextKey = "uiCapability"

// MCP protocol constants
const (
	MCPProtocolVersion   = "2024-11-05"
	UICapabilityMimeType = "text/html;profile=mcp-app"
	UIExtensionKey       = "io.modelcontextprotocol/ui"
)

// WithUICapability adds UI capability information to the context.
// This allows tools to adapt their output based on client UI support.
func WithUICapability(ctx context.Context, supportsUI bool) context.Context {
	return context.WithValue(ctx, uiCapabilityKey, supportsUI)
}

// ClientSupportsUIFromContext retrieves UI capability from context.
// Returns true if the client supports MCP Apps UI (text/html;profile=mcp-app).
func ClientSupportsUIFromContext(ctx context.Context) bool {
	if val := ctx.Value(uiCapabilityKey); val != nil {
		if supportsUI, ok := val.(bool); ok {
			return supportsUI
		}
	}
	return false
}

var embeddedPipelineResultsHTML string
var embeddedStaticFindingsHTML string
var embeddedDynamicFindingsHTML string

// SetUIResources sets the embedded UI resources from the main package
func SetUIResources(pipeline, staticFindings, dynamicFindings string) {
	embeddedPipelineResultsHTML = pipeline
	embeddedStaticFindingsHTML = staticFindings
	embeddedDynamicFindingsHTML = dynamicFindings
}

// MCPServer represents the core MCP server that handles protocol communication
// and coordinates between tool registries, handlers, and transport layers.
type MCPServer struct {
	initialized      bool
	clientSupportsUI bool
	capabilities     ServerCapabilities
	tools            []types.Tool
	toolManager      *tools.ToolManager
}

// NewMCPServer creates a new MCP server instance with all necessary registries.
// It loads tool definitions, initializes registries, and prepares tool handlers.
func NewMCPServer() (*MCPServer, error) {
	// Create unified tool manager
	toolManager, err := tools.NewToolManager()
	if err != nil {
		return nil, fmt.Errorf("failed to create tool manager: %w", err)
	}

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
		toolManager: toolManager,
	}

	// Initialize tool implementations and register their handlers
	if err := s.toolManager.LoadAllTools(); err != nil {
		return nil, fmt.Errorf("failed to load tool implementations: %w", err)
	}

	// Convert tool definitions to MCP tools for tools/list
	s.tools = s.toolManager.GetAllMCPTools()

	return s, nil
}

// ServeStdio starts the MCP server using stdio transport.
// This is the standard mode for MCP client integrations.
func (s *MCPServer) ServeStdio() error {
	t := transport.NewStdioTransport(s)
	return t.Start()
}

// ServeHTTP starts the MCP server using HTTP transport with Server-Sent Events.
// This allows remote connections and web-based integrations.
func (s *MCPServer) ServeHTTP(addr string) error {
	t := transport.NewHTTPTransport(s)
	return t.Start(addr)
}

// HandleRequest processes incoming MCP protocol requests and routes them
// to the appropriate handler methods. This is the main request dispatcher.
func (s *MCPServer) HandleRequest(req *types.JSONRPCRequest) *types.JSONRPCResponse {
	log.Printf("Handling request: %s (id: %v)", req.Method, req.ID)

	resp := &types.JSONRPCResponse{
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
		resp.Error = &types.RPCError{
			Code:    -32601,
			Message: fmt.Sprintf("Method not found: %s", req.Method),
		}
	}

	return resp
}

// ClientSupportsUI returns whether the current client supports MCP Apps UI.
// This information is determined during the initialize handshake.
func (s *MCPServer) ClientSupportsUI() bool {
	return s.clientSupportsUI
}

// Shutdown gracefully shuts down the MCP server and all tool implementations.
// This should be called when the server is terminating to ensure proper cleanup.
func (s *MCPServer) Shutdown() {
	if s.toolManager != nil {
		s.toolManager.Shutdown()
	}
}

// GetToolStats returns statistics about the loaded tools for monitoring and debugging.
// This provides insight into the tool manager's state and can be useful for diagnostics.
func (s *MCPServer) GetToolStats() tools.ToolManagerStats {
	if s.toolManager != nil {
		return s.toolManager.GetStats()
	}
	return tools.ToolManagerStats{}
}

// getAvailableToolNames returns a formatted list of available tool names
// for error messages and debugging.
func (s *MCPServer) getAvailableToolNames() string {
	names := s.toolManager.GetAvailableToolNames()
	return fmt.Sprintf("[%s]", fmt.Sprint(names))
}
