# CWE-601: Open Redirect - C\# / ASP.NET

## LLM Guidance

Open Redirect vulnerabilities occur when an application redirects users to URLs controlled by attackers, enabling phishing attacks and credential theft. Use `Url.IsLocalUrl()` (ASP.NET Core) or `Request.IsUrlLocalToHost()` (ASP.NET MVC 5) to validate that redirect URLs are local before redirecting. These framework methods reject external URLs, protocol-relative URLs, and JavaScript URLs.

## Key Principles

- Always validate redirect URLs before using them in redirect responses
- Prefer allowlists of known-safe redirect destinations over validation
- Use framework-provided validation methods rather than custom regex patterns
- Reject URLs with absolute paths, external domains, or protocol handlers
- Default to safe fallback URLs when validation fails

## Remediation Steps

- Identify all redirect operations using `Response.Redirect()`, `RedirectToAction()`, or similar methods
- Check if redirect URLs come from user input (query parameters, form data, headers)
- Replace direct redirects with validation using `Url.IsLocalUrl()` or allowlist checks
- Add null/empty checks before validation to prevent exceptions
- Implement fallback redirects to safe defaults (e.g., home page) when validation fails
- Test with malicious URLs like `//evil.com`, `javascript -alert(1)`, and `http -//attacker.com`

## Safe Pattern

```csharp
public IActionResult Login(string returnUrl)
{
    // Validate authentication logic here
    
    if (!string.IsNullOrEmpty(returnUrl) && Url.IsLocalUrl(returnUrl))
    {
        return Redirect(returnUrl);
    }
    
    return RedirectToAction("Index", "Home");
}
```
