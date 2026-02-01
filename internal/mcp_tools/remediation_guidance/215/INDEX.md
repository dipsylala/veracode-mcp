# CWE-215: Insertion of Sensitive Information Into Debugging Code

## LLM Guidance

Insertion of sensitive information into debugging code occurs when debug statements, verbose logging, stack traces, or development features expose passwords, tokens, internal paths, SQL queries, or system architecture in production environments, enabling information disclosure and attack reconnaissance.

## Key Principles

- Never expose diagnostic or debugging instrumentation to untrusted clients
- Runtime responses must be constructed independently of debug or developer-only state
- Separate development-time debugging from production error handling
- Implement environment-specific configurations that disable debug features in production
- Log sensitive operations securely without exposing credential values or internal paths

## Remediation Steps

- Identify debug mode flags, verbose logging, and debug endpoints in production code
- Locate error handlers that reveal stack traces or internal details to users
- Disable all debug modes and verbose error output in production environments
- Replace detailed error messages with generic user-facing responses
- Remove or conditionally disable debug logging that records passwords, tokens, or API keys
- Review exception handling to ensure stack traces and system details are logged server-side only
