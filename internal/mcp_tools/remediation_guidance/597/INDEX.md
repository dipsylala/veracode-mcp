# CWE-597: Use of Wrong Operator in String Comparison

## LLM Guidance

Using reference equality (`==`) instead of value equality (`.equals()`) for string comparison in **Java** compares memory addresses, not content, causing security checks to fail unpredictably when strings are dynamically created vs literals, enabling authentication bypass and logic errors. **Note:** In C#, `==` is overloaded for strings and performs value comparison, making it safe to use.

**Primary Defence:** Always use `.equals()` in Java for string content comparison; use constant-first pattern for null safety; prefer enums over strings for security-critical values like roles and permissions. In C#, both `==` and `.Equals()` work for string comparison, though `.Equals()` with `StringComparison` options provides more control.

## Key Principles

- Use `.equals()` in Java for all string content comparisons, never `==` or `!=`
- Apply constant-first pattern (e.g., `"admin".equals(userRole)`) to prevent NullPointerException
- Use constant-time comparison functions for sensitive data (passwords, tokens, secrets)
- In C#, `==` is safe for strings but use `String.Equals()` for explicit case-sensitivity control
- Prioritize security-critical code: authentication, authorization, token validation

## Remediation Steps

- Search codebase for `== "` and `!= "` patterns in Java files to locate vulnerable comparisons
- Identify security-critical comparisons - passwords, roles, tokens, API keys, session IDs
- Replace `str1 == str2` with `str1.equals(str2)` or constant-first `"constant".equals(variable)`
- For secrets, use `MessageDigest.isEqual()` or constant-time comparison libraries
- Add null-safety checks or use `Objects.equals(str1, str2)` for null-safe comparison
- Verify fixes with unit tests covering both matching and non-matching string scenarios
