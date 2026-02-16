package mcp_tools

import (
	"testing"
)

// ============================================================================
// extractRequiredString tests
// ============================================================================

func TestExtractRequiredString_Success(t *testing.T) {
	args := map[string]interface{}{
		"field": "value",
	}

	result, err := extractRequiredString(args, "field")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != "value" {
		t.Errorf("Expected 'value', got '%s'", result)
	}
}

func TestExtractRequiredString_Missing(t *testing.T) {
	args := map[string]interface{}{}

	_, err := extractRequiredString(args, "field")
	if err == nil {
		t.Fatal("Expected error for missing field")
	}
	expectedMsg := "field is required"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestExtractRequiredString_Empty(t *testing.T) {
	args := map[string]interface{}{
		"field": "",
	}

	_, err := extractRequiredString(args, "field")
	if err == nil {
		t.Fatal("Expected error for empty string")
	}
	expectedMsg := "field is required"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestExtractRequiredString_WrongType(t *testing.T) {
	args := map[string]interface{}{
		"field": 123,
	}

	_, err := extractRequiredString(args, "field")
	if err == nil {
		t.Fatal("Expected error for wrong type")
	}
}

// ============================================================================
// extractOptionalString tests
// ============================================================================

func TestExtractOptionalString_Present(t *testing.T) {
	args := map[string]interface{}{
		"field": "value",
	}

	result, present := extractOptionalString(args, "field")
	if !present {
		t.Fatal("Expected present=true")
	}
	if result != "value" {
		t.Errorf("Expected 'value', got '%s'", result)
	}
}

func TestExtractOptionalString_PresentButEmpty(t *testing.T) {
	args := map[string]interface{}{
		"field": "",
	}

	result, present := extractOptionalString(args, "field")
	if !present {
		t.Fatal("Expected present=true for empty string")
	}
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestExtractOptionalString_NotPresent(t *testing.T) {
	args := map[string]interface{}{}

	result, present := extractOptionalString(args, "field")
	if present {
		t.Fatal("Expected present=false")
	}
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestExtractOptionalString_WrongType(t *testing.T) {
	args := map[string]interface{}{
		"field": 123,
	}

	result, present := extractOptionalString(args, "field")
	if present {
		t.Fatal("Expected present=false for wrong type")
	}
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

// ============================================================================
// extractInt tests
// ============================================================================

func TestExtractInt_SuccessFromFloat64(t *testing.T) {
	args := map[string]interface{}{
		"field": float64(42),
	}

	result := extractInt(args, "field", 10)
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
}

func TestExtractInt_UsesDefault(t *testing.T) {
	args := map[string]interface{}{}

	result := extractInt(args, "field", 99)
	if result != 99 {
		t.Errorf("Expected default 99, got %d", result)
	}
}

func TestExtractInt_LargeValue(t *testing.T) {
	args := map[string]interface{}{
		"field": float64(1000000),
	}

	result := extractInt(args, "field", 0)
	if result != 1000000 {
		t.Errorf("Expected 1000000, got %d", result)
	}
}

func TestExtractInt_WrongType(t *testing.T) {
	args := map[string]interface{}{
		"field": "not a number",
	}

	result := extractInt(args, "field", 42)
	if result != 42 {
		t.Errorf("Expected default 42, got %d", result)
	}
}

// ============================================================================
// extractOptionalInt32Ptr tests
// ============================================================================

func TestExtractOptionalInt32Ptr_Success(t *testing.T) {
	args := map[string]interface{}{
		"field": float64(5),
	}

	result, present, err := extractOptionalInt32Ptr(args, "field")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if !present {
		t.Fatal("Expected present=true")
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if *result != 5 {
		t.Errorf("Expected 5, got %d", *result)
	}
}

func TestExtractOptionalInt32Ptr_NotPresent(t *testing.T) {
	args := map[string]interface{}{}

	result, present, err := extractOptionalInt32Ptr(args, "field")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if present {
		t.Fatal("Expected present=false")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

func TestExtractOptionalInt32Ptr_Overflow(t *testing.T) {
	args := map[string]interface{}{
		"field": float64(2147483648), // MaxInt32 + 1
	}

	result, present, err := extractOptionalInt32Ptr(args, "field")
	if err == nil {
		t.Fatal("Expected error for overflow")
	}
	if !present {
		t.Error("Expected present=true even with overflow")
	}
	if result != nil {
		t.Errorf("Expected nil result on overflow, got %v", result)
	}
}

func TestExtractOptionalInt32Ptr_Underflow(t *testing.T) {
	args := map[string]interface{}{
		"field": float64(-2147483649), // MinInt32 - 1
	}

	result, present, err := extractOptionalInt32Ptr(args, "field")
	if err == nil {
		t.Fatal("Expected error for underflow")
	}
	if !present {
		t.Error("Expected present=true even with underflow")
	}
	if result != nil {
		t.Errorf("Expected nil result on underflow, got %v", result)
	}
}

func TestExtractOptionalInt32Ptr_WrongType(t *testing.T) {
	args := map[string]interface{}{
		"field": "not a number",
	}

	result, present, err := extractOptionalInt32Ptr(args, "field")
	if err != nil {
		t.Fatalf("Expected no error for wrong type, got: %v", err)
	}
	if present {
		t.Fatal("Expected present=false")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

// ============================================================================
// extractOptionalBool tests
// ============================================================================

func TestExtractOptionalBool_True(t *testing.T) {
	args := map[string]interface{}{
		"field": true,
	}

	result, present := extractOptionalBool(args, "field")
	if !present {
		t.Fatal("Expected present=true")
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if !*result {
		t.Error("Expected true, got false")
	}
}

func TestExtractOptionalBool_False(t *testing.T) {
	args := map[string]interface{}{
		"field": false,
	}

	result, present := extractOptionalBool(args, "field")
	if !present {
		t.Fatal("Expected present=true")
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if *result {
		t.Error("Expected false, got true")
	}
}

func TestExtractOptionalBool_NotPresent(t *testing.T) {
	args := map[string]interface{}{}

	result, present := extractOptionalBool(args, "field")
	if present {
		t.Fatal("Expected present=false")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

func TestExtractOptionalBool_WrongType(t *testing.T) {
	args := map[string]interface{}{
		"field": "not a bool",
	}

	result, present := extractOptionalBool(args, "field")
	if present {
		t.Fatal("Expected present=false for wrong type")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

// ============================================================================
// extractFlawID tests
// ============================================================================

func TestExtractFlawID_SuccessFromFloat64(t *testing.T) {
	args := map[string]interface{}{
		"flaw_id": float64(12345),
	}

	result, err := extractFlawID(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != 12345 {
		t.Errorf("Expected 12345, got %d", result)
	}
}

func TestExtractFlawID_SuccessFromInt(t *testing.T) {
	args := map[string]interface{}{
		"flaw_id": 67890,
	}

	result, err := extractFlawID(args)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != 67890 {
		t.Errorf("Expected 67890, got %d", result)
	}
}

func TestExtractFlawID_ZeroFromFloat64(t *testing.T) {
	args := map[string]interface{}{
		"flaw_id": float64(0),
	}

	_, err := extractFlawID(args)
	if err == nil {
		t.Fatal("Expected error for zero flaw_id")
	}
	expectedMsg := "flaw_id must be a positive integer"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestExtractFlawID_ZeroFromInt(t *testing.T) {
	args := map[string]interface{}{
		"flaw_id": 0,
	}

	_, err := extractFlawID(args)
	if err == nil {
		t.Fatal("Expected error for zero flaw_id")
	}
}

func TestExtractFlawID_Missing(t *testing.T) {
	args := map[string]interface{}{}

	_, err := extractFlawID(args)
	if err == nil {
		t.Fatal("Expected error for missing flaw_id")
	}
	expectedMsg := "flaw_id is required"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestExtractFlawID_WrongType(t *testing.T) {
	args := map[string]interface{}{
		"flaw_id": "not a number",
	}

	_, err := extractFlawID(args)
	if err == nil {
		t.Fatal("Expected error for wrong type")
	}
}

// ============================================================================
// validateIntRange tests
// ============================================================================

func TestValidateIntRange_Valid(t *testing.T) {
	err := validateIntRange(50, 1, 100, "test")
	if err != nil {
		t.Errorf("Expected no error for valid value, got: %v", err)
	}
}

func TestValidateIntRange_AtMinimum(t *testing.T) {
	err := validateIntRange(1, 1, 100, "test")
	if err != nil {
		t.Errorf("Expected no error at minimum, got: %v", err)
	}
}

func TestValidateIntRange_AtMaximum(t *testing.T) {
	err := validateIntRange(100, 1, 100, "test")
	if err != nil {
		t.Errorf("Expected no error at maximum, got: %v", err)
	}
}

func TestValidateIntRange_BelowMinimum(t *testing.T) {
	err := validateIntRange(0, 1, 100, "test")
	if err == nil {
		t.Fatal("Expected error for value below minimum")
	}
	expectedMsg := "test must be between 1 and 100, got 0"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestValidateIntRange_AboveMaximum(t *testing.T) {
	err := validateIntRange(101, 1, 100, "test")
	if err == nil {
		t.Fatal("Expected error for value above maximum")
	}
	expectedMsg := "test must be between 1 and 100, got 101"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

// ============================================================================
// validateInt32Range tests
// ============================================================================

func TestValidateInt32Range_Valid(t *testing.T) {
	val := int32(3)
	err := validateInt32Range(&val, 0, 5, "severity")
	if err != nil {
		t.Errorf("Expected no error for valid value, got: %v", err)
	}
}

func TestValidateInt32Range_Nil(t *testing.T) {
	err := validateInt32Range(nil, 0, 5, "severity")
	if err != nil {
		t.Errorf("Expected no error for nil value, got: %v", err)
	}
}

func TestValidateInt32Range_AtMinimum(t *testing.T) {
	val := int32(0)
	err := validateInt32Range(&val, 0, 5, "severity")
	if err != nil {
		t.Errorf("Expected no error at minimum, got: %v", err)
	}
}

func TestValidateInt32Range_AtMaximum(t *testing.T) {
	val := int32(5)
	err := validateInt32Range(&val, 0, 5, "severity")
	if err != nil {
		t.Errorf("Expected no error at maximum, got: %v", err)
	}
}

func TestValidateInt32Range_BelowMinimum(t *testing.T) {
	val := int32(-1)
	err := validateInt32Range(&val, 0, 5, "severity")
	if err == nil {
		t.Fatal("Expected error for value below minimum")
	}
	expectedMsg := "severity must be between 0 and 5, got -1"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestValidateInt32Range_AboveMaximum(t *testing.T) {
	val := int32(6)
	err := validateInt32Range(&val, 0, 5, "severity")
	if err == nil {
		t.Fatal("Expected error for value above maximum")
	}
	expectedMsg := "severity must be between 0 and 5, got 6"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

// ============================================================================
// validatePaginationParams tests
// ============================================================================

func TestValidatePaginationParams_Valid(t *testing.T) {
	err := validatePaginationParams(10, 0)
	if err != nil {
		t.Errorf("Expected no error for valid pagination, got: %v", err)
	}
}

func TestValidatePaginationParams_MaxValues(t *testing.T) {
	err := validatePaginationParams(500, 500)
	if err != nil {
		t.Errorf("Expected no error for max values, got: %v", err)
	}
}

func TestValidatePaginationParams_SizeTooSmall(t *testing.T) {
	err := validatePaginationParams(0, 0)
	if err == nil {
		t.Fatal("Expected error for size=0")
	}
	expectedMsg := "size must be between 1 and 500, got 0"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestValidatePaginationParams_SizeTooLarge(t *testing.T) {
	err := validatePaginationParams(501, 0)
	if err == nil {
		t.Fatal("Expected error for size=501")
	}
	expectedMsg := "size must be between 1 and 500, got 501"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestValidatePaginationParams_PageNegative(t *testing.T) {
	err := validatePaginationParams(10, -1)
	if err == nil {
		t.Fatal("Expected error for page=-1")
	}
	expectedMsg := "page must be between 0 and 500, got -1"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestValidatePaginationParams_PageTooLarge(t *testing.T) {
	err := validatePaginationParams(10, 501)
	if err == nil {
		t.Fatal("Expected error for page=501")
	}
	expectedMsg := "page must be between 0 and 500, got 501"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

// ============================================================================
// validateSeverity tests
// ============================================================================

func TestValidateSeverity_Valid(t *testing.T) {
	val := int32(3)
	err := validateSeverity(&val, "severity")
	if err != nil {
		t.Errorf("Expected no error for valid severity, got: %v", err)
	}
}

func TestValidateSeverity_Nil(t *testing.T) {
	err := validateSeverity(nil, "severity")
	if err != nil {
		t.Errorf("Expected no error for nil severity, got: %v", err)
	}
}

func TestValidateSeverity_Zero(t *testing.T) {
	val := int32(0)
	err := validateSeverity(&val, "severity")
	if err != nil {
		t.Errorf("Expected no error for severity=0, got: %v", err)
	}
}

func TestValidateSeverity_Five(t *testing.T) {
	val := int32(5)
	err := validateSeverity(&val, "severity")
	if err != nil {
		t.Errorf("Expected no error for severity=5, got: %v", err)
	}
}

func TestValidateSeverity_BelowZero(t *testing.T) {
	val := int32(-1)
	err := validateSeverity(&val, "severity")
	if err == nil {
		t.Fatal("Expected error for severity=-1")
	}
}

func TestValidateSeverity_AboveFive(t *testing.T) {
	val := int32(6)
	err := validateSeverity(&val, "severity")
	if err == nil {
		t.Fatal("Expected error for severity=6")
	}
}

// ============================================================================
// extractCWEIDs tests
// ============================================================================

func TestExtractCWEIDs_Float64Array(t *testing.T) {
	args := map[string]interface{}{
		"cwe_ids": []interface{}{float64(79), float64(89), float64(78)},
	}

	result := extractCWEIDs(args)
	if len(result) != 3 {
		t.Fatalf("Expected 3 CWE IDs, got %d", len(result))
	}
	if result[0] != "79" || result[1] != "89" || result[2] != "78" {
		t.Errorf("Expected ['79', '89', '78'], got %v", result)
	}
}

func TestExtractCWEIDs_IntArray(t *testing.T) {
	args := map[string]interface{}{
		"cwe_ids": []interface{}{79, 89, 78},
	}

	result := extractCWEIDs(args)
	if len(result) != 3 {
		t.Fatalf("Expected 3 CWE IDs, got %d", len(result))
	}
	if result[0] != "79" || result[1] != "89" || result[2] != "78" {
		t.Errorf("Expected ['79', '89', '78'], got %v", result)
	}
}

func TestExtractCWEIDs_StringArray(t *testing.T) {
	args := map[string]interface{}{
		"cwe_ids": []interface{}{"79", "89", "78"},
	}

	result := extractCWEIDs(args)
	if len(result) != 3 {
		t.Fatalf("Expected 3 CWE IDs, got %d", len(result))
	}
	if result[0] != "79" || result[1] != "89" || result[2] != "78" {
		t.Errorf("Expected ['79', '89', '78'], got %v", result)
	}
}

func TestExtractCWEIDs_MixedTypes(t *testing.T) {
	args := map[string]interface{}{
		"cwe_ids": []interface{}{float64(79), "89", 78},
	}

	result := extractCWEIDs(args)
	if len(result) != 3 {
		t.Fatalf("Expected 3 CWE IDs, got %d", len(result))
	}
	if result[0] != "79" || result[1] != "89" || result[2] != "78" {
		t.Errorf("Expected ['79', '89', '78'], got %v", result)
	}
}

func TestExtractCWEIDs_NotPresent(t *testing.T) {
	args := map[string]interface{}{}

	result := extractCWEIDs(args)
	if result != nil {
		t.Errorf("Expected nil for missing cwe_ids, got %v", result)
	}
}

func TestExtractCWEIDs_WrongType(t *testing.T) {
	args := map[string]interface{}{
		"cwe_ids": "not an array",
	}

	result := extractCWEIDs(args)
	if result != nil {
		t.Errorf("Expected nil for wrong type, got %v", result)
	}
}

func TestExtractCWEIDs_EmptyArray(t *testing.T) {
	args := map[string]interface{}{
		"cwe_ids": []interface{}{},
	}

	result := extractCWEIDs(args)
	if result == nil {
		t.Fatal("Expected non-nil result for empty array")
	}
	if len(result) != 0 {
		t.Errorf("Expected empty array, got %v", result)
	}
}
