# CWE-200: Exposure of Sensitive Information to an Unauthorized Actor

## LLM Guidance

Information exposure occurs when applications reveal sensitive data to unauthorized users through error messages, APIs, configuration files, or exposed resources. Information that appears harmless individually (user enumeration, timing differences, stack traces) can be combined to enable sophisticated attacks like account compromise. Never return sensitive or internal information to clients unless explicitly required for their authorized function.

## Key Principles

- Use allowlisted exposure model: construct responses from explicitly approved data, not internal state
- Sanitize all user-facing outputs including error messages, API responses, headers, and logs
- Implement generic error handling that reveals no internal system details
- Validate what information each user role legitimately needs before exposing it
- Assume attackers will combine multiple small leaks to build attack chains

## Remediation Steps

- Review flaw details to identify the specific file, line number, and code pattern exposing information
- Identify what sensitive data is leaking - stack traces, credentials, file paths, user data, internal IDs, system details
- Trace the data flow from source to exposure point through error handlers, API responses, logs, or HTTP headers
- Determine the audience - who can access this information (authenticated users, anonymous users, public internet)
- Replace detailed error messages with generic user-facing messages; log full details server-side only
- Implement response filtering to strip sensitive fields before returning data to clients
