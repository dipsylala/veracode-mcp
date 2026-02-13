package api

import (
	"context"
	"fmt"

	"github.com/dipsylala/veracodemcp-go/api/rest"
	applications "github.com/dipsylala/veracodemcp-go/api/rest/generated/applications"
	dynamic_flaw "github.com/dipsylala/veracodemcp-go/api/rest/generated/dynamic_flaw"
	static_finding_data_path "github.com/dipsylala/veracodemcp-go/api/rest/generated/static_finding_data_path"
	"github.com/dipsylala/veracodemcp-go/api/xml"
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
	// Convert xml.MitigationInfo to api.MitigationInfo
	return convertXMLMitigationInfo(xmlInfo), nil
}

func (c *unifiedClient) GetMitigationInfoForSingleFlaw(ctx context.Context, buildID, flawID int64) (*MitigationIssue, error) {
	xmlIssue, err := c.xmlClient.GetMitigationInfoForSingleFlaw(ctx, buildID, flawID)
	if err != nil {
		return nil, err
	}
	// Convert xml.Issue to api.MitigationIssue
	return convertXMLIssue(xmlIssue), nil
}

// Helper functions to convert XML types to API types

func convertXMLMitigationInfo(xmlInfo *xml.MitigationInfo) *MitigationInfo {
	if xmlInfo == nil {
		return nil
	}

	issues := make([]MitigationIssue, len(xmlInfo.Issues))
	for i := range xmlInfo.Issues {
		issues[i] = *convertXMLIssue(&xmlInfo.Issues[i])
	}

	errors := make([]MitigationError, len(xmlInfo.Errors))
	for i, err := range xmlInfo.Errors {
		errors[i] = MitigationError{
			Type:       err.Type,
			FlawIDList: err.FlawIDList,
		}
	}

	return &MitigationInfo{
		Version: xmlInfo.Version,
		BuildID: xmlInfo.BuildID,
		Issues:  issues,
		Errors:  errors,
	}
}

func convertXMLIssue(xmlIssue *xml.Issue) *MitigationIssue {
	if xmlIssue == nil {
		return nil
	}

	actions := make([]MitigationAction, len(xmlIssue.MitigationActions))
	for i, action := range xmlIssue.MitigationActions {
		actions[i] = MitigationAction{
			Action:   action.Action,
			Desc:     action.Desc,
			Reviewer: action.Reviewer,
			Date:     action.Date,
			Comment:  action.Comment,
		}
	}

	return &MitigationIssue{
		FlawID:            xmlIssue.FlawID,
		Category:          xmlIssue.Category,
		MitigationActions: actions,
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

// XML API types (re-exported for convenience)
type (
	// MitigationInfo represents the root element of the mitigation info response
	MitigationInfo struct {
		Version string
		BuildID int64
		Issues  []MitigationIssue
		Errors  []MitigationError
	}

	// MitigationIssue represents a flaw with its mitigation actions
	MitigationIssue struct {
		FlawID            int64
		Category          string
		MitigationActions []MitigationAction
	}

	// MitigationAction represents a single mitigation action for a flaw
	MitigationAction struct {
		Action   string
		Desc     string
		Reviewer string
		Date     string
		Comment  string
	}

	// MitigationError represents an error for flaws that could not be processed
	MitigationError struct {
		Type       string
		FlawIDList string
	}
)
