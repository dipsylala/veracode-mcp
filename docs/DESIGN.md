# Architecture & Design

This document describes the architecture and design decisions of the Veracode MCP Server.

## Overview

The Veracode MCP Server is a Model Context Protocol (MCP) server implementation in Go that provides Veracode security scanning capabilities to AI assistants and LLMs. It supports both stdio and HTTP/SSE transport modes.

## Architecture

### High-Level Design

```text
┌─────────────────────────────────────────────────────────────┐
│           MCP Client (AI Assistant)                         │
│          (Claude, VS Code, Codex, etc.)                     │
└────────────────────────┬────────────────────────────────────┘
                         │ JSON-RPC 2.0
                         │ (stdio or HTTP/SSE)
┌────────────────────────▼────────────────────────────────────┐
│                    MCP Server (Go)                          │
│  ┌──────────────────────────────────────────────────────┐   │
│  │   Transport Layer (stdio / HTTP+SSE)                 │   │
│  └──────────────────────┬───────────────────────────────┘   │
│  ┌──────────────────────▼───────────────────────────────┐   │
│  │   Protocol Handler (JSON-RPC 2.0)                    │   │
│  │   - initialize, tools/list, tools/call, etc.         │   │
│  └──────────────────────┬───────────────────────────────┘   │
│  ┌──────────────────────▼───────────────────────────────┐   │
│  │   Tool Registry & Dispatcher                         │   │
│  │   - Loads tools.json                                 │   │
│  │   - Routes calls to implementations                  │   │
│  └──────────────────────┬───────────────────────────────┘   │
│  ┌──────────────────────▼───────────────────────────────┐   │
│  │   Tool Implementations (mcp_tools/*.go)              │   │
│  │   - Auto-registered via init()                       │   │
│  │   - Type-safe parameter handling                     │   │
│  └──────────────────────┬───────────────────────────────┘   │
│  ┌──────────────────────▼───────────────────────────────┐   │
│  │   API Integration Layer (api/)                       │   │
│  │   - Veracode REST API clients                        │   │
│  │   - HMAC authentication                              │   │
│  └──────────────────────┬───────────────────────────────┘   │
└─────────────────────────┼───────────────────────────────────┘
                          │ HTTPS
                          │ (HMAC-signed requests)
                 ┌────────▼────────┐
                 │  Veracode APIs  │
                 │   (Platform)    │
                 └─────────────────┘
```

### File Structure

**Core Server:**

- `main.go` - Entry point and mode selection
- `internal/server/server.go` - Core MCP server logic and request handling
- `internal/types/` - MCP protocol type definitions
- `internal/transport/stdio.go` - Stdio transport implementation
- `internal/transport/http.go` - HTTP/SSE transport implementation

**Tool System:**

- `tools.json` - JSON-driven tool definitions and schemas
- `internal/tools/registry.go` - Tool registry and auto-registration system
- `mcp_tools/` - MCP tool implementations (auto-registered)
  - `registry.go` - Tool implementation registry
  - `*_findings.go` - Security findings tools
  - `pipeline_*.go` - Pipeline scan tools
  - `*_sca_*.go` - SCA analysis tools

**API Integration:**

- `api/` - API integration layer
  - `client.go` - Client orchestrator
  - `findings.go` - Findings API wrapper
  - `applications.go` - Applications API wrapper
  - `health.go` - Health check API
  - `generated/` - Swagger-generated API clients

**Supporting Modules:**

- `credentials/` - Credential management (file + env vars)
- `hmac/` - HMAC authentication for Veracode APIs
- `workspace/` - Workspace configuration management

**UI Components:**

- `ui/static-findings-app/` - React UI for static findings
- `ui/dynamic-findings-app/` - React UI for dynamic findings
- `ui/pipeline-results-app/` - React UI for pipeline results

## Tool Architecture

### Auto-Registration Pattern

The server uses an **auto-registration architecture** where tools register themselves on import:

**Tool Definition** (`tools.json`)

```json
{
  "name": "static-findings",
  "description": "Get source code vulnerabilities...",
  "params": [
    {
      "name": "application_path",
      "type": "string",
      "isRequired": true,
      "description": "..."
    }
  ]
}
```

**Tool Implementation** (`mcp_tools/static_findings.go`)

```go
package mcp_tools

// Auto-register on package import
func init() {
    RegisterTool("static-findings", func() ToolImplementation {
        return NewStaticFindingsTool()
    })
}

type StaticFindingsTool struct{}

func (t *StaticFindingsTool) Initialize() error { ... }
func (t *StaticFindingsTool) RegisterHandlers(registry HandlerRegistry) error { ... }
func (t *StaticFindingsTool) Shutdown() error { ... }
```

**Registries:**

1. **ToolRegistry** (`internal/tools/registry.go`)
   - Loads tool definitions from `tools.json`
   - Validates parameter schemas
   - Provides tool metadata to clients

2. **ToolHandlerRegistry** (`mcp_tools/registry.go`)
   - Maps tool names to handler functions
   - Thread-safe concurrent access
   - Runtime tool invocation

3. **ToolImplRegistry** (`mcp_tools/registry.go`)
   - Manages tool implementation lifecycle
   - Initialize → RegisterHandlers → Shutdown
   - Factory-based instantiation

**Benefits:**

- ✅ **No manual registration** - Just create the file and import the package
- ✅ **Type-safe** - Go implementations with compile-time checking
- ✅ **Self-documenting** - JSON schemas describe parameters
- ✅ **Testable** - Each tool can be tested independently
- ✅ **Modular** - Tools are independent, loosely coupled

### Tool Execution Flow

```text
1. Client sends tools/call request
        ↓
2. Server validates against tools.json schema
        ↓
3. Server dispatches to registered handler
        ↓
4. Handler parses typed parameters
        ↓
5. Handler calls Veracode API (if needed)
        ↓
6. Handler formats response (JSON + UI)
        ↓
7. Server sends response to client
```

## MCP Protocol Implementation

### Supported Methods

The server implements these core MCP protocol methods:

- **`initialize`** - Protocol handshake and capability negotiation
  - Supports protocol versions >= 2024-11-05
  - Returns server capabilities and metadata
  
- **`notifications/initialized`** - Client initialization confirmation
  - Properly handled as notification (no response per JSON-RPC 2.0)
  
- **`tools/list`** - List available tools
  - Loaded from `tools.json`
  - Rich parameter schemas with validation rules
  
- **`tools/call`** - Execute a tool
  - Parameter validation against schemas
  - Type coercion (strings, numbers, booleans, arrays)
  - Error handling and reporting
  
- **`resources/list`** - List available resources
  - Currently returns empty (reserved for future use)
  
- **`resources/read`** - Read resource content
  - Currently not implemented (reserved for future use)

### Protocol Compatibility

- **Version Negotiation**: Automatically negotiates protocol version with client
- **Tested Clients**: Codex (protocol 2025-06-18), VS Code MCP extension
- **JSON-RPC 2.0 Compliance**: Strict adherence to spec
  - Notifications receive no response
  - Error codes follow JSON-RPC standard
  - Batching not currently supported
- **Transport Reliability**:
  - Buffered stdio transport with explicit flushing
  - SSE with automatic reconnection support
  - Message framing and error recovery

### MCP Apps UI Support

The server includes **MCP Apps** support for rich UI experiences:

**Detection:**

- Checks client capabilities in `initialize` request
- Looks for `experimental.mcpApps` capability
- Can be forced with `-force-mcp-app` flag

**Response Format:**

```json
{
  "content": [
    {
      "type": "text",
      "text": "JSON data for LLM"
    }
  ],
  "structuredContent": {
    "application": {...},
    "findings": [...],
    "summary": {...}
  }
}
```

**UI Apps:**

- React + TypeScript applications
- Vite build system
- Embedded via `go:embed` directives
- Rendered in MCP-capable clients (Codex, VS Code)

## Transport Modes

### Stdio Mode (Default)

**Use Case:** Local integrations where MCP server runs as subprocess

**Implementation:**

- Reads JSON-RPC messages from stdin
- Writes responses to stdout
- Logs to file (via `-log` flag) or discards
- Buffered I/O with explicit flushing

**Client Configuration:**

```toml
[mcp_servers.Veracode]
command = "/path/to/veracode-mcp.exe"
args = ["-log", "debug.log"]
```

**Advantages:**

- Simple process model
- No network configuration
- Secure (no exposed ports)

**Limitations:**

- Single client per process
- No remote access

### HTTP Mode

**Use Case:** Remote connections, multiple clients, cloud deployments

**Implementation:**

- HTTP server with SSE for server→client messages
- POST endpoint for client→server messages
- Session-based message routing

**Endpoints:**

- `GET /sse?sessionId=<id>` - Establish SSE connection
- `POST /message?sessionId=<id>` - Send JSON-RPC request
- `GET /health` - Health check

**Advantages:**

- Remote access
- Multiple concurrent clients
- Standard HTTP infrastructure

**Limitations:**

- Network configuration required
- More complex deployment

## API Integration

### Veracode REST APIs

The server integrates with multiple Veracode APIs:

1. **Applications API** - Application metadata and GUID lookup
2. **Findings API** - Static, dynamic, and SCA findings
3. **Identity API** - Health checks and authentication

### Authentication

**HMAC Signing:**

- All API requests signed with HMAC-SHA256
- Credentials from `~/.veracode/veracode.yml` or env vars
- Custom `Authorization` header with signature

**Credential Sources (checked in order):**

1. File: `~/.veracode/veracode.yml`
2. Environment: `VERACODE_API_ID` + `VERACODE_API_KEY`

### API Client Architecture

**Generated Clients** (`api/generated/`)

- Swagger-generated from OpenAPI specs
- Type-safe request/response structs
- Automatic serialization/deserialization

**Wrapper Layer** (`api/*.go`)

- Business logic and error handling
- Response transformation
- Pagination support
- Retry logic (future)

**Client Orchestration** (`api/client.go`)

- Single entry point for all APIs
- Credential management
- HMAC authentication injection
- HTTP client configuration

## Data Flow Examples

### Static Findings Request

1. Client: tools/call("static-findings", `{application_path: "/app"}`)
2. Server: Validate parameters against tools.json schema
3. Server: Lookup application GUID from workspace config
4. API: GET /appsec/v1/applications?name=`<app_profile>`
5. API: GET /appsec/v2/applications/`<guid>`/findings?type=STATIC
6. Handler: Transform API response to MCP format
7. Handler: Sort by severity, build summary, pagination
8. Server: Return {content: JSON, structuredContent: UI_data}
9. Client: Render UI or process JSON

### Pipeline Scan Workflow

1. package-workspace → Creates .zip artifact
2. pipeline-scan → Runs local scan, generates results.json
3. pipeline-status → Checks scan completion
4. pipeline-results → Reads results.json, returns findings
5. pipeline-detailed-results → Gets specific flaw with data flow

## Design Decisions

### Why Go?

- **Performance**: Fast startup, low memory footprint
- **Concurrency**: Built-in goroutines for async operations
- **Type Safety**: Compile-time checking, robust error handling
- **Single Binary**: Easy deployment, no runtime dependencies
- **Standard Library**: Excellent HTTP, JSON, and I/O support

### Why Auto-Registration?

- **Developer Experience**: Add tool by creating file, automatic discovery
- **Maintainability**: No central registry to update
- **Modularity**: Tools are independent units
- **Testing**: Each tool can be tested in isolation

### Why JSON Schema in tools.json?

- **Single Source of Truth**: Tool metadata in one place
- **Client Integration**: Tools automatically appear in client
- **Validation**: Parameter checking without code changes
- **Documentation**: Self-documenting for LLMs

### Why Embedded UI?

- **Single Binary**: No separate UI server or static file serving
- **Versioning**: UI and server always in sync
- **Deployment**: Copy one file, everything works
- **Security**: No CDN dependencies, works offline

## Future Enhancements

### Planned Features

- **Async Tool Execution**: Long-running scans don't block
- **Caching**: Cache API responses for performance
- **Retry Logic**: Automatic retry on transient failures
- **WebSocket Transport**: Alternative to SSE for bidirectional streaming
- **Resource Support**: Expose scan results as MCP resources
- **Batch Operations**: Process multiple applications at once

### Extensibility Points

- **New Tools**: Add `mcp_tools/<tool>.go` + entry in `tools.json`
- **New APIs**: Add client in `api/generated/`, wrapper in `api/`
- **New Transports**: Implement `Transport` interface
- **Custom Auth**: Replace HMAC provider in `api/client.go`
- **UI Customization**: Modify React apps in `ui/*/`

## Related Documentation

- **[Adding Tools](ADDING_TOOLS.md)** - Create new MCP tools
- **[API Integration](API_INTEGRATION.md)** - Integrate Veracode APIs
- **[Tool Structure](TOOLS_STRUCTURE.md)** - Tool system internals
- **[Code Quality](CODE_QUALITY.md)** - Build and quality checks
- **[MCP Testing](MCP_TESTING.md)** - Testing strategies
