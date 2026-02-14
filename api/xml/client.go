package xml

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/dipsylala/veracode-mcp/credentials"
	veracodehmac "github.com/dipsylala/veracode-mcp/hmac"
)

const (
	// Base URL for the Veracode XML API
	DefaultBaseURL = "https://analysiscenter.veracode.com/api"
)

// Client wraps HTTP client with HMAC authentication for Veracode XML API
type Client struct {
	apiID   string
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewClient creates a new Veracode XML API client
// Credentials are loaded from:
// 1. ~/.veracode/veracode.yml (preferred)
// 2. Environment variables VERACODE_API_ID and VERACODE_API_KEY (fallback)
func NewClient() (*Client, error) {
	apiID, apiKey, _, err := credentials.GetCredentials()
	if err != nil {
		return nil, err
	}

	return &Client{
		apiID:   apiID,
		apiKey:  apiKey,
		baseURL: DefaultBaseURL,
		client:  http.DefaultClient,
	}, nil
}

// NewClientUnconfigured creates a client without checking credentials
// Useful for testing or when credentials will be set later
func NewClientUnconfigured() *Client {
	apiID, apiKey, _, _ := credentials.GetCredentials()

	return &Client{
		apiID:   apiID,
		apiKey:  apiKey,
		baseURL: DefaultBaseURL,
		client:  http.DefaultClient,
	}
}

// IsConfigured returns true if API credentials are set
func (c *Client) IsConfigured() bool {
	return c.apiID != "" && c.apiKey != ""
}

// doRequest performs an HTTP request with HMAC authentication
func (c *Client) doRequest(ctx context.Context, method, endpoint string, params url.Values) ([]byte, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured")
	}

	// Build full URL with query parameters
	fullURL := c.baseURL + endpoint
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	// Parse URL for HMAC calculation
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	// Normalize query string to use %20 instead of + for spaces
	if parsedURL.RawQuery != "" {
		normalizedQuery := strings.ReplaceAll(parsedURL.RawQuery, "+", "%20")
		parsedURL.RawQuery = normalizedQuery
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, parsedURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Calculate HMAC authorization header
	authHeader, err := veracodehmac.CalculateAuthorizationHeader(parsedURL, method, c.apiID, c.apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate HMAC authorization: %w", err)
	}

	// Add authorization header
	req.Header.Set("Authorization", authHeader)

	// Execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// GetMitigationInfo retrieves mitigation information for specific flaws in a build
// buildID: the build ID to query
// flawIDs: list of flaw IDs to retrieve mitigation info for
func (c *Client) GetMitigationInfo(ctx context.Context, buildID int64, flawIDs []int64) (*MitigationInfo, error) {
	if len(flawIDs) == 0 {
		return nil, fmt.Errorf("at least one flaw ID is required")
	}

	// Convert flaw IDs to comma-delimited string
	flawIDStrs := make([]string, len(flawIDs))
	for i, id := range flawIDs {
		flawIDStrs[i] = fmt.Sprintf("%d", id)
	}
	flawIDList := strings.Join(flawIDStrs, ",")

	// Build query parameters
	params := url.Values{}
	params.Set("build_id", fmt.Sprintf("%d", buildID))
	params.Set("flaw_id_list", flawIDList)

	// Make request
	body, err := c.doRequest(ctx, "GET", "/getmitigationinfo.do", params)
	if err != nil {
		return nil, err
	}

	// Parse XML response
	var mitigationInfo MitigationInfo
	if err := xml.Unmarshal(body, &mitigationInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML response: %w", err)
	}

	return &mitigationInfo, nil
}

// GetMitigationInfoForSingleFlaw is a convenience method to get mitigation info for a single flaw
func (c *Client) GetMitigationInfoForSingleFlaw(ctx context.Context, buildID, flawID int64) (*Issue, error) {
	mitigationInfo, err := c.GetMitigationInfo(ctx, buildID, []int64{flawID})
	if err != nil {
		return nil, err
	}

	// Check for errors
	if len(mitigationInfo.Errors) > 0 {
		return nil, fmt.Errorf("API error: %s (flaw IDs: %s)",
			mitigationInfo.Errors[0].Type,
			mitigationInfo.Errors[0].FlawIDList)
	}

	// Find the issue for the requested flaw
	for i := range mitigationInfo.Issues {
		if mitigationInfo.Issues[i].FlawID == flawID {
			return &mitigationInfo.Issues[i], nil
		}
	}

	return nil, fmt.Errorf("flaw ID %d not found in response", flawID)
}
