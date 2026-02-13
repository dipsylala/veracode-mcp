package mcp_tools

import (
	"context"
	"fmt"
	"time"

	"github.com/dipsylala/veracodemcp-go/api"
)

const APIHealthToolName = "api-health"

// Auto-register this tool when the package is imported
func init() {
	RegisterMCPTool(APIHealthToolName, handleAPIHealth)
}

// handleAPIHealth checks the health of Veracode API endpoints
func handleAPIHealth(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	timestamp := time.Now().Format(time.RFC3339)

	// Try to create API client (will check credentials)
	client, err := api.NewClient()
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
