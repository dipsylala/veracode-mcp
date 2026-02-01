# CWE-114: Process Control - Python

## LLM Guidance

Process control vulnerabilities occur when untrusted input influences process creation, termination, or management operations. Attackers can spawn malicious processes, kill critical services, or exhaust system resources. Always validate process identifiers, use allowlists for process operations, and avoid passing user input directly to process control functions.

## Remediation Strategy

- Validate all process identifiers (PIDs) against an allowlist of expected/authorized processes
- Never pass unsanitized user input to subprocess calls, signal operations, or process management functions
- Use least-privilege principles: restrict process control operations to specific authorized users/roles
- Implement resource limits and monitoring to detect abnormal process spawning or termination patterns
- Prefer built-in APIs over shell execution; avoid shell=True in subprocess calls

## Remediation Steps

- Identify all code paths where user input affects process control (subprocess, os.kill, signal operations)
- Implement strict allowlists mapping user inputs to predefined safe process operations
- Replace dynamic process calls with parameterized safe alternatives using validated inputs
- Add authorization checks before any process control operation
- Set resource limits using setrlimit() to prevent process exhaustion attacks
- Log and monitor all process control operations for security auditing

## Minimal Safe Pattern

```python
import subprocess

ALLOWED_COMMANDS = {
    'backup': ['/usr/bin/backup.sh', '--safe'],
    'report': ['/usr/bin/generate_report.py']
}

def execute_task(task_name):
    if task_name not in ALLOWED_COMMANDS:
        raise ValueError("Unauthorized task")
    
    result = subprocess.run(
        ALLOWED_COMMANDS[task_name],
        capture_output=True,
        timeout=30,
        shell=False
    )
    return result.stdout
```
