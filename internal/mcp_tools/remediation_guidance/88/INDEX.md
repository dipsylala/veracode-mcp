# CWE-88: Argument Injection

## LLM Guidance

Argument Injection occurs when untrusted user input is used to construct command-line arguments, function parameters, or system calls, allowing attackers to inject malicious arguments and alter program behavior. Innocuous installed executables (LOLBins) can be subverted through command-line arguments to perform code execution or filesystem manipulation.

Remediation Strategy:

- Never allow untrusted input to be parsed as command options or flags
- Use flag terminators (`--`) to separate options from user-controlled arguments
- Prefer parameterized APIs and safe libraries over string concatenation
- Validate input rigorously using allowlists of permitted values
- Avoid invoking system commands when safer alternatives exist

Remediation Steps:

- Trace the data path: identify where untrusted data enters (source), how it's used in argument construction, and where it reaches command execution (sink)
- Review scan results for string concatenation or direct use of user input in command arguments
- Locate all command execution functions (`system()`, `exec()`, `subprocess`, `Runtime.exec()`)
- Implement input validation with strict allowlists for permitted characters and values
- Use parameterized command execution methods that separate arguments from the command
- Place user input after `--` flag terminator to prevent interpretation as options
