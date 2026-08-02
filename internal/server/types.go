package server

// MCP Protocol types
type InitializeParams struct {
	ProtocolVersion string                 `json:"protocolVersion"`
	Capabilities    ClientCapabilities     `json:"capabilities"`
	ClientInfo      Implementation         `json:"clientInfo"`
	Meta            map[string]interface{} `json:"meta,omitempty"`
}

type ClientCapabilities struct {
	Roots      *RootsCapability       `json:"roots,omitempty"`
	Sampling   *SamplingCapability    `json:"sampling,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
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
	Instructions    string             `json:"instructions,omitempty"`
}

type ServerCapabilities struct {
	Tools     *ToolsCapability     `json:"tools,omitempty"`
	Resources *ResourcesCapability `json:"resources,omitempty"`
	Prompts   *PromptsCapability   `json:"prompts,omitempty"`
	Tasks     *TasksCapability     `json:"tasks,omitempty"`
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

// TasksCapability declares support for the MCP Tasks utility (spec 2025-11-25).
// Fields are present/omitted rather than boolean-valued, per the spec's examples.
type TasksCapability struct {
	Cancel   *TasksCancelCapability   `json:"cancel,omitempty"`
	Requests *TasksRequestsCapability `json:"requests,omitempty"`
}

// TasksCancelCapability marks support for the tasks/cancel operation.
type TasksCancelCapability struct{}

// TasksRequestsCapability lists which request types can be task-augmented.
type TasksRequestsCapability struct {
	Tools *TasksToolsRequests `json:"tools,omitempty"`
}

// TasksToolsRequests marks support for task-augmented tools/call requests.
type TasksToolsRequests struct {
	Call *TasksCallCapability `json:"call,omitempty"`
}

// TasksCallCapability is an empty marker (present = supported).
type TasksCallCapability struct{}

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
