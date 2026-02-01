# CWE-79: Cross-Site Scripting (XSS) - PHP

## LLM Guidance

XSS occurs when untrusted data is rendered in web pages without proper encoding, allowing attackers to inject malicious scripts. Always use `htmlspecialchars()` with `ENT_QUOTES | ENT_HTML5` flags and UTF-8 encoding for HTML output, or leverage framework auto-escaping (Laravel Blade `{{ }}`, Twig `{{ }}`). Apply context-specific encoding for JavaScript, URLs, and CSS contexts.

## Remediation Strategy

- Use output encoding appropriate to the context (HTML, JavaScript, URL, CSS)
- Enable auto-escaping in templating engines by default
- Never trust user input or data from external sources
- Implement Content Security Policy (CSP) headers as defense-in-depth
- Validate and sanitize input at application boundaries

## Remediation Steps

- Replace all unencoded output with `htmlspecialchars($data, ENT_QUOTES | ENT_HTML5, 'UTF-8')`
- Use framework escaping - Laravel `{{ $var }}` instead of `{!! $var !!}`, Twig `{{ var }}` not `{{ var|raw }}`
- For JavaScript contexts, use `json_encode($data, JSON_HEX_TAG | JSON_HEX_AMP)`
- For URLs, apply `urlencode()` or `rawurlencode()` to user data
- Review all instances of `echo`, `print`, and template rendering
- Add CSP header - `Content-Security-Policy - default-src 'self'`

## Minimal Safe Pattern

```php
<?php
// HTML context
$username = $_GET['name'];
echo htmlspecialchars($username, ENT_QUOTES | ENT_HTML5, 'UTF-8');

// Laravel Blade template
// {{ $username }} - auto-escaped

// JavaScript context
echo '<script>var user = ' . json_encode($data, JSON_HEX_TAG | JSON_HEX_AMP) . ';</script>';

// URL context
echo '<a href="profile.php?id=' . urlencode($userId) . '">Profile</a>';
?>
```
