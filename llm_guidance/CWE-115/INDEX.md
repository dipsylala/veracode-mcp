# CWE-115: Misinterpretation of Input

## LLM Guidance

Misinterpretation of input occurs when different system layers (web server, application framework, backend) process the same data with conflicting parsers or validation logic, allowing attackers to craft payloads that bypass security controls in one layer while being maliciously interpreted by another. The core fix: ensure all layers interpret input identically through consistent parsing, explicit encoding, and validation that matches actual interpretation semantics.

## Key Principles

- Use consistent parsing logic across all application layers
- Explicitly declare character encoding (UTF-8) and reject ambiguous encodings
- Normalize/canonicalize input before performing security validation
- Enforce strict validation that matches actual parser behavior
- Handle edge cases (multi-line headers, duplicate headers, URL encoding) consistently

## Remediation Steps

- Declare UTF-8 encoding explicitly at all layers and reject ambiguous encodings
- Use the same URL parser for both validation and routing operations
- Implement canonicalization before security checks to ensure consistent interpretation
- Handle HTTP headers consistently (multi-line, folding, duplicates) across components
- Enforce Content-Type validation that matches actual parser behavior
- Test with ambiguous inputs (mixed encodings, parser differentials) to verify consistent handling
