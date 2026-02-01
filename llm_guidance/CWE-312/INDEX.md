# CWE-312 - Cleartext Storage of Sensitive Information

## LLM Guidance

Cleartext storage occurs when sensitive information (credentials, PII, financial data, cryptographic keys) is stored without encryption in databases, files, logs, cache, backups, or memory dumps, making it readable to anyone who gains access. Unlike cleartext transmission (CWE-319), this vulnerability affects data at rest. Core fix - Encrypt sensitive data at rest or redact before persisting.

## Remediation Strategy

- Never store sensitive information in cleartext; always encrypt or redact before persisting
- Use strong encryption algorithms (AES-256) for data at rest
- Implement defense-in-depth - combine database encryption, column-level encryption, and application-level encryption
- Protect encryption keys separately from encrypted data
- Redact sensitive data from logs, cache, and temporary files

## Remediation Steps

- Identify storage locations - Review flaw details to locate where sensitive data is stored (database columns, configuration files, logs, cache, temporary files, cloud storage)
- Classify data type - Determine what's exposed (credentials, API keys, PII, financial data, health records, cryptographic keys)
- Implement database encryption - Use transparent data encryption (TDE) for entire database; apply column-level encryption with AES-256 for extra-sensitive fields
- Encrypt files and configuration - Use filesystem-level or application-level encryption for sensitive configuration files and local storage
- Redact from logs - Remove or mask sensitive data before writing to logs, error messages, or debug output
- Secure encryption keys - Store keys in dedicated key management systems (AWS KMS, Azure Key Vault, HashiCorp Vault), never alongside encrypted data
