# CWE-597: Use of Wrong Operator in String Comparison - Java

## LLM Guidance

Using reference equality (`==`) instead of value equality (`.equals()`) for string comparison in Java compares memory addresses, not content, causing security checks to fail unpredictably when strings are dynamically created vs literals, enabling authentication bypass and logic errors.

**Primary Defence:** Always use `.equals()` for string content comparison; use constant-first pattern for null safety.

## Key Principles

- Replace `==` with `.equals()` for all string content comparisons
- Use `"expected".equals(variable)` pattern to prevent NullPointerException
- Use `Objects.equals(a, b)` when both strings may be null
- Consider `equalsIgnoreCase()` for case-insensitive comparisons
- Never use `==` except for explicit null checks

## Remediation Steps

- Scan codebase for `string1 == string2` patterns in conditionals and authentication logic
- Replace with `.equals()`, putting known constant first where possible
- Add null checks or use `Objects.equals()` where both values may be null
- Prioritize authentication, authorization, and security-critical paths first
- Run comprehensive test suite to verify logic correctness after changes
- Enable static analysis rules (e.g., PMD, SpotBugs) to prevent future violations

## Safe Pattern

```java
// Authentication check - safe string comparison
public boolean authenticate(String username, String password) {
    String validUser = "admin";
    String validPass = retrieveHashedPassword(username);
    
    // Constant-first pattern prevents NPE, compares content not references
    if (validUser.equals(username) && validPass != null) {
        return validPass.equals(hashPassword(password));
    }
    return false;
}
```
