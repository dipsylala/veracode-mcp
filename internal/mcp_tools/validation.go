package mcp_tools

import (
	"fmt"
	"strconv"
	"strings"
)

// extractRequiredString extracts and validates a required non-empty string parameter.
// Returns an error if the parameter is missing or empty.
func extractRequiredString(args map[string]interface{}, field string) (string, error) {
	val, ok := args[field].(string)
	if !ok || val == "" {
		return "", fmt.Errorf("%s is required", field)
	}
	return val, nil
}

// extractOptionalString extracts an optional string parameter.
// Returns (value, true) if present, ("", false) if not present.
// This allows distinguishing between "not provided" and "provided but empty".
func extractOptionalString(args map[string]interface{}, field string) (string, bool) {
	val, ok := args[field].(string)
	return val, ok
}

// extractInt safely extracts an integer from JSON (which provides float64).
// Returns the defaultValue if the parameter is not present.
func extractInt(args map[string]interface{}, field string, defaultValue int) int {
	if val, ok := args[field].(float64); ok {
		return int(val)
	}
	return defaultValue
}

// extractOptionalInt32Ptr extracts an optional int32 pointer parameter (for severity fields).
// Returns (pointer to value, true) if present, (nil, false) if not present.
// Returns error if value is present but not a valid number.
func extractOptionalInt32Ptr(args map[string]interface{}, field string) (*int32, bool, error) {
	val, ok := args[field].(float64)
	if !ok {
		return nil, false, nil
	}

	// Check for overflow when converting to int32
	if val > 2147483647 || val < -2147483648 {
		return nil, true, fmt.Errorf("%s value %v exceeds int32 range", field, val)
	}

	result := int32(val)
	return &result, true, nil
}

// extractOptionalBool extracts an optional boolean parameter.
// Returns (pointer to value, true) if present, (nil, false) if not present.
func extractOptionalBool(args map[string]interface{}, field string) (*bool, bool) {
	val, ok := args[field].(bool)
	if !ok {
		return nil, false
	}
	return &val, true
}

// extractFlawID extracts a flaw_id parameter as an integer.
// This is used for platform API tools where flaw IDs are always unique integers.
// For pipeline scanner tools, use extractFlawIDString which supports the "1000-2" format.
func extractFlawID(args map[string]interface{}) (int, error) {
	val, exists := args["flaw_id"]
	if !exists {
		return 0, fmt.Errorf("flaw_id is required")
	}

	switch v := val.(type) {
	case float64:
		flawID := int(v)
		if flawID <= 0 {
			return 0, fmt.Errorf("flaw_id must be a positive integer")
		}
		return flawID, nil
	case int:
		if v <= 0 {
			return 0, fmt.Errorf("flaw_id must be a positive integer")
		}
		return v, nil
	case string:
		// Support string numbers for consistency
		flawID, err := strconv.Atoi(v)
		if err != nil || flawID <= 0 {
			return 0, fmt.Errorf("flaw_id must be a positive integer")
		}
		return flawID, nil
	default:
		return 0, fmt.Errorf("flaw_id must be an integer")
	}
}

// FlawIDComponents represents a parsed flaw_id with issue_id and occurrence
type FlawIDComponents struct {
	IssueID    int
	Occurrence int // 1-based, defaults to 1 for non-suffixed IDs
}

// extractFlawIDString extracts and parses a flaw_id parameter (pipeline tools only).
// Only accepts strings in format "1000" or "1000-2".
// Returns the parsed components (issue_id and occurrence) or an error if invalid.
func extractFlawIDString(args map[string]interface{}) (*FlawIDComponents, error) {
	val, exists := args["flaw_id"]
	if !exists {
		return nil, fmt.Errorf("flaw_id is required")
	}

	flawIDStr, ok := val.(string)
	if !ok {
		return nil, fmt.Errorf("flaw_id must be a string (e.g., \"1000\" or \"1000-2\")")
	}

	if flawIDStr == "" {
		return nil, fmt.Errorf("flaw_id cannot be empty")
	}

	// Parse string format "1000" or "1000-2"
	parts := strings.Split(flawIDStr, "-")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid flaw_id format: %s (expected '1000' or '1000-2')", flawIDStr)
	}

	// Parse the base issue_id
	issueID, err := strconv.Atoi(parts[0])
	if err != nil || issueID <= 0 {
		return nil, fmt.Errorf("invalid flaw_id: must be a positive integer (got '%s')", parts[0])
	}

	// Parse the occurrence suffix if present
	occurrence := 1
	if len(parts) == 2 {
		occurrence, err = strconv.Atoi(parts[1])
		if err != nil || occurrence <= 0 {
			return nil, fmt.Errorf("invalid flaw_id: occurrence must be a positive integer (got '%s')", parts[1])
		}
	}

	return &FlawIDComponents{IssueID: issueID, Occurrence: occurrence}, nil
}

// validateIntRange validates that an integer is within the specified range (inclusive).
// Returns an error with a descriptive message if the value is out of range.
func validateIntRange(value int, min, max int, fieldName string) error {
	if value < min || value > max {
		return fmt.Errorf("%s must be between %d and %d, got %d", fieldName, min, max, value)
	}
	return nil
}

// validateInt32Range validates that an optional int32 pointer is within the specified range (inclusive).
// Returns nil if the pointer is nil (value not provided).
// Returns an error with a descriptive message if the value is out of range.
func validateInt32Range(value *int32, min, max int32, fieldName string) error {
	if value != nil && (*value < min || *value > max) {
		return fmt.Errorf("%s must be between %d and %d, got %d", fieldName, min, max, *value)
	}
	return nil
}

// validatePaginationParams validates size and page parameters against schema constraints.
// size must be between 1 and 500, page must be between 0 and 500.
func validatePaginationParams(size, page int) error {
	if err := validateIntRange(size, 1, 500, "size"); err != nil {
		return err
	}
	if err := validateIntRange(page, 0, 500, "page"); err != nil {
		return err
	}
	return nil
}

// validateSeverity validates that a severity value is between 0 and 5 (inclusive).
// The fieldName parameter is used in error messages (e.g., "severity" or "severity_gte").
func validateSeverity(severity *int32, fieldName string) error {
	return validateInt32Range(severity, 0, 5, fieldName)
}

// extractCWEIDs converts CWE IDs from various JSON types to a string slice.
// Handles arrays of numbers, strings, or mixed types.
// Returns nil if the parameter is not present or not an array.
func extractCWEIDs(args map[string]interface{}) []string {
	cweArray, ok := args["cwe_ids"].([]interface{})
	if !ok {
		return nil
	}

	cweIDs := make([]string, len(cweArray))
	for i, cwe := range cweArray {
		switch v := cwe.(type) {
		case float64:
			cweIDs[i] = fmt.Sprintf("%.0f", v)
		case int:
			cweIDs[i] = fmt.Sprintf("%d", v)
		case string:
			cweIDs[i] = v
		default:
			cweIDs[i] = fmt.Sprintf("%v", v)
		}
	}
	return cweIDs
}
