# CWE-601: Open Redirect - Java

## LLM Guidance

Open redirect vulnerabilities in Java web applications occur when user-controlled input is used in `sendRedirect()`, `forward()`, or `<meta>` refresh tags without proper validation, enabling phishing attacks and credential theft. The core fix is to validate redirect URLs against an allowlist of trusted destinations or use relative paths only. For Spring MVC and Jakarta EE applications, use path-based routing instead of accepting arbitrary URLs.

## Key Principles

- Allowlist validation: Match redirect URLs against a predefined list of permitted domains or paths
- Relative paths only: Use context-relative paths (`/dashboard`) instead of accepting full URLs
- Avoid user input in redirects: Use indirect references (e.g., enums, IDs mapped to destinations)
- Strict URL parsing: Validate protocol, domain, and path components before redirecting
- Framework-level guards: Configure Spring Security or servlet filters to block external redirects

## Remediation Steps

- Identify all `response.sendRedirect()`, `return "redirect -"`, and `<meta http-equiv="refresh">` usage
- Replace user-controlled URLs with enum/ID mapping to predefined destinations
- For necessary external redirects, validate against an allowlist of trusted domains
- Use `URI` class to parse and validate URL components (scheme, host, path)
- Implement a centralized redirect validator for consistent enforcement
- Add unit tests verifying that malicious URLs (`//evil.com`, `https -//attacker.com`) are rejected

## Safe Pattern

```java
// Allowlist-based redirect validator
private static final Set<String> ALLOWED_HOSTS = Set.of("example.com", "app.example.com");

public String safeRedirect(String userUrl, HttpServletResponse response) {
    try {
        URI uri = new URI(userUrl);
        String host = uri.getHost();
        
        if (host != null && !ALLOWED_HOSTS.contains(host)) {
            return "redirect:/error"; // Reject external redirects
        }
        return "redirect:" + uri.getPath(); // Use path only
    } catch (URISyntaxException e) {
        return "redirect:/error";
    }
}
```
