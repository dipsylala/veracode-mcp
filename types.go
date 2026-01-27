package main

import "encoding/json"

// JSON-RPC 2.0 types
type JSONRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
}

type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// MCP Protocol types
type InitializeParams struct {
	ProtocolVersion string                 `json:"protocolVersion"`
	Capabilities    ClientCapabilities     `json:"capabilities"`
	ClientInfo      Implementation         `json:"clientInfo"`
	Meta            map[string]interface{} `json:"meta,omitempty"`
}

type ClientCapabilities struct {
	Roots    *RootsCapability    `json:"roots,omitempty"`
	Sampling *SamplingCapability `json:"sampling,omitempty"`
}

type RootsCapability struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

type SamplingCapability struct{}

type Implementation struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResult struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      Implementation     `json:"serverInfo"`
}

type ServerCapabilities struct {
	Tools     *ToolsCapability     `json:"tools,omitempty"`
	Resources *ResourcesCapability `json:"resources,omitempty"`
	Prompts   *PromptsCapability   `json:"prompts,omitempty"`
}

type ToolsCapability struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

type ResourcesCapability struct {
	Subscribe   bool `json:"subscribe,omitempty"`
	ListChanged bool `json:"listChanged,omitempty"`
}

type PromptsCapability struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

// Tool types
type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	InputSchema map[string]interface{} `json:"inputSchema"`
	Meta        map[string]interface{} `json:"_meta,omitempty"`
}

type ListToolsResult struct {
	Tools []Tool `json:"tools"`
}

type CallToolParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`
}

type CallToolResult struct {
	Content           []Content              `json:"content"`
	IsError           bool                   `json:"isError,omitempty"`
	StructuredContent map[string]interface{} `json:"structuredContent,omitempty"`
}

type Content struct {
	Type     string            `json:"type"`
	Text     string            `json:"text,omitempty"`
	Resource *ResourceContents `json:"resource,omitempty"`
}

// Resource types
type Resource struct {
	URI         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
}

type ListResourcesResult struct {
	Resources []Resource `json:"resources"`
}

type ReadResourceParams struct {
	URI string `json:"uri"`
}

type ReadResourceResult struct {
	Contents []ResourceContents `json:"contents"`
}

type ResourceContents struct {
	URI      string                 `json:"uri"`
	MimeType string                 `json:"mimeType,omitempty"`
	Text     string                 `json:"text,omitempty"`
	Meta     map[string]interface{} `json:"_meta,omitempty"`
}
