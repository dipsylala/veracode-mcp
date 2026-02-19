package api

import (
	"context"
	"fmt"

	"github.com/dipsylala/veracode-mcp/api/rest"
	applications "github.com/dipsylala/veracode-mcp/api/rest/generated/applications"
	dynamic_flaw "github.com/dipsylala/veracode-mcp/api/rest/generated/dynamic_flaw"
	policy "github.com/dipsylala/veracode-mcp/api/rest/generated/policy"
	static_finding_data_path "github.com/dipsylala/veracode-mcp/api/rest/generated/static_finding_data_path"
	"github.com/dipsylala/veracode-mcp/api/xml"
)

// Compile-time check to ensure unifiedClient implements the Client interface
var _ Client = (*unifiedClient)(nil)

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

	// Policy
	GetPolicy(ctx context.Context, policyName string) (*policy.PagedResourceOfPolicyVersion, error)

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

	// XML API Methods
	// GetMitigationInfo retrieves mitigation information for specific flaws in a build
	GetMitigationInfo(ctx context.Context, buildID int64, flawIDs []int64) (*MitigationInfo, error)

	// GetMitigationInfoForSingleFlaw is a convenience method to get mitigation info for a single flaw
	GetMitigationInfoForSingleFlaw(ctx context.Context, buildID, flawID int64) (*MitigationIssue, error)
}

// unifiedClient wraps both REST and XML API clients to provide transparent access
type unifiedClient struct {
	restClient *rest.Client
	xmlClient  *xml.Client
}

// NewClient creates a new Veracode API client
// The client transparently uses both REST and XML APIs as needed.
//
// Credentials are loaded from:
// 1. ~/.veracode/veracode.yml (preferred)
// 2. Environment variables VERACODE_API_ID and VERACODE_API_KEY (fallback)
func NewClient() (Client, error) {
	restClient, err := rest.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create REST client: %w", err)
	}

	xmlClient, err := xml.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create XML client: %w", err)
	}

	return &unifiedClient{
		restClient: restClient,
		xmlClient:  xmlClient,
	}, nil
}

// NewClientUnconfigured creates a client without checking credentials
// Useful for testing or when credentials will be set later
func NewClientUnconfigured() Client {
	return &unifiedClient{
		restClient: rest.NewClientUnconfigured(),
		xmlClient:  xml.NewClientUnconfigured(),
	}
}

// REST API methods - delegate to restClient

func (c *unifiedClient) IsConfigured() bool {
	return c.restClient.IsConfigured()
}

func (c *unifiedClient) GetAuthContext(ctx context.Context) context.Context {
	return c.restClient.GetAuthContext(ctx)
}

func (c *unifiedClient) CheckHealth(ctx context.Context) (*HealthStatus, error) {
	return c.restClient.CheckHealth(ctx)
}

func (c *unifiedClient) GetApplication(ctx context.Context, applicationGUID string) (*applications.Application, error) {
	return c.restClient.GetApplication(ctx, applicationGUID)
}

func (c *unifiedClient) GetApplicationByName(ctx context.Context, name string) (*applications.Application, error) {
	return c.restClient.GetApplicationByName(ctx, name)
}

func (c *unifiedClient) ListApplications(ctx context.Context, page, size int) (*applications.PagedResourceOfApplication, error) {
	return c.restClient.ListApplications(ctx, page, size)
}

func (c *unifiedClient) GetPolicy(ctx context.Context, policyName string) (*policy.PagedResourceOfPolicyVersion, error) {
	return c.restClient.GetPolicy(ctx, policyName)
}

func (c *unifiedClient) GetStaticFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error) {
	return c.restClient.GetStaticFindings(ctx, req)
}

func (c *unifiedClient) GetDynamicFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error) {
	return c.restClient.GetDynamicFindings(ctx, req)
}

func (c *unifiedClient) GetScaFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error) {
	return c.restClient.GetScaFindings(ctx, req)
}

func (c *unifiedClient) GetFindingByID(ctx context.Context, findingID string, isDynamic bool) (*Finding, error) {
	return c.restClient.GetFindingByID(ctx, findingID, isDynamic)
}

func (c *unifiedClient) StaticFindingDataPathClient() *static_finding_data_path.APIClient {
	return c.restClient.StaticFindingDataPathClient()
}

func (c *unifiedClient) DynamicFlawClient() *dynamic_flaw.APIClient {
	return c.restClient.DynamicFlawClient()
}

func (c *unifiedClient) RawGet(ctx context.Context, endpoint string) (string, error) {
	return c.restClient.RawGet(ctx, endpoint)
}

// XML API methods - delegate to xmlClient

func (c *unifiedClient) GetMitigationInfo(ctx context.Context, buildID int64, flawIDs []int64) (*MitigationInfo, error) {
	xmlInfo, err := c.xmlClient.GetMitigationInfo(ctx, buildID, flawIDs)
	if err != nil {
		return nil, err
	}
	return convertXMLMitigationInfo(xmlInfo), nil
}

func (c *unifiedClient) GetMitigationInfoForSingleFlaw(ctx context.Context, buildID, flawID int64) (*MitigationIssue, error) {
	xmlIssue, err := c.xmlClient.GetMitigationInfoForSingleFlaw(ctx, buildID, flawID)
	if err != nil {
		return nil, err
	}
	return convertXMLIssue(xmlIssue), nil
}
