# CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes - Python

## LLM Guidance

Mass assignment vulnerabilities occur when user input is directly mapped to object attributes without validation, allowing attackers to modify unintended fields like `is_admin` or `role`. In Python, this happens through unsafe patterns like `fields = "__all__"` in Django forms/serializers, `setattr()` loops, or `**request.data` unpacking. Always use explicit allowlists to control which attributes can be modified.

## Key Principles

- Use explicit field allowlists in forms/serializers; never use `fields = "__all__"`
- Validate and filter input before assigning to model attributes
- Mark sensitive fields as `read_only` in serializers or `editable=False` in models
- Use separate serializers/forms for create vs update operations with different field sets

## Remediation Steps

- Replace `fields = "__all__"` with explicit field lists - `fields = ['name', 'email', 'description']`
- Set `read_only_fields = ['is_admin', 'created_by', 'role']` for protected attributes
- Use `editable=False` on model fields that should never be user-modified
- Avoid `for k,v in request.data.items() - setattr(obj, k, v)` patterns
- Create separate `CreateSerializer` and `UpdateSerializer` classes limiting exposed fields
- Validate input with `clean()` methods before mass operations

## Safe Pattern

```python
# Django REST Framework - Explicit field control
class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ['username', 'email', 'bio']  # Explicit allowlist
        read_only_fields = ['is_staff', 'is_superuser', 'date_joined']
```
