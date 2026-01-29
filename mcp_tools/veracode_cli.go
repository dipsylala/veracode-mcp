package mcp_tools

// VeracodeExitCodeInfo contains information about a Veracode CLI exit code
type VeracodeExitCodeInfo struct {
	ExitCode  int
	Icon      string
	Message   string
	NextSteps string
	IsWarning bool // True for exit codes that should be treated as warnings rather than errors (3, 4)
}

// InterpretVeracodeExitCode interprets a Veracode CLI exit code and returns structured information
func InterpretVeracodeExitCode(exitCode int) VeracodeExitCodeInfo {
	info := VeracodeExitCodeInfo{
		ExitCode: exitCode,
	}

	switch exitCode {
	case 0:
		info.Icon = "✓"
		info.Message = "Command completed successfully"
		info.NextSteps = "Review the output and proceed with next steps"
		info.IsWarning = false

	case 1:
		info.Icon = "❌"
		info.Message = "Generic error occurred"
		info.NextSteps = "Next steps:\n- Review the error output above\n- Check Veracode CLI documentation\n- Verify input parameters and retry"
		info.IsWarning = false

	case 2:
		info.Icon = "❌"
		info.Message = "Parser error - check command input and user permissions"
		info.NextSteps = "Next steps:\n- Verify paths are accessible\n- Check file and directory permissions\n- Ensure all required parameters are valid"
		info.IsWarning = false

	case 3:
		info.Icon = "⚠"
		info.Message = "Command completed, but application did not pass policy"
		info.NextSteps = "Next steps:\n- Review the results\n- Check policy violations\n- Address policy issues before submission"
		info.IsWarning = true

	case 4:
		info.Icon = "⚠"
		info.Message = "Command completed with warnings or no artifacts found"
		info.NextSteps = "Next steps:\n- Check if --strict flag was used (causes warnings to fail)\n- Verify source files are present\n- Review build output for warnings\n- Check for compilation errors"
		info.IsWarning = true

	case 125:
		info.Icon = "❌"
		info.Message = "Out of memory error"
		info.NextSteps = "Next steps:\n- Reduce workspace size or exclude large files\n- Increase available system memory\n- Contact system administrator"
		info.IsWarning = false

	case 126:
		info.Icon = "❌"
		info.Message = "Command failed - check local system configuration"
		info.NextSteps = "Next steps:\n- Verify Veracode CLI is properly installed\n- Check system dependencies\n- Review system logs for configuration issues"
		info.IsWarning = false

	case 127:
		info.Icon = "❌"
		info.Message = "Command not found - check PATH configuration"
		info.NextSteps = "Next steps:\n- Verify 'veracode' command is in system PATH\n- Reinstall Veracode CLI if necessary\n- Check command syntax is correct"
		info.IsWarning = false

	case 128:
		info.Icon = "❌"
		info.Message = "Invalid argument - retry the command"
		info.NextSteps = "Next steps:\n- Review command arguments\n- Check for typos in parameters\n- Verify argument format matches CLI expectations"
		info.IsWarning = false

	case 130:
		info.Icon = "❌"
		info.Message = "Command terminated by user"
		info.NextSteps = "Next steps:\n- Retry the operation\n- Allow sufficient time for completion\n- Check if any prompts need to be answered"
		info.IsWarning = false

	default:
		info.Icon = "❌"
		info.Message = "Command failed with unexpected exit code"
		info.NextSteps = "Next steps:\n- Review error output\n- Consult Veracode CLI documentation\n- Contact Veracode support if issue persists"
		info.IsWarning = false
	}

	return info
}
