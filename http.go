package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

// HTTPTransport handles JSON-RPC over HTTP with SSE
type HTTPTransport struct {
	server  *MCPServer
	clients map[string]*SSEClient
	mu      sync.RWMutex
}

type SSEClient struct {
	id       string
	messages chan *JSONRPCResponse
	done     chan struct{}
}

func NewHTTPTransport(server *MCPServer) *HTTPTransport {
	return &HTTPTransport{
		server:  server,
		clients: make(map[string]*SSEClient),
	}
}

func (t *HTTPTransport) Start(addr string) error {
	http.HandleFunc("/sse", t.handleSSE)
	http.HandleFunc("/message", t.handleMessage)
	http.HandleFunc("/health", t.handleHealth)

	log.Printf("HTTP server listening on %s", addr)
	return http.ListenAndServe(addr, nil)
}

// handleSSE establishes Server-Sent Events connection
func (t *HTTPTransport) handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Create new client
	client := &SSEClient{
		id:       uuid.New().String(),
		messages: make(chan *JSONRPCResponse, 10),
		done:     make(chan struct{}),
	}

	t.mu.Lock()
	t.clients[client.id] = client
	t.mu.Unlock()

	defer func() {
		t.mu.Lock()
		delete(t.clients, client.id)
		t.mu.Unlock()
		close(client.done)
	}()

	// Send client ID
	fmt.Fprintf(w, "event: endpoint\ndata: /message?sessionId=%s\n\n", client.id)
	flusher.Flush()

	// Keep connection alive and send messages
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			return
		case msg := <-client.messages:
			data, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Failed to marshal message: %v", err)
				continue
			}
			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		case <-ticker.C:
			// Send keep-alive ping
			fmt.Fprintf(w, ": ping\n\n")
			flusher.Flush()
		}
	}
}

// handleMessage processes incoming JSON-RPC requests
func (t *HTTPTransport) handleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sessionID := r.URL.Query().Get("sessionId")
	if sessionID == "" {
		http.Error(w, "Missing sessionId", http.StatusBadRequest)
		return
	}

	t.mu.RLock()
	client, exists := t.clients[sessionID]
	t.mu.RUnlock()

	if !exists {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	var req JSONRPCRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON-RPC request", http.StatusBadRequest)
		return
	}

	// Process request
	resp := t.server.HandleRequest(&req)

	// Send response via SSE
	select {
	case client.messages <- resp:
		w.WriteHeader(http.StatusAccepted)
	case <-time.After(5 * time.Second):
		http.Error(w, "Timeout sending response", http.StatusRequestTimeout)
	}
}

// handleHealth provides a simple health check endpoint
func (t *HTTPTransport) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
