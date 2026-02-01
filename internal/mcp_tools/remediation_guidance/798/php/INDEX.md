# CWE-798: Hard-coded Credentials - PHP

## LLM Guidance

Hard-coded credentials (passwords, API keys, database credentials, encryption keys) in PHP code or configuration files create critical security vulnerabilities. The core fix is to externalize all secrets using environment variables, secure configuration files outside version control, or dedicated secrets managers. Use `getenv()`, `$_ENV`, vlucas/phpdotenv for development, or cloud secrets managers for production.

## Remediation Strategy

- Never commit credentials to version control (use .gitignore for .env files)
- Store secrets in environment variables or external configuration
- Use different credentials per environment (dev/staging/production)
- Implement secrets rotation and access controls
- Validate all credentials come from secure external sources

## Remediation Steps

- Identify all hard-coded credentials in code and configuration files
- Create environment variables or use a secrets manager for each credential
- Replace hard-coded values with `getenv()` or `$_ENV` calls
- Add .env files to .gitignore and remove any committed secrets from git history
- Test credential loading in all environments
- Document the secrets management approach for the team

## Minimal Safe Pattern

```php
// Load from environment variable
$dbPassword = getenv('DB_PASSWORD');
$apiKey = $_ENV['API_KEY'];

// Database connection
$conn = new PDO(
    "mysql:host=" . getenv('DB_HOST'),
    getenv('DB_USER'),
    getenv('DB_PASSWORD')
);

// Never do this:
// $password = "hardcoded123";
```
