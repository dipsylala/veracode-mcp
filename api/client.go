package api

import (
	"context"
	"fmt"
	"net/http"

	applications "github.com/dipsylala/veracodemcp-go/api/generated/applications"
	dynamic_flaw "github.com/dipsylala/veracodemcp-go/api/generated/dynamic_flaw"
	findings "github.com/dipsylala/veracodemcp-go/api/generated/findings"
	healthcheck "github.com/dipsylala/veracodemcp-go/api/generated/healthcheck"
	static_finding_data_path "github.com/dipsylala/veracodemcp-go/api/generated/static_finding_data_path"
	"github.com/dipsylala/veracodemcp-go/credentials"
	"github.com/dipsylala/veracodemcp-go/hmac"
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
	// Calculate the HMAC Authorization header
	authHeader, err := hmac.CalculateAuthorizationHeader(req.URL, req.Method, t.apiID, t.apiKey)
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
// Base URL defaults to https://api.veracode.com but can be overridden via:
// - api-base-url in veracode.yml
// - VERACODE_API_BASE_URL environment variable
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
