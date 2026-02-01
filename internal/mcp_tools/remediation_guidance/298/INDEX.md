# CWE-298: Improper Certificate Validation (Expired)

## LLM Guidance

Improper validation of expired certificates occurs when applications accept SSL/TLS certificates past their validity period, creating security risks where certificate revocation becomes meaningless and compromised private keys cannot be detected. X.509 certificates contain "Not Before" and "Not After" dates that must be enforced. Always validate certificate expiration dates, reject connections with expired certificates, and implement automated renewal before expiration.

## Key Principles

- Always validate certificate expiration dates using TLS library defaults
- Reject connections to services presenting expired certificates
- Never disable certificate validation (avoid `CERT_NONE`, `check_hostname = False`, or custom validation bypasses)
- Use secure default contexts that validate chain of trust, expiration, hostname, and revocation
- Implement automated certificate renewal processes before expiration

## Remediation Steps

- Use `ssl.create_default_context()` in Python which validates expiration, chain of trust, hostname, and revocation
- Avoid `ssl.CERT_NONE` and `context.check_hostname = False` settings
- Don't override `checkServerIdentity` to skip validation in Node.js/JavaScript TLS connections
- Enable expiration checking in all TLS/SSL library configurations
- Monitor certificate expiration dates and set up alerts
- Automate certificate renewal workflows 30-60 days before expiration
