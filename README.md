# Veracode MCP Server (Go)

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
  - Remediation Guidance for Pipeline results
  - Finding details - Data paths and dynamic request/responses
  

## Installation

### Download from Releases

Download the latest pre-built binary from the [Releases page](https://github.com/dipsylala/veracode-mcp/releases):

Windows · macOS · Linux | x64 · ARM64

Extract and place the executable in a directory of your choice (e.g., `C:\Program Files\VeracodeMCP\` on Windows or `/usr/local/bin/` on macOS/Linux).

### Install Veracode CLI (Required)

Some tools (such as `package-workspace`, `pipeline-scan`, `run-sca-scan`) require the Veracode CLI to be installed and available in your system PATH.

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
.\path\to\veracode-mcp.exe [options]

Options:
  -verbose
        Enable verbose logging to stderr (disabled by default)
  -log string
        Log file path for debugging (recommended for stdio mode)
  -version
        Display version information
```

**Usage Examples:**

Bear in mind, this will typically not be run from the command-line directly, but part of the IDE configuration.

```bash
# Basic usage (silent mode, stdio transport)
.\path\to\veracode-mcp.exe

# With debug logging to file (recommended for troubleshooting)
.\path\path\to\veracode-mcp.exe -log \path\to\veracode-mcp.log

# With verbose logging to stderr (avoid in stdio mode as some MCP clients can react badly)
.\path\to\veracode-mcp.exe -verbose

```

**Important:** When using stdio mode with MCP clients (like VS Code or Claude Desktop), avoid using `-verbose` as stderr output can interfere with JSON-RPC communication. Instead, add `-log <filepath>` to write debug information to a file.

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

### VS Code: Veracode Analyst Agent

A pre-built VS Code agent is available that works with the MCP to provide AI-powered analysis of your Veracode findings. Where the MCP provides structured, LLM-optimised data retrieval, the agent provides the non-deterministic layer: risk prioritisation, cross-scan correlation, root cause grouping, and remediation planning.

**Check the feature is enabled in VS Code:**


**Install for a project:**

Download `veracode-analyst.agent.md` from the [Releases page](https://github.com/dipsylala/veracode-mcp/releases) and place it in your project:

```text
<your-project>/.github/agents/veracode-analyst.agent.md
```

**Install for personal use across all projects (VS Code Insiders):**

```powershell
# Windows
Copy-Item veracode-analyst.agent.md "$env:APPDATA\Code\User\agents\"
```

```bash
# macOS/Linux
cp veracode-analyst.agent.md "$HOME/Library/Application Support/Code/User/agents/"
```

After placing the file, reload VS Code (`Developer: Reload Window`). The **Veracode Analyst** agent will appear in the chat mode selector in the Copilot Chat panel.

**Usage:**

Select **Veracode Analyst** from the chat mode selector, or let Copilot spawn it automatically as a subagent when it determines it's appropriate:

```text
Analyse the security posture of /path/to/my/project and tell me what I should fix first.
```

The agent will check for findings, retrieve findings across SCA and Pipeline static scans, and synthesise a prioritised remediation plan. It requires the MCP server to be configured in VS Code settings as shown above.

### VS Code: Copilot Skills

Five task-focused skills are available for GitHub Copilot in VS Code. Each skill is a pre-built prompt that knows which MCP tools to call — they require the Veracode MCP server to be configured in VS Code.

| Skill | Trigger phrase | What it does |
|---|---|---|
| **scanit** | `#scanit` | Packages the workspace and starts a pipeline SAST scan |
| **thirdit** | `#thirdit` | Runs a local SCA scan on third-party dependencies |
| **reportit** | `#reportit` | Retrieves findings and produces a prioritised executive summary |
| **fixit** | `#fixit` | Retrieves remediation guidance and applies fixes for a specific flaw or CVE |
| **explainit** | `#explainit` | Explains a specific flaw or CVE in plain language |

#### Installing the skills

Skills are distributed as individual `.md` files on the [Releases page](https://github.com/dipsylala/veracode-mcp/releases).

**Install for a project** (team-shared, checked in to source control):

```text
<your-project>/.github/skills/scanit.skill.md
<your-project>/.github/skills/thirdit.skill.md
<your-project>/.github/skills/reportit.skill.md
<your-project>/.github/skills/fixit.skill.md
<your-project>/.github/skills/explainit.skill.md
```

**Install for personal use across all projects:**

*Windows:*
```powershell
# Copy all skill files to your user skills directory
Copy-Item *.skill.md "$env:APPDATA\Code\User\skills\"
```

*macOS/Linux:*
```bash
cp *.skill.md "$HOME/Library/Application Support/Code/User/skills/"
```

After placing the files, reload VS Code (`Developer: Reload Window`). Skills appear as `#scanit`, `#thirdit`, etc. in the Copilot Chat input.

#### Typical workflow

```text
# 1. Package and scan
@workspace #scanit scan /path/to/my/project

# 2. Check status (after scan completes)
@workspace #reportit summarise findings in /path/to/my/project

# 3. Fix a specific flaw
@workspace #fixit fix flaw 1026-1 in /path/to/my/project

# 4. Scan dependencies separately
@workspace #thirdit scan /path/to/my/project
```

> **Note:** Skills call MCP tools automatically based on your request. The MCP server must be running and configured in VS Code (see VS Code configuration above) for skills to work.

---

## Available MCP Tools

The server provides these Veracode-specific tools:

- **api-health** - Verify Veracode API connectivity and credentials
- **dynamic-findings** - Retrieve runtime security vulnerabilities from Dynamic Analysis (DAST) scans
- **static-findings** - Retrieve source code vulnerabilities from Static Analysis (SAST) scans
- **sca-findings** - Retrieve third-party component vulnerabilities from Software Composition Analysis
- **finding-details** - Get detailed information about a specific finding

- **package-workspace** - Package workspace files for Veracode upload
- **pipeline-scan** - Start an asynchronous pipeline scan, with the largest packaged file as default
- **pipeline-status** - Check the status of a Pipeline Scan
- **pipeline-findings** - Get results from Veracode Pipeline Scans
- **pipeline-detailed-results** - Get detailed results from Pipeline Scans with full flaw information
- **run-sca-scan** - Run Software Composition Analysis scan on a directory to identify vulnerable dependencies
- **local-sca-findings** - Read and parse local SCA scan results from veracode.json file

> **Note:** Use the `tools/list` MCP method to see all available tools with their complete parameter schemas and documentation.

### Remediation Guidance

The `remediation-guidance` tool provides CWE-specific, language-aware security guidance with code examples. It returns structured JSON containing:

- Vulnerability summary and key principles
- Step-by-step remediation instructions  
- Safe code patterns and examples
- Data flow information from the scan

> [!NOTE]
> **Quality Expectations:** The usefulness of remediation guidance depends heavily on how the AI assistant (LLM) interprets and applies the returned information. The tool provides structured security best practices and code samples, but the quality of the final code suggestions is determined by the capabilities of the LLM you're using (e.g., Claude, GPT-4, etc.). More capable models will better understand context and provide more accurate, applicable fixes.

---

## For Developers

See [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md) for building from source, running tests, adding new tools, and the full developer documentation index.

## License

MIT

## Contributing

⚠️ This is alpha software. Contributions are welcome — see [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md). APIs and architecture may change significantly.
