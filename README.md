# Veracode MCP Server (Go)

A basic Model Context Protocol (MCP) server implementation in Go that supports both stdio and HTTP/SSE transport modes.

## Features

- **Dual Transport Support**
  - stdio mode for local process communication
  - HTTP mode with Server-Sent Events (SSE) for network communication
  
- **MCP Protocol Support**
  - JSON-RPC 2.0 message handling
  - Tool invocation capabilities
  - Resource access
  - Protocol version negotiation

- **Example Implementations**
  - Echo tool (demonstrates basic tool functionality)
  - Example resource (demonstrates resource access)

## Installation

```bash
# Clone the repository
git clone https://github.com/dipsylala/veracodemcp-go.git
cd mcp-go

# Install dependencies
go mod download

# Build the server
.\build.ps1    # Windows (builds to dist/mcp-server.exe)
# or
go build -o dist/mcp-server .  # Manual build
```

## Configuration

### Veracode API Credentials

The server requires Veracode API credentials to access the Veracode platform. Credentials can be provided in two ways (checked in order):

1. **File-based configuration** (Recommended)
   
   Create `~/.veracode/veracode.yml`:
   ```yaml
   api:
     key-id: YOUR_API_KEY_ID
     key-secret: YOUR_API_KEY_SECRET
   ```

   **Setup commands:**
   
   *Linux/macOS:*
   ```bash
   mkdir -p ~/.veracode
   cat > ~/.veracode/veracode.yml << EOF
   api:
     key-id: YOUR_API_KEY_ID
     key-secret: YOUR_API_KEY_SECRET
   EOF
   chmod 600 ~/.veracode/veracode.yml
   ```

   *Windows PowerShell:*
   ```powershell
   New-Item -ItemType Directory -Path "$env:USERPROFILE\.veracode" -Force
   @"
   api:
     key-id: YOUR_API_KEY_ID
     key-secret: YOUR_API_KEY_SECRET
   "@ | Out-File -FilePath "$env:USERPROFILE\.veracode\veracode.yml" -Encoding UTF8
   ```

2. **Environment variables** (Fallback)
   
   ```bash
   export VERACODE_API_ID="YOUR_API_KEY_ID"
   export VERACODE_API_KEY="YOUR_API_KEY_SECRET"
   ```

See [credentials/README.md](credentials/README.md) for detailed information.

## Usage

### Stdio Mode (Default)

The stdio mode is ideal for local integrations where the MCP server runs as a subprocess:

```bash
.\dist\mcp-server.exe -mode stdio
```

Or simply:
```bash
.\dist\mcp-server.exe
```

### HTTP Mode

The HTTP mode allows remote connections via HTTP with Server-Sent Events:

```bash
.\dist\mcp-server.exe -mode http -addr :8080
```

The HTTP server provides these endpoints:
- `GET /sse` - Establish SSE connection for receiving responses
- `POST /message?sessionId=<id>` - Send JSON-RPC requests
- `GET /health` - Health check endpoint

## Architecture

### Plugin Architecture

The server uses a **plugin-based architecture** following the **open/closed principle**:
- **Open for extension**: Add new tools via plugins without modifying core code
- **Closed for modification**: Core server logic remains stable

**Benefits:**
- ✅ Modular tool organization
- ✅ Independent plugin development
- ✅ Clean separation of concerns
- ✅ Easy testing and maintenance
- ✅ Optional feature enablement

See [PLUGINS.md](PLUGINS.md) for detailed plugin development guide.

### File Structure

- `main.go` - Entry point and mode selection
- `server.go` - Core MCP server logic and request handling
- `types.go` - MCP protocol type definitions
- `stdio.go` - Stdio transport implementation
- `http.go` - HTTP/SSE transport implementation
- `api/` - API integration layer
  - `client.go` - Client orchestrator
  - `helpers/` - Business logic wrappers
  - `generated/` - Swagger-generated API clients
- `tools/` - Tool implementations (auto-registered)
- `docs/` - Documentation
  - `ADDING_TOOLS.md` - Guide for adding new tools
  - `API_INTEGRATION.md` - API architecture guide
  - `CODE_QUALITY.md` - Build tools and quality checks
  - `TOOLS_STRUCTURE.md` - Tool system architecture
  - `QUICKSTART.md` - Quick start guide
  - `PLUGINS.md` - Plugin development guide

### Auto-Registration Tool Architecture

The server uses **auto-registration** - tools register themselves on import:

**Benefits:**
- ✅ No manual registration - just create the file
- ✅ Type-safe Go implementations
- ✅ Self-documenting via GetInputSchema()
- ✅ Thread-safe registry
- ✅ Independent testing

See [docs/ADDING_TOOLS.md](docs/ADDING_TOOLS.md) for the complete guide.

## Available Tools (Examples)

### API Health Check
**Name:** `api-health`  
**Purpose:** Verify Veracode API connectivity  
**Parameters:** None

### Dynamic Findings
**Name:** `dynamic-findings`  
**Purpose:** Retrieve DAST scan findings  
**Parameters:** 
- `app_profile` (required) - Application GUID or name
- `severity` (optional) - Filter by severity levels
- `status` (optional) - Filter by finding status

### Static Findings
**Name:** `static-findings`  
**Purpose:** Retrieve SAST scan findings  
**Parameters:** Same as dynamic-findings

## Adding New Tools

See [docs/ADDING_TOOLS.md](docs/ADDING_TOOLS.md) for the complete guide.

### Quick Example

Create `tools/your_tool.go`:

```go
package tools

import "context"

func init() {
    RegisterTool("your-tool", func() ToolImplementation {
        return &YourTool{}
    })
}

type YourTool struct{}

func (t *YourTool) GetName() string { return "your-tool" }

func (t *YourTool) GetDescription() string {
    return "Description for LLMs about when and how to use this tool"
}

func (t *YourTool) GetInputSchema() interface{} {
    return map[string]interface{}{
        "type": "object",
        "properties": map[string]interface{}{
            "param": map[string]interface{}{
                "type": "string",
                "description": "Parameter description",
            },
        },
        "required": []string{"param"},
    }
}

func (t *YourTool) Handle(ctx context.Context, params map[string]interface{}) (interface{}, error) {
    // Your tool logic here
    return map[string]interface{}{
        "success": true,
        "message": "Result",
    }, nil
}
```

Build and it's automatically available!

**JSON Definitions** (`tools.json`):
- Tool metadata (name, description, parameters)
- Parameter types, validation rules, and constraints
- Rich descriptions optimized for LLM consumption
- Enum values and allowed ranges

**Go Implementations** (`tool_handlers.go`):
- Actual tool logic and API integrations
- Type-safe parameter handling
- Business logic and data processing

**ToolHandlerRegistry**:
- Maps tool names to handler functions
- Automatically populated by tool implementations
- Validates parameters before execution

**ToolImplRegistry**:
- Manages tool implementation lifecycle (Initialize → RegisterHandlers → Shutdown)
- Auto-discovers tool implementations
- Isolated tool initialization and error handling

This provides:
- ✅ Consistent, LLM-friendly tool descriptions
- ✅ Centralized parameter definitions
- ✅ Easy addition of new tools without recompilation (JSON)
- ✅Create or update a plugin in `plugins/`
3. Implement handler method in plugin
4. Register handler in plugin's `RegisterHandlers()` method
5. Server automatically loads plugin, validates,
Example workflow:
1. Define tool in `tools.json` with rich metadata
2. Implement handler function in `tool_handlers.go`
3. Register handler in `registerDefaultHandlers()`
4. Server automatically validates and routes calls

### JSON-Driven Tool Definitions

Tools are defined in `tools.json` following a standardized schema. This provides:
- Consistent, LLM-friendly tool descriptions
- Centralized parameter definitions with validation rules
- Easy addition of new tools without recompilation
- Rich metadata for parameter types and constraints

Example tool definition:
```json
{
  "name": "get-dynamic-findings",
  "description": "Get runtime vulnerabilities from DYNAMIC analysis scans...",
  "category": "findings",
  "params": [
    {
      "name": "app_profile",
      "type": "string",
      "isRequired": false,
      "description": "Application Profile GUID or name..."
    }
  ]
}
```

### Key Components

**MCPServer**: Core server handling MCP protocol methods:
- `initialize` - Protocol handshake
- `tools/list` - List available tools (loaded from tools.json)
- `tools/call` - Execute a tool
- `resources/list` - List available resources
- `resources/read` - Read a resource

**ToolRegistry**: Loads and manages tool definitions from `tools.json`
PluginRegistry**: Manages plugin lifecycle and handler registration

**
**StdioTransport**: Handles line-delimited JSON-RPC over stdin/stdout

**HTTPTransport**: Manages HTTP connections with SSE for async responses

## Available Tools

The server includes these Veracode-specific tools (defined in `tools.json`):

### get-dynamic-findings
Get runtime vulnerabilities from DYNAMIC analysis scans.

**Parameters:**
- `app_profile` (string, optional) - Application Profile GUID or name
- `sandbox` (string, optional) - Sandbox GUID or name
- `severity` (array, optional) - Filter by severity levels
- `status` (array, optional) - Filter by remediation status
- `cwe_ids` (array, optional) - Filter by CWE IDs
- `violates_policy` (boolean, optional) - Filter by policy violation
- `page` (number, optional) - Page number (0-based)
- `size` (number, optional) - Findings per page (1-500)
- `include_mitigations` (boolean, optional) - Include mitigation comments

### get-static-findings
Get source code vulnerabilities from STATIC analysis scans.

## Documentation

- **[Quick Start](docs/QUICKSTART.md)** - Get up and running quickly
- **[Adding Tools](docs/ADDING_TOOLS.md)** - Create new MCP tools
- **[API Integration](docs/API_INTEGRATION.md)** - Integrate Veracode REST APIs
- **[Code Quality](docs/CODE_QUALITY.md)** - Build tools and quality checks
- **[Tool Structure](docs/TOOLS_STRUCTURE.md)** - Tool system architecture
- **[Plugin Development](docs/PLUGINS.md)** - Advanced plugin patterns
3. Supports the core MCP methods (initialize, tools, resources)
4. Uses protocol version `2024-11-05`

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.
