# CWE-80: Cross-Site Scripting (XSS) - Python

## LLM Guidance

XSS occurs when untrusted data is included in web output without proper encoding, allowing attackers to inject malicious scripts. Python frameworks like Django and Flask provide auto-escaping in templates; always use these built-in protections or explicit encoding with `html.escape()` for manual HTML construction.

**Primary Defence:** Use framework auto-escaping (Django templates, Flask/Jinja2) or explicit encoding functions.

## Key Principles

- Enable auto-escaping: Ensure Django/Jinja2 templates have auto-escaping enabled (default in most cases)
- Never mark untrusted data as safe: Avoid `mark_safe()`, `|safe` filter, or `Markup()` on user input
- Context-aware encoding: Use appropriate escaping for HTML attributes, JavaScript, CSS, or URLs
- Validate input types: Restrict input to expected formats before rendering
- Content Security Policy: Implement CSP headers as defence-in-depth

## Remediation Steps

- Audit template code for `|safe`, `mark_safe()`, or manual HTML concatenation with user data
- Replace manual HTML construction with template rendering or `html.escape()`
- Review JavaScript contexts where Python variables are embedded; use JSON encoding
- Enable auto-escaping in templates - `{% autoescape on %}` or verify it's enabled globally
- For unavoidable raw HTML, use sanitization libraries like `bleach` with strict allowlists
- Add CSP headers to restrict inline scripts

## Safe Pattern

```python
from flask import Flask, render_template_string, escape
import html

app = Flask(__name__)

@app.route('/user/<username>')
def show_user(username):
    # Template auto-escaping (safe)
    return render_template_string("<h1>Hello {{ name }}</h1>", name=username)

@app.route('/manual/<input>')
def manual_output(input):
    # Manual escaping when constructing HTML directly
    safe_html = f"<div>{html.escape(input)}</div>"
    return safe_html
```
