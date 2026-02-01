# CWE-93: CRLF Injection - C# / ASP.NET

## LLM Guidance

CRLF Injection occurs when attackers inject `\r\n` characters to manipulate HTTP headers, log files, or line-based formats, enabling HTTP Response Splitting, log forgery, or header manipulation. ASP.NET Core provides built-in header validation that automatically rejects invalid header values. Strip or reject newline characters from user input before using in headers or logs.

## Key Principles

- Never use raw user input in HTTP headers or log messages
- Leverage ASP.NET Core's automatic header validation (rejects `\r\n` by default)
- Sanitize input by removing or encoding `\r`, `\n`, and other control characters
- Use structured logging libraries that handle escaping automatically
- Validate redirect URLs and Location headers to prevent header injection

## Remediation Steps

- Replace or remove `\r` and `\n` characters from user input before header usage
- Use `Response.Headers.Append()` in ASP.NET Core which validates values
- For logging, use structured logging (e.g., Serilog, NLog) instead of string concatenation
- Validate and sanitize redirect URLs with `Uri.TryCreate()` and allowlists
- For legacy code, use `Regex.Replace(input, @"[\r\n]", "")` to strip newlines
- Test with payloads containing `%0d%0a` to verify protection

## Safe Pattern

```csharp
// Safe header setting with validation
public IActionResult SetCustomHeader(string userInput)
{
    // Remove CRLF characters
    var sanitized = Regex.Replace(userInput, @"[\r\n]", "");
    
    // ASP.NET Core validates automatically, but explicit sanitization is defense-in-depth
    Response.Headers.Append("X-Custom-Header", sanitized);
    
    return Ok();
}

// Safe logging with structured approach
_logger.LogInformation("User action: {UserInput}", userInput); // Library escapes automatically
```
