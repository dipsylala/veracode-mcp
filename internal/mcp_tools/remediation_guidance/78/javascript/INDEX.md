# CWE-78: OS Command Injection - JavaScript

## LLM Guidance

OS Command Injection occurs when untrusted input is passed to system command execution functions like `child_process.exec()`, allowing attackers to execute arbitrary commands. The primary fix is to avoid shell invocation entirely by using `execFile()` or `spawn()` with argument arrays, ensuring commands and arguments are separated and cannot be manipulated.

## Remediation Strategy

- Never pass user input directly to shell command interpreters
- Use argument arrays instead of string concatenation for commands
- Prefer `execFile()` or `spawn()` over `exec()` to bypass the shell
- Validate and sanitize all external input against strict allowlists
- Apply principle of least privilege to child processes

## Remediation Steps

- Replace `child_process.exec()` with `execFile()` or `spawn()`
- Pass command arguments as array elements, not concatenated strings
- Implement strict input validation using allowlists for expected values
- Escape special shell characters if shell usage is unavoidable
- Run commands with minimal privileges using appropriate user contexts
- Log all command executions for security monitoring

## Minimal Safe Pattern

```javascript
const { execFile } = require('child_process');

// SAFE: Arguments passed as array, no shell interpretation
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
