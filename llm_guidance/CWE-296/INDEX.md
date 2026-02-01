# CWE-296: Improper Certificate Validation (Trust Chain)

## LLM Guidance

Improper certificate validation occurs when applications fail to verify that X.509 certificates chain properly to a trusted root Certificate Authority (CA), enabling man-in-the-middle attacks using self-signed or untrusted certificates. This commonly happens when certificate validation is disabled during development and mistakenly remains disabled in production. Always validate the complete certificate trust chain to known trusted root CAs.

## Key Principles

- Validate complete certificate trust chain to known trusted root CAs in all environments
- Never disable certificate validation or trust unverified certificates in production
- Verify certificate expiration, revocation status (CRL/OCSP), and intermediate certificates
- Use TLS/SSL library defaults that enforce proper chain validation

## Remediation Steps

- Enable certificate chain validation in TLS/SSL libraries (set verify=True, never verify=False)
- Ensure certificates chain to trusted root CAs with valid intermediate certificates included
- Check certificate expiration dates and revocation status via CRL/OCSP
- Remove any code disabling verification (verify=False, rejectUnauthorized - false, SSL_VERIFY_NONE)
- Use platform/library default certificate stores for trusted root CAs
- Test TLS configurations to confirm proper chain validation is enforced
