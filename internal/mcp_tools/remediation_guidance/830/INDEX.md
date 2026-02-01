# CWE-830: Inclusion of Functionality from Untrusted Control Sphere

## LLM Guidance

Applications load code, libraries, or resources from attacker-controlled sources without verifying integrity or authenticity. This includes dynamic code loading from user-controlled paths, remote scripts without integrity checks, untrusted plugins/extensions, and compromised package managers. Fix by loading only from trusted sources with cryptographic verification and blocking user control of inclusion paths.

## Key Principles

- Load code and resources only from trusted, verified sources; implement integrity checks using cryptographic hashes; never allow user input to control code inclusion paths.
- Implement Subresource Integrity (SRI) for all external scripts and stylesheets
- Use allowlists to restrict code sources to specific trusted domains and paths
- Validate cryptographic signatures for dynamically loaded modules and plugins
- Lock dependency versions with integrity hashes in package managers
- Isolate dynamic code execution in sandboxed environments with restricted permissions

## Remediation Steps

- Add SRI `integrity` and `crossorigin` attributes to all `<script>` and `<link>` tags loading external resources
- Configure Content Security Policy (CSP) headers to restrict script sources to trusted domains only
- Verify cryptographic signatures before loading plugins, extensions, or dynamically imported modules
- Use lock files (package-lock.json, yarn.lock) with integrity hashes for all dependencies
- Implement server-side path validation that rejects user-supplied paths containing traversal sequences
- Review and audit all dynamically loaded code sources; eliminate unnecessary dynamic loading
