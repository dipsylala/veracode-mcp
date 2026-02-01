# CWE-297 - Improper Validation of Certificate with Host Mismatch

## LLM Guidance

Improper certificate hostname validation occurs when TLS/SSL clients don't verify the certificate's Common Name (CN) or Subject Alternative Name (SAN) matches the hostname being connected to. This enables man-in-the-middle attacks by allowing attackers to present valid certificates for different domains.

## Key Principles

- Never establish a TLS connection unless the certificate's hostname is validated against the requested peer identity
- Applications must not bypass or weaken hostname verification (avoid SSL_VERIFY_NONE, check_hostname=False)
- Use default SSL/TLS library configurations which enable hostname verification by default
- Custom certificate validators must explicitly check CN/SAN fields against the target hostname
- Hostname verification is separate from certificate chain validation—both are required

## Remediation Steps

- Locate the vulnerability - Review flaw details for specific file, line number, and code pattern where hostname verification is disabled or weakened
- Identify bypass patterns - Search for SSL_VERIFY_NONE, check_hostname=False, InsecureRequestWarning suppressions, or custom validators
- Enable default verification - Use standard library defaults with hostname checking enabled (e.g., requests with verify=True, Python ssl with check_hostname=True)
- Remove verification bypasses - Delete or refactor any code that disables certificate or hostname validation
- Test secure connections - Verify connections fail appropriately with invalid/mismatched certificates
- Review all TLS clients - Audit entire codebase for similar patterns in other connection points
