# CWE-798: Hard-coded Credentials - C\#

## LLM Guidance

Hard-coded credentials (passwords, API keys, connection strings, encryption keys) in C# source code create critical security vulnerabilities by exposing secrets in version control and compiled assemblies. Never embed credentials directly in code or configuration files. Use environment variables, .NET User Secrets for development, Azure Key Vault for production, or secure `IConfiguration` providers to externalize and protect sensitive data.

## Key Principles

- Externalize all credentials to secure storage outside source code and version control
- Use .NET User Secrets during development and Key Vault or environment variables in production
- Implement least-privilege access with managed identities where possible
- Rotate credentials regularly and audit all secret access

## Remediation Steps

- Identify all hard-coded credentials using code scanning tools or manual review
- Move secrets to User Secrets (`dotnet user-secrets set`) for local development
- Configure Azure Key Vault or environment variables for production environments
- Update code to retrieve credentials via `IConfiguration` or Key Vault client
- Remove hard-coded values and scrub from version control history
- Implement secret scanning in CI/CD pipelines to prevent future violations

## Safe Pattern

```csharp
// appsettings.json - no secrets here
public class Startup {
    public void ConfigureServices(IServiceCollection services) {
        var config = new ConfigurationBuilder()
            .AddUserSecrets<Startup>()      // Development
            .AddEnvironmentVariables()       // Production
            .Build();
        
        var apiKey = config["ApiKey"];
        var connStr = config["ConnectionStrings:Database"];
        services.AddDbContext<AppDbContext>(options =>
            options.UseSqlServer(connStr));
    }
}
```
