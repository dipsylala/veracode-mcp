# CWE-117: Log Injection / Log Forging - Java

## LLM Guidance

Log Injection occurs when untrusted data is written to log files without encoding, allowing attackers to forge log entries or inject malicious content. The core fix is to use structured logging (JSON/ECS format) with SLF4J/Logback or Log4j2, which automatically encodes control characters within field values, preventing log forgery. Parameterized logging alone is insufficient without structured output formats.

## Remediation Strategy

- Use structured logging formats (JSON, ECS) that encode control characters automatically within field boundaries
- Employ parameterized logging with SLF4J `{}` placeholders to separate untrusted data from log messages
- Sanitize newlines (`\n`, `\r`) and control characters from user input before logging if structured logging is unavailable
- Avoid string concatenation when building log messages with user-controlled data
- Configure logstash-logback-encoder or Log4j2 JsonLayout for production environments

## Remediation Steps

- Add logstash-logback-encoder or log4j-layout-template-json dependency to your project
- Configure Logback/Log4j2 to use JSON encoder/layout in logback.xml or log4j2.xml
- Replace string concatenation in log statements with parameterized logging using `{}`
- Pass user input as parameters, not concatenated into the message string
- For legacy systems without JSON support, strip `\n`, `\r`, and control characters using `replaceAll("[\\r\\n]", "")`
- Test by attempting to inject newlines into logged fields and verify they appear encoded in output

## Minimal Safe Pattern

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
