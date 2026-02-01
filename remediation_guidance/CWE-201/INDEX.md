# CWE-201: Insertion of Sensitive Information Into Sent Data

## LLM Guidance

Insertion of sensitive information into sent data occurs when applications include confidential data (passwords, tokens, internal paths, stack traces, PII) in HTTP responses, error messages, logs, or API responses transmitted to users or external systems. This enables information disclosure that attackers can exploit to gain unauthorized access or escalate privileges.

## Remediation Principles

- Minimize data exposure: Only include necessary data in responses; exclude internal details, debug info, and sensitive fields
- Sanitize error messages: Return generic errors to clients; log detailed errors server-side only
- Filter serialization: Explicitly control which object properties are serialized to JSON/XML/etc.
- Scrub credentials: Never send passwords, tokens, API keys, or session identifiers in responses or logs
- Strip metadata: Remove stack traces, file paths, database schemas, and system configuration from outbound data

## Remediation Steps

- Identify sensitive data exposure - Review responses for passwords, tokens, session IDs, API keys, PII, internal paths, stack traces, database structure
- Locate transmission points - Check HTTP responses, error messages, API responses, logs, debug output, client-side JavaScript
- Audit serialization - Examine where objects convert to JSON/XML; use whitelists to control exposed fields
- Review error handlers - Replace detailed exception messages with generic errors; log full details server-side only
- Implement response filtering - Add middleware to strip sensitive fields before sending data
- Test data flow - Trace sensitive data from database to user-facing output; verify scrubbing at each layer
