# CWE-326: Inadequate Encryption Strength - Java

## LLM Guidance

Inadequate Encryption Strength in Java occurs when weak cryptographic algorithms (DES, RC4, MD5), insufficient key sizes (<128-bit for symmetric, <2048-bit for RSA), or deprecated ciphers are used, leaving data vulnerable to brute-force and cryptanalytic attacks. The core fix is to use strong, modern algorithms (AES-256, RSA-2048+, SHA-256+) with proper key generation from secure random sources via Java's JCA/JCE framework.

## Key Principles

- Use AES with 256-bit keys for symmetric encryption and RSA with minimum 2048-bit keys for asymmetric encryption
- Generate cryptographic keys using `SecureRandom` with proper entropy, never hardcode or derive from weak sources
- Specify complete cipher transformations including mode and padding (e.g., "AES/GCM/NoPadding") to avoid insecure defaults
- Use authenticated encryption modes (GCM, CCM) that provide both confidentiality and integrity protection
- Regularly update to latest JDK versions to benefit from security patches and modern algorithm support

## Remediation Steps

- Replace DES, 3DES, RC4, Blowfish with AES-256; replace MD5, SHA-1 with SHA-256 or SHA-512
- Update `KeyGenerator.getInstance()` calls to specify key sizes - 256 for AES, 2048+ for RSA
- Change cipher initialization to use explicit modes - prefer "AES/GCM/NoPadding" over "AES"
- Replace `new Random()` or `Math.random()` with `SecureRandom` for all cryptographic operations
- Review and update key storage mechanisms to use Java KeyStore with strong passwords
- Add cipher strength validation in security configuration or startup checks

## Safe Pattern

```java
import javax.crypto.*;
import javax.crypto.spec.*;
import java.security.SecureRandom;

// Generate strong AES-256 key
KeyGenerator keyGen = KeyGenerator.getInstance("AES");
keyGen.init(256, new SecureRandom());
SecretKey key = keyGen.generateKey();

// Use authenticated encryption with GCM mode
Cipher cipher = Cipher.getInstance("AES/GCM/NoPadding");
byte[] iv = new byte[12];
new SecureRandom().nextBytes(iv);
GCMParameterSpec spec = new GCMParameterSpec(128, iv);
cipher.init(Cipher.ENCRYPT_MODE, key, spec);
byte[] encrypted = cipher.doFinal(plaintext);
```
