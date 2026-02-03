# CWE-89: SQL Injection - C\#

## LLM Guidance

SQL Injection occurs when untrusted data is incorporated into SQL queries without proper validation or parameterization, allowing attackers to manipulate queries to bypass authentication, access unauthorized data, modify/delete records, or execute administrative operations.

**Primary Defence:** Use parameterized queries with `SqlCommand.Parameters`, Entity Framework LINQ queries, or Dapper with parameter binding.

## Key Principles

- Always use parameterized queries; never concatenate user input into SQL strings
- Prefer Entity Framework LINQ queries which automatically parameterize
- Use stored procedures with parameters when direct SQL is unavoidable
- Validate and sanitize all user input at entry points
- Apply least privilege to database accounts

## Remediation Steps

- Review data flow from source (Request.QueryString, Request.Form, route parameters) to sink (ExecuteReader, FromSqlRaw, ExecuteSqlRaw)
- Identify string concatenation or interpolation in SQL query construction
- Replace concatenated queries with parameterized alternatives using `SqlCommand.Parameters.AddWithValue()`
- For Entity Framework, replace `FromSqlRaw` with `FromSqlInterpolated` or LINQ queries
- Verify all user-supplied values are passed as parameters, not embedded in query strings
- Test the fix to ensure functionality and confirm scan results are resolved

## Safe Pattern

```csharp
// Parameterized query with ADO.NET
string query = "SELECT * FROM Users WHERE Username = @username AND Status = @status";
using (SqlCommand cmd = new SqlCommand(query, connection))
{
    cmd.Parameters.AddWithValue("@username", userInput);
    cmd.Parameters.AddWithValue("@status", statusValue);
    using (SqlDataReader reader = cmd.ExecuteReader())
    {
        // Process results
    }
}
```
