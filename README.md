# Veracode MCP Server (Go)

⚠️ **ALPHA SOFTWARE** - This is early-stage software under active development. APIs and functionality may change without notice. This is not production-ready code.

A Model Context Protocol (MCP) server implementation in Go that provides Veracode security scanning capabilities to AI assistants and LLMs. Supports both stdio and HTTP/SSE transport modes.

## Features

- **Dual Transport Support**
  - stdio mode for local process communication
  - HTTP mode with Server-Sent Events (SSE) for network communication
  
- **MCP Protocol Support**
  - JSON-RPC 2.0 message handling
  - Tool invocation capabilities
  - Resource access
  - Protocol version negotiation (supports 2024-11-05 and newer including 2025-06-18)
  
- **Veracode Integration**
  - Dynamic (DAST) findings
  - Static (SAST) findings
  - SCA (Software Composition Analysis) findings
  - Pipeline scan results
  - API health checks
  - Finding details
  - Workspace packaging

## Installation

```bash
# Clone the repository
git clone https://github.com/dipsylala/veracodemcp-go.git
cd veracodemcp-go

# Install dependencies
go mod download

# Build the server
.\build.ps1    # Windows (builds to dist/veracode-mcp.exe)
# or
go build -o dist/veracode-mcp.exe .  # Manual build
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

### Command Line Options

```bash
.\dist\veracode-mcp.exe [options]

# Basic usage (silent mode, no logging)
.\dist\veracode-mcp.exe

# With debug logging to file (recommended for troubleshooting)
.\dist\veracode-mcp.exe -log veracode-mcp.log

# With verbose logging to stderr
.\dist\veracode-mcp.exe -verbose
```

**Important:** When using stdio mode with MCP clients (like VS Code or Codex), avoid using `-verbose` as stderr output can interfere with JSON-RPC communication. Instead, use `-log <filepath>` to write debug information to a file.     Enable verbose logging to stderr (disabled by default)
  -log string
        Log file path for debugging (recommended for stdio mode)
  -version
        Display version information
```

### Stdio Mode (Default)

The stdio mode is ideal for local integrations where the MCP server runs as a subprocess:

```bash
.\dist\veracode-mcp.exe -mode stdio
```

Or simply:
```bMCP Client Configuration

**VS Code / Codex:**

Add to your MCP client configuration (e.g., `~/.codex/config.toml`):

```toml
[mcp_servers.Veracode]
command = "/path/to/veracode-mcp.exe"
args = ["-log", "/path/to/veracode-mcp.log"]  # Optional but recommended for debugging
```

**Claude Desktop:**

Add to `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "veracode": {
      "command": "/path/to/veracode-mcp.exe",
      "args": ["-log", "/path/to/veracode-mcp.log"]
    }
  }
}
```

### ash
.\dist\veracode-mcp.exe
```

### HTTP Mode

The HTTP mode allows remote connections via HTTP with Server-Sent Events:

```bash
.\dist\veracode-mcp.exe -mode http -addr :8080
```

The HTTP server provides these endpoints:
- `GET /sse` - Establish SSE connection for receiving responses
- `POST /message?sessionId=<id>` - Send JSON-RPC requests
- `GET /health` - Health check endpoint

## Architecture

### File Structure

- `main.go` - Entry point and mode selection
- `server.go` - Core MCP server logic and request handling
- `types.go` - MCP protocol type definitions
- `stdio.go` - Stdio transport implementation
- `http.go` - HTTP/SSE transport implementation
- `tool_loader.go` - Tool registry and auto-registration system
- `tool_handlers.go` - Tool handler registry
- `tool_implementations.go` - Tool implementation registry
- `tools.json` - JSON-driven tool definitions and schemas
- `api/` - API integration layer
  - `client.go` - Client orchestrator
  - `helpers/` - Business logic wrappers
  - `generated/` - Swagger-generated API clients
- `tools/` - Tool implementations (auto-registered)
- `credentials/` - Credential management
- `hmac/` - HMAC authentication utilities
- `docs/` - Documentation

### Tool Architecture

The server uses an **auto-registration architecture** where tools register themselves on import:
 (supports protocol versions >= 2024-11-05)
- `notifications/initialized` - Client initialization confirmation (properly handled as notification per JSON-RPC 2.0 spec)
- `tools/list` - List available tools (loaded from tools.json)
- `tools/call` - Execute a tool with validated parameters
- `resources/list` - List available resources
- `resources/read` - Read resource content

**Protocol Compatibility:**
- Automatically negotiates protocol version with client
- Tested with Codex (protocol 2025-06-18) and VS Code MCP clients
- Strict JSON-RPC 2.0 compliance (notifications receive no response)
- Buffered stdio transport with explicit flushing for reliable communic
   - Rich descriptions optimized for LLM consumption
   - Enum values and allowed ranges

2. **Tool Implementations** (`tools/*.go`)
   - Type-safe Go implementations
   - Self-registering via `init()` functions
   - Lifecycle management (Initialize → RegisterHandlers → Shutdown)

3. **Registries**
   - **ToolRegistry**: Loads tool definitions from `tools.json`
   - **ToolHandlerRegistry**: Maps tool names to handler functions
   - **ToolImplRegistry**: Manages tool implementation lifecycle

**Benefits:**
- ✅ No manual registration - just create the file
- ✅ Type-safe Go implementations
- ✅ Self-documenting via JSON schemas
- ✅ Thread-safe registry
- ✅ Independent testing
- ✅ Modular tool organization

### MCP Server Methods

The server implements these core MCP protocol methods:

- `initialize` - Protocol handshake and capability negotiation
- `tools/list` - List available tools (loaded from tools.json)
- `tools/call` - Execute a tool with validated parameters
- `resources/list` - List available resources
- `resources/read` - Read resource content
- `notifications/initialized` - Client initialization confirmation

## Available Tools

The server provides these Veracode-specific tools:

- **api-health** - Verify Veracode API connectivity and credentials
- **dynamic-findings** - Retrieve runtime security vulnerabilities from Dynamic Analysis (DAST) scans
- **static-findings** - Retrieve source code vulnerabilities from Static Analysis (SAST) scans
- **sca-findings** - Retrieve third-party component vulnerabilities from Software Composition Analysis
- **finding-details** - Get detailed information about a specific finding

- **package-workspace** - Package workspace files for Veracode upload
- **pipeline-scan** - Start an asynchronous pipeline scan, with the largest packaged file as default
- **pipeline-status** - Check the status of a Pipeline Scan
- **pipeline-results** - Get results from Veracode Pipeline Scans
- **pipeline-detailed-results** - Get detailed results from Pipeline Scans with full flaw information

> **Note:** Use the `tools/list` MCP method to see all available tools with their complete parameter schemas and documentation.

## Adding New Tools

See [docs/ADDING_TOOLS.md](docs/ADDING_TOOLS.md) for the complete guide on implementing new MCP tools.

## Documentation

- **[Quick Start](docs/QUICKSTART.md)** - Get up and running quickly
- **[Adding Tools](docs/ADDING_TOOLS.md)** - Create new MCP tools (extensibility guide)
- **[API Integration](docs/API_INTEGRATION.md)** - Integrate Veracode REST APIs
- **[Code Quality](docs/CODE_QUALITY.md)** - Build tools and quality checks
- **[Tool Structure](docs/TOOLS_STRUCTURE.md)** - Tool system architecture
- **[MCP Testing](docs/MCP_TESTING.md)** - Testing MCP implementations

## Building

The project includes a PowerShell build script with quality checks:

```powershell
.\build.ps1          # Full build with all checks
.\build.ps1 -Quick   # Fast build, skip checks
.\build.ps1 -NoTest  # Build without running tests
.\build.ps1 -Verbose # Show detailed test output
```

The build script performs:
1. Code formatting (`go fmt`)
2. Static analysis (`go vet`)
3. Linting (`golangci-lint` if available)
4. Unit tests
5. Binary compilation to `dist/veracode-mcp.exe`

## Testing

Run tests with:

```bash
go test ./...        # Run all tests
go test ./... -v     # Verbose output
go test ./tools/...  # Test only tools package
```

See [docs/MCP_TESTING.md](docs/MCP_TESTING.md) for MCP integration testing.

## License

MIT

## Contributing

⚠️ This is alpha software. Contributions are welcome, but please be aware that APIs and architecture may change significantly.

Please feel free to submit issues and pull requests.
