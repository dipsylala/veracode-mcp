# Contributing & Building from Source

## Building from Source

```bash
# Clone the repository
git clone https://github.com/dipsylala/veracode-mcp.git
cd veracode-mcp

# Install dependencies
go mod download

# Build the server
.\build.ps1    # Windows (builds to dist/veracode-mcp.exe)
# or
go build -o dist/veracode-mcp.exe .  # Manual build
```

The PowerShell build script runs quality checks before compiling:

```powershell
.\build.ps1          # Full build with all checks
.\build.ps1 -Quick   # Fast build, skip checks
.\build.ps1 -NoTest  # Build without running tests
.\build.ps1 -Verbose # Show detailed test output
```

Steps performed:

1. Code formatting (`go fmt`)
2. Static analysis (`go vet`)
3. Linting (`golangci-lint` if available)
4. Unit tests
5. Binary compilation to `dist/veracode-mcp.exe`

## Testing

```bash
go test ./...        # Run all tests
go test ./... -v     # Verbose output
go test ./tools/...  # Test only tools package
```

See [MCP_TESTING.md](MCP_TESTING.md) for MCP integration testing.

## Adding New Tools

See [ADDING_TOOLS.md](ADDING_TOOLS.md) for the complete guide on implementing new MCP tools.

## Developer Documentation

- **[Architecture & Design](DESIGN.md)** - System architecture and design decisions
- **[Adding Tools](ADDING_TOOLS.md)** - Quick start and full guide for adding new tools
- **[API Integration](API_INTEGRATION.md)** - Integrate Veracode REST APIs
- **[Code Quality](CODE_QUALITY.md)** - Build tools and quality checks
- **[Tool Structure](TOOLS_STRUCTURE.md)** - Tool system architecture
- **[MCP Testing](MCP_TESTING.md)** - Testing MCP implementations

## Contributing

⚠️ This is alpha software. Contributions are welcome, but please be aware that APIs and architecture may change significantly.

Please feel free to submit issues and pull requests.
