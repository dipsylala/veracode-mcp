# CWE-78: OS Command Injection - Python

## LLM Guidance

**STOP. Before doing anything else: DO NOT execute system processes.**

OS Command Injection occurs when OS commands are executed with untrusted data. The ONLY correct remediation is to **ELIMINATE ALL calls to subprocess, os.system(), os.popen(), and process execution entirely**.

## ⛔ FORBIDDEN APPROACHES

These are **NEVER acceptable remediations**:

- ❌ Input validation/sanitization
- ❌ Input filtering or allowlists
- ❌ Using `subprocess.run()` with `shell=False`
- ❌ Escaping shell characters
- ❌ Any approach that still executes system processes

**Why?** Because Python has native library alternatives for virtually every operation. Process execution is almost never necessary.

## ✅ REQUIRED APPROACH: Replace with Python Libraries

**Your task:** Find the Python library that replaces the system command, then delete the process execution code entirely.

## Common Command → Library Replacements

**Use this table to find the replacement.** If the code executes ANY of these commands, replace them with the Python library shown:

| System Command | Python Library Alternative | Method |
| ---------------- | ------------------------- | -------- |
| `ping` | `socket` module | `socket.create_connection((host, port), timeout)` |
| `curl`, `wget` | `requests` or `urllib` | `requests.get(url, timeout=10)` |
| `cp`, `copy` | `shutil` module | `shutil.copy2(source, dest)` |
| `mv`, `move` | `shutil` module | `shutil.move(source, dest)` |
| `rm`, `del` | `pathlib` module | `Path(file).unlink()` |
| `mkdir` | `pathlib` module | `Path(dir).mkdir(parents=True)` |
| `unzip`, `tar -xf` | `zipfile`, `tarfile` | `zipfile.ZipFile(path).extractall()` |
| `zip`, `tar -cf` | `zipfile`, `tarfile` | `zipfile.ZipFile(path, 'w').write()` |
| `ps`, `tasklist` | `psutil` (3rd party) | `psutil.Process().name()` |
| `cat`, `type` | `pathlib` module | `Path(file).read_text()` |
| `grep`, `findstr` | Python string/regex | `'text' in str`, `re.search()` |

**If the command is not in this table:** Research the Python library that provides the same functionality. There is almost certainly one.

## Example: Replacing ping with socket

```python
# ❌ WRONG: Executing ping command
import subprocess
output = subprocess.check_output(f"ping -c 1 {host}", shell=True)

# ❌ STILL WRONG: Adding validation doesn't fix the root problem
if re.match(r'^[a-zA-Z0-9.-]+$', host):
    output = subprocess.check_output(f"ping -c 1 {host}", shell=True)  # Still executing a process!

# ❌ STILL WRONG: Using subprocess.run with shell=False is still executing a process
result = subprocess.run(["ping", "-c", "1", host], capture_output=True)  # NO! Don't execute processes!

# ✅ CORRECT: No process execution at all - use Python library
import socket

def is_host_reachable(host, port=80, timeout=5):
    try:
        socket.create_connection((host, port), timeout=timeout)
        return True
    except (socket.timeout, socket.error) as e:
        print(f"Host unreachable: {e}")
        return False
```

**Notice:** The correct solution deletes all subprocess, os.system(), and os.popen() code completely.

## More Examples: File Operations

```python
# ❌ WRONG: Executing cp command
os.system(f"cp {source} {dest}")

# ✅ CORRECT: Use shutil module
import shutil
shutil.copy2(source, dest)
```

## More Examples: HTTP Requests

```python
# ❌ WRONG: Executing curl command
subprocess.run(["curl", url])

# ✅ CORRECT: Use requests library
import requests
response = requests.get(url, timeout=10)
content = response.text
```

---

## ⚠️ EXTREMELY RARE EXCEPTION: No Python Library Exists

**STOP.** Before reading this section, ask yourself: "Have I thoroughly researched Python standard library, requests, pathlib, shutil, zipfile, tarfile, and psutil?"

**If you're executing a standard command like ping, curl, cp, mv, tar, zip, ps, cat, grep - STOP. Go back to the replacement table above. These all have Python library alternatives.**

Only proceed if:

1. The operation is highly specialized (e.g., hardware-specific tool, vendor CLI)
2. You've confirmed no Python library exists
3. You cannot install a Python library that provides this functionality

In this extremely rare case, use argument lists with `shell=False`:

```python
# ⚠️ LAST RESORT ONLY - Almost always wrong to use this
result = subprocess.run(
    ['/usr/bin/specialized-tool', validated_arg1, validated_arg2],
    shell=False,  # NEVER shell=True
    capture_output=True,
    timeout=30
)

# Note: NO shell invocation, NO string concatenation, NO f-strings with commands
```

**Re-check:** Are you absolutely certain there's no Python library? subprocess should appear in less than 1% of CWE-78 remediations.
