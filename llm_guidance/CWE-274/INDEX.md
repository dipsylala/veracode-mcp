# CWE-274 - Improper Handling of Insufficient Privileges

## LLM Guidance

Improper handling of insufficient privileges occurs when applications don't gracefully handle permission denied errors, continuing operation with partial functionality, exposing error details, or failing insecurely. Applications must check permissions before operations and handle authorization failures properly by failing closed when privileges are insufficient.

## Key Principles

- Fail closed - Do not proceed with operations when privileges are insufficient
- Check before access - Verify permissions before attempting file/resource operations
- Handle errors securely - Catch permission exceptions without exposing sensitive paths or details
- Avoid insecure fallbacks - Don't silently continue with reduced functionality on authorization failure
- Enforce consistently - Apply permission checks uniformly across all resource access points

## Remediation Steps

- Locate missing permission checks - Find file/resource access operations without prior authorization verification
- Implement pre-operation checks - Add permission validation before attempting privileged operations
- Catch permission exceptions - Wrap operations in try-catch blocks that handle PermissionError/access denied errors
- Sanitize error messages - Remove file paths, system details, and permission information from user-facing errors
- Fail securely - Terminate or deny access rather than continuing with partial functionality
- Review fallback logic - Ensure degraded modes don't bypass security requirements
