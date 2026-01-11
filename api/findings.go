package api

import (
	"context"
	"fmt"

	findings "github.com/dipsylala/veracodemcp-go/api/generated/findings"
)

// FindingsRequest represents common parameters for findings queries
type FindingsRequest struct {
	AppProfile         string   `json:"app_profile"`
	Sandbox            string   `json:"sandbox,omitempty"`
	Severity           []string `json:"severity,omitempty"`
	Status             []string `json:"status,omitempty"`
	CWEIDs             []string `json:"cwe_ids,omitempty"`
	ViolatesPolicy     *bool    `json:"violates_policy,omitempty"`
	Page               int      `json:"page,omitempty"`
	Size               int      `json:"size,omitempty"`
	IncludeMitigations bool     `json:"include_mitigations,omitempty"`
}

// Finding represents a security finding (SAST or DAST)
type Finding struct {
	ID             string                 `json:"id"`
	Severity       string                 `json:"severity"`
	CWE            string                 `json:"cwe"`
	Status         string                 `json:"status"`
	Description    string                 `json:"description"`
	FilePath       string                 `json:"file_path,omitempty"`   // SAST only
	LineNumber     int                    `json:"line_number,omitempty"` // SAST only
	URL            string                 `json:"url,omitempty"`         // DAST only
	ViolatesPolicy bool                   `json:"violates_policy"`
	Mitigations    []Mitigation           `json:"mitigations,omitempty"`
	AdditionalInfo map[string]interface{} `json:"additional_info,omitempty"`
}

// Mitigation represents a proposed fix or risk acceptance
type Mitigation struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// FindingsResponse represents a paginated list of findings
type FindingsResponse struct {
	Findings   []Finding `json:"findings"`
	TotalCount int       `json:"total_count"`
	Page       int       `json:"page"`
	Size       int       `json:"size"`
}

// GetDynamicFindings retrieves DAST (Dynamic Analysis) findings
func (c *VeracodeClient) GetDynamicFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)

	// Build the API request using the fluent builder pattern
	apiReq := c.findingsClient.ApplicationFindingsInformationAPI.GetFindingsUsingGET(authCtx, req.AppProfile).
		ScanType([]string{"DYNAMIC"})

	// Add CWE filter if provided
	if len(req.CWEIDs) > 0 {
		cweInts := make([]int32, len(req.CWEIDs))
		for i, cwe := range req.CWEIDs {
			var cweInt int32
			fmt.Sscanf(cwe, "%d", &cweInt)
			cweInts[i] = cweInt
		}
		apiReq = apiReq.Cwe(cweInts)
	}

	// Add policy violation filter
	if req.ViolatesPolicy != nil {
		apiReq = apiReq.ViolatesPolicy(*req.ViolatesPolicy)
	}

	// Call the Findings API
	resp, httpResp, err := apiReq.Execute()

	if err != nil {
		if httpResp != nil {
			return nil, fmt.Errorf("API returned status %d: %w", httpResp.StatusCode, err)
		}
		return nil, fmt.Errorf("failed to get dynamic findings: %w", err)
	}

	// Check if response is valid
	if resp == nil || resp.Embedded == nil {
		return &FindingsResponse{
			Findings:   []Finding{},
			TotalCount: 0,
			Page:       req.Page,
			Size:       req.Size,
		}, nil
	}

	// Convert API response to our Finding type
	convertedFindings := convertFindings(resp.Embedded.Findings, "DYNAMIC")

	// Apply post-response filters
	filteredFindings := applyFilters(convertedFindings, req)

	return &FindingsResponse{
		Findings:   filteredFindings,
		TotalCount: len(filteredFindings),
		Page:       req.Page,
		Size:       req.Size,
	}, nil
}

// GetStaticFindings retrieves SAST (Static Analysis) findings
func (c *VeracodeClient) GetStaticFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)

	// Build the API request using the fluent builder pattern
	apiReq := c.findingsClient.ApplicationFindingsInformationAPI.GetFindingsUsingGET(authCtx, req.AppProfile).
		ScanType([]string{"STATIC"})

	// Add CWE filter if provided
	if len(req.CWEIDs) > 0 {
		cweInts := make([]int32, len(req.CWEIDs))
		for i, cwe := range req.CWEIDs {
			var cweInt int32
			fmt.Sscanf(cwe, "%d", &cweInt)
			cweInts[i] = cweInt
		}
		apiReq = apiReq.Cwe(cweInts)
	}

	// Add policy violation filter
	if req.ViolatesPolicy != nil {
		apiReq = apiReq.ViolatesPolicy(*req.ViolatesPolicy)
	}

	// Call the Findings API
	resp, httpResp, err := apiReq.Execute()

	if err != nil {
		if httpResp != nil {
			return nil, fmt.Errorf("API returned status %d: %w", httpResp.StatusCode, err)
		}
		return nil, fmt.Errorf("failed to get static findings: %w", err)
	}

	// Check if response is valid
	if resp == nil || resp.Embedded == nil {
		return &FindingsResponse{
			Findings:   []Finding{},
			TotalCount: 0,
			Page:       req.Page,
			Size:       req.Size,
		}, nil
	}

	// Convert API response to our Finding type
	convertedFindings := convertFindings(resp.Embedded.Findings, "STATIC")

	// Apply post-response filters
	filteredFindings := applyFilters(convertedFindings, req)

	return &FindingsResponse{
		Findings:   filteredFindings,
		TotalCount: len(filteredFindings),
		Page:       req.Page,
		Size:       req.Size,
	}, nil
}

// GetFindingByID retrieves a specific finding by ID
func (c *VeracodeClient) GetFindingByID(ctx context.Context, findingID string, isDynamic bool) (*Finding, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured")
	}

	// TODO: Implement actual API call
	return nil, fmt.Errorf("not implemented")
}

// convertFindings converts the generated API findings to our Finding type
func convertFindings(apiFindings []findings.Finding, scanType string) []Finding {
	result := make([]Finding, 0, len(apiFindings))

	for _, apiFinding := range apiFindings {
		finding := Finding{}

		// Handle pointer fields from generated code
		if apiFinding.IssueId != nil {
			finding.ID = fmt.Sprintf("%d", *apiFinding.IssueId)
		}
		if apiFinding.Description != nil {
			finding.Description = *apiFinding.Description
		}
		if apiFinding.ViolatesPolicy != nil {
			finding.ViolatesPolicy = *apiFinding.ViolatesPolicy
		}

		// Extract severity, CWE, and other details from FindingDetails
		// The FindingDetails is a OneOf type that can be StaticFinding, DynamicFinding, or ScaFinding
		// For now, we'll extract what we can from the basic Finding fields
		// TODO: Properly handle the OneOfFindingFindingDetails polymorphic type

		// Extract status if available
		if apiFinding.FindingStatus != nil && apiFinding.FindingStatus.Status != nil {
			finding.Status = *apiFinding.FindingStatus.Status
		}

		// Extract scan-type-specific fields
		if scanType == "STATIC" {
			// Static findings have file paths and line numbers
			if apiFinding.FindingDetails != nil {
				// Extract from StaticFinding details
				// TODO: Type assert to StaticFinding and extract FilePath, LineNumber
				// finding.FilePath = ...
				// finding.LineNumber = ...
			}
		} else if scanType == "DYNAMIC" {
			// Dynamic findings have URLs
			if apiFinding.FindingDetails != nil {
				// Extract from DynamicFinding details
				// TODO: Type assert to DynamicFinding and extract URL
				// finding.URL = ...
			}
		}

		result = append(result, finding)
	}

	return result
}

// applyFilters applies client-side filters to findings
func applyFilters(findings []Finding, req FindingsRequest) []Finding {
	result := make([]Finding, 0, len(findings))

	for _, finding := range findings {
		// Filter by severity if specified
		if len(req.Severity) > 0 {
			matchesSeverity := false
			for _, sev := range req.Severity {
				if finding.Severity == sev {
					matchesSeverity = true
					break
				}
			}
			if !matchesSeverity {
				continue
			}
		}

		// Filter by status if specified
		if len(req.Status) > 0 {
			matchesStatus := false
			for _, status := range req.Status {
				if finding.Status == status {
					matchesStatus = true
					break
				}
			}
			if !matchesStatus {
				continue
			}
		}

		result = append(result, finding)
	}

	// Apply pagination
	startIdx := req.Page * req.Size
	endIdx := startIdx + req.Size

	if startIdx >= len(result) {
		return []Finding{}
	}

	if endIdx > len(result) {
		endIdx = len(result)
	}

	if req.Size > 0 {
		return result[startIdx:endIdx]
	}

	return result
}
