# CWE-93: CRLF Injection - Java

## LLM Guidance

CRLF Injection occurs when attackers inject `\r\n` characters to manipulate HTTP headers, log files, or line-based formats, potentially enabling HTTP response splitting or log forgery. The core fix is to strip or reject newline characters (`\r`, `\n`, `\r\n`) from all user input before using it in HTTP headers or logs. Use Spring Framework's `ResponseEntity` and `HttpHeaders` which provide built-in validation, or implement strict allowlist validation for header values.

## Key Principles

- Input Sanitization: Remove or reject all CR/LF characters from user-controlled data before using in headers or logs
- Framework Protection: Leverage Spring's `HttpHeaders` and `ResponseEntity` which validate header values automatically
- Structured Logging: Use SLF4J with parameterized logging instead of string concatenation to prevent log injection
- Allowlist Validation: Validate header values against strict patterns (alphanumeric, specific safe characters only)
- Avoid Direct Manipulation: Never directly construct HTTP headers or log entries from untrusted input

## Remediation Steps

- Identify all locations where user input flows into HTTP headers or log statements
- Replace direct header manipulation with Spring's `HttpHeaders` API
- Sanitize input by removing `\r` and `\n` characters - `value.replaceAll("[\\r\\n]", "")`
- Implement regex validation for expected header value formats before use
- Convert log statements to use SLF4J parameterized logging - `log.info("User - {}", username)`
- Test with CRLF payloads (`%0d%0a`, `\r\n`) to verify protection

## Safe Pattern

```java
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;

public ResponseEntity<String> setCustomHeader(String userInput) {
    // Sanitize input - remove CRLF characters
    String sanitized = userInput.replaceAll("[\\r\\n]", "");
    
    // Use Spring's HttpHeaders for built-in validation
    HttpHeaders headers = new HttpHeaders();
    headers.add("X-Custom-Header", sanitized);
    
    return ResponseEntity.ok()
        .headers(headers)
        .body("Response");
}
```
