# CWE-93: CRLF Injection

## LLM Guidance

CRLF Injection occurs when untrusted user input is included in HTTP headers or protocol fields without validation, allowing attackers to inject carriage return (CR) and line feed (LF) characters to manipulate protocol behavior. The core issue is allowing untrusted data to inject CRLF or protocol delimiters into structured text like headers, logs, or protocol messages.

## Key Principles

- Never allow untrusted data to directly inject CRLF sequences or protocol delimiters
- Validate and encode all user input before generating structured text (headers, logs, protocols)
- Trace data flow from untrusted sources to header/protocol construction points
- Prefer framework-provided header APIs over manual string concatenation
- Remove or encode CR (`\r`) and LF (`\n`) characters from all untrusted input

## Remediation Steps

- Identify the vulnerability - Review flaw details for file, line number, and code pattern where untrusted data reaches header construction
- Trace data flow - Map the path from source (user input, database, external file) to sink (header-setting functions like `setHeader()`, `response.headers[]`)
- Locate dangerous patterns - Find string concatenation or interpolation where untrusted data is inserted into header values
- Encode CRLF characters - Strip or encode `\r` and `\n` characters from all untrusted input before use
- Use safe APIs - Replace manual header construction with framework-provided methods that handle encoding automatically
- Validate input - Implement allowlists for expected header values and reject inputs containing control characters
