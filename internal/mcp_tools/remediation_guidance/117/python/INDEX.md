# CWE-117: Log Injection - Python

## LLM Guidance

Log injection occurs when untrusted data is written to logs without proper encoding, allowing attackers to inject newline characters to forge entries, CRLF sequences to split entries, or escape sequences to manipulate output. Use structured logging with JSON output (python-json-logger or structlog) to automatically encode control characters within fields, preventing log forging while preserving evidence.

## Key Principles

- Structured logging automatically handles encoding and prevents injection
- Never concatenate user input directly into log messages
- Encode ALL control characters if manual logging required: ASCII controls (0x00-0x1F), DEL (0x7F), C1 controls (0x80-0x9F), Unicode line separators (U+0085, U+2028, U+2029), and ANSI escape sequences
- Parameterize log messages using `%` or `{}` formatting
- Treat all external input (user data, headers, API responses) as untrusted

## Remediation Steps

- Install structured logging library - `pip install python-json-logger` or `structlog`
- Configure logger to use JSON formatter for all handlers
- Replace string concatenation with structured fields - `logger.info("Login", extra={"user": username})`
- Use parameterized messages with format strings, not f-strings in log calls
- For legacy systems without JSON logging, implement comprehensive control character encoding (see Manual Encoding Pattern below)
- Test by attempting to inject newlines, null bytes, ANSI codes, and Unicode line separators - verify proper encoding

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
# Output: {"message":"User login","username":"admin\\nINJECTED","ip":"1.2.3.4"}
```

## Safe Pattern (Legacy Systems)

If structured logging is unavailable, use comprehensive control character encoding:

```python
def encode_for_log(input_str):
    """Encode all control characters for safe logging."""
    if not isinstance(input_str, str):
        return str(input_str)
    
    output = []
    for ch in input_str:
        code = ord(ch)
        
        # Encode ASCII control chars (0x00-0x1F) + DEL (0x7F) + C1 controls (0x80-0x9F)
        if code <= 0x1F or code == 0x7F or (0x80 <= code <= 0x9F):
            if ch == '\\':
                output.append('\\\\')
            elif ch == '\r':
                output.append('\\r')
            elif ch == '\n':
                output.append('\\n')
            elif ch == '\t':
                output.append('\\t')
            else:
                output.append(f'\\u{code:04x}')
            continue
        
        # Encode Unicode line separators
        if code in (0x0085, 0x2028, 0x2029):
            output.append(f'\\u{code:04x}')
            continue
        
        output.append(ch)
    
    return ''.join(output)

# Usage
logger.info(f"User login: {encode_for_log(username)}")
logger.warning(f"Suspicious activity from {encode_for_log(user_input)}")
```
