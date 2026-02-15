package rest

// HealthStatus represents the result of a health check
type HealthStatus struct {
	Available  bool   `json:"available"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// FindingsRequest represents common parameters for findings queries
type FindingsRequest struct {
	AppProfile     string   `json:"app_profile"`
	Sandbox        string   `json:"sandbox,omitempty"`
	Severity       *int32   `json:"severity,omitempty"`     // Filter for exact severity value (0-5)
	SeverityGte    *int32   `json:"severity_gte,omitempty"` // Filter for severity >= value (0-5)
	Status         []string `json:"status,omitempty"`
	CWEIDs         []string `json:"cwe_ids,omitempty"`
	ViolatesPolicy *bool    `json:"violates_policy,omitempty"`
	Page           int      `json:"page,omitempty"`
	Size           int      `json:"size,omitempty"`
}

// Finding represents a security finding (SAST, DAST, or SCA)
type Finding struct {
	ID                string                 `json:"id"`
	BuildID           int64                  `json:"build_id,omitempty"` // Build/scan ID where this finding was discovered
	Severity          string                 `json:"severity"`           // Numeric severity as string for backward compatibility
	SeverityScore     int32                  `json:"severity_score"`     // Numeric severity score (0-5)
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

// FindingsResponse represents a paginated list of findings
type FindingsResponse struct {
	Findings   []Finding `json:"findings"`
	TotalCount int       `json:"total_count"`
	Page       int       `json:"page"`
	Size       int       `json:"size"`
}
