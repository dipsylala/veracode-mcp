# CWE-223: Omission of Security-Relevant Information

## LLM Guidance

Omission of security-relevant information occurs when applications fail to log critical security events such as login failures, access denials, privilege escalations, and data modifications. This prevents effective security monitoring, incident response, compliance auditing, and attack detection. The core fix is implementing comprehensive logging of all security-relevant events while ensuring logged information doesn't expose sensitive data.

## Remediation Strategy

- Ensure exceptions and failure modes do not disclose sensitive data or bypass security checks; fail closed
- Log all authentication and authorization events including both successes and failures
- Implement structured logging with sufficient context for security analysis and forensics
- Comply with regulatory audit requirements (PCI-DSS, HIPAA, SOX, GDPR)

## Remediation Steps

- Find unlogged security events - Identify missing logs for authentication attempts (login/logout), authorization failures, privilege changes, sensitive data access, and configuration modifications
- Check authentication flows - Ensure logging covers login, logout, password reset, MFA, and session lifecycle events
- Review authorization points - Log all access denied events, role/permission checks, and resource access attempts
- Identify sensitive operations - Track data modifications (create, update, delete), administrative actions, and configuration changes
- Trace audit requirements - Map logging to compliance obligations and regulatory standards
