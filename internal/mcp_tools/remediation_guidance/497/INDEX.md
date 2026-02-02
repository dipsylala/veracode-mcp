# CWE-497: Exposure of Sensitive System Information to an Unauthorized Control Sphere

## LLM Guidance

System information exposure occurs when applications leak internal details (server versions, paths, database names, framework versions, OS info) through error messages, headers, APIs, or pages, providing attackers reconnaissance data for targeted exploits. The core fix is to minimize disclosure of system internals and restrict information to what users legitimately need.

## Key Principles

- Do not expose system internals (paths, versions, stack traces) beyond what is necessary
- Remove or sanitize identifying HTTP headers in production
- Disable debug modes and verbose error messages in production environments
- Implement generic error pages that don't reveal technical details
- Limit information disclosure in API endpoints like /info, /status, /health

## Remediation Steps

- Check HTTP headers - Remove Server, X-Powered-By, X-AspNet-Version, and framework-specific headers
- Review error messages - Replace stack traces and database errors with generic messages for users
- Disable debug output - Turn off debug pages, version endpoints, and verbose logging in production
- Sanitize API responses - Remove version numbers, internal paths, and technology details from public APIs
- Examine data_paths in findings - Trace where system information flows to unauthorized users
- Configure error handling - Implement custom error pages that log details internally without exposing them externally
