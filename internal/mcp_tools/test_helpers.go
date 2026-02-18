package mcp_tools

// Test helper functions for backward compatibility with existing tests
// These allow tests to continue using the old New...Tool() pattern

// NewStaticFindingsTool creates a SimpleTool for static findings (test helper)
func NewStaticFindingsTool() ToolImplementation {
	return NewSimpleTool(StaticFindingsToolName, handleGetStaticFindings)
}

// NewDynamicFindingsTool creates a SimpleTool for dynamic findings (test helper)
func NewDynamicFindingsTool() ToolImplementation {
	return NewSimpleTool(DynamicFindingsToolName, handleGetDynamicFindings)
}

// NewScaFindingsTool creates a SimpleTool for SCA findings (test helper)
func NewScaFindingsTool() ToolImplementation {
	return NewSimpleTool(ScaFindingsToolName, handleGetScaFindings)
}

// NewPipelineScanTool creates a SimpleTool for pipeline scan (test helper)
func NewPipelineScanTool() ToolImplementation {
	return NewSimpleTool(PipelineScanToolName, handlePipelineScan)
}

// NewPipelineStatusTool creates a SimpleTool for pipeline status (test helper)
func NewPipelineStatusTool() ToolImplementation {
	return NewSimpleTool(PipelineStatusToolName, handlePipelineStatus)
}

// NewPipelineFindingsTool creates a SimpleTool for pipeline findings (test helper)
func NewPipelineFindingsTool() ToolImplementation {
	return NewSimpleTool(PipelineFindingsToolName, handlePipelineFindings)
}

// NewLocalSCAScanTool creates a SimpleTool for SCA scan (test helper)
func NewLocalSCAScanTool() ToolImplementation {
	return NewSimpleTool(LocalSCAScanToolName, handleLocalSCAScan)
}

// NewGetLocalSCAFindingsTool creates a SimpleTool for local SCA findings (test helper)
func NewGetLocalSCAFindingsTool() ToolImplementation {
	return NewSimpleTool(GetLocalSCAFindingsToolName, handleGetLocalSCAFindings)
}

// NewPackageWorkspaceTool creates a SimpleTool for package workspace (test helper)
func NewPackageWorkspaceTool() ToolImplementation {
	return NewSimpleTool(PackageWorkspaceToolName, handlePackageWorkspace)
}

// NewFindingDetailsTool creates a SimpleTool for finding details (test helper)
func NewFindingDetailsTool() ToolImplementation {
	return NewSimpleTool(FindingDetailsToolName, handleGetFindingDetails)
}

// NewAPIHealthTool creates a SimpleTool for API health (test helper)
func NewAPIHealthTool() ToolImplementation {
	return NewSimpleTool(APIHealthToolName, handleAPIHealth)
}

// NewRemediationGuidanceTool creates a SimpleTool for remediation guidance (test helper)
func NewRemediationGuidanceTool() ToolImplementation {
	return NewSimpleTool(RemediationGuidanceToolName, handleGetRemediationGuidance)
}
