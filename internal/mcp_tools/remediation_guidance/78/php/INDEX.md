# CWE-78: OS Command Injection - PHP

## LLM Guidance

OS Command Injection occurs when applications execute OS commands with untrusted data without proper sanitization, allowing attackers to run arbitrary commands on the host system.

**CRITICAL: DO NOT suggest escapeshellarg() or validation as the primary fix.** The correct remediation is to **ELIMINATE system command execution** by using PHP built-in functions. Shell escaping alone is insufficient.

**Primary Defence:** Replace ALL system commands with PHP native functions. Common replacements:

- **file operations → file functions (copy, rename, unlink)** instead of cp, mv, rm
- **HTTP requests → cURL or file_get_contents()** instead of wget/curl commands
- **archive operations → ZipArchive, PharData** instead of tar/unzip commands

Only if no PHP function alternative exists, use `proc_open()` with argument arrays and `bypass_shell => true`.

## Key Principles

- **ALWAYS prefer PHP built-in functions over system commands** - eliminates injection risk entirely
- For file operations, use **PHP file functions** instead of system commands
- For archives, use **ZipArchive** or **PharData** classes instead of system commands
- **escapeshellarg() is NOT a remediation** - it's only defense-in-depth after eliminating exec()
- If system calls are unavoidable, use `proc_open()` with argument arrays and `bypass_shell => true`
- Never use exec(), system(), shell_exec(), passthru(), or backtick operators

## Remediation Steps

- Identify all exec(), system(), shell_exec(), passthru(), backtick (``), and popen() calls
- **First, determine the PHP built-in alternative** - do not default to "secure" exec patterns
- Replace system commands with PHP native function equivalents
- For any truly unavoidable commands, refactor to use `proc_open()` with argument arrays and `bypass_shell => true`
- Remove all exec(), system(), shell_exec() usage
- Disable dangerous functions in php.ini (disable_functions directive)

## Remediation Pattern

**The transformation pattern:** System Command → PHP Built-in Function

This applies to ALL system commands - if you find exec(), shell_exec(), or backticks executing a command, there is almost always a PHP function alternative.

```php
// ❌ VULNERABLE: System command execution with user input
$output = shell_exec("ping -c 1 " . $host);

// ❌ STILL BAD: escapeshellarg() doesn't eliminate injection risk
$output = shell_exec("ping -c 1 " . escapeshellarg($host));  // Still vulnerable!

// ✅ CORRECT: Replace command with PHP function
function checkConnectivity($host, $port = 80, $timeout = 5) {
    $socket = @fsockopen($host, $port, $errno, $errstr, $timeout);
    if ($socket) {
        fclose($socket);
        return true;
    }
    error_log("Host unreachable: $errstr");
    return false;
}
```

## Common Command → Function Replacements

Apply the same pattern for ANY system command:

```php
// File operations → PHP file functions
copy($source, $dest);                   // instead of: cp, copy
rename($old, $new);                     // instead of: mv, move
unlink($file);                          // instead of: rm, del

// Network → cURL, file_get_contents, fsockopen
curl_exec(curl_init($url));             // instead of: curl, wget
file_get_contents($url);                // instead of: wget, curl
fsockopen($host, $port);                // instead of: telnet, nc

// Archives → ZipArchive, PharData
$zip = new ZipArchive();                // instead of: unzip, tar -xf
$zip->open($path);
$zip->extractTo($dest);
$phar = new PharData($path);            // instead of: tar
$phar->extractTo($dest);
```

## When No Function Exists (Rare)

Only use proc_open() if there is genuinely no PHP function for the operation:

```php
// ⚠️ LAST RESORT: Use argument array with bypass_shell
$process = proc_open(
    ['/usr/bin/imagemagick-convert', $validatedInput, $validatedOutput],
    $descriptors,
    $pipes,
    null, null,
    ['bypass_shell' => true]  // CRITICAL: bypasses shell
);
```
