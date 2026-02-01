# CWE-114: Process Control - C

## LLM Guidance

In C, CWE-114 occurs when loading shared libraries (`dlopen()`, `LoadLibrary()`) or executing processes (`exec*()`, `CreateProcess()`) without proper validation. Attackers exploit this through DLL hijacking, LD_PRELOAD attacks, and path manipulation.

**Primary Defense:** Use absolute paths, validate all inputs, disable unsafe search paths, and set secure file permissions.

## Remediation Strategy

- Use absolute paths for all library loads and process executions
- Validate and sanitize all external inputs before use in library/process calls
- Disable current directory library search on Windows (`SetDllDirectory("")`)
- Set restrictive file permissions (755 for executables, owned by trusted user)
- Verify library signatures and integrity before loading

## Remediation Steps

- Replace all relative paths with absolute, validated paths
- Implement whitelist validation for library/executable names
- Use `secure_getenv()` on Linux or sanitize environment variables
- Call `SetDllDirectory("")` before any `LoadLibrary()` on Windows
- Set `LD_LIBRARY_PATH` restrictions and use RPATH with `$ORIGIN` carefully
- Verify file ownership and permissions before loading

## Minimal Safe Pattern

```c
#include <dlfcn.h>
#include <limits.h>

// Secure library loading with absolute path
int load_library_safe(const char *lib_name) {
    char abs_path[PATH_MAX];
    const char *safe_dir = "/usr/lib/myapp";
    
    // Whitelist validation
    if (strstr(lib_name, "..") || strchr(lib_name, '/'))
        return -1;
    
    snprintf(abs_path, sizeof(abs_path), "%s/%s", safe_dir, lib_name);
    void *handle = dlopen(abs_path, RTLD_NOW | RTLD_LOCAL);
    return handle ? 0 : -1;
}
```
