package rest

import (
	"context"
	"fmt"
	"log"

	"github.com/dipsylala/veracode-mcp/api/rest/generated/policy"
)

// GetPolicy retrieves policies matching the exact name, returning the full paged response
func (c *Client) GetPolicy(ctx context.Context, policyName string) (*policy.PagedResourceOfPolicyVersion, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)

	pagedResult, httpResp, err := c.policyClient.PolicyInformationAPIAPI.
		GetPoliciesUsingGET(authCtx).
		Name(policyName).
		NameExact(true).
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
		return nil, fmt.Errorf("failed to list policies: %w", err)
	}

	if pagedResult == nil || pagedResult.Embedded == nil || len(pagedResult.Embedded.PolicyVersions) == 0 {
		return nil, fmt.Errorf("policy not found: %q", policyName)
	}

	return pagedResult, nil
}
