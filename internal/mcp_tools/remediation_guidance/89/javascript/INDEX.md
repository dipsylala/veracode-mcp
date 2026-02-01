# CWE-89: SQL Injection - JavaScript

## LLM Guidance

SQL Injection occurs when untrusted user input is incorporated into SQL queries without proper sanitization, allowing attackers to manipulate query logic, extract data, or execute administrative operations. Node.js database libraries (mysql, pg, better-sqlite3, etc.) all support parameterized queries as the primary defense against this vulnerability.

## Key Principles

- Always use parameterized queries (prepared statements) instead of string concatenation or template literals
- Validate and sanitize user input before use, enforcing strict type checking and whitelist validation
- Apply principle of least privilege to database accounts used by the application
- Use ORMs cautiously, ensuring raw query methods still use parameterization
- Never trust client-side validation alone; always validate on the server

## Remediation Steps

- Identify all locations where user input flows into SQL queries
- Replace string concatenation and template literals with parameterized queries using placeholders (`?` or `$1`, `$2`)
- For dynamic column/table names, use strict whitelist validation rather than parameterization
- Review ORM usage to ensure `.raw()` or similar methods use proper parameterization
- Test with SQL injection payloads to verify fixes
- Implement input validation layers before data reaches queries

## Safe Pattern

```javascript
// Unsafe: String concatenation
const unsafe = `SELECT * FROM users WHERE id = ${userId}`;

// Safe: Parameterized query (mysql)
db.query('SELECT * FROM users WHERE id = ?', [userId], (err, results) => {
  // Handle results
});

// Safe: Parameterized query (pg)
await client.query('SELECT * FROM users WHERE id = $1', [userId]);

// Safe: Prepared statement (better-sqlite3)
const stmt = db.prepare('SELECT * FROM users WHERE id = ?');
const user = stmt.get(userId);
```
