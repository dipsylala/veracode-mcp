# CWE-22: Path Traversal

## LLM Guidance

Path Traversal occurs when applications use user-supplied input to construct file paths without proper validation, allowing attackers to access files outside the intended directory using sequences like `../`. The core fix is to never allow untrusted input to directly control filesystem paths; always canonicalize paths and enforce containment within an allowlisted root directory.

## Key Principles

- Reject direct user input in paths: Never use untrusted data directly as file paths or path components
- Canonicalize before validation: Resolve symlinks and relative paths (`.`, `..`) to absolute form before security checks
- Enforce allowlist containment: Verify canonicalized paths stay within permitted root directories
- Use indirect references: Map user input to internal IDs, then lookup actual paths from a secure database or allowlist
- Validate after canonicalization: String filtering alone fails; validate the resolved absolute path

## Remediation Steps

- Trace data flow - Identify where file path data enters (user input, APIs, databases), how it's constructed (string concatenation with `/` or `\`), and where it reaches file operations (`open()`, `readFile()`, etc.)
- Implement indirect references - Replace direct path usage with user-provided IDs/names that map to system-controlled paths via database lookup or allowlist
- Canonicalize paths - Convert all paths to absolute canonical form, resolving symlinks and relative references before any validation
- Enforce root containment - After canonicalization, verify the path starts with an approved root directory prefix
- Reject suspicious patterns - Block or sanitize inputs containing `..`, absolute paths, or encoded traversal sequences before path construction
- Apply defense in depth - Combine indirect references, canonicalization, allowlist validation, and filesystem permissions as layered controls
