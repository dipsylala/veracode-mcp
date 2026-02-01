# CWE-614: Sensitive Cookie Without 'Secure' Flag - Java

## LLM Guidance

Sensitive cookies in Java web applications transmitted without the `secure` attribute can be intercepted over HTTP connections, enabling man-in-the-middle attacks and session hijacking. The fix requires setting `setSecure(true)` on all cookies containing authentication tokens, session IDs, or other sensitive data to ensure transmission only over HTTPS.

## Key Principles

- Set `secure` attribute to `true` on all cookies containing sensitive information
- Enforce HTTPS-only transmission for authentication and session cookies
- Apply `HttpOnly` flag alongside `Secure` to prevent client-side script access
- Configure framework-level defaults for secure cookie creation
- Validate that production environments use HTTPS exclusively

## Remediation Steps

- Identify all cookie creation points (Servlet `Cookie`, Spring `ResponseCookie`, framework configurations)
- Add `cookie.setSecure(true)` to every sensitive cookie instantiation
- Set `HttpOnly` flag with `cookie.setHttpOnly(true)` for additional protection
- Configure Spring Security or servlet container defaults for secure cookies
- Test in HTTPS environment to verify cookies transmit only over secure channels
- Review session management configuration in `web.xml` or application properties

## Safe Pattern

```java
// Servlet API
Cookie sessionCookie = new Cookie("JSESSIONID", sessionId);
sessionCookie.setSecure(true);
sessionCookie.setHttpOnly(true);
sessionCookie.setPath("/");
response.addCookie(sessionCookie);

// Spring Framework
ResponseCookie cookie = ResponseCookie.from("authToken", token)
    .secure(true)
    .httpOnly(true)
    .path("/")
    .maxAge(3600)
    .build();
response.addHeader("Set-Cookie", cookie.toString());
```
