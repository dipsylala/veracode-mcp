# CWE-117: Log Injection - JavaScript/Node.js

## LLM Guidance

Log Injection in JavaScript/Node.js occurs when untrusted user input is written to logs without sanitization, allowing attackers to forge log entries, inject newlines to create fake events, or manipulate log analysis tools. Node.js applications using winston, bunyan, pino, or console.log are vulnerable when logging user-controlled data containing control characters like `\n`, `\r`, or ANSI escape codes. Fix by replacing or encoding newlines and control characters before logging.

## Remediation Strategy

- Replace newline characters (`\n`, `\r`) and control characters with safe representations before logging
- Use structured logging (JSON format) instead of string concatenation to prevent injection
- Configure logging frameworks to automatically escape or sanitize user input fields
- Validate and sanitize all user-controlled data at application boundaries before it reaches logging code
- Implement log output encoding that escapes special characters in user-provided values

## Remediation Steps

- Identify all locations where user input (params, headers, body) is logged
- Replace control characters - `userInput.replace(/[\n\r]/g, '_')` or use encodeURIComponent
- Switch to structured logging with separate fields for user data instead of string interpolation
- Configure logger settings to auto-sanitize or use libraries like `validator` to clean inputs
- Add input validation at entry points to reject or sanitize data with control characters
- Review log outputs to verify no user input can create fake log entries

## Minimal Safe Pattern

```javascript
const sanitizeForLog = (input) => {
  if (typeof input !== 'string') return input;
  return input.replace(/[\n\r\t]/g, '_');
};

// Structured logging with sanitized fields
logger.info({
  event: 'user_login',
  username: sanitizeForLog(req.body.username),
  ip: sanitizeForLog(req.ip)
});

// Or with string interpolation
console.log(`Login attempt: ${sanitizeForLog(username)}`);
```
