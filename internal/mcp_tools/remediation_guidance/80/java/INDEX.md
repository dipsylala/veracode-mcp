# CWE-80: Cross-Site Scripting (XSS) - Java

## LLM Guidance

Cross-Site Scripting (CWE-80) occurs when untrusted data is included in web pages without proper encoding, allowing attackers to inject malicious scripts that execute in victim browsers. The core fix is to encode all user-controlled output using context-appropriate methods before rendering in HTML, JavaScript, URLs, or CSS contexts.

## Key Principles

- Always encode output based on context (HTML, JavaScript, URL, CSS)
- Use framework-provided encoding mechanisms (JSTL `<c:out>`, Spring escaping, Thymeleaf auto-escaping)
- Validate and sanitize input when encoding alone is insufficient
- Apply Content Security Policy headers to restrict script execution
- Never trust user input, even from authenticated users

## Remediation Steps

- Replace direct JSP expressions (`<%= %>`) with JSTL `<c -out>` tags
- Enable auto-escaping in templating frameworks (Thymeleaf, Freemarker)
- Use OWASP Java Encoder for manual encoding in servlets
- Set `HttpOnly` and `Secure` flags on session cookies
- Implement Content Security Policy headers to block inline scripts
- Review all output points where user data appears

## Safe Pattern

```jsp
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>

<!-- Safe: Auto-escaped output -->
<c:out value="${userInput}" />

<!-- Safe: Explicit escaping in attributes -->
<input type="text" value="<c:out value='${userInput}' />" />

<!-- Servlet-side encoding -->
import org.owasp.encoder.Encode;
response.getWriter().write(Encode.forHtml(userInput));
```
