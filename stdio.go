package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
			// Send parse error without logging to avoid stdio interference
			_ = t.sendError(nil, -32700, "Parse error", nil)
			continue
		}

		go t.handleRequest(&req)
	}
}

func (t *StdioTransport) handleRequest(req *JSONRPCRequest) {
	resp := t.server.HandleRequest(req)
	_ = t.sendResponse(resp)
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
