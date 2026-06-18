# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Model Context Protocol (MCP) server implementation in Go that provides Veracode security scanning capabilities to AI assistants and LLMs. It supports stdio transport for local filesystem operations and uses JSON-RPC 2.0 for communication.

## Building and Testing

### Build Commands

```powershell
# Full build with UI, quality checks, and tests
.\build.ps1

# Quick build (skip quality checks)
.\build.ps1 -Quick

# Build without tests
.\build.ps1 -NoTest

# Skip UI builds (Go server only)
.\build.ps1 -SkipUI

# UI only build
.\build.ps1 -UIOnly

# Verbose test output
.\build.ps1 -Verbose
```

The build script performs these steps:

1. Builds 5 React UI apps (pipeline/static/dynamic/local-sca/local-iac findings)
2. Runs `go fmt ./...`
3. Runs `go vet` (excludes `/api/generated/`)
4. Runs `golangci-lint run` (if available)
5. Runs `go test ./...`
6. Builds binary to `dist/veracode-mcp.exe`

### Manual Build

```bash
go mod download
go build -o dist/veracode-mcp.exe .
```

### Testing

```bash
go test ./...              # All tests
go test ./... -v           # Verbose output
go test ./internal/mcp_tools/...  # Specific package
```

## Architecture

### Auto-Registration Pattern

This codebase uses an **auto-registration architecture** where tools register themselves on package import via `init()` functions. There are three coordinated registries:

1. **ToolRegistry** (`internal/tool_registry/tool_loader.go`)
   - Loads tool definitions from `tools.json` (embedded at compile time)
   - Validates parameter schemas
   - Provides tool metadata to MCP clients

2. **ToolHandlerRegistry** (`internal/tool_registry/tool_handlers.go`)
   - Maps tool names to handler functions
   - Runtime tool invocation

3. **ToolImplRegistry** (`internal/tool_registry/tool_implementations.go`)
   - Manages tool lifecycle: Initialize → RegisterHandlers → Shutdown
   - Factory-based instantiation

The auto-registration mechanism itself (`RegisterTool`, `RegisterMCPTool`, `GetAllTools`) lives in `internal/mcp_tools/registry.go` and is called by each tool's `init()` function.

**Adding a new tool requires:**

- Add definition to `tools.json` with name, description, and parameter schema
- Create `internal/mcp_tools/<tool_name>.go` with an `init()` function that calls `RegisterTool()`
- Implement the `ToolImplementation` interface (Initialize, RegisterHandlers, Shutdown)
- The tool is automatically discovered and available to clients

### Core Architecture Layers

```text
MCP Client (Claude, VS Code, Codex)
         ↓ JSON-RPC 2.0 (stdio)
Transport Layer (internal/transport/)
         ↓
Protocol Handler (internal/server/server.go)
         ↓
Tool Registry & Dispatcher (internal/tool_registry/)
         ↓
Tool Implementations (internal/mcp_tools/)
         ↓
API Integration Layer (api/)
         ↓
Veracode APIs (HMAC-signed HTTPS)
```

### Key Directories

- `main.go` - Entry point, embeds `tools.json`, `instructions.json`, and UI HTML via `go:embed`
- `internal/server/` - MCP server core, request handling, protocol negotiation
- `internal/transport/` - Stdio transport implementation with buffered I/O
- `internal/types/` - MCP protocol type definitions (JSON-RPC 2.0)
- `internal/tool_registry/` - Tool definition loading and schema validation
- `internal/mcp_tools/` - Tool implementations (auto-registered via `init()`)
- `internal/cli/` - Logging configuration and server run helper
- `api/` - Veracode API integration
  - `client.go` - Client orchestrator
  - `rest/` - API wrappers with business logic
  - `rest/generated/` - Swagger-generated clients (DO NOT EDIT)
  - `xml/` - XML request/response types
  - `converters.go` / `types.go` - Cross-layer type conversions and shared types
- `hmac/` - HMAC-SHA256 authentication for Veracode APIs
- `credentials/` - Credential loading from `~/.veracode/veracode.yml` or env vars
- `workspace/` - `.veracode-workspace.json` configuration management
- `ui/` - React/TypeScript UI apps (Vite build, embedded via `go:embed`)

## MCP Protocol Implementation

### Supported MCP Methods

- `initialize` - Handshake, capability negotiation (supports protocol >= 2024-11-05)
- `notifications/initialized` - No response (per JSON-RPC 2.0 spec)
- `tools/list` - Returns tools from `tools.json`
- `tools/call` - Validates params against schema, dispatches to handler
- `resources/list` - Returns the 5 UI app resources (`ui://*/app.html`)
- `resources/read` - Serves embedded UI HTML for a requested `ui://` URI

### MCP Apps UI Support

The server detects UI support via the `io.modelcontextprotocol/ui` key in the client's `capabilities.extensions` object, looking for the `text/html;profile=mcp-app` MIME type in `mimeTypes`. When supported:

- Tools return dual content: `content` (JSON for LLM) and `structuredContent` (data for UI)
- UI apps are React apps embedded as single HTML strings via `go:embed`
- Five UI apps: pipeline findings, static findings, dynamic findings, local SCA findings, local IaC findings

## Veracode Integration

### Workspace Configuration

Create `.veracode-workspace.json` in your project root to associate code with a Veracode application profile:

```json
{
  "name": "My_Profile_Name"
}
```

Tools like `static-findings`, `dynamic-findings`, and `sca-findings` auto-detect the profile name from this file.

### Credential Sources (checked in order)

1. `~/.veracode/veracode.yml` (see Veracode CLI docs)
2. Environment variables: `VERACODE_API_ID` and `VERACODE_API_KEY`

### API Authentication

All API requests use HMAC-SHA256 signing (`hmac/` package) with a custom `Authorization` header. The `api-health` tool validates connectivity and credentials via `GET /api/authn/v2/principal`.

## Code Quality and Security

### Linting

The codebase uses `golangci-lint` with configuration in `.golangci.yml`. Generated code in `api/rest/generated/` is excluded from `go vet` checks.

### Veracode Scan Findings

This codebase is scanned with Veracode Pipeline Scan. Findings and mitigations are documented in `VERACODE_FINDINGS.md`. Key findings include:

- **Log injection (CWE-117)**: Input validation via `isValidMethod()` (transport layer) and `ValidateMethod()` (server layer) runs before any method-name logging
- **Error message disclosure (CWE-209)**: Error details only written to operator logs (stderr/file), never in JSON-RPC responses to client
- **Path traversal (CWE-73)**: Log file path is a CLI flag intentionally controlled by the operator (marked `#nosec G304`)

When modifying logging or error handling, ensure error details remain in operator logs and never leak to JSON-RPC client responses.

## Development Notes

### go:embed and Build Dependencies

The binary embeds these resources at compile time:

- `tools.json` - Tool definitions and schemas
- `instructions.json` - LLM instructions
- `ui/*/dist/mcp-app.html` - Five React UI apps

**When changing UI or tool definitions:**

1. UI changes: Run `.\build.ps1 -UIOnly` to rebuild UI apps first
2. Tool changes: Update `tools.json` and the corresponding `internal/mcp_tools/<tool>.go`
3. Full rebuild: Run `.\build.ps1` to embed all updated resources

### Transport Mode

This server only supports **stdio mode** (no HTTP/SSE mode). It operates as a subprocess communicating via stdin/stdout. Use `-log <filepath>` to write debug logs to a file rather than stderr (which can interfere with JSON-RPC).

### Logging

- In stdio mode, avoid `stderr` output during normal operation (use `-log` flag instead)
- Verbose logging (`-verbose`) should only be used for debugging
- Log messages should never include untrusted input without validation
- Error details belong in operator logs, not client responses

### Generated Code

Do not manually edit files in `api/rest/generated/`. These are Swagger-generated from OpenAPI specs. Instead, modify the wrapper layer in `api/rest/*.go` for business logic, error handling, and response transformation.
