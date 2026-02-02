# CWE-284: Improper Access Control

## LLM Guidance

Improper access control occurs when applications fail to restrict what authenticated users can access or do. While authentication verifies identity ("who you are"), authorization verifies permissions ("what you're allowed to access") - strong authentication doesn't prevent authenticated users from accessing unauthorized data. Fix by implementing deny-by-default, server-side authorization checks on every protected resource and operation.

## Key Principles

- Implement deny-by-default access control - explicitly grant access rather than blocking known threats
- Validate authorization server-side on every request to protected resources
- Never trust client-side controls, hidden fields, or user input for authorization decisions
- Use role-based or attribute-based access control (RBAC/ABAC) frameworks
- Apply principle of least privilege - grant minimum necessary permissions

## Remediation Steps

- Identify the protected resource lacking authorization checks (admin panel, user data, API endpoint, sensitive operation)
- Trace data flow to understand how the resource is accessed (direct URL, API call, parameter manipulation)
- Implement server-side authorization checks before granting access to any protected resource
- Validate user permissions against the requested operation and specific resource (not just resource type)
- Apply authorization checks consistently across all entry points (UI, API, direct access, background jobs)
- Test by attempting unauthorized access with different user roles and privilege levels
