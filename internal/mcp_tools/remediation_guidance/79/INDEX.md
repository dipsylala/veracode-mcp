# CWE-79: Cross-Site Scripting (XSS) - LLM Guidance

## LLM Guidance

Cross-Site Scripting (XSS) occurs when untrusted data is included in web pages without proper validation or encoding, allowing attackers to inject malicious scripts that execute in victims' browsers. The vulnerability can appear in various contexts including HTML content, attributes, JavaScript, CSS, or URLs.

## Key Principles

- Never render untrusted input directly into executable browser contexts-ensure data remains data, not code
- Apply context-aware output encoding specific to where data appears (HTML encoding differs from JavaScript/URL encoding)
- Treat all external sources as untrusted: user input, databases, external APIs, cookies, headers
- Use defence-in-depth with Content Security Policy (CSP) as a secondary layer
- Validate input format where possible, but rely on output encoding as primary defence

## Actionable Steps

- Identify sources: Locate all untrusted data entry points (user input, external files, databases, network requests, cookies, headers)
- Trace data flow: Follow transformations from source through the application to output
- Determine output context: Identify where data is rendered (HTML body, attribute, JavaScript, CSS, URL)
- Apply context-specific encoding: Use appropriate encoding functions for each context at every data sink
- Verify encoding presence: Audit code for missing encoding/escaping at rendering points
- Test defences: Validate with XSS payloads to confirm proper protection
