package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const principalEndpoint = "/api/authn/v2/principal"

// CheckHealth verifies that Veracode API services are operational
// and that the configured credentials can authenticate.
func (c *Client) CheckHealth(ctx context.Context) (*HealthStatus, error) {
	ctx = c.GetAuthContext(ctx)

	endpoint := strings.TrimRight(c.baseURL, "/") + principalEndpoint
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return &HealthStatus{
			Available:  false,
			Message:    fmt.Sprintf("Failed to create principal request: %v", err),
			StatusCode: 0,
		}, nil
	}

	resp, err := newHMACHTTPClient(c.apiID, c.apiKey).Do(req)
	if err != nil {
		return &HealthStatus{
			Available:  false,
			Message:    fmt.Sprintf("Authenticated principal request failed: %v", err),
			StatusCode: 0,
		}, nil
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("Failed to close response body: %v", closeErr)
		}
	}()

	statusCode := resp.StatusCode
	available := statusCode == http.StatusOK

	message := "API is operational and authentication succeeded"
	if !available {
		message = fmt.Sprintf("Authenticated principal request returned status %d", statusCode)
	}

	return &HealthStatus{
		Available:  available,
		Message:    message,
		StatusCode: statusCode,
	}, nil
}

// CheckHealthSimple returns just a boolean indicating if the API is available
func (c *Client) CheckHealthSimple(ctx context.Context) bool {
	status, err := c.CheckHealth(ctx)
	if err != nil {
		return false
	}
	return status.Available
}
