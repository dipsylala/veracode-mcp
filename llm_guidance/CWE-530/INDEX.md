# CWE-530: Information Exposure Through Source Code

## LLM Guidance

Information exposure through source code occurs when applications unintentionally expose source files, version control directories (.git, .svn), configuration files (.env), or artifacts (backup files like .bak, .old, .swp) to unauthorized users. This reveals business logic, credentials, algorithms, and security controls. The core fix is to prevent deployment of development artifacts to production and configure web servers to deny access to sensitive file types.

## Key Principles

- Never deploy version control directories or development artifacts to production web servers
- Configure web servers to explicitly deny access to sensitive file extensions and hidden files
- Remove backup files and temporary files from web-accessible directories before deployment
- Use build processes that exclude source files and only deploy compiled/minified production assets
- Implement automated scanning to detect exposed sensitive files in production

## Remediation Steps

- Delete all version control directories (.git, .svn, .hg) from web roots using find commands or deployment scripts
- Configure Apache with `<DirectoryMatch>` and `<FilesMatch>` rules to block .git directories and sensitive extensions (.env, .bak, .old, .swp)
- Configure Nginx location blocks to deny access to files starting with dots and sensitive extensions
- Remove commented-out code containing credentials or sensitive logic before deployment
- Use .gitignore and deployment tools to prevent accidental inclusion of config files
- Scan production servers regularly for exposed .git directories, .env files, and backup files using automated tools
