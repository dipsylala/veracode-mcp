# CWE-114: Process Control

## LLM Guidance

Process control vulnerabilities occur when applications accept untrusted input to control process execution, library loading, or behavior (start, stop, kill, priority, resource limits). This includes insecure dynamic library loading (DLL hijacking, LD_PRELOAD attacks), command injection through process execution, and unauthorized process control. Attackers exploit weak validation to load malicious libraries, execute arbitrary commands, or manipulate running processes.

## Key Principles

Only load components from trusted, integrity-checked locations and never allow user input to directly control process operations or library search paths.

- Hardcode all library paths and process execution commands; never construct them from user input
- Validate process identifiers against an allowlist before performing control operations
- Use absolute paths for all library loading and disable dynamic search path manipulation
- Implement strict input validation and sanitization before any process control operation
- Run processes with least privilege and enforce OS-level security policies

## Remediation Steps

- Identify the vulnerability. Locate where untrusted data controls process operations (`os.kill()`, `dlopen()`, `LoadLibrary()`, process spawning)
- Replace user-controlled process parameters with hardcoded values or validated allowlists
- Use absolute paths for libraries and set secure search paths (`DLL safe search mode`, `RPATH` restrictions)
- If process control is required, validate PIDs/names against authorized processes owned by the application
- Apply least privilege. Ensure processes run with minimal permissions and cannot manipulate critical system processes

```python
# VULNERABLE - User controls which process to kill
pid = int(request.GET['pid'])
os.kill(pid, signal.SIGTERM)

# SECURE - Only allow killing application-owned processes
ALLOWED_PROCESSES = {'worker' - worker_pid, 'job' - job_pid}
process_name = request.GET['process']
if process_name in ALLOWED_PROCESSES -
    os.kill(ALLOWED_PROCESSES[process_name], signal.SIGTERM)
```
