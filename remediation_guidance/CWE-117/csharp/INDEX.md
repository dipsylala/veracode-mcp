# CWE-117: Log Injection / Log Forging - C# / ASP.NET

## LLM Guidance

Log Injection occurs when untrusted data is written to log files without sanitization, allowing attackers to forge log entries, hide malicious activity, or inject malicious content into log viewing tools. The primary defense is structured logging with JSON formatters (Microsoft.Extensions.Logging with JSON output or Serilog) that automatically encode control characters within structured fields. For manual logging, encode newlines (`\n` → `\\n`) and carriage returns (`\r` → `\\r`) before writing user input to logs.

## Remediation Strategy

- Use structured logging frameworks with JSON formatters to isolate user data in properly escaped fields
- Never concatenate user input directly into log message strings
- Encode or strip control characters (`\n`, `\r`, form feeds) from user input before logging
- Validate and sanitize log inputs at application boundaries
- Use parameterized logging with message templates instead of string interpolation

## Remediation Steps

- Adopt structured logging with `Microsoft.Extensions.Logging` and JSON formatter or Serilog with JSON sinks
- Replace string concatenation/interpolation with parameterized logging - `logger.LogInformation("User {Username} logged in", username)`
- If manual logging is required, encode control characters - replace `\n` with `\\n`, `\r` with `\\r`
- Configure log outputs to use structured formats (JSON, ECS) rather than plain text
- Review existing logging statements to ensure user input is passed as parameters, not embedded in format strings
- Implement input validation to reject or sanitize suspicious patterns before logging

## Minimal Safe Pattern

```csharp
using Microsoft.Extensions.Logging;

public class UserService
{
    private readonly ILogger<UserService> _logger;
    
    public void ProcessLogin(string username)
    {
        // SAFE: Structured logging with parameters
        _logger.LogInformation("User {Username} attempted login", username);
        
        // UNSAFE: String interpolation allows injection
        // _logger.LogInformation($"User {username} attempted login");
    }
}
```
