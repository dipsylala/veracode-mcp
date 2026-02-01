# CWE-311: Missing Encryption of Sensitive Data

## LLM Guidance

Missing Encryption occurs when sensitive data is transmitted or stored without proper cryptographic protection, exposing it to unauthorized access or interception. The core fix is to encrypt sensitive data in transit using TLS 1.2+ and at rest when storage systems are untrusted or exposure is plausible.

## Key Principles

- Treat storage as untrusted—encrypt sensitive data at rest when exposure is plausible
- Use TLS 1.2+ for all network communications transmitting sensitive data
- Never transmit credentials, PII, or secrets over unencrypted channels
- Apply defense-in-depth: combine encryption with access controls and secure key management

## Remediation Steps

- Identify unencrypted sensitive data by reviewing flaw details (file, line, data type - passwords, PII, tokens, API keys)
- Trace data flow to determine if exposure is in transit (network) or at rest (storage, database, files)
- Enforce TLS 1.2+ for all network communications - HTTPS for web traffic, TLS for database connections, secure WebSocket (wss -//)
- Encrypt sensitive data at rest using AES-256 or equivalent when stored in databases, files, or untrusted systems
- Implement secure key management - use hardware security modules, key vaults, or secrets management services—never hardcode keys
- Validate encryption coverage - verify no plaintext sensitive data exists in logs, backups, or temporary files
