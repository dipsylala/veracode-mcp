package transport

import (
	"github.com/dipsylala/veracodemcp-go/internal/types"
)

// RequestHandler defines the interface that MCP servers must implement
// to handle JSON-RPC requests. This allows the transport layer to be
// decoupled from the specific server implementation.
type RequestHandler interface {
	HandleRequest(req *types.JSONRPCRequest) *types.JSONRPCResponse
	ClientSupportsUI() bool
}
