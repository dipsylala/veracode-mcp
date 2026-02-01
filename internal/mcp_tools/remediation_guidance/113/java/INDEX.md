# CWE-113: HTTP Response Splitting - Java

## LLM Guidance

HTTP Response Splitting occurs when attackers inject CRLF characters (`\r\n`) into HTTP headers, enabling them to inject additional headers or response bodies, leading to cache poisoning, XSS, or session hijacking. Use Spring Framework's built-in redirect methods and header builders that automatically sanitize inputs; never manually construct headers with untrusted data.

## Key Principles

- Use framework-provided abstractions: Spring's `redirect:` prefix, `RedirectView`, `ResponseCookie.from()`, and `ContentDisposition.builder()` handle encoding automatically
- Validate and sanitize user input: Reject CRLF sequences (`\r`, `\n`) from any data used in headers
- Apply allowlisting: Restrict header values to expected character sets (alphanumeric, safe punctuation)
- Avoid manual header construction: Never concatenate user input directly into headers
- Use safe APIs: Prefer `UriComponentsBuilder` for URL construction with proper encoding

## Remediation Steps

- Replace manual `response.setHeader()` calls with Spring's `ResponseCookie.from()` builder for cookies
- Use `redirect -` prefix in controller return values instead of manually setting `Location` headers
- Apply input validation to reject `\r` and `\n` characters before any header operations
- Use `ContentDisposition.builder()` for file download headers instead of string concatenation
- Implement allowlist validation for redirect URLs using `UrlValidator` or regex patterns
- Review all response header manipulations and replace with framework methods

## Safe Pattern

```java
@Controller
public class SafeRedirectController {
    
    @GetMapping("/redirect")
    public String safeRedirect(@RequestParam String url) {
        // Validate against allowlist
        if (!url.matches("^/[a-zA-Z0-9/_-]+$")) {
            throw new IllegalArgumentException("Invalid URL");
        }
        return "redirect:" + url; // Spring handles encoding
    }
    
    @GetMapping("/cookie")
    public ResponseEntity<Void> safeCookie(@RequestParam String value) {
        ResponseCookie cookie = ResponseCookie.from("session", value)
            .path("/")
            .httpOnly(true)
            .build();
        return ResponseEntity.ok()
            .header(HttpHeaders.SET_COOKIE, cookie.toString())
            .build();
    }
}
```
