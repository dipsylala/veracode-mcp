# CWE-352: Cross-Site Request Forgery (CSRF)

## LLM Guidance

CSRF attacks force authenticated users to perform unwanted actions by exploiting the website's trust in the user's browser. Attackers craft malicious requests that leverage the victim's active session to execute state-changing operations. The core fix is verifying request origin and authenticity using server-controlled CSRF tokens.

## Key Principles

- Never process authenticated state-changing requests without verifying origin and authenticity
- Require server-controlled validation mechanisms (tokens) for all POST, PUT, DELETE operations
- Do not rely solely on session cookies for authentication of state-changing actions
- Validate CSRF tokens on the server side before processing any data modifications
- Use framework-provided CSRF protection rather than custom implementations

## Remediation Steps

- Identify all state-changing endpoints (POST, PUT, DELETE) in security findings by file and line number
- Locate forms and AJAX requests that lack CSRF token inclusion
- Review framework configuration to ensure CSRF protection is enabled globally
- Implement Synchronizer Token Pattern with cryptographically random tokens (minimum 32 bytes)
- Include CSRF tokens in all forms and AJAX requests that modify data
- Validate tokens server-side before processing any state-changing request
