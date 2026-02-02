# CWE-77: Command Injection - PHP

## LLM Guidance

Command injection in PHP occurs when applications construct system commands using untrusted input through functions like `system()`, `exec()`, `shell_exec()`, or backticks. The primary defense is eliminating shell execution entirely by using PHP built-in functions or libraries for specific tasks. When shell execution is unavoidable, properly escape all user input with `escapeshellarg()` for arguments and `escapeshellcmd()` for entire commands.

## Key Principles

- Eliminate shell execution functions entirely; use native PHP functions for file operations, network requests, and system tasks
- Implement strict input validation with whitelists of allowed values before any processing
- Apply proper escaping with `escapeshellarg()` for all user-supplied arguments when shell execution is unavoidable
- Enforce least privilege by running PHP processes with minimal system permissions
- Avoid dynamic command construction through string concatenation or interpolation

## Remediation Steps

- Audit code for all instances of `system()`, `exec()`, `shell_exec()`, backticks, `passthru()`, `proc_open()`, and `popen()`
- Replace shell commands with PHP alternatives (e.g., `file_get_contents()` instead of `cat`, `unlink()` instead of `rm`)
- For unavoidable shell usage, wrap all user input with `escapeshellarg()` before passing to command functions
- Implement whitelist validation for any parameters that determine command behavior
- Remove or restrict user control over command structure, file paths, and executable names
- Configure PHP with `disable_functions` in php.ini to block dangerous functions in production

## Safe Pattern

```php
// UNSAFE: Direct user input in command
$file = $_GET['file'];
$output = shell_exec("cat $file");

// SAFE: Use PHP built-in instead
$file = $_GET['file'];
$allowedFiles = ['report.txt', 'data.csv'];
if (in_array($file, $allowedFiles, true)) {
    $output = file_get_contents($file);
}

// SAFE: If shell required, escape properly
$userArg = escapeshellarg($_GET['input']);
$output = shell_exec("grep search_term " . $userArg);
```
