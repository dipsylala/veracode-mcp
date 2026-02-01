# CWE-352: Cross-Site Request Forgery (CSRF) - Java

## LLM Guidance

CSRF vulnerabilities in Java web applications occur when state-changing endpoints don't verify that requests originated from the application itself, allowing attackers to execute unauthorized actions on behalf of authenticated users. Spring Security provides built-in CSRF protection using synchronizer tokens that must be included in state-changing requests. The defense requires enabling CSRF protection, including tokens in forms/AJAX calls, and using appropriate HTTP methods (POST/PUT/DELETE for state changes).

## Key Principles

- Enable Spring Security CSRF protection (enabled by default in Spring Boot)
- Include CSRF tokens in all state-changing requests (POST, PUT, DELETE, PATCH)
- Use SameSite cookie attributes to provide defense-in-depth
- Validate CSRF tokens on the server side for all non-safe HTTP methods
- Never disable CSRF protection globally without explicit security review

## Remediation Steps

- Add Spring Security dependency and enable CSRF protection in configuration
- Include `${_csrf.token}` hidden field in all HTML forms or add CSRF header to AJAX requests
- Configure SameSite=Strict or Lax on session cookies
- Use POST/PUT/DELETE for state-changing operations (never GET)
- Ensure Spring Security's CSRF filter is active in the filter chain
- Test that requests without valid tokens are rejected with 403 Forbidden

## Safe Pattern

```java
// Spring Boot Controller with CSRF protection
@Controller
public class AccountController {
    
    @PostMapping("/transfer")
    public String transfer(@RequestParam String recipient, 
                          @RequestParam BigDecimal amount) {
        // Spring Security automatically validates CSRF token
        accountService.transfer(recipient, amount);
        return "redirect:/success";
    }
}

// Thymeleaf template with CSRF token
// <form method="post" action="/transfer">
//   <input type="hidden" th:name="${_csrf.parameterName}" th:value="${_csrf.token}"/>
// </form>
```
