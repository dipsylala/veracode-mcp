# CWE-601: Open Redirect - PHP

## LLM Guidance

Open redirect vulnerabilities occur when user-controlled input is used in `header("Location: ...")`, `<meta>` refresh tags, or JavaScript redirects without validation, enabling phishing and credential theft. The core fix is to validate redirect destinations using allowlists for external URLs or ensuring local redirects use relative paths starting with `/` but not `//`.

## Key Principles

- Prefer allowlist validation for external URLs against known safe domains
- For internal redirects, validate paths are relative (start with `/` not `//`) or match expected patterns
- Use framework-provided redirect methods that include built-in protections
- Never directly insert user input into `Location` headers or redirect mechanisms
- Implement URL parsing to verify scheme, host, and path components before redirecting

## Remediation Steps

- Identify all redirect points using `header()`, framework redirect methods, or client-side redirects
- Replace direct user input with allowlist validation for external URLs
- For local redirects, verify paths start with `/` and don't contain `//` or absolute URLs
- Use `parse_url()` to extract and validate URL components before redirecting
- Apply framework security features (e.g., Laravel's `$request->validate()` with URL rules)
- Add unit tests verifying malicious redirect attempts are blocked

## Safe Pattern

```php
function safeRedirect($userUrl, $allowedDomains = ['example.com']) {
    $parsed = parse_url($userUrl);
    
    // Allow relative paths only (local redirects)
    if (!isset($parsed['host']) && isset($parsed['path']) && 
        $parsed['path'][0] === '/' && substr($parsed['path'], 0, 2) !== '//') {
        header("Location: " . $parsed['path']);
        exit;
    }
    
    // For absolute URLs, validate against allowlist
    if (isset($parsed['host']) && in_array($parsed['host'], $allowedDomains, true)) {
        header("Location: " . $userUrl);
        exit;
    }
    
    // Default safe redirect
    header("Location: /");
    exit;
}
```
