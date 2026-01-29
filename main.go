package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dipsylala/veracodemcp-go/internal/server"
	"github.com/dipsylala/veracodemcp-go/internal/tools"
)

//go:embed tools.json
var toolsJSONData []byte

//go:embed ui/pipeline-results-app/mcp-app.html
var pipelineResultsHTML string

//go:embed ui/static-findings-app/mcp-app.html
var staticFindingsHTML string

//go:embed ui/dynamic-findings-app/mcp-app.html
var dynamicFindingsHTML string

func init() {
	// Set embedded resources in the internal packages
	tools.SetToolsJSON(toolsJSONData)
	server.SetUIResources(pipelineResultsHTML, staticFindingsHTML, dynamicFindingsHTML)
}

// version can be set at build time with -ldflags="-X main.version=x.y.z"
var version = "dev"

// configureLogging sets up logging based on command line flags
func configureLogging(logFilePath string, verbose bool) error {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	if logFilePath != "" {
		// Write logs to specified file
		// #nosec G304 -- logFilePath is from command-line flag, user controls log destination
		f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return fmt.Errorf("failed to open log file %s: %w", logFilePath, err)
		}
		log.SetOutput(f)
	} else if !verbose {
		// Disable logging if not verbose and no log file
		log.SetOutput(io.Discard)
	}
	// If verbose and no log file, logging goes to stderr (default)
	return nil
}

// runServer starts the MCP server in the specified mode
func runServer(s *server.MCPServer, mode, addr string, verbose bool) error {
	switch mode {
	case "stdio":
		log.Println("Starting MCP server in stdio mode...")
		if err := s.ServeStdio(); err != nil {
			if verbose {
				log.Fatalf("stdio server error: %v", err)
			}
			return err
		}
	case "http":
		log.Printf("Starting MCP server in http mode on %s...\n", addr)
		if err := s.ServeHTTP(addr); err != nil {
			if verbose {
				log.Fatalf("http server error: %v", err)
			}
			return err
		}
	default:
		return fmt.Errorf("invalid mode: %s. Use 'stdio' or 'http'", mode)
	}
	return nil
}

func main() {
	showVersion := flag.Bool("version", false, "Display version information")
	mode := flag.String("mode", "stdio", "Server mode: stdio or http")
	addr := flag.String("addr", ":8080", "HTTP server address (only for http mode)")
	verbose := flag.Bool("verbose", false, "Enable verbose logging (disabled by default)")
	logFile := flag.String("log", "", "Log file path (if not specified, logs go to stderr when verbose)")
	flag.Parse()

	if *showVersion {
		fmt.Fprintf(os.Stderr, "veracode-mcp-server version %s\n", version)
		os.Exit(0)
	}

	// Configure logging based on flags
	if err := configureLogging(*logFile, *verbose); err != nil {
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

	if err := runServer(mcpServer, *mode, *addr, *verbose); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
