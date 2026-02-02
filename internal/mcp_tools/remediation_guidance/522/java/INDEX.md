# CWE-522: Insufficiently Protected Credentials - Java

## LLM Guidance

Insufficiently Protected Credentials in Java occurs when passwords, API keys, or authentication tokens are hardcoded, stored in plaintext, weakly encrypted, or transmitted insecurely. The core fix is to externalize credentials to secure vaults (AWS Secrets Manager, HashiCorp Vault, Azure Key Vault), use environment variables or encrypted configuration files, and leverage Java's KeyStore or strong encryption (AES-256) when storage is necessary. Never log, hardcode, or commit credentials to version control.

## Key Principles

- Store credentials in external secret management systems (AWS Secrets Manager, Azure Key Vault, HashiCorp Vault)
- Use environment variables or encrypted configuration files with restricted file permissions
- Encrypt credentials at rest using AES-256 with Java KeyStore; never store plaintext passwords
- Transmit credentials only over TLS/SSL; use char[] for passwords in memory and clear immediately after use
- Implement credential rotation policies and audit access logs regularly

## Remediation Steps

- Remove hardcoded credentials from source code; scan with tools like git-secrets or TruffleHog
- Migrate credentials to AWS Secrets Manager, Azure Key Vault, or HashiCorp Vault
- Configure application to retrieve credentials at runtime from secret manager or environment variables
- Use JCEKS KeyStore with strong passwords for local encrypted credential storage if needed
- Replace String passwords with char[] and zero out arrays after authentication
- Enable TLS 1.3 for all credential transmission; never send credentials in URLs or logs

## Safe Pattern

```java
// Retrieve credentials from AWS Secrets Manager
import com.amazonaws.services.secretsmanager.*;
import com.amazonaws.services.secretsmanager.model.*;

AWSSecretsManager client = AWSSecretsManagerClientBuilder.standard()
    .withRegion("us-east-1").build();

GetSecretValueRequest request = new GetSecretValueRequest()
    .withSecretId("prod/db/password");
GetSecretValueResult result = client.getSecretValue(request);
String secret = result.getSecretString();

// Use credential immediately, then clear
char[] password = secret.toCharArray();
authenticateUser(username, password);
Arrays.fill(password, '\0'); // Clear sensitive data
```
