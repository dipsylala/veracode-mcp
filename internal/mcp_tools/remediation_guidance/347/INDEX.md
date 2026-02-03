# CWE-347: Improper Verification of Cryptographic Signature

## LLM Guidance

Improper signature verification occurs when applications accept unsigned data, fail to validate signatures, use weak signature algorithms, or implement flawed verification logic. This enables attackers to forge signatures, tamper with signed data, and bypass authentication mechanisms.

## Key Principles

- Verify all signatures before trusting data - Never skip verification or accept unsigned data
- Use strong, approved signature algorithms - Reject weak algorithms (MD5, SHA1) and "alg=none"
- Validate complete certificate chains - Check validity, revocation status, and trust anchors
- Fail securely on verification errors - Reject data immediately on any verification failure
- Apply canonical forms before verification - Prevent signature bypass via data manipulation

## Remediation Steps

- Locate the vulnerability - Review flaw details to identify missing or improper signature verification in your code
- Identify the signature type - Determine what's affected - JWTs, API signatures, certificates, software updates, or documents
- Implement mandatory verification - Always verify signatures before trusting data; never accept unsigned content
- Use approved algorithms - Enforce strong algorithms (RSA-PSS, ECDSA with SHA-256+); reject weak or "none" algorithms
- Validate certificate chains - Check issuer validity, expiration dates, and revocation status (CRL/OCSP)
- Fail securely - Reject data immediately on any verification failure; log security events for monitoring
