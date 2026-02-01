# CWE-79: Cross-Site Scripting (XSS) - C# / ASP.NET

## LLM Guidance

XSS occurs when untrusted data is included in web output without proper encoding, allowing attackers to inject malicious scripts into victim browsers. In C#/ASP.NET, leverage built-in auto-encoding features and context-specific encoders to prevent malicious content from executing.

## Remediation Strategy

- Use Razor's automatic HTML encoding with `@variable` syntax (ASP.NET Core/MVC)
- Apply context-specific encoders: `HtmlEncoder` for HTML, `JavaScriptEncoder` for JS contexts, `UrlEncoder` for URLs
- Sanitize rich HTML with HtmlSanitizer library before using `@Html.Raw()`
- Implement Content Security Policy headers for defense-in-depth
- Validate input format as secondary defense, never rely on it alone

## Remediation Steps

- Replace `@Html.Raw()`, `Response.Write()`, and Literal controls with Razor's `@variable` auto-encoding
- Use `HtmlEncoder.Default.Encode()` for explicit HTML encoding in controllers or classic ASP.NET
- Apply `JavaScriptEncoder.Default.Encode()` when embedding data in `<script>` blocks or JS event handlers
- Use `UrlEncoder.Default.Encode()` for query parameters and URL components
- For rich text editors, integrate HtmlSanitizer with allowlist of safe tags before rendering with `@Html.Raw()`
- Add CSP middleware to restrict script sources and enforce 'self' policy

## Minimal Safe Pattern

```csharp
// Razor view - auto-encoded by default
<div>Welcome, @Model.UserName</div>

// Explicit HTML encoding in controller
using System.Text.Encodings.Web;
public IActionResult Display(string input)
{
    ViewBag.Safe = HtmlEncoder.Default.Encode(input);
    return View();
}

// JavaScript context encoding
<script>
    var msg = '@JavaScriptEncoder.Default.Encode(ViewBag.Message)';
</script>
```
