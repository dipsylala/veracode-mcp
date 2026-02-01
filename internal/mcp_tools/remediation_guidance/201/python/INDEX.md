# CWE-201: Insertion of Sensitive Information Into Sent Data - Python

## LLM Guidance

Python web applications commonly expose sensitive data through HTTP responses, error messages, logs, or API responses. Frameworks like Django, Flask, and FastAPI make it easy to serialize entire objects or return detailed error traces, which can leak passwords, tokens, internal paths, PII, and configuration details. The core fix is to explicitly control what data is returned and sanitize error messages.

## Key Principles

- Explicitly whitelist response fields rather than serializing entire objects
- Disable debug mode and detailed error traces in production
- Sanitize exceptions before exposing them to clients
- Use structured logging that excludes sensitive fields
- Implement response DTOs or serializers with only necessary fields

## Remediation Steps

- Set `DEBUG = False` in Django/Flask production settings
- Define explicit serializer fields instead of using `fields = '__all__'`
- Catch exceptions and return generic error messages to clients
- Configure logging filters to redact sensitive data (passwords, tokens, keys)
- Use environment variables for secrets, never hardcode in responses
- Review API responses to ensure only required data is included

## Safe Pattern

```python
# FastAPI/Pydantic example - explicit response model
from pydantic import BaseModel

class UserResponse(BaseModel):
    id: int
    username: str
    email: str
    # Exclude: password_hash, api_key, internal_id

@app.get("/users/{user_id}", response_model=UserResponse)
async def get_user(user_id: int):
    user = db.get_user(user_id)
    return UserResponse(**user.dict())  # Only whitelisted fields returned
```
