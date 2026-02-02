# CWE-295: Improper Certificate Validation

## LLM Guidance

Improper certificate validation occurs when an application fails to correctly verify the authenticity of SSL/TLS certificates during secure connections. This undermines the entire purpose of HTTPS by allowing attackers to intercept encrypted communications (man-in-the-middle attacks), impersonate legitimate servers, and steal credentials or sensitive data in transit.

SSL/TLS certificates serve two critical purposes: authentication (prove the server is who it claims to be) and encryption (establish a secure encrypted channel). When certificate validation is disabled or improperly implemented, encryption alone is insufficient - you may be encrypting data to an attacker's server.

## Key Principles

1. Replace unsafe sinks with safe native APIs or library functions.
2. Apply the primary safe pattern for this CWE.
3. Validate untrusted data with strict allowlists and type checks.
4. Apply least privilege and safe defaults.

## Remediation Steps

- Identify the sink and confirm the data path from untrusted data
- Apply the primary safe pattern for this CWE
- Add allowlist validation or encoding where required
- Verify behavior with normal and boundary cases
