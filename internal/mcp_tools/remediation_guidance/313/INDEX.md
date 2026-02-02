# CWE-313: Cleartext Storage in a File or on Disk

## LLM Guidance

Storing sensitive data (passwords, API keys, PII, credit cards, session tokens) unencrypted in files or databases enables data theft through backup exposure, filesystem access, database dumps, stolen devices, or insider threats. The core fix is to encrypt sensitive data at rest using strong encryption algorithms and store encryption keys in secure key management systems, never alongside the encrypted data.

## Key Principles

- Never store sensitive data in cleartext on disk; always use encryption at rest
- Use strong authenticated encryption (AES-256-GCM) for data protection
- Store encryption keys separately in dedicated key management systems (AWS KMS, Azure Key Vault, HashiCorp Vault)
- Apply least privilege access controls to encrypted data and key storage
- Avoid storing sensitive data unnecessarily; minimize what you persist

## Remediation Steps

- Review flaw details to identify the specific file, line number, and storage location where sensitive data is written unencrypted
- Determine what sensitive data is at risk - passwords, API keys, PII, session tokens, or credentials
- Implement AES-256-GCM encryption for all sensitive data before writing to disk or database
- Store encryption keys in a secure key management system (AWS KMS, Azure Key Vault, HashiCorp Vault), never in code or configuration files
- Apply file system permissions and access controls to restrict who can read encrypted files
- Audit existing files, databases, backups, and logs to remove any cleartext sensitive data already persisted
