# CWE-91: XML Injection - LLM Guidance

## LLM Guidance

XML Injection occurs when untrusted user input is incorporated into XML documents without proper validation or escaping, allowing attackers to modify XML structure or content. The core fix is to never construct XML through string concatenation; instead, use XML libraries that automatically escape user input as data.

## Key Principles

- Never concatenate untrusted input directly into XML strings
- Use XML libraries with built-in escaping and serialization
- Treat all user input as data, not XML structure
- Validate and sanitize input before XML processing
- Implement allowlist validation for XML element/attribute names

## Remediation Steps

- Trace the data flow - Identify where untrusted data enters (source) and where XML is constructed (sink)
- Replace string concatenation - Use safe XML APIs (e.g., `DocumentBuilder`, `XMLStreamWriter`, `lxml.etree`) that treat input as data
- Escape special characters - If concatenation is unavoidable, escape `<`, `>`, `&`, `"`, `'` using library functions
- Validate input format - Apply strict allowlist validation on data that determines XML structure
- Use parameterized APIs - Leverage DOM methods like `createElement()` and `createTextNode()` to safely build XML
- Test edge cases - Verify fix with malicious payloads containing XML metacharacters
