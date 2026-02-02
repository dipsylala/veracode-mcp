# CWE-79: Cross-Site Scripting (XSS) - Python

## LLM Guidance

XSS occurs when untrusted data is included in web output without proper encoding, allowing attackers to inject malicious scripts. Python frameworks like Django and Flask provide auto-escaping in templates-use `{{ variable }}` syntax and keep auto-escaping enabled. For manual encoding, use `html.escape()` or `bleach.clean()` with allowlists for rich content.

## Key Principles

- Use framework auto-escaping: Django templates and Jinja2 `.html` files escape by default
- Never mark untrusted input as safe: Avoid `|safe`, `mark_safe()`, or `Markup()` on user data
- Context-aware encoding: Use HTML escaping for HTML context, JavaScript encoding for `<script>` blocks
- Sanitize rich content: Use `bleach.clean()` with strict allowlists for permitted HTML tags/attributes
- Validate input format: Reject unexpected formats early before rendering

## Remediation Steps

- Enable and verify template auto-escaping is active (Django - `TEMPLATES['OPTIONS']['autoescape'] = True`)
- Replace `mark_safe()` or `|safe` filters on user-controlled variables with proper escaping
- Use `html.escape()` when rendering user input in non-template contexts
- For rich HTML, use `bleach.clean(user_input, tags=['b', 'i', 'u'], attributes={})` with minimal allowlists
- Set `Content-Security-Policy` headers to restrict script execution
- Audit all template rendering and ensure no raw user input reaches the DOM

## Safe Pattern

```python
from flask import Flask, render_template_string
import html

app = Flask(__name__)

@app.route('/profile')
def profile():
    username = request.args.get('name', '')
    # Auto-escaped in template
    return render_template_string('<h1>Welcome {{ name }}</h1>', name=username)
    
    # Or manual escaping
    safe_name = html.escape(username)
    return f'<h1>Welcome {safe_name}</h1>'
```
