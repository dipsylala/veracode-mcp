# CWE-285: Improper Authorization - C# / ASP.NET Core

## LLM Guidance

In ASP.NET Core applications, improper authorization occurs when authentication middleware is configured via `UseAuthorization()` but controllers or actions lack `[Authorize]` or `[AllowAnonymous]` attributes, allowing unauthenticated access to protected resources. Missing or misconfigured attributes bypass authorization checks even when authentication middleware is enabled.

## Key Principles

- Apply `[Authorize]` at controller level to secure all endpoints by default
- Use `[AllowAnonymous]` only for explicitly public endpoints
- Avoid relying solely on middleware configuration without attribute enforcement
- Combine authentication and role/policy-based authorization for sensitive operations
- Never assume endpoints are protected without explicit authorization attributes

## Remediation Steps

- Add `[Authorize]` attribute to all controller classes requiring authentication
- Mark public endpoints with `[AllowAnonymous]` attribute explicitly
- Apply role-based authorization (`[Authorize(Roles = "Admin")]`) for privileged operations
- Use policy-based authorization for complex access control requirements
- Verify authorization configuration in `Program.cs` includes both `UseAuthentication()` and `UseAuthorization()`
- Test each endpoint to confirm unauthorized access is blocked

## Safe Pattern

```csharp
[Authorize]  // Protect all actions by default
[ApiController]
[Route("api/[controller]")]
public class UserController : ControllerBase
{
    [AllowAnonymous]  // Explicitly allow public access
    [HttpGet("public")]
    public IActionResult GetPublicInfo() => Ok("Public data");
    
    [HttpGet("profile/{id}")]  // Requires authentication
    public IActionResult GetProfile(int id) => Ok(_userService.GetUser(id));
    
    [Authorize(Roles = "Admin")]  // Requires admin role
    [HttpDelete("{id}")]
    public IActionResult DeleteUser(int id) => NoContent();
}
```
