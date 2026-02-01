# CWE-377: Insecure Temporary File - JavaScript/Node.js

## LLM Guidance

Insecure temporary file creation in Node.js occurs when applications create files with predictable names, insecure permissions, or without proper cleanup in shared directories. Attackers can exploit race conditions, overwrite files, or access sensitive data. Use secure libraries like `tmp`, `temp`, or `fs.mkdtemp()` with proper permissions and automatic cleanup.

## Remediation Strategy

- Use dedicated temporary file libraries (`tmp`, `temp`) or Node.js built-in `fs.mkdtemp()` with secure random naming
- Set restrictive file permissions (0600) to prevent unauthorized access
- Ensure automatic cleanup of temporary files using library callbacks or process exit handlers
- Never create temporary files in world-writable directories with predictable names
- Validate and sanitize any user input used in temporary file operations

## Remediation Steps

- Replace manual `fs.writeFile()` in `/tmp` with `tmp.file()` or `fs.mkdtemp()` for secure random names
- Configure file permissions to owner-only (mode `0600`) when creating temporary files
- Use library cleanup callbacks or `tmp.setGracefulCleanup()` to remove files on exit
- Avoid race conditions by using exclusive file creation flags (`wx` mode)
- Store sensitive temporary data in user-specific directories or system secure temp locations
- Implement error handling to ensure cleanup occurs even when operations fail

## Minimal Safe Pattern

```javascript
const tmp = require('tmp');

// Create secure temporary file with automatic cleanup
tmp.file({ mode: 0600, prefix: 'myapp-', postfix: '.tmp' }, (err, path, fd, cleanupCallback) => {
  if (err) throw err;
  
  // Use the file
  fs.writeSync(fd, 'sensitive data');
  
  // Cleanup when done
  cleanupCallback();
});

// Enable automatic cleanup on process exit
tmp.setGracefulCleanup();
```
