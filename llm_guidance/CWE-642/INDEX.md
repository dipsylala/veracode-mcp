# CWE-642: External Control of Critical State Data

## LLM Guidance

External control of critical state data occurs when applications allow untrusted input to directly control state variables determining program behavior, authentication status, authorization levels, or security decisions. This vulnerability arises from storing security-critical state (user roles, permissions, prices, session attributes) in client-side storage (cookies, hidden form fields, URL parameters) without server-side validation or cryptographic protection.

## Key Principles

- Never trust client-side data for security decisions
- Store all security-critical state server-side in protected storage
- Validate any client-provided state against authoritative server records
- Use cryptographically signed tokens if client-side state is unavoidable
- Enforce authorization checks on the server for every security-sensitive operation

## Remediation Steps

- Move all security-critical state (roles, permissions, prices) from client storage to server-side sessions or databases
- Implement server-side session management using secure session stores (Redis, database, encrypted filesystem)
- Validate and authorize every request based on server-side state, not client-provided values
- If client-side state is necessary, use cryptographically signed and encrypted tokens (JWT with strong algorithms)
- Add input validation and sanitization for any client-provided parameters
- Implement comprehensive authorization checks at the application layer for all security-sensitive operations
