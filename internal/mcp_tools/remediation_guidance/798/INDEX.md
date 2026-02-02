# CWE-798: Use of Hard-coded Credentials

## LLM Guidance

Hard-coded credentials occur when authentication secrets (passwords, API keys, encryption keys, tokens) are embedded directly in source code, configuration files, or binaries. This violates the principle of separation of code and configuration - credentials become visible to anyone with code access, changing them requires redeployment, they persist in version control history, and credential rotation becomes nearly impossible.

## Remediation Steps

- Identify the sink and confirm the data path from untrusted data
- Apply the primary safe pattern for this CWE
- Add allowlist validation or encoding where required
- Verify behavior with normal and boundary cases

## Key Principles

1. Remove hard-coded secrets from source and config.
2. Load secrets from environment variables or a secrets manager.
3. Rotate and revoke exposed credentials.
4. Restrict secret access by least privilege.
