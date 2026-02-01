# CWE-183: Permissive List of Allowed Inputs - Java

## LLM Guidance

Permissive input validation occurs when regex patterns or validation logic fails to match the complete input, allowing malicious data to bypass checks. In Java, use fully anchored regex patterns with `^` and `$` anchors and call `matches()` instead of `find()` to ensure complete input validation. Leverage secure APIs like `URI`, `Path`, and `InetAddress` for complex validation scenarios and always enforce length limits to prevent injection attacks.

## Key Principles

- Use `Pattern.matches()` with fully anchored patterns (`^...$`) to validate entire input strings
- Avoid `Matcher.find()` which only matches substrings, creating bypass opportunities
- Implement strict length validation before regex processing to prevent ReDoS attacks
- Use specialized validation classes (`URI`, `Path`, `InetAddress`) for structured data instead of regex
- Apply defense-in-depth with multiple validation layers for critical inputs

## Remediation Steps

- Replace `Pattern.compile(regex).matcher(input).find()` with `input.matches("^" + regex + "$")`
- Add explicit length checks before validation - `if (input.length() > MAX_LENGTH) throw new ValidationException()`
- Use `URI.create()` for URLs, `Paths.get()` for file paths, and `InetAddress.getByName()` for IPs instead of regex
- Compile patterns with anchors - `Pattern pattern = Pattern.compile("^[a-zA-Z0-9]{3,20}$")`
- Test validation with boundary cases including partial matches, empty strings, and oversized inputs
- Log validation failures for security monitoring

## Safe Pattern

```java
private static final Pattern USERNAME_PATTERN = Pattern.compile("^[a-zA-Z0-9_]{3,20}$");
private static final int MAX_LENGTH = 20;

public boolean validateUsername(String username) {
    if (username == null || username.length() > MAX_LENGTH) {
        return false;
    }
    return USERNAME_PATTERN.matcher(username).matches();
}

// For URLs, use specialized classes instead of regex
public boolean validateUrl(String urlString) {
    try {
        URI uri = URI.create(urlString);
        return "https".equals(uri.getScheme());
    } catch (IllegalArgumentException e) {
        return false;
    }
}
```
