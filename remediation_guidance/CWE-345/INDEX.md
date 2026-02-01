# CWE-345: Insufficient Verification of Data Authenticity

## LLM Guidance

Insufficient verification of data authenticity occurs when applications don't validate that data hasn't been tampered with during storage or transmission. This includes accepting unsigned data, not verifying signatures/MACs, or trusting unauthenticated sources, enabling data tampering, message forgery, and MITM attacks.

## Key Principles

- Never allow untrusted data to influence security or control decisions unless its authenticity is verified by a server-controlled integrity mechanism
- Use cryptographic signatures or MACs to verify data has not been modified
- Validate data authenticity at trust boundaries before processing
- Implement server-side verification; never rely solely on client-side checks
- Apply defense-in-depth with multiple verification layers for critical data

## Remediation Steps

- Review flaw details to identify files, line numbers, and code patterns where data authenticity verification is missing
- Identify what data lacks authentication - cookies, tokens, API messages, file uploads, database records, or configuration data
- Determine the data source and trace the flow to understand where unauthenticated data is accepted and used
- Implement HMAC-SHA256 or similar message authentication codes for verifying data integrity
- Add cryptographic signature verification for critical operations and data exchanges
- Validate all signatures/MACs server-side before trusting or processing the data
