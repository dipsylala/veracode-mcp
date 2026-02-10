package mcp_tools

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/dipsylala/veracodemcp-go/api"
	findings "github.com/dipsylala/veracodemcp-go/api/generated/findings"
)

// SeverityLevel represents the severity of a finding
type SeverityLevel string

const (
	SeverityCritical      SeverityLevel = "CRITICAL"
	SeverityVeryHigh      SeverityLevel = "VERY HIGH"
	SeverityHigh          SeverityLevel = "HIGH"
	SeverityMedium        SeverityLevel = "MEDIUM"
	SeverityLow           SeverityLevel = "LOW"
	SeverityVeryLow       SeverityLevel = "VERY LOW"
	SeverityInformational SeverityLevel = "INFO"
)

// FindingStatus represents the status of a finding
type FindingStatus string

const (
	StatusOpen    FindingStatus = "OPEN"
	StatusClosed  FindingStatus = "CLOSED"
	StatusUnknown FindingStatus = "UNKNOWN"
)

// MitigationStatus represents the mitigation status
type MitigationStatus string

const (
	MitigationNone     MitigationStatus = "NONE"
	MitigationProposed MitigationStatus = "PROPOSED"
	MitigationApproved MitigationStatus = "APPROVED"
	MitigationRejected MitigationStatus = "REJECTED"
)

// TransformSeverity converts numeric severity to text representation
func TransformSeverity(severity *int32) SeverityLevel {
	if severity == nil {
		return SeverityInformational
	}

	switch *severity {
	case 5:
		return SeverityVeryHigh
	case 4:
		return SeverityHigh
	case 3:
		return SeverityMedium
	case 2:
		return SeverityLow
	case 1:
		return SeverityVeryLow
	case 0:
		return SeverityInformational
	default:
		return SeverityInformational
	}
}

// TransformStatus converts raw finding status to clean MCP status
func TransformStatus(findingStatus *findings.FindingStatus) FindingStatus {
	if findingStatus == nil || findingStatus.Status == nil {
		return StatusUnknown
	}

	status := strings.ToUpper(*findingStatus.Status)
	switch status {
	case "OPEN", "NEW":
		return StatusOpen
	case "CLOSED":
		return StatusClosed
	default:
		return StatusUnknown
	}
}

// TransformMitigationStatus converts resolution status to mitigation status
func TransformMitigationStatus(findingStatus *findings.FindingStatus) MitigationStatus {
	if findingStatus == nil || findingStatus.ResolutionStatus == nil {
		return MitigationNone
	}

	status := strings.ToUpper(*findingStatus.ResolutionStatus)
	switch status {
	case "PROPOSED":
		return MitigationProposed
	case "APPROVED":
		return MitigationApproved
	case "REJECTED":
		return MitigationRejected
	default:
		return MitigationNone
	}
}

// DeterminesPolicyViolation determines if a finding violates policy
// Business Rule: CLOSED findings no longer violate policy, regardless of API response.
func DeterminesPolicyViolation(status FindingStatus, resolutionStatus string, originalViolatesPolicy *bool) bool {
	// CLOSED findings never violate policy (override API response)
	if status == StatusClosed {
		return false
	}

	// For non-closed findings, use original API value if available
	if originalViolatesPolicy != nil {
		return *originalViolatesPolicy
	}

	// Default: OPEN findings violate policy
	return status == StatusOpen
}

// TransformMitigationAction converts action code to user-friendly description
func TransformMitigationAction(actionCode string) string {
	switch strings.ToUpper(actionCode) {
	case "FP":
		return "Potential False Positive"
	case "APPDESIGN":
		return "Mitigate by Design"
	case "OSENV":
		return "Mitigate by OS Environment"
	case "NETENV":
		return "Mitigate by Network Environment"
	case "ACCEPTRISK":
		return "Accept Risk"
	case "DEFER":
		return "Deferred"
	case "CONFORMS":
		return "Conforms"
	case "DEVIATES":
		return "Deviates"
	case "CUSTOMCLEANSERAPPROVED":
		return "Approved due to Custom Cleanser"
	case "APPROVED":
		return "Approved"
	case "COMMENT":
		return "Comment"
	default:
		if actionCode == "" {
			return "Unknown"
		}
		return actionCode
	}
}

// ProcessBase64Description decodes base64 encoded descriptions (for DAST findings)
func ProcessBase64Description(description string) string {
	// Try to detect if the description is base64 encoded
	if decoded, err := base64.StdEncoding.DecodeString(description); err == nil {
		// Check if decoded string is valid UTF-8
		if decodedStr := string(decoded); len(decodedStr) > 0 {
			return decodedStr
		}
	}
	return description
}

// CleanDescription removes HTML tags and cleans up the description
func CleanDescription(description string) string {
	// First, remove the entire "References:" section with all its links
	cleaned := referencesSectionPattern.ReplaceAllString(description, "")

	// Then remove all remaining HTML tags
	cleaned = htmlTagPattern.ReplaceAllString(cleaned, "")

	// Finally, clean up whitespace
	cleaned = strings.TrimSpace(cleaned)
	cleaned = whitespacePattern.ReplaceAllString(cleaned, " ")

	return cleaned
}

// Pre-compiled regex patterns for performance
var (
	// Matches HTML anchor tags: <a href=\"url\">name</a> or <a href="url">name</a>
	anchorTagPattern = regexp.MustCompile(`<a\s+href=\\?"([^"\\]+)\\?"[^>]*>([^<]+)</a>`)
	// Matches plain URLs as fallback
	plainURLPattern = regexp.MustCompile(`https?://[^\s<>"\\]+`)
	// Matches "References:" section with anchor tags (to be removed from descriptions)
	// Example: "References: <a href=\"...\">CWE</a> <a href=\"...\">OWASP</a>"
	// Or: "<span>References: <a ...>...</a> <a ...>...</a></span>"
	referencesSectionPattern = regexp.MustCompile(`(?:<[^>]*>)?\s*(?i)references?:\s*(?:<a\s+[^>]*>.*?</a>\s*)+\s*(?:</[^>]*>)?`)
	// Matches all HTML tags
	htmlTagPattern = regexp.MustCompile(`<[^>]*>`)
	// Matches multiple whitespace characters
	whitespacePattern = regexp.MustCompile(`\s+`)
)

// ExtractReferences extracts reference links from descriptions
// Handles both HTML anchor tags and plain URLs. Does not modify the input.
func ExtractReferences(description string) []Reference {
	references := []Reference{}

	// First, try to extract HTML anchor tags: <a href=\"url\">name</a>
	// Handle both escaped quotes (\") and regular quotes (")
	matches := anchorTagPattern.FindAllStringSubmatch(description, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			if len(match) >= 3 {
				references = append(references, Reference{
					URL:  match[1],
					Name: match[2],
				})
			}
		}
	} else {
		// Fallback: Extract plain URLs without names
		urlMatches := plainURLPattern.FindAllString(description, -1)
		for _, url := range urlMatches {
			references = append(references, Reference{
				URL:  url,
				Name: url,
			})
		}
	}

	return references
}

// TransformDescription processes description based on scan type
func TransformDescription(description string, scanType string) (cleanedDesc string, references []Reference) {
	if description == "" {
		return "No description available", nil
	}

	switch strings.ToUpper(scanType) {
	case "DAST", "DYNAMIC":
		// For DAST: decode base64 first, then extract references and clean
		decoded := ProcessBase64Description(description)
		references = ExtractReferences(decoded)
		cleanedDesc = CleanDescription(decoded)

	case "STATIC", "SAST":
		// For STATIC: extract references and clean
		references = ExtractReferences(description)
		cleanedDesc = CleanDescription(description)

	default:
		// For SCA and others: just clean
		cleanedDesc = CleanDescription(description)
	}

	return cleanedDesc, references
}

// TransformDate formats date to ISO string
func TransformDate(dateStr *string) string {
	if dateStr == nil || *dateStr == "" {
		return ""
	}

	// Try to parse and format as ISO
	if t, err := time.Parse(time.RFC3339, *dateStr); err == nil {
		return t.Format(time.RFC3339)
	}

	// Return original if parsing fails
	return *dateStr
}

// TransformWeaknessType generates weakness type string (e.g., "CWE-79")
func TransformWeaknessType(cweID *int32) string {
	if cweID == nil {
		return "Unknown"
	}
	return fmt.Sprintf("CWE-%d", *cweID)
}

// FindingsSummary represents summary statistics
type FindingsSummary struct {
	TotalFindings    int            `json:"total_findings"`
	OpenFindings     int            `json:"open_findings"`
	PolicyViolations int            `json:"policy_violations"`
	BySeverity       map[string]int `json:"by_severity"`
	ByScanType       map[string]int `json:"by_scan_type"`
	ByStatus         map[string]int `json:"by_status"`
	ByMitigation     map[string]int `json:"by_mitigation_status"`
}

// GenerateSummary generates summary statistics from findings array
func GenerateSummary(findings []api.Finding) FindingsSummary {
	summary := FindingsSummary{
		TotalFindings: len(findings),
		BySeverity: map[string]int{
			"very high": 0,
			"high":      0,
			"medium":    0,
			"low":       0,
			"very low":  0,
			"info":      0,
		},
		ByScanType: map[string]int{
			"static": 0,
			"sca":    0,
			"dast":   0,
		},
		ByStatus: map[string]int{
			"open":   0,
			"closed": 0,
		},
		ByMitigation: map[string]int{
			"none":     0,
			"proposed": 0,
			"approved": 0,
			"rejected": 0,
		},
	}

	for _, finding := range findings {
		// Count open findings
		if finding.Status == string(StatusOpen) {
			summary.OpenFindings++
		}

		// Count policy violations
		if finding.ViolatesPolicy {
			summary.PolicyViolations++
		}

		// Count by severity
		severity := strings.ToLower(finding.Severity)
		if count, ok := summary.BySeverity[severity]; ok {
			summary.BySeverity[severity] = count + 1
		}

		// Count by scan type (if we add it to Finding struct)
		// This would require updating the Finding struct

		// Count by status
		status := strings.ToLower(finding.Status)
		if count, ok := summary.ByStatus[status]; ok {
			summary.ByStatus[status] = count + 1
		}
	}

	return summary
}
