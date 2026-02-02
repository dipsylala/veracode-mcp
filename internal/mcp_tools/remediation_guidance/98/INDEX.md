# CWE-98: PHP Remote File Inclusion

## LLM Guidance

Remote File Inclusion (RFI) in PHP occurs when untrusted input is used in file inclusion functions (`include`, `require`, `include_once`, `require_once`) without proper validation, allowing attackers to execute arbitrary code from remote sources. The core fix is to never allow untrusted input to select files for inclusion-use allowlists and disable remote file inclusion entirely.

## Key Principles

- Never use untrusted input directly in `include`/`require` functions
- Disable `allow_url_include` and `allow_url_fopen` in php.ini
- Map user input to predefined allowlists, not file paths
- Validate that resolved paths stay within expected directories
- Prefer autoloading over dynamic file inclusion

## Remediation Steps

- Trace data flow - Identify where untrusted data (HTTP params, cookies, external APIs) reaches file inclusion functions
- Implement allowlists - Map user input to predefined file paths using arrays or switch statements
- Disable remote inclusion - Set `allow_url_include=0` and `allow_url_fopen=0` in php.ini
- Validate paths - Use `realpath()` to resolve paths and verify they're within allowed directories
- Remove dynamic inclusion - Replace variable-based includes with explicit imports or autoloading
- Sanitize input - If dynamic inclusion is unavoidable, strip path separators and validate against strict patterns
