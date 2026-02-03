# CWE-114: Process Control - JavaScript/Node.js

## LLM Guidance

Process control vulnerabilities in JavaScript/Node.js applications occur when untrusted user input controls process execution, lifecycle, or behavior. Node.js's `child_process` module and process management capabilities make these vulnerabilities particularly dangerous as attackers can spawn malicious processes, terminate critical services, or exhaust system resources. Always validate and sanitize input before using it in process-related operations, and use allowlists to restrict which processes can be controlled.

Key Security Issues:

- User input directly controls `child_process.spawn()`, `exec()`, or `fork()` parameters
- Unsanitized input used in process arguments, environment variables, or working directories
- Allowing process termination via user-controlled PID values
- Command injection through shell metacharacters in process execution

## Key Principles

- Use strict allowlists for executable paths and process arguments
- Disable shell interpretation by using `spawn()` with array arguments instead of `exec()`
- Validate all user input against expected formats before process operations
- Implement least privilege principles for process execution permissions
- Use safe APIs like `execFile()` that bypass shell interpretation

## Remediation Steps

- Replace `exec()` and `execSync()` with `spawn()` or `execFile()` to avoid shell interpretation
- Create allowlists of permitted executables and validate against them before execution
- Sanitize and validate all user inputs used in process arguments or environment variables
- Use argument arrays instead of concatenating strings for command execution
- Implement timeout and resource limits for spawned processes
- Log all process execution attempts with security monitoring

## Safe Pattern

```javascript
const { execFile } = require('child_process');
const path = require('path');

// Allowlist of permitted tools
const ALLOWED_TOOLS = { 'imagemagick': '/usr/bin/convert', 'ffmpeg': '/usr/bin/ffmpeg' };

function processFile(toolName, userFile) {
  const executablePath = ALLOWED_TOOLS[toolName];
  if (!executablePath) throw new Error('Invalid tool');
  
  // Validate filename against pattern
  if (!/^[a-zA-Z0-9_-]+\.(jpg|png)$/.test(userFile)) throw new Error('Invalid file');
  
  // Use execFile with argument array - no shell interpretation
  execFile(executablePath, ['-resize', '800x600', userFile], { timeout: 5000 }, (err, stdout) => {
    if (err) console.error('Process failed:', err);
  });
}
```
