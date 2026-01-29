package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/dipsylala/veracodemcp-go/internal/cli"
	"github.com/dipsylala/veracodemcp-go/internal/server"
	"github.com/dipsylala/veracodemcp-go/internal/tools"
)

//go:embed tools.json
var toolsJSONData []byte

//go:embed ui/pipeline-results-app/dist/mcp-app.html
var pipelineResultsHTML string

//go:embed ui/static-findings-app/dist/mcp-app.html
var staticFindingsHTML string

//go:embed ui/dynamic-findings-app/dist/mcp-app.html
var dynamicFindingsHTML string

func init() {
	// Set embedded resources in the internal packages
	tools.SetToolsJSON(toolsJSONData)
	server.SetUIResources(pipelineResultsHTML, staticFindingsHTML, dynamicFindingsHTML)
}

// version can be set at build time with -ldflags="-X main.version=x.y.z"
var version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "Display version information")
	mode := flag.String("mode", "stdio", "Server mode: stdio or http")
	addr := flag.String("addr", ":8080", "HTTP server address (only for http mode)")
	verbose := flag.Bool("verbose", false, "Enable verbose logging (disabled by default)")
	logFile := flag.String("log", "", "Log file path (if not specified, logs go to stderr when verbose)")
	forceMCPApp := flag.Bool("force-mcp-app", false, "Force MCP Apps mode (always send structuredContent regardless of client capabilities)")
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

	mcpServer, err := server.NewMCPServer(*forceMCPApp)
	if err != nil {
		// Always show server creation errors to stderr, even in non-verbose mode
		// This is before stdio transport starts, so it won't interfere with JSON-RPC
		fmt.Fprintf(os.Stderr, "Failed to create MCP server: %v\n", err)
		os.Exit(1)
	}

	if err := cli.RunServer(mcpServer, *mode, *addr, *verbose); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
