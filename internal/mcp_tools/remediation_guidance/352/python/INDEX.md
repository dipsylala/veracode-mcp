# CWE-352: Cross-Site Request Forgery (CSRF) - Python

## LLM Guidance

CSRF occurs when state-changing endpoints don't validate that requests originated from the application itself, allowing attackers to forge authenticated requests from malicious sites. Enable Django's `CsrfViewMiddleware` with `{% csrf_token %}` in forms, or use Flask-WTF's `CSRFProtect` for automatic token validation. Configure `SESSION_COOKIE_SAMESITE='Strict'` as defence-in-depth.

## Key Principles

- Enable framework-native CSRF protection globally (Django middleware, Flask-WTF `CSRFProtect`)
- Validate cryptographic tokens on all state-changing operations (POST/PUT/PATCH/DELETE)
- Use SameSite=Strict cookies to prevent cross-site cookie transmission
- Include tokens in forms via template tags and AJAX via custom headers
- Never disable CSRF protection with `@csrf_exempt` or similar decorators

## Remediation Steps

- Add `django.middleware.csrf.CsrfViewMiddleware` to Django `MIDDLEWARE` settings
- Include `{% csrf_token %}` in all Django forms or initialize `CSRFProtect(app)` in Flask
- Configure secure cookies - `SESSION_COOKIE_SAMESITE='Strict'`, `CSRF_COOKIE_SECURE=True`
- For AJAX - read CSRF token from cookie and send in `X-CSRFToken` header
- Remove any `@csrf_exempt` decorators from state-changing endpoints
- Verify all POST/PUT/PATCH/DELETE routes validate tokens automatically

## Safe Pattern

```python
# Django settings.py
MIDDLEWARE = [
    'django.middleware.csrf.CsrfViewMiddleware',  # Enable CSRF
]
SESSION_COOKIE_SAMESITE = 'Strict'
CSRF_COOKIE_SECURE = True

# views.py
@login_required
@require_http_methods(["POST"])
def transfer_funds(request):
    # Token validated automatically by middleware
    perform_transfer(request.user, request.POST['to_account'], request.POST['amount'])
    return redirect('success')

# Template
# <form method="post">{% csrf_token %}<input name="to_account"><button>Transfer</button></form>
```
