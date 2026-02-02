# CWE-209: Error Message Information Leak

## LLM Guidance

Error Message Information Leak occurs when detailed error messages expose sensitive information about the application's internal structure, configuration, or data to users. This includes stack traces, file paths, database errors, SQL queries, and system configuration details. The core fix is to display only generic error messages to users while logging detailed information server-side.

## Key Principles

- Never expose internal error details to clients; error responses must come from a fixed, server-controlled contract
- Separate user-facing messages from internal diagnostic information
- Use generic messages in production: "An error occurred", "Invalid credentials", "Request failed"
- Log detailed errors server-side for debugging and monitoring
- Avoid user enumeration through error message differences

## Remediation Steps

- Review flaw details to identify where detailed error messages are exposed
- Trace the error flow from exception handling to user response
- Implement generic error messages for all production error responses
- Configure error handlers (404, 500, API errors) to return sanitized messages
- Add server-side logging for full error details including stack traces
- Validate that error responses don't leak paths, queries, or configuration details
