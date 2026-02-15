# CWE-78: OS Command Injection - Python

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands on the host. In Python, eliminate subprocess, os.system(), and os.popen() calls entirely by using native Python libraries (pathlib, shutil, requests, socket) for file operations, HTTP requests, and network operations.

## Key Principles

- Replace all subprocess, os.system(), and os.popen() calls with Python standard library alternatives
- Use pathlib and shutil for file operations (copy, move, delete) instead of system commands
- Use requests or urllib for HTTP requests instead of curl/wget
- Use socket for network checks instead of ping commands
- Never concatenate user input into command strings
- Never use shell=True - it enables shell injection
- Only use subprocess as a last resort with argument lists and shell=False

## Remediation Steps

- Locate command execution - Identify all subprocess, os.system(), os.popen() instances
- Determine the operation's purpose - Understand what the command is trying to accomplish
- Find the Python library alternative - Use pathlib/shutil for file ops, requests for HTTP, socket for network
- Replace process execution - Delete subprocess/os.system code and use the appropriate Python library
- For unavoidable commands - Use subprocess.run() with argument list and shell=False, validate all inputs
- Test thoroughly - Verify the Python library replacement provides the same functionality

## Safe Pattern

```python
# UNSAFE: Executing ping command
import subprocess
output = subprocess.check_output(f"ping -c 1 {host}", shell=True)

# UNSAFE: Even with shell=False
result = subprocess.run(["ping", "-c", "1", host], capture_output=True)

# SAFE: Use socket for reachability check
import socket

def is_host_reachable(host, port=80, timeout=5):
    try:
        socket.create_connection((host, port), timeout=timeout)
        return True
    except (socket.timeout, socket.error) as e:
        print(f"Host unreachable: {e}")
        return False

# SAFE: File copy with shutil
import shutil
shutil.copy2(source, dest)

# SAFE: HTTP request with requests
import requests
response = requests.get(url, timeout=10)
content = response.text

# SAFE: File operations with pathlib
from pathlib import Path
Path(file).unlink()  # Delete file
Path(directory).mkdir(parents=True, exist_ok=True)  # Create directory
content = Path(file).read_text()  # Read file
```
