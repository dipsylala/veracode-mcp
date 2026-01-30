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

### Veracode CLI

Some tools (such as `package-workspace`, `pipeline-scan`, `run-sca-scan`) require the Veracode CLI to be installed and available in your system PATH.

**Install the Veracode CLI:**

*Windows (PowerShell):*

```powershell
iex (iwr https://tools.veracode.com/veracode-cli/install.ps1)
```

*macOS/Linux:*

```bash
curl -fsS https://tools.veracode.com/veracode-cli/install | sh
```

**Authenticate the CLI:**

After installation, configure your API credentials:

```bash
veracode configure
```

For detailed installation instructions and alternative methods, see the [official Veracode CLI installation guide](https://docs.veracode.com/r/Install_the_Veracode_CLI).

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

Options:
  -mode string
        Server mode: stdio or http (default "stdio")
  -addr string
        HTTP server address, only used in http mode (default ":8080")
  -verbose
        Enable verbose logging to stderr (disabled by default)
  -log string
        Log file path for debugging (recommended for stdio mode)
  -force-mcp-app
        Force MCP Apps mode (always send structuredContent regardless of client capabilities)
  -version
        Display version information
```

**Usage Examples:**

```bash
# Basic usage (silent mode, stdio transport)
.\dist\veracode-mcp.exe

# With debug logging to file (recommended for troubleshooting)
.\dist\veracode-mcp.exe -log veracode-mcp.log

# With verbose logging to stderr (avoid in stdio mode with MCP clients)
.\dist\veracode-mcp.exe -verbose

# HTTP mode on custom port
.\dist\veracode-mcp.exe -mode http -addr :9000

# Force MCP Apps UI mode
.\dist\veracode-mcp.exe -force-mcp-app -log debug.log
```

**Important:** When using stdio mode with MCP clients (like VS Code or Claude Desktop), avoid using `-verbose` as stderr output can interfere with JSON-RPC communication. Instead, use `-log <filepath>` to write debug information to a file.

### Stdio Mode (Default)

The stdio mode is ideal for local integrations where the MCP server runs as a subprocess:

```bash
.\dist\veracode-mcp.exe -mode stdio
```

Or simply:

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

### HTTP Mode

The HTTP mode allows remote connections via HTTP with Server-Sent Events:

```bash
.\dist\veracode-mcp.exe -mode http -addr :8080
```

The HTTP server provides these endpoints:

- `GET /sse` - Establish SSE connection for receiving responses
- `POST /message?sessionId=<id>` - Send JSON-RPC requests
- `GET /health` - Health check endpoint

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
- **run-sca-scan** - Run Software Composition Analysis scan on a directory to identify vulnerable dependencies
- **get-local-sca-results** - Read and parse local SCA scan results from veracode.json file

> **Note:** Use the `tools/list` MCP method to see all available tools with their complete parameter schemas and documentation.

## Adding New Tools

See [docs/ADDING_TOOLS.md](docs/ADDING_TOOLS.md) for the complete guide on implementing new MCP tools.

## Documentation

- **[Architecture & Design](docs/DESIGN.md)** - System architecture and design decisions
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
