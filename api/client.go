package api

import (
	"context"
	"fmt"

	"github.com/dipsylala/veracodemcp-go/api/rest"
	applications "github.com/dipsylala/veracodemcp-go/api/rest/generated/applications"
	dynamic_flaw "github.com/dipsylala/veracodemcp-go/api/rest/generated/dynamic_flaw"
	static_finding_data_path "github.com/dipsylala/veracodemcp-go/api/rest/generated/static_finding_data_path"
)

// Client is the main interface for interacting with Veracode APIs
// It abstracts the underlying implementation (REST, XML, etc.)
type Client interface {
	// Authentication & Configuration
	IsConfigured() bool
	GetAuthContext(ctx context.Context) context.Context

	// Health Check
	CheckHealth(ctx context.Context) (*HealthStatus, error)

	// Applications
	GetApplication(ctx context.Context, applicationGUID string) (*applications.Application, error)
	GetApplicationByName(ctx context.Context, name string) (*applications.Application, error)
	ListApplications(ctx context.Context, page, size int) (*applications.PagedResourceOfApplication, error)

	// Findings (SAST/DAST/SCA)
	GetStaticFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error)
	GetDynamicFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error)
	GetScaFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error)
	GetFindingByID(ctx context.Context, findingID string, isDynamic bool) (*Finding, error)

	// Direct access to generated clients (for advanced use cases)
	StaticFindingDataPathClient() *static_finding_data_path.APIClient
	DynamicFlawClient() *dynamic_flaw.APIClient

	// Raw API access
	RawGet(ctx context.Context, endpoint string) (string, error)
}

// ClientOption configures the client
type ClientOption func(*clientConfig)

type clientConfig struct {
	protocol string // "rest", "xml", "auto" (future)
}

// WithProtocol specifies which API protocol to use
func WithProtocol(protocol string) ClientOption {
	return func(c *clientConfig) {
		c.protocol = protocol
	}
}

// NewClient creates a new Veracode API client
// By default, uses REST API. In the future, can auto-select or use XML.
//
// Credentials are loaded from:
// 1. ~/.veracode/veracode.yml (preferred)
// 2. Environment variables VERACODE_API_ID and VERACODE_API_KEY (fallback)
func NewClient(opts ...ClientOption) (Client, error) {
	cfg := &clientConfig{
		protocol: "rest", // Default to REST
	}

	for _, opt := range opts {
		opt(cfg)
	}

	switch cfg.protocol {
	case "rest":
		return rest.NewClient()
	case "xml":
		return nil, fmt.Errorf("XML API not yet implemented")
	default:
		return nil, fmt.Errorf("unknown protocol: %s", cfg.protocol)
	}
}

// NewClientUnconfigured creates a client without checking credentials
// Useful for testing or when credentials will be set later
func NewClientUnconfigured(opts ...ClientOption) Client {
	cfg := &clientConfig{
		protocol: "rest",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	switch cfg.protocol {
	case "rest":
		return rest.NewClientUnconfigured()
	default:
		// Fall back to REST
		return rest.NewClientUnconfigured()
	}
}

// Re-export types from rest package so tools can import just "api"
type (
	// HealthStatus represents the health check response
	HealthStatus = rest.HealthStatus

	// FindingsResponse represents a response containing security findings
	FindingsResponse = rest.FindingsResponse

	// FindingsRequest contains parameters for querying findings
	FindingsRequest = rest.FindingsRequest

	// Finding represents a single security finding (SAST/DAST/SCA)
	Finding = rest.Finding

	// Mitigation represents mitigation information for a finding
	Mitigation = rest.Mitigation

	// License represents license information for SCA findings
	License = rest.License
)
