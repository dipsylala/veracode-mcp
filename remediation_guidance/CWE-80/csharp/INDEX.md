# CWE-80: Cross-Site Scripting (XSS) - C# / ASP.NET

## LLM Guidance

XSS occurs when untrusted data is included in web output without proper encoding, allowing attackers to inject malicious scripts that execute in victims' browsers. In C#/ASP.NET applications, leverage Razor's automatic HTML encoding and avoid raw HTML output. Always encode dynamic content using `HttpUtility.HtmlEncode()` or `System.Net.WebUtility.HtmlEncode()`, implement Content Security Policy headers, and validate input using allowlists.

## Remediation Strategy

- Use automatic encoding: Rely on Razor's `@variable` syntax which HTML-encodes by default
- Avoid raw output: Never use `@Html.Raw()`, `HtmlString`, or `MvcHtmlString` with untrusted data
- Context-aware encoding: Apply JavaScript encoding for script contexts, URL encoding for URLs, and attribute encoding for HTML attributes
- Defense in depth: Combine output encoding with CSP headers and input validation
- Sanitize rich content: Use libraries like HtmlSanitizer for user-submitted HTML

## Remediation Steps

- Replace all `@Html.Raw(userInput)` with `@userInput` to enable automatic encoding
- For JavaScript contexts, use `@System.Text.Encodings.Web.JavaScriptEncoder.Default.Encode(value)`
- Implement CSP headers via middleware - `context.Response.Headers.Add("Content-Security-Policy", "default-src 'self'")`
- Use `[ValidateInput(true)]` on controllers and validate with allowlists for expected input formats
- For rich text editors, sanitize HTML with `Ganss.Xss.HtmlSanitizer` before storage
- Audit existing code for `InnerHtml`, `Write()`, and string concatenation in JavaScript

## Minimal Safe Pattern

```csharp
// Razor view - automatic encoding
<div>
    <h1>Welcome @Model.UserName</h1>
    <p>@Model.UserComment</p>
</div>

// Controller with CSP header
public IActionResult Index()
{
    Response.Headers.Add("Content-Security-Policy", "default-src 'self'");
    return View(model);
}

// Manual encoding when needed
var encoded = System.Net.WebUtility.HtmlEncode(userInput);
```
