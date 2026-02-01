# CWE-780: Use of RSA Without OAEP - Java

## LLM Guidance

Using RSA encryption without OAEP (Optimal Asymmetric Encryption Padding) enables padding oracle attacks, chosen ciphertext attacks, and message malleability. In Java, this commonly occurs when using `Cipher.getInstance("RSA")` without specifying the padding mode, which defaults to the insecure PKCS#1 v1.5 padding. Always use `Cipher.getInstance("RSA/ECB/OAEPWithSHA-256AndMGF1Padding")` with explicit OAEP parameters.

## Remediation Strategy

- Always specify the complete cipher transformation string including OAEP padding mode
- Use SHA-256 or stronger hash functions for OAEP (avoid SHA-1)
- Configure OAEPParameterSpec explicitly with MGF1 and appropriate hash algorithm
- Use RSA key sizes of 2048 bits minimum (4096 recommended for sensitive data)

## Remediation Steps

- Replace `Cipher.getInstance("RSA")` with `Cipher.getInstance("RSA/ECB/OAEPWithSHA-256AndMGF1Padding")`
- Create OAEPParameterSpec with SHA-256, MGF1ParameterSpec.SHA256, and PSource.PSpecified.DEFAULT
- Initialize cipher with the OAEP parameter spec using `cipher.init()` with AlgorithmParameterSpec
- Generate RSA keys with minimum 2048-bit key size using KeyPairGenerator
- Test encryption/decryption with sample data to verify proper OAEP implementation

## Minimal Safe Pattern

```java
import javax.crypto.Cipher;
import java.security.spec.OAEPParameterSpec;
import java.security.spec.MGF1ParameterSpec;
import javax.crypto.spec.PSource;

// Secure RSA encryption with OAEP
Cipher cipher = Cipher.getInstance("RSA/ECB/OAEPWithSHA-256AndMGF1Padding");
OAEPParameterSpec oaepParams = new OAEPParameterSpec(
    "SHA-256", "MGF1", MGF1ParameterSpec.SHA256, PSource.PSpecified.DEFAULT
);
cipher.init(Cipher.ENCRYPT_MODE, publicKey, oaepParams);
byte[] ciphertext = cipher.doFinal(plaintext);
```
