package api

// NOTE ON POLYMORPHIC UNMARSHALING WORKAROUND:
// The generated OpenAPI client has issues with FindingFindingDetails (oneOf: StaticFinding, DynamicFinding, ManualFinding, ScaFinding).
// The UnmarshalJSON tries each type and accepts the first where the marshaled result != "{}".
// This causes frequent misidentification:
//   - Dynamic findings data is often unmarshaled into ScaFinding (most common case)
//   - Static findings data can sometimes go into DynamicFinding
//
// SOLUTION:
// We use the scan_type field from the parent Finding object to determine the correct extraction path,
// and implement fallback extraction functions that can extract common fields (CWE, Severity) from
// any variant type, even when the data is in the wrong variant.
//
// This is not ideal but necessary until the OpenAPI generator is fixed or we implement custom unmarshaling.

import (
	"context"
	"fmt"
	"log"
	"math"

	findings "github.com/dipsylala/veracodemcp-go/api/generated/findings"
)

// safeInt64ToInt safely converts int64 to int, capping at MaxInt if overflow would occur
func safeInt64ToInt(val int64) int {
	if val > int64(math.MaxInt) {
		return math.MaxInt
	}
	if val < int64(math.MinInt) {
		return math.MinInt
	}
	return int(val)
}

// buildFindingsAPIRequest creates a configured API request with common parameters
func buildFindingsAPIRequest(client *VeracodeClient, ctx context.Context, req FindingsRequest, scanType string) findings.ApiGetFindingsUsingGETRequest {
	apiReq := client.findingsClient.ApplicationFindingsInformationAPI.GetFindingsUsingGET(ctx, req.AppProfile).
		ScanType([]string{scanType}).
		IncludeAnnot(req.IncludeMitigations)

	// Add pagination parameters

	if req.Page > math.MaxInt32 {
		req.Page = math.MaxInt32
	}

	if req.Size > math.MaxInt32 {
		req.Size = math.MaxInt32
	}

	if req.Page >= 0 {
		apiReq = apiReq.Page(int32(req.Page)) // #nosec G115 - validated above
	}
	if req.Size > 0 {
		apiReq = apiReq.Size(int32(req.Size)) // #nosec G115 - validated above
	}

	// Add CWE filter if provided
	if len(req.CWEIDs) > 0 {
		cweInts := make([]int32, len(req.CWEIDs))
		for i, cwe := range req.CWEIDs {
			var cweInt int32
			_, _ = fmt.Sscanf(cwe, "%d", &cweInt) // nolint:errcheck
			cweInts[i] = cweInt
		}
		apiReq = apiReq.Cwe(cweInts)
	}

	// Add severity filters
	if req.Severity != nil {
		apiReq = apiReq.Severity(*req.Severity)
	}
	if req.SeverityGte != nil {
		apiReq = apiReq.SeverityGte(*req.SeverityGte)
	}

	// Add policy violation filter
	if req.ViolatesPolicy != nil {
		apiReq = apiReq.ViolatesPolicy(*req.ViolatesPolicy)
	}

	return apiReq
}

// executeFindingsRequest executes the API request and handles response
func executeFindingsRequest(apiReq findings.ApiGetFindingsUsingGETRequest, req FindingsRequest, scanType string) (*FindingsResponse, error) {
	// Call the Findings API
	resp, httpResp, err := apiReq.Execute()
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
		return nil, fmt.Errorf("failed to get %s findings: %w", scanType, err)
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
	convertedFindings := convertFindings(resp.Embedded.Findings, scanType)

	// Apply post-response filters
	filteredFindings := applyFilters(convertedFindings, req)

	// Get total count from pagination metadata
	totalCount := 0
	if resp.Page != nil && resp.Page.TotalElements != nil {
		totalCount = safeInt64ToInt(*resp.Page.TotalElements)
	}

	return &FindingsResponse{
		Findings:   filteredFindings,
		TotalCount: totalCount,
		Page:       req.Page,
		Size:       req.Size,
	}, nil
}

// FindingsRequest represents common parameters for findings queries
type FindingsRequest struct {
	AppProfile         string   `json:"app_profile"`
	Sandbox            string   `json:"sandbox,omitempty"`
	Severity           *int32   `json:"severity,omitempty"`     // Filter for exact severity value (0-5)
	SeverityGte        *int32   `json:"severity_gte,omitempty"` // Filter for severity >= value (0-5)
	Status             []string `json:"status,omitempty"`
	CWEIDs             []string `json:"cwe_ids,omitempty"`
	ViolatesPolicy     *bool    `json:"violates_policy,omitempty"`
	Page               int      `json:"page,omitempty"`
	Size               int      `json:"size,omitempty"`
	IncludeMitigations bool     `json:"include_mitigations,omitempty"`
}

// Finding represents a security finding (SAST, DAST, or SCA)
type Finding struct {
	ID                string                 `json:"id"`
	Severity          string                 `json:"severity"`       // Numeric severity as string for backward compatibility
	SeverityScore     int32                  `json:"severity_score"` // Numeric severity score (0-5)
	CWE               string                 `json:"cwe"`
	Status            string                 `json:"status"`
	ResolutionStatus  string                 `json:"resolution_status"` // Mitigation status (PROPOSED, APPROVED, REJECTED, etc.)
	Description       string                 `json:"description"`
	FilePath          string                 `json:"file_path,omitempty"`     // SAST only
	LineNumber        int                    `json:"line_number,omitempty"`   // SAST only
	Module            string                 `json:"module,omitempty"`        // SAST only
	Procedure         string                 `json:"procedure,omitempty"`     // SAST only
	AttackVector      string                 `json:"attack_vector,omitempty"` // SAST only
	URL               string                 `json:"url,omitempty"`           // DAST only
	ViolatesPolicy    bool                   `json:"violates_policy"`
	Mitigations       []Mitigation           `json:"mitigations,omitempty"`
	AdditionalInfo    map[string]interface{} `json:"additional_info,omitempty"`
	ComponentFilename string                 `json:"component_filename,omitempty"` // SCA only
	ComponentVersion  string                 `json:"component_version,omitempty"`  // SCA only
	CVE               string                 `json:"cve,omitempty"`                // SCA only
	Licenses          []License              `json:"licenses,omitempty"`           // SCA only
}

// License represents SCA component license information
type License struct {
	LicenseID  string `json:"license_id"`
	RiskRating string `json:"risk_rating,omitempty"`
}

// Mitigation represents a proposed fix or risk acceptance
type Mitigation struct {
	Action    string `json:"action"`
	Comment   string `json:"comment"`
	Submitter string `json:"submitter"`
	Date      string `json:"date"`
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
	apiReq := buildFindingsAPIRequest(c, authCtx, req, "DYNAMIC")
	return executeFindingsRequest(apiReq, req, "DYNAMIC")
}

// GetStaticFindings retrieves SAST (Static Analysis) findings
func (c *VeracodeClient) GetStaticFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)
	apiReq := buildFindingsAPIRequest(c, authCtx, req, "STATIC")
	return executeFindingsRequest(apiReq, req, "STATIC")
}

// GetScaFindings retrieves SCA findings for an application
func (c *VeracodeClient) GetScaFindings(ctx context.Context, req FindingsRequest) (*FindingsResponse, error) {
	if !c.IsConfigured() {
		return nil, fmt.Errorf("API credentials not configured. Set VERACODE_API_ID and VERACODE_API_KEY")
	}

	authCtx := c.GetAuthContext(ctx)
	apiReq := buildFindingsAPIRequest(c, authCtx, req, "SCA")
	return executeFindingsRequest(apiReq, req, "SCA")
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
		finding := convertSingleFinding(apiFinding, scanType)
		result = append(result, finding)
	}

	return result
}

// convertSingleFinding converts a single API finding to our Finding type
func convertSingleFinding(apiFinding findings.Finding, scanType string) Finding {
	finding := Finding{}

	// Extract basic fields
	extractBasicFields(&finding, &apiFinding)

	// The generated client's polymorphic unmarshaling is unreliable.
	// Use scan_type to determine which variant to extract from FindingDetails.
	if apiFinding.FindingDetails != nil {
		switch scanType {
		case "STATIC":
			// For static findings, try StaticFinding first, then DynamicFinding as fallback
			// (due to broken unmarshaling that sometimes puts data in wrong field)
			if apiFinding.FindingDetails.StaticFinding != nil {
				extractStaticFindingDetails(&finding, apiFinding.FindingDetails.StaticFinding)
			} else if apiFinding.FindingDetails.DynamicFinding != nil {
				extractStaticFromDynamic(&finding, apiFinding.FindingDetails.DynamicFinding)
			}
		case "DYNAMIC":
			// For dynamic findings, try DynamicFinding first, then fall back to ScaFinding or StaticFinding
			// NOTE: The generated client frequently misidentifies dynamic findings as ScaFinding
			if apiFinding.FindingDetails.DynamicFinding != nil {
				extractDynamicFindingDetails(&finding, apiFinding.FindingDetails.DynamicFinding)
			} else if apiFinding.FindingDetails.ScaFinding != nil {
				// Common case: UnmarshalJSON puts dynamic data in ScaFinding
				extractDynamicFromSca(&finding, apiFinding.FindingDetails.ScaFinding)
			} else if apiFinding.FindingDetails.StaticFinding != nil {
				extractDynamicFromMismarshaled(&finding, apiFinding.FindingDetails.StaticFinding)
			}
		case "SCA":
			// For SCA findings, try ScaFinding first, then fall back to other types
			if apiFinding.FindingDetails.ScaFinding != nil {
				extractScaFindingDetails(&finding, apiFinding.FindingDetails.ScaFinding)
			} else if apiFinding.FindingDetails.StaticFinding != nil {
				// Fallback: extract common fields from StaticFinding
				extractStaticFindingDetails(&finding, apiFinding.FindingDetails.StaticFinding)
			} else if apiFinding.FindingDetails.DynamicFinding != nil {
				// Fallback: extract common fields from DynamicFinding
				extractDynamicFindingDetails(&finding, apiFinding.FindingDetails.DynamicFinding)
			}
		}
	}

	return finding
}

// extractBasicFields extracts common fields from the API finding
func extractBasicFields(finding *Finding, apiFinding *findings.Finding) {
	if apiFinding.IssueId != nil {
		finding.ID = fmt.Sprintf("%d", *apiFinding.IssueId)
	}
	if apiFinding.Description != nil {
		finding.Description = *apiFinding.Description
	}
	if apiFinding.ViolatesPolicy != nil {
		finding.ViolatesPolicy = *apiFinding.ViolatesPolicy
	}

	// Extract status and resolution status if available
	if apiFinding.FindingStatus != nil {
		if apiFinding.FindingStatus.Status != nil {
			finding.Status = *apiFinding.FindingStatus.Status
		}
		if apiFinding.FindingStatus.ResolutionStatus != nil {
			finding.ResolutionStatus = *apiFinding.FindingStatus.ResolutionStatus
		}
	}

	// Extract annotations (mitigations)
	if len(apiFinding.Annotations) > 0 {
		finding.Mitigations = make([]Mitigation, 0, len(apiFinding.Annotations))
		for _, annotation := range apiFinding.Annotations {
			mitigation := Mitigation{}
			if annotation.Action != nil {
				mitigation.Action = *annotation.Action
			}
			if annotation.Comment != nil {
				mitigation.Comment = *annotation.Comment
			}
			if annotation.UserName != nil {
				mitigation.Submitter = *annotation.UserName
			}
			if annotation.Created != nil {
				mitigation.Date = annotation.Created.Format("2006-01-02T15:04:05Z")
			}
			finding.Mitigations = append(finding.Mitigations, mitigation)
		}
	}
}

// extractStaticFindingDetails extracts fields from static finding details
func extractStaticFindingDetails(finding *Finding, staticDetails *findings.StaticFinding) {
	if staticDetails == nil {
		return
	}

	// Extract CWE information
	if staticDetails.Cwe != nil && staticDetails.Cwe.Id != nil {
		finding.CWE = fmt.Sprintf("CWE-%d", *staticDetails.Cwe.Id)
	}

	// Extract severity
	if staticDetails.Severity != nil {
		finding.SeverityScore = *staticDetails.Severity
		finding.Severity = fmt.Sprintf("%d", *staticDetails.Severity)
	}

	// Extract file path and line number
	if staticDetails.FilePath != nil {
		finding.FilePath = *staticDetails.FilePath
	}
	if staticDetails.FileLineNumber != nil {
		finding.LineNumber = int(*staticDetails.FileLineNumber)
	}

	// Extract module and procedure
	if staticDetails.Module != nil {
		finding.Module = *staticDetails.Module
	}
	if staticDetails.Procedure != nil {
		finding.Procedure = *staticDetails.Procedure
	}
	if staticDetails.AttackVector != nil {
		finding.AttackVector = *staticDetails.AttackVector
	}
}

// extractStaticFindingDetails extracts fields from static finding details stored in DynamicFinding
// when the generated client incorrectly unmarshaled static data as dynamic
func extractStaticFromDynamic(finding *Finding, dynamicDetails *findings.DynamicFinding) {
	if dynamicDetails == nil {
		return
	}

	// Extract CWE - both types have this field
	if dynamicDetails.Cwe != nil && dynamicDetails.Cwe.Id != nil {
		finding.CWE = fmt.Sprintf("CWE-%d", *dynamicDetails.Cwe.Id)
	}

	// Extract severity - both types have this field
	if dynamicDetails.Severity != nil {
		finding.SeverityScore = *dynamicDetails.Severity
		finding.Severity = fmt.Sprintf("%d", *dynamicDetails.Severity)
	}

	// FilePath and FileLineNumber are not in DynamicFinding, so they won't be extracted
	// This is a limitation of the broken unmarshaling
}

// extractDynamicFindingDetails extracts fields from dynamic finding details
func extractDynamicFindingDetails(finding *Finding, dynamicDetails *findings.DynamicFinding) {
	if dynamicDetails == nil {
		return
	}

	// Extract CWE information
	if dynamicDetails.Cwe != nil && dynamicDetails.Cwe.Id != nil {
		finding.CWE = fmt.Sprintf("CWE-%d", *dynamicDetails.Cwe.Id)
	}

	// Extract severity
	if dynamicDetails.Severity != nil {
		finding.SeverityScore = *dynamicDetails.Severity
		finding.Severity = fmt.Sprintf("%d", *dynamicDetails.Severity)
	}

	// Extract URL
	if dynamicDetails.URL != nil {
		finding.URL = *dynamicDetails.URL
	}

	// Extract attack vector
	if dynamicDetails.AttackVector != nil {
		finding.AttackVector = *dynamicDetails.AttackVector
	}
}

// extractDynamicFromMismarshaled extracts dynamic finding data from a StaticFinding
// when the generated client incorrectly unmarshaled dynamic data as static
func extractDynamicFromMismarshaled(finding *Finding, staticDetails *findings.StaticFinding) {
	if staticDetails == nil {
		return
	}

	// Extract CWE - both types have this field
	if staticDetails.Cwe != nil && staticDetails.Cwe.Id != nil {
		finding.CWE = fmt.Sprintf("CWE-%d", *staticDetails.Cwe.Id)
	}

	// Extract severity - both types have this field
	if staticDetails.Severity != nil {
		finding.SeverityScore = *staticDetails.Severity
		finding.Severity = fmt.Sprintf("%d", *staticDetails.Severity)
	}

	// URL is not in StaticFinding, so it won't be extracted
	// This is a limitation of the broken unmarshaling
}

// extractDynamicFromSca extracts dynamic finding fields from ScaFinding
// when the generated client incorrectly unmarshaled dynamic data as SCA
func extractDynamicFromSca(finding *Finding, scaDetails *findings.ScaFinding) {
	if scaDetails == nil {
		return
	}

	// Extract CWE - ScaFinding has CWE with ID as float32
	if scaDetails.Cwe != nil && scaDetails.Cwe.Id != nil {
		finding.CWE = fmt.Sprintf("CWE-%d", int32(*scaDetails.Cwe.Id))
	}

	// Extract severity - ScaFinding has severity as float32
	if scaDetails.Severity != nil {
		finding.SeverityScore = int32(*scaDetails.Severity)
		finding.Severity = fmt.Sprintf("%d", int32(*scaDetails.Severity))
	}

	// URL is not in ScaFinding, so it won't be extracted
	// This is a limitation of the broken unmarshaling
}

// extractScaFindingDetails extracts fields from SCA finding details
func extractScaFindingDetails(finding *Finding, scaDetails *findings.ScaFinding) {
	if scaDetails == nil {
		return
	}

	// Extract CWE - ScaFinding has CWE with ID as float32
	if scaDetails.Cwe != nil && scaDetails.Cwe.Id != nil {
		finding.CWE = fmt.Sprintf("CWE-%d", int32(*scaDetails.Cwe.Id))
	}

	// Extract severity - ScaFinding has severity as float32
	if scaDetails.Severity != nil {
		finding.SeverityScore = int32(*scaDetails.Severity)
		finding.Severity = fmt.Sprintf("%d", int32(*scaDetails.Severity))
	}

	// Extract SCA-specific component information
	if scaDetails.ComponentFilename != nil {
		finding.ComponentFilename = *scaDetails.ComponentFilename
	}

	if scaDetails.Version != nil {
		finding.ComponentVersion = *scaDetails.Version
	}

	// Extract CVE information if available
	if scaDetails.Cve != nil && scaDetails.Cve.Name != nil {
		finding.CVE = *scaDetails.Cve.Name
	}

	// Extract license information
	if len(scaDetails.Licenses) > 0 {
		finding.Licenses = make([]License, 0, len(scaDetails.Licenses))
		for _, lic := range scaDetails.Licenses {
			license := License{}
			if lic.LicenseId != nil {
				license.LicenseID = *lic.LicenseId
			}
			if lic.RiskRating != nil {
				license.RiskRating = *lic.RiskRating
			}
			finding.Licenses = append(finding.Licenses, license)
		}
	}
}

// applyFilters applies client-side filters to findings
// applyFilters applies client-side filters that are not supported by the API
// Currently only Status filtering is done client-side, as the API doesn't support it.
// Pagination is now handled server-side via Page and Size parameters.
func applyFilters(findings []Finding, req FindingsRequest) []Finding {
	// Filter by status if specified (not supported by API)
	if len(req.Status) == 0 {
		return findings
	}

	result := make([]Finding, 0, len(findings))
	for _, finding := range findings {
		for _, status := range req.Status {
			if finding.Status == status {
				result = append(result, finding)
				break
			}
		}
	}

	return result
}
