# Veracode MCP Server

> [!NOTE]
> This is not associated with Veracode, and does not fall under their support

⚠️ **BETA** - After running this regularly for months on multiple applications, I'm happy to move this to Beta. This is still early-stage software under active development, but the tools and general functionality have stabilised.

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

For detailed installation instructions and alternative methods, see the [official Veracode CLI installation guide](https://docs.veracode.com/r/Install_the_Veracode_CLI).

## Configuration

### Veracode API Credentials

**veracode.yml**

Follow the instructions as per the [Veracode Docs](https://docs.veracode.com/r/Install_the_Veracode_CLI#store-credentials-with-the-cli).

**Environment variables** (Fallback)

If you wish to use the MCP server without the `veracode.yml`, it can also accept credentials [via the environment](https://docs.veracode.com/r/Install_the_Veracode_CLI?current-os=windows#configure-credentials-as-environment-variables)

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

### Veracode platform work

To assign a Veracode profile to code you're working on, create a `.veracode-workspace.json` file with the following contents in your code area:

<img width="742" height="502" alt="image" src="https://github.com/user-attachments/assets/6f67f0fe-cb10-45c4-8fd0-f80dffc870f5" />

This will allow the MCP to know which profile you're focusing on in your IDE/TUI.

The contents are as follows:

```json
{
  "name": "{profile name}"
}
```

example:

```json
{
  "name": "Verademo"
}
```

---

## Available MCP Tools

The server provides these Veracode-specific tools:

API:

- **api-health** - Verify API connectivity and HMAC credentials with the Identity API principal endpoint

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
- **local-sca-summary** - Group local SCA findings by component showing the minimum upgrade version to fix all CVEs
- **local-sca-findings** - Read and parse local SCA scan results from veracode.json file
- **local-iac-findings** - Read and parse local IaC scan results (Dockerfile and configuration misconfigurations)

> **Note:** Use the `tools/list` MCP method to see all available tools with their complete parameter schemas and documentation.

---

## But MCP and my context window…

If you want some of the benefits of the MCP but without the Context window overhead, I wrote some skills that can help out instead. They're more results-focused, and you can pick and choose which ones you want. So if you're happy packaging, and maybe your scanning is being handled elsewhere (CI/CD, pipeline, platform scanning), they'll handle the final hurdle: getting the results into your LLM.

[veracode-pipeline-results](https://github.com/dipsylala/veracode-pipeline-results) - designed to work on the json results output by the pipeline scan.

[veracode-local-sca-results](https://github.com/dipsylala/veracode-local-sca-results) - doing a local SCA/IaC scan? This can provide a summary, as well as helping you dig deeper on individual results.

[veracode-platform-results](https://github.com/dipsylala/veracode-platform-results) - this pulls SCA, Dynamic and Static data from the platform.

In each case, these are designed to look after your context window - `veracode-pipeline-results` and `veracode-local-sca-results` use Python scripts to translate the JSON into something more LLM-convenient. `veracode-platform-results` uses my [veracode-api](https://github.com/dipsylala/veracode-api) CLI tool to use your existing Veracode credentials to authenticate to the platform and query the API. 

---

## For Developers

See [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md) for building from source, running tests, adding new tools, and the full developer documentation index.

## Contributing

⚠️ This is beta software. Contributions are welcome — see [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md).
