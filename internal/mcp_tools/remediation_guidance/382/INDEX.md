# CWE-382: J2EE Bad Practices: Use of System.exit()

## LLM Guidance

Calling System.exit() in J2EE applications terminates the entire application server, affecting all deployed applications and users. This violates the J2EE threading model, prevents proper cleanup of container-managed resources, and causes denial of service.

## Key Principles

- Never expose dangerous functionality like System.exit() in untrusted J2EE contexts
- Use exception-based error handling instead of process termination
- Rely on container-managed lifecycle for application shutdown
- Keep privileged operations gated and isolated from application code
- Allow the J2EE container to manage resource cleanup and thread lifecycle

## Remediation Steps

- Locate calls - Search for `System.exit` in servlets, EJBs, filters, and JSPs using grep or IDE search
- Identify context - Determine why System.exit() is called (error handling, validation failure, shutdown logic)
- Replace with exceptions - Use `throw new ServletException("message")` or appropriate exception types instead of System.exit()
- Use container shutdown - If legitimate shutdown needed, use container management tools or JMX beans
- Implement proper error handling - Return error responses, log failures, and set HTTP status codes appropriately
- Verify fix - Test error paths to confirm server remains operational after handled failures

## Wrong vs. Right Pattern

- Wrong: `if (error) { System.exit(1); }` - kills entire application server
- Right: `if (error) { throw new ServletException("Error message"); }` - terminates request, logs error, server continues
