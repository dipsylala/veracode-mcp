# CWE-321 - Use of Hard-coded Cryptographic Key

## LLM Guidance

Hard-coded cryptographic keys in source code, configuration files, or binaries are exposed to anyone with codebase access. This compromises all encrypted data and prevents key rotation without redeployment.

## Key Principles

- Never embed keys in application artifacts (source, binaries, images, client-side config)
- Keys must be provided at runtime from dedicated key-management systems
- Treat keys as replaceable secrets with rotation capabilities
- Separate key storage from application deployment pipeline
- Use environment-specific keys with proper access controls

## Remediation Steps

- Locate hard-coded keys - Search for patterns like `const KEY =`, `SECRET_KEY =`, base64 strings, or keys in config files committed to version control
- Remove keys from code - Delete hard-coded values and replace with runtime retrieval from secure sources
- Implement secure key storage - Use environment variables, vault services (AWS KMS, Azure Key Vault, HashiCorp Vault), or platform-native secret managers
- Load keys at runtime - Retrieve keys from secure storage during application initialization
- Rotate compromised keys - Generate new keys, update all encrypted data, and revoke old keys
- Audit version control - Remove keys from git history using tools like BFG Repo-Cleaner or git-filter-repo
