package cli

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dipsylala/veracode-mcp/internal/server"
	tools "github.com/dipsylala/veracode-mcp/internal/tool_registry"
)

// AppConfig holds the application configuration
type AppConfig struct {
	Version             string
	ToolsJSON           []byte
	PipelineResultsHTML string
	StaticFindingsHTML  string
	DynamicFindingsHTML string
}

// ConfigureLogging sets up logging based on command line flags
func ConfigureLogging(logFilePath string, verbose bool) error {
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

// RunServer starts the MCP server in stdio mode
func RunServer(s *server.MCPServer, verbose bool) error {
	log.Println("Starting MCP server in stdio mode...")
	if err := s.ServeStdio(); err != nil {
		if verbose {
			log.Fatalf("stdio server error: %v", err)
		}
		return err
	}
	return nil
}

// InitializeResources sets the embedded resources in internal packages
func InitializeResources(config AppConfig) {
	tools.SetToolsJSON(config.ToolsJSON)
	server.SetUIResources(config.PipelineResultsHTML, config.StaticFindingsHTML, config.DynamicFindingsHTML)
}
