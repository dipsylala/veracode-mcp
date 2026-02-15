package api

import "github.com/dipsylala/veracode-mcp/api/rest"

// Re-export types from rest package so consumers can import just "api"
type (
	// HealthStatus represents the health check response
	HealthStatus = rest.HealthStatus

	// FindingsResponse represents a response containing security findings
	FindingsResponse = rest.FindingsResponse

	// FindingsRequest contains parameters for querying findings
	FindingsRequest = rest.FindingsRequest

	// Finding represents a single security finding (SAST/DAST/SCA)
	Finding = rest.Finding

	// License represents license information for SCA findings
	License = rest.License
)

// MitigationInfo represents the root element of the mitigation info response
// This is a wrapper around xml.MitigationInfo that provides a cleaner API
// without XML serialization tags.
type MitigationInfo struct {
	Version string
	BuildID int64
	Issues  []MitigationIssue
	Errors  []MitigationError
}

// MitigationIssue represents a flaw with its mitigation actions
type MitigationIssue struct {
	FlawID            int64
	Category          string
	MitigationActions []MitigationAction
}

// MitigationAction represents a single mitigation action for a flaw
type MitigationAction struct {
	Action   string
	Desc     string
	Reviewer string
	Date     string
	Comment  string
}

// MitigationError represents an error for flaws that could not be processed
type MitigationError struct {
	Type       string
	FlawIDList string
}
