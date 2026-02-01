# CWE-749: Exposed Dangerous Method or Function

## LLM Guidance

Exposed dangerous methods (admin functions, debug endpoints, internal APIs, privileged operations) accessible without proper authorization enable attackers to invoke sensitive functionality directly, bypassing normal access controls. The core fix is to make dangerous methods private/internal and gate high-risk operations behind strong authentication and authorization controls.

## Key Principles

- **Make methods private/internal**: Change visibility modifiers, remove from public APIs, move to internal classes
- **Require multi-layered authorization**: Implement authentication tokens, role verification, and permission checks before execution
- **Limit operation scope**: Replace batch operations with single-item operations that require individual authorization
- **Add audit logging**: Track all invocations of sensitive operations for security monitoring

## Remediation Steps

- Identify exposed dangerous methods - Scan for admin functions, delete operations, debug endpoints, privileged APIs
- Determine exposure paths - Check which methods are publicly accessible via API routes, HTTP endpoints, or RPC calls
- Check authorization gaps - Verify if dangerous methods lack authentication, role checks, or permission validation
- Make methods private/internal - Remove from public API surface, add access modifiers, move to internal classes
- Add access controls - Require authentication tokens, verify user roles (e.g., `@require_role('SUPER_ADMIN')`), implement permission checks
- Test unauthorized access - Attempt to invoke dangerous methods without proper credentials—should return 401/403 errors
