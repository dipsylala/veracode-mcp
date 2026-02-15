# CWE-78: OS Command Injection - JavaScript/Node.js

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands on the host. In Node.js, eliminate child_process module usage entirely by using native Node.js modules (fs, net, http/https) for file operations, HTTP requests, and network operations.

## Key Principles

- Replace all child_process.exec(), child_process.spawn(), and child_process.execFile() calls with Node.js module alternatives
- Use fs or fs.promises for file operations instead of system commands
- Use fetch, http, or https modules for HTTP requests instead of curl/wget
- Use net module for network checks instead of ping commands
- Never concatenate user input into command strings
- Never use shell: true - it enables shell injection
- Only use child_process as a last resort with argument arrays and shell: false

## Remediation Steps

- Locate command execution - Identify all child_process.exec(), spawn(), execFile() instances
- Determine the operation's purpose - Understand what the command is trying to accomplish
- Find the Node.js module alternative - Use fs for file ops, fetch/https for HTTP, net for network
- Replace process execution - Delete child_process code and use the appropriate Node.js module
- For unavoidable commands - Use execFile() with argument array and no shell option, validate all inputs
- Test thoroughly - Verify the Node.js module replacement provides the same functionality

## Safe Pattern

```javascript
// UNSAFE: Executing ping command
const { exec } = require('child_process');
exec(`ping -c 1 ${host}`, (error, stdout) => {
  console.log(stdout);
});

// UNSAFE: Even with execFile
const { execFile } = require('child_process');
execFile('ping', ['-c', '1', host], callback);

// SAFE: Use net module for reachability check
const net = require('net');

function isHostReachable(host, port = 80, timeout = 5000) {
  return new Promise((resolve) => {
    const socket = new net.Socket();
    socket.setTimeout(timeout);
    socket.on('connect', () => { socket.destroy(); resolve(true); });
    socket.on('timeout', () => { socket.destroy(); resolve(false); });
    socket.on('error', () => { socket.destroy(); resolve(false); });
    socket.connect(port, host);
  });
}

// SAFE: File copy with fs.promises
const fs = require('fs').promises;
await fs.copyFile(source, dest);

// SAFE: HTTP request with fetch
const response = await fetch(url, { timeout: 10000 });
const text = await response.text();

// SAFE: File operations with fs
const fs = require('fs');
await fs.promises.unlink(file);  // Delete file
await fs.promises.mkdir(directory, { recursive: true });  // Create directory
const content = await fs.promises.readFile(file, 'utf8');  // Read file
```
