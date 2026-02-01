# CWE-494: Download of Code Without Integrity Check

## LLM Guidance

This vulnerability occurs when applications download code or executables from external sources without verifying their integrity, allowing attackers to inject malicious code. Only download and install code with integrity verification (cryptographic signatures or hashes) from trusted sources over secure transport.

## Key Principles

- Verify integrity using cryptographic hashes (SHA-256/SHA-512) or digital signatures before executing downloaded code
- Download code only from trusted, authenticated sources over secure channels (HTTPS/TLS)
- Implement checksum verification for all packages, plugins, scripts, and executables
- Use package managers with built-in integrity checking and signed repositories
- Fail securely if integrity verification fails—never execute unverified code

## Remediation Steps

- Identify download operations - Review scan data_paths to find HTTP downloads, package installations, plugin loading, and script fetching
- Check for missing verification - Look for absent hash checks, signature validation, or checksum verification in download code
- Add integrity checks - Implement cryptographic hash verification (SHA-256+) before executing or importing downloaded code
- Use secure sources - Replace HTTP with HTTPS; use official package repositories with signature verification
- Validate before execution - Ensure downloaded files match expected hashes/signatures before `exec()`, `import`, or plugin load operations
- Handle failures securely - Reject and log downloads that fail integrity checks; do not fall back to unverified execution
