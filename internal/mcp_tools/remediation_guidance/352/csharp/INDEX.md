# CWE-352: Cross-Site Request Forgery (CSRF) - C\#

## LLM Guidance

CSRF vulnerabilities occur when state-changing endpoints don't verify requests originated from the application itself, allowing attackers to perform actions on behalf of authenticated users. ASP.NET Core provides built-in anti-forgery token support that should be enabled on all state-changing operations. The primary defense is using `[ValidateAntiForgeryToken]` attribute on POST/PUT/DELETE actions or enabling automatic validation globally.

## Key Principles

- Apply anti-forgery tokens to all state-changing HTTP methods (POST, PUT, DELETE, PATCH)
- Use automatic token validation globally rather than relying on per-action attributes
- Ensure tokens are included in forms via `@Html.AntiForgeryToken()` or auto-generated
- Validate SameSite cookie attributes are set to `Strict` or `Lax`
- Never disable CSRF protection for authenticated endpoints

## Remediation Steps

- Add `services.AddAntiforgery()` to `ConfigureServices` in Startup.cs
- Enable automatic validation with `options.Filters.Add(new AutoValidateAntiforgeryTokenAttribute())`
- Add `@Html.AntiForgeryToken()` to all forms or use Tag Helpers with `method="post"`
- Apply `[ValidateAntiForgeryToken]` to individual controllers/actions if not using global filters
- Configure SameSite cookies - `options.Cookie.SameSite = SameSiteMode.Strict`
- Test protected endpoints reject requests without valid tokens

## Safe Pattern

```csharp
// Startup.cs - Global configuration
public void ConfigureServices(IServiceCollection services)
{
    services.AddControllersWithViews(options =>
    {
        options.Filters.Add(new AutoValidateAntiforgeryTokenAttribute());
    });
}

// Controller action (token validated automatically)
[HttpPost]
public IActionResult UpdateProfile(ProfileModel model)
{
    // Token validation happens automatically
    return View();
}
```
