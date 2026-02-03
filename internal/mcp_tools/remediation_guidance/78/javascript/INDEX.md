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

## Safe Pattern

```javascript
const { exec } = require('child_process');
const net = require('net');
const dns = require('dns').promises;

// ❌ VULNERABLE: System command with user input
exec(`ping -c 1 ${host}`, (error, stdout) => {
  console.log(stdout);
});

// ❌ STILL BAD: Validation doesn't fix the root problem
if (/^[a-zA-Z0-9.-]+$/.test(host)) {  // Insufficient protection!
  exec(`ping -c 1 ${host}`, callback);
}

// ✅ CORRECT: Use Node.js net module instead of ping command
function isHostReachable(host, port = 80, timeout = 5000) {
  return new Promise((resolve) => {
    const socket = new net.Socket();
    
    socket.setTimeout(timeout);
    socket.on('connect', () => {
      socket.destroy();
      resolve(true);
    });
    socket.on('timeout', () => {
      socket.destroy();
      resolve(false);
    });
    socket.on('error', () => {
      resolve(false);
    });
    
    socket.connect(port, host);
  });
}

// Alternative: DNS resolution test
async function canResolveHost(host) {
  try {
    await dns.lookup(host);
    return true;
  } catch (error) {
    console.error(`Cannot resolve ${host}:`, error.message);
    return false;
  }
}

// ✅ OTHER MODULE-BASED SOLUTIONS:

// File operations: Use fs module, not cp/mv commands
const fs = require('fs').promises;
await fs.copyFile(userSourcePath, '/backup/data.txt');

// HTTP requests: Use fetch or https, not curl command
const response = await fetch(validatedUrl, { timeout: 10000 });
const data = await response.text();

// Alternative HTTP with https module
const https = require('https');
https.get(validatedUrl, (res) => {
  let data = '';
  res.on('data', chunk => data += chunk);
  res.on('end', () => console.log(data));
});

// Archive extraction: Use adm-zip package, not unzip command
// npm install adm-zip
const AdmZip = require('adm-zip');
const zip = new AdmZip(validatedZipPath);
zip.extractAllTo('/extract/', true);

// Archive creation: Use archiver package, not tar command
// npm install archiver
const archiver = require('archiver');
const output = fs.createWriteStream('/backup/archive.tar.gz');
const archive = archiver('tar', { gzip: true });
archive.pipe(output);
archive.directory(validatedDirectory, false);
await archive.finalize();

// ⚠️ LAST RESORT: If no module alternative exists
// Use execFile with argument array (never exec with strings)
const { execFile } = require('child_process');
execFile('/usr/bin/convert', [validatedInput, validatedOutput], (error, stdout) => {
  if (error) throw error;
  console.log(stdout);
});
```
