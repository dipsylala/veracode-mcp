package transport

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"unicode"

	"github.com/dipsylala/veracodemcp-go/internal/types"
)

// StdioTransport handles JSON-RPC over stdin/stdout
type StdioTransport struct {
	handler RequestHandler
	reader  *bufio.Reader
	writer  io.Writer
	mu      sync.Mutex
}

func NewStdioTransport(handler RequestHandler) *StdioTransport {
	return &StdioTransport{
		handler: handler,
		reader:  bufio.NewReader(os.Stdin),
		writer:  bufio.NewWriter(os.Stdout),
	}
}

func (t *StdioTransport) Start() error {
	for {
		line, err := t.reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("stdio: received EOF, client disconnected")
				return nil
			}
			log.Printf("stdio: read error: %v", err)
			return fmt.Errorf("read error: %w", err)
		}

		// Skip empty lines
		if len(line) <= 1 {
			continue
		}

		log.Printf("stdio: received message: %s", string(line[:min(len(line), 100)]))

		var req types.JSONRPCRequest
		if err := json.Unmarshal(line, &req); err != nil {
			log.Printf("stdio: parse error for message: %v", err)
			// Send parse error without logging to avoid stdio interference
			_ = t.sendError(nil, -32700, "Parse error", nil)
			continue
		}

		go t.handleRequest(&req)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// isValidMethod checks if a method name contains only alphanumeric characters and forward slashes
// to prevent log forging attacks via CRLF injection
func isValidMethod(method string) bool {
	for _, r := range method {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '/' {
			return false
		}
	}
	return true
}

func (t *StdioTransport) handleRequest(req *types.JSONRPCRequest) {
	// Validate method name before logging to prevent log forging attacks
	if !isValidMethod(req.Method) {
		log.Printf("=== INCOMING REQUEST - REJECTED ===")
		log.Printf("Invalid method name (ID: %v)", req.ID)
		log.Printf("========================")
		// Send error response for invalid method
		_ = t.sendError(req.ID, -32600, "Invalid request: method name contains invalid characters", nil)
		return
	}

	log.Printf("=== INCOMING REQUEST ===")
	log.Printf("Method: %s (ID: %v)", req.Method, req.ID)
	if req.Params != nil {
		paramsJSON, _ := json.MarshalIndent(req.Params, "", "  ")
		log.Printf("Params:\n%s", string(paramsJSON))
	}
	log.Printf("========================")

	resp := t.handler.HandleRequest(req)

	// Only send response if one was returned (notifications return nil)
	if resp != nil {
		log.Printf("=== OUTGOING RESPONSE ===")
		log.Printf("ID: %v", resp.ID)
		responseJSON, _ := json.MarshalIndent(resp, "", "  ")
		log.Printf("Response:\n%s", string(responseJSON))
		log.Printf("========================")

		if err := t.sendResponse(resp); err != nil {
			log.Printf("Failed to send response: %v", err)
		}
	}
}

func (t *StdioTransport) sendResponse(resp *types.JSONRPCResponse) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Debug: Check for structuredContent before marshaling
	if resp.Result != nil {
		if ctr, ok := resp.Result.(*types.CallToolResult); ok {
			log.Printf("[STDIO] Before JSON marshal: hasStructuredContent=%v", ctr.StructuredContent != nil)
		}
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("stdio: marshal error: %v", err)
		return fmt.Errorf("marshal error: %w", err)
	}

	// Debug: Check if structuredContent is in the JSON
	if resp.Result != nil {
		log.Printf("[STDIO] JSON output length: %d bytes", len(data))
		jsonStr := string(data)
		hasStructuredContent := contains(jsonStr, `"structuredContent":`)
		log.Printf("[STDIO] JSON actually contains 'structuredContent': %v", hasStructuredContent)
		if len(data) < 2000 {
			log.Printf("[STDIO] JSON output: %s", jsonStr)
		} else {
			log.Printf("[STDIO] JSON preview (first 2000 chars): %s", jsonStr[0:min(2000, len(data))])
		}
	}

	log.Printf("stdio: sending response (first 100 chars): %s", string(data[:min(len(data), 100)]))

	data = append(data, '\n')
	_, err = t.writer.Write(data)
	if err != nil {
		log.Printf("stdio: write error: %v", err)
		return err
	}

	// Ensure response is flushed immediately
	if flusher, ok := t.writer.(interface{ Flush() error }); ok {
		if err := flusher.Flush(); err != nil {
			log.Printf("stdio: flush error: %v", err)
			return err
		}
		log.Println("stdio: response flushed successfully")
	}

	return nil
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func (t *StdioTransport) sendError(id interface{}, code int, message string, data interface{}) error {
	resp := &types.JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &types.RPCError{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
	return t.sendResponse(resp)
}
