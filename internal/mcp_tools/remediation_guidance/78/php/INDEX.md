# CWE-78: OS Command Injection - PHP

## LLM Guidance

**STOP. Before doing anything else: DO NOT execute system processes.**

OS Command Injection occurs when OS commands are executed with untrusted data. The ONLY correct remediation is to **ELIMINATE ALL calls to exec(), system(), shell_exec(), passthru(), backticks, and process execution entirely**.

## ⛔ FORBIDDEN APPROACHES

These are **NEVER acceptable remediations**:

- ❌ `escapeshellarg()` or `escapeshellcmd()`
- ❌ Input validation/sanitization
- ❌ Input filtering or allowlists
- ❌ Using `proc_open()` with `bypass_shell`
- ❌ Escaping shell characters
- ❌ Any approach that still executes system processes

**Why?** Because PHP has built-in function alternatives for virtually every operation. Process execution is almost never necessary.

## ✅ REQUIRED APPROACH: Replace with PHP Built-in Functions

**Your task:** Find the PHP function that replaces the system command, then delete the process execution code entirely.

## Common Command → Function Replacements

**Use this table to find the replacement.** If the code executes ANY of these commands, replace them with the PHP function shown:

| System Command | PHP Function Alternative | Method |
| ---------------- | ------------------------- | -------- |
| `ping` | `fsockopen()` | `fsockopen($host, $port, $errno, $errstr, $timeout)` |
| `curl`, `wget` | `curl_*()` or `file_get_contents()` | `curl_exec(curl_init($url))` |
| `cp`, `copy` | `copy()` | `copy($source, $dest)` |
| `mv`, `move` | `rename()` | `rename($old, $new)` |
| `rm`, `del` | `unlink()` | `unlink($file)` |
| `mkdir` | `mkdir()` | `mkdir($dir, 0755, true)` |
| `unzip`, `tar -xf` | `ZipArchive`, `PharData` | `$zip->extractTo($dest)` |
| `zip`, `tar -cf` | `ZipArchive` | `$zip->addFile($file)` |
| `cat`, `type` | `file_get_contents()` | `file_get_contents($file)` |
| `grep`, `findstr` | PHP string functions | `str_contains()`, `preg_match()` |

**If the command is not in this table:** Research the PHP function that provides the same functionality. There is almost certainly one.

## Example: Replacing ping with fsockopen

```php
// ❌ WRONG: Executing ping command
$output = shell_exec("ping -c 1 " . $host);

// ❌ STILL WRONG: Adding escapeshellarg() doesn't fix the root problem
$output = shell_exec("ping -c 1 " . escapeshellarg($host));  // Still executing a process!

// ❌ STILL WRONG: Using proc_open with bypass_shell is still executing a process
$process = proc_open(['ping', '-c', '1', $host], $descriptors, $pipes);  // NO! Don't execute processes!

// ✅ CORRECT: No process execution at all - use PHP function
function isHostReachable($host, $port = 80, $timeout = 5) {
    $socket = @fsockopen($host, $port, $errno, $errstr, $timeout);
    if ($socket) {
        fclose($socket);
        return true;
    }
    error_log("Host unreachable: $errstr ($errno)");
    return false;
}
```

**Notice:** The correct solution deletes all exec(), shell_exec(), system(), and backtick code completely.

## More Examples: File Operations

```php
// ❌ WRONG: Executing cp command
exec("cp $source $dest");

// ✅ CORRECT: Use copy() function
copy($source, $dest);
```

## More Examples: HTTP Requests

```php
// ❌ WRONG: Executing curl command
$output = shell_exec("curl " . escapeshellarg($url));

// ✅ CORRECT: Use cURL functions
$ch = curl_init($url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_TIMEOUT, 10);
$response = curl_exec($ch);
curl_close($ch);
```

---

## ⚠️ EXTREMELY RARE EXCEPTION: No PHP Function Exists

**STOP.** Before reading this section, ask yourself: "Have I thoroughly researched PHP built-in functions, cURL, file functions, and PEAR/Composer packages?"

**If you're executing a standard command like ping, curl, cp, mv, tar, zip, cat, grep - STOP. Go back to the replacement table above. These all have PHP function alternatives.**

Only proceed if:

1. The operation is highly specialized (e.g., hardware-specific tool, vendor CLI)
2. You've confirmed no PHP function exists
3. You cannot install a Composer package that provides this functionality

In this extremely rare case, use argument arrays with `bypass_shell`:

```php
// ⚠️ LAST RESORT ONLY - Almost always wrong to use this
$descriptorspec = [
    0 => ['pipe', 'r'],
    1 => ['pipe', 'w'],
    2 => ['pipe', 'w']
];

$process = proc_open(
    ['/usr/bin/specialized-tool', $validated_arg1, $validated_arg2],
    $descriptorspec,
    $pipes,
    null,
    null,
    ['bypass_shell' => true]  // CRITICAL: bypasses shell
);

if (is_resource($process)) {
    fclose($pipes[0]);
    $output = stream_get_contents($pipes[1]);
    fclose($pipes[1]);
    proc_close($process);
}

// Note: NO exec(), NO shell_exec(), NO system(), NO backticks, NO passthru()
```

**Re-check:** Are you absolutely certain there's no PHP function? proc_open() should appear in less than 1% of CWE-78 remediations.

## Additional Security

After eliminating all system command execution:

1. **Disable dangerous functions in php.ini:**

   ```ini
   disable_functions = exec,passthru,shell_exec,system,proc_open,popen,curl_exec,curl_multi_exec,parse_ini_file,show_source
   ```

2. **Use strict type checking and validation on all inputs** (defense-in-depth only)
