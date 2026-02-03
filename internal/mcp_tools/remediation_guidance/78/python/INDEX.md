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

## Safe Pattern

```python
# ❌ VULNERABLE: System command with user input
import subprocess
output = subprocess.check_output(f"ping -c 1 {host}", shell=True)

# ❌ STILL BAD: Validation doesn't fix the root problem
if re.match(r'^[a-zA-Z0-9.-]+$', host):  # Insufficient protection!
    output = subprocess.check_output(f"ping -c 1 {host}", shell=True)

# ✅ CORRECT: Use Python socket library instead of ping command
import socket

def is_host_reachable(host, port=80, timeout=5):
    """Test network connectivity using socket, not ping command"""
    try:
        socket.create_connection((host, port), timeout=timeout)
        return True
    except (socket.timeout, socket.error) as e:
        print(f"Host unreachable: {e}")
        return False

# Alternative: ICMP ping using third-party library (no subprocess needed)
# pip install ping3
from ping3 import ping
response_time = ping(host, timeout=5)
is_reachable = response_time is not None

# ✅ OTHER LIBRARY-BASED SOLUTIONS:

# File operations: Use pathlib/shutil, not cp/mv commands
from pathlib import Path
import shutil

source = Path(user_path).resolve()
dest = Path("/backup/data.txt")
shutil.copy2(source, dest)

# HTTP requests: Use requests, not curl/wget
import requests
response = requests.get(validated_url, timeout=10)

# Archive extraction: Use zipfile, not unzip command
import zipfile
with zipfile.ZipFile(validated_zip_path, 'r') as zip_ref:
    zip_ref.extractall('/extract/')

# Archive creation: Use tarfile, not tar command  
import tarfile
with tarfile.open('/backup/archive.tar.gz', 'w:gz') as tar:
    tar.add(validated_directory)

# ⚠️ LAST RESORT: If no library alternative exists
# Use subprocess with list args and shell=False
result = subprocess.run(
    ['/usr/bin/convert', validated_input, validated_output],  # argument list
    shell=False,  # NEVER use shell=True
    capture_output=True,
    text=True,
    timeout=30
)
```
