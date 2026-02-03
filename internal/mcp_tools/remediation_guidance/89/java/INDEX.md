# CWE-89: SQL Injection - Java

## LLM Guidance

SQL Injection occurs when untrusted data is incorporated into SQL queries without proper validation or parameterization, allowing attackers to manipulate queries to bypass authentication, access unauthorized data, or modify records. Use `PreparedStatement` with parameterized queries, JPA/Hibernate named parameters, or query builder frameworks that automatically parameterize values.

## Key Principles

- Always use parameterized queries or prepared statements-never concatenate user input into SQL strings
- Prefer ORM frameworks (JPA/Hibernate) with bound parameters over raw JDBC
- Validate and sanitize input as a secondary defence layer
- Apply least privilege principles to database accounts
- Use stored procedures with parameterized inputs where appropriate

## Remediation Steps

- Locate - Identify the source (user input entry points like `request.getParameter()`, `@RequestParam`) and sink (SQL execution like `executeQuery()`, `createNativeQuery()`)
- Trace data flow - Check if string concatenation or `String.format()` is used to build SQL queries
- Replace concatenation - Convert string-based queries to `PreparedStatement` with `?` placeholders or JPA named parameters (`-paramName`)
- Bind parameters - Use `setString()`, `setInt()` or similar methods to bind user input to placeholders
- Test - Verify the fix handles special characters and injection attempts correctly
- Review - Ensure all similar patterns in the codebase are addressed

## Safe Pattern

```java
// SAFE: PreparedStatement with parameters
String sql = "SELECT * FROM users WHERE username = ? AND status = ?";
PreparedStatement stmt = connection.prepareStatement(sql);
stmt.setString(1, userInput);
stmt.setString(2, statusFilter);
ResultSet rs = stmt.executeQuery();

// SAFE: JPA named parameters
String jpql = "SELECT u FROM User u WHERE u.username = :username";
TypedQuery<User> query = em.createQuery(jpql, User.class);
query.setParameter("username", userInput);
List<User> results = query.getResultList();
```
