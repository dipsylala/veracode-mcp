# CWE-359: Exposure of Private Personal Information to an Unauthorized Actor

## LLM Guidance

Privacy violations occur when applications expose PII (names, SSN, medical data, financial info) through logs, error messages, APIs, URLs, or insufficient access controls, violating GDPR/CCPA/HIPAA and enabling identity theft, fraud, and legal liability. Core fix: minimize PII collection and enforce strict authorization with least disclosure principles.

## Key Principles

- Minimize PII collection and retention—only collect what's necessary
- Enforce authorization and least privilege for all PII access
- Control exposure channels: logs, errors, APIs, URLs, analytics, UI
- Delete PII promptly when no longer needed
- Question every PII field in forms and data models

## Remediation Steps

- Review flaw details to identify specific exposure location (file, line, code pattern)
- Identify exposed PII types - names, SSN, medical data, financial info, contact details
- Determine exposure channel - logs, error messages, APIs, URLs, database queries, analytics, UI
- Trace data flow from collection through storage to exposure point
- Implement data minimization - remove unnecessary PII fields from collection
- Add access controls - validate user authorization before returning PII in responses
