# CWE-566: Authorization Bypass Through User-Controlled Key - Java

## LLM Guidance

Authorization bypass through user-controlled keys (IDOR) occurs when Java applications use user-supplied identifiers (order IDs, user IDs, document IDs) in database queries without verifying the authenticated user has authorization to access those resources. This enables horizontal privilege escalation where attackers can access, modify, or delete other users' data. The core fix is to validate ownership or permissions before granting access to any resource identified by user input.

## Key Principles

- Never trust user-supplied resource identifiers without authorization checks
- Implement ownership validation by comparing resource owner with authenticated user
- Use authorization frameworks (Spring Security, Apache Shiro) for consistent enforcement
- Apply principle of least privilege in database queries by filtering on user context
- Validate access at the service layer, not just presentation layer

## Remediation Steps

- Locate user-controlled inputs (`@PathVariable`, `@RequestParam`, `@PathParam`) used as resource identifiers
- Trace the data flow to database queries or resource access methods
- Identify missing authorization checks between resource retrieval and usage
- Add ownership validation comparing `entity.getUserId()` with authenticated user ID
- Implement access control checks before returning or modifying resources
- Test with different authenticated users attempting to access each other's resources

## Safe Pattern

```java
@GetMapping("/orders/{orderId}")
public Order getOrder(@PathVariable Long orderId, Authentication auth) {
    Order order = orderRepository.findById(orderId)
        .orElseThrow(() -> new ResourceNotFoundException());
    
    String currentUser = auth.getName();
    if (!order.getUserId().equals(currentUser)) {
        throw new AccessDeniedException("Not authorized");
    }
    
    return order;
}
```
