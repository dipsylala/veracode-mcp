# CWE-80: Cross-Site Scripting (XSS) - PHP

## LLM Guidance

XSS occurs when untrusted data is included in web output without proper encoding, allowing attackers to inject malicious scripts. The core fix is context-aware output encoding using `htmlspecialchars()` with `ENT_QUOTES | ENT_HTML5` flags for HTML contexts, `json_encode()` for JavaScript contexts, or framework auto-escaping features.

## Key Principles

- Apply output encoding at the point of rendering, not at input
- Use context-appropriate encoding (HTML, JavaScript, URL, CSS)
- Leverage framework auto-escaping (Laravel Blade `{{ }}`, Twig `{{ }}`)
- Set Content-Security-Policy headers to limit script execution
- Never trust user input, even if previously validated

## Remediation Steps

- Identify all user-controlled data rendered in responses
- Replace raw output (`<?= ?>`, `{!! !!}`) with encoded output
- Use `htmlspecialchars($data, ENT_QUOTES | ENT_HTML5, 'UTF-8')` for HTML
- Use `json_encode($data, JSON_HEX_TAG | JSON_HEX_AMP)` for JavaScript
- Enable framework auto-escaping globally (Blade, Twig default behavior)
- Validate the fix by testing with payloads like `<script>alert(1)</script>`

## Safe Pattern

```php
// HTML context - safe encoding
<?php
$userName = $_GET['name'];
?>
<h1>Welcome, <?= htmlspecialchars($userName, ENT_QUOTES | ENT_HTML5, 'UTF-8') ?></h1>

// JavaScript context - safe encoding
<script>
var userName = <?= json_encode($userName, JSON_HEX_TAG | JSON_HEX_AMP) ?>;
</script>

// Laravel Blade - auto-escaped
<h1>Welcome, {{ $userName }}</h1>
```
