# CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes - Java

## LLM Guidance

Mass assignment vulnerabilities in Java occur when Spring MVC/Boot automatically binds HTTP request parameters to object fields, allowing attackers to modify security-critical fields like `isAdmin`, `role`, or `balance`. Use DTOs with only permitted fields for user input, apply `@JsonIgnoreProperties` or allowlist binding with `@InitBinder`, and validate with Bean Validation annotations. Never bind request data directly to JPA entities or domain objects.

## Remediation Strategy

- Use separate DTOs for user input that expose only modifiable fields
- Apply `@JsonIgnoreProperties(ignoreUnknown = true)` and exclude sensitive fields with `@JsonIgnore`
- Restrict form parameter binding using `@InitBinder` with `setAllowedFields()` or `setDisallowedFields()`
- Validate all input with Bean Validation constraints (`@Valid`, `@NotNull`, `@Size`, etc.)
- Map DTO fields explicitly to entities rather than using reflection-based copiers

## Remediation Steps

- Create a DTO class containing only fields users should modify (e.g., `UpdateUserDTO` with `name`, `email`)
- Annotate sensitive entity fields with `@JsonIgnore` to prevent JSON binding
- Add `@InitBinder` method to controllers restricting allowed form fields
- Apply `@Valid` to controller method parameters and handle `BindingResult` errors
- Use explicit field mapping when transferring DTO data to entities (avoid `BeanUtils.copyProperties`)
- Never expose JPA entities directly as `@RequestBody` or form-backing objects

## Minimal Safe Pattern

```java
// DTO with only permitted fields
public class UpdateUserDTO {
    @NotBlank
    private String name;
    @Email
    private String email;
    // Getters/setters - NO isAdmin, role, or other sensitive fields
}

@PostMapping("/users/{id}")
public ResponseEntity<?> updateUser(@PathVariable Long id, 
                                     @Valid @RequestBody UpdateUserDTO dto) {
    User user = userRepository.findById(id).orElseThrow();
    user.setName(dto.getName());  // Explicit mapping only
    user.setEmail(dto.getEmail());
    userRepository.save(user);
    return ResponseEntity.ok(user);
}
```
