# CWE-798: Hard-coded Credentials - Java

## LLM Guidance

Hard-coded credentials in source code create severe security risks as they are exposed in version control, decompiled bytecode, and configuration files. Attackers gaining access to the codebase can extract these credentials to compromise systems. Store credentials externally using environment variables, secrets management services (AWS Secrets Manager, Azure Key Vault), or encrypted configuration files with restricted access.

## Key Principles

- Never embed passwords, API keys, tokens, or secrets directly in source code or properties files committed to version control
- Use environment variables or system properties for runtime credential injection
- Leverage cloud-native secrets managers (AWS Secrets Manager, Azure Key Vault, HashiCorp Vault) for production environments
- Implement proper access controls and encryption for configuration files containing sensitive data
- Rotate credentials regularly and revoke any previously hard-coded credentials immediately

## Remediation Steps

- Identify all hard-coded credentials using static analysis tools or code review
- Move credentials to environment variables or a secrets management service
- Update code to retrieve credentials at runtime from external sources
- Add credential files to .gitignore to prevent accidental commits
- Remove credential history from version control using tools like git-filter-repo
- Rotate all exposed credentials immediately after removal

## Safe Pattern

```java
// Retrieve credentials from environment variables at runtime
public class DatabaseConnection {
    private static final String DB_URL = System.getenv("DB_URL");
    private static final String DB_USER = System.getenv("DB_USER");
    private static final String DB_PASSWORD = System.getenv("DB_PASSWORD");
    
    public Connection connect() throws SQLException {
        if (DB_PASSWORD == null) {
            throw new IllegalStateException("DB_PASSWORD not set");
        }
        return DriverManager.getConnection(DB_URL, DB_USER, DB_PASSWORD);
    }
}
```
