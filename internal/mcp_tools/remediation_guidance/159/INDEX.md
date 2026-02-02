# CWE-159. Improper Handling of Invalid Use of Special Elements

## LLM Guidance

Improper handling of invalid special elements occurs when applications fail to properly validate, encode, or reject special characters (metacharacters) that have meaning in specific contexts (SQL, shell, HTML, regex, file paths), enabling injection attacks.

## Key Principles

- Canonicalize and validate special characters/encodings before processing
- Never assume downstream code handles special elements safely
- Apply context-specific encoding for each target context (SQL, shell, HTML, XML, regex, file paths)
- Validate input against allowlists rather than blocklists of special characters
- Reject or encode metacharacters at input boundaries before they reach sensitive operations

## Remediation Steps

- Identify vulnerable code. Locate where untrusted data with special characters reaches sensitive operations (queries, commands, templates)
- Determine target context. Identify whether data flows to SQL, shell, HTML, XML, regex, file path, or other interpreters
- Trace data flow. Map how untrusted input (user data, files, network requests) flows from source to sink
- List dangerous metacharacters. Document which special characters are dangerous in the specific context
- Apply context-specific encoding. Use parameterized queries (SQL), command arrays (shell), HTML entity encoding, or context-appropriate escaping
- Validate and canonicalize. Normalize encodings and validate against allowlists before use
