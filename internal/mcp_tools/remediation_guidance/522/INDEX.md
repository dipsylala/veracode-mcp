# CWE-522: Insufficiently Protected Credentials

## LLM Guidance

This vulnerability occurs when credentials (passwords, API keys, tokens) are stored or transmitted without adequate protection, making them susceptible to theft or misuse. The core fix is to ensure credentials are protected in transit, at rest, and during processing, never appearing in logs, URLs, or client-visible data.

## Key Principles

- Never hardcode credentials in source code or configuration files committed to version control
- Protect credentials in transit using TLS/HTTPS and at rest using strong encryption
- Use secure credential storage systems (secrets managers, hardware security modules, encrypted vaults)
- Never expose credentials in logs, error messages, URLs, or client-side code
- Apply least-privilege access and rotate credentials regularly

## Remediation Steps

- Identify all credential locations - review code, configuration files, databases, and transmission points for exposed passwords, API keys, tokens, or private keys
- Remove hardcoded credentials from source code and move to environment variables or secrets management systems (AWS Secrets Manager, Azure Key Vault, HashiCorp Vault)
- Implement strong encryption - use bcrypt, Argon2, or PBKDF2 for password hashing; encrypt sensitive data at rest with AES-256
- Secure transmission - enforce HTTPS/TLS for all credential transmission; never send credentials in URL parameters or GET requests
- Implement proper access controls - use multi-factor authentication, enforce least-privilege principles, and rotate credentials on a regular schedule
- Remove credentials from version control history using tools like git-filter-repo or BFG Repo-Cleaner
