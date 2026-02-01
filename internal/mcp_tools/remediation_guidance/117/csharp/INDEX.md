# CWE-117: Log Injection / Log Forging - C# / ASP.NET

## LLM Guidance

Log Injection occurs when untrusted data is written to log files without sanitization, allowing attackers to forge log entries, hide malicious activity, or inject malicious content into log viewing tools. The primary defense is structured logging with JSON formatters (Microsoft.Extensions.Logging with JSON output or Serilog) that automatically encode control characters within structured fields. For manual logging, encode all control characters: ASCII controls (0x00-0x1F), DEL (0x7F), C1 controls (0x80-0x9F), and Unicode line separators (U+0085, U+2028, U+2029).

## Key Principles

- Use structured logging frameworks with JSON formatters to isolate user data in properly escaped fields
- Never concatenate user input directly into log message strings
- If structured logging is unavailable, encode ALL control characters: ASCII controls (0x00-0x1F), DEL (0x7F), C1 controls (0x80-0x9F), Unicode line separators (U+0085, U+2028, U+2029)
- Validate and sanitize log inputs at application boundaries
- Use parameterized logging with message templates instead of string interpolation

## Remediation Steps

- Adopt structured logging with `Microsoft.Extensions.Logging` and JSON formatter or Serilog with JSON sinks
- Replace string concatenation/interpolation with parameterized logging - `logger.LogInformation("User {Username} logged in", username)`
- For legacy systems without structured logging, implement comprehensive control character encoding (see Manual Encoding Pattern below)
- Configure log outputs to use structured formats (JSON, ECS) rather than plain text
- Review existing logging statements to ensure user input is passed as parameters, not embedded in format strings
- Test by attempting to inject newlines, null bytes, and Unicode line separators - verify proper encoding

## Safe Pattern (Structured Logging)

```csharp
using Microsoft.Extensions.Logging;

public class UserService
{
    private readonly ILogger<UserService> _logger;
    
    public void ProcessLogin(string username)
    {
        // SAFE: Structured logging with parameters
        // JSON formatter automatically encodes control characters
        _logger.LogInformation("User {Username} attempted login", username);
        
        // UNSAFE: String interpolation allows injection
        // _logger.LogInformation($"User {username} attempted login");
    }
}
```

## Safe Pattern (Legacy Systems)

If structured logging is unavailable, use comprehensive control character encoding:

```csharp
using System;
using System.Text;

public static class LogEncoder
{
    public static string EncodeForLog(string input)
    {
        if (string.IsNullOrEmpty(input))
            return input;
        
        var output = new StringBuilder(input.Length + 16);
        
        foreach (char ch in input)
        {
            int code = (int)ch;
            
            // Encode ASCII control chars (0x00-0x1F) + DEL (0x7F) + C1 controls (0x80-0x9F)
            if (code <= 0x1F || code == 0x7F || (code >= 0x80 && code <= 0x9F))
            {
                switch (ch)
                {
                    case '\\':
                        output.Append("\\\\");
                        break;
                    case '\r':
                        output.Append("\\r");
                        break;
                    case '\n':
                        output.Append("\\n");
                        break;
                    case '\t':
                        output.Append("\\t");
                        break;
                    default:
                        output.AppendFormat("\\u{0:x4}", code);
                        break;
                }
                continue;
            }
            
            // Encode Unicode line separators
            if (code == 0x0085 || code == 0x2028 || code == 0x2029)
            {
                output.AppendFormat("\\u{0:x4}", code);
                continue;
            }
            
            output.Append(ch);
        }
        
        return output.ToString();
    }
}

// Usage
logger.LogInformation($"User {LogEncoder.EncodeForLog(username)} attempted login");
Console.WriteLine($"Activity: {LogEncoder.EncodeForLog(userInput)}");
```
