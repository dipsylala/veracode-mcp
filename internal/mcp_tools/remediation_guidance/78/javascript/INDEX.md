# CWE-78: OS Command Injection - JavaScript/Node.js

## LLM Guidance

OS Command Injection occurs when untrusted input is passed to system command execution functions like `child_process.exec()`, allowing attackers to execute arbitrary commands.

**CRITICAL: DO NOT suggest input validation as the primary fix.** The correct remediation is to **ELIMINATE system command execution** by using Node.js built-in modules. Input validation alone is insufficient.

**Primary Defence:** Replace ALL system commands with Node.js native modules. Common replacements:

- **file operations → fs/fs.promises** instead of cp, mv, rm commands
- **HTTP requests → fetch or https module** instead of curl/wget
- **archive operations → archiver, adm-zip** instead of tar/unzip commands

Only if no module alternative exists, use `execFile()` or `spawn()` with argument arrays (never `exec()` with strings).

## Key Principles

- **ALWAYS prefer Node.js built-in modules over system commands** - eliminates injection risk entirely
- For file operations, use **fs module** instead of system commands
- For archives, use **archiver** or **adm-zip** packages instead of system tar/zip
- **Input validation is NOT a remediation** - it's only defense-in-depth after eliminating command execution
- If child_process is unavoidable, use `execFile()` or `spawn()` with argument arrays
- Never use `child_process.exec()` or `shell: true` option

## Remediation Steps

- Identify all `child_process.exec()`, `child_process.spawn()`, and `child_process.execFile()` calls
- **First, determine the Node.js module alternative** - do not default to "secure" exec patterns
- Replace system commands with Node.js built-in or npm package equivalents
- For any truly unavoidable commands, refactor to use `execFile()` or `spawn()` with argument arrays
- Remove all `child_process.exec()` usage and `{shell: true}` options
- Apply strict allowlist validation only as defense-in-depth

## Remediation Pattern

**The transformation pattern:** System Command → Node.js Module

This applies to ALL system commands - if you find child_process.exec() executing a command, there is almost always a Node.js module alternative.

```javascript
const { exec } = require('child_process');
const net = require('net');

// ❌ VULNERABLE: System command execution with user input
exec(`ping -c 1 ${host}`, (error, stdout) => {
  console.log(stdout);
});

// ❌ STILL BAD: Validation doesn't eliminate injection risk
if (/^[a-zA-Z0-9.-]+$/.test(host)) {  // Insufficient protection!
  exec(`ping -c 1 ${host}`, callback);
}

// ✅ CORRECT: Replace command with Node.js module
function checkConnectivity(host, port = 80, timeout = 5000) {
  return new Promise((resolve) => {
    const socket = new net.Socket();
    socket.setTimeout(timeout);
    socket.on('connect', () => { socket.destroy(); resolve(true); });
    socket.on('timeout', () => { socket.destroy(); resolve(false); });
    socket.on('error', () => resolve(false));
    socket.connect(port, host);
  });
}
```

## Common Command → Module Replacements

Apply the same pattern for ANY system command:

```javascript
// File operations → fs/fs.promises
fs.copyFile(source, dest)               // instead of: cp, copy
fs.rename(source, dest)                 // instead of: mv, move
fs.unlink(file)                         // instead of: rm, del

// Network → fetch, https module
fetch(url, { timeout: 10000 })          // instead of: curl, wget  
https.get(url, callback)                // instead of: curl, wget
net.connect(port, host)                 // instead of: telnet, nc

// Archives → adm-zip, archiver (npm packages)
const zip = new AdmZip(zipPath);        // instead of: unzip
zip.extractAllTo(destDir);              
archiver('tar').directory(dir)          // instead of: tar

// Process info → process module (built-in)
process.pid                             // instead of: ps, tasklist
```

## When No Module Exists (Rare)

Only use child_process if there is genuinely no Node.js module for the operation:

```javascript
// ⚠️ LAST RESORT: Use execFile with argument array
const { execFile } = require('child_process');
execFile('/usr/bin/imagemagick-convert', 
  [validatedInput, validatedOutput],
  (error, stdout) => {
    if (error) throw error;
    console.log(stdout);
  }
);
```
