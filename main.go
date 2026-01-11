package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	mode := flag.String("mode", "stdio", "Server mode: stdio or http")
	addr := flag.String("addr", ":8080", "HTTP server address (only for http mode)")
	flag.Parse()

	server, err := NewMCPServer()
	if err != nil {
		log.Fatalf("Failed to create MCP server: %v", err)
	}

	switch *mode {
	case "stdio":
		log.Println("Starting MCP server in stdio mode...")
		if err := server.ServeStdio(); err != nil {
			log.Fatalf("stdio server error: %v", err)
		}
	case "http":
		log.Printf("Starting MCP server in http mode on %s...\n", *addr)
		if err := server.ServeHTTP(*addr); err != nil {
			log.Fatalf("http server error: %v", err)
		}
	default:
		fmt.Fprintf(os.Stderr, "Invalid mode: %s. Use 'stdio' or 'http'\n", *mode)
		os.Exit(1)
	}
}
