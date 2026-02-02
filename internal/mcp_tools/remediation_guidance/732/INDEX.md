# CWE-732: Incorrect Permission Assignment for Critical Resource

## LLM Guidance

This vulnerability occurs when critical resources (files, directories, services) are assigned overly broad or incorrect permissions, allowing unauthorized users to access or modify them. The core fix is applying least privilege principles: default-deny access and grant only minimum necessary permissions for legitimate operations.

## Key Principles

- Apply least privilege to all permission assignments and ACLs
- Use default-deny approach: start restrictive, grant only required access
- Minimize scope: limit who can read, write, or execute resources
- Validate permissions match resource sensitivity (config files, credentials, user data require stricter controls)
- Review and audit permissions regularly to prevent drift

## Remediation Steps

- Review flaw details to identify the specific resource (file, directory, service) with incorrect permissions
- Identify current permissions using `ls -l` (Unix), `Get-Acl` (Windows), or application-specific tools
- Determine resource sensitivity and required access patterns
- Assign minimum necessary permissions - grant only what's needed for legitimate operations
- Remove world-readable/writable permissions; use user/group-specific grants
- Validate changes with security tests - attempt unauthorized access to confirm restrictions work
