# Veracode MCP Server

⚠️ **ALPHA SOFTWARE** - This is early-stage software under active development. APIs and functionality may change without notice. This is not production-ready code.

A Model Context Protocol (MCP) server implementation in Go that provides Veracode security scanning capabilities to AI assistants and LLMs. Uses stdio transport for local filesystem operations.

This is my 4th version, after writing it in TypeScript, Python, as a set of [Agent Skills](https://agentskills.io/home), and now - in Go. Go ultimately makes it easier to distribute, and I wanted more practice in it, so here we go.

## Features

- **MCP Protocol Support**
  - stdio transport for local process communication
  - JSON-RPC 2.0 message handling
  - Tool invocation capabilities
  - Resource access
  - Protocol version negotiation (supports 2024-11-05 and newer including 2025-06-18)
  
- **Veracode Integration**
  - Platform Dynamic (DAST) findings
  - Platform Static (SAST) findings
  - Platform SCA (Software Composition Analysis) findings
  - Workspace packaging for scan preparation
  - Static Pipeline and scan results
  - Finding details - Data paths and dynamic request/responses

## Installation

### Download from Releases

Download the latest pre-built binary from the [Releases page](https://github.com/dipsylala/veracode-mcp/releases):

Windows · macOS · Linux | x64 · ARM64

Extract and place the executable in a directory of your choice (e.g., `C:\Program Files\VeracodeMCP\` on Windows or `/usr/local/bin/` on macOS/Linux).

### Install Veracode CLI (Required)

Some tools (such as `package-workspace`, `pipeline-scan`, `run-sca-scan`) require the Veracode CLI to be installed and available in your system PATH.

Given the Veracode installation process requires elevated privileges, we took the decision for the user to perform the installation themselves, rather than an MCP requesting elevated privileges and installing software on a machine.

**Install the Veracode CLI:**

*Windows (Admin PowerShell):*

```powershell
Set-ExecutionPolicy AllSigned -Scope Process -Force
iex (iwr https://tools.veracode.com/veracode-cli/install.ps1)
```

*macOS/Linux:*

```bash
curl -fsS https://tools.veracode.com/veracode-cli/install | sh
```

For detailed installation instructions and alternative methods, see the [official Veracode CLI installation guide](https://docs.veracode.com/r/Install_the_Veracode_CLI).

## Configuration

### Veracode API Credentials

**Authenticate the CLI:**

After installation, configure your API credentials via:

1. **Veracode CLI-based configuration** (Recommended)

```bash
veracode configure
```

2. **File-based configuration** (Recommended)

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

3. **Environment variables** (Fallback)

   ```bash
   export VERACODE_API_ID="YOUR_API_KEY_ID"
   export VERACODE_API_KEY="YOUR_API_KEY_SECRET"
   ```

See [credentials/README.md](credentials/README.md) for detailed information.

## Usage

### Command Line Options

```bash
Options:
  -verbose
        Enable verbose logging to stderr (disabled by default)
  -log string
        Log file path for debugging (recommended for stdio mode)
  -version
        Display version information
```

**Important:** When using stdio mode with MCP clients (like VS Code or Claude Desktop), `-verbose` generates stderr output which can interfere with some JSON-RPC clients. If necessary, add `-log <filepath>` to write debug information to a file.

### Stdio Mode

The server runs in stdio mode for local integrations where it operates as a subprocess. This is the only supported mode as the server requires local filesystem access for workspace operations.

**Codex:**

via the command-line:
```bash
codex mcp add veracode -- "\path\to\veracode-mcp.exe"
```

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

**Claude CLI**

```bash
claude mcp add --transport stdio veracode "\path\to\veracode-mcp.exe"
```

**VS Code:**

```json
{
  "servers": {
    "veracode": {
      "command": "/path/to/veracode-mcp.exe",
      "cwd": "${workspaceFolder}",
      "args": ["-log", "/path/to/veracode-mcp.log"]
    },
  }
}
```

## Available MCP Tools

The server provides these Veracode-specific tools:

API:
- **api-health** - Verify Veracode API connectivity and credentials

Platform:
- **dynamic-findings** - Retrieve runtime security vulnerabilities from Dynamic Analysis (DAST) scans
- **static-findings** - Retrieve source code vulnerabilities from Static Analysis (SAST) scans
- **sca-findings** - Retrieve third-party component vulnerabilities from Software Composition Analysis
- **finding-details** - Get detailed information about a specific finding

Pipeline:
- **package-workspace** - Package workspace files for Veracode upload
- **pipeline-scan** - Start an asynchronous pipeline scan, with the largest packaged file as default
- **pipeline-status** - Check the status of a Pipeline Scan
- **pipeline-findings** - Get results from Veracode Pipeline Scans
- **pipeline-detailed-results** - Get detailed results from Pipeline Scans with full flaw information

SCA:
- **run-sca-scan** - Run Software Composition Analysis scan on a directory to identify vulnerable dependencies
- **local-sca-findings** - Read and parse local SCA scan results from veracode.json file
- **local-iac-findings** - Read and parse local IaC scan results (Dockerfile and configuration misconfigurations)

> **Note:** Use the `tools/list` MCP method to see all available tools with their complete parameter schemas and documentation.

---

## For Developers

See [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md) for building from source, running tests, adding new tools, and the full developer documentation index.

## Contributing

⚠️ This is alpha software. Contributions are welcome — see [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md). APIs and architecture may change significantly.
