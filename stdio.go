package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// StdioTransport handles JSON-RPC over stdin/stdout
type StdioTransport struct {
	server *MCPServer
	reader *bufio.Reader
	writer io.Writer
	mu     sync.Mutex
}

func NewStdioTransport(server *MCPServer) *StdioTransport {
	return &StdioTransport{
		server: server,
		reader: bufio.NewReader(os.Stdin),
		writer: os.Stdout,
	}
}

func (t *StdioTransport) Start() error {
	for {
		line, err := t.reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("read error: %w", err)
		}

		// Skip empty lines
		if len(line) <= 1 {
			continue
		}

		var req JSONRPCRequest
		if err := json.Unmarshal(line, &req); err != nil {
			log.Printf("Failed to parse JSON-RPC request: %v", err)
			if sendErr := t.sendError(nil, -32700, "Parse error", nil); sendErr != nil {
				log.Printf("Failed to send parse error: %v", sendErr)
			}
			continue
		}

		go t.handleRequest(&req)
	}
}

func (t *StdioTransport) handleRequest(req *JSONRPCRequest) {
	log.Printf("=== INCOMING REQUEST ===")
	log.Printf("Method: %s (ID: %v)", req.Method, req.ID)
	if req.Params != nil {
		paramsJSON, _ := json.MarshalIndent(req.Params, "", "  ")
		log.Printf("Params:\n%s", string(paramsJSON))
	}
	log.Printf("========================")

	resp := t.server.HandleRequest(req)

	log.Printf("=== OUTGOING RESPONSE ===")
	log.Printf("ID: %v", resp.ID)
	responseJSON, _ := json.MarshalIndent(resp, "", "  ")
	log.Printf("Response:\n%s", string(responseJSON))
	log.Printf("========================")

	if err := t.sendResponse(resp); err != nil {
		log.Printf("Failed to send response: %v", err)
	}
}

func (t *StdioTransport) sendResponse(resp *JSONRPCResponse) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	data, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	data = append(data, '\n')
	_, err = t.writer.Write(data)
	return err
}

func (t *StdioTransport) sendError(id interface{}, code int, message string, data interface{}) error {
	resp := &JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &RPCError{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
	return t.sendResponse(resp)
}
