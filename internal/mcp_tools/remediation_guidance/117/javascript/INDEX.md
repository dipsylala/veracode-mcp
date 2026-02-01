# CWE-117: Log Injection - JavaScript/Node.js

## LLM Guidance

Log Injection in JavaScript/Node.js occurs when untrusted user input is written to logs without sanitization, allowing attackers to forge log entries, inject newlines to create fake events, or manipulate log analysis tools. Node.js applications using winston, bunyan, pino, or console.log are vulnerable when logging user-controlled data containing control characters (ASCII 0x00-0x1F, DEL 0x7F, C1 controls 0x80-0x9F, Unicode line separators U+0085/U+2028/U+2029), or ANSI escape codes. Fix by encoding all control characters before logging or using structured JSON logging.

## Key Principles

- Use structured JSON logging (winston, bunyan, pino) which automatically encodes control characters within field values
- Encode ALL control characters if using plain text logging: ASCII controls (0x00-0x1F), DEL (0x7F), C1 controls (0x80-0x9F), Unicode line separators (U+0085, U+2028, U+2029), and ANSI escape sequences
- Avoid string concatenation/interpolation when building log messages with user-controlled data
- Configure logging frameworks to automatically escape or sanitize user input fields
- Validate and sanitize all user-controlled data at application boundaries before it reaches logging code

## Remediation Steps

- Identify all locations where user input (params, headers, body) is logged
- Switch to structured JSON logging (winston/bunyan/pino) with separate fields for user data instead of string interpolation
- For plain text logging, implement comprehensive control character encoding (see Manual Encoding Pattern below)
- Configure logger settings to auto-sanitize or use libraries like `validator` to clean inputs
- Add input validation at entry points to reject or sanitize data with control characters
- Test by attempting to inject newlines, null bytes, ANSI codes, and Unicode line separators - verify proper encoding

## Safe Pattern (Structured Logging)

```javascript
// Use winston/bunyan/pino with JSON formatting
const winston = require('winston');
const logger = winston.createLogger({
  format: winston.format.json(),  // Automatically encodes control chars
  transports: [new winston.transports.Console()]
});

// User input is automatically encoded in JSON fields
logger.info({
  event: 'user_login',
  username: req.body.username,  // Control chars encoded automatically
  ip: req.ip
});
// Output: {"event":"user_login","username":"admin\\nfake_entry","ip":"127.0.0.1"}
```

## Safe Pattern (Legacy/Plain Text Logging)

If structured logging is unavailable, use comprehensive control character encoding:

```javascript
function encodeForLog(input) {
  if (typeof input !== 'string') return String(input);
  
  let output = '';
  for (let i = 0; i < input.length; i++) {
    const code = input.charCodeAt(i);
    const ch = input[i];
    
    // Encode ASCII control chars (0x00-0x1F) + DEL (0x7F) + C1 controls (0x80-0x9F)
    if (code <= 0x1F || code === 0x7F || (code >= 0x80 && code <= 0x9F)) {
      switch (ch) {
        case '\\': output += '\\\\'; break;
        case '\r': output += '\\r'; break;
        case '\n': output += '\\n'; break;
        case '\t': output += '\\t'; break;
        default: output += `\\u${code.toString(16).padStart(4, '0')}`; break;
      }
      continue;
    }
    
    // Encode Unicode line separators
    if (code === 0x0085 || code === 0x2028 || code === 0x2029) {
      output += `\\u${code.toString(16).padStart(4, '0')}`;
      continue;
    }
    
    output += ch;
  }
  return output;
}

// Usage
console.log(`Login attempt: ${encodeForLog(username)}`);
logger.info(`User ${encodeForLog(userId)} performed action: ${encodeForLog(action)}`);
```
