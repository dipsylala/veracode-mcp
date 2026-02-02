# CWE-472: External Control of Assumed-Immutable Web Parameter

## LLM Guidance

Assumed-immutable parameter vulnerabilities occur when applications trust that client-controlled data (hidden form fields, cookies, disabled inputs, URL parameters) remains unchanged, failing to validate on server-side. This enables price manipulation, privilege escalation, and business logic bypass.

## Key Principles

- Never trust client-side data as immutable-always validate server-side
- Recompute critical values (prices, permissions, quotas) from authoritative sources
- Store authoritative state server-side (sessions, databases), not in client parameters
- Validate all inputs even if marked "disabled" or "hidden" in UI
- Use cryptographic signatures or HMACs for parameters that must round-trip to client

## Remediation Steps

- Locate vulnerabilities - Review data paths where client-controlled parameters affect pricing, authorization, or business logic without server-side validation
- Identify assumed-immutable fields - Find hidden inputs, cookies, URL parameters, disabled form fields used as trusted data
- Implement server-side validation - Verify all client-supplied values against authoritative server-side sources
- Recompute critical values - Calculate prices, quotas, and permissions from database/session data, not client parameters
- Use secure tokens - For parameters requiring client storage, apply cryptographic signatures (HMAC) to detect tampering
- Apply defence-in-depth - Combine server-side validation with session management and access controls
