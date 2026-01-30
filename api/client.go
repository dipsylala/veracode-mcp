package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	applications "github.com/dipsylala/veracodemcp-go/api/generated/applications"
	dynamic_flaw "github.com/dipsylala/veracodemcp-go/api/generated/dynamic_flaw"
	findings "github.com/dipsylala/veracodemcp-go/api/generated/findings"
	healthcheck "github.com/dipsylala/veracodemcp-go/api/generated/healthcheck"
	static_finding_data_path "github.com/dipsylala/veracodemcp-go/api/generated/static_finding_data_path"
	"github.com/dipsylala/veracodemcp-go/credentials"
	veracodehmac "github.com/dipsylala/veracodemcp-go/hmac"
)

// VeracodeClient wraps the generated API clients with authentication and configuration
type VeracodeClient struct {
	healthcheckClient           *healthcheck.APIClient
	findingsClient              *findings.APIClient
	dynamicFlawClient           *dynamic_flaw.APIClient
	staticFindingDataPathClient *static_finding_data_path.APIClient
	applicationsClient          *applications.APIClient
	apiID                       string
	apiKey                      string
	baseURL                     string
}

// hmacTransport is an HTTP RoundTripper that adds Veracode HMAC authentication to every request
type hmacTransport struct {
	apiID     string
	apiKey    string
	transport http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface, adding HMAC authentication to each request
func (t *hmacTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Veracode's API expects query parameters to use %20 for spaces, not +
	// Go's url.Values.Encode() uses + for spaces (per application/x-www-form-urlencoded spec)
	// We need to normalize the query string to use %20 for proper HMAC calculation and HTTP transmission
	if req.URL.RawQuery != "" {
		normalizedQuery := strings.ReplaceAll(req.URL.RawQuery, "+", "%20")
		req.URL.RawQuery = normalizedQuery
	}

	// Calculate the HMAC Authorization header
	authHeader, err := veracodehmac.CalculateAuthorizationHeader(req.URL, req.Method, t.apiID, t.apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate HMAC authorization: %w", err)
	}

	// Add the Authorization header to the request
	req.Header.Set("Authorization", authHeader)

	// Use the wrapped transport (or default) to execute the request
	transport := t.transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	return transport.RoundTrip(req)
}

// newHMACHTTPClient creates an HTTP client that automatically adds HMAC authentication to all requests
func newHMACHTTPClient(apiID, apiKey string) *http.Client {
	return &http.Client{
		Transport: &hmacTransport{
			apiID:     apiID,
			apiKey:    apiKey,
			transport: http.DefaultTransport,
		},
	}
}

// NewVeracodeClient creates a new Veracode API client
// Credentials are loaded from:
// 1. ~/.veracode/veracode.yml (preferred)
// 2. Environment variables VERACODE_API_ID and VERACODE_API_KEY (fallback)
// Base URL is auto-detected from API key ID prefix (vera01ei-* → EU, otherwise → US)
// Can be overridden via override-api-base-url in veracode.yml or VERACODE_OVERRIDE_API_BASE_URL environment variable
func NewVeracodeClient() (*VeracodeClient, error) {
	apiID, apiKey, baseURL, err := credentials.GetCredentials()
	if err != nil {
		return nil, err
	}

	// Create HTTP client with HMAC authentication
	httpClient := newHMACHTTPClient(apiID, apiKey)

	// Configure all API clients to use the HMAC-authenticated HTTP client and base URL
	healthcheckCfg := healthcheck.NewConfiguration()
	healthcheckCfg.HTTPClient = httpClient
	healthcheckCfg.Servers[0].URL = baseURL

	findingsCfg := findings.NewConfiguration()
	findingsCfg.HTTPClient = httpClient
	findingsCfg.Servers[0].URL = baseURL

	dynamicFlawCfg := dynamic_flaw.NewConfiguration()
	dynamicFlawCfg.HTTPClient = httpClient
	dynamicFlawCfg.Servers[0].URL = baseURL

	staticFindingDataPathCfg := static_finding_data_path.NewConfiguration()
	staticFindingDataPathCfg.HTTPClient = httpClient
	staticFindingDataPathCfg.Servers[0].URL = baseURL

	applicationsCfg := applications.NewConfiguration()
	applicationsCfg.HTTPClient = httpClient
	applicationsCfg.Servers[0].URL = baseURL

	return &VeracodeClient{
		healthcheckClient:           healthcheck.NewAPIClient(healthcheckCfg),
		findingsClient:              findings.NewAPIClient(findingsCfg),
		dynamicFlawClient:           dynamic_flaw.NewAPIClient(dynamicFlawCfg),
		staticFindingDataPathClient: static_finding_data_path.NewAPIClient(staticFindingDataPathCfg),
		applicationsClient:          applications.NewAPIClient(applicationsCfg),
		apiID:                       apiID,
		apiKey:                      apiKey,
		baseURL:                     baseURL,
	}, nil
}

// NewVeracodeClientUnconfigured creates a client without checking credentials
// Useful for testing or when credentials will be set later
func NewVeracodeClientUnconfigured() *VeracodeClient {
	apiID, apiKey, baseURL, err := credentials.GetCredentials()

	// Create HTTP client with HMAC authentication if credentials are available
	var httpClient *http.Client
	if err == nil && apiID != "" && apiKey != "" {
		httpClient = newHMACHTTPClient(apiID, apiKey)
	} else {
		httpClient = http.DefaultClient
		apiID = ""
		apiKey = ""
		baseURL = credentials.DefaultBaseURL
	}

	// Configure all API clients
	healthcheckCfg := healthcheck.NewConfiguration()
	healthcheckCfg.HTTPClient = httpClient
	healthcheckCfg.Servers[0].URL = baseURL

	findingsCfg := findings.NewConfiguration()
	findingsCfg.HTTPClient = httpClient
	findingsCfg.Servers[0].URL = baseURL

	dynamicFlawCfg := dynamic_flaw.NewConfiguration()
	dynamicFlawCfg.HTTPClient = httpClient
	dynamicFlawCfg.Servers[0].URL = baseURL

	staticFindingDataPathCfg := static_finding_data_path.NewConfiguration()
	staticFindingDataPathCfg.HTTPClient = httpClient
	staticFindingDataPathCfg.Servers[0].URL = baseURL

	applicationsCfg := applications.NewConfiguration()
	applicationsCfg.HTTPClient = httpClient
	applicationsCfg.Servers[0].URL = baseURL

	return &VeracodeClient{
		healthcheckClient:           healthcheck.NewAPIClient(healthcheckCfg),
		findingsClient:              findings.NewAPIClient(findingsCfg),
		dynamicFlawClient:           dynamic_flaw.NewAPIClient(dynamicFlawCfg),
		staticFindingDataPathClient: static_finding_data_path.NewAPIClient(staticFindingDataPathCfg),
		applicationsClient:          applications.NewAPIClient(applicationsCfg),
		apiID:                       apiID,
		apiKey:                      apiKey,
		baseURL:                     baseURL,
	}
}

// IsConfigured returns true if API credentials are set
func (c *VeracodeClient) IsConfigured() bool {
	return c.apiID != "" && c.apiKey != ""
}

// GetAuthContext returns a context with Veracode API authentication
// This should be passed to all API calls
// Note: HMAC authentication is handled automatically by the HTTP client transport
func (c *VeracodeClient) GetAuthContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	// HMAC authentication is now handled automatically by the hmacTransport
	// in the HTTP client, so no need to add anything to the context
	return ctx
}

// StaticFindingDataPathClient returns the static finding data path API client
func (c *VeracodeClient) StaticFindingDataPathClient() *static_finding_data_path.APIClient {
	return c.staticFindingDataPathClient
}

// DynamicFlawClient returns the dynamic flaw API client
func (c *VeracodeClient) DynamicFlawClient() *dynamic_flaw.APIClient {
	return c.dynamicFlawClient
}

// RawGet performs a raw GET request to the specified endpoint
func (c *VeracodeClient) RawGet(ctx context.Context, endpoint string) (string, error) {
	if !c.IsConfigured() {
		return "", fmt.Errorf("API credentials not configured")
	}

	// Create HTTP client with HMAC authentication
	httpClient := newHMACHTTPClient(c.apiID, c.apiKey)

	// Build the full URL
	fullURL := c.baseURL + endpoint

	// Create the request
	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Execute the request
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			body = append(body, buf[:n]...)
		}
		if err != nil {
			break
		}
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		bodyStr := string(body)
		if len(bodyStr) == 0 {
			bodyStr = "(empty response)"
		}
		return bodyStr, fmt.Errorf("API returned status %d: %s", resp.StatusCode, bodyStr)
	}

	return string(body), nil
}
