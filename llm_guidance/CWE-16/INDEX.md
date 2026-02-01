# CWE-16: Configuration

## LLM Guidance

Configuration vulnerabilities occur when applications are deployed with insecure settings, missing security controls, or default configurations that expose attack surfaces. This includes misconfigured security headers, debug mode in production, exposed administrative interfaces, default credentials, and weak cryptographic settings. These issues stem from incomplete security hardening and failure to follow secure deployment practices.

## Key Principles

- Never deploy with default or insecure configurations in production
- Disable all debug and development modes before deployment
- Implement proper security headers and least-privilege access controls
- Follow security hardening guidelines for all production systems
- Remove or secure administrative interfaces and sample applications

## Remediation Steps

- Review security scan results to identify specific misconfigurations (missing headers, debug mode, defaults)
- Disable debug/verbose modes that expose stack traces or detailed error messages
- Remove or secure default credentials, sample applications, and unnecessary services
- Implement required security headers (CSP, X-Frame-Options, HSTS, X-Content-Type-Options)
- Restrict access to administrative interfaces and API documentation with authentication
- Configure strong TLS/SSL settings and update weak cryptographic configurations
