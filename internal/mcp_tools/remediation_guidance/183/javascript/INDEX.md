# CWE-183: Permissive List of Allowed Inputs - JavaScript/TypeScript

## LLM Guidance

CWE-183 occurs when input validation uses overly permissive patterns that fail to reject malicious inputs, allowing attackers to bypass security controls through crafted strings. In JavaScript/TypeScript, this commonly happens with unanchored regex, incomplete URL validation, or loose string matching. Use fully anchored regex patterns (`^...$`), native validation APIs (URL constructor, path.resolve()), Set-based allowlists, and strict length limits to ensure complete input validation.

## Key Principles

- Always anchor regex patterns with `^` and `$` to match entire input
- Prefer native APIs (URL, path) over regex for structured data validation
- Use Set-based lookups for discrete allowlists instead of pattern matching
- Enforce strict length limits before validation
- Validate normalized/canonical forms to prevent bypass techniques

## Remediation Steps

- Replace unanchored regex with fully anchored patterns using `^` and `$`
- Implement Set-based allowlists for known valid values
- Use URL constructor for URL validation and path.resolve() for file paths
- Add maximum length checks before validation logic
- Normalize inputs (lowercase, trim) before allowlist comparison
- Throw errors on validation failure; never fall back to permissive defaults

## Safe Pattern

```javascript
function validateAllowedDomain(input) {
  const ALLOWED_DOMAINS = new Set(['example.com', 'trusted.org']);
  const MAX_LENGTH = 253;
  
  if (!input || input.length > MAX_LENGTH) return false;
  
  try {
    const url = new URL(input);
    return ALLOWED_DOMAINS.has(url.hostname.toLowerCase());
  } catch {
    return false;
  }
}
```
