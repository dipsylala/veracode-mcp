# CWE-89: SQL Injection - PHP

## LLM Guidance

SQL Injection occurs when untrusted data is incorporated into SQL queries without proper parameterization, allowing attackers to manipulate queries to bypass authentication, access unauthorized data, or execute administrative operations.

**Primary Defence:** Use prepared statements with PDO or MySQLi, or modern ORMs like Laravel Eloquent that use prepared statements internally. Escape functions alone are insufficient and must not be relied upon.

## Key Principles

- Always use prepared statements with bound parameters for all SQL queries containing user input
- Employ parameterized queries through PDO or MySQLi, never string concatenation
- Use ORMs (Laravel Eloquent, Doctrine) that handle parameterization automatically
- Apply input validation as a secondary defence layer, but never as the primary protection
- Reject escape functions (mysql_real_escape_string) as the sole defence mechanism

## Remediation Steps

- Locate the sink (SQL execution point) and source (user input) in the data flow report
- Trace how data flows from source to sink, identifying any concatenation or interpolation
- Replace concatenated queries with prepared statements using PDO or MySQLi
- Bind all user-supplied values as parameters, never interpolate them into query strings
- Validate input types and formats as secondary defence (e.g., verify IDs are numeric)
- Test the fix by attempting injection attacks and reviewing query logs

## Safe Pattern

```php
// PDO prepared statement with named parameters
$stmt = $pdo->prepare("SELECT * FROM users WHERE username = :username AND status = :status");
$stmt->execute([
    ':username' => $_POST['username'],
    ':status' => $_POST['status']
]);
$results = $stmt->fetchAll();

// MySQLi prepared statement with positional parameters
$stmt = $mysqli->prepare("SELECT * FROM users WHERE username = ? AND status = ?");
$stmt->bind_param("ss", $_POST['username'], $_POST['status']);
$stmt->execute();
$results = $stmt->get_result()->fetch_all();
```
