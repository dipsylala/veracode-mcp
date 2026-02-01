# CWE-434: Unrestricted File Upload - LLM Guidance

## LLM Guidance

Unrestricted file upload vulnerabilities occur when applications accept file uploads without properly validating file type, content, size, or destination. Attackers exploit this to upload malicious files including web shells (PHP, JSP, ASPX), executable malware, HTML/SVG files with XSS payloads, or oversized files causing denial of service. This is particularly dangerous when uploaded files are stored within the webroot and can be directly accessed or executed.

## Key Principles

- Validate file type through content inspection (magic bytes/MIME type), never trust file extensions
- Store uploads outside the webroot to prevent direct execution
- Rename uploaded files using random identifiers to prevent path traversal and direct access
- Never execute or serve uploaded files in their original format without sanitization
- Implement strict file size limits to prevent denial of service attacks

## Remediation Steps

- Use libraries like `python-magic` or `libmagic` to validate actual file content via magic bytes, not extensions
- Configure upload directory outside the webroot (e.g., `/var/app_data/uploads` instead of `/var/www/html/uploads`)
- Generate random filenames (UUIDs) and maintain a database mapping to prevent predictable file paths
- Enforce maximum file size limits appropriate to your use case
- Serve files through download handlers that set proper `Content-Type` and `Content-Disposition` headers
- For images, re-encode through trusted libraries to strip potential malicious payloads before serving
