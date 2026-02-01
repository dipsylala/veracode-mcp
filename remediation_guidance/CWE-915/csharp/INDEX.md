# CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes - C\#

## LLM Guidance

Mass assignment vulnerabilities in C# occur when ASP.NET model binding automatically maps user input to object properties, allowing attackers to modify security-critical fields like `IsAdmin`, `Role`, or `Balance`. This guidance focuses on ASP.NET Core and newer. Always use ViewModels/DTOs with only permitted properties, apply `[Bind]` attributes to restrict binding, and validate all model state.

## Remediation Strategy

- Use dedicated ViewModels/DTOs containing only properties safe for user modification
- Never bind directly to domain models or entities with sensitive properties
- Apply `[BindNever]` or `[BindRequired]` attributes to control property binding explicitly
- Validate `ModelState.IsValid` before processing any bound data
- Use explicit property assignment instead of automatic model binding for sensitive operations

## Remediation Steps

- Create a ViewModel/DTO class with only the properties users should modify
- Apply `[BindNever]` attribute to sensitive properties if direct entity binding is unavoidable
- Replace controller action parameters from domain entities to ViewModels
- Use `[Bind(Include = "Prop1,Prop2")]` on action parameters to whitelist bindable properties
- Map ViewModel properties explicitly to domain entities using manual assignment or AutoMapper
- Always check `ModelState.IsValid` before processing bound data

## Minimal Safe Pattern

```csharp
// ViewModel with only safe properties
public class UpdateUserViewModel
{
    public string Name { get; set; }
    public string Email { get; set; }
    // IsAdmin, Role excluded - cannot be mass-assigned
}

[HttpPost]
public IActionResult Update(UpdateUserViewModel model)
{
    if (!ModelState.IsValid) return BadRequest(ModelState);
    
    var user = _db.Users.Find(User.Identity.Name);
    user.Name = model.Name;  // Explicit assignment
    user.Email = model.Email;
    _db.SaveChanges();
    return Ok();
}
```
