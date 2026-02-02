# CWE-89: SQL Injection - Python

## LLM Guidance

SQL Injection occurs when untrusted user input is incorporated into SQL queries without proper sanitization, allowing attackers to manipulate query logic, extract data, or execute unauthorized database operations. Python's database libraries (sqlite3, psycopg2, mysql-connector) all support parameterized queries, which is the primary defense mechanism. Always use parameterized queries instead of string concatenation or f-strings when building SQL statements.

## Key Principles

- Use parameterized queries exclusively - Never concatenate user input into SQL strings
- Employ ORM frameworks - Use SQLAlchemy, Django ORM, or similar frameworks that handle parameterization
- Apply input validation - Validate data types, formats, and ranges as a secondary defense layer
- Use least privilege - Database accounts should have minimal necessary permissions
- Escape dynamic identifiers - When table/column names must be dynamic, use allowlisting

## Remediation Steps

- Replace all string concatenation and f-strings in SQL queries with parameterized placeholders
- Use library-specific parameter syntax (?, %s, or  -name depending on the database driver)
- Pass user input as separate arguments to execute() methods, never embedded in query strings
- For dynamic table/column names, validate against a predefined allowlist of safe values
- Review all database interaction code for direct string manipulation
- Implement automated testing with SQL injection payloads

## Safe Pattern

```python
import sqlite3

# UNSAFE: String concatenation
# query = f"SELECT * FROM users WHERE username = '{user_input}'"

# SAFE: Parameterized query
conn = sqlite3.connect('database.db')
cursor = conn.cursor()

user_input = request.get('username')
query = "SELECT * FROM users WHERE username = ?"
cursor.execute(query, (user_input,))

results = cursor.fetchall()
```
