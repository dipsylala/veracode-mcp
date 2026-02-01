# CWE-566: Authorization Bypass Through User-Controlled Key

## LLM Guidance

CWE-566 (Insecure Direct Object Reference/IDOR) occurs when applications use user-controlled identifiers directly in database lookups or access decisions without verifying the authenticated user is authorized to access that specific resource. This enables horizontal privilege escalation where attackers access other users' data by manipulating ID parameters. **Core fix:** Enforce object-level access control on every resource lookup—never trust user-controlled keys to bypass authorization.

## Key Principles

- Validate ownership/authorization before every resource access, not just authentication
- Use session-based ownership verification rather than trusting user-supplied IDs
- Return consistent error codes (403 for unauthorized, 404 for not found) to prevent enumeration
- Implement defense-in-depth with indirect references, access control lists, and audit logging
- Apply authorization checks at the data access layer, not just application layer

## Remediation Steps

- Trace data flow - Identify where resource IDs enter the system (URL params, POST data, API paths) and where they're used in database queries
- Locate authorization gaps - Find direct lookups without ownership checks between user input and data access
- Add ownership validation - Before resource access, verify the authenticated user owns/can access the requested resource ID
- Implement consistent responses - Return 403 for unauthorized access, 404 for non-existent resources; avoid leaking existence information
- Use indirect references - Replace predictable IDs with session-mapped tokens or UUIDs where possible
- Test thoroughly - Attempt cross-user access with different accounts to verify authorization enforcement
