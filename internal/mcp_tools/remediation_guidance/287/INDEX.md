# CWE-287: Improper Authentication

## LLM Guidance

Improper authentication occurs when an application fails to correctly verify user, service, or system identity through flawed authentication logic itself-not just weak credentials. These flaws allow attackers to bypass authentication entirely, impersonate legitimate users, or escalate privileges. The core fix is never trusting client-supplied identity and requiring server-validated authentication for every request.

## Key Principles

- Never trust client-supplied identity claims-always validate server-side
- Implement complete authentication checks before accepting any session or identity
- Use defence-in-depth with multiple authentication factors
- Enforce authentication uniformly across all protected resources
- Properly manage session lifecycle (generation, timeout, invalidation)

## Remediation Steps

- Identify all authentication entry points (login forms, API auth, SSO, OAuth, password reset)
- Examine credential validation logic for bypass conditions or incomplete checks
- Review session token generation, storage, timeout, and invalidation mechanisms
- Audit all endpoints to ensure authentication is required before access
- Find and protect resources lacking authentication checks, especially admin functions and APIs
- Implement multi-factor authentication for sensitive operations
