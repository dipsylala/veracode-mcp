# CWE-829: Inclusion of Functionality from Untrusted Control Sphere

## LLM Guidance

This vulnerability occurs when applications include code, libraries, or functionality from untrusted sources, allowing attackers to inject malicious behavior. The core fix is to verify the integrity and provenance of all executable code before inclusion, ensuring it originates only from trusted, controlled sources.

## Key Principles

- Never execute or load functionality without explicitly verifying its integrity and provenance
- Restrict all executable code to trusted, controlled sources with verified authenticity
- Implement strict validation and verification mechanisms for any dynamically loaded components
- Apply defence-in-depth by combining source verification, integrity checks, and runtime controls

## Remediation Steps

- Identify the vulnerability - Review flaw details for specific file, line number, and code pattern where untrusted inclusion occurs (dynamic imports, eval(), exec(), third-party libraries)
- Trace data flow - Determine where the code source originates (URL, user input, external API, package manager) and assess what privileges and data access the included code has
- Use verified sources only - Download libraries exclusively from official package repositories and trusted vendors with established reputations
- Implement integrity verification - Use cryptographic checksums, digital signatures, or subresource integrity (SRI) to validate code before execution
- Remove dynamic code execution - Eliminate eval(), exec(), and dynamic imports; replace with static alternatives or whitelist-based approaches
- Apply allowlisting - If dynamic loading is required, maintain a strict allowlist of approved components and block all others by default
