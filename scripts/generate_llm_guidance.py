"""
Generate LLM Guidance - Non-destructive CWE Guidance Processor

This script processes CWE security guidance files from ./guidance/ and generates
enhanced LLM-assisted versions in ./llm_guidance/, preserving the original files.

Features:
- Two-mode operation: LLM-first (via GitHub Copilot CLI) with deterministic fallback
- Categorizes CWEs by vulnerability type (SQL, XSS, command injection, etc.)
- Adds language-specific safe code patterns (Python, Java, C#, JS, PHP, Ruby, Go, etc.)
- Normalizes structure and removes deprecated sections
- Tracks progress in LLM_GUIDANCE_SPEC.md with checkboxes
- Logs all changes to LLM_REVIEW_LOG.md
- Saves LLM audit logs to llm_audit/
- Skips existing files (unless --force used)

Usage:
    python scripts/generate_llm_guidance.py              # Process all (skip existing)
    python scripts/generate_llm_guidance.py 78 80        # Process only CWE-78 and CWE-80
    python scripts/generate_llm_guidance.py CWE-89       # Process only CWE-89
    python scripts/generate_llm_guidance.py CWE-94/python  # Process only CWE-94/python/INDEX.md
    python scripts/generate_llm_guidance.py --force      # Re-process all files

Configuration:
- RW_LLM=0 to disable LLM mode (use deterministic transforms only)
- RW_COPILOT_CMD to override the Copilot command (default: 'copilot')

Output:
- ./llm_guidance/ - Enhanced guidance files (mirrors ./guidance/ structure)
- ./llm_guidance/LLM_GUIDANCE_SPEC.md - Progress tracking
- ./llm_guidance/LLM_REVIEW_LOG.md - Change log
- ./llm_guidance/llm_audit/ - LLM interaction audit logs
"""
from pathlib import Path
import re
from datetime import datetime
import os
import subprocess
import shutil
import shlex
import sys

root = Path('guidance')
output_root = Path('llm_guidance')
spec_path = output_root / 'LLM_GUIDANCE_SPEC.md'
log_path = output_root / 'LLM_REVIEW_LOG.md'
llm_audit_dir = output_root / 'llm_audit'

# LLM usage (default on). Set RW_LLM=0 to disable.
LLM_ENABLED = os.getenv('RW_LLM', '1') != '0'
# Command used to invoke Copilot. Override with RW_COPILOT_CMD.
# Default to copilot (will use -i interactive mode).
COPILOT_CMD = os.getenv('RW_COPILOT_CMD', 'copilot')
COPILOT_BIN = COPILOT_CMD.split()[0]
COPILOT_BIN_PATH = Path(COPILOT_BIN)

# --- Categorization ---
patterns = [
    ('sql', re.compile(r'\bSQL Injection\b', re.I)),
    ('xss', re.compile(r'\bCross-Site Scripting|\bXSS\b', re.I)),
    ('cmd', re.compile(r'\bCommand Injection\b|\bOS Command Injection\b', re.I)),
    ('ldap', re.compile(r'\bLDAP Injection\b', re.I)),
    ('xpath', re.compile(r'\bXPath Injection\b', re.I)),
    ('xxe', re.compile(r'\bXML External Entity\b|\bXXE\b', re.I)),
    ('deser', re.compile(r'\bDeserialization\b', re.I)),
    ('path', re.compile(r'\bPath Traversal\b|\bDirectory Traversal\b', re.I)),
    ('csrf', re.compile(r'\bCSRF\b|\bCross-Site Request Forgery\b', re.I)),
    ('ssrf', re.compile(r'\bSSRF\b|\bServer-Side Request Forgery\b', re.I)),
    ('secrets', re.compile(r'Hard[- ]coded Credentials|Hard[- ]coded Password|Hard[- ]coded', re.I)),
    ('redirect', re.compile(r'\bOpen Redirect\b', re.I)),
    ('upload', re.compile(r'\bFile Upload\b', re.I)),
    ('crypto', re.compile(r'Cryptograph|Encryption|Crypto|Cipher|Key|Hash|Random|PRNG', re.I)),
    ('auth', re.compile(r'Authentication', re.I)),
    ('access', re.compile(r'Authorization|Access Control|IDOR|Insecure Direct Object', re.I)),
    ('logging', re.compile(r'Log|Logging', re.I)),
]

# --- Remediation Strategy steps ---
STRATEGY = {
    'sql': [
        'Use parameterized queries or prepared statements for all database access.',
        'Avoid string concatenation for SQL; use ORM parameter binding.',
        'Validate untrusted data with strict allowlists where feasible.',
        'Apply least privilege to database accounts.',
    ],
    'xss': [
        'Encode untrusted data at the output sink (HTML, attribute, JS, URL).',
        'Avoid raw HTML rendering unless content is sanitized with a safe allowlist.',
        'Prefer framework auto-escaping features.',
        'Apply CSP as defense in depth when applicable.',
    ],
    'cmd': [
        'Replace OS commands with native APIs.',
        'If a process is required, disable shell execution and pass args as a list.',
        'Allowlist untrusted data used in arguments or paths.',
        'Run with least privilege.',
    ],
    'ldap': [
        'Use LDAP parameterization or safe filter escaping.',
        'Avoid constructing LDAP filters with string concatenation.',
        'Allowlist untrusted data when possible.',
        'Apply least privilege to directory access.',
    ],
    'xpath': [
        'Use safe XPath APIs or parameterization where available.',
        'Avoid string concatenation for XPath expressions.',
        'Allowlist or strictly validate untrusted data used in queries.',
        'Limit query scope to necessary nodes.',
    ],
    'xxe': [
        'Disable DTDs and external entity resolution in XML parsers.',
        'Use hardened XML libraries where available.',
        'Reject or sanitize untrusted XML inputs.',
        'Apply least privilege to any file/network access in parsers.',
    ],
    'deser': [
        'Avoid unsafe native serialization formats; prefer JSON with schemas.',
        'Allowlist types/classes when deserializing.',
        'Reject untrusted serialized data by default.',
        'Apply least privilege to deserialization endpoints.',
    ],
    'path': [
        'Canonicalize paths and enforce a fixed base directory.',
        'Allowlist filenames or identifiers used to select files.',
        'Avoid direct filesystem access from untrusted data.',
        'Apply least privilege on file permissions.',
    ],
    'csrf': [
        'Require anti-CSRF tokens on state-changing requests.',
        'Enforce same-site cookie settings and use POST/PUT for mutations.',
        'Validate origin or referer when appropriate.',
        'Avoid sensitive actions via GET.',
    ],
    'ssrf': [
        'Allowlist destinations (host, scheme, port) for outbound requests.',
        'Block access to internal IP ranges and metadata endpoints.',
        'Use fixed URLs when possible instead of untrusted data.',
        'Apply network egress controls.',
    ],
    'secrets': [
        'Remove hard-coded secrets from source and config.',
        'Load secrets from environment variables or a secrets manager.',
        'Rotate and revoke exposed credentials.',
        'Restrict secret access by least privilege.',
    ],
    'redirect': [
        'Allowlist redirect targets or use relative paths only.',
        'Avoid reflecting untrusted URLs into redirect responses.',
        'Normalize and validate redirect destinations.',
        'Use server-side routing where possible.',
    ],
    'upload': [
        'Validate file type and content; reject dangerous formats.',
        'Store uploads outside the web root and use randomized names.',
        'Enforce size limits and scan content if required.',
        'Apply least privilege to upload handling.',
    ],
    'crypto': [
        'Use modern algorithms and safe defaults (e.g., AES-GCM, Argon2/bcrypt).',
        'Avoid custom crypto or weak algorithms/modes.',
        'Use secure random number generators.',
        'Store keys in a secrets manager.',
    ],
    'auth': [
        'Perform server-side authorization checks for each request.',
        'Do not trust client-supplied role/permission claims.',
        'Use centralized access control logic.',
        'Apply least privilege by default.',
    ],
    'access': [
        'Perform server-side authorization checks for each request.',
        'Do not trust client-supplied identifiers or permissions.',
        'Use centralized access control logic.',
        'Apply least privilege by default.',
    ],
    'logging': [
        'Sanitize or encode untrusted data before logging.',
        'Use structured logging to avoid injection into log formats.',
        'Avoid logging secrets or sensitive data.',
        'Apply log access controls and retention limits.',
    ],
}

DEFAULT_STRATEGY = [
    'Replace unsafe sinks with safe native APIs or library functions.',
    'Apply the primary safe pattern for this CWE.',
    'Validate untrusted data with strict allowlists and type checks.',
    'Apply least privilege and safe defaults.',
]

# --- Minimal safe patterns (best-practice per category) ---
SNIP = {
    'csharp': {
        'sql': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - parameterized query\nusing var cmd = new SqlCommand(\"SELECT * FROM Users WHERE Id = @id\", conn);\ncmd.Parameters.Add(\"@id\", SqlDbType.Int).Value = id;\n```\n""",
        'xss': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - output encoding\nvar safe = HtmlEncoder.Default.Encode(value);\n```\n""",
        'cmd': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - no shell, argument list\nvar psi = new ProcessStartInfo(cmd) { UseShellExecute = false };\npsi.ArgumentList.Add(arg);\nProcess.Start(psi);\n```\n""",
        'path': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - canonicalize and enforce base path\nvar full = Path.GetFullPath(Path.Combine(baseDir, name));\nif (!full.StartsWith(baseDir, StringComparison.Ordinal)) throw new SecurityException();\n```\n""",
        'csrf': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - anti-forgery token\n[ValidateAntiForgeryToken]\npublic IActionResult Post(Model m) { ... }\n```\n""",
        'ssrf': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - allowlist host and scheme\nvar uri = new Uri(url);\nif (!allowedHosts.Contains(uri.Host)) throw new SecurityException();\n```\n""",
        'secrets': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - load secrets from environment/secret store\nvar apiKey = Environment.GetEnvironmentVariable(\"API_KEY\");\n```\n""",
        'deser': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - safe JSON deserialization\nvar obj = JsonSerializer.Deserialize<MyType>(json);\n```\n""",
        'crypto': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - strong password hashing\nvar hash = BCrypt.Net.BCrypt.HashPassword(password);\n```\n""",
        'auth': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - server-side authorization check\nif (!User.HasClaim(\"perm\", \"resource:read\")) return Forbid();\n```\n""",
        'access': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - server-side authorization check\nif (!User.HasClaim(\"perm\", \"resource:read\")) return Forbid();\n```\n""",
    },
    'java': {
        'sql': """### Minimal Safe Pattern\n\n```java\n// SECURE - parameterized query\nPreparedStatement ps = conn.prepareStatement(\"SELECT * FROM users WHERE id = ?\");\nps.setInt(1, id);\n```\n""",
        'xss': """### Minimal Safe Pattern\n\n```java\n// SECURE - output encoding\nString safe = Encode.forHtml(value);\n```\n""",
        'cmd': """### Minimal Safe Pattern\n\n```java\n// SECURE - no shell, explicit args\nnew ProcessBuilder(cmd, arg).start();\n```\n""",
        'path': """### Minimal Safe Pattern\n\n```java\n// SECURE - canonicalize and enforce base path\nPath full = base.resolve(name).normalize();\nif (!full.startsWith(base)) throw new SecurityException();\n```\n""",
        'csrf': """### Minimal Safe Pattern\n\n```java\n// SECURE - enable CSRF protection\nhttp.csrf();\n```\n""",
        'ssrf': """### Minimal Safe Pattern\n\n```java\n// SECURE - allowlist host and scheme\nURI uri = new URI(url);\nif (!allowedHosts.contains(uri.getHost())) throw new SecurityException();\n```\n""",
        'secrets': """### Minimal Safe Pattern\n\n```java\n// SECURE - load secrets from environment/secret store\nString apiKey = System.getenv(\"API_KEY\");\n```\n""",
        'deser': """### Minimal Safe Pattern\n\n```java\n// SECURE - safe JSON deserialization\nMyType obj = new ObjectMapper().readValue(json, MyType.class);\n```\n""",
        'crypto': """### Minimal Safe Pattern\n\n```java\n// SECURE - strong password hashing\nString hash = BCrypt.hashpw(password, BCrypt.gensalt());\n```\n""",
        'auth': """### Minimal Safe Pattern\n\n```java\n// SECURE - server-side authorization check\nif (!authz.canRead(user, resource)) throw new ForbiddenException();\n```\n""",
        'access': """### Minimal Safe Pattern\n\n```java\n// SECURE - server-side authorization check\nif (!authz.canRead(user, resource)) throw new ForbiddenException();\n```\n""",
        'ldap': """### Minimal Safe Pattern\n\n```java\n// SECURE - escape LDAP filter values\nString safe = LdapEncoder.filterEncode(value);\n```\n""",
        'xxe': """### Minimal Safe Pattern\n\n```java\n// SECURE - disable DTDs/external entities\nfactory.setFeature(\"http://apache.org/xml/features/disallow-doctype-decl\", true);\n```\n""",
    },
    'python': {
        'sql': """### Minimal Safe Pattern\n\n```python\n# SECURE - parameterized query\ncursor.execute(\"SELECT * FROM users WHERE id = %s\", (user_id,))\n```\n""",
        'xss': """### Minimal Safe Pattern\n\n```python\n# SECURE - output encoding\nsafe = html.escape(value, quote=True)\n```\n""",
        'cmd': """### Minimal Safe Pattern\n\n```python\n# SECURE - no shell, args list\nsubprocess.run([cmd, arg], check=True)\n```\n""",
        'path': """### Minimal Safe Pattern\n\n```python\n# SECURE - canonicalize and enforce base path\nfull = (base / name).resolve()\nif not str(full).startswith(str(base)): raise ValueError()\n```\n""",
        'csrf': """### Minimal Safe Pattern\n\n```python\n# SECURE - enable CSRF protection\n@csrf_protect\ndef post(request): ...\n```\n""",
        'ssrf': """### Minimal Safe Pattern\n\n```python\n# SECURE - allowlist host and scheme\nuri = urlparse(url)\nif uri.hostname not in allowed_hosts: raise ValueError()\n```\n""",
        'secrets': """### Minimal Safe Pattern\n\n```python\n# SECURE - load secrets from environment/secret store\napi_key = os.environ[\"API_KEY\"]\n```\n""",
        'deser': """### Minimal Safe Pattern\n\n```python\n# SECURE - safe JSON deserialization\nobj = json.loads(body)\n```\n""",
        'crypto': """### Minimal Safe Pattern\n\n```python\n# SECURE - strong password hashing\nhash = bcrypt.hashpw(password.encode(), bcrypt.gensalt())\n```\n""",
        'auth': """### Minimal Safe Pattern\n\n```python\n# SECURE - server-side authorization check\nif not user.can(\"read\", resource): raise PermissionError()\n```\n""",
        'access': """### Minimal Safe Pattern\n\n```python\n# SECURE - server-side authorization check\nif not user.can(\"read\", resource): raise PermissionError()\n```\n""",
        'ldap': """### Minimal Safe Pattern\n\n```python\n# SECURE - escape LDAP filter values\nsafe = escape_filter_chars(value)\n```\n""",
        'xxe': """### Minimal Safe Pattern\n\n```python\n# SECURE - defused XML parsing\nroot = defusedxml.ElementTree.fromstring(xml)\n```\n""",
    },
    'javascript': {
        'sql': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - parameterized query\nawait db.query('SELECT * FROM users WHERE id = ?', [id]);\n```\n""",
        'xss': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - output encoding\nel.textContent = value;\n```\n""",
        'cmd': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - no shell, args list\nspawn(cmd, [arg], { shell: false });\n```\n""",
        'path': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - canonicalize and enforce base path\nconst full = path.resolve(base, name);\nif (!full.startsWith(base)) throw new Error('invalid');\n```\n""",
        'csrf': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - send CSRF token\nfetch(url, { method: 'POST', headers: { 'X-CSRF-Token': token } });\n```\n""",
        'ssrf': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - allowlist host and scheme\nconst u = new URL(url);\nif (!allowed.has(u.hostname)) throw new Error('invalid');\n```\n""",
        'secrets': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - load secrets from environment/secret store\nconst apiKey = process.env.API_KEY;\n```\n""",
        'deser': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - safe JSON parsing\nconst obj = JSON.parse(body);\n```\n""",
        'crypto': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - strong password hashing\nconst hash = await bcrypt.hash(password, 12);\n```\n""",
        'auth': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - server-side authorization check\nif (!can(user, 'read', resource)) throw new Error('forbidden');\n```\n""",
        'access': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - server-side authorization check\nif (!can(user, 'read', resource)) throw new Error('forbidden');\n```\n""",
        'ldap': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - escape LDAP filter values\nconst safe = escapeLDAPFilter(value);\n```\n""",
        'xxe': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - disable external entities\nconst parser = new XMLParser({ processEntities: false });\n```\n""",
        'redirect': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - allowlist redirect targets\nconst u = new URL(next, base);\nif (u.origin !== base) throw new Error('invalid');\n```\n""",
    },
    'php': {
        'sql': """### Minimal Safe Pattern\n\n```php\n// SECURE - parameterized query\n$stmt = $pdo->prepare('SELECT * FROM users WHERE id = ?');\n$stmt->execute([$id]);\n```\n""",
        'xss': """### Minimal Safe Pattern\n\n```php\n// SECURE - output encoding\n$safe = htmlspecialchars($value, ENT_QUOTES, 'UTF-8');\n```\n""",
        'cmd': """### Minimal Safe Pattern\n\n```php\n// SECURE - no shell, fixed executable and args\n$proc = new Symfony\\Component\\Process\\Process([$cmd, $arg]);\n$proc->run();\n```\n""",
        'path': """### Minimal Safe Pattern\n\n```php\n// SECURE - canonicalize and enforce base path\n$full = realpath($base . DIRECTORY_SEPARATOR . $name);\nif (strpos($full, $base) !== 0) { throw new Exception('invalid'); }\n```\n""",
        'csrf': """### Minimal Safe Pattern\n\n```php\n// SECURE - verify CSRF token\nif (!hash_equals($_SESSION['csrf'], $_POST['csrf'])) { http_response_code(403); }\n```\n""",
        'ssrf': """### Minimal Safe Pattern\n\n```php\n// SECURE - allowlist host and scheme\n$u = parse_url($url);\nif (!in_array($u['host'], $allowed, true)) { throw new Exception('invalid'); }\n```\n""",
        'secrets': """### Minimal Safe Pattern\n\n```php\n// SECURE - load secrets from environment/secret store\n$apiKey = getenv('API_KEY');\n```\n""",
        'deser': """### Minimal Safe Pattern\n\n```php\n// SECURE - safe JSON parsing\n$obj = json_decode($json, true, 512, JSON_THROW_ON_ERROR);\n```\n""",
        'crypto': """### Minimal Safe Pattern\n\n```php\n// SECURE - strong password hashing\n$hash = password_hash($password, PASSWORD_ARGON2ID);\n```\n""",
        'auth': """### Minimal Safe Pattern\n\n```php\n// SECURE - server-side authorization check\nif (!$user->can('read', $resource)) { http_response_code(403); }\n```\n""",
        'access': """### Minimal Safe Pattern\n\n```php\n// SECURE - server-side authorization check\nif (!$user->can('read', $resource)) { http_response_code(403); }\n```\n""",
    },
    'ruby': {
        'sql': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - parameterized query\nUser.where(id: id)\n```\n""",
        'xss': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - output encoding\nsafe = ERB::Util.html_escape(value)\n```\n""",
        'cmd': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - no shell, args list\nOpen3.capture3(cmd, arg)\n```\n""",
        'path': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - canonicalize and enforce base path\nfull = File.expand_path(name, base)\nraise 'invalid' unless full.start_with?(base)\n```\n""",
        'csrf': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - enable CSRF protection\nprotect_from_forgery with: :exception\n```\n""",
        'ssrf': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - allowlist host and scheme\nuri = URI.parse(url)\nraise 'invalid' unless allowed.include?(uri.host)\n```\n""",
        'secrets': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - load secrets from environment/secret store\napi_key = ENV.fetch('API_KEY')\n```\n""",
        'deser': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - safe JSON parsing\nobj = JSON.parse(body)\n```\n""",
        'crypto': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - strong password hashing\nhash = BCrypt::Password.create(password)\n```\n""",
        'auth': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - server-side authorization check\nraise 'forbidden' unless can?(:read, resource)\n```\n""",
        'access': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - server-side authorization check\nraise 'forbidden' unless can?(:read, resource)\n```\n""",
        'xxe': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - disable external entities\nNokogiri::XML(xml) { |cfg| cfg.nonet }\n```\n""",
    },
    'go': {
        'sql': """### Minimal Safe Pattern\n\n```go\n// SECURE - parameterized query\nrows, err := db.Query(\"SELECT * FROM users WHERE id = ?\", id)\n```\n""",
        'xss': """### Minimal Safe Pattern\n\n```go\n// SECURE - output encoding\nsafe := template.HTMLEscapeString(value)\n```\n""",
        'cmd': """### Minimal Safe Pattern\n\n```go\n// SECURE - no shell, args list\ncmd := exec.Command(command, arg)\n```\n""",
        'path': """### Minimal Safe Pattern\n\n```go\n// SECURE - canonicalize and enforce base path\nfull := filepath.Clean(filepath.Join(base, name))\nif !strings.HasPrefix(full, base) { return err }\n```\n""",
        'csrf': """### Minimal Safe Pattern\n\n```go\n// SECURE - enable CSRF protection\nhandler = csrf.Protect(key)(handler)\n```\n""",
        'ssrf': """### Minimal Safe Pattern\n\n```go\n// SECURE - allowlist host and scheme\nu, _ := url.Parse(target)\nif !allowed[u.Hostname()] { return err }\n```\n""",
        'secrets': """### Minimal Safe Pattern\n\n```go\n// SECURE - load secrets from environment/secret store\napiKey := os.Getenv(\"API_KEY\")\n```\n""",
        'deser': """### Minimal Safe Pattern\n\n```go\n// SECURE - safe JSON parsing\nerr := json.Unmarshal(body, &obj)\n```\n""",
        'crypto': """### Minimal Safe Pattern\n\n```go\n// SECURE - strong password hashing\nhash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)\n```\n""",
        'auth': """### Minimal Safe Pattern\n\n```go\n// SECURE - server-side authorization check\nif !can(user, resource) { return err }\n```\n""",
        'access': """### Minimal Safe Pattern\n\n```go\n// SECURE - server-side authorization check\nif !can(user, resource) { return err }\n```\n""",
    },
    'perl': {
        'sql': """### Minimal Safe Pattern\n\n```perl\n# SECURE - parameterized query\nmy $sth = $dbh->prepare('SELECT * FROM users WHERE id = ?');\n$sth->execute($id);\n```\n""",
        'xss': """### Minimal Safe Pattern\n\n```perl\n# SECURE - output encoding\nmy $safe = encode_entities($value);\n```\n""",
        'secrets': """### Minimal Safe Pattern\n\n```perl\n# SECURE - load secrets from environment/secret store\nmy $api_key = $ENV{'API_KEY'};\n```\n""",
        'auth': """### Minimal Safe Pattern\n\n```perl\n# SECURE - server-side authorization check\ndie 'forbidden' unless can($user, $resource);\n```\n""",
        'access': """### Minimal Safe Pattern\n\n```perl\n# SECURE - server-side authorization check\ndie 'forbidden' unless can($user, $resource);\n```\n""",
    },
    'c': {
        'sql': """### Minimal Safe Pattern\n\n```c\n// SECURE - parameterized query (SQLite)\nsqlite3_prepare_v2(db, \"SELECT * FROM users WHERE id = ?\", -1, &stmt, 0);\nsqlite3_bind_int(stmt, 1, id);\n```\n""",
        'path': """### Minimal Safe Pattern\n\n```c\n// SECURE - canonicalize and enforce base path\nrealpath(path, full);\nif (strncmp(full, base, strlen(base)) != 0) return ERROR;\n```\n""",
    },
    'cpp': {
        'sql': """### Minimal Safe Pattern\n\n```cpp\n// SECURE - parameterized query\nauto stmt = conn.prepare(\"SELECT * FROM users WHERE id = ?\");\nstmt.bind(1, id);\n```\n""",
        'path': """### Minimal Safe Pattern\n\n```cpp\n// SECURE - canonicalize and enforce base path\nauto full = fs::weakly_canonical(base / name);\nif (full.string().rfind(base.string(), 0) != 0) throw std::runtime_error(\"invalid\");\n```\n""",
    },
}

FALLBACK = {
    'csharp': """### Minimal Safe Pattern\n\n```csharp\n// SECURE - allowlist untrusted data before use\nif (!Regex.IsMatch(value, @\"^[A-Za-z0-9._-]+$\")) throw new ValidationException();\n```\n""",
    'java': """### Minimal Safe Pattern\n\n```java\n// SECURE - allowlist untrusted data before use\nif (!value.matches(\"^[A-Za-z0-9._-]+$\")) throw new IllegalArgumentException(\"invalid\");\n```\n""",
    'python': """### Minimal Safe Pattern\n\n```python\n# SECURE - allowlist untrusted data before use\nif not re.fullmatch(r\"[A-Za-z0-9._-]+\", value):\n    raise ValueError(\"invalid\")\n```\n""",
    'javascript': """### Minimal Safe Pattern\n\n```javascript\n// SECURE - allowlist untrusted data before use\nif (!/^[A-Za-z0-9._-]+$/.test(value)) throw new Error('invalid');\n```\n""",
    'php': """### Minimal Safe Pattern\n\n```php\n// SECURE - allowlist untrusted data before use\nif (!preg_match('/^[A-Za-z0-9._-]+$/', $value)) { throw new InvalidArgumentException(); }\n```\n""",
    'ruby': """### Minimal Safe Pattern\n\n```ruby\n# SECURE - allowlist untrusted data before use\nraise 'invalid' unless value.match?(/\\A[a-zA-Z0-9._-]+\\z/)\n```\n""",
    'go': """### Minimal Safe Pattern\n\n```go\n// SECURE - allowlist untrusted data before use\nif !regexp.MustCompile(`^[A-Za-z0-9._-]+$`).MatchString(value) { return err }\n```\n""",
    'perl': """### Minimal Safe Pattern\n\n```perl\n# SECURE - allowlist untrusted data before use\ndie \"invalid\" unless $value =~ /^[A-Za-z0-9._-]+$/;\n```\n""",
    'c': """### Minimal Safe Pattern\n\n```c\n// SECURE - allowlist untrusted data before use\nfor (const char *p = value; *p; ++p) { if (!isalnum((unsigned char)*p)) return ERROR; }\n```\n""",
    'cpp': """### Minimal Safe Pattern\n\n```cpp\n// SECURE - allowlist untrusted data before use\nfor (char c : value) { if (!std::isalnum((unsigned char)c)) throw std::invalid_argument(\"invalid\"); }\n```\n""",
}

REMOVE_TITLES = {
    'overview',
    'data path focus',
    'testing & verification',
    'references',
    'common pitfalls to avoid',
    'common vulnerable patterns',
}

section_heading_re = re.compile(r'^(#{2,6})\s+(.+?)\s*$', re.M)


def categorize(title):
    for cat, rx in patterns:
        if rx.search(title):
            return cat
    return None


def build_strategy(title):
    cat = categorize(title)
    steps = STRATEGY.get(cat, DEFAULT_STRATEGY)
    return "## Remediation Strategy\n\n" + "\n".join([f"{i+1}. {s}" for i, s in enumerate(steps)]) + "\n"


def build_remediation_steps():
    return (
        "## Remediation Steps\n\n"
        "- Identify the sink and confirm the data path from untrusted data\n"
        "- Apply the primary safe pattern for this CWE\n"
        "- Add allowlist validation or encoding where required\n"
        "- Verify behavior with normal and boundary cases\n"
    )


def remove_sections(text):
    lines = text.splitlines(keepends=True)
    out = []
    i = 0
    while i < len(lines):
        m = section_heading_re.match(lines[i])
        if m:
            level = len(m.group(1))
            title = m.group(2).strip().lower()
            if title in REMOVE_TITLES:
                i += 1
                while i < len(lines):
                    m2 = section_heading_re.match(lines[i])
                    if m2 and len(m2.group(1)) <= level:
                        break
                    i += 1
                continue
        out.append(lines[i])
        i += 1
    return ''.join(out)


def ensure_minimal_safe_pattern(text, lang, title):
    cat = categorize(title)
    snippet = SNIP.get(lang, {}).get(cat) or FALLBACK.get(lang)
    if not snippet:
        return text
    if '### Minimal Safe Pattern' in text:
        return re.sub(r'### Minimal Safe Pattern\n\n```[\s\S]*?```\n', lambda m: snippet, text, count=1)
    m = re.search(r'^##\s+Secure Patterns\s*$', text, re.M)
    if not m:
        return text
    start = m.end()
    next_m = re.search(r'^##\s+[^#].*$', text[start:], re.M)
    end = start + (next_m.start() if next_m else len(text[start:]))
    section = text[start:end].rstrip()
    new_section = section + "\n\n" + snippet + "\n"
    return text[:start] + new_section + text[end:]


def normalize_file(input_path, output_path):
    text = input_path.read_text(encoding='utf-8')
    lines = text.splitlines()
    title_line = lines[0] if lines else ''
    title = title_line.lstrip('#').strip()

    # LLM-first pass (default). Fall back to deterministic transform if LLM fails.
    if LLM_ENABLED:
        llm_text = llm_update_file(input_path, text)
        if llm_text:
            output_path.parent.mkdir(parents=True, exist_ok=True)
            output_path.write_text(llm_text.rstrip('\n') + '\n', encoding='utf-8')
            return 'llm'

    text = remove_sections(text)

    if '## Remediation Strategy' in text:
        text = re.sub(r'## Remediation Strategy[\s\S]*?(?=^##\s+|\Z)', build_strategy(title) + "\n", text, flags=re.M)
    else:
        text += "\n" + build_strategy(title) + "\n"

    rel = input_path.relative_to(root)
    if len(rel.parts) == 2:
        if '## Remediation Steps' in text:
            text = re.sub(r'## Remediation Steps[\s\S]*?(?=^##\s+|\Z)', build_remediation_steps() + "\n", text, flags=re.M)
        else:
            text += "\n" + build_remediation_steps() + "\n"

    if len(rel.parts) >= 3:
        lang = rel.parts[1]
        text = ensure_minimal_safe_pattern(text, lang, title)

    text = text.rstrip('\n') + '\n'
    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.write_text(text, encoding='utf-8')
    return 'deterministic'


def update_progress(spec_text, file_paths):
    if '## Progress' not in spec_text:
        spec_text = spec_text.rstrip('\n') + '\n\n## Progress\n\n' + '\n'.join([f"- [ ] {p}" for p in file_paths]) + '\n'
        return spec_text

    lines = spec_text.splitlines()
    listed = {line[6:] for line in lines if line.startswith('- [')}
    missing = [p for p in file_paths if p not in listed]
    if missing:
        insert_idx = len(lines)
        for i, line in enumerate(lines):
            if line.startswith('## Progress'):
                insert_idx = i + 1
                break
        for p in missing:
            lines.insert(insert_idx, f"- [ ] {p}")
            insert_idx += 1
        spec_text = "\n".join(lines) + '\n'
    return spec_text


def mark_done(spec_text, file_path):
    return spec_text.replace(f"- [ ] {file_path}", f"- [x] {file_path}")


def log_review_entry(file_path, mode):
    ts = datetime.utcnow().isoformat() + 'Z'
    entry = f"- {ts} {file_path} ({mode})"
    if log_path.exists():
        current = log_path.read_text(encoding='utf-8').rstrip('\n')
        log_path.write_text(current + "\n" + entry + "\n", encoding='utf-8')
    else:
        log_path.write_text("# LLM Review Log\n\n" + entry + "\n", encoding='utf-8')

def extract_section(text, section_name):
    """Extract a specific ## section from markdown text."""
    lines = text.split('\n')
    in_section = False
    section_lines = []
    
    for line in lines:
        if line.startswith('## ') and section_name.lower() in line.lower():
            in_section = True
            continue
        elif line.startswith('## ') and in_section:
            # Hit next section, stop
            break
        elif in_section:
            section_lines.append(line)
    
    return '\n'.join(section_lines).strip()


def llm_update_file(path, text):
    """
    Generate concise LLM-focused guidance from human-focused documentation.
    
    Expected output format (200-400 tokens):
    
    # CWE-XXX: Title
    
    ## LLM Guidance
    [2-3 sentences: what the vuln is + core fix principle]
    
    ## Remediation Strategy
    - [3-5 bullet points with key principles]
    
    ## Remediation Steps
    - [4-6 bullet points with actionable steps]
    
    ## Minimal Safe Pattern (language files only)
    ```lang
    [10-15 lines of secure code]
    ```
    """
    # If copilot isn't available, skip. For a .ps1, try to run even if the path can't be resolved here.
    if COPILOT_BIN_PATH.suffix.lower() == '.ps1':
        copilot_exists = True
    elif COPILOT_BIN_PATH.is_absolute():
        copilot_exists = COPILOT_BIN_PATH.exists()
    else:
        copilot_exists = bool(shutil.which(COPILOT_BIN))
    if not copilot_exists:
        llm_audit_dir.mkdir(parents=True, exist_ok=True)
        rel = path.relative_to(root)
        audit_path = llm_audit_dir / (rel.as_posix().replace('/', '__') + '.txt')
        audit_path.write_text(
            "SKIPPED: copilot not found on PATH for command "
            f"'{COPILOT_CMD}'\n\n"
            "DIAGNOSTICS:\n"
            f"- COPILOT_CMD: {COPILOT_CMD}\n"
            f"- COPILOT_BIN: {COPILOT_BIN}\n"
            f"- COPILOT_BIN_PATH: {COPILOT_BIN_PATH}\n"
            f"- COPILOT_BIN_PATH exists: {COPILOT_BIN_PATH.exists()}\n"
            f"- COPILOT_BIN_PATH absolute: {COPILOT_BIN_PATH.is_absolute()}\n"
            f"- shutil.which(COPILOT_BIN): {shutil.which(COPILOT_BIN)}\n"
            f"- PATH: {os.environ.get('PATH','')}\n",
            encoding='utf-8',
        )
        return None

    spec = spec_path.read_text(encoding='utf-8') if spec_path.exists() else ''
    rel = path.relative_to(root)
    
    # Extract key sections from source
    lines = text.split('\n')
    title = lines[0] if lines else ''
    overview = extract_section(text, 'Overview')
    strategy = extract_section(text, 'Remediation Strategy')
    steps = extract_section(text, 'Remediation Steps')
    
    # Check if this is a language-specific file (has parent dirs beyond CWE-XXX)
    is_language_file = len(rel.parts) >= 3
    
    prompt = f"""
Create ULTRA-CONCISE LLM guidance. Maximum 400 tokens total.

{title}

## LLM Guidance
Write 2-3 sentences: what is the vulnerability + core fix principle
Source: {overview[:400]}

## Remediation Strategy  
Extract 3-5 key principles as SHORT bullets (one line each, no sub-bullets)
Source: {strategy[:600]}

## Remediation Steps
Extract 4-6 actionable steps as SHORT bullets (one line each)
Source: {steps[:600]}

{"## Minimal Safe Pattern\nProvide EXACTLY ONE code example showing the primary safe pattern.\nMaximum 15 lines of code. No multiple examples. No VULNERABLE code.\nShow the bare minimum to demonstrate the fix." if is_language_file else ""}

CRITICAL RULES:
- NO markdown code fences around the entire output (````markdown)
- ONE code example maximum (not 3, not 5, just 1)
- Keep code to 10-15 lines maximum
- Use bullets, not paragraphs
- Total output must be under 400 tokens (~1600 chars)

Return ONLY the markdown content. No diagnostics, no wrapper fences.
"""
    cmd_parts = shlex.split(COPILOT_CMD)
    # Interactive mode: use -i flag with prompt
    try:
        if COPILOT_BIN_PATH.suffix.lower() == '.ps1':
            cmd_display = [
                "powershell",
                "-NoProfile",
                "-ExecutionPolicy",
                "Bypass",
                "-File",
                COPILOT_BIN,
                *cmd_parts[1:],
                "-i",
                "<PROMPT>",
            ]
            result = subprocess.run(
                [
                    "powershell",
                    "-NoProfile",
                    "-ExecutionPolicy",
                    "Bypass",
                    "-File",
                    COPILOT_BIN,
                    *cmd_parts[1:],
                    "-i",
                    prompt,
                ],
                text=True,
                encoding="utf-8",
                errors="replace",
                capture_output=True,
                timeout=300,
            )
        else:
            cmd_display = [*cmd_parts, "-i", "<PROMPT>"]
            result = subprocess.run(
                [*cmd_parts, "-i", prompt],
                text=True,
                encoding="utf-8",
                errors="replace",
                capture_output=True,
                timeout=300,
            )
    except Exception:
        return None

    llm_audit_dir.mkdir(parents=True, exist_ok=True)
    audit_path = llm_audit_dir / (rel.as_posix().replace('/', '__') + '.txt')
    audit_path.write_text(
        f"""COMMAND:
{cmd_display}

RETURN CODE:
{result.returncode}

PROMPT:
{prompt}

STDOUT:
{result.stdout}

STDERR:
{result.stderr}
""",
        encoding='utf-8',
    )

    if result.returncode != 0:
        return None
    output = result.stdout.strip()
    if not output:
        return None
    
    # Extract markdown content from Copilot CLI output
    # Copilot CLI outputs diagnostic info (● lines) before the actual content
    lines = output.split('\n')
    
    # Check if content is wrapped in markdown code fence
    if lines and lines[0].strip().startswith('```'):
        # Find the closing fence
        start_idx = 1
        end_idx = len(lines)
        for i in range(1, len(lines)):
            if lines[i].strip() == '```':
                end_idx = i
                break
        markdown_content = '\n'.join(lines[start_idx:end_idx]).strip()
    else:
        # Find the start of actual markdown content (after all diagnostic lines)
        content_start = 0
        for i, line in enumerate(lines):
            # Skip lines that start with ● (bullet point diagnostics)
            # Skip lines that are indented with └ or other tree characters
            stripped = line.lstrip()
            if not stripped.startswith('●') and not stripped.startswith('└') and not line.startswith('  └'):
                # Check if this looks like markdown (starts with # or has content)
                if stripped.startswith('#') or (stripped and not stripped.startswith('Now I')):
                    content_start = i
                    break
        
        # Extract only the markdown content
        markdown_lines = lines[content_start:]
        
        # Remove any trailing "Now I'll" messages or similar
        while markdown_lines and (
            markdown_lines[-1].strip().startswith('Now I') or 
            not markdown_lines[-1].strip()
        ):
            markdown_lines.pop()
        
        markdown_content = '\n'.join(markdown_lines).strip()
    
    # Validate we got markdown (should start with # heading)
    if not markdown_content or not markdown_content.lstrip().startswith('#'):
        return None
    
    return markdown_content


def main():
    """
    Non-destructively process guidance files:
    - Reads INDEX.md files from ./guidance/
    - Applies LLM or deterministic transformations
    - Writes output to ./llm_guidance/ (preserving directory structure)
    - Original ./guidance/ files remain unchanged
    
    Usage:
        python scripts/generate_llm_guidance.py                    # Process all CWEs
        python scripts/generate_llm_guidance.py 78 80              # Process only CWE-78 and CWE-80
        python scripts/generate_llm_guidance.py CWE-94/python      # Process only CWE-94/python/INDEX.md
        python scripts/generate_llm_guidance.py --force CWE-114/c  # Re-process specific file
    """

    # Parse command-line arguments
    file_filter = None
    cwe_filter = None
    force_reprocess = '--force' in sys.argv or '-f' in sys.argv
    
    args = [arg for arg in sys.argv[1:] if arg not in ('--force', '-f')]
    
    if args:
        # Check if arguments contain paths (have /)
        path_args = [arg for arg in args if '/' in arg or '\\' in arg]
        cwe_args = [arg for arg in args if '/' not in arg and '\\' not in arg]
        
        if path_args:
            # File-level filtering
            file_filter = set()
            for arg in path_args:
                # Normalize path separators
                normalized = arg.replace('\\', '/').strip('/')
                file_filter.add(normalized)
            print(f"Processing specific files: {', '.join(sorted(file_filter))}")
        
        if cwe_args:
            # CWE-level filtering
            cwe_filter = set()
            for arg in cwe_args:
                # Remove 'CWE-' prefix if provided
                cwe_num = arg.replace('CWE-', '').replace('cwe-', '')
                cwe_filter.add(f'CWE-{cwe_num}')
            print(f"Processing CWEs: {', '.join(sorted(cwe_filter))}")
    
    if force_reprocess:
        print("Force mode: will re-process existing files")

    files = []
    for p in root.rglob('INDEX.md'):
        if '_dynamic' in p.parts:
            continue
        
        rel_path = p.relative_to(root)
        rel_path_str = str(rel_path).replace('\\', '/')
        
        # Filter by specific files if provided
        if file_filter:
            # Check if this file matches any of the file filters
            matches = False
            for filter_path in file_filter:
                # Match if the file path starts with or equals the filter
                if rel_path_str.startswith(filter_path) or rel_path_str == filter_path + '/INDEX.md':
                    matches = True
                    break
            if not matches:
                continue
        
        # Filter by CWE numbers if provided (only if no file filter)
        elif cwe_filter:
            if not any(cwe in p.parts for cwe in cwe_filter):
                continue
        
        files.append(rel_path)

    files.sort(key=lambda p: str(p))

    # Create output directory
    output_root.mkdir(parents=True, exist_ok=True)

    spec = spec_path.read_text(encoding='utf-8') if spec_path.exists() else '# LLM Guidance Spec\n'
    spec = update_progress(spec, [str(p) for p in files])

    processed = 0
    skipped = 0
    
    for rel_path in files:
        input_path = root / rel_path
        output_path = output_root / rel_path
        
        # Skip if output already exists (unless force mode)
        if not force_reprocess and output_path.exists():
            print(f"Skipped {rel_path} (already exists)")
            skipped += 1
            continue
        
        mode = normalize_file(input_path, output_path)
        spec = mark_done(spec, str(rel_path))
        log_review_entry(str(rel_path), mode)
        print(f"Processed {rel_path} -> llm_guidance/{rel_path}")
        processed += 1

    spec = spec.rstrip('\n') + '\n'
    spec_path.write_text(spec, encoding='utf-8')
    
    if file_filter:
        print(f"\nProcessed {processed} files, skipped {skipped} for paths: {', '.join(sorted(file_filter))}")
    elif cwe_filter:
        print(f"\nProcessed {processed} files, skipped {skipped} for CWEs: {', '.join(sorted(cwe_filter))}")
    else:
        print(f"\nProcessed {processed} files, skipped {skipped} from guidance/ to llm_guidance/")
    print(f"Updated {spec_path}")


if __name__ == '__main__':
    main()
