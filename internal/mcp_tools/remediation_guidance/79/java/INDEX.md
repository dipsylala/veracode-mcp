# CWE-79: Cross-Site Scripting (XSS) - Java

## LLM Guidance

Cross-Site Scripting (CWE-79) occurs when untrusted data is included in web pages without proper encoding, allowing attackers to inject malicious scripts that execute in victim browsers to steal sessions, harvest credentials, or distribute malware. Java applications must encode all user-controlled output using context-appropriate methods (HTML, JavaScript, URL, CSS). Fix by applying output encoding at every injection point based on the context where data appears.

## Key Principles

- Always encode output based on context (HTML entity encoding for HTML content, JavaScript encoding for JS contexts, URL encoding for URLs)
- Use established encoding libraries like OWASP Java Encoder or Spring's HtmlUtils rather than custom sanitization
- Validate and sanitize input as defense-in-depth, but never rely on input validation alone
- Set Content Security Policy (CSP) headers to restrict script execution sources
- Use HTML templating engines with auto-escaping enabled by default

## Remediation Steps

- Identify all locations where user input or external data is rendered in responses
- Apply context-appropriate encoding using OWASP Java Encoder at each output point
- Replace innerHTML assignments with textContent or properly encoded alternatives
- Enable auto-escaping in your templating engine (JSP, Thymeleaf, FreeMarker)
- Implement CSP headers with strict policies (`default-src 'self'`)
- Review and test all dynamic content rendering paths

## Safe Pattern

```java
import org.owasp.encoder.Encode;

@GetMapping("/profile")
public String showProfile(@RequestParam String name, Model model) {
    // Encode for HTML context before rendering
    String safeName = Encode.forHtml(name);
    model.addAttribute("userName", safeName);
    return "profile";
}

// In JSP with JSTL c:out (auto-escapes):
// <c:out value="${userName}" />

// Or in Thymeleaf (auto-escapes by default):
// <span th:text="${userName}"></span>
```
