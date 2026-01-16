package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// version can be set at build time with -ldflags="-X main.version=x.y.z"
var version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "Display version information")
	mode := flag.String("mode", "stdio", "Server mode: stdio or http")
	addr := flag.String("addr", ":8080", "HTTP server address (only for http mode)")
	verbose := flag.Bool("verbose", false, "Enable verbose logging (disabled by default)")
	logFile := flag.String("log", "", "Log file path (if not specified, logs go to stderr when verbose)")
	flag.Parse()

	// Setup logging based on flags
	if *logFile != "" {
		// Write logs to specified file
		f, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open log file %s: %v\n", *logFile, err)
			os.Exit(1)
		}
		defer f.Close()
		log.SetOutput(f)
	} else if !*verbose {
		// Disable logging by default if no log file and not verbose
		log.SetOutput(io.Discard)
	}
	// If verbose and no log file, logging goes to stderr (default)
	if *showVersion {
		fmt.Fprintf(os.Stderr, "veracode-mcp-server version %s\n", version)
		os.Exit(0)
	}

	server, err := NewMCPServer()
	if err != nil {
		// Always show server creation errors to stderr, even in non-verbose mode
		// This is before stdio transport starts, so it won't interfere with JSON-RPC
		fmt.Fprintf(os.Stderr, "Failed to create MCP server: %v\n", err)
		os.Exit(1)
	}

	switch *mode {
	case "stdio":
		log.Println("Starting MCP server in stdio mode...")
		if err := server.ServeStdio(); err != nil {
			if *verbose {
				log.Fatalf("stdio server error: %v", err)
			} else {
				os.Exit(1)
			}
		}
	case "http":
		log.Printf("Starting MCP server in http mode on %s...\n", *addr)
		if err := server.ServeHTTP(*addr); err != nil {
			if *verbose {
				log.Fatalf("http server error: %v", err)
			} else {
				fmt.Fprintf(os.Stderr, "http server error: %v\n", err)
				os.Exit(1)
			}
		}
	default:
		fmt.Fprintf(os.Stderr, "Invalid mode: %s. Use 'stdio' or 'http'\n", *mode)
		os.Exit(1)
	}
}
