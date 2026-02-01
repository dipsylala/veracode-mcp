# CWE-74: Injection (Generic)

## LLM Guidance

Injection vulnerabilities occur when untrusted data is sent to an interpreter as part of a command or query without proper validation, sanitization, or parameterization. This encompasses SQL injection, command injection, LDAP injection, NoSQL injection, and expression language injection. Attackers exploit these by manipulating input to break out of data context and execute arbitrary commands or code within the interpreter.

## Key Principles

- Never concatenate untrusted input into commands, queries, or executable expressions
- Use parameterized APIs, prepared statements, and safe libraries exclusively
- Treat user data as data only, never as executable code
- Apply input validation using allowlists over denylists
- Implement least privilege for interpreter access

## Remediation Steps

- Identify injection point - Locate source (HTTP parameters, headers, cookies), sink (database, OS shell, LDAP), data flow, and injection type
- Use parameterized queries/safe APIs - Replace string concatenation with prepared statements or parameterized methods
- Validate input - Apply strict allowlist validation; reject malformed or unexpected data
- Sanitize/escape - When parameterization isn't possible, use context-specific escaping functions
- Apply least privilege - Restrict interpreter permissions; use dedicated service accounts with minimal access
- Test thoroughly - Verify fixes with injection payloads; use SAST/DAST tools to confirm remediation
