# CWE-285: Improper Authorization

## LLM Guidance

Improper Authorization occurs when applications fail to enforce or incorrectly implement authorization checks, allowing users to access resources or perform actions beyond their intended permissions. The core fix is to explicitly validate that the authenticated user has permission to access the specific resource or perform the requested operation before allowing the action.

## Key Principles

- Never infer authorization from authentication, role, or prior checks alone
- Every security-sensitive action must be explicitly authorized against the specific resource and operation
- Authorization checks must occur server-side and cannot be bypassed by client manipulation
- Use centralized authorization logic to ensure consistent enforcement across the application
- Default to deny: require explicit permission grants rather than assuming access

## Remediation Steps

- Review flaw details to identify the specific file, line number, and vulnerable code pattern
- Identify the protected resource or operation lacking authorization checks
- Trace the data flow to determine if ANY authorization check exists before the operation
- Implement explicit permission checks for every request that accesses protected resources
- Validate that the current user owns or has explicit access to the specific resource being accessed
- Use established authorization frameworks or patterns rather than custom logic where possible
