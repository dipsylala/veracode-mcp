# CWE-547: Use of Hard-coded, Security-relevant Constants

## LLM Guidance

Hard-coded security-relevant constants (paths, ports, IPs, credentials, encryption keys) in source code create deployment failures, environment coupling, and security risks by preventing configuration changes without recompilation. These values often leak sensitive information and fail when moved between environments. The core fix is externalizing all configuration to environment variables, config files, or secret management systems.

## Key Principles

- Never hard-code credentials, keys, tokens, bypass flags, or security parameters in source code
- Use environment variables or configuration files for all deployment-specific values (hosts, ports, paths)
- Implement centralized secret management (AWS Secrets Manager, Azure Key Vault, HashiCorp Vault) for sensitive data
- Ensure configuration changes don't require code recompilation or redeployment
- Apply principle of least privilege to configuration access and rotation

## Remediation Steps

- Identify hard-coded constants - Database hosts/ports, file paths, IP addresses, URLs, admin credentials, timeout values, encryption algorithms
- Search codebase for patterns - IP addresses (`\d+\.\d+\.\d+\.\d+`), platform-specific paths (`C -\`, `/var/`), port numbers, connection strings
- Extract to external configuration - Move hard-coded values to environment variables, `.env` files, or config management systems
- Replace with dynamic lookups - Use `os.getenv()`, config readers, or dependency injection to load values at runtime
- Secure sensitive values - Store secrets in vaults, never in version control; rotate credentials regularly
- Validate at startup - Check required configuration exists on application start with clear error messages for missing values
