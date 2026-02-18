package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/dipsylala/veracode-mcp/internal/cli"
	"github.com/dipsylala/veracode-mcp/internal/server"
	tools "github.com/dipsylala/veracode-mcp/internal/tool_registry"
)

//go:embed tools.json
var toolsJSONData []byte

//go:embed instructions.json
var instructionsJSONData []byte

//go:embed ui/pipeline-findings-app/dist/mcp-app.html
var pipelineFindingsHTML string

//go:embed ui/static-findings-app/dist/mcp-app.html
var staticFindingsHTML string

//go:embed ui/dynamic-findings-app/dist/mcp-app.html
var dynamicFindingsHTML string

//go:embed ui/local-sca-findings-app/dist/mcp-app.html
var localSCAResultsHTML string

func init() {
	// Set embedded resources in the internal packages
	tools.SetToolsJSON(toolsJSONData)
	server.SetInstructions(instructionsJSONData)
	server.SetUIResources(pipelineFindingsHTML, staticFindingsHTML, dynamicFindingsHTML, localSCAResultsHTML)
}

// version can be set at build time with -ldflags="-X main.version=x.y.z"
var version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "Display version information")
	verbose := flag.Bool("verbose", false, "Enable verbose logging (disabled by default)")
	logFile := flag.String("log", "", "Log file path (if not specified, logs go to stderr when verbose)")
	flag.Parse()

	if *showVersion {
		fmt.Fprintf(os.Stderr, "veracode-mcp-server version %s\n", version)
		os.Exit(0)
	}

	// Configure logging based on flags
	if err := cli.ConfigureLogging(*logFile, *verbose); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	mcpServer, err := server.NewMCPServer()
	if err != nil {
		// Always show server creation errors to stderr, even in non-verbose mode
		// This is before stdio transport starts, so it won't interfere with JSON-RPC
		fmt.Fprintf(os.Stderr, "Failed to create MCP server: %v\n", err)
		os.Exit(1)
	}

	if err := cli.RunServer(mcpServer, *verbose); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
