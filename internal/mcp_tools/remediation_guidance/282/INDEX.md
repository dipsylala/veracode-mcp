# CWE-282: Improper Ownership Management

## LLM Guidance

Improper ownership management occurs when files, directories, or resources are created with incorrect ownership, allowing unauthorized modification, privilege escalation, or data tampering by unintended users. Core fix: Set explicit ownership on all resources and verify it—never assume safe defaults.

## Key Principles

- Set explicit ownership on all created resources using chown/chgrp
- Use least-privilege ownership (service users, not root) for application files
- Verify ownership after creation in deployment scripts
- Never create world-writable files or directories
- Apply correct group ownership to enable proper access control

## Remediation Steps

- Check file creation code for explicit ownership setting (chown commands)
- Review service files to find resources owned by root that should be service-owned
- Identify world-writable files (chmod 777) and restrict permissions
- Review sensitive configuration files for incorrect ownership
- Add chown commands immediately after file creation in deployment scripts
- Use install command with -o/-g flags instead of touch + chown
