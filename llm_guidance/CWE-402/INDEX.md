# CWE-402: Transmission of Private Resources

## LLM Guidance

Applications unintentionally expose private files, internal data structures, or restricted resources to unauthorized parties through web responses, API endpoints, or error messages. This includes serving files outside the webroot, exposing backup files (.bak, ~, .swp), returning internal data in API responses, or leaking data through misconfigured proxies. Core fix: restrict to designated public directories with default-deny access controls.

## Key Principles

- Default-deny access; explicitly allowlist only public resources
- Never serve files or data from outside designated public directories
- Validate and canonicalize all file paths before access
- Block exposure of backup files, internal data structures, and configuration files
- Configure proxies and load balancers to prevent header/data leakage

## Remediation Steps

- Restrict file serving to a designated public directory using absolute path validation
- Use secure filename functions to sanitize user input and prevent path traversal
- Verify resolved paths remain within allowed boundaries using canonical path comparison
- Implement extension blocklists for backup files (.bak, .swp, ~, .old)
- Remove or protect internal API endpoints that expose system data structures
- Review proxy/load balancer configurations to strip internal headers and prevent info leakage
