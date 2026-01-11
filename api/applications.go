package api

import (
	"context"
	"fmt"

	applications "github.com/dipsylala/veracodemcp-go/api/generated/applications"
)

// GetApplication retrieves a single application by its GUID
func (c *VeracodeClient) GetApplication(ctx context.Context, applicationGUID string) (*applications.Application, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)

	app, httpResp, err := c.applicationsClient.ApplicationInformationAPIAPI.GetApplicationUsingGET(authCtx, applicationGUID).Execute()

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
func (c *VeracodeClient) GetApplicationByName(ctx context.Context, name string) (*applications.Application, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)

	// Search for the application by name
	resp, httpResp, err := c.applicationsClient.ApplicationInformationAPIAPI.GetApplicationsUsingGET(authCtx).
		Name(name).
		Execute()

	if err != nil {
		if httpResp != nil {
			return nil, fmt.Errorf("API returned status %d: %w", httpResp.StatusCode, err)
		}
		return nil, fmt.Errorf("failed to search applications: %w", err)
	}

	if resp == nil || resp.Embedded == nil || resp.Embedded.Applications == nil || len(resp.Embedded.Applications) == 0 {
		return nil, fmt.Errorf("application not found: %s", name)
	}

	return &resp.Embedded.Applications[0], nil
}

// ListApplications retrieves a paginated list of applications
func (c *VeracodeClient) ListApplications(ctx context.Context, page, size int) (*applications.PagedResourceOfApplication, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)

	req := c.applicationsClient.ApplicationInformationAPIAPI.GetApplicationsUsingGET(authCtx)

	if page >= 0 {
		req = req.Page(int32(page))
	}
	if size > 0 {
		req = req.Size(int32(size))
	}

	resp, httpResp, err := req.Execute()

	if err != nil {
		if httpResp != nil {
			return nil, fmt.Errorf("API returned status %d: %w", httpResp.StatusCode, err)
		}
		return nil, fmt.Errorf("failed to list applications: %w", err)
	}

	return resp, nil
}
