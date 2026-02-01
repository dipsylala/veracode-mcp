# CWE-548: Information Exposure Through Directory Listing

## LLM Guidance

Information exposure through directory listing occurs when web servers or applications allow users to browse directory contents without an index file, revealing file names, directory structure, and potentially sensitive files. This vulnerability results from misconfigured web servers (missing index.html, disabled auto-index protection), frameworks exposing static directories, or intentional browsing features without proper access controls. Core fix: Disable directory browsing on all web servers and require explicit index files for publicly accessible directories.

## Key Principles

- Disable directory browsing/auto-indexing on all web servers and application frameworks
- Require explicit index files (index.html, index.php) for all publicly accessible directories
- Configure proper access controls on static file directories and framework configurations
- Implement defense in depth with both server-level and application-level protections
- Audit regularly for exposed directories and missing index files

## Remediation Steps

- Disable server directory listing - Apache - `Options -Indexes` in .htaccess/httpd.conf; Nginx - `autoindex off;` in nginx.conf; IIS - `<directoryBrowse enabled="false" />` in web.config
- Create index files - Place index.html or index.php in all directories, including empty placeholder files for directories without content
- Review framework configurations - Disable or restrict static file serving in application frameworks (Django, Express, Spring Boot)
- Audit existing directories - Scan all web-accessible paths for directory listing exposure
- Implement automated checks - Add CI/CD tests to detect missing index files or enabled directory browsing
- Test and validate - Use security scanners and manual testing to confirm directory listings are disabled across all paths
