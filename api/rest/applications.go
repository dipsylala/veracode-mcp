package rest

import (
	"context"
	"fmt"
	"log"
	"math"
	"strings"

	applications "github.com/dipsylala/veracode-mcp/api/rest/generated/applications"
)

// GetApplication retrieves a single application by its GUID
func (c *Client) GetApplication(ctx context.Context, applicationGUID string) (*applications.Application, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)

	app, httpResp, err := c.applicationsClient.ApplicationInformationAPIAPI.GetApplicationUsingGET(authCtx, applicationGUID).Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() {
			if closeErr := httpResp.Body.Close(); closeErr != nil {
				log.Printf("Failed to close response body: %v", closeErr)
			}
		}()
	}

	if err != nil {
		if httpResp != nil {
			return nil, fmt.Errorf("API returned status %d: %w", httpResp.StatusCode, err)
		}
		return nil, fmt.Errorf("failed to get application: %w", err)
	}

	return app, nil
}

// GetApplicationByName retrieves an application by searching for its name
// Returns the first matching application or nil if not found
func (c *Client) GetApplicationByName(ctx context.Context, name string) (*applications.Application, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)

	// Search for the application by name
	resp, httpResp, err := c.applicationsClient.ApplicationInformationAPIAPI.GetApplicationsUsingGET(authCtx).
		Name(name).
		Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() {
			if closeErr := httpResp.Body.Close(); closeErr != nil {
				log.Printf("Failed to close response body: %v", closeErr)
			}
		}()
	}

	if err != nil {
		if httpResp != nil {
			return nil, fmt.Errorf("API returned status %d: %w", httpResp.StatusCode, err)
		}
		return nil, fmt.Errorf("failed to search applications: %w", err)
	}

	if resp == nil || resp.Embedded == nil || resp.Embedded.Applications == nil || len(resp.Embedded.Applications) == 0 {
		return nil, fmt.Errorf("application not found: %s", name)
	}

	// If multiple results returned (substring search), find exact match (case-insensitive)
	if len(resp.Embedded.Applications) > 1 {
		app := findExactApplicationMatch(resp.Embedded.Applications, name)
		if app != nil {
			return app, nil
		}
		return nil, fmt.Errorf("no exact match found for application: %s (found %d partial matches)", name, len(resp.Embedded.Applications))
	}

	return &resp.Embedded.Applications[0], nil
}

// equalFoldStrings compares two strings case-insensitively
func equalFoldStrings(a, b string) bool {
	return len(a) == len(b) && strings.EqualFold(a, b)
}

// findExactApplicationMatch searches for an application with an exact name match (case-insensitive)
func findExactApplicationMatch(apps []applications.Application, name string) *applications.Application {
	for _, app := range apps {
		if app.Profile != nil && app.Profile.Name != nil {
			if equalFoldStrings(*app.Profile.Name, name) {
				return &app
			}
		}
	}
	return nil
}

// ListApplications retrieves a paginated list of applications
func (c *Client) ListApplications(ctx context.Context, page, size int) (*applications.PagedResourceOfApplication, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)

	req := c.applicationsClient.ApplicationInformationAPIAPI.GetApplicationsUsingGET(authCtx)

	if page >= 0 {
		if page > math.MaxInt32 {
			return nil, fmt.Errorf("page value %d exceeds maximum allowed value", page)
		}
		req = req.Page(int32(page)) // #nosec G115 - validated above
	}
	if size > 0 {
		if size > math.MaxInt32 {
			return nil, fmt.Errorf("size value %d exceeds maximum allowed value", size)
		}
		req = req.Size(int32(size)) // #nosec G115 - validated above
	}

	resp, httpResp, err := req.Execute()
	if httpResp != nil && httpResp.Body != nil {
		defer func() {
			if closeErr := httpResp.Body.Close(); closeErr != nil {
				log.Printf("Failed to close response body: %v", closeErr)
			}
		}()
	}

	if err != nil {
		if httpResp != nil {
			return nil, fmt.Errorf("API returned status %d: %w", httpResp.StatusCode, err)
		}
		return nil, fmt.Errorf("failed to list applications: %w", err)
	}

	return resp, nil
}
