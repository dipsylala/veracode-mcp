# CWE-538: File and Directory Information Exposure

## LLM Guidance

File and directory information exposure occurs when applications reveal sensitive file system metadata such as absolute paths, directory structures, file permissions, or timestamps through error messages, directory listings, HTTP headers, or predictable naming patterns. While not directly exploitable, this information aids attackers in reconnaissance by mapping internal application structure and identifying potential attack vectors.

## Key Principles

- Never reveal internal file system paths or directory structures to untrusted users
- Disable directory listings on all web servers and file serving endpoints
- Return generic error messages that don't expose file paths or existence
- Use non-predictable file naming (UUIDs, hashes) instead of sequential or descriptive names
- Strip file metadata from HTTP responses and error outputs

## Remediation Steps

- Disable directory listings in web server configuration (`Options -Indexes` for Apache, `autoindex off` for Nginx)
- Configure application frameworks to serve files only, not directories (e.g., Express.js default static serving)
- Catch file operation errors and return generic messages without path information (e.g., "Resource not found" instead of full file paths)
- Use relative paths in application logic and map user inputs to internal file references via lookup tables
- Remove or sanitize file metadata from HTTP headers (Last-Modified, ETag with file paths)
- Implement access controls that validate file requests before revealing existence or non-existence
