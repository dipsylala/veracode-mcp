# CWE-113: HTTP Response Splitting - C# / ASP.NET

## LLM Guidance

HTTP Response Splitting occurs when attackers inject CRLF (`\r\n`) characters into HTTP headers, enabling them to inject additional headers or response bodies, potentially leading to cache poisoning, XSS, or session hijacking. The vulnerability arises when user input is directly concatenated into HTTP headers without sanitization. Always use ASP.NET Core's built-in methods that automatically sanitize headers and avoid manual header construction.

## Remediation Strategy

- Use ASP.NET Core framework methods (`Redirect()`, `RedirectToAction()`, `Response.Cookies.Append()`) that automatically encode/sanitize values
- Never manually concatenate user input into `Response.Headers` or construct raw HTTP responses
- Reject or strip CRLF characters (`\r`, `\n`) from any user input destined for headers
- Use `Uri.IsWellFormedUriString()` to validate URLs before redirects
- Enable response header validation in web.config or through middleware

## Remediation Steps

- Replace manual `Response.AddHeader()` or `Response.Headers.Add()` calls with framework methods
- Use `Redirect()` or `RedirectToAction()` instead of setting `Location` header manually
- Validate redirect URLs with `Uri.TryCreate()` and whitelist allowed domains
- Remove `\r`, `\n`, `%0d`, `%0a` from inputs using `input.Replace("\r", "").Replace("\n", "")`
- Set `cookieOptions.HttpOnly = true` and use `Response.Cookies.Append()` for cookies
- Enable ASP.NET Core's built-in header validation (enabled by default in Core 2.1+)

## Minimal Safe Pattern

```csharp
// Safe redirect with ASP.NET Core built-in validation
public IActionResult SafeRedirect(string returnUrl)
{
    // Validate URL is well-formed and local
    if (!string.IsNullOrEmpty(returnUrl) && Url.IsLocalUrl(returnUrl))
    {
        return Redirect(returnUrl); // Framework sanitizes automatically
    }
    return RedirectToAction("Index", "Home");
}

// Safe cookie setting
Response.Cookies.Append("UserPref", userValue, new CookieOptions 
{ 
    HttpOnly = true, 
    Secure = true 
});
```
