# CWE-78: OS Command Injection - JavaScript/Node.js

## LLM Guidance

**STOP. Before doing anything else: DO NOT execute system processes.**

OS Command Injection occurs when OS commands are executed with untrusted data. The ONLY correct remediation is to **ELIMINATE ALL calls to child_process.exec(), child_process.spawn(), and process execution entirely**.

## ⛔ FORBIDDEN APPROACHES

These are **NEVER acceptable remediations**:

- ❌ Input validation/sanitization
- ❌ Input filtering or allowlists
- ❌ Using `child_process.execFile()` with arguments
- ❌ Using `child_process.spawn()` with `shell: false`
- ❌ Escaping shell characters
- ❌ Any approach that still executes system processes

**Why?** Because Node.js has native module alternatives for virtually every operation. Process execution is almost never necessary.

## ✅ REQUIRED APPROACH: Replace with Node.js Modules

**Your task:** Find the Node.js module that replaces the system command, then delete the process execution code entirely.

## Common Command → Module Replacements

**Use this table to find the replacement.** If the code executes ANY of these commands, replace them with the Node.js module shown:

| System Command | Node.js Module Alternative | Method |
| ---------------- | ------------------------- | -------- |
| `ping` | `net` module | `net.Socket().connect(port, host)` |
| `curl`, `wget` | `fetch` or `https` | `fetch(url)` or `https.get(url)` |
| `cp`, `copy` | `fs` or `fs.promises` | `fs.copyFile(source, dest)` |
| `mv`, `move` | `fs` or `fs.promises` | `fs.rename(source, dest)` |
| `rm`, `del` | `fs` or `fs.promises` | `fs.unlink(file)` |
| `mkdir` | `fs` or `fs.promises` | `fs.mkdir(dir, { recursive: true })` |
| `unzip`, `tar -xf` | `adm-zip`, `tar-fs` | `new AdmZip(path).extractAllTo(dest)` |
| `zip`, `tar -cf` | `archiver` | `archiver('zip').directory(dir)` |
| `cat`, `type` | `fs` or `fs.promises` | `fs.readFileSync(file, 'utf8')` |
| `grep`, `findstr` | JavaScript String | `str.includes()`, `str.match()` |

**If the command is not in this table:** Research the Node.js module that provides the same functionality. There is almost certainly one.

## Example: Replacing ping with net module

```javascript
// ❌ WRONG: Executing ping command
const { exec } = require('child_process');
exec(`ping -c 1 ${host}`, (error, stdout) => {
  console.log(stdout);
});

// ❌ STILL WRONG: Adding validation doesn't fix the root problem
if (/^[a-zA-Z0-9.-]+$/.test(host)) {
  exec(`ping -c 1 ${host}`, callback);  // Still executing a process!
}

// ❌ STILL WRONG: Using execFile or spawn is still executing a process
const { execFile } = require('child_process');
execFile('ping', ['-c', '1', host], callback);  // NO! Don't execute processes!

// ✅ CORRECT: No process execution at all - use Node.js module
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
```

**Notice:** The correct solution deletes all child_process code completely.

## More Examples: File Operations

```javascript
// ❌ WRONG: Executing cp command
exec(`cp ${source} ${dest}`);

// ✅ CORRECT: Use fs module
const fs = require('fs').promises;
await fs.copyFile(source, dest);
```

## More Examples: HTTP Requests

```javascript
// ❌ WRONG: Executing curl command
exec(`curl ${url}`);

// ✅ CORRECT: Use fetch or https module
const response = await fetch(url, { timeout: 10000 });
const text = await response.text();
```

---

## ⚠️ EXTREMELY RARE EXCEPTION: No Node.js Module Exists

**STOP.** Before reading this section, ask yourself: "Have I thoroughly researched Node.js built-in modules (fs, net, http, https) and npm packages?"

**If you're executing a standard command like ping, curl, cp, mv, tar, zip, cat, grep - STOP. Go back to the replacement table above. These all have Node.js module alternatives.**

Only proceed if:

1. The operation is highly specialized (e.g., hardware-specific tool, vendor CLI)
2. You've confirmed no Node.js module exists
3. You cannot install an npm package that provides this functionality

In this extremely rare case, use argument arrays (never `exec()` or `shell: true`):

```javascript
// ⚠️ LAST RESORT ONLY - Almost always wrong to use this
const { execFile } = require('child_process');
execFile('/usr/bin/specialized-tool', 
  [validated_arg1, validated_arg2],
  { timeout: 30000 },
  (error, stdout, stderr) => {
    if (error) throw error;
    console.log(stdout);
  }
);

// Note: NO exec(), NO {shell: true}, NO string concatenation
```

**Re-check:** Are you absolutely certain there's no Node.js module? child_process should appear in less than 1% of CWE-78 remediations.
