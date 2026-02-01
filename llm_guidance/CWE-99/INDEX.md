# CWE-99: Resource Injection

## LLM Guidance

Resource Injection occurs when untrusted input selects system resources (files, ports, class names, URLs) without validation, allowing attackers to manipulate which resources the application accesses. The core fix is to never let untrusted input directly select resources; instead, use allowlisted mappings and canonical path validation to ensure only permitted resources are accessed.

## Remediation Strategy

- Never let untrusted input select resources by name or path directly
- Canonicalize all resource identifiers (resolve paths, normalize names) before validation
- Map user-controlled input to allowlisted resources using indirect references
- Validate against strict allowlists of permitted resources, not denylists
- Use resource IDs or tokens that map server-side to actual resource locations

## Remediation Steps

- Identify sources - Find all untrusted data sources (HTTP parameters, external APIs, databases, file uploads, network requests)
- Trace resource selection - Follow how untrusted data flows to resource access points (file paths, port numbers, class names, URLs)
- Locate sinks - Identify where resources are accessed (file operations, network connections, class loading, database connections)
- Detect missing validation - Find resource selection without allowlist checks or canonicalization
- Implement allowlists - Create strict allowlists of permitted resources and validate all input against them before access
- Use indirect references - Replace direct resource names with IDs or tokens mapped server-side to actual resources
