package api

import (
	"github.com/dipsylala/veracode-mcp/api/xml"
)

// convertXMLMitigationInfo converts xml.MitigationInfo to api.MitigationInfo wrapper type
func convertXMLMitigationInfo(xmlInfo *xml.MitigationInfo) *MitigationInfo {
	if xmlInfo == nil {
		return nil
	}

	issues := make([]MitigationIssue, len(xmlInfo.Issues))
	for i := range xmlInfo.Issues {
		issues[i] = *convertXMLIssue(&xmlInfo.Issues[i])
	}

	errors := make([]MitigationError, len(xmlInfo.Errors))
	for i, err := range xmlInfo.Errors {
		errors[i] = MitigationError{
			Type:       err.Type,
			FlawIDList: err.FlawIDList,
		}
	}

	return &MitigationInfo{
		Version: xmlInfo.Version,
		BuildID: xmlInfo.BuildID,
		Issues:  issues,
		Errors:  errors,
	}
}

// convertXMLIssue converts xml.Issue to api.MitigationIssue wrapper type
func convertXMLIssue(xmlIssue *xml.Issue) *MitigationIssue {
	if xmlIssue == nil {
		return nil
	}

	actions := make([]MitigationAction, len(xmlIssue.MitigationActions))
	for i, action := range xmlIssue.MitigationActions {
		actions[i] = MitigationAction{
			Action:   action.Action,
			Desc:     action.Desc,
			Reviewer: action.Reviewer,
			Date:     action.Date,
			Comment:  action.Comment,
		}
	}

	return &MitigationIssue{
		FlawID:            xmlIssue.FlawID,
		Category:          xmlIssue.Category,
		MitigationActions: actions,
	}
}
