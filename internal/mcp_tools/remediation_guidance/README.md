# LLM Guidance Structure

This directory contains remediation guidance for security vulnerabilities organized by CWE (Common Weakness Enumeration) ID. The guidance is structured to provide both general vulnerability information and language-specific remediation instructions.

## Directory Structure

```text
remediation_guidance/
├── README.md (this file)
├── {CWE_ID}/
│   ├── INDEX.md (general CWE guidance)
│   ├── java/
│   │   └── INDEX.md (Java-specific guidance)
│   ├── python/
│   │   └── INDEX.md (Python-specific guidance)
│   ├── javascript/
│   │   └── INDEX.md (JavaScript-specific guidance)
│   ├── csharp/
│   │   └── INDEX.md (C#-specific guidance)
│   ├── php/
│   │   └── INDEX.md (PHP-specific guidance)
│   └── ... (other languages as needed)
```

---

## General CWE INDEX.md Structure

Location: `{CWE_ID}/INDEX.md`

Each general CWE INDEX.md file provides language-agnostic guidance and follows this structure:

### Template

```markdown
# CWE-{ID}: {Vulnerability Name}

## LLM Guidance

{2-4 sentence overview explaining what the vulnerability is, how it occurs, and the core fix/approach. This should be actionable and vendor-neutral.}

## Key Principles

- {Principle 1: Primary Defence mechanism}
- {Principle 2: What NOT to do}
- {Principle 3: Data handling approach}
- {Principle 4: Defence-in-depth considerations}
- {Principle 5: Additional security controls}

## Remediation Steps

- {Step 1: Identify/locate the vulnerability - describe sources and sinks}
- {Step 2: Find the problematic pattern - describe what to look for}
- {Step 3: Replace with secure pattern - describe the fix approach}
- {Step 4: Implement the fix - describe technical implementation}
- {Step 5: Add secondary defences - describe validation/additional layers}
- {Step 6: Test - describe verification approach}
```

### Guidelines

- LLM Guidance: Write for an LLM that will help developers fix code. Be direct and actionable
- Key Principles: List 4-6 fundamental security principles that apply to this vulnerability class
- Remediation Steps: Provide 4-8 sequential steps that guide from detection to verification
- Tone: Technical, concise, and prescriptive
- Length: Keep total content under 500 words for efficient LLM processing

### Example

See [89/INDEX.md](89/INDEX.md) for SQL Injection as a reference implementation.

---

## Language-Specific INDEX.md Structure

Location: `{CWE_ID}/{language}/INDEX.md`

Each language-specific INDEX.md file provides concrete, language-specific remediation guidance and code examples:

### Template

```markdown
# CWE-{ID}: {Vulnerability Name} - {Language}

## LLM Guidance

{2-4 sentence overview specific to this language: how the vulnerability manifests in this language, common frameworks/libraries involved, and the primary remediation approach using language-specific APIs/patterns.}

## Key Principles

- {Principle 1: Language-specific Primary Defence with API/framework names}
- {Principle 2: Preferred language patterns/frameworks}
- {Principle 3: Language-specific validation approaches}
- {Principle 4: Framework-specific best practices}
- {Principle 5: Language/ecosystem conventions}

## Remediation Steps

- Locate - {Identify source and sink with language-specific function names like `request.getParameter()`, `executeQuery()`}
- Trace data flow - {Describe language-specific patterns to look for like string concatenation operators, formatting functions}
- Replace {anti-pattern} - {Describe how to convert to secure pattern with specific API names and syntax}
- Bind parameters - {Show how to bind/pass parameters using language-specific methods}
- Test - {Describe testing approach with language-specific considerations}
- Review - {Additional review steps or patterns to check}

## Safe Pattern

```{language}
// SAFE: {Description of first safe pattern}
{code example showing the secure implementation}

// SAFE: {Description of second safe pattern (if applicable)}
{alternative secure implementation if relevant}
```

### Guidelines

- Language/Framework Specificity: Reference actual APIs, methods, classes, and packages
- Code Examples: Include runnable code showing SAFE patterns (and optionally UNSAFE for contrast)
- Function Names: Use actual function names from the language/framework (e.g., `PreparedStatement`, `setString()`, `createQuery()`)
- Syntax: Use correct language syntax with proper code fencing
- Imports: Include necessary import statements when they're not obvious
- Comments: Mark examples clearly as SAFE or UNSAFE
- Length: Keep under 800 words; focus on the most common/important patterns
- Version Awareness: Prefer modern language versions but note if older patterns are still common

### Example

See [89/java/INDEX.md](89/java/INDEX.md) for SQL Injection in Java as a reference implementation.

---

## Adding New CWE Guidance

1. Create directory: `mkdir {CWE_ID}`
2. Create general guidance: `{CWE_ID}/INDEX.md`
3. Create language directories as needed: `{CWE_ID}/{language}/`
4. Create language-specific guidance: `{CWE_ID}/{language}/INDEX.md`
5. Follow the templates above
6. Test that the guidance is helpful for LLM-assisted remediation

## Supported Languages

Current language directories in use:

- `java` - Java
- `python` - Python
- `javascript` - JavaScript/Node.js
- `csharp` - C#/.NET
- `php` - PHP
- `perl` - Perl
- (More languages can be added as needed)

## Usage

The remediation guidance tool loads these files automatically and provides them to LLMs when developers need help fixing security vulnerabilities identified by Veracode scans.
