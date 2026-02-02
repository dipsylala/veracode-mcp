# CWE-597: Use of Wrong Operator in String Comparison - C\#

## LLM Guidance

Using reference equality (`==` on object types) instead of value equality (`.Equals()`) for string comparison in C# can compare memory addresses rather than content when strings are not interned, causing security checks to fail unpredictably and enabling authentication bypass and logic errors. Always use `.Equals()` method or `String.Equals()` with explicit `StringComparison` parameter for security-critical comparisons.

## Key Principles

- Use `.Equals()` or `String.Equals()` for all string content comparisons, never `==` for validation logic
- Specify `StringComparison.Ordinal` for case-sensitive or `StringComparison.OrdinalIgnoreCase` for case-insensitive comparisons
- Avoid culture-sensitive comparisons (`CurrentCulture`) in authentication, authorization, and security checks
- Never rely on string interning for security decisions

## Remediation Steps

- Identify all string comparisons using `==` or `!=` operators in security-sensitive code paths
- Replace with `.Equals()` method calls with explicit `StringComparison` parameter
- For null-safe comparisons, use `String.Equals(str1, str2, StringComparison.Ordinal)`
- Review authentication, authorization, token validation, and input validation logic
- Add unit tests covering non-interned strings to verify correct comparison behavior
- Use static analysis tools to flag `==` usage on string types

## Safe Pattern

```csharp
// UNSAFE: Reference comparison
if (userRole == "Admin") { /* grant access */ }

// SAFE: Value comparison with explicit mode
if (userRole.Equals("Admin", StringComparison.Ordinal)) { /* grant access */ }

// SAFE: Null-safe static method
if (String.Equals(userRole, "Admin", StringComparison.Ordinal)) { /* grant access */ }

// SAFE: Case-insensitive when appropriate
if (fileExtension.Equals(".exe", StringComparison.OrdinalIgnoreCase)) { /* block */ }
```
