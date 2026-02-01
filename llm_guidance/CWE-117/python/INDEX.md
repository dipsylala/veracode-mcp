# CWE-117: Log Injection - Python

## LLM Guidance

Log injection occurs when untrusted data is written to logs without proper encoding, allowing attackers to inject newline characters to forge entries, CRLF sequences to split entries, or escape sequences to manipulate output. Use structured logging with JSON output (python-json-logger or structlog) to automatically encode control characters within fields, preventing log forging while preserving evidence.

## Key Principles

- Structured logging automatically handles encoding and prevents injection
- Never concatenate user input directly into log messages
- Encode control characters (`\n`, `\r`, ANSI escapes) if manual logging required
- Parameterize log messages using `%` or `{}` formatting
- Treat all external input (user data, headers, API responses) as untrusted

## Remediation Steps

- Install structured logging library - `pip install python-json-logger` or `structlog`
- Configure logger to use JSON formatter for all handlers
- Replace string concatenation with structured fields - `logger.info("Login", extra={"user" - username})`
- Use parameterized messages with format strings, not f-strings in log calls
- For legacy systems, strip `\r\n` and ANSI codes from input before logging
- Validate log output contains no attacker-controlled newlines

## Safe Pattern

```python
import logging
from pythonjsonlogger import jsonlogger

logger = logging.getLogger()
handler = logging.StreamHandler()
handler.setFormatter(jsonlogger.JsonFormatter())
logger.addHandler(handler)

# Safe: user input automatically encoded in JSON field
username = request.get("username")  # May contain \n or \r
logger.info("User login", extra={"username": username, "ip": request.remote_addr})
# Output: {"message":"User login","username":"admin\nINJECTED","ip":"1.2.3.4"}
```
