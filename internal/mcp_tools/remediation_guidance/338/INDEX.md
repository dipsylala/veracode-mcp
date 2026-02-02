# CWE-338: Use of Cryptographically Weak Pseudo-Random Number Generator (PRNG)

## LLM Guidance

Weak PRNG vulnerabilities occur when applications use cryptographically insecure random number generators (like `random`, `Math.random`, `java.util.Random`) for security-sensitive operations such as session tokens, cryptographic keys, or nonces. Attackers can predict or reproduce these values, compromising security. Fix: Replace weak PRNGs with cryptographically secure alternatives.

## Key Principles

- Always use cryptographically secure PRNGs for security-sensitive randomness
- Never use standard PRNGs (`random`, `Math.random`) for tokens, keys, salts, or nonces
- Never implement custom random number generators
- Avoid predictable seeding patterns (time-based, PID-based, static seeds)

## Remediation Steps

- Identify weak PRNG usage - Review flaw details for file/line number and trace data flow to determine if random values are used for security purposes
- Replace with secure alternatives -
  - Python - Use `secrets` module or `os.urandom()` instead of `random`
  - Java - Use `java.security.SecureRandom` instead of `java.util.Random`
  - JavaScript - Use `crypto.getRandomValues()` or `crypto.randomBytes()` instead of `Math.random()`
- Verify proper initialization - Ensure secure PRNGs are properly seeded by the OS
- Review all usages - Search codebase for other instances of weak PRNGs in security contexts
- Test the fix - Verify random values are unpredictable and non-reproducible across sessions
