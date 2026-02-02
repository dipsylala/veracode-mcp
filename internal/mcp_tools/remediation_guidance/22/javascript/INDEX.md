# CWE-22: Path Traversal - JavaScript/Node.js

## LLM Guidance

Path Traversal in JavaScript/Node.js occurs when applications use unsanitized user input to construct file paths, allowing attackers to access files outside intended directories using sequences like `../`. The primary defense is indirect reference mapping (mapping user IDs to files) rather than accepting direct file paths. When direct paths are necessary, validate against an allowlist and resolve paths to ensure they remain within the intended directory.

## Key Principles

- Use indirect reference mapping with IDs/tokens instead of accepting file paths from users
- Validate all path inputs against strict allowlists of permitted files/directories
- Resolve and normalize paths with `path.resolve()` and verify they start with the expected base directory
- Reject inputs containing path traversal sequences (`../`, `..\\`, encoded variants)
- Apply principle of least privilege to file system permissions

## Remediation Steps

- Replace direct file path parameters with indirect references (database IDs, UUIDs)
- Use `path.resolve()` to normalize user input and base directory to absolute paths
- Verify resolved path starts with intended base directory using `path.relative()` or `startsWith()`
- Implement allowlist validation for permitted file extensions and names
- Sanitize input by rejecting `..`, null bytes, and encoded traversal attempts
- Configure Express static middleware with `dotfiles - 'deny'` and strict root directories

## Safe Pattern

```javascript
const path = require('path');
const fs = require('fs');

const BASE_DIR = path.resolve('./uploads');

function safeReadFile(userFilename) {
  const requestedPath = path.resolve(BASE_DIR, userFilename);
  
  if (!requestedPath.startsWith(BASE_DIR + path.sep)) {
    throw new Error('Access denied');
  }
  
  return fs.readFileSync(requestedPath);
}
```
