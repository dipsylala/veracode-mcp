# CWE-346: Origin Validation Error

## LLM Guidance

Origin validation errors occur when applications fail to properly verify the source of requests, accepting cross-origin requests without validation, trusting Referer headers, or misconfiguring CORS. This enables CSRF attacks, cross-site data theft, and unauthorized cross-domain access. The core fix is to validate the origin of security-relevant data and bind it to a trusted identity or channel before acting on it.

## Key Remediation Principles

- Validate request origin before processing state-changing operations
- Implement anti-CSRF tokens for all sensitive forms and endpoints
- Configure CORS policies restrictively with explicit origin allowlists
- Never trust client-supplied headers (Referer, Origin) alone for security decisions
- Bind sensitive operations to authenticated sessions with origin validation

## Remediation Steps

- Identify missing origin checks - Review scan results for endpoints accepting cross-origin requests without validation (CORS misconfiguration, missing CSRF tokens, no Origin/Referer validation)
- Find vulnerable endpoints - Locate state-changing operations (POST, PUT, DELETE) that don't verify request origin
- Fix CORS configuration - Replace `Access-Control-Allow-Origin - *` with specific trusted origins; remove credentials support from wildcard CORS
- Implement CSRF protection - Add anti-CSRF tokens to forms and validate them server-side; use SameSite cookie attributes
- Validate Origin header - Check Origin/Referer headers against allowlists for sensitive endpoints
- Test cross-origin scenarios - Verify protection by attempting requests from untrusted origins
