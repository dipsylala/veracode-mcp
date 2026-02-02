# CWE-117: Log Injection / Log Forging - Java

## LLM Guidance

Log Injection occurs when untrusted data is written to log files without encoding, allowing attackers to forge log entries or inject malicious content. The core fix is to use structured logging (JSON/ECS format) with SLF4J/Logback or Log4j2, which automatically encodes control characters within field values, preventing log forgery. Parameterized logging alone is insufficient without structured output formats.

## Key Principles

- Use structured logging formats (JSON, ECS) that encode control characters automatically within field boundaries
- Employ parameterized logging with SLF4J `{}` placeholders to separate untrusted data from log messages
- If structured logging is unavailable, encode ALL control characters: ASCII controls (0x00-0x1F), DEL (0x7F), C1 controls (0x80-0x9F), and Unicode line separators (U+0085, U+2028, U+2029)
- Avoid string concatenation when building log messages with user-controlled data
- Configure logstash-logback-encoder or Log4j2 JsonLayout for production environments

## Remediation Steps

- Add logstash-logback-encoder or log4j-layout-template-json dependency to your project
- Configure Logback/Log4j2 to use JSON encoder/layout in logback.xml or log4j2.xml
- Replace string concatenation in log statements with parameterized logging using `{}`
- Pass user input as parameters, not concatenated into the message string
- For legacy systems without JSON support, implement comprehensive control character encoding (see Manual Encoding Pattern below)
- Test by attempting to inject newlines, null bytes, and Unicode line separators into logged fields and verify proper encoding

## Safe Pattern

```java
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SecureLogging {
    private static final Logger logger = LoggerFactory.getLogger(SecureLogging.class);
    
    public void logUserAction(String username, String action) {
        // Use parameterized logging with JSON encoder configured
        logger.info("User action performed: username={}, action={}", username, action);
        // JSON output: {"message":"User action performed","username":"admin","action":"login"}
        // Newlines in username/action are automatically encoded as \n
    }
}
```

## Safe Pattern (Legacy Systems)

If structured logging is not available, use comprehensive control character encoding:

```java
public class LogEncoder {
    public static String encodeForLog(String input) {
        if (input == null) return null;
        StringBuilder out = new StringBuilder(input.length() + 16);
        
        for (int i = 0; i < input.length(); i++) {
            char ch = input.charAt(i);
            // Encode full ASCII control range + DEL + C1 controls
            if (ch <= 0x1F || ch == 0x7F || (ch >= 0x80 && ch <= 0x9F)) {
                switch (ch) {
                    case '\\': out.append("\\\\"); break;
                    case '\r': out.append("\\r"); break;
                    case '\n': out.append("\\n"); break;
                    case '\t': out.append("\\t"); break;
                    default: out.append(String.format("\\u%04x", (int) ch)); break;
                }
                continue;
            }
            // Encode Unicode line separators
            if (ch == '\u0085') { out.append("\\u0085"); continue; }
            if (ch == '\u2028') { out.append("\\u2028"); continue; }
            if (ch == '\u2029') { out.append("\\u2029"); continue; }
            out.append(ch);
        }
        return out.toString();
    }
}

// Usage
logger.info("User action: {}", LogEncoder.encodeForLog(userInput));
```
