# Duplicate Issue ID Fix

## Problem

Veracode Pipeline Scanner reuses `issue_id` values for different flaws within the same scan. This caused confusion when:
- The UI showed one CWE for flaw 1000
- The remediation-guidance tool returned a different CWE for the same issue_id

### Example from Real Data

Issue ID 1000 appeared **twice** in the same scan:
1. CWE-201 (Information Exposure) at `BlabController.java:57`
2. CWE-80 (XSS) at `profile.jsp:247`

## Root Cause

The `issue_id` field in Veracode pipeline results is **not unique**. The truly unique identifier is `flaw_match.flaw_hash`.

## Solution

### Approach: Human-Readable Suffixes

Instead of exposing the internal flaw_hash, we use a human-readable suffix approach:
- First occurrence: `flaw_id = "1000"`
- Second occurrence: `flaw_id = "1000-2"`
- Third occurrence: `flaw_id = "1000-3"`

This maintains:
- ✅ Uniqueness (each flaw gets a distinct flaw_id)
- ✅ Human readability (easy to reference in LLM conversations)
- ✅ Backward compatibility (no breaking JSON schema changes)
- ✅ Intuitive ordering (suffixes show which occurrence)

### Implementation Details

#### pipeline-results Tool
- Tracks occurrence count for each `issue_id`
- Assigns unique `flaw_id` with suffix for duplicates
- Both occurrences appear in the results list

#### remediation-guidance Tool
- When looking up by `flaw_id` (e.g., 1000), finds the **first** occurrence
- Logs a warning when duplicates exist
- Includes a note in the response listing other occurrences with their unique `flaw_id` values

### User Experience

**Before:**
```
User: "What's the CWE for flaw 1000?"
Tool: Shows CWE-80 in UI, but returns CWE-201 guidance
User: Confused 😕
```

**After:**
```
User: "Show me pipeline results"
Tool: Lists two distinct flaws:
  - flaw_id: 1000 (CWE-201 at BlabController.java:57)
  - flaw_id: 1000-2 (CWE-80 at profile.jsp:247)

User: "Get remediation guidance for flaw 1000-2"
Tool: Returns correct CWE-80 guidance for the XSS issue
User: Happy! 🎉
```

## Testing

See `cwe_consistency_test.go`:
- `TestCWEConsistency_IssueID1000` - Validates both tools return consistent CWE values
- `TestDuplicateIssueIDs` - Specifically tests the duplicate handling logic
- `TestCWEConsistency_WithPagination` - Ensures pagination doesn't affect flaw_id mapping

## Alternative Approaches Considered

1. **Use flaw_hash as flaw_id**: More accurate but not human-readable (e.g., "659586723")
2. **Composite key**: `{issue_id}-{file}-{line}` - Too verbose for LLM interactions
3. **Ignore duplicates**: Would hide flaws from users

## Future Enhancements

If Veracode changes their pipeline scanner to guarantee unique issue_ids, we can remove the suffix logic while maintaining backward compatibility.
