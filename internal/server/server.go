// Package server implements a Model Context Protocol (MCP) server for Veracode security tools.
// It uses stdio transport for local filesystem operations and provides auto-registered tool discovery.
package server

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dipsylala/veracodemcp-go/internal/tools"
	"github.com/dipsylala/veracodemcp-go/internal/transport"
	"github.com/dipsylala/veracodemcp-go/internal/types"
	"github.com/dipsylala/veracodemcp-go/mcp_tools"
)

// Context key for UI capability (using plain string for cross-package compatibility)
const uiCapabilityKey = mcp_tools.UICapabilityKey

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
	forceMCPApp      bool // Override to always send structuredContent
	capabilities     ServerCapabilities
	tools            []types.Tool
	toolManager      *tools.ToolManager
}

// NewMCPServer creates a new MCP server instance with all necessary registries.
// It loads tool definitions, initializes registries, and prepares tool handlers.
func NewMCPServer(forceMCPApp bool) (*MCPServer, error) {
	// Create tool manager
	toolManager, err := tools.NewToolManager()
	if err != nil {
		return nil, fmt.Errorf("failed to create tool manager: %w", err)
	}

	s := &MCPServer{
		forceMCPApp: forceMCPApp,
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
// This is the only supported mode as the server requires local filesystem access.
func (s *MCPServer) ServeStdio() error {
	t := transport.NewStdioTransport(s)
	return t.Start()
}

var (
	ErrInvalidMethodName = errors.New("json-rpc Method Name is invalid")
	ErrInvalidIDType     = errors.New("json-rpc id must be a string or integer")
	ErrInvalidIDString   = errors.New("json-rpc id must be printable ASCII")
	ErrInvalidIDLength   = errors.New("json-rpc id length out of bounds")
)

// isValidMethod checks if a method name contains only alphanumeric characters and forward slashes
// to prevent log forging attacks via CRLF injection
func ValidateMethod(method string) error {
	for _, r := range method {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '/' {
			return ErrInvalidMethodName
		}
	}
	return nil
}

func ValidateID(id any, maxLen int) error {
	switch v := id.(type) {

	case string:
		if len(v) == 0 || len(v) > maxLen {
			return ErrInvalidIDLength
		}

		// Fast path: ASCII-only check
		for i := 0; i < len(v); i++ {
			b := v[i]
			if b < 0x20 || b > 0x7E {
				return ErrInvalidIDString
			}
		}

		// Defensive: ensure no sneaky UTF-8 (should be redundant)
		if !utf8.ValidString(v) {
			return ErrInvalidIDString
		}

		return nil

	case int, int32, int64, uint, uint32, uint64, float64:
		// JSON numbers are unmarshaled as float64 by default
		// Allow integers and floats that represent whole numbers
		return nil

	case nil:
		// "id": null is invalid in requests per JSON-RPC 2.0
		return fmt.Errorf("json-rpc id must not be null")

	default:
		return ErrInvalidIDType
	}
}

// HandleRequest processes incoming MCP protocol requests and routes them
// to the appropriate handler methods. This is the main request dispatcher.
func (s *MCPServer) HandleRequest(req *types.JSONRPCRequest) *types.JSONRPCResponse {

	// Notifications (no ID) don't require responses
	if req.ID == nil {
		// Only allow nil ID for notification methods
		if strings.HasPrefix(req.Method, "notifications/") {
			log.Printf("Handling notification: %s (no response needed)", req.Method)
			return nil
		}
		// Non-notification methods MUST have an ID
		log.Printf("ERROR: Non-notification method '%s' missing ID", req.Method)
		return &types.JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      nil,
			Error: &types.RPCError{
				Code:    -32600,
				Message: fmt.Sprintf("Invalid request: method '%s' requires an id", req.Method),
			},
		}
	}

	// Check if ID is explicitly the JSON null value
	if len(*req.ID) == 4 && string(*req.ID) == "null" {
		return &types.JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      nil,
			Error: &types.RPCError{
				Code:    -32600,
				Message: "Invalid request: id must not be null",
			},
		}
	}

	// Unmarshal the ID to validate it properly
	var idValue interface{}
	if err := json.Unmarshal(*req.ID, &idValue); err != nil {
		log.Printf("Rejecting request with malformed id: %v", err)
		return &types.JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Error: &types.RPCError{
				Code:    -32600,
				Message: "Invalid request: id is malformed",
			},
		}
	}

	errID := ValidateID(idValue, 64)
	if errID != nil {
		log.Printf("Rejecting request with invalid id: %v (error: %v)", idValue, errID)
		return &types.JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Error: &types.RPCError{
				Code:    -32600,
				Message: "Invalid request: id MUST be a non-empty string consisting solely of printable ASCII characters in the range U+0020 through U+007E, with a maximum length of 64 characters.",
			},
		}
	}

	// Validate method name to prevent log forging attacks
	errMethod := ValidateMethod(req.Method)
	if errMethod != nil {
		log.Printf("Rejecting request with invalid id: %v)", req.ID)
		return &types.JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Error: &types.RPCError{
				Code:    -32600,
				Message: "Invalid request: method name contains invalid characters",
			},
		}
	}

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
		resp.Result = nil
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
