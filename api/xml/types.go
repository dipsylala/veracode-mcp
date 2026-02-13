package xml

import "encoding/xml"

// MitigationInfo represents the root element of the mitigation info XML response
type MitigationInfo struct {
	XMLName xml.Name `xml:"https://analysiscenter.veracode.com/schema/mitigationinfo/1.0 mitigationinfo"`
	Version string   `xml:"mitigationinfo_version,attr,omitempty"`
	BuildID int64    `xml:"build_id,attr"`
	Issues  []Issue  `xml:"issue"`
	Errors  []Error  `xml:"error"`
}

// Issue represents a flaw with its mitigation actions
type Issue struct {
	FlawID            int64              `xml:"flaw_id,attr"`
	Category          string             `xml:"category,attr"`
	MitigationActions []MitigationAction `xml:"mitigation_action"`
}

// MitigationAction represents a single mitigation action for a flaw
type MitigationAction struct {
	Action   string `xml:"action,attr"`
	Desc     string `xml:"desc,attr,omitempty"`
	Reviewer string `xml:"reviewer,attr,omitempty"`
	Date     string `xml:"date,attr,omitempty"`
	Comment  string `xml:"comment,attr,omitempty"`
}

// Error represents an error for flaws that could not be processed
type Error struct {
	Type       string `xml:"type,attr,omitempty"`
	FlawIDList string `xml:"flaw_id_list,attr,omitempty"`
}

// ActionType represents the allowed mitigation action types
type ActionType string

const (
	ActionComment       ActionType = "comment"
	ActionFP            ActionType = "fp" // False Positive
	ActionLibrary       ActionType = "library"
	ActionAcceptRisk    ActionType = "acceptrisk"
	ActionAppDesign     ActionType = "appdesign"
	ActionOSEnv         ActionType = "osenv"
	ActionNetEnv        ActionType = "netenv"
	ActionRejected      ActionType = "rejected"
	ActionAccepted      ActionType = "accepted"
	ActionRemediated    ActionType = "remediated"
	ActionNoActionTaken ActionType = "noactiontaken"
	ActionConforms      ActionType = "conforms"
	ActionDeviates      ActionType = "deviates"
	ActionDefer         ActionType = "defer"
)
