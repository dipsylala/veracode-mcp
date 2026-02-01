# CWE-526: Information Exposure Through Environment Variables

## LLM Guidance

Information exposure through environment variables occurs when applications unintentionally leak sensitive configuration data, credentials, API keys, or system information stored in environment variables. This typically happens through error pages displaying environment dumps, debugging endpoints exposing process state, server-side template injection rendering environment context, or misconfigured logging that captures full environment. The core fix is to never expose environment variables to untrusted users and filter environment data from all outputs.

## Key Principles

- Never expose environment variables to untrusted users through any interface
- Use secure secret management systems instead of environment variables for sensitive data
- Filter environment data from all user-facing outputs including errors, logs, and debug endpoints
- Disable debug mode and verbose error reporting in production environments
- Implement least-privilege access controls for environment variable visibility

## Remediation Steps

- Remove or protect debug endpoints that expose environment variables with authentication
- Set `DEBUG=False` in production to prevent stack traces from revealing environment state
- Implement custom error handlers that log full details server-side but return generic messages to clients
- Audit codebase for any environment variable exposure points (error pages, logging, templates, API responses)
- Migrate secrets from environment variables to dedicated secret management systems (AWS Secrets Manager, HashiCorp Vault, Azure Key Vault)
- Configure logging frameworks to explicitly exclude environment variables from captured context
