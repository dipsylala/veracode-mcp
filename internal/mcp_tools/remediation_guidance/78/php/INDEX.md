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

## Safe Pattern

```php
// ❌ VULNERABLE: System command with user input
$output = shell_exec("ping -c 1 " . $host);

// ❌ STILL BAD: escapeshellarg() doesn't fix the root problem
$output = shell_exec("ping -c 1 " . escapeshellarg($host));  // Still vulnerable!

// ✅ CORRECT: Use PHP socket functions instead of ping command
function isHostReachable($host, $port = 80, $timeout = 5) {
    $errno = 0;
    $errstr = '';
    
    // Suppress warnings and handle errors explicitly
    $socket = @fsockopen($host, $port, $errno, $errstr, $timeout);
    
    if ($socket) {
        fclose($socket);
        return true;
    }
    
    error_log("Host unreachable: $errstr ($errno)");
    return false;
}

// Alternative: Socket extension for more control
function checkConnection($host, $port = 80, $timeout = 5) {
    $socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
    socket_set_option($socket, SOL_SOCKET, SO_RCVTIMEO, ['sec' => $timeout, 'usec' => 0]);
    
    $result = @socket_connect($socket, $host, $port);
    socket_close($socket);
    
    return $result !== false;
}

// ✅ OTHER FUNCTION-BASED SOLUTIONS:

// File operations: Use PHP functions, not cp/mv commands
copy($userSourcePath, '/backup/data.txt');
rename($oldPath, $newPath);
unlink($filePath);

// HTTP requests: Use cURL, not wget/curl commands  
$ch = curl_init($validatedUrl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_TIMEOUT, 10);
$response = curl_exec($ch);
curl_close($ch);

// Alternative HTTP: file_get_contents with context
$context = stream_context_create([
    'http' => [
        'timeout' => 10,
        'ignore_errors' => true
    ]
]);
$response = file_get_contents($validatedUrl, false, $context);

// Archive extraction: Use ZipArchive, not unzip command
$zip = new ZipArchive();
if ($zip->open($validatedZipPath) === true) {
    $zip->extractTo('/extract/');
    $zip->close();
}

// Archive creation: Use ZipArchive, not tar/zip commands
$zip = new ZipArchive();
if ($zip->open('/backup/archive.zip', ZipArchive::CREATE) === true) {
    $zip->addFile($validatedFilePath, basename($validatedFilePath));
    $zip->close();
}

// Tar archives: Use PharData class
$phar = new PharData('/backup/archive.tar');
$phar->buildFromDirectory($validatedDirectory);
$phar->compress(Phar::GZ);

// ⚠️ LAST RESORT: If no PHP function alternative exists
// Use proc_open with argument array and bypass_shell
$descriptors = [
    0 => ['pipe', 'r'],
    1 => ['pipe', 'w'],
    2 => ['pipe', 'w']
];

$process = proc_open(
    ['/usr/bin/convert', $validatedInput, $validatedOutput],  // Argument array
    $descriptors,
    $pipes,
    null,
    null,
    ['bypass_shell' => true]  // CRITICAL: bypasses shell interpretation
);

if (is_resource($process)) {
    $output = stream_get_contents($pipes[1]);
    fclose($pipes[1]);
    proc_close($process);
}
```
