# CWE-78: OS Command Injection - Python

## LLM Guidance

OS Command Injection occurs when applications incorporate untrusted data into OS commands without proper validation, allowing attackers to execute arbitrary commands on the host system.

**CRITICAL: DO NOT suggest input validation as the primary fix.** The correct remediation is to **ELIMINATE system command execution** by using Python's standard library. Input validation alone is insufficient.

**Primary Defence:** Replace ALL system commands with Python native libraries. Common replacements:

- **file operations → pathlib, shutil** instead of cp, mv, rm commands
- **HTTP requests → requests or urllib** instead of curl/wget
- **archive operations → zipfile, tarfile** instead of tar/unzip commands

Only if no library alternative exists, use `subprocess.run()` with argument lists and `shell=False`.

## Key Principles

- **ALWAYS prefer Python standard library over system commands** - eliminates injection risk entirely
- For file operations, use **pathlib.Path** and **shutil** instead of system commands
- For archives, use **zipfile/tarfile** modules instead of system tar/zip commands
- **Input validation is NOT a remediation** - it's only defense-in-depth after eliminating command execution
- If subprocess is unavoidable, use argument lists with `shell=False` (NEVER `shell=True` or string commands)
- Never use `os.system()`, `os.popen()`, or any shell invocation

## Remediation Steps

- Identify all subprocess, os.system(), os.popen(), and exec family calls
- **First, determine the library-based replacement** - do not default to "secure" subprocess usage
- Replace system commands with Python standard library equivalents
- For any truly unavoidable commands, refactor to use `subprocess.run()` with list arguments and `shell=False`
- Remove all `shell=True` usage and string-based command construction
- Apply strict allowlist validation only as defense-in-depth

## Remediation Pattern

**The transformation pattern:** System Command → Python Standard Library

This applies to ALL system commands - if you find subprocess, os.system(), or os.popen() executing a command, there is almost always a Python library alternative.

```python
# ❌ VULNERABLE: System command execution with user input
import subprocess
output = subprocess.check_output(f"ping -c 1 {host}", shell=True)

# ❌ STILL BAD: Validation doesn't eliminate injection risk
if re.match(r'^[a-zA-Z0-9.-]+$', host):  # Insufficient protection!
    output = subprocess.check_output(f"ping -c 1 {host}", shell=True)

# ✅ CORRECT: Replace command with Python library
import socket

def check_connectivity(host, port=80, timeout=5):
    try:
        socket.create_connection((host, port), timeout=timeout)
        return True
    except (socket.timeout, socket.error):
        return False
```

## Common Command → Library Replacements

Apply the same pattern for ANY system command:

```python
# File operations → pathlib, shutil
shutil.copy2(source, dest)              # instead of: cp, copy
shutil.move(source, dest)               # instead of: mv, move  
Path(file).unlink()                     # instead of: rm, del

# Network → requests, urllib, socket
requests.get(url, timeout=10)           # instead of: curl, wget
socket.create_connection((host, port))  # instead of: ping, telnet

# Archives → zipfile, tarfile
zipfile.ZipFile(path).extractall()      # instead of: unzip, tar -xf
tarfile.open(path).extractall()         # instead of: tar -xf
zipfile.ZipFile(path, 'w').write()      # instead of: zip, tar -cf

# Process info → psutil (third-party)
import psutil
psutil.Process(pid).name()              # instead of: ps, tasklist
```

## When No Library Exists (Rare)

Only use subprocess if there is genuinely no Python library for the operation:

```python
# ⚠️ LAST RESORT: Use argument list with shell=False
result = subprocess.run(
    ['/usr/bin/imagemagick-convert', validated_input, validated_output],
    shell=False,  # NEVER shell=True
    capture_output=True,
    timeout=30
)
```
