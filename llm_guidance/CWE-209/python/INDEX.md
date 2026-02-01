# CWE-209: Error Message Information Leak - Python

## LLM Guidance

Error Message Information Leak occurs when Python applications expose sensitive details through exception tracebacks, debug output, or verbose error messages in responses. These leaks reveal file paths, code structure, library versions, SQL queries, and internal logic to attackers. Return generic error messages to users while logging detailed exceptions securely server-side.

## Remediation Strategy

- Separate user-facing and internal error handling: Show generic messages to clients, log full details securely
- Disable debug mode in production: Set `DEBUG=False` in Django/Flask and disable verbose tracebacks
- Sanitize all error responses: Never expose stack traces, file paths, or internal state in API/web responses
- Use structured logging: Log exceptions with context to secure locations inaccessible to users

## Remediation Steps

- Configure production settings to disable debug mode and detailed error pages
- Implement custom exception handlers that return generic HTTP error responses
- Add centralized logging for all exceptions with full traceback details
- Review all try-except blocks to ensure user-facing messages are generic
- Implement error monitoring with tools that capture exceptions server-side
- Test error scenarios to verify no sensitive information leaks through responses

## Minimal Safe Pattern

```python
import logging

logger = logging.getLogger(__name__)

def process_request(data):
    try:
        result = perform_operation(data)
        return {"status": "success", "data": result}
    except Exception as e:
        # Log full details server-side
        logger.error(f"Operation failed: {e}", exc_info=True)
        # Return generic message to user
        return {"status": "error", "message": "An error occurred processing your request"}, 500
```
