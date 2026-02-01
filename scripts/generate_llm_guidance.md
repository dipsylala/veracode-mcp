# generate_llm_guidance.py

Non-destructive CWE guidance processor that generates enhanced security guidance using LLM assistance.

## Purpose

Processes CWE security guidance files from `./guidance/` and generates enhanced versions in `./llm_guidance/` with:

- LLM-enhanced content (via GitHub Copilot CLI)
- Language-specific safe code patterns
- Normalized structure and remediation strategies
- Categorized vulnerability patterns

## Usage

```bash
# Process all CWEs (skips existing files in llm_guidance/)
python scripts/generate_llm_guidance.py

# Process specific CWE numbers only
python scripts/generate_llm_guidance.py 78 80
python scripts/generate_llm_guidance.py CWE-89 CWE-79

# Process specific files only (useful for regenerating long files)
python scripts/generate_llm_guidance.py CWE-94/python CWE-114/c
python scripts/generate_llm_guidance.py CWE-89/java CWE-89/python

# Force re-process all (ignore existing files)
python scripts/generate_llm_guidance.py --force

# Force re-process specific files
python scripts/generate_llm_guidance.py --force CWE-94/python CWE-114/c

# Deterministic mode only (no LLM)
RW_LLM=0 python scripts/generate_llm_guidance.py

# Specific CWEs with deterministic mode
RW_LLM=0 python scripts/generate_llm_guidance.py 78 80

# Custom Copilot command
RW_COPILOT_CMD="copilot-custom" python scripts/generate_llm_guidance.py
```

## Output

All generated content goes to `./llm_guidance/`:

- **CWE-*/INDEX.md** - Enhanced guidance files
- **LLM_GUIDANCE_SPEC.md** - Progress tracking with checkboxes
- **LLM_REVIEW_LOG.md** - Processing log with timestamps
- **llm_audit/** - Detailed LLM interaction audit logs

## Features

**Two-Mode Operation:**

1. **LLM Mode** (default): Uses GitHub Copilot CLI for intelligent enhancement
2. **Deterministic Mode**: Rule-based transformations if LLM unavailable

**Smart Processing:**

- Skips files that already exist in `./llm_guidance/` (use `--force` to override)
- Processes only specific CWEs when provided as arguments
- Shows progress with processed/skipped counts

**Vulnerability Categorization:**

- SQL Injection, XSS, Command Injection, LDAP, XPath, XXE
- Deserialization, Path Traversal, CSRF, SSRF
- Secrets, Redirects, File Upload, Crypto, Auth, Access Control, Logging

**Language Coverage:**

- Python, Java, C#, JavaScript, PHP
- Ruby, Go, Perl, C, C++

**Safe Patterns:**

- Parameterized queries for SQL
- Output encoding for XSS
- Safe subprocess execution for command injection
- Path canonicalization for traversal
- And many more category-specific patterns

## Non-Destructive

Original files in `./guidance/` are **never modified**. All output is written to `./llm_guidance/`.

## Environment Variables

- `RW_LLM` - Set to `0` to disable LLM mode (default: `1`)
- `RW_COPILOT_CMD` - Copilot CLI command (default: `copilot`)

## Review Workflow

1. Run script to generate enhanced guidance
2. Review output in `./llm_guidance/`
3. Compare with originals in `./guidance/`
4. Manually merge desired changes back to `./guidance/`

## Requirements

**For LLM mode:**

- GitHub Copilot CLI (`copilot` command available)

**For deterministic mode:**

- Python 3.7+ (no external dependencies)

## Related Tools

### find_long_files.py

Identifies markdown files in `llm_guidance/` that exceed the target token count.

#### Purpose

Analyzes all generated files and categorizes them by size:

- **Target:** 200-400 tokens (~800-1600 chars)
- **Warning:** Files 500-750 tokens (~2000-3000 chars)  
- **Error:** Files >750 tokens (~3000+ chars)

Helps identify files that may be too verbose for LLM consumption and need regeneration.

#### Usage

```bash
# Analyze all generated files
python scripts/find_long_files.py
```

#### Output

```
Analyzing 350 files in llm_guidance/

❌ 5 files are TOO LONG (>3000 chars, ~750 tokens):

  3842 chars (~ 960 tokens)  llm_guidance/CWE-89/python/INDEX.md
  3156 chars (~ 789 tokens)  llm_guidance/CWE-79/java/INDEX.md
  ...

⚠️  23 files are LONGER than target (2000-3000 chars, ~500-750 tokens):

  2456 chars (~ 614 tokens)  llm_guidance/CWE-78/php/INDEX.md
  ...

✅ 322 files within target size (<2000 chars, ~500 tokens)

Statistics:
  Average: 1456 chars (~364 tokens)
  Largest: 3842 chars (~960 tokens) - llm_guidance/CWE-89/python/INDEX.md
  Smallest: 892 chars (~223 tokens) - llm_guidance/CWE-103/INDEX.md
```

Then regenerate problematic files:

```bash
# Regenerate specific too-long files
python scripts/generate_llm_guidance.py --force CWE-94/python CWE-114/c CWE-114/python
```

**Workflow for fixing long files:**

```bash
# 1. Find files that are too long
python scripts/find_long_files.py

# 2. Copy the paths from output (e.g., CWE-94/python, CWE-114/c)

# 3. Regenerate only those specific files
python scripts/generate_llm_guidance.py --force CWE-94/python CWE-114/c

# 4. Verify they're now within target size
python scripts/find_long_files.py
```

#### Requirements

- Python 3.7+ (no external dependencies)

### Markdown Validation

Use `pymarkdown` to lint all generated files:

```bash
# Lint all guidance files
pymarkdown scan llm_guidance/

# Lint specific directory
pymarkdown scan llm_guidance/CWE-78/
```
