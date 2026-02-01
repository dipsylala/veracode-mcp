# CWE-183: Permissive List of Allowed Inputs

## LLM Guidance

Permissive allowlists occur when input validation accepts too broad a range of values, allowing attackers to bypass security controls with edge cases, encoding variations, or unexpected but technically valid inputs. This enables injection attacks, path traversal, or logic bypass vulnerabilities. The core fix is implementing strict allowlists with default-deny policies.

## Key Principles

- Use strict allowlists with default-deny; reject everything not explicitly permitted
- Anchor regex patterns and avoid overly broad wildcards (`.+`, `.*`, `\w+`)
- Validate against exact expected formats, not permissive patterns
- Apply validation at all trust boundaries before dangerous operations
- Consider canonicalization and encoding variations when defining allowed inputs

## Remediation Steps

- Locate validation logic in your code (regex patterns, allowlists, format checks)
- Review allowlists for overly broad patterns or missing constraints
- Replace permissive patterns with strict, anchored validation (e.g., `^[a-z]{3,10}$` instead of `\w+`)
- Add validation for inputs that lack security checks
- Test edge cases - special characters, encoding variations, path traversal sequences
- Ensure validated data is used safely in downstream operations (SQL, file paths, commands)
