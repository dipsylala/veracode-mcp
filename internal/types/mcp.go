package types

// Tool represents a tool available through the MCP protocol.
// This structure defines how tools are presented to MCP clients.
type Tool struct {
	Name        string      `json:"name"`           // Unique identifier for the tool
	Description string      `json:"description"`    // Human-readable description
	InputSchema interface{} `json:"inputSchema"`    // JSON Schema for tool parameters
	Meta        interface{} `json:"meta,omitempty"` // Optional metadata (UI hints, etc.)
}

// JSONRPCRequest represents an incoming JSON-RPC 2.0 request.
// This is the standard format for MCP protocol messages.
type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`          // Always "2.0" for JSON-RPC 2.0
	ID      interface{} `json:"id"`               // Request identifier for matching responses
	Method  string      `json:"method"`           // MCP method name (e.g., "tools/call")
	Params  interface{} `json:"params,omitempty"` // Method-specific parameters
}

// JSONRPCResponse represents an outgoing JSON-RPC 2.0 response.
// Contains either a result or an error, never both.
type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`          // Always "2.0"
	ID      interface{} `json:"id"`               // Matches the request ID
	Result  interface{} `json:"result,omitempty"` // Success result
	Error   *RPCError   `json:"error,omitempty"`  // Error information
}

// RPCError represents a JSON-RPC 2.0 error object.
// Used when a request cannot be processed successfully.
type RPCError struct {
	Code    int         `json:"code"`           // Standard JSON-RPC error code
	Message string      `json:"message"`        // Human-readable error message
	Data    interface{} `json:"data,omitempty"` // Additional error details
}

// ListToolsResult is the response format for the tools/list method.
// Returns all tools available on this MCP server.
type ListToolsResult struct {
	Tools []Tool `json:"tools"`
}

// CallToolParams represents the parameters for a tools/call request.
// Specifies which tool to invoke and what arguments to pass.
type CallToolParams struct {
	Name      string                 `json:"name"`      // Tool name to invoke
	Arguments map[string]interface{} `json:"arguments"` // Tool-specific parameters
}

// CallToolResult represents the response from a tools/call request.
// Contains the tool's output in a structured format.
type CallToolResult struct {
	Content []Content   `json:"content"`           // Tool output content
	IsError bool        `json:"isError,omitempty"` // Whether this represents an error
	Meta    interface{} `json:"meta,omitempty"`    // Additional metadata
}

// Content represents a piece of content in MCP responses.
// Supports different content types (text, binary data, etc.).
type Content struct {
	Type     string `json:"type"`               // Content type ("text", "image", etc.)
	Text     string `json:"text,omitempty"`     // Text content
	MimeType string `json:"mimeType,omitempty"` // MIME type for binary content
	Data     string `json:"data,omitempty"`     // Base64-encoded binary data
}
