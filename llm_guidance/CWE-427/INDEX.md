# CWE-427: Uncontrolled Search Path Element

## LLM Guidance

Uncontrolled Search Path vulnerabilities occur when applications unsafely modify search paths (PATH, LD_LIBRARY_PATH, etc.), allowing attackers to inject malicious directories and control which executables or libraries are loaded. The core fix is to use absolute paths for all critical resources and never let external input influence search path behavior.

## Key Principles

- Use absolute paths for all executables and libraries—avoid relying on search path resolution
- Never modify global search paths (PATH, LD_LIBRARY_PATH, PYTHONPATH) with untrusted data
- If search paths must be modified, validate and sanitize all path elements against an allowlist
- Set minimal, explicit search paths with only trusted directories
- Load resources directly by full path rather than depending on loader search behavior

## Remediation Steps

- Review flaw details to identify unsafe PATH modifications - `os.environ['PATH'] +=`, `setenv("PATH")`, `System.setProperty("java.library.path")`
- Trace the source of path elements—determine if they come from user input, configuration files, or external data
- Replace search-path-dependent code with absolute paths to executables and libraries
- If PATH modification is unavoidable, validate path elements against an allowlist of trusted directories
- Remove or sanitize any user-controllable input before appending to search paths
- Test with restricted environments to ensure malicious path injection cannot load unintended resources
