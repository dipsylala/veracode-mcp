# CWE-89: SQL Injection

## LLM Guidance

SQL Injection occurs when untrusted data is incorporated into SQL queries without proper sanitization, allowing attackers to manipulate query logic, access unauthorized data, or execute administrative operations. The core fix is to use parameterized queries (prepared statements) so user input is always treated as data, not query structure.

## Key Principles

- Never build SQL by concatenating untrusted input directly into queries
- Use parameterized queries/prepared statements as the primary defence
- Treat all user input as untrusted data, not executable SQL code
- Apply defence-in-depth: combine parameterized queries with input validation and least privilege
- Avoid dynamic query construction; use static SQL with parameters

## Remediation Steps

- Trace the data path - Identify the source (user input, external data) and sink (SQL execution function like `.execute()`, `.query()`)
- Locate string concatenation - Find instances of `+`, `concat()`, `format()`, or template literals building SQL with untrusted data
- Replace with parameterized queries - Convert concatenated SQL to prepared statements with placeholders (`?`, `$1`, `-param`)
- Bind parameters separately - Pass untrusted data as separate parameters to the query execution function
- Validate input - Add input validation as a secondary defence layer (whitelist allowed values, validate types/formats)
- Test thoroughly - Verify the fix prevents injection by testing with malicious payloads (e.g., `' OR '1'='1`)
