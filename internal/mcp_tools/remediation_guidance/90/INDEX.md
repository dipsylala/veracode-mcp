# CWE-90: LDAP Injection - LLM Guidance

## LLM Guidance

LDAP Injection occurs when untrusted user input is used to construct LDAP queries without proper validation or escaping, allowing attackers to modify queries and access or manipulate directory data. Never concatenate untrusted input into LDAP filters; use safe LDAP APIs and strict allowlists.

## Key Remediation Principles

- Use parameterized LDAP APIs that separate query structure from user data
- Escape special LDAP characters using framework-specific encoding functions
- Apply strict allowlist validation for filter components
- Minimize search scope and restrict returned attributes
- Implement defence-in-depth with input validation, output encoding, and least privilege

## Remediation Steps

- Trace the vulnerability - Identify the source of untrusted data and trace its flow to the LDAP query construction sink
- Replace string concatenation - Use parameterized LDAP APIs that prevent injection by design
- Escape user input - Apply LDAP-specific escaping for special characters - `*`, `(`, `)`, `\`, `NUL`, `/`
- Validate with allowlists - Restrict input to known-safe patterns using allowlists for attribute names and values
- Limit query scope - Use specific base DNs and restrict search depth to minimize exposure
- Test thoroughly - Verify fixes with injection payloads like `*)(objectClass=*)` and review authentication/authorization logic
