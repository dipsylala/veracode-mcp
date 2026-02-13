package rest

import (
	"context"
	"fmt"
	"log"
)

// HealthStatus represents the result of a health check
type HealthStatus struct {
	Available  bool   `json:"available"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// CheckHealth verifies that Veracode API services are operational
// Returns the health status of the API
func (c *VeracodeClient) CheckHealth(ctx context.Context) (*HealthStatus, error) {
	// Get authenticated context
	authCtx := c.GetAuthContext(ctx)

	// Call the health check endpoint
	resp, err := c.healthcheckClient.HealthcheckAPIsAPI.HealthcheckStatusGet(authCtx).Execute()
	if err != nil {
		return &HealthStatus{
			Available:  false,
			Message:    fmt.Sprintf("Health check failed: %v", err),
			StatusCode: 0,
		}, nil // Return nil error so tools can handle gracefully
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("Failed to close response body: %v", closeErr)
		}
	}()

	// Parse response
	statusCode := resp.StatusCode
	available := statusCode == 200

	message := "API is operational"
	if !available {
		message = fmt.Sprintf("API returned status %d", statusCode)
	}

	return &HealthStatus{
		Available:  available,
		Message:    message,
		StatusCode: statusCode,
	}, nil
}

// CheckHealthSimple returns just a boolean indicating if the API is available
func (c *VeracodeClient) CheckHealthSimple(ctx context.Context) bool {
	status, err := c.CheckHealth(ctx)
	if err != nil {
		return false
	}
	return status.Available
}
