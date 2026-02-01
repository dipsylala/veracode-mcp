# CWE-566: Authorization Bypass Through User-Controlled Key - Python

## LLM Guidance

Authorization bypass through user-controlled keys (IDOR) occurs when Python web applications use user-supplied identifiers to retrieve resources without verifying the authenticated user has permission to access them. This enables horizontal privilege escalation where attackers access other users' data by manipulating URL parameters or request values. The core fix is to implement ownership verification before returning resources.

## Key Principles

- **Never trust user-controlled resource identifiers** - Always validate ownership before returning data
- **Implement access control checks at the data layer** - Verify `resource.owner_id == current_user.id`
- **Use framework-level authorization decorators** - Enforce permissions consistently across endpoints
- **Query by composite keys** - Filter by both resource ID and user ID: `filter(id=doc_id, user_id=current_user.id)`
- **Return 404 instead of 403** - Avoid leaking resource existence information

## Remediation Steps

- Locate the vulnerability - Find where user-controlled IDs flow into database queries without authorization checks
- Add ownership verification - Before returning resources, confirm `resource.owner_id == authenticated_user.id`
- Use filtered queries - Query with both resource ID and user ID - `Document.query.filter_by(id=doc_id, user_id=current_user.id).first_or_404()`
- Implement authorization middleware - Use decorators or middleware to enforce checks consistently
- Handle missing resources uniformly - Return 404 for both non-existent and unauthorized resources
- Test privilege escalation - Verify users cannot access others' resources by ID manipulation

## Minimal Safe Pattern

```python
from flask import abort
from flask_login import current_user

@app.route('/documents/<int:doc_id>')
@login_required
def get_document(doc_id):
    # Query with BOTH resource ID and user ID
    document = Document.query.filter_by(
        id=doc_id,
        user_id=current_user.id
    ).first_or_404()
    
    return jsonify(document.to_dict())
```
