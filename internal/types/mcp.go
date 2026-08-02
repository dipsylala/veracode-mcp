package types

import "encoding/json"

// Tool represents a tool available through the MCP protocol.
// This structure defines how tools are presented to MCP clients.
type Tool struct {
	Name        string                 `json:"name"`                // Unique identifier for the tool
	Description string                 `json:"description"`         // Human-readable description
	InputSchema interface{}            `json:"inputSchema"`         // JSON Schema for tool parameters
	Execution   *ToolExecution         `json:"execution,omitempty"` // Task-augmentation support (MCP 2025-11-25 Tasks)
	Meta        map[string]interface{} `json:"_meta,omitempty"`     // Optional metadata (UI hints, etc.) - underscore prefix per MCP spec
}

// ToolExecution describes a tool's support for task-augmented invocation.
type ToolExecution struct {
	// TaskSupport is one of "required", "optional", or "forbidden" (default).
	TaskSupport string `json:"taskSupport,omitempty"`
}

// JSONRPCRequest represents an incoming JSON-RPC 2.0 request.
// This is the standard format for MCP protocol messages.
type JSONRPCRequest struct {
	JSONRPC string           `json:"jsonrpc"`
	ID      *json.RawMessage `json:"id,omitempty"`
	Method  string           `json:"method"`
	Params  json.RawMessage  `json:"params,omitempty"`
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
	Name      string                 `json:"name"`           // Tool name to invoke
	Arguments map[string]interface{} `json:"arguments"`      // Tool-specific parameters
	Task      *TaskParams            `json:"task,omitempty"` // Present when the request is task-augmented (MCP Tasks)
}

// TaskParams carries task-augmentation options included on a task-augmented request.
type TaskParams struct {
	TTL *int64 `json:"ttl,omitempty"` // Requested task lifetime in milliseconds since creation
}

// Task status values, per the MCP Tasks specification (2025-11-25).
const (
	TaskStatusWorking       = "working"
	TaskStatusInputRequired = "input_required"
	TaskStatusCompleted     = "completed"
	TaskStatusFailed        = "failed"
	TaskStatusCancelled     = "cancelled"
)

// Task represents the execution state of a task-augmented request.
type Task struct {
	TaskID        string `json:"taskId"`
	Status        string `json:"status"`
	StatusMessage string `json:"statusMessage,omitempty"`
	CreatedAt     string `json:"createdAt"`     // ISO 8601
	LastUpdatedAt string `json:"lastUpdatedAt"` // ISO 8601
	TTL           *int64 `json:"ttl,omitempty"`
	PollInterval  *int64 `json:"pollInterval,omitempty"`
}

// CreateTaskResult is returned in place of the normal result when a receiver
// accepts a task-augmented request.
type CreateTaskResult struct {
	Task Task `json:"task"`
}

// TaskIDParams is the params shape for tasks/get, tasks/result, and tasks/cancel.
type TaskIDParams struct {
	TaskID string `json:"taskId"`
}

// CallToolResult represents the response from a tools/call request.
// Contains the tool's output in a structured format.
type CallToolResult struct {
	Content           []Content   `json:"content"`                     // Tool output content
	IsError           bool        `json:"isError,omitempty"`           // Whether this represents an error
	Meta              interface{} `json:"meta,omitempty"`              // Additional metadata
	StructuredContent interface{} `json:"structuredContent,omitempty"` // Structured data for MCP Apps UI
}

// Content represents a piece of content in MCP responses.
// Supports different content types (text, binary data, etc.).
type Content struct {
	Type     string `json:"type"`               // Content type ("text", "image", etc.)
	Text     string `json:"text,omitempty"`     // Text content
	MimeType string `json:"mimeType,omitempty"` // MIME type for binary content
	Data     string `json:"data,omitempty"`     // Base64-encoded binary data
}
