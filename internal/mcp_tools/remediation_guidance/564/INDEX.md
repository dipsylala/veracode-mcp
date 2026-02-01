# CWE-564: SQL Injection: Hibernate

## LLM Guidance

Hibernate injection occurs when user input is concatenated into HQL (Hibernate Query Language) or native SQL queries without parameterization, enabling SQL injection attacks despite using an ORM framework. This allows attackers to steal data, gain unauthorized access, and manipulate databases. Never build Hibernate/ORM queries by concatenation; use parameter binding and avoid dynamic query fragments from untrusted input.

## Key Principles

- Use parameterized queries with named or positional parameters (`:param` or `?1`) instead of string concatenation
- Avoid `createNativeQuery()` when HQL alternatives exist; native SQL bypasses ORM protections
- Never concatenate user input into query strings, even for ORDER BY or IN clauses
- Use Criteria API or QueryDSL for dynamic queries requiring type safety
- Validate and whitelist user input for non-parameterizable elements like column names

## Remediation Steps

- Search codebase for `createQuery()`, `createNativeQuery()`, and `createSQLQuery()` with string concatenation (`+` operator or `String.format()`)
- Identify dynamic query construction patterns - ORDER BY clauses, IN clause lists, and WHERE conditions built with user input
- Replace concatenated queries with `setParameter()` binding for all user-controlled values
- For dynamic sorting, use whitelisted column names mapped to safe values rather than direct user input
- Refactor complex dynamic queries to use Criteria API or Specifications for type-safe construction
- Test all changes to verify query logic remains correct and parameters bind properly
