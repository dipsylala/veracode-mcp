# CWE-566: Authorization Bypass Through User-Controlled Key - C\#

## LLM Guidance

Authorization bypass through user-controlled keys (IDOR) occurs when C# applications use user-supplied identifiers (order IDs, user IDs, document IDs) in database queries without verifying the authenticated user has authorization to access those resources. This enables horizontal privilege escalation where attackers can access, modify, or delete other users' data. Always validate ownership or permissions before allowing resource access.

## Key Principles

- Never trust user-supplied resource identifiers without authorization checks
- Filter queries by authenticated user ID or verify ownership after retrieval
- Implement centralized authorization policies using ASP.NET Core Authorization or repository patterns
- Use strongly-typed claims (`ClaimsPrincipal.FindFirstValue(ClaimTypes.NameIdentifier)`) to get authenticated user context
- Return 404 instead of 403 for unauthorized resources to avoid information disclosure

## Remediation Steps

- Identify user-controlled inputs - route parameters (`[FromRoute]`), query strings (`[FromQuery]`), request body properties
- Trace the ID to database queries (`FindAsync`, `FirstOrDefaultAsync`, LINQ) lacking user filters
- Extract authenticated user ID from `User.FindFirstValue(ClaimTypes.NameIdentifier)`
- Add authorization check - filter queries with `.Where(e => e.UserId == currentUserId)` or verify `entity.UserId == currentUserId` after retrieval
- Return `NotFound()` if entity doesn't exist or user lacks access
- Apply authorization attributes (`[Authorize]`) and consider policy-based authorization for complex scenarios

## Safe Pattern

```csharp
[Authorize]
public async Task<IActionResult> GetOrder(int orderId)
{
    var userId = User.FindFirstValue(ClaimTypes.NameIdentifier);
    var order = await _context.Orders
        .Where(o => o.Id == orderId && o.UserId == userId)
        .FirstOrDefaultAsync();
    
    if (order == null)
        return NotFound();
    
    return Ok(order);
}
```
