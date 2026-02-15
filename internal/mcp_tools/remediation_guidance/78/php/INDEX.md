# CWE-78: OS Command Injection - PHP

## LLM Guidance

OS Command Injection occurs when untrusted data is incorporated into operating system commands without proper validation, allowing attackers to execute arbitrary commands on the host. In PHP, eliminate exec(), system(), shell_exec(), passthru(), and backtick calls entirely by using native PHP functions (copy(), rename(), file_get_contents(), cURL functions) for file operations and HTTP requests.

## Key Principles

- Replace all exec(), system(), shell_exec(), passthru(), and backtick calls with PHP built-in function alternatives
- Use copy(), rename(), unlink(), mkdir() for file operations instead of system commands
- Use cURL functions or file_get_contents() for HTTP requests instead of curl/wget commands
- Use fsockopen() for network checks instead of ping command
- Never use escapeshellarg() or escapeshellcmd() as a primary defence - they are insufficient
- Never concatenate user input into command strings
- Only use proc_open() as a last resort with bypass_shell option and argument arrays

## Remediation Steps

- Locate command execution - Identify all exec(), system(), shell_exec(), passthru(), backtick, and proc_open() instances
- Determine the operation's purpose - Understand what the command is trying to accomplish
- Find the PHP function alternative - Use copy/rename for file ops, cURL for HTTP, fsockopen for network
- Replace process execution - Delete exec()/system()/shell_exec() code and use the appropriate PHP function
- For unavoidable commands - Use proc_open() with argument array and bypass_shell option, validate all inputs
- Test thoroughly - Verify the PHP function replacement provides the same functionality

## Safe Pattern

```php
// UNSAFE: Executing ping command
$output = shell_exec("ping -c 1 " . $host);

// UNSAFE: Even with escapeshellarg()
$output = shell_exec("ping -c 1 " . escapeshellarg($host));

// SAFE: Use fsockopen for reachability check
function isHostReachable($host, $port = 80, $timeout = 5) {
    $socket = @fsockopen($host, $port, $errno, $errstr, $timeout);
    if ($socket) {
        fclose($socket);
        return true;
    }
    error_log("Host unreachable: $errstr ($errno)");
    return false;
}

// SAFE: File copy with built-in function
copy($source, $dest);

// SAFE: HTTP request with cURL
$ch = curl_init($url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_TIMEOUT, 10);
$response = curl_exec($ch);
curl_close($ch);

// SAFE: File operations with built-in functions
unlink($file);                          // Delete file
rename($old, $new);                     // Move/rename file
mkdir($dir, 0755, true);                // Create directory
$content = file_get_contents($file);    // Read file
```
