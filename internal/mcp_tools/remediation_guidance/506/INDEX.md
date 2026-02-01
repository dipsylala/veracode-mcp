# CWE-506: Embedded Malicious Code

## LLM Guidance

Embedded malicious code includes backdoors, logic bombs, time bombs, and trojans intentionally placed in source code or dependencies. Compromised npm packages, malicious gems, and supply chain attacks inject code that steals credentials, creates backdoors, or triggers on specific conditions.

## Remediation Strategy

- Enforce mandatory code review for all changes, especially from new contributors
- Implement dependency signing and provenance verification
- Scan for obfuscated code patterns (base64, eval/exec, encoded payloads)
- Monitor for unexpected behaviors like unauthorized network calls or file access
- Establish supply chain security controls for all third-party dependencies

## Remediation Steps

- Examine all dependencies - Review npm packages, pip packages, gems, JARs for suspicious activity
- Review recent code changes - Check commits for obfuscation, unexpected network calls, or process execution
- Search for red flags - `eval(base64.b64decode(...))`, `exec(__import__('base64').b64decode(...))`, hardcoded credentials
- Identify backdoors - Look for hidden authentication bypasses, unknown domain connections
- Remove malicious code - Delete compromised dependencies, revert suspicious commits, replace with verified versions
- Verify provenance - Use lock files, check package signatures, audit dependency chains
