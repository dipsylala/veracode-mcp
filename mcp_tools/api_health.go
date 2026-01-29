package tools

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dipsylala/veracodemcp-go/api"
)

const APIHealthToolName = "api-health"

// Auto-register this tool when the package is imported
func init() {
	RegisterTool(APIHealthToolName, func() ToolImplementation {
		return NewAPIHealthTool()
	})
}

// APIHealthTool provides the api-health tool
type APIHealthTool struct{}

// NewAPIHealthTool creates a new API health check tool
func NewAPIHealthTool() *APIHealthTool {
	return &APIHealthTool{}
}

// Initialize sets up the tool
func (t *APIHealthTool) Initialize() error {
	log.Printf("Initializing tool: %s", APIHealthToolName)
	return nil
}

// RegisterHandlers registers the API health check handler
func (t *APIHealthTool) RegisterHandlers(registry HandlerRegistry) error {
	log.Printf("Registering handlers for tool: %s", APIHealthToolName)
	registry.RegisterHandler(APIHealthToolName, t.handleAPIHealth)
	return nil
}

// Shutdown cleans up tool resources
func (t *APIHealthTool) Shutdown() error {
	log.Printf("Shutting down tool: %s", APIHealthToolName)
	return nil
}

// APIHealthRequest represents the parsed parameters for api-health
type APIHealthRequest struct {
	// Currently no parameters, but struct provides consistency for future extension
}

// parseAPIHealthRequest extracts and validates parameters from the raw args map
func parseAPIHealthRequest(args map[string]interface{}) (*APIHealthRequest, error) {
	req := &APIHealthRequest{}
	// No parameters currently required
	return req, nil
}

// handleAPIHealth checks the health of Veracode API endpoints
func (t *APIHealthTool) handleAPIHealth(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	// Parse and validate request parameters
	_, err := parseAPIHealthRequest(args)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, nil
	}
	timestamp := time.Now().Format(time.RFC3339)

	// Try to create API client (will check credentials)
	client, err := api.NewVeracodeClient()
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Veracode API Health Check
========================

Timestamp: %s

❌ Authentication: Not configured
Error: %v

Required credentials:
- ~/.veracode/veracode.yml with key-id and key-secret
- OR environment variables VERACODE_API_ID and VERACODE_API_KEY

Please configure credentials and try again.`, timestamp, err),
			}},
		}, nil
	}

	// Perform actual health check
	healthStatus, err := client.CheckHealth(ctx)
	if err != nil {
		return map[string]interface{}{
			"content": []map[string]string{{
				"type": "text",
				"text": fmt.Sprintf(`Veracode API Health Check
========================

Timestamp: %s

❌ Health check failed: %v`, timestamp, err),
			}},
		}, nil
	}

	availableIcon := "✓"
	if !healthStatus.Available {
		availableIcon = "❌"
	}

	return map[string]interface{}{
		"content": []map[string]string{{
			"type": "text",
			"text": fmt.Sprintf(`Veracode API Health Check
========================

Timestamp: %s

%s Veracode API (api.veracode.com)
Status: %s
HTTP Status: %d

✓ Authentication: Configured

Next steps:
- Run dynamic-findings or static-findings to fetch data
- Check application access permissions
- Review API rate limits and quotas`, timestamp, availableIcon, healthStatus.Message, healthStatus.StatusCode),
		}},
	}, nil
}
