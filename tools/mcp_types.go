package tools

// MCPFindingsResponse represents the complete MCP response for findings
type MCPFindingsResponse struct {
	Application MCPApplication     `json:"application"`
	Sandbox     *MCPSandbox        `json:"sandbox,omitempty"`
	Summary     MCPFindingsSummary `json:"summary"`
	Findings    []MCPFinding       `json:"findings"`
	Pagination  *MCPPagination     `json:"pagination,omitempty"`
}

// MCPApplication represents application information
type MCPApplication struct {
	Name                string `json:"name"`
	ID                  string `json:"id"`
	BusinessCriticality string `json:"business_criticality,omitempty"`
}

// MCPSandbox represents sandbox information
type MCPSandbox struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// MCPFinding represents a security finding in MCP format
type MCPFinding struct {
	// Core identification
	FlawID   string `json:"flaw_id"`
	ScanType string `json:"scan_type"` // STATIC, DAST, SCA

	// Status
	Status           string `json:"status"`            // OPEN, CLOSED, UNKNOWN
	MitigationStatus string `json:"mitigation_status"` // NONE, PROPOSED, APPROVED, REJECTED
	ViolatesPolicy   bool   `json:"violates_policy"`

	// Security details
	Severity      string   `json:"severity"`       // CRITICAL, HIGH, MEDIUM, LOW, INFORMATIONAL
	SeverityScore int32    `json:"severity_score"` // 0-5
	CweId         int32    `json:"cwe_id"`         // CWE ID number (e.g., 78 for CWE-78)
	Description   string   `json:"description"`
	References    []string `json:"references,omitempty"`

	// Location (STATIC only)
	FilePath   string `json:"file_path,omitempty"`
	LineNumber int    `json:"line_number,omitempty"`

	// URL (DAST only)
	URL string `json:"url,omitempty"`

	// SCA-specific
	Component     *MCPComponent     `json:"component,omitempty"`
	Vulnerability *MCPVulnerability `json:"vulnerability,omitempty"`

	// Mitigations
	Mitigations []MCPMitigation `json:"mitigations,omitempty"`

	// Dates
	FirstFound string `json:"first_found,omitempty"`
	LastSeen   string `json:"last_seen,omitempty"`

	// Additional context
	Module       string `json:"module,omitempty"`
	Procedure    string `json:"procedure,omitempty"`
	AttackVector string `json:"attack_vector,omitempty"`
	ContextType  string `json:"context_type,omitempty"`
	Count        int    `json:"count,omitempty"`
}

// MCPComponent represents SCA component information
type MCPComponent struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Library string `json:"library,omitempty"`
}

// MCPVulnerability represents SCA vulnerability information
type MCPVulnerability struct {
	CVEID       string  `json:"cve_id"`
	CVSSScore   float64 `json:"cvss_score"`
	Exploitable bool    `json:"exploitable"`
}

// MCPMitigation represents a mitigation/annotation
type MCPMitigation struct {
	Action    string `json:"action"`
	Comment   string `json:"comment"`
	Submitter string `json:"submitter"`
	Date      string `json:"date"`
}

// MCPFindingsSummary represents summary statistics
type MCPFindingsSummary struct {
	TotalFindings    int            `json:"total_findings"`
	OpenFindings     int            `json:"open_findings"`
	PolicyViolations int            `json:"policy_violations"`
	BySeverity       map[string]int `json:"by_severity"`
	ByStatus         map[string]int `json:"by_status"`
	ByMitigation     map[string]int `json:"by_mitigation_status"`
}

// MCPPagination represents pagination information
type MCPPagination struct {
	CurrentPage   int  `json:"current_page"`
	TotalPages    int  `json:"total_pages"`
	PageSize      int  `json:"page_size"`
	TotalElements int  `json:"total_elements"`
	HasNext       bool `json:"has_next"`
	HasPrevious   bool `json:"has_previous"`
}
