# CWE-80: Improper Neutralization of Script-Related HTML Tags (Basic XSS)

## LLM Guidance

CWE-80 occurs when applications fail to properly neutralize script-related HTML tags (`<script>`, `<img>`, `<iframe>`) in web output, allowing attackers to inject malicious scripts that execute in victims' browsers. This is a specific subset of CWE-79 focusing on basic HTML tag injection. The core fix is applying context-appropriate output encoding to prevent untrusted input from being interpreted as executable markup.

## Key Principles

- Never include untrusted input in HTML output without context-appropriate encoding
- Ensure data cannot be interpreted as executable markup or script
- Apply encoding based on specific output context (HTML body, attribute, JavaScript, CSS, URL)
- Use context-aware output encoding as the primary defence layer

## Remediation Steps

- Identify all sources of untrusted data (user input, databases, external files, network requests, cookies, headers)
- Trace data transformations from source to output sink
- Locate where data is rendered (response writing, template rendering, DOM manipulation)
- Determine the specific output context where data appears
- Check for missing or inadequate encoding/escaping functions
- Apply context-aware output encoding at all output points
