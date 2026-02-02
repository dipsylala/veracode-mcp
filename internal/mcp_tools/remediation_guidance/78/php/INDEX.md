# CWE-78: OS Command Injection - PHP

## LLM Guidance

OS Command Injection occurs when applications execute OS commands with untrusted data without proper sanitization, allowing attackers to run arbitrary commands on the host system.

**Primary Defence:** Use PHP native functions (scandir, cURL, ZipArchive) instead of system commands. If system calls are unavoidable, use `proc_open()` with argument arrays and `bypass_shell => true`.

## Key Principles

- Replace system calls entirely: Use PHP native functions instead of exec()/system()/shell_exec() to eliminate the vulnerability
- Use proc_open() with bypass_shell: If system calls are required, pass arguments as arrays with `bypass_shell => true` option
- Apply defence-in-depth: Implement allowlist validation and `escapeshellarg()` as additional layers, never as primary defence
- Harden environment: Disable dangerous functions in php.ini and apply least privilege permissions

## Remediation Steps

- Identify all exec(), system(), shell_exec(), passthru(), and backtick calls in codebase
- Replace with PHP native alternatives (file operations, cURL, ZipArchive, etc.)
- For unavoidable cases, refactor to `proc_open()` with argument arrays and bypass_shell option
- Add allowlist validation for user inputs that influence commands
- Apply `escapeshellarg()` to dynamic arguments as supplementary protection
- Test fixes with command injection payloads (e.g., `; ls`, `| whoami`)

## Safe Pattern

```php
// SAFE: proc_open with bypass_shell and argument array
$descriptors = [
    0 => ['pipe', 'r'],
    1 => ['pipe', 'w'],
    2 => ['pipe', 'w']
];

$process = proc_open(
    ['/usr/bin/command', $userInput],  // Array prevents injection
    $descriptors,
    $pipes,
    null,
    null,
    ['bypass_shell' => true]  // Critical: bypasses shell interpretation
);
```
