# CWE-78: OS Command Injection - JavaScript

## LLM Guidance

OS Command Injection occurs when untrusted input is passed to system command execution functions like `child_process.exec()`, allowing attackers to execute arbitrary commands.

**Primary Defence:** Use Node.js native modules (fs, https, stream, path, etc.) instead of executing system commands to eliminate the vulnerability entirely. If system commands are unavoidable, use `execFile()` or `spawn()` with argument arrays to avoid shell invocation.

## Key Principles

- **BEST:** Use Node.js native modules (fs, https, path, zlib) instead of system commands to eliminate command injection risk
- **If commands unavoidable:** Use `execFile()` or `spawn()` with argument arrays to bypass the shell
- Never pass user input directly to shell command interpreters
- Use argument arrays instead of string concatenation for commands
- Validate and sanitize all external input against strict allowlists as defence-in-depth
- Apply principle of least privilege to child processes

## Remediation Steps

- Identify all `child_process.exec()`, `child_process.spawn()`, and `child_process.execFile()` calls
- **Replace system commands with Node.js native modules** (fs for files, https for requests, etc.) to eliminate vulnerability
- If commands are unavoidable, replace `child_process.exec()` with `execFile()` or `spawn()`
- Pass command arguments as array elements, not concatenated strings
- Implement strict input validation using allowlists as additional defence
- Run commands with minimal privileges using appropriate user contexts
- Log all command executions for security monitoring

## Safe Pattern

```javascript
const fs = require('fs');
const https = require('https');

// BEST: Use Node.js native modules
function processFile(userFilename) {
  const allowedChars = /^[a-zA-Z0-9_.-]+$/;
  if (!allowedChars.test(userFilename)) {
    throw new Error('Invalid filename');
  }
  const data = fs.readFileSync(userFilename, 'utf8');
  // Process with native APIs
}

// If unavoidable: Arguments passed as array, no shell interpretation
const { execFile } = require('child_process');
function convertFile(userFilename) {
  const allowedChars = /^[a-zA-Z0-9_.-]+$/;
  if (!allowedChars.test(userFilename)) {
    throw new Error('Invalid filename');
  }
  execFile('convert', [userFilename, 'output.png'], (error, stdout) => {
    if (error) throw error;
    console.log(stdout);
  });
}
```
